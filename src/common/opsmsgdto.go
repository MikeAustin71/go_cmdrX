package common

import (
	"time"
	"strings"
	"fmt"
	"errors"
	"unicode/utf8"
	"os"
)

/*  'opsmsgdto.go' is located in source code
		repository:

		https://github.com/MikeAustin71/ErrHandlerGo.git

		Dependencies:
		=============

			timezoneutility.go - is part of the date time operations library. The source code repository
 														for this file is located at:

														https://github.com/MikeAustin71/datetimeopsgo.git

		NOTE:
		=====
		The 'OpsMsgDto' type supersedes and replaces the 'SpecErr' type also
		found in this source code repository.


		Message Formats:
		================
		Depending on the setting for OpsMsgDto field 'UseFormattedMsg', two
		types of message formatting are employed. This is an example of a
		fully formatted message for display output:
   --------------------------------------------------------------------------------

   !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
   FATAL ERROR Message                             								 Error No: 6974
   !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
   Is Error: true       Is Panic\Fatal Error: true
   ------------------------------------------------------------------------------
   Message: This is FATAL Error message text.
   ------------------------------------------------------------------------------
   Parent Context History:
   SrcFile: TSource01 -ParentObj: PObj01 -FuncName: Func001 -BaseMsgId: 1000
   SrcFile: TSource02 -ParentObj: PObj02 -FuncName: Func002 -BaseMsgId: 2000
   SrcFile: TSource03 -ParentObj: PObj03 -FuncName: Func003 -BaseMsgId: 3000
   SrcFile: TSource04 -ParentObj: PObj04 -FuncName: Func004 -BaseMsgId: 4000
   SrcFile: TSource05 -ParentObj: PObj05 -FuncName: Func005 -BaseMsgId: 5000
   ------------------------------------------------------------------------------
   Current Message Context:
   SrcFile: TSource06 -ParentObj: PObj06 -FuncName: Func006 -BaseMsgId: 6000
   ------------------------------------------------------------------------------
     Message Time UTC: 2017-12-16 Sat 22:12:28.458551100 +0000 UTC
   Message Time Local: 2017-12-16 Sat 16:12:28.458551100 -0600 CST
   !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!

   --------------------------------------------------------------------------------
		The second type of message output mode is an abbreviated message text. This is
		an example of an abbreviated message text
   --------------------------------------------------------------------------------

	 FATAL ERROR Msg No: 972 - 12/16/2017 16:18:24.904997000 -0600 CST - Test Serious Error.

   --------------------------------------------------------------------------------
		Reference Method OpsMsgDto.SetMessageOutputMode() below for additional information.

*/

// OpsMsgCollection - A collection of Operations Message Dto
// objects
type OpsMsgCollection struct {
	OpsMessages [] OpsMsgDto
}

// AddOpsMsg - Adds an Operations Message (OpsMsgDto object)
// to the end of the OpsMessages array.
func (omc *OpsMsgCollection) AddOpsMsg(opsMsg OpsMsgDto) {
	omc.OpsMessages = append(omc.OpsMessages, opsMsg)
}

// AddOpsMsgCollection - Adds another collection of OpsMsgDto
// objects to the current collection.
func (omc *OpsMsgCollection) AddOpsMsgCollection(omc2 *OpsMsgCollection) {

	lOmc2 := len(omc2.OpsMessages)

	if lOmc2 == 0 {
		return
	}

	for i:=0; i< lOmc2 ; i++ {
		omc.AddOpsMsg(omc2.OpsMessages[i].CopyOut())
	}

	return
}

// CopyOut - Returns an OpsMsgCollection which is an
// exact duplicate of the current OpsMsgCollection
func (omc *OpsMsgCollection) CopyOut() OpsMsgCollection {
	omc2 := OpsMsgCollection{}

	lOmc := len(omc.OpsMessages)

	for i:= 0; i < lOmc; i++ {
		omc2.AddOpsMsg(omc.OpsMessages[i].CopyOut())
	}

	return omc2
}

// GetArrayLength - returns the array length of the
// OpsMessages array.
func (omc *OpsMsgCollection) GetArrayLength() int {
	return len(omc.OpsMessages)
}

// PopLastMsg - Removes the last OpsMsgDto object
// from the collections array, and returns it to
// the calling method.
func (omc *OpsMsgCollection) PopLastMsg() OpsMsgDto {

	l1 := len(omc.OpsMessages)

	om := omc.OpsMessages[l1-1].CopyOut()

	omc.OpsMessages = omc.OpsMessages[0:l1-1]

	return om
}

// PopFirstMsg - Removes the first OpsMsgDto object
// from the collections array, and returns it to
// the calling method.
func (omc *OpsMsgCollection) PopFirstMsg() OpsMsgDto {

	l1 := len(omc.OpsMessages)

	om := omc.OpsMessages[0].CopyOut()

	omc.OpsMessages = omc.OpsMessages[1:l1]

	return om
}

// PeekFirstMsg - Returns the first element from the
// Operation Messages Collection, but does NOT remove
// it from the OpsMessages array.
func (omc *OpsMsgCollection) PeekFirstMsg() OpsMsgDto {

	return omc.OpsMessages[0].CopyOut()
}

// PeekLastMsg - Returns the last element from the
// Operation Messages Collection, but does NOT remove
// it from the OpsMessages array.
func (omc *OpsMsgCollection) PeekLastMsg() OpsMsgDto {

	l1 := len(omc.OpsMessages)

	return omc.OpsMessages[l1-1].CopyOut()
}


// OpsMsgType - Designates type of Message being logged
type OpsMsgType int

// String - Method used to display the text
// name of an Operations Message Type.
func (opstype OpsMsgType) String() string {
	return OpsMsgTypeNames[opstype]
}

const (

	// OpsMsgTypeNOERRORNOMSG - 0 Uninitialized -
	// no errors and no messages
	OpsMsgTypeNOERRORNOMSGS OpsMsgType = iota

	// OpsMsgTypeERRORMSG - 1 Error Message
	OpsMsgTypeERRORMSG

	// OpsMsgTypeFATALERRORMSG 2 Fatal Error Message
	OpsMsgTypeFATALERRORMSG

	// OpsMsgTypeINFOMSG - 3 Information Message Type
	OpsMsgTypeINFOMSG

	// OpsMsgTypeWARNINGMSG - 4 Warning Message Type
	OpsMsgTypeWARNINGMSG

	// OpsMsgTypeDEBUGMSG - 5 Debug Message
	OpsMsgTypeDEBUGMSG

	// OpsMsgTypeSUCCESSFULCOMPLETION - 6 Message signalling
	// successful completion of the operation.
	OpsMsgTypeSUCCESSFULCOMPLETION

)

// OpsMsgTypeNames - String Array holding Message Type names.
var OpsMsgTypeNames = [...]string{"NOERRORSNOMSGS","ERROR", "FATALERROR", "INFO", "WARNING", "DEBUG", "SUCCESS"}


/*
// OpsMsgClass - Holds the Message level indicating the relative importance of a specific log Message.
type OpsMsgClass int

// String - Returns the name of the OpsMsgClass element
// formatted as a string.
func (opsmsgclass OpsMsgClass) String() string {
	return OpsMsgClassNames[opsmsgclass]
}


const (
	// OpsMsgClassNOERRORSNOMESSAGES - 0 Signals uninitialized message
	// with no errors and no messages
	OpsMsgClassNOERRORSNOMESSAGES OpsMsgClass = iota

	// OpsMsgClassOPERROR - 1 Message is an Error Message
	OpsMsgClassOPERROR

	// OpsMsgClassFATAL - 2 Message is a Fatal Error Message
	OpsMsgClassFATAL

	// OpsMsgClassINFO - 3 Message is an Informational Message
	OpsMsgClassINFO

	// OpsMsgClassWARNING - 4 Message is a warning Message
	OpsMsgClassWARNING

	// OpsMsgClassDEBUG - 5 Message is a Debug Message
	OpsMsgClassDEBUG

	// OpsMsgClassSUCCESSFULCOMPLETION - 6 Message signalling successful
	// completion of the operation
	OpsMsgClassSUCCESSFULCOMPLETION
)

// OpsMsgClassNames - string array containing names of Log Levels
var OpsMsgClassNames = [...]string{"NOERRORSNOMESSAGES", "OPERROR", "FATAL", "INFO", "WARNING", "DEBUG", "SUCCESS"}

*/

// OpsMsgContextInfo - Contains context information describing
// the current environment in which the message was generated.
type OpsMsgContextInfo struct {
	SourceFileName   string
	ParentObjectName string
	FuncName         string
	BaseMessageId    int64
}

// New - returns a new, populated OpsMsgContextInfo Structure
func (ci OpsMsgContextInfo) New(srcFile, parentObject, funcName string, baseMsgID int64) OpsMsgContextInfo {
	return OpsMsgContextInfo{SourceFileName: srcFile, ParentObjectName:parentObject, FuncName: funcName, BaseMessageId: baseMsgID}
}

// NewFuncName - Returns a New OpsMsgContextInfo structure with a different Func Name
func (ci OpsMsgContextInfo) NewFuncName(funcName string) OpsMsgContextInfo {
	return OpsMsgContextInfo{SourceFileName: ci.SourceFileName, FuncName: funcName, BaseMessageId: ci.BaseMessageId}
}

// NewOpsMsgContextInfo - returns a deep copy of the current
// OpsMsgContextInfo structure.
func (ci OpsMsgContextInfo) NewOpsMsgContextInfo() OpsMsgContextInfo {
	return OpsMsgContextInfo{SourceFileName: ci.SourceFileName, ParentObjectName: ci.ParentObjectName, FuncName: ci.FuncName, BaseMessageId: ci.BaseMessageId}
}

