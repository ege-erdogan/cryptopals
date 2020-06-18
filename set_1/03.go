// Set 1 - Basics
// Challenge 3
// Single-byte XOR cipher

package main

import (
	"encoding/hex"
	"fmt"
	"math"
	"regexp"
	"strings"
)

func decryptSingleByteXOR(ct string) string {
	bytes, _ := hex.DecodeString(ct)

	var minError float64 = math.MaxFloat64
	var minMessage string

	// for each possible key, calculate the error resulting from frequency analysis
	for i := 0; i < 26; i++ {
		key := byte('A' + i)
		message := applyKey(bytes, key)
		error := calculateFrequencyError(string(message))

		if error < minError {
			minError = error
			minMessage = string(message)
		}

		fmt.Printf("[KEY=%s] %s\t%f\n", string(key), message, error)
	}

	return minMessage
}

// mean absolute error
func calculateFrequencyError(msg string) float64 {
	englishFreqs := map[string]float64{
		"A": 8.49,
		"B": 1.49,
		"C": 2.2,
		"D": 4.25,
		"E": 11.16,
		"F": 2.22,
		"G": 2.01,
		"H": 6.09,
		"I": 7.54,
		"J": 0.15,
		"K": 1.29,
		"L": 4.02,
		"M": 2.4,
		"N": 6.75,
		"O": 7.5,
		"P": 1.93,
		"Q": 0.09,
		"R": 7.59,
		"S": 6.33,
		"T": 9.36,
		"U": 2.76,
		"V": 0.97,
		"W": 2.56,
		"X": 0.15,
		"Y": 1.99,
		"Z": 0.07,
	}
	freqs, length := getLetterFreqs(msg)

	error := 0.0
	for char, freq := range freqs {
		error += math.Abs(freq - englishFreqs[char])
	}

	return error / float64(length)
}

func getLetterFreqs(msg string) (map[string]float64, int) {
	freqs := make(map[string]float64)
	reg, _ := regexp.Compile("[^A-Z]+")

	// make string uppercase and remove non-alphanumeric characters
	modified := reg.ReplaceAllString(strings.ToUpper(msg), "")

	for _, char := range modified {
		freqs[string(char)]++
	}

	for letter, count := range freqs {
		freq := (count / float64(len(msg))) * 100
		freqs[letter] = freq
	}

	for i := 0; i < 26; i++ {
		char := string('A' + i)
		if _, exists := freqs[char]; !exists {
			freqs[char] = 0.0
		}
	}

	return freqs, len(modified)
}

func applyKey(msg []byte, key byte) []byte {
	ciphertext := make([]byte, len(msg))
	for i, val := range msg {
		ciphertext[i] = val ^ key
	}
	return ciphertext
}

func main() {
	fmt.Println("Plaintext: " + decryptSingleByteXOR("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"))
}
