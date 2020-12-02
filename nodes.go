package main

import "fmt"

// Node is an interface that generalizes all types of Nodes in the program
type Node interface {
	getOutputChan() chan Bit
}

// LogicNode is the main type of Node in the program
// it takes in a variable number of Bit(s), has a function, and has an output
// channel
type LogicNode struct {
	input  []chan Bit
	f      func(...Bit) Bit
	output chan Bit
}

// SourceNode is a special type of node for creating the input Bits into graph
type SourceNode struct {
	output chan Bit
}

// Connect takes a list of Nodes and connects the channels to a new Node
func Connect(ins []Node, out *LogicNode) *LogicNode {
	if len(out.input) != len(ins) {
		panic("Mismatch expected inputs")
	}
	for i := 0; i < len(ins); i++ {
		out.input[i] = ins[i].getOutputChan()
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
	}
}

// getOutputChan to satisfy the Node interface
func (n LogicNode) getOutputChan() chan Bit { return n.output }

// EvaluateNode is
func (n *LogicNode) EvaluateNode() {
	inputs := []Bit{}
	for _, ch := range n.input {
		inputs = append(inputs, <-ch)
	}
	fmt.Println(inputs)
	n.output <- n.f(inputs...)
}

// NewSourceNode is a constructor for the SourceNode type
func NewSourceNode() SourceNode {
	return SourceNode{
		make(chan Bit),
	}
}

// getOutputChan to satisfy the Node interface
func (n SourceNode) getOutputChan() chan Bit { return n.output }
