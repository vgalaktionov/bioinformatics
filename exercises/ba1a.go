package exercises

import (
	"github.com/vgalaktionov/bioinformatics/lib"
)

func (ex Exercise) BA1A(params []string) int {
	text := []byte(params[0])
	pattern := []byte(params[1])

	result := lib.PatternCount(text, pattern)

	return result
}
