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
	// SearchByPrefix finds and returns words by prefix.
	SearchByPrefix(prefix string) []string
	// StartsWith returns true if there is any word in the trie that starts
	// with the given prefix.
	StartsWith(prefix string) bool
	// Size returns the number of keys in the tree.
	Size() int
}

// New creates a new trie, by default rune-wise.
func New() Trie {
	return NewRuneTrie()
}
