package internal

// Node represents a data that 
type Node struct {
	data interface{}
	freq int
}

func NewNode(data interface{}) *Node {
	return &Node{
		data: data,
		freq: 0,
	}
}

func (n *Node) IncFreq() {
	n.freq++
}

func (n *Node) DecFreq() {
	n.freq--
}

func (n *Node) Data() interface{} {
	return n.data
}

func (n *Node) Freq() int {
	return n.freq
}

func (n *Node) ResetFreq() {
	n.freq = 0
}