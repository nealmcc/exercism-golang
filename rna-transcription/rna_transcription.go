package strand

import "strings"

var translation = strings.NewReplacer(
	"C", "G",
	"G", "C",
	"T", "A",
	"A", "U",
)

func ToRNA(dna string) string {
	return translation.Replace(dna)
}
