package nodes

import (
	"fmt"
	gates "go-logic-circuits/gates"
	"sync"
)

// Node is an interface that generalizes all types of Nodes in the program
// A Node is defined as something that can be Evaluated and has an outputz
type Node interface {
	OutputChan() chan gates.Bit
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
	bitValue   gates.Bit
	output     chan gates.Bit
	numOutputs int
}

// Connect takes a list of Nodes and connects the channels to a new Node
func Connect(ins []Node, node *LogicNode) *LogicNode {
	if len(node.input) != len(ins) {
		panic(fmt.Sprintf("Got %v input(s) and expected %v input(s)", len(ins), len(node.input)))
	}

	// connect receiving node's input channels to the output channels of the input nodes passed in
	// increment input nodes number of outputs required
	for i := 0; i < len(ins); i++ {
		node.input[i] = ins[i].OutputChan()
		ins[i].IncrementOutput()
	}

	return node
}

// NewLogicNode is a constructor for the LogicNode type
func NewLogicNode(inputShape int, f func(...gates.Bit) gates.Bit) LogicNode {
	// validate the input shape
	if inputShape < 1 {
		panic("inputShape must be greater than 1")
	}

	//create a slice of input gates.Bit channels
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

// OutputChan to satisfy the Node interface
func (n *LogicNode) OutputChan() chan gates.Bit { return n.output }

// NumOutputs to get numOutputs
func (n *LogicNode) NumOutputs() int { return n.numOutputs }

// IncrementOutput is used to increment number of outputs node is connected to
func (n *LogicNode) IncrementOutput() { n.numOutputs++ }

// EvaluateNode is function to activate Node
func (n *LogicNode) EvaluateNode() {

	var wg sync.WaitGroup

	// must recieve all inputs to calculate result
	inputs := make([]gates.Bit, len(n.input))

	// for all input channels wait for inputs
	for i, ch := range n.input {
		wg.Add(1)
		go func(i int, ch chan gates.Bit) {
			defer wg.Done()
			inputs[i] = <-ch
		}(i, ch)
	}

	// wait for all inputs and calculate results
	wg.Wait()
	result := n.f(inputs...)

	// push the result out numOutputs times
	for i := 0; i < n.numOutputs; i++ {
		n.output <- result
	}
}

// OutputChan to satisfy the Node interface
func (n *SourceNode) OutputChan() chan gates.Bit { return n.output }

// NumOutputs to get numOutputs
func (n *SourceNode) NumOutputs() int { return n.numOutputs }

// IncrementOutput is used to increment number of outputs node is connected to
func (n *SourceNode) IncrementOutput() { n.numOutputs++ }

// EvaluateNode for SourceNode activates Node to push static value out
func (n *SourceNode) EvaluateNode() {
	for i := 0; i < n.numOutputs; i++ {
		go func(n *SourceNode) {
			n.output <- n.bitValue
		}(n)
	}
}

// SetBitValue sets gates.BitValue for SourceNode
func (n *SourceNode) SetBitValue(b gates.Bit) { n.bitValue = b }

// BitValue gets gates.BitValue
func (n *SourceNode) BitValue() gates.Bit { return n.bitValue }
