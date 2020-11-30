package main

type LogicNode struct {
	input  []chan Bit
	f      func(...Bit) Bit
	output chan Bit
}

type SourceNode struct {
	output chan Bit
}

func NewLogicNode(inputShape int, f func(...Bit) Bit) LogicNode {

	if inputShape < 1 {
		panic("inputShape must be greater than 1")
	}

	inChs := make([]chan Bit, inputShape)
	for i := 0; i < inputShape; i++ {
		inChs[i] = make(chan Bit, 1)
	}

	return LogicNode{
		inChs,
		f,
		make(chan Bit, 1),
	}
}

func NewSourceNode() SourceNode {
	return SourceNode{
		make(chan Bit, 1),
	}
}

func (n LogicNode) getOutputChan() chan Bit  { return n.output }
func (n SourceNode) getOutputChan() chan Bit { return n.output }

type Node interface {
	getOutputChan() chan Bit
}

func (n *LogicNode) EvaluateNode() {
	inputs := make([]Bit, 2)
	for _, ch := range n.input {
		inputs = append(inputs, <-ch)
	}
	n.output <- n.f(inputs...)
}

func Connect(ins []Node, out LogicNode) *LogicNode {
	if len(out.input) != len(ins) {
		panic("Mismatch expected inputs")
	}
	for i := 0; i < len(ins); i++ {
		out.input[i] = ins[i].getOutputChan()
	}
	return &out
}
