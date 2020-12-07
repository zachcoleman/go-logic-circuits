package main

import (
	"fmt"
	"go-logic-circuits/circuit"
	"go-logic-circuits/gates"
	"log"
	"math/rand"
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
	rand.Seed(time.Now().UnixNano())

	circ := circuit.NewCircuit()

	fmt.Println(circ.ParseString("INPUT A B"))
	fmt.Println(circ.ParseString("C = A OR B"))

	fmt.Println(circ.ParseString("OUTPUT C"))

	fmt.Println(circ.ParseString("INPUT D"))
	fmt.Println(circ.ParseString("E = C AND D"))

	fmt.Println(circ.ParseString("OUTPUT E"))

	input := make(map[string]gates.Bit)
	output := make(map[string]gates.Bit)

	runs := 100
	start := time.Now()
	for i := 0; i < runs; i++ {
		for name := range circ.InputNodes {
			input[name] = gates.Bit(rand.Intn(2) == 1)
		}
		output = circ.Evaluate(input)
	}

	avgTime := time.Since(start) / time.Duration(runs)

	// log.Printf("\n\t" + strings.Join(circ.Log, "\n\t"))
	log.Printf(`
	Input: %v
	Result: %v
	Num input nodes: %v
	Num output nodes: %v
	Total number of operations: %v
	Avg. Execution time (%v runs): %v`, input, output, len(circ.InputNodes), len(circ.OutputNodes),
		len(circ.CircuitNodes), runs, avgTime)

	// Duplicating circuit
	// newCirc := circuit.DuplicateCircuit(circ)

	// start = time.Now()
	// output = newCirc.Evaluate(input)
	// elapsed = time.Since(start)

	// add a test that verifies that the Duplicate pointers point to different
	// memory locations but produce same output (for random input)
	// fmt.Println(circ.InputNodes)
	// fmt.Println(newCirc.InputNodes)
	// fmt.Println(circ.CircuitNodes)
	// fmt.Println(newCirc.CircuitNodes)
	// fmt.Println(circ.OutputNodes)
	// fmt.Println(newCirc.OutputNodes)

	// log.Printf("\n\t" + strings.Join(newCirc.Log, "\n\t"))
	// log.Printf(`
	// Input: %v
	// Result: %v
	// Num input nodes: %v
	// Num output nodes: %v
	// Total number of operations: %v
	// Execution time: %v`, input, output, len(newCirc.InputNodes), len(newCirc.OutputNodes),
	// 	len(newCirc.CircuitNodes), elapsed)

}
