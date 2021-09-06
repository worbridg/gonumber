package gonumber

// Numbers is a collection of Number.
type Numbers []*Number

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
