package common

import (
	"testing"
	"strings"
	"errors"
)

func TestOpsMsgDto_NewSuccessfulCompletionMsg_01(t *testing.T) {

	testParentHistory := testOpsMsgDtoCreateParentHistory()
	testMsgContext := testOpsMsgDtoCreateContextInfoObj()

	xMsg := "Successful Completion"
	msgId := int64(64)
	msgNo := int64(6064)
	msgType := OpsMsgTypeSUCCESSFULCOMPLETION

	om := OpsMsgDto{}.InitializeAllContextInfo(testParentHistory, testMsgContext).NewSuccessfulCompletionMsg("",msgId)

	l1 := len(testParentHistory)

	l2 := len(om.ParentContextHistory)

	if l1 != l2 {
		t.Errorf("Expected om.ParentContextHistory to equal length= '%v'. It did NOT! actual length= '%v'",l1, l2)
	}

	for i:=0; i<l1; i++ {
		if !testParentHistory[i].Equal(&om.ParentContextHistory[i]) {
			t.Errorf("Expected om.ParentContextHistory to Equal testParentContext History. It did NOT!. i= '%v'",i)
		}
	}

	if !testMsgContext.Equal(&om.MsgContext) {
		t.Error("Expected testMsgContext to EQUAL om.MsgContext. It did NOT!")
	}

	if om.MsgType != msgType {
		t.Errorf("Expected Messgage Type == '%v'. Instead, Message Type == '%v'.", msgType, om.MsgType)
	}

	if om.IsError() != false {
		t.Error("Expected Successful Completion Message to generate IsError='false'. It did NOT! IsError='true'.")
	}

	if om.IsFatalError() != false {
		t.Errorf("Expected Successful Completion Message to generate IsFatalError()='false'. It did NOT! IsFatalError()='%v'", om.IsFatalError())
	}

	mId := om.GetMessageId()

	if mId != msgId {
		t.Errorf("Expected message id = '%v'. Instead message id = '%v'.", msgId, mId)
	}

	mNo := om.GetMessageNumber()

	if msgNo != mNo {
		t.Errorf("Expected message number = '%v'. Instead message number = '%v'.", msgNo, mNo)
	}

	actMsg := om.GetFmtMessage()

	if !strings.Contains(actMsg, xMsg) {
		t.Errorf("Expected message to contain '%v'. It did NOT! Actual Message = '%v'",xMsg, actMsg)
	}

	if om.MsgTimeUTC.IsZero()  {
		t.Errorf("Error: om.MsgTimeUTC == Zero. om.MsgTimeUTC== '%v'", om.MsgTimeUTC)
	}

	if om.MsgTimeLocal.IsZero()  {
		t.Errorf("Error: om.MsgTimeLocal == Zero. om.MsgTimeLocal== '%v'",om.MsgTimeLocal)
	}

	if om.MsgLocalTimeZone != "Local" {
		t.Errorf("Error: om.MsgLocalTimeZone is NOT set to 'Local'. om.MsgLocalTimeZone== '%v' ", om.MsgLocalTimeZone)
	}

}

func TestOpsMsgDto_NewSuccessfulCompletionMsg_02(t *testing.T) {

	testMsgContext := testOpsMsgDtoCreateContextInfoObj()

	xMsg := "Successful Completion"
	msgId := int64(64)
	msgNo := int64(6064)
	msgType := OpsMsgTypeSUCCESSFULCOMPLETION

	om := OpsMsgDto{}.InitializeWithMessageContext(testMsgContext).NewSuccessfulCompletionMsg("",msgId)

	if !testMsgContext.Equal(&om.MsgContext) {
		t.Error("Expected testMsgContext to EQUAL om.MsgContext. It did NOT!")
	}

	if om.MsgType != msgType {
		t.Errorf("Expected Messgage Type == '%v'. Instead, Message Type == '%v'.", msgType, om.MsgType)
	}

	if om.IsError() != false {
		t.Error("Expected Successful Completion Message to generate IsError='false'. It did NOT! IsError='true'.")
	}

	if om.IsFatalError() != false {
		t.Errorf("Expected Successful Completion Message to generate IsFatalError()='false'. It did NOT! IsFatalError()='%v'", om.IsFatalError())
	}

	mId := om.GetMessageId()

	if mId != msgId {
		t.Errorf("Expected message id = '%v'. Instead message id = '%v'.", msgId, mId)
	}

	mNo := om.GetMessageNumber()

	if msgNo != mNo {
		t.Errorf("Expected message number = '%v'. Instead message number = '%v'.", msgNo, mNo)
	}

	actMsg := om.GetFmtMessage()

	if !strings.Contains(actMsg, xMsg) {
		t.Errorf("Expected message to contain '%v'. It did NOT! Actual Message = '%v'",xMsg, actMsg)
	}

	if om.MsgTimeUTC.IsZero()  {
		t.Errorf("Error: om.MsgTimeUTC == Zero. om.MsgTimeUTC== '%v'", om.MsgTimeUTC)
	}

	if om.MsgTimeLocal.IsZero()  {
		t.Errorf("Error: om.MsgTimeLocal == Zero. om.MsgTimeLocal== '%v'",om.MsgTimeLocal)
	}

	if om.MsgLocalTimeZone != "Local" {
		t.Errorf("Error: om.MsgLocalTimeZone is NOT set to 'Local'. om.MsgLocalTimeZone== '%v' ", om.MsgLocalTimeZone)
	}

}

