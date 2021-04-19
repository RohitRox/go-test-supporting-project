package controllers_test

import (
	"go-test-supporting-project/controllers"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestStatusHandler(t *testing.T) {
	t.Run("status check", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/anyroute", nil)
		if err != nil {
			t.Fatal(err)
		}

		handlers := controllers.Handler{}

		rec := httptest.NewRecorder()
		handlers.Status(rec, req)

		if status := rec.Code; status != http.StatusOK {
			t.Errorf("expected sattus code: %v got: %v",
				http.StatusOK, status)
		}
	})
}
