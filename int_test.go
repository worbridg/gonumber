package rtype

import "testing"

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