// DeepCopyOpsMsgContextInfo - Same as NewOpsMsgContextInfo()
func (ci OpsMsgContextInfo) DeepCopyOpsMsgContextInfo() OpsMsgContextInfo {
	return OpsMsgContextInfo{SourceFileName: ci.SourceFileName, ParentObjectName: ci.ParentObjectName, FuncName: ci.FuncName, BaseMessageId: ci.BaseMessageId}
}

// Equal - Compares two OpsMsgContextInfo objects
// to determine if they are equivalent.
func (ci *OpsMsgContextInfo) Equal(ci2 *OpsMsgContextInfo) bool {
	if ci.SourceFileName 		!= ci2.SourceFileName 	||
			ci.ParentObjectName != ci2.ParentObjectName ||
			ci.FuncName 				!= ci2.FuncName					||
			ci.BaseMessageId    != ci2.BaseMessageId		{

				return false
	}

	return true
}

// GetBaseOpsMsgDto - Returns an empty
// OpsMsgDto structure populated with
// Base Message Context Information
func (ci OpsMsgContextInfo) GetBaseOpsMsgDto() OpsMsgDto {

	return OpsMsgDto{MsgContext: ci.NewOpsMsgContextInfo()}
}

// GetNewParentInfo - Returns a slice of OpsMsgContextInfo
// structures with the first element initialized to a
// new OpsMsgContextInfo structure.
func (ci OpsMsgContextInfo) GetNewParentInfo(srcFile, parentObject, funcName string, baseMsgID int64) []OpsMsgContextInfo {
	var parent []OpsMsgContextInfo

	newCi := ci.New(srcFile, parentObject, funcName, baseMsgID)

	return append(parent, newCi)
}


// OpsMsgDto - Data Transfer Object
// containing information about an
// operations Message
type OpsMsgDto struct {
	ParentContextHistory 	[] OpsMsgContextInfo // Function tree showing the execution path leading to this method
	MsgContext           	OpsMsgContextInfo
	Message              	string // The original message sent to OpsMsgDto
	MsgType              	OpsMsgType
	MsgTimeUTC           	time.Time
	MsgTimeLocal         	time.Time
	MsgLocalTimeZone     	string
	UseFormattedMsg			 	bool		// If true, message output methods will return the fully formatted message.
															// If false, message output methods will return an abbreviated from of the message
															// By default, the fully formatted version of the message will be displayed.
	fmtMessage   					string   // The formatted message
	abbrvMessage 					string   // An Abbreviated Form of the message
	msgId        					int64    // The identifying number for this message
	msgNumber    					int64    //  Message Number = msgId + MsgContext.BaseMessageId. This is the number displayed in the message

}


// AddParentContextHistory - Adds ParentInfo elements to
// the current SpecErr ParentInfo slice
func (opsMsg *OpsMsgDto) AddParentContextHistory(parent []OpsMsgContextInfo) {

	if len(parent) == 0 {
		return
	}

	x := opsMsg.DeepCopyParentContextHistory(parent)

	for _, bi := range x {
		opsMsg.ParentContextHistory = append(opsMsg.ParentContextHistory, bi.NewOpsMsgContextInfo())
	}

	return
}

// AddOpsMsgContextInfoToParentHistory - Adds an OpsMsgContextInfo object to the Parent
// Context History maintained by the current or host OpsMsgDto object.
func (opsMsg *OpsMsgDto) AddOpsMsgContextInfoToParentHistory(newContextInfo OpsMsgContextInfo) {
	ci := newContextInfo.DeepCopyOpsMsgContextInfo()
	opsMsg.ParentContextHistory = append(opsMsg.ParentContextHistory, ci)
}


// ChangeMsg - Change the existing message text.
// Note: this will NOT change the message type.
// Only the message text is affected.
func (opsMsg *OpsMsgDto) ChangeMsg(msg string) {

	opsMsg.setMessageText(msg, opsMsg.msgId)

}

// ChangeMsgId - Change the existing message Id.
//
// Input Parameter:
// ================
//
// msgId int64		- The message Id number to be associated with this message. If this
//									Id number is set to zero ('0'), no message number will be displayed
//									in the final message text.
//
//									Message Id differs from Message Number. The final Message Number is
//									calculated by adding Message Id to the opsMsg.MsgContext.BaseMessageId.
//									Again, if Message Id is zero ('0'), no message number is displayed in
//									the final message text.
//
func (opsMsg *OpsMsgDto) ChangeMsgId(msgId int64) {
	opsMsg.setMsgIdAndMsgNumber(msgId)
}

// ConfigureContextHistoryFromParentOpsMsgDto - Receives an OpsMsgDto object as
// an input parameter ('parentOpsMsgDto'). The parent context history from this
// second OpsMsgDto object ('parentOpsMsgDto') is added to the parent history
// of the current or host OpsMsgDto Object. In addition, the MsgContext object
// from 'parentOpsMsgDto' is also added to the parent history of the current
// or host OpsMsgDto Object.
func (opsMsg *OpsMsgDto) ConfigureContextHistoryFromParentOpsMsgDto(parentOpsMsgDto OpsMsgDto) {
	opsMsg.AddParentContextHistory(parentOpsMsgDto.ParentContextHistory)
	opsMsg.ParentContextHistory = append(opsMsg.ParentContextHistory, parentOpsMsgDto.DeepCopyMsgContext())

}

// ConfigureMessageContext - Sets MsgContext for the current or host OpsMsgDto object
// to the input parameter 'newMsgContext'
func (opsMsg *OpsMsgDto) ConfigureMessageContext(newMsgContext OpsMsgContextInfo) {
	opsMsg.MsgContext = newMsgContext.DeepCopyOpsMsgContextInfo()
}

// CopyIn - Receives an OpsMsgDto object as input.
// Then a deep copy is created and used to populate
// the current OpsMsgDto object.
func (opsMsg *OpsMsgDto) CopyIn(opsMsg2 *OpsMsgDto) {
	opsMsg.Empty()
	opsMsg.ParentContextHistory = opsMsg2.DeepCopyParentContextHistory(opsMsg2.ParentContextHistory)
	opsMsg.MsgContext 			= opsMsg2.MsgContext.DeepCopyOpsMsgContextInfo()
	opsMsg.Message       		= opsMsg2.Message
	opsMsg.MsgType          = opsMsg2.MsgType
	opsMsg.MsgTimeUTC       = opsMsg2.MsgTimeUTC
	opsMsg.MsgTimeLocal     = opsMsg2.MsgTimeLocal
	opsMsg.MsgLocalTimeZone = opsMsg2.MsgLocalTimeZone
	opsMsg.msgId            = opsMsg2.GetMessageId()
	opsMsg.msgNumber        = opsMsg2.GetMessageNumber()
	opsMsg.fmtMessage 			=	opsMsg2.GetFmtMessage()
	opsMsg.abbrvMessage			= opsMsg2.GetAbbrvMessage()

}

// CopyOut - Makes a deep copy of the current OpsMsgDto
// and returns it as a new OpsMsgDto with equivalent
// field values.
func (opsMsg *OpsMsgDto) CopyOut() OpsMsgDto {
	
	opsMsg2 := OpsMsgDto{}

	opsMsg2.ParentContextHistory = opsMsg.DeepCopyParentContextHistory(opsMsg.ParentContextHistory)
	opsMsg2.MsgContext = opsMsg.MsgContext.DeepCopyOpsMsgContextInfo()

	opsMsg2.Message       		= opsMsg.Message
	opsMsg2.MsgType          	= opsMsg.MsgType
	opsMsg2.MsgTimeUTC       	= opsMsg.MsgTimeUTC
	opsMsg2.MsgTimeLocal     	= opsMsg.MsgTimeLocal
	opsMsg2.MsgLocalTimeZone 	= opsMsg.MsgLocalTimeZone
	opsMsg2.fmtMessage 				= opsMsg.GetFmtMessage()
	opsMsg2.abbrvMessage			= opsMsg.GetAbbrvMessage()
	opsMsg2.msgId            	= opsMsg.GetMessageId()
	opsMsg2.msgNumber        	= opsMsg.GetMessageNumber()

	return opsMsg2
}

// DeepCopyOpsMsgContextInfo - Returns a deep copy of the
// current MsgContext (OpsMsgContextInfo structure).
func (opsMsg *OpsMsgDto) DeepCopyMsgContext() OpsMsgContextInfo {
	return opsMsg.MsgContext.DeepCopyOpsMsgContextInfo()
}


// DeepCopyParentContextHistory - Receives an array of slices
// type OpsMsgContextInfo and appends deep copies
// of those slices to the OpsMsgDto ParentContextHistory
// field.
func (opsMsg *OpsMsgDto) DeepCopyParentContextHistory(pi []OpsMsgContextInfo) []OpsMsgContextInfo {

	if len(pi) == 0 {
		return pi
	}

	newHistory := make([]OpsMsgContextInfo, 0, len(pi)+10)
	for _, ci := range pi {
		newHistory = append(newHistory, ci.NewOpsMsgContextInfo())
	}

	return newHistory
}

// Empty - Resets the current OpsMsgDto object to
// an uninitialized or 'empty' state.
func (opsMsg *OpsMsgDto) Empty() {
	opsMsg.EmptyParentHistory()
	opsMsg.EmptyMessageContext()
	opsMsg.EmptyMsgData()
}


// EmptyParentHistory - Deletes the current ParentHistory ([] OpsMsgContextInfo)
// and resets it to an 'empty' or uninitialized state (zero length array).
func (opsMsg *OpsMsgDto) EmptyParentHistory() {
	opsMsg.ParentContextHistory = make([] OpsMsgContextInfo, 0, 30)
}

// EmptyMessageContext - Deletes the current message context
// (OpsMsgDto.MsgContext) and resets it to an uninitialized state.
func (opsMsg *OpsMsgDto) EmptyMessageContext() {
	opsMsg.MsgContext = OpsMsgContextInfo{}
}

