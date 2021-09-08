package gonumber

import "strings"

// Numbers is a collection of Number.
type Numbers []*Number

// NumbersState stands for a state of Numbers.
type AllNumbersState struct {
	numbers Numbers
}

type OneOfNumbersState struct {
	numbers Numbers
}

// All wraps numbers with AllNumbersState and returns it.
func All(numbers Numbers) *AllNumbersState {
	return &AllNumbersState{numbers: numbers}
}

// OneOf wraps numbers with OneOfNumbersState and returns it.
func OneOf(numbers Numbers) *OneOfNumbersState {
	return &OneOfNumbersState{numbers: numbers}
}

// Are checks if all numbers are n or not. it returns true if same otherwise
// false.
func (s *AllNumbersState) Are(n int) bool {
	for i := 0; i < len(s.numbers); i++ {
		if !s.numbers[i].Equal(n) {
			return false
		}
	}
	return true
}

// IsZero checks if there is 0 in numbers or not. it returns true if same
// otherwise false.
func (s *OneOfNumbersState) IsZero() (*Number, bool) {
	if number, ok := s.numbers.Have(0); ok {
		return number, ok
	}
	return nil, false
}

// Is is equivalent to Numbers.Have().
func (s *OneOfNumbersState) Is(n int) (*Number, bool) {
	if number, ok := s.numbers.Have(n); ok {
		return number, ok
	}
	return nil, false
}

// Are compares numbers to n slice. it returns true if same otherwise false.
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

// AreNot equals to !Numbers.Are().
func (numbers Numbers) AreNot(n ...int) bool {
	return !numbers.Are(n...)
}

// AreSame checks if all numbers are same or not.
func (numbers Numbers) AreSame() bool {
	for i := 0; i < len(numbers); i++ {
		if !numbers[0].Equal(numbers[i].n) {
			return false
		}
	}
	return true
}

// AreNotSame equals to !Numbers.AreSame().
func (numbers Numbers) AreNotSame() bool {
	return !numbers.AreSame()
}

// Have checks if numbers contains n or not.
func (numbers Numbers) Have(n int) (*Number, bool) {
	for i := 0; i < len(numbers); i++ {
		if numbers[i].Equal(n) {
			return numbers[i], true
		}
	}
	return nil, false
}

// AreNothing returns true if the length of the numbers is 0 otherwise false.
func (numbers Numbers) AreNothing() bool {
	return len(numbers) == 0
}

// Strings returns a string separated by conma.
func (numbers Numbers) String() string {
	s := make([]string, len(numbers))
	for i := 0; i < len(numbers); i++ {
		s[i] = numbers[i].String()
	}
	return strings.Join(s, ",")
}
