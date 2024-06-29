package editdistance

import "errors"

var ErrLengthNotEqual = errors.New("inputs must be of equal length")

func EditDistance(a []byte, b []byte) (int, error) {

	if len(a) != len(b) {
		return 0, ErrLengthNotEqual
	}

	var output int
	for i := 0; i < len(a); i++ {
		result := uint8(a[i] ^ b[i])
		count := 0
		for result > 0 {
			// 1 in binary 00000001 so if the rightmost bit of result is 1
			// then when we mask result with 1 we will get a decimal value of 1
			// keep doing this until there are no 1 ones left in result
			// (e.g. result will be 00000000) and for loop will end
			if (result & 1) == 1 {
				count++
			}
			result = result >> 1
		}
		output += count
	}
	return output, nil
}
