package trie

// runeTrie is a rune-wise trie implementation.
type runeTrie struct {
	root *runeNode
}

type runeNode struct {
	children map[rune]*runeNode
	end      bool
}

// NewRuneTrie creates a rune-wise new trie.
func NewRuneTrie() Trie {
	return &runeTrie{root: &runeNode{make(map[rune]*runeNode), false}}
}

// Insert inserts a word into the trie and returns true
// if the word already exists.
func (t *runeTrie) Insert(word string) bool {
	exists := true
	current := t.root
	for i, letter := range word {
		n, ok := current.children[letter]
		if !ok {
			exists = false

			n = &runeNode{make(map[rune]*runeNode), false}
			current.children[letter] = n
		}

		current = n

		if i == (len(word) - 1) {
			current.end = true
		}
	}

	return exists
}

// Contains returns true if an exact match of the word is found, otherwise false.
func (t *runeTrie) Contains(word string) bool {
	current := t.root
	for _, letter := range word {
		n, ok := current.children[letter]
		if !ok {
			return false
		}

		current = n
	}

	return current.end
}
