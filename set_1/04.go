// Set 1 - Basics
// Challenge 4
// Detect single-character XOR

package main

import (
	"fmt"
	"math"

	"../util"
)

const inputTextURL = "https://cryptopals.com/static/challenge-data/4.txt"

func detectXOR() []byte {
	lines := util.ReadLines("../inputs/04.txt")

	// for each line, find the plaintext with min. error
	// then find the min. error plaintext out of all the lines
	minError := math.MaxFloat64
	var message, minCiphertext []byte
	for _, ciphertext := range lines {
		possibleMessage, errorVal, _ := decryptSingleByteXOR([]byte(ciphertext))
		if errorVal < minError {
			minError = errorVal
			message = possibleMessage
			minCiphertext = ciphertext
		}
	}

	fmt.Println("Ciphertext: " + string(minCiphertext))
	fmt.Println("Plaintext: " + string(message))
	return minCiphertext
}
