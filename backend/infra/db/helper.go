package db

import "strings"

// containsWord checks if a map contains a specific key.
func containsWord(words map[string]int32, word string) bool {
	_, ok := words[word]
	return ok
}

// containsIgnoreCase verifica se um texto contém outro ignorando maiúsculas e minúsculas
func containsIgnoreCase(text, substr string) bool {
	return strings.Contains(strings.ToLower(text), strings.ToLower(substr))
}
func containsIgnoreCaseSlice(text string, substr []string) bool {
	for _, sub := range substr {
		return containsIgnoreCase(text, sub)
	}
	return false
}
