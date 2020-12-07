package circuit

import (
	"go-logic-circuits/gates"
	nodes "go-logic-circuits/nodes"
	"sync"
)

// Evaluate the Circuit object on a set of inputs
func (c *Circuit) Evaluate(ins map[string]gates.Bit) map[string]gates.Bit {

	// synchronously call set BitValue for every InputNode
	for s, node := range c.InputNodes {
		node.(*nodes.SourceNode).SetBitValue(ins[s])
	}

	// create WaitGroup
	var wg sync.WaitGroup

	// push Bit onto all the sources
	for _, node := range c.InputNodes {
		node.EvaluateNode() // not a blocking call
	}

	// Start threads waiting to evaluate
	for _, node := range c.CircuitNodes {
		node.EvaluateNode() // blocks as necessary
	}

	// collect the outputs
	var outVals = sync.Map{}

	for s, node := range c.OutputNodes {
		wg.Add(1)
		go func(s string, node nodes.Node) {
			defer wg.Done()
			outVals.Store(s, <-node.GetOutputChan())
		}(s, node)
	}

	// wait on all output channels
	wg.Wait()

	tmp := make(map[string]gates.Bit)
	outVals.Range(
		func(key interface{}, value interface{}) bool {
			tmp[key.(string)] = value.(gates.Bit)
			return true
		},
	)
	return tmp
}
