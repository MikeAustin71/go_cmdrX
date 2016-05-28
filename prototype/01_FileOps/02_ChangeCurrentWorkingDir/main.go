package main

import (
	"fmt"
	"os"

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

func main() {
	d1, err := GetCurrentWorkingDirectory()

	if err != nil {
		fmt.Println("Error Getting Current Working Directory!\n", err)
		os.Exit(-1)
	}

	fmt.Println("Current Working Directory:", d1)

	err = ChangeCWD("../01_GetCurrentDir")

	if err != nil {
		fmt.Println("Error Changing Current Working Directory!\n", err)
		os.Exit(-2)
	}

	d2, err := GetCurrentWorkingDirectory()

	if err != nil {
		fmt.Println("Error Getting Current Working Directory 01_GetCurrentDir !\n", err)
		os.Exit(-3)
	}

	fmt.Println("Current Working Directory:", d2)

	err = ChangeCWD("../02_ChangeCurrentWorkingDir")

	if err != nil {
		fmt.Println("Error Changing Current Working Directory Back To Original!\n", err)
		os.Exit(-4)
	}

	d3, err := GetCurrentWorkingDirectory()

	if err != nil {
		fmt.Println("Error Getting Last Curren Working Directory (Original)")
	}

	fmt.Println("Final Current Working Directory:", d3)

}
