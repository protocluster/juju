package k8s

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	k8stest "github.com/juju/juju/caas/kubernetes/provider/test"

	"github.com/juju/loggo"
	gc "gopkg.in/check.v1"
	admission "k8s.io/api/admission/v1beta1"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

type HandlerSuite struct {
	logger Logger
}

var _ = gc.Suite(&HandlerSuite{})

func TestHandlerSuite(t *testing.T) { gc.TestingT(t) }

func (h *HandlerSuite) SetupTest(c *gc.C) {
	h.logger = loggo.Logger{}
}

func (h *HandlerSuite) TestCompareGroupVersionKind(c *gc.C) {
	tests := []struct {
		A           *schema.GroupVersionKind
		B           *schema.GroupVersionKind
		ShouldMatch bool
	}{
		{
			A: &schema.GroupVersionKind{
				Group:   admission.SchemeGroupVersion.Group,
				Version: admission.SchemeGroupVersion.Version,
				Kind:    "AdmissionReview",
			},
			B: &schema.GroupVersionKind{
				Group:   admission.SchemeGroupVersion.Group,
				Version: admission.SchemeGroupVersion.Version,
				Kind:    "AdmissionReview",
			},
			ShouldMatch: true,
		},
		{
			A: &schema.GroupVersionKind{
				Group:   admission.SchemeGroupVersion.Group,
				Version: admission.SchemeGroupVersion.Version,
				Kind:    "AdmissionReview",
			},
			B: &schema.GroupVersionKind{
				Group:   admission.SchemeGroupVersion.Group,
				Version: admission.SchemeGroupVersion.Version,
				Kind:    "Junk",
			},
			ShouldMatch: false,
		},
		{
			A: &schema.GroupVersionKind{
				Group:   admission.SchemeGroupVersion.Group,
				Version: admission.SchemeGroupVersion.Version,
				Kind:    "AdmissionReview",
			},
			B:           nil,
			ShouldMatch: false,
		},
	}

	for _, test := range tests {
		c.Assert(compareGroupVersionKind(test.A, test.B), gc.Equals, test.ShouldMatch)
	}
}

func (h *HandlerSuite) TestEmptyBodyFails(c *gc.C) {
	req := httptest.NewRequest(http.MethodPost, "/", nil)
	recorder := httptest.NewRecorder()

	admissionHandler(h.logger, &k8stest.RBACMapper{}).ServeHTTP(recorder, req)

	c.Assert(recorder.Code, gc.Equals, http.StatusBadRequest)
}

func (h *HandlerSuite) TestUnknownContentType(c *gc.C) {
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader("junk"))
	req.Header.Set("junk", "junk")
	recorder := httptest.NewRecorder()

	admissionHandler(h.logger, &k8stest.RBACMapper{}).ServeHTTP(recorder, req)

	c.Assert(recorder.Code, gc.Equals, http.StatusUnsupportedMediaType)
}

func (h *HandlerSuite) TestUnsupportedGroupKindVersion(c *gc.C) {
	admissionRequest := &admission.AdmissionRequest{
		Kind: meta.GroupVersionKind{
			Group:   AdmissionGVK.Group,
			Version: AdmissionGVK.Version,
			Kind:    "junk",
		},
	}

	body, err := json.Marshal(admissionRequest)
	c.Assert(err, gc.IsNil)

	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(body))
	req.Header.Set(HeaderContentType, ExpectedContentType)
	recorder := httptest.NewRecorder()

	admissionHandler(h.logger, &k8stest.RBACMapper{}).ServeHTTP(recorder, req)
	c.Assert(recorder.Code, gc.Equals, http.StatusOK)
}
