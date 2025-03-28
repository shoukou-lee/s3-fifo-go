package internal

type Main struct {
	q *Queue
}

func NewMain(capacity int) *Main {	
	return &Main{
		NewQueue(capacity),
	}
}

// M uses an algorithm similar to FIFO-Reinsertion but tracks access count
func (m *Main) Insert(n *Node) *Node {
	var evicted *Node = nil
	if m.q.Full() {
		evicted = m.Evict()
	}
	m.q.Enqueue(n)
	return evicted
}

// If M's tail is accessed more than once, it is moved to head decreasing its frequency. Otherwise, it is evicted.
func (m *Main) Evict() *Node {
	for {
		if m.q.Size() > 0 {
			node := m.q.Dequeue()
			if node.Freq() > 1 {
				node.DecFreq()
				m.q.Enqueue(node)
				continue
			}
			return node
		}
		return nil
	}
}

func (m *Main) Full() bool {
	return m.q.Full()
}
