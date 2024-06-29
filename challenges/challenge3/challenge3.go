package challenge3

import (
	"sort"

	"github.com/danieldavies99/CryptopalsGoV2/utilities/englishscore"
	"github.com/danieldavies99/CryptopalsGoV2/utilities/hex"
	"github.com/danieldavies99/CryptopalsGoV2/utilities/xor"
)

func Challenge3(input []byte) ([]byte, byte, error) {
	decoded, err := hex.HexDecode(input)
	if err != nil {
		return nil, 0, err
	}
	type CharScore struct {
		character byte
		score     float32
	}

	// store the score for each character
	characterScores := []CharScore{
		{'a', 0}, {'b', 0}, {'c', 0}, {'d', 0}, {'e', 0},
		{'f', 0}, {'g', 0}, {'h', 0}, {'i', 0}, {'j', 0},
		{'k', 0}, {'l', 0}, {'m', 0}, {'n', 0}, {'o', 0},
		{'p', 0}, {'q', 0}, {'r', 0}, {'s', 0}, {'t', 0},
		{'u', 0}, {'v', 0}, {'w', 0}, {'x', 0}, {'y', 0},
		{'z', 0}, {'A', 0}, {'B', 0}, {'C', 0}, {'D', 0},
		{'E', 0}, {'F', 0}, {'G', 0}, {'H', 0}, {'I', 0},
		{'J', 0}, {'K', 0}, {'L', 0}, {'M', 0}, {'N', 0},
		{'O', 0}, {'P', 0}, {'Q', 0}, {'R', 0}, {'S', 0},
		{'T', 0}, {'U', 0}, {'V', 0}, {'W', 0}, {'X', 0},
		{'Y', 0}, {'Z', 0}, {' ', 0}, {'0', 0}, {'1', 0},
		{'2', 0}, {'3', 0}, {'4', 0}, {'5', 0}, {'6', 0},
		{'7', 0}, {'8', 0}, {'9', 0}, {'.', 0}, {',', 0},
		{'!', 0}, {'?', 0}, {':', 0}, {';', 0}, {'-', 0},
	}

	for i, cs := range characterScores {
		resultText, err := xor.XORAgainstOneByte(decoded, cs.character)
		if err != nil {
			return nil, 0, err
		}
		characterScores[i].score = englishscore.ScoreText(resultText)
	}

	// sort by score
	sort.Slice(characterScores, func(i, j int) bool {
		return characterScores[i].score > characterScores[j].score
	})

	// we could decode the top n scoring texts, but
	// for this solution method I'm just going to decode the
	// best scoring text
	xorDecoded, err := xor.XORAgainstOneByte(decoded, characterScores[0].character)
	if err != nil {
		return nil, byte(0), err
	}
	return xorDecoded, characterScores[0].character, nil
}
