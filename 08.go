package main

import (
	"math"

	"./util"
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

	for _, ciphertext := range ciphertexts {
		numBlocks := int(math.Ceil(float64(len(ciphertext)) / float64(blockSize)))
		blocks := util.GetFirstNBlocks(ciphertext, blockSize, numBlocks)
		numUnique := util.UniqueCount(blocks)
		score := float64(numBlocks) / float64(numUnique)

		if score > maxScore {
			maxScore = score
			maxCiphertext = ciphertext
		}
	}

	return maxCiphertext
}