func TestOpsMsgDto_NewSuccessfulCompletionMsg_03(t *testing.T) {

	xMsg := "Successful Completion"
	msgId := int64(64)
	msgNo := int64(64)
	msgType := OpsMsgTypeSUCCESSFULCOMPLETION

	om := OpsMsgDto{}.NewSuccessfulCompletionMsg("",msgId)

	if om.MsgType != msgType {
		t.Errorf("Expected Messgage Type == '%v'. Instead, Message Type == '%v'.", msgType, om.MsgType)
	}

	if om.IsError() != false {
		t.Error("Expected Successful Completion Message to generate IsError='false'. It did NOT! IsError='true'.")
	}

	if om.IsFatalError() != false {
		t.Errorf("Expected Successful Completion Message to generate IsFatalError()='false'. It did NOT! IsFatalError()='%v'", om.IsFatalError())
	}

	mId := om.GetMessageId()

	if mId != msgId {
		t.Errorf("Expected message id = '%v'. Instead message id = '%v'.", msgId, mId)
	}

	mNo := om.GetMessageNumber()

	if msgNo != mNo {
		t.Errorf("Expected message number = '%v'. Instead message number = '%v'.", msgNo, mNo)
	}

	actMsg := om.GetFmtMessage()

	if !strings.Contains(actMsg, xMsg) {
		t.Errorf("Expected message to contain '%v'. It did NOT! Actual Message = '%v'",xMsg, actMsg)
	}

	if om.MsgTimeUTC.IsZero()  {
		t.Errorf("Error: om.MsgTimeUTC == Zero. om.MsgTimeUTC== '%v'", om.MsgTimeUTC)
	}

	if om.MsgTimeLocal.IsZero()  {
		t.Errorf("Error: om.MsgTimeLocal == Zero. om.MsgTimeLocal== '%v'",om.MsgTimeLocal)
	}

	if om.MsgLocalTimeZone != "Local" {
		t.Errorf("Error: om.MsgLocalTimeZone is NOT set to 'Local'. om.MsgLocalTimeZone== '%v' ", om.MsgLocalTimeZone)
	}

}

func TestOpsMsgDto_NewWarningMsg_01(t *testing.T) {

	testParentHistory := testOpsMsgDtoCreateParentHistory()
	testMsgContext := testOpsMsgDtoCreateContextInfoObj()

	xMsg := "This is Warning Message for test object."
	msgId := int64(67)
	msgNo := int64(6067)
	msgType := OpsMsgTypeWARNINGMSG

	om := OpsMsgDto{}.InitializeAllContextInfo(testParentHistory, testMsgContext).NewWarningMsg(xMsg, msgId)

	l1 := len(testParentHistory)

	l2 := len(om.ParentContextHistory)

	if l1 != l2 {
		t.Errorf("Expected om.ParentContextHistory to equal length= '%v'. It did NOT! actual length= '%v'",l1, l2)
	}

	for i:=0; i<l1; i++ {
		if !testParentHistory[i].Equal(&om.ParentContextHistory[i]) {
			t.Errorf("Expected om.ParentContextHistory to Equal testParentContext History. It did NOT!. i= '%v'",i)
		}
	}

	if !testMsgContext.Equal(&om.MsgContext) {
		t.Error("Expected testMsgContext to EQUAL om.MsgContext. It did NOT!")
	}


	if om.MsgType != msgType {
		t.Errorf("Expected Messgage Type == '%v'. Instead, Message Type == '%v'.", msgType, om.MsgType)
	}

	if om.IsError() != false {
		t.Error("Expected Warning Message to generate IsError='false'. It did NOT! IsError='true'.")
	}

	if om.IsFatalError() != false {
		t.Errorf("Expected Warning to generate IsFatalError()='false'. It did NOT! IsFatalError()='%v'", om.IsFatalError())
	}

	mId := om.GetMessageId()

	if mId != msgId {
		t.Errorf("Expected message id = '%v'. Instead message id = '%v'.", msgId, mId)
	}

	mNo := om.GetMessageNumber()

	if msgNo != mNo {
		t.Errorf("Expected message number = '%v'. Instead message number = '%v'.", msgNo, mNo)
	}

	actMsg := om.GetFmtMessage()

	if !strings.Contains(actMsg, xMsg) {
		t.Errorf("Expected message to contain '%v'. It did NOT! Actual Message = '%v'",xMsg, actMsg)
	}

	if om.MsgTimeUTC.IsZero()  {
		t.Errorf("Error: om.MsgTimeUTC == Zero. om.MsgTimeUTC== '%v'", om.MsgTimeUTC)
	}

	if om.MsgTimeLocal.IsZero()  {
		t.Errorf("Error: om.MsgTimeLocal == Zero. om.MsgTimeLocal== '%v'",om.MsgTimeLocal)
	}

	if om.MsgLocalTimeZone != "Local" {
		t.Errorf("Error: om.MsgLocalTimeZone is NOT set to 'Local'. om.MsgLocalTimeZone== '%v' ", om.MsgLocalTimeZone)
	}

}

