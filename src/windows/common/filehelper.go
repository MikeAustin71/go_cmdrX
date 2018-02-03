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
This source code file is located in the source code repository:
https://github.com/MikeAustin71/pathfilego.git
*/

type FileWalkInfo struct {
	Path string
	Info os.FileInfo
}

// DirWalkInfo - structure used to
// encapsulate information gleaned from
// 'walking' a directory tree using
// filepath.Walk(root string, walkFn WalkFunc) error
type DirWalkInfo struct {
	StartPath            string
	Directories          []string
	FoundFiles           []FileWalkInfo
	ErrReturns           []string
	PatternMatch         string
	DeleteFilesOlderThan time.Time
	DeletedFiles         []FileWalkInfo
}

// FileHelper - structure used
// to encapsulate FileHelper utility
// methods.
type FileHelper struct {
	IsInitialized                   bool
	OriginalPathFileName            string
	AbsolutePath                    string
	AbsolutePathIsPopulated         bool
	AbsolutePathDoesExist           bool
	AbsolutePathDifferentFromPath   bool
	AbsolutePathFileName            string
	AbsolutePathFileNameIsPopulated bool
	AbsolutePathFileNameDoesExist   bool
	Path                            string
	PathIsPopulated                 bool
	PathDoesExist                   bool
	FileName                        string
	FileNameIsPopulated             bool
	FileExt                         string
	FileExtIsPopulated              bool
	FileNameExt                     string
	FileNameExtIsPopulated          bool
	VolumeName                      string
	VolumeIsPopulated               bool
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

// CopyFile - Copies file from source path & File Name
// to destination path & File Name.
// See: https://stackoverflow.com/questions/21060945/simple-way-to-copy-a-file-in-golang
func (fh FileHelper) CopyFile(src, dst string) (err error) {

	sfi, err := os.Stat(src)
	if err != nil {
		return
	}
	if !sfi.Mode().IsRegular() {
		// cannot copy non-regular files (e.g., directories,
		// symlinks, devices, etc.)
		return fmt.Errorf("FileHelper:CopyFile(): non-regular source file %s (%q)", sfi.Name(), sfi.Mode().String())
	}

	dfi, err := os.Stat(dst)

	if err != nil {

		if !os.IsNotExist(err) {
			// Must be PathError - Path does not exist
			return fmt.Errorf("FileHelper:CopyFile(): Destination File Path Error - Path does NOT exist. Error: %v", err.Error())
		}
	} else {
		if !(dfi.Mode().IsRegular()) {
			return fmt.Errorf("FileHelper:CopyFile(): non-regular destination file %s (%q) - Cannot Overwrite destination file.", dfi.Name(), dfi.Mode().String())
		}
		if os.SameFile(sfi, dfi) {
			return nil
		}
	}
	if err = os.Link(src, dst); err == nil {
		return
	}

	err = fh.CopyFileContents(src, dst)
	return
}

// CopyFileContents - Copies file contents from source to destination file.
// Note: No validity checks are performed on 'src' and 'dest' files.
// This method is called by FileHelper:CopyFile(). Use FileHelper:CopyFile() for
// ordinary file copy operations since it provides validity checks on 'src' and 'dest'
// files.
// See: https://stackoverflow.com/questions/21060945/simple-way-to-copy-a-file-in-golang
func (fh FileHelper) CopyFileContents(src, dst string) (err error) {
	in, err := os.Open(src)
	if err != nil {
		return
	}
	defer in.Close()
	out, err := os.Create(dst)
	if err != nil {
		return
	}

	defer func() {
		cerr := out.Close()
		if err == nil {
			err = cerr
		}
	}()

	if _, err = io.Copy(out, in); err != nil {
		return
	}
	err = out.Sync()

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

// CopyToThis - Copies a second FileHelper data structure
// to the current FileHelper data structure
func (fh *FileHelper) CopyToThis(fh2 FileHelper) {

	fh.IsInitialized = fh2.IsInitialized
	fh.OriginalPathFileName = fh2.OriginalPathFileName
	fh.AbsolutePath = fh2.AbsolutePath
	fh.AbsolutePathIsPopulated = fh2.AbsolutePathIsPopulated
	fh.AbsolutePathDoesExist = fh2.AbsolutePathDoesExist
	fh.AbsolutePathDifferentFromPath = fh2.AbsolutePathDifferentFromPath
	fh.AbsolutePathFileName = fh2.AbsolutePathFileName
	fh.AbsolutePathFileNameIsPopulated = fh2.AbsolutePathFileNameIsPopulated
	fh.AbsolutePathFileNameDoesExist = fh2.AbsolutePathFileNameDoesExist
	fh.Path = fh2.Path
	fh.PathIsPopulated = fh2.PathIsPopulated
	fh.PathDoesExist = fh2.PathDoesExist
	fh.FileName = fh2.FileName
	fh.FileNameIsPopulated = fh2.FileNameIsPopulated
	fh.FileExt = fh2.FileExt
	fh.FileExtIsPopulated = fh2.FileExtIsPopulated
	fh.FileNameExt = fh2.FileNameExt
	fh.FileNameExtIsPopulated = fh2.FileNameExtIsPopulated
	fh.VolumeName = fh2.VolumeName
	fh.VolumeIsPopulated = fh2.VolumeIsPopulated

	return
}

// CreateDirAndFile - Performs two operations:
// 1. This is a Wrapper function for os.Create - Create a file.
// 2. If the home directory does not currently exist, this method
// 		will first create the directory tree, before creating the new
//		file.
func (fh FileHelper) CreateDirAndFile(pathFileName string) (*os.File, error) {

	ePrefix := "FileHelper:CreateDirAndFile() Error - "

	if pathFileName == "" {
		return nil, errors.New(ePrefix + "pathFileName is EMPTY!")
	}

	fh2, err := fh.GetPathFileNameElements(pathFileName)

	if err != nil {
		return nil, errors.New(ePrefix + " from FileHelper:GetPathFileNameElements(): " + err.Error())
	}

	if !fh2.AbsolutePathIsPopulated {
		return nil, errors.New(ePrefix + " from FileHelper:GetPathFileNameElements(): Could NOT extract absolute path/directory!")
	}

	if !fh2.AbsolutePathDoesExist {
		// Directory does NOT exist, create it!
		err := fh.MakeDirAll(fh2.AbsolutePath)

		if err != nil {
			return nil, errors.New(ePrefix + " from FileHelper:MakeDirAll(absolutePath) Error: " + err.Error())
		}

	}

	// The path now exists - create the file.
	if !fh2.AbsolutePathFileNameIsPopulated {
		return nil, errors.New(ePrefix + " from FileHelper:GetPathFileNameElements(): Could NOT extract absolute path and file name!")
	}

	return fh.CreateFile(fh2.AbsolutePathFileName)

}

// CreateFile - Wrapper function for os.Create
func (fh FileHelper) CreateFile(pathFileName string) (*os.File, error) {
	return os.Create(pathFileName)
}

// DeleteDirFile - Wrapper function for Remove.
// Remove removes the named file or directory.
// If there is an error, it will be of type *PathError.
func (fh FileHelper) DeleteDirFile(pathFile string) error {
	return os.Remove(pathFile)
}

// DeleteDirPathAll - Wrapper function for RemoveAll
// RemoveAll removes path and any children it contains.
// It removes everything it can but returns the first
// error it encounters. If the path does not exist,
// RemoveAll returns nil (no error).
func (fh FileHelper) DeleteDirPathAll(path string) error {
	return os.RemoveAll(path)
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

// Equal - Compares a second FileHelper data structure
// to the current FileHelpter data structure and returns
// a boolean value indicating whether they are equal
// in all respects.
func (fh *FileHelper) Equal(fh2 FileHelper) bool {

	if fh.IsInitialized != fh2.IsInitialized ||
		fh.OriginalPathFileName != fh2.OriginalPathFileName ||
		fh.AbsolutePath != fh2.AbsolutePath ||
		fh.AbsolutePathIsPopulated != fh2.AbsolutePathIsPopulated ||
		fh.AbsolutePathDoesExist != fh2.AbsolutePathDoesExist ||
		fh.AbsolutePathDifferentFromPath != fh2.AbsolutePathDifferentFromPath ||
		fh.AbsolutePathFileNameIsPopulated != fh2.AbsolutePathFileNameIsPopulated ||
		fh.AbsolutePathFileNameDoesExist != fh2.AbsolutePathFileNameDoesExist ||
		fh.AbsolutePathFileName != fh2.AbsolutePathFileName ||
		fh.Path != fh2.Path ||
		fh.PathIsPopulated != fh2.PathIsPopulated ||
		fh.PathDoesExist != fh2.PathDoesExist ||
		fh.FileName != fh2.FileName ||
		fh.FileNameIsPopulated != fh2.FileNameIsPopulated ||
		fh.FileExt != fh2.FileExt ||
		fh.FileExtIsPopulated != fh2.FileExtIsPopulated ||
		fh.FileNameExt != fh2.FileNameExt ||
		fh.FileNameExtIsPopulated != fh2.FileNameExtIsPopulated ||
		fh.VolumeName != fh2.VolumeName ||
		fh.VolumeIsPopulated != fh2.VolumeIsPopulated {
		return false
	}

	return true
}

// GetAbsPathFromFilePath - Supply a string containing both
// the path file name and extension and return the path
// element.
func (fh FileHelper) GetAbsPathFromFilePath(filePath string) (string, error) {

	return fh.MakeAbsolutePath(path.Dir(filePath))

}

// GetAbsCurrDir - returns
// the absolute path of the
// current working directory
func (fh FileHelper) GetAbsCurrDir() (string, error) {

	dir, err := fh.GetCurrentDir()

	if err != nil {
		return dir, err
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

// GetExecutablePathFileName - Gets the file name
// and path of the executable that started the
// current process
func (fh FileHelper) GetExecutablePathFileName() (string, error) {
	ex, err := os.Executable()

	return ex, err

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

// GetPathAndFileNameExt - Breaks out path and FileName+Ext elements from
// a path string. If both path and fileName are empty strings, this method
// returns an error.
func (fh FileHelper) GetPathAndFileNameExt(pathFileNameExt string) (path, fileNameExt string, bothAreEmpty bool) {

	bothAreEmpty = true
	if pathFileNameExt == "" {
		return "", "", bothAreEmpty
	}

	path, fileNameExt = fp.Split(pathFileNameExt)

	if path == "" && fileNameExt == "" {
		return path, fileNameExt, bothAreEmpty
	}

	bothAreEmpty = false

	return path, fileNameExt, bothAreEmpty

}

// GetPathFromPathFileName - Returns the path from a path and file name string.
// If the returned path is an empty string, return parameter 'isEmpty' is set to
// 'true'
func (fh FileHelper) GetPathFromPathFileName(pathFileNameExt string) (path string, isEmpty bool) {
	path = ""
	isEmpty = true

	if pathFileNameExt == "" {
		return path, isEmpty
	}

	path, _ = fp.Split(pathFileNameExt)

	if path != "" {
		isEmpty = false
	}

	return path, isEmpty

}

// GetFileNameWithoutExt - returns the file name
// without the path or extension. If the returned
// File Name is an empty string, isEmpty is set to true.
func (fh FileHelper) GetFileNameWithoutExt(pathFileNameExt string) (fName string, isEmpty bool) {
	isEmpty = true
	fName = ""

	if pathFileNameExt == "" {
		return fName, isEmpty
	}

	_, f := fp.Split(pathFileNameExt)

	if f == "" {
		return fName, isEmpty
	}

	sArray := strings.Split(f, ".")

	if len(sArray) == 0 {
		return fName, isEmpty
	}

	isEmpty = false
	fName = sArray[0]

	return fName, isEmpty

}

// GetFileExtension - Returns the File Extension with
// the dot. If there is no File Extension an empty
// string is returned (NO dot included). If the returned
// File Extension is an empty string, the returned
// parameter 'isEmpty' is set equal to 'true'.
func (fh FileHelper) GetFileExtension(pathFileNameExt string) (ext string, isEmpty bool) {
	isEmpty = true
	ext = fp.Ext(pathFileNameExt)

	if ext == "" {
		return ext, isEmpty
	}

	isEmpty = false

	return ext, isEmpty
}

// GetPathFileNameElements - Parses out the path, file name and file extension
func (fh FileHelper) GetPathFileNameElements(pathFileNameExt string) (FileHelper, error) {

	var s string
	var pathIsEmpty bool
	var err error

	fInfo := FileHelper{}

	if pathFileNameExt == "" {
		return fInfo, errors.New("FileHelper:GetPathFileNameElements()-Error: pathFileNameExt is Empty!")
	}

	adjustedPathFileNameExt := fh.AdjustPathSlash(pathFileNameExt)
	s, pathIsEmpty = fh.GetPathFromPathFileName(adjustedPathFileNameExt)

	fInfo.IsInitialized = true
	fInfo.OriginalPathFileName = adjustedPathFileNameExt

	if !pathIsEmpty {
		fInfo.PathIsPopulated = true
		fInfo.Path = s
		fInfo.PathDoesExist = fh.DoesFileExist(fInfo.Path)
		s, err = fh.MakeAbsolutePath(fp.Clean(s))
		if err != nil {
			return FileHelper{}, errors.New("FileHelper:GetPathFileNameElements()-Error: " + err.Error())
		}

		fInfo.AbsolutePathIsPopulated = true
		fInfo.AbsolutePath = fp.Clean(s)
		fInfo.AbsolutePathDoesExist = fh.DoesFileExist(fInfo.AbsolutePath)
		if fInfo.AbsolutePath != fInfo.Path {
			fInfo.AbsolutePathDifferentFromPath = true
		}

		s = ""

		if fInfo.AbsolutePathIsPopulated {
			s = fp.VolumeName(fInfo.AbsolutePath)
		} else if fInfo.PathIsPopulated {
			s = fp.VolumeName(fInfo.Path)
		}

		if s != "" {
			fInfo.VolumeIsPopulated = true
			fInfo.VolumeName = s
		}

	}

	s, pathIsEmpty = fh.GetFileNameWithoutExt(adjustedPathFileNameExt)

	if !pathIsEmpty {
		fInfo.FileNameIsPopulated = true
		fInfo.FileName = s
	}

	s, pathIsEmpty = fh.GetFileExtension(adjustedPathFileNameExt)

	if !pathIsEmpty {
		fInfo.FileExtIsPopulated = true
		fInfo.FileExt = s
	}

	if fInfo.FileNameIsPopulated {
		fInfo.FileNameExtIsPopulated = true
		fInfo.FileNameExt = fInfo.FileName + fInfo.FileExt
	}

	fInfo.AbsolutePathFileName = fh.JoinPathsAdjustSeparators(fInfo.AbsolutePath, fInfo.FileNameExt)
	fInfo.AbsolutePathFileNameIsPopulated = true
	fInfo.AbsolutePathFileNameDoesExist = fh.DoesFileExist(fInfo.AbsolutePathFileName)

	return fInfo, nil
}

// IsAbsolutePath - Wrapper function for path.IsAbs()
// https://golang.org/pkg/path/#IsAbs
// This method reports whether the input parameter is
// an absolute path.
func (fh FileHelper) IsAbsolutePath(pathStr string) bool {
	return path.IsAbs(pathStr)
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

	p, err := fp.Abs(relPath)

	if err != nil {
		return "Invalid p!", err
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
func (fh FileHelper) MoveFile(src, dst string) (err error) {

	err = fh.CopyFile(src, dst)

	if err != nil {
		return
	}

	err = fh.DeleteDirFile(src)

	if err != nil {
		return fmt.Errorf("FileHelper:MoveFile() Successfully copied file from source, '%v', to destination '%v'; however deletion of source file failed! Error: %v", src, dst, err.Error())
	}

	return
}

// OpenFile - Wrapper function for os.Open() method which opens
// files on disk. Open opens the named file for reading.
// If successful, methods on the returned file can be used for reading;
// the associated file descriptor has mode O_RDONLY. If there is an error,
// it will be of type *PathError. (See CreateFile() above.
func (fh FileHelper) OpenFile(fileName string) (*os.File, error) {
	return os.Open(fileName)
}

// WalkDirPurgeFilesOlderThan - Walks a directory tree specified by
// input parameter dInfo.StartPath and purges (a.k.a. 'Deletes) all
// files older than t.Time, dInfo.DeleteFilesOlderThan. Note: The
// list of file eligible for deletion can be further constrained
// through use of a file pattern (i.e. "*.log") which is stored
// in input parameter 'dInfo.PatternMatch'.
// BE CAREFUL, THIS METHOD DELETES FILES!!!
func (fh FileHelper) WalkDirPurgeFilesOlderThan(dInfo *DirWalkInfo) error {

	err := fp.Walk(dInfo.StartPath, fh.MakePurgeFilesFunc(dInfo))

	if err != nil {

		return err
	}

	return nil
}

// WalkDirGetFileInfo - Walks a directory tree specified by input
// parameter, 'dInfo.StartPath'. All file information found in
// the directory tree is then stored and returned through
// (dInfo *DirWalkInfo).
func (fh FileHelper) WalkDirGetFileInfo(dInfo *DirWalkInfo) error {

	err := fp.Walk(dInfo.StartPath, fh.MakeWalkDirGetFilesFunc(dInfo))

	if err != nil {

		return err
	}

	return nil
}

// MakeWalkDirGetFilesFunc - Creates a function used in conjunction with
// method FileHelper.WalkDirGetFileInfo(), above. The signature of the
// method below is designed to coordinate with filepath.Walk() method.
// See https://golang.org/pkg/path/filepath/#Walk
func (fh FileHelper) MakeWalkDirGetFilesFunc(dInfo *DirWalkInfo) func(string, os.FileInfo, error) error {
	return func(path string, info os.FileInfo, erIn error) error {

		if erIn != nil {
			dInfo.ErrReturns = append(dInfo.ErrReturns, erIn.Error())
		}

		if info.IsDir() {
			dInfo.Directories = append(dInfo.Directories, path)
		} else if dInfo.PatternMatch != "" {

			matched, err := fp.Match(dInfo.PatternMatch, info.Name())

			if err != nil {
				panic(err)
			}

			if matched {
				dInfo.FoundFiles = append(dInfo.FoundFiles, FileWalkInfo{Path: path, Info: info})
			}

		} else {
			dInfo.FoundFiles = append(dInfo.FoundFiles, FileWalkInfo{Path: path, Info: info})
		}

		return nil
	}
}

// MakePurgeFilesFunc - Creates a function used in conjunction with
// method FileHelper.WalkDirPurgeFilesOlderThan(), above. The signature of the
// method below is designed to coordinate with the filepath.Walk() method.
//          See https://golang.org/pkg/path/filepath/#Walk
// When used with FileHelper.WalkDirPurgeFilesOlderThan(), the method below may
// be used to delete old files.
func (fh FileHelper) MakePurgeFilesFunc(dInfo *DirWalkInfo) func(string, os.FileInfo, error) error {
	return func(path string, info os.FileInfo, _ error) error {

		if info.IsDir() {
			return nil
		}

		if dInfo.PatternMatch != "" {

			matched, err := fp.Match(dInfo.PatternMatch, info.Name())

			if err != nil {
				panic(err)
			}

			if matched {
				if info.ModTime().Before(dInfo.DeleteFilesOlderThan) {

					err := os.Remove(path)

					if err != nil {
						dInfo.ErrReturns = append(dInfo.ErrReturns, err.Error())
					} else {
						// No Errors! - Successful File Deletion Confirmed
						dInfo.DeletedFiles = append(dInfo.DeletedFiles, FileWalkInfo{Path: path, Info: info})
					}
				}
			}
		} else {

			if info.ModTime().Before(dInfo.DeleteFilesOlderThan) {
				err := os.Remove(path)

				if err != nil {
					dInfo.ErrReturns = append(dInfo.ErrReturns, err.Error())
				} else {
					// No Errors! - Successful File Deletion Confirmed
					dInfo.DeletedFiles = append(dInfo.DeletedFiles, FileWalkInfo{Path: path, Info: info})
				}
			}

		}

		return nil
	}
}

func (fh FileHelper) ReadFileBytes(rFile *os.File, byteBuff []byte) (int, error) {
	return rFile.Read(byteBuff)
}

// WriteFileStr - Wrapper for *os.File.WriteString. Writes a string
// to an open file pointed to by *os.File.
func (fh FileHelper) WriteFileStr(str string, fPtr *os.File) (int, error) {

	return fPtr.WriteString(str)

}
