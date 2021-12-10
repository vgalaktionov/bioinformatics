package lib

import (
	"bytes"
	"sort"
)

func PatternCount(text []byte, pattern []byte) int {
	n := 0
	lastI := len(text) - len(pattern)
	for i := 0; i <= lastI; i++ {
		window := text[i : i+len(pattern)]
		if bytes.Equal(window, pattern) {
			n++
		}
	}
	return n
}

func FrequentWords(text []byte, k int, minFreq int) []string {
	frequencies := make(map[string]int)
	lastI := len(text) - k
	for i := 0; i <= lastI; i++ {
		kmer := text[i : i+k]
		_, seen := frequencies[string(kmer)]
		if !seen {
			frequencies[string(kmer)] = PatternCount(text[i:], kmer)
		}
	}

	keys := make([]string, 0, len(frequencies))
	for k := range frequencies {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		return frequencies[keys[i]] < frequencies[keys[j]]
	})
	highest := frequencies[keys[0]]
	mostFrequent := []string{}
	for _, k := range keys {
		freq := frequencies[k]
		if freq != highest || freq < minFreq {
			mostFrequent = append(mostFrequent, k)
		}
	}
	return mostFrequent
}