func TestOpsMsgDto_NewWarningMsg_02(t *testing.T) {

	testMsgContext := testOpsMsgDtoCreateContextInfoObj()

	xMsg := "This is Warning Message for test object."
	msgId := int64(67)
	msgNo := int64(6067)
	msgType := OpsMsgTypeWARNINGMSG

	om := OpsMsgDto{}.InitializeWithMessageContext(testMsgContext).NewWarningMsg(xMsg, msgId)

	if !testMsgContext.Equal(&om.MsgContext) {
		t.Error("Expected testMsgContext to EQUAL om.MsgContext. It did NOT!")
	}

	if om.MsgType != msgType {
		t.Errorf("Expected Messgage Type == '%v'. Instead, Message Type == '%v'.", msgType, om.MsgType)
	}

	if om.IsError() != false {
		t.Error("Expected Warning Message to generate IsError='false'. It did NOT! IsError='true'.")
	}

	if om.IsFatalError() != false {
		t.Errorf("Expected Warning to generate IsFatalError()='false'. It did NOT! IsFatalError()='%v'", om.IsFatalError())
	}

	mId := om.GetMessageId()

	if mId != msgId {
		t.Errorf("Expected message id = '%v'. Instead message id = '%v'.", msgId, mId)
	}

	mNo := om.GetMessageNumber()

	if msgNo != mNo {
		t.Errorf("Expected message number = '%v'. Instead message number = '%v'.", msgNo, mNo)
	}

	actMsg := om.GetFmtMessage()

	if !strings.Contains(actMsg, xMsg) {
		t.Errorf("Expected message to contain '%v'. It did NOT! Actual Message = '%v'",xMsg, actMsg)
	}

	if om.MsgTimeUTC.IsZero()  {
		t.Errorf("Error: om.MsgTimeUTC == Zero. om.MsgTimeUTC== '%v'", om.MsgTimeUTC)
	}

	if om.MsgTimeLocal.IsZero()  {
		t.Errorf("Error: om.MsgTimeLocal == Zero. om.MsgTimeLocal== '%v'",om.MsgTimeLocal)
	}

	if om.MsgLocalTimeZone != "Local" {
		t.Errorf("Error: om.MsgLocalTimeZone is NOT set to 'Local'. om.MsgLocalTimeZone== '%v' ", om.MsgLocalTimeZone)
	}

}

func TestOpsMsgDto_NewWarningMsg_03(t *testing.T) {

	xMsg := "This is Warning Message for test object."
	msgId := int64(67)
	msgNo := int64(67)
	msgType := OpsMsgTypeWARNINGMSG

	om := OpsMsgDto{}.NewWarningMsg(xMsg, msgId)

	if om.MsgType != msgType {
		t.Errorf("Expected Messgage Type == '%v'. Instead, Message Type == '%v'.", msgType, om.MsgType)
	}

	if om.IsError() != false {
		t.Error("Expected Warning Message to generate IsError='false'. It did NOT! IsError='true'.")
	}

	if om.IsFatalError() != false {
		t.Errorf("Expected Warning to generate IsFatalError()='false'. It did NOT! IsFatalError()='%v'", om.IsFatalError())
	}

	mId := om.GetMessageId()

	if mId != msgId {
		t.Errorf("Expected message id = '%v'. Instead message id = '%v'.", msgId, mId)
	}

	mNo := om.GetMessageNumber()

	if msgNo != mNo {
		t.Errorf("Expected message number = '%v'. Instead message number = '%v'.", msgNo, mNo)
	}

	actMsg := om.GetFmtMessage()

	if !strings.Contains(actMsg, xMsg) {
		t.Errorf("Expected message to contain '%v'. It did NOT! Actual Message = '%v'",xMsg, actMsg)
	}

	if om.MsgTimeUTC.IsZero()  {
		t.Errorf("Error: om.MsgTimeUTC == Zero. om.MsgTimeUTC== '%v'", om.MsgTimeUTC)
	}

	if om.MsgTimeLocal.IsZero()  {
		t.Errorf("Error: om.MsgTimeLocal == Zero. om.MsgTimeLocal== '%v'",om.MsgTimeLocal)
	}

	if om.MsgLocalTimeZone != "Local" {
		t.Errorf("Error: om.MsgLocalTimeZone is NOT set to 'Local'. om.MsgLocalTimeZone== '%v' ", om.MsgLocalTimeZone)
	}

}

func TestOpsMsgDto_ParentHistory_01(t *testing.T) {

	ci := OpsMsgContextInfo{}

	x1 := ci.New("TSource01", "PObj01", "Func001", 1000)
	x2 := ci.New("TSource02", "PObj02", "Func002", 2000)
	x3 := ci.New("TSource03", "PObj02", "Func003", 3000)

	om := OpsMsgDto{}
	om.ParentContextHistory = append(om.ParentContextHistory, x1)
	om.ParentContextHistory = append(om.ParentContextHistory, x2)
	om.ParentContextHistory = append(om.ParentContextHistory, x3)

	l := len(om.ParentContextHistory)

	if l != 3 {
		t.Errorf("Expected Parent Context History Length = 3. Instead, Parent Context History Lenth = '%v'", l)
	}

	if om.ParentContextHistory[2].FuncName != "Func003" {
		t.Errorf("Expected 3rd OpsMsgContextInfo object in Parent Context History FuncName= 'Func003'. Instead, FuncName= '%v'",om.ParentContextHistory[2].FuncName)
	}

}

