package editdistance

import (
	"errors"
	"testing"
)

func TestEditDistance(t *testing.T) {
	type returnType struct {
		editDistance int
		err          error
	}
	type testCase struct {
		inputA []byte
		inputB []byte
		want   returnType
	}

	tests := []testCase{
		{inputA: []byte("this is a test"), inputB: []byte("wokka wokka!!!"), want: returnType{editDistance: 37, err: nil}},
		{inputA: []byte("England"), inputB: []byte("dnalgnE"), want: returnType{editDistance: 8, err: nil}},
		{inputA: []byte("one"), inputB: []byte("four"), want: returnType{editDistance: 0, err: ErrLengthNotEqual}},
	}

	for _, tc := range tests {
		got, err := EditDistance(tc.inputA, tc.inputB)
		if !errors.Is(err, tc.want.err) {
			t.Errorf("EditDistance(%q %q) error = %v\n\n %v", tc.inputA, tc.inputB, err, tc.want.err)
		}
		if got != tc.want.editDistance {
			t.Errorf("EditDistance(%q %q) = %d, want %d", tc.inputA, tc.inputB, got, tc.want.editDistance)
		}
	}
}
