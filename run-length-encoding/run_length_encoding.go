package encode

import (
	"bufio"
	"log"
	"strconv"
	"strings"
)

// RunLengthEncode compresses an input string as described in the exercse. See:
// https://en.wikipedia.org/wiki/Run-length_encoding
func RunLengthEncode(plain string) string {
	var (
		s *bufio.Scanner
		b strings.Builder
	)

	s = bufio.NewScanner(strings.NewReader(plain))
	s.Split(onNextSymbol)

	for s.Scan() {
		if err := s.Err(); err != nil {
			log.Fatal(err)
		}
		bytes := s.Bytes()
		if len(bytes) > 1 {
			b.WriteString(strconv.Itoa(len(bytes)))
		}
		b.WriteByte(bytes[0])
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}

	return b.String()
}

// onNextSymbol is a bufio.Split function which splits the data into tokens
// where each token is a sequence of identical bytes
func onNextSymbol(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}

	var (
		symbol  byte = data[0]
		i       int  = 1
		isMatch bool
	)

	for ; i < len(data); i++ {
		next := data[i]
		isMatch = next == symbol
		if !isMatch {
			break
		}
	}

	if isMatch {
		if atEOF {
			// return the final token
			return i, data, nil
		}
		// get more data and try again
		return 0, nil, nil
	}

	return i, data[:i], nil
}

// RunLengthDecode extracts a string that was compressed with RunLengthEncode
func RunLengthDecode(comp string) string {
	return comp
}
