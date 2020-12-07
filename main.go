package main

import (
	"fmt"
	"go-logic-circuits/circuit"
	"go-logic-circuits/gates"
	"log"
	"strings"
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

	fmt.Println(circ.ParseString("INPUT D E"))
	fmt.Println(circ.ParseString("F = A XOR E"))
	fmt.Println(circ.ParseString("G = NOT B"))
	fmt.Println(circ.ParseString("H = F NOR G"))

	fmt.Println(circ.ParseString("OUTPUT H"))

	var input = map[string]gates.Bit{
		"A": gates.Bit(true),
		"B": gates.Bit(true),
		"D": gates.Bit(true),
		"E": gates.Bit(true),
	}
	output := circ.Evaluate(input)

	elapsed := time.Since(start)

	log.Printf("\n\t" + strings.Join(circ.Log, "\n\t"))
	log.Printf(`
	Input: %v
	Result: %v
	Num input nodes: %v
	Num output nodes: %v
	Total number of operations: %v
	Execution time: %v`, input, output, len(circ.InputNodes), len(circ.OutputNodes),
		len(circ.CircuitNodes), elapsed)

	newCirc := circuit.DuplicateCircuit(circ)
	fmt.Println(newCirc.Log)
	fmt.Println(newCirc.Evaluate(input))

}
