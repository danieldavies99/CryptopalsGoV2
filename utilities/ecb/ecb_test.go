package ecb

import (
	"crypto/aes"
	"errors"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/danieldavies99/CryptopalsGoV2/utilities/base64"
	"github.com/danieldavies99/CryptopalsGoV2/utilities/pkcs7"
)

func TestECBDecrypt(t *testing.T) {
	type returnType struct {
		decoded []byte
		err     error
	}
	type testCase struct {
		input []byte
		key   []byte
		want  returnType
	}

	fileContent, err := ioutil.ReadFile("./test_encrypted.txt")
	if err != nil {
		fmt.Println("errror reading file", err)
	}

	base64Decoded, err := base64.Base64Decode(fileContent)
	if err != nil {
		fmt.Println("errror b64 decoding file", err)
	}

	decoded, err := ioutil.ReadFile("./test_decrypted.txt")
	if err != nil {
		fmt.Println("errror reading file", err)
	}

	tests := []testCase{
		{input: base64Decoded, key: []byte("YELLOW SUBMARINE"), want: returnType{decoded: decoded, err: nil}},
		{input: append(base64Decoded, byte('0')), key: []byte("YELLOW SUBMARINE"), want: returnType{decoded: nil, err: ErrCipherTextLength}},
		{input: base64Decoded, key: []byte("YELLOW SUBMARINES"), want: returnType{decoded: nil, err: aes.KeySizeError(17)}},
	}

	for _, tc := range tests {
		got, err := ECBDecrypt(tc.input, tc.key)
		// strip padding of output before compare
		if got != nil {
			got = pkcs7.StripPadding(got)
		}
		if !errors.Is(err, tc.want.err) {
			t.Errorf("ECBDecrypt(%q, %q) error = %v\n\n %v", tc.input, tc.key, err, tc.want.err)
		}
		if string(got) != string(tc.want.decoded) {
			t.Errorf("ECBDecrypt(%q, %q) = %q\n\n %q", tc.input, tc.key, got, tc.want.decoded)
		}
	}
}

func TestECBEncrypt(t *testing.T) {
	type returnType struct {
		encrypted []byte
		err       error
	}
	type testCase struct {
		input []byte
		key   []byte
		want  returnType
	}

	encryptedInput, err := ioutil.ReadFile("./test_encrypted.txt")
	if err != nil {
		fmt.Println("errror reading file", err)
	}

	base64DecodedInput, err := base64.Base64Decode(encryptedInput)
	if err != nil {
		fmt.Println("errror b64 decoding file", err)
	}

	plaintext, err := ioutil.ReadFile("./test_decrypted.txt")
	if err != nil {
		fmt.Println("errror reading file", err)
	}
	lengthToPadTo := (16 - len(plaintext)%16) + len(plaintext)
	fmt.Println(lengthToPadTo)
	paddedPlaintext, err := pkcs7.Pad(plaintext, lengthToPadTo)
	if err != nil {
		fmt.Println("error padding", err)
	}

	tests := []testCase{
		{input: paddedPlaintext, key: []byte("YELLOW SUBMARINE"), want: returnType{encrypted: base64DecodedInput, err: nil}},
		{input: append(paddedPlaintext, byte('0')), key: []byte("YELLOW SUBMARINE"), want: returnType{encrypted: nil, err: ErrPlainTextLength}},
		{input: paddedPlaintext, key: []byte("YELLOW SUBMARINES"), want: returnType{encrypted: nil, err: aes.KeySizeError(17)}},
	}

	for _, tc := range tests {
		got, err := ECBEncrypt(tc.input, tc.key)
		if !errors.Is(err, tc.want.err) {
			t.Errorf("ECBEncrypt(%q, %q) error = %v\n\n %v", tc.input, tc.key, err, tc.want.err)
		}
		if string(got) != string(tc.want.encrypted) {
			t.Errorf("ECBEncrypt(%q, %q) = %q\n\n %q", tc.input, tc.key, got, tc.want.encrypted)
		}
	}
}
