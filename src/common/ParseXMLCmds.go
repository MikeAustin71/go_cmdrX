package common

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

// ParseXML - Reads and Parses command
// data from an XML file.
func ParseXML(xmlPathFileName string) CommandsBatch {

	var cmds CommandsBatch

	xmlFile, err1 := os.Open(xmlPathFileName)

	if err1 != nil {
		fmt.Println("File Name: ", xmlPathFileName)
		fmt.Println("Error opening file:")
		panic(err1)
	}

	defer xmlFile.Close()

	b, err2 := ioutil.ReadAll(xmlFile)

	if err2 != nil {
		fmt.Println("Error Reading XML File")
		fmt.Println("XML File: ", xmlPathFileName)
		panic(err2)
	}

	xml.Unmarshal(b, &cmds)

	return cmds
}
