package main

import (
	"fmt"
	"math"
)

// Stack - stack
type Stack struct {
	items []int
}

// Push - push
func (s *Stack) Push(i int) {
	s.items = append(s.items, i)
}

// Pop - pop
func (s *Stack) Pop() int {
	l := len(s.items)

	if l == 0 {
		return math.MinInt32
	}

	toRemove := s.items[l-1]
	s.items = s.items[:l-1]
	return toRemove
}

func main() {
	myStack := Stack{}
	myStack.Push(100)
	myStack.Push(200)
	myStack.Push(300)
	fmt.Println(myStack)

	myStack.Pop()
	fmt.Println(myStack)
}
