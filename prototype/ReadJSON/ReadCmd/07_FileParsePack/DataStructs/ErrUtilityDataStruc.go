package DataStructs

type SpecError struct{
	PrefixMsg string
	ErrMsg string
}

func (s SpecError) Error() string {
	return s.PrefixMsg + s.ErrMsg
}

