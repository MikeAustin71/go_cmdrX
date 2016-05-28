package main

import (
	"fmt"
	"os"
)

func GetCurrentWorkingDirectory() (string, error) {
	dir, err := os.Getwd()

	if err != nil {
		return "" , err
	}

	return dir , nil

}


func main() {
	dir, err := GetCurrentWorkingDirectory()
	if err != nil {
		fmt.Println("Error on os.Getwd() ", err)
		os.Exit(1)
	}

	fmt.Println("Current Working Directory:", dir)

}
