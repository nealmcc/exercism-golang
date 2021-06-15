package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
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
		sem     semaphore = make(semaphore, len(texts))
		results []FreqMap = make([]FreqMap, len(texts))
	)

	for i, s := range texts {
		// the anonymous closure captures the values of i, s for each loop
		go func(i int, s string) {
			defer sem.signal()
			results[i] = Frequency(s)
		}(i, s)
	}
	sem.wait(len(texts))

	return sum(results)
}

func sum(freq []FreqMap) FreqMap {
	total := make(FreqMap, 52)
	for _, fm := range freq {
		for r, n := range fm {
			total[r] += n
		}
	}
	return total
}

type empty struct{}

type semaphore chan empty

// send a signal on the channel
func (s semaphore) signal() {
	s <- empty{}
}

// wait until n signals have been sent
func (s semaphore) wait(n int) {
	for i := 0; i < n; i++ {
		<-s
	}
}
