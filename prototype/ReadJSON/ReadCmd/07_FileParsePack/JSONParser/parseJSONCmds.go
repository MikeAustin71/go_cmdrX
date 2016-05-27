package JSONParser

import(
	ds "go_cmdrX/prototype/ReadJSON/ReadCmd/07_FileParsePack/DataStructs"
	eu "go_cmdrX/prototype/ReadJSON/ReadCmd/07_FileParsePack/ErrUtility"
	"os"
	"io"
	"encoding/json"
)


func ParseJSONCmds(fileNamePath string) ds.JsonCmdBatch {
	f, err := os.Open(fileNamePath)
	defer f.Close()
	eu.SpecCheckErr("Command File Error: " + fileNamePath + "\n", err)
	var JObj ds.JsonCmdBatch
	rdr := io.Reader(f)
	err = json.NewDecoder(rdr).Decode(&JObj)
	eu.SpecCheckErr("JSON Parsing Error Cmd File: "  + fileNamePath + "\n", err)
	return JObj
}

