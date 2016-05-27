package DataStructs

type JsonCmdBatch struct {
	Batch CmdHdr `json:"commands_batch"`
}

type CmdHdr struct {
	Hdr  CmdHdrDat `json:"jobs_header"`
	Jobs []CmdJob  `json:"command_jobs"`
}

type CmdHdrDat struct {
	// JSON Conversion works as long as first letter of field
	//   is capitalized
	LogFileRetentionInDays int    `json:"log_file_retention_in_days"`
	CmdExeDirectory        string `json:"command_exe_directory"`
	LogPathFileName        string `json:"log_path_file_name"`
}

type CmdJob struct {
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

