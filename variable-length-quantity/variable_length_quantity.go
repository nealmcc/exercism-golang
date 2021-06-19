package variablelengthquantity

import (
	"bytes"
	"io"
	"log"
	"variablelengthquantity/internal/varint"
)

func EncodeVarint(input []uint32) []byte {
	var output bytes.Buffer
	for _, n := range input {
		v := varint.FromUint(n)
		if _, err := v.WriteTo(&output); err != nil {
			log.Fatal(err)
		}
	}
	return output.Bytes()
}

func DecodeVarint(input []byte) ([]uint32, error) {
	var (
		r    io.Reader = bytes.NewReader(input)
		ints []uint32
	)

	for {
		v, err := varint.FromReader(r)
		if err != nil {
			if err == io.EOF {
				return ints, nil
			}
			return nil, err
		}
		ints = append(ints, v.ToUint())
	}
}
