package cbc

import (
	"errors"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/danieldavies99/CryptopalsGoV2/utilities/base64"
	"github.com/danieldavies99/CryptopalsGoV2/utilities/pkcs7"
)

func TestCBCDecrypt(t *testing.T) {
	type returnType struct {
		decoded []byte
		err     error
	}
	type testCase struct {
		input []byte
		key   []byte
		iv    []byte
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
		{
			input: base64Decoded,
			key:   []byte("YELLOW SUBMARINE"),
			iv:    []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			want:  returnType{decoded: decoded, err: nil},
		},
		{
			input: base64Decoded,
			key:   []byte("YELLOW SUBMARIN"),
			iv:    []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			want:  returnType{decoded: nil, err: ErrKeyLength},
		},
		{
			input: base64Decoded,
			key:   []byte("YELLOW SUBMARINE"),
			iv:    []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			want:  returnType{decoded: nil, err: ErrIVLength},
		},
		{
			input: append(base64Decoded, 0),
			key:   []byte("YELLOW SUBMARINE"),
			iv:    []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			want:  returnType{decoded: nil, err: ErrCypherTextLength},
		},
	}

	for _, tc := range tests {
		got, err := CBCDecrypt(tc.input, tc.key, tc.iv)
		// strip padding of output before compare
		if got != nil {
			got = pkcs7.StripPadding(got)
		}
		if !errors.Is(err, tc.want.err) {
			t.Errorf("CBCDecrypt(%q, %q, %q) error = %v\n\n %v", tc.input, tc.key, tc.iv, err, tc.want.err)
		}
		if string(got) != string(tc.want.decoded) {
			t.Errorf("CBCDecrypt(%q, %q, %q) = %q\n\n %q", tc.input, tc.key, tc.iv, got, tc.want.decoded)
		}
	}
}

func TestCBCEncrypt(t *testing.T) {
	type returnType struct {
		encrypted []byte
		err       error
	}
	type testCase struct {
		input []byte
		key   []byte
		iv    []byte
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
		{
			input: paddedPlaintext,
			key:   []byte("YELLOW SUBMARINE"),
			iv:    []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			want:  returnType{encrypted: base64DecodedInput, err: nil},
		},
		{
			input: paddedPlaintext,
			key:   []byte("YELLOW SUBMARIN"),
			iv:    []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			want:  returnType{encrypted: nil, err: ErrKeyLength},
		},
		{
			input: paddedPlaintext,
			key:   []byte("YELLOW SUBMARINE"),
			iv:    []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			want:  returnType{encrypted: nil, err: ErrIVLength},
		},
	}

	for _, tc := range tests {
		got, err := CBCEncrypt(tc.input, tc.key, tc.iv)
		if !errors.Is(err, tc.want.err) {
			t.Errorf("CBCEncrypt(%q, %q, %q) error = %v\n\n %v", tc.input, tc.key, tc.iv, err, tc.want.err)
		}
		if string(got) != string(tc.want.encrypted) {
			t.Errorf("CBCEncrypt(%q, %q, %q) = %q\n\n %q", tc.input, tc.key, tc.iv, got, tc.want.encrypted)
		}
	}
}
