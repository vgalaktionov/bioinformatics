package exercises

import "github.com/vgalaktionov/bioinformatics/lib"

func (ex Exercise) BA1G(params []string) int {
	a := params[0]
	b := params[1]

	result := lib.HammingDistance(a, b)

	return result
}
