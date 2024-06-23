package crypt

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncryptDecrypt(t *testing.T) {

	// mock key
	key := []byte("e21f63fbcfbc73c8813b69f71c4c05eb44fe33998dc01665dbe5e6d7e512ea19")
	key = key[:32]

	plainText := PlainText("Plaintext stuff goes here")

	cypherText, err := Encrypt(plainText, key)

	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	decryptedCypherText, err := Decrypt(cypherText, key)

	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	assert.Equal(t, plainText, decryptedCypherText)

}
