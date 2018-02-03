package common

import (
	"testing"
	"strings"
	"errors"
)


func TestOpsMsgDto_NewError_01(t *testing.T) {

	testParentHistory := testOpsMsgDtoCreateParentHistory()
	testMsgContext := testOpsMsgDtoCreateContextInfoObj()

	prefixMsgText := "Prefix Msg: "
	errMsgText := "This is FATAL Error Msg for test object"
	err := errors.New(errMsgText)
	xMsg := prefixMsgText + errMsgText
	msgId := int64(152)
	msgNo := int64(6152)
	msgType := OpsMsgTypeFATALERRORMSG

	om := OpsMsgDto{}.InitializeAllContextInfo(testParentHistory, testMsgContext).NewError(prefixMsgText, err, msgType , msgId)

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

func TestOpsMsgDto_NewError_02(t *testing.T) {

	testParentHistory := testOpsMsgDtoCreateParentHistory()
	testMsgContext := testOpsMsgDtoCreateContextInfoObj()

	prefixMsgText := "Prefix Msg: "
	errMsgText := "This is Standard Error Msg for test object"
	err := errors.New(errMsgText)
	xMsg := prefixMsgText + errMsgText
	msgId := int64(152)
	msgNo := int64(6152)
	msgType := OpsMsgTypeERRORMSG

	om := OpsMsgDto{}.InitializeAllContextInfo(testParentHistory, testMsgContext).NewError(prefixMsgText, err, msgType , msgId)

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

}

func TestOpsMsgDto_NewError_03(t *testing.T) {

	testMsgContext := testOpsMsgDtoCreateContextInfoObj()

	prefixMsgText := "Prefix Msg: "
	errMsgText := "This is Standard Error Msg for test object"
	err := errors.New(errMsgText)
	xMsg := prefixMsgText + errMsgText
	msgId := int64(152)
	msgNo := int64(6152)
	msgType := OpsMsgTypeERRORMSG

	om := OpsMsgDto{}.InitializeWithMessageContext(testMsgContext).NewError(prefixMsgText, err, msgType , msgId)

	if !testMsgContext.Equal(&om.MsgContext) {
		t.Error("Expected testMsgContext to EQUAL om.MsgContext. It did NOT!")
	}

	if om.MsgType != msgType {
		t.Errorf("Expected Messgage Type == '%v'. Instead, Message Type == '%v'.", msgType, om.MsgType)
	}

	if om.IsError() != true {
		t.Errorf("Expected Fatal Error Message to generate IsError='true'. It did NOT! IsError='false'.")
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

}

func TestOpsMsgDto_NewError_04(t *testing.T) {

	prefixMsgText := "Prefix Msg: "
	errMsgText := "This is Standard Error Msg for test object"
	err := errors.New(errMsgText)
	xMsg := prefixMsgText + errMsgText
	msgId := int64(152)
	msgNo := int64(152)
	msgType := OpsMsgTypeERRORMSG

	om := OpsMsgDto{}.NewError(prefixMsgText, err, msgType , msgId)


	if om.MsgType != msgType {
		t.Errorf("Expected Messgage Type == '%v'. Instead, Message Type == '%v'.", msgType, om.MsgType)
	}

	if om.IsError() != true {
		t.Errorf("Expected Fatal Error Message to generate IsError='true'. It did NOT! IsError='false'.")
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

}

func TestOpsMsgDto_NewFatalErrorMessage_01(t *testing.T) {

	testParentHistory := testOpsMsgDtoCreateParentHistory()
	testMsgContext := testOpsMsgDtoCreateContextInfoObj()

	xMsg := "This is FATAL Error Msg for test object"
	msgId := int64(152)
	msgNo := int64(6152)
	msgType := OpsMsgTypeFATALERRORMSG

	om := OpsMsgDto{}.InitializeAllContextInfo(testParentHistory, testMsgContext).NewFatalErrorMsg(xMsg, msgId)

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

func TestOpsMsgDto_NewFatalErrorMessage_02(t *testing.T) {

	testMsgContext := testOpsMsgDtoCreateContextInfoObj()

	xMsg := "This is FATAL Error Msg for test object"
	msgId := int64(152)
	msgNo := int64(6152)
	msgType := OpsMsgTypeFATALERRORMSG

	om := OpsMsgDto{}.InitializeWithMessageContext(testMsgContext).NewFatalErrorMsg(xMsg, msgId)

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

func TestOpsMsgDto_NewFatalErrorMessage_03(t *testing.T) {

	xMsg := "This is FATAL Error Msg for test object"
	msgId := int64(152)
	msgNo := int64(152)
	msgType := OpsMsgTypeFATALERRORMSG

	om := OpsMsgDto{}.NewFatalErrorMsg(xMsg, msgId)

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

func TestOpsMsgDto_NewInfoMsg_01(t *testing.T) {

	testParentHistory := testOpsMsgDtoCreateParentHistory()
	testMsgContext := testOpsMsgDtoCreateContextInfoObj()

	xMsg := "This is Information Message for test object"
	msgId := int64(19)
	msgNo := int64(6019)
	msgType := OpsMsgTypeINFOMSG

	om := OpsMsgDto{}.InitializeAllContextInfo(testParentHistory, testMsgContext).NewInfoMsg(xMsg, msgId)

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

func TestOpsMsgDto_NewInfoMsg_02(t *testing.T) {

	testMsgContext := testOpsMsgDtoCreateContextInfoObj()

	xMsg := "This is Information Message for test object"
	msgId := int64(19)
	msgNo := int64(6019)
	msgType := OpsMsgTypeINFOMSG

	om := OpsMsgDto{}.InitializeWithMessageContext(testMsgContext).NewInfoMsg(xMsg, msgId)

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

func TestOpsMsgDto_NewInfoMsg_03(t *testing.T) {

	xMsg := "This is Information Message for test object"
	msgId := int64(19)
	msgNo := int64(19)
	msgType := OpsMsgTypeINFOMSG

	om := OpsMsgDto{}.NewInfoMsg(xMsg, msgId)

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


func TestOpsMsgDto_NewNoErrorsNoMessagesMsg_01(t *testing.T) {

	testParentHistory := testOpsMsgDtoCreateParentHistory()
	testMsgContext := testOpsMsgDtoCreateContextInfoObj()

	xMsg := "No Errors and No Messages"
	msgId := int64(28)
	msgNo := int64(6028)
	msgType := OpsMsgTypeNOERRORNOMSGS

	om := OpsMsgDto{}.InitializeAllContextInfo(testParentHistory, testMsgContext).NewNoErrorsNoMessagesMsg("",msgId)

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
		t.Error("Expected No Errors-No Messages Message to generate IsError='false'. It did NOT! IsError='true'.")
	}

	if om.IsFatalError() != false {
		t.Errorf("Expected No Errors-No Messages Message to generate IsFatalError()='false'. It did NOT! IsFatalError()='%v'", om.IsFatalError())
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

func TestOpsMsgDto_NewNoErrorsNoMessagesMsg_02(t *testing.T) {

	testMsgContext := testOpsMsgDtoCreateContextInfoObj()

	xMsg := "No Errors and No Messages"
	msgId := int64(28)
	msgNo := int64(6028)
	msgType := OpsMsgTypeNOERRORNOMSGS

	om := OpsMsgDto{}.InitializeWithMessageContext(testMsgContext).NewNoErrorsNoMessagesMsg("",msgId)

	if !testMsgContext.Equal(&om.MsgContext) {
		t.Error("Expected testMsgContext to EQUAL om.MsgContext. It did NOT!")
	}

	if om.MsgType != msgType {
		t.Errorf("Expected Messgage Type == '%v'. Instead, Message Type == '%v'.", msgType, om.MsgType)
	}

	if om.IsError() != false {
		t.Error("Expected No Errors-No Messages Message to generate IsError='false'. It did NOT! IsError='true'.")
	}

	if om.IsFatalError() != false {
		t.Errorf("Expected No Errors-No Messages Message to generate IsFatalError()='false'. It did NOT! IsFatalError()='%v'", om.IsFatalError())
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

func TestOpsMsgDto_NewNoErrorsNoMessagesMsg_03(t *testing.T) {

	xMsg := "No Errors and No Messages"
	msgId := int64(28)
	msgNo := int64(28)
	msgType := OpsMsgTypeNOERRORNOMSGS

	om := OpsMsgDto{}.NewNoErrorsNoMessagesMsg("",msgId)

	if om.MsgType != msgType {
		t.Errorf("Expected Messgage Type == '%v'. Instead, Message Type == '%v'.", msgType, om.MsgType)
	}

	if om.IsError() != false {
		t.Error("Expected No Errors-No Messages Message to generate IsError='false'. It did NOT! IsError='true'.")
	}

	if om.IsFatalError() != false {
		t.Errorf("Expected No Errors-No Messages Message to generate IsFatalError()='false'. It did NOT! IsFatalError()='%v'", om.IsFatalError())
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

func TestOpsMsgDto_NewStdErrorMsg_01 (t *testing.T) {

	testParentHistory := testOpsMsgDtoCreateParentHistory()
	testMsgContext := testOpsMsgDtoCreateContextInfoObj()

	xMsg := "This is Standard Error Msg for test object"
	msgId := int64(429)
	msgNo := int64(6429)
	msgType := OpsMsgTypeERRORMSG


	om := OpsMsgDto{}.InitializeAllContextInfo(testParentHistory, testMsgContext).NewStdErrorMsg(xMsg, msgId)


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

func TestOpsMsgDto_NewStdErrorMsg_02 (t *testing.T) {

	testMsgContext := testOpsMsgDtoCreateContextInfoObj()

	xMsg := "This is Standard Error Msg for test object"
	msgId := int64(429)
	msgNo := int64(6429)
	msgType := OpsMsgTypeERRORMSG

	om := OpsMsgDto{}.InitializeWithMessageContext(testMsgContext).NewStdErrorMsg(xMsg, msgId)
	om.SetMessageContext(testMsgContext)

	if !testMsgContext.Equal(&om.MsgContext) {
		t.Error("Expected testMsgContext to EQUAL om.MsgContext. It did NOT!")
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

func TestOpsMsgDto_NewStdErrorMsg_03 (t *testing.T) {

	xMsg := "This is Standard Error Msg for test object"
	msgId := int64(429)
	msgNo := int64(429)
	msgType := OpsMsgTypeERRORMSG

	om := OpsMsgDto{}.NewStdErrorMsg(xMsg, msgId)


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



