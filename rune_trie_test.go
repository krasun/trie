package trie

import (
	"reflect"
	"sort"
	"testing"
)

var words = []string{
	"c", "ce", "banana", "app", "aple", "ban", "apple1", "aaapple",
	"1c", "1ce", "1banana", "1app", "1aple", "1ban", "1apple1", "1aaapple",
}

var filledTrieHashSet = (func() map[string]struct{} {
	s := make(map[string]struct{})
	for _, w := range words {
		s[w] = struct{}{}
	}

	return s
})()

var filledRuneTrie = (func() Trie {
	t := NewRuneTrie()
	for _, w := range words {
		t.Insert(w)
	}

	return t
})()

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
		{"cc", false},
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

func TestRuneTrieStartsWith(t *testing.T) {
	trie := NewRuneTrie()

	trie.Insert("apple")
	trie.Insert("aplhabet")
	trie.Insert("tree")

	if !trie.StartsWith("a") {
		t.Error("trie StartsWith(a) must return true, but returned false")
	}

	if trie.StartsWith("try") {
		t.Error("trie StartsWith(try) must return false, but returned true")
	}
}

func TestRuneTrieSearchByPrefix(t *testing.T) {
	trie := NewRuneTrie()

	trie.Insert("c")
	trie.Insert("apple")
	trie.Insert("banana")
	trie.Insert("alphabet")
	trie.Insert("alcohol")

	actual := trie.SearchByPrefix("a")
	sort.Strings(actual)

	expected := []string{"alcohol", "alphabet", "apple"}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("%v != %v", actual, expected)
	}
}

func BenchmarkRuneTrieInsert(b *testing.B) {
	var r bool
	for i := 0; i < b.N; i++ {
		t := NewRuneTrie()
		for _, w := range words {
			r = t.Insert(w)
		}
	}

	benchmarkResult = r
}

func BenchmarkRuneTrieContains(b *testing.B) {
	var r bool
	for i := 0; i < b.N; i++ {
		for _, w := range words {
			r = filledRuneTrie.Contains(w)
		}
	}

	benchmarkResult = r
}

func BenchmarkWordHashSetInsert(b *testing.B) {
	var r map[string]struct{}
	for i := 0; i < b.N; i++ {
		s := make(map[string]struct{})
		for _, w := range words {
			s[w] = struct{}{}
		}
		r = s
	}

	benchmarkSetResult = r
}

func BenchmarkWordHashSetContains(b *testing.B) {
	var r bool
	for i := 0; i < b.N; i++ {
		for _, w := range words {
			_, ok := filledTrieHashSet[w]
			r = ok
		}
	}

	benchmarkResult = r
}
