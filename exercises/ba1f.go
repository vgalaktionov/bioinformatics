package exercises

import "github.com/vgalaktionov/bioinformatics/lib"

func (ex Exercise) BA1F(params []string) []int {
	genome := params[0]

	result := lib.MinSkew(genome)

	return result
}
