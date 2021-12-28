package exercises_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vgalaktionov/bioinformatics/exercises"
)

func TestBA1J(t *testing.T) {
	expected := []string{"ATGT", "ACAT"}
	result := exercises.Exercise{}.BA1J([]string{"ACGTTGCATGTCGCATGATGCATGAGAGCT", "4", "1"})

	assert.ElementsMatch(t, result, expected)
}

func BenchmarkBA1J(b *testing.B) {
	for i := 0; i < b.N; i++ {
		exercises.Exercise{}.BA1J([]string{"ACGTTGCATGTCGCATGATGCATGAGAGCT", "4", "1"})
	}
}
