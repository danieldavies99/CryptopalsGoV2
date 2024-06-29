package xor

import (
	"errors"
	"testing"

	"github.com/danieldavies99/CryptopalsGoV2/utilities/base64"
	"github.com/danieldavies99/CryptopalsGoV2/utilities/hex"
)

func TestXOR(t *testing.T) {
	type returnType struct {
		result []byte
		err    error
	}
	type testCase struct {
		inputA []byte
		inputB []byte
		want   returnType
	}

	inputA, _ := hex.HexDecode([]byte("1c0111001f010100061a024b53535009181c"))
	inputB, _ := hex.HexDecode([]byte("686974207468652062756c6c277320657965"))

	tests := []testCase{
		{
			inputA: inputA,
			inputB: inputB,
			want:   returnType{result: []byte("the kid don't play"), err: nil},
		},
		{
			inputA: inputA,
			inputB: append(inputB, byte(0)),
			want:   returnType{result: []byte{}, err: ErrLengthNotEqual},
		},
	}

	for _, tc := range tests {
		got, err := XOR(tc.inputA, tc.inputB)
		if !errors.Is(err, tc.want.err) {
			t.Errorf("XOR(%q, %q) error = %v\n\n %v", tc.inputA, tc.inputB, err, tc.want.err)
		}
		if string(got) != string(tc.want.result) {
			t.Errorf("XOR(%q, %q) = %q\n\n %q", tc.inputA, tc.inputB, got, tc.want.result)
		}
	}
}

func TestXORAgainstOneByte(t *testing.T) {

	type returnType struct {
		result []byte
		err    error
	}

	type testCase struct {
		inputA []byte
		key    byte
		want   returnType
	}

	inputA, _ := base64.Base64Decode([]byte("FTk2IXgwOTY8K3g1OTM9eDQxPzAseC83KjN2"))
	tests := []testCase{
		{inputA: []byte(inputA), key: byte('X'), want: returnType{result: []byte("Many hands make light work."), err: nil}},
		{inputA: []byte{}, key: byte('X'), want: returnType{result: []byte{}, err: ErrTextLengthZero}},
	}

	for _, tc := range tests {
		got, err := XORAgainstOneByte(tc.inputA, tc.key)
		if !errors.Is(err, tc.want.err) {
			t.Errorf("XORAgainstOneByte(%q, %q) error = %v\n\n %v", tc.inputA, tc.key, err, tc.want.err)
		}
		if string(got) != string(tc.want.result) {
			t.Errorf("XORAgainstOneByte(%q, %q) = %q\n\n %q", tc.inputA, tc.key, got, tc.want.result)
		}
	}
}

func TestXORRepeatingKey(t *testing.T) {

	type returnType struct {
		result []byte
		err    error
	}

	type testCase struct {
		input []byte
		key   []byte
		want  returnType
	}

	res, _ := hex.HexDecode([]byte("0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f"))
	tests := []testCase{
		{input: []byte("Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal"), key: []byte("ICE"), want: returnType{result: res, err: nil}},
		{input: []byte("Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal"), key: []byte{}, want: returnType{result: []byte{}, err: ErrKeyLengthZero}},
		{input: []byte{}, key: []byte("ICE"), want: returnType{result: []byte{}, err: ErrTextLengthZero}},
		{input: []byte{}, key: []byte{}, want: returnType{result: []byte{}, err: ErrTextLengthZero}},
	}

	for _, tc := range tests {
		got, err := XORRepeatingKey(tc.input, tc.key)
		if !errors.Is(err, tc.want.err) {
			t.Errorf("XORAgainstOneByte(%q, %q) error = %v\n\n %v", tc.input, tc.key, err, tc.want.err)
		}
		if string(got) != string(tc.want.result) {
			t.Errorf("XORAgainstOneByte(%q, %q) = %q\n\n %q", tc.input, tc.key, got, tc.want.result)
		}
	}
}
