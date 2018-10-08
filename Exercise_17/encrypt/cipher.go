package cipher

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"fmt"
)

func encrypt(key, plainText string) (string, error) {
	// cipherBlock, err := newCipherBlock(key)
	// if err != nil {
	// 	fmt.Println("Error : ", err)
	// }
	// cipherText := make([]byte, aes.BlockSize+len(plainText))
	// iv := cipherText[:aes.BlockSize]
	// stream := cipher.NewCBCEncrypter(cipherBlock, cipherText)
	//stream.BlockSize

	return "", nil
}

func decrypt() string {

	return ""
}

func newCipherBlock(key string) (cipher.Block, error) {
	hasher := md5.New()
	fmt.Fprint(hasher, key)
	cipherKey := hasher.Sum(nil)
	return aes.NewCipher([]byte(cipherKey))
}
