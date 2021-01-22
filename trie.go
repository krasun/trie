package trie

// Trie represents a Trie data structure.
//
// https://en.wikipedia.org/wiki/RuneTrie
type Trie interface {
	Insert(word string) bool
	Contains(word string) bool
}

// New creates a new trie, by default rune-wise.
func New() Trie {
	return NewRuneTrie()
}
