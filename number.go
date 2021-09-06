package gonumber

import (
	"fmt"
	"strconv"
)

// Number is a wrapper of int and provide you code readability in your codes.
type Number struct {
	// allowedN holds what numerics must be in `n`.
	allowedN []int
	n        int
	prev     int
	next     int
}

// NewInt returns new Number. return a number if n is numeric otherwise an error.
func New(n interface{}) (*Number, error) {
	// TODO: other numeric types(int8 and so on) will be support with generics go1.18.
	return &Number{n: n.(int)}, nil
}

// NewInt is a type safe initialization for int.
func NewInt(n int) *Number {
	number, _ := New(n)
	return number
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
	if !number.canUpdate(number.n + 1) {
		return fmt.Errorf("next value is expected to be %d", number.next)
	}
	if !number.isAllowed(number.n + 1) {
		return fmt.Errorf("unexpected value")
	}
	number.prev = number.n
	number.next = 0
	number.n = number.n + 1
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
		if number.n == n[i] {
			number.allowedN = n
			return number, nil
		}
	}
	return nil, fmt.Errorf("the number should be one of %v", n)
}

// Strings returns a numeric string.
func (number *Number) String() string {
	return strconv.Itoa(number.n)
}

// Strings returns int type n.
func (number *Number) Number() int {
	return number.n
}

func (number *Number) canUpdate(n int) bool {
	return number.next == 0 || number.next == n
}

// Add plus n to n in this. See also Number.Increment.
func (number *Number) Add(n int) error {
	if !number.canUpdate(number.n + n) {
		return fmt.Errorf("next value is expected to be %d", number.next)
	}
	if !number.isAllowed(number.n + n) {
		return fmt.Errorf("unexpected value")
	}
	number.prev = number.n
	number.n = number.n + n
	number.next = 0
	return nil
}

// Sub subtracts n from n in this. See also Int.Increment.
func (number *Number) Sub(n int) error {
	if !number.canUpdate(number.n - n) {
		return fmt.Errorf("next value is expected to be %d", number.next)
	}
	if !number.isAllowed(number.n - n) {
		return fmt.Errorf("unexpected value")
	}
	number.prev = number.n
	number.n = number.n - n
	number.next = 0
	return nil
}

// IsZero checks if the number is zero or not.
func (number *Number) IsZero() bool {
	return number.n == 0
}

// IsNotZero checks if the number isn't zero or not.
func (number *Number) IsNotZero() bool {
	return number.n != 0
}

// ChangeTo changes n to new one. if either WillBe() or ShouldBe() is called
// before, n is checked if allowed to update it or not. it returns nil if
// allowed otherwise an error.
func (number *Number) ChangeTo(n int) error {
	if !number.canUpdate(n) {
		return fmt.Errorf("next value is expected to be %d", number.next)
	}
	if !number.isAllowed(n) {
		return fmt.Errorf("unexpected value")
	}
	number.prev = number.n
	number.n = n
	number.next = 0
	return nil
}

// IsGreaterThan evaluates "number > n".
func (number *Number) IsGreaterThan(n int) bool {
	return number.n > n
}

// IsGreaterThan evaluates "number < n".
func (number *Number) IsLessThan(n int) bool {
	return number.n < n
}

// IsGreaterThan evaluates "number == n".
func (number *Number) Equal(n int) bool {
	return number.n == n
}

// Int returns the number as int.
func (number *Number) Int() int {
	return number.n
}

// IsNegative checks if number is positive or not.
func (number *Number) IsPositive() bool {
	return number.n > 0
}

// IsNegative checks if number is negative or not.
func (number *Number) IsNegative() bool {
	return number.n < 0
}
