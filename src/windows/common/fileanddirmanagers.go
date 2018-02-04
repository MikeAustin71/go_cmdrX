package common

import (
	"fmt"
	fp "path/filepath"
	"errors"
	"os"
	"time"
	"strings"
	"io/ioutil"
)

/*
		This source code file contains types 'DirMgr' and
		'FileMgr'.

		The Source Repository for this source code file is :
			https://github.com/MikeAustin71/pathfilego.git

		Dependencies:
		-------------

		Types 'DirMgr' and 'FileMgr' depend on type, 'FileHellper'
		which is contained in source code file, 'filehelper.go' located
		in this directory.

		To use type 'DirMgr' or type 'FileMgr' you will need the following
		two source code files:

			1. pathfilego/003_filehelper/common/filehelper.go

			2. pathfilego/003_filehelper/common/fileanddirmanagers.go


 */


type DirMgrCollection struct {
	DirMgrs []DirMgr
}

// AddDirMgr - Adds a DirMgr object to the collection
func (dMgrs *DirMgrCollection) AddDirMgr(dMgr DirMgr) {
	dMgrs.DirMgrs = append(dMgrs.DirMgrs, dMgr.CopyOut())
}


func (dMgrs *DirMgrCollection) AddDirMgrByPathFile(pathFileName string) error {
	ePrefix := "DirMgrCollection.AddDirMgrByPathFile() "

	dMgr, err := DirMgr{}.New(pathFileName)

	if err!=nil {
		s:= ePrefix + "Error returned from DirMgr{}.New(pathFileName). pathFileName='%v' Error='%v'"
		return fmt.Errorf(s, pathFileName, err.Error())
	}

	dMgrs.DirMgrs = append(dMgrs.DirMgrs, dMgr)

	return nil
}

// AddFileInfo - Adds a File Manager object to the collection based on input from
// a directory path string and a os.FileInfo object.
func (dMgrs *DirMgrCollection) AddFileInfo(pathFile string, info os.FileInfo) error {

	ePrefix := "DirMgrCollection) AddFileInfo() "

	dMgr, err := DirMgr{}.NewFromFileInfo(pathFile, info)

	if err != nil {
		return fmt.Errorf(ePrefix + "Error retrned from DirMgr{}.NewFromFileInfo(pathFile, info). pathFile='%v' info.Name()='%v'  Error='%v'", pathFile, info.Name(), err.Error())
	}

	dMgrs.DirMgrs = append(dMgrs.DirMgrs, dMgr)

	return nil
}

// AddDirMgrCollection - Adds another collection of File Manager (DirMgr)
// objects to the current collection.
func (dMgrs *DirMgrCollection) AddDirMgrCollection(dMgrs2 *DirMgrCollection) {

	lOmc2 := len(dMgrs2.DirMgrs)

	if lOmc2 == 0 {
		return
	}

	for i:=0; i< lOmc2 ; i++ {
		dMgrs.AddDirMgr(dMgrs2.DirMgrs[i].CopyOut())
	}

	return
}

// CopyOut - Returns an DirMgrCollection which is an
// exact duplicate of the current DirMgrCollection
func (dMgrs *DirMgrCollection) CopyOut() (DirMgrCollection, error) {

	ePrefix := "DirMgrCollection.CopyOut() "

	dMgrs2 := DirMgrCollection{}

	lOmc := len(dMgrs.DirMgrs)

	if lOmc == 0 {
		return DirMgrCollection{}, errors.New(ePrefix + "Error: Empty DirMgrCollection. No messages available!")
	}

	for i:= 0; i < lOmc; i++ {
		dMgrs2.AddDirMgr(dMgrs.DirMgrs[i].CopyOut())
	}

	return dMgrs2, nil
}

// FindDirectories - searches through the DirMgrCollection to find DirMgr objects matching specified
// search criteria.
func (dMgrs *DirMgrCollection) FindDirectories(fileSelectionCriteria FileSelectionCriteria) (DirMgrCollection, error) {
	
	ePrefix := "DirMgrCollection.FindDirectories() "
	
	lDirCol := len(dMgrs.DirMgrs)

	if lDirCol == 0 {
		return DirMgrCollection{}, nil
	}

	fh := FileHelper{}

	var isMatchedFile bool
	var err error

	dMgrs2 := DirMgrCollection{}

	for i:=0; i < lDirCol; i++ {
		dMgr := dMgrs.DirMgrs[i]

		if dMgr.ActualDirFileInfo.IsFInfoInitialized {

			isMatchedFile, err = fh.FilterFileName(dMgr.ActualDirFileInfo, fileSelectionCriteria)

			if err!=nil {
				return 	DirMgrCollection{}, fmt.Errorf(ePrefix + "Error returned by fh.FilterFileName(dMgr.ActualDirFileInfo, fileSelectionCriteria) dMgr.ActualDirFileInfo.Name()='%v'  Error='%v'",dMgr.ActualDirFileInfo.Name(), err.Error())
			}
			
		} else {
			
			fip := FileInfoPlus{}
			
			fip.SetName(dMgr.DirectoryName)
			
			isMatchedFile, err = fh.FilterFileName(fip, fileSelectionCriteria)
			
			if err!=nil {
				s:= ePrefix + "Error returned by fh.FilterFileName(fip, fileSelectionCriteria) fip.Name()='%v'  Error='%v'"
				return 	DirMgrCollection{}, fmt.Errorf(s,fip.Name(), err.Error())
			}
			
		}

		if isMatchedFile && err == nil {
			dMgrs2.AddDirMgr(dMgr)
		} 

	}
	
	return dMgrs2, nil
}

// GetArrayLength - returns the array length of the
// DirMgrCollection File Managers (FMgrs) array.
func (dMgrs *DirMgrCollection) GetArrayLength() int {
	return len(dMgrs.DirMgrs)
}


// PopLastDirMgr - Removes the last File Manager (DirMgr) object
// from the collections array, and returns it to the calling method.
func (dMgrs *DirMgrCollection) PopLastDirMgr() (DirMgr, error) {

	ePrefix := "DirMgrCollection.PopLastDirMgr() "

	l1 := len(dMgrs.DirMgrs)

	if l1 == 0 {
		return DirMgr{}, errors.New(ePrefix + "Error: Empty DirMgrCollection. No messages available!")
	}

	fmgr := dMgrs.DirMgrs[l1-1].CopyOut()

	dMgrs.DirMgrs = dMgrs.DirMgrs[0:l1-1]

	return fmgr, nil
}

// PopFirstDirMgr - Removes the first OpsMsgDto object
// from the collections array, and returns it to
// the calling method.
func (dMgrs *DirMgrCollection) PopFirstDirMgr() (DirMgr, error) {

	ePrefix := "DirMgrCollection.PopFirstDirMgr() "

	l1 := len(dMgrs.DirMgrs)

	if l1 == 0 {
		return DirMgr{}, errors.New(ePrefix + "Error: Empty DirMgrCollection. No messages available!")
	}

	om := dMgrs.DirMgrs[0].CopyOut()

	dMgrs.DirMgrs = dMgrs.DirMgrs[1:l1]

	return om, nil
}

// PopDirMgrAtIndex - Returns a copy of the File Manager (DirMgr) object located
// at index, 'idx', in the DirMgrCollection array. As a 'Pop' method, the original
// DirMgr object is deleted from the DirMgrCollection array.
func (dMgrs *DirMgrCollection) PopDirMgrAtIndex(idx int) (DirMgr, error) {
	ePrefix := "DirMgrCollection.PopDirMgrAtIndex() "

	if idx < 0 {
		return DirMgr{}, fmt.Errorf(ePrefix + "Error: Input Parameter is less than zero. Index Out-Of-Range! idx='%v'", idx)
	}

	lDirMgrs := len(dMgrs.DirMgrs)

	if idx >= lDirMgrs {
		return DirMgr{}, fmt.Errorf(ePrefix + "Error: Input Parameter is greater than the length of the collection index. Index Out-Of-Range! idx='%v' Array Length='%v' ", idx, lDirMgrs)
	}

	if idx == 0 {
		return dMgrs.PopFirstDirMgr()
	}

	if idx == lDirMgrs- 1 {
		return dMgrs.PopLastDirMgr()
	}

	dMgr := dMgrs.DirMgrs[idx].CopyOut()

	dirCol2 := DirMgrCollection{}

	for i:=0; i < lDirMgrs; i++ {

		if i != idx {
			dirCol2.DirMgrs = append(dirCol2.DirMgrs, dMgrs.DirMgrs[i].CopyOut())
		}

	}

	dMgrs.DirMgrs = dirCol2.DirMgrs

	return dMgr, nil
}

// PeekFirstDirMgr - Returns the first element from the
// DirMgrCollection, but does NOT remove
// it from the OpsMessages array.
func (dMgrs *DirMgrCollection) PeekFirstDirMgr() (DirMgr, error) {

	ePrefix := "DirMgrCollection.PeekFirstDirMgr() "

	if len(dMgrs.DirMgrs) == 0 {
		return DirMgr{}, errors.New(ePrefix + "Error: Empty DirMgrCollection. No messages available!")
	}

	return dMgrs.DirMgrs[0].CopyOut(), nil
}

// PeekLastDirMgr - Returns the last element from the
// Operation Messages Collection, but does NOT remove
// it from the OpsMessages array.
func (dMgrs *DirMgrCollection) PeekLastDirMgr() (DirMgr, error) {

	ePrefix := "DirMgrCollection.PeekLastDirMgr()"

	l1 := len(dMgrs.DirMgrs)

	if l1==0 {
		return DirMgr{}, errors.New(ePrefix + "Error: Empty DirMgrCollection. No messages available!")
	}

	return dMgrs.DirMgrs[l1-1].CopyOut(), nil
}

// PeekDirMgrAtIndex - Returns a copy of the File Manager (DirMgr) object located
// at array index 'idx' in the DirMgrCollection. This is a 'Peek' method and the
// original DirMgr object is not deleted from the DirMgrCollection array.
func (dMgrs *DirMgrCollection) PeekDirMgrAtIndex(idx int) (DirMgr, error){

	ePrefix := "DirMgrCollection.PeekDirMgrAtIndex() "

	if idx < 0 {
		return DirMgr{}, fmt.Errorf(ePrefix + "Error: Input Parameter is less than zero. Index Out-Of-Range! idx='%v'", idx)
	}

	if idx >= len(dMgrs.DirMgrs) {
		return DirMgr{}, fmt.Errorf(ePrefix + "Error: Input Parameter is greater than the length of the collection index. Index Out-Of-Range! idx='%v' Array Length='%v' ", idx, len(dMgrs.DirMgrs))
	}

	return dMgrs.DirMgrs[idx].CopyOut(), nil
}



// DirMgr - This structure and associated methods
// are used to manage a specific directory.
type DirMgr struct {
	IsInitialized                 bool
	OriginalPath                  string
	Path                          string 	// Stored with no trailing path separator
	PathIsPopulated               bool
	PathDoesExist                 bool
	ParentPath                    string 	// Stored with no trailing path separator
	ParentPathIsPopulated         bool
	RelativePath                  string 	// Stored with no preceding path separator
	RelativePathIsPopulated       bool
	AbsolutePath                  string
	AbsolutePathIsPopulated       bool
	AbsolutePathDoesExist         bool
	AbsolutePathDifferentFromPath bool
	DirectoryName									string	// Name of directory with out parent path.
	VolumeName                    string
	VolumeIsPopulated             bool
	ActualDirFileInfo							FileInfoPlus
}

// CopyIn - Receives a pointer to a DirMgr object as an
// input parameter and copies the values from the incoming
// object to the current DirMgr object. When the copy operation
// is completed, the current DirMgr object is a duplicate of the
// incoming DirMgr object.
func (dMgr *DirMgr) CopyIn(dmgrIn *DirMgr) {

	dMgr.IsInitialized 									= dmgrIn.IsInitialized
	dMgr.OriginalPath										= dmgrIn.OriginalPath
	dMgr.Path														= dmgrIn.Path
	dMgr.PathIsPopulated								= dmgrIn.PathIsPopulated
	dMgr.PathDoesExist 									= dmgrIn.PathDoesExist
	dMgr.ParentPath							 				= dmgrIn.ParentPath
	dMgr.ParentPathIsPopulated					= dmgrIn.ParentPathIsPopulated
	dMgr.RelativePath						 				= dmgrIn.RelativePath
	dMgr.RelativePathIsPopulated				= dmgrIn.RelativePathIsPopulated
	dMgr.AbsolutePath										= dmgrIn.AbsolutePath
	dMgr.AbsolutePathIsPopulated				= dmgrIn.AbsolutePathIsPopulated
	dMgr.AbsolutePathDoesExist					= dmgrIn.AbsolutePathDoesExist
	dMgr.AbsolutePathDifferentFromPath	= dmgrIn.AbsolutePathDifferentFromPath
	dMgr.DirectoryName									= dmgrIn.DirectoryName
	dMgr.VolumeName											= dmgrIn.VolumeName
	dMgr.VolumeIsPopulated							= dmgrIn.VolumeIsPopulated
	dMgr.ActualDirFileInfo							= dmgrIn.ActualDirFileInfo.CopyOut()
}

// CopyOut - Makes a duplicate copy of the current DirMgr values and
// returns them in a new DirMgr object.
func (dMgr *DirMgr) CopyOut() DirMgr {

	dOut := DirMgr{}

	dOut.IsInitialized 									= dMgr.IsInitialized
	dOut.OriginalPath										= dMgr.OriginalPath
	dOut.Path														= dMgr.Path
	dOut.PathIsPopulated								= dMgr.PathIsPopulated
	dOut.PathDoesExist 									= dMgr.PathDoesExist
	dOut.ParentPath							 				= dMgr.ParentPath
	dOut.ParentPathIsPopulated					= dMgr.ParentPathIsPopulated
	dOut.RelativePath						 				= dMgr.RelativePath
	dOut.RelativePathIsPopulated				= dMgr.RelativePathIsPopulated
	dOut.AbsolutePath										= dMgr.AbsolutePath
	dOut.AbsolutePathIsPopulated				= dMgr.AbsolutePathIsPopulated
	dOut.AbsolutePathDoesExist					= dMgr.AbsolutePathDoesExist
	dOut.AbsolutePathDifferentFromPath	= dMgr.AbsolutePathDifferentFromPath
	dOut.DirectoryName									= dMgr.DirectoryName
	dOut.VolumeName											= dMgr.VolumeName
	dOut.VolumeIsPopulated							= dMgr.VolumeIsPopulated
	dOut.ActualDirFileInfo							= dMgr.ActualDirFileInfo.CopyOut()

	return dOut
}

