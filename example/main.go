package main

import (
	"fmt"

	"github.com/worbridg/gonumber"
)

func main() {
	number := gonumber.New(0)
	if number.IsZero() {
		number.Increment()
	}
	number.WillBe(3)
	if err := number.Increment(); err != nil {
		fmt.Println(err)
		// Output: next value must be 3
	}

	number, _ = gonumber.New(2).ShouldBe(1, 2, 3)
	if err := number.Add(2); err != nil {
		fmt.Println(err)
		// unexpected value
	}
}
