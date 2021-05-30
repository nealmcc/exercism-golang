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
	case "UUU":
		return "Phenylalanine", nil
	case "UUC":
		return "Phenylalanine", nil
	case "UUA":
		return "Leucine", nil
	case "UUG":
		return "Leucine", nil
	case "UCU":
		return "Serine", nil
	case "UCC":
		return "Serine", nil
	case "UCA":
		return "Serine", nil
	case "UCG":
		return "Serine", nil
	case "UAU":
		return "Tyrosine", nil
	case "UAC":
		return "Tyrosine", nil
	case "UGU":
		return "Cysteine", nil
	case "UGC":
		return "Cysteine", nil
	case "UGG":
		return "Tryptophan", nil
	case "UAA":
		return "", ErrStop
	case "UAG":
		return "", ErrStop
	case "UGA":
		return "", ErrStop
	default:
		return "", ErrInvalidBase
	}
}
