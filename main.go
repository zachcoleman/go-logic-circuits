package main

import "fmt"

func main() {

	a := []Node{
		NewSourceNode(),
		NewSourceNode(),
	}

	b := []*LogicNode{
		Connect(a, NewLogicNode(2, And)),
	}

	// want to implement output nodes

	// push all the sources
	for _, node := range a {
		node.getOutputChan() <- Bit(true)
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
