package common

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path"
	fp "path/filepath"
	"strings"
	"time"
)


/*
		'filehelper.go' - Contains type 'FileHelper' and related data structures.

		The Source Repository for this source code file is :
			https://github.com/MikeAustin71/pathfilego.git

		'FileHelper' is dependent of 'DirMgr' and 'FileMgr'.  'DirMgr' and 'FileMgr'
		are located in source file 'fileanddirmanagers.go' found in this same
	 	directory: '003_filehelper/common/fileanddirmanagers.go'


 */

// FileInfoPlus - Conforms to the os.FileInfo interface. This structure will store
// FileInfo information plus additional information related to a file or directory.
type FileInfoPlus struct {
	IsFInfoInitialized	bool        // Not part of FileInfo interface.
																	// 'true' = structure fields have been properly initialized
	IsDirPathInitialized	bool			// Not part of FileInfo interface.
																	// 'true' = structure field 'dirPath' has been successfully initialized
	CreateTimeStamp    	time.Time   // Not part of FileInfo interface.
																	// Date time at which this instance was initialized
	dirPath							string			// Not part of FileInfo interface. Directory path associated with file name
	fName              	string      // FileInfo.Name() base name of the file
	fSize              	int64       // FileInfo.Size() length in bytes for regular files; system-dependent for others
	fMode              	os.FileMode // FileInfo.Mode() file mode bits
	fModTime           	time.Time   // FileInfo.ModTime() file modification time
	isDir              	bool        // FileInfo.IsDir() 'true'= this is a directory not a file
	dataSrc							interface{} // FileInfo.Sys() underlying data source (can return nil)
}

// Name - base name of the file
func(fip FileInfoPlus) Name() string {

	return fip.fName
}

//Size - file length in bytes for regular files; system-dependent for others
func (fip FileInfoPlus) Size() int64 {
	return fip.fSize
}

// Mode - file mode bits. See os.FileMode
// A FileMode represents a file's mode and permission bits.
// The bits have the same definition on all systems, so that
// information about files can be moved from one system
// to another portably. Not all bits apply to all systems.
// The only required bit is ModeDir for directories.
//
// type FileMode uint32
//
// The defined file mode bits are the most significant bits of the FileMode.
// The nine least-significant bits are the standard Unix rwxrwxrwx permissions.
// The values of these bits should be considered part of the public API and
// may be used in wire protocols or disk representations: they must not be
// changed, although new bits might be added.
// const (
//  // The single letters are the abbreviations
//  // used by the String method's formatting.
// 	ModeDir        FileMode = 1 << (32 - 1 - iota) // d: is a directory
// 	ModeAppend                                     // a: append-only
// 	ModeExclusive                                  // l: exclusive use
// 	ModeTemporary                                  // T: temporary file; Plan 9 only
// 	ModeSymlink                                    // L: symbolic link
// 	ModeDevice                                     // D: device file
// 	ModeNamedPipe                                  // p: named pipe (FIFO)
// 	ModeSocket                                     // S: Unix domain socket
// 	ModeSetuid                                     // u: setuid
// 	ModeSetgid                                     // g: setgid
// 	ModeCharDevice                                 // c: Unix character device, when ModeDevice is set
// 	ModeSticky                                     // t: sticky
//
// 	// Mask for the type bits. For regular files, none will be set.
// 	ModeType = ModeDir | ModeSymlink | ModeNamedPipe | ModeSocket | ModeDevice
//
// 	ModePerm FileMode = 0777 // Unix permission bits
// )
//
func (fip FileInfoPlus) Mode() os.FileMode {
	return fip.fMode
}

// ModTime - file modification time
func (fip FileInfoPlus) ModTime() time.Time {
	return fip.fModTime
}

// IsDir - 'true' = this is a directory,
// not a file.
//
// abbreviation for Mode().IsDir()
//
func (fip FileInfoPlus) IsDir() bool {
	return fip.isDir
}

// Sys - underlying data source (can return nil)
func (fip FileInfoPlus) Sys() interface{} {
	return fip.dataSrc
}

// CopyOut - Creates a copy of the current FileInfoPlus
// instance and returns it.
func (fip *FileInfoPlus) CopyOut() FileInfoPlus{
	newInfo := FileInfoPlus{}

	newInfo.SetName(fip.Name())
	newInfo.SetSize(fip.Size())
	newInfo.SetMode(fip.Mode())
	newInfo.SetModTime(fip.ModTime())
	newInfo.SetIsDir(fip.IsDir())
	newInfo.SetSysDataSrc(fip.Sys())
	newInfo.SetDirectoryPath(fip.DirPath())
	newInfo.IsFInfoInitialized = fip.IsFInfoInitialized
	newInfo.IsDirPathInitialized = fip.IsDirPathInitialized
	newInfo.CreateTimeStamp = fip.CreateTimeStamp
	return newInfo
}

// DirPath - Returns the directory path. This field, FileInfoPlus.dirPath,
// is not part of the standard FileInfo interface.
func (fip FileInfoPlus) DirPath() string {
	return fip.dirPath
}

// Equal - Compares two FileInfoPlus objects to determine
// if they are equal.
func (fip *FileInfoPlus) Equal(fip2 *FileInfoPlus) bool {

	if fip.Name() 	!= fip2.Name()			||
		fip.Size() 		!= fip2.Size()			||
		fip.Mode() 		!= fip2.Mode()			||
		fip.ModTime() != fip2.ModTime() 	||
		fip.IsDir()		!= fip2.IsDir()			||
		fip.Sys()			!= fip2.Sys()				||
		fip.DirPath() != fip2.DirPath()		{
			return false
	}

	return true

}

// NewFromFileInfo - Creates and returns a new FileInfoPlus object
// populated with FileInfo data received from the input parameter.
// Notice that this version of the 'New' method does NOT set the
// Directory Path. This method is NOT part of the FileInfo interface.
//
// Example Usage:
//	fip := FileInfoPlus{}.NewFromFileInfo(info)
//  -- fip is now a newly populated FileInfoPlus instance.
//
func (fip FileInfoPlus) NewFromFileInfo(info os.FileInfo) FileInfoPlus {
	newInfo := FileInfoPlus{}

	newInfo.SetName(info.Name())
	newInfo.SetSize(info.Size())
	newInfo.SetMode(info.Mode())
	newInfo.SetModTime(info.ModTime())
	newInfo.SetIsDir(info.IsDir())
	newInfo.SetSysDataSrc(info.Sys())
	newInfo.SetIsFInfoInitialized(true)
	return newInfo
}

// NewPathFileInfo - Creates and returns a new FileInfoPlus object
// populated with directory path and FileInfo data received from
// the input parameters.
//
// Example Usage:
//	fip := FileInfoPlus{}.NewPathFileInfo(dirPath, info)
//  -- fip is now a newly populated FileInfoPlus instance.
//
func (fip FileInfoPlus) NewPathFileInfo(dirPath string, info os.FileInfo) FileInfoPlus {
	newInfo := FileInfoPlus{}.NewFromFileInfo(info)
	newInfo.SetDirectoryPath(dirPath)
	return newInfo
}

// SetDirectoryPath - Sets the dirPath field. This
// field is not part of the standard FileInfo data structure.
func (fip *FileInfoPlus) SetDirectoryPath(dirPath string) error {
	fh := FileHelper{}
	dirPath = strings.TrimLeft(strings.TrimRight(dirPath, " "), " ")

	if len(dirPath) == 0 {
		return fmt.Errorf("FileInfoPlus.SetDirectoryPath() Error: 'dirPath' is a Zero Length String!")
	}

	dirPath =  fh.RemovePathSeparatorFromEndOfPathString(dirPath)
	fip.dirPath = dirPath
	fip.IsDirPathInitialized = true
	return nil
}

// SetName - Sets the file name field.
func (fip *FileInfoPlus) SetName(name string) {
	fip.fName = name
}

// SetSize - Sets the file size field
func (fip *FileInfoPlus) SetSize(fileSize int64) {
	fip.fSize = fileSize
}

// SetMode - Sets the file Mode
func (fip *FileInfoPlus) SetMode(fileMode os.FileMode) {
	fip.fMode = fileMode
}

// SetModTime - Sets the file modification time
func (fip *FileInfoPlus) SetModTime(fileModTime time.Time) {
	fip.fModTime = fileModTime
}

// SetIsDir - Sets is directory field.
func (fip *FileInfoPlus) SetIsDir(isDir bool) {
	fip.isDir = isDir
}

// SetSysDataSrc - Sets the dataSrc field
func (fip *FileInfoPlus) SetSysDataSrc(sysDataSrc interface{}) {
	fip.dataSrc = sysDataSrc
}

// SetIsFInfoInitialized - Sets the flag for 'Is File Info Initialized'
// If set to 'true' it means that all of the File Info fields have
// been initialized.
func (fip *FileInfoPlus) SetIsFInfoInitialized(isInitialized bool) {
	if !isInitialized {
		fip.IsFInfoInitialized = false
		fip.CreateTimeStamp = time.Time{}
		return
	}

	fip.IsFInfoInitialized = true
	fip.CreateTimeStamp = time.Now().Local()
	return
}

// FileSelectCriterionMode - Used in conjunction with the
// FileSelectionCriteria structure, below
type FileSelectCriterionMode int

// String - Method used to display the text
// name of an Operations Message Type.
func (fSelectMode FileSelectCriterionMode) String() string {
	return FileSelectCriterionModeNames[fSelectMode]
}

const (

	// ANDFILESELECTCRITERION - 0 File Selection Criterion are And'ed
	// together. If there are three file selection criterion then
	// all three must be satisfied before a file is selected.
	ANDFILESELECTCRITERION FileSelectCriterionMode = iota

	// ORFILESELECTCRITERION - 1 File Selection Criterion are Or'd together.
	// If there are three file selection criterion then satisfying any
	// one of the three criterion will cause the file to be selected.
	ORFILESELECTCRITERION

)

// FileSelectCriterionModeNames - String Array holding File Select Criteria  names.
var FileSelectCriterionModeNames = [...]string{"AND File Select Criterion","OR File Select Criterion"}


// FileSelectionCriteria - Used is selecting file names. These
// data fields specify the crierion used to determine if a
// file should be selected for some type of operation.
// Example: find files or delete files operations
type FileSelectionCriteria struct {
	FileNamePatterns						[]string		// a string array containing one or more file matching
																					// patterns. Example '*.txt' '*.log' 'common*.*'
	FilesOlderThan							time.Time 	// Used to select files with a modification less than this date time
	FilesNewerThan							time.Time 	// Used to select files with a modification greater than this date time
	SelectByFileMode						os.FileMode // Used to select files with equivalent FileMode values
																					//   Note: os.FileMode is an uint32 type
	SelectCriterionMode					FileSelectCriterionMode // Can be one of two values:
																											// ANDFILESELECTCRITERION or ORFILESELECTCRITERION
																											//
																											// ANDFILESELECTCRITERION = select a file only if ALL
																											//										      the selection crietrion
																											//                          are satisfied.
																											//
																											// ORFILESELECTCRITERION  = select a file if only ONE
																											//													of the selection crierion
																											//													are satisfied.
}

// ArePatternsActive - surveys the FileNamePatterns string
// array to determine if there currently any active search
// file pattern string.
//
// A search file pattern is considered active if the string
// length of the pattern string is greater than zero.
func (fsc *FileSelectionCriteria) ArePatternsActive() bool {

	lPats := len(fsc.FileNamePatterns)

	if lPats == 0 {
		return false
	}

	isActive := false

	for i:=0; i < lPats; i++ {
		fsc.FileNamePatterns[i] = strings.TrimRight(strings.TrimLeft(fsc.FileNamePatterns[i], " "), " ")
		if fsc.FileNamePatterns[i] != "" {
			isActive = true
		}

	}

	return isActive
}


// DirectoryTreeInfo - structure used
// to 'Find' files in a directory specified
// by 'StartPath'. The file search will be
// filtered by a 'FileSelectCriteria' object.
//
// 'FileSelectCriteria' is a FileSelectionCriteria type
// which contains FileNamePatterns strings and
//'FilesOlderThan' or 'FilesNewerThan' date time
// parameters which can be used as a selection
// criteria.
//
type DirectoryTreeInfo struct {
	StartPath            	string
	Directories          	DirMgrCollection
	FoundFiles           	FileMgrCollection
	ErrReturns           	[]string
	FileSelectCriteria    FileSelectionCriteria
}

