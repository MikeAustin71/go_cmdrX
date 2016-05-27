package ErrUtility

import (
	"go_cmdrX/prototype/ReadJSON/ReadCmd/07_FileParsePack/DataStructs"
)


func CheckErr(e error) {
	if e != nil {
		panic(e)
	}
}

func SpecCheckErr (prefix string, err error){
	if(err == nil){
		return
	}
	e := DataStructs.SpecError{prefix, err.Error()}
	panic(e)
}
