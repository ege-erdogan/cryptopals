// Set 2 - Block crypto
// Challenge 11
// ECB/CBC Detection Oracle

package main

import (
	"math/rand"
	"time"

	"./util"
)

// should use "crypto/rand" but this is enough for our purposes
func generateRandomKey(length int) []byte {
	key := make([]byte, length)
	for i := 0; i < length; i++ {
		rand.Seed(time.Now().UnixNano())
		key[i] = byte(rand.Intn(256))
	}
	return key
}

func encryptWithRandomKey(text []byte) []byte {
	rand.Seed(time.Now().UnixNano())
	appendLength := rand.Intn(6) + 5
	for i := 0; i < appendLength; i++ {
		text = append([]byte{15}, text...)
		text = append(text, byte(15))
	}

	key := generateRandomKey(16)
	if rand.Intn(2) == 0 {
		iv := generateRandomKey(16)
		return encryptCBC(text, iv, key)
	} else {
		return encryptAESECB128(text, key)
	}
}

// we send an 256-byte message of 0's
// if ecb, there should be lots of identical blocks
// if cbc, probability of two ct blocks being equal is negligible
func detectEncryptionMode() string {
	message := make([]byte, 256)
	ciphertext := encryptWithRandomKey(message)
	blocks := util.GetFirstNBlocks(ciphertext, 16, -1)

	uniqueCount := util.UniqueCount(blocks)
	if uniqueCount <= len(blocks)/2 {
		return "ECB"
	} else {
		return "CBC"
	}
}
