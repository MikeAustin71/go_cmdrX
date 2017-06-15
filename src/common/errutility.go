package common

import (
	"fmt"
)

// ErrBaseInfo is intended for use with
// the SpecErr Structure. It sets up base
// error info to be used repeatedly.
type ErrBaseInfo struct {
	SourceFileName string
	FuncName       string
	BaseErrorID    int64
}

// New - returns a new, populated ErrBaseInfo Structure
func (b ErrBaseInfo) New(srcFile, funcName string, baseErrID int64) ErrBaseInfo {
	return ErrBaseInfo{SourceFileName: srcFile, FuncName: funcName, BaseErrorID: baseErrID}
}

// NewFunc - Returns a New ErrBaseInfo structure with a different Func Name
func (b ErrBaseInfo) NewFunc(funcName string) ErrBaseInfo {
	return ErrBaseInfo{SourceFileName: b.SourceFileName, FuncName: funcName, BaseErrorID: b.BaseErrorID}
}

// NewBaseInfo - returns a deep copy of the current
// ErrBaseInfo structure.
func (b ErrBaseInfo) NewBaseInfo() ErrBaseInfo {
	return ErrBaseInfo{SourceFileName: b.SourceFileName, FuncName: b.FuncName, BaseErrorID: b.BaseErrorID}
}

// DeepCopyBaseInfo - Same as NewBaseInfo()
func (b ErrBaseInfo) DeepCopyBaseInfo() ErrBaseInfo {
	return ErrBaseInfo{SourceFileName: b.SourceFileName, FuncName: b.FuncName, BaseErrorID: b.BaseErrorID}
}

// GetBaseSpecErr - Returns an empty
// SpecErr structure populated with
// Base Error Info
func (b ErrBaseInfo) GetBaseSpecErr() SpecErr {

	return SpecErr{BaseInfo: b.NewBaseInfo()}
}

// GetNewParentInfo - Returns a slice of ErrBaseInfo
// structures with the first element initialized to a
// new ErrBaseInfo structure.
func (b ErrBaseInfo) GetNewParentInfo(srcFile, funcName string, baseErrID int64) []ErrBaseInfo {
	var parent []ErrBaseInfo

	bi := b.New(srcFile, funcName, baseErrID)

	return append(parent, bi)
}

// SpecErr - A data structure used
// to hold custom error information
type SpecErr struct {
	ParentInfo []ErrBaseInfo
	BaseInfo   ErrBaseInfo
	IsErr      bool
	IsPanic    bool
	PrefixMsg  string
	ErrMsg     string
	ErrNo      int64
}

// InitializeBaseInfo - Initializes a SpecErr Structure
// from a ParentInfo array and a ErrBaseInfo
// structure
func (s SpecErr) InitializeBaseInfo(parent []ErrBaseInfo, bi ErrBaseInfo) SpecErr {

	return SpecErr{
		ParentInfo: s.DeepCopyParentInfo(parent),
		BaseInfo:   bi.DeepCopyBaseInfo()}
}

// Initialize - Initializes all elements of
// the SpecErr structure
func (s SpecErr) Initialize(parent []ErrBaseInfo, bi ErrBaseInfo, prefix string, err error, isPanic bool, errNo int64) SpecErr {
	return s.InitializeBaseInfo(parent, bi).New(prefix, err, isPanic, errNo)

}

// New - Creates new SpecErr Type. Uses existing
// Parent and ErrBaseInfo data
func (s SpecErr) New(prefix string, err error, isPanic bool, errNo int64) SpecErr {

	x := SpecErr{
		ParentInfo: s.DeepCopyParentInfo(s.ParentInfo),
		BaseInfo:   s.BaseInfo.DeepCopyBaseInfo(),
		PrefixMsg:  prefix,
		IsPanic:    isPanic}

	x.ErrNo = errNo + x.BaseInfo.BaseErrorID

	if err != nil {
		x.ErrMsg = err.Error()
		x.IsErr = true
	} else {
		x.ErrMsg = ""
		x.IsErr = false
		x.IsPanic = false
	}

	return x
}

// SignalNoErrors - Returns a SpecErr
// structure with IsErr set to false.
func (s SpecErr) SignalNoErrors() SpecErr {
	return SpecErr{IsErr: false, IsPanic: false}
}


