package k8s

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

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

func admissionHandler(logger Logger) http.Handler {
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

		var reviewResponse *admission.AdmissionResponse
		var admissionReview *admission.AdmissionReview
		obj, gvk, err := codecFactory.UniversalDecoder().Decode(data, nil, nil)
		if err != nil {
			reviewResponse = errToAdmissionResponse(err)
		}

		if !compareGroupVersionKind(&AdmissionGVK, gvk) {
			reviewResponse = errToAdmissionResponse(errors.NewNotValid(nil, "unsupported group kind version"))
		} else {
			var ok bool
			if admissionReview, ok = obj.(*admission.AdmissionReview); !ok {
				reviewResponse = errToAdmissionResponse(errors.NewNotValid(nil, "converting admission request"))
			}
		}

		var uid types.UID
		if admissionReview != nil {
			uid = admissionReview.Request.UID
		}

		if reviewResponse == nil {
			reviewResponse = &admission.AdmissionResponse{}
		}

		reviewResponse.UID = uid
		response := admission.AdmissionReview{
			Response: reviewResponse,
		}

		body, err := json.Marshal(response)
		logger.Infof("got here")
		if err != nil {
			logger.Errorf("marshaling admission request response body: %v", err)
			http.Error(res, fmt.Sprintf("%s: building response body",
				http.StatusText(http.StatusInternalServerError)), http.StatusInternalServerError)
		}
		if _, err := res.Write(body); err != nil {
			logger.Errorf("writing admission request response body: %v", err)
		}
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
