package exercises

import "github.com/vgalaktionov/bioinformatics/lib"

func (ex Exercise) BA1D(params []string) []int {
	pattern := params[0]
	text := params[1]

	result := lib.PatternFind(text, pattern)

	return result
}
