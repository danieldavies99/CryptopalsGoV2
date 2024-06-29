package challenge5

import (
	"github.com/danieldavies99/CryptopalsGoV2/utilities/hex"
	"github.com/danieldavies99/CryptopalsGoV2/utilities/xor"
)

func Challenge5(input []byte, key []byte) ([]byte, error) {
	xorRes, err := xor.XORRepeatingKey(input, key)
	if err != nil {
		return nil, err
	}
	hexEncoded := hex.HexEncode(xorRes)

	return hexEncoded, nil
}