// DeleteAll - BE CAREFUL!!! - This method will remove the directory identified by
// this DirMgr object. It will also delete all child directories and files in the
// directory tree.
//
// Example:
// Run DeleteAll on Directory: "../pathfilego/003_filehelper/testdestdir/destdir"
// All files and all subdirectories will be deleted.
// Only the parent path will remain: "../pathfilego/003_filehelper/testdestdir"
//
func (dMgr *DirMgr) DeleteAll() error {
	ePrefix := "DirMgr.DeleteAll() "
	var err error

	err = dMgr.IsDirMgrValid(ePrefix)

	if err != nil {
		return err
	}

	if dMgr.AbsolutePathDoesExist {

		err = os.RemoveAll(dMgr.AbsolutePath)

		if err != nil {
			return fmt.Errorf(ePrefix + "Error returned by os.RemoveAll(dMgr.AbsolutePath) returned error. dMgr.AbsolutePath='%v' Error='%v' ", dMgr.AbsolutePath, err.Error())
		}

		dMgr.DoesDirMgrPathExist()
		dMgr.DoesDirMgrAbsolutePathExist()

		return nil

	} else if dMgr.PathDoesExist {

		time.Sleep(time.Millisecond * 500)

		err = os.RemoveAll(dMgr.Path)

		if err != nil {
			return fmt.Errorf(ePrefix + "Error returned by os.RemoveAll(dMgr.AbsolutePath) returned error. dMgr.Path='%v' Error='%v' ", dMgr.Path, err.Error())
		}

		dMgr.AbsolutePathDoesExist = false
		dMgr.PathDoesExist = false

		return nil

	} else {
		dMgr.AbsolutePathDoesExist = false
		dMgr.PathDoesExist = false
		return nil
	}

}

func(dMgr *DirMgr) DoesDirectoryExist() (doesPathExist, doesAbsolutePathExist bool) {

	doesPathExist = dMgr.DoesDirMgrPathExist()

	doesAbsolutePathExist = dMgr.DoesDirMgrAbsolutePathExist()

	return
}

// DoesDirMgrAbsolutePathExist - Performs two operations.
// First the method determine whether the directory
// path indicated by the DirMgr.AbsolutePath field
// actually does exist on disk and returns a 'true'
// or 'false' boolean value accordingly. In addition,
// it also updates the DirMgr field
// 'DirMgr.AbsolutePathDoesExist'.
//
func (dMgr *DirMgr) DoesDirMgrAbsolutePathExist() bool {

	if dMgr.AbsolutePath == "" {
		dMgr.AbsolutePathDoesExist = false
		return false
	}

	info, err := os.Stat(dMgr.AbsolutePath)

	if err!=nil {
		dMgr.AbsolutePathDoesExist = false
	} else {
		dMgr.AbsolutePathDoesExist = true
		dMgr.ActualDirFileInfo = FileInfoPlus{}.NewPathFileInfo(dMgr.AbsolutePath, info)
	}

	return dMgr.AbsolutePathDoesExist

}

// DoesDirMgrPathExist - Performs two operations.
// First the method determine whether the directory
// path indicated by the DirMgr.Path field actually
// does exist on disk and returns a 'true' or 'false'
// boolean value accordingly. In addition it also
// updates the DirMgr field DirMgr.PathDoesExist field.
//
func (dMgr *DirMgr) DoesDirMgrPathExist() bool {

	if dMgr.Path == "" {
		dMgr.PathIsPopulated = false
		return false
	}

	info, err := os.Stat(dMgr.Path)

	if err!=nil {
		dMgr.PathDoesExist = false
	} else {
		dMgr.PathDoesExist = true
		dMgr.ActualDirFileInfo = FileInfoPlus{}.NewPathFileInfo(dMgr.Path, info)
	}

	return dMgr.PathDoesExist

}

// Empty - Returns all DirMgr field values to their uninitialized
// or original zero values.
func (dMgr *DirMgr) Empty() {

	dMgr.IsInitialized 									= false
	dMgr.OriginalPath										= ""
	dMgr.Path														= ""
	dMgr.PathIsPopulated								= false
	dMgr.PathDoesExist 									= false
	dMgr.ParentPath							 				= ""
	dMgr.ParentPathIsPopulated					= false
	dMgr.RelativePath						 				= ""
	dMgr.RelativePathIsPopulated				= false
	dMgr.AbsolutePath										= ""
	dMgr.AbsolutePathIsPopulated				= false
	dMgr.AbsolutePathDoesExist					= false
	dMgr.AbsolutePathDifferentFromPath	= false
	dMgr.DirectoryName									= ""
	dMgr.VolumeName											= ""
	dMgr.VolumeIsPopulated							= false
	dMgr.ActualDirFileInfo							= FileInfoPlus{}

}

// Equal - Compares two DirMgr objects to determine if
// they are equal.
func (dMgr *DirMgr) Equal(dmgr2 *DirMgr) bool {

	if dMgr.IsInitialized != dmgr2.IsInitialized ||
		dMgr.OriginalPath										!= dmgr2.OriginalPath 									||
		dMgr.Path														!= dmgr2.Path														||
		dMgr.PathIsPopulated								!= dmgr2.PathIsPopulated								||
		dMgr.PathDoesExist != dmgr2.PathDoesExist ||
		dMgr.ParentPath							 				!= dmgr2.ParentPath											||
		dMgr.ParentPathIsPopulated					!= dmgr2.ParentPathIsPopulated					||
		dMgr.RelativePath						 				!= dmgr2.RelativePath										||
		dMgr.RelativePathIsPopulated				!= dmgr2.RelativePathIsPopulated				||
		dMgr.AbsolutePath										!= dmgr2.AbsolutePath										||
		dMgr.AbsolutePathIsPopulated				!= dmgr2.AbsolutePathIsPopulated				||
		dMgr.AbsolutePathDoesExist					!= dmgr2.AbsolutePathDoesExist					||
		dMgr.AbsolutePathDifferentFromPath	!= dmgr2.AbsolutePathDifferentFromPath	||
		dMgr.DirectoryName									!= dmgr2.DirectoryName									||
		dMgr.VolumeName											!= dmgr2.VolumeName											||
		dMgr.VolumeIsPopulated							!= dmgr2.VolumeIsPopulated	{

		return false
	}

	if !dMgr.ActualDirFileInfo.Equal(&dmgr2.ActualDirFileInfo) {
		return false
	}

	return true
}


// EqualPaths - Compares two DirMgr objects to determine
// if their paths are equal.
func (dMgr *DirMgr) EqualPaths(dMgr2 *DirMgr) bool {

	if dMgr.IsInitialized != dMgr2.IsInitialized {
		return false
	}

	if dMgr.AbsolutePath == dMgr2.AbsolutePath &&
			dMgr.Path == dMgr2.Path {
		return true
	}

	return false
}

// DeleteWalkDirFiles - !!! BE CAREFUL !!! This method
// deletes files in a specified directory tree.
//
// This method searches for files residing in the directory tree
// identified by the current DirMgr object. The method 'walks the
// directory tree' locating all files in the directory tree which
// match the file selection criteria submitted as method input parameter,
// 'deleteFileSelectionCriteria'.
//
// If a file matches the File Selection Criteria, it is DELETED. By the way,
// if ALL the file selection criterion are set to zero values or 'Inactive',
// then ALL FILES IN THE DIRECTORY ARE DELETED!!!
//
// A record of file deletions is included in the returned DirectoryDeleteFileInfo
// structure (DirectoryDeleteFileInfo.DeletedFiles).
//
// Input Parameter:
// ================
//
// deleteFileSelectionCriteria FileSelectionCriteria
//			This input parameter should be configured with the desired file
//      selection criteria. Files matching this criteria will be deleted.
//
//			type FileSelectionCriteria struct {
//					FileNamePatterns						[]string
//					FilesOlderThan							time.Time
//					FilesNewerThan							time.Time
//					SelectByFileMode						os.FileMode
//					SelectCriterionMode					FileSelectCriterionMode // Default is AND Criterion Together
//				}
//
//			The FileSelectionCriteria type allows for configuration of
//			single or multiple file selection criterion. The 'SelectCriterionMode'
//			can be used to specify whether the file must match all
//			or any one of the active file selection criterion.
//
//			Elements of the FileSelectionCriteria are described below:
//
// 			FileNamePatterns []string		- An array of strings which may define
//																		one or more search patterns. If a file
//																		matches any one of the search pattern
//																		strings, it is deemed to be a 'match'
//																		for the search pattern criterion.
//																		Example Patterns:
//																				"*.log"
//																				"current*.txt"
//
//														  			If this string array has zero length or if
//																		all the strings are empty strings, then this
//																		file search criterion is considered 'Inactive'
//																		or 'Not Set'.
//
//
//        FilesOlderThan	time.Time	- This date time type is compared to file
//																		modification date times in order to determine
//																		whether the file is older than the 'FilesOlderThan'
//																		file selection criterion. If the file is older than
//																		the 'FilesOlderThan' date time, that file is considered
// 																		a 'match'	for this file selection criterion.
//
//																	  If the value of 'FilesOlderThan' is set to time zero,
//																		the default value for type time.Time{}, then this
//																		file selection criterion is considered to be 'Inactive'
//																		or 'Not Set'.
//
//        FilesNewerThan	time.Time	- This date time type is compared to the file
//																		modification date time in order to determine
//																		whether the file is newer than the 'FilesNewerThan'
//																		file selection criterion. If the file modification date time
// 																		is newer than the 'FilesNewerThan' date time, that file is
// 																		considered a 'match' for this file selection criterion.
//
//																	  If the value of 'FilesNewerThan' is set to time zero,
//																		the default value for type time.Time{}, then this
//																		file selection criterion is considered to be 'Inactive'
//																		or 'Not Set'.
//
// 		 SelectByFileMode os.FileMode - os.FileMode is a uint32 value. This file selection criterion
// 																		allows for the selection of files by File Mode. File modes
// 																		are compared to the value	of 'SelectByFileMode'. If the File
// 																		Mode for a given file is equal to the value of 'SelectByFileMode',
//																		that file is considered to be a 'match' for this file selection
// 																		criterion.
//
//																		If the value of 'SelectByFileMode' is set equal to zero, then
//																		this file selection criterion is considered 'Inactive' or
//																		'Not Set'.
//
//	SelectCriterionMode	FileSelectCriterionMode -
//																		This parameter selects the manner in which the file selection
//																		criteria above are applied in determining a 'match' for file
// 																		selection purposes. 'SelectCriterionMode' may be set to one of
//																		two constant values:
//
//																		ANDFILESELECTCRITERION	- File selected if all active selection criteria
//																			are satisfied.
//
// 																			If this constant value is specified for the file selection mode,
// 																			then	a given file will not be judged as 'selected' unless all of
// 																			the active selection criterion are satisfied. In other words, if
// 																			three active search criterion are provided for 'FileNamePatterns',
//																			'FilesOlderThan' and 'FilesNewerThan', then a file will NOT be
//																			selected unless it has satisfied all three criterion in this example.
//
//																		ORFILESELECTCRITERION 	- File selected if any active selection criterion
//																			is satisfied.
//
// 																			If this constant value is specified for the file selection mode,
// 																			then a given file will be selected if any one of the active file
// 																			selection criterion is satisfied. In other words, if three active
// 																			search criterion are provided for 'FileNamePatterns', 'FilesOlderThan'
// 																			and 'FilesNewerThan', then a file will be selected if it satisfies any
// 																			one of the three criterion in this example.
//
// IMPORTANT
// *********
// If all of the file selection criterion in the FileSelectionCriteria object are
// 'Inactive' or 'Not Set' (set to their zero or default values), then all of
// the files processed will be selected and DELETED.
//
// 			Example:
//					FileNamePatterns 	= ZERO Length Array
//          filesOlderThan 		= time.Time{}
//					filesNewerThan 		= time.Time{}
//					SelectByFileMode 	= uint32(0)
//
//					In this example, all of the selection criterion are
//					'Inactive' and therefore all of the files encountered
//					will be SELECTED FOR DELETION!
//
//
// Return Value:
// =============
//
// 				type DirectoryDeleteFileInfo struct {
//									StartPath            	string
//									DirMgrs          	[]DirMgr
//									FoundFiles           	[]FileWalkInfo
//									ErrReturns           	[]string
//									DeleteFileSelectCriteria FileSelectionCriteria
//									DeletedFiles         	[]FileWalkInfo
//					}
//
//					If successful, files matching the file selection criteria
//					specified in input parameter 'deleteFileSelectionCriteria'
//  				will be DELETED and returned in a 'DirectoryDeleteFileInfo'
//  				structure field, 'DirectoryDeleteFileInfo.DeletedFiles.'
//
//					Note: It is a good idea to check the returned field
// 								DirectoryDeleteFileInfo.ErrReturns to determine if any
// 								system errors were encountered during file processing.
//
//	error	- If a program execution error is encountered during processing, it will
//					returned as an 'error' type. Also, see the comment on
// 					DirectoryDeleteFileInfo.ErrReturns, above.
//
func (dMgr *DirMgr) DeleteWalkDirFiles(deleteFileSelectionCriteria FileSelectionCriteria) (DirectoryDeleteFileInfo, error) {

	ePrefix := "DirMgr.DeleteWalkDirFiles() "
	deleteFilesInfo := DirectoryDeleteFileInfo{}

	err := dMgr.IsDirMgrValid(ePrefix)

	if err!= nil {
		return deleteFilesInfo, err
	}

	if dMgr.AbsolutePathDoesExist {

		deleteFilesInfo.StartPath = dMgr.AbsolutePath

	} else if dMgr.DoesDirMgrPathExist() {

		deleteFilesInfo.StartPath = dMgr.Path

	} else {

		return deleteFilesInfo, fmt.Errorf(ePrefix + "Path and AbsolutePath - PATH DOES NOT EXIST! dMgr.AbsolutePath='%v' dMgr.Path='%v'", dMgr.AbsolutePath, dMgr.Path)
	}

	deleteFilesInfo.DeleteFileSelectCriteria = deleteFileSelectionCriteria

	fh := FileHelper{}

	err =  fp.Walk(deleteFilesInfo.StartPath, fh.makeFileHelperWalkDirDeleteFilesFunc(&deleteFilesInfo))

	if err != nil {
		return deleteFilesInfo, fmt.Errorf(ePrefix+"Error returned by  FileHelper.makeFileHelperWalkDirDeleteFilesFunc(&dWalkInfo). dWalkInfo.StartPath='%v' Error='%v' ",deleteFilesInfo.StartPath, err.Error())
	}

	return deleteFilesInfo, nil

}

