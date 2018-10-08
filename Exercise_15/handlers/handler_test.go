package handlers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPanicDemo(t *testing.T) {
	req, err := http.NewRequest("GET", "localhost:3000", nil)
	if err != nil {
		t.Error("Cound not create new Request ", err)
	}
	rec := httptest.NewRecorder()
	defer func() {
		if err := recover(); err != nil {
		}
	}()
	PanicDemo(rec, req)
	res := rec.Result()
	if res.StatusCode != http.StatusOK {
		t.Errorf("Error occured, expected status OK but recieved %s ", string(res.StatusCode))
	}
}

func TestSourceCodeNavigator(t *testing.T) {
	testSuits := []struct {
		testName string
		path     string
		status   int
	}{
		{
			testName: "Test1",
			path:     "line=aws&path=/usr/local/go/src/runtime/debug/stack.go",
			status:   500,
		},
		{
			testName: "Test2",
			path:     "line=aw17&path=/home/gslab/go/src/github.com/Gophercises/Exercise_15/middleware/MwRecovery.go",
			status:   500,
		},
		{
			testName: "Test3",
			path:     "line=65&path=%2Fhome%2Fgslab%2Fgo%2Fsrc%2Fgithub.com%2FGophercises%2FExercise_15%2Fhandlers%2Fhandler.go",
			status:   200,
		},
	}

	for index := 0; index < len(testSuits); index++ {
		req, err := http.NewRequest("GET", "localhost:3000/debug/?"+testSuits[index].path, nil)
		if err != nil {
			t.Error("Could not create new Request ", err)
		}
		rec := httptest.NewRecorder()
		SourceCodeNavigator(rec, req)
		res := rec.Result()
		if res.StatusCode != testSuits[index].status {
			t.Errorf("Expected Status OK but recieved :%v ", res.StatusCode)
		}
	}
}

func TestWelcome(t *testing.T) {
	req, err := http.NewRequest("GET", "localhost:3000", nil)
	if err != nil {
		t.Errorf("Error occured while creating new request %v", err)
	}
	rec := httptest.NewRecorder()
	Welcome(rec, req)
	res := rec.Result()
	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected status ok but got %v ", res.StatusCode)
	}
}

func TestHandler(t *testing.T) {
	server := httptest.NewServer(Handler())
	defer server.Close()
	testSuits := []struct {
		testName string
		url      string
		status   int
	}{
		{
			testName: "Test1",
			url:      "/debug/?line=65&path=%2Fhome%2Fgslab%2Fgo%2Fsrc%2Fgithub.com%2FGophercises%2FExercise_15%2Fhandlers%2Fhandler.go",
			status:   500,
		},
		{
			testName: "Test2",
			url:      "/debug/?line=aws&path=%2Fhome%2Fgslab%2Fgo%2Fsrc%2Fgithub.com%2FGophercises%2FExercise_15%2Fhandlers%2Fhandler.go",
			status:   500,
		},
		{
			testName: "Test3",
			url:      "/",
			status:   200,
		},
		{
			testName: "Test4",
			url:      "/debug/",
			status:   500,
		},
		{
			testName: "Test2",
			url:      "/debug/?line=aws&path=%2%2FGophercises%2FExercise_15%2Fhandlers%2Fhandler.go",
			status:   500,
		},
	}
	for index := 0; index < len(testSuits); index++ {
		res, err := http.Get(fmt.Sprintf(server.URL + testSuits[index].url))
		if err != nil {
			t.Error("Test case falied here due to GET method..", err)
		}
		defer res.Body.Close()
		if res.StatusCode != testSuits[index].status {
			t.Errorf("Unexpected Status code Actual :%v  and expected : %v", res.StatusCode, http.StatusOK)
		}
	}
}
