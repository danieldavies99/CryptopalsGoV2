package ecb

import (
	"crypto/aes"
	"errors"
)

var ErrCipherTextLength = errors.New("cipher text length must be multiple of 16")
var ErrPlainTextLength = errors.New("plain text length must be multiple of 16")

func ECBDecrypt(cipherText []byte, key []byte) ([]byte, error) {
	if len(cipherText)%16 != 0 {
		return nil, ErrCipherTextLength
	}

	cipher, err := aes.NewCipher([]byte(key))
	if err != nil {
		return nil, err
	}
	decrypted := make([]byte, len(cipherText))
	size := 16

	for bs, be := 0, size; bs < len(cipherText); bs, be = bs+size, be+size {
		cipher.Decrypt(decrypted[bs:be], cipherText[bs:be])
	}

	return decrypted, nil
}

func ECBEncrypt(cipherText []byte, key []byte) ([]byte, error) {
	cipher, err := aes.NewCipher([]byte(key))
	if err != nil {
		return nil, err
	}
	encrypted := make([]byte, len(cipherText))
	size := 16

	if len(cipherText)%size != 0 {
		return nil, ErrPlainTextLength
	}

	for bs, be := 0, size; bs < len(cipherText); bs, be = bs+size, be+size {
		cipher.Encrypt(encrypted[bs:be], cipherText[bs:be])
	}

	return encrypted, nil
}
