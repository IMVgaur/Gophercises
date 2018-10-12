package handlers

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
	"time"

	homedir "github.com/mitchellh/go-homedir"
)

func TestWelcome(t *testing.T) {
	srv := httptest.NewServer(Handler())
	defer srv.Close()
	client := &http.Client{
		Timeout: 50 * time.Second,
	}
	r, _ := http.NewRequest("GET", srv.URL, nil)
	res, _ := client.Do(r)
	if res.StatusCode != http.StatusOK {
		t.Error("Expected status ok but got different status")
	}
}

func TestUpload(t *testing.T) {
	srv := httptest.NewServer(Handler())
	defer srv.Close()
	h, _ := homedir.Dir()
	imgPath := filepath.Join(h, "img/test.jpg")
	file, err := os.Open(imgPath)
	if err != nil {
		t.Error("error in opening file")
	}
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("image", file.Name())
	if err != nil {
		t.Error("error in copy")
	}
	_, err = io.Copy(part, file)
	if err != nil {
		t.Error("error in copy")
	}
	err = writer.Close()
	if err != nil {
		t.Error("error in close writer")
	}
	req, _ := http.NewRequest("POST", srv.URL+"/upload", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, _ := http.DefaultClient.Do(req)
	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected status ok but got different status %v", res.Status)
	}
}

func TestTempFile(t *testing.T) {
	_, err := tempFile("test", "txt")
	if err != nil {
		t.Errorf("Expected no error but got error : %v", err)
	}
}

func TestTempFileNegetive(t *testing.T) {
	_, err := tempFile("/invalid/Name/file/", "txt")
	if err == nil {
		t.Error("Expected error but got no error", err)
	}
}

func TestTempFileNilFile(t *testing.T) {
	_, err := tempFile("", "/invalid/Name/file/")
	if err == nil {
		t.Error("Expected error but got no error", err)
	}
}

func TestModify(t *testing.T) {
	srv := httptest.NewServer(Handler())
	defer srv.Close()
	req, _ := http.NewRequest("GET", srv.URL+"/modify/test.jpg?mode=3", nil)
	res, _ := http.DefaultClient.Do(req)
	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected status ok but got different status %v", res.Status)
	}
}

func TestModifyMode(t *testing.T) {
	srv := httptest.NewServer(Handler())
	defer srv.Close()
	req, _ := http.NewRequest("GET", srv.URL+"/modify/test.jpg?mode=3", nil)
	res, _ := http.DefaultClient.Do(req)
	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected status ok but got different status %v", res.Status)
	}
}

func TestModifyInvalidMode(t *testing.T) {
	srv := httptest.NewServer(Handler())
	defer srv.Close()
	req, _ := http.NewRequest("GET", srv.URL+"/modify/test.jpg?mode=a", nil)
	res, _ := http.DefaultClient.Do(req)
	if res.StatusCode == http.StatusOK {
		t.Errorf("Expected status Internal server error but got different status %v", res.Status)
	}
}

func TestModifyModeShapes(t *testing.T) {
	srv := httptest.NewServer(Handler())
	defer srv.Close()
	req, _ := http.NewRequest("GET", srv.URL+"/modify/test.jpg?mode=3&n=5", nil)
	res, _ := http.DefaultClient.Do(req)
	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected status ok but got different status %v", res.Status)
	}
}
func TestModifyModeNegativeExt(t *testing.T) {
	srv := httptest.NewServer(Handler())
	defer srv.Close()
	req, _ := http.NewRequest("GET", srv.URL+"/modify/secret.txt?mode=2", nil)
	res, _ := http.DefaultClient.Do(req)
	if res.StatusCode == http.StatusOK {
		t.Errorf("Expected status internal server error but got different status %v", res.Status)
	}
}