// EmptyMsgData - Resets all OpsMsgDto fields, with
// the exception of ParentContextHistory and MsgContext,
// to an uninitialized or 'empty' state.
func (opsMsg *OpsMsgDto) EmptyMsgData() {
	opsMsg.Message 					= ""
	opsMsg.fmtMessage 			= ""
	opsMsg.abbrvMessage			= ""
	opsMsg.msgId          	= int64(0) // The identifying number for this message
	opsMsg.msgNumber      	= int64(0) //  Message Number = msgId + MsgContext.BaseMessageId. This is the number displayed in the message
	opsMsg.MsgType        	= OpsMsgTypeNOERRORNOMSGS
	opsMsg.MsgTimeUTC     	= time.Time{}
	opsMsg.MsgTimeLocal   	= time.Time{}
	opsMsg.MsgLocalTimeZone	= ""
}

// Equal - Compares an incoming OpsMsgDto object
// with the current OpsMsgDto object to determine
// if they are equivalent.
func (opsMsg *OpsMsgDto) Equal(opsMsg2 *OpsMsgDto) bool {

	l1 := len(opsMsg.ParentContextHistory)
	l2 := len(opsMsg2.ParentContextHistory)

	if l1 != l2 {
		return false
	}

	for i:= 0; i < l1; i++ {
		if !opsMsg.ParentContextHistory[i].Equal(&opsMsg2.ParentContextHistory[i]) {
			return false
		}
	}

	if !opsMsg.MsgContext.Equal(&opsMsg2.MsgContext) {
		return false
	}

	if  opsMsg.Message     			!= opsMsg2.Message 						||
			opsMsg.MsgType          != opsMsg2.MsgType						||
			opsMsg.MsgTimeUTC       != opsMsg2.MsgTimeUTC					||
			opsMsg.MsgTimeLocal     != opsMsg2.MsgTimeLocal				||
			opsMsg.MsgLocalTimeZone != opsMsg2.MsgLocalTimeZone   ||
			opsMsg.fmtMessage 			!= opsMsg2.GetFmtMessage()		||
			opsMsg.abbrvMessage			!= opsMsg2.GetAbbrvMessage()	||
			opsMsg.msgId            != opsMsg2.GetMessageId()			||
			opsMsg.msgNumber        != opsMsg2.GetMessageNumber()	{


		return false

	}

	return true

}

// Error - implements the 'Error' interface
// for golang errors. You can therefore pass
// the OpsMsgDto structure to any golang
// method that supports the 'Error' interface.
//
// Notice that the error object is created
// with one of two string types. If the
// OpsMsgDto field 'UseFormattedMsg' is set
// to true, the fully formatted message
// string is used. Otherwise, a shorter
// or abbreviated version of the the
// message string is used.
//
// This message does not filter by message
// type. The Error() method will create
// and return an error object for any type
// of message object, including Information,
// Warning, NoErrorsNoMessages and Successful
// Completion Messages.
func (opsMsg OpsMsgDto) Error() string {

	if opsMsg.UseFormattedMsg {
		return opsMsg.fmtMessage
	}

	return opsMsg.abbrvMessage

}

// GetError - If the current OpsMsgDto is
// configured as either a Standard Error or
// Fatal Error, this method will return
// an 'error' type containing the error
// message. If OpsMsgDto is configured as
// a non-error type message, this method
// will return 'nil'.
//
// The error string returned by this method
// is determined by the boolean OpsMsgDto
// field, opsMsg.UseFormattedMsg. If true,
// the fully formatted message string will
// be configured in the returned error type.
// If false, the abbreviated or short version
// of the message string will be configured
// in the error.
func (opsMsg *OpsMsgDto) GetError() error {

	if opsMsg.IsError() {

		if opsMsg.UseFormattedMsg {
			return errors.New(opsMsg.GetFmtMessage())
		}

		return errors.New(opsMsg.GetAbbrvMessage())
	}

	return nil

}

// GetFmtMessage - Returns the formatted Operations
// Message string for this OpsMsgDto object.
//
// For an abbreviated form of this message, see
// GetAbbrvMessage()
func (opsMsg *OpsMsgDto) GetFmtMessage() string {

	return opsMsg.fmtMessage
}

// GetAbbrvMessage - returns a shorter form of the
// message associated with this OpsMsgDto object.
// For the fully formatted message, reference method
// GetFmtMessage()
func (opsMsg *OpsMsgDto) GetAbbrvMessage() string {
	return opsMsg.abbrvMessage
}


// GetMessageId - returns data field 'msgId' for
// the current OpsMsgDto object.
func (opsMsg *OpsMsgDto) GetMessageId() int64 {
	return opsMsg.msgId
}

// GetMessageNumber - returns the data field 'msgNumber'
// for the current OpsMsgDto object. The 'msgNumber' is
// equal to 'msgId' plus 'opsMsg.MsgContext.BaseMessageId'
func (opsMsg *OpsMsgDto) GetMessageNumber() int64 {

	return opsMsg.msgNumber
}

// GetNewParentHistory - Returns a new Parent History Array consisting
// of the original Parent History plus the current Message Context
func (opsMsg *OpsMsgDto) GetNewParentHistory() [] OpsMsgContextInfo {

	var newParentHistory [] OpsMsgContextInfo

	if len(opsMsg.ParentContextHistory) == 0 {
		newParentHistory = make([] OpsMsgContextInfo, 0, 5)
	} else {
		newParentHistory = opsMsg.DeepCopyParentContextHistory(opsMsg.ParentContextHistory)
	}

	if opsMsg.MsgContext.SourceFileName != "" 		||
			opsMsg.MsgContext.ParentObjectName != "" 	||
				opsMsg.MsgContext.FuncName != ""				||
					opsMsg.MsgContext.BaseMessageId != 0	{

		newParentHistory = append(newParentHistory, opsMsg.MsgContext.DeepCopyOpsMsgContextInfo())
	}



	return newParentHistory
}

// InitializeAllContextInfo - Initializes Parent Context History and Message Context Info for a new
// OpsMsgDto object.
//
// Input Parameters:
// =================
//
// parentHistory []OpsMsgContextInfo - An array of OpsMsgContextInfo objects
// 											documenting execution path that led to the execution
//											of this method.
//
// msgContext OpsMsgContextInfo - This object records the current context in
//											which the new OpsMsgDto returned by this method will
//											will be operating.
//
// 											It allows the newly created OpsMsgDto to return data
// 											on the execution path which	led to the generation of
// 											the Operations Message.
//
// Example Usage:
// ==============
//
// oMsg := OpsMsgDto{}.InitializeAllContextInfo(parentHistory, msgContext)
//
// Parent Context History and current Message Context serve as an important
// purpose. It allows one to maintain a record of the function execution tree
// that led to the generation of this message.
//
func(opsMsg OpsMsgDto) InitializeAllContextInfo(parentHistory []OpsMsgContextInfo, msgContext OpsMsgContextInfo) OpsMsgDto {
	om := OpsMsgDto{}
	om.ParentContextHistory = om.DeepCopyParentContextHistory(parentHistory)
	om.MsgContext = msgContext.DeepCopyOpsMsgContextInfo()

	return om
}


// InitializeWithMessageContext - Initialize a new OpsMsgDto object
// with only a Message Context - No ParentHistory.
//
// Input Parameters:
// =================
//
// msgContext OpsMsgContextInfo - This object records the current context in
//											which the new OpsMsgDto returned by this method will
//											will be operating.
//
// 											It allows the newly created OpsMsgDto to return data
// 											on the execution path which	led to the generation of
// 											the Operations Message.
//
//
// Example Usage:
// ==============
//
// oMsg := OpsMsgDto{}.InitializeAllContextInfo(msgContext)
//
func(opsMsg OpsMsgDto) InitializeWithMessageContext(msgContext OpsMsgContextInfo) OpsMsgDto {
	om := OpsMsgDto{}
	om.MsgContext = msgContext.DeepCopyOpsMsgContextInfo()
	return om
}

// InitializeWithParentHistory - Initialize a new OpsMsgDto object
// with Parent History Only - No MessageContext
//
// Input Parameters:
// =================
//
// parentHistory []OpsMsgContextInfo - An array of OpsMsgContextInfo objects
// 											documenting the execution path that led to the execution
//											of this method.
//
//
// Example Usage:
// ==============
//
// oMsg := OpsMsgDto{}.InitializeAllContextInfo(parentHistory)
//
func(opsMsg OpsMsgDto) InitializeWithParentHistory(parentHistory []OpsMsgContextInfo) OpsMsgDto {

	om := OpsMsgDto{}

	om.ParentContextHistory = om.DeepCopyParentContextHistory(parentHistory)

	return  om
}

