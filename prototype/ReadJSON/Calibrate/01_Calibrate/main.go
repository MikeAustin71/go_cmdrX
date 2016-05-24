package main

import (
	"fmt"
)

type cmdHdrDat struct {
	logFileRetentionInDays    int    `json:"defaultlogfileretentionindays"`
	defaultCmdExeDir          string `json:"defaultcommandexedirectory"`
	defaultCmdLogPathFileName string `json:"default_cmd_log_path_file_name"`
}

func main() {
	n := "cmdrX"
	fmt.Println("Hello World - ", n)
}
