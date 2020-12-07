package nodes

import (
	"fmt"
	gates "go-logic-circuits/gates"
)

// Node is an interface that generalizes all types of Nodes in the program
// A Node is defined as something that can be Evaluated and has an outputz
type Node interface {
	GetOutputChan() chan gates.Bit
	EvaluateNode()
	IncrementOutput()
}

// LogicNode is the main type of Node in the program
// it takes in a variable number of gates.Bit(s), has a function, and has an output
// channel
type LogicNode struct {
	input      []chan gates.Bit
	f          func(...gates.Bit) gates.Bit
	output     chan gates.Bit
	numOutputs int
}

// SourceNode is a special type of node for creating the input gates.Bits into graph
type SourceNode struct {
	BitValue   gates.Bit
	output     chan gates.Bit
	numOutputs int
}

// Connect takes a list of Nodes and connects the channels to a new Node
func Connect(ins []Node, out *LogicNode) *LogicNode {
	if len(out.input) != len(ins) {
		panic(fmt.Sprintf("Got %v input(s) and expected %v input(s)", len(ins), len(out.input)))
	}

	for i := 0; i < len(ins); i++ {
		out.input[i] = ins[i].GetOutputChan()
		ins[i].IncrementOutput()
	}

	return out
}

// NewLogicNode is a constructor for the LogicNode type
func NewLogicNode(inputShape int, f func(...gates.Bit) gates.Bit) LogicNode {
	// validate the input shape
	if inputShape < 1 {
		panic("inputShape must be greater than 1")
	}

	// create a slice of input gates.Bit channels
	inChs := make([]chan gates.Bit, inputShape)
	for i := 0; i < inputShape; i++ {
		inChs[i] = make(chan gates.Bit)
	}

	return LogicNode{
		inChs,
		f,
		make(chan gates.Bit),
		0,
	}
}

// NewSourceNode is a constructor for the SourceNode type
func NewSourceNode() SourceNode {
	return SourceNode{
		gates.Bit(false),
		make(chan gates.Bit),
		0,
	}
}

// GetOutputChan to satisfy the Node interface
func (n *LogicNode) GetOutputChan() chan gates.Bit { return n.output }

// IncrementOutput is used to increment number of outputs node is connected to
func (n *LogicNode) IncrementOutput() { n.numOutputs++ }

// EvaluateNode is function to activate Node
func (n *LogicNode) EvaluateNode() {

	// must recieve all inputs to calculate result
	inputs := []gates.Bit{}
	for _, ch := range n.input {
		inputs = append(inputs, <-ch)
	}
	result := n.f(inputs...)

	// push the result numOutputs times
	for i := 0; i < n.numOutputs; i++ {
		n.output <- result
	}
}

// GetOutputChan to satisfy the Node interface
func (n *SourceNode) GetOutputChan() chan gates.Bit { return n.output }

// IncrementOutput is used to increment number of outputs node is connected to
func (n *SourceNode) IncrementOutput() { n.numOutputs++ }

// EvaluateNode for SourceNode activates Node to push static value out
func (n *SourceNode) EvaluateNode() {
	for i := 0; i < n.numOutputs; i++ {
		n.output <- n.BitValue
	}
}

// Setgates.BitValue sets gates.BitValue for SourceNode
func (n *SourceNode) SetBitValue(b gates.Bit) { n.BitValue = b }

// Getgates.BitValue gets gates.BitValue
func (n *SourceNode) GetBitValue() gates.Bit { return n.BitValue }
