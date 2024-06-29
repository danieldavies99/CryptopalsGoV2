package challenge9

import "github.com/danieldavies99/CryptopalsGoV2/utilities/pkcs7"

// this file is a bit pointless because I'm just wrapping
// the pkcs7.Pad function, but I'm going to keep it to keep my
// structure consistent and easy to follow
func Challenge9(input []byte, padTo int) ([]byte, error) {
	return pkcs7.Pad(input, padTo)
}
