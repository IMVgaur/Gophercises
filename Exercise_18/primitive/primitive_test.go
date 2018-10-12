package primitive

import (
	"os"
	"path/filepath"
	"testing"

	"github.ibm.com/CloudBroker/dash_utils/dashtest"

	"github.com/mitchellh/go-homedir"
)

func TestMain(m *testing.M) {
	dashtest.ControlCoverage(m)
}

func TestTempFile(t *testing.T) {
	_, err := tempFile("test", "txt")
	if err != nil {
		t.Errorf("Expected no error but got error : %v", err)
	}
}

func TestTempFileNegetive(t *testing.T) {
	_, err := tempFile("///////////invalid///", "txt")
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

func TestWithMode(t *testing.T) {
	ret := WithMode(ModeCombo)
	if ret == nil {
		t.Error("expected []string of modes but got nil")
	}
}

func TestWithModeNegetive(t *testing.T) {
	ret := WithMode(5555555)
	if ret == nil {
		t.Error("expected []string of modes but got nil")
	}
}

func TestPrimitive(t *testing.T) {
	args := WithMode(ModeCircle)
	h, _ := homedir.Dir()
	imgPath := filepath.Join(h, "img/test.jpg")
	outPath := filepath.Join(h, "img/test1.jpg")
	_, err := primitive(imgPath, outPath, 1, args...)
	if err != nil {
		t.Errorf("Expected no error but got error:: %v", err)
	}
}

func TestPrimitiveNegetive(t *testing.T) {
	args := WithMode(ModeCircle)
	home, _ := homedir.Dir()
	inFile := filepath.Join(home, "img/invalidFile")
	outFile := filepath.Join(home, "img/out/out.png")
	_, err := primitive(inFile, outFile, 2, args...)
	if err == nil {
		t.Error("Expected error but got no error")
	}
}

func TestTransform(t *testing.T) {
	home, _ := homedir.Dir()
	fp := filepath.Join(home, "img/test.jpg")
	ext := filepath.Ext(fp)
	file, _ := os.Open(fp)
	defer file.Close()
	_, err := Transform(file, ext, 10, nil)
	if err != nil {
		t.Errorf("Ecpected no error but got error : %v", err)
	}
}

func TestTransformNExtension(t *testing.T) {
	home, _ := homedir.Dir()
	fp := filepath.Join(home, "img/test.jpg")
	_ = filepath.Ext(fp)
	file, _ := os.Open(fp)
	defer file.Close()
	_, err := Transform(file, "/////////", 10, nil)
	if err == nil {
		t.Errorf("Ecpected no error but got error : %v", err)
	}
}

func TestTransformNegetivePrimitive(t *testing.T) {
	home, _ := homedir.Dir()
	fp := filepath.Join(home, "img/out.png")
	ext := filepath.Ext(fp)
	file, _ := os.Open(fp)
	_, err := Transform(file, ext, -2, nil)
	if err == nil {
		t.Errorf("Expected error but got no error%v", err)
	}
}

func TestDataBuffer(t *testing.T) {
	home, _ := homedir.Dir()
	fp := filepath.Join(home, "img/test.jpg")
	file, _ := os.Open(fp)
	_, err := dataBuffer(file)
	if err != nil {
		t.Errorf("Expected no error but got : %v", err)
	}
}

func TestDataBufferNegetive(t *testing.T) {
	home, _ := homedir.Dir()
	fp := filepath.Join(home, "img/abcout.png")
	file, _ := os.Open(fp)
	_, err := dataBuffer(file)
	if err == nil {
		t.Errorf("Expected no error but got : %v", err)
	}
}