func TestOpsMsgDto_ParentHistory_02(t *testing.T) {

	parent := testOpsMsgDtoCreateParentHistory()

	om := OpsMsgDto{}
	om.AddParentContextHistory(parent)

	l := len(om.ParentContextHistory)

	if l != 5 {
		t.Errorf("Expected Parent Context History Length = 5. Instead, Parent Context History Lenth = '%v'", l)
	}

	if om.ParentContextHistory[0].SourceFileName != "TSource01" {
		t.Errorf("Expected 1st OpsMsgContextInfo object in Parent Context History SourceFileName= 'TSource01'. Instead, SourceFileName= '%v'",om.ParentContextHistory[0].SourceFileName)
	}

	if om.ParentContextHistory[0].ParentObjectName != "PObj01" {
		t.Errorf("Expected 1st OpsMsgContextInfo object in Parent Context History ParentObjectName= 'PObj01'. Instead, ParentObjectName= '%v'",om.ParentContextHistory[0].ParentObjectName)
	}

	if om.ParentContextHistory[0].FuncName != "Func001" {
		t.Errorf("Expected 1st OpsMsgContextInfo object in Parent Context History FuncName= 'Func001'. Instead, ParentObjectName= '%v'",om.ParentContextHistory[0].FuncName)
	}

	if om.ParentContextHistory[0].BaseMessageId != int64(1000) {
		t.Errorf("Expected 1st OpsMsgContextInfo object in Parent Context History BaseMessageId = 1000. Instead, BaseMessageId = '%v'", om.ParentContextHistory[0].BaseMessageId)
	}

	if om.ParentContextHistory[1].SourceFileName != "TSource02" {
		t.Errorf("Expected 2nd OpsMsgContextInfo object in Parent Context History SourceFileName= 'TSource02'. Instead, SourceFileName= '%v'",om.ParentContextHistory[1].SourceFileName)
	}


	if om.ParentContextHistory[1].ParentObjectName != "PObj02" {
		t.Errorf("Expected 2nd OpsMsgContextInfo object in Parent Context History ParentObjectName= 'PObj02'. Instead, ParentObjectName= '%v'",om.ParentContextHistory[1].ParentObjectName)
	}

	if om.ParentContextHistory[1].FuncName != "Func002" {
		t.Errorf("Expected 2nd OpsMsgContextInfo object in Parent Context History FuncName= 'Func002'. Instead, ParentObjectName= '%v'",om.ParentContextHistory[1].FuncName)
	}

	if om.ParentContextHistory[1].BaseMessageId != int64(2000) {
		t.Errorf("Expected 2nd OpsMsgContextInfo object in Parent Context History BaseMessageId = 2000. Instead, BaseMessageId = '%v'", om.ParentContextHistory[1].BaseMessageId)
	}

	if om.ParentContextHistory[2].SourceFileName != "TSource03" {
		t.Errorf("Expected 3rd OpsMsgContextInfo object in Parent Context History SourceFileName= 'TSource03'. Instead, SourceFileName= '%v'",om.ParentContextHistory[2].SourceFileName)
	}

	if om.ParentContextHistory[2].ParentObjectName != "PObj03" {
		t.Errorf("Expected 3rd OpsMsgContextInfo object in Parent Context History ParentObjectName= 'PObj03'. Instead, ParentObjectName= '%v'",om.ParentContextHistory[2].ParentObjectName)
	}

	if om.ParentContextHistory[2].FuncName != "Func003" {
		t.Errorf("Expected 3rd OpsMsgContextInfo object in Parent Context History FuncName= 'Func003'. Instead, FuncName= '%v'",om.ParentContextHistory[2].FuncName)
	}

	if om.ParentContextHistory[2].BaseMessageId != int64(3000) {
		t.Errorf("Expected 3rd OpsMsgContextInfo object in Parent Context History BaseMessageId = 3000. Instead, BaseMessageId = '%v'", om.ParentContextHistory[2].BaseMessageId)
	}


	if om.ParentContextHistory[3].SourceFileName != "TSource04" {
		t.Errorf("Expected 4th OpsMsgContextInfo object in Parent Context History SourceFileName= 'TSource04'. Instead, SourceFileName= '%v'",om.ParentContextHistory[3].SourceFileName)
	}

	if om.ParentContextHistory[3].ParentObjectName != "PObj04" {
		t.Errorf("Expected 4th OpsMsgContextInfo object in Parent Context History ParentObjectName= 'PObj04'. Instead, ParentObjectName= '%v'",om.ParentContextHistory[3].ParentObjectName)
	}

	if om.ParentContextHistory[3].FuncName != "Func004" {
		t.Errorf("Expected 4th OpsMsgContextInfo object in Parent Context History FuncName= 'Func004'. Instead, FuncName= '%v'",om.ParentContextHistory[3].FuncName)
	}

	if om.ParentContextHistory[3].BaseMessageId != int64(4000) {
		t.Errorf("Expected 4th OpsMsgContextInfo object in Parent Context History BaseMessageId = 4000. Instead, BaseMessageId = '%v'", om.ParentContextHistory[3].BaseMessageId)
	}

	if om.ParentContextHistory[4].SourceFileName != "TSource05" {
		t.Errorf("Expected 5th OpsMsgContextInfo object in Parent Context History SourceFileName= 'TSource05'. Instead, SourceFileName= '%v'",om.ParentContextHistory[4].SourceFileName)
	}

	if om.ParentContextHistory[4].ParentObjectName != "PObj05" {
		t.Errorf("Expected 5th OpsMsgContextInfo object in Parent Context History ParentObjectName = 'PObj05'. Instead, ParentObjectName = '%v'", om.ParentContextHistory[4].ParentObjectName)
	}

	if om.ParentContextHistory[4].FuncName != "Func005" {
		t.Errorf("Expected 5th OpsMsgContextInfo object in Parent Context History FuncName= 'Func005'. Instead, FuncName= '%v'",om.ParentContextHistory[4].FuncName)
	}

	if om.ParentContextHistory[4].BaseMessageId != int64(5000) {
		t.Errorf("Expected 5th OpsMsgContextInfo object in Parent Context History BaseMessageId = 5000. Instead, BaseMessageId = '%v'", om.ParentContextHistory[4].BaseMessageId)
	}

}

