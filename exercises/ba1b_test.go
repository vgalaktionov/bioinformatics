package exercises_test

import (
	"testing"

	"github.com/vgalaktionov/bioinformatics/exercises"
)

func TestBA1B(t *testing.T) {
	expected := "CATG GCAT"
	result := exercises.Exercise{}.BA1B([][]byte{[]byte("ACGTTGCATGTCGCATGATGCATGAGAGCT"), []byte("4")})

	if result != expected {
		t.Errorf("expected: %s; got: %s", expected, result)
	}
}

func BenchmarkBA1B(b *testing.B) {
	for i := 0; i < b.N; i++ {
		exercises.Exercise{}.BA1B([][]byte{[]byte("ACGTTGCATGTCGCATGATGCATGAGAGCT"), []byte("4")})
	}
}
