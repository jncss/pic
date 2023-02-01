package pic

import (
	"fmt"
	"testing"
)

type PicTest struct {
	TestInt    int     `pic:"9(12)V9(2)S"`
	TestFloat  float64 `pic:"9999V99"`
	TestString string  `pic:"X(20)"`
}

// Test
func TestMarshall(t *testing.T) {
	picTest := PicTest{
		TestInt:    1234,
		TestFloat:  1234.5678,
		TestString: "1234567890",
	}

	fmt.Println(Marshall(picTest))
}
