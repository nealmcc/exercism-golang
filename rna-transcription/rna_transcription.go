package strand

type replacer []byte

var dnarna = replacer{
	'C', 'G',
	'G', 'C',
	'T', 'A',
	'A', 'U',
}

var maxD = len(dnarna) - 2

func ToRNA(dna string) string {
	length := len(dna)
	rna := make([]byte, length)
	for i := 0; i < length; i++ {
		for d := 0; d <= maxD; d += 2 {
			if dna[i] == dnarna[d] {
				rna[i] = dnarna[d+1]
			}
		}
	}
	return string(rna)
}