// FindWalkDirFiles - This method returns file information on files residing in a specific
// directory tree identified by the current DirMgr object.
//
// This method 'walks the directory tree' locating all files in the directory tree which match
// the file selection criteria submitted as input parameter, 'fileSelectCriteria'.
//
// If a file matches the File Selection Criteria, it is included in the returned field,
// 'DirectoryTreeInfo.FoundFiles'. By the way, if ALL the file selection criterion are set to
// zero values or 'Inactive', then ALL FILES in the directory are selected and returned in the field,
// 'DirectoryTreeInfo.FoundFiles'.
//
// Input Parameter:
// ================
//
// fileSelectCriteria FileSelectionCriteria
//			This input parameter should be configured with the desired file
//      selection criteria. Files matching this criteria will be returned as
// 			'Found Files'.
//
//			type FileSelectionCriteria struct {
//					FileNamePatterns						[]string		// An array of strings containing File Name Patterns
//					FilesOlderThan							time.Time		// Match files with older modification date times
//					FilesNewerThan							time.Time		// Match files with newer modification date times
//					SelectByFileMode						os.FileMode	// Match file mode. Zero if inactive
//					SelectCriterionMode					FileSelectCriterionMode // Specifies 'AND' or 'OR' selection mode
//				}
//
//			The FileSelectionCriteria type allows for configuration of single or multiple file
// 			selection criterion. The 'SelectCriterionMode' can be used to specify whether the
// 			file must match all, or any one, of the active file selection criterion.
//
//			Elements of the FileSelectionCriteria are described below:
//
// 			FileNamePatterns []string		- An array of strings which may define one or more
//																		search patterns. If a file name matches any one of the
// 																		search pattern strings, it is deemed to be a 'match'
//																		for the search pattern criterion.
//																		Example Patterns:
//																				"*.log"
//																				"current*.txt"
//
//														  			If this string array has zero length or if
//																		all the strings are empty strings, then this
//																		file search criterion is considered 'Inactive'
//																		or 'Not Set'.
//
//
//        FilesOlderThan	time.Time	- This date time type is compared to file
//																		modification date times in order to determine
//																		whether the file is older than the 'FilesOlderThan'
//																		file selection criterion. If the file modification
// 																		date time is older than the 'FilesOlderThan' date time,
// 																		that file is considered a 'match'	for this file selection
// 																		criterion.
//
//																	  If the value of 'FilesOlderThan' is set to time zero,
//																		the default value for type time.Time{}, then this
//																		file selection criterion is considered to be 'Inactive'
//																		or 'Not Set'.
//
//        FilesNewerThan	time.Time	- This date time type is compared to the file
//																		modification date time in order to determine
//																		whether the file is newer than the 'FilesNewerThan'
//																		file selection criterion. If the file modification date time
// 																		is newer than the 'FilesNewerThan' date time, that file is
// 																		considered a 'match' for this file selection criterion.
//
//																	  If the value of 'FilesNewerThan' is set to time zero,
//																		the default value for type time.Time{}, then this
//																		file selection criterion is considered to be 'Inactive'
//																		or 'Not Set'.
//
// 		 SelectByFileMode os.FileMode - os.FileMode is an uint32 value. This file selection criterion
// 																		allows for the selection of files by File Mode. File Modes
// 																		are compared to the value	of 'SelectByFileMode'. If the File
// 																		Mode for a given file is equal to the value of 'SelectByFileMode',
//																		that file is considered to be a 'match' for this file selection
// 																		criterion.
//
//																		If the value of 'SelectByFileMode' is set equal to zero, then
//																		this file selection criterion is considered 'Inactive' or
//																		'Not Set'.
//
//	SelectCriterionMode	FileSelectCriterionMode -
//																		This parameter selects the manner in which the file selection
//																		criteria above are applied in determining a 'match' for file
// 																		selection purposes. 'SelectCriterionMode' may be set to one of
//																		two constant values:
//
//																		ANDFILESELECTCRITERION	- File selected if all active selection criteria
//																			are satisfied.
//
// 																			If this constant value is specified for the file selection mode,
// 																			then a given file will not be judged as 'selected' unless all of
// 																			the active selection criterion are satisfied. In other words, if
// 																			three active search criterion are provided for 'FileNamePatterns',
//																			'FilesOlderThan' and 'FilesNewerThan', then a file will NOT be
//																			selected unless it has satisfied all three criterion in this example.
//
//																		ORFILESELECTCRITERION 	- File selected if any active selection criterion
//																			is satisfied.
//
// 																			If this constant value is specified for the file selection mode,
// 																			then a given file will be selected if any one of the active file
// 																			selection criterion is satisfied. In other words, if three active
// 																			search criterion are provided for 'FileNamePatterns', 'FilesOlderThan'
// 																			and 'FilesNewerThan', then a file will be selected if it satisfies any
// 																			one of the three criterion in this example.
//
// IMPORTANT
// *********
// If all of the file selection criterion in the FileSelectionCriteria object are
// 'Inactive' or 'Not Set' (set to their zero or default values), then all of
// the files processed in the directory tree will be selected and returned as
// 'Found Files'.
//
// 			Example:
//					FileNamePatterns 	= ZERO Length Array
//          filesOlderThan 		= time.Time{}
//					filesNewerThan 		= time.Time{}
//					SelectByFileMode 	= uint32(0)
//
//					In this example, all of the selection criterion are
//					'Inactive' and therefore all of the files encountered
//					in the target directory will be selected and returned
//					as 'Found Files'.
//
//
// Return Value:
// =============
//
//	DirectoryTreeInfo structure	-
//					type DirectoryTreeInfo struct {
//						StartPath            	string								// The starting path or directory for the file search
//						DirMgrs          	[]DirMgr							// DirMgrs found during directory tree search
//						FoundFiles           	[]FileWalkInfo				// Found Files matching file selection criteria
//						ErrReturns           	[]string							// Internal System errors encountered
//						FileSelectCriteria    FileSelectionCriteria // The File Selection Criteria submitted as an
// 																												// input parameter to this method.
//					}
//
//					If successful, files matching the file selection criteria input
//  				parameter shown above will be returned in a 'DirectoryTreeInfo'
//  				object. The field 'DirectoryTreeInfo.FoundFiles' contains information
// 					on all the files in the specified directory tree which match the file selection
// 					criteria.
//
//					Note: It is a good idea to check the returned field 'DirectoryTreeInfo.ErrReturns'
// 								to determine if any internal system errors were encountered while processing
// 								the directory tree.
//
//	error	- If a program execution error is encountered during processing, it will
//					be returned as an 'error' type. Also, see the comment on
// 					'DirectoryTreeInfo.ErrReturns', above.
//
func (dMgr *DirMgr) FindWalkDirFiles(fileSelectCriteria FileSelectionCriteria) (DirectoryTreeInfo, error) {

	ePrefix := "DirMgr.GetWalkDirInfo() "
	findFilesInfo := DirectoryTreeInfo{}

	err := dMgr.IsDirMgrValid(ePrefix)

	if err!= nil {
		return findFilesInfo, err
	}

	if dMgr.AbsolutePathDoesExist {

		findFilesInfo.StartPath = dMgr.AbsolutePath

	} else if dMgr.DoesDirMgrPathExist() {

		findFilesInfo.StartPath = dMgr.Path

	} else {

		return findFilesInfo, fmt.Errorf(ePrefix + "Path and AbsolutePath - PATH DOES NOT EXIST! dMgr.AbsolutePath='%v' dMgr.Path='%v'", dMgr.AbsolutePath, dMgr.Path)
	}

	findFilesInfo.FileSelectCriteria = fileSelectCriteria

	fh := FileHelper{}

	err =  fp.Walk(findFilesInfo.StartPath, fh.makeFileHelperWalkDirFindFilesFunc(&findFilesInfo))

	if err != nil {
		return findFilesInfo, fmt.Errorf(ePrefix+"Error returned by  FileHelper.FindFilesWalkDirectory(&dWalkInfo). dWalkInfo.StartPath='%v' Error='%v' ",findFilesInfo.StartPath, err.Error())
	}

	return findFilesInfo, nil
}

// GetAbsolutePathWithSeparator - Returns the current
// DirMgr.AbsolutePath with a trailing os.PathSeparator
// character.
func (dMgr *DirMgr) GetAbsolutePathWithSeparator() string {
	lPath := len(dMgr.AbsolutePath)

	if lPath == 0 {
		return ""
	}

	if dMgr.AbsolutePath[lPath-1] != os.PathSeparator {
		return dMgr.AbsolutePath + string(os.PathSeparator)
	}

	return dMgr.AbsolutePath
}

// GetPathWithSeparator - Returns the current
// DirMgr.AbsolutePath with a trailing os.PathSeparator
// character.
func (dMgr *DirMgr) GetPathWithSeparator() string {
	lPath := len(dMgr.Path)

	if lPath == 0 {
		return ""
	}

	if dMgr.Path[lPath-1] != os.PathSeparator {
		return dMgr.Path + string(os.PathSeparator)
	}

	return dMgr.Path
}

// IsDirMgrValid - This method examines the current DirMgr object
// to determine whether it has been properly configured.
// If the current DirMgr object is valid, the method returns
// 'nil' for no errors.
//
// Otherwise, if the DirMgr object is INVALID, an error is
// returned.
func (dMgr *DirMgr) IsDirMgrValid(errPrefixStr string) error {

	ePrefix := strings.TrimRight(errPrefixStr, " ") + " DirMgr.IsDirMgrValid() "

	if !dMgr.IsInitialized {
		return fmt.Errorf(ePrefix+"Error: DirMgr is NOT Initialized.")
	}

	if dMgr.AbsolutePath == "" {
		return fmt.Errorf(ePrefix+"Error: DirMgr.AbsolutePath is EMPTY!.")
	}

	dMgr.AbsolutePathIsPopulated = true

	if dMgr.Path == "" {
		return fmt.Errorf(ePrefix+"Error: DirMgr.AbsolutePath is EMPTY!.")
	}

	dMgr.PathIsPopulated = true

	dMgr.DoesDirMgrAbsolutePathExist()
	dMgr.DoesDirMgrPathExist()

	return nil
}

// MakeDir - If the directory path identified
// by the current DirMgr object does not exist,
// this method will create that directory path.
func (dMgr *DirMgr) MakeDir() error {

	ePrefix := "DirMgr.MakeDir() "
	var err error

	err = dMgr.IsDirMgrValid(ePrefix)

	if err != nil {
		return nil
	}

	var ModePerm os.FileMode = 0777

	if dMgr.AbsolutePathIsPopulated {

		if dMgr.AbsolutePathDoesExist {
			// No need to create directory, it already
			// exists.
			return nil
		}

		err = os.MkdirAll(dMgr.AbsolutePath, ModePerm)

		if err != nil {
			return fmt.Errorf(ePrefix + "Error returned from os.MkdirAll(dMgr.AbsolutePath, ModePerm) dMgr.AbsolutePath='%v' ModePerm='%v'  Error='%v'", dMgr.AbsolutePath, ModePerm, err.Error())
		}

		dMgr.DoesDirMgrPathExist()
		dMgr.DoesDirMgrAbsolutePathExist()

		// No errors - directory created.
		return nil
	}

	// dMgr.PathIsPopulated MUST equal 'true'

	if dMgr.PathDoesExist {
		// No need to create directory, it already
		// exists.
		return nil
	}

	err = os.MkdirAll(dMgr.Path, ModePerm)

	if err != nil {
		return fmt.Errorf(ePrefix + "Error returned from os.MkdirAll(dMgr.Path, ModePerm) dMgr.Path='%v' ModePerm='%v'  Error='%v'", dMgr.Path, ModePerm, err.Error())
	}

	dMgr.DoesDirMgrPathExist()
	dMgr.DoesDirMgrAbsolutePathExist()

	// No errors - directory created.
	return nil
}

// New - Returns a new DirMgr object and populates the
// the data fields.
//
// Input Parameters:
// =================
//
// pathStr string 	- A path string designating a path or directory.
// 										To reduce errors, the 'pathStr' should be terminated
//										with an appropriate path separator ('/' or '\')
//										Example 'pathStr': "C:\dirA\dirB\dirC\"
//
// Example Output After DirMgr Configuration:
//
//     ----------------------------
//     DirMgr Fields
//     ----------------------------
//                
//		IsInitialized:  true
//		Original Path:  D:\go\work\src\MikeAustin71\pathfilego\003_filehelper\logTest\testoverwrite
//		Path:  D:\go\work\src\MikeAustin71\pathfilego\003_filehelper\logTest\testoverwrite
//		PathIsPopuslated:  true
//		PathDoesExist:  true
//		ParentPath:  D:\go\work\src\MikeAustin71\pathfilego\003_filehelper\logTest
//		ParentPathIsPopulated:  true
//		RelativePath:  testoverwrite
//		RelativePathIsPopulated:  true
//		AbsolutePath:  D:\go\work\src\MikeAustin71\pathfilego\003_filehelper\logTest\testoverwrite
//		AbsolutePathIsPopulated:  true
//		AbsolutePathDifferentFromPath:  false
//		AbsolutePathDoesExist:  true
//		Directory Name:  testoverwrite
//		VolumeName:  D:
//		VolumeIsPopulated:  true
//		========== File Info Data ==========
//		File Info IsDir():  true
//		File Info Name():  testoverwrite
//		File Info Size():  0
//		File Info ModTime():  2018-01-06 Sat 00:06:56.421236800 -0600 CST
//		File Info Mode():  drwxrwxrwx
//		File Info     Sys():  &{16 {617269082 30594119} {2388100752 30639796} {2388100752 30639796} 0 0}
//		Dir Path:  D:\go\work\src\MikeAustin71\pathfilego\003_filehelper\logTest\testoverwrite
//
func (dMgr DirMgr) New(pathStr string) (DirMgr, error) {

	ePrefix := "DirMgr.New() "

	if len(pathStr) == 0 {
		return DirMgr{}, errors.New(ePrefix + "Error: Input parameter 'pathStr' is Zero Length!")
	}


	newDirMgr := DirMgr{}

	_, err := newDirMgr.SetDirMgr(pathStr)

	if err != nil {
		return DirMgr{}, fmt.Errorf(ePrefix + "Error returned by newDirMgr.SetDirMgr(pathStr) pathStr='%v' Error='%v'", pathStr, err.Error())
	}

	return newDirMgr, nil
}


// NewFromFileInfo - Returns a new DirMgr object based on two input parameters:
// 		- A directory path string
//		- An os.FileInfo object
func (dMgr DirMgr) NewFromFileInfo(pathStr string, info os.FileInfo) (DirMgr, error) {

	ePrefix := "DirMgr) NewFromFileInfo() "

	if len(pathStr) == 0 {
		return DirMgr{}, errors.New(ePrefix + "Error: Input parameter 'pathStr' is Zero Length!")
	}

	newDirMgr := DirMgr{}

	err := newDirMgr.SetDirMgrWithFileInfo(pathStr, info)

	if err !=nil {
		fmt.Errorf(ePrefix + "Error returned from ")
	}

	return newDirMgr, nil
}

