package pkcs7

import "errors"

var ErrInvalidPadding = errors.New("invalid padding")
var ErrExceedsMaxPadding = errors.New("padding exceed maximum (255 bytes)")
var ErrPaddingTooSmall = errors.New("padded length cannot be shorter than original")

func StripPadding(input []byte) []byte {
	lastByte := input[len(input)-1]
	numPadding := int(lastByte)
	if numPadding > len(input) {
		// padding is longer than input, something has gone wrong, just return input
		return input
	}
	paddingBytes := input[len(input)-numPadding:]
	for i := 0; i < numPadding; i++ {
		if paddingBytes[i] != lastByte {
			// padding not detected
			return input
		}
	}
	return input[:len(input)-numPadding]
}

func Pad(input []byte, padTo int) ([]byte, error) {
	if padTo < len(input) {
		return nil, ErrPaddingTooSmall
	}
	paddingAmount := padTo - len(input)
	if paddingAmount > 255 {
		return nil, ErrExceedsMaxPadding
	}
	paddingByte := byte(paddingAmount)
	for i := 0; i < paddingAmount; i++ {
		input = append(input, paddingByte)
	}
	return input, nil
}