// InitializeContextWithOpsMsgDto - Initialize a new OpsMsgDto object with a parent
// OpsMsgDto object plus the OpsMsgContextInfo object passed as an input
// parameter, 'newMsgContext'.
//
// Input Parameters:
// =================
//
//	parentOpsMsg OpsMsgDto 	- The Parent History context from the incoming
//														OpsMsgDto will be added to the new OpsMsgDto
//														object being created by this method. In addition,
//														the Parent OpsMsgDto MsgContext will be added to
//														current OpsMsgDto ParentContextHistory.
//
//	newMsgContext OpsMsgContextInfo - This new OpsMsgContextInfo object will
//										be configured as the 'MsgContext' field in the new
// 										OpsMsgDto object created by this method.
//
//										It allows the newly created OpsMsgDto to return data
// 										on the execution path which	led to the generation of
// 										the Operations Message.
//
//	Example Usage:
//  ==============
//
//	parentOpsMsgDto   = OpsMsgDto object created in the parent function
//	currentMsgContext = OpsMsgContextInfo object describing the current context
//											(i.e. Function Name)
//
// omParent := OpsMsgDto{}.InitializeAllContextInfo(parentHistory, parentMsgContextInfo)
// currentContext := OpsMsgContextInfo{SourceFileName:"xray.go",ParentObjectName: "stringutil",
// 												FuncName:"DoSomeWork",BaseMessgeId:int64(8000)}
//
// om := OpsMsgDto{}.InitializeContextWithOpsMsgDto(omParent, currentContext)
//
// Background:
// ===========
// Parent Context History and current Message Context serve as an important
// purpose. It allows one to maintain a record of the function execution tree
// that led to the generation of this message.
//
func(opsMsg OpsMsgDto) InitializeContextWithOpsMsgDto(parentOpsMsg OpsMsgDto, newMsgContext OpsMsgContextInfo) OpsMsgDto {

	om := OpsMsgDto{}

	om.ParentContextHistory = om.DeepCopyParentContextHistory(parentOpsMsg.ParentContextHistory)
	om.AddOpsMsgContextInfoToParentHistory(parentOpsMsg.MsgContext)
	om.MsgContext = newMsgContext.DeepCopyOpsMsgContextInfo()

	return om
}

// Initialize - Create a new OpsMsgDto from a parent OpsMsgDto. This
// method will add the parent Message Context to ParentContext History
// array and return it in a new OpsMsgDto object.
//
// Note: The returned OpsMsgDto object will have an 'empty' Message
// Context object.
func (opsMsg OpsMsgDto) Initialize(parentOpsMsg OpsMsgDto) OpsMsgDto {
	om := OpsMsgDto{}

	om.ParentContextHistory = om.DeepCopyParentContextHistory(parentOpsMsg.ParentContextHistory)
	om.AddOpsMsgContextInfoToParentHistory(parentOpsMsg.MsgContext)

	return om
}


// IsDebugMsg - Returns a boolean value indicating
// whether or not this message is a 'Debug'
// type message.
func(opsMsg OpsMsgDto) IsDebugMsg() bool {

	if opsMsg.MsgType == OpsMsgTypeDEBUGMSG {
		return true
	}

	return false

}

// IsError - Returns a boolean value signaling
// whether the current OpsMsgDto is an 'error'
// message.
//
// If the return value is true, the OpsMsgDto
// will be configured as either a Standard Error
// or a Fatal Error. (See Method IsFatalError())
func (opsMsg *OpsMsgDto) IsError() bool {

	if opsMsg.MsgType == OpsMsgTypeERRORMSG ||
			opsMsg.MsgType == OpsMsgTypeFATALERRORMSG {
		return true
	}

	return false
}

// IsFatalError - If the current OpsMsgDto object is configured
// as a fatal error, this method will return true. A fatal error
// is the equivalent of a 'panic' error which immediately terminates
// program execution.
func (opsMsg *OpsMsgDto) IsFatalError() bool {

	if opsMsg.MsgType == OpsMsgTypeFATALERRORMSG {
		return true
	}

	return false

}

// IsInfoMsg - Returns a boolean value indicating
// whether or not this message is an 'Information'
// type message.
func (opsMsg *OpsMsgDto) IsInfoMsg() bool {

	if opsMsg.MsgType == OpsMsgTypeINFOMSG {
		return true
	}

	return false


}

// IsNoErrorsNoMessages  - Returns a boolean value indicating
// whether or not this message is a 'No Errors No Messages'
// type message.
//
// 'No Errors No Messages' is the type of message assigned to
// uninitialized OpsMsgDto objects.
func (opsMsg *OpsMsgDto) IsNoErrorsNoMessages() bool {

	if opsMsg.MsgType == OpsMsgTypeNOERRORNOMSGS {
		return true
	}

	return false

}

// IsSuccessfulCompletion - Returns a boolean value indicating
// whether or not this message is a 'Successful Completion' type
// message.
func (opsMsg *OpsMsgDto) IsSuccessfulCompletionMsg() bool {

	if opsMsg.MsgType == OpsMsgTypeSUCCESSFULCOMPLETION {
		return true
	}

	return false

}

// IsWarningMsg - Returns a boolean value indicating
// whether or not this message is a 'Warning' type
// message.
func (opsMsg *OpsMsgDto) IsWarningMsg() bool {

	if opsMsg.MsgType == OpsMsgTypeWARNINGMSG {
		return true
	}

	return false

}

// LogMsgToFile - Logs the OpsMsgDto message text to a file.
func (opsMsg *OpsMsgDto) LogMsgToFile(fPtr *os.File) (int, error) {
	return fPtr.WriteString(opsMsg.String())
}

// LogMsgAndPanic - This method will write the operations message text
// to disk file. Then, if the OpsMsgDto object is configured as a 'FATAL'
// error, it will call the panic function with the operations message text
// and terminate program execution.
//
// If the OpsMsgDto object is NOT configured as a 'FATAL' Error Message,
// the program will continue execution normally after writing the operations
// message to the disk file.
//
// ************************************************************************
// BE CAREFUL! - Understand the implications before calling this method.
// ************************************************************************
//
func (opsMsg *OpsMsgDto) LogMsgAndPanic(fPtr *os.File) {

	opsMsg.LogMsgToFile(fPtr)

	if opsMsg.MsgType == OpsMsgTypeFATALERRORMSG {

		defer fPtr.Close()

		panic(opsMsg.String())

	}

}

// NewDebugMsg - Create a new Debug Message
//
// Input Parameters
// ****************
//
// msg string			- This is the text message that will be formatted and displayed
//									for this OpsMsgDto object. This is the Debug Message.
//
// msgId int64		- The message Id number to be associated with this message. If this
//									Id number is set to zero ('0'), no message number will be displayed
//									in the final message text.
//
//									Message Id differs from Message Number. The final Message Number is
//									calculated by adding Message Id to the opsMsg.MsgContext.BaseMessageId.
//									Again, if Message Id is zero ('0'), no message number is displayed in
//									the final message text.
//
func(opsMsg OpsMsgDto) NewDebugMsg(msg string, msgId int64) OpsMsgDto {

	om := OpsMsgDto{}
	om.SetParentMessageContextHistory(opsMsg.DeepCopyParentContextHistory(opsMsg.ParentContextHistory))
	om.SetMessageContext(opsMsg.MsgContext)
	om.SetDebugMessage(msg, msgId)

	return om
}

// NewError - Create a new OpsMsgDto by inputting an error type and a message type.
//
// Input Parameters
// ****************
//
//  prefixMsg string	- The message text will be prefixed to the error text to create the final
//											message text.
//
//	err error					- An error type containing error text will be added to the end of the prefixMsg
//											to create the final message text.
//
//
//	msgType OpsMsgType- This Type code specifies the type of message to be created. Available types
//											are:
// 												1. No Errors - No Messages 		= OpsMsgTypeNOERRORNOMSGS
//												2. Standard Errors 						= OpsMsgTypeERRORMSG
//												3. Fatal Errors (Panic Errors)= OpsMsgTypeFATALERRORMSG
//												4. Information Messages				= OpsMsgTypeINFOMSG
//												5. Warning Messages						= OpsMsgTypeWARNINGMSG
//												6. Debug Messages							= OpsMsgTypeDEBUGMSG
//												7. Successful Completion Msg  = OpsMsgTypeSUCCESSFULCOMPLETION
//
//
//	errId	int64		- The error Id number to be associated with
//									this message. If 'errId' is equal to zero,
//									no error number will be displayed in the
//									final error message
//
//									Error Id differs from Error Number. The final Error Number is
//									calculated by adding Error Id to the opsMsg.MsgContext.BaseMessageId.
//									Again, if Error Id is zero ('0'), no Error Number is displayed in
//									the final message text.
//
func (opsMsg OpsMsgDto) NewError(prefixMsg string, err error, msgType OpsMsgType, errId int64) OpsMsgDto {

	om := OpsMsgDto{}
	om.SetParentMessageContextHistory(opsMsg.DeepCopyParentContextHistory(opsMsg.ParentContextHistory))
	om.SetMessageContext(opsMsg.MsgContext)
	om.SetFromError(prefixMsg, err, msgType, errId)

	return om
}

// NewFatalErrorMsg - Creates a New FATAL Error Message
//
// Input Parameters
// ****************
//
//	errMsg string	- The text of the Error Message
//
//	errId	int64		- The error Id number to be associated with
//									this message. If 'errId' is equal to zero,
//									no error number will be displayed in the
//									final error message
//
//									Error Id differs from Error Number. The final Error Number is
//									calculated by adding Error Id to the opsMsg.MsgContext.BaseMessageId.
//									Again, if Error Id is zero ('0'), no Error Number is displayed in
//									the final message text.
//
func (opsMsg OpsMsgDto) NewFatalErrorMsg(errMsg string, errId int64) OpsMsgDto {

	om := OpsMsgDto{}
	om.SetParentMessageContextHistory(opsMsg.DeepCopyParentContextHistory(opsMsg.ParentContextHistory))
	om.SetMessageContext(opsMsg.MsgContext)
	om.SetFatalErrorMessage(errMsg, errId)
	return om

}

// NewInfoMsg - Create a new Operations Message which is
// an Informational Message.
//
// Input Parameters
// ****************
//
//	msg string 		- The text of the Information Message
//
// msgId int64		- The message Id number to be associated with this message. If this
//									Id number is set to zero ('0'), no message number will be displayed
//									in the final message text.
//
//									Message Id differs from Message Number. The final Message Number is
//									calculated by adding Message Id to the opsMsg.MsgContext.BaseMessageId.
//									Again, if Message Id is zero ('0'), no message number is displayed in
//									the final message text.
//
func(opsMsg OpsMsgDto) NewInfoMsg(msg string, msgId int64) OpsMsgDto {

	om := OpsMsgDto{}
	om.SetParentMessageContextHistory(opsMsg.DeepCopyParentContextHistory(opsMsg.ParentContextHistory))
	om.SetMessageContext(opsMsg.MsgContext)
	om.SetInfoMessage(msg, msgId)

	return om
}

