package gates

// Bit is alias for bool to make code more readable and easier to reason about
type Bit bool

// Not Bit function
func Not(b ...Bit) Bit { return !b[0] }

// And Bit function
func And(b ...Bit) Bit { return Bit(b[0] && b[1]) }

// Or Bit function
func Or(b ...Bit) Bit { return Bit(b[0] || b[1]) }

// Nand Bit function
func Nand(b ...Bit) Bit { return Not(And(b...)) }

// Nor Bit function
func Nor(b ...Bit) Bit { return Not(Or(b...)) }

// Xor Bit function
func Xor(b ...Bit) Bit { return Bit((b[0] && !b[1]) || (!b[0] && b[1])) }

// Xnor Bit function
func Xnor(b ...Bit) Bit { return Not(Xor(b...)) }

// Old Code
// type UnaryBitFunction func(Bit) Bit
// type BinaryBitFunction func(Bit, Bit) Bit
// type MultiBinaryBitFunction func([]Bit, []Bit) []Bit
// type MultiUnaryBitFunction func([]Bit) []Bit

// type BitLengthError struct {
// 	a []Bit
// 	b []Bit
// }

// func (e *BitLengthError) Error() string {
// 	return fmt.Sprintf("Mismatched bit slice lengths: %v, %v", len(e.a), len(e.b))
// }

// func newMultiBinaryBitFunction(f BinaryBitFunction) MultiBinaryBitFunction {
// 	return func(a, b []Bit) []Bit {
// 		if len(a) != len(b) {
// 			err := BitLengthError{a, b}
// 			panic(err)
// 		}
// 		ret := make([]Bit, len(a))
// 		for i := range a {
// 			ret[i] = f(a[i], b[i])
// 		}
// 		return ret
// 	}
// }

// func newMultiUnaryBitFunction(f UnaryBitFunction) MultiUnaryBitFunction {
// 	return func(a []Bit) []Bit {
// 		ret := make([]Bit, len(a))
// 		for i := range a {
// 			ret[i] = f(a[i])
// 		}
// 		return ret
// 	}
// }

// Old Bit functions
// func Not(a Bit) Bit     { return !a }
// func And(a, b Bit) Bit  { return Bit(a && b) }
// func Or(a, b Bit) Bit   { return Bit(a || b) }
// func Nand(a, b Bit) Bit { return Not(And(a, b)) }
// func Nor(a, b Bit) Bit  { return Not(Or(a, b)) }
// func Xor(a, b Bit) Bit  { return Bit((a && !b) || (!a && b)) }
// func Xnor(a, b Bit) Bit { return Not(Xor(a, b)) }
