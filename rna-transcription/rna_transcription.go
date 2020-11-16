package strand

var translate = map[byte]byte{
	'C': 'G',
	'G': 'C',
	'T': 'A',
	'A': 'U',
}

func ToRNA(dna string) string {
	length := len(dna)
	rna := make([]byte, length)
	for i := 0; i < length; i++ {
		rna[i] = translate[dna[i]]
	}
	return string(rna)
}
