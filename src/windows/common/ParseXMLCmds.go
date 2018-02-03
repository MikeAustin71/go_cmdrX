package common

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
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
func ParseXML(xmlPathFileName string, parent []OpsMsgContextInfo) (CommandBatch, OpsMsgDto) {

	//bi := ErrBaseInfo{}.New(parseXMLSrcFile, "ParseXML", parseSMLErrBlockNo)
	msgCtx := OpsMsgContextInfo{
						SourceFileName: parseXMLSrcFile,
						ParentObjectName:"",
						FuncName:"ParseXML",
						BaseMessageId: parseSMLErrBlockNo,
						}


	om := OpsMsgDto{}.InitializeAllContextInfo(parent, msgCtx)

	var cmds CommandBatch

	xmlFile, err1 := os.Open(xmlPathFileName)

	if err1 != nil {
		s := fmt.Sprintf("File Name: %v  - Error opening file: %v", xmlPathFileName, err1.Error())
		om.SetFatalError(s, err1, 1005)
		return cmds, om
	}

	defer xmlFile.Close()

	b, err2 := ioutil.ReadAll(xmlFile)

	if err2 != nil {
		s := fmt.Sprintf("Error Reading XML File: %v. Error - %v", xmlPathFileName, err2.Error())
		om.SetFatalError(s,err2,1006)
		return cmds, om
	}

	xml.Unmarshal(b, &cmds)

	om.SetSuccessfulCompletionMessage("Finished Function: ParseXML", 1009)
	return cmds, om
}
