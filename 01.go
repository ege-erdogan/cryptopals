// Set 1 - Basics
// Challenge 1
// Convert hex to base64

package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// table[i] = integer i's encoding in base64
var table = [64]string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "+", "/"}

func hexToBase64(in []byte) string {
	result := ""
	bits := byteArrayToBits(in)

	// #bits should be multiple of 6
	for len(bits)%6 != 0 {
		bits = append(bits, 0)
	}

	numBlocks := len(bits) / 6

	// for each 6-bit block, convert to decimal and find that value's base64 encoding
	for i := 0; i < numBlocks; i++ {
		blockStart := i * 6
		block := bits[blockStart : blockStart+6]
		decimal := binaryToDecimal(block)
		base64 := table[decimal]
		result += base64
	}

	// padding
	for len(result)%4 != 0 {
		result += "="
	}

	return result
}

// converts binary number to decimal
// first bit in array is the MSB
func binaryToDecimal(binary []int) int {
	result := 0
	for i, val := range binary {
		power := float64(len(binary) - i - 1)
		result += int(val) * int(math.Pow(2, power))
	}
	return result
}

// convert an array of bytes into one contiguous array of 1's and 0's
func byteArrayToBits(bytes []byte) []int {
	result := make([]int, len(bytes)*8)
	for i, val := range bytes {
		bits := byteToBits(val)
		for j := 0; j < 8; j++ {
			result[i*8+j] = bits[j]
		}
	}
	return result
}

// convert byte to array of 1's and 0's
// result[0] -> MSB
func byteToBits(in byte) [8]int {
	arr := strings.Split(fmt.Sprintf("%08b", in), "")
	var intArr [8]int
	for i, el := range arr {
		intArr[i], _ = strconv.Atoi(el)
	}
	return intArr
}
