package main

import "fmt"

// ArraySize the array size
const ArraySize = 7

// HashTable st
type HashTable struct {
	array [ArraySize]*bucket
}

// Insert insert words
func (h *HashTable) Insert(w string) {
	ind := hash(w)
	h.array[ind].Insert(w)
}

// Search search words
func (h *HashTable) Search(w string) bool {
	ind := hash(w)
	return h.array[ind].Search(w)
}

// Delete delete words
func (h *HashTable) Delete(w string) {
	ind := hash(w)
	h.array[ind].Delete(w)
}

// bucket struc
type bucket struct {
	head *bucketNode
}

type bucketNode struct {
	key  string
	next *bucketNode
}

// Insert insert words
func (b *bucket) Insert(w string) {
	if b.Search(w) {
		fmt.Println("Already exists")
		return
	}

	newNode := &bucketNode{key: w}
	newNode.next = b.head
	b.head = newNode
}

// Search search words
func (b *bucket) Search(w string) bool {
	curr := b.head

	for curr != nil {
		if curr.key == w {
			return true
		}

		curr = curr.next
	}

	return false
}

// Delete delete words
func (b *bucket) Delete(w string) {
	curr := b.head

	if curr.key == w {
		b.head = b.head.next
		return
	}

	for curr.next != nil {
		if curr.next.key == w {
			curr.next = curr.next.next
			return
		}

		curr = curr.next
	}
}

// Hash hash
func hash(w string) int {
	sum := 0
	for _, v := range w {
		sum += int(v)
	}
	return sum % ArraySize
}

// Init init
func Init() *HashTable {
	result := &HashTable{}
	for i := range result.array {
		result.array[i] = &bucket{}
	}

	return result
}

func main() {
	test := Init()

	list := []string{
		"ERIC",
		"KENNY",
		"KYLE",
		"STAN",
		"RANDY",
		"BUTTERS",
		"TOKEN",
	}

	for _, v := range list {
		test.Insert(v)
	}

	fmt.Println(test)
	fmt.Println(test.Search("BUTTERS"))
	test.Delete("BUTTERS")
	fmt.Println(test.Search("BUTTERS"))
	fmt.Println(test.Search("STAN"))
}
