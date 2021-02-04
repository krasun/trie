package trie

import "testing"

// to prevent compiler optimizations, all benchmark results are stored
// in the package-level variable
var benchmarkResult bool
var benchmarkSetResult map[string]struct{}

func TestNew(t *testing.T) {
	trie := New()

	_, ok := trie.(*runeTrie)
	if !ok {
		t.Errorf("New function must return rune-wise by default")
	}
}
