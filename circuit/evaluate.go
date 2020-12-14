package circuit

import (
	"go-logic-circuits/gates"
	nodes "go-logic-circuits/nodes"
	"sync"
)

// Evaluate the Circuit object on a set of inputs
func (c *Circuit) Evaluate(ins map[string]gates.Bit) map[string]gates.Bit {

	// set bitValue for every InputNode
	for s, node := range c.InputNodes {
		node.(*nodes.SourceNode).SetBitValue(ins[s])
	}

	// push Bit onto all the source output channels
	for _, node := range c.InputNodes {
		go node.EvaluateNode()
	}

	// Start threads waiting to evaluate inputs
	for _, node := range c.CircuitNodes {
		go node.EvaluateNode()
	}

	// collect the outputs
	var outVals = sync.Map{}

	var wg sync.WaitGroup
	for s, node := range c.OutputNodes {
		wg.Add(1)
		go func(s string, node nodes.Node) {
			defer wg.Done()
			outVals.Store(s, <-node.OutputChan())
		}(s, node)
	}
	wg.Wait()

	// put outputs into normal map
	tmp := make(map[string]gates.Bit)
	outVals.Range(
		func(key interface{}, value interface{}) bool {
			tmp[key.(string)] = value.(gates.Bit)
			return true
		},
	)
	return tmp
}