// CopyToDirectoryTree - Copies an entire directory tree to an alternate location.
// The copy operation includes all files and all directories in the designated directory
// tree.
func (dirTree *DirectoryTreeInfo) CopyToDirectoryTree(baseDir, newBaseDir DirMgr) (DirectoryTreeInfo, error) {

	ePrefix := "DirectoryTreeInfo.CopyToDirectoryTree() "

	newDirTree := DirectoryTreeInfo{}

	if !baseDir.IsInitialized {
		return newDirTree, errors.New(ePrefix + "Error: Input parameter 'baseDir' is NOT initialized. It is EMPTY!")
	}

	err2 := baseDir.IsDirMgrValid("")

	if err2 != nil {
		return newDirTree, fmt.Errorf(ePrefix + "Error: Input Parameter 'baseDir' is INVALID! Error='%v'", err2.Error())
	}

	if !newBaseDir.IsInitialized {
		return newDirTree, errors.New(ePrefix + "Error: Input parameter 'newBaseDir' is NOT initialized. It is EMPTY!")

	}

	err2 = newBaseDir.IsDirMgrValid("")

	if err2 != nil {
		return newDirTree, fmt.Errorf(ePrefix + "Error: Input Parameter 'newBaseDir' is INVALID! Error='%v'", err2.Error())
	}

	err2 = newBaseDir.MakeDir()

	if err2 != nil {
		return newDirTree, fmt.Errorf(ePrefix + "Error returned from  newBaseDir.MakeDir(). newBaseDir.AbsolutePath='%v'  Error='%v'", newBaseDir.AbsolutePath, err2.Error())
	}

	lAry := len(dirTree.Directories.DirMgrs)

	// Make the new Directory Tree
	for i:=0 ; i < lAry; i++ {

		newDMgr, err2 := dirTree.Directories.DirMgrs[i].SubstituteBaseDir(baseDir, newBaseDir)

		if err2 != nil {
			return DirectoryTreeInfo{}, fmt.Errorf(ePrefix + "Error returned from SubstituteBaseDir(baseDir, newBaseDir). i='%v' Error='%v'", i, err2.Error())
		}

		err2 = newDMgr.MakeDir()

		if err2 != nil {
			return DirectoryTreeInfo{},  fmt.Errorf(ePrefix + "Error returned fromnewDMgr.MakeDir()  Error='%v'", err2.Error())

		}

		newDirTree.Directories.AddDirMgr(newDMgr)

	}

	lAry = len(dirTree.FoundFiles.FMgrs)

	for j:=0; j < lAry; j++ {

		fileDMgr, err2 := dirTree.FoundFiles.FMgrs[j].DMgr.SubstituteBaseDir(baseDir, newBaseDir)

		if err2 != nil {
			return DirectoryTreeInfo{}, fmt.Errorf(ePrefix + "Error returned by dirTree.FoundFiles.FMgrs[j].DMgr.SubstituteBaseDir(baseDir, newBaseDir). Error='%v'", err2.Error())
		}


		newFileMgr, err2 := FileMgr{}.NewFromDirMgrFileNameExt(fileDMgr, dirTree.FoundFiles.FMgrs[j].FileNameExt)

		if err2 != nil {
			return DirectoryTreeInfo{}, fmt.Errorf(ePrefix + "Error returned by FileMgr{}.NewFromDirMgrFileNameExt(dMgr, dirTree.FoundFiles.FMgrs[j].FileNameExt) dirTree.FoundFiles.FMgrs[j].FileNameExt='%v' j='%v' Error='%v'", dirTree.FoundFiles.FMgrs[j].FileNameExt, j, err2.Error())
		}

		err2 = dirTree.FoundFiles.FMgrs[j].CopyFileMgr(&newFileMgr)

		if err2 != nil {
			return DirectoryTreeInfo{}, fmt.Errorf(ePrefix + "Error returned by FMgrs[j].CopyFileMgr(&newFileMgr) SrcFileName:'%v'  DestFileName:'%v' Error='%v'", dirTree.FoundFiles.FMgrs[j].FileNameExt, newFileMgr.FileNameExt, err2.Error())

		}

		newDirTree.FoundFiles.AddFileMgr(newFileMgr)
	}

	return newDirTree, nil
}

// DirectoryDeleteFileInfo - structure used
// to delete files in a directory specified
// by 'StartPath'. Deleted files will be selected
// based on 'DeleteFileSelectCriteria' value.
//
// 'DeleteFileSelectCriteria' is a 'FileSelectionCriteria'
// type which contains  FileNamePatterns strings and the
// FilesOlderThan or FilesNewerThan date time parameters
// which can be used as file selection criteria.
type DirectoryDeleteFileInfo struct {
	StartPath            	string
	Directories          	DirMgrCollection
	ErrReturns           	[]string
	DeleteFileSelectCriteria FileSelectionCriteria
	DeletedFiles         	FileMgrCollection
}


// FileHelper - structure used
// to encapsulate FileHelper utility
// methods.
type FileHelper struct {
	Input		string
	Output	string
}

// AddPathSeparatorToEndOfPathStr - Receives a path string as an input
// parameter. If the last character of the path string is not a path
// separator, this method will add a path separator to the end of that
// path string and return it to the calling method.
func (fh FileHelper) AddPathSeparatorToEndOfPathStr(pathStr string) (string, error) {
	lStr := len(pathStr)

	ePrefix := "FileHelper.AddPathSeparatorToEndOfPathStr() "

	if lStr == 0 {
		return "", errors.New(ePrefix + "Error: Zero length input parameter, 'pathStr'!")
	}

	if pathStr[lStr-1] == os.PathSeparator || pathStr[lStr-1] == '/' || pathStr[lStr-1] == '\\'  {
		return pathStr, nil
	}

	newPathStr := pathStr + string(os.PathSeparator)

	return newPathStr, nil
}

// AdjustPathSlash standardize path
// separators according to operating system
func (fh FileHelper) AdjustPathSlash(path string) string {

	return fp.FromSlash(path)
}

// ChangeDir - Chdir changes the current working directory to the named directory. If there is an error, it will be of type *PathError.
func (fh FileHelper) ChangeDir(dirPath string) error {

	err := os.Chdir(dirPath)

	if err != nil {
		return err
	}

	return nil
}

// CopyFileByLink - Copies a file from source to destination
// by means of creating a 'hard link' to the source file.
// See: https://stackoverflow.com/questions/21060945/simple-way-to-copy-a-file-in-golang
//
// Note: This method of copying files does not create a new
// destination file and write the contents of the source file
// to destination file. (See CopyToNewFile Below).  Instead, this
// method performs the copy operation by creating a hard symbolic
// link to the source file.
//
func (fh FileHelper) CopyFileByLink(src, dst string) (err error) {

	ePrefix := "FileHelper.CopyFileByLink() "

	if len(src) == 0 {
		err = errors.New(ePrefix + "Error: Input parameter 'src' is ZERO length string!")
	}

	if len(dst) == 0 {
		err = errors.New(ePrefix + "Error: Input parameter 'dst' is a ZERO length string!")
		return
	}

	if !fh.DoesFileExist(src) {
		err = fmt.Errorf(ePrefix + "Error: Input parameter 'src' file DOES NOT EXIST! src='%v'", src)
		return
	}

	sfi, err2 := os.Stat(src)

	if err2 != nil {
		err = fmt.Errorf(ePrefix + "Error returned from os.Stat(src). src='%v'  Error='%v'", src, err2.Error())
		return
	}

	if !sfi.Mode().IsRegular() {
		// cannot copy non-regular files (e.g., directories,
		// symlinks, devices, etc.)
		err = fmt.Errorf(ePrefix + "Error: non-regular source file. Source File Name='%v'  Source File Mode='%v' ", sfi.Name(), sfi.Mode().String())
		return
	}

	dfi, err2 := os.Stat(dst)

	if err2 != nil {

		if !os.IsNotExist(err2) {
			// Must be PathError - Path does not exist
			err = fmt.Errorf(ePrefix + "Destination File Path Error - Path does NOT exist. Destination File='%v' Error: %v", dst,  err2.Error())
			return
		}

	} else {

		if !(dfi.Mode().IsRegular()) {
			err = fmt.Errorf(ePrefix + "non-regular destination file - Cannot Overwrite destination file. Destination File='%v'  Destination File Mode= '%v'", dfi.Name(), dfi.Mode().String())
			return
		}

		// Source and destination have the same path
		// and file names. They are one in the same
		// file. Nothing to do.
		if os.SameFile(sfi, dfi) {
			err = nil
			return
		}

	}

	err2 = os.Link(src, dst)

	if err2 != nil {
		err = fmt.Errorf(ePrefix + "- os.Link(src, dst) FAILED! src='%v' dst='%v'  Error='%v'", src, dst, err2.Error())
		return
	}

	err = nil

	return
}

// CopyToNewFile - Copies file from source path & File Name
// to destination path & File Name.
// See: https://stackoverflow.com/questions/21060945/simple-way-to-copy-a-file-in-golang
//
// Note: Unlike the method CopyFileByLink above, this method
// does NOT rely on the creation of symbolic links. Instead,
// a new destination file is created and the contents of the source
// file are written to the new destination file.
//
func (fh FileHelper) CopyToNewFile(src, dst string) (err error) {
	ePrefix := "FileHelper.CopyToNewFile() "
	err = nil

	if len(src) == 0 {
		err = errors.New(ePrefix + "Error: Input parameter 'src' is a ZERO length string!")
		return
	}

	if len(dst) == 0 {
		err = errors.New(ePrefix + "Error: Input parameter 'dst' is a ZERO length string!")
		return
	}

	if !fh.DoesFileExist(src) {
		err = fmt.Errorf(ePrefix + "Error: Input parameter 'src' file DOES NOT EXIST! src='%v'", src)
		return
	}

	sfi, err2 := os.Stat(src)

	if err2 != nil {
		err = fmt.Errorf(ePrefix + "Error returned from os.Stat(src). src='%v'  Error='%v'", src, err2.Error())
		return
	}

	if !sfi.Mode().IsRegular() {
		// cannot copy non-regular files (e.g., directories,
		// symlinks, devices, etc.)
		err = fmt.Errorf(ePrefix + "Error non-regular source file ='%v' source file Mode='%v'", sfi.Name(), sfi.Mode().String())
		return
	}

	dfi, err2 := os.Stat(dst)

	if err2 != nil {

		if !os.IsNotExist(err2) {
			// Must be PathError - Path does not exist
			err = fmt.Errorf(ePrefix + "Destination File Path Error - Path does NOT exist. Destination File='%v' Error: %v", dst, err.Error())
			return
		}

	} else {

		if !dfi.Mode().IsRegular() {
			err = fmt.Errorf(ePrefix + "Error: non-regular destination file. Cannot Overwrite destination file. Destination file='%v' destination file mode='%v'", dfi.Name(), dfi.Mode().String())
			return
		}
		if os.SameFile(sfi, dfi) {
			// Source and destination are the same
			// path and file name.
			err = nil
			return
		}

	}

	// Create a new destination file and copy source
	// file contents to the destination file.
	err = fh.CopyFileContents(src, dst)
	return
}

// CopyFileContents - Copies file contents from source to destination file.
// Note: If 'src' file does NOT exist, an error will be returned.
//
// No validity checks are performed on 'dest' file.
//
// This method is called by FileHelper:CopyFileStr(). Use FileHelper:CopyFileStr() for
// ordinary file copy operations since it provides validity checks on 'src' and 'dest'
// files.
//
// Reference: https://stackoverflow.com/questions/21060945/simple-way-to-copy-a-file-in-golang
//
func (fh FileHelper) CopyFileContents(src, dst string) (err error) {
	ePrefix := "FileHelper.CopyFileContents() "

	err = nil

	if len(src) == 0 {
		err = errors.New(ePrefix + "Error: Input parameter 'src' is a ZERO length string!")
		return
	}

	if len(dst) == 0 {
		err = errors.New(ePrefix + "Error: Input parameter 'dst' is a ZERO length string!")
		return
	}

	if !fh.DoesFileExist(src) {
		err = fmt.Errorf(ePrefix + "Error: Source file does NOT exist! src='%v'", src)
		return
	}

	in, err2 := os.Open(src)

	if err2 != nil {
		err = fmt.Errorf(ePrefix + "Error returned from os.Open(src) src='%v'  Error='%v'", src, err2.Error())
		return
	}

	defer in.Close()

	out, err2 := os.Create(dst)

	if err2 != nil {
		err = fmt.Errorf(ePrefix + "Error returned from os.Create(dst) dst='%v'  Error='%v'", dst, err2.Error())
		return
	}

	defer func() {
		cerr := out.Close()
		if cerr != nil {
			err = fmt.Errorf(ePrefix + "Error returned by defer func() out.Close() out=dst='%v'  Error='%v'",dst, cerr.Error())
			return
		}
	}()

	if _, err2 = io.Copy(out, in); err2 != nil {
		err = fmt.Errorf(ePrefix + "Error returned from io.Copy(dst, src) dst='%v'  src='%v'  Error='%v' ", dst, src, err2.Error())
		return
	}

	// flush file buffers in memory
	err2 = out.Sync()

	if err2 != nil {
		err = fmt.Errorf(ePrefix + "Error returned from out.Sync() out=dst='%v' Error='%v'", dst, err2.Error())
		return
	}

	err = nil

	return
}

// CleanPathStr - Wrapper Function for filepath.Clean()
// See: https://golang.org/pkg/path/filepath/#Clean
// Clean returns the shortest path name equivalent to path
// by purely lexical processing. It applies the following rules
// iteratively until no further processing can be done:
// 1. Replace multiple Separator elements with a single one.
// 2. Eliminate each . path name element (the current directory).
// 3. Eliminate each inner .. path name element (the parent directory)
// 		along with the non-.. element that precedes it.
// 4. Eliminate .. elements that begin a rooted path:
// 		that is, replace "/.." by "/" at the beginning of a path,
// 		assuming Separator is '/'.'
// The returned path ends in a slash only if it represents a root
// directory, such as "/" on Unix or `C:\` on Windows.
// Finally, any occurrences of slash are replaced by Separator.
// If the result of this process is an empty string,
// Clean returns the string ".".

func (fh FileHelper) CleanPathStr(pathStr string) string {

	return fp.Clean(pathStr)
}


