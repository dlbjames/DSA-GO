package main

import "fmt"

var count int

// Node node
type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

// Insert add node
func (n *Node) Insert(k int) {
	if n.Key < k {
		if n.Right == nil {
			n.Right = &Node{Key: k}
		} else {
			n.Right.Insert(k)
		}
	} else {
		if n.Left == nil {
			n.Left = &Node{Key: k}
		} else {
			n.Left.Insert(k)
		}
	}
}

// Search for the element
func (n *Node) Search(k int) bool {
	count++

	if n == nil {
		return false
	}

	if n.Key < k {
		return n.Right.Search(k)
	} else if n.Key > k {

		return n.Left.Search(k)
	}

	return true
}

func main() {
	tree := &Node{Key: 100}
	tree.Insert(50)
	tree.Insert(150)
	tree.Insert(10)
	tree.Insert(520)
	tree.Insert(151)
	tree.Insert(5)
	tree.Insert(4)
	tree.Insert(11)
	tree.Insert(52)
	fmt.Println(tree)
	fmt.Println(tree.Search(151))
	fmt.Println(count)
}
