package gonumber

import (
	"fmt"
	"strconv"
)

// Int is a wrapper of int and provide you code readability in your codes.
type Int struct {
	// allowedN holds what numerics must be in `n`.
	allowedN []int
	n        int
	prev     int
	next     int
}

// NewInt returns new Int.
func NewInt(n int) Int {
	return Int{n: n}
}

// Is is equal to "=="
func (num Int) Is(n int) bool {
	return num.n == n
}

// IsNot is equal to "!="
func (num Int) IsNot(n int) bool {
	return num.n != n
}

// Increment add 1 to n and holds last value of n in myself.
// if next value is designated by WillBe(), n+1 must be same to it.
// if not, returns an error. and more if a value to n is restricted
// by ShouldBe(), it also must be obey. cleared both, it return nil.
func (num *Int) Increment() error {
	if num.next != 0 && num.next != num.n+1 {
		return fmt.Errorf("next value must be %d", num.next)
	}
	if !num.isAllowed(num.n + 1) {
		return fmt.Errorf("unexpected value")
	}
	num.prev = num.n
	num.next = 0
	num.n++
	return nil
}

// isAllowed always return true if allowedN isn't set.
func (num Int) isAllowed(n int) bool {
	if len(num.allowedN) > 0 {
		for i := 0; i < len(num.allowedN); i++ {
			if num.allowedN[i] == n {
				return true
			}
		}
		return false
	}
	return true
}

// Was checks if last value is equal to n.
func (num Int) Was(n int) bool {
	return num.prev == n
}

// WasNot checks if last value isn't equal to n.
func (num Int) WasNot(n int) bool {
	return num.prev != n
}

// WillBe expects that next value is n. See Also Int.Increment.
func (num *Int) WillBe(n int) {
	num.next = n
}

// ShouldBe restricts values that an user can set. See Also Int.Increment.
func (num *Int) ShouldBe(n ...int) {
	num.allowedN = n
}

// Strings returns a numeric string.
func (num *Int) String() string {
	return strconv.Itoa(num.n)
}

// Strings returns int type n.
func (num *Int) Int() int {
	return num.n
}

// Add plus n to n in this. See also Int.Increment.
func (num *Int) Add(n int) error {
	if !num.isAllowed(num.n + n) {
		return fmt.Errorf("unexpected value")
	}
	num.n = num.n + n
	return nil
}

// Sub subtracts n from n in this. See also Int.Increment.
func (num *Int) Sub(n int) error {
	if !num.isAllowed(num.n - n) {
		return fmt.Errorf("unexpected value")
	}
	num.n = num.n - n
	return nil
}
