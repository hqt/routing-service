package datastructure

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestIntHeap(t *testing.T) {
	t.Parallel()

	pq := NewPriorityQueue()

	pq.Push(&HeapNode{Key: "4", Weight: 4})
	pq.Push(&HeapNode{Key: "3", Weight: 3})
	pq.Push(&HeapNode{Key: "2", Weight: 2})
	pq.Push(&HeapNode{Key: "1", Weight: 1})

	node := pq.Pop()
	require.Equal(t, 1, node.Weight)

	node = pq.Pop()
	require.Equal(t, 2, node.Weight)

	pq.Push(&HeapNode{Key: "0", Weight: 0})
	node = pq.Pop()
	require.Equal(t, 0, node.Weight)
}
