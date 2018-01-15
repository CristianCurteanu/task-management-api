package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
)

type Hashing struct {
	Key   []byte
	Value []byte
}

func (e *Hashing) Encrypt() ([]byte, error) {
	block, err := aes.NewCipher(e.Key)
	if err != nil {
		return nil, err
	}
	b := base64.StdEncoding.EncodeToString(e.Value)
	ciphertext := make([]byte, aes.BlockSize+len(b))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}
	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(ciphertext[aes.BlockSize:], []byte(b))
	return ciphertext, nil
}

func (e *Hashing) Decrypt() ([]byte, error) {
	block, err := aes.NewCipher(e.Key)
	if err != nil {
		return nil, err
	}
	if len(e.Value) < aes.BlockSize {
		return nil, errors.New("ciphertext too short")
	}
	iv := e.Value[:aes.BlockSize]
	e.Value = e.Value[aes.BlockSize:]
	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(e.Value, e.Value)
	data, err := base64.StdEncoding.DecodeString(string(e.Value))
	if err != nil {
		return nil, err
	}
	return data, nil
}
