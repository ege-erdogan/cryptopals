// Set 1 - Basics
// Challenge 2
// Fixed XOR

package main

import (
	"encoding/hex"
)

func fixedXOR(s1 string) string {
	const s2 = "686974207468652062756c6c277320657965"

	// decode from hex to byte array
	d1, _ := hex.DecodeString(s1)
	d2, _ := hex.DecodeString(s2)

	// can XOR individual bytes, not slices
	result := make([]byte, len(d1))
	for i := 0; i < len(d1); i++ {
		result[i] = d1[i] ^ d2[i]
	}

	// encode back to string, may also return byte array instead
	return hex.EncodeToString(result)
}
