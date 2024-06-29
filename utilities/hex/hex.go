package hex

import (
	"errors"
	"strings"
)

var ErrInvalidLength = errors.New("input length must be divisible by 2")

// as with my base64 encoder/decoder
// In the real world I would use a hex package
// to do this but I figured it would be interesting
// to try and write my own encoder/decoder using
// bitwise operations

func HexDecode(input []byte) ([]byte, error) {
	hex2ByteMap := map[byte]int{
		'0': 0, '1': 1, '2': 2, '3': 3, '4': 4, '5': 5, '6': 6, '7': 7,
		'8': 8, '9': 9, 'A': 10, 'B': 11, 'C': 12, 'D': 13, 'E': 14, 'F': 15,
	}

	// convert hex string to uppercase
	// as there is no standard for upper vs
	// lower case characters in hex
	hexString := string(input)
	hexString = strings.ToUpper(hexString)

	// sometimes hex strings are supplied with spaces between the pairs
	// but our function assumes there aren't any
	hexString = strings.ReplaceAll(hexString, " ", "")

	hexBytes := []byte(hexString)

	if len(hexBytes)%2 != 0 {
		return nil, ErrInvalidLength
	}

	var output []byte
	for i := 0; i < len(hexBytes); i += 2 {
		var convertedByte uint8

		firstHalf := hex2ByteMap[hexBytes[i]] << 4
		secondHalf := hex2ByteMap[hexBytes[i+1]]
		convertedByte |= uint8(firstHalf) | uint8(secondHalf)

		output = append(output, convertedByte)
	}

	return output, nil
}

func HexEncode(input []byte) []byte {

	byte2HexMap := map[int]byte{
		0: '0', 1: '1', 2: '2', 3: '3', 4: '4', 5: '5', 6: '6', 7: '7',
		8: '8', 9: '9', 10: 'A', 11: 'B', 12: 'C', 13: 'D', 14: 'E', 15: 'F',
	}

	var output []byte
	for i := 0; i < len(input); i++ {
		p1 := input[i] >> 4
		var mask = byte(15) // 1111
		p2 := input[i] & mask

		output = append(output, byte2HexMap[int(p1)])
		output = append(output, byte2HexMap[int(p2)])
	}

	return output
}
