package challenge7

import (
	"errors"
	"fmt"
	"io/ioutil"

	"github.com/danieldavies99/CryptopalsGoV2/utilities/base64"
	"github.com/danieldavies99/CryptopalsGoV2/utilities/ecb"
	"github.com/danieldavies99/CryptopalsGoV2/utilities/pkcs7"
)

var ErrReadingFile = errors.New("failed to read file")

func Challenge7(filepath string, key []byte) ([]byte, error) {
	// read challenge input
	fileContent, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrReadingFile, err)
	}

	// decode base64
	base64Decoded, err := base64.Base64Decode(fileContent)
	if err != nil {
		return nil, err
	}

	// decode ecb
	decoded, err := ecb.ECBDecrypt(base64Decoded, key)
	if err != nil {
		return nil, err
	}

	// strip padding
	decoded = pkcs7.StripPadding(decoded)

	return decoded, nil
}
