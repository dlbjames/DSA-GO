package main

import "fmt"

// ALPHABETSIZE The character in trie
const ALPHABETSIZE = 26

// Node node
type Node struct {
	children [ALPHABETSIZE]*Node
	isEnd    bool
}

//Trie trie
type Trie struct {
	root *Node
}

// InitTrie wil create a Trie
func InitTrie() *Trie {
	result := &Trie{root: &Node{}}
	return result
}

//Insert insert
func (t *Trie) Insert(w string) {
	len := len(w)
	curr := t.root
	for i := 0; i < len; i++ {
		charInd := w[i] - 'a'
		if curr.children[charInd] == nil {
			curr.children[charInd] = &Node{}
		}
		curr = curr.children[charInd]
	}

	curr.isEnd = true
}

//Search search
func (t *Trie) Search(w string) bool {
	len := len(w)
	curr := t.root

	for i := 0; i < len; i++ {
		charInd := w[i] - 'a'
		if curr.children[charInd] == nil {
			return false
		}

		curr = curr.children[charInd]
	}

	return curr.isEnd
}

func main() {
	testTrie := InitTrie()
	testTrie.Insert("hello")
	testTrie.Insert("halo")
	testTrie.Insert("hawaii")
	testTrie.Insert("hungry")

	fmt.Println(testTrie.root)
	fmt.Println(testTrie.Search("hawaii"))
	fmt.Println(testTrie.Search("hawai"))
}
