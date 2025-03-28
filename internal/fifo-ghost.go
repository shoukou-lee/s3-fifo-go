package internal

type Ghost struct {
	q *Queue
	m *Main
}

func NewGhost(capacity int, main *Main) *Ghost {
	return &Ghost{
		q: NewQueue(capacity),
		m: main,
	}
}

// 
func (g *Ghost) In(n *Node) bool {
	return false
}

// When G is full, it evicts objects in FIFO order.
func (g *Ghost) Insert(n *Node) {
	if g.Full() {
		
	}
	n.ResetFreq()
	g.q.Enqueue(n)
}

func (g *Ghost) Evict() {
	g.q.Dequeue()
}

func (g *Ghost) Full() bool {
	return g.q.Full()
}