// CleanDirStr - Cleans and formats a directory string.
//
// Example:
// dirName = '../dir1/dir2/fileName.ext' returns "../dir1/dir2"
// dirName = 'fileName.ext' returns "" isEmpty = true
// dirName = '../dir1/dir2/' returns '../dir1/dir2'
// dirName = '../dir1/dir2/filename.ext' returns '../dir1/dir2'
//
func (fh FileHelper) CleanDirStr(dirNameStr string) (dirName string, isEmpty bool, err error) {

	ePrefix := "FileHelper.CleanDirStr() "
	dirName = ""
	isEmpty = true
	err = nil

	if len(dirNameStr) == 0 {
		err = errors.New(ePrefix + "Error: Input parameter 'dirNameStr' is a Zero Length String!")
		return
	}

	dirNameStr = strings.TrimLeft( strings.TrimRight(dirNameStr, " "), " " )

	if len(dirNameStr) == 0 {
		err = errors.New(ePrefix + "Error: After trimming white space, input parameter 'dirNameStr' is a Zero length string!")
		return
	}

	adjustedDirName := fh.AdjustPathSlash(dirNameStr)

	lAdjustedDirName := len(adjustedDirName)

	if lAdjustedDirName == 0 {
		err = errors.New(ePrefix + "Error: After adjusting for path separators, input parameter 'dirNameStr' is an empty string!")
		return
	}

	if strings.Contains(adjustedDirName, "...") {
		err = fmt.Errorf(ePrefix + "Error: Invalid Directory string. Contains invalid dots. adjustedDirName='%v' ", adjustedDirName)
		return
	}

	volName := fp.VolumeName(adjustedDirName)

	if volName == adjustedDirName {
		dirName = adjustedDirName
		isEmpty = false
		err = nil
		return
	}

	// Find out if the directory path
	// actually exists.
	fInfo, err2 := os.Stat(adjustedDirName)

	if err2==nil {
		// The path exists

		if fInfo.IsDir() {
			// The path exists and it is a directory
			if adjustedDirName[lAdjustedDirName - 1] == os.PathSeparator {
				dirName = adjustedDirName[0:lAdjustedDirName - 1]
			} else {
				dirName = adjustedDirName
			}

			if len(dirName) == 0 {
				isEmpty = true
			} else {
				isEmpty = false
			}

			err = nil
			return

		} else {
			// The Path exists but it is
			// a File Name and NOT a directory name.
			adjustedDirName = strings.TrimSuffix(adjustedDirName, fInfo.Name())
			lAdjustedDirName = len(adjustedDirName)

			if lAdjustedDirName < 1 {
				dirName = ""
				isEmpty = true
				err = nil
				return
			}

			if adjustedDirName[lAdjustedDirName-1] == os.PathSeparator {
				dirName = adjustedDirName[0:lAdjustedDirName-1]
			} else {

				dirName = adjustedDirName
			}

			if len(dirName) == 0 {
				isEmpty = true
			} else {
				isEmpty = false
			}

			err = nil
			return
		}
	}

	firstCharIdx, lastCharIdx, err2 := fh.GetFirstLastNonSeparatorCharIndexInPathStr(adjustedDirName)

	if err2 != nil {
		err = fmt.Errorf(ePrefix + "Error returned by fh.GetFirstLastNonSeparatorCharIndexInPathStr(adjustedDirName). adjustedDirName='%v'  Error='%v'", adjustedDirName, err2.Error())
		return
	}

	if firstCharIdx == -1 || lastCharIdx == -1 {
		if adjustedDirName[lAdjustedDirName-1] == os.PathSeparator {
			dirName = adjustedDirName[0:lAdjustedDirName-1]
		} else {
			dirName = adjustedDirName
		}

		isEmpty = false
		err = nil
		return
	}

	interiorDotPathIdx := strings.LastIndex(adjustedDirName, "." + string(os.PathSeparator))

	if interiorDotPathIdx > firstCharIdx {
		err = fmt.Errorf(ePrefix + "Error: INVALID PATH. Invalid interior relative path detected! adjustedDirName='%v'", adjustedDirName)
		return
	}

	slashIdxs, err2 := fh.GetPathSeparatorIndexesInPathStr(adjustedDirName)

	if err2 != nil {
		err = fmt.Errorf("Error returned by fh.GetPathSeparatorIndexesInPathStr(adjustedDirName). adjusteDirName='%v'  Error='%v'", adjustedDirName, err2.Error())
		return
	}

	lSlashIdxs := len(slashIdxs)

	if lSlashIdxs == 0 {
		dirName = adjustedDirName
		isEmpty = false
		err = nil
		return
	}

	dotIdxs, err2 := fh.GetDotSeparatorIndexesInPathStr(adjustedDirName)

	if err2 != nil {
		err = fmt.Errorf("Error returned by fh.GetDotSeparatorIndexesInPathStr(adjustedDirName). adjustedDirName='%v'  Error='%v'", adjustedDirName, err2.Error())
		return
	}

	lDotIdxs := len(dotIdxs)


	// If a path separator is the last character
	if slashIdxs[lSlashIdxs-1] == lAdjustedDirName - 1 {
			dirName = adjustedDirName[0: slashIdxs[lSlashIdxs-1]]
			if len(dirName) == 0 {
				isEmpty = true
			} else {
				isEmpty = false
			}

			err = nil
			return
	}

	// If there is a dot after the last path separator,
	// this is a filename extension, NOT a directory
	if lDotIdxs > 0 && dotIdxs[lDotIdxs-1] > slashIdxs[lSlashIdxs-1] {

		dirName = adjustedDirName[0: slashIdxs[lSlashIdxs-1]]

		if len(dirName) == 0 {
			isEmpty = true
		} else {
			isEmpty = false
		}

		err = nil
		return
	}

	dirName = adjustedDirName
	isEmpty = false
	err = nil
	return
}

// CleanFileNameExtStr - Cleans up a file name extension string.
//
// Example:
// fileNameExt = '../dir1/dir2/fileName.ext' returns "fileName.ext"
// fileNameExt = 'fileName.ext" returns "fileName.ext"
// fileNameExt = '../dir1/dir2/' returns "" and isEmpty=true
//
func (fh FileHelper) CleanFileNameExtStr(fileNameExtStr string) (fileNameExt string, isEmpty bool, err error) {

	ePrefix := "FileHelper.CleanFileNameExtStr() "
	fileNameExt = ""
	isEmpty = true
	err = nil

	if len(fileNameExtStr) == 0 {
		err = errors.New(ePrefix + "Error: Input parameter fileNameExtStr is a Zero Length String!")
		return
	}

	adjustedFileNameExt := fh.AdjustPathSlash(fileNameExtStr)

	if strings.Contains(adjustedFileNameExt, "...") {
		err = fmt.Errorf(ePrefix + "Error: Invalid Directory string. Contains invalid dots. adjustedFileNameExt='%v' ", adjustedFileNameExt)
		return
	}

	// Find out if the file name extension path
	// actually exists.
	fInfo, err2 := os.Stat(adjustedFileNameExt)

	if err2==nil {
		// The path exists

		if fInfo.IsDir() {
			// The path exists and it is a directory.
			// There is no File Name present.
			fileNameExt = ""
			isEmpty = true
			err = fmt.Errorf(ePrefix + "Error: adjustedFileNameExt exists as a 'Directory' - NOT A FILE NAME! adjustedFileNameExt='%v'", adjustedFileNameExt)
			return
		} else {
			// The Path exists and it is a valid
			// file name.
			fileNameExt = fInfo.Name()
			isEmpty = false

			err = nil
			return
		}
	}

	firstCharIdx, lastCharIdx, err := fh.GetFirstLastNonSeparatorCharIndexInPathStr(adjustedFileNameExt)

	if firstCharIdx == -1 || lastCharIdx == -1 {
		err = fmt.Errorf(ePrefix + "File Name Extension string contains no valid file name characters! adjustedFileNameExt='%v'", adjustedFileNameExt)
		return
	}


	// The file name extension path does not exist

	interiorDotPathIdx := strings.LastIndex(adjustedFileNameExt, "." + string(os.PathSeparator))

	if interiorDotPathIdx > firstCharIdx {
		err = fmt.Errorf(ePrefix + "Error: INVALID PATH. Invalid interior relative path detected! adjustedFileNameExt='%v'", adjustedFileNameExt)
		return
	}

	slashIdxs, err := fh.GetPathSeparatorIndexesInPathStr(adjustedFileNameExt)

	if err != nil {
		err = fmt.Errorf(ePrefix+ "Error returned from fh.GetPathSeparatorIndexesInPathStr(adjustedFileNameExt). adustedFileNameExt='%v'  Error='%v'", adjustedFileNameExt, err.Error())
		return
	}

	lSlashIdxs := len(slashIdxs)

	if lSlashIdxs == 0 {
		fileNameExt = adjustedFileNameExt
		isEmpty = false
		err = nil
		return
	}

	if lastCharIdx < slashIdxs[lSlashIdxs-1] {
		// Example: ../dir1/dir2/
		fileNameExt = ""
		isEmpty = true
		err = nil
		return
	}

	result := adjustedFileNameExt[slashIdxs[lSlashIdxs-1]+1:]

	fileNameExt = result

	if len(result) == 0 {
		isEmpty = true
	} else {
		isEmpty = false
	}

	err = nil
	return
}

// CreateFile - Wrapper function for os.Create
func (fh FileHelper) CreateFile(pathFileName string) (*os.File, error) {
	return os.Create(pathFileName)
}


// DeleteFilesWalkDir - !!! BE CAREFUL !!! This method deletes files in
// a specified directory tree.
//
// This method searches for files residing in the directory tree
// identified by the input parameter 'startPath'. The method 'walks the
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
// Input Parameters:
// =================
//
// startPath string					- This directory path string specifies the
//														directory containing files which will be matched
//														for deletion according to the file selection criteria.
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
//					SelectCriterionMode					FileSelectCriterionMode
//				}
//
//			The FileSelectionCriteria type allows for configuration of single or multiple
// 			file selection criterion. The 'SelectCriterionMode' can be used to specify
// 			whether the file must match all or any one of the active file selection criterion.
//
//			Elements of the FileSelectionCriteria are described below:
//
// 			FileNamePatterns []string		- An array of strings which may define
//																		one or more search patterns. If a file
//																		name matches any one of the search pattern
//																		strings, it is deemed to be a 'match'
//																		for this search pattern criterion.
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
// 																		that file is considered a 'match'	for this file selection criterion.
//
//																	  If the value of 'FilesOlderThan' is set to time zero,
//																		the default value for type time.Time{}, then this
//																		file selection criterion is considered to be 'Inactive'
//																		or 'Not Set'.
//
//        FilesNewerThan	time.Time	- This date time type is compared to the file
//																		modification date time in order to determine
//																		whether the file is newer than the 'FilesNewerThan'
//																		file selection criterion. If the file modification
// 																		date time is newer than the 'FilesNewerThan' date time,
// 																		that file is considered a 'match' for this file selection
// 																		criterion.
//
//																	  If the value of 'FilesNewerThan' is set to time zero,
//																		the default value for type time.Time{}, then this file
//																		selection criterion is considered to be 'Inactive' or
//																		'Not Set'.
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
//					in the target directory tree will be SELECTED FOR DELETION!
//
//
// Return Value:
// =============
//
// 				type DirectoryDeleteFileInfo struct {
//							StartPath            	string			// Starting directory path submitted as an input parameter
// 																								//     to this method
//							DirMgrs          	[]DirMgr		// Returned information on directories found in directory tree
//							ErrReturns           	[]string		// Internal System Errors encountered
//							DeleteFileSelectCriteria FileSelectionCriteria // File Selection Criteria submitted as an
//																														 //   input parameter to this method.
//							DeletedFiles         	[]FileWalkInfo // Infomation on the files deleted by this method.
//					}
//
//					If successful, files matching the file selection criteria specified in input
// 					parameter 'deleteFileSelectionCriteria' will be DELETED and returned in a
// 					'DirectoryDeleteFileInfo' structure field, DirectoryDeleteFileInfo.DeletedFiles.
//
//					Note: It is a good idea to check the returned field DirectoryDeleteFileInfo.ErrReturns
// 								to determine if any internal system errors were encountered during file processing.
//
//	error	- If a program execution error is encountered during processing, it will
//					returned as an 'error' type. Also, see the comment on DirectoryDeleteFileInfo.ErrReturns,
// 					above.
//
func (fh FileHelper) DeleteFilesWalkDir(startPath string, deleteFileSelectionCriteria FileSelectionCriteria) (DirectoryDeleteFileInfo, error) {
	ePrefix := "FileHelper.DeleteFilesWalkDir() "

	deleteFilesInfo := DirectoryDeleteFileInfo{}

	startPath = fh.RemovePathSeparatorFromEndOfPathString(startPath)

	if ! fh.DoesFileExist(startPath) {
		return deleteFilesInfo, fmt.Errorf(ePrefix+"Error: startPath DOES NOT EXIST! startPath='%v'", startPath)
	}

	deleteFilesInfo.StartPath = startPath
	deleteFilesInfo.DeleteFileSelectCriteria = deleteFileSelectionCriteria

	err :=  fp.Walk(deleteFilesInfo.StartPath, fh.makeFileHelperWalkDirDeleteFilesFunc(&deleteFilesInfo))

	if err != nil {
		return deleteFilesInfo, fmt.Errorf(ePrefix+"Error returned by  FileHelper.makeFileHelperWalkDirDeleteFilesFunc(&dWalkInfo). dWalkInfo.StartPath='%v' Error='%v' ",deleteFilesInfo.StartPath, err.Error())
	}

	return deleteFilesInfo, nil
}


// DeleteDirFile - Wrapper function for Remove.
// Remove removes the named file or directory.
// If there is an error, it will be of type *PathError.
func (fh FileHelper) DeleteDirFile(pathFile string) error {
	ePrefix := "FileHelper.DeleteDirFile() "

	if len(pathFile) == 0 {
		return fmt.Errorf(ePrefix + "Error: Input parameter 'pathFile' is a Zero length string!")
	}

	if !fh.DoesFileExist(pathFile) {
		// Doesn't exist. Nothing to do.
		return nil
	}

	err := os.Remove(pathFile)

	if err != nil {
		return fmt.Errorf(ePrefix + "Error returned from os.Remove(pathFile). pathFile='%v' Error='%v'", pathFile, err.Error())
	}

	return nil
}

// DeleteDirPathAll - Wrapper function for RemoveAll
// RemoveAll removes path and any children it contains.
// It removes everything it can but returns the first
// error it encounters. If the path does not exist,
// RemoveAll returns nil (no error).
func (fh FileHelper) DeleteDirPathAll(pathDir string) error {
	ePrefix := "FileHelper.DeleteDirPathAll() "

	if len(pathDir) == 0 {
		return fmt.Errorf(ePrefix + "Error: Input parameter 'pathDir' is a ZERO length string!")
	}

	// If the path does NOT exist,
	// 'RemoveAll()' returns 'nil'.
	err := os.RemoveAll(pathDir)

	if err != nil {
		return fmt.Errorf(ePrefix + "Error returned by os.RemoveAll(pathDir). pathDir='%v'  Error='%v'", pathDir, err.Error())
	}

	return nil
}

// DoesFileExist - Returns a boolean value
// designating whether the passed file name
// exists.
func (fh FileHelper) DoesFileExist(pathFileName string) bool {

	status, _, _ := fh.DoesFileInfoExist(pathFileName)

	return status
}

