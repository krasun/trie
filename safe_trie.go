package trie

import "sync"

// safeTrie is safe to use concurrently
type safeTrie struct {
	t  Trie
	mu sync.RWMutex
}

// Safe returns goroutine-safe (thread-safe) version of
// a given trie.
func Safe(trie Trie) Trie {
	return &safeTrie{t: trie}
}

// Insert inserts word into the trie and returns
// true if the trie already contains a record.
//
// It is safe to use concurrently.
func (st *safeTrie) Insert(word string) bool {
	st.mu.Lock()
	defer st.mu.Unlock()

	return st.t.Insert(word)
}

// Contains returns true, if the trie contains a word.
//
// It is safe to use concurrently.
func (st *safeTrie) Contains(word string) bool {
	st.mu.RLock()
	defer st.mu.RUnlock()

	return st.t.Contains(word)
}

// StartsWith returns true if there is any word in the trie that starts
// with the given prefix.
func (st *safeTrie) StartsWith(prefix string) bool {
	st.mu.RLock()
	defer st.mu.RUnlock()

	return st.t.StartsWith(prefix)
}


// Finds and returns words by prefix.
func (st *safeTrie) SearchByPrefix(prefix string) []string {
	st.mu.RLock()
	defer st.mu.RUnlock()

	return st.t.SearchByPrefix(prefix)
}

// Size returns the number of keys in the tree.
func (st *safeTrie) Size() int {
	st.mu.RLock()
	defer st.mu.RUnlock()

	return st.t.Size()
}
