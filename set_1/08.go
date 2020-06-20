package main

import (
	"bytes"
	"fmt"
	"math"
)

// Set 1 - Basics
// Challenge 8
// Detect AES in ECB mode

// ECB mode encrypts identical plaintext blocks to identical ciphertext blocks
// the ciphertext with the least amount of unique blocks is likely to be encrypted with ECB
// this heuristic holds in this case, as there is only one ciphertext with repeating blocks (line 133 of input file, 4 identical blocks)
func detectAESECB(ciphertexts [][]byte) []byte {
	maxScore := 0.0
	blockSize := 16
	var maxCiphertext []byte

	for index, ciphertext := range ciphertexts {
		numBlocks := int(math.Ceil(float64(len(ciphertext)) / float64(blockSize)))
		blocks := getFirstNBlocks(ciphertext, blockSize, numBlocks)
		numUnique := uniqueCount(blocks)
		score := float64(numBlocks) / float64(numUnique)

		if score > maxScore {
			maxScore = score
			maxCiphertext = ciphertext
			fmt.Println(index)
		}
	}

	return maxCiphertext
}

func uniqueCount(list [][]byte) int {
	var uniques [][]byte
	for _, val := range list {
		if !contains(uniques, val) {
			uniques = append(uniques, val)
		}
	}
	return len(uniques)
}

func contains(list [][]byte, target []byte) bool {
	for _, val := range list {
		if bytes.Equal(val, target) {
			return true
		}
	}
	return false
}
