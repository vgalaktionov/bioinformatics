package exercises_test

import (
	"testing"

	"github.com/vgalaktionov/bioinformatics/exercises"
)

func TestBA1A(t *testing.T) {
	expected := 2
	result := exercises.Exercise{}.BA1A([][]byte{[]byte("GCGCG"), []byte("GCG")})

	if result != expected {
		t.Errorf("expected: %d; got: %d", expected, result)
	}
}

func BenchmarkBA1A(b *testing.B) {
	for i := 0; i < b.N; i++ {
		exercises.Exercise{}.BA1A([][]byte{[]byte("GCGCG"), []byte("GCG")})
	}
}