// DoesFileInfoExist - returns a boolean value indicating
// whether the path and file name passed to the function
// actually exists. Note: If the file actually exists,
// the function will return the associated FileInfo structure.
func (fh FileHelper) DoesFileInfoExist(pathFileName string) (doesFInfoExist bool, fInfo os.FileInfo, err error) {

	doesFInfoExist = false

	if fInfo, err = os.Stat(pathFileName); os.IsNotExist(err) {
		return doesFInfoExist, fInfo, err
	}

	doesFInfoExist = true

	return doesFInfoExist, fInfo, nil

}

func (fh FileHelper) DoesStringEndWithPathSeparator(pathStr string) bool {

	lenStr := len(pathStr)

	if lenStr < 1 {
		return false
	}

	if pathStr[lenStr-1] == '\\' || pathStr[lenStr-1] == '/' || pathStr[lenStr-1] == os.PathSeparator {
		return true
	}

	return false
}

// FilterFileName - Utility method designed to determine whether a file described by a filePath string
// and an os.FileInfo object meets any one of three criteria: A string pattern match, a modification time
// which is older than the 'findFileOlderThan' parameter or a modification time which is newer than the
// 'findFileNewerThan' parameter.
//
// If the three search criteria are all set the their 'zero' or default values, the no selection filter is
// applied and all files are deemed to be a match for the selection criteria ('isMatchedFile=true').
//
// Three selection criterion are applied to the file name (info.Name()).
//
// If a given selection criterion is set to a zero value, then that criterion is defined as 'not set'
// and therefore not used in determining determining whether a file is a 'match'.
//
// If a given criterion is set to a non-zero value, then that criterion is defined as 'set' and the file
// information must comply with that criterion in order to be judged as a match ('isMatchedFile=true').
//
// If none of the three criterion are 'set', then all files are judged as matched ('isMatchedFile=true').
//
// If one of the three criterion is 'set', then a file must comply with that one criterion in order to
// be judged as matched ('isMatchedFile=true').
//
// If two criteria are 'set', then the file must comply with both of those criterion in order to be judged
// as matched ('isMatchedFile=true').
//
// If three criteria are 'set', then the file must comply with all three criterion in order to be judged
// as matched ('isMatchedFile=true').
//
func(fh *FileHelper) FilterFileName(info os.FileInfo, fileSelectionCriteria FileSelectionCriteria ) (isMatchedFile bool, err error) {

	ePrefix := "FileHelper.FilterFileName() "
	isMatchedFile = false
	err = nil

	isPatternSet, isPatternMatch, err2 := fh.SearchFilePatternMatch(info, fileSelectionCriteria)

	if err2 != nil {
		err = fmt.Errorf(ePrefix + "Error returned from dMgr.searchFilePatternMatch(info, fileSelectionCriteria) info.Name()='%v' Error='%v'", info.Name(), err.Error() )
		isMatchedFile = false
		return
	}

	isFileOlderThanSet, isFileOlderThanMatch, err2 := fh.SearchFileOlderThan(info, fileSelectionCriteria)

	if err2 != nil {
		err = fmt.Errorf(ePrefix + "Error returned from dMgr.searchFileOlderThan(info, fileSelectionCriteria) fileSelectionCriteria.FilesOlderThan='%v' info.Name()='%v' Error='%v'", fileSelectionCriteria.FilesOlderThan, info.Name(), err.Error() )
		isMatchedFile = false
		return
	}


	isFileNewerThanSet, isFileNewerThanMatch, err2 := fh.SearchFileNewerThan(info, fileSelectionCriteria)

	if err2 != nil {
		err = fmt.Errorf(ePrefix + "Error returned from dMgr.searchFileNewerThan(info, fileSelectionCriteria) fileSelectionCriteria.FilesNewerThan='%v' info.Name()='%v' Error='%v'", fileSelectionCriteria.FilesNewerThan, info.Name(), err.Error() )
		isMatchedFile = false
		return
	}

	isFileModeSearchSet, isFileModeSearchMatch, err2 := fh.SearchFileModeMatch(info, fileSelectionCriteria)

	if err2 != nil {
		err = fmt.Errorf(ePrefix + "Error returned from dMgr.searchFileModeMatch(info, fileSelectionCriteria) fileSelectionCriteria.SelectByFileMode='%v' info.Name()='%v' Error='%v'", fileSelectionCriteria.SelectByFileMode, info.Name(), err.Error() )
		isMatchedFile = false
		return
	}


	// If no file selection criterion are set, then always select the file
	if !isPatternSet && !isFileOlderThanSet && !isFileNewerThanSet && !isFileModeSearchSet {
		isMatchedFile = true
		err = nil
		return
	}

	// If using the AND File Select Criterion Mode, then for criteria that
	// are set and active, they must all be 'matched'.
	if fileSelectionCriteria.SelectCriterionMode == ANDFILESELECTCRITERION {

		if isPatternSet && !isPatternMatch {
			isMatchedFile = false
			err = nil
			return
		}

		if isFileOlderThanSet && !isFileOlderThanMatch {
			isMatchedFile = false
			err = nil
			return
		}

		if isFileNewerThanSet && !isFileNewerThanMatch {
			isMatchedFile = false
			err = nil
			return
		}

		if isFileModeSearchSet && !isFileModeSearchMatch {
			isMatchedFile = false
			err = nil
			return
		}

		isMatchedFile = true
		err = nil
		return

	} // End of ANDFILESELECTCRITERION


	// Must be ORFILESELECTCRITERION Mode
	// If ANY of the section criterion are active and 'matched', then
	// classify the file as matched.

	if isPatternSet && isPatternMatch {
		isMatchedFile = true
		err = nil
		return
	}

	if isFileOlderThanSet && isFileOlderThanMatch {
		isMatchedFile = true
		err = nil
		return
	}

	if isFileNewerThanSet && isFileNewerThanMatch {
		isMatchedFile = true
		err = nil
		return
	}

	if isFileModeSearchSet && isFileModeSearchMatch {
		isMatchedFile = true
		err = nil
		return
	}

	isMatchedFile = false
	err = nil
	return
}

// FindFilesWalkDirectory - This method returns file information on files residing in a specified
// directory tree identified by the input parameter, 'startPath'.
//
// This method 'walks the directory tree' locating all files in the directory tree which match
// the file selection criteria submitted as input parameter, 'fileSelectCriteria'.
//
// If a file matches the File Selection Criteria, it is included in the returned field,
// 'DirectoryTreeInfo.FoundFiles'. By the way, if ALL the file selection criterion are set to zero values
// or 'Inactive', then ALL FILES in the directory are selected and returned in the field,
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
//					be returned as an 'error' type. Also, see the comment on 'DirectoryTreeInfo.ErrReturns',
// 					above.
//
func (fh FileHelper) FindFilesWalkDirectory(startPath string, fileSelectCriteria FileSelectionCriteria) (DirectoryTreeInfo, error) {

	ePrefix := "FileHelper.FindFilesWalkDirectory() "

	findFilesInfo := DirectoryTreeInfo{}

	startPath = fh.RemovePathSeparatorFromEndOfPathString(startPath)

	if !fh.DoesFileExist(startPath) {
		return findFilesInfo, fmt.Errorf(ePrefix + "Error - startPath DOES NOT EXIST! startPath='%v'", startPath)
	}

	findFilesInfo.StartPath = startPath

	findFilesInfo.FileSelectCriteria = fileSelectCriteria


	// err := fp.Walk(findFilesInfo.StartPath, fh.MakeWalkDirGetFilesFunc(dInfo))
	err :=  fp.Walk(findFilesInfo.StartPath, fh.makeFileHelperWalkDirFindFilesFunc(&findFilesInfo))

	if err != nil {

		return findFilesInfo, fmt.Errorf(ePrefix+"Error returned from fp.Walk(findFilesInfo.StartPath, fh.makeFileHelperWalkDirFindFilesFunc(&findFilesInfo)). startPath='%v' Error='%v'",startPath, err.Error())
	}


	return findFilesInfo, nil
}

// GetAbsPathFromFilePath - Supply a string containing both
// the path file name and extension and return the path
// element.
func (fh FileHelper) GetAbsPathFromFilePath(filePath string) (string, error) {
	ePrefix := "FileHelper.GetAbsPathFromFilePath() "

	if len(filePath) == 0 {
		return "", errors.New(ePrefix+ "Error: Input parameter 'filePath' is an EMPTY string!")
	}

	testFilePath := fh.AdjustPathSlash(filePath)

	if len(testFilePath) == 0 {
		return "", errors.New(ePrefix + "Error: After adjusting Path Separators, filePath resolves to an empty string!")
	}

	absPath, err := fh.MakeAbsolutePath(testFilePath)

	if err!=nil {
		return "", fmt.Errorf(ePrefix+"Error returned from ")
	}

	return absPath, nil
}

// GetAbsCurrDir - returns
// the absolute path of the
// current working directory
func (fh FileHelper) GetAbsCurrDir() (string, error) {
	ePrefix := "FileHelper.GetAbsCurrDir() "

	dir, err := fh.GetCurrentDir()

	if err != nil {
		return dir, fmt.Errorf(ePrefix + "Error returned from fh.GetCurrentDir(). Error='%v'", err.Error())
	}

	return fh.MakeAbsolutePath(dir)
}

// GetCurrentDir - Wrapper function for
// Getwd(). Getwd returns a rooted path name
// corresponding to the current directory.
// If the current directory can be reached via
// multiple paths (due to symbolic links),
// Getwd may return any one of them.
func (fh FileHelper) GetCurrentDir() (string, error) {
	return os.Getwd()
}

// GetDotSeparatorIndexesInPathStr - Returns an array of integers representing the
// indexes of dots ('.') located in input paramter 'pathStr'
func (fh FileHelper) GetDotSeparatorIndexesInPathStr(pathStr string) ([]int, error) {

	ePrefix := "FileHelper.GetDotSeparatorIndexesInPathStr() "
	lPathStr := len(pathStr)

	if lPathStr == 0 {
		return []int{}, fmt.Errorf(ePrefix + "Error: Zero length 'pathStr' passed to this method!")
	}

	var dotIdxs [] int

	for i:=0; i < lPathStr; i++ {

		rChar := pathStr[i]

		if rChar == '.' {

			dotIdxs = append(dotIdxs, i)
		}

	}

	return dotIdxs, nil

}

// GetExecutablePathFileName - Gets the file name
// and path of the executable that started the
// current process
func (fh FileHelper) GetExecutablePathFileName() (string, error) {
	ex, err := os.Executable()

	return ex, err

}

// GetFileNameWithExt - This method expects to receive a valid directory path and file name or file name plus
// extension. It then extracts the File Name and Extension from the file path and returns it as a string.
//
// Input Parameter:
// ================
//
// pathFileNameExt string		- This input parameter is expected to contain a properly formatted directory
//														path and File Name.  The File Name may or may not include a File Extension.
//														The directory path must include the correct delimiters such as Path Separators
//														('/' or'\'), dots ('.') and in the case of Windows, a volume designation
// 														(Example: 'F:').
//
// Return Values:
// ==============
//
// fNameExt	string					- If successful, this method will the return value 'fNameExt' equal to the
//													  File Name and File Extension extracted from the input file path, 'pathFileNameExt'.
//														Example 'fNameExt' return value: 'somefilename.txt'
//
//														If the File Extension is not present, only the File Name will be returned.
//														Example return value with no file extension: 'somefilename'.
//
// isEmpty bool							- If this method CAN NOT parse out a valid File Name and Extension from
//														input parameter string 'pathFileNameExt', return value 'fNameExt' will
//														be set to an empty string and return value 'isEmpty' will be set to 'true'.
//
// err error								- If an error is encountered during processing, 'err' will return a properly
//														formatted 'error' type.
//
// 														Note that if this method cannot parse out a valid	File Name and Extension
// 														due to an improperly formatted directory path and	file name string
// 														(Input Parameter: 'pathFileNameExt'), 'fNameExt' will be set to an empty string,
// 														'isEmpty' will be set to 'true' and 'err' return 'nil'. In this situation, no
//														error will be returned.
//
func (fh FileHelper) GetFileNameWithExt(pathFileNameExt string) (fNameExt string, isEmpty bool, err error) {

	ePrefix := "FileHelper.GetFileNameWithExt"
	fNameExt = ""
	isEmpty = true
	err = nil

	if len(pathFileNameExt) == 0 {
		err = errors.New(ePrefix + "Error: Input parameter 'pathFileNameExt' is a ZERO Length string!")
		return
	}

	pathFileNameExt = strings.TrimLeft(strings.TrimRight(pathFileNameExt, " "), " ")

	if pathFileNameExt == "" {
		err = errors.New(ePrefix + "Error: After trimming 'pathFileNameExt', input parameter 'pathFileNameExt' is a ZERO Length string!")
		return
	}

	testPathFileNameExt := fh.AdjustPathSlash(pathFileNameExt)

	volName := fh.GetVolumeName(testPathFileNameExt)

	if volName != "" {
		testPathFileNameExt = strings.TrimPrefix(testPathFileNameExt, volName)
	}

	lTestPathFileNameExt := len(testPathFileNameExt)

	if lTestPathFileNameExt == 0 {
		err = errors.New(ePrefix + "Error: Cleaned version of 'pathFileNameExt', 'testPathFileNameExt' is a ZERO Length string!")
		return
	}

	firstCharIdx, lastCharIdx, err2 := fh.GetFirstLastNonSeparatorCharIndexInPathStr(testPathFileNameExt)

	if err2 != nil {
		err = fmt.Errorf(ePrefix + "Error returned by fh.GetFirstLastNonSeparatorCharIndexInPathStr(testPathFileNameExt). testPathFileNameExt='%v'  Error='%v'", testPathFileNameExt, err2.Error())
		return
	}

	// There are no alpha numeric characters present.
	// Therefore, there is no file name and extension
	if firstCharIdx == -1  || lastCharIdx == -1 {
		isEmpty = true
		err = nil
		return
	}

	slashIdxs, err2 := fh.GetPathSeparatorIndexesInPathStr(testPathFileNameExt)

	if err2 != nil {
		err = fmt.Errorf(ePrefix + "Error returned by fh.GetPathSeparatorIndexesInPathStr(testPathFileNameExt). testPathFileNameExt='%v'  Error='%v'", testPathFileNameExt, err2.Error())
		return
	}

	dotIdxs, err2 := fh.GetDotSeparatorIndexesInPathStr(testPathFileNameExt)

	if err2 != nil {
		err = fmt.Errorf(ePrefix + "Error returned by fh.GetDotSeparatorIndexesInPathStr(testPathFileNameExt). testPathFileNameExt='%v'  Error='%v'", testPathFileNameExt, err2.Error())
		return

	}

	lSlashIdxs := len(slashIdxs)
	lDotIdxs := len(dotIdxs)

	if lSlashIdxs > 0  {
		// This string has path separators

		// Last char is a path separator. Therefore,
		// there is no file name and extension.
		if slashIdxs[lSlashIdxs-1] == lTestPathFileNameExt - 1  {
			fNameExt = ""
		} else if lastCharIdx > slashIdxs[lSlashIdxs-1] {

			fNameExt = testPathFileNameExt[slashIdxs[lSlashIdxs-1]+1:]

		} else {
			fNameExt = ""
		}

		if len(fNameExt) == 0 {
			isEmpty = true
		} else {
			isEmpty = false
		}

		err = nil
		return
	}

	// There are no path separators lSlashIdxs == 0

	if lDotIdxs > 0 {
		// This string has one or more dot separators ('.')

		fNameExt = ""

		if firstCharIdx > dotIdxs[lDotIdxs-1] {
			// Example '.txt' - Invalid File name and extension
			isEmpty = true
			err = fmt.Errorf(ePrefix + "Error: File extension exists but no file name. result='%v'", testPathFileNameExt)
			return

		} else if firstCharIdx < dotIdxs[lDotIdxs-1] {
			fNameExt = testPathFileNameExt[firstCharIdx:]
		}

		if len(fNameExt) == 0 {
			isEmpty = true
		} else {
			isEmpty = false
		}

		err = nil
		return
	}

	// Must be lSlashIdxs == 0 && lDotIdxs ==  0
	// There are no Path Separators and there are
	// no dot separators ('.').

	fNameExt = testPathFileNameExt[firstCharIdx:]

	if len(fNameExt) == 0 {
		isEmpty = true
	} else {
		isEmpty = false
	}

	err = nil

	return
}

