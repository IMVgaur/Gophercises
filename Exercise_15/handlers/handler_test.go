package handlers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.ibm.com/CloudBroker/dash_utils/dashtest"
)

func TestMain(m *testing.M) {
	dashtest.ControlCoverage(m)
}
func TestWelcome(t *testing.T) {
	req, err := http.NewRequest("GET", "localhost:8888", nil)
	if err != nil {
		t.Fatalf("could not create request: %v", err)
	}
	rec := httptest.NewRecorder()
	welcome(rec, req)
	res := rec.Result()
	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected Status OK, recieved status %s", string(res.StatusCode))
	}
}

func TestPanicDemo(t *testing.T) {
	req, err := http.NewRequest("GET", "localhost:8888/panic", nil)
	if err != nil {
		t.Fatalf("could not create request: %v", err)
	}
	rec := httptest.NewRecorder()
	defer func() {
		if err := recover(); err != nil {
		}
	}()
	PanicDemo(rec, req)
	res := rec.Result()
	if res.StatusCode != http.StatusInternalServerError {
		t.Errorf("Expected Status StatusInternalServerError, recieved status %s", string(res.StatusCode))
	}
}

func TestSourceCodeNavigator(t *testing.T) {

	testSuit := []struct {
		testName string
		url      string
		status   int
	}{
		{
			testName: "test1",
			url:      "line=24&path=/usr/local/go/src/runtime/debug/stack.go",
			status:   200,
		}, {
			testName: "test2",
			url:      "line=ewr&path=/usr/local/go/src/runtime/debug/stack.go",
			status:   200,
		},
		{
			testName: "test3",
			url:      "line=24&path=/usr/local/go/src/debug/stack.go",
			status:   500,
		},
	}
	for i := 0; i < len(testSuit); i++ {
		req, err := http.NewRequest("GET", "http://localhost:8888/debug/?"+testSuit[i].url, nil)
		if err != nil {
			t.Fatalf("could not create request: %v", err)
		}
		rec := httptest.NewRecorder()
		sourceCodeNavigator(rec, req)
		res := rec.Result()
		if res.StatusCode != testSuit[i].status {
			t.Errorf("Test case Number: %v Expected %v , Actual status %v", testSuit[i].testName, testSuit[i].status, res.StatusCode)
		}
	}

}

func TestHandler(t *testing.T) {
	srv := httptest.NewServer(Handler())
	defer srv.Close()
	testSuit := []struct {
		testName string
		url      string
		status   int
	}{
		{
			testName: "test1",
			url:      "/",
			status:   200,
		},
		{
			testName: "test2",
			url:      "/debug",
			status:   500,
		},
	}
	for i := 0; i < len(testSuit); i++ {
		res, err := http.Get(fmt.Sprintf(srv.URL + "/" + testSuit[i].url))
		if err != nil {
			t.Fatalf("could not send GET request: %v", err)
		}
		defer res.Body.Close()
		if res.StatusCode != testSuit[i].status {
			t.Errorf("test case name : %v expected %v; got %v", testSuit[i].testName, testSuit[i].status, res.Status)
		}
	}
}
