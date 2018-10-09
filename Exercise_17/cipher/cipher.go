package cipher

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
)

//Ecrypt the input hex value
//input : key string, plaintext string
//returns : Ecrypted stirng, error
func Encrypt(key, plainText string) (string, error) {
	cipherBlock, err := newCipherBlock(key)
	if err != nil {
		return "", err
	}
	cipherText := make([]byte, aes.BlockSize+len(plainText))
	iv := cipherText[:aes.BlockSize]
	_, err = io.ReadFull(rand.Reader, iv)
	if err != nil {
		return "", err
	}
	stream := cipher.NewCFBEncrypter(cipherBlock, iv)
	stream.XORKeyStream(cipherText[aes.BlockSize:], []byte(plainText))
	return fmt.Sprintf("%x", cipherText), nil
}

//Decrypt the input hex value
//input : key string, cipherHex string
//returns : decrypted stirng, error
func Decrypt(key, cipherHex string) (string, error) {
	cipherBlock, err := newCipherBlock(key)
	if err != nil {
		fmt.Printf("Error : %v", err)
	}
	cipherText, err := hex.DecodeString(cipherHex)
	if len(cipherText) < aes.BlockSize {
		return "", errors.New("Cipher text is too short")
	}
	iv := cipherText[:aes.BlockSize]
	stream := cipher.NewCFBDecrypter(cipherBlock, iv)
	stream.XORKeyStream(cipherText, cipherText)
	return string(cipherText), nil
}

//helper function
//input : key string
//returns cipher.Block, error
func newCipherBlock(key string) (cipher.Block, error) {
	hasher := md5.New()
	fmt.Fprint(hasher, key)
	cipherKey := hasher.Sum(nil)
	return aes.NewCipher([]byte(cipherKey))
}
