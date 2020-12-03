package main

// Node is an interface that generalizes all types of Nodes in the program
// A Node is defined as something that can be Evaluated and has an output
type Node interface {
	getOutputChan() chan Bit
	EvaluateNode()
	incrementOutput()
}

// LogicNode is the main type of Node in the program
// it takes in a variable number of Bit(s), has a function, and has an output
// channel
type LogicNode struct {
	input      []chan Bit
	f          func(...Bit) Bit
	output     chan Bit
	numOutputs int
}

// SourceNode is a special type of node for creating the input Bits into graph
type SourceNode struct {
	bitValue   Bit
	output     chan Bit
	numOutputs int
}

// Connect takes a list of Nodes and connects the channels to a new Node
func Connect(ins []Node, out *LogicNode) *LogicNode {
	if len(out.input) != len(ins) {
		panic("Mismatching # of expected inputs and input channels")
	}

	for i := 0; i < len(ins); i++ {
		out.input[i] = ins[i].getOutputChan()
		ins[i].incrementOutput()
	}

	return out
}

// NewLogicNode is a constructor for the LogicNode type
func NewLogicNode(inputShape int, f func(...Bit) Bit) LogicNode {
	// validate the input shape
	if inputShape < 1 {
		panic("inputShape must be greater than 1")
	}

	// create a slice of input Bit channels
	inChs := make([]chan Bit, inputShape)
	for i := 0; i < inputShape; i++ {
		inChs[i] = make(chan Bit)
	}

	return LogicNode{
		inChs,
		f,
		make(chan Bit),
		0,
	}
}

// NewSourceNode is a constructor for the SourceNode type
func NewSourceNode() SourceNode {
	return SourceNode{
		Bit(false),
		make(chan Bit),
		0,
	}
}

// getOutputChan to satisfy the Node interface
func (n *LogicNode) getOutputChan() chan Bit { return n.output }

// incrementOutput is used to increment number of outputs node is connected to
func (n *LogicNode) incrementOutput() { n.numOutputs++ }

// EvaluateNode is function to activate Node
func (n *LogicNode) EvaluateNode() {

	// must recieve all inputs to calculate result
	inputs := []Bit{}
	for _, ch := range n.input {
		inputs = append(inputs, <-ch)
	}
	result := n.f(inputs...)

	// push the result numOutputs times
	for i := 0; i < n.numOutputs; i++ {
		n.output <- result
	}
}

// getOutputChan to satisfy the Node interface
func (n *SourceNode) getOutputChan() chan Bit { return n.output }

// incrementOutput is used to increment number of outputs node is connected to
func (n *SourceNode) incrementOutput() { n.numOutputs++ }

// EvaluateNode for SourceNode activates Node to push static value out
func (n *SourceNode) EvaluateNode() {
	for i := 0; i < n.numOutputs; i++ {
		n.output <- n.bitValue
	}
}

// SetBitValue sets bitValue for SourceNode
func (n *SourceNode) SetBitValue(b Bit) { n.bitValue = b }
