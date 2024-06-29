package base64

import (
	"errors"
	"testing"
)

func TestBase64Encode(t *testing.T) {
	type testCase struct {
		input []byte
		want  []byte
	}

	tests := []testCase{
		{input: []byte("Many hands make light work."), want: []byte("TWFueSBoYW5kcyBtYWtlIGxpZ2h0IHdvcmsu")},
		{input: []byte("Wake me up when September ends."), want: []byte("V2FrZSBtZSB1cCB3aGVuIFNlcHRlbWJlciBlbmRzLg==")},
		{input: []byte("This sentence has a single padding character"), want: []byte("VGhpcyBzZW50ZW5jZSBoYXMgYSBzaW5nbGUgcGFkZGluZyBjaGFyYWN0ZXI=")},
	}

	for _, tc := range tests {
		got := Base64Encode(tc.input)
		if string(got) != string(tc.want) {
			t.Errorf("Base64Encode(%q) = %q\n\n %q", tc.input, got, tc.want)
		}
	}
}

func TestBase64Decode(t *testing.T) {
	type returnType struct {
		decoded []byte
		err     error
	}
	type testCase struct {
		input []byte
		want  returnType
	}

	tests := []testCase{
		{input: []byte("TWFueSBoYW5kcyBtYWtlIGxpZ2h0IHdvcmsu"), want: returnType{decoded: []byte("Many hands make light work."), err: nil}},
		{input: []byte("V2FrZSBtZSB1cCB3aGVuIFNlcHRlbWJlciBlbmRzLg=="), want: returnType{decoded: []byte("Wake me up when September ends."), err: nil}},
		{input: []byte("VGhpcyBzZW50ZW5jZSBoYXMgYSBzaW5nbGUgcGFkZGluZyBjaGFyYWN0ZXI="), want: returnType{decoded: []byte("This sentence has a single padding character"), err: nil}},
		{input: []byte("VGhpcyBzZW50ZW5jZSBoYXMgYSBzaW5nbGUgcGFkZGluZyBjaGFyYWN0ZXI"), want: returnType{decoded: []byte{}, err: ErrDecodeInvalidLength}},
	}

	for _, tc := range tests {
		got, err := Base64Decode(tc.input)
		if !errors.Is(err, tc.want.err) {
			t.Errorf("Base64Decode(%q) error = %v\n\n %v", tc.input, err, tc.want.err)
		}
		if string(got) != string(tc.want.decoded) {
			t.Errorf("Base64Decode(%q) = %q\n\n %q", tc.input, got, tc.want.decoded)
		}
	}
}
