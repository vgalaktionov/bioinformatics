package exercises

import (
	"log"
	"strconv"

	"github.com/vgalaktionov/bioinformatics/lib"
)

func (ex Exercise) BA1B(params []string) []string {
	text := []byte(params[0])
	k, err := strconv.Atoi(params[1])
	if err != nil {
		log.Fatalf("%s", err)
	}

	result := lib.FrequentWords(text, k, 0)

	return result
}
