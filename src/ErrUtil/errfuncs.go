package ErrUtil

import (
	"go_cmdrX/src/DataStrucs"
)

/*
CheckErr ests for errors and issues panic(e) in case of error
*/
func CheckErr(e error) {
	if e != nil {
		panic(e)
	}
}

/*
SpecCheckErr Prints formatted errors for DataStructs
*/
func SpecCheckErr(prefix string, err error) {
	if err == nil {
		return
	}

	e := DataStructs.SpecError{PrefixMsg: prefix, ErrMsg: err.Error()}
	panic(e)
}
