package challenge7

import (
	"crypto/aes"
	"errors"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/danieldavies99/CryptopalsGoV2/utilities/base64"
)

func TestChallenge7(t *testing.T) {
	type returnType struct {
		output []byte
		err    error
	}
	type testCase struct {
		filepath string
		key      []byte
		want     returnType
	}

	decodedoutput, err := ioutil.ReadFile("./decrypted.txt")
	if err != nil {
		t.Errorf("failed to read decoded.txt: %v", err)
	}

	tests := []testCase{
		{filepath: "./input.txt", key: []byte("YELLOW SUBMARINE"), want: returnType{output: decodedoutput, err: nil}},
		{filepath: "./inp.txt", key: []byte("YELLOW SUBMARINE"), want: returnType{output: nil, err: ErrReadingFile}},
		{filepath: "./input_invalid_b64.txt", key: []byte("YELLOW SUBMARINE"), want: returnType{output: nil, err: base64.ErrDecodeInvalidLength}},
		{filepath: "./input.txt", key: []byte("YELLOW SUBMARINES"), want: returnType{output: nil, err: aes.KeySizeError(17)}},
	}

	for _, tc := range tests {
		got, err := Challenge7(tc.filepath, tc.key)
		if !errors.Is(err, tc.want.err) {
			fmt.Println(err.Error())
			t.Errorf("Challenge7(%q, %q) error = %v\n\n %v", tc.filepath, tc.key, err, tc.want.err)
		}
		if string(got) != string(tc.want.output) {
			t.Errorf("Challenge7(%q, %q) = %q\n\n %q", tc.filepath, tc.key, got, tc.want.output)
		}
	}
}
