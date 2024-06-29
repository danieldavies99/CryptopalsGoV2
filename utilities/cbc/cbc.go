package cbc

import (
	"errors"

	"github.com/danieldavies99/CryptopalsGoV2/utilities/ecb"
	"github.com/danieldavies99/CryptopalsGoV2/utilities/xor"
)

var ErrKeyLength = errors.New("key length must be 16 bytes")
var ErrIVLength = errors.New("iv length must be 16 bytes")
var ErrCypherTextLength = errors.New("cypher text length must be a multiple of 16 bytes")

// I wouldn't not have been able to write this without the info here
// https://en.wikipedia.org/wiki/Block_cipher_mode_of_operation
func CBCEncrypt(plainText []byte, key []byte, iv []byte) ([]byte, error) {
	if len(key) != 16 {
		return nil, ErrKeyLength
	}
	if len(iv) != 16 {
		return nil, ErrIVLength
	}

	// split plaintext into blocks and add iv to start
	plainTextBlocks := [][]byte{}
	// plainTextBlocks = append(plainTextBlocks, iv)
	for i := 0; i < len(plainText); i += 16 {
		plainTextBlocks = append(plainTextBlocks, plainText[i:i+16])
	}

	// iterate over blocks
	cipherBlocks := [][]byte{}
	cipherBlocks = append(cipherBlocks, iv)
	for i := 0; i < len(plainTextBlocks); i++ {
		previousBlock := cipherBlocks[i]
		currentBlock := plainTextBlocks[i]
		// xor previous block with current block
		xorBlock, err := xor.XOR(previousBlock, currentBlock)
		if err != nil {
			return nil, err
		}
		// encrypt xor block
		encryptedBlock, err := ecb.ECBEncrypt(xorBlock, key)
		if err != nil {
			return nil, err
		}
		cipherBlocks = append(cipherBlocks, encryptedBlock)
	}

	// join cipher blocks
	cipherText := []byte{}
	for i, block := range cipherBlocks {
		// skip iv
		if i == 0 {
			continue
		}
		cipherText = append(cipherText, block...)
	}
	return cipherText, nil
}

func CBCDecrypt(cypherText []byte, key []byte, iv []byte) ([]byte, error) {
	if len(key) != 16 {
		return nil, ErrKeyLength
	}
	if len(iv) != 16 {
		return nil, ErrIVLength
	}
	if len(cypherText)%16 != 0 {
		return nil, ErrCypherTextLength
	}

	// split cyphertext into blocks
	cypherTextBlocks := [][]byte{}
	cypherTextBlocks = append(cypherTextBlocks, iv)
	for i := 0; i < len(cypherText); i += 16 {
		cypherTextBlocks = append(cypherTextBlocks, cypherText[i:i+16])
	}

	// iterate over blocks
	plainTextBlocks := [][]byte{}
	for i := 1; i < len(cypherTextBlocks); i++ {
		previousBlock := cypherTextBlocks[i-1]
		currentBlock := cypherTextBlocks[i]
		decryptedBlock, err := ecb.ECBDecrypt(currentBlock, key)
		if err != nil {
			return nil, err
		}
		// xor previous block with current block
		xorBlock, err := xor.XOR(previousBlock, decryptedBlock)
		if err != nil {
			return nil, err
		}
		plainTextBlocks = append(plainTextBlocks, xorBlock)

	}

	// join cipher blocks
	plainText := []byte{}
	for _, block := range plainTextBlocks {
		plainText = append(plainText, block...)
	}
	return plainText, nil
}
