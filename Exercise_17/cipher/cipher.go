package cipher

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"errors"
	"fmt"
	"io"
)

//EncryptWriter takes key, io.writer and writes the encoded value on *cipher.StreamWriter
//returns error if any
func EncryptWriter(key string, w io.Writer) (*cipher.StreamWriter, error) {
	iv := make([]byte, aes.BlockSize)
	io.ReadFull(rand.Reader, iv)
	stream, _ := encrytStream(key, iv)
	n, err := w.Write(iv)
	err = checkInitVector(n, iv, err)
	return &cipher.StreamWriter{S: stream, W: w}, err
}

//checking whether all data from init vaector has been read or not
//return : error
func checkInitVector(n int, iv []byte, err error) error {
	if len(iv) != n || err != nil {
		return errors.New("Unable to write IV into writer")
	}
	return nil
}

//Func DecryptReader decrypts the ecrypted data
//input : key string - based on which encryption was done
//return reader which holds the decrypted data and error
func DecryptReader(key string, r io.Reader) (*cipher.StreamReader, error) {
	iv := make([]byte, aes.BlockSize)
	n, err := r.Read(iv)
	if n < len(iv) || err != nil {
		return nil, errors.New("Encrypt : unable to read full IV")
	}
	stream, err := decryptStream(key, iv)
	return &cipher.StreamReader{S: stream, R: r}, nil
}

//encryptStream : Encrypts data
//input : key string- based on which encryption has to be done
//return : stream of encrypted data
func encrytStream(key string, iv []byte) (cipher.Stream, error) {
	block, err := newCipherBlock(key)
	return cipher.NewCFBEncrypter(block, iv), err
}

//func decryptStream decrypts the data
//input : key string - based on which ecryption was done
//return :  stream of decrypted data, error
func decryptStream(key string, iv []byte) (cipher.Stream, error) {
	cipherBlock, err := newCipherBlock(key)
	return cipher.NewCFBDecrypter(cipherBlock, iv), err

}

//function newCipherBlock returns the fixed size cipher.Block
func newCipherBlock(key string) (cipher.Block, error) {
	hash := md5.New()
	fmt.Fprint(hash, key)
	cipherKey := hash.Sum(nil)
	return aes.NewCipher(cipherKey)
}
