package trie

import (
	"fmt"
	"testing"
)

func TestSafeTrieContains(t *testing.T) {
	trie := Safe(NewRuneTrie())

	for i := 0; i < 10; i++ {
		trie.Insert("c")
		trie.Insert("apple")
		trie.Insert("banana")

		go func(index int) {
			trie.Insert(fmt.Sprintf("c%d", index))
			trie.Insert(fmt.Sprintf("apple%d", index))
			trie.Insert(fmt.Sprintf("banana%d", index))
		}(i)
	}

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
		{"apple25", false},
		{"aaapple", false},
	}

	done := make(chan bool)
	go func() {
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
		done <- true
	}()

	<-done
}
