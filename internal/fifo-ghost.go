package internal

type Ghost struct {
	q         *Queue
	m         *Main
	hashTable map[string]bool // To evaluate if the key to be inserted exists in the ghost queue
}

func NewGhost(capacity int, main *Main) *Ghost {
	return &Ghost{
		q: NewQueue(capacity),
		m: main,
		hashTable: make(map[string]bool),
	}
}

func (g *Ghost) In(n *Node) bool {
	_, ok := g.hashTable[n.Key()]
	return ok
}

// When G is full, it evicts objects in FIFO order.
func (g *Ghost) Insert(n *Node) {
	if g.Full() {
		g.evict()
	}
	n.ResetFreq()
	g.q.Enqueue(n)
	g.hashTable[n.Key()] = true
}

func (g *Ghost) evict() {
	node := g.q.Dequeue()
	delete(g.hashTable, node.Key())
}

func (g *Ghost) Full() bool {
	return g.q.Full()
}