package exercises_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vgalaktionov/bioinformatics/exercises"
)

func TestBA1D(t *testing.T) {
	expected := []int{1, 3, 9}
	result := exercises.Exercise{}.BA1D([]string{"ATAT", "GATATATGCATATACTT"})

	assert.ElementsMatch(t, result, expected)
}

func BenchmarkBA1D(b *testing.B) {
	for i := 0; i < b.N; i++ {
		exercises.Exercise{}.BA1D([]string{"ATAT", "GATATATGCATATACTT"})
	}
}
