package lib_test

import (
	"testing"

	"github.com/vgalaktionov/bioinformatics/lib"
)

func TestPatternCount(t *testing.T) {
	expected := 2
	result := lib.PatternCount([]byte("GCGCG"), []byte("GCG"))

	if result != expected {
		t.Errorf("expected: %d; got: %d", expected, result)
	}
}

func BenchmarkPatternCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lib.PatternCount([]byte("GCGCG"), []byte("GCG"))
	}
}
