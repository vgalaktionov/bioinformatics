package exercises

import (
	"log"
	"strconv"

	"github.com/vgalaktionov/bioinformatics/lib"
)

func (ex Exercise) BA1H(params []string) []int {
	pattern := params[0]
	text := params[1]
	d, err := strconv.Atoi(params[2])
	if err != nil {
		log.Fatalf("%s", err)
	}

	result := lib.PatternFindApprox(text, pattern, d)

	return result
}
