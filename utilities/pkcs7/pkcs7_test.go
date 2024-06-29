package pkcs7

import (
	"errors"
	"testing"
)

func TestStripPadding(t *testing.T) {
	type returnType struct {
		output []byte
	}
	type testCase struct {
		input []byte
		want  returnType
	}

	paddingX4 := []byte{0x04, 0x04, 0x04, 0x04}
	paddingX7 := []byte{0x07, 0x07, 0x07, 0x07, 0x07, 0x07, 0x07}

	tests := []testCase{
		{input: append([]byte("Hello World"), paddingX4...), want: returnType{output: []byte("Hello World")}},
		{input: append([]byte("Hello World"), paddingX7...), want: returnType{output: []byte("Hello World")}},
		{input: []byte("Hello World"), want: returnType{output: []byte("Hello World")}},
	}

	for _, tc := range tests {
		got := StripPadding(tc.input)
		if string(got) != string(tc.want.output) {
			t.Errorf("StripPadding(%q) = %q\n\n %q", tc.input, got, tc.want.output)
		}
	}
}

func TestPad(t *testing.T) {
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
		{input: []byte("YELLOW SUBMARINE"), padTo: 1000, want: returnType{output: nil, err: ErrExceedsMaxPadding}},
		{input: []byte("YELLOW SUBMARINE"), padTo: 2, want: returnType{output: nil, err: ErrPaddingTooSmall}},
	}

	for _, tc := range tests {
		got, err := Pad(tc.input, tc.padTo)
		if !errors.Is(err, tc.want.err) {
			t.Errorf("Pad(%q, %q) error = %v\n\n %v", tc.input, tc.padTo, err, tc.want.err)
		}
		if string(got) != string(tc.want.output) {
			t.Errorf("Pad(%q, %q) = %q\n\n %q", tc.input, tc.padTo, got, tc.want.output)
		}
	}
}
