package internal

type Small struct {
	q *Queue
	m *Main
	g *Ghost
}

func NewSmall(capacity int, main *Main, ghost *Ghost) *Small {
	return &Small{
		q: NewQueue(capacity),
		m: main,
		g: ghost,
	}
}

// When S is full, the object at the tail is either moved to M if it is accessed more than once or G if not.
// And its access bits are cleared during the move.
func (s *Small) Insert(n *Node) *Node {
	var evicted *Node = nil
	if s.Full() {
		evicted = s.Evict()
	}
	s.q.Enqueue(n)
	return evicted
}

func (s *Small) Evict() *Node {
	node := s.q.Dequeue()

	if node.Freq() == 0 {
		s.g.Insert(node)
		return node
	}
	node.ResetFreq()
	s.m.Insert(node)
	return nil
}

func (s *Small) Full() bool {
	return s.q.Full()
}