// SetDirMgr - Sets the DirMgr fields and path strings for the current DirMgr object.
//
// Input Parameters:
// =================
//
// pathStr string 	- A path string designating a path or directory.
// 										To reduce errors, the 'pathStr' should be terminated
//										with an appropriate path separator ('/' or '\')
//										Example 'pathStr': "C:\dirA\dirB\dirC\"
//
// Example Output After DirMgr Configuration:
//
//     ----------------------------
//     DirMgr Fields
//     ----------------------------
//
//		IsInitialized:  true
//		Original Path:  D:\go\work\src\MikeAustin71\pathfilego\003_filehelper\logTest\testoverwrite
//		Path:  D:\go\work\src\MikeAustin71\pathfilego\003_filehelper\logTest\testoverwrite
//		PathIsPopuslated:  true
//		PathDoesExist:  true
//		ParentPath:  D:\go\work\src\MikeAustin71\pathfilego\003_filehelper\logTest
//		ParentPathIsPopulated:  true
//		RelativePath:  testoverwrite
//		RelativePathIsPopulated:  true
//		AbsolutePath:  D:\go\work\src\MikeAustin71\pathfilego\003_filehelper\logTest\testoverwrite
//		AbsolutePathIsPopulated:  true
//		AbsolutePathDifferentFromPath:  false
//		AbsolutePathDoesExist:  true
//		Directory Name:  testoverwrite
//		VolumeName:  D:
//		VolumeIsPopulated:  true
//		========== File Info Data ==========
//		File Info IsDir():  true
//		File Info Name():  testoverwrite
//		File Info Size():  0
//		File Info ModTime():  2018-01-06 Sat 00:06:56.421236800 -0600 CST
//		File Info Mode():  drwxrwxrwx
//		File Info     Sys():  &{16 {617269082 30594119} {2388100752 30639796} {2388100752 30639796} 0 0}
//		Dir Path:  D:\go\work\src\MikeAustin71\pathfilego\003_filehelper\logTest\testoverwrite
//
func (dMgr *DirMgr) SetDirMgr(pathStr string) (isEmpty bool, err error) {
	ePrefix := "DirMgr.SetDirMgr() "

	dMgr.Empty()

	fh := FileHelper{}

	err = nil
	isEmpty = true

	if len(pathStr) == 0 {
		isEmpty = true
		err = errors.New(ePrefix + "Error: Input parameter 'pathStr' is a Zero length string!")
		return
	}

	trimmedPathStr := strings.TrimLeft(strings.TrimRight(pathStr, " "), " ")

	if len(trimmedPathStr) == 0 {
		err = errors.New(ePrefix + "Error: Trimmed Input Parameter 'pathStr' is a Zero length string!")
		isEmpty = true
		return
	}

	adjustedTrimmedPathStr := fh.AdjustPathSlash(trimmedPathStr)

	finalPathStr, isEmptyPath, err2 := fh.GetPathFromPathFileName(adjustedTrimmedPathStr)

	if err2 != nil {
		err = fmt.Errorf(ePrefix + "Error: INVALID PATH. fh.GetPathFromPathFileName(pathStr) pathStr='%v'  Error='%v'", pathStr, err2.Error())
		isEmpty = isEmptyPath
		return
	}

	if isEmptyPath {
		isEmpty = true
		err = fmt.Errorf(ePrefix + "Error: INVALID PATH. 'pathStr' generated an Empty Path! pathStr='%v' ", pathStr)
		return
	}

	if len(finalPathStr) == 0 {
		err = fmt.Errorf(ePrefix + "Error: Path returned from fh.GetPathFromPathFileName(pathStr) is EMPTY! pathStr='%v'", pathStr)
		isEmpty = true
		return
	}

	dMgr.OriginalPath = adjustedTrimmedPathStr

	dMgr.Path = finalPathStr

	dMgr.PathIsPopulated = true
	dMgr.DoesDirMgrPathExist()

	if dMgr.Path == fp.VolumeName(dMgr.Path) {

		dMgr.AbsolutePath = dMgr.Path

	} else {

		dMgr.AbsolutePath, err2 = fh.MakeAbsolutePath(dMgr.Path)

		if err2 != nil {
			dMgr.Empty()
			err = fmt.Errorf(ePrefix + "- fh.MakeAbsolutePath(dMgr.Path) returned error. dMgr.Path='%v' Error='%v'", dMgr.Path, err2.Error())
			isEmpty = true
			return
		}

	}


	dMgr.AbsolutePath = fh.AdjustPathSlash(dMgr.AbsolutePath)

	dMgr.AbsolutePathIsPopulated = true
	dMgr.DoesDirMgrAbsolutePathExist()

	strAry := strings.Split(dMgr.AbsolutePath, string(os.PathSeparator))
	lStr := len(strAry)
	idxStr := strAry[lStr-1]

	idx := strings.Index(dMgr.AbsolutePath, idxStr)
	dMgr.ParentPath = fh.RemovePathSeparatorFromEndOfPathString(dMgr.AbsolutePath[0:idx])

	dMgr.ParentPathIsPopulated = true

	if dMgr.AbsolutePathIsPopulated && dMgr.ParentPathIsPopulated {

		dMgr.RelativePath, err2 =	fp.Rel(dMgr.ParentPath, dMgr.AbsolutePath)

		if err2!=nil {
			dMgr.RelativePath = ""
			dMgr.ParentPathIsPopulated = false
		} else {
			dMgr.ParentPathIsPopulated = true
			dMgr.RelativePathIsPopulated = true

		}

	}

	if idxStr != "" {
		dMgr.DirectoryName = idxStr
	} else {
		dMgr.DirectoryName = dMgr.AbsolutePath
	}


	if dMgr.Path != dMgr.AbsolutePath {
		dMgr.AbsolutePathDifferentFromPath = true
	}

	var vn string
	if dMgr.AbsolutePathIsPopulated {
		vn = fp.VolumeName(dMgr.AbsolutePath)
	} else if dMgr.PathIsPopulated {
		vn = fp.VolumeName(dMgr.Path)
	}

	if vn != "" {
		dMgr.VolumeIsPopulated = true
		dMgr.VolumeName = vn
	}

	if dMgr.AbsolutePathIsPopulated && dMgr.PathIsPopulated {
		dMgr.IsInitialized = true
		isEmpty = false
	} else {
		isEmpty = true
	}

	err = nil

	return
}

// SetDirMgrWithFileInfo - Sets the DirMgr fields and path strings for the current DirMgr object
// based on an input 'pathStr' parameter and an os.FileInfo input parameter ('info').
func (dMgr *DirMgr) SetDirMgrWithFileInfo(pathStr string, info os.FileInfo) error {
	ePrefix := "DirMgr.SetDirMgrWithFileInfo() "

	if len(pathStr) == 0 {
		return errors.New(ePrefix + "Error: Input parameter 'pathStr' is Zero Length!")
	}

	_, err := dMgr.SetDirMgr(pathStr)

	if err!=nil {
		return fmt.Errorf(ePrefix + "Error returned from dMgr.SetDirMgr(pathStr). pathStr='%v'  Error='%v'", pathStr, err.Error())
	}

	dMgr.ActualDirFileInfo = FileInfoPlus{}.NewPathFileInfo(pathStr, info)
	dMgr.DirectoryName = info.Name()

	return nil
}


// SubstituteBaseDir - Substitute baseDir segment of the current DirMgr with a new
// parent directory identified by input parameter 'substituteBaseDir'. This is useful
// in copying files to new directory trees.
func (dMgr *DirMgr) SubstituteBaseDir(baseDir DirMgr, substituteBaseDir DirMgr) (newDMgr DirMgr, err error) {

	ePrefix := "DirMgr.SubstituteBaseDir() "

	newDMgr = DirMgr{}
	err = nil


	if !baseDir.IsInitialized {
		err = errors.New(ePrefix + "Error: baseDir DirMgr is NOT Initialized!")
		return
	}


	if !substituteBaseDir.IsInitialized {
		err = errors.New(ePrefix + "Error: substituteBaseDir DirMgr is NOT Initialized!")
		return
	}

	err2 := baseDir.IsDirMgrValid("")

	if err2!=nil {
		err = fmt.Errorf(ePrefix + "Error: baseDir DirMgr object is INVALID!")
		return
	}

	err2 = substituteBaseDir.IsDirMgrValid("")

	if err2!=nil {
		err = fmt.Errorf(ePrefix + "Error: substituteBaseDir DirMgr object is INVALID!")
		return
	}

	var trimmedRemainingPath string
	var newPath string

	if strings.HasPrefix(dMgr.Path, baseDir.Path) {

		trimmedRemainingPath = strings.TrimPrefix(dMgr.Path, baseDir.Path)

		lPath := len(trimmedRemainingPath)

		if lPath > 0 && trimmedRemainingPath[0] == os.PathSeparator {
			trimmedRemainingPath = trimmedRemainingPath[1:]
		}

		newPath = substituteBaseDir.GetPathWithSeparator() + trimmedRemainingPath

	} else if strings.HasPrefix(dMgr.AbsolutePath, baseDir.AbsolutePath) {

		trimmedRemainingPath = strings.TrimPrefix(dMgr.AbsolutePath, baseDir.AbsolutePath)

		lPath := len(trimmedRemainingPath)

		if lPath > 0 && trimmedRemainingPath[0] == os.PathSeparator {
			trimmedRemainingPath = trimmedRemainingPath[1:]
		}

		newPath = substituteBaseDir.GetAbsolutePathWithSeparator() + trimmedRemainingPath

	} else {
		err = fmt.Errorf(ePrefix + "Error: Could not locate baseDir.Path or baseDir.AbsolutePath in this dMgr. dMgr.Path='%v' dMgr.AbsolutePath='%v'", dMgr.Path, dMgr.AbsolutePath)
		return
	}

	newDMgr, err2 = DirMgr{}.New(newPath)

	if err2 != nil {
		newDMgr = DirMgr{}
		err = fmt.Errorf(ePrefix + "Error returned from DirMgr{}.New(newPath). newPath='%v'  Error='%v'", newPath, err2.Error())
		return
	}

	err = nil
	return
}

// FileMgrCollection - A collection of FileMgr objects
type FileMgrCollection struct{
	FMgrs				[]FileMgr
}

// AddFileMgr - Adds a FileMgr object to the collection
func (fMgrs *FileMgrCollection) AddFileMgr(fMgr FileMgr) {
	fMgrs.FMgrs = append(fMgrs.FMgrs, fMgr.CopyOut())
}


func (fMgrs *FileMgrCollection) AddFileMgrByPathFile(pathFileName string) error {
	ePrefix := "FileMgrCollection.AddFileMgrByPathFile() "

	fMgr, err := FileMgr{}.New(pathFileName)

	if err!=nil {
		return fmt.Errorf(ePrefix + "Error returned from FileMgr{}.New(pathFileName). pathFileName='%v' Error='%v'", pathFileName, err.Error())
	}

	fMgrs.FMgrs = append(fMgrs.FMgrs, fMgr)

	return nil
}

// AddFileInfo - Adds a File Manager object to the collection based on input from
// a directory path string and a os.FileInfo object.
func (fMgrs *FileMgrCollection) AddFileInfo(pathFile string, info os.FileInfo) error {

	ePrefix := "FileMgrCollection) AddFileInfo() "

	fMgr, err := FileMgr{}.NewFromFileInfo(pathFile, info)

	if err != nil {
		return fmt.Errorf(ePrefix + "Error retrned from FileMgr{}.NewFromFileInfo(pathFile, info). pathFile='%v' info.Name()='%v'  Error='%v'", pathFile, info.Name(), err.Error())
	}

	fMgrs.FMgrs = append(fMgrs.FMgrs, fMgr)

	return nil
}

// AddFileMgrCollection - Adds another collection of File Manager (FileMgr)
// objects to the current collection.
func (fMgrs *FileMgrCollection) AddFileMgrCollection(fMgrs2 *FileMgrCollection) {

	lOmc2 := len(fMgrs2.FMgrs)

	if lOmc2 == 0 {
		return
	}

	for i:=0; i< lOmc2 ; i++ {
		fMgrs.AddFileMgr(fMgrs2.FMgrs[i].CopyOut())
	}

	return
}

// CopyOut - Returns an FileMgrCollection which is an
// exact duplicate of the current FileMgrCollection
func (fMgrs *FileMgrCollection) CopyOut() (FileMgrCollection, error) {

	ePrefix := "FileMgrCollection.CopyOut() "

	fMgrs2 := FileMgrCollection{}

	lOmc := len(fMgrs.FMgrs)

	if lOmc == 0 {
		return FileMgrCollection{}, errors.New(ePrefix + "Error: Empty FileMgrCollection. No messages available!")
	}

	for i:= 0; i < lOmc; i++ {
		fMgrs2.AddFileMgr(fMgrs.FMgrs[i].CopyOut())
	}

	return fMgrs2, nil
}

// FindFiles - Searches the current FileMgrCollection and returns a new FileMgrCollection containing
// FileMgr objects which match the specified search criteria.
func (fMgrs *FileMgrCollection) FindFiles(fileSelectionCriteria FileSelectionCriteria) (FileMgrCollection, error) {

	ePrefix := "FileMgrCollection.FindFiles() "

	lDirCol := len(fMgrs.FMgrs)

	if lDirCol == 0 {
		return FileMgrCollection{}, nil
	}

	fh := FileHelper{}

	var isMatchedFile bool
	var err error

	fMgrs2 := FileMgrCollection{}

	for i:=0; i < lDirCol; i++ {
		fMgr := fMgrs.FMgrs[i]

		if fMgr.ActualFileInfo.IsFInfoInitialized {

			isMatchedFile, err = fh.FilterFileName(fMgr.ActualFileInfo, fileSelectionCriteria)

			if err!=nil {
				return 	FileMgrCollection{}, fmt.Errorf(ePrefix + "Error returned by fh.FilterFileName(fMgr.ActualFileInfo, fileSelectionCriteria) fMgr.ActualFileInfo.Name()='%v'  Error='%v'",fMgr.ActualFileInfo.Name(), err.Error())
			}

		} else {

			fip := FileInfoPlus{}

			fip.SetName(fMgr.FileNameExt)

			isMatchedFile, err = fh.FilterFileName(fip, fileSelectionCriteria)

			if err!=nil {
				return 	FileMgrCollection{}, fmt.Errorf(ePrefix + "Error returned by fh.FilterFileName(fip, fileSelectionCriteria) fip.Name()='%v'  Error='%v'",fip.Name(), err.Error())
			}

		}

		if isMatchedFile && err == nil {
			fMgrs2.AddFileMgr(fMgr)
		}

	}

	return fMgrs2, nil
}


// GetArrayLength - returns the array length of the
// FileMgrCollection File Managers (FMgrs) array.
func (fMgrs *FileMgrCollection) GetArrayLength() int {
	return len(fMgrs.FMgrs)
}

