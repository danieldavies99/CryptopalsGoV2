package challenge8

import (
	"errors"
	"fmt"
	"io/ioutil"

	"github.com/danieldavies99/CryptopalsGoV2/utilities/hex"
)

var ErrReadingFile = errors.New("failed to read file")

func Challenge8(filepath string) (int, error) {
	// read challenge input
	fileContent, err := ioutil.ReadFile(filepath)
	if err != nil {
		return 0, fmt.Errorf("%w: %v", ErrReadingFile, err)
	}

	// split each line of input into byte slices
	lines := [][]byte{}
	newLine := []byte{}
	for _, character := range fileContent {
		if character == '\n' {
			lines = append(lines, newLine)
			newLine = []byte{}
		} else {
			newLine = append(newLine, character)
		}
	}

	// hex decode each line
	hexDecodedLines := [][]byte{}
	for _, line := range lines {
		decodedLine, err := hex.HexDecode([]byte(string(line)))
		if err != nil {
			return 0, err
		}
		hexDecodedLines = append(hexDecodedLines, decodedLine)
	}

	highestDupeCount := 0
	highestDupeCountIndex := 0

	for i, line := range hexDecodedLines {
		// split line into blocks of 16 bytes
		blocks := splitBytesIntoBlocks(line, 16)

		// count duplicates
		count := countDuplicates(blocks)

		if count > highestDupeCount {
			highestDupeCount = count
			highestDupeCountIndex = i
		}
	}

	return highestDupeCountIndex + 1, nil
}

func countDuplicates(input [][]byte) int {
	count := 0
	foundBlocks := map[string]bool{}

	for i, block := range input {
		if foundBlocks[string(block)] {
			continue // already counted dupes
		}
		foundBlocks[string(block)] = true

		for j, otherBlock := range input {
			if i == j {
				continue
			}
			if isSame(block, otherBlock) {
				count++
			}
		}
	}
	return count
}

func isSame(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func splitBytesIntoBlocks(input []byte, blockLength int) [][]byte {

	blocks := [][]byte{}
	for i := 0; i < len(input); i += blockLength {
		blocks = append(blocks, input[i:i+blockLength])
	}

	return blocks
}
