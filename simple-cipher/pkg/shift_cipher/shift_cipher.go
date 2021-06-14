package shiftcipher

// a Cipher is able to encode and decode ASCII text
type Cipher interface {
	Encode(string) string
	Decode(string) string
}

// New creates a new Cipher with the ability to override individual Options.
// The default cipher removes non letters, and converts the input to lowercase.
func New(opts ...Option) Cipher {
	options := options{
		keep:    KeepLetters,
		encoder: ToLower,
		decoder: ToLower,
	}

	for _, o := range opts {
		o.apply(&options)
	}

	return cipher{options}
}

type cipher struct {
	opts options
}

// compile-time interface check
var _ Cipher = cipher{}

type Option interface {
	apply(o *options)
}

type options struct {
	keep    FilterFn
	encoder TransposeFn
	decoder TransposeFn
}

func (c cipher) Encode(plain string) string {
	var (
		keep   FilterFn    = c.opts.keep
		encode TransposeFn = c.opts.encoder
		curr   int
		next   int
	)

	// filter and transpose the plain text one byte at a time:
	text := []byte(plain)
	for ; next < len(text); next++ {
		ch := text[next]
		if keep(ch) {
			text[curr] = encode(curr, ch)
			curr++
		}
	}
	text = text[:curr]

	return string(text)
}

func (c cipher) Decode(text string) string {
	var (
		decode TransposeFn = c.opts.decoder
		plain  []byte      = make([]byte, len(text))
	)

	for i, ch := range []byte(text) {
		plain[i] = decode(i, ch)
	}

	return string(plain)
}

// a FilterFn decides which bytes to keep vs discard before transposing
type FilterFn func(byte) bool

// a TransposeFn converts a plaintext byte to a ciphertext byte or vice versa
type TransposeFn func(index int, ch byte) byte

// WithFilter is used when creating a new Cipher to override the default FilterFn
func WithFilter(f func(byte) bool) Option {
	return FilterFn(f)
}

func (f FilterFn) apply(o *options) {
	o.keep = f
}

// KeepAll is a FilterFn which keeps all of the input bytes
func KeepAll(ch byte) bool {
	return true
}

// KeepLetters is the FilterFn that only keeps ASCII letters
func KeepLetters(ch byte) bool {
	switch {
	case 'a' <= ch && ch <= 'z':
		return true
	case 'A' <= ch && ch <= 'Z':
		return true
	default:
		return false
	}
}

// WithEncoder can be used when creating a new Cipher.  This option overrides
// the default encoder.
// Transposition only happens for bytes that have been kept.
// See also: WithDecoder()
func WithEncoder(fn func(index int, ch byte) byte) Option {
	return encodeOpt(fn)
}

type encodeOpt TransposeFn

func (enc encodeOpt) apply(o *options) {
	o.encoder = TransposeFn(enc)
}

// WithDecoder can be used when creating a new Cipher.  This option overrides
// the default decoder.
// See also: WithEncoder()
func WithDecoder(fn func(index int, ch byte) byte) Option {
	return decodeOpt(fn)
}

type decodeOpt TransposeFn

func (enc decodeOpt) apply(o *options) {
	o.decoder = TransposeFn(enc)
}

// Pass is a TransposeFn which does not alter the bytes
func Pass(_ int, ch byte) byte {
	return ch
}

// ToLower is the default TransposeFn which converts ASCII letters to lowercase
func ToLower(_ int, ch byte) byte {
	switch {
	case 'a' <= ch && ch <= 'z':
		return ch - 'a'
	case 'A' <= ch && ch <= 'Z':
		return ch - 'A'
	default:
		return ch
	}
}

// WithShift provides an option that overrides both the encoder and decoder.
// The encoder shifts ASCII letters by n, and the decoder unshifts them by
// the same amount. Anything other than an ASCII letter is unaffected.
func WithShift(n int) Option {
	n = n % 26
	if n < 0 {
		n += 26
	}

	return codecOpt{
		enc: makeShifter(byte(n)),
		dec: makeShifter(byte(26 - n)),
	}
}

func makeShifter(n byte) TransposeFn {
	shift := func(b byte) byte {
		b = (b + n) % 26
		return b
	}

	return func(_ int, ch byte) byte {
		switch {
		case 'a' <= ch && ch <= 'z':
			return shift(ch-'a') + 'a'
		case 'A' <= ch && ch <= 'Z':
			return shift(ch-'A') + 'a'
		}
		return ch
	}
}

type codecOpt struct {
	enc TransposeFn
	dec TransposeFn
}

func (c codecOpt) apply(o *options) {
	o.encoder = c.enc
	o.decoder = c.dec
}

// WithPattern provides an option that overrides both the encoder and decoder.
// The given pattern is used when encoding, and the inverse of this pattern is
// used when decoding.
func WithPattern(pattern []byte) Option {
	var (
		penc []byte = make([]byte, len(pattern))
		pdec []byte = make([]byte, len(pattern))
	)

	for i, x := range pattern {
		penc[i] = x - 'a'
		pdec[i] = 26 - x + 'a'
	}

	return codecOpt{
		enc: makePatternShifter(penc),
		dec: makePatternShifter(pdec),
	}
}

func makePatternShifter(offsets []byte) TransposeFn {
	shift := func(i int, b byte) byte {
		n := offsets[i%len(offsets)]
		b = (b + n) % 26
		return b
	}

	return func(i int, ch byte) byte {
		switch {
		case 'a' <= ch && ch <= 'z':
			return shift(i, ch-'a') + 'a'
		case 'A' <= ch && ch <= 'Z':
			return shift(i, ch-'A') + 'a'
		}
		return ch
	}
}
