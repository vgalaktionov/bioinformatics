package exercises

import "github.com/vgalaktionov/bioinformatics/lib"

func (ex Exercise) BA1C(params []string) string {
	pattern := params[0]

	result := lib.ReverseComplement(pattern)

	return result
}
