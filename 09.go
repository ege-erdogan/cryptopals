// Set 2 - Block crypto
// Challenge 9
// Implement PKCS#7 padding

package main

// pads the input until the target length
func pkcsPadding(in []byte, targetLength int) []byte {
	paddingLength := targetLength - len(in)
	for i := 0; i < paddingLength; i++ {
		in = append(in, byte(paddingLength))
	}
	return in
}
