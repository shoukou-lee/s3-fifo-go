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
func (s *Small) Insert(n *Node) {
	if s.Full() {
		node := s.q.Dequeue()
		node.ResetFreq()
		if node.Freq() > 1 {
			s.m.Insert(node)
			return
		}
		s.g.Insert(node)
		return
	}
}

func (s *Small) Evict() {
	node := s.q.Dequeue()

	if node.Freq() == 0 {
		s.g.Insert(node)
		return
	}

	if s.m.Full() {
		s.m.Evict()
	}
	node.ResetFreq()
	s.m.Insert(node)
}

func (s *Small) Full() bool {
	return s.q.Full()
}

// TODO: implement this
func (s *Small) In(n *Node) bool {
	return false
}