package challenge6

import (
	"io/ioutil"
	"sort"

	"github.com/danieldavies99/CryptopalsGoV2/challenges/challenge3"
	"github.com/danieldavies99/CryptopalsGoV2/utilities/base64"
	"github.com/danieldavies99/CryptopalsGoV2/utilities/editdistance"
	"github.com/danieldavies99/CryptopalsGoV2/utilities/hex"
	"github.com/danieldavies99/CryptopalsGoV2/utilities/xor"
)

type keysizeEditDistance struct {
	keysize int
	ed      float32
}

func Challenge6(filepath string) ([]byte, []byte, error) {
	// read challenge input
	fileContent, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, nil, err
	}

	// the challenge input is base64 encoded,
	// so decode it before processing
	b64Decoded, err := base64.Base64Decode(fileContent)
	if err != nil {
		return nil, nil, err
	}

	// Let KEYSIZE be the guessed length of the key; try values from 2 to (say) 40.
	// store the score for each character
	keysizeEditDistances := []keysizeEditDistance{
		{2, 0.0}, {3, 0.0}, {4, 0.0}, {5, 0.0},
		{6, 0.0}, {7, 0.0}, {8, 0.0}, {9, 0.0},
		{10, 0.0}, {11, 0.0}, {12, 0.0}, {13, 0.0},
		{14, 0.0}, {15, 0.0}, {16, 0.0}, {17, 0.0},
		{18, 0.0}, {19, 0.0}, {20, 0.0}, {21, 0.0},
		{22, 0.0}, {23, 0.0}, {24, 0.0}, {25, 0.0},
		{26, 0.0}, {27, 0.0}, {28, 0.0}, {29, 0.0},
		{30, 0.0}, {31, 0.0}, {32, 0.0}, {33, 0.0},
		{34, 0.0}, {35, 0.0}, {36, 0.0}, {37, 0.0},
		{38, 0.0}, {39, 0.0}, {40, 0.0},
	}

	orderedKeySizeEditDistances, err := getBestKeysizes(b64Decoded, keysizeEditDistances)
	if err != nil {
		return nil, nil, err
	}
	
	// The KEYSIZE with the smallest normalized edit distance is probably the key.
	bestKeysize := orderedKeySizeEditDistances[0].keysize

	// Now that you probably know the KEYSIZE: break the ciphertext into blocks of KEYSIZE length.
	blocks := splitCypherTextIntoBlocks(b64Decoded, bestKeysize)

	// Now transpose the blocks: make a block that is the first byte of every block,
	// and a block that is the second byte of every block, and so on.
	transposedBlocks := transposeBlocks(blocks, bestKeysize)

	key, err := findKey(transposedBlocks)
	if err != nil {
		return nil, nil, err
	}

	// Now that you know the key, you can decipher the entire message.
	decodedRes, err := xor.XORRepeatingKey(b64Decoded, key)
	if err != nil {
		return nil, nil, err
	}

	return key, decodedRes, nil
}

func getNormalizedEditDistanceForKeySize(input []byte, keysize int, nblocks int) (float32, error) {
	editDistances := []float32{}

	// split the input into blocks of keysize length
	blocks := [][]byte{}
	for i := 0; i < len(input); i += keysize {
		blocks = append(blocks, input[i:i+keysize])
	}

	// get the edit distance between each block
	for i := 0; i < len(blocks); i += nblocks {

		compBlock := blocks[i]
		for j := i; j < len(blocks); j++ {
			ed, err := editdistance.EditDistance(compBlock, blocks[j])
			if err != nil {
				return 0.0, err
			}
			editDistances = append(editDistances, float32(ed))
		}

	}

	// average the edit distances
	var sum float32
	for _, ed := range editDistances {
		sum += ed
	}
	return (sum / float32(len(editDistances))) / float32(keysize), nil
}

func splitCypherTextIntoBlocks(input []byte, keysize int) [][]byte {
	blocks := [][]byte{}
	for i := 0; i < len(input); i += keysize {
		blocks = append(blocks, input[i:i+keysize])
	}

	return blocks
}

func transposeBlocks(blocks [][]byte, keysize int) [][]byte {
	transposedBlocks := [][]byte{}
	for i := 0; i < keysize; i++ {
		newBlock := []byte{}
		for j := 0; j < len(blocks); j++ {
			newBlock = append(newBlock, blocks[j][i])
		}
		transposedBlocks = append(transposedBlocks, newBlock)
	}
	return transposedBlocks
}

func findKey(transposedBlocks [][]byte) ([]byte, error) {
	key := []byte{}
	for _, block := range transposedBlocks {
		// hacky but challenge 3 method expects input to be hex encoded
		hexBlock := hex.HexEncode(block)
		_, keybyte, err := challenge3.Challenge3(hexBlock)
		if err != nil {
			return nil, err
		}
		key = append(key, keybyte)
	}
	return key, nil
}

func getBestKeysizes(input []byte, keysizeEditDistances []keysizeEditDistance) ([]keysizeEditDistance, error) {
	for i, ksed := range keysizeEditDistances {
		ed, err := getNormalizedEditDistanceForKeySize(input, ksed.keysize, 2)
		if err != nil {
			return nil, err
		}
		keysizeEditDistances[i].ed = ed
	}

	// sort by edit distance
	sort.Slice(keysizeEditDistances, func(i, j int) bool {
		return keysizeEditDistances[i].ed < keysizeEditDistances[j].ed
	})

	return keysizeEditDistances, nil
}
