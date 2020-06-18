// Set 1 - Basics
// Challenge 5
// Implement repeating-key XOR

package main

func repeatingKeyXOR(message, key []byte) []byte {
	ciphertext := make([]byte, len(message))
	for pos, plainByte := range message {
		keyByte := key[pos%len(key)]
		encryptedByte := plainByte ^ keyByte
		ciphertext[pos] = encryptedByte
	}
	return ciphertext
}
