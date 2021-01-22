package trie

import "testing"

func TestRuneTrieContains(t *testing.T) {
	trie := NewRuneTrie()

	trie.Insert("c")
	trie.Insert("apple")
	trie.Insert("banana")

	cases := []struct {
		word   string
		exists bool
	}{
		{"c", true},
		{"ce", false},
		{"banana", true},
		{"app", false},
		{"aple", false},
		{"ban", false},
		{"apple", true},
		{"apple1", false},
		{"aaapple", false},
	}

	for _, c := range cases {
		actual := trie.Contains(c.word)
		if actual != c.exists {
			if c.exists {
				t.Errorf("%s is expected to be found in the trie", c.word)
			} else {
				t.Errorf("%s is not expected to be found in the trie", c.word)
			}
		}
	}
}
