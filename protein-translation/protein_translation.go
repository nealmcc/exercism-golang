package protein

import (
	"errors"
	"io"
	"strings"
)

var (
	ErrStop        = errors.New("stop")
	ErrInvalidBase = errors.New("invalid base")
)

type (
	codon     = string
	aminoAcid = string
	protein   = []string
)

// FromRNA converts an RNA sequence to the corresponding protein
func FromRNA(rna string) (protein, error) {
	protein := make([]aminoAcid, 0, len(rna)/3)
	buf := make([]byte, 3)
	r := strings.NewReader(rna)
	for {
		_, err := io.ReadFull(r, buf)
		if err != nil {
			if err == io.EOF {
				return protein, nil
			}
			return protein, err
		}
		acid, err := FromCodon(string(buf))
		if err != nil {
			if err == ErrStop {
				return protein, nil
			}
			return protein, err
		}
		protein = append(protein, acid)
	}
}

// FromCodon translates a single codon to a single amino acid
func FromCodon(c codon) (aminoAcid, error) {
	switch c {
	case "AUG":
		return "Methionine", nil

	case "UUU", "UUC":
		return "Phenylalanine", nil

	case "UUA", "UUG":
		return "Leucine", nil

	case "UCU", "UCC", "UCA", "UCG":
		return "Serine", nil

	case "UAU", "UAC":
		return "Tyrosine", nil

	case "UGU", "UGC":
		return "Cysteine", nil

	case "UGG":
		return "Tryptophan", nil

	case "UAA", "UAG", "UGA":
		return "", ErrStop

	default:
		return "", ErrInvalidBase
	}
}
