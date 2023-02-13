package src

import (
	"encoding/base32"
	"os"
)

type Secret string

func LoadFromFile(fileName string) (*Secret, error) {
	bin, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	secret := Secret(bin)
	return &secret, nil
}

func (s Secret) Base32() (*[]byte, error) {
	bin, err := base32.StdEncoding.DecodeString(string(s))
	if err != nil {
		return nil, err
	}

	return &bin, err
}
