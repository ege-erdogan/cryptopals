// Set 1 - Basics
// Challenge 4
// Detect single-character XOR

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

const inputTextURL = "https://cryptopals.com/static/challenge-data/4.txt"

func detectXOR() string {
	lines := readLines()

	// for each line, find the plaintext with min. error
	// then find the min. error plaintext out of all the lines
	minError := math.MaxFloat64
	var message, minCiphertext string
	for _, ciphertext := range lines {
		possibleMessage, errorVal, _ := decryptSingleByteXOR([]byte(ciphertext))
		if errorVal < minError {
			minError = errorVal
			message = possibleMessage
			minCiphertext = ciphertext
		}
	}

	fmt.Println("Ciphertext: " + minCiphertext)
	fmt.Println("Plaintext: " + message)
	return minCiphertext
}

// read lines from file to an array of strings
func readLines() []string {
	file, _ := os.Open("04_input.txt")

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, strings.TrimSuffix(scanner.Text(), "\n"))
	}

	return lines
}
