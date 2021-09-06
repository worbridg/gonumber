package gonumber

// Numbers is a collection of Number.
type Numbers []*Number

// NumbersState stands for a state of Numbers.
type NumbersState struct {
	numbers Numbers
}

type OneOfNumbersState struct {
	numbers Numbers
}

func All(numbers Numbers) *NumbersState {
	return &NumbersState{numbers: numbers}
}

func OneOf(numbers Numbers) *OneOfNumbersState {
	return &OneOfNumbersState{numbers: numbers}
}

func (ns *NumbersState) Are(n int) bool {
	for i := 0; i < len(ns.numbers); i++ {
		if ns.numbers[i].n != n {
			return false
		}
	}
	return true
}

func (s *OneOfNumbersState) IsZero() bool {
	for i := 0; i < len(s.numbers); i++ {
		if s.numbers[i].IsZero() {
			return true
		}
	}
	return false
}

func (numbers Numbers) Are(n ...int) bool {
	if len(numbers) != len(n) {
		return false
	}
	for i := 0; i < len(numbers); i++ {
		if numbers[i].n != n[i] {
			return false
		}
	}
	return true
}

func (numbers Numbers) AreNot(n ...int) bool {
	return !numbers.Are(n...)
}

func (numbers Numbers) AreSame() bool {
	for i := 0; i < len(numbers); i++ {
		if numbers[0].n != numbers[i].n {
			return false
		}
	}
	return true
}

func (numbers Numbers) AreNotSame() bool {
	return !numbers.AreSame()
}
