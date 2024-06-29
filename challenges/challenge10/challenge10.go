package challenge10

import (
	"errors"
	"fmt"
	"io/ioutil"

	"github.com/danieldavies99/CryptopalsGoV2/utilities/base64"
	"github.com/danieldavies99/CryptopalsGoV2/utilities/cbc"
)

var ErrReadingFile = errors.New("failed to read file")

func Challenge10(filepath string, key []byte, iv []byte) ([]byte, error) {
	fileContent, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrReadingFile, err)
	}
	decoded, err := base64.Base64Decode(fileContent)
	if err != nil {
		return nil, err
	}
	return cbc.CBCDecrypt(decoded, key, iv)
}