// GetFileNameWithoutExt - returns the file name
// without the path or extension. If the returned
// File Name is an empty string, isEmpty is set to true.
//
// Example
// pathFileNameExt = ./pathfilego/003_filehelper/common/dirmgr_01_test.go
// Returned 'fName' = dirmgr_01_test
//
func (fh FileHelper) GetFileNameWithoutExt(pathFileNameExt string) (fName string, isEmpty bool, err error) {
	ePrefix:= "FileHelper.GetFileNameWithoutExt()"

	isEmpty = true
	fName = ""
	err = nil

	if len(pathFileNameExt) == 0 {
		err = errors.New(ePrefix + "Error: Input parameter 'pathFileNameExt' is a ZERO Length string!")
		return
	}

	pathFileNameExt = strings.TrimLeft(strings.TrimRight(pathFileNameExt, " "), " ")

	if pathFileNameExt == "" {
		err = errors.New(ePrefix + "Error: After trimming 'pathFileNameExt', input parameter 'pathFileNameExt' is a ZERO Length string!")
		return
	}

	testPathFileNameExt := fh.AdjustPathSlash(pathFileNameExt)

	if len(testPathFileNameExt) == 0 {
		err = errors.New(ePrefix + "Error: Adjusted Path version of 'pathFileNameExt', 'testPathFileNameExt' is a ZERO Length string!")
		return
	}

	fileNameExt, isFileNameExtEmpty, err2 := fh.GetFileNameWithExt(testPathFileNameExt)

	if err2 != nil {
		err = fmt.Errorf(ePrefix + "Error returned from fh.GetFileNameWithExt(testPathFileNameExt) testPathFileNameExt='%v'  Error='%v'", testPathFileNameExt, err2.Error())
		return
	}

	if isFileNameExtEmpty {
		isEmpty = true
		fName = ""
		err = nil
		return
	}

	dotIdxs, err2 := fh.GetDotSeparatorIndexesInPathStr(fileNameExt)

	if err2 != nil {
		err = fmt.Errorf(ePrefix + "Error returned from fh.GetDotSeparatorIndexesInPathStr(fileNameExt). fileNameExt='%v'  Error='%v'",fileNameExt, err2.Error())
		return
	}

	lDotIdxs := len(dotIdxs)

	// Primary Case: filename.ext
	if lDotIdxs > 0 {
		fName = fileNameExt[0:dotIdxs[lDotIdxs-1]]

		if fName == "" {
			isEmpty = true
		} else {
			isEmpty = false
		}
		err = nil
		return
	}

	// Secondary Case: filename
	fName = fileNameExt

	if fName == "" {
		isEmpty = true
	} else {
		isEmpty = false
	}

	err = nil
	return
}

// GetFileExtension - Returns the File Extension with
// the dot. If there is no File Extension an empty
// string is returned (NO dot included). If the returned
// File Extension is an empty string, the returned
// parameter 'isEmpty' is set equal to 'true'.
//
// When an extension is returned in the 'ext' variable, this
// extension includes a leading dot. Example: '.txt'
func (fh FileHelper) GetFileExtension(pathFileNameExt string) (ext string, isEmpty bool, err error) {
	ePrefix := "FileHelper.GetFileExtension() "

	ext = ""
	isEmpty = true
	err = nil

	pathFileNameExt = strings.TrimLeft(strings.TrimRight(pathFileNameExt, " "), " ")

	if len(pathFileNameExt) == 0 {
		err = errors.New(ePrefix + "Error: After trimming 'pathFileNameExt'. Input parameter 'pathFileNameExt' is a Zero length string!")
		return
	}

	testPathFileNameExt := fh.AdjustPathSlash(pathFileNameExt)

	lenTestPathFileNameExt := len(testPathFileNameExt)

	if lenTestPathFileNameExt == 0 {
		err = errors.New(ePrefix + "Error: Cleaned version of 'pathFileNameExt', 'testPathFileNameExt' is a ZERO length string!")
		return
	}

	dotIdxs, err2 := fh.GetDotSeparatorIndexesInPathStr(testPathFileNameExt)

	if err2 != nil {
		ext = ""
		isEmpty = true
		err = fmt.Errorf(ePrefix + "Error returned from fh.GetDotSeparatorIndexesInPathStr(testPathFileNameExt). testPathFileNameExt='%v'  Error='%v'", testPathFileNameExt, err2)
		return
	}

	lenDotIdxs := len(dotIdxs)

	// Deal with case where the pathFileNameExt contains
	// no dots.
	if lenDotIdxs == 0 {
		ext = ""
		isEmpty = true
		err = nil
		return

	}


	firstGoodCharIdx, lastGoodCharIdx, err2 := fh.GetFirstLastNonSeparatorCharIndexInPathStr(testPathFileNameExt)

	if err2 != nil {
		ext = ""
		isEmpty = true
		err = fmt.Errorf(ePrefix + "Error returned from fh.GetFirstLastNonSeparatorCharIndexInPathStr(testPathFileNameExt). testPathFileNameExt='%v'  Error='%v'", testPathFileNameExt, err2)
		return
	}

	// Deal with the case where pathFileNameExt contains no
	// valid alpha numeric characters
	if firstGoodCharIdx== -1 || lastGoodCharIdx == -1 {
		ext = ""
		isEmpty = true
		err = nil
		return
	}


	slashIdxs, err2 := fh.GetPathSeparatorIndexesInPathStr(testPathFileNameExt)

	if err2 != nil {
		ext = ""
		isEmpty = true
		err = fmt.Errorf(ePrefix + "Error returned from fh.GetPathSeparatorIndexesInPathStr(testPathFileNameExt). testPathFileNameExt='%v'  Error='%v'", testPathFileNameExt, err2)
		return
	}

	lenSlashIdxs := len(slashIdxs)

	if lenSlashIdxs == 0 {
		ext = testPathFileNameExt[dotIdxs[lenDotIdxs-1]:]
		isEmpty = false
		err = nil
		return
	}

	// lenDotIdxs and lenSlasIdxs both greater than zero
	if dotIdxs[lenDotIdxs-1] > slashIdxs[lenSlashIdxs-1] &&
		dotIdxs[lenDotIdxs-1] < lastGoodCharIdx {

		ext = testPathFileNameExt[dotIdxs[lenDotIdxs-1]:]
		isEmpty = false
		err = nil
		return

	}

	ext = ""
	isEmpty = true
	err = nil
	return
}

// GetFileInfoFromPath - Wrapper function for os.Stat(). This method
// can be used to return FileInfo data on a specific file. If the file
// does NOT exist, an error will be triggered. This method is similar to
// FileHelpter.DoesFileInfoExist().
//
// type FileInfo interface {
// 	Name() string       // base name of the file
// 	Size() int64        // length in bytes for regular files; system-dependent for others
// 	Mode() FileMode     // file mode bits
// 	ModTime() time.Time // modification time
// 	IsDir() bool        // abbreviation for Mode().IsDir()
// 	Sys() interface{}   // underlying data source (can return nil)
// }
func (fh FileHelper) GetFileInfoFromPath(pathFileName string) (os.FileInfo, error) {

	return os.Stat(pathFileName)

}

// GetFileLastModificationDate - Returns the last modification'
// date/time on a specific file. If input parameter 'customTimeFmt'
// string is empty, default time format will be used to format the
// returned time string.
func (fh FileHelper) GetFileLastModificationDate(pathFileName string, customTimeFmt string) (time.Time, string, error) {

	const fmtDateTimeNanoSecondStr = "2006-01-02 15:04:05.000000000"
	var zeroTime time.Time

	if pathFileName == "" {
		return zeroTime, "", errors.New("FileHelper:GetFileLastModificationDate() Error: Input parameter 'pathFileName' is empty string!")
	}

	fmtStr := customTimeFmt

	if len(fmtStr) == 0 {
		fmtStr = fmtDateTimeNanoSecondStr
	}

	fInfo, err := fh.GetFileInfoFromPath(pathFileName)

	if err != nil {
		return zeroTime, "", errors.New(fmt.Sprintf("FileHelper:GetFileLastModificationDate() Error Getting FileInfo on %v Error on GetFileInfoFromPath(): %v", pathFileName, err.Error()))
	}

	return fInfo.ModTime(), fInfo.ModTime().Format(fmtStr), nil
}

// GetFirstLastNonSeparatorCharIndexInPathStr - Basically this method returns
// the first index of the first alpha numeric character in a path string.
//
// Specifically, the character must not be a Path Separator ('\', '/') and
// it must not be a dot ('.').
//
// If the first Non-Separator char is found, this method will return
// an integer index which is greater than or equal to zero plus an
// error value of nil.
//
// The first character found will never be part of the volume name.
// Example On Windows: "D:\fDir1\fDir2" - first character index will
// be 3 denoting character 'f'.
//
func (fh FileHelper) GetFirstLastNonSeparatorCharIndexInPathStr(pathStr string) (firstIdx, lastIdx int, err error) {

	ePrefix := "FileHelper.GetFirstNonSeparatorCharIndexInPathStr() "
	lPathStr := len(pathStr)
	firstIdx = -1
	lastIdx = -1

	if lPathStr == 0 {
		err = fmt.Errorf(ePrefix + "Error: Zero length 'pathStr' passed to this method!")
		return
	}

	pathStr = fp.FromSlash(pathStr)

	lPathStr = len(pathStr)

	if lPathStr == 0 {
		err = fmt.Errorf(ePrefix + "Error: After Path Separator adjustment, 'pathStr' is a Zero length string!")
		return
	}

	// skip the volume name. Don't count
	// first characters in the volume name
	volName := fp.VolumeName(pathStr)
	lVolName := len(volName)

	startIdx := 0

	if lVolName > 0 {
		startIdx = lVolName
	}


	var rChar rune

	forbiddenTextChars := []rune {os.PathSeparator,
		'\\',
		'/',
		'|',
		'.',
		'&',
		'!',
		'%',
		'$',
		'#',
		'@',
		'^',
		'*',
		'(',
		')',
		'-',
		'_',
		'+',
		'=',
		'[',
		'{',
		']',
		'}',
		'|',
		'<',
		'>',
		',',
		'~',
		'`',
		':',
		';',
		'"',
		'\'',
		'\n',
		'\t',
		'\r'}

	lForbiddenTextChars := len(forbiddenTextChars)

	for i:=startIdx; i < lPathStr; i++ {
		rChar = rune(pathStr[i])
		isForbidden := false

		for j:= 0; j<lForbiddenTextChars; j++ {
			if rChar == forbiddenTextChars[j] {
				isForbidden = true
			}

		}

		if isForbidden == false {

			if firstIdx == -1 {
				firstIdx = i
			}

			lastIdx = i
		}

	}

	err = nil

	return
}

// GetLastPathElement - Analyzes a 'pathName' string and returns the last
// element in the path. If 'pathName' ends in a path separator ('/'), this
// method returns an empty string.
//
// Example:
// pathName = '../dir1/dir2/fileName.ext' will return "fileName.ext"
// pathName = '../dir1/dir2/' will return ""
// pathName = 'fileName.ext' will return "fileName.ext"
// pathName = '../dir1/dir2/dir3' will return "dir3"
//
func (fh FileHelper) GetLastPathElement(pathName string) (string, error) {
	ePrefix := "FileHelper.GetLastPathElement() "

	if len(pathName) == 0 {
		return "", errors.New(ePrefix + "Error: Input parameter 'pathName' is a Zero Length String!")
	}

	adjustedPath := fh.AdjustPathSlash(pathName)

	resultAry := strings.Split(adjustedPath, string(os.PathSeparator))

	lResultAry := len(resultAry)

	if lResultAry == 0 {
		return adjustedPath, nil
	}

	return resultAry[lResultAry-1], nil
}

