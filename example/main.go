package main

import (
	"fmt"

	"github.com/worbridg/gonumber"
)

func all(numbers gonumber.Numbers) *gonumber.NumbersState {
	return gonumber.All(numbers)
}

func main() {
	number := gonumber.New(0)
	if number.IsZero() {
		number.Increment()
	}
	number.WillBe(3)
	if err := number.Increment(); err != nil {
		fmt.Println(err)
		// Output: next is expected to be 3
	}

	number, _ = gonumber.New(2).ShouldBe(1, 2, 3)
	if err := number.Add(2); err != nil {
		fmt.Println(err)
		// unexpected value
	}

	numbers := gonumber.Numbers{gonumber.New(1), gonumber.New(1)}
	if all(numbers).Are(1) {
		fmt.Println("they are same")
	}
}
