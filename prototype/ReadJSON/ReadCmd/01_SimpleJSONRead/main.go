package main

import (
	"io/ioutil"
	"fmt"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	dat, err := ioutil.ReadFile("./CmdrXCmds001.json")
	check(err)
	fmt.Print(string(dat))
}