// GetPathAndFileNameExt - Breaks out path and FileName+Ext elements from
// a path string. If both path and fileName are empty strings, this method
// returns an error.
func (fh FileHelper) GetPathAndFileNameExt(pathFileNameExt string) (pathDir, fileNameExt string, bothAreEmpty bool, err error) {

	ePrefix := "FileHelper.GetPathAndFileNameExt() "
	pathDir = ""
	fileNameExt = ""
	bothAreEmpty = true
	err = nil

	if pathFileNameExt == "" {
		err = errors.New(ePrefix + "Error: Input parameter 'pathFileNameExt' is a Zero length string!")
		return
	}

	trimmedFileNameExt := strings.TrimLeft(strings.TrimRight(pathFileNameExt," ")," ")

	if len(trimmedFileNameExt) == 0 {
		err = errors.New(ePrefix + "Error: Trimmed input parameter 'pathFileNameExt' is a Zero length string!")
		return
	}

	xFnameExt, isEmpty, err2 := fh.GetFileNameWithExt(trimmedFileNameExt)

	if err2 != nil {
		err = fmt.Errorf(ePrefix+"Error returned from fh.GetFileNameWithExt(pathFileNameExt). pathFileNameExt='%v' Error='%v'", pathFileNameExt, err2.Error())
		return
	}

	if isEmpty {
		fileNameExt = ""
	} else {
		fileNameExt = xFnameExt
	}

	remainingPathStr := strings.TrimSuffix(trimmedFileNameExt, fileNameExt)

	if len(remainingPathStr) == 0 {
		pathDir = ""

		if pathDir == "" && fileNameExt == "" {
			bothAreEmpty = true
		} else {
			bothAreEmpty = false
		}

		err = nil

		return

	}

	xPath, isEmpty, err2 := fh.GetPathFromPathFileName(remainingPathStr)

	if err2 != nil {
		err = fmt.Errorf(ePrefix+"Error returned from fh.GetPathFromPathFileName(remainingPathStr). remainingPathStr='%v' Error='%v'", remainingPathStr, err2.Error())
		return
	}

	if isEmpty {
		pathDir = ""
	} else {
		pathDir = xPath
	}

	if pathDir == "" && fileNameExt == "" {
		bothAreEmpty = true
	} else {
		bothAreEmpty = false
	}

	err = nil

	return
}

// GetPathFromPathFileName - Returns the path from a path and file name string.
// If the returned path is an empty string, return parameter 'isEmpty' is set to
// 'true'.
//
// Input Parameter:
// ================
//
// pathFileNameExt string		- This is an input parameter. The method expects to
// 														receive a single, properly formatted path and file
//														name string delimited by dots ('.') and Path Separators
//														('/' or '\'). On Windows the 'pathFileNameExt' string
//														valid volume designations (Example: "D:")
// Return Values:
// ==============
//
// path string							- This is the directory path extracted from the input parameter
//														'pathFileNameExt'. If successful, the 'path' string that is returned
// 														by this method WILL NOT include a trailing path separator ('/' or '\'
// 														depending on the os). Example 'path': "./pathfile/003_filehelper"
//
// isEmpty bool							- If the method determines that it cannot extract a valid directory
//														path from input parameter 'pathFileNameExt', this boolean value
//														will be set to 'true'. Failure to extract a valid directory path
//														will occur if the input parameter 'pathFileNameExt' is not properly
//														formatted as a valid path and file name.
//
// err error								- If a processing error is detected, an error will be returned. Note that
//														in the event that this method fails to extract a valid directory path
//														'pathFileNameExt' due to the fact that 'pathFileNameExt' was improperly
//														formatted, 'isEmpty' will be set to 'true', but no error will be returned.
//
//														If no error is returned, 'err' is set to 'nil'.
//
// Examples:
// =========
//
// pathFileNameExt = ""					returns isEmpty==true  err==nil
// pathFileNameExt = "D:\"  		returns "D:\"
// pathFileNameExt = "."				returns "."
// pathFileNameExt = "..\"			returns "..\"
// pathFileNameExt = "...\"			returns ERROR
// pathFileNameExt = ".\pathfile\003_filehelper\HowToRunTests.md"		returns ".\pathfile\003_filehelper"
//
func (fh FileHelper) GetPathFromPathFileName(pathFileNameExt string) (dirPath string, isEmpty bool, err error) {
	ePrefix := "FileHelper.GetPathFromPathFileName() "
	dirPath = ""
	isEmpty = true
	err = nil

	if len(pathFileNameExt) == 0 {
		err = errors.New(ePrefix + "Error: Input parameter 'pathFileNameExt' is a ZERO length string!")
		return
	}

	testPathStr, isDirEmpty , err2 := fh.CleanDirStr(pathFileNameExt)

	if err2 != nil {
		err = fmt.Errorf(ePrefix + "Error returned by fh.CleanDirStr(pathFileNameExt). pathFileNameExt='%v'  Error='%v'", pathFileNameExt, err2.Error())
		return
	}

	if isDirEmpty {
		dirPath = ""
		isEmpty = true
		err = nil
		return
	}

	lTestPathStr := len(testPathStr)

	if lTestPathStr == 0 {
		err = errors.New(ePrefix + "Error: AdjustPathSlash was applied to 'pathStr'. The 'testPathStr' string is a Zero Length string!")
		return
	}

	slashIdxs, err2 := fh.GetPathSeparatorIndexesInPathStr(testPathStr)

	if err2 != nil {
		err = fmt.Errorf(ePrefix + "Error returned by fh.GetPathSeparatorIndexesInPathStr(testPathStr). testPathStr='%v'  Error='%v'", testPathStr, err2.Error())
		return
	}

	lSlashIdxs := len(slashIdxs)

	firstGoodChar, lastGoodChar, err2:= fh.GetFirstLastNonSeparatorCharIndexInPathStr(testPathStr)

	if err2 != nil {
		err = fmt.Errorf(ePrefix + "Error returned by fh.GetFirstLastNonSeparatorCharIndexInPathStr(testPathStr). testPathStr='%v'  Error='%v'", testPathStr, err2.Error())
	}

	dotIdxs, err2 := fh.GetDotSeparatorIndexesInPathStr(testPathStr)

	if err2 != nil {
		err = fmt.Errorf(ePrefix + "Error returned by fh.GetDotSeparatorIndexesInPathStr(testPathStr). testPathStr='%v'  Error='%v'", testPathStr, err2.Error())
		return
	}

	lDotIdxs := len(dotIdxs)

	var finalPathStr string

	volName := fp.VolumeName(testPathStr)

	if testPathStr == volName {

		finalPathStr = testPathStr

	} else if strings.Contains(testPathStr, "...") {

		err = fmt.Errorf(ePrefix + "Error: PATH CONTAINS INVALID Dot Characters! testPathStr='%v'", testPathStr)
		return

	}else	if firstGoodChar == -1 || lastGoodChar == -1 {

		absPath, err2 := fh.MakeAbsolutePath(testPathStr)

		if err2!= nil  {
			err = fmt.Errorf(ePrefix + "Error returned from fh.MakeAbsolutePath(testPathStr). testPathStr='%v' Error='%v'", testPathStr, err2.Error())
			return
		}

		if absPath == "" {
			err = fmt.Errorf(ePrefix + "Error: Could not convert 'testPathStr' to Absolute Path! tesPathStr='%v'", testPathStr)
			return
		}

		finalPathStr = testPathStr

	} else if lSlashIdxs==0 {
		// No path separators but alpha numeric chars are present
		dirPath = ""
		isEmpty = true
		err = nil
		return

	} else if lDotIdxs == 0 {
		//Path separators are present but there are no dots in the string

		if slashIdxs[lSlashIdxs-1] == lTestPathStr-1 {
			// Trailing path separator
			finalPathStr = testPathStr[0: slashIdxs[lSlashIdxs-2]]
		} else {
			finalPathStr = testPathStr
		}

	}else	if dotIdxs[lDotIdxs-1] > slashIdxs[lSlashIdxs-1] {
		// format: ./dir1/dir2/fileName.ext
		finalPathStr = testPathStr[0: slashIdxs[lSlashIdxs-1]]

	} else if dotIdxs[lDotIdxs-1] < slashIdxs[lSlashIdxs-1] {

		finalPathStr = testPathStr

	} else {
		err = fmt.Errorf(ePrefix + "Error: INVALID PATH STRING. testPathStr='%v'", testPathStr)
		return
	}

	if len(finalPathStr) == 0 {
		err = fmt.Errorf(ePrefix + "Error: Processed Path is a Zero Length String!")
		return
	}



	//Successfully isolated and returned a valid
	// directory path from 'pathFileNameExt'
	dirPath = finalPathStr

	if len(dirPath) == 0 {
		isEmpty = true
	} else {
		isEmpty = false
	}

	err = nil

	return

}


// GetPathSeparatorIndexesInPathStr - Returns an array containing the indexes of
// Path Separators (Forward slashes or backward slashes depending on operating
// system).
func (fh FileHelper) GetPathSeparatorIndexesInPathStr(pathStr string) ([]int, error) {

	ePrefix := "FileHelper.GetPathSeparatorIndexesInPathStr() "
	lPathStr := len(pathStr)

	if lPathStr == 0 {
		return []int{}, fmt.Errorf(ePrefix + "Error: Zero length 'pathStr' passed to this method!")
	}

	var slashIdxs [] int

	for i:=0; i < lPathStr; i++ {

		rChar := pathStr[i]

		if rChar == os.PathSeparator ||
			rChar == '\\' ||
			rChar == '/'  {

			slashIdxs = append(slashIdxs, i)
		}

	}

	return slashIdxs, nil
}

// GetVolumeName - Returns the volume name of associated with
// a given directory path.
func(fh FileHelper)GetVolumeName(pathStr string) string {

	return fp.VolumeName(pathStr)
}

// GetVolumeSeparatorIdxInPathStr - Returns the index of the
// Windows volume separator from an Path string.
func (fh FileHelper) GetVolumeSeparatorIdxInPathStr(pathStr string) (volIdx int, err error) {

	ePrefix := "FileHelper.GetVolumeSeparatorIdxInPathStr()"

	volIdx = -1
	err = nil

	lPathStr := len(pathStr)

	if lPathStr == 0 {
		err = fmt.Errorf(ePrefix + "Error: Input parameter pathStr is a Zero Length string!")
		return
	}

	for i:=0; i < lPathStr ; i++ {

		if rune(pathStr[i]) == ':' {
			volIdx = i
			err = nil
			return
		}

	}

	volIdx = -1
	err = nil
	return
}

// IsAbsolutePath - Wrapper function for path.IsAbs()
// https://golang.org/pkg/path/#IsAbs
// This method reports whether the input parameter is
// an absolute path.
func (fh FileHelper) IsAbsolutePath(pathStr string) bool {
	return path.IsAbs(pathStr)
}


