package exercises

import (
	"log"
	"strconv"

	"github.com/vgalaktionov/bioinformatics/lib"
)

func (ex Exercise) BA1E(params []string) []string {
	genome := params[0]
	k, err := strconv.Atoi(params[1])
	if err != nil {
		log.Fatalf("%s", err)
	}
	L, err := strconv.Atoi(params[2])
	if err != nil {
		log.Fatalf("%s", err)
	}
	t, err := strconv.Atoi(params[3])
	if err != nil {
		log.Fatalf("%s", err)
	}

	result := lib.FindClumps(genome, k, L, t)

	return result
}
