package main

import "fmt"

type Bit bool

type UnaryBitFunction func(Bit) Bit
type BinaryBitFunction func(Bit, Bit) Bit
type MultiBinaryBitFunction func([]Bit, []Bit) []Bit
type MultiUnaryBitFunction func([]Bit) []Bit

type BitLengthError struct {
	a []Bit
	b []Bit
}

func (e *BitLengthError) Error() string {
	return fmt.Sprintf("Mismatched bit slice lengths: %v, %v", len(e.a), len(e.b))
}

func newMultiBinaryBitFunction(f BinaryBitFunction) MultiBinaryBitFunction {
	return func(a, b []Bit) []Bit {
		if len(a) != len(b) {
			err := BitLengthError{a, b}
			panic(err)
		}
		ret := make([]Bit, len(a))
		for i := range a {
			ret[i] = f(a[i], b[i])
		}
		return ret
	}
}

func newMultiUnaryBitFunction(f UnaryBitFunction) MultiUnaryBitFunction {
	return func(a []Bit) []Bit {
		ret := make([]Bit, len(a))
		for i := range a {
			ret[i] = f(a[i])
		}
		return ret
	}
}

// func Not(a Bit) Bit     { return !a }
func Not(b ...Bit) Bit { return !b[0] }

// func And(a, b Bit) Bit  { return Bit(a && b) }
func And(b ...Bit) Bit { return Bit(b[0] && b[1]) }

// func Or(a, b Bit) Bit   { return Bit(a || b) }
func Or(b ...Bit) Bit { return Bit(b[0] || b[1]) }

// func Nand(a, b Bit) Bit { return Not(And(a, b)) }
func Nand(b ...Bit) Bit { return Not(And(b...)) }

// func Nor(a, b Bit) Bit  { return Not(Or(a, b)) }
func Nor(b ...Bit) Bit { return Not(Or(b...)) }

// func Xor(a, b Bit) Bit  { return Bit((a && !b) || (!a && b)) }
func Xor(b ...Bit) Bit { return Bit((b[0] && !b[1]) || (!b[0] && b[1])) }

// func Xnor(a, b Bit) Bit { return Not(Xor(a, b)) }
func Xnor(b ...Bit) Bit { return Not(Xor(b...)) }
