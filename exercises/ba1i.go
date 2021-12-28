package exercises

import (
	"strconv"

	"github.com/vgalaktionov/bioinformatics/lib"
)

func frequentWordsWithMismatches(text string, k int, d int) []string {
	frequencies := make(map[string]int)
	kmers := lib.StringProduct("ACTG", k)
	highest := 0
	for _, kmer := range kmers {
		freq := lib.PatternCountApprox(text, kmer, d)
		frequencies[kmer] = freq
		if freq > highest {
			highest = freq
		}
	}

	mostFrequent := []string{}
	for kmer, freq := range frequencies {
		if freq == highest {
			mostFrequent = append(mostFrequent, kmer)
		}
	}
	return mostFrequent
}

func (ex Exercise) BA1I(params []string) []string {
	text := params[0]
	k, err := strconv.Atoi(params[1])
	lib.Check(err)
	d, err := strconv.Atoi(params[2])
	lib.Check(err)

	result := frequentWordsWithMismatches(text, k, d)

	return result
}
