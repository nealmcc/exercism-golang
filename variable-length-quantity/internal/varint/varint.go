package varint

import (
	"encoding/hex"
	"fmt"
	"io"
)

// Varint is a variable length unsigned integer
// the zero value represents 0 and is ready to use
type Varint struct {
	// octets are the encoded octets for this Varint.
	// index 0 is the most significant byte.
	// all bytes except the last will have their continuation markers set
	octets []byte
}

// FromUint creates a new Varint from the given uint32
func FromUint(val uint32) Varint {
	octets := make([]byte, 0, 32/7+1)
	octets = append(octets, byte(val&0x7f))
	val >>= 7
	for val > 0 {
		octets = append(octets, byte((val&0x7f)|0x80))
		val >>= 7
	}

	// reverse the order of the octets, so index 0 is most significant
	for i, j := 0, len(octets)-1; i < j; i, j = i+1, j-1 {
		octets[i], octets[j] = octets[j], octets[i]
	}

	return Varint{octets}
}

// ToUint decodes this varint as an unsigned 32-bit integer.
// If this varint is larger than a uint32, the result will overflow.
func (v Varint) ToUint() uint32 {
	var total uint32
	for _, oct := range v.octets {
		total = total<<7 | uint32(oct&0x7f)
	}
	return total
}

// FromReader creates a new Varint, consuming data from the reader.
// Stops reading after the current byte no longer has the continuation marker,
// or when the input runs out.  Returns an error if the input ends without
// a termination marker.
func FromReader(r io.Reader) (Varint, error) {
	var (
		b      []byte = make([]byte, 1)
		octets []byte = make([]byte, 0, 8)
		done   bool
		err    error
	)
	for !done {
		_, err = r.Read(b)
		done = b[0]&0x80 == 0
		if err != nil {
			if err != io.EOF {
				return Varint{}, err
			}
			if !done {
				return Varint{}, io.ErrUnexpectedEOF
			}
		}
		octets = append(octets, b[0])
	}
	return Varint{octets}, err
}

// WriteTo implements io.WriterTo
//
// WriteTo writes the binary representation of this Varint to the given writer
func (v Varint) WriteTo(w io.Writer) (int64, error) {
	if len(v.octets) == 0 {
		n, err := w.Write([]byte{0})
		return int64(n), err
	}

	n, err := w.Write(v.octets)
	return int64(n), err
}

var _ io.WriterTo = Varint{}

// String implements fmt.Stringer
//
// String includes the byte length and hexadecimal octets
func (v Varint) String() string {
	return fmt.Sprintf("[%d] {%s}", len(v.octets), hex.EncodeToString(v.octets))
}

var _ fmt.Stringer = Varint{}