func TestOpsMsgDto_SetDebugMessage_01(t *testing.T) {

	testParentHistory := testOpsMsgDtoCreateParentHistory()
	testMsgContext := testOpsMsgDtoCreateContextInfoObj()

	om := testOpsMsgDtoCreateDebugMsg()

	xMsg := "This is DEBUG Message for test object."
	msgId := int64(238)
	msgNo := int64(6238)
	msgType := OpsMsgTypeDEBUGMSG

	l1 := len(testParentHistory)

	l2 := len(om.ParentContextHistory)

	if l1 != l2 {
		t.Errorf("Expected om.ParentContextHistory to equal length= '%v'. It did NOT! actual length= '%v'",l1, l2)
	}

	for i:=0; i<l1; i++ {
		if !testParentHistory[i].Equal(&om.ParentContextHistory[i]) {
			t.Errorf("Expected om.ParentContextHistory to Equal testParentContext History. It did NOT!. i= '%v'",i)
		}
	}

	if !testMsgContext.Equal(&om.MsgContext) {
		t.Error("Expected testMsgContext to EQUAL om.MsgContext. It did NOT!")
	}

	if om.MsgType != msgType {
		t.Errorf("Expected Messgage Type == '%v'. Instead, Message Type == '%v'.", msgType, om.MsgType)
	}

	if om.IsError() != false {
		t.Error("Expected Debug Message to generate IsError='false'. It did NOT! IsError='true'.")
	}

	if om.IsFatalError() != false {
		t.Errorf("Expected Debug to generate IsFatalError()='false'. It did NOT! IsFatalError()='%v'", om.IsFatalError())
	}

	mId := om.GetMessageId()

	if mId != msgId {
		t.Errorf("Expected message id = '%v'. Instead message id = '%v'.", msgId, mId)
	}

	mNo := om.GetMessageNumber()

	if msgNo != mNo {
		t.Errorf("Expected message number = '%v'. Instead message number = '%v'.", msgNo, mNo)
	}

	actMsg := om.GetFmtMessage()

	if !strings.Contains(actMsg, xMsg) {
		t.Errorf("Expected message to contain '%v'. It did NOT! Actual Message = '%v'",xMsg, actMsg)
	}

	if om.MsgTimeUTC.IsZero()  {
		t.Errorf("Error: om.MsgTimeUTC == Zero. om.MsgTimeUTC== '%v'", om.MsgTimeUTC)
	}

	if om.MsgTimeLocal.IsZero()  {
		t.Errorf("Error: om.MsgTimeLocal == Zero. om.MsgTimeLocal== '%v'",om.MsgTimeLocal)
	}

	if om.MsgLocalTimeZone != "Local" {
		t.Errorf("Error: om.MsgLocalTimeZone is NOT set to 'Local'. om.MsgLocalTimeZone== '%v' ", om.MsgLocalTimeZone)
	}

}

func TestOpsMsgDto_SetDebugMessage_02(t *testing.T) {

	testMsgContext := testOpsMsgDtoCreateContextInfoObj()

	om := OpsMsgDto{}

	xMsg := "This is DEBUG Message for test object."
	msgId := int64(238)
	msgNo := int64(6238)
	msgType := OpsMsgTypeDEBUGMSG

	om.SetMessageContext(testOpsMsgDtoCreateContextInfoObj())
	om.SetDebugMessage(xMsg, msgId)

	if !testMsgContext.Equal(&om.MsgContext) {
		t.Error("Expected testMsgContext to EQUAL om.MsgContext. It did NOT!")
	}

	if om.MsgType != msgType {
		t.Errorf("Expected Messgage Type == '%v'. Instead, Message Type == '%v'.", msgType, om.MsgType)
	}

	if om.IsError() != false {
		t.Error("Expected Debug Message to generate IsError='false'. It did NOT! IsError='true'.")
	}

	if om.IsFatalError() != false {
		t.Errorf("Expected Debug to generate IsFatalError()='false'. It did NOT! IsFatalError()='%v'", om.IsFatalError())
	}

	mId := om.GetMessageId()

	if mId != msgId {
		t.Errorf("Expected message id = '%v'. Instead message id = '%v'.", msgId, mId)
	}

	mNo := om.GetMessageNumber()

	if msgNo != mNo {
		t.Errorf("Expected message number = '%v'. Instead message number = '%v'.", msgNo, mNo)
	}

	actMsg := om.GetFmtMessage()

	if !strings.Contains(actMsg, xMsg) {
		t.Errorf("Expected message to contain '%v'. It did NOT! Actual Message = '%v'",xMsg, actMsg)
	}

	if om.MsgTimeUTC.IsZero()  {
		t.Errorf("Error: om.MsgTimeUTC == Zero. om.MsgTimeUTC== '%v'", om.MsgTimeUTC)
	}

	if om.MsgTimeLocal.IsZero()  {
		t.Errorf("Error: om.MsgTimeLocal == Zero. om.MsgTimeLocal== '%v'",om.MsgTimeLocal)
	}

	if om.MsgLocalTimeZone != "Local" {
		t.Errorf("Error: om.MsgLocalTimeZone is NOT set to 'Local'. om.MsgLocalTimeZone== '%v' ", om.MsgLocalTimeZone)
	}

}

