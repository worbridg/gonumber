package rtype

type Int int

func NewInt(n int) Int {
	return Int(n)
}

func (i Int) Is(n int) bool {
	return int(i) == n
}

func (i Int) IsNot(n int) bool {
	return int(i) != n
}
