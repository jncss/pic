package pic

import (
	"fmt"
	"testing"
)

type PicTest struct {
	TestInt    int     `pic:"9(6)S"`
	TestFloat  float64 `pic:"999999V9(2)S"`
	TestString string  `pic:"X(20)"`
}

// Test
func TestMarshall(t *testing.T) {
	picTest := PicTest{
		TestInt:    -1234,
		TestFloat:  -1234.5678,
		TestString: "1234567890AAA",
	}

	// Marshall
	val, _ := Marshall(picTest)

	fmt.Println("Marshall:", val)

	// Unmarshall
	var picTest2 PicTest
	Unmarshall(val, &picTest2)

	fmt.Println("Unmarshall:", picTest2)
}
