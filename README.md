# trie

[![Build Status]](https://github.com/<OWNER>/<REPOSITORY>/actions/workflows/build/badge.svg)
[![codecov](https://codecov.io/gh/krasun/trie/branch/main/graph/badge.svg?token=rh8BDdHc2v)](https://codecov.io/gh/krasun/trie)
[![Go Report Card](https://goreportcard.com/badge/github.com/krasun/trie)](https://goreportcard.com/report/github.com/krasun/trie)
[![GoDoc](https://godoc.org/https://godoc.org/github.com/krasun/trie?status.svg)](https://godoc.org/github.com/krasun/trie)

A missing [trie](https://en.wikipedia.org/wiki/Trie) implementation for Go. 

## Installation

```shell
go get github.com/krasun/trie
```

## Usage 

Known limitations: 
1. Empty string keys are **not supported**. And functions won't return an error. 
2. Keys are not ordered, so do not expect any ordering. 

Without any stress:

```go 
t := trie.New() // or trie.NewRuneTrie()  

t.Insert("apple")
t.Insert("alphabet")
t.Insert("banana")

t.Contains("apple") // true
t.Contains("app") // false
t.Contains("banana") // true
t.Contains("ban") // false

t.SearchByPrefix("a") // []string{"apple", "alphabet"}

t.StartsWith("a") // true
```

### A goroutine-safe (thread-safe) trie

By default, rune-wise trie are not safe to use concurrently, 
but it is easy to make them safe. You can wrap any trie into the safe version by: 
```go
safeTrie := trie.Safe(trie.NewTrie())

// the same interface as for a regular trie
```

## Tests 

To run tests, just execute: 
```
$ go test . -race
```

## Benchmarking 

Zero heap allocations for `Contains` function:
```
$ go test -bench=. -benchmem -benchtime=1000x
goos: darwin
goarch: amd64
op	       0 allocs/op
BenchmarkRuneTrieInsert-8          	    1000	      7196 ns/op	    6984 B/op	     129 allocs/op
BenchmarkRuneTrieContains-8        	    1000	       517 ns/op	       0 B/op	       0 allocs/op
BenchmarkWordHashSetInsert-8       	    1000	      1406 ns/op	    1100 B/op	       4 allocs/op
BenchmarkWordHashSetContains-8     	    1000	       178 ns/op	       0 B/op	       0 allocs/op
PASS
```

But the hash set (map[string]{}struct) is much much more efficient than the trie. Carefully 
weigh when to use the trie.

## Applications

Use the trie in cases when it only outperforms hash tables or hash sets (map[string]{}struct):
1. Search key by the prefix. 

## License 

trie is released under [the MIT license](LICENSE).

