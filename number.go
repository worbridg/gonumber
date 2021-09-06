package gonumber

import (
	"fmt"
	"strconv"
)

// Number is a wrapper of int and provide you code readability in your codes.
type Number struct {
	// allowedN holds what numerics must be in `n`.
	allowedN []int
	n        interface{}
	prev     int
	next     int
}

// NewInt returns new Number.
func New(n interface{}) *Number {
	return &Number{n: n}
}

// Is is equal to "=="
func (number *Number) Is(n int) bool {
	return number.n == n
}

// IsNot is equal to "!="
func (number *Number) IsNot(n int) bool {
	return number.n != n
}

// Increment add 1 to n and holds last value of n in myself.
// if next value is designated by WillBe(), n+1 must be same to it.
// if not, returns an error. and more if a value to n is restricted
// by ShouldBe(), it also must be obey. cleared both, it return nil.
func (number *Number) Increment() error {
	if !number.canUpdate(number.n.(int) + 1) {
		return fmt.Errorf("next value must be %d", number.next)
	}
	if !number.isAllowed(number.n.(int) + 1) {
		return fmt.Errorf("unexpected value")
	}
	number.prev = number.n.(int)
	number.next = 0
	number.n = number.n.(int) + 1
	return nil
}

// isAllowed always return true if allowedN isn't set.
func (number *Number) isAllowed(n int) bool {
	if len(number.allowedN) > 0 {
		for i := 0; i < len(number.allowedN); i++ {
			if number.allowedN[i] == n {
				return true
			}
		}
		return false
	}
	return true
}

// Was checks if last value is equal to n.
func (number *Number) Was(n int) bool {
	return number.prev == n
}

// WasNot checks if last value isn't equal to n.
func (number *Number) WasNot(n int) bool {
	return number.prev != n
}

// WillBe expects that next value is n. See Also Number.Increment.
func (number *Number) WillBe(n int) {
	number.next = n
}

// ShouldBe restricts values that an user can set. See Also Int.Increment.
func (number *Number) ShouldBe(n ...int) (*Number, error) {
	for i := 0; i < len(n); i++ {
		if number.n.(int) == n[i] {
			number.allowedN = n
			return number, nil
		}
	}
	return nil, fmt.Errorf("the number should be one of %v", n)
}

// Strings returns a numeric string.
func (number *Number) String() string {
	return strconv.Itoa(number.n.(int))
}

// Strings returns int type n.
func (number *Number) Number() int {
	return number.n.(int)
}

func (number *Number) canUpdate(n int) bool {
	return number.next == 0 || number.next == n
}

// Add plus n to n in this. See also Number.Increment.
func (number *Number) Add(n int) error {
	if !number.canUpdate(number.n.(int) + n) {
		return fmt.Errorf("next value is expected to be %d", number.next)
	}
	if !number.isAllowed(number.n.(int) + n) {
		return fmt.Errorf("unexpected value")
	}
	number.prev = number.n.(int)
	number.n = number.n.(int) + n
	number.next = 0
	return nil
}

// Sub subtracts n from n in this. See also Int.Increment.
func (number *Number) Sub(n int) error {
	if !number.isAllowed(number.n.(int) - n) {
		return fmt.Errorf("unexpected value")
	}
	number.prev = number.n.(int)
	number.n = number.n.(int) - n
	number.next = 0
	return nil
}
