package encode

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"strconv"
	"strings"
)

// RunLengthEncode compresses an input string as described in the exercse.
func RunLengthEncode(plain string) string {
	r := strings.NewReader(plain)
	b := strings.Builder{}
	if err := EncodeStream(r, &b); err != nil {
		log.Fatal(err)
	}
	return b.String()
}

// RunLengthDecode extracts a string that was compressed with RunLengthEncode
func RunLengthDecode(compressed string) string {
	r := strings.NewReader(compressed)
	b := strings.Builder{}
	if err := DecodeStream(r, &b); err != nil {
		log.Fatal(err)
	}
	return b.String()
}

// EncodeStream reads from r, and writes the compressed text to w
func EncodeStream(r io.Reader, w io.Writer) error {
	s := bufio.NewScanner(r)
	s.Split(onNextSymbol)
	for s.Scan() {
		if err := s.Err(); err != nil {
			return err
		}
		bytes := s.Bytes()
		if len(bytes) > 1 {
			w.Write([]byte(strconv.Itoa(len(bytes))))
		}
		if _, err := w.Write(bytes[:1]); err != nil {
			return err
		}
	}
	return s.Err()
}

// DecodeStream reads from r, and writes the plain text to w
func DecodeStream(r io.Reader, w io.Writer) error {
	s := bufio.NewScanner(r)
	s.Split(onNonDigit)
	for s.Scan() {
		if err := s.Err(); err != nil {
			return err
		}
		token := s.Bytes()
		if len(token) == 1 {
			if _, err := w.Write(token); err != nil {
				return err
			}
			continue
		}
		char := token[len(token)-1:]
		count, err := strconv.Atoi(string(token[:len(token)-1]))
		if err != nil {
			return err
		}
		for i := 0; i < count; i++ {
			if _, err := w.Write(char); err != nil {
				return err
			}
		}
	}
	return s.Err()
}

// onNextSymbol is a bufio.Split function which splits the data into tokens
// where each token is a sequence of identical bytes
func onNextSymbol(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}

	var (
		symbol      byte = data[0]
		i           int  = 1
		isDifferent bool
	)

	for ; i < len(data); i++ {
		next := data[i]
		isDifferent = next != symbol
		if isDifferent {
			return i, data[:i], nil
		}
	}

	if atEOF {
		return i, data, nil
	}

	return 0, nil, nil
}

// onNonDigit is a bufio.Split function which splits the data into tokens
// where each token consumes everything up to and including the first non-digit
func onNonDigit(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}

	for i := 0; i < len(data); i++ {
		if data[i]-'0' > 9 {
			return i + 1, data[:i+1], nil
		}
	}

	if atEOF {
		return 0, nil, fmt.Errorf("%w", ErrCompressionInvalid)
	}

	return 0, nil, nil
}

var ErrCompressionInvalid = errors.New("the compressed string ended in a digit")
