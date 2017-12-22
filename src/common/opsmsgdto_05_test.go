package common

import (
	"testing"
	"strings"
	"errors"
)


func TestOpsMsgDto_SetFatalErrorMessage_01(t *testing.T) {

	om := testOpsMsgDtoCreateFatalErrorMsg()

	xMsg := "This is FATAL Error Msg for test object"
	msgId := int64(152)
	msgNo := int64(6152)
	msgType := OpsMsgTypeFATALERRORMSG

	if om.MsgType != msgType {
		t.Errorf("Expected Messgage Type == '%v'. Instead, Message Type == '%v'.", msgType, om.MsgType)
	}

	if om.IsError() != true {
		t.Errorf("Expected Fatal Error Message to generate IsError='true'. It did NOT! IsError='false'.")
	}

	if om.IsFatalError() != true {
		t.Errorf("Expected Fatal Error Message to generate IsFatalError()='true'. It did NOT! IsFatalError()='%v'", om.IsFatalError())
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

func TestOpsMsgDto_SetFatalErrorMessage_02(t *testing.T) {

	om := OpsMsgDto{}


	xMsg := "This is FATAL Error Msg for test object"
	msgId := int64(152)
	msgNo := int64(6152)
	msgType := OpsMsgTypeFATALERRORMSG

	om.SetMessageContext(testOpsMsgDtoCreateContextInfoObj())
	om.SetFatalErrorMessage(xMsg, msgId)

	if om.MsgType != msgType {
		t.Errorf("Expected Messgage Type == '%v'. Instead, Message Type == '%v'.", msgType, om.MsgType)
	}

	if om.IsError() != true {
		t.Errorf("Expected Fatal Error Message to generate IsError='true'. It did NOT! IsError='false'.")
	}

	if om.IsFatalError() != true {
		t.Errorf("Expected Fatal Error Message to generate IsFatalError()='true'. It did NOT! IsFatalError()='%v'", om.IsFatalError())
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

func TestOpsMsgDto_SetFatalErrorMessage_03(t *testing.T) {

	om := OpsMsgDto{}


	xMsg := "This is FATAL Error Msg for test object"
	msgId := int64(152)
	msgNo := int64(152)
	msgType := OpsMsgTypeFATALERRORMSG

	om.SetFatalErrorMessage(xMsg, msgId)

	if om.MsgType != msgType {
		t.Errorf("Expected Messgage Type == '%v'. Instead, Message Type == '%v'.", msgType, om.MsgType)
	}

	if om.IsError() != true {
		t.Errorf("Expected Fatal Error Message to generate IsError='true'. It did NOT! IsError='false'.")
	}

	if om.IsFatalError() != true {
		t.Errorf("Expected Fatal Error Message to generate IsFatalError()='true'. It did NOT! IsFatalError()='%v'", om.IsFatalError())
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

func TestOpsMsgDto_SetFatalError_01(t *testing.T) {

	errText := "New FATAL error from error type."
	prefixMsgText := "Prefix Msg: "
	xMsg := prefixMsgText + errText
	err := errors.New(errText)
	testParentHistory := testOpsMsgDtoCreateParentHistory()
	testMsgContext := testOpsMsgDtoCreateContextInfoObj()
	msgId := int64(152)
	msgNo := int64(6152)
	msgType := OpsMsgTypeFATALERRORMSG
	om := OpsMsgDto{}.InitializeAllContextInfo(testParentHistory, testMsgContext)
	om.SetFatalError(prefixMsgText, err, msgId)

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

	if om.IsError() != true {
		t.Errorf("Expected Fatal Error Message to generate IsError='true'. It did NOT! IsError='false'.")
	}

	if om.IsFatalError() != true {
		t.Errorf("Expected Fatal Error Message to generate IsFatalError()='true'. It did NOT! IsFatalError()='%v'", om.IsFatalError())
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

func TestOpsMsgDto_SetFatalError_02(t *testing.T) {

	errText := "New FATAL error from error type."
	prefixMsgText := ""
	xMsg := errText
	err := errors.New(errText)
	testParentHistory := testOpsMsgDtoCreateParentHistory()
	testMsgContext := testOpsMsgDtoCreateContextInfoObj()
	msgId := int64(152)
	msgNo := int64(6152)
	msgType := OpsMsgTypeFATALERRORMSG
	om := OpsMsgDto{}.InitializeAllContextInfo(testParentHistory, testMsgContext)
	om.SetFatalError(prefixMsgText, err, msgId)

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

	if om.IsError() != true {
		t.Errorf("Expected Fatal Error Message to generate IsError='true'. It did NOT! IsError='false'.")
	}

	if om.IsFatalError() != true {
		t.Errorf("Expected Fatal Error Message to generate IsFatalError()='true'. It did NOT! IsFatalError()='%v'", om.IsFatalError())
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

func TestOpsMsgDto_SetFromSpecErrMessage_01(t *testing.T) {

	xMsg := "FATAL ERROR - Configured in SpecErr"
	msgId := int64(123)
	msgNo := int64(6123)
	msgType := OpsMsgTypeFATALERRORMSG
	parentHistory := testOpsMsgDtoCreateSpecErrParentBaseInfo5Elements()
	baseInfo:= testOpsMsgDtoCreateSpecErrBaseInfoObject()
	se := SpecErr{}.InitializeBaseInfo(parentHistory, baseInfo)
	se.SetFatalError(xMsg, msgId)

	om := OpsMsgDto{}
	om.SetFromSpecErrMessage(se)

	if om.MsgType != msgType {
		t.Errorf("Expected Messgage Type == '%v'. Instead, Message Type == '%v'.", msgType, om.MsgType)
	}

	if om.IsError() != true {
		t.Errorf("Expected Fatal Error Message to generate IsError='true'. It did NOT! IsError='false'.")
	}

	if om.IsFatalError() != true {
		t.Errorf("Expected Fatal Error Message to generate IsFatalError()='true'. It did NOT! IsFatalError()='%v'", om.IsFatalError())
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

	if om.MsgContext.SourceFileName != "TSource06" {
		t.Errorf("Expected MsgContext.SourceFileName == 'TSource06'. Instead, SourceFileName== '%v'", om.MsgContext.SourceFileName)
	}

	if om.MsgContext.ParentObjectName != "PObj06" {
		t.Errorf("Expected MsgContext.ParentObjectName == 'PObj06'. Instead, ParentObjectName== '%v'", om.MsgContext.ParentObjectName)
	}

	if om.MsgContext.FuncName != "Func006" {
		t.Errorf("Expected MsgContext.FuncName == 'Func006'. Instead, FuncName== '%v'", om.MsgContext.FuncName)
	}

	if om.MsgContext.BaseMessageId != 6000 {
		t.Errorf("Expected MsgContext.BaseMessageId == '6000'. Instead, BaseMessageId== '%v'", om.MsgContext.BaseMessageId)
	}

}

func TestOpsMsgDto_SetFromSpecErrMessage_02(t *testing.T) {

	xMsg := "STANDARD ERROR - Configured in SpecErr"
	msgId := int64(123)
	msgNo := int64(6123)
	msgType := OpsMsgTypeERRORMSG
	parentHistory := testOpsMsgDtoCreateSpecErrParentBaseInfo5Elements()
	baseInfo:= testOpsMsgDtoCreateSpecErrBaseInfoObject()
	se := SpecErr{}.InitializeBaseInfo(parentHistory, baseInfo)
	se.SetStdError(xMsg, msgId)

	om := OpsMsgDto{}
	om.SetFromSpecErrMessage(se)

	if om.MsgType != msgType {
		t.Errorf("Expected Messgage Type == '%v'. Instead, Message Type == '%v'.", msgType, om.MsgType)
	}

	if om.IsError() != true {
		t.Errorf("Expected Standard Error Message to generate IsError='true'. It did NOT! IsError='false'.")
	}

	if om.IsFatalError() != false {
		t.Errorf("Expected Standard Error Message to generate IsFatalError()='false'. It did NOT! IsFatalError()='%v'", om.IsFatalError())
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

	if om.MsgContext.SourceFileName != "TSource06" {
		t.Errorf("Expected MsgContext.SourceFileName == 'TSource06'. Instead, SourceFileName== '%v'", om.MsgContext.SourceFileName)
	}

	if om.MsgContext.ParentObjectName != "PObj06" {
		t.Errorf("Expected MsgContext.ParentObjectName == 'PObj06'. Instead, ParentObjectName== '%v'", om.MsgContext.ParentObjectName)
	}

	if om.MsgContext.FuncName != "Func006" {
		t.Errorf("Expected MsgContext.FuncName == 'Func006'. Instead, FuncName== '%v'", om.MsgContext.FuncName)
	}

	if om.MsgContext.BaseMessageId != 6000 {
		t.Errorf("Expected MsgContext.BaseMessageId == '6000'. Instead, BaseMessageId== '%v'", om.MsgContext.BaseMessageId)
	}

}

func TestOpsMsgDto_SetFromSpecErrMessage_03(t *testing.T) {

	xMsg := "Information Message - Configured in SpecErr"
	msgId := int64(123)
	msgNo := int64(6123)
	msgType := OpsMsgTypeINFOMSG
	parentHistory := testOpsMsgDtoCreateSpecErrParentBaseInfo5Elements()
	baseInfo:= testOpsMsgDtoCreateSpecErrBaseInfoObject()
	se := SpecErr{}.InitializeBaseInfo(parentHistory, baseInfo)
	se.SetInfoMessage(xMsg, msgId)

	om := OpsMsgDto{}
	om.SetFromSpecErrMessage(se)

	if om.MsgType != msgType {
		t.Errorf("Expected Messgage Type == '%v'. Instead, Message Type == '%v'.", msgType, om.MsgType)
	}

	if om.IsError() != false {
		t.Errorf("Expected Information Message to generate IsError='false'. It did NOT! IsError='true'.")
	}

	if om.IsFatalError() != false {
		t.Errorf("Expected Information Message to generate IsFatalError()='false'. It did NOT! IsFatalError()='%v'", om.IsFatalError())
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

	if om.MsgContext.SourceFileName != "TSource06" {
		t.Errorf("Expected MsgContext.SourceFileName == 'TSource06'. Instead, SourceFileName== '%v'", om.MsgContext.SourceFileName)
	}

	if om.MsgContext.ParentObjectName != "PObj06" {
		t.Errorf("Expected MsgContext.ParentObjectName == 'PObj06'. Instead, ParentObjectName== '%v'", om.MsgContext.ParentObjectName)
	}

	if om.MsgContext.FuncName != "Func006" {
		t.Errorf("Expected MsgContext.FuncName == 'Func006'. Instead, FuncName== '%v'", om.MsgContext.FuncName)
	}

	if om.MsgContext.BaseMessageId != 6000 {
		t.Errorf("Expected MsgContext.BaseMessageId == '6000'. Instead, BaseMessageId== '%v'", om.MsgContext.BaseMessageId)
	}

}

func TestOpsMsgDto_SetInfoMessage_01(t *testing.T) {

	testParentHistory := testOpsMsgDtoCreateParentHistory()
	testMsgContext := testOpsMsgDtoCreateContextInfoObj()

	xMsg := "This is Information Message for test object"
	msgId := int64(19)
	msgNo := int64(6019)
	msgType := OpsMsgTypeINFOMSG

	om := testOpsMsgDtoCreateInfoMsg()

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
		t.Error("Expected Information Message to generate IsError='false'. It did NOT! IsError='true'.")
	}

	if om.IsFatalError() != false {
		t.Errorf("Expected Information to generate IsFatalError()='false'. It did NOT! IsFatalError()='%v'", om.IsFatalError())
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

func TestOpsMsgDto_SetInfoMessage_02(t *testing.T) {

	om := OpsMsgDto{}

	xMsg := "This is Information Message for test object"
	msgId := int64(19)
	msgNo := int64(6019)
	msgType := OpsMsgTypeINFOMSG

	om.SetMessageContext(testOpsMsgDtoCreateContextInfoObj())
	om.SetInfoMessage(xMsg, msgId)


	if om.MsgType != msgType {
		t.Errorf("Expected Messgage Type == '%v'. Instead, Message Type == '%v'.", msgType, om.MsgType)
	}

	if om.IsError() != false {
		t.Error("Expected Information Message to generate IsError='false'. It did NOT! IsError='true'.")
	}

	if om.IsFatalError() != false {
		t.Errorf("Expected Information to generate IsFatalError()='false'. It did NOT! IsFatalError()='%v'", om.IsFatalError())
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

func TestOpsMsgDto_SetInfoMessage_03(t *testing.T) {

	om := OpsMsgDto{}

	xMsg := "This is Information Message for test object"
	msgId := int64(19)
	msgNo := int64(19)
	msgType := OpsMsgTypeINFOMSG

	om.SetInfoMessage(xMsg, msgId)

	if om.MsgType != msgType {
		t.Errorf("Expected Messgage Type == '%v'. Instead, Message Type == '%v'.", msgType, om.MsgType)
	}

	if om.IsError() != false {
		t.Error("Expected Information Message to generate IsError='false'. It did NOT! IsError='true'.")
	}

	if om.IsFatalError() != false {
		t.Errorf("Expected Information to generate IsFatalError()='false'. It did NOT! IsFatalError()='%v'", om.IsFatalError())
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
