package exercises

import (
	"strconv"

	"github.com/vgalaktionov/bioinformatics/lib"
)

func (ex Exercise) BA1E(params []string) []string {
	genome := params[0]
	k, err := strconv.Atoi(params[1])
	lib.Check(err)
	L, err := strconv.Atoi(params[2])
	lib.Check(err)
	t, err := strconv.Atoi(params[3])
	lib.Check(err)

	result := lib.FindClumps(genome, k, L, t)

	return result
}