// PopLastFMgr - Removes the last File Manager (FileMgr) object
// from the collections array, and returns it to the calling method.
func (fMgrs *FileMgrCollection) PopLastFMgr() (FileMgr, error) {

	ePrefix := "FileMgrCollection.PopLastFMgr() "

	l1 := len(fMgrs.FMgrs)

	if l1 == 0 {
		return FileMgr{}, errors.New(ePrefix + "Error: Empty FileMgrCollection. No messages available!")
	}

	fmgr := fMgrs.FMgrs[l1-1].CopyOut()

	fMgrs.FMgrs = fMgrs.FMgrs[0:l1-1]

	return fmgr, nil
}

// PopFirstFMgr - Removes the first OpsMsgDto object
// from the collections array, and returns it to
// the calling method.
func (fMgrs *FileMgrCollection) PopFirstFMgr() (FileMgr, error) {

	ePrefix := "FileMgrCollection.PopFirstFMgr() "

	l1 := len(fMgrs.FMgrs)

	if l1 == 0 {
		return FileMgr{}, errors.New(ePrefix + "Error: Empty FileMgrCollection. No messages available!")
	}

	om := fMgrs.FMgrs[0].CopyOut()

	fMgrs.FMgrs = fMgrs.FMgrs[1:l1]

	return om, nil
}

// PopFMgrAtIndex - Returns a copy of the File Manager (FileMgr) object located
// at index, 'idx', in the FileMgrCollection array. As a 'Pop' method, the original
// FileMgr object is deleted from the FileMgrCollection array.
func (fMgrs *FileMgrCollection) PopFMgrAtIndex(idx int) (FileMgr, error) {

	ePrefix := "FileMgrCollection.PopFMgrAtIndex() "

	if idx < 0 {
		return FileMgr{}, fmt.Errorf(ePrefix + "Error: Input Parameter is less than zero. Index Out-Of-Range! idx='%v'", idx)
	}

	lFMgrs := len(fMgrs.FMgrs)

	if idx >= lFMgrs {
		return FileMgr{}, fmt.Errorf(ePrefix + "Error: Input Parameter is greater than the length of the collection index. Index Out-Of-Range! idx='%v' Array Length='%v' ", idx, lFMgrs)
	}

	if idx == 0 {
		return fMgrs.PopFirstFMgr()
	}

	if idx == lFMgrs - 1 {
		return fMgrs.PopLastFMgr()
	}

	fmgr := fMgrs.FMgrs[idx].CopyOut()

	fCol := FileMgrCollection{}

	for i:=0; i < lFMgrs; i++ {

		if i != idx {
			fCol.FMgrs = append(fCol.FMgrs, fMgrs.FMgrs[i].CopyOut())
		}

	}

	fMgrs.FMgrs = fCol.FMgrs

	return fmgr, nil
}

// PeekFirstFMgr - Returns the first element from the
// FileMgrCollection, but does NOT remove
// it from the OpsMessages array.
func (fMgrs *FileMgrCollection) PeekFirstFMgr() (FileMgr, error) {

	ePrefix := "FileMgrCollection.PeekFirstFMgr() "

	if len(fMgrs.FMgrs) == 0 {
		return FileMgr{}, errors.New(ePrefix + "Error: Empty FileMgrCollection. No messages available!")
	}

	return fMgrs.FMgrs[0].CopyOut(), nil
}

// PeekLastFMgr - Returns the last element from the
// Operation Messages Collection, but does NOT remove
// it from the OpsMessages array.
func (fMgrs *FileMgrCollection) PeekLastFMgr() (FileMgr, error) {

	ePrefix := "FileMgrCollection.PeekLastFMgr()"

	l1 := len(fMgrs.FMgrs)

	if l1==0 {
		return FileMgr{}, errors.New(ePrefix + "Error: Empty FileMgrCollection. No messages available!")
	}

	return fMgrs.FMgrs[l1-1].CopyOut(), nil
}

// PeekFMgrAtIndex - Returns a copy of the File Manager (FileMgr) object located
// at array index 'idx' in the FileMgrCollection. This is a 'Peek' method and the
// original FileMgr object is not deleted from the FileMgrCollection array.
func (fMgrs *FileMgrCollection) PeekFMgrAtIndex(idx int) (FileMgr, error){

	ePrefix := "FileMgrCollection.PeekFMgrAtIndex() "

	if idx < 0 {
		return FileMgr{}, fmt.Errorf(ePrefix + "Error: Input Parameter is less than zero. Index Out-Of-Range! idx='%v'", idx)
	}

	if idx >= len(fMgrs.FMgrs) {
		return FileMgr{}, fmt.Errorf(ePrefix + "Error: Input Parameter is greater than the length of the collection index. Index Out-Of-Range! idx='%v' Array Length='%v' ", idx, len(fMgrs.FMgrs))
	}

	return fMgrs.FMgrs[idx].CopyOut(), nil
}

// FileMgr - This structure and associated methods
// are used to manage a specific file.
type FileMgr struct {
	IsInitialized                   bool
	OriginalPathFileName            string
	DMgr														DirMgr
	AbsolutePathFileName						string
	AbsolutePathFileNameIsPopulated bool
	AbsolutePathFileNameDoesExist		bool
	FileName                        string
	FileNameIsPopulated             bool
	FileExt                         string
	FileExtIsPopulated              bool
	FileNameExt                     string
	FileNameExtIsPopulated          bool
	FilePtr													*os.File
	IsFilePtrOpen										bool
	ActualFileInfo									FileInfoPlus
}



// CopyFileToDir - Copies the file identified by the current File Manager
// (FileMgr) object to another directory specified by input parameter
// 'dir', a 'DirMgr' object.
func (fMgr *FileMgr) CopyFileToDir(dir DirMgr) error {

	ePrefix := "FileMgr.CopyFileToDir() "
	err := dir.IsDirMgrValid("")

	if err != nil {
		return fmt.Errorf(ePrefix + "Error: Input parmater dir is INVALID! Error='%v'", err.Error())
	}


	newFMgr, err := FileMgr{}.NewFromDirMgrFileNameExt(dir, fMgr.FileNameExt)

	if err != nil {
		return fmt.Errorf(ePrefix + "Error returned from FileMgr{}.NewFromDirMgrFileNameExt(dir, fMgr.FileNameExt) dir.AbsolutePath='%v'  fMgr.FileNameExt='%v'  Error='%v'", dir.AbsolutePath, fMgr.FileNameExt, err.Error())
	}

	err = fMgr.CopyFileMgr(&newFMgr)

	if err != nil {
		return fmt.Errorf(ePrefix + "Error returned from fMgr.CopyFileMgr(&newFMgr) newFMgr.AbsolutePathFileName='%v'  Error='%v'", newFMgr.AbsolutePathFileName, err.Error())
	}

	return nil
}

// CopyFileMgr - Copies the file represented by the current File Manager object
// to a location specified by an input parameter File Manager.
//
// Note that if the destination directory does not exist, this method will
// attempt to create it.
func(fMgr *FileMgr) CopyFileMgr(fMgrDest *FileMgr) error {
	ePrefix:= "FileMgr.CopyFileMgr() "

	err := fMgr.IsFileMgrValid("")

	if err!= nil {
		return fmt.Errorf(ePrefix + " Error: This FileMgr object is INVALID! Error='%v'", err.Error())
	}

	if fMgrDest == nil {
		return errors.New(ePrefix + "Error: Input parameter fMgrDest is a nil pointer!")
	}

	err = fMgrDest.IsFileMgrValid("")

	if err!=nil {
		return fmt.Errorf(ePrefix + " Error: The Destination FileMgr object is INVALID! Error='%v'", err.Error())
	}


	if !fMgr.AbsolutePathFileNameDoesExist {
		return fmt.Errorf(ePrefix + "Error: Source file-  %v  DOES NOT EXIST!", fMgr.AbsolutePathFileName)
	}

	if !fMgr.ActualFileInfo.Mode().IsRegular() {
		return fmt.Errorf(ePrefix + "Error: Source file is a Non-Regular File and cannot be compied. File='%v'",fMgr.AbsolutePathFileName)
	}

	err = fMgrDest.DMgr.MakeDir()

	if err != nil {
		return fmt.Errorf(ePrefix + "Atempted creation of destination directory FAILED! Error= '%v'", err.Error())
	}

	destFileExists, err := fMgrDest.DoesThisFileExist()

	if err !=nil {
		return fmt.Errorf(ePrefix + "Error returned by fMgrDest.DoesThisFileExist(). Error= '%v'", err.Error())
	}

	if destFileExists && !fMgrDest.ActualFileInfo.Mode().IsRegular() {
		return fmt.Errorf(ePrefix + "Error: Destination file exists and it is NOT a 'regular' file. Copy operation aborted! Destination File='%v' ", fMgrDest.AbsolutePathFileName)
	}

	// Source and destination files are the same file.
	// No need to copy over source. Just return - Copy
	// operation complete.
	if destFileExists && os.SameFile(fMgr.ActualFileInfo, fMgrDest.ActualFileInfo) {
		return nil
	}

	fMgrDest.DMgr.MakeDir()

	// See Reference:
	// https://stackoverflow.com/questions/21060945/simple-way-to-copy-a-file-in-golang
	/*
	if err = os.Link(fMgr.AbsolutePathFileName, fMgrDest.AbsolutePathFileName); err == nil {

		destFileExists, err = fMgrDest.DoesThisFileExist()

		if err != nil {
			return fmt.Errorf(ePrefix + "Error returned from LinkSet fMgrDest.DoesThisFileExist(). fMgrDest.AbsolutePathFileName='%v'  Error='%v'", fMgrDest.AbsolutePathFileName, err.Error())
		}

		if !destFileExists {
			return fmt.Errorf(ePrefix + "Error from LinkSet: After attempted file copy to destination file. Destination file does NOT exist! fMgrDest.AbsolutePathFileName='%v'",fMgrDest.AbsolutePathFileName)
		}

	 	return nil
	}
	*/

	fh := FileHelper{}

	err = fh.CopyFileContents(fMgr.AbsolutePathFileName, fMgrDest.AbsolutePathFileName)

	if err != nil {
		return fmt.Errorf(ePrefix + "Error returned by fh.CopyFileContents(fMgr.AbsolutePathFileName, fMgrDest.AbsolutePathFileName) fMgr.AbsolutePathFileName='%v'  fMgrDest.AbsolutePathFileName='%v'  Error='%v'",fMgr.AbsolutePathFileName, fMgrDest.AbsolutePathFileName, err.Error())
	}

	destFileExists, err = fMgrDest.DoesThisFileExist()

	if err != nil {
		return fmt.Errorf(ePrefix + "Error returned from fMgrDest.DoesThisFileExist(). fMgrDest.AbsolutePathFileName='%v'  Error='%v'", fMgrDest.AbsolutePathFileName, err.Error())
	}

	if !destFileExists {
		return fmt.Errorf(ePrefix + "Error: After attempted file copy to destination file. Destination file does NOT exist! fMgrDest.AbsolutePathFileName='%v'",fMgrDest.AbsolutePathFileName)
	}

	return nil
}

// CopyFileStr - Copies file from fMgr.AbsolutePathFileName to
// to destination path & File Name.
// See: https://stackoverflow.com/questions/21060945/simple-way-to-copy-a-file-in-golang
func (fMgr *FileMgr) CopyFileStr(dstPathFileName string) error {

	ePrefix := "FileMgr.CopyFileStr "

	fMgrDest, err := FileMgr{}.New(dstPathFileName)

	if err != nil {
		return fmt.Errorf(ePrefix + "Error returned by FileMgr{}.New(dstPathFileName). dstPathFileName='%v' Error='%v'", dstPathFileName, err.Error())
	}

	err = fMgr.CopyFileMgr(&fMgrDest)

	if err != nil {
		return fmt.Errorf(ePrefix + "Error returned from fMgr.CopyFileMgr(&fMgrDest) fMgrDest.AbsolutePathFileName='%v'  Error='%v'", fMgrDest.AbsolutePathFileName, err.Error())
	}


	return nil
}

// CloseFile - This method will call the Close()
// method on the current file pointer, FileHelper.FilePtr
func (fMgr *FileMgr) CloseFile() error {

	ePrefix := "FileMgr.CloseFile() "
	var err error

	if fMgr.FilePtr == nil {
		fMgr.IsFilePtrOpen = false
		return nil
	}

	err = fMgr.FlushBytesToDisk()

	if err!=nil {
		return fmt.Errorf(ePrefix + "Error returned from fMgr.FlushBytesToDisk().  Error='%v'", err.Error())
	}

	err = fMgr.FilePtr.Close()

	if err != nil {
		return fmt.Errorf(ePrefix + "Received Error from fMgr.FilePtr.Close(). fMgr.AbsolutePathFileName")
	}

	fMgr.IsFilePtrOpen = false
	fMgr.FilePtr = nil

	return nil
}

// CopyIn - Copies data from an incoming FileMgr object
// into the current FileMgr object.
func (fMgr *FileMgr) CopyIn(fmgr2 *FileMgr) {

	if fmgr2 == nil {
		panic("FileMgr.CopyIn() - Input parameter is a nil pointer!")
	}

	fMgr.IsInitialized = fmgr2.IsInitialized
	fMgr.DMgr.CopyIn(&fmgr2.DMgr)
	fMgr.OriginalPathFileName = fmgr2.OriginalPathFileName
	fMgr.AbsolutePathFileName = fmgr2.AbsolutePathFileName
	fMgr.AbsolutePathFileNameIsPopulated = fmgr2.AbsolutePathFileNameIsPopulated
	fMgr.AbsolutePathFileNameDoesExist = fmgr2.AbsolutePathFileNameDoesExist
	fMgr.FileName = fmgr2.FileName
	fMgr.FileNameIsPopulated = fmgr2.FileNameIsPopulated
	fMgr.FileExt = fmgr2.FileExt
	fMgr.FileExtIsPopulated = fmgr2.FileExtIsPopulated
	fMgr.FileNameExt = fmgr2.FileNameExt
	fMgr.FileNameExtIsPopulated = fmgr2.FileNameExtIsPopulated
	fMgr.FilePtr = fmgr2.FilePtr
	fMgr.IsFilePtrOpen = fmgr2.IsFilePtrOpen
	fMgr.ActualFileInfo = fmgr2.ActualFileInfo.CopyOut()

	return
}

// CopyOut - Duplicates the file information in the current
// FileMgr object and returns it as a new FileMgr object.
func (fMgr *FileMgr) CopyOut() FileMgr {

	fmgr2 := FileMgr{}

	fmgr2.IsInitialized = fMgr.IsInitialized
	fmgr2.DMgr = fMgr.DMgr.CopyOut()
	fmgr2.OriginalPathFileName = fMgr.OriginalPathFileName
	fmgr2.AbsolutePathFileName = fMgr.AbsolutePathFileName
	fmgr2.AbsolutePathFileNameIsPopulated = fMgr.AbsolutePathFileNameIsPopulated
	fmgr2.AbsolutePathFileNameDoesExist = fMgr.AbsolutePathFileNameDoesExist
	fmgr2.FileName = fMgr.FileName
	fmgr2.FileNameIsPopulated = fMgr.FileNameIsPopulated
	fmgr2.FileExt = fMgr.FileExt
	fmgr2.FileExtIsPopulated = fMgr.FileExtIsPopulated
	fmgr2.FileNameExt = fMgr.FileNameExt
	fmgr2.FileNameExtIsPopulated = fMgr.FileNameExtIsPopulated
	fmgr2.FilePtr = fMgr.FilePtr
	fmgr2.IsFilePtrOpen = fMgr.IsFilePtrOpen
	fmgr2.ActualFileInfo = fMgr.ActualFileInfo.CopyOut()

	return fmgr2
}

