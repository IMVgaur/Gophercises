package middleware

import (
	"net/http"
	"net/http/httptest"
	"runtime/debug"
	"strings"
	"testing"

	handlers "github.com/Gophercises/Exercise_15/handlers"
	"github.ibm.com/CloudBroker/dash_utils/dashtest"
)

func TestMain(m *testing.M) {
	dashtest.ControlCoverage(m)
}

func TestRecoveryMID(t *testing.T) {
	handler := http.HandlerFunc(handlers.PanicDemo)
	executeRequest("GET", "/panic", RecoveryMid(handler))
}

func executeRequest(method string, url string, handler http.Handler) (*httptest.ResponseRecorder, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	rec.Result()
	return rec, err
}

func TestErrorLinks(t *testing.T) {
	stack := debug.Stack()
	output := ErrLinks(string(stack))
	if !strings.Contains(output, "<a href=") {
		t.Error("Response is not expected ...", output)
	}
}
