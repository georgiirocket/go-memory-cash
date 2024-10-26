Go memory cash
================================

A simple package for data caching

```bash
go get github.com/georgiirocket/gomemorycash
```

## Example #1

```go
package main

import (
	"fmt"
	"github.com/georgiirocket/gomemorycash"
)

func main() {
	cash := gomemorycash.NewCash()

	//Set data
	cash.Set("1", 20)
	cash.Set("2", "tester")
	cash.Set("3", []int{1, 2, 4})

	//Get data
	data := cash.Get("1")

	//Delete data
	cash.Delete("2")

	fmt.Print(data)
}
```
