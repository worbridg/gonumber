package rtype

import (
	"fmt"
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

	if len(num.allowedN) > 0 {
		for i := 0; i < len(num.allowedN); i++ {
			if num.allowedN[i] == num.n+1 {
				goto INCREMENT
			}
		}
		return fmt.Errorf("unexpected value")
	}
INCREMENT:
	num.prev = num.n
	num.next = 0
	num.n++
	return nil
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
