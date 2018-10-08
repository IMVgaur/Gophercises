package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMain(m *testing.M) {
	main()
}

func TestWelcome(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		log.Fatal("Error occured while testing : ", err)
	}
	rec := httptest.NewRecorder()
	welcome(rec, req)
	res := rec.Result()
	if res.StatusCode == http.StatusOK {
		t.Error("Unexpected Results....")
	}
	t.Error("Here is the error...")
}

func TestPanicDemo(t *testing.T) {
	req, err := http.NewRequest("GET", "/panic/", nil)
	if err != nil {
		t.Fatal("Error occured while testing...")
	}
	rec := httptest.NewRecorder()
	panicDemo(rec, req)
	res := rec.Result()
	if status := res.StatusCode; status != http.StatusOK {
		t.Errorf("Error occured while testing : ")
		fmt.Println("Error occured...")
	}
	t.Error("Error occured while testing : ", rec.Result().Status)
}

func TestAfterPanic(t *testing.T) {
	req, err := http.NewRequest("GET", "localhost:3000/afterPanic/", nil)
	if err != nil {
		t.Error("Error occured while testing...")
	}
	rec := httptest.NewRecorder()
	afterPanicDemo(rec, req)
	res := rec.Result()
	if res.StatusCode != http.StatusOK {
		t.Error("Mismatch in the expected status code and actual status code...")
	}
}
