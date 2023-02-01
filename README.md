# pic
Cobol picture tag for golang
```
package main

import (
	"fmt"

	"github.com/jncss/pic"
)

type TestPic struct {
	testNumInt1  int     `pic:"999999"`
	testNumInt2  int     `pic:"9(6)S"`
	testNumFloat float64 `pic:"9(6)V99S"`
	testString   string  `pic:"X(20)"`
}

func main() {
	test := TestPic{
		testNumInt1:  1234,
		testNumInt2:  -123,
		testNumFloat: 1234.78,
		testString:   "Hello World",
	}

	result, err := pic.Marshall(test)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(result)
}
```
