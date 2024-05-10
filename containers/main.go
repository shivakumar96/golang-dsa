package main

import (
	"container/heap"
	"container/list"
	"fmt"
	"sort"
)

func listimp() {
	l := list.New()
	l.Init()
	l.PushBack(2)
	l.PushBack(3)
	l.PushBack(4)
	l.PushBack(5)
	l.PushFront(1)

	for curr := l.Front(); curr != nil; curr = curr.Next() {
		fmt.Println(curr.Value.(int) == 2)
	}

}

type myHeap []int

func (h myHeap) Len() int {
	return len(h)
}

func (h myHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h myHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h *myHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *myHeap) Pop() any {
	old := *h
	n := len(*h)
	*h = (*h)[:n-1]
	return old[n-1]

}

func heapImp() {
	h := &myHeap{}
	heap.Init(h)
	heap.Push(h, 6)
	heap.Push(h, 3)
	heap.Push(h, 5)
	val := heap.Pop(h)
	fmt.Println(h)
	fmt.Println(val == 3)
}

func customSort() {
}

func main() {
	//listimp()
	//heapImp()

	val := []int{3, 2, 1}
	var temp myHeap = val
	sort.Sort(temp)
	fmt.Println(temp)
}
