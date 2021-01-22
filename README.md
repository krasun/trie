# trie
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

t.Exists("apple") // true
t.Exists("app") // false
t.Exists("banana") // true
t.Exists("ban") // false
```

### A domain-optimized trie

Yes, it is that simple:

```go 
t := trie.NewDomainTrie()

t.Insert("apple.com")
t.Insert("google.com")

t.Exists("apple.com") // true
t.Exists(".com") // false
t.Exists("mail.google.com") // false
t.Exists("google.com") // true
```

## License 

trie is released under [the MIT license](LICENSE).



[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fkrasun%2Ftrie.svg?type=large)](https://app.fossa.com/projects/git%2Bgithub.com%2Fkrasun%2Ftrie?ref=badge_large)