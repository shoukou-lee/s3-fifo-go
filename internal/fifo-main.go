package internal

type Main struct {
	q *Queue
}

func NewMain(capacity int) *Main {	
	return &Main{
		NewQueue(capacity),
	}
}

// M uses an algorithm similar to FIFO-Reinsertion but tracks access information using two bits.
func (m *Main) Insert(n *Node) {
	m.q.Enqueue(n)
}

// If M's tail is accessed more than once, it is moved to head decreasing its frequency. Otherwise, it is evicted.
func (m *Main) Evict() {
	for {
		if m.q.Size() > 0 {
			node := m.q.Dequeue()
			if node.Freq() > 1 {
				node.DecFreq()
				m.q.Enqueue(node)
				continue
			}
			return
		}
	}
}

func (m *Main) Full() bool {
	return m.q.Full()
}

// TODO: implement this
func (m *Main) In(n *Node) bool {
	return false
}