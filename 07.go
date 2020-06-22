// Set 1 - Basics
// Challenge 7
// AES in ECB mode

package main

import (
	"crypto/aes"

	"./util"
)

func decryptAESECB128(ciphertext, key []byte) []byte {
	var plaintext []byte

	cipher, _ := aes.NewCipher(key)
	ciphertextBlocks := util.GetFirstNBlocks(ciphertext, 16, -1)

	for _, block := range ciphertextBlocks {
		plainBlock := make([]byte, 16)
		block = pkcsPadding(block, 16)
		cipher.Decrypt(plainBlock, block)
		for _, val := range plainBlock {
			plaintext = append(plaintext, val)
		}
	}

	return plaintext
}
