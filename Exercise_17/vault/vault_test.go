package vault

import (
	"os"
	"path/filepath"
	"testing"

	homedir "github.com/mitchellh/go-homedir"
	"github.ibm.com/CloudBroker/dash_utils/dashtest"
)

func TestMain(m *testing.M) {
	dashtest.ControlCoverage(m)
}
func TestSet(t *testing.T) {
	home, _ := homedir.Dir()
	fp := filepath.Join(home, "secretTest.txt")
	v := File("abc", fp)
	err := v.Set("xyz", "testing")
	if err != nil {
		t.Errorf("Expected no err but got err %v", err)
	}
}
func TestSetNegative(t *testing.T) {
	home, _ := homedir.Dir()
	fp := filepath.Join(home, "secretTest.txt")
	v := File("", fp)
	err := v.Set("xyz", "testing")
	if err == nil {
		t.Error("Expected  Error but got suucessful execution")
	}
}

func TestGet(t *testing.T) {
	home, _ := homedir.Dir()
	fp := filepath.Join(home, "secretTest.txt")
	v := File("abc", fp)
	_, err := v.Get("xyz")
	if err != nil {
		t.Errorf("Expected no err but got err %v", err)
	}
}
func TestGetNegative(t *testing.T) {
	home, _ := homedir.Dir()
	fp := filepath.Join(home, "secretTest.txt")
	v := File("abc", fp)
	_, err := v.Get("x")
	if err == nil {
		t.Error("Expected Error but got NO error ")
	}
	v = File("", fp)
	_, err = v.Get("x")
	if err == nil {
		t.Error("Expected Error but got NO error ")
	}
}

func TestLoad(t *testing.T) {
	home, _ := homedir.Dir()
	fp := filepath.Join(home, "testload.txt")
	v := File("abc", fp)
	err := v.load()
	if err != nil {
		t.Errorf("Expected no err but got err %v", err)
	}
	os.Remove(fp)
}

func TestLoadNegative(t *testing.T) {
	home, _ := homedir.Dir()
	fp := filepath.Join(home, "secretTest.txt")
	v := File("", fp)
	err := v.load()
	if err == nil {
		t.Error("Expected  error but got NO error ")
	}
	f := filepath.Join(home, "loadtest.txt")
	v = File("abc", f)
	err = v.load()
	if err == nil {
		t.Error("Expected Error but got NO error ")
	}
}

func TestSave(t *testing.T) {
	var v Vault
	err := v.save()
	if err == nil {
		t.Error("Expected Error but got NO error ")
	}
}
