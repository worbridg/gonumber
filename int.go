package rtype

import (
	"fmt"
)

type Int struct {
	n int
	p int
	q int
}

func NewInt(n int) Int {
	return Int{n: n}
}

func (i Int) Is(n int) bool {
	return i.n == n
}

func (i Int) IsNot(n int) bool {
	return i.n != n
}

func (i *Int) Increment() error {
	if i.q != 0 && i.q != i.n+1 {
		return fmt.Errorf("next value must be %d", i.q)
	}
	i.p = i.n
	i.q = 0
	i.n++
	return nil
}

func (i Int) Was(n int) bool {
	return i.p == n
}

func (i Int) WasNot(n int) bool {
	return i.p != n
}

func (i *Int) WillBe(n int) {
	i.q = n
}
