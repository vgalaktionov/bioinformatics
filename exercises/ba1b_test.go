package exercises_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vgalaktionov/bioinformatics/exercises"
)

func TestBA1B(t *testing.T) {
	expected := []string{"CATG", "GCAT"}
	result := exercises.Exercise{}.BA1B([]string{"ACGTTGCATGTCGCATGATGCATGAGAGCT", "4"})

	assert.ElementsMatch(t, result, expected)
}

func BenchmarkBA1B(b *testing.B) {
	for i := 0; i < b.N; i++ {
		exercises.Exercise{}.BA1B([]string{"ACGTTGCATGTCGCATGATGCATGAGAGCT", "4"})
	}
}
