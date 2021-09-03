package rtype

import (
	"testing"
)

func TestInt_Is(t *testing.T) {
	n := NewInt(3)
	if n.Is(2) {
		t.Errorf("n must be 3")
	}
}

func TestInt_IsNot(t *testing.T) {
	n := NewInt(3)
	if n.IsNot(3) {
		t.Errorf("n must be 3")
	}
}

func TestInt_Increment(t *testing.T) {
	n := NewInt(0)
	n.Increment()
	if n.IsNot(1) {
		t.Errorf("n must be 1")
	}
}

func TestInt_Was(t *testing.T) {
	n := NewInt(0)
	n.Increment()
	if n.Was(1) {
		t.Errorf("previous n must be 0")
	}
}

func TestInt_WasNot(t *testing.T) {
	n := NewInt(0)
	n.Increment()
	if n.WasNot(0) {
		t.Errorf("previous n must be 0")
	}
}
