package challenge1

import (
	"github.com/danieldavies99/CryptopalsGoV2/utilities/base64"
	"github.com/danieldavies99/CryptopalsGoV2/utilities/hex"
)

func Challenge1(input []byte) ([]byte, error) {
	// convert input from hex to bytes
	hexDecoded, err := hex.HexDecode(input)
	if err != nil {
		return nil, err
	}

	// encode the bytes with b64
	return base64.Base64Encode(hexDecoded), nil
}
