package challenge6

import (
	"errors"
	"io/ioutil"
	"testing"
)

func TestChallenge4(t *testing.T) {
	type returnType struct {
		key    []byte
		output []byte
		err    error
	}
	type testCase struct {
		filepath string
		want     returnType
	}

	decodedoutput, err := ioutil.ReadFile("./decrypted.txt")
	if err != nil {
		t.Errorf("failed to read decrypted.txt: %v", err)
	}

	tests := []testCase{
		{filepath: "./input.txt", want: returnType{key: []byte("Terminator X: Bring the noise"), output: decodedoutput, err: nil}},
	}

	for _, tc := range tests {
		key, text, err := Challenge6(tc.filepath)
		if !errors.Is(err, tc.want.err) {
			t.Errorf("Challenge6(%q) error = %v\n\n %v", tc.filepath, err, tc.want.err)
		}
		if string(text) != string(tc.want.output) || string(key) != string(tc.want.key) {
			t.Errorf("Challenge6(%q) = %q, %q\n\n %q, %q", tc.filepath, key, text, tc.want.key, tc.want.output)
		}
	}
}
