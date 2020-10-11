package main

import "fmt"

type node struct {
	data int
	next *node
}

// LinkedList - The linkedlist
type LinkedList struct {
	head   *node
	length int
}

func (l *LinkedList) prepend(n *node) {
	second := l.head
	l.head = n
	l.head.next = second
	l.length++
}

func (l LinkedList) printData() {
	len := l.length
	temp := l.head
	for len != 0 {
		fmt.Printf("%d ", temp.data)
		temp = temp.next
		len--
	}
	fmt.Println()
}

func (l *LinkedList) remove(value int) {
	temp := l.head

	if l.head.data == value {
		l.head = l.head.next
		l.length--
		return
	}

	for temp.next.data != value {
		if temp.next.next == nil {
			return
		}

		temp = temp.next
	}

	temp.next = temp.next.next
	l.length--
}

func main() {
	myList := LinkedList{}
	node1 := &node{data: 40}
	node2 := &node{data: 30}
	node3 := &node{data: 20}
	node4 := &node{data: 10}
	node5 := &node{data: 3}
	node6 := &node{data: 2}
	node7 := &node{data: 4}
	node8 := &node{data: 12}
	node9 := &node{data: 19}
	myList.prepend(node1)
	myList.prepend(node2)
	myList.prepend(node3)
	myList.prepend(node4)
	myList.prepend(node5)
	myList.prepend(node6)
	myList.prepend(node7)
	myList.prepend(node8)
	myList.prepend(node9)
	myList.printData()
	myList.remove(10)
	myList.remove(40)
	myList.remove(19)
	myList.remove(100)
	myList.printData()
}
