// Set 2 - Block crypto
// Challenge 13
// ECB cut-and-paste

package main

import (
	"errors"
	"fmt"
	"strings"
)

// 16-byte key variable is declared in 12.go

func keyValueParse(input string) map[string]string {
	result := make(map[string]string)
	tokens := strings.Split(input, "&")
	for _, token := range tokens {
		pair := strings.Split(token, "=")
		result[pair[0]] = pair[1]
	}
	return result
}

func profileFor(email string) ([]byte, error) {
	if strings.ContainsAny(email, "&=") {
		return nil, errors.New("characters & and = not allowed")
	}

	result := fmt.Sprintf("email=%s&uid=10&role=user", email)

	return encryptAESECB128([]byte(result), key), nil
}

func decryptAndParse(profile []byte) map[string]string {
	decrypted := decryptAESECB128(profile, key)
	return keyValueParse(string(decrypted))
}

// this is what the attacker does
func generateAdminProfile() []byte {
	blockSize := 16

	input := []byte{24, 24, 24, 24, 24, 24, 24, 24, 24, 24, 24}
	input = append([]byte("admin"), input...)
	input = append([]byte("eg@bar.com"), input...)

	temp, _ := profileFor(string(input))
	adminBlock := temp[blockSize : 2*blockSize]

	ciphertext, _ := profileFor("hello@bar.com")

	result := append(ciphertext[0:2*blockSize], adminBlock...)

	return result
}
