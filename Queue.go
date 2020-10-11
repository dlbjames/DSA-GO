package main

import (
	"fmt"
	"math"
)

// Queue - q
type Queue struct {
	items []int
}

// Enqueue - enqueue
func (q *Queue) Enqueue(i int) {
	q.items = append(q.items, i)
}

// Dequeue - dequeue
func (q *Queue) Dequeue() int {
	if len(q.items) == 0 {
		return math.MinInt32
	}

	toRemove := q.items[0]
	q.items = q.items[1:]
	return toRemove
}

func main() {
	myQ := Queue{}
	myQ.Enqueue(100)
	myQ.Enqueue(200)
	myQ.Enqueue(300)
	fmt.Println(myQ)
	fmt.Println(myQ.Dequeue())
	fmt.Println(myQ)
	fmt.Println(myQ.Dequeue())
	fmt.Println(myQ.Dequeue())
	fmt.Println(myQ.Dequeue())
	fmt.Println(myQ)
}
