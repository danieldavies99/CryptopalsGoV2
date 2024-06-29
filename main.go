package main

import (
	"io/ioutil"

	"github.com/danieldavies99/CryptopalsGoV2/utilities/base64"
	"github.com/danieldavies99/CryptopalsGoV2/utilities/cbc"
	"github.com/danieldavies99/CryptopalsGoV2/utilities/pkcs7"
)

func main() {
	fileContent, err := ioutil.ReadFile("./10.txt")
	if err != nil {
		panic(err)
	}
	cipherText, err := base64.Base64Decode(fileContent)
	if err != nil {
		panic(err)
	}
	decoded, err := cbc.CBCDecrypt(
		cipherText,
		[]byte("YELLOW SUBMARINE"),
		[]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	)
	if err != nil {
		panic(err)
	}

	decoded = pkcs7.StripPadding(decoded)


	println("Decoded: ", string(decoded))
}

// Sometimes I feel
//  like life is ju
// st a game fr fr.
