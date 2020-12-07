package tests

import (
	"fmt"
	"go-logic-circuits/gates"
	"go-logic-circuits/nodes"
	"math/rand"
	"testing"
	"time"
)

func TestParallel(t *testing.T) {
	// set random seed
	rand.Seed(time.Now().UnixNano())

	// start time
	start := time.Now()

	// slices to hold all input Nodes
	inputNodes := []nodes.Node{}
	logicNodes := []nodes.Node{}
	outputNodes := []nodes.Node{}

	// large scale paralell circuit
	n := 100000

	for i := 0; i < n; i++ {
		tmp := nodes.NewSourceNode()
		inputNodes = append(inputNodes, &tmp)
	}

	for i := 0; i < n-1; i++ {
		tmp1 := nodes.NewLogicNode(2, gates.Nand)
		tmp2 := nodes.NewLogicNode(2, gates.Xor)
		logicNodes = append(logicNodes, nodes.Connect(inputNodes[i:i+2], &tmp1))
		tmpSlice := []nodes.Node{inputNodes[i], &tmp1}
		logicNodes = append(logicNodes, nodes.Connect(tmpSlice, &tmp2))
		outputNodes = append(outputNodes, &tmp2)
	}

	// Set all output nodes to push
	for _, node := range outputNodes {
		node.IncrementOutput()
	}
	// push Bit onto all the sources
	// right now all sources can only be used once!!!
	for _, node := range inputNodes {
		go func(n nodes.Node) {
			n.(*nodes.SourceNode).SetBitValue(gates.Bit(rand.Intn(2) == 1)) // set random Bit
			n.EvaluateNode()                                                // Start pushing them out
		}(node)
	}

	// Start threads waiting to evaluate
	for _, node := range logicNodes {
		go node.EvaluateNode()
	}

	// print inputs
	inVals := []gates.Bit{}
	for _, node := range inputNodes {
		inVals = append(inVals, node.(*nodes.SourceNode).GetBitValue())
	}

	// collect outputs
	outVals := make([]gates.Bit, len(outputNodes))
	for i, node := range outputNodes {
		outVals[i] = <-node.GetOutputChan()
	}

	elapsed := time.Since(start)
	outstr := fmt.Sprintf("Executed %v Nodes in %v.", len(inputNodes)+len(logicNodes), elapsed)
	t.Log(outstr)

}

func TestSeries(t *testing.T) {

	// set random seed
	rand.Seed(time.Now().UnixNano())

	// start time
	start := time.Now()

	// slice to hold all input Nodes
	inputNodes := []nodes.Node{}
	logicNodes := []nodes.Node{}
	outputNodes := []nodes.Node{}

	// large scale series circuit
	n := 100000

	// Starting large chain event
	tmp1 := nodes.NewSourceNode()
	inputNodes = append(inputNodes, &tmp1)

	tmp2 := nodes.NewSourceNode()
	inputNodes = append(inputNodes, &tmp2)

	logicTmp1 := nodes.NewLogicNode(2, gates.Or)
	logicNodes = append(logicNodes, nodes.Connect(inputNodes, &logicTmp1))

	logicTmp2 := nodes.NewLogicNode(2, gates.Or)
	logicNodes = append(logicNodes, nodes.Connect([]nodes.Node{&tmp2, &logicTmp1}, &logicTmp2))

	for i := 2; i < n; i++ {
		logicTmp := nodes.NewLogicNode(2, gates.Or)
		logicNodes = append(logicNodes, nodes.Connect(logicNodes[i-2:i], &logicTmp))
	}

	outputNodes = append(outputNodes, logicNodes[len(logicNodes)-1])

	// Set all output nodes to push
	for _, node := range outputNodes {
		node.IncrementOutput()
	}

	// push Bit onto all the sources
	// right now all sources can only be used once!!!
	for _, node := range inputNodes {
		go func(n nodes.Node) {
			n.(*nodes.SourceNode).SetBitValue(gates.Bit(rand.Intn(2) == 1)) // set random Bit
			n.EvaluateNode()                                                // Start pushing them out
		}(node)
	}

	// Start threads waiting to evaluate
	for _, node := range logicNodes {
		go node.EvaluateNode()
	}

	// print inputs
	inVals := []gates.Bit{}
	for _, node := range inputNodes {
		inVals = append(inVals, node.(*nodes.SourceNode).GetBitValue())
	}

	// collect outputs
	outVals := make([]gates.Bit, len(outputNodes))
	for i, node := range outputNodes {
		outVals[i] = <-node.GetOutputChan()
	}

	elapsed := time.Since(start)
	outstr := fmt.Sprintf("Executed %v Nodes in %v.", len(inputNodes)+len(logicNodes), elapsed)
	t.Log(outstr)
}
