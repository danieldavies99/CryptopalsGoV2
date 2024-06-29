package hex

import (
	"errors"
	"testing"
)

func TestHexEncode(t *testing.T) {
	type testCase struct {
		input []byte
		want  []byte
	}

	tests := []testCase{
		{input: []byte("I'm killing your brain like a poisonous mushroom"), want: []byte("49276D206B696C6C696E6720796F757220627261696E206C696B65206120706F69736F6E6F7573206D757368726F6F6D")},
		{input: []byte("Wake me up when September ends."), want: []byte("57616B65206D65207570207768656E2053657074656D62657220656E64732E")},
	}

	for _, tc := range tests {
		got := HexEncode(tc.input)
		if string(got) != string(tc.want) {
			t.Errorf("HexEncode(%q) = %q\n\n %q", tc.input, got, tc.want)
		}
	}
}

func TestHexDecode(t *testing.T) {
	type returnType struct {
		decoded []byte
		err     error
	}
	type testCase struct {
		input []byte
		want  returnType
	}

	tests := []testCase{
		{input: []byte("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"), want: returnType{decoded: []byte("I'm killing your brain like a poisonous mushroom"), err: nil}},
		{input: []byte("59 6F 75 20 67 6F 74 74 61 20 70 61 79 20 74 68 65 20 74 72 6F 6C 6C 20 74 6F 6C 6C 2E"), want: returnType{decoded: []byte("You gotta pay the troll toll."), err: nil}},
		{input: []byte("492"), want: returnType{decoded: []byte{}, err: ErrInvalidLength}},
	}

	for _, tc := range tests {
		got, err := HexDecode(tc.input)
		if !errors.Is(err, tc.want.err) {
			t.Errorf("HexStringToBytes(%q) error = %v\n\n %v", tc.input, err, tc.want.err)
		}
		if string(got) != string(tc.want.decoded) {
			t.Errorf("HexStringToBytes(%q) = %q\n\n %q", tc.input, got, tc.want.decoded)
		}
	}
}
