// Set 1 - Basics
// Challenge 3
// Single-byte XOR cipher

package main

import (
	"encoding/hex"
	"fmt"
	"math"
	"strings"
	"unicode/utf8"
)

var letterFrequencies = map[string]float64{
	"A": 0.0849,
	"B": 0.0149,
	"C": 0.022,
	"D": 0.0425,
	"E": 0.1116,
	"F": 0.0222,
	"G": 0.0201,
	"H": 0.0609,
	"I": 0.0754,
	"J": 0.0015,
	"K": 0.0129,
	"L": 0.0402,
	"M": 0.024,
	"N": 0.0675,
	"O": 0.075,
	"P": 0.0193,
	"Q": 0.0009,
	"R": 0.0759,
	"S": 0.0633,
	"T": 0.0936,
	"U": 0.0276,
	"V": 0.0097,
	"W": 0.0256,
	"X": 0.0015,
	"Y": 0.0199,
	"Z": 0.0007,
}

func decryptSingleByteXOR(ct string) (string, float64) {
	bytes, _ := hex.DecodeString(ct)

	minError := math.MaxFloat64
	var message string

	for i := 0; i < 26; i++ {
		key := byte('A' + i)
		possibleMessage := xorBytes(bytes, key)
		errorValue := getChi2(string(possibleMessage))

		fmt.Printf("%s\t%f\n", possibleMessage, errorValue)
		if errorValue < minError {
			minError = errorValue
			message = string(possibleMessage)
		}
	}

	return message, minError
}

func getChi2(msg string) float64 {
	msg = strings.ToUpper(msg)

	var counts [26]int
	for _, char := range msg {
		if 65 <= char && char <= 90 {
			counts[char%65]++
		}
	}

	errorValue := 0.0
	length := float64(utf8.RuneCountInString(msg))
	for i, observed := range counts {
		expected := length * letterFrequencies[string('A'+i)]
		errorValue += math.Pow(float64(observed)-expected, 2) / expected
	}

	return errorValue
}

func xorBytes(msg []byte, key byte) []byte {
	ciphertext := make([]byte, len(msg))
	for i, val := range msg {
		ciphertext[i] = val ^ key
	}
	return ciphertext
}
