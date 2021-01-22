package trie

// domainTrie is a trie implementation optimized for storing and searching domains.
type domainTrie struct {
	root *domainNode
}

type domainNode struct {
	children map[byte]*domainNode
	end      bool
}

// NewDomainTrie creates a new trie optimized for storing and search domains.
func NewDomainTrie() Trie {
	return &domainTrie{root: &domainNode{make(map[byte]*domainNode), false}}
}

// Insert inserts a domain into the trie and returns true
// if the word already exists.
func (t *domainTrie) Insert(domain string) bool {
	exists := true
	current := t.root

	for i := len(domain) - 1; i >= 0; i-- {
		letter := domain[i]
		n, ok := current.children[letter]
		if !ok {
			exists = false

			n = &domainNode{make(map[byte]*domainNode), false}
			current.children[letter] = n
		}

		current = n

		if i == 0 {
			current.end = true
		}
	}

	return exists
}

// Contains returns true if an exact match of the domain is found, otherwise false.
func (t *domainTrie) Contains(domain string) bool {
	current := t.root
	for i := len(domain) - 1; i >= 0; i-- {
		letter := domain[i]
		n, ok := current.children[letter]
		if !ok {
			return false
		}

		current = n
	}

	return current.end
}
