package rtype

import (
	"fmt"
)

type Int struct {
	a []int
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

	if len(i.a) > 0 {
		for a := 0; a < len(i.a); a++ {
			if i.a[a] == i.n+1 {
				goto INCREMENT
			}
		}
		return fmt.Errorf("unexpected value")
	}
INCREMENT:
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

func (i *Int) ShouldBe(n ...int) {
	i.a = n
}
