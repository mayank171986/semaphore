package util

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestGetIntParam(t *testing.T) {
	req, _ := http.NewRequest("GET", "/test/123", nil)
	rr := httptest.NewRecorder()

	r := mux.NewRouter()
	r.NewRoute().Path("/test/{test_id}").Subrouter().Use(mockParam)
	r.ServeHTTP(rr, req)

	if rr.Code != 200 {
		t.Errorf("Response code should be 200 %d", rr.Code)
	}
}

func mockParam(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := GetIntParam("test_id", w, r)
		if err != nil {
			return
		}

		w.WriteHeader(200)

		if (next != nil) {
      next.ServeHTTP(w, r)
    }
	})
}
