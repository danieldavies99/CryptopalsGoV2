package challenge3

import (
	"errors"
	"testing"

	"github.com/danieldavies99/CryptopalsGoV2/utilities/hex"
	"github.com/danieldavies99/CryptopalsGoV2/utilities/xor"
)

func TestChallenge3(t *testing.T) {
	type returnType struct {
		output []byte
		key    byte
		err    error
	}
	type testCase struct {
		input []byte
		want  returnType
	}

	tests := []testCase{
		{input: []byte("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"), want: returnType{output: []byte("Cooking MC's like a pound of bacon"), key: byte('X'), err: nil}},
		{input: []byte("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b373"), want: returnType{output: nil, key: 0, err: hex.ErrInvalidLength}},
		{input: []byte{}, want: returnType{output: nil, key: 0, err: xor.ErrTextLengthZero}},
	}

	for _, tc := range tests {
		got, key, err := Challenge3(tc.input)
		if !errors.Is(err, tc.want.err) {
			t.Errorf("Challenge3(%q) error = %v\n\n %v", tc.input, err, tc.want.err)
		}
		if string(got) != string(tc.want.output) || string(key) != string(tc.want.key) {
			t.Errorf("Challenge3(%q) = %q, %q\n\n %q, %q", tc.input, got, key, tc.want.output, tc.want.key)
		}
	}
}
