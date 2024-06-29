package xor

import (
	"errors"
)

var ErrLengthNotEqual = errors.New("inputs must be of equal length")
var ErrTextLengthZero = errors.New("text must have length of greater than 0")
var ErrKeyLengthZero = errors.New("key must have length of greater than 0")

func XOR(a []byte, b []byte) ([]byte, error) {

	if len(a) != len(b) {
		return nil, ErrLengthNotEqual
	}

	var output []byte
	for i := 0; i < len(a); i++ {
		output = append(output, a[i]^b[i])
	}
	return output, nil
}

func XORAgainstOneByte(text []byte, key byte) ([]byte, error) {

	if len(text) == 0 {
		return []byte{}, ErrTextLengthZero
	}

	var output []byte
	for i := 0; i < len(text); i++ {
		output = append(output, text[i]^key)
	}

	return output, nil
}

func XORRepeatingKey(text []byte, key []byte) ([]byte, error) {

	if len(text) == 0 {
		return []byte{}, ErrTextLengthZero
	}

	if len(key) == 0 {
		return []byte{}, ErrKeyLengthZero
	}

	var output []byte
	keyIndex := 0
	for i := 0; i < len(text); i++ {
		output = append(output, text[i]^key[keyIndex])
		keyIndex = (keyIndex + 1) % len(key)
	}
	return output, nil
}