func TestOpsMsgDto_SetDebugMessage_03(t *testing.T) {

	om := OpsMsgDto{}

	xMsg := "This is DEBUG Message for test object."
	msgId := int64(238)
	msgNo := int64(238)
	msgType := OpsMsgTypeDEBUGMSG

	om.SetDebugMessage(xMsg, msgId)


	if om.MsgType != msgType {
		t.Errorf("Expected Messgage Type == '%v'. Instead, Message Type == '%v'.", msgType, om.MsgType)
	}

	if om.IsError() != false {
		t.Error("Expected Debug Message to generate IsError='false'. It did NOT! IsError='true'.")
	}

	if om.IsFatalError() != false {
		t.Errorf("Expected Debug to generate IsFatalError()='false'. It did NOT! IsFatalError()='%v'", om.IsFatalError())
	}

	mId := om.GetMessageId()

	if mId != msgId {
		t.Errorf("Expected message id = '%v'. Instead message id = '%v'.", msgId, mId)
	}

	mNo := om.GetMessageNumber()

	if msgNo != mNo {
		t.Errorf("Expected message number = '%v'. Instead message number = '%v'.", msgNo, mNo)
	}

	actMsg := om.GetFmtMessage()

	if !strings.Contains(actMsg, xMsg) {
		t.Errorf("Expected message to contain '%v'. It did NOT! Actual Message = '%v'",xMsg, actMsg)
	}

	if om.MsgTimeUTC.IsZero()  {
		t.Errorf("Error: om.MsgTimeUTC == Zero. om.MsgTimeUTC== '%v'", om.MsgTimeUTC)
	}

	if om.MsgTimeLocal.IsZero()  {
		t.Errorf("Error: om.MsgTimeLocal == Zero. om.MsgTimeLocal== '%v'",om.MsgTimeLocal)
	}

	if om.MsgLocalTimeZone != "Local" {
		t.Errorf("Error: om.MsgLocalTimeZone is NOT set to 'Local'. om.MsgLocalTimeZone== '%v' ", om.MsgLocalTimeZone)
	}

}

func TestOpsMsgDto_SetStdError_01(t *testing.T) {

	testParentHistory := testOpsMsgDtoCreateParentHistory()

	testMsgCtx := testOpsMsgDtoCreateContextInfoObj()

	prefixMsgText := "This is Prefix Msg: "
	errMsgText := "This is Standard Error Msg for test object"
	err := errors.New(errMsgText)
	xMsg := prefixMsgText + errMsgText
	msgId := int64(429)
	msgNo := int64(6429)
	msgType := OpsMsgTypeERRORMSG

	om := OpsMsgDto{}.InitializeAllContextInfo(testParentHistory, testMsgCtx)
	om.SetStdError(prefixMsgText, err, msgId)

	l1 := len(testParentHistory)

	l2 := len(om.ParentContextHistory)

	if l1 != l2 {
		t.Errorf("Expected om.ParentContextHistory to equal length= '%v'. It did NOT! actual length= '%v'",l1, l2)
	}

	for i:=0; i<l1; i++ {
		if !testParentHistory[i].Equal(&om.ParentContextHistory[i]) {
			t.Errorf("Expected om.ParentContextHistory to Equal testParentContext History. It did NOT!. i= '%v'",i)
		}
	}

	if !testMsgCtx.Equal(&om.MsgContext) {
		t.Error("Expected testMsgCtx to EQUAL om.MsgContext. It did NOT!")
	}

	if om.MsgType != msgType {
		t.Errorf("Expected Messgage Type == '%v'. Instead, Message Type == '%v'.", msgType, om.MsgType)
	}

	if om.IsError() != true {
		t.Error("Expected error msg to generate IsError='true'. It did NOT! IsError='false'.")
	}

	if om.IsFatalError() == true {
		t.Error("Expected standard error msg to generate IsFatalError()='false'. It did NOT! IsFatalError()='true'")
	}

	mId := om.GetMessageId()

	if mId != msgId {
		t.Errorf("Expected message id = '%v'. Instead message id = '%v'.", msgId, mId)
	}

	mNo := om.GetMessageNumber()

	if msgNo != mNo {
		t.Errorf("Expected message number = '%v'. Instead message number = '%v'.", msgNo, mNo)
	}

	actMsg := om.GetFmtMessage()

	if !strings.Contains(actMsg, xMsg) {
		t.Errorf("Expected message to contain '%v'. It did NOT! Actual Message = '%v'",xMsg, actMsg)
	}

	if om.MsgTimeUTC.IsZero()  {
		t.Errorf("Error: om.MsgTimeUTC == Zero. om.MsgTimeUTC== '%v'", om.MsgTimeUTC)
	}

	if om.MsgTimeLocal.IsZero()  {
		t.Errorf("Error: om.MsgTimeLocal == Zero. om.MsgTimeLocal== '%v'",om.MsgTimeLocal)
	}

	if om.MsgLocalTimeZone != "Local" {
		t.Errorf("Error: om.MsgLocalTimeZone is NOT set to 'Local'. om.MsgLocalTimeZone== '%v' ", om.MsgLocalTimeZone)
	}

}

