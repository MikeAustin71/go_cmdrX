package common

// OpsMsgDto - Data Transfer Object
// containing information about an
// operations message
type OpsMsgDto struct {
	message  []string
	msgType  LogMsgType
	msgLevel LogLevel
}
