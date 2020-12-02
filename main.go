package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	// set random seed for pushing random values onto the sources
	// for testing primarily
	rand.Seed(time.Now().UnixNano())

	// will need to switch to map implementations (unique variable stuff for when text is parsed)

	inputNodes := []Node{
		NewSourceNode(),
		NewSourceNode(),
		NewSourceNode(),
		NewSourceNode(),
	}

	logicNodes := []Node{}
	outputNodes := []Node{}

	var newNode LogicNode

	// this is pattern for putting node in graph
	newNode = NewLogicNode(2, And)
	logicNodes = append(logicNodes, Connect(inputNodes[:2], &newNode))

	newNode = NewLogicNode(2, And)
	logicNodes = append(logicNodes, Connect(inputNodes[2:], &newNode))

	newNode = NewLogicNode(2, Or)
	logicNodes = append(logicNodes, Connect(logicNodes[:2], &newNode))
	outputNodes = append(outputNodes, &newNode)

	// push Bit onto all the sources
	for _, node := range inputNodes {
		go func(n Node) {
			n.getOutputChan() <- Bit(rand.Intn(2) == 1)
		}(node)
	}

	// Start threads waiting to evaluate
	for _, node := range logicNodes {
		go node.EvaluateNode()
	}

	// collect outputs
	for _, node := range outputNodes {
		fmt.Println(<-node.getOutputChan())
	}

}