func TestOpsMsgDto_SetStdError_02(t *testing.T) {

	testParentHistory := testOpsMsgDtoCreateParentHistory()

	testMsgCtx := testOpsMsgDtoCreateContextInfoObj()

	prefixMsgText := ""
	errMsgText := "This is Standard Error Msg for test object"
	err := errors.New(errMsgText)
	xMsg := errMsgText
	msgId := int64(429)
	msgNo := int64(6429)
	msgType := OpsMsgTypeERRORMSG

	om := OpsMsgDto{}.InitializeAllContextInfo(testParentHistory, testMsgCtx)
	om.SetStdError(prefixMsgText, err, msgId)

	l1 := len(testParentHistory)

	l2 := len(om.ParentContextHistory)

	if l1 != l2 {
		t.Errorf("Expected om.ParentContextHistory to equal length= '%v'. It did NOT! actual length= '%v'",l1, l2)
	}

	for i:=0; i<l1; i++ {
		if !testParentHistory[i].Equal(&om.ParentContextHistory[i]) {
			t.Errorf("Expected om.ParentContextHistory to Equal testParentContext History. It did NOT!. i= '%v'",i)
		}
	}

	if !testMsgCtx.Equal(&om.MsgContext) {
		t.Error("Expected testMsgCtx to EQUAL om.MsgContext. It did NOT!")
	}

	if om.MsgType != msgType {
		t.Errorf("Expected Messgage Type == '%v'. Instead, Message Type == '%v'.", msgType, om.MsgType)
	}

	if om.IsError() != true {
		t.Error("Expected error msg to generate IsError='true'. It did NOT! IsError='false'.")
	}

	if om.IsFatalError() == true {
		t.Error("Expected standard error msg to generate IsFatalError()='false'. It did NOT! IsFatalError()='true'")
	}

	mId := om.GetMessageId()

	if mId != msgId {
		t.Errorf("Expected message id = '%v'. Instead message id = '%v'.", msgId, mId)
	}

	mNo := om.GetMessageNumber()

	if msgNo != mNo {
		t.Errorf("Expected message number = '%v'. Instead message number = '%v'.", msgNo, mNo)
	}

	actMsg := om.GetFmtMessage()

	if !strings.Contains(actMsg, xMsg) {
		t.Errorf("Expected message to contain '%v'. It did NOT! Actual Message = '%v'",xMsg, actMsg)
	}

	if om.MsgTimeUTC.IsZero()  {
		t.Errorf("Error: om.MsgTimeUTC == Zero. om.MsgTimeUTC== '%v'", om.MsgTimeUTC)
	}

	if om.MsgTimeLocal.IsZero()  {
		t.Errorf("Error: om.MsgTimeLocal == Zero. om.MsgTimeLocal== '%v'",om.MsgTimeLocal)
	}

	if om.MsgLocalTimeZone != "Local" {
		t.Errorf("Error: om.MsgLocalTimeZone is NOT set to 'Local'. om.MsgLocalTimeZone== '%v' ", om.MsgLocalTimeZone)
	}

}

func TestOpsMsgDto_SetStdErrorMessage_01(t *testing.T) {

	testParentHistory := testOpsMsgDtoCreateParentHistory()

	testMsgCtx := testOpsMsgDtoCreateContextInfoObj()

	om := testOpsMsgDtoCreateStdErrorMsg()

	xMsg := "This is Standard Error Msg for test object"
	msgId := int64(429)
	msgNo := int64(6429)
	msgType := OpsMsgTypeERRORMSG

	l1 := len(testParentHistory)

	l2 := len(om.ParentContextHistory)

	if l1 != l2 {
		t.Errorf("Expected om.ParentContextHistory to equal length= '%v'. It did NOT! actual length= '%v'",l1, l2)
	}

	for i:=0; i<l1; i++ {
		if !testParentHistory[i].Equal(&om.ParentContextHistory[i]) {
			t.Errorf("Expected om.ParentContextHistory to Equal testParentContext History. It did NOT!. i= '%v'",i)
		}
	}

	if !testMsgCtx.Equal(&om.MsgContext) {
		t.Error("Expected testMsgCtx to EQUAL om.MsgContext. It did NOT!")
	}

	if om.MsgType != msgType {
		t.Errorf("Expected Messgage Type == '%v'. Instead, Message Type == '%v'.", msgType, om.MsgType)
	}

	if om.IsError() != true {
		t.Error("Expected error msg to generate IsError='true'. It did NOT! IsError='false'.")
	}

	if om.IsFatalError() == true {
		t.Error("Expected standard error msg to generate IsFatalError()='false'. It did NOT! IsFatalError()='true'")
	}

	mId := om.GetMessageId()

	if mId != msgId {
		t.Errorf("Expected message id = '%v'. Instead message id = '%v'.", msgId, mId)
	}

	mNo := om.GetMessageNumber()

	if msgNo != mNo {
		t.Errorf("Expected message number = '%v'. Instead message number = '%v'.", msgNo, mNo)
	}

	actMsg := om.GetFmtMessage()

	if !strings.Contains(actMsg, xMsg) {
		t.Errorf("Expected message to contain '%v'. It did NOT! Actual Message = '%v'",xMsg, actMsg)
	}

	if om.MsgTimeUTC.IsZero()  {
		t.Errorf("Error: om.MsgTimeUTC == Zero. om.MsgTimeUTC== '%v'", om.MsgTimeUTC)
	}

	if om.MsgTimeLocal.IsZero()  {
		t.Errorf("Error: om.MsgTimeLocal == Zero. om.MsgTimeLocal== '%v'",om.MsgTimeLocal)
	}

	if om.MsgLocalTimeZone != "Local" {
		t.Errorf("Error: om.MsgLocalTimeZone is NOT set to 'Local'. om.MsgLocalTimeZone== '%v' ", om.MsgLocalTimeZone)
	}

}

