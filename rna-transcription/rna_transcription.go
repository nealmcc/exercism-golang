package strand

type replacer []byte

var oldnew = replacer{
	'C', 'G',
	'G', 'C',
	'T', 'A',
	'A', 'U',
}

var maxK = len(oldnew) - 2

func ToRNA(old string) string {
	length := len(old)
	new := make([]byte, length)
	for i := 0; i < length; i++ {
		for k := 0; k <= maxK; k += 2 {
			if old[i] == oldnew[k] {
				new[i] = oldnew[k+1]
			}
		}
	}
	return string(new)
}
