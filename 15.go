// Set 2 - Block crypto
// Challenge 15
// PKCS#7 padding validation

package main

import (
	"errors"
)

func removePkcs7Padding(text []byte) ([]byte, error) {
	paddingLength := int(text[len(text)-1])
	for i := 0; i < paddingLength; i++ {
		if text[len(text)-i-1] != byte(paddingLength) {
			return nil, errors.New("invalid padding")
		}
	}
	return text[:len(text)-paddingLength], nil
}
