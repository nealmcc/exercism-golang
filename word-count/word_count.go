package wordcount

import (
	"bufio"
	"log"
	"strings"
)

type Frequency map[string]int

// WordCount returns the number of times each word appears in the phrase.
// treats the input as ansi, since all the test cases are, and the problem
// doesn't specify otherwise.
func WordCount(phrase string) Frequency {
	count := make(Frequency, 8)

	s := bufio.NewScanner(strings.NewReader(phrase))
	s.Split(atNextWord)

	for s.Scan() {
		// avoid allocating a string more than once per word
		bytes := s.Bytes()
		word := string(lowercase(bytes))
		count[word]++
	}

	if err := s.Err(); err != nil {
		log.Fatal(err)
	}

	return count
}

// atNextWord is a bufio.SplitFunc used by bufio.Scanner to tokenize the
// phrase into words as as specified in the problem.
// Adapted from bufio.ScanWords
func atNextWord(data []byte, atEOF bool) (advance int, token []byte, err error) {
	var (
		start      int
		haveLetter bool
		haveDigit  bool
	)

	// Skip leading non-word characters
	for start < len(data) {
		b := data[start]
		if isLetter(b) {
			haveLetter = true
			break
		}
		if isDigit(b) {
			haveDigit = true
			break
		}
		start++
	}

	allowApos := true

	// Scan until we reach a non-word character, marking end of word.
	for end := start; end < len(data); end++ {
		b := data[end]

		// if we currently have letters allow more letters
		if haveLetter && isLetter(b) {
			continue
		}

		// if we currently have digits, allow more digits
		if haveDigit && isDigit(b) {
			continue
		}

		// possibly allow an apostrophe
		if allowApos && isContraction(data, end) {
			allowApos = false
			// possibly swap between letters and digits
			haveLetter = isLetter(data[end+1])
			haveDigit = !haveLetter
			continue
		}

		// we've reached a non-word character
		return end, data[start:end], nil
	}

	// If we're at EOF, we have a final, non-empty, non-terminated word. Return it.
	if atEOF && len(data) > start {
		return len(data), data[start:], nil
	}

	// Request more data.
	return start, nil, nil
}

func isLetter(b byte) bool {
	if 'a' <= b && b <= 'z' {
		return true
	}

	if 'A' <= b && b <= 'Z' {
		return true
	}

	return false
}

func isDigit(b byte) bool {
	return '0' <= b && b <= '9'
}

func isContraction(data []byte, i int) bool {
	if i <= 0 {
		return false
	}

	if i+1 >= len(data) {
		return false
	}

	if data[i] != '\'' {
		return false
	}

	leftIsWord := isLetter(data[i-1]) || isDigit(data[i-1])
	rightIsWord := isLetter(data[i+1]) || isDigit(data[i+1])

	return leftIsWord && rightIsWord
}

func lowercase(bytes []byte) []byte {
	for i, b := range bytes {
		if 'A' <= b && b <= 'Z' {
			b += 'a' - 'A'
			bytes[i] = b
		}
	}
	return bytes
}
