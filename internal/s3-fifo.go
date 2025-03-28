package internal

type S3Fifo struct {
	ghost *Ghost
	small *Small
	main *Main
	maxCapacity int
	smallCapacity int
}

func NewS3Fifo(maxCap int) *S3Fifo {
	if maxCap <= 0 {
		panic("totalCapacity must be greater than 0")
	}

	smallCap := maxCap / 10
	mainCap := maxCap - smallCap

	main := NewMain(mainCap)
	ghost := NewGhost(maxCap, main)
	small := NewSmall(smallCap, main, ghost)

	return &S3Fifo{
		ghost: ghost,
		small: small,
		main:  main,
		maxCapacity: maxCap,
		smallCapacity: smallCap,
	}
}

// todo: arg n should be another type
func (s *S3Fifo) Read(n *Node) *Node {
	// If n in S or M, increase n's freq and return it 
	if s.small.In(n) || s.main.In(n) {
		n.IncFreq()
		return n
	}

	// Otherwise, insert it
	s.Insert(n)
	return n
}

// New objects are inserted into S if not in G. Otherwise, it is inserted into M.
func (s *S3Fifo) Insert(n *Node) {
	// TODO: ambigous description: while cache is full, do evict
	if s.small.Full() && s.main.Full() {
		s.Evict()
	}

	if s.ghost.In(n) {
		s.main.Insert(n)
		return
	}
	s.small.Insert(n)
}

func (s *S3Fifo) Evict() {
	if s.small.q.Size() > s.smallCapacity {
		s.small.Evict()
		return
	}
	s.main.Evict()
}