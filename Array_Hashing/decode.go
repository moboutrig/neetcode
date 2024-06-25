package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Codec struct {
	b strings.Builder
}

// Encodes a list of strings to a single string.
func (codec *Codec) Encode(strs []string) string {
	defer codec.b.Reset()
	for _, word := range strs {
		codec.b.WriteString(strconv.Itoa(len(word)))
		codec.b.WriteRune('|')
		codec.b.WriteString(word)
	}
	return codec.b.String()
}

func (codec *Codec) Decode(strs string) []string {
    var words []string
    for len(strs) > 0 {
        // Find the index of the first "|"
        start := strings.Index(strs, "|")+1
        // Extract the number of characters for the next word
        length, _ := strconv.Atoi(strs[:start-1])
        // Calculate the end index of the next word
        end := start + length
        // Append the extracted word to the words slice
        words = append(words, strs[start:end])
        // Update strs to remove the processed part
        strs = strs[end:]
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
