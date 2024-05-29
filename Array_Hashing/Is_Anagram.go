package main

import (
	"fmt"
)

// isAnagram function to check if two strings are anagrams
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	s_track, t_track := make(map[rune]int), make(map[rune]int)
	populateMaps(s, t, s_track, t_track, 0)
	return mapsAreEqual(s_track, t_track)
}

// populateMaps function to populate the character count maps recursively for both strings
func populateMaps(s, t string, s_track, t_track map[rune]int, index int) {
	if index == len(s) {
		return
	}
	s_track[rune(s[index])]++
	t_track[rune(t[index])]++
	populateMaps(s, t, s_track, t_track, index+1)
}

func mapsAreEqual(m1, m2 map[rune]int) bool {
	for k, v := range m1 {
		if v2, ok := m2[k]; !ok || v != v2 {
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
