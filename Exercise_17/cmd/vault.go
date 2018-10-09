package vault

import (
	"errors"
)

type Vault struct {
	key       string
	keyValues map[string]string
}

func (v *Vault) Get(key string) (string, error) {
	value := v.keyValues[key]
	if value == "" {
		return "", errors.New("No value found.")
	}
	return value, nil
}

func (v *Vault) Set(key string) error {
	return nil
}
