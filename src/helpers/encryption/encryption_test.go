package encryption

import (
	"testing"
)

func TestEncryption(t *testing.T) {
	key := "70C488A7-9D99-4043-9765-C07917C6"
	enc := Hashing{Key: []byte(key), Value: []byte("some value")}
	encrypted, err := enc.Encrypt()
	if err != nil {
		t.Fatalf(err.Error())
	}
	if len(encrypted) == 0 {
		t.Fatalf("Encryption erronated")
	}

	decr := Hashing{Key: []byte(key), Value: encrypted}
	decrypted, err := decr.Decrypt()
	if err != nil {
		t.Fatalf(err.Error())
	}

	if string(enc.Value[:]) != string(decrypted[:]) {
		t.Fatalf("Decryption gone wrong")
	}
}
