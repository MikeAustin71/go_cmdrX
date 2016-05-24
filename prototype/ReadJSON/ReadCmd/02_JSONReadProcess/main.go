package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type jsonobject struct {
	Object cmdHdrDat
}

type cmdHdrDat struct {
	Id    		int	`json:"id_no"`
	FirstName   string	`json:"first_name"`
	LastName 	string  `json:"last_name"`
}

func main() {
	dat, err := ioutil.ReadFile("./CmdrXCmds002.json")
	check(err)
	// fmt.Print(string(dat))
	var jsontype jsonobject
	err = json.Unmarshal(dat, &jsontype)
	check(err)
	fmt.Print(string(dat))
	fmt.Println(jsontype)
	fmt.Println("ID No:", jsontype.Object.Id)
	fmt.Println("First Name:", jsontype.Object.FirstName)
	fmt.Println("Last Name:", jsontype.Object.LastName)

}
