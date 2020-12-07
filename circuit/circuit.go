package circuit

import (
	nodes "go-logic-circuits/nodes"
)

// Circuit is type to manage 3 main sets of Nodes
// InputNodes, LogicNodes, and OutputNodes
type Circuit struct {
	InputNodes   map[string]nodes.Node
	CircuitNodes map[string]nodes.Node
	OutputNodes  map[string]nodes.Node
	Log          []string
}

// NewCircuit will make a new circuit
func NewCircuit() *Circuit {
	return &Circuit{
		make(map[string]nodes.Node),
		make(map[string]nodes.Node),
		make(map[string]nodes.Node),
		[]string{},
	}
}

// DuplicateCircuit will duplicate and make a new circuit
// this utilizes the Log so you must use the ParseString function
// and not manually manipulate the Circuit
func DuplicateCircuit(input *Circuit) *Circuit {
	ret := NewCircuit()
	for _, s := range input.Log {
		ret.ParseString(s)
	}
	return ret
}

func (c *Circuit) validNewInput(input string) bool {
	if _, ok := c.InputNodes[input]; ok {
		return false
	}
	return true
}

func (c *Circuit) validNewNode(input string) bool {
	if _, ok := c.CircuitNodes[input]; ok {
		return false
	}
	return true
}

func (c *Circuit) validNewOutput(input string) bool {
	// Is already an OutputNode
	if _, ok := c.OutputNodes[input]; ok {
		return false
	}
	// Has to be an existing Node
	if _, ok := c.CircuitNodes[input]; ok {
		return true
	}
	if _, ok := c.InputNodes[input]; ok {
		return true
	}
	// Not an existing Node in Circuit
	return false
}
