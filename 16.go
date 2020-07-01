// Set 2 - Block crypto
// Challenge 16
// CBC bit flipping attacks

package main

import (
	"fmt"
	"strings"
)

// 'key' variable is the random AES key

func modifyAndEncrypt(text []byte) []byte {
	prefix := []byte("comment1=cooking%20MCs;userdata=")
	suffix := []byte(";comment2=%20like%20a%20pound%20of%20bacon")
	text = []byte(strings.ReplaceAll(string(text), ";", ""))
	text = []byte(strings.ReplaceAll(string(text), "=", ""))
	text = append(append(prefix, text...), suffix...)

	for len(text)%16 != 0 {
		text = pkcsPadding(text, len(text)+1)
	}

	iv := make([]byte, 16)
	return encryptCBC(text, iv, key)
}

func decryptAndFindAdmin(ct []byte) bool {
	iv := make([]byte, 16)
	text := decryptCBC(ct, iv, key)
	tokens := strings.Split(string(text), ";")
	for _, token := range tokens {
		if token == "admin=true" {
			return true
		}
	}

	return false
}

func next(in []byte) []byte {
	for i, val := range in {
		if val == 255 {
			in[i] = 0
		} else {
			in[i]++
			return in
		}
	}
	return nil
}

// call as main function
func solveChallenge16() {
	prefixLength := findPrefixLength(modifyAndEncrypt)
	s := []byte(";admin=true;")
	inputLength := 16 - (prefixLength % 16) + len(s)
	input := make([]byte, inputLength)
	ct := modifyAndEncrypt(input)
	p := input[len(input)-len(s):]

	start := prefixLength + inputLength - len(s)
	end := start + len(s)
	temp, _ := xor(p, ct[start:end])
	block, _ := xor(temp, s)

	ct = append(ct[0:start], append(block, ct[end:]...)...)
	fmt.Println(decryptAndFindAdmin(ct)) // prints true
}
