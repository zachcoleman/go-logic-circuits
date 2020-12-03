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

	// will want to switch to map implementations (unique variable stuff for when text is parsed)

	inputNodes := []Node{}
	for i := 0; i < 4; i++ {
		tmp := NewSourceNode()
		inputNodes = append(inputNodes, &tmp)
	}

	// slices to hold nodes
	logicNodes := []Node{}
	outputNodes := []Node{}

	// this is pattern for putting node in graph
	// can not use same variable since need new addresses
	tmp1 := NewLogicNode(2, And)
	logicNodes = append(logicNodes, Connect(inputNodes[:2], &tmp1))

	tmp2 := NewLogicNode(2, Or)
	logicNodes = append(logicNodes, Connect(inputNodes[2:], &tmp2))

	tmp3 := NewLogicNode(2, Xor)
	logicNodes = append(logicNodes, Connect(logicNodes[:2], &tmp3))
	outputNodes = append(outputNodes, &tmp3)

	// Set all output nodes to push
	for _, node := range outputNodes {
		node.incrementOutput()
	}

	// push Bit onto all the sources
	// right now all sources can only be used once!!!
	for _, node := range inputNodes {
		go func(n Node) {
			n.(*SourceNode).SetBitValue(Bit(rand.Intn(2) == 1)) // set random Bit
			n.EvaluateNode()                                    // Start pushing them out
		}(node)
	}

	// Start threads waiting to evaluate
	for _, node := range logicNodes {
		go node.EvaluateNode()
	}

	// print inputs
	inVals := []Bit{}
	for _, node := range inputNodes {
		inVals = append(inVals, node.(*SourceNode).bitValue)
	}
	fmt.Println(inVals)

	// collect outputs
	for _, node := range outputNodes {
		fmt.Println(<-node.getOutputChan())
	}

}
