package rtype

type Int struct {
	n int
	p int
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

func (i *Int) Increment() {
	i.p = i.n
	i.n++
}

func (i Int) Was(n int) bool {
	return i.p == n
}

func (i Int) WasNot(n int) bool {
	return i.p != n
}
