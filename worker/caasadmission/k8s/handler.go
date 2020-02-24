package k8s

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/juju/juju/caas/kubernetes/provider"

	"github.com/juju/errors"
	admission "k8s.io/api/admission/v1beta1"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/apimachinery/pkg/types"
)

const (
	ExpectedContentType = "application/json"
	HeaderContentType   = "Content-Type"
)

var (
	AdmissionGVK = schema.GroupVersionKind{
		Group:   admission.SchemeGroupVersion.Group,
		Version: admission.SchemeGroupVersion.Version,
		Kind:    "AdmissionReview",
	}
)

func admissionHandler(logger Logger, rbacMapper provider.RBACMapper) http.Handler {
	codecFactory := serializer.NewCodecFactory(runtime.NewScheme())

	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		data, err := ioutil.ReadAll(req.Body)
		if err != nil {
			logger.Errorf("digesting admission request body: %v", err)
			http.Error(res, fmt.Sprintf("%s: reading request body",
				http.StatusText(http.StatusInternalServerError)), http.StatusInternalServerError)
			return
		}

		if len(data) == 0 {
			http.Error(res, fmt.Sprintf("%s: empty request body",
				http.StatusText(http.StatusBadRequest)), http.StatusBadRequest)
			return
		}

		if req.Header.Get(HeaderContentType) != ExpectedContentType {
			http.Error(res, fmt.Sprintf("%s: supported content types = [%s]",
				http.StatusText(http.StatusUnsupportedMediaType),
				ExpectedContentType), http.StatusUnsupportedMediaType)
			return
		}

		finalise := func(review *admission.AdmissionReview, response *admission.AdmissionResponse) {
			var uid types.UID
			if review != nil {
				uid = review.Request.UID
			}
			response.UID = uid

			body, err := json.Marshal(admission.AdmissionReview{
				Response: response,
			})
			if err != nil {
				logger.Errorf("marshaling admission request response body: %v", err)
				http.Error(res, fmt.Sprintf("%s: building response body",
					http.StatusText(http.StatusInternalServerError)), http.StatusInternalServerError)
			}
			if _, err := res.Write(body); err != nil {
				logger.Errorf("writing admission request response body: %v", err)
			}
		}

		var admissionReview *admission.AdmissionReview
		obj, gvk, err := codecFactory.UniversalDecoder().Decode(data, nil, nil)
		if err != nil {
			finalise(admissionReview, errToAdmissionResponse(err))
			return
		}

		if !compareGroupVersionKind(&AdmissionGVK, gvk) {
			finalise(admissionReview,
				errToAdmissionResponse(errors.NewNotValid(nil, "unsupported group kind version")))
			return
		} else {
			var ok bool
			if admissionReview, ok = obj.(*admission.AdmissionReview); !ok {
				finalise(admissionReview,
					errToAdmissionResponse(errors.NewNotValid(nil, "converting admission request")))
				return
			}
		}

		_, err = rbacMapper.AppNameForServiceAccount(
			types.UID(admissionReview.Request.UserInfo.UID))
		if err != nil && !errors.IsNotValid(err) {
			finalise(admissionReview, errToAdmissionResponse(err))
			return
		}

		reviewResponse := &admission.AdmissionResponse{}
		finalise(admissionReview, reviewResponse)
	})
}

func compareGroupVersionKind(a, b *schema.GroupVersionKind) bool {
	if a == nil || b == nil {
		return false
	}
	return a.Group == b.Group && a.Version == b.Version && a.Kind == b.Kind
}

func errToAdmissionResponse(err error) *admission.AdmissionResponse {
	return &admission.AdmissionResponse{
		Result: &meta.Status{
			Message: err.Error(),
		},
	}
}
