package lib

import (
	"bytes"
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
