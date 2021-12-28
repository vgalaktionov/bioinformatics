package exercises_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vgalaktionov/bioinformatics/exercises"
)

func TestBA1H(t *testing.T) {
	expected := []int{6, 7, 26, 27, 78}
	result := exercises.Exercise{}.BA1H([]string{"ATTCTGGA", "CGCCCGAATCCAGAACGCATTCCCATATTTCGGGACCACTGGCCTCCACGGTACGGACGTCAATCAAATGCCTAGCGGCTTGTGGTTTCTCCTACGCTCC", "3"})

	assert.ElementsMatch(t, result, expected)
}

func BenchmarkBA1H(b *testing.B) {
	for i := 0; i < b.N; i++ {
		exercises.Exercise{}.BA1H([]string{"ATTCTGGA", "CGCCCGAATCCAGAACGCATTCCCATATTTCGGGACCACTGGCCTCCACGGTACGGACGTCAATCAAATGCCTAGCGGCTTGTGGTTTCTCCTACGCTCC", "3"})
	}
}
