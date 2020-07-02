// Set 2 - Block crypto
// Challenge 14
// byte-at-a-time ECB decryption (harder)

package main

import (
	"encoding/base64"
	"fmt"
	"math"
	"math/rand"
	"time"

	"./util"
)

var prefix []byte = generateRandomPrefix()

func generateRandomPrefix() []byte {
	rand.Seed(time.Now().UnixNano())
	length := rand.Intn(100)
	result := make([]byte, length)
	for i := range result {
		result[i] = byte(rand.Intn(256))
	}

	return result
}

// AES-128-ECB(random-prefix || text || fixed-message, random-key)
func harderRandomEncryptOracle(text []byte) []byte {
	end, _ := base64.StdEncoding.DecodeString(util.ReadFileToString("inputs/12.txt"))
	input := append(append(prefix, text...), []byte(end)...)

	return encryptAESECB128(input, key)
}

// functions below for the attacker

type cipher func([]byte) []byte

func findPrefixLength(encryptionMethod cipher) int {
	input1 := []byte{0}
	input2 := []byte{1}
	count := math.MaxInt32
	for i := 0; i < 17; i++ {
		ct1 := util.GetFirstNBlocks(encryptionMethod(input1), 16, -1)
		ct2 := util.GetFirstNBlocks(encryptionMethod(input2), 16, -1)
		input1 = append(input1, byte(0))
		input2 = append(input2, byte(1))
		changed := blockChangeCount(ct1, ct2)
		if changed > count {
			temp := util.GetFirstNBlocks(harderRandomEncryptOracle(append(input1, byte(1))), 16, -1)
			for pos, block := range temp {
				if string(block) != string(ct1[pos]) {
					return pos*16 - i
				}
			}
		}
		count = changed
	}

	// never reaches here
	return -1
}

func blockChangeCount(b1, b2 [][]byte) int {
	count := 0
	for i, val := range b1 {
		if string(val) != string(b2[i]) {
			count++
		}
	}
	return count
}

// break msg

func harderBreakByteAtATime(blockSize, decodeIndex int, knownMessage []byte) byte {
	prefixLength := findPrefixLength(harderRandomEncryptOracle)
	prefixBlocks := int(math.Ceil(float64(prefixLength) / 16.0))
	prefixInputLength := 16 - (prefixLength % 16)
	inputLength := prefixInputLength + blockSize - (decodeIndex % blockSize) - 1

	input := make([]byte, inputLength)
	blockIndex := decodeIndex / blockSize
	start := (prefixBlocks + blockIndex) * blockSize
	ctBlock := harderRandomEncryptOracle(input)[start : start+blockSize]

	var prefix []byte // first (blocksize - 1) characters of the block we search
	if blockIndex == 0 {
		prefix = append(input[prefixInputLength:], knownMessage...)
	} else {
		// last (blockSize - 1) characters of the known message
		prefix = knownMessage[len(knownMessage)-blockSize+1:]
	}
	nextByte := harderMatchOneByte(ctBlock, prefix, prefixBlocks, blockSize, prefixInputLength)
	return nextByte
}

func harderMatchOneByte(targetBlock, prefix []byte, blocksRemoved, blockSize, prefixInputLength int) byte {
	for i := 1; i < 256; i++ { // trying every possible byte
		temp := make([]byte, prefixInputLength)
		message := append(append(temp, prefix...), byte(i))
		ct := harderRandomEncryptOracle(message)[blocksRemoved*blockSize:]
		candidateBlock := ct[0:blockSize]
		if string(candidateBlock) == string(targetBlock) {
			return byte(i)
		}
	}
	return byte(0)
}

// call this as main function to run
func solveChallenge14() {
	blockSize := 16
	var message []byte

	for i := 0; i < 137; i++ {
		nextByte := harderBreakByteAtATime(blockSize, i, message)
		message = append(message, nextByte)
	}
	fmt.Println(string(message))
}
