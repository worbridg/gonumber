[![ci](https://github.com/worbridg/gonumber/actions/workflows/go.yml/badge.svg)](https://github.com/worbridg/gonumber/actions)

# gonumber

Let's improve code readability!

```golang
package main

import (
	"fmt"

	"github.com/worbridg/gonumber"
)

func main() {
	number := gonumber.New(0)
	if number.Is(0) {
		number.Increment()
	}
	number.WillBe(3)
	if err := number.Increment(); err != nil {
		fmt.Println(err)
		// Output: next value must be 3
	}

	number = gonumber.New(2)
	number.ShouldBe(1, 2, 3)
	if err := number.Add(2); err != nil {
		fmt.Println(err)
		// unexpected value
	}
}

```

# LICENSE

MIT

# Author

worbridg