// Equal - Compares a second FileHelper data structure
// to the current FileHelper data structure and returns
// a boolean value indicating whether they are equal
// in all respects.
func (fMgr *FileMgr) Equal(fmgr2 *FileMgr) bool {

	if fMgr.IsInitialized != fmgr2.IsInitialized ||
		fMgr.OriginalPathFileName != fmgr2.OriginalPathFileName ||
		fMgr.AbsolutePathFileNameIsPopulated != fmgr2.AbsolutePathFileNameIsPopulated ||
		fMgr.AbsolutePathFileNameDoesExist != fmgr2.AbsolutePathFileNameDoesExist ||
		fMgr.AbsolutePathFileName != fmgr2.AbsolutePathFileName ||
		fMgr.FileName != fmgr2.FileName ||
		fMgr.FileNameIsPopulated != fmgr2.FileNameIsPopulated ||
		fMgr.FileExt != fmgr2.FileExt ||
		fMgr.FileExtIsPopulated != fmgr2.FileExtIsPopulated ||
		fMgr.FileNameExt != fmgr2.FileNameExt ||
		fMgr.FileNameExtIsPopulated != fmgr2.FileNameExtIsPopulated ||
		fMgr.FilePtr != fmgr2.FilePtr ||
		fMgr.IsFilePtrOpen != fmgr2.IsFilePtrOpen {

		return false
	}

	if !fMgr.DMgr.Equal(&fmgr2.DMgr) {
		return false
	}

	if !fMgr.ActualFileInfo.Equal(&fmgr2.ActualFileInfo) {
		return false
	}

	return true
}

// Empty - resets all data fields in the FileMgr structure to 
// their uninitialized or zero state.
func(fMgr *FileMgr) Empty() {
	fMgr.IsInitialized 										= false
	fMgr.DMgr 														= DirMgr{}
	fMgr.OriginalPathFileName 						= ""
	fMgr.AbsolutePathFileName 						= ""
	fMgr.AbsolutePathFileNameIsPopulated 	= false
	fMgr.AbsolutePathFileNameDoesExist 		= false
	fMgr.FileName 												= ""
	fMgr.FileNameIsPopulated 							= false
	fMgr.FileExt 													= ""
	fMgr.FileExtIsPopulated 							= false
	fMgr.FileNameExt 											= ""
	fMgr.FileNameExtIsPopulated 					= false
	fMgr.FilePtr 													= nil
	fMgr.IsFilePtrOpen 										= false
	fMgr.ActualFileInfo 									= FileInfoPlus{}

}

// CreateDirAndFile - Performs two operations:
// This is a Wrapper function for os.Create - Create a file.
//
//  If the home directory does not currently exist, this method
//  will first create the directory tree, before creating the new
//	file.
//
// If the file previously exists, it will be truncated.
//
// Note that if this method successfully creates the file, a
// File Pointer (*File) will be stored in the FileMgr field
// fMgr.FilePtr. Be sure to close the File Pointer when
// finished with it. See FileMgr.CloseFile().
//
func (fMgr *FileMgr) CreateDirAndFile()  error {

	ePrefix := "FileMgr:CreateDirAndFile() "
	var err error

	err = fMgr.IsFileMgrValid(ePrefix)

	if err!=nil {
		return err
	}

	fh := FileHelper{}

	if !fh.DoesFileExist(fMgr.DMgr.AbsolutePath) {
		// Directory does NOT exist, create it!

		err := fh.MakeDirAll(fMgr.DMgr.AbsolutePath)

		if err != nil {
			return fmt.Errorf(ePrefix + "Errors from FileHelper:MakeDirAll(fMgr.DMgr.AbsolutePath). fMgr.DMgr.AbsolutePath='%v'  Error='%v' ",fMgr.DMgr.AbsolutePath, err.Error())
		}

		fMgr.DMgr.AbsolutePathDoesExist = true

	} else {

		fMgr.DMgr.AbsolutePathDoesExist = true

	}


	fMgr.FilePtr, err = os.Create(fMgr.AbsolutePathFileName)

	if err != nil {
		return fmt.Errorf(ePrefix + "Error creating File. Error returned from os.Create(fMgr.AbsolutePathFileName). fMgr.AbsolutePathFileName='%v' Error='%v' ",fMgr.AbsolutePathFileName, err.Error() )
	}

	fMgr.IsFilePtrOpen = true

	return nil

}

// CreateFile - Creates the File identified by FileMgr.AbsolutePathFileName.
// If the directory in the path file name designation does not exist, this
// method will throw an error.
//
// See Method CreateDirAndFile() which will create both the directory and the file
// as required.
//
// Note that if the file is actually created, the returned file pointer (*File)
// is stored in the FileMgr field, fMgr.FilePtr. Be sure to 'close' the File Pointer
// when finished with it. See FileMgr.CloseFile()
//
func (fMgr *FileMgr) CreateFile() error {

	ePrefix := "FileMgr:CreateFile() Error - "

	fh := FileHelper{}

	if !fMgr.IsInitialized {
		return errors.New(ePrefix + " FileMgr is NOT Initialized!")
	}

	if !fMgr.DMgr.AbsolutePathIsPopulated {
		return errors.New(ePrefix + " FileMgrDMgr.AbsolutePathIsPopulated is NOT populated!")
	}

	if fMgr.AbsolutePathFileName == "" {
		fMgr.AbsolutePathFileNameIsPopulated = false
		return errors.New(ePrefix + " FileMgr.AbsolutePathFileName is EMPTY!")
	}

	if !fh.DoesFileExist(fMgr.DMgr.AbsolutePath) {
		fMgr.DMgr.AbsolutePathDoesExist = false
		return fmt.Errorf(ePrefix + "FileMgr.DMgr.AbsolutePath Does NOT exist! Create the path. FileMgr.DMgr.AbsolutePath='%v'", fMgr.DMgr.AbsolutePath)
	}

	var err error

	fMgr.FilePtr, err = os.Create(fMgr.AbsolutePathFileName)

	if err != nil {
		return fmt.Errorf(ePrefix + "Error creating File. Error returned from os.Create(fMgr.AbsolutePathFileName). fMgr.AbsolutePathFileName='%v' Error='%v' ",fMgr.AbsolutePathFileName, err.Error())
	}

	fMgr.IsFilePtrOpen = true

	return nil

}

// DeleteThisFile - Deletes the file identified by FileMgr.AbsolutePathFileName
// in the current FileHelper structure.
func (fMgr *FileMgr) DeleteThisFile() error {

	ePrefix := "FileMgr.DeleteThisFile()"

	err := fMgr.IsFileMgrValid("")

	if err!=nil {
		return fmt.Errorf(ePrefix + " Error: This FileMgr object is INVALID!  Error='%v'", err.Error())
	}


	if fMgr.FilePtr != nil {
		fMgr.FilePtr.Close()
	}

	fMgr.IsFilePtrOpen = false


	err = os.Remove(fMgr.AbsolutePathFileName)

	if err != nil {
		return fmt.Errorf(ePrefix + " - os.Remove(fMgr.AbsolutePathFileName) returned an error. AbsolutePathFileName='%v'   Error='%v'", fMgr.AbsolutePathFileName, err.Error())
	}

	fileExists, err := fMgr.DoesThisFileExist()

	if err!=nil {
		return fmt.Errorf(ePrefix + "Error returned by fMgr.DoesThisFileExist() fMgr.AbsolutePathFileName='%v'  Error='%v'", fMgr.AbsolutePathFileName, err.Error())
	}

	if fileExists {
		return fmt.Errorf(ePrefix + "Error: Attempted file deletion FAILED!. File still exists. fMgr.AbsolutePathFileName='%v'", fMgr.AbsolutePathFileName)
	}

	fMgr.ActualFileInfo = FileInfoPlus{}

	return nil
}


// DoesThisFileExist - Returns a boolean value
// designated whether the file specified by the
// current FileHelper.AbsolutePathFileName field
// exists.
func(fMgr *FileMgr) DoesThisFileExist() (bool, error) {

	ePrefix := "FileMgr.DoesThisFileExist() "

	if !fMgr.IsInitialized {
		return false,  errors.New(ePrefix + " Error: The File Manager data structure has NOT been initialized.")
	}

	if fMgr.AbsolutePathFileNameIsPopulated == false {
		return false, errors.New(ePrefix + " Error: AbsolutePathFileName is NOT POPULATED!")
	}

	if fMgr.AbsolutePathFileName == "" {
		fMgr.AbsolutePathFileNameIsPopulated = false
		return false, errors.New(ePrefix + " Error: AbsolutePathFileName is EMPTY!")
	}

	info, err := os.Stat(fMgr.AbsolutePathFileName)


	if err != nil {
		fMgr.ActualFileInfo = FileInfoPlus{}
		fMgr.AbsolutePathFileNameDoesExist = false
	} else {
		fMgr.AbsolutePathFileNameIsPopulated = true
		fMgr.AbsolutePathFileNameDoesExist = true
		fMgr.ActualFileInfo = FileInfoPlus{}.NewPathFileInfo(fMgr.AbsolutePathFileName, info)
	}

	fMgr.DMgr.DoesDirMgrPathExist()
	fMgr.DMgr.DoesDirMgrPathExist()



	return fMgr.AbsolutePathFileNameDoesExist, nil
}

// FlushBytesToDisk - After Writing bytes to a file, use this
// method to commit the contents of the current file to
// stable storage.
func (fMgr *FileMgr) FlushBytesToDisk() error {

	ePrefix := "FileMgr.FlushBytesToDisk() "

	var err error

	if fMgr.IsFilePtrOpen && fMgr.FilePtr!= nil {

		err = fMgr.FilePtr.Sync()

		if err!=nil {
			return fmt.Errorf(ePrefix + "Error returned from fMgr.FilePtr.Sync()")
		}

	}

	return nil
}

// GetFileInfo - Wrapper function for os.Stat(). This method
// can be used to return FileInfo data on the specific file identified
// by FileMgr.AbsolutePathFileName. If the file does NOT exist,
// an error will be triggered.
//
// type FileInfo interface {
// 	Name() string       // base name of the file
// 	Size() int64        // length in bytes for regular files; system-dependent for others
// 	Mode() FileMode     // file mode bits
// 	ModTime() time.Time // modification time
// 	IsDir() bool        // abbreviation for Mode().IsDir()
// 	Sys() interface{}   // underlying data source (can return nil)
// }
func (fMgr *FileMgr) GetFileInfo() (os.FileInfo, error) {

	ePrefix := "FileMgr.DeleteThisFile()"

	if !fMgr.IsInitialized {
		return nil, errors.New(ePrefix + " Error: This data structure is NOT initialized.")
	}

	if !fMgr.AbsolutePathFileNameIsPopulated {
		return nil, errors.New(ePrefix + " Error: AbsolutePathFileName is NOT populated/initialized.")
	}

	if fMgr.AbsolutePathFileName == "" {
		fMgr.AbsolutePathFileNameIsPopulated = false
		return nil, errors.New(ePrefix + " Error: AbsolutePathFileName is EMPTY!")
	}


	info, err := os.Stat(fMgr.AbsolutePathFileName)

	if err != nil {
		return nil, fmt.Errorf(ePrefix + "Error returned by os.Stat(fMgr.AbsolutePathFileName). fMgr.AbsolutePathFileName='%v'  Error='%v'",fMgr.AbsolutePathFileName, err.Error())
	}

	fMgr.ActualFileInfo = FileInfoPlus{}.NewPathFileInfo(fMgr.DMgr.AbsolutePath, info)

	return info, nil
}

// IsFileMgrValid - Analyzes the current FileMgr object. If the
// current FileMgr object is INVALID, an error is returned.
//
// If the current FileMgr is VALID, this method returns 'nil'
//
func(fMgr *FileMgr) IsFileMgrValid(errorPrefixStr string) error {

	ePrefix := strings.TrimRight(errorPrefixStr, " ") + "FileMgr.IsFileMgrValid()"

	if !fMgr.IsInitialized {
		return errors.New(ePrefix + " Error: This data structure is NOT initialized.")
	}

	if !fMgr.AbsolutePathFileNameIsPopulated {
		return errors.New(ePrefix + " Error: AbsolutePathFileName is NOT populated/initialized.")
	}

	if fMgr.AbsolutePathFileName == "" {
		fMgr.AbsolutePathFileNameIsPopulated = false
		return errors.New(ePrefix + " Error: AbsolutePathFileName is EMPTY!")
	}

	fMgr.DoesThisFileExist()

	err := fMgr.DMgr.IsDirMgrValid(ePrefix)

	if err != nil {
		return fmt.Errorf("FileMgr Directory Manager INVALID - %v", err.Error())
	}

	return nil
}


