package exercises_test

import (
	"testing"

	"github.com/vgalaktionov/bioinformatics/exercises"
)

func TestBA1C(t *testing.T) {
	expected := "ACCGGGTTTT"
	result := exercises.Exercise{}.BA1C([]string{"AAAACCCGGT"})

	if result != expected {
		t.Errorf("expected: %s; got: %s", expected, result)
	}
}

func BenchmarkBA1C(b *testing.B) {
	for i := 0; i < b.N; i++ {
		exercises.Exercise{}.BA1C([]string{"AAAACCCGGT"})
	}
}
