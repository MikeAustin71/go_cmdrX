package common

import (
	"encoding/xml"
	"fmt"
)

/*  'ParseXMLCmds.go' is located in source code
		repository:

		https://github.com/MikeAustin71/go_cmdrX.git

*/


const (
	parseXMLSrcFile    = "ParseXMLCmds.go"
	parseSMLErrBlockNo = int64(219610000)
)

// ParseXML - Reads and Parses command
// data from an XML file.
func ParseXML(xmlCmdFile FileMgr, parent []OpsMsgContextInfo) (CommandBatch, OpsMsgDto) {

	//bi := ErrBaseInfo{}.New(parseXMLSrcFile, "ParseXML", parseSMLErrBlockNo)
	msgCtx := OpsMsgContextInfo{
							SourceFileName: parseXMLSrcFile,
							ParentObjectName:"",
							FuncName:"ParseXML",
							BaseMessageId: parseSMLErrBlockNo,
						}


	om := OpsMsgDto{}.InitializeAllContextInfo(parent, msgCtx)

	cmds := CommandBatch{}

	err1 := xmlCmdFile.IsFileMgrValid("")

	if err1 != nil {
		s := fmt.Sprintf("Error: XML Command File Manager is INVALID. xmlCmdFile='%v'\n", xmlCmdFile.AbsolutePathFileName)
		om.SetFatalError(s, err1, 1001)
		return cmds, om

	}

	doesFileExist, err1 := xmlCmdFile.DoesThisFileExist()

	if err1 != nil {
		s := fmt.Sprintf("Error returned by xmlCmdFile.DoesThisFileExist(). xmlCmdFile='%v'\n", xmlCmdFile.AbsolutePathFileName)
		om.SetFatalError(s, err1, 1003)
		return cmds, om
	}

	if !doesFileExist {
		s := "Error: xmlCmdFile DOES NOT EXIST!"
		err1 = fmt.Errorf("xmlCmdFile DOES NOT EXIST! xmlCmdFile='%v'\n", xmlCmdFile.AbsolutePathFileName)
		om.SetFatalError(s, err1, 1005)
		return cmds, om
	}


	b, err1 :=  xmlCmdFile.ReadAllFile()

	if err1 != nil {
		s := fmt.Sprintf("Error returned by xmlCmdFile.ReadAllFile(). xmlCmdFile='%v'.\n", xmlCmdFile.AbsolutePathFileName)
		om.SetFatalError(s,err1,1007)
		return cmds, om
	}

	err1 = xml.Unmarshal(b, &cmds)

	if err1 != nil {
		s := fmt.Sprintf("Error returned by xml.Unmarshal(b, &cmds). xmlCmdFile='%v'.\n", xmlCmdFile.AbsolutePathFileName)
		om.SetFatalError(s,err1,1009)
		return cmds, om
	}

	om.SetSuccessfulCompletionMessage("Finished Function: ParseXML", 1099)
	return cmds, om
}
