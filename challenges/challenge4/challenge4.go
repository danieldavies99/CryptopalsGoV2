package challenge4

import (
	"io/ioutil"

	"github.com/danieldavies99/CryptopalsGoV2/challenges/challenge3"
	"github.com/danieldavies99/CryptopalsGoV2/utilities/englishscore"
)

func Challenge4(filepath string) ([]byte, error) {
	fileContent, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	hexEncodedLines := splitFileLines(fileContent)

	// get best scoring text for each line
	bestTexts := [][]byte{}
	for _, line := range hexEncodedLines {
		bestText, _, err := challenge3.Challenge3(line)
		if err != nil {
			return nil, err
		}
		bestTexts = append(bestTexts, bestText)
	}

	bestText := []byte{}
	var bestTextScore float32 = 0.0

	for i := 0; i < len(bestTexts); i++ {
		score := englishscore.ScoreText(bestTexts[i])
		if score > bestTextScore {
			bestText = bestTexts[i]
			bestTextScore = score
		}
	}

	return bestText, err
}

func splitFileLines(input []byte) [][]byte {
	output := [][]byte{}
	newLine := []byte{}
	for i := 0; i < len(input); i++ {
		if input[i] == '\n' {
			output = append(output, newLine)
			newLine = []byte{}
		} else {
			newLine = append(newLine, input[i])
		}
	}
	return output
}
