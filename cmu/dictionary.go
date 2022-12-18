// Provides a structure and support functions for
// creating a CMU-style phonetic dictionary.
package cmu

import (
	"strings"
	"unicode"
)

// A custom type of map[string]string which
// holds CMU translations.
type Dictionary map[string]string

func splitPunc(s string) (word string, punc string) {
	if s == "" {
		return "", ""
	}

	runes := []rune(s)
	puncRune := runes[len(runes)-1]
	if !unicode.IsPunct(puncRune) {
		return s, ""
	}

	punc = string(puncRune)
	word = string(runes[:len(runes)-1])
	return
}

func scrubWord(word string) string {
	var sb strings.Builder
	for _, r := range []rune(word) {
		if unicode.IsPunct(r) && r != '\'' {
			continue
		}
		sb.WriteRune(unicode.ToLower(r))
	}
	return sb.String()
}

// Upserts a given Dictionary into the owner.
func (this Dictionary) Merge(src Dictionary) {
	for key, value := range src {
		this[key] = value
	}
}

// Takes a human readable word and return the equivalent
// CMU representation (if found).
func (this Dictionary) Lookup(src string) (output string, ok bool) {
	word, punc := splitPunc(src)
	word = scrubWord(word)
	translation, ok := this[word]
	if !ok {
		return "", false
	}

	return translation + punc, true
}
