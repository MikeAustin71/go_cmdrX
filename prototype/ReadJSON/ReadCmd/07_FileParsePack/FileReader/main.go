package main

import(
	// "go_cmdrX/prototype/ReadJSON/ReadCmd/07_FileParsePack/DataStructs"
	"go_cmdrX/prototype/ReadJSON/ReadCmd/07_FileParsePack/JSONParser"

	"fmt"
)

func main() {
	// Note: relative JSON file path is determined by reference
	// to main.go path.
	fileName := "../JSONParser/CmdrX_JSON_070.json"
	jObj := JSONParser.ParseJSONCmds(fileName)
	fmt.Println(jObj)
	fmt.Println("=======================================")
	fmt.Println("Header")
	fmt.Println("Log File Retention In Days:", jObj.Batch.Hdr.LogFileRetentionInDays)
	fmt.Println("=======================================")
	fmt.Println("Command 1")
	fmt.Println("Cmd-1 Display Name:", jObj.Batch.Jobs[0].DisplayName)
	fmt.Println("Cmn-1 Cmd Element-1:", jObj.Batch.Jobs[0].CmdElements[0].CmdUnit)
	fmt.Println("=======================================")
	fmt.Println("Command 2")
	fmt.Println("Cmd-2 Display Name:", jObj.Batch.Jobs[1].DisplayName)
	fmt.Println("Cmn-2 Cmd Element-1:", jObj.Batch.Jobs[1].CmdElements[0].CmdUnit)


}