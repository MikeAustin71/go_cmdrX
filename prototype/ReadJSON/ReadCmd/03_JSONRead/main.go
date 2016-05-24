package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type jsonobject struct {
	Cmds cmdHdr  `json:"commands"`
}

type cmdHdr struct {
	Hdr cmdHdrDat `json:"header"`
}

type cmdHdrDat struct {
	Id    		int	`json:"id_no"`
	FirstName   string	`json:"first_name"`
	LastName 	string  `json:"last_name"`
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	dat, err := ioutil.ReadFile("./CmdrXCmds003.json")
	check(err)

	var jsontype jsonobject
	err = json.Unmarshal(dat, &jsontype)
	check(err)
	fmt.Print(string(dat))
	fmt.Println(jsontype)
	fmt.Println("=======================================")
	fmt.Println("ID No:", jsontype.Cmds.Hdr.Id)
	fmt.Println("First Name:", jsontype.Cmds.Hdr.FirstName)
	fmt.Println("Last Name:", jsontype.Cmds.Hdr.LastName)

}
