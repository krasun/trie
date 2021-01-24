package trie

// Trie represents a Trie data structure.
//
// https://en.wikipedia.org/wiki/Trie
type Trie interface {
	// Insert inserts a word into the trie and returns true
	// if the word already exists.
	Insert(word string) bool
	// Contains returns true if an exact match of the word is found, otherwise false.
	Contains(word string) bool
}

// New creates a new trie, by default rune-wise.
func New() Trie {
	return NewRuneTrie()
}
