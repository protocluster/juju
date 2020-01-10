package k8s

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	admission "k8s.io/api/admission/v1beta1"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	gc "gopkg.in/check.v1"
)

type HandlerSuite struct {
}

var _ = gc.Suite(&HandlerSuite{})

func TestHandlerSuite(t *testing.T) { gc.TestingT(t) }

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

func (k *HandlerSuite) TestEmptyBodyFails(c *gc.C) {
	req := httptest.NewRequest(http.MethodPost, "/", nil)
	recorder := httptest.NewRecorder()

	admissionHandler().ServeHTTP(recorder, req)

	c.Assert(recorder.Code, gc.Equals, http.StatusBadRequest)
}

func (k *HandlerSuite) TestUnknownContentType(c *gc.C) {
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader("junk"))
	req.Header.Set("junk", "junk")
	recorder := httptest.NewRecorder()

	admissionHandler().ServeHTTP(recorder, req)

	c.Assert(recorder.Code, gc.Equals, http.StatusUnsupportedMediaType)
}

func (k *HandlerSuite) TestUnsupportedGroupKindVersion(c *gc.C) {
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

	admissionHandler().ServeHTTP(recorder, req)
	c.Assert(recorder.Code, gc.Equals, http.StatusOK)
}