// IsPathString - Attempts to determine whether a string is a
// path string designating a directory (and not a path file name
// file extension string).
//
// Input Parameter
// ===============
// pathStr string 	- The path string to be analyzed.
//
// Return Values
// =============
//
// isPathStr bool				- If the input parameter, 'pathStr'
//												is determined to be a directory
//												path, this return value is set to
//												true. Here, a 'directory path' is defined
//												as a true directory and the path does NOT
//												contain a file name.
//
// cannotDetermine bool	- If the method cannot determine whether
//												the input parameter 'pathStr' is or
// 												is NOT a valid directory path, this
//												this return value will be set to 'true'.
//												The 'cannotDetermine=true' condition occurs
//												with path names like 'D:\DirA\common'. The
//												cannot determine whether 'common' is a file
//												name or a directory name.
//
// 
// testPathStr string		- Input parameter 'pathStr' is subjected to cleaning routines
//												designed to exclude extraneous characters from the analysis.
//												'testPathFileStr' is the actual string on which the analysis was
//												performed.
// 
// err error						- If an error occurs this return value will
//												be populated. If no errors occur, this return
//												value is set to nil.
//
//
// If the path exists on disk, this method will examine the
// associated file information and render a definitive and
// accurate determination as to whether the path string represents
// a directory.
//
func (fh FileHelper) IsPathString(pathStr string) (isPathStr bool, cannotDetermine bool, testPathStr string, err error) {

	ePrefix := "FileHelper.IsPathString() "

	var fInfo os.FileInfo

	lpathStr := len(pathStr)

	if lpathStr == 0 {
		isPathStr = false
		cannotDetermine = false
		testPathStr = ""
		err = errors.New(ePrefix + "Error - Zero Length input parameter 'pathStr'.")
		return
	}

	testPathStr = fp.FromSlash(pathStr)
	lTestPathStr := len(testPathStr)

	if lTestPathStr == 0 {
		isPathStr = false
		cannotDetermine = false
		testPathStr = ""
		err = fmt.Errorf(ePrefix + "Error - fp.FromSlash(pathStr) yielded a Zero Length String. pathStr='%v'", pathStr)
		return
	}

	// See if path actually exists on disk and
	// then examine the File Info object returned.
	fInfo, err = os.Stat(testPathStr)

	if err==nil  {

		if fInfo.IsDir() {
			isPathStr = true
			cannotDetermine = false
			err = nil
			return

		} else {
			isPathStr = false
			cannotDetermine = false
			err = nil
			return
		}

	}


	// Ok - We know the testPathStr does NOT exist on disk

	if strings.Contains(testPathStr, "...") {
		// This is an INVALID Path String
		isPathStr = false
		cannotDetermine = false
		err = fmt.Errorf("Error: INVALID PATH String! testPathStr='%v' ", testPathStr)
		return
	}

	volName := fp.VolumeName(testPathStr)

	if testPathStr == volName {
		isPathStr = true
		cannotDetermine = false
		err = nil
		return
	}

	_, checkPathIsEmpty, err2 := fh.GetPathFromPathFileName(testPathStr)

	if err2 != nil {
		isPathStr = false
		cannotDetermine = false
		err = fmt.Errorf(ePrefix + "fh.GetPathFromPathFileName(testPathStr) returned error. testPathStr='%v' Error='%v'",testPathStr, err2.Error())
		return
	}

	if checkPathIsEmpty {
		isPathStr = false
		cannotDetermine = false
		err = nil
		return
	}

	slashIdxs, err2 := fh.GetPathSeparatorIndexesInPathStr(testPathStr)

	if err2!=nil {
		isPathStr = false
		cannotDetermine = false
		err = fmt.Errorf(ePrefix + "fh.GetPathSeparatorIndexesInPathStr(testPathStr) returned error. testPathStr='%v' Error='%v'",testPathStr, err2.Error())
		return
	}

	dotIdxs, err2 := fh.GetDotSeparatorIndexesInPathStr(testPathStr)

	if err2!=nil {
		isPathStr = false
		cannotDetermine = true // Uncertain outcome. Cannot determine if this is a path string
		err = fmt.Errorf(ePrefix + "fh.GetDotSeparatorIndexesInPathStr(testPathStr) retured error. testPathStr='%v' Error='%v'",testPathStr, err2.Error())
		return
	}


	firstNonSepCharIdx, lastNonSepCharIdx, err2 := fh.GetFirstLastNonSeparatorCharIndexInPathStr(testPathStr)


	if err2!= nil {
		isPathStr = false
		cannotDetermine = true // Uncertain outcome.
		err = fmt.Errorf(ePrefix + "fh.GetFirstLastNonSeparatorCharIndexInPathStr(testPathStr) retured error. testPathStr='%v' Error='%v'",testPathStr, err2.Error())
		return
	}

	if firstNonSepCharIdx == -1 || lastNonSepCharIdx == -1 {
		// All the characters are separator characters.
		isPathStr = true
		cannotDetermine = false // High confidence in result
		err = nil
		return
	}


	// *******************************
	// From here on fristNonSepCharIdx
	// and lastNonSepCharIdx MUST be
	// greater than -1
	// *******************************

	lenDotIdx := len(dotIdxs)
	lenSlashIdx := len(slashIdxs)

	// Address case "../common"
	if strings.HasPrefix(testPathStr,"..") {

		if lenDotIdx == 2 {
			isPathStr = true
			cannotDetermine = false // High confidence in result
			err = nil
			return
		}

	}


	if strings.HasPrefix(testPathStr,".") {

		if lenDotIdx == 1 {
			isPathStr = true
			cannotDetermine = false // High confidence in result
			err = nil
			return
		}

	}


	if lenDotIdx == 0 && lenSlashIdx == 0 {
		// Just text name with no path separators
		// and no dots. This is not a Path
		isPathStr = false
		cannotDetermine = false // High confidence in result
		err = nil
		return
	}

	// Address Case No Slashes only Dots
	if lenDotIdx > 0 && lenSlashIdx == 0 {

		// .common
		if dotIdxs[lenDotIdx-1] < firstNonSepCharIdx {
			isPathStr = true
			cannotDetermine = false // High confidence in result
			err = nil
			return
		}

		// common.
		if dotIdxs[lenDotIdx-1] > firstNonSepCharIdx {
			isPathStr = false
			cannotDetermine = false // High confidence in result
			err = nil
			return
		}

	}

	// Address Case No Dots only slashes
	if lenDotIdx == 0 && lenSlashIdx > 0 {

			isPathStr = true
			cannotDetermine = false // High confidence in result
			err = nil
			return

	}

	// ***********************************
	// Both lenDotIdx and lenSlashIdx are
	// greater than zero. Therefore, both
	// path separators and dot separators
	// are present.
	// ***********************************

	// Address Case: PathSeparator at end of PathStr
	if 	lTestPathStr - 1 == slashIdxs[lenSlashIdx-1]{

					// There is a slash at the end of the path string.
					// This is definitely a path string.
					isPathStr = true
					cannotDetermine = false // High confidence in result
					err = nil
					return

	}


	// Address Case where last dot comes after last path separator
	if 	dotIdxs[lenDotIdx - 1] > firstNonSepCharIdx &&
				dotIdxs[lenDotIdx - 1] > slashIdxs[lenSlashIdx-1] {

						// This is a PathFileName string - NOT a PathStr
						isPathStr = false
						cannotDetermine = false // High confidence in result
						err = nil
						return

	}

	// Address case where last path separator comes after
	// the last dot.
	if slashIdxs[lenSlashIdx-1] > dotIdxs[lenDotIdx-1] {
					// This is a PathStr
					isPathStr = true
					cannotDetermine = false // High confidence in result
					err = nil
					return
	}

	// Address case "/common/xray.txt"
	if	lastNonSepCharIdx > dotIdxs[lenDotIdx-1] &&
				lastNonSepCharIdx > slashIdxs[lenDotIdx-1] &&
					dotIdxs[lenDotIdx-1] > slashIdxs[lenDotIdx-1] {

							isPathStr = false
							cannotDetermine = false // High confidence in result
							err = nil
							return
	}


	// Address case "..\dirA\dirB\xray"
	// In this method we will assume that
	// xray is a directory

	if dotIdxs[lenDotIdx-1] < slashIdxs[lenDotIdx-1] &&
		slashIdxs[lenDotIdx-1] < lastNonSepCharIdx {

			isPathStr = true
			cannotDetermine = true // Can't be 100% certain that xray
														 // is not a file name.
			err = nil
			return
	}


	// Can't be certain what this string is.
	// could be either directory path or
	// directory path and file name. Let
	// calling method make the call.
	// Example D:\\DirA\\common
	// Is common a file name or a directory name.
	isPathStr = false
	cannotDetermine = true
	err = nil
	return
}


// IsPathFileString - Returns 'true' if the it is determined that
// input parameter, 'pathFileStr', represents a directory path, 
// file name and optionally, a file extension.
// 
// If 'pathFileStr' is judged to be a directory path and file name,
// by definition it cannot be solely a directory path.
//
// Input Parameter:
// ================
//
// pathFileStr string		- The string to be analyzed.
//
// Return Values:
// ==============
//
// isPathFileStr bool		- A boolean indicating whether the input parameter
//												'pathFileStr' is in fact both a directory path and file name.
//
// cannotDetermine bool	- A boolean value indicating whether the method could or
//												could NOT determine whether input parameter 'pathFileStr' 
//												is a valid directory path and file name.
//		
//												'cannotDetermine' will be set to 'true' if 'pathFileStr'
// 												does not currently exist on disk and 'pathFileStr' is formatted
//												like the following example:
//														"D:\\dirA\\common"
//												In this example, the method cannot determine if 'common'
//												is a file name or a directory name. 
// 
// testPathFileStr string	- Input parameter 'pathFileStr' is subjected to cleaning routines
//												designed to exclude extraneous characters from the analysis.
//												'testPathFileStr' is the actual string on which the analysis was
//												performed.
// 
// 
// err error						- If an error is encountered during processing, it is returned here.
//
func (fh FileHelper) IsPathFileString(pathFileStr string) (isPathFileStr bool, cannotDetermine bool, testPathFileStr string, err error) {
	
	ePrefix := "FileHelper.IsPathFileString() "

	if len(pathFileStr) == 0 {
		isPathFileStr = false
		cannotDetermine = false // High confidence in result
		testPathFileStr = ""
		err = errors.New(ePrefix + "Error - Zero Length input parameter 'pathFileStr'.")
		return
	}

	testPathFileStr = fp.FromSlash(pathFileStr)
	lTestPathStr := len(testPathFileStr)

	if lTestPathStr == 0 {
		isPathFileStr = false
		cannotDetermine = false // High confidence in result
		testPathFileStr = ""
		err = fmt.Errorf(ePrefix + "Error - fp.Clean(fp.FromSlash(pathFileStr)) yielded a Zero Length String. pathFileStr='%v'", pathFileStr)
		return
	}

	// See if path actually exists on disk and
	// then examine the File Info object returned.
	fInfo, err2 := os.Stat(testPathFileStr)

	if err2==nil  {

		if !fInfo.IsDir() {
			isPathFileStr = true
			cannotDetermine = false // High confidence in result
			err = nil
			return

		} else {
			isPathFileStr = false
			cannotDetermine = false // High confidence in result
			err = nil
			return
		}

	}

	// Ok - We know the testPathFileStr does NOT exist on disk

	if strings.Contains(testPathFileStr, "...") {
		isPathFileStr = false
		cannotDetermine = false // High confidence in result
		err = fmt.Errorf(ePrefix + "Error: INVALID PATH STRING! testPathFileStr='%v'", testPathFileStr)
		return

	}


	firstCharIdx, lastCharIdx, err2 := fh.GetFirstLastNonSeparatorCharIndexInPathStr(testPathFileStr)

	if err2!=nil {
		isPathFileStr = false
		cannotDetermine = false // High confidence in result
		err = fmt.Errorf(ePrefix + "Error returned from fh.GetFirstLastNonSeparatorCharIndexInPathStr(testPathFileStr) testPathFileStr='%v'  Error='%v'", testPathFileStr, err2.Error())
		return
	}

	if firstCharIdx==-1 || lastCharIdx == -1 {
		// The pathfilestring contains no alpha numeric characters.
		// Therefore, it does NOT contain a file name!
		isPathFileStr = false
		cannotDetermine = false // High confidence in result
		err = nil
		return
	}


	volName := fp.VolumeName(testPathFileStr)

	if volName == testPathFileStr {
		// This is a volume name not a file Name!
		isPathFileStr = false
		cannotDetermine = false // High confidence in result
		err = nil
		return
	}


	slashIdxs, err2 := fh.GetPathSeparatorIndexesInPathStr(testPathFileStr)

	if err2!=nil {
		isPathFileStr = false
		cannotDetermine = true
		err = fmt.Errorf(ePrefix + "fh.GetPathSeparatorIndexesInPathStr(testPathFileStr) returned error. testPathFileStr='%v' Error='%v'",testPathFileStr, err2.Error())
		return
	}

	dotIdxs, err2 := fh.GetDotSeparatorIndexesInPathStr(testPathFileStr)

	if err2!=nil {
		isPathFileStr = false
		cannotDetermine = true // Uncertain outcome. Cannot determine if this is a path string
		err = fmt.Errorf(ePrefix + "fh.GetDotSeparatorIndexesInPathStr(testPathFileStr) retured error. testPathFileStr='%v' Error='%v'",testPathFileStr, err2.Error())
		return
	}

	lenDotIdx := len(dotIdxs)

	lenSlashIdx := len(slashIdxs)

	if lenSlashIdx == 0 {

		isPathFileStr = false
		cannotDetermine = false // high degree of certainty
		err = nil
		return

	}

	// We know the string contains one or more path separators

	if lenDotIdx == 0 {

		if lastCharIdx > slashIdxs[lenSlashIdx-1] {
			// Example D:\dir1\dir2\xray

			isPathFileStr = true
			cannotDetermine = true // Maybe, really can't tell if xray is a directory or a file!
			err = nil
			return

		}

		isPathFileStr = false
		cannotDetermine = false // high degree of certainty
		err = nil
		return

	}

	// We know that the test string contains both path separators and
	// dot separators ('.')

	if dotIdxs[lenDotIdx-1] > slashIdxs[lenSlashIdx-1] &&
		lastCharIdx > slashIdxs[lenSlashIdx-1] {
		isPathFileStr = true
		cannotDetermine = false // high degree of certainty
		err = nil
		return
	}


	// Check to determine if last character in testPathFileStr is a PathSeparator
	if slashIdxs[lenSlashIdx-1] == lTestPathStr - 1 {
		// Yes, last char in testPathFileStr is a PathSeparator. This must be a directory.
		isPathFileStr = false
		cannotDetermine = false // high degree of certainty
		err = nil
		return
	}

	// Cannot be certain of the result.
	// Don't know for sure what this string is
	isPathFileStr = false
	cannotDetermine = true
	err = nil
	return
}

// JoinPathsAdjustSeparators - Joins two
// path strings and standardizes the
// path separators according to the
// current operating system.
func (fh FileHelper) JoinPathsAdjustSeparators(p1 string, p2 string) string {
	ps1 := fp.FromSlash(fp.Clean(p1))
	ps2 := fp.FromSlash(fp.Clean(p2))
	return fp.Clean(fp.FromSlash(path.Join(ps1, ps2)))

}

// JoinPaths - correctly joins 2-paths
func (fh FileHelper) JoinPaths(p1 string, p2 string) string {

	return fp.Clean(path.Join(fp.Clean(p1), fp.Clean(p2)))

}

// MakeAbsolutePath - Supply a relative path or any path
// string and resolve that path to an Absolute Path.
// Note: Clean() is called on result by fp.Abs().
func (fh FileHelper) MakeAbsolutePath(relPath string) (string, error) {

	ePrefix := "FileHelper.MakeAbsolutePath() "

	if len(relPath) == 0 {
		return "", errors.New(ePrefix + "Error: Input Parameter 'relPath' is an EMPTY string!")
	}

	testRelPath := fh.AdjustPathSlash(relPath)

	if len(testRelPath) == 0 {
		return "", errors.New(ePrefix + "Error: Input Parameter 'relPath' adjusted for Path Separators is an EMPTY string!")
	}

	p, err := fp.Abs(testRelPath)

	if err != nil {
		return "Invalid p!", fmt.Errorf(ePrefix + "Error returned from  fp.Abs(testRelPath). testRelPath='%v'  Error='%v'", testRelPath, err.Error())
	}

	return p, err
}

// MakeDirAll - creates a directory named path,
// along with any necessary parents, and returns nil,
// or else returns an error. The permission bits perm
// are used for all directories that MkdirAll creates.
// If path is already a directory, MkdirAll does nothing
// and returns nil.
func (fh FileHelper) MakeDirAll(dirPath string) error {
	var ModePerm os.FileMode = 0777
	return os.MkdirAll(dirPath, ModePerm)
}

// MakeDir - Makes a directory. Returns
// boolean value of false plus error if
// the operation fails. If successful,
// the function returns true.
func (fh FileHelper) MakeDir(dirPath string) (bool, error) {
	var ModePerm os.FileMode = 0777
	err := os.Mkdir(dirPath, ModePerm)

	if err != nil {
		return false, err
	}

	return true, nil
}

