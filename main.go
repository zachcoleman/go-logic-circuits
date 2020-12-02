package main

import "fmt"

func main() {

	a := []Node{
		NewSourceNode(),
		NewSourceNode(),
		NewSourceNode(),
	}

	newNode := NewLogicNode(2, And)
	b := []*LogicNode{
		Connect(a, &newNode),
	}

	// push all the sources
	for _, node := range a {
		go func(n Node) {
			n.getOutputChan() <- Bit(true)
		}(node)
	}

	// Start threads waiting to evaluate
	for _, node := range b {
		go node.EvaluateNode()
	}

	// collect outputs
	for _, node := range b {
		fmt.Println(<-node.getOutputChan())
	}

}
