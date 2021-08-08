// Package twelve solves the exercism problem 'Twelve Days of Christmas'
package twelve

import (
	"bytes"
	"os"
	"strings"
)

const songFile = "twelve-days.txt"

// Song returns the full song.
func Song() string {
	verses, err := readSong(songFile)
	if err != nil {
		panic(err)
	}
	return strings.Join(verses, "\n")
}

// Verse returns the requested verse, starting at 1.
func Verse(n int) string {
	verses, err := readSong(songFile)
	if err != nil {
		panic(err)
	}
	return verses[n-1]
}

// readSong reads the entire file and returns the verses of the song.
// Each verse in the song must be separated by a blank line.
// Each verse may be broken down into multiple lines in the file.  If so,
// this function will join them back together so that the result of this function
// contains a single string for each verse.
func readSong(filename string) ([]string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	data = bytes.TrimSpace(data)
	verses := bytes.Split(data, []byte{'\n', '\n'})
	song := make([]string, len(verses))
	for i, v := range verses {
		song[i] = string(bytes.ReplaceAll(v, []byte{'\n'}, []byte{' '}))
	}
	return song, nil
}
