// Set 2 - Block crypto
// Challenge 10
// Implement CBC mode

package main

import (
	"crypto/aes"
	"encoding/base64"
	"errors"
	"fmt"
	"math"

	"./util"
)

// This is the same code as exercise 7, except it encrypts instead of decrypting

func encryptAESECB128(plaintext, key []byte) []byte {
	var ciphertext []byte

	cipher, _ := aes.NewCipher(key)
	plaintextBlocks := util.GetFirstNBlocks(plaintext, 16, -1)

	for _, block := range plaintextBlocks {
		ctBlock := make([]byte, 16)
		block = pkcsPadding(block, 16)
		cipher.Encrypt(ctBlock, block)
		for _, val := range ctBlock {
			ciphertext = append(ciphertext, val)
		}
	}

	return ciphertext
}

func encryptCBC(message, iv, key []byte) []byte {
	targetLength := len(message)
	for needPadding(targetLength) {
		targetLength++
	}

	message = pkcsPadding(message, targetLength)
	messageBlocks := util.GetFirstNBlocks(message, 16, -1)
	var ciphertext []byte

	xorInput := iv
	for _, block := range messageBlocks {
		aesInput, _ := xor(xorInput, block)
		ciphertextBlock := encryptAESECB128(aesInput, key)
		xorInput = ciphertextBlock
		for _, val := range ciphertextBlock[0:16] {
			ciphertext = append(ciphertext, val)
		}
	}

	return ciphertext
}

func decryptCBC(ciphertext, iv, key []byte) []byte {
	var plaintext []byte
	ciphertextBlocks := util.GetFirstNBlocks(ciphertext, 16, -1)

	xorInput := iv
	for _, block := range ciphertextBlocks {
		plaintextBlock, _ := xor(xorInput, decryptAESECB128(block, key))
		xorInput = block

		for _, val := range plaintextBlock {
			plaintext = append(plaintext, val)
		}
	}

	return plaintext
}

func needPadding(length int) bool {
	n := int(math.Floor((float64(length) / 16.0)))
	return n == 0 || n%2 != 0
}

func xor(b1, b2 []byte) ([]byte, error) {
	if len(b1) != len(b2) {
		return nil, errors.New("inputs must be of equal length")
	}

	result := make([]byte, len(b1))
	for i, val := range b1 {
		result[i] = val ^ b2[i]
	}
	return result, nil
}

func main() {
	key := []byte("YELLOW SUBMARINE")
	temp, _ := base64.StdEncoding.DecodeString(util.ReadFileToString("inputs/10.txt"))
	text := []byte(temp)
	iv := []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	result := decryptCBC(text, iv, key)
	fmt.Println(string(result))
}
