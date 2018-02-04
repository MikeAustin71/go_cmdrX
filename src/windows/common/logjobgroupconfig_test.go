package common

import (
	"testing"
	"time"
)

func TestLogJobGroupConfig_New(t *testing.T) {

	lg := LogJobGroup{}

	thisSrcFileName := "logjobgroupconfig_test.go"
	thisMethodName := "TestLogJobGroupConfig_New"
	thisErrBlockNo := int64(80000)
	lg.AppPathFileNameExt ,_ = FileMgr{}.New("../app/cmdrX.exe")
	lg.CmdPathFileNameExt, _ = FileMgr{}.New("../app/cmdrXCmds.xml")
	lg.AppStartTimeTzu, _ = TimeZoneUtility{}.New(time.Now().UTC(), "Local")
	lg.BatchStartTimeTzu, _ = TimeZoneUtility{}.New(lg.AppStartTimeTzu.TimeUTC, "Local")
	lg.AppVersion = "2.0.0"
	lg.LogMode = LogVERBOSE
	dt := DateTimeUtility{}
	dateTimeStamp := dt.GetDateTimeStr(lg.AppStartTimeTzu.TimeOut)
	logFileNameExt :=  lg.AppPathFileNameExt.FileName + "_" + dateTimeStamp + ".log"
	appErrFileNameExt := lg.AppPathFileNameExt.FileName + "_Errors_" + dateTimeStamp + ".txt"
	logPath, _ := DirMgr{}.New( "./cmdrXLog")
	lg.LogPathFileNameExt, _ = FileMgr{}.NewFromDirMgrFileNameExt(logPath, logFileNameExt)
	lg.AppErrPathFileNameExt, _ = FileMgr{}.NewFromDirMgrFileNameExt(logPath, appErrFileNameExt)
	lg.NoOfJobs = 37
	lg.BaseStartDir = lg.AppPathFileNameExt.DMgr.CopyOut()

	//parent := ErrBaseInfo{}.GetNewParentInfo(thisSrcFileName, thisMethodName, thisErrBlockNo)
	msgCtx := OpsMsgContextInfo{
							SourceFileName: thisSrcFileName,
							ParentObjectName: "",
							FuncName: thisMethodName,
							BaseMessageId: thisErrBlockNo,
						}

	parentHistory := []OpsMsgContextInfo{msgCtx}

	om := lg.New(parentHistory)

	if om.IsError() {
		t.Errorf("Expected om.IsError() == false, got %v. Error='%v'", om.IsError(), om.Error())
	}

}
