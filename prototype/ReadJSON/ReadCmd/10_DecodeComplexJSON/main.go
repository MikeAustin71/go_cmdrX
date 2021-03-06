package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type jsonobject struct {
	Batch cmdHdr `json:"commands_batch"`
}

type cmdHdr struct {
	Hdr  cmdHdrDat `json:"jobs_header"`
	Jobs []cmdJob  `json:"command_jobs"`
}

type cmdHdrDat struct {
	// JSON Conversion works as long as first letter of field
	//   is capitalized
	LogFileRetentionInDays int    `json:"log_file_retention_in_days"`
	CmdExeDirectory        string `json:"command_exe_directory"`
	LogPathFileName        string `json:"log_path_file_name"`
}

type cmdJob struct {
	DisplayName               string       `json:"cmd_display_name"`
	Desc                      string       `json:"cmd_description"`
	Type                      string       `json:"cmd_type"`
	ExeDir                    string       `json:"execute_cmd_in_dir"`
	DelayStartSecs            string       `json:"delay_cmd_start_seconds"`
	StartAtDateTime           string       `json:"start_cmd_date_time"`
	KillOnExitCodeGreaterThan string       `json:"kill_jobs_on_exit_code_greater_than"`
	KillOnExitCodeLessThan    string       `json:"kill_jobs_on_exit_code_less_than"`
	TimeOutMinutes            string       `json:"cmd_timeout_in_minutes"`
	CmdElements               []CmdElement `json:"cmd_elements"`
}

type CmdElement struct {
	CmdUnit string `json:"cmdelement"`
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	f, err := os.Open("./CmdrX_JSON_010.json")
	defer f.Close()
	check(err)
	var JObj jsonobject
	rdr := io.Reader(f)
	err = json.NewDecoder(rdr).Decode(&JObj)
	check(err)
	fmt.Println(JObj)
	fmt.Println("=======================================")
	fmt.Println("Header")
	fmt.Println("Log File Retention In Days:", JObj.Batch.Hdr.LogFileRetentionInDays)
	fmt.Println("=======================================")
	fmt.Println("Command 1")
	fmt.Println("Cmd-1 Display Name:", JObj.Batch.Jobs[0].DisplayName)
	fmt.Println("Cmn-1 Cmd Element-1:", JObj.Batch.Jobs[0].CmdElements[0].CmdUnit)
	fmt.Println("=======================================")
	fmt.Println("Command 2")
	fmt.Println("Cmd-2 Display Name:", JObj.Batch.Jobs[1].DisplayName)
	fmt.Println("Cmn-2 Cmd Element-1:", JObj.Batch.Jobs[1].CmdElements[0].CmdUnit)

}
