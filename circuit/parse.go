package circuit

import (
	"fmt"
	"go-logic-circuits/gates"
	nodes "go-logic-circuits/nodes"
	"strings"
)

var parseMap = map[string]func(...gates.Bit) gates.Bit{
	"NOT":  gates.Not,
	"AND":  gates.And,
	"OR":   gates.Or,
	"XOR":  gates.Xor,
	"NAND": gates.Nand,
	"NOR":  gates.Nor,
	"XNOR": gates.Xnor,
}

// ParseString parses an input string and applies the logic to the Circuit
// e.g
// 		INPUT A B C
// 		D = B OR C
//		E = A AND C
//		F = D XOR E
//		G = NOT F
//		OUTPUT G
//
// returns a boolean, error (sucessfully changed circuit, error)
func (c *Circuit) ParseString(input string) (bool, error) {

	// split the input into tokens
	tokens := strings.Split(input, " ")

	// add input(s)
	if strings.ToUpper(tokens[0]) == "INPUT" {
		for _, name := range tokens[1:] {
			if c.validNewInput(name) {
				tmp := nodes.NewSourceNode()
				c.InputNodes[name] = &tmp
			} else {
				// at least one input failed return false and error
				return false, fmt.Errorf("%v already defined : \n\t %v", name, input)
			}
		}
		// successfully looped through all inputs so log and return
		c.Log = append(c.Log, input)
		return true, nil
	}

	// add output node(s)
	if strings.ToUpper(tokens[0]) == "OUTPUT" {
		for _, name := range tokens[1:] {
			if c.validNewOutput(name) {
				c.OutputNodes[name] = c.CircuitNodes[name]
			} else {
				// at least one output failed return false and error
				if _, ok := c.CircuitNodes[name]; ok {
					return false, fmt.Errorf("%v already output : \n\t %v", name, input)
				}
				return false, fmt.Errorf("%v not defined : \n\t %v", name, input)
			}
		}
		// successfully looped through all outputs so log and return
		c.Log = append(c.Log, input)
		return true, nil
	}

	// Expecting a string of style:
	//		C = A OR B or C = NOT A
	// Unexpected length for a declarative statement
	if len(tokens) < 4 && len(tokens) > 5 {
		return false, fmt.Errorf("unexpected number of tokens %v: \n\t %v", len(tokens), input)
	}

	// add circuit node
	if len(tokens) == 5 {
		if c.validNewNode(strings.ToUpper(tokens[0])) && tokens[1] == "=" {
			newNode := tokens[0]
			inputs := []string{tokens[2], tokens[4]}
			operation := parseMap[tokens[3]]

			// get the two nodes from either InputNodes or CircuitNodes
			ins := []nodes.Node{}
			for _, name := range inputs {
				if _, ok := c.InputNodes[name]; ok {
					ins = append(ins, c.InputNodes[name])
				}
				if _, ok := c.CircuitNodes[name]; ok {
					ins = append(ins, c.CircuitNodes[name])
				}
			}
			// create new node, connect to circuit, add to log, return true
			tmp := nodes.NewLogicNode(2, operation)
			c.CircuitNodes[newNode] = nodes.Connect(ins, &tmp)
			c.Log = append(c.Log, input)
			return true, nil
		}

	} else if len(tokens) == 4 {
		// assumption is that this is of form B = NOT A
		if c.validNewNode(strings.ToUpper(tokens[0])) && tokens[1] == "=" {

			if tokens[2] != "NOT" {
				return false, fmt.Errorf("expected NOT and got %v: \n\t %v", tokens[2], input)
			}

			newNode := tokens[0]
			inputs := []string{tokens[3]}

			// get the two nodes from either InputNodes or CircuitNodes
			ins := []nodes.Node{}
			for _, name := range inputs {
				if _, ok := c.InputNodes[name]; ok {
					ins = append(ins, c.InputNodes[name])
				}
				if _, ok := c.CircuitNodes[name]; ok {
					ins = append(ins, c.CircuitNodes[name])
				}
			}

			// create new node, connect to circuit, add to log, return true
			tmp := nodes.NewLogicNode(1, parseMap["NOT"])
			c.CircuitNodes[newNode] = nodes.Connect(ins, &tmp)
			c.Log = append(c.Log, input)
			return true, nil
		}
	}

	// input fell through all possible cases
	return false, fmt.Errorf("Invalid statement: %v", input)
}
