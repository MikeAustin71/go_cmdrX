package ErrUtil

type SpecError struct{
	PrefixMsg string
	ErrMsg string
}

func (s SpecError) Error() string {
	return s.PrefixMsg +"\n"+ s.ErrMsg
}


func CheckErr(e error) {
	if e != nil {
		panic(e)
	}
}

func SpecCheckErr (prefix string, err error){
	if(err == nil){
		return
	}
	e := SpecError{prefix, err.Error()}
	panic(e)
}
