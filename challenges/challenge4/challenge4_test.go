package challenge4

import (
	"errors"
	"testing"
)

func TestChallenge4(t *testing.T) {
	type returnType struct {
		output []byte
		err    error
	}
	type testCase struct {
		filepath string
		want     returnType
	}

	tests := []testCase{
		{filepath: "./input.txt", want: returnType{output: []byte("Now that the party is jumping\n"), err: nil}},
	}

	for _, tc := range tests {
		got, err := Challenge4(tc.filepath)
		if !errors.Is(err, tc.want.err) {
			t.Errorf("Challenge3(%q) error = %v\n\n %v", tc.filepath, err, tc.want.err)
		}
		if string(got) != string(tc.want.output) {
			t.Errorf("Challenge3(%q) = %q\n\n %q", tc.filepath, got, tc.want.output)
		}
	}
}
