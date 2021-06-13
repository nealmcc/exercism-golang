package cipher

import (
	"strings"
)

// A Cipher provides an algorithm to read from a stream of bytes,
// keeping and transposing some of those bytes while discarding others.
// The Cipher can then write the cipher text to an output stream.
type Cipher struct {
	opts options
}

// New creates a new Cipher with the ability to override individual Options.
// The default Options implement the Atbash cipher from the Exercism.io problem.
func New(opts ...Option) Cipher {
	options := options{
		keep:      AtbashKeep,
		transpose: AtbashTranspose,
		blockSize: 5,
	}

	for _, o := range opts {
		o.apply(&options)
	}

	return Cipher{opts: options}
}

// Each Option is used to override the default behaviour of a Cipher.
// Use WithKeep(), WithTranspose() and WithBlockSize() to create Options
// that can be passed in when creating a new Cipher.
type Option interface {
	apply(*options)
}

type options struct {
	keep      KeepFn
	transpose TransposeFn
	blockSize int
}

// a KeepFn decides which bytes to keep and which to discard, before transposing
type KeepFn func(byte) bool

// WithKeep is used when creating a new Cipher to override the default KeepFn
func WithKeep(fn func(ch byte) bool) Option {
	return KeepFn(fn)
}

func (k KeepFn) apply(o *options) {
	o.keep = k
}

// AtbashKeep is the default KeepFn.
// It keeps ASCII letters and digits and discards anythnig else.
func AtbashKeep(ch byte) bool {
	switch {
	case 'a' <= ch && ch <= 'z':
		return true
	case 'A' <= ch && ch <= 'Z':
		return true
	case '0' <= ch && ch <= '9':
		return true
	default:
		return false
	}
}

// a TransposeFn converts a plaintext byte to a ciphertext byte
type TransposeFn func(byte) byte

// WithTranspose is used when creating a new Cipher to override the default.
// The given function performs the character substitution for the Cipher.
// Transposition only happens for bytes that have been kept.
func WithTranspose(fn func(ch byte) byte) Option {
	return TransposeFn(fn)
}

func (t TransposeFn) apply(o *options) {
	o.transpose = t
}

var inverse = []byte{'z', 'y', 'x', 'w', 'v', 'u', 't', 's', 'r', 'q', 'p', 'o', 'n', 'm', 'l', 'k', 'j', 'i', 'h', 'g', 'f', 'e', 'd', 'c', 'b', 'a'}

// AtbashTranspose is the default TransposeFn.
// It converts letters to lowercase, with 'a' <-> 'z', 'b' <-> 'y' etc..
func AtbashTranspose(ch byte) byte {
	switch {
	case 'a' <= ch && ch <= 'z':
		return inverse[ch-'a']
	case 'A' <= ch && ch <= 'Z':
		return inverse[ch-'A']
	default:
		return ch
	}
}

// WithBlockSize is used when creating a new Cipher to
// override the default size of 'words' when writing the ciphertext.
// If size <= 0 then the ciphertext will not be broken up into 'words'.
func WithBlockSize(size int) Option {
	return sizeOpt(size)
}

type sizeOpt int

func (sz sizeOpt) apply(o *options) {
	o.blockSize = int(sz)
}

// Encode returns the ciphertext for the given plaintext
func (c Cipher) Encode(plain string) string {
	var (
		keep      KeepFn      = c.opts.keep
		transpose TransposeFn = c.opts.transpose
		blockSize int         = c.opts.blockSize
		curr      int
		next      int
	)

	// filter and transpose the plain text
	text := []byte(plain)
	for ; next < len(text); next++ {
		ch := text[next]
		if keep(ch) {
			text[curr] = transpose(ch)
			curr++
		}
	}
	text = text[:curr]

	if blockSize <= 0 {
		return string(text)
	}

	// allocate a string builder of sufficient size
	b := new(strings.Builder)
	b.Grow(len(text) + len(text)/blockSize)

	// write each chunk to the string builder
	for start, end := 0, 0; end < len(text); start = end {
		end += blockSize

		// the last chunk could be shorter than blockSize
		if end > len(text) {
			end = len(text)
		}

		b.Write(text[start:end])

		// add a space if we have at least one more chunk
		if end < len(text) {
			b.WriteByte(' ')
		}
	}
	return b.String()
}
