package rtype

type Int int

func NewInt(n int) Int {
	return Int(n)
}