// DeepCopyBaseInfo - Returns a deep copy of the
// current BaseInfo structure.
func (s SpecErr) DeepCopyBaseInfo() ErrBaseInfo {
	return s.BaseInfo.DeepCopyBaseInfo()
}

// SetBaseInfo - Sets the SpecErr ErrBaseInfo internal
// structure. This data is used for creating repetitive
// error information.
func (s SpecErr) SetBaseInfo(bi ErrBaseInfo) {
	s.BaseInfo = bi.NewBaseInfo()
}

// SetParentInfo - Sets the ParentInfo Slice for
// the current SpecErr structure
func (s SpecErr) SetParentInfo(parent []ErrBaseInfo) {
	if len(parent) == 0 {
		return
	}

	s.ParentInfo = s.DeepCopyParentInfo(parent)
}

// AddParentInfo - Adds ParentInfo elements to
// the current SpecErr ParentInfo slice
func (s SpecErr) AddParentInfo(parent []ErrBaseInfo) {
	if len(parent) == 0 {
		return
	}

	x := s.DeepCopyParentInfo(parent)

	for _, bi := range x {
		s.ParentInfo = append(s.ParentInfo, bi.NewBaseInfo())
	}

	return

}

// AddBaseToParentInfo - Adds the structure's
// ErrBaseInfo data to ParentInfo and returns a
// new ParentInfo Array
func (s SpecErr) AddBaseToParentInfo() []ErrBaseInfo {

	a := s.DeepCopyParentInfo(s.ParentInfo)
	return append(a, s.BaseInfo.DeepCopyBaseInfo())
}

// DeepCopyParentInfo - Receives an array of slices
// type ErrBaseInfo and appends deep copies
// of those slices to the SpecErr ParentInfo
// field.
func (s SpecErr) DeepCopyParentInfo(pi []ErrBaseInfo) []ErrBaseInfo {

	if len(pi) == 0 {
		return pi
	}

	a := make([]ErrBaseInfo, 0, len(pi)+10)
	for _, bi := range pi {
		a = append(a, bi.NewBaseInfo())
	}

	return a
}

// Panic - Executes 'panic' command
// if IsPanic == 'true'
func (s SpecErr) Panic() {
	if s.IsPanic {
		panic(s)
	}
}

// Error - Implements Error Interface
func (s SpecErr) Error() string {

	m := "\nError Message:"
	m += "\n---------------------"
	if s.PrefixMsg != "" {
		m += "\n"
		m += s.PrefixMsg
	}

	m += "\n" + s.ErrMsg
	m += "\n---------------------"

	if s.BaseInfo.SourceFileName != "" {
		m += "\nSourceFile: " + s.BaseInfo.SourceFileName
	}

	if s.BaseInfo.FuncName != "" {
		m += "\nFuncName: " + s.BaseInfo.FuncName
	}

	m += fmt.Sprintf("\nErrNo: %v", s.ErrNo)
	m += fmt.Sprintf("\nIsErr: %v", s.IsErr)
	m += fmt.Sprintf("\nIsPanic: %v", s.IsPanic)

	// If parent Function Info Exists
	// Print it out.
	if len(s.ParentInfo) > 0 {
		m += "\n---------------------"
		m += "\n  Parent Func Info"
		m += "\n---------------------"

		for _, bi := range s.ParentInfo {
			m += "\n" + bi.SourceFileName + "-" + bi.FuncName
			if bi.BaseErrorID != 0 {
				m += fmt.Sprintf(" ErrorID: %v", bi.BaseErrorID)
			}
		}
	}

	return m
}

var blankErrBaseInfo = ErrBaseInfo{}
var blankParentInfo = make([]ErrBaseInfo, 0, 10)

// CheckErrPanic - Checks for error and then
// executes 'panic'
func CheckErrPanic(e error) {
	if e != nil {
		panic(e)
	}
}

// CheckIsSpecErr - If error is present,
// returns 'true'.  If NO Error, returns
// 'false'.
func CheckIsSpecErr(eSpec SpecErr) bool {

	if eSpec.IsErr {
		return true
	}

	return false

}

// PanicOnSpecErr - Issues a 'panic'
// command if SpecErr IsPanic flag is set
func PanicOnSpecErr(eSpec SpecErr) {

	if eSpec.IsPanic {
		panic(eSpec)
	}
}
