package trie

import "testing"

func TestDomainTrieContains(t *testing.T) {
	trie := NewDomainTrie()

	trie.Insert("apple.com")
	trie.Insert("google.com")
	trie.Insert("yahoo.com")

	cases := []struct {
		domain   string
		exists bool
	}{
		{"apple", false},
		{"apple.com", true},
		{"google.com", true},
		{"mail.google.com", false},
	}

	for _, c := range cases {
		actual := trie.Contains(c.domain)
		if actual != c.exists {
			if c.exists {
				t.Errorf("%s is expected to be found in the trie", c.domain)
			} else {
				t.Errorf("%s is not expected to be found in the trie", c.domain)
			}
		}
	}
}
