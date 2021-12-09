package exercises

import (
	"github.com/vgalaktionov/bioinformatics/lib"
)

func (ex Exercise) BA1A(params [][]byte) int {
	text := params[0]
	pattern := params[1]

	result := lib.PatternCount(text, pattern)

	return result
}
