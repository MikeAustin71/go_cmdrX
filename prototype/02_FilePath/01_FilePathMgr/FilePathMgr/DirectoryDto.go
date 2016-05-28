package FilePathMgr

import (
	"go_cmdrX/prototype/02_FilePath/01_FilePathMgr/ErrUtil"
	fp "path/filepath"
	"fmt"
)

type DirectoryDto struct {
	Directory string
	AbsDirectoryPath string
	DirIsValid bool
}

func (d DirectoryDto) SetDirInfo(dirPath string) {
	aps, err := fp.Abs(dirPath)
	ErrUtil.SpecCheckErr("Failed To Get Absolute Path For " + dirPath, err)
	fmt.Println("Initial Dir Path: ", dirPath)
	fmt.Println("Absolute Dir Path: ", aps)
	d.AbsDirectoryPath = aps
	d.Directory = dirPath
}

func GetDirDto(dirPath string) DirectoryDto {
	d:= DirectoryDto{}
	d.SetDirInfo(dirPath)
	return d
}
