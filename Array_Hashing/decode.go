package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Codec struct {
	b strings.Builder // Builder to construct the encoded string
}

// Encodes a list of strings to a single string.
func (codec *Codec) Encode(strs []string) string {
	defer codec.b.Reset() // Ensure builder is reset after encoding
	for _, word := range strs {
		codec.b.WriteString(strconv.Itoa(len(word))) // Write length of the word
		codec.b.WriteRune('|')                        // Write delimiter
		codec.b.WriteString(word)                     // Write the word itself
	}
	return codec.b.String() // Return the complete encoded string
}

// Decodes a single string to a list of strings.
func (codec *Codec) Decode(strs string) []string {
	var words []string
	for len(strs) > 0 {
		start := strings.Index(strs, "|") + 1         // Find the position of the delimiter
		length, _ := strconv.Atoi(strs[:start-1])     // Extract and convert length
		words = append(words, strs[start:start+length]) // Extract the word and append to list
		strs = strs[start+length:]                    // Update the string to remove the processed part
	}
	return words
}
func main() {
	var codec Codec
	original := []string{"hello","hello","hello","hello","hello"}
	encoded := codec.Encode(original)
	fmt.Println("Encoded:", encoded)

	decoded := codec.Decode(encoded)
	fmt.Println("Decoded:", decoded)
}

// Your Codec object will be instantiated and called as such:
// var codec Codec
// codec.Decode(codec.Encode(strs));
