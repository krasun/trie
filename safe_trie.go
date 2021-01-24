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
