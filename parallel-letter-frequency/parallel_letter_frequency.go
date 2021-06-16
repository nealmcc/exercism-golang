package letter

import "sync"

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := make(FreqMap, 32)
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency counts the frequency of each rune in all of the given
// texts and returns this data as a FreqMap.
// uses the parallel for-loop pattern described here:
// http://www.golangpatterns.info/concurrency/parallel-for-loop
func ConcurrentFrequency(texts []string) FreqMap {
	var (
		wg      sync.WaitGroup
		results []FreqMap = make([]FreqMap, len(texts))
	)

	wg.Add(len(texts))

	for i, s := range texts {
		// the anonymous closure captures the values of i, s for each loop
		go func(i int, s string) {
			results[i] = Frequency(s)
			wg.Done()
		}(i, s)
	}

	wg.Wait()
	return sum(results)
}

func sum(freq []FreqMap) FreqMap {
	total := make(FreqMap, 32)
	for _, fm := range freq {
		for r, n := range fm {
			total[r] += n
		}
	}
	return total
}
