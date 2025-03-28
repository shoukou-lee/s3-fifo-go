package internal

// Node represents a data that 
type Node struct {
	key string
	value interface{}
	freq int
}

func NewNode(k string, v interface{}) *Node {
	return &Node{
		key: k,
		value: v,
		freq: 0,
	}
}

func (n *Node) IncFreq() {
	if n.freq >= 3 {
		return
	}
	n.freq++
}

func (n *Node) DecFreq() {
	if n.freq <= 0 {
		return
	}
	n.freq--
}

func (n *Node) Key() string {
	return n.key
}

func (n *Node) Value() interface{} {
	return n.value
}

func (n *Node) Freq() int {
	return n.freq
}

func (n *Node) ResetFreq() {
	n.freq = 0
}