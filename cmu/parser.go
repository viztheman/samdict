package cmu

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func adjustPitch(r rune, adj int) (newPitch rune, err error) {
	pitch, err := strconv.Atoi(string(r))
	if err != nil {
		return 0, err
	}

	pitch += adj
	switch {
	case pitch < 0:
		newPitch = '0'
	case pitch > 9:
		newPitch = '9'
	default:
		newPitch = rune('0' + pitch)
	}
	return
}

func parseSyllables(syllable string, pitchAdj int) string {
	var sb strings.Builder
	for _, r := range []rune(syllable) {
		switch {
		case unicode.IsLetter(r):
			sb.WriteRune(unicode.ToLower(r))
		case unicode.IsNumber(r):
			newPitch, _ := adjustPitch(r, pitchAdj)
			sb.WriteRune(newPitch)
		}
	}
	result := strings.Replace(sb.String(), "hh", "/h", -1)
	return strings.Replace(result, "jh", "j", -1)
}

func parseLine(line string, pitchAdj int) (word string, translation string, valid bool) {
	if line == "" {
		return "", "", false
	}
	if unicode.IsPunct([]rune(line)[0]) {
		return "", "", false
	}

	fields := strings.Fields(line)
	if len(fields) < 2 {
		return "", "", false
	}

	word = strings.ToLower(fields[0])
	syllables := strings.Join(fields[1:], "")
	translation = parseSyllables(syllables, pitchAdj)
	valid = true
	return
}

// Parses the given file into a Dictionary.
func Read(path string, pitchAdj int) (dict Dictionary, err error) {
	dict = make(Dictionary)

	file, err := os.Open(path)
	if err != nil {
		return dict, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		word, translation, valid := parseLine(scanner.Text(), pitchAdj)
		if valid {
			dict[word] = translation
		}
	}

	err = scanner.Err()
	return
}