// NewMsgFromSpecErrMsg - Create a new Operations Message based on
// the error information contained in a Type SpecErr passed
// into the method. The new message will be designated as
// an error message.
//
// Input Parameters:
// se SpecErr 		- A SpecErr object. See the SpecErr type in source
//									file errutility.go in this repository.
func (opsMsg *OpsMsgDto) NewMsgFromSpecErrMsg(se SpecErr) OpsMsgDto {

	om := OpsMsgDto{}
	om.SetFromSpecErrMessage(se)

	return om
}

// NewStdErrorMsg - Creates a new non-fatal error message
//
// Input Parameters
// ****************
//
//	errMsg string	- The text of the Error Message
//
//	errId	int64		- The error Id number to be associated with
//									this message. If 'errId' is equal to zero,
//									no error number will be displayed in the
//									final error message.
//
//									Error Id differs from Error Number. The final Error Number is
//									calculated by adding Error Id to the opsMsg.MsgContext.BaseMessageId.
//									Again, if Error Id is zero ('0'), no Error Number is displayed in
//									the final message text.
func (opsMsg OpsMsgDto) NewStdErrorMsg(errMsg string, errId int64) OpsMsgDto {

	om := OpsMsgDto{}
	om.SetParentMessageContextHistory(opsMsg.ParentContextHistory)
	om.SetMessageContext(opsMsg.MsgContext)
	om.SetStdErrorMessage(errMsg, errId)
	return om
}

// NewSuccessfulCompletionMsg - Creates a new Successful Completion
// Message and returns it as a new OpsMsgDto object.
//
// Input Parameters:
// =================
//
// msg string			- This is the text message that will be formatted and displayed
//									for this OpsMsgDto object.
//
// msgId int64		- The message Id number to be associated with this message. If this
//									Id number is set to zero ('0'), no message number will be displayed
//									in the final message text.
//
//									Message Id differs from Message Number. The final Message Number is
//									calculated by adding Message Id to the opsMsg.MsgContext.BaseMessageId.
//									Again, if Message Id is zero ('0'), no message number is displayed in
//									the final message text.
func (opsMsg OpsMsgDto) NewSuccessfulCompletionMsg(msg string, msgId int64) OpsMsgDto {

	om := OpsMsgDto{}
	om.SetParentMessageContextHistory(opsMsg.ParentContextHistory)
	om.SetMessageContext(opsMsg.MsgContext)
	om.SetSuccessfulCompletionMessage(msg, msgId)
	return om
}

// NewWarningMsg - Creates a new Warning Message
// and returns it as a new OpsMsgDto object.
func (opsMsg OpsMsgDto) NewWarningMsg(msg string, msgNo int64) OpsMsgDto {

	om := OpsMsgDto{}
	om.SetParentMessageContextHistory(opsMsg.ParentContextHistory)
	om.SetMessageContext(opsMsg.MsgContext)

	om.SetWarningMessage(msg, msgNo)

	return om

}

// NewNoErrorsNoMessagesMsg - Creates a new No Errors and No
// Messages Message and returns it as a new OpsMsgDto object.
//
// Input Parameters:
// =================
//
// msg string			- This is the text message that will be formatted and displayed
//									for this OpsMsgDto object.
//
// msgId int64		- The message Id number to be associated with this message. If this
//									Id number is set to zero ('0'), no message number will be displayed
//									in the final message text.
//
//									Message Id differs from Message Number. The final Message Number is
//									calculated by adding Message Id to the opsMsg.MsgContext.BaseMessageId.
//									Again, if Message Id is zero ('0'), no message number is displayed in
//									the final message text.
//
func (opsMsg OpsMsgDto) NewNoErrorsNoMessagesMsg(msg string, msgId int64) OpsMsgDto {

	om := OpsMsgDto{}
	om.SetParentMessageContextHistory(opsMsg.ParentContextHistory)
	om.SetMessageContext(opsMsg.MsgContext)

	om.SetNoErrorsNoMessages(msg, msgId)

	return om

}

// PanicOnFatalError - If the current OpsMsgDto is configured
// as a FATAL Error - this method will issue a call to
// 'panic' thereby halting program execution.
func (opsMsg *OpsMsgDto) PanicOnFatalError() {

	if opsMsg.MsgType == OpsMsgTypeFATALERRORMSG {
		panic(opsMsg.String())
	}

}

// SetDebugMessage - Configures the current or host
// OpsMsgDto object as a 'DEBUG' message.
//
// Input Parameters:
// =================
//
// msg string			- This is the text message that will be formatted and displayed
//									for this OpsMsgDto object.
//
// msgId int64		- The message Id number to be associated with this message. If this
//									Id number is set to zero ('0'), no message number will be displayed
//									in the final message text.
//
//									Message Id differs from Message Number. The final Message Number is
//									calculated by adding Message Id to the opsMsg.MsgContext.BaseMessageId.
//									Again, if Message Id is zero ('0'), no message number is displayed in
//									the final message text.
//
func (opsMsg *OpsMsgDto) SetDebugMessage(msg string, msgId int64){
	opsMsg.EmptyMsgData()
	opsMsg.MsgType = OpsMsgTypeDEBUGMSG

	opsMsg.setMessageText(msg, msgId)

}

// PrintToConsole - Prints Operations Message to the
// Console
func (opsMsg *OpsMsgDto) PrintToConsole() {

	fmt.Printf(opsMsg.String())

}


// SetFatalError - Creates a Fatal Error message from an error type.
//
// Input Parameters
// ****************
//
//  prefixMsg string	- The message text will be prefixed to the error text to create the final
//											message text.
//
//	err error					- An error type containing error text which will be added to the end of the
//											'prefixMsg' to create the final message text.
//
//	errMsg string			- The text of the Error Message
//
//	errId	int64				- The number to be associated with this message. If 'errId' is equal to zero,
// 											no error number will be displayed in the final error message text.
//
//											Error Id differs from Error Number. The final Error Number is
//											calculated by adding Error Id to the opsMsg.MsgContext.BaseMessageId.
//											Again, if Error Id is zero ('0'), no error number is displayed in
//											the final message text.
//
//
func(opsMsg *OpsMsgDto) SetFatalError(prefixMsg string, err error, errId int64) {

	var msg string
	var errMsg string

	if err == nil {
		errMsg = ""
	} else {
		errMsg = err.Error()
	}

	if prefixMsg != "" {
		msg = prefixMsg + errMsg
	} else {
		msg = errMsg
	}

	opsMsg.EmptyMsgData()
	opsMsg.MsgType = OpsMsgTypeFATALERRORMSG

	opsMsg.setMessageText(msg, errId)

}

// SetFatalErrorMessage - Configures the current or host
// OpsMsgDto object as an information message.
//
// Input Parameters:
// =================
//
// errMsg string	- This is the text message that will be formatted and displayed
//									for this OpsMsgDto object.
//
// errId int64		- The error Id number to be associated with this message. If this
//									Id number is set to zero ('0'), no message number will be displayed
//									in the final message text.
//
//									Error Id differs from Error Number. The final Error Number is
//									calculated by adding Error Id to the opsMsg.MsgContext.BaseMessageId.
//									Again, if Error Id is zero ('0'), no error number is displayed in
//									the final message text.
//
func (opsMsg *OpsMsgDto) SetFatalErrorMessage(errMsg string, errId int64) {

	opsMsg.EmptyMsgData()
	opsMsg.MsgType = OpsMsgTypeFATALERRORMSG

	opsMsg.setMessageText(errMsg, errId)

}

// SetFromError - Configure this OpsMsgDto object by
// inputting an error type.
//
// Input Parameters
// ****************
//
//  prefixMsg string	- The message text will be prefixed to the error text to create the final
//											message text.
//
//	err error					- An error type containing error text which will be added to the end of the
//											'prefixMsg' to create the final message text.
//
//	msgType OpsMsgType- This Type code specifies the type of message to be created. Available types
//											are:
// 												1. No Errors - No Messages 		= OpsMsgTypeNOERRORNOMSGS
//												2. Standard Errors 						= OpsMsgTypeERRORMSG
//												3. Fatal Errors (Panic Errors)= OpsMsgTypeFATALERRORMSG
//												4. Information Messages				= OpsMsgTypeINFOMSG
//												5. Warning Messages						= OpsMsgTypeWARNINGMSG
//												6. Debug Messages							= OpsMsgTypeDEBUGMSG
//												7. Successful Completion Msg  = OpsMsgTypeSUCCESSFULCOMPLETION
//
//	errId	int64				- The error id number to be associated with this message. If 'errNo' is equal
// 											to zero, no error number will be displayed in the final error message.
//
//											Error Id differs from Error Number. The final Error Number is
//											calculated by adding Error Id to the opsMsg.MsgContext.BaseMessageId.
//											Again, if Message Id is zero ('0'), no message number is displayed in
//											the final message text.
//
func (opsMsg *OpsMsgDto) SetFromError(prefixMsg string, err error, msgType OpsMsgType, errId int64) {

	var msg, errMsg string

	if err == nil {
		errMsg = ""
	} else {
		errMsg = err.Error()
	}


	if prefixMsg != "" {
		msg = prefixMsg + errMsg
	} else {
		msg = errMsg
	}

	switch msgType {

	case OpsMsgTypeNOERRORNOMSGS:
		opsMsg.SetNoErrorsNoMessages(msg, errId)

	case OpsMsgTypeERRORMSG:
		opsMsg.SetStdErrorMessage(msg, errId)

	case OpsMsgTypeFATALERRORMSG:
		opsMsg.SetFatalErrorMessage(msg, errId)

	case OpsMsgTypeINFOMSG:
		opsMsg.SetInfoMessage(msg, errId)

	case OpsMsgTypeWARNINGMSG:
		opsMsg.SetWarningMessage(msg, errId)

	case OpsMsgTypeDEBUGMSG:
		opsMsg.SetWarningMessage(msg, errId)

	case OpsMsgTypeSUCCESSFULCOMPLETION:
		opsMsg.SetSuccessfulCompletionMessage(msg, errId)

	default:
		panic("OpsMsgDto.SetFromError() Error - INVALID OpsMsgType!")
	}

}

