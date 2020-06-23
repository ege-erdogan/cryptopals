// Set 2 - Block crypto
// Challenge 12
// Byte-at-a-time ECB decryption (simple)

package main

import (
	"encoding/base64"
	"fmt"
	"math"

	"./util"
)

var key = generateRandomKey(16)

func encryptECBUnknownKey(text []byte) []byte {
	end, _ := base64.StdEncoding.DecodeString(util.ReadFileToString("inputs/12.txt"))
	return encryptAESECB128(append(text, []byte(end)...), key)
}

func discoverCipherBlockSize() int {
	input := []byte{1}
	lastLength := math.MaxInt32
	for {
		ciphertext := encryptECBUnknownKey(input)
		if len(ciphertext) > lastLength {
			return len(ciphertext) - lastLength
		}
		lastLength = len(ciphertext)
		input = append(input, byte(1))
	}
}

func breakByteAtATime(blockSize, decodeIndex int, knownMessage []byte) byte {
	inputLength := blockSize - (decodeIndex % blockSize) - 1
	input := make([]byte, inputLength)

	blockIndex := decodeIndex / blockSize
	start := blockIndex * blockSize
	ctBlock := encryptECBUnknownKey(input)[start : start+blockSize]

	var prefix []byte // first (blocksize - 1) characters of the block we search
	if blockIndex == 0 {
		prefix = append(input, knownMessage...)
	} else {
		// last (blockSize - 1) characters of the known message
		prefix = knownMessage[len(knownMessage)-blockSize+1:]
	}
	nextByte := matchOneByte(ctBlock, prefix, blockSize)
	return nextByte
}

func matchOneByte(targetBlock, prefix []byte, blockSize int) byte {
	for i := 0; i < 256; i++ { // trying every possible byte
		message := append(prefix, byte(i))
		candidateBlock := encryptECBUnknownKey(message)[0:blockSize]
		if string(candidateBlock) == string(targetBlock) {
			return byte(i)
		}
	}
	return byte(0)
}

// call this as main function to run
func solveChallenge12() {
	blockSize := discoverCipherBlockSize()
	messageLength := len(encryptECBUnknownKey([]byte{}))
	var message []byte

	for i := 0; i < messageLength; i++ {
		nextByte := breakByteAtATime(blockSize, i, message)
		message = append(message, nextByte)
	}
	fmt.Println(string(message))
}
