package main

import (

	fp "go_cmdrX/prototype/02_FilePath/01_FilePathMgr/FilePathMgr"
	"fmt"
)

func main() {

	d:= fp.GetDirDto("./")
	fmt.Println("")
	fmt.Println("Structure Absolute Directory = ", d.AbsDirectoryPath)

}
