package pic

import (
	"fmt"
	"testing"
)

type PicTest struct {
	TestInt    int     `pic:"999999V99"`
	TestFloat  float64 `pic:"9999V99"`
	TestString string  `pic:"XXXXXXXXXXXXXX"`
}

// Test
func TestMarshall(t *testing.T) {
	picTest := PicTest{
		TestInt:    1234,
		TestFloat:  1234.5678,
		TestString: "12345678901234567890123456789012345678901234567890123456789012345678901234567890",
	}

	fmt.Println(Marshall(picTest))
}
