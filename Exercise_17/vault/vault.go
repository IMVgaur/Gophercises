package vault

import (
	"encoding/json"
	"errors"
	"io"
	"os"
	"sync"

	"github.com/Gophercises/Exercise_17/cipher"
)

//Represents bunch of required data structure
type Vault struct {
	encodingKey string
	filePath    string
	mutex       sync.Mutex
	keyValues   map[string]string
}

//File function Initialize task struct
//input : Encoding key for encryption, filepath for file io
//return : Initialized Vault struct
func File(encodingKey, filePath string) *Vault {
	return &Vault{
		encodingKey: encodingKey,
		filePath:    filePath,
	}
}

//Set function sets encrypted data into the file
//input : key string - used for encrypting password
// return : error
func (v *Vault) Set(key, value string) error {
	v.mutex.Lock()
	defer v.mutex.Unlock()
	err := v.load()
	if err != nil {
		return err
	}
	v.keyValues[key] = value
	return v.save()
}

//Get method gets the key and value in plain text format from secret file
func (v *Vault) Get(key string) (string, error) {
	v.mutex.Lock()
	defer v.mutex.Unlock()
	err := v.load()
	if err != nil {
		return "", err
	}
	value, ok := v.keyValues[key]
	if !ok {
		return "", errors.New("Get : Element not found for this Key : " + key)
	}
	return value, nil
}

//load func loads the data from file
//return stream of decoded key value pairs
func (v *Vault) load() error {
	file, err := os.Open(v.filePath)
	if err != nil {
		v.keyValues = make(map[string]string)
		return err
	}
	defer file.Close()
	sr, err := cipher.DecryptReader(v.encodingKey, file)
	if err != nil {
		return err
	}
	return v.readKeyValues(sr)
}

//func save represents logic of writting encrypted data on file
//return error if any
func (v *Vault) save() error {
	f, err := os.OpenFile(v.filePath, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return err
	}
	defer f.Close()
	w, err := cipher.EncryptWriter(v.encodingKey, f)
	if err != nil {
		return err
	}
	return v.writeKeyValues(w)
}

//func readKeyValues takes stream of decrypted key values
//input : reader object which having data written on it
//return : Readable form of provided input params
func (v *Vault) readKeyValues(r io.Reader) error {
	decoder := json.NewDecoder(r)
	return decoder.Decode(&v.keyValues)
}

func (v *Vault) writeKeyValues(w io.Writer) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(&v.keyValues)
}
