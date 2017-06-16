package common

import "strings"

// AssembleCmdElements - Assembles
// Command Elements, or Command Fragments,
// and stores them in struct
// CommandJob.ExeCommand
func AssembleCmdElements(cmds *CommandsBatch) {
	var exCmd, s string

	cmds.CmdJobsHdr.NoOfCmdJobs = len(cmds.CmdJobs.CmdJobArray)

	for idx, cmdJob := range cmds.CmdJobs.CmdJobArray {
		exCmd = ""
		for _, cmdElement := range cmdJob.CmdElements.CmdFragments {
			s = strings.TrimRight(strings.TrimLeft(cmdElement, " "), " ")
			exCmd += (s + " ")
		}

		cmds.CmdJobs.CmdJobArray[idx].ExeCommand = strings.TrimRight(exCmd, " ")

	}

	return
}
