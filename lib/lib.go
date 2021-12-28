package lib

import (
	"log"
)

func PatternCount(text string, pattern string) int {
	n := 0
	lastI := len(text) - len(pattern)
	for i := 0; i <= lastI; i++ {
		window := text[i : i+len(pattern)]
		if window == pattern {
			n++
		}
	}
	return n
}

func PatternFind(text string, pattern string) []int {
	lastI := len(text) - len(pattern)
	positions := make([]int, 0, lastI)
	for i := 0; i <= lastI; i++ {
		window := text[i : i+len(pattern)]
		if window == pattern {
			positions = append(positions, i)
		}
	}
	return positions
}

func FrequentWords(text string, k int, minFreq int) []string {
	frequencies := make(map[string]int)
	lastI := len(text) - k
	highest := 0
	for i := 0; i <= lastI; i++ {
		kmer := text[i : i+k]
		if _, seen := frequencies[kmer]; !seen {
			freq := PatternCount(text[i:], kmer)
			frequencies[kmer] = freq
			if freq > highest {
				highest = freq
			}
		}
	}

	mostFrequent := []string{}
	for kmer, freq := range frequencies {
		if freq == highest && freq >= minFreq {
			mostFrequent = append(mostFrequent, kmer)
		}
	}
	return mostFrequent
}

func ReverseComplement(pattern string) string {
	n := len(pattern)
	runes := make([]rune, n)
	for _, nucleotide := range pattern {
		n--
		var complement rune
		switch nucleotide {
		case 'A':
			complement = 'T'
		case 'C':
			complement = 'G'
		case 'G':
			complement = 'C'
		case 'T':
			complement = 'A'
		}
		runes[n] = complement
	}
	return string(runes[n:])
}

func uniqueStrings(stringSlice []string) []string {
	keys := make(map[string]bool)
	out := make([]string, 0, len(stringSlice))
	for _, entry := range stringSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			out = append(out, entry)
		}
	}
	return out
}

func FindClumps(genome string, k int, L int, t int) []string {
	lastI := len(genome) - k - L
	ch := make(chan []string, lastI)
	for i := 0; i <= lastI; i++ {
		window := genome[i : i+L]
		go func() {
			ch <- FrequentWords(window, k, t)
		}()
	}
	clumps := make([]string, 0, lastI*lastI)
	for i := 0; i <= lastI; i++ {
		clumps = append(clumps, <-ch...)
	}
	return uniqueStrings(clumps)
}

func MinSkew(genome string) []int {
	skew := make([]int, len(genome)+1)
	min := 0
	for i, nucleotide := range genome {
		var newSkew int
		if nucleotide == 'G' {
			newSkew = skew[i] + 1
		} else if nucleotide == 'C' {
			newSkew = skew[i] - 1
		} else {
			newSkew = skew[i]
		}
		skew[i+1] = newSkew
		if newSkew < min {
			min = newSkew
		}
	}
	positions := []int{}
	for i, s := range skew {
		if s == min {
			positions = append(positions, i)
		}
	}
	return positions
}

func HammingDistance(textA string, textB string) int {
	if len(textA) != len(textB) {
		log.Fatalf("strings must be of equal length")
	}
	distance := 0
	for i := 0; i < len(textA); i++ {
		if textA[i] != textB[i] {
			distance++
		}
	}
	return distance
}

func PatternCountApprox(text string, pattern string, d int) int {
	lastI := len(text) - len(pattern)
	count := 0
	for i := 0; i <= lastI; i++ {
		window := text[i : i+len(pattern)]
		if HammingDistance(window, pattern) <= d {
			count++
		}
	}
	return count
}

func PatternFindApprox(text string, pattern string, d int) []int {
	lastI := len(text) - len(pattern)
	positions := make([]int, 0, lastI)
	for i := 0; i <= lastI; i++ {
		window := text[i : i+len(pattern)]
		if HammingDistance(window, pattern) <= d {
			positions = append(positions, i)
		}
	}
	return positions
}

func cartN(a ...string) []string {
	if len(a) == 0 {
		return []string{""}
	}
	acc := cartN(a[1:]...)
	out := []string{}
	for _, e := range a[0] {
		for _, p := range acc {
			comb := string(p) + string(e)
			if comb != "" {
				out = append(out, string(p)+string(e))
			}
		}
	}
	return out
}

func StringProduct(a string, r int) []string {
	pools := make([]string, r)
	for i := 0; i < r; i++ {
		pools[i] = a
	}
	result := cartN(pools...)
	return result
}
