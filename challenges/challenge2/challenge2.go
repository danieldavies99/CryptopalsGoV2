package challenge2

import (
	"github.com/danieldavies99/CryptopalsGoV2/utilities/hex"
	"github.com/danieldavies99/CryptopalsGoV2/utilities/xor"
)

func Challenge2(inputA []byte, inputB []byte) ([]byte, error) {
	hexDecodedA, err := hex.HexDecode(inputA)
	if err != nil {
		return nil, err
	}
	hexDecodedB, err := hex.HexDecode(inputB)
	if err != nil {
		return nil, err
	}
	xorRes, err := xor.XOR(hexDecodedA, hexDecodedB)
	if err != nil {
		return nil, err
	}
	return hex.HexEncode(xorRes), nil
}
