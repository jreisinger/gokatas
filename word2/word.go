// Package word provides utilities for word games. It also shows how to write
// tests and examples.
//
// Adapted from github.com/adonovan/gopl.io/tree/master/ch11/word2
//
// Level: intermediate
// Topics: unicode, testing, rand
package word

import (
	"unicode"
)

// IsPalindrome reports whether s reads the same forward and backward.
// Letter case is ingored, as are non-letters.
func IsPalindrome(s string) bool {
	var letters []rune
	for _, r := range s {
		if unicode.IsLetter(r) {
			letters = append(letters, unicode.ToLower(r))
		}
	}
	for i := range letters {
		if letters[i] != letters[len(letters)-1-i] {
			return false
		}
	}
	return true
}
