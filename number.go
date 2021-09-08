package gonumber

import (
	"fmt"
	"strconv"
)

var (
	MaxUpTo = 100
)

// Number is a wrapper of int and provide you code readability in your codes.
type Number struct {
	min      int
	max      int
	positive bool
	// allowedN holds what numerics must be in `n`.
	allowedN []int
	n        int
	prev     int
	next     int
	changed  bool
}

// NewInt returns new Number. return a number if n is numeric otherwise an error.
func New(n interface{}) (*Number, error) {
	// TODO: other numeric types(int8 and so on) will be support with generics go1.18.
	return &Number{n: n.(int)}, nil
}

// NewInt is a type safe initialization for int.
func NewInt(n int) *Number {
	number, _ := New(n)
	return number
}

// Is is equal to "=="
func (number *Number) Is(n int) bool {
	return number.n == n
}

// IsNot is equal to "!="
func (number *Number) IsNot(n int) bool {
	return number.n != n
}

// changeTo add n to n of the number and holds last value of the number in
// myself. if next value is designated by WillBe(), they must be same. if not,
// returns an error. and more if a value to n is restricted by ShouldBe(), it
// also must be obey. cleared both, it return nil.
func (number *Number) changeTo(n int) error {
	if err, ok := number.canUpdate(n); !ok {
		return fmt.Errorf("unexpected value: %v", err)
	}
	number.update(n)
	return nil
}

func (number *Number) update(n int) {
	number.changed = true
	number.prev = number.n
	number.next = 0
	number.n = n
}

// Increment add 1 to the number.
func (number *Number) Increment() error {
	return number.changeTo(number.n + 1)
}

// Decrement subtract 1 from the number.
func (number *Number) Decrement() error {
	return number.changeTo(number.n - 1)
}

// isAllowed checks if the number is allowed to update with n.
func (number *Number) canUpdate(n int) (error, bool) {
	if number.next > 0 && number.next != n {
		return fmt.Errorf("next value is expected to be %d", number.next), false
	}
	if number.positive {
		return fmt.Errorf("value should be positive"), n > 0
	}
	if number.max-number.min != 0 {
		return fmt.Errorf("the number should be between %d and %d", number.min, number.max), between(n, number.min, number.max)
	}
	if len(number.allowedN) == 0 {
		return nil, true
	}
	return number.isAllowedToSet(n)
}

func between(n, min, max int) bool {
	return min <= n && n <= max
}

// IsBetween checks if the number is between min and max.
func (number *Number) IsBetween(min, max int) bool {
	return between(number.n, min, max)
}

func (number *Number) isAllowedToSet(n int) (error, bool) {
	for i := 0; i < len(number.allowedN); i++ {
		if number.allowedN[i] == n {
			return nil, true
		}
	}
	return fmt.Errorf("%d isn't allowed to set", n), false
}

// Was checks if last value is equal to n.
func (number *Number) Was(n int) bool {
	return number.prev == n
}

// WasNot checks if last value isn't equal to n.
func (number *Number) WasNot(n int) bool {
	return number.prev != n
}

// WillBe expects that next value is n. See Also Number.Increment.
func (number *Number) WillBe(n int) {
	number.next = n
}

// ShouldBe restricts values that an user can set. See Also Int.Increment.
func (number *Number) ShouldBe(n ...int) (*Number, error) {
	for i := 0; i < len(n); i++ {
		if number.n == n[i] {
			number.allowedN = n
			return number, nil
		}
	}
	return nil, fmt.Errorf("the number should be one of %v", n)
}

// ShouldBePositive restricts value than an user can set to positive.
func (number *Number) ShouldBePositive() (*Number, error) {
	number.positive = true
	return number, nil
}

// ShouldBeBetween restricts value that an user can set to between min and max.
func (number *Number) ShouldBeBetween(min, max int) (*Number, error) {
	number.min = min
	number.max = max
	return number, nil
}

// Strings returns a numeric string.
func (number *Number) String() string {
	return strconv.Itoa(number.n)
}

// Strings returns int type n.
func (number *Number) Number() int {
	return number.n
}

// Add plus n to n in this. See also Number.Increment.
func (number *Number) Add(n int) error {
	return number.changeTo(number.n + n)
}

// Sub subtracts n from n in this. See also Int.Increment.
func (number *Number) Sub(n int) error {
	return number.changeTo(number.n - n)
}

// IsZero checks if the number is zero or not.
func (number *Number) IsZero() bool {
	return number.n == 0
}

// IsNotZero checks if the number isn't zero or not.
func (number *Number) IsNotZero() bool {
	return number.n != 0
}

// ChangeTo changes n to new one. if either WillBe() or ShouldBe() is called
// before, n is checked if allowed to update it or not. it returns nil if
// allowed otherwise an error.
func (number *Number) ChangeTo(n int) error {
	return number.changeTo(n)
}

// IsGreaterThan evaluates "number > n".
func (number *Number) IsGreaterThan(n int) bool {
	return number.n > n
}

// IsLessThan evaluates "number < n".
func (number *Number) IsLessThan(n int) bool {
	return number.n < n
}

// Equal evaluates "number == n".
func (number *Number) Equal(n int) bool {
	return number.n == n
}

// Int returns the number as int.
func (number *Number) Int() int {
	return number.n
}

// IsNegative checks if number is positive or not.
func (number *Number) IsPositive() bool {
	return number.n > 0
}

// IsNegative checks if number is negative or not.
func (number *Number) IsNegative() bool {
	return number.n < 0
}

// A safe logic to protect when too long n is supplied unexpectedly.
func upto(n int) int {
	if n > MaxUpTo {
		return MaxUpTo
	}
	return n
}

// UpTo increments the number up to n and returns numbers.
func (number *Number) UpTo(n int) Numbers {
	if number.IsGreaterThan(n) {
		return Numbers{number}
	}
	numbers := make(Numbers, upto(n-number.n+1))
	for i := number.n; i <= n; i++ {
		numbers[i-number.n] = NewInt(i)
	}
	return numbers
}

// IsEven checks if the number is even or not.
func (number *Number) IsEven() bool {
	return number.n%2 == 0
}

// IsOdd checks if the number is odd or not.
func (number *Number) IsOdd() bool {
	return number.n%2 == 1
}

// In checks if the number is in numbers.
func (number *Number) In(numbers Numbers) bool {
	_, ok := numbers.Have(number.n)
	return ok
}

// HasBeenChanged checks if the number have been changed.
func (number *Number) HasBeenChanged() bool {
	return number.changed
}

// Copy copys the number to new one without changed status.
func (number *Number) Copy(n *Number) *Number {
	return &Number{
		max:      number.max,
		min:      number.min,
		positive: number.positive,
		allowedN: number.allowedN,
		n:        number.n,
		prev:     number.prev,
		next:     number.next,
		changed:  false,
	}
}

// Rollback swaps current value and previous one.
func (number *Number) Rollback() {
	number.prev, number.n = number.n, number.prev
}

type EveryNumberState struct {
	number *Number
}

func Every(number *Number) *EveryNumberState {
	return &EveryNumberState{number: number}
}

// every doing in todoLists
func (s *EveryNumberState) In(numbers Numbers) Numbers {
	everyNumbers := Numbers{}
	for i := 0; i < len(numbers); i++ {
		if numbers.At(i).Equal(s.number.Int()) {
			everyNumbers = append(everyNumbers, numbers.At(i))
		}
	}
	return everyNumbers
}
