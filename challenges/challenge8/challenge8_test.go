package challenge8

import (
	"errors"
	"testing"

	"github.com/danieldavies99/CryptopalsGoV2/utilities/hex"
)

func TestChallenge8(t *testing.T) {
	type returnType struct {
		output int
		err    error
	}
	type testCase struct {
		filepath string
		want     returnType
	}

	tests := []testCase{
		{filepath: "./input.txt", want: returnType{output: 133, err: nil}},
		{filepath: "./inp.txt", want: returnType{output: 0, err: ErrReadingFile}},
		{filepath: "./input_invalid_hex.txt", want: returnType{output: 0, err: hex.ErrInvalidLength}},
	}

	for _, tc := range tests {
		output, err := Challenge8(tc.filepath)
		if !errors.Is(err, tc.want.err) {
			t.Errorf("Challenge8(%q) error = %v\n\n %v", tc.filepath, err, tc.want.err)
		}
		if output != tc.want.output {
			t.Errorf("Challenge8(%q) = %q\n\n %q", tc.filepath, output, tc.want.output)
		}
	}
}
