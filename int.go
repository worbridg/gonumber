package rtype

import (
	"fmt"
	"strconv"
)

type Int struct {
	allowedN []int
	n        int
	prev     int
	next     int
}

func NewInt(n int) Int {
	return Int{n: n}
}

func (num Int) Is(n int) bool {
	return num.n == n
}

func (num Int) IsNot(n int) bool {
	return num.n != n
}

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

func (num Int) Was(n int) bool {
	return num.prev == n
}

func (num Int) WasNot(n int) bool {
	return num.prev != n
}

func (num *Int) WillBe(n int) {
	num.next = n
}

func (num *Int) ShouldBe(n ...int) {
	num.allowedN = n
}

func (num *Int) String() string {
	return strconv.Itoa(num.n)
}

func (num *Int) Int() int {
	return num.n
}

func (num *Int) Add(n int) error {
	if !num.isAllowed(num.n + n) {
		return fmt.Errorf("unexpected value")
	}
	num.n = num.n + n
	return nil
}

func (num *Int) Sub(n int) error {
	if !num.isAllowed(num.n - n) {
		return fmt.Errorf("unexpected value")
	}
	num.n = num.n - n
	return nil
}
