package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"go-logic-circuits/circuit"
	"go-logic-circuits/gates"
	"log"
	"math/rand"
	"os"
	"strings"
)

func main() {

	// set random seed for pushing random values onto the sources
	// for testing primarily
	// rand.Seed(time.Now().UnixNano())

	// parse input
	fp := flag.String("filepath", "",
		"string filepath to circuit definition")
	flag.Parse()

	// open file
	f, err := os.Open(*fp)
	if errors.Is(err, os.ErrNotExist) {
		log.Fatal(err)
	}
	defer f.Close()

	// create circuit
	circ := circuit.NewCircuit()

	// read in the input
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		_, err := circ.ParseString(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
	}

	// get a counter
	outputData := make(map[string]float32)
	for key := range circ.OutputNodes {
		outputData[key] = 0.0
	}

	// if number of input nodes > 10 then run random
	// sample on circuit
	if len(circ.CircuitNodes) > 10 {

		// make maps
		input := make(map[string]gates.Bit)
		output := make(map[string]gates.Bit)

		// do runs
		runs := 1_000
		for i := 0; i < runs; i++ {
			for name := range circ.InputNodes {
				input[name] = gates.Bit(rand.Intn(2) == 1)
			}
			output = circ.Evaluate(input)

			for k, v := range output {
				if v {
					outputData[k] += 1.0
				}
			}
		}
		// get the freq data
		for key := range circ.OutputNodes {
			outputData[key] = outputData[key] / float32(runs)
		}

		inputKeys := make([]string, 0, len(circ.InputNodes))
		for key := range circ.InputNodes {
			inputKeys = append(inputKeys, key)
		}

		log.Printf("\n\t" + strings.Join(circ.Log, "\n\t"))
		log.Printf("\n\tInput(s): %v \n\tResults: %v",
			inputKeys, outputData)

	} else {
		// run an enumeration
		tree := makeTree(len(circ.InputNodes) + 1)
		perms := traverseTree(tree, []int{})

		// get one set ordering of key
		keys := []string{}
		for key := range circ.InputNodes {
			keys = append(keys, key)
		}

		// get pairs of results
		pairMap := [][]map[string]gates.Bit{}
		for _, perm := range perms {
			input := make(map[string]gates.Bit)
			for i, key := range keys {
				fmt.Println(i, gates.Bit(perm[i] == 1))
				input[key] = gates.Bit(perm[i] == 1)
			}
			output := circ.Evaluate(input)
			pairMap = append(pairMap, []map[string]gates.Bit{input, output})
		}

		log.Printf("\n\t" + strings.Join(circ.Log, "\n\t"))
		logMsg := "\nInput:" + strings.Repeat(" ", len(pairMap[0][0])*5) + "\t\t\tResult:"
		for _, pair := range pairMap {
			logMsg += fmt.Sprintf("\n%v \t\t%v", pair[0], pair[1])
		}
		log.Printf(logMsg)
	}
}

type Node struct {
	val   int
	left  *Node
	right *Node
}

func fillTree(root *Node, num int) *Node {
	if num < 1 {
		return root
	}

	if num > 1 {
		root.left = &Node{
			0,
			nil,
			nil,
		}
		root.right = &Node{
			1,
			nil,
			nil,
		}
	}
	fillTree(root.left, num-1)
	fillTree(root.right, num-1)

	return root
}

func makeTree(n int) *Node {
	root := &Node{
		-1,
		nil,
		nil,
	}
	fillTree(root, n)
	return root
}

func traverseTree(root *Node, curr []int) [][]int {
	res := [][]int{}

	if root.left == nil && root.right == nil {
		curr = append(curr, root.val)
		res = append(res, curr)
		return res
	} else {
		if root.val != -1 {
			curr = append(curr, root.val)
		}
		res = append(res, traverseTree(root.left, curr)...)
		res = append(res, traverseTree(root.right, curr)...)
	}

	return res
}
