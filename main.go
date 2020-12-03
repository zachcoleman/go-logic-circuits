package main

import (
	"log"
	"math/rand"
	"time"
)

/*
TODO:
	() create tests
	() make text parser
	() create map implementation instead of slice
*/

func main() {

	// set random seed for pushing random values onto the sources
	// for testing primarily
	rand.Seed(time.Now().UnixNano())

	// start timer
	start := time.Now()

	inputNodes := []Node{}
	// for i := 0; i < 10000; i++ {
	// 	tmp := NewSourceNode()
	// 	inputNodes = append(inputNodes, &tmp)
	// }

	tmp1 := NewSourceNode()
	inputNodes = append(inputNodes, &tmp1)

	tmp2 := NewSourceNode()
	inputNodes = append(inputNodes, &tmp2)

	// slices to hold nodes
	logicNodes := []Node{}
	outputNodes := []Node{}

	logicTmp1 := NewLogicNode(2, Or)
	logicNodes = append(logicNodes, Connect(inputNodes, &logicTmp1))

	logicTmp2 := NewLogicNode(2, Or)
	logicNodes = append(logicNodes, Connect([]Node{&tmp2, &logicTmp1}, &logicTmp2))

	for i := 2; i < 10000; i++ {
		logicTmp := NewLogicNode(2, Or)
		logicNodes = append(logicNodes, Connect(logicNodes[i-2:i], &logicTmp))
	}

	outputNodes = append(outputNodes, logicNodes[len(logicNodes)-1])

	// for i := 0; i < 9999; i++ {
	// 	tmp1 := NewLogicNode(2, Nand)
	// 	tmp2 := NewLogicNode(2, Xor)
	// 	logicNodes = append(logicNodes, Connect(inputNodes[i:i+2], &tmp1))
	// 	tmpSlice := []Node{inputNodes[i], &tmp1}
	// 	logicNodes = append(logicNodes, Connect(tmpSlice, &tmp2))
	// 	outputNodes = append(outputNodes, &tmp2)
	// }

	// this is pattern for putting node in graph
	// can not use same variable since need new addresses
	// tmp1 := NewLogicNode(2, And)
	// logicNodes = append(logicNodes, Connect(inputNodes[:2], &tmp1))

	// tmp2 := NewLogicNode(2, Or)
	// logicNodes = append(logicNodes, Connect(inputNodes[2:], &tmp2))

	// tmp3 := NewLogicNode(2, Xor)
	// logicNodes = append(logicNodes, Connect(logicNodes[:2], &tmp3))
	// outputNodes = append(outputNodes, &tmp3)

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

	// collect outputs
	outVals := make([]Bit, len(outputNodes))
	for i, node := range outputNodes {
		outVals[i] = <-node.getOutputChan()
	}

	elapsed := time.Since(start)
	log.Print(inVals)
	log.Print(outVals)
	log.Printf(`
	Executed logic graph
	Num input nodes: %v
	Num output nodes: %v
	Total number of operations: %v 
	Execution time: %v`, len(inputNodes), len(outputNodes), len(logicNodes), elapsed)

}
