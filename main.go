package main

import (
	"fmt"
	"go-logic-circuits/circuit"
	"go-logic-circuits/gates"
	"log"
	"time"
)

/*
TODO:
	(x) create performance tests
	() make text parser
	() create map implementation instead of slice
*/

func main() {

	// set random seed for pushing random values onto the sources
	// for testing primarily
	// rand.Seed(time.Now().UnixNano())

	// start timer
	start := time.Now()

	circ := circuit.NewCircuit()

	fmt.Println(circ.ParseString("INPUT A B"))
	fmt.Println(circ.ParseString("C = A OR B"))
	fmt.Println(circ.ParseString("OUTPUT C"))

	var input = map[string]gates.Bit{
		"A": gates.Bit(true),
		"B": gates.Bit(true),
	}
	output := circ.Evaluate(input)

	elapsed := time.Since(start)

	log.Printf(`
	Executed logic graph
	Result: %v
	Num input nodes: %v
	Num output nodes: %v
	Total number of operations: %v
	Execution time: %v`, output, len(circ.InputNodes), len(circ.OutputNodes), len(circ.CircuitNodes), elapsed)

}
