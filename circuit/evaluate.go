package circuit

import (
	"go-logic-circuits/gates"
	nodes "go-logic-circuits/nodes"
)

// Evaluate the Circuit object on a set of inputs
func (c *Circuit) Evaluate(ins map[string]gates.Bit) map[string]gates.Bit {
	// Set all output nodes to push
	for _, node := range c.OutputNodes {
		node.IncrementOutput()
	}

	// push Bit onto all the sources
	for s, node := range c.InputNodes {
		go func(s string, n nodes.Node) {
			n.(*nodes.SourceNode).SetBitValue(ins[s])
			n.EvaluateNode()
		}(s, node)
	}

	// Start threads waiting to evaluate
	for _, node := range c.CircuitNodes {
		go node.EvaluateNode()
	}

	// collect outputs
	outVals := make(map[string]gates.Bit)
	for s, node := range c.OutputNodes {
		outVals[s] = <-node.GetOutputChan()
	}

	return outVals
}
