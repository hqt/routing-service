package datastructure

import "container/heap"

// HeapNode represents the node for the priority queue
type HeapNode struct {
	Key    string
	Weight int
}

// NewHeapNode creates HeapNode object
func NewHeapNode(key string, val int) *HeapNode {
	return &HeapNode{
		Key:    key,
		Weight: val,
	}
}

// A heapArr implements heap.Interface and holds Items.
type heapArr []*HeapNode

func (pq heapArr) Len() int { return len(pq) }

func (pq heapArr) Less(i, j int) bool {
	return pq[i].Weight < pq[j].Weight
}
func (pq heapArr) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *heapArr) Push(x interface{}) {
	item := x.(*HeapNode)
	*pq = append(*pq, item)
}

func (pq *heapArr) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil // avoid memory leak
	*pq = old[0 : n-1]
	return item
}

// PriorityQueue represents the priority queue
type PriorityQueue struct {
	arr heapArr
}

// NewPriorityQueue returns a new priority queue
func NewPriorityQueue() *PriorityQueue {
	arr := heapArr{}
	heap.Init(&arr)
	return &PriorityQueue{
		arr: arr,
	}
}

// Len returns size of this queue
func (pq *PriorityQueue) Len() int {
	return len(pq.arr)
}

// Push pushes a node to queue
func (pq *PriorityQueue) Push(node *HeapNode) {
	heap.Push(&pq.arr, node)
}

// Pop pops a node out of queue
func (pq *PriorityQueue) Pop() *HeapNode {
	return heap.Pop(&pq.arr).(*HeapNode)
}
