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

func TestNumbers_Have(t *testing.T) {
	numbers := Numbers{NewInt(1), NewInt(2), NewInt(3)}
	if _, exist := numbers.Have(2); !exist {
		t.Errorf("2 must be there")
	}
	if _, exist := numbers.Have(4); exist {
		t.Errorf("4 isn't expected")
	}
}

func TestNumbers_String(t *testing.T) {
	numbers := Numbers{NewInt(1), NewInt(2), NewInt(3)}
	if numbers.String() != "1,2,3" {
		t.Errorf("strings must be 1,2,3")
	}
}
