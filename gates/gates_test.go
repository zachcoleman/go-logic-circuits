package gates

import (
	"testing"
)

func TestNot(t *testing.T) {
	if Not(Bit(true)) != false {
		t.Error("Not func failed")
	}
	if Not(Bit(false)) != true {
		t.Error("Not func failed")
	}

}

func TestAnd(t *testing.T) {
	if And(Bit(true), Bit(true)) != true {
		t.Error("And func failed")
	}
	if And(Bit(true), Bit(false)) != false {
		t.Error("And func failed")
	}
	if And(Bit(false), Bit(true)) != false {
		t.Error("And func failed")
	}
	if And(Bit(false), Bit(false)) != false {
		t.Error("And func failed")
	}
}

func TestOr(t *testing.T) {
	if Or(Bit(true), Bit(true)) != true {
		t.Error("Or func failed")
	}
	if Or(Bit(true), Bit(false)) != true {
		t.Error("Or func failed")
	}
	if Or(Bit(false), Bit(true)) != true {
		t.Error("Or func failed")
	}
	if Or(Bit(false), Bit(false)) != false {
		t.Error("Or func failed")
	}
}

func TestNor(t *testing.T) {
	if Nor(Bit(true), Bit(true)) != false {
		t.Error("Nand func failed")
	}
	if Nor(Bit(true), Bit(false)) != false {
		t.Error("Nand func failed")
	}
	if Nor(Bit(false), Bit(true)) != false {
		t.Error("Nand func failed")
	}
	if Nor(Bit(false), Bit(false)) != true {
		t.Error("Nand func failed")
	}
}

func TestNand(t *testing.T) {
	if Nand(Bit(true), Bit(true)) != false {
		t.Error("Nand func failed")
	}
	if Nand(Bit(true), Bit(false)) != true {
		t.Error("Nand func failed")
	}
	if Nand(Bit(false), Bit(true)) != true {
		t.Error("Nand func failed")
	}
	if Nand(Bit(false), Bit(false)) != true {
		t.Error("Nand func failed")
	}
}

func TestXor(t *testing.T) {
	if Xor(Bit(true), Bit(true)) != false {
		t.Error("Xor func failed")
	}
	if Xor(Bit(true), Bit(false)) != true {
		t.Error("Xor func failed")
	}
	if Xor(Bit(false), Bit(true)) != true {
		t.Error("Xor func failed")
	}
	if Xor(Bit(false), Bit(false)) != false {
		t.Error("Xor func failed")
	}
}

func TestXnor(t *testing.T) {
	if Xnor(Bit(true), Bit(true)) != true {
		t.Error("Xnor func failed")
	}
	if Xnor(Bit(true), Bit(false)) != false {
		t.Error("Xnor func failed")
	}
	if Xnor(Bit(false), Bit(true)) != false {
		t.Error("Xnor func failed")
	}
	if Xnor(Bit(false), Bit(false)) != true {
		t.Error("Xnor func failed")
	}
}
