package etl

func Transform(in map[int][]string) map[string]int {
	out := map[string]int{}
	for score, letters := range in {
		for _, s := range letters {
			ch := []byte(s)
			if 'a' > []byte(ch)[0] {
				ch[0] += 'a' - 'A'
			}
			out[string(ch)] = score
		}
	}
	return out
}
