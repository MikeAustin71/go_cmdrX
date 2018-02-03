package common

/*  'logmsgtype.go' is located in source code
		repository:

		https://github.com/MikeAustin71/logopsgo.git

*/

// LogMsgType - Designates type of Message being logged
type LogMsgType int

func (mtype LogMsgType) String() string {
	return LogMsgTypeNames[mtype]
}

const (
	// LogERRORMSGTYPE - Message Type
	LogERRORMSGTYPE LogMsgType = iota
	// LogINFOMSGTYPE - Information Message Type
	LogINFOMSGTYPE
)

// LogMsgTypeNames - String Array holding Message Type names.
var LogMsgTypeNames = [...]string{"ERROR", "INFO"}
