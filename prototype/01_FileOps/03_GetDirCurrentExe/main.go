package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// GetExeDirectory ...
// Returns the directory of the currently running
//  executable. Note this only works if you have
//  compiled code to an executable and run the exe.
func GetExeDirectory() (string, error) {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))

	if err != nil {
		return "", err
	}

	return dir, nil
}

func main() {
	d1, err := GetExeDirectory()

	if err != nil {
		fmt.Println("Error Getting Exe Directroy\n", err)
		os.Exit(-1)
	}

	fmt.Println("Exe Directory:", d1)
}
