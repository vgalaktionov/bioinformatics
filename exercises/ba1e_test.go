package exercises_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vgalaktionov/bioinformatics/exercises"
)

func TestBA1E(t *testing.T) {
	expected := []string{"CGACA", "GAAGA", "AATGT"}
	result := exercises.Exercise{}.BA1E([]string{"CGGACTCGACAGATGTGAAGAAATGTGAAGACTGAGTGAAGAGAAGAGGAAACACGACACGACATTGCGACATAATGTACGAATGTAATGTGCCTATGGC", "5", "75", "4"})

	assert.ElementsMatch(t, result, expected)
}

func BenchmarkBA1E(b *testing.B) {
	for i := 0; i < b.N; i++ {
		exercises.Exercise{}.BA1E([]string{"CGGACTCGACAGATGTGAAGAAATGTGAAGACTGAGTGAAGAGAAGAGGAAACACGACACGACATTGCGACATAATGTACGAATGTAATGTGCCTATGGC", "5", "75", "4"})
	}
}