// SetFromSpecErrMessage - Sets an instance of OpsMsgDto type based
// on a SpecErr object passed as an input parameter.
//
// Input Parameters:
// se SpecErr 		- A SpecErr object. See the SpecErr type in source
//									file errutility.go in this repository.
//
func (opsMsg *OpsMsgDto) SetFromSpecErrMessage(se SpecErr) {

	opsMsg.Empty()
	
	x := se.DeepCopyParentInfo(se.ParentInfo)

	for _, bi := range x {
		ci := OpsMsgContextInfo{SourceFileName:bi.SourceFileName, ParentObjectName: bi.ParentObjectName, FuncName: bi.FuncName, BaseMessageId: bi.BaseErrorId}
		opsMsg.ParentContextHistory = append(opsMsg.ParentContextHistory, ci)
	}

	y := se.DeepCopyBaseInfo()

	opsMsg.MsgContext = OpsMsgContextInfo{SourceFileName:y.SourceFileName, ParentObjectName: y.ParentObjectName, FuncName: y.FuncName, BaseMessageId: y.BaseErrorId}

	errId := se.GetErrorId()

	switch se.ErrorMsgType {

	case SpecErrTypeNOERRORSALLCLEAR:
		opsMsg.SetNoErrorsNoMessages(se.ErrMsg, errId)

	case SpecErrTypeERROR:
		opsMsg.SetStdErrorMessage(se.ErrMsg, errId)

	case SpecErrTypeFATAL:
		opsMsg.SetFatalErrorMessage(se.ErrMsg, errId)

	case SpecErrTypeINFO:
		opsMsg.SetInfoMessage(se.ErrMsg, errId)

	case SpecErrTypeWARNING:
		opsMsg.SetWarningMessage(se.ErrMsg, errId)

	case SpecErrTypeDEBUG:
		opsMsg.SetDebugMessage(se.ErrMsg, errId)

	case SpecErrTypeSUCCESSFULCOMPLETION:
		opsMsg.SetSuccessfulCompletionMessage(se.ErrMsg, errId)

	default:
		panic("OpsMsgDto.SetFromSpecErrMessage() - INVALID SpecErrType Code")
	}

}

// SetInfoMessage - Configures the current or host
// OpsMsgDto object as an information message.
//
// Input Parameters:
// =================
//
// msg string			- This is the text message that will be formatted and displayed
//									for this OpsMsgDto object.
//
// msgId int64		- The message Id number to be associated with this message. If this
//									Id number is set to zero ('0'), no message number will be displayed
//									in the final message text.
//
//									Message Id differs from Message Number. The final Message Number is
//									calculated by adding Message Id to the opsMsg.MsgContext.BaseMessageId.
//									Again, if Message Id is zero ('0'), no message number is displayed in
//									the final message text.
//
func (opsMsg *OpsMsgDto) SetInfoMessage(msg string, msgId int64) {
	opsMsg.EmptyMsgData()
	opsMsg.MsgType = OpsMsgTypeINFOMSG

	opsMsg.setMessageText(msg, msgId)
}


// SetMsgContext - Receives an OpsMsgContextInfo object as
// an input parameter and configures the current OpsMsgDto
// MessageContext.
func (opsMsg *OpsMsgDto) SetMessageContext(msgContext OpsMsgContextInfo) {
	opsMsg.MsgContext = msgContext.DeepCopyOpsMsgContextInfo()
}

// SetMessageOutputMode - This method is used to set the message output mode.
// If the input parameter, 'isFullyFormattedMsg' is set to 'true', the methods
// String() and Error() will return fully formatted messages like this:
//
// !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
// FATAL ERROR Message                             								 Error No: 6974
// !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
// Is Error: true       Is Panic\Fatal Error: true
// ------------------------------------------------------------------------------
// Message: This is FATAL Error message text.
// ------------------------------------------------------------------------------
// Parent Context History:
// SrcFile: TSource01 -ParentObj: PObj01 -FuncName: Func001 -BaseMsgId: 1000
// SrcFile: TSource02 -ParentObj: PObj02 -FuncName: Func002 -BaseMsgId: 2000
// SrcFile: TSource03 -ParentObj: PObj03 -FuncName: Func003 -BaseMsgId: 3000
// SrcFile: TSource04 -ParentObj: PObj04 -FuncName: Func004 -BaseMsgId: 4000
// SrcFile: TSource05 -ParentObj: PObj05 -FuncName: Func005 -BaseMsgId: 5000
// ------------------------------------------------------------------------------
// Current Message Context:
// SrcFile: TSource06 -ParentObj: PObj06 -FuncName: Func006 -BaseMsgId: 6000
// ------------------------------------------------------------------------------
//   Message Time UTC: 2017-12-16 Sat 22:12:28.458551100 +0000 UTC
// Message Time Local: 2017-12-16 Sat 16:12:28.458551100 -0600 CST
// !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
//
// If input parameter 'isFullyFormattedMsg' is set to 'false', an abbreviated
// message text will be returned by methods Error() and String(). Example:
//
// FATAL ERROR Msg No: 972 - 12/16/2017 16:18:24.904997000 -0600 CST - Test Serious Error.
//
// Use this method, 'SetMessageOutputMode', to determine the message format you wish
// to display.
//
func (opsMsg *OpsMsgDto) SetMessageOutputMode(isFullyFormattedMsg bool) {

	opsMsg.UseFormattedMsg = isFullyFormattedMsg
}


// SetParentMessageContextHistory - Deletes the current opsMsg.ParentContextHistory
// and replaeces it with the input parameter, 'parentHistory',
func (opsMsg *OpsMsgDto) SetParentMessageContextHistory( parentHistory []OpsMsgContextInfo) {
	opsMsg.ParentContextHistory = make([] OpsMsgContextInfo, 0, 30)
	l1 := len(parentHistory)

	for i:= 0; i < l1; i++ {
		opsMsg.ParentContextHistory = append(opsMsg.ParentContextHistory, parentHistory[i])
	}

}

// SetStdError - Configures a Standard Error message from an input error type.
//
// Input Parameters:
// =================
//
// prefixMsg string		- This string will be prefixed the error string to produce
//											the final message text.
//
// err error					- The text contained in this error type will be added to the
//											prefixMsg string to produce the final message text.
//
// errId int64				- This error Id number will be added to the opsMsg.MsgContext.BaseMessageId
//											to create the final Error or Message Number.
//
//											Error Id differs from Error Number. The final Error Number is
//											calculated by adding Error Id to the opsMsg.MsgContext.BaseMessageId.
//											Again, if Error Id is zero ('0'), no Error Number is displayed in
//											the final message text.
func (opsMsg *OpsMsgDto) SetStdError(prefixMsg string, err error, errId int64){

	var msg string
	var errMsg string

	if err == nil {
		errMsg = ""
	} else {
		errMsg = err.Error()
	}

	if prefixMsg != "" {
		msg = prefixMsg + errMsg
	} else {
		msg = errMsg
	}

	opsMsg.EmptyMsgData()
	opsMsg.MsgType = OpsMsgTypeERRORMSG
	opsMsg.setMessageText(msg, errId)


}

// SetStdErrorMessage - Configures the current or host
// OpsMsgDto object as a standard error message.
//
// Input Parameters
// ****************
//
//	errMsg string	- The text of the Error Message
//
//	errId	int64		- The error Id number to be associated with
//									this message. If 'errId' is equal to zero,
//									no error number will be displayed in the
//									final error message.
//
//									Error Id differs from Error Number. The final Error Number is
//									calculated by adding Error Id to the opsMsg.MsgContext.BaseMessageId.
//									Again, if Error Id is zero ('0'), no Error Number is displayed in
//									the final message text.
func (opsMsg *OpsMsgDto) SetStdErrorMessage(errMsg string, errId int64){
	opsMsg.EmptyMsgData()
	opsMsg.MsgType = OpsMsgTypeERRORMSG

	opsMsg.setMessageText(errMsg, errId)

}

// SetNoErrorsNoMessages - Configures the current or host
// OpsMsgDto object for the default message type,
// 'No Errors and No Messages'.
//
// Input Parameters:
// =================
//
// msg string			- This is the text message that will be formatted and displayed
//									for this OpsMsgDto object.
//
// msgId int64		- The message Id number to be associated with this message. If this
//									Id number is set to zero ('0'), no message number will be displayed
//									in the final message text.
//
//									Message Id differs from Message Number. The final Message Number is
//									calculated by adding Message Id to the opsMsg.MsgContext.BaseMessageId.
//									Again, if Message Id is zero ('0'), no message number is displayed in
//									the final message text.
//
func (opsMsg *OpsMsgDto) SetNoErrorsNoMessages(msg string, msgId int64) {

	opsMsg.EmptyMsgData()
	opsMsg.MsgType = OpsMsgTypeNOERRORNOMSGS

	opsMsg.setMessageText(msg, msgId)

}

