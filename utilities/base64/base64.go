package base64

import (
	"bytes"
	"errors"
)

var ErrDecodeInvalidLength = errors.New("input length must be divisible by 4")

// In the real world I would use a base64 package
// to do this but I figured it would be interesting
// to try and write my own encoder/decoder using
// bitwise operations

func Base64Encode(input []byte) []byte {
	encodeMap := map[int]byte{
		0: 'A', 1: 'B', 2: 'C', 3: 'D', 4: 'E', 5: 'F', 6: 'G',
		7: 'H', 8: 'I', 9: 'J', 10: 'K', 11: 'L', 12: 'M', 13: 'N',
		14: 'O', 15: 'P', 16: 'Q', 17: 'R', 18: 'S', 19: 'T', 20: 'U',
		21: 'V', 22: 'W', 23: 'X', 24: 'Y', 25: 'Z', 26: 'a', 27: 'b',
		28: 'c', 29: 'd', 30: 'e', 31: 'f', 32: 'g', 33: 'h', 34: 'i',
		35: 'j', 36: 'k', 37: 'l', 38: 'm', 39: 'n', 40: 'o', 41: 'p',
		42: 'q', 43: 'r', 44: 's', 45: 't', 46: 'u', 47: 'v', 48: 'w',
		49: 'x', 50: 'y', 51: 'z', 52: '0', 53: '1', 54: '2', 55: '3',
		56: '4', 57: '5', 58: '6', 59: '7', 60: '8', 61: '9', 62: '+',
		63: '/',
	}

	var output []byte
	for i := 0; i < len(input); i += 3 {
		// put 3 bytes into a 24-bit chunk
		// uint32  is actually used because
		// there is no uint24 type so we just
		// ignore the first 8 bits
		var chunk uint32
		chunk |= uint32(input[i])
		chunk = chunk << 8
		if i+1 < len(input) {
			chunk |= uint32(input[i+1])
		}
		chunk = chunk << 8
		if i+2 < len(input) {
			chunk |= uint32(input[i+2])
		}

		var mask uint32 = 63 // 111111
		// fmt.Printf("%024b\n", chunk)
		p4 := chunk & mask
		chunk = chunk >> 6
		p3 := chunk & mask
		chunk = chunk >> 6
		p2 := chunk & mask
		chunk = chunk >> 6
		p1 := chunk & mask

		// fmt.Printf("%06b %06b %06b %06b\n", p1, p2, p3, p4)

		// three if statements are needed for
		// the three possible states of padding
		if i+1 >= len(input) {
			output = append(
				output,
				encodeMap[int(p1)],
				encodeMap[int(p2)],
				'=',
				'=',
			)
		} else if i+2 >= len(input) {
			output = append(
				output,
				encodeMap[int(p1)],
				encodeMap[int(p2)],
				encodeMap[int(p3)],
				'=',
			)
		} else {
			output = append(
				output,
				encodeMap[int(p1)],
				encodeMap[int(p2)],
				encodeMap[int(p3)],
				encodeMap[int(p4)],
			)
		}

	}

	return output
}

func Base64Decode(input []byte) ([]byte, error) {
	decodeMap := map[byte]int{
		'A': 0, 'B': 1, 'C': 2, 'D': 3, 'E': 4, 'F': 5, 'G': 6,
		'H': 7, 'I': 8, 'J': 9, 'K': 10, 'L': 11, 'M': 12, 'N': 13,
		'O': 14, 'P': 15, 'Q': 16, 'R': 17, 'S': 18, 'T': 19, 'U': 20,
		'V': 21, 'W': 22, 'X': 23, 'Y': 24, 'Z': 25, 'a': 26, 'b': 27,
		'c': 28, 'd': 29, 'e': 30, 'f': 31, 'g': 32, 'h': 33, 'i': 34,
		'j': 35, 'k': 36, 'l': 37, 'm': 38, 'n': 39, 'o': 40, 'p': 41,
		'q': 42, 'r': 43, 's': 44, 't': 45, 'u': 46, 'v': 47, 'w': 48,
		'x': 49, 'y': 50, 'z': 51, '0': 52, '1': 53, '2': 54, '3': 55,
		'4': 56, '5': 57, '6': 58, '7': 59, '8': 60, '9': 61, '+': 62,
		'/': 63,
	}

	// remove whitespace
	input = bytes.ReplaceAll(input, []byte{'\r'}, []byte{})
	input = bytes.ReplaceAll(input, []byte{'\n'}, []byte{})
	input = bytes.ReplaceAll(input, []byte{' '}, []byte{})

	if len(input)%4 != 0 {
		return nil, ErrDecodeInvalidLength
	}

	var output []byte
	for i := 0; i < len(input); i += 4 {
		var chunk uint32
		chunk |= uint32(decodeMap[input[i]])
		chunk = chunk << 6
		chunk |= uint32(decodeMap[input[i+1]])
		chunk = chunk << 6
		chunk |= uint32(decodeMap[input[i+2]])
		chunk = chunk << 6
		chunk |= uint32(decodeMap[input[i+3]])

		// fmt.Printf("%24b\n", chunk)

		var mask uint32 = 255 // 11111111
		p3 := chunk & mask
		chunk = chunk >> 8
		p2 := chunk & mask
		chunk = chunk >> 8
		p1 := chunk & mask
		if input[i+2] == '=' {
			output = append(output, byte(p1))
		} else if input[i+3] == '=' {
			output = append(output, byte(p1), byte(p2))
		} else {
			output = append(output, byte(p1), byte(p2), byte(p3))
		}
	}
	return output, nil
}
