package FilePathMgr

import (
	"os"
	"path/filepath"
)

// Changes the Current Working Directory
//  to directory parameter 'dirPath'
func ChangeCWD(dirPath string) error {
	err:= os.Chdir(dirPath)
	return err
}

// Returns the current working directory
func GetCurrentWorkingDirectory() (string, error) {
	dir, err := os.Getwd()

	if err != nil {
		return "", err
	}

	return dir, nil

}

// Returns the directory of the currently running
//  executable
func GetExeDirectory() (string, error) {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))

	if err != nil {
		return "", err
	}

	return dir, nil
}

