package main

import (
	"fmt"
	pf "go_cmdrX/src/stringmgr/printfmtr"
)

func main() {
	a := int64(500300000)
	fmt.Println("Original Int64", a)
	aa := pf.CommasInt64(a)
	fmt.Println("Formatted Int64", aa)
	b := 8500300
	bb := pf.CommasInt(b)
	fmt.Println("Original int", b)
	fmt.Println("Formatted int", bb)
}
