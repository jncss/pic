# pic
Cobol picture tag for golang

## Install:
```
go get github.com/jncss/pic
```
## Example:
```
package main

import (
	"fmt"

	"github.com/jncss/pic"
)

type TestPic struct {
	TestNumInt1  int     `pic:"999999"`
	TestNumInt2  int     `pic:"9(6)S"`
	TestNumFloat float64 `pic:"9(6)V99S"`
	TestString   string  `pic:"X(20)"`
}

func main() {
	test := TestPic{
		TestNumInt1:  1234,
		TestNumInt2:  -123,
		TestNumFloat: 1234.78,
		TestString:   "Hello World",
	}

	// Original
	fmt.Println(test)

	// Marshall
	result, err := pic.Marshall(test)
	fmt.Println(result, err)

	// Unmarshall
	var test2 TestPic
	err = pic.Unmarshall(result, &test2)
	fmt.Println(test2, err)
}

```
## Example output:
```
Original struct: {1234 -123 1234.78 Hello World}
Marshall result: 001234000123-00123478+Hello World         
Unmarshall result: {1234 -123 1234.78 Hello World         }
```