// SetSuccessfulCompletionMessage - Configures the current or host
// OpsMsgDto object as a Successful Completion Message.
//
// Input Parameters:
// =================
//
// msg string			- This is the text message that will be formatted and displayed
//									for this OpsMsgDto object.
//
// msgId int64		- The message Id number to be associated with this message. If this
//									Id number is set to zero ('0'), no message number will be displayed
//									in the final message text.
//
//									Message Id differs from Message Number. The final Message Number is
//									calculated by adding Message Id to the opsMsg.MsgContext.BaseMessageId.
//									Again, if Message Id is zero ('0'), no message number is displayed in
//									the final message text.
//
func (opsMsg *OpsMsgDto) SetSuccessfulCompletionMessage(msg string, msgId int64){
	opsMsg.EmptyMsgData()
	opsMsg.MsgType = OpsMsgTypeSUCCESSFULCOMPLETION

	opsMsg.setMessageText( msg, msgId)

}

// SetWarningMessage - Configures the current or host
// OpsMsgDto object as a Warning Message.
//
// Input Parameters:
// =================
//
// msg string			- This is the text message that will be formatted and displayed
//									for this OpsMsgDto object.
//
// msgId int64		- The message Id number to be associated with this message. If this
//									Id number is set to zero ('0'), no message number will be displayed
//									in the final message text.
//
//									Message Id differs from Message Number. The final Message Number is
//									calculated by adding Message Id to the opsMsg.MsgContext.BaseMessageId.
//									Again, if Message Id is zero ('0'), no message number is displayed in
//									the final message text.
//
func (opsMsg *OpsMsgDto) SetWarningMessage(msg string, msgId int64) {
	opsMsg.EmptyMsgData()
	opsMsg.MsgType = OpsMsgTypeWARNINGMSG

	opsMsg.setMessageText(msg, msgId)

}

// SignalNoErrors - Creates a 'No Errors - No Messages' message and
// returns it as a new OpsMsgDto object.
//
// Can be used to quickly return a 'No Errors - No Messages' message
// to a calling function. A better approach might be to use the
// 'SignalSuccessfulCompletion() method shown below.
//
// Input Parameters:
// =================
//
// msgId int64		- The message Id number to be associated with this message. If this
//									Id number is set to zero ('0'), no message number will be displayed
//									in the final message text.
//
//									Message Id differs from Message Number. The final Message Number is
//									calculated by adding Message Id to the opsMsg.MsgContext.BaseMessageId.
//									Again, if Message Id is zero ('0'), no message number is displayed in
//									the final message text.
//
func (opsMsg OpsMsgDto) SignalNoErrors(msgId int64) OpsMsgDto {

	om := OpsMsgDto{}.InitializeAllContextInfo(opsMsg.DeepCopyParentContextHistory(opsMsg.ParentContextHistory), opsMsg.MsgContext.DeepCopyOpsMsgContextInfo())
	om.SetNoErrorsNoMessages("", msgId)

	return om
}

// SignalSuccessfulCompletion - Creates a Successful Completion message and
// returns it as a new OpsMsgDto object.
//
// Can be used to quickly return Successful Completion message to a calling function.
//
// Input Parameters:
// =================
//
// msgId int64		- The message Id number to be associated with this message. If this
//									Id number is set to zero ('0'), no message number will be displayed
//									in the final message text.
//
//									Message Id differs from Message Number. The final Message Number is
//									calculated by adding Message Id to the opsMsg.MsgContext.BaseMessageId.
//									Again, if Message Id is zero ('0'), no message number is displayed in
//									the final message text.
//
func (opsMsg OpsMsgDto) SignalSuccessfulCompletion(msgId int64) OpsMsgDto {

	om := OpsMsgDto{}.InitializeAllContextInfo(opsMsg.DeepCopyParentContextHistory(opsMsg.ParentContextHistory), opsMsg.MsgContext.DeepCopyOpsMsgContextInfo())
	om.SetSuccessfulCompletionMessage("", msgId)
	return om
}



// String - returns the operations message as a
// string.
func (opsMsg *OpsMsgDto) String() string {

	if opsMsg.UseFormattedMsg {
		return opsMsg.GetFmtMessage()
	}

	return opsMsg.GetAbbrvMessage()
}

// SetTimeZone - This method sets the local time zone for
// use in time stamping messages.
//
// Input Parameters:
// =================
//
// ianaTimeZone string - 	The IANA Time Zone as specified in the IANA Time Zone
// 												Data base. Reference https://www.iana.org/time-zones.
//												If this parameter is judged to be INVALID, the time
//												zone will default to 'Local' which is the local time
//												zone designated on the host computer.
//
func (opsMsg *OpsMsgDto) SetTimeZone(ianaTimeZone string) {

	if opsMsg.MsgLocalTimeZone == ianaTimeZone {
		return
	}

	opsMsg.MsgLocalTimeZone = ianaTimeZone

	opsMsg.setMessageText(opsMsg.Message, opsMsg.msgId)
}

// ***********************************************
// private methods
// ***********************************************

// setMsgParms - Returns the Message title, message number and the
// banner line separator based on value of OpsMsgDto.MsgClass
func (opsMsg *OpsMsgDto) setMsgParms() (banner1, banner2, title, numTitle, abbrvTitle string, ) {

	switch opsMsg.MsgType {

	case OpsMsgTypeNOERRORNOMSGS:
		// OpsMsgTypeNOERRORNOMSGS - 0 Signals uninitialized message
		// with no errors and no messages
		title = "No Errors and No Messages"
		abbrvTitle = "No Errors-No Messages"
		numTitle = "Msg Number"
		banner1 = strings.Repeat("&", 78)
		banner2 = strings.Repeat("-", 78)

	case OpsMsgTypeERRORMSG:
		// OpsMsgTypeERRORMSG - 1 Message is an Error Message
		title = "Standard ERROR Message"
		abbrvTitle = "Standard ERROR Msg"
		numTitle = "Error No"
		banner1 = strings.Repeat("#", 78)
		banner2 = strings.Repeat("-", 78)

	case OpsMsgTypeFATALERRORMSG:
		// OpsMsgTypeFATALERRORMSG - 2 Message is a Fatal Error Message
		title = "FATAL ERROR Message"
		abbrvTitle = "FATAL ERROR Msg"
		numTitle = "Error No"
		banner1 = strings.Repeat("!", 78)
		banner2 = strings.Repeat("-", 78)

	case OpsMsgTypeINFOMSG:
		// OpsMsgTypeINFOMSG - 3 Message is an Informational Message
		title = "Information Message"
		abbrvTitle = "Information Msg"
		numTitle = "Msg No"
		banner1 = strings.Repeat("*", 78)
		banner2 = strings.Repeat("-", 78)

	case OpsMsgTypeWARNINGMSG:
		// OpsMsgTypeWARNINGMSG - 4 Message is a warning Message
		title = "WARNING Message"
		abbrvTitle = "WARNING Msg"
		numTitle = "Msg No"
		banner1 = strings.Repeat("?", 78)
		banner2 = strings.Repeat("-", 78)

	case OpsMsgTypeDEBUGMSG:
		// OpsMsgTypeDEBUGMSG - 5 Message is a Debug Message
		title = "DEBUG Message"
		abbrvTitle = "DEBUG Msg"
		numTitle = " Number"
		banner1 = strings.Repeat("@", 78)
		banner2 = strings.Repeat("-", 78)

	case OpsMsgTypeSUCCESSFULCOMPLETION:
		// OpsMsgTypeSUCCESSFULCOMPLETION - 6 Message signalling successful
		// completion of the operation
		title = "Successful Completion"
		abbrvTitle = "Successful Completion Msg"
		numTitle = "Msg No"
		banner1 = strings.Repeat("$", 78)
		banner2 = strings.Repeat("-", 78)

	default:
		// This should never happen
		panic("OpsMsgDto.setMsgParms() - Invalid opsMsg.MsgClass")
	}

	return banner1, banner2, title, numTitle, abbrvTitle
}


func(opsMsg *OpsMsgDto) setDebugMsgText(banner1, banner2, title, numTitle string) {

	m := "\n\n"
	m += "\n" + banner1

	if opsMsg.msgNumber != 0 {
		m += fmt.Sprintf("\n %v %v: %v", title, numTitle, opsMsg.msgNumber)
	} else {
		m += fmt.Sprintf("\n %v -", title)
	}

	m += "\n  " + opsMsg.Message

	l1 := len(opsMsg.ParentContextHistory)
	if l1 > 0 {
		m += "\n" + banner2
		m += "\n Parent Context History:"
		for i:=0; i < l1; i++ {
			m+= "\n  Src File: " + opsMsg.ParentContextHistory[i].SourceFileName
			m+= "   Parent Obj: " + opsMsg.ParentContextHistory[i].ParentObjectName
			m+= "   Func Name: " + opsMsg.ParentContextHistory[i].FuncName
		}

	}

	if opsMsg.MsgContext.SourceFileName != "" ||
		opsMsg.MsgContext.ParentObjectName != "" ||
		opsMsg.MsgContext.FuncName != "" {
		m += "\n" + banner2
		m += "\n Current Message Context:"
		if opsMsg.MsgContext.SourceFileName != "" {
			m+= "\n  Src File: " + opsMsg.MsgContext.SourceFileName
		}

		if opsMsg.MsgContext.ParentObjectName != "" {
			m+= "   Parent Obj: " + opsMsg.MsgContext.ParentObjectName
		}

		if opsMsg.MsgContext.FuncName != "" {
			m+= "   Func Name: " + opsMsg.MsgContext.FuncName
		}
	}

	// FmtDateTimeTzNanoYMD

	localTz := opsMsg.MsgLocalTimeZone
	dtFmt := "01/02/2006 15:04:05.000000000 -0700 MST"
	if localTz == "Local" || localTz == "local" {
		localZone, _ := time.Now().Zone()
		localTz += " - " + localZone
	}
	m += "\n" + banner2
	m += "\n   UTC Time: " + opsMsg.MsgTimeUTC.Format(dtFmt)
	m += "\n Local Time: " + opsMsg.MsgTimeLocal.Format(dtFmt) + "   Time Zone: " + localTz

	m += "\n" + banner1


	opsMsg.fmtMessage =  m
}

