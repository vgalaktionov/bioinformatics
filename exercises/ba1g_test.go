package exercises_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vgalaktionov/bioinformatics/exercises"
)

func TestBA1G(t *testing.T) {
	expected := 3
	result := exercises.Exercise{}.BA1G([]string{"GGGCCGTTGGT", "GGACCGTTGAC"})

	assert.Equal(t, result, expected)
}

func BenchmarkBA1G(b *testing.B) {
	for i := 0; i < b.N; i++ {
		exercises.Exercise{}.BA1G([]string{"GGGCCGTTGGT", "GGACCGTTGAC"})
	}
}
