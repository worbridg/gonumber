[![ci](https://github.com/worbridg/rtype/actions/workflows/go.yml/badge.svg)](https://github.com/worbridg/rtype/actions)

# rtype

Let's improve code readability!

```golang
package main

import (
	"fmt"

	"github.com/worbridg/rtype"
)

func main() {
	number := rtype.NewInt(0)
	if number.Is(0) {
		number.Increment()
	}
	number.WillBe(3)
	if err := number.Increment(); err != nil {
		fmt.Println(err)
		// Output: next value must be 3
	}

	number = rtype.NewInt(2)
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
