package gonumber

import (
	"testing"
)

func TestNumbers_All(t *testing.T) {
	numbers := Numbers{NewInt(1), NewInt(1), NewInt(1)}
	if !All(numbers).Are(1) {
		t.Errorf("all numbers are 1")
	}
	numbers = append(numbers, NewInt(2))
	if All(numbers).Are(1) {
		t.Errorf("2 is added")
	}
}

func TestNumbers_OneOf(t *testing.T) {
	numbers := Numbers{NewInt(1), NewInt(1), NewInt(1)}
	if _, exist := OneOf(numbers).IsZero(); exist {
		t.Errorf("0 should be none")
	}
	numbers = append(numbers, NewInt(0))
	if _, exist := OneOf(numbers).IsZero(); !exist {
		t.Errorf("0 is added")
	}
}
