[![ci](https://github.com/worbridg/gonumber/actions/workflows/go.yml/badge.svg)](https://github.com/worbridg/gonumber/actions)

# gonumber

Let's improve code readability!

```golang
package main

import (
	"fmt"

	"github.com/worbridg/gonumber"
)

func all(numbers gonumber.Numbers) *gonumber.AllNumbersState {
	return gonumber.All(numbers)
}

func oneOf(numbers gonumber.Numbers) *gonumber.OneOfNumbersState {
	return gonumber.OneOf(numbers)
}

func main() {
	number := gonumber.NewInt(0)
	if number.IsZero() {
		number.Increment()
	}
	number.WillBe(3)
	if err := number.Increment(); err != nil {
		fmt.Println(err)
		// Output: next is expected to be 3
	}

	number, _ = gonumber.NewInt(2).ShouldBe(1, 2, 3)
	if err := number.Add(2); err != nil {
		fmt.Println(err)
		// unexpected value
	}

	numbers := gonumber.Numbers{gonumber.NewInt(1), gonumber.NewInt(1), gonumber.NewInt(0)}
	if it, exist := oneOf(numbers).IsZero(); exist {
		it.WillBe(1)
		it.ChangeTo(1)
	}
	if numbers.Are(1, 1, 1) {
		fmt.Println("they are same")
	}
	if all(numbers).Are(1) {
		fmt.Println("they also are same")
	}
}

```

# LICENSE

MIT

# Author

worbridg
