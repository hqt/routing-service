package routingservice

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewFIFO(t *testing.T) {
	t.Parallel()

	queue := NewFIFO()

	queue.Push(1)
	queue.Push(2)
	queue.Push(3)
	require.Equal(t, 3, queue.Len())

	e := queue.Front()
	require.Equal(t, 1, e)
	e = queue.Front()
	require.Equal(t, 2, e)
	e = queue.Front()
	require.Equal(t, 3, e)

	e = queue.Front()
	require.Nil(t, e)

	queue.Push(4)
	queue.Push(5)
	queue.Push(6)
	require.Equal(t, 3, queue.Len())

	e = queue.Front()
	require.Equal(t, 4, e)
}
