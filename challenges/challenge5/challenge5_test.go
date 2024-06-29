package challenge5

import (
	"errors"
	"testing"
)

func TestChallenge5(t *testing.T) {
	type returnType struct {
		output []byte
		err    error
	}
	type testCase struct {
		input []byte
		key   []byte
		want  returnType
	}

	tests := []testCase{
		{
			input: []byte("Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal"),
			key:   []byte("ICE"),
			want:  returnType{output: []byte("0B3637272A2B2E63622C2E69692A23693A2A3C6324202D623D63343C2A26226324272765272A282B2F20430A652E2C652A3124333A653E2B2027630C692B20283165286326302E27282F"), err: nil}},
	}

	for _, tc := range tests {
		got, err := Challenge5(tc.input, tc.key)
		if !errors.Is(err, tc.want.err) {
			t.Errorf("Challenge5(%q, %q) error = %v\n\n %v", tc.input, tc.key, err, tc.want.err)
		}
		if string(got) != string(tc.want.output) {
			t.Errorf("Challenge5(%q, %q) = %q\n\n %q", tc.input, tc.key, got, tc.want.output)
		}
	}
}