// MoveFile - Copies file from source to destination and, if
// successful, then deletes the original source file.
//
// The copy procedure will first attempt to the the 'Copy By Link' technique.
// See FileHelper.CopyFileByLink().  If this fails, the method will seamlessly
// attempt to copy the file the source file to the destination file by means
// of writing the contents of the source file to a newly created destination
// file. Reference Method FileHelper.CopyToNewFile().
//
// If an error is encountered during this procedure it will be by means of the
// return parameter 'err'.
//
// A boolean value is also returned. If 'copyByLink' is 'true', it signals that
// the move operation was accomplished using the 'CopyFileByLink' technique. If
// the return parameter 'copyByLink' is 'false', it signals that the 'CopyToNewFile'
// technique was used.
//
func (fh FileHelper) MoveFile(src, dst string) (copyByLink bool, err error) {
	ePrefix := "FileHelper.MoveFile() "
	copyByLink = true

	if len(src) == 0 {
		err = errors.New(ePrefix + "Error: Input parameter 'src' is ZERO length string!")
		return
	}

	if len(dst) == 0 {
		err = errors.New(ePrefix + "Error: Input parameter 'dst' is a ZERO length string!")
		return
	}

	if !fh.DoesFileExist(src) {
		err = fmt.Errorf(ePrefix + "Error: Input parameter 'src' file DOES NOT EXIST! src='%v'", src)
		return
	}

	err2 := fh.CopyFileByLink(src, dst)

	if err2 != nil {
		copyByLink = false

		err2 = fh.CopyToNewFile(src, dst)

		if err2 != nil {

			err = fmt.Errorf(ePrefix + "Error returned from fh.CopyToNewFile(src, dst). Error='%v'", err2.Error())
		}

		return
	}

	copyByLink = true

	err2 = fh.DeleteDirFile(src)

	if err2 != nil {
		err = fmt.Errorf("Successfully copied file from source, '%v', to destination '%v'; however deletion of source file failed! Error: %v", src, dst, err2.Error())
		return
	}

	err = nil
	return
}

// OpenFileForReading - Wrapper function for os.Open() method which opens
// files on disk. 'Open' opens the named file for reading.
// If successful, methods on the returned file can be used for reading;
// the associated file descriptor has mode O_RDONLY. If there is an error,
// it will be of type *PathError. (See CreateFile() above.
func (fh FileHelper) OpenFileForReading(fileName string) (*os.File, error) {
	return os.Open(fileName)
}

// RemovePathSeparatorFromEndOfPathString - Remove Trailing Path Separator from
// a path string - if said trailing Path Separator exists.
func (fh FileHelper) RemovePathSeparatorFromEndOfPathString(pathStr string) string {
	lPathStr := len(pathStr)

	if lPathStr == 0 {
		return ""
	}

	lastChar := rune(pathStr[lPathStr-1])

	if lastChar == os.PathSeparator ||
			lastChar == '\\' 						||
				lastChar == '/' {

					if lPathStr < 2 {
						return ""
					}

					return pathStr[0:lPathStr-1]
	}

	return pathStr
}

// ReadFileBytes - Read bytes from file into a File Buffer.
func (fh FileHelper) ReadFileBytes(rFile *os.File, byteBuff []byte) (int, error) {
	return rFile.Read(byteBuff)
}

// SearchFileModeMatch - This method determines whether the file mode of the file described by input
// parameter, 'info', is match for the File Selection Criteria 'fileSelectCriteria.SelectByFileMode'.
// If the file's FileMode matches the 'fileSelectCriteria.SelectByFileMode' value, the return value,
// 'isFileModeMatch' is set to 'true'.
//
// If 'fileSelectCriteria.SelectByFileMode' is set to zero, the return value 'isFileModeSet' set to 'false'
// signaling the File Mode File Selection Criterion is NOT active.
//
// Note: Input parameter 'info' is of type os.FileInfo.  You can substitute a type 'FileInfoPlus' object
// for the 'info' parameter because 'FileInfoPlus' implements the 'os.FileInfo' interface.
//
func(fh *FileHelper) SearchFileModeMatch(info os.FileInfo,fileSelectCriteria FileSelectionCriteria) (isFileModeSet, isFileModeMatch bool, err error) {

	if fileSelectCriteria.SelectByFileMode == 0 {
		isFileModeSet = false
		isFileModeMatch = false
		err = nil
		return
	}

	if fileSelectCriteria.SelectByFileMode == info.Mode() {
		isFileModeSet = true
		isFileModeMatch = true
		err = nil
		return

	}

	isFileModeSet = true
	isFileModeMatch = false
	err = nil
	return
}

// SearchFileNewerThan - This method is called to determine whether the file described by the
// input parameter 'info' is a 'match' for the File Selection Criteria, 'fileSelectCriteria.FilesNewerThan'.
// If the file modification date time occurs after the 'fileSelectCriteria.FilesNewerThan' date time,
// the return value 'isFileNewerThanMatch' is set to 'true'.
//
// If 'fileSelectCriteria.FilesNewerThan' is set to time.Time zero ( the default or zero value for this type),
// the return value 'isFileNewerThanSet' is set to 'false' signaling that this search criterion is NOT active.
//
// Note: Input parameter 'info' is of type os.FileInfo.  You can substitute a type 'FileInfoPlus' object
// for the 'info' parameter because 'FileInfoPlus' implements the 'os.FileInfo' interface.
//
func(fh *FileHelper) SearchFileNewerThan(info os.FileInfo,fileSelectCriteria FileSelectionCriteria) (isFileNewerThanSet, isFileNewerThanMatch bool, err error) {

	isFileNewerThanSet = false
	isFileNewerThanMatch = false
	err = nil

	if fileSelectCriteria.FilesNewerThan.IsZero() {
		isFileNewerThanSet = false
		isFileNewerThanMatch = false
		err = nil
		return
	}

	if fileSelectCriteria.FilesNewerThan.Before(info.ModTime()) {
		isFileNewerThanSet = true
		isFileNewerThanMatch = true
		err = nil
		return

	}

	isFileNewerThanSet = true
	isFileNewerThanMatch = false
	err = nil

	return
}

// SearchFileOlderThan - This method is called to determine whether the file described by the
// input parameter 'info' is a 'match' for the File Selection Criteria, 'fileSelectCriteria.FilesOlderThan'.
// If the file modification date time occurs before the 'fileSelectCriteria.FilesOlderThan' date time,
// the return value 'isFileOlderThanMatch' is set to 'true'.
//
// If 'fileSelectCriteria.FilesOlderThan' is set to time.Time zero ( the default or zero value for this type),
// the return value 'isFileOlderThanSet' is set to 'false' signaling that this search criterion is NOT active.
//
// Note: Input parameter 'info' is of type os.FileInfo.  You can substitute a type 'FileInfoPlus' object
// for the 'info' parameter because 'FileInfoPlus' implements the 'os.FileInfo' interface.
//
func(fh *FileHelper) SearchFileOlderThan(info os.FileInfo,fileSelectCriteria FileSelectionCriteria) (isFileOlderThanSet, isFileOlderThanMatch bool, err error) {

	if fileSelectCriteria.FilesOlderThan.IsZero() {
		isFileOlderThanSet = false
		isFileOlderThanMatch = false
		err = nil
		return
	}

	if fileSelectCriteria.FilesOlderThan.After(info.ModTime()) {
		isFileOlderThanSet = true
		isFileOlderThanMatch = true
		err = nil
		return
	}

	isFileOlderThanSet = true
	isFileOlderThanMatch = false
	err = nil
	return

}

// SearchFilePatternMatch - used to determine whether a file described by the
// 'info' parameter meets the specified File Selection Criteria and is judged
// to be a match for the fileSelectCriteria FileNamePattern. 'fileSelectCriteria.FileNamePatterns'
// consists of a string array. If the pattern signified by any element in the string array
// is a 'match', the return value 'isPatternMatch' is set to true.
//
// If the 'fileSelectCriteria.FileNamePatterns' array is empty or if it contains only empty strings,
// the return value isPatternSet is set to 'false' signaling that the pattern file search selection
// criterion is NOT active.
//
// Note: Input parameter 'info' is of type os.FileInfo.  You can substitute a type 'FileInfoPlus' object
// for the 'info' parameter because 'FileInfoPlus' implements the 'os.FileInfo' interface.
//
func(fh *FileHelper) SearchFilePatternMatch(info os.FileInfo, fileSelectCriteria FileSelectionCriteria) (isPatternSet, isPatternMatch bool, err error) {

	ePrefix := "DirMgr.SearchFilePatternMatch()"

	isPatternMatch = false
	isPatternSet = false
	err = nil

	isPatternSet = fileSelectCriteria.ArePatternsActive()

	if !isPatternSet {
		isPatternSet = false
		isPatternMatch = false
		err = nil
		return
	}

	lPats := len(fileSelectCriteria.FileNamePatterns)

	for i:=0; i < lPats; i++ {

		matched, err2 := fp.Match(fileSelectCriteria.FileNamePatterns[i] , info.Name())

		if err2 != nil {
			isPatternSet = true
			err = fmt.Errorf(ePrefix + "Error returned from filepath.Match(fileSelectCriteria.FileNamePatterns[i] , info.Name()) fileSelectCriteria.FileNamePatterns[i]='%v' info.Name()='%v' Error='%v'", fileSelectCriteria.FileNamePatterns[i], info.Name(), err.Error() )
			isPatternMatch = false
			return
		}

		if matched {
			isPatternSet = true
			isPatternMatch = true
			err = nil
			return
		}
	}


	isPatternSet = true
	isPatternMatch = false
	err = nil
	return
}

// WriteFileStr - Wrapper for *os.File.WriteString. Writes a string
// to an open file pointed to by *os.File.
func (fh FileHelper) WriteFileStr(str string, fPtr *os.File) (int, error) {

	return fPtr.WriteString(str)

}

/*
		FileHelper private methods
 */


// makeFileHelperWalkDirDeleteFilesFunc - Used in conjunction with DirMgr.DeleteWalDirFiles to select and delete
// files residing the directory tree identified by the current DirMgr object.
func(fh *FileHelper) makeFileHelperWalkDirDeleteFilesFunc(dInfo *DirectoryDeleteFileInfo) func(string, os.FileInfo, error) error {
	return func(pathFile string, info os.FileInfo, erIn error) error {

		ePrefix := "DirMgr.makeFileHelperWalkDirDeleteFilesFunc"

		if erIn != nil {
			dInfo.ErrReturns = append(dInfo.ErrReturns, erIn.Error())
			return nil
		}

		if info.IsDir() {

			subDir, err := DirMgr{}.New(pathFile)

			if err !=nil {
				ex := fmt.Errorf(ePrefix + "Error returned from DirMgr{}.New(pathFile). pathFile:='%v' Error='%v'", pathFile, err.Error())

				dInfo.ErrReturns = append(dInfo.ErrReturns, ex.Error())

				if subDir.IsInitialized {
					subDir.ActualDirFileInfo = FileInfoPlus{}.NewPathFileInfo(pathFile, info)
					dInfo.Directories.AddDirMgr(subDir)
				}

				return nil
			}

			subDir.ActualDirFileInfo = FileInfoPlus{}.NewPathFileInfo(pathFile, info)
			dInfo.Directories.AddDirMgr(subDir)

			return nil
		}

		fh := FileHelper{}

		isFoundFile, err := fh.FilterFileName(info, dInfo.DeleteFileSelectCriteria)

		if err!=nil {

			ex := fmt.Errorf(ePrefix + "Error returned from dMgr.FilterFileName(info, dInfo.DeleteFileSelectCriteria) pathFile='%v' info.Name()='%v' Error='%v' ", pathFile, info.Name(), err.Error())
			dInfo.ErrReturns = append(dInfo.ErrReturns, ex.Error())
			return nil
		}

		if isFoundFile {

			err := os.Remove(pathFile)

			if err != nil {
				ex := fmt.Errorf(ePrefix + "Error returned from os.Remove(pathFile). pathFile='%v' Error='%v'", pathFile, err.Error())
				dInfo.ErrReturns = append(dInfo.ErrReturns, ex.Error())
				return nil
			}

			dInfo.DeletedFiles.AddFileInfo( pathFile,  info)
		}

		return nil
	}
}

// makeFileHelperWalkDirFindFilesFunc - This function is designed to work in conjunction
// with a walk directory function like FindWalkDirFiles. It will process
// files extracted from a 'Directory Walk' operation initiated by the 'filepath.Walk' method.
func(fh *FileHelper) makeFileHelperWalkDirFindFilesFunc(dInfo *DirectoryTreeInfo) func(string, os.FileInfo, error) error {
	return func(pathFile string, info os.FileInfo, erIn error) error {

		ePrefix := "DirMgr.makeFileHelperWalkDirFindFilesFunc() "

		if erIn != nil {
			ex2 := fmt.Errorf(ePrefix + "Error returned from directory walk function. pathFile= '%v' Error='%v'", pathFile, erIn.Error())
			dInfo.ErrReturns = append(dInfo.ErrReturns, ex2.Error())
			return nil
		}

		if info.IsDir() {
			subDir, err := DirMgr{}.NewFromFileInfo(pathFile, info)
			if err != nil {

				if subDir.IsInitialized {
					dInfo.Directories.AddDirMgr(subDir)
				}

				er2 := fmt.Errorf(ePrefix + "Error returned by DirMgr{}.New(pathFile). pathFile='%v' Error='%v'", pathFile, err.Error() )
				dInfo.ErrReturns = append(dInfo.ErrReturns, er2.Error())
				return nil
			}

			dInfo.Directories.AddDirMgr(subDir)
			return nil
		}

		fh := FileHelper{}

		// This is not a directory. It is a file.
		// Determine if it matches the find file criteria.
		isFoundFile, err := fh.FilterFileName(info, dInfo.FileSelectCriteria )

		if err != nil {

			er2 := fmt.Errorf(ePrefix + "Error returned from dMgr.FilterFileName(info, dInfo.FileSelectCriteria) pathFile='%v' info.Name()='%v' Error='%v' ", pathFile, info.Name(), err.Error())
			dInfo.ErrReturns = append(dInfo.ErrReturns, er2.Error())
			return nil
		}

		if isFoundFile {
			dInfo.FoundFiles.AddFileInfo( pathFile,  info)
		}

		return nil
	}
}
