package gonumber

import "testing"

func TestNumbers_All(t *testing.T) {
	numbers := Numbers{New(1), New(1), New(1)}
	if !All(numbers).Are(1) {
		t.Errorf("all numbers are 1")
	}
	numbers = append(numbers, New(2))
	if All(numbers).Are(1) {
		t.Errorf("2 is added")
	}
}
