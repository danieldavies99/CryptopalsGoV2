package challenge9

import (
	"errors"
	"testing"

	"github.com/danieldavies99/CryptopalsGoV2/utilities/pkcs7"
)

func TestChallenge9(t *testing.T) {
	type returnType struct {
		output []byte
		err    error
	}
	type testCase struct {
		input []byte
		padTo int
		want  returnType
	}

	paddingX4 := []byte{0x04, 0x04, 0x04, 0x04}

	tests := []testCase{
		{input: []byte("YELLOW SUBMARINE"), padTo: 20, want: returnType{output: append([]byte("YELLOW SUBMARINE"), paddingX4...), err: nil}},
		{input: []byte("YELLOW SUBMARINE"), padTo: 16, want: returnType{output: []byte("YELLOW SUBMARINE"), err: nil}},
		{input: []byte("YELLOW SUBMARINE"), padTo: 1000, want: returnType{output: nil, err: pkcs7.ErrExceedsMaxPadding}},
		{input: []byte("YELLOW SUBMARINE"), padTo: 2, want: returnType{output: nil, err: pkcs7.ErrPaddingTooSmall}},
	}

	for _, tc := range tests {
		got, err := Challenge9(tc.input, tc.padTo)
		if !errors.Is(err, tc.want.err) {
			t.Errorf("Challenge9(%q, %q) error = %v\n\n %v", tc.input, tc.padTo, err, tc.want.err)
		}
		if string(got) != string(tc.want.output) {
			t.Errorf("Challenge9(%q, %q) = %q\n\n %q", tc.input, tc.padTo, got, tc.want.output)
		}
	}
}
