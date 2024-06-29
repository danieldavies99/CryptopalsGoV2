package challenge2

import (
	"errors"
	"testing"

	"github.com/danieldavies99/CryptopalsGoV2/utilities/hex"
	"github.com/danieldavies99/CryptopalsGoV2/utilities/xor"
)

func TestChallenge2(t *testing.T) {
	type returnType struct {
		output []byte
		err    error
	}
	type testCase struct {
		inputA []byte
		inputB []byte
		want   returnType
	}

	tests := []testCase{
		{
			inputA: []byte("1c0111001f010100061a024b53535009181c"),
			inputB: []byte("686974207468652062756c6c277320657965"),
			want:   returnType{output: []byte("746865206B696420646F6E277420706C6179"), err: nil},
		},
		{
			inputA: []byte("1c0111001f010100061a024b53535009181"),
			inputB: []byte("686974207468652062756c6c277320657965"),
			want:   returnType{output: nil, err: hex.ErrInvalidLength},
		},
		{
			inputA: []byte("1c0111001f010100061a024b53535009181c"),
			inputB: []byte("686974207468652062756c6c27732065796"),
			want:   returnType{output: nil, err: hex.ErrInvalidLength},
		},
		{
			inputA: []byte("1c0111001f010100061a024b53535009181c"),
			inputB: []byte("686974207468652062756c6c27732065791c1c"),
			want:   returnType{output: nil, err: xor.ErrLengthNotEqual},
		},
	}

	for _, tc := range tests {
		got, err := Challenge2(tc.inputA, tc.inputB)
		if !errors.Is(err, tc.want.err) {
			t.Errorf("Challenge2(%q, %q) error = %v\n\n %v", tc.inputA, tc.inputB, err, tc.want.err)
		}
		if string(got) != string(tc.want.output) {
			t.Errorf("Challenge2(%q, %q) = %q\n\n %q", tc.inputA, tc.inputB, got, tc.want.output)
		}
	}
}
