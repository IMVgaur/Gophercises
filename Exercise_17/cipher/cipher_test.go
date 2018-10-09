package cipher

import (
	"testing"

	"github.ibm.com/CloudBroker/dash_utils/dashtest"
)

func TestMain(m *testing.M) {
	dashtest.ControlCoverage(m)
}

func TestEncrypt(t *testing.T) {
	plainText := "IMVGaur"
	key := "aBC"
	_, err := Encrypt(key, plainText)
	if err != nil {
		t.Errorf("Error : %v", err)
	}
}

func TestDecrypt(t *testing.T) {
	cipherHex := "6749571c3e4c16917a25f4502a444563564044f3389917"
	key := "aBC"
	_, err := Decrypt(key, cipherHex)
	if err != nil {
		t.Errorf("Error  : %v", err)
	}
}

func TestNewCipherBlock(t *testing.T) {
	key := "aBC"
	_, err := newCipherBlock(key)
	if err != nil {
		t.Errorf("Error : %v", err)
	}
}