// MoveFileToNewDirMgr - This method will move the file identified
// by the current FileMgr to a new path contained in the input parameter
// 'dMgr'.
//
// IMPORTANT:
// ==========
// The current file identified by the current FileMgr object will
// be DELETED!
//
// The new file located in the new directory will be returned in the return
// parameter 'newFMgr'.
//
func (fMgr *FileMgr) MoveFileToNewDirMgr(dMgr DirMgr) (newFMgr FileMgr, err error) {
	ePrefix := "FileMgr.MoveFileToNewDirMgr() "
	newFMgr = FileMgr{}
	err = nil

	if !fMgr.IsInitialized {
		err = errors.New(ePrefix + "Error: The current FileMgr object is NOT initialized. It is Empty!")
		return
	}

	err2 := fMgr.IsFileMgrValid("")

	if err2 != nil {
		err = fmt.Errorf(ePrefix + "Error: Current FileMgr object is INVALID!. Error='%v'", err2.Error())
		return
	}

	if !dMgr.IsInitialized {
		err = errors.New(ePrefix + "Error: Input parameter 'dMgr' is NOT initialized. It is Empty!")
		return
	}

	err2 = dMgr.IsDirMgrValid("")

	if err2 != nil {
		err = errors.New(ePrefix + "Error: Input parameter 'dMgr' reports as INVALID!")
		return
	}

	if !dMgr.AbsolutePathDoesExist && !dMgr.PathDoesExist {
		err = fmt.Errorf(ePrefix + "Error: Destination Path DOES NOT EXIST!. For this DirMgr object, both AbsolutePath and Path DO NOT EXIST! dMgr.AbsolutePath='%v'  dMgr.Path='%v'", dMgr.AbsolutePath, dMgr.Path)
		return
	}

	srcFileExists, err2 := fMgr.DoesThisFileExist()

	if err2 != nil {
		err = fmt.Errorf("Error returned by fMgr.DoesThisFileExist(). fMgr.AbsolutePathFileName='%v'  Error='%v'", fMgr.AbsolutePathFileName, err2.Error())
		return
	}

	if !srcFileExists {
		err = fmt.Errorf("Error: The source file identified by the current FileMgr object DOES NOT EXIST! fMgr.AbsolutePathFileName='%v'", fMgr.AbsolutePathFileName)
		return
	}

	var destPathFileName string

	if dMgr.AbsolutePathDoesExist {
		destPathFileName = dMgr.GetAbsolutePathWithSeparator() + fMgr.FileNameExt
	} else {
		destPathFileName = dMgr.GetPathWithSeparator() + fMgr.FileNameExt
	}


	fh := FileHelper{}
	_, err = fh.MoveFile(fMgr.AbsolutePathFileName, destPathFileName)

	if err!=nil {
		return FileMgr{}, fmt.Errorf(ePrefix + "Error returned from fh.MoveFile(fMgr.AbsolutePathFileName, destPathFileName). fMgr.AbsolutePathFileName='%v'  pathFile='%v' Error='%v'", fMgr.AbsolutePathFileName, destPathFileName, err.Error())
	}

	newFMgr, err2 = FileMgr{}.New(destPathFileName)

	if err2 != nil {
		newFMgr = FileMgr{}
		err = fmt.Errorf(ePrefix + "Error returned by FileMgr{}.New(destPathFileName). destPathFileName='%v' Error='%v'", destPathFileName, err2.Error())
		return
	}

	fMgr.DoesThisFileExist()

	err = nil

	return

}

// MoveFileToNewDir - This method will move the current file
// identified by this FileMgr object to a new path designated
// by input parameter string, 'dirPath'.
//
// IMPORTANT:
// ==========
// The current file identified by the current FileMgr object will
// be DELETED!
//
// The new file located in the new directory will be returned in the return
// parameter 'newFMgr'.
//
func (fMgr *FileMgr) MoveFileToNewDir(dirPath string) (newFMgr FileMgr, err error) {

	newFMgr = FileMgr{}
	err = nil

	ePrefix := "FileMgr.MoveFileToNewDir() "

	lPath := len(dirPath)

	if lPath==0 {
		err = errors.New(ePrefix + "Error: Input parameter 'dirPath' is a Zero length string!")
		return
	}

	if !fMgr.IsInitialized {
		err = errors.New(ePrefix + "Error: The current FileMgr object is NOT Initialized! It is EMPTY!")
		return
	}

	err2 := fMgr.IsFileMgrValid("")

	if err2 != nil {
		err = fmt.Errorf(ePrefix + "Error: Current FileMgr object is INVALID!. Error='%v'", err2.Error())
		return
	}


	if !fMgr.AbsolutePathFileNameDoesExist {
		err = fmt.Errorf(ePrefix + "Error: The source files does NOT exist. srcFile='%v' ", fMgr.AbsolutePathFileName)
		return
	}

	dMgr, err2 := DirMgr{}.New(dirPath)

	if err2 != nil {
		err = fmt.Errorf(ePrefix+"Error returned from DirMgr{}.New(dirPath). dirPath='%v'  Error='%v'", dirPath, err2.Error())
		return
	}

	if !dMgr.IsInitialized {
		err = fmt.Errorf(ePrefix + "Error: Input parameter 'dirPath' generated an empty DirMgr object. dirPath='%v'", dirPath)
		return
	}

	pathExists, absPathExists := dMgr.DoesDirectoryExist()

	if !pathExists && !absPathExists {
		err = fmt.Errorf(ePrefix + "Error: Target Destination Path DOES NOT EXIST! dirPath='%v'", dirPath)
		return
	}

	newFMgr, err2 = fMgr.MoveFileToNewDirMgr(dMgr)

	if err2 != nil {
		newFMgr = FileMgr{}
		err = fmt.Errorf(ePrefix + "Error returned by fMgr.MoveFileToNewDirMgr(dMgr). dMgr.Path='%v' Error='%v'", dMgr.Path, err2.Error())
		return
	}

	return newFMgr, nil
}

// New - Creates a new FileMgr object. Input parameter parses out the
// path, file name and file extension. The file data is returned in
// the data fields of the new FileMgr object.
//
// Input Parameter
// ===============
//
// pathFileNameExt string - Must consist of a valid path, file name
// 													and file extension. The file need not exist.
//													Failure to provide a properly formatted path
//													path, file name will result in an error.
//
// Example Usage:
// fmgr := FileMgr{}.New("../common/FileName.ext")
//
func (fMgr FileMgr) New(pathFileNameExt string) (FileMgr, error) {

	ePrefix := "FileMgr.New() "

	if pathFileNameExt == "" {
		return FileMgr{}, errors.New(ePrefix +"-Error: pathFileNameExt is Empty!")
	}

	fMgrOut := FileMgr{}

	isEmpty, err := fMgrOut.SetFileMgrFromPathFileName(pathFileNameExt)

	if err != nil {
		return FileMgr{}, fmt.Errorf(ePrefix + "Error returned from fMgrOut.SetFileMgrFromPathFileName(pathFileNameExt) pathFileNameExt='%v'  Error='%v'", pathFileNameExt, err.Error())
	}

	if isEmpty {
		return FileMgr{}, fmt.Errorf(ePrefix + "Error: Returned FileMgr is Empty! pathFileNameExt='%v' ", pathFileNameExt)
	}

	return fMgrOut, nil
}


// NewFromFileInfo - Creates and returns a new FileMgr object based on input from a
// directory path string and an os.FileInfo object.
func (fMgr FileMgr) NewFromFileInfo(pathStr string, info os.FileInfo) (FileMgr, error) {

	ePrefix := "FileMgr.NewFromFInfo() "
	var err error

	fh := FileHelper{}

	pathStr, err = fh.AddPathSeparatorToEndOfPathStr(pathStr)

	if err!= nil {
		return FileMgr{}, fmt.Errorf(ePrefix + "Error returned from fh.AddPathSeparatorToEndOfPathStr(pathStr). pathStr='%v' Error='%v'", pathStr, err.Error())
	}

	pathFileName := pathStr + info.Name()

	fmgr2 := FileMgr{}

	isEmpty, err := fmgr2.SetFileMgrFromPathFileName(pathFileName)

	if err!= nil {
		return FileMgr{}, fmt.Errorf(ePrefix + "Error returned from fmgr2.SetFileMgrFromPathFileName(pathFileName). pathFileName='%v' Error='%v'", pathFileName, err.Error())
	}

	if isEmpty {
		return FileMgr{}, fmt.Errorf(ePrefix + "Error returned FileMgr is Empty! pathFileName='%v'", pathFileName)
	}

	fmgr2.ActualFileInfo = FileInfoPlus{}.NewPathFileInfo(pathStr , info)

	return  fmgr2, nil
}


// NewFromDirMgrFileNameExt - this method is designed to create a new File Manager (FileMgr)
// object from two input parameters.
//
// Input parameter 'dirMgr' is a valid and correctly populated Directory Manager object containing
// the file path.
//
// Input parameter 'fileNameExt' is a string containing the file name and the file extension.
//
// 'dirMgr' and 'fileNameExt' will be combined to create a new File Manager (FileMgr) with a
// properly configured path, file name and file extension.
//
func (fMgr FileMgr) NewFromDirMgrFileNameExt(dirMgr DirMgr, fileNameExt string) (FileMgr, error) {
	ePrefix := "FileMgr.NewFromDirMgrFileNameExt() "

	if len(fileNameExt) == 0 {
		return FileMgr{}, errors.New(ePrefix + "Error: Input Parameter fileNameExt is a zero length string!")
	}

	if dirMgr.IsInitialized {
		err := dirMgr.IsDirMgrValid("")

		if err != nil {
			return FileMgr{}, errors.New(ePrefix + "Error: DirMgr object, 'dirMgr', passed as input parameter is INVALID!")
		}

	}

	fmgr2 := FileMgr{}

	isEmpty, err := fmgr2.SetFileMgrFromDirMgrFileName(dirMgr, fileNameExt)

	if err != nil {
		return FileMgr{}, fmt.Errorf(ePrefix + "Error returned by fmgr2.SetFileMgrFromDirMgrFileName(dirMgr, fileNameExt). Error='%v'", err.Error())
	}

	if isEmpty {
		return FileMgr{}, fmt.Errorf(ePrefix + "Error: Empty FileMgr object returned by fmgr2.SetFileMgrFromDirMgrFileName(dirMgr, fileNameExt) dirMgr.Path='%v' fileNameExt='%v'", dirMgr.Path, fileNameExt)
	}

	return  fmgr2, nil
}

// NewFromDirStrFileNameStr - Creates a new file manager object (FileMgr) from a directory
// string and a File Name and Extension string passed as input parameters.
func (fMgr FileMgr) NewFromDirStrFileNameStr(dirStr, fileNameExtStr string) (FileMgr, error) {

	ePrefix := "FileMgr.NewFromDirStrFileNameStr() "

	if len(dirStr) == 0 {
		return FileMgr{}, errors.New(ePrefix + "Error: Input parameter 'dirStr' is a Zero Length String!")
	}


	if len(fileNameExtStr) == 0 {
		return FileMgr{}, errors.New(ePrefix + "Error: Input parameter 'fileNameExtStr' is a Zero Length String!")
	}

	dMgr, err := DirMgr{}.New(dirStr)

	if err != nil {
		return FileMgr{}, fmt.Errorf(ePrefix + "Error returned by DirMgr{}.New(dirStr). dirStr='%v'  Error='%v'", dirStr, err.Error())
	}

	fMgr2, err := FileMgr{}.NewFromDirMgrFileNameExt(dMgr, fileNameExtStr)

	if err != nil {
		return FileMgr{}, fmt.Errorf(ePrefix + "Error returned from FileMgr{}.NewFromDirMgrFileNameExt(dMgr, fileNameExtStr). fileNameExtStr='%v'  Error='%v'", fileNameExtStr, err.Error())
	}

	return fMgr2, nil

}

// OpenThisFileReadOnly - Opens the file identified by the current
// FileMgr object as a 'Read-Only' File. Subsequent operations may
// read from this file but may NOT write to this file.
//
//
func (fMgr *FileMgr) OpenThisFileReadOnly() error {
	ePrefix := "FileMgr.OpenThisFileReadWrite()"
	var err error

	if !fMgr.IsInitialized {
		return errors.New(ePrefix + " Error: The File Manager data structure has NOT been initialized.")
	}

	if !fMgr.AbsolutePathFileNameIsPopulated {
		return errors.New(ePrefix + " Error: FileMgr.AbsolutePathFileName has NOT been initialized and populated.")
	}

	if fMgr.AbsolutePathFileName == "" {
		fMgr.AbsolutePathFileNameIsPopulated = false
		return errors.New(ePrefix + " Error: FileMgr.AbsolutePathFileName is EMPTY!")
	}

	if fMgr.IsFilePtrOpen {
		err = fMgr.CloseFile()
		if err!=nil {
			return fmt.Errorf(ePrefix + "Error returned by fMgr.CloseFile(). AbsolutePathFileName='%v'  Error='%v'",fMgr.AbsolutePathFileName, err.Error() )
		}
	}

	doesFileExist, err := fMgr.DoesThisFileExist()

	if err != nil {
		return fmt.Errorf(ePrefix + " Error returned from fMgr.DoesThisFileExist() - %v", err.Error())
	}

	if !doesFileExist {

		err = fMgr.CreateDirAndFile()

		if err != nil {
			return fmt.Errorf(ePrefix + " Error from fMgr.CreateDirAndFile(fMgr.AbsolutePathFileName). AbsolutePathFileName='%v'. Error='%v'", fMgr.AbsolutePathFileName, err.Error())
		}

		fMgr.AbsolutePathFileNameDoesExist = true
		fMgr.AbsolutePathFileNameIsPopulated = true
	}

	if err != nil {
		return fmt.Errorf(ePrefix + "Error creating File. Error returned from os.Create(fMgr.AbsolutePathFileName). fMgr.AbsolutePathFileName='%v' Error='%v' ",fMgr.AbsolutePathFileName, err.Error() )
	}

	fMgr.FilePtr, err = os.OpenFile(fMgr.AbsolutePathFileName,os.O_RDONLY,0666)

	if err != nil {
		return fmt.Errorf(ePrefix + " Error opening file: '%v' Error= '%v'", fMgr.AbsolutePathFileName, err.Error())
	}

	fMgr.IsFilePtrOpen = true

	return nil

}

// OpenThisFileReadWrite - Opens a file using file data in the
// current FileHelper data fields. If successful, this method
// will use FileHelper.AbsolutePathFileName to open an *os.File
// or File Pointer.
//
// As the method's name implies, the 'FileHelper.AbsolutePathFileName'
// will be opened for reading and writing. If FileHelper.AbsolutePathFileName
// does not exist, it will be created. The FileMode is set to'rwxrwxrwx' and
// the permission Mode= '0666'
//
func (fMgr *FileMgr) OpenThisFileReadWrite() error {
	var err error

	ePrefix := "FileMgr.OpenThisFileReadWrite()"

	if !fMgr.IsInitialized {
		return errors.New(ePrefix + " Error: The File Manager data structure has NOT been initialized.")
	}

	if !fMgr.AbsolutePathFileNameIsPopulated {
		return errors.New(ePrefix + " Error: FileMgr.AbsolutePathFileName has NOT been initialized and populated.")
	}

	if fMgr.AbsolutePathFileName == "" {
		fMgr.AbsolutePathFileNameIsPopulated = false
		return errors.New(ePrefix + " Error: FileMgr.AbsolutePathFileName is EMPTY!")
	}

	if fMgr.IsFilePtrOpen || fMgr.FilePtr != nil {
		err = fMgr.CloseFile()
		if err!=nil {
			return fmt.Errorf(ePrefix + "Error returned by fMgr.CloseFile(). AbsolutePathFileName='%v'  Error='%v'",fMgr.AbsolutePathFileName, err.Error() )
		}
	}

	doesFileExist, err := fMgr.DoesThisFileExist()

	if err != nil {
		return fmt.Errorf(ePrefix + " Error returned from fMgr.DoesThisFileExist() - %v", err.Error())
	}

	if !doesFileExist {

		err = fMgr.CreateDirAndFile()

		if err != nil {
			return fmt.Errorf(ePrefix + " Error from fMgr.CreateDirAndFile(fMgr.AbsolutePathFileName). AbsolutePathFileName='%v'. Error='%v'", fMgr.AbsolutePathFileName, err.Error())
		}

		fMgr.AbsolutePathFileNameDoesExist = true
		fMgr.AbsolutePathFileNameIsPopulated = true

		return nil
	}


	fMgr.FilePtr, err = os.OpenFile(fMgr.AbsolutePathFileName,os.O_RDWR,0666)

	if err != nil {
		return fmt.Errorf(ePrefix + " Error opening file: '%v' Error= '%v'", fMgr.AbsolutePathFileName, err.Error())
	}


	fMgr.IsFilePtrOpen = true

	return nil
}


