package main

import "fmt"

// MaxHeap max
type MaxHeap struct {
	array []int
}

// Insert insert
func (h *MaxHeap) Insert(key int) {
	h.array = append(h.array, key)
	h.maxHeapifyUp(len(h.array) - 1)
}

// Extract extract
func (h *MaxHeap) Extract() int {
	ex := h.array[0]
	l := len(h.array)

	if l == 0 {
		fmt.Println("No elements")
		return -1
	}

	h.array[0] = h.array[l-1]
	h.array = h.array[:l-1]

	h.maxHeapifyDown(0)
	return ex
}

func (h *MaxHeap) maxHeapifyUp(ind int) {
	for h.array[parent(ind)] < h.array[ind] {
		h.swap(parent(ind), ind)
		ind = parent(ind)
	}
}

func (h *MaxHeap) maxHeapifyDown(ind int) {
	last := len(h.array) - 1
	l, r := left(ind), right(ind)
	childToComp := 0

	for l <= last {
		if l == last || h.array[l] > h.array[r] {
			childToComp = l
		} else {
			childToComp = r
		}

		if h.array[ind] < h.array[childToComp] {
			h.swap(ind, childToComp)
			ind = childToComp
			l, r = left(ind), right(ind)
		} else {
			return
		}
	}
}

// get the parent index
func parent(i int) int {
	return (i - 1) / 2
}

// left child
func left(i int) int {
	return i*2 + 1
}

// right child
func right(i int) int {
	return i*2 + 2
}

// Swaps 2 nodes
func (h *MaxHeap) swap(i1, i2 int) {
	h.array[i1], h.array[i2] = h.array[i2], h.array[i1]
}

func main() {
	fmt.Println("yeet")
	m := &MaxHeap{}
	fmt.Println(m)

	build := []int{10, 20, 30, 1, 2, 3, 100, 200, 300}
	for _, v := range build {
		m.Insert(v)
	}

	fmt.Println(m)

	fmt.Println(m.Extract())
	fmt.Println(m)
}
