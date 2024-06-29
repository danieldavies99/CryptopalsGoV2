package challenge10

import (
	"errors"
	"io/ioutil"
	"testing"

	"github.com/danieldavies99/CryptopalsGoV2/utilities/pkcs7"
)

func TestChallenge10(t *testing.T) {
	type returnType struct {
		decoded []byte
		err     error
	}
	type testCase struct {
		filepath string
		key      []byte
		iv       []byte
		want     returnType
	}

	expected, err := ioutil.ReadFile("./decrypted.txt")
	if err != nil {
		t.Errorf("failed to read decrypted: %v", err)
	}

	tests := []testCase{
		{
			filepath: "./input.txt",
			key:      []byte("YELLOW SUBMARINE"),
			iv:       []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			want:     returnType{decoded: expected, err: nil},
		},
	}

	for _, tc := range tests {
		got, err := Challenge10(tc.filepath, tc.key, tc.iv)
		// strip padding of output before compare
		if got != nil {
			got = pkcs7.StripPadding(got)
		}
		if !errors.Is(err, tc.want.err) {
			t.Errorf("Challenge10(%q, %q, %q) error = %v\n\n %v", tc.filepath, tc.key, tc.iv, err, tc.want.err)
		}
		if string(got) != string(tc.want.decoded) {
			t.Errorf("Challenge10(%q, %q, %q) = %q\n\n %q", tc.filepath, tc.key, tc.iv, got, tc.want.decoded)
		}
	}
}
