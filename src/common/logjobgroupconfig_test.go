package common

import (
	"fmt"
	"testing"
	"time"
)

func TestLogJobGroupConfig_New(t *testing.T) {
	parms := StartupParameters{}
	thisSrcFileName := "logjobgroupconfig_test.go"
	thisMethodName := "TestLogJobGroupConfig_New"
	thisErrBlockNo := int64(80000)
	parms.StartTime = time.Now()
	parms.AppVersion = "2.0.0"
	parms.LogMode = LogVERBOSE
	parms.AppLogPath = "./CmdrX"
	parms.AppName = "CmdrX"
	parms.AppExeFileName = "CmdrX.exe"
	parms.NoOfJobs = 37
	parms.CommandFileName = "CmdrX.xml"

	//parent := ErrBaseInfo{}.GetNewParentInfo(thisSrcFileName, thisMethodName, thisErrBlockNo)
	msgCtx := OpsMsgContextInfo{
							SourceFileName: thisSrcFileName,
							ParentObjectName: "",
							FuncName: thisMethodName,
							BaseMessageId: thisErrBlockNo,
						}

	parentHistory := []OpsMsgContextInfo{msgCtx}

	lg := LogJobGroup{}

	om := lg.New(parms, parentHistory)

	if om.IsError() {
		t.Errorf("Expected om.IsError() == false, got %v", om.IsError())
	}

	if parms.CommandFileName != lg.CommandFileName {
		t.Error(fmt.Sprintf("Expected CommandFileName = %v, got ", parms.CommandFileName), lg.CommandFileName)
	}

	if parms.StartTime != lg.StartTime {
		t.Error(fmt.Sprintf("Expected Start Time = %v, got", parms.StartTime.String()), lg.StartTime.String())
	}

	if parms.AppVersion != lg.AppVersion {
		t.Error(fmt.Sprintf("Expected App Version = %v, got", parms.AppVersion), lg.AppVersion)
	}

	if parms.AppName != lg.AppName {
		t.Error(fmt.Sprintf("Expected App Name = %v, got", parms.AppName), lg.AppName)
	}

	if parms.LogMode != lg.LogMode {
		t.Error(fmt.Sprintf("Expected Log Mode = %v, got", parms.LogMode), lg.LogMode)
	}

	if parms.NoOfJobs != lg.AppNoOfJobs {
		t.Error(fmt.Sprintf("Expected Number Of Jobs = %v, got", parms.NoOfJobs), lg.AppNoOfJobs)
	}

}
