package exercises_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vgalaktionov/bioinformatics/exercises"
)

func TestBA1F(t *testing.T) {
	expected := []int{53, 97}
	result := exercises.Exercise{}.BA1F([]string{"CCTATCGGTGGATTAGCATGTCCCTGTACGTTTCGCCGCGAACTAGTTCACACGGCTTGATGGCAAATGGTTTTTCCGGCGACCGTAATCGTCCACCGAG"})

	assert.ElementsMatch(t, result, expected)
}

func BenchmarkBA1F(b *testing.B) {
	for i := 0; i < b.N; i++ {
		exercises.Exercise{}.BA1F([]string{"CCTATCGGTGGATTAGCATGTCCCTGTACGTTTCGCCGCGAACTAGTTCACACGGCTTGATGGCAAATGGTTTTTCCGGCGACCGTAATCGTCCACCGAG"})
	}
}
