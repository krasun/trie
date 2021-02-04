package trie

import "testing"

func TestNew(t *testing.T) {
	trie := New()

	_, ok := trie.(*runeTrie)
	if !ok {
		t.Errorf("New function must return rune-wise by default")
	}
}
