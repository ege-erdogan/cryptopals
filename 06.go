// Set 1 - Basics
// Challenge 6
// Break repeating-key XOR

package main

import (
	"fmt"
	"math"
	"math/bits"

	"./util"
)

func breakRepeatingKeyXOR(ciphertext []byte) []byte {
	keysize := findBestKeySize(ciphertext, 64, 8)

	numBlocks := int(math.Ceil(float64(len(ciphertext)) / float64(keysize)))
	blocks := util.GetFirstNBlocks(ciphertext, keysize, numBlocks)

	transposed := make([][]byte, keysize)
	for i := 0; i < keysize; i++ {
		var temp []byte
		for _, block := range blocks {
			if i < len(block) {
				temp = append(temp, block[i])
			}
		}
		transposed[i] = temp
	}

	key := make([]byte, keysize)
	for pos, block := range transposed {
		_, _, keyByte := decryptSingleByteXOR(block)
		key[pos] = keyByte
	}

	fmt.Println(keysize)
	return util.RepeatingKeyXOR(ciphertext, key)
}

func findBestKeySize(ciphertext []byte, maxKeysize, numBlocks int) int {
	minDistance := math.MaxFloat64
	var bestKeysize int

	for keysize := 2; keysize <= maxKeysize; keysize++ {
		blocks := util.GetFirstNBlocks(ciphertext, keysize, numBlocks)

		distance := 0.0
		for i, b1 := range blocks {
			for j, b2 := range blocks {
				if i != j {
					distance += float64(getHammingDistance(b1, b2)) / float64(keysize)
				}
			}
		}

		if distance < minDistance {
			minDistance = distance
			bestKeysize = keysize
		}
	}

	return bestKeysize
}

func getHammingDistance(b1, b2 []byte) int {
	distance := 0
	for pos, val := range b1 {
		distance += bits.OnesCount(uint(val ^ b2[pos]))
	}
	return distance
}
