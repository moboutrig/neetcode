package main

import (
	"fmt"
)

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sCharCount := make(map[rune]int, len(s))
	tCharCount := make(map[rune]int, len(t))

	for i := 0; i < len(s); i++ {
		sCharCount[rune(s[i])]++
		tCharCount[rune(t[i])]++
	}

	return mapsAreEqual(sCharCount, tCharCount)
}

func mapsAreEqual(map1, map2 map[rune]int) bool {
	for k, v := range map1 {
		if v2, ok := map2[k]; !ok || v != v2 {
			return false
		}
	}
	return true
}

func main() {
	var s, t string

	// Taking input for the strings
	fmt.Println("Enter the first string:")
	fmt.Scanln(&s)
	fmt.Println("Enter the second string:")
	fmt.Scanln(&t)

	// Check if the strings are anagrams
	result := isAnagram(s, t)
	if result {
		fmt.Printf("%q and %q are anagrams.\n", s, t)
	} else {
		fmt.Printf("%q and %q are not anagrams.\n", s, t)
	}
}