func TestOpsMsgDto_SetStdErrorMessage_02(t *testing.T) {
	xMsg := "This is Standard Error Msg for test object"
	msgId := int64(429)
	msgNo := int64(6429)
	msgType := OpsMsgTypeERRORMSG
	baseMsgId := int64(6000)
	srcFile := "TSource06"

	om := OpsMsgDto{}
	mCtx := testOpsMsgDtoCreateContextInfoObj()
	om.SetMessageContext(mCtx)
	om.SetStdErrorMessage(xMsg, msgId)

	if om.MsgContext.BaseMessageId != baseMsgId {
		t.Errorf("Expected om.MsgContext.BaseMessageId == '%v'. Instead, om.MsgContext.BaseMessageId == '%v'", baseMsgId, om.MsgContext.BaseMessageId )
	}

	if om.MsgContext.SourceFileName != srcFile{
		t.Errorf("Expected om.MsgContext.SourceFileName == '%v'. Instead, om.MsgContext.BaseMessageId == '%v'", srcFile, om.MsgContext.SourceFileName )
	}

	if !om.MsgContext.Equal(&mCtx) {
		t.Error("Expected om.MsgContext==mCtx. It did NOT!")
	}

	if om.MsgType != msgType {
		t.Errorf("Expected Messgage Type == '%v'. Instead, Message Type == '%v'.", msgType, om.MsgType)
	}

	if om.IsError() != true {
		t.Error("Expected error msg to generate IsError='true'. It did NOT! IsError='false'.")
	}

	if om.IsFatalError() == true {
		t.Error("Expected standard error msg to generate IsFatalError()='false'. It did NOT! IsFatalError()='true'")
	}

	mId := om.GetMessageId()

	if mId != msgId {
		t.Errorf("Expected message id = '%v'. Instead message id = '%v'.", msgId, mId)
	}

	mNo := om.GetMessageNumber()

	if msgNo != mNo {
		t.Errorf("Expected message number = '%v'. Instead message number = '%v'.", msgNo, mNo)
	}

	actMsg := om.GetFmtMessage()

	if !strings.Contains(actMsg, xMsg) {
		t.Errorf("Expected message to contain '%v'. It did NOT! Actual Message = '%v'",xMsg, actMsg)
	}

	if om.MsgTimeUTC.IsZero()  {
		t.Errorf("Error: om.MsgTimeUTC == Zero. om.MsgTimeUTC== '%v'", om.MsgTimeUTC)
	}

	if om.MsgTimeLocal.IsZero()  {
		t.Errorf("Error: om.MsgTimeLocal == Zero. om.MsgTimeLocal== '%v'",om.MsgTimeLocal)
	}

	if om.MsgLocalTimeZone != "Local" {
		t.Errorf("Error: om.MsgLocalTimeZone is NOT set to 'Local'. om.MsgLocalTimeZone== '%v' ", om.MsgLocalTimeZone)
	}

}

func TestOpsMsgDto_SetStdErrorMessage_03(t *testing.T) {
	xMsg := "This is Standard Error Msg for test object"
	msgId := int64(429)
	msgNo := int64(429)
	msgType := OpsMsgTypeERRORMSG

	om := OpsMsgDto{}
	om.SetStdErrorMessage(xMsg, msgId)

	if om.MsgType != msgType {
		t.Errorf("Expected Messgage Type == '%v'. Instead, Message Type == '%v'.", msgType, om.MsgType)
	}

	if om.IsError() != true {
		t.Error("Expected error msg to generate IsError='true'. It did NOT! IsError='false'.")
	}

	if om.IsFatalError() == true {
		t.Error("Expected standard error msg to generate IsFatalError()='false'. It did NOT! IsFatalError()='true'")
	}

	mId := om.GetMessageId()

	if mId != msgId {
		t.Errorf("Expected message id = '%v'. Instead message id = '%v'.", msgId, mId)
	}

	mNo := om.GetMessageNumber()

	if msgNo != mNo {
		t.Errorf("Expected message number = '%v'. Instead message number = '%v'.", msgNo, mNo)
	}

	actMsg := om.GetFmtMessage()

	if !strings.Contains(actMsg, xMsg) {
		t.Errorf("Expected message to contain '%v'. It did NOT! Actual Message = '%v'",xMsg, actMsg)
	}

	if om.MsgTimeUTC.IsZero()  {
		t.Errorf("Error: om.MsgTimeUTC == Zero. om.MsgTimeUTC== '%v'", om.MsgTimeUTC)
	}

	if om.MsgTimeLocal.IsZero()  {
		t.Errorf("Error: om.MsgTimeLocal == Zero. om.MsgTimeLocal== '%v'",om.MsgTimeLocal)
	}

	if om.MsgLocalTimeZone != "Local" {
		t.Errorf("Error: om.MsgLocalTimeZone is NOT set to 'Local'. om.MsgLocalTimeZone== '%v' ", om.MsgLocalTimeZone)
	}

}
