package crypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"io"
)

type PlainText string
type CypherText string

func Encrypt(text PlainText, key []byte) (CypherText, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// convert text to raw bytes, maybe just pass this as a byte array
	plainText := []byte(text)
	cipherText := make([]byte, aes.BlockSize+len(plainText))
	iv := cipherText[:aes.BlockSize]

	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cipherText[aes.BlockSize:], plainText)

	return CypherText(hex.EncodeToString(cipherText)), nil
}

func Decrypt(text CypherText, key []byte) (PlainText, error) {
	cipherText, err := hex.DecodeString(string(text))
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	if len(cipherText) < aes.BlockSize {
		return "", errors.New("ciphertext too short")
	}

	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(cipherText, cipherText)

	return PlainText(cipherText), nil
}
