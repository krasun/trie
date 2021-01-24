# trie

[![Build Status](https://travis-ci.com/krasun/trie.svg?branch=main)](https://travis-ci.com/krasun/trie)
[![codecov](https://codecov.io/gh/krasun/trie/branch/main/graph/badge.svg?token=rh8BDdHc2v)](https://codecov.io/gh/krasun/trie)
[![Go Report Card](https://goreportcard.com/badge/github.com/krasun/trie)](https://goreportcard.com/report/github.com/krasun/trie)
[![GoDoc](https://godoc.org/https://godoc.org/github.com/krasun/trie?status.svg)](https://godoc.org/github.com/krasun/trie)
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fkrasun%2Ftrie.svg?type=shield)](https://app.fossa.com/projects/git%2Bgithub.com%2Fkrasun%2Ftrie?ref=badge_shield)

A missing [Trie](https://en.wikipedia.org/wiki/Trie) implementation for Go. 

## Install 

```shell
go get github.com/krasun/trie
```

## Usage 

### A rune-wise trie

Without any stress:

```go 
t := trie.New() // or trie.NewRuneTrie()  

t.Insert("apple")
t.Insert("banana")

t.Contains("apple") // true
t.Contains("app") // false
t.Contains("banana") // true
t.Contains("ban") // false
```

### A domain-optimized trie

Yes, it is that simple:

```go 
t := trie.NewDomainTrie()

t.Insert("apple.com")
t.Insert("google.com")

t.Contains("apple.com") // true
t.Contains(".com") // false
t.Contains("mail.google.com") // false
t.Contains("google.com") // true
```

### A goroutine-safe (thread-safe) trie

By default rune-wise and domain-optimized tries are not safe to use 
concurrently, but it is easy to make them safe. You can wrap any trie into the safe version by: 
```go
safeTrie := trie.Safe(trie.NewDomainTrie())

// the same interface as for a regular trie
```

## License 

trie is released under [the MIT license](LICENSE).

