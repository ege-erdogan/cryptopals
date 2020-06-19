// Set 1 - Basics
// Challenge 7
// AES in ECB mode

package main

import (
	"github.com/spacemonkeygo/openssl"
)

func decryptAESECB128(ciphertext, key []byte) []byte {
	cipher, _ := openssl.GetCipherByName("aes-128-ecb")

	ctx, _ := openssl.NewDecryptionCipherCtx(cipher, nil, key, nil)
	cipherBytes, _ := ctx.DecryptUpdate(ciphertext)
	finalBytes, _ := ctx.DecryptFinal()
	cipherBytes = append(cipherBytes, finalBytes...)
	return cipherBytes
}
