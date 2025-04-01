package internal

import "fmt"

type S3Fifo struct {
	ghost *Ghost
	small *Small
	main *Main
	maxCapacity int
	smallCapacity int
	hashTable map[string]*Node
}

func NewS3Fifo(maxCap int) *S3Fifo {
	if maxCap <= 0 {
		panic("totalCapacity must be greater than 0")
	}

	// We choose S to use 10% of the cache space based on experiments
	// with 10 traces and find that 10% generalizes well. 
	// M then uses 90% of the cache space. 
	// The ghost queue G stores the same number of ghost entries (no data) as M
	smallCap := maxCap / 10
	mainCap := maxCap - smallCap

	main := NewMain(mainCap)
	ghost := NewGhost(mainCap, main)
	small := NewSmall(smallCap, main, ghost)

	return &S3Fifo{
		ghost: ghost,
		small: small,
		main:  main,
		maxCapacity: maxCap,
		smallCapacity: smallCap,
		hashTable: make(map[string]*Node),
	}
}

func (s *S3Fifo) GetOrElsePut(key string, value interface{}) interface{} {
	if found := s.Get(key); found != nil {
		return found
	}
	s.Put(key, value)
	return value
}

// Get retrieves the value associated with the given key from the cache if exists, otherwise nil.
func (s *S3Fifo) Get(key string) interface{} {
	if found, ok := s.hashTable[key]; ok {
		found.IncFreq()
		return found.Value()
	}
	return nil
}

// Put inserts a new key-value pair into the cache.
func (s *S3Fifo) Put(key string, value interface{}) {
	n := NewNode(key, value)
	s.insert(n)
}

// New objects are inserted into S if not in G. Otherwise, it is inserted into M.
func (s *S3Fifo) insert(n *Node) {
	var evicted *Node = nil

	if s.ghost.In(n) {
		evicted = s.main.Insert(n)
	} else {
		evicted = s.small.Insert(n)
	}

	if evicted != nil {
		delete(s.hashTable, evicted.Key())
	}
	s.hashTable[n.Key()] = n
}

func (s *S3Fifo) Log() {
	fmt.Printf("small: %d | ", s.small.q.Size())
	for i := len(s.small.q.items) - 1; i >= 0; i-- {
		n := s.small.q.items[i]
		fmt.Printf("%s(%d) ", n.Key(), n.Freq())
	}
	fmt.Println()

	fmt.Printf("main: %d | ", s.main.q.Size())
	for i := len(s.main.q.items) - 1; i >= 0; i-- {
		n := s.main.q.items[i]
		fmt.Printf("%s(%d) ", n.Key(), n.Freq())
	}
	fmt.Println()

	fmt.Printf("ghost: %d | ", s.ghost.q.Size())
	for i := len(s.ghost.q.items) - 1; i >= 0; i-- {
		n := s.ghost.q.items[i]
		fmt.Printf("%s(%d) ", n.Key(), n.Freq())
	}
	fmt.Println()
	fmt.Printf("==========")
	fmt.Println()
}