// ReadAllFile - Reads the file identified by the current FileMgr
// and returns the contents in a byte array.
//
// If no errors are encountered the returned 'error' value is
// nil.
func (fMgr *FileMgr) ReadAllFile() ([]byte, error) {
	ePrefix := "FileMgr.ReadAllFile() "
	var err error

	if !fMgr.IsInitialized {
		return []byte{}, errors.New(ePrefix + " Error: The File Manager data structure has NOT been initialized.")
	}

	err = fMgr.IsFileMgrValid("")

	if err != nil {
		return []byte{}, fmt.Errorf(ePrefix + "Error - This File Manager is INVALID! Error='%v'", err.Error())
	}

	if !fMgr.AbsolutePathFileNameIsPopulated {
		return []byte{}, errors.New(ePrefix + "Error: FileMgr.AbsolutePathFileName has NOT been initialized and populated.")
	}

	if fMgr.AbsolutePathFileName == "" {
		fMgr.AbsolutePathFileNameIsPopulated = false
		return []byte{}, errors.New(ePrefix + "Error: FileMgr.AbsolutePathFileName is EMPTY!")
	}


	if !fMgr.IsFilePtrOpen || fMgr.FilePtr == nil {

		if fMgr.FilePtr != nil {
			fMgr.CloseFile()
		}

		// If the path and file name do not exist, this method will
		// attempt to create said path and file name.
		err = fMgr.OpenThisFileReadWrite()

		if err != nil {
			return []byte{}, fmt.Errorf(ePrefix + "Error returned from fMgr.OpenThisFileReadWrite() fileNameExt='%v' Error='%v'", fMgr.AbsolutePathFileName, err.Error())
		}

		fMgr.IsFilePtrOpen = true
	}

	bytesRead, err := ioutil.ReadAll(fMgr.FilePtr)

	if err!=nil {
		return []byte{}, fmt.Errorf(ePrefix + "Error returned by ioutil.ReadAll(fMgr.FilePtr). FileName='%v' Errors='%v'", fMgr.AbsolutePathFileName, err.Error())
	}

	return bytesRead, nil
}

// ReadFileBytes - Reads bytes from the file identified by the current FileMgr
// object. Bytes are stored in 'byteBuff', a byte array passed in as an input
// parameter.
//
// If successful, the returned error value is 'nil'. The returned value 'int'
// contains the number of bytes read from the current file.
func (fMgr *FileMgr) ReadFileBytes(byteBuff []byte) (int, error) {

	ePrefix := "FileMgr.ReadFileBytes() "
	var err error

	if !fMgr.IsInitialized {
		return 0, errors.New(ePrefix + " Error: The File Manager data structure has NOT been initialized.")
	}

	err = fMgr.IsFileMgrValid("")

	if err != nil {
		return 0, fmt.Errorf(ePrefix + "Error - This File Manager is INVALID! Error='%v'", err.Error())
	}

	if !fMgr.AbsolutePathFileNameIsPopulated {
		return 0, errors.New(ePrefix + " Error: FileMgr.AbsolutePathFileName has NOT been initialized and populated.")
	}

	if fMgr.AbsolutePathFileName == "" {
		fMgr.AbsolutePathFileNameIsPopulated = false
		return 0, errors.New(ePrefix + " Error: FileMgr.AbsolutePathFileName is EMPTY!")
	}

	if !fMgr.IsFilePtrOpen || fMgr.FilePtr == nil {

		if fMgr.FilePtr != nil {
			fMgr.CloseFile()
		}

		// If the path and file name do not exist, this method will
		// attempt to create said path and file name.
		err = fMgr.OpenThisFileReadWrite()

		if err != nil {
			return 0, fmt.Errorf(ePrefix + "Error returned by fMgr.OpenThisFileReadWrite()  FileName='%v'  Error='%v'", fMgr.AbsolutePathFileName, err.Error())
		}

		fMgr.IsFilePtrOpen = true
	}

	bytesRead, err := fMgr.FilePtr.Read(byteBuff)

	if err!=nil {
		return bytesRead, fmt.Errorf(ePrefix + "Error returned by fMgr.FilePtr.Read(byteBuff). FileName='%v'  Error='%v'", fMgr.AbsolutePathFileName, err.Error())
	}

	return bytesRead, nil
}

// SetFileMgrFromDirMgrFileName - Sets the data fields of the current FileMgr object
// based on a DirMgr object and a File Name string which are passed as input parameters.
func (fMgr *FileMgr) SetFileMgrFromDirMgrFileName(dMgr DirMgr, fileNameExt string) (isEmpty bool, err error){
	ePrefix := "FileMgr.SetFileMgrFromDirMgrFileName() "
	isEmpty = true
	err = nil
	fMgr.Empty()
	fh:= FileHelper{}

	if dMgr.IsInitialized {

		err2 := dMgr.IsDirMgrValid("")

		if err2 != nil {
			err = fmt.Errorf(ePrefix + "Error: Input parameter 'dMgr' is INVALID! dMgr.Path='%v'  Error='%v'", dMgr.Path, err2.Error())
			return
		}

	}

	if len(fileNameExt) == 0 {
		err = errors.New(ePrefix + "Error: Input parameter 'fileNameExt' is a Zero length string!")
		return
	}

	adjustedFileNameExt, isFileNameEmpty, err2 := fh.CleanFileNameExtStr(fileNameExt)

	if err2 != nil {
		err = fmt.Errorf(ePrefix + "Error returned from fh.CleanFileNameExtStr(fileNameExt). fileNameExt='%v' Error='%v'", fileNameExt, err2.Error())
		return
	}

	if isFileNameEmpty {
		err = fmt.Errorf(ePrefix + "Error: fileName returned from fh.CleanFileNameExtStr(fileNameExt) is a ZERO length string! fileNameExt='%v'", fileNameExt)
		return
	}

	fMgr.DMgr = dMgr.CopyOut()

	s, fNameIsEmpty, err2 := fh.GetFileNameWithoutExt(adjustedFileNameExt)

	if err2!=nil{
		err = fmt.Errorf(ePrefix + "Error returned from fh.GetFileNameWithoutExt(adjustedFileNameExt). adjustedFileNameExt='%v'  Error='%v' ", adjustedFileNameExt, err2.Error())
		fMgr.Empty()
		isEmpty = true
		return
	}

	if fNameIsEmpty {
		err = fmt.Errorf(ePrefix + "Error: FileName returned from fh.GetFileNameWithoutExt(adjustedFileNameExt) is Zero length string! adjustedFileNameExt='%v'  ", adjustedFileNameExt)
		fMgr.Empty()
		isEmpty = true
		return

	}

	fMgr.FileNameIsPopulated = true
	fMgr.FileName = s

	s, extIsEmpty, err2 := fh.GetFileExtension(adjustedFileNameExt)

	if err2 != nil {
		err = fmt.Errorf(ePrefix + "Error returned from fh.GetFileExtension(fileNameAndExt). fileNameAndExt='%v'  Error='%v' ", adjustedFileNameExt, err2.Error())
		fMgr.Empty()
		isEmpty = true
		return
	}

	if !extIsEmpty {
		fMgr.FileExtIsPopulated = true
		fMgr.FileExt = s
	}

	if fMgr.FileNameIsPopulated {
		fMgr.FileNameExtIsPopulated = true
		fMgr.FileNameExt = fMgr.FileName + fMgr.FileExt
	}

	lPath := len(fMgr.DMgr.AbsolutePath)
	if lPath==0 {
		fMgr.AbsolutePathFileName = fMgr.FileNameExt

	} else if fMgr.DMgr.AbsolutePath[lPath-1] == os.PathSeparator {
		fMgr.AbsolutePathFileName = fMgr.DMgr.AbsolutePath + fMgr.FileNameExt

	} else {
		fMgr.AbsolutePathFileName = fMgr.DMgr.AbsolutePath + string(os.PathSeparator) + fMgr.FileNameExt

	}

	lPath = len(fMgr.DMgr.Path)

	if lPath==0 {
		fMgr.OriginalPathFileName = fMgr.FileNameExt

	} else if fMgr.DMgr.Path[lPath-1] == os.PathSeparator {
		fMgr.OriginalPathFileName = fMgr.DMgr.Path + fMgr.FileNameExt

	} else {
		fMgr.OriginalPathFileName = fMgr.DMgr.Path + string(os.PathSeparator) + fMgr.FileNameExt

	}

	fMgr.AbsolutePathFileNameIsPopulated = true

	fInfo, err2 := os.Stat(fMgr.AbsolutePathFileName)

	if err2 == nil {
		fMgr.AbsolutePathFileNameDoesExist = true
		fMgr.ActualFileInfo = FileInfoPlus{}.NewPathFileInfo(fMgr.DMgr.AbsolutePath, fInfo)
	} else {
		fMgr.AbsolutePathFileNameDoesExist = false
		fMgr.ActualFileInfo = FileInfoPlus{}
	}

	fMgr.IsInitialized = true

	err = nil
	isEmpty = false

	return

}

// SetFileMgrFromPathFileName - Initializes all the data fields of the
// current FileMgr object based on the path file name string passed to
// this method as an input parameter.
func (fMgr *FileMgr) SetFileMgrFromPathFileName(pathFileNameExt string) (isEmpty bool, err error) {

	ePrefix := "FileMgr.SetFileMgrFromPathFileName() "
	isEmpty = true
	err = nil
	fh:= FileHelper{}
	
	if len(pathFileNameExt) == 0 {
		err = errors.New(ePrefix + "Error: Input parameter 'pathFileNameExt' is a zero length or empty string!")
		return
	}

	adjustedPathFileNameExt := fh.AdjustPathSlash(pathFileNameExt)

	adjustedFileNameExt, isEmptyFileName, err2 := fh.CleanFileNameExtStr(adjustedPathFileNameExt)

	if err2 != nil {
		err = fmt.Errorf(ePrefix + "Error returned from fh.CleanFileNameExtStr(adjustedPathFileNameExt). adjustedPathFileNameExt='%v' Error='%v'", adjustedPathFileNameExt, err2.Error() )
		return
	}

	if isEmptyFileName {
		err = fmt.Errorf(ePrefix + "Error: File Name returned from fh.CleanFileNameExtStr(adjustedPathFileNameExt) is a Zero Length String!. pathFileNameExt='%v'", adjustedPathFileNameExt )
	}

	remainingPathStr := strings.TrimSuffix(adjustedPathFileNameExt, adjustedFileNameExt)

	var dMgr DirMgr

	if len(remainingPathStr) == 0 {
		dMgr = DirMgr{}
	} else {

		dMgr, err2 = DirMgr{}.New(remainingPathStr)

		if err2 != nil {
			err = fmt.Errorf(ePrefix + "Error returned from DirMgr{}.New(remainingPathStr). remainingPathStr='%v'  Error='%v'", remainingPathStr, err2.Error())
			return
		}

	}

	isEmptyFMgr, err2 := fMgr.SetFileMgrFromDirMgrFileName(dMgr, adjustedFileNameExt)

	if err2 != nil {
		err = fmt.Errorf(ePrefix + "Error returned from fMgr.SetFileMgrFromDirMgrFileName(dMgr, adjustedFileNameExt). adjustedFileNameExt='%v'  Error='%v'", adjustedFileNameExt, err2.Error())
		return
	}

	if isEmptyFMgr {
		err = fmt.Errorf(ePrefix + "Error: Empty FileMgr returned from fMgr.SetFileMgrFromDirMgrFileName(dMgr, adjustedFileNameExt). dMgr.Path='%v'   adjustedFileNameExt='%v' ", dMgr.Path, adjustedFileNameExt)
		return
	}

	isEmpty = false
	err = nil

	return
}


// SetFileInfo - Used to initilize the os.FileInfo structure maintained as
// part of the current FileMgr object.
func (fMgr *FileMgr) SetFileInfo(info os.FileInfo) error {
	ePrefix := "FileMgr.SetFileInfo() "

	if info.Name() == "" {
		return errors.New(ePrefix + "Error: info.Name() is an EMPTY string!")
	}

	if info.IsDir() {
		return errors.New(ePrefix + "info.IsDir()=='true'. This is a Directory NOT A FILE!")
	}

	fMgr.ActualFileInfo = FileInfoPlus{}.NewFromFileInfo(info)

	if !fMgr.ActualFileInfo.IsFInfoInitialized {
		return fmt.Errorf(ePrefix + "Error: Failed to initialize fMgr.ActualFileInfo object. info.Name()='%v'", info.Name())
	}

	return nil
}

// WriteStrToFile - Writes a string to the File identified by
// FileMgr.AbsolutePathFileName. If the file is not open, this
// method will attempt to open it.
func (fMgr *FileMgr) WriteStrToFile(str string) (int, error) {

	ePrefix := "FileMgr.WriteStrToFile() "
	var err error

	if !fMgr.IsInitialized {
		return 0, errors.New(ePrefix + "Error: The File Manager data structure has NOT been initialized.")
	}

	err = fMgr.IsFileMgrValid("")

	if err != nil {
		return 0, fmt.Errorf(ePrefix + "Error: This File Manger is INVALID! fileNameExt='%v'  Error='%v'", fMgr.AbsolutePathFileName, err.Error())
	}

	if !fMgr.AbsolutePathFileNameIsPopulated {
		return 0, errors.New(ePrefix + "Error: FileMgr.AbsolutePathFileName has NOT been initialized and populated.")
	}

	if fMgr.AbsolutePathFileName == "" {
		fMgr.AbsolutePathFileNameIsPopulated = false
		return 0, errors.New(ePrefix + "Error: FileMgr.AbsolutePathFileName is EMPTY!")
	}

	if !fMgr.IsFilePtrOpen || fMgr.FilePtr == nil {

		if fMgr.FilePtr != nil {
			fMgr.CloseFile()
		}

		// If the path and file name do not exist, this method will
		// attempt to create said path and file name.
		err = fMgr.OpenThisFileReadWrite()

		if err != nil {
			return 0, fmt.Errorf(ePrefix + " - fMgr.OpenThisFileReadWrite() returned errors: %v", err.Error())
		}

		fMgr.IsFilePtrOpen = true
	}

	bytesWritten, err := fMgr.FilePtr.WriteString(str)

	if err!= nil {
		return bytesWritten, fmt.Errorf("Error returned from fMgr.FilePtr.WriteString(str). Output File='%v'. Error='%v'", fMgr.AbsolutePathFileName, err.Error())
	}

	return bytesWritten, nil
}

