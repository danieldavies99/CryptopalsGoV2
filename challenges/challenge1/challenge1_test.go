package challenge1

import (
	"errors"
	"testing"

	"github.com/danieldavies99/CryptopalsGoV2/utilities/hex"
)

func TestChallenge1(t *testing.T) {
	type returnType struct {
		output []byte
		err    error
	}
	type testCase struct {
		input []byte
		want  returnType
	}

	tests := []testCase{
		{input: []byte("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"), want: returnType{output: []byte("SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"), err: nil}},
		{input: []byte("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6"), want: returnType{output: nil, err: hex.ErrInvalidLength}},
	}

	for _, tc := range tests {
		got, err := Challenge1(tc.input)
		if !errors.Is(err, tc.want.err) {
			t.Errorf("Challenge1(%q) error = %v\n\n %v", tc.input, err, tc.want.err)
		}
		if string(got) != string(tc.want.output) {
			t.Errorf("Challenge1(%q) = %q\n\n %q", tc.input, got, tc.want.output)
		}
	}
}
