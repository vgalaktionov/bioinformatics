package exercises_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vgalaktionov/bioinformatics/exercises"
)

func TestBA1I(t *testing.T) {
	expected := []string{"GATG", "ATGC", "ATGT"}
	result := exercises.Exercise{}.BA1I([]string{"ACGTTGCATGTCGCATGATGCATGAGAGCT", "4", "1"})

	assert.ElementsMatch(t, result, expected)
}

func BenchmarkBA1I(b *testing.B) {
	for i := 0; i < b.N; i++ {
		exercises.Exercise{}.BA1I([]string{"ACGTTGCATGTCGCATGATGCATGAGAGCT", "4", "1"})
	}
}
