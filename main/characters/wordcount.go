package characters

import "strings"

// CountWords returns a map containing words and their frequency
func CountWords(text string) map[string]int {
	counts := make(map[string]int)
	words := strings.Fields(text)

	for _, word := range words {
		// accessing a map on a non existent key gives 0
		counts[strings.ToLower(word)]++
	}

	return counts
}