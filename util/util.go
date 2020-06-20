package util

import (
	"bufio"
	"bytes"
	"os"
	"strings"
)

// Collection of utility methods not related to specific problems for reading files etc.s

// ReadLines read lines lines from file as byte slices
func ReadLines(path string) [][]byte {
	file, _ := os.Open(path)

	var lines [][]byte
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, []byte(strings.TrimSuffix(scanner.Text(), "\n")))
	}

	return lines
}

// ReadFileToString read file contents as a single string
func ReadFileToString(path string) string {
	file, _ := os.Open(path)
	var result string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		result += scanner.Text()
	}
	return result
}

// UniqueCount number of unique byte slices
func UniqueCount(list [][]byte) int {
	var uniques [][]byte
	for _, val := range list {
		if !Contains(uniques, val) {
			uniques = append(uniques, val)
		}
	}
	return len(uniques)
}

// Contains true if byte slice contained in list
func Contains(list [][]byte, target []byte) bool {
	for _, val := range list {
		if bytes.Equal(val, target) {
			return true
		}
	}
	return false
}

// GetFirstNBlocks first `count` blocks of given size from the slice
func GetFirstNBlocks(ct []byte, size, count int) [][]byte {
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
