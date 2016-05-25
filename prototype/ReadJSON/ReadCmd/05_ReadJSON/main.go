package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type jsonobject struct {
	Batch cmdHdr `json:"batch_commands"`
}

type cmdHdr struct {
	Hdr cmdHdrDat `json:"jobs_header"`
}

type cmdHdrDat struct {
	// JSON Conversion works as long as first letter of field
	//   is capitalized
	LogFileRetentionInDays int    `json:"log_file_retention_in_days"`
	CmdExeDirectory string `json:"command_exe_directory"`
	LogPathFileName  string `json:"log_path_file_name"`
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	dat, err := ioutil.ReadFile("./CmdrX_JSON_005.json")
	check(err)

	var jsontype jsonobject
	err = json.Unmarshal(dat, &jsontype)
	check(err)
	fmt.Print(string(dat))
	fmt.Println(jsontype)
	fmt.Println("=======================================")
	fmt.Println("Log File Retention In Days:", jsontype.Batch.Hdr.LogFileRetentionInDays)
	fmt.Println("First Name:", jsontype.Batch.Hdr.CmdExeDirectory)
	fmt.Println("Last Name:", jsontype.Batch.Hdr.LogPathFileName)

}
