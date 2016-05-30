package main

import (
	"fmt"
)

// CommasInt receives an int and returns a string representing
// the whole number with comma grouping.
func CommasInt(x int) string {
	value := fmt.Sprint(x)
	for i := len(value) - 3; i > 0; i -= 3 {
		value = value[:i] + "," + value[i:]
	}
	return value
}

// CommasInt64 receives an int64 and returns a string representing
// the whole number with comma grouping.
func CommasInt64(x int64) string {
	value := fmt.Sprint(x)
	for i := len(value) - 3; i > 0; i -= 3 {
		value = value[:i] + "," + value[i:]
	}
	return value
}

func main() {
	a := int64(500300000)
	fmt.Println("Original Int64", a)
	fmt.Println("Formatted Int64", CommasInt64(a))
	b := 8500300
	fmt.Println("Original int", b)
	fmt.Println("Formatted int", CommasInt(b))

}
