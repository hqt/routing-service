package datastructure

// Queue interface for queue
type Queue interface {
	Push(node interface{})
	Front() interface{}
	Len() int
}

// FIFO is a FIFO queue
type FIFO struct {
	queue []interface{}
}

// NewFIFO creates new FIFO and returns it
func NewFIFO() *FIFO {
	return &FIFO{
		queue: make([]interface{}, 0),
	}
}

// Push pushed Node to the back of the queue
func (f *FIFO) Push(node interface{}) {
	f.queue = append(f.queue, node)
}

// Front takes a value from the front of the queue and returns it
func (f *FIFO) Front() interface{} {
	if len(f.queue) == 0 {
		return nil
	}

	node := f.queue[0]
	f.queue[0] = nil
	f.queue = f.queue[1:]

	return node
}

// Len returns total elements of FIFO
func (f *FIFO) Len() int {
	return len(f.queue)
}
