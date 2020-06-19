// Set 1 - Basics
// Challenge 6
// Break repeating-key XOR

package main

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"math"
	"math/bits"
	"os"
)

func breakRepeatingKeyXOR(ciphertext []byte) []byte {
	keysize := findBestKeySize(ciphertext, 64, 8)

	numBlocks := int(math.Ceil(float64(len(ciphertext)) / float64(keysize)))
	blocks := getFirstNBlocks(ciphertext, keysize, numBlocks)

	transposed := make([][]byte, keysize)
	for i := 0; i < keysize; i++ {
		var temp []byte
		for _, block := range blocks {
			if i < len(block) {
				temp = append(temp, block[i])
			}
		}
		transposed[i] = temp
	}

	key := make([]byte, keysize)
	for pos, block := range transposed {
		_, _, keyByte := decryptSingleByteXOR(block)
		key[pos] = keyByte
	}

	fmt.Println(keysize)
	return repeatingKeyXOR(ciphertext, key)
}

func findBestKeySize(ciphertext []byte, maxKeysize, numBlocks int) int {
	minDistance := math.MaxFloat64
	var bestKeysize int

	for keysize := 2; keysize <= maxKeysize; keysize++ {
		blocks := getFirstNBlocks(ciphertext, keysize, numBlocks)

		distance := 0.0
		for i, b1 := range blocks {
			for j, b2 := range blocks {
				if i != j {
					distance += float64(getHammingDistance(b1, b2)) / float64(keysize)
				}
			}
		}

		fmt.Printf("%d\t%f\n", keysize, distance)
		if distance < minDistance {
			minDistance = distance
			bestKeysize = keysize
		}
	}

	return bestKeysize
}

func getFirstNBlocks(ct []byte, size, count int) [][]byte {
	result := make([][]byte, count)
	for i := 0; i < count; i++ {
		var block []byte
		end := (i + 1) * size
		if end > len(ct) {
			block = ct[i*size:]
		} else {
			block = ct[i*size : end]
		}
		result[i] = block
	}
	return result
}

func getHammingDistance(b1, b2 []byte) int {
	distance := 0
	for pos, val := range b1 {
		distance += bits.OnesCount(uint(val ^ b2[pos]))
	}
	return distance
}

func readFileToString(path string) string {
	file, _ := os.Open(path)
	var result string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		result += scanner.Text()
	}
	return result
}

func main() {
	decoded := readFileToString("06_input.txt")
	ciphertext, _ := base64.StdEncoding.DecodeString(decoded)
	fmt.Println(string(breakRepeatingKeyXOR([]byte(ciphertext))))
}