// setMessageText - This method is called internally to set
// and format the text message for specific message types.
func(opsMsg *OpsMsgDto) setMessageText(msg string, msgId int64) {

	opsMsg.setMsgIdAndMsgNumber(msgId)

	opsMsg.setTime(opsMsg.MsgLocalTimeZone)

	opsMsg.UseFormattedMsg = true

	opsMsg.Message = msg

	banner1, banner2, title, numTitle, abbrvTitle := opsMsg.setMsgParms()

	if opsMsg.MsgType == OpsMsgTypeDEBUGMSG {
		opsMsg.setDebugMsgText(banner1, banner2, title, numTitle)
		opsMsg.setAbbreviatedMessageText(abbrvTitle)
		return
	}

	opsMsg.setFormatMessageText(banner1, banner2, title, numTitle)

	opsMsg.setAbbreviatedMessageText(abbrvTitle)
}


// - setAbbreviatedMessageText - Sets the abbreviated or
// or 'short form' text message for all messages.
func(opsMsg *OpsMsgDto) setAbbreviatedMessageText(abbrvTitle string) {

	var m string

	m = "\n\n"
	m += "\n" + abbrvTitle

	if opsMsg.msgNumber != 0 {
		m+= fmt.Sprintf(" No: %v - ", opsMsg.msgNumber)
	} else {
		m+= " - "
	}

	dtFmt := "01/02/2006 15:04:05.000000000 -0700 MST"
	m += opsMsg.MsgTimeLocal.Format(dtFmt)
	m += " - "
	m += opsMsg.Message

	opsMsg.abbrvMessage = m

}

func(opsMsg *OpsMsgDto) setFormatMessageText(banner1, banner2, title, numTitle string){

	var m string
	lineWidth := len(banner1)

	dtFmt := "2006-01-02 Mon 15:04:05.000000000 -0700 MST"

	m= "\n\n"
	m += "\n" + banner1
	nextBanner := banner1
	s1 := (lineWidth / 3) * 2
	s2 := lineWidth - s1

	if opsMsg.msgNumber != 0 {
		sNo:= fmt.Sprintf("%v: %v", numTitle, opsMsg.msgNumber)
		str1, _ := opsMsg.strCenterInStr(title, s1)
		str2, _ := opsMsg.strRightJustify(sNo, s2)
		m+= "\n" + str1 + str2
	} else {
		str1, _ := opsMsg.strCenterInStr(title, s1)
		m+= "\n" + str1
	}

	if opsMsg.MsgType == OpsMsgTypeERRORMSG ||
		opsMsg.MsgType == OpsMsgTypeFATALERRORMSG {

		m += "\n" + nextBanner
		nextBanner = banner2

		str1 := fmt.Sprintf(" Is Error: %v       Is Panic\\Fatal Error: %v", opsMsg.IsError(), opsMsg.IsFatalError())
		m += "\n" + str1

	}

	if opsMsg.Message != "" {
		m += "\n" + nextBanner
		nextBanner = banner2

		m += "\n Message: "

		if len(opsMsg.Message) > 67 {
			m += "\n  "
		}

		m += opsMsg.Message

	} else {
		m += "\n" + nextBanner
		nextBanner = banner2

		m += "\n Message: NO MESSAGE TEXT PROVIDED!!"

	}

	l1 := len(opsMsg.ParentContextHistory)
	if l1 > 0 {
		m += "\n" + nextBanner
		m += "\n Parent Context History:"
		for i:=0; i < l1; i++ {
			m+= "\n  SrcFile: " + opsMsg.ParentContextHistory[i].SourceFileName
			m+= " -ParentObj: " + opsMsg.ParentContextHistory[i].ParentObjectName
			m+= " -FuncName: " + opsMsg.ParentContextHistory[i].FuncName
			m+= " -BaseMsgId: " + fmt.Sprintf("%v",opsMsg.ParentContextHistory[i].BaseMessageId)
		}

		nextBanner = banner2
	}

	if opsMsg.MsgContext.SourceFileName != "" ||
		opsMsg.MsgContext.ParentObjectName != "" ||
		opsMsg.MsgContext.FuncName != "" {
		m += "\n" + nextBanner
		nextBanner = banner2
		m += "\n Current Message Context:"
		if opsMsg.MsgContext.SourceFileName != "" {
			m+= "\n  SrcFile: " + opsMsg.MsgContext.SourceFileName
		}

		if opsMsg.MsgContext.ParentObjectName != "" {
			m+= " -ParentObj: " + opsMsg.MsgContext.ParentObjectName
		}

		if opsMsg.MsgContext.FuncName != "" {
			m+= " -FuncName: " + opsMsg.MsgContext.FuncName
		}

		if opsMsg.MsgContext.BaseMessageId != 0 {
			m+= " -BaseMsgId: " + fmt.Sprintf("%v", opsMsg.MsgContext.BaseMessageId)
		}
	}

	m += "\n" + nextBanner
	m += fmt.Sprintf("\n   Message Time UTC: %v ", opsMsg.MsgTimeUTC.Format(dtFmt))
	m += fmt.Sprintf("\n Message Time Local: %v ", opsMsg.MsgTimeLocal.Format(dtFmt))
	m += "\n" + banner1

	opsMsg.fmtMessage =  m
}

// setMsgIdAndMsgNumber - This method is called internally
// to set the OpsMsgDto.msgId and OpsMsgDto.msgNumber fields.
func (opsMsg *OpsMsgDto) setMsgIdAndMsgNumber(msgId int64) {

	if msgId == 0 {
		opsMsg.msgId = 0
		opsMsg.msgNumber = 0
	} else {
		opsMsg.msgId = msgId
		opsMsg.msgNumber = msgId + opsMsg.MsgContext.BaseMessageId
	}


}


// setTime - Sets the time stamp for this Operations
// Message. Notice that the input parameter 'localTimeZone'
// is the Iana Time Zone for local time.
//
// Reference Iana Time Zones: https://www.iana.org/time-zones
//
// If the 'localTimeZone' parameter string is empty or an
// invalid time zone, local time zone will default to 'Local'.
// The 'Local' time zone is determined by the host computer.
func(opsMsg *OpsMsgDto) setTime(localTimeZone string){

	if localTimeZone == "" {
		localTimeZone = "Local"
	}

	tz := TimeZoneUtility{}

	isValid, _, _ := tz.IsValidTimeZone(localTimeZone)

	if !isValid {
		localTimeZone = "Local"
	}

	opsMsg.MsgTimeUTC = time.Now().UTC()
	opsMsg.MsgLocalTimeZone = localTimeZone

	tzLocal, _ := tz.ConvertTz(opsMsg.MsgTimeUTC, opsMsg.MsgLocalTimeZone)
	opsMsg.MsgTimeLocal = tzLocal.TimeOut

}


/*

Private String Management Methods

*/

// strCenterInStr - returns a string which includes
// a left pad blank string plus the original string.
// The complete string will effectively center the
// original string is a field of specified length.
func (opsMsg *OpsMsgDto) strCenterInStr(strToCenter string, fieldLen int) (string, error) {

	sLen := len(strToCenter)

	if sLen > fieldLen {
		return strToCenter,  fmt.Errorf("'fieldLen' = '%v' strToCenter Length= '%v'. 'fieldLen is shorter than strToCenter Length!", fieldLen, sLen)
	}

	if sLen == fieldLen {
		return strToCenter, nil
	}

	leftPadCnt := (fieldLen-sLen)/2

	leftPadStr := strings.Repeat(" ", leftPadCnt)

	rightPadCnt := fieldLen - sLen - leftPadCnt

	rightPadStr := ""

	if rightPadCnt > 0 {
		rightPadStr = strings.Repeat(" ", rightPadCnt)
	}


	return leftPadStr + strToCenter	+ rightPadStr, nil

}

// strRightJustify - Returns a string where input parameter
// 'strToJustify' is right justified. The length of the returned
// string is determined by input parameter 'fieldlen'.
func (opsMsg *OpsMsgDto) strRightJustify(strToJustify string, fieldLen int) (string, error) {

	strLen := len(strToJustify)

	if fieldLen == strLen {
		return strToJustify, nil
	}

	if fieldLen < strLen {
		return strToJustify, fmt.Errorf("Length of string to right justify is '%v'. 'fieldLen' is less. 'fieldLen'= '%v'", strLen, fieldLen)
	}

	// fieldLen must be greater than strLen
	lefPadCnt := fieldLen - strLen

	leftPadStr := strings.Repeat(" ", lefPadCnt)



	return leftPadStr + strToJustify, nil
}

// strPadLeftToCenter - Returns a blank string
// which allows centering of the target string
// in a fixed length field.
func (opsMsg *OpsMsgDto) strPadLeftToCenter(strToCenter string, fieldLen int) (string, error) {

	sLen := opsMsg.strGetRuneCnt(strToCenter)

	if sLen > fieldLen {
		return "", errors.New("StringUtility:StrPadLeftToCenter() - String To Center is longer than Field Length")
	}

	if sLen == fieldLen {
		return "", nil
	}

	margin := (fieldLen - sLen) / 2

	return strings.Repeat(" ", margin), nil
}

// strGetRuneCnt - Uses utf8 Rune Count
// function to return the number of characters
// in a string.
func (opsMsg *OpsMsgDto) strGetRuneCnt(targetStr string) int {
	return utf8.RuneCountInString(targetStr)
}

// strGetCharCnt - Uses the 'len' method to
// return the number of characters in a
// string.
func (opsMsg *OpsMsgDto) strGetCharCnt(targetStr string) int {
	return len([]rune(targetStr))
}


