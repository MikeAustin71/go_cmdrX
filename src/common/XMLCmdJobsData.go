package common

// CommandsBatch - Xml Root and Parent Element
type CommandsBatch struct {
	CmdJobsHdr CommandJobsHdr  `xml:"CommandJobsHeader"`
	CmdJobs    CommandJobArray `xml:"CommandJobs"`
}

// CommandJobsHdr - Holds base info related to
// command jobs
type CommandJobsHdr struct {
	Version                 string `xml:"Version"`
	LogFileRetentionInDays  int    `xml:"LogFileRetentionInDays"`
	KillAllJobsOnFirstError bool   `xml:"KillAllJobsOnFirstError"`
	IanaTimeZone						string `xml:"IanaTimeZone"`
	NoOfCmdJobs							int
}

// CommandJobArray - Holds individual
// CommandJob structs
type CommandJobArray struct {
	CmdJobArray []CmdJob `xml:"CommandJob"`
}

// CmdJob - Command Job information
type CmdJob struct {
	CmdDisplayName                string  `xml:"CommandDisplayName"`
	CmdDescription                string  `xml:"CommandDescription"`
	CmdType                       string  `xml:"CommandType"`
	ExeCmdInDir                   string  `xml:"ExecuteCmdInDir"`
	DelayCmdStart                 string  `xml:"DelayCmdStartSeconds"`
	StartCmdDateTime              string  `xml:"StartCmdDateTime"`
	KillJobsOnExitCodeGreaterThan int     `xml:"KillJobsOnExitCdeGreaterThan"`
	KillJobsOnExitCodeLessThan    int     `xml:"KillJobsOnExitCdeLessThan"`
	CommandTimeOutInMinutes       float64 `xml:"CmdTimeOutInMinutes"`
	ExeCommand                    string
	CmdElements                   CommandElementsArray `xml:"CmdElements"`
}

// CommandElementsArray - Holds CmdElement structures
type CommandElementsArray struct {
	CmdFragments []string `xml:"CmdElement"`
}


