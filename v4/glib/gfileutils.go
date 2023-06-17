// Package glib was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package glib

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
)

// Flags to pass to g_file_set_contents_full() to affect its safety and
// performance.
type FileSetContentsFlags int

const (

	// No guarantees about file consistency or durability.
	//   The most dangerous setting, which is slightly faster than other settings.
	GFileSetContentsNoneValue FileSetContentsFlags = 0
	// Guarantee file consistency: after a crash,
	//   either the old version of the file or the new version of the file will be
	//   available, but not a mixture. On Unix systems this equates to an `fsync()`
	//   on the file and use of an atomic `rename()` of the new version of the file
	//   over the old.
	GFileSetContentsConsistentValue FileSetContentsFlags = 1
	// Guarantee file durability: after a crash, the
	//   new version of the file will be available. On Unix systems this equates to
	//   an `fsync()` on the file (if %G_FILE_SET_CONTENTS_CONSISTENT is unset), or
	//   the effects of %G_FILE_SET_CONTENTS_CONSISTENT plus an `fsync()` on the
	//   directory containing the file after calling `rename()`.
	GFileSetContentsDurableValue FileSetContentsFlags = 2
	// Only apply consistency and durability
	//   guarantees if the file already exists. This may speed up file operations
	//   if the file doesn’t currently exist, but may result in a corrupted version
	//   of the new file if the system crashes while writing it.
	GFileSetContentsOnlyExistingValue FileSetContentsFlags = 4
)

// A test to perform on a file using g_file_test().
type FileTest int

const (

	// %TRUE if the file is a regular file
	//     (not a directory). Note that this test will also return %TRUE
	//     if the tested file is a symlink to a regular file.
	GFileTestIsRegularValue FileTest = 1
	// %TRUE if the file is a symlink.
	GFileTestIsSymlinkValue FileTest = 2
	// %TRUE if the file is a directory.
	GFileTestIsDirValue FileTest = 4
	// %TRUE if the file is executable.
	GFileTestIsExecutableValue FileTest = 8
	// %TRUE if the file exists. It may or may not
	//     be a regular file.
	GFileTestExistsValue FileTest = 16
)

// Values corresponding to @errno codes returned from file operations
// on UNIX. Unlike @errno codes, GFileError values are available on
// all systems, even Windows. The exact meaning of each code depends
// on what sort of file operation you were performing; the UNIX
// documentation gives more details. The following error code descriptions
// come from the GNU C Library manual, and are under the copyright
// of that manual.
//
// It's not very portable to make detailed assumptions about exactly
// which errors will be returned from a given operation. Some errors
// don't occur on some systems, etc., sometimes there are subtle
// differences in when a system will report a given error, etc.
type FileError int

const (

	// Operation not permitted; only the owner of
	//     the file (or other resource) or processes with special privileges
	//     can perform the operation.
	GFileErrorExistValue FileError = 0
	// File is a directory; you cannot open a directory
	//     for writing, or create or remove hard links to it.
	GFileErrorIsdirValue FileError = 1
	// Permission denied; the file permissions do not
	//     allow the attempted operation.
	GFileErrorAccesValue FileError = 2
	// Filename too long.
	GFileErrorNametoolongValue FileError = 3
	// No such file or directory. This is a "file
	//     doesn't exist" error for ordinary files that are referenced in
	//     contexts where they are expected to already exist.
	GFileErrorNoentValue FileError = 4
	// A file that isn't a directory was specified when
	//     a directory is required.
	GFileErrorNotdirValue FileError = 5
	// No such device or address. The system tried to
	//     use the device represented by a file you specified, and it
	//     couldn't find the device. This can mean that the device file was
	//     installed incorrectly, or that the physical device is missing or
	//     not correctly attached to the computer.
	GFileErrorNxioValue FileError = 6
	// The underlying file system of the specified file
	//     does not support memory mapping.
	GFileErrorNodevValue FileError = 7
	// The directory containing the new link can't be
	//     modified because it's on a read-only file system.
	GFileErrorRofsValue FileError = 8
	// Text file busy.
	GFileErrorTxtbsyValue FileError = 9
	// You passed in a pointer to bad memory.
	//     (GLib won't reliably return this, don't pass in pointers to bad
	//     memory.)
	GFileErrorFaultValue FileError = 10
	// Too many levels of symbolic links were encountered
	//     in looking up a file name. This often indicates a cycle of symbolic
	//     links.
	GFileErrorLoopValue FileError = 11
	// No space left on device; write operation on a
	//     file failed because the disk is full.
	GFileErrorNospcValue FileError = 12
	// No memory available. The system cannot allocate
	//     more virtual memory because its capacity is full.
	GFileErrorNomemValue FileError = 13
	// The current process has too many files open and
	//     can't open any more. Duplicate descriptors do count toward this
	//     limit.
	GFileErrorMfileValue FileError = 14
	// There are too many distinct file openings in the
	//     entire system.
	GFileErrorNfileValue FileError = 15
	// Bad file descriptor; for example, I/O on a
	//     descriptor that has been closed or reading from a descriptor open
	//     only for writing (or vice versa).
	GFileErrorBadfValue FileError = 16
	// Invalid argument. This is used to indicate
	//     various kinds of problems with passing the wrong argument to a
	//     library function.
	GFileErrorInvalValue FileError = 17
	// Broken pipe; there is no process reading from the
	//     other end of a pipe. Every library function that returns this
	//     error code also generates a 'SIGPIPE' signal; this signal
	//     terminates the program if not handled or blocked. Thus, your
	//     program will never actually see this code unless it has handled
	//     or blocked 'SIGPIPE'.
	GFileErrorPipeValue FileError = 18
	// Resource temporarily unavailable; the call might
	//     work if you try again later.
	GFileErrorAgainValue FileError = 19
	// Interrupted function call; an asynchronous signal
	//     occurred and prevented completion of the call. When this
	//     happens, you should try the call again.
	GFileErrorIntrValue FileError = 20
	// Input/output error; usually used for physical read
	//    or write errors. i.e. the disk or other physical device hardware
	//    is returning errors.
	GFileErrorIoValue FileError = 21
	// Operation not permitted; only the owner of the
	//    file (or other resource) or processes with special privileges can
	//    perform the operation.
	GFileErrorPermValue FileError = 22
	// Function not implemented; this indicates that
	//    the system is missing some functionality.
	GFileErrorNosysValue FileError = 23
	// Does not correspond to a UNIX error code; this
	//    is the standard "failed for unspecified reason" error code present
	//    in all #GError error code enumerations. Returned if no specific
	//    code applies.
	GFileErrorFailedValue FileError = 24
)

var xBasename func(string) string

// Gets the name of the file without any leading directory
// components. It returns a pointer into the given file name
// string.
func Basename(FileNameVar string) string {

	return xBasename(FileNameVar)

}

var xBuildFilename func(string, ...interface{}) string

// Creates a filename from a series of elements using the correct
// separator for filenames.
//
// On Unix, this function behaves identically to `g_build_path
// (G_DIR_SEPARATOR_S, first_element, ....)`.
//
// On Windows, it takes into account that either the backslash
// (`\` or slash (`/`) can be used as separator in filenames, but
// otherwise behaves as on UNIX. When file pathname separators need
// to be inserted, the one that last previously occurred in the
// parameters (reading from left to right) is used.
//
// No attempt is made to force the resulting filename to be an absolute
// path. If the first element is a relative path, the result will
// be a relative path.
func BuildFilename(FirstElementVar string, varArgs ...interface{}) string {

	return xBuildFilename(FirstElementVar, varArgs...)

}

var xBuildFilenameValist func(string, []interface{}) string

// Behaves exactly like g_build_filename(), but takes the path elements
// as a va_list. This function is mainly meant for language bindings.
func BuildFilenameValist(FirstElementVar string, ArgsVar []interface{}) string {

	return xBuildFilenameValist(FirstElementVar, ArgsVar)

}

var xBuildFilenamev func(uintptr) string

// Behaves exactly like g_build_filename(), but takes the path elements
// as a string array, instead of varargs. This function is mainly
// meant for language bindings.
func BuildFilenamev(ArgsVar uintptr) string {

	return xBuildFilenamev(ArgsVar)

}

var xBuildPath func(string, string, ...interface{}) string

// Creates a path from a series of elements using @separator as the
// separator between elements. At the boundary between two elements,
// any trailing occurrences of separator in the first element, or
// leading occurrences of separator in the second element are removed
// and exactly one copy of the separator is inserted.
//
// Empty elements are ignored.
//
// The number of leading copies of the separator on the result is
// the same as the number of leading copies of the separator on
// the first non-empty element.
//
// The number of trailing copies of the separator on the result is
// the same as the number of trailing copies of the separator on
// the last non-empty element. (Determination of the number of
// trailing copies is done without stripping leading copies, so
// if the separator is `ABA`, then `ABABA` has 1 trailing copy.)
//
// However, if there is only a single non-empty element, and there
// are no characters in that element not part of the leading or
// trailing separators, then the result is exactly the original value
// of that element.
//
// Other than for determination of the number of leading and trailing
// copies of the separator, elements consisting only of copies
// of the separator are ignored.
func BuildPath(SeparatorVar string, FirstElementVar string, varArgs ...interface{}) string {

	return xBuildPath(SeparatorVar, FirstElementVar, varArgs...)

}

var xBuildPathv func(string, uintptr) string

// Behaves exactly like g_build_path(), but takes the path elements
// as a string array, instead of varargs. This function is mainly
// meant for language bindings.
func BuildPathv(SeparatorVar string, ArgsVar uintptr) string {

	return xBuildPathv(SeparatorVar, ArgsVar)

}

var xCanonicalizeFilename func(string, string) string

// Gets the canonical file name from @filename. All triple slashes are turned into
// single slashes, and all `..` and `.`s resolved against @relative_to.
//
// Symlinks are not followed, and the returned path is guaranteed to be absolute.
//
// If @filename is an absolute path, @relative_to is ignored. Otherwise,
// @relative_to will be prepended to @filename to make it absolute. @relative_to
// must be an absolute path, or %NULL. If @relative_to is %NULL, it'll fallback
// to g_get_current_dir().
//
// This function never fails, and will canonicalize file paths even if they don't
// exist.
//
// No file system I/O is done.
func CanonicalizeFilename(FilenameVar string, RelativeToVar string) string {

	return xCanonicalizeFilename(FilenameVar, RelativeToVar)

}

var xDirMakeTmp func(string) string

// Creates a subdirectory in the preferred directory for temporary
// files (as returned by g_get_tmp_dir()).
//
// @tmpl should be a string in the GLib file name encoding containing
// a sequence of six 'X' characters, as the parameter to g_mkstemp().
// However, unlike these functions, the template should only be a
// basename, no directory components are allowed. If template is
// %NULL, a default template is used.
//
// Note that in contrast to g_mkdtemp() (and mkdtemp()) @tmpl is not
// modified, and might thus be a read-only literal string.
func DirMakeTmp(TmplVar string) string {

	return xDirMakeTmp(TmplVar)

}

var xFileErrorFromErrno func(int32) FileError

// Gets a #GFileError constant based on the passed-in @err_no.
//
// For example, if you pass in `EEXIST` this function returns
// %G_FILE_ERROR_EXIST. Unlike `errno` values, you can portably
// assume that all #GFileError values will exist.
//
// Normally a #GFileError value goes into a #GError returned
// from a function that manipulates files. So you would use
// g_file_error_from_errno() when constructing a #GError.
func FileErrorFromErrno(ErrNoVar int32) FileError {

	return xFileErrorFromErrno(ErrNoVar)

}

var xFileGetContents func(string, uintptr, uint) bool

// Reads an entire file into allocated memory, with good error
// checking.
//
// If the call was successful, it returns %TRUE and sets @contents to the file
// contents and @length to the length of the file contents in bytes. The string
// stored in @contents will be nul-terminated, so for text files you can pass
// %NULL for the @length argument. If the call was not successful, it returns
// %FALSE and sets @error. The error domain is %G_FILE_ERROR. Possible error
// codes are those in the #GFileError enumeration. In the error case,
// @contents is set to %NULL and @length is set to zero.
func FileGetContents(FilenameVar string, ContentsVar uintptr, LengthVar uint) bool {

	return xFileGetContents(FilenameVar, ContentsVar, LengthVar)

}

var xFileOpenTmp func(string, string) int32

// Opens a file for writing in the preferred directory for temporary
// files (as returned by g_get_tmp_dir()).
//
// @tmpl should be a string in the GLib file name encoding containing
// a sequence of six 'X' characters, as the parameter to g_mkstemp().
// However, unlike these functions, the template should only be a
// basename, no directory components are allowed. If template is
// %NULL, a default template is used.
//
// Note that in contrast to g_mkstemp() (and mkstemp()) @tmpl is not
// modified, and might thus be a read-only literal string.
//
// Upon success, and if @name_used is non-%NULL, the actual name used
// is returned in @name_used. This string should be freed with g_free()
// when not needed any longer. The returned name is in the GLib file
// name encoding.
func FileOpenTmp(TmplVar string, NameUsedVar string) int32 {

	return xFileOpenTmp(TmplVar, NameUsedVar)

}

var xFileReadLink func(string) string

// Reads the contents of the symbolic link @filename like the POSIX
// readlink() function.  The returned string is in the encoding used
// for filenames. Use g_filename_to_utf8() to convert it to UTF-8.
func FileReadLink(FilenameVar string) string {

	return xFileReadLink(FilenameVar)

}

var xFileSetContents func(string, uintptr, int) bool

// Writes all of @contents to a file named @filename. This is a convenience
// wrapper around calling g_file_set_contents_full() with `flags` set to
// `G_FILE_SET_CONTENTS_CONSISTENT | G_FILE_SET_CONTENTS_ONLY_EXISTING` and
// `mode` set to `0666`.
func FileSetContents(FilenameVar string, ContentsVar uintptr, LengthVar int) bool {

	return xFileSetContents(FilenameVar, ContentsVar, LengthVar)

}

var xFileSetContentsFull func(string, uintptr, int, FileSetContentsFlags, int32) bool

// Writes all of @contents to a file named @filename, with good error checking.
// If a file called @filename already exists it will be overwritten.
//
// @flags control the properties of the write operation: whether it’s atomic,
// and what the tradeoff is between returning quickly or being resilient to
// system crashes.
//
// As this function performs file I/O, it is recommended to not call it anywhere
// where blocking would cause problems, such as in the main loop of a graphical
// application. In particular, if @flags has any value other than
// %G_FILE_SET_CONTENTS_NONE then this function may call `fsync()`.
//
// If %G_FILE_SET_CONTENTS_CONSISTENT is set in @flags, the operation is atomic
// in the sense that it is first written to a temporary file which is then
// renamed to the final name.
//
// Notes:
//
//   - On UNIX, if @filename already exists hard links to @filename will break.
//     Also since the file is recreated, existing permissions, access control
//     lists, metadata etc. may be lost. If @filename is a symbolic link,
//     the link itself will be replaced, not the linked file.
//
//   - On UNIX, if @filename already exists and is non-empty, and if the system
//     supports it (via a journalling filesystem or equivalent), and if
//     %G_FILE_SET_CONTENTS_CONSISTENT is set in @flags, the `fsync()` call (or
//     equivalent) will be used to ensure atomic replacement: @filename
//     will contain either its old contents or @contents, even in the face of
//     system power loss, the disk being unsafely removed, etc.
//
//   - On UNIX, if @filename does not already exist or is empty, there is a
//     possibility that system power loss etc. after calling this function will
//     leave @filename empty or full of NUL bytes, depending on the underlying
//     filesystem, unless %G_FILE_SET_CONTENTS_DURABLE and
//     %G_FILE_SET_CONTENTS_CONSISTENT are set in @flags.
//
//   - On Windows renaming a file will not remove an existing file with the
//     new name, so on Windows there is a race condition between the existing
//     file being removed and the temporary file being renamed.
//
//   - On Windows there is no way to remove a file that is open to some
//     process, or mapped into memory. Thus, this function will fail if
//     @filename already exists and is open.
//
// If the call was successful, it returns %TRUE. If the call was not successful,
// it returns %FALSE and sets @error. The error domain is %G_FILE_ERROR.
// Possible error codes are those in the #GFileError enumeration.
//
// Note that the name for the temporary file is constructed by appending up
// to 7 characters to @filename.
//
// If the file didn’t exist before and is created, it will be given the
// permissions from @mode. Otherwise, the permissions of the existing file may
// be changed to @mode depending on @flags, or they may remain unchanged.
func FileSetContentsFull(FilenameVar string, ContentsVar uintptr, LengthVar int, FlagsVar FileSetContentsFlags, ModeVar int32) bool {

	return xFileSetContentsFull(FilenameVar, ContentsVar, LengthVar, FlagsVar, ModeVar)

}

var xNewFileTest func(string, FileTest) bool

// Returns %TRUE if any of the tests in the bitfield @test are
// %TRUE. For example, `(G_FILE_TEST_EXISTS | G_FILE_TEST_IS_DIR)`
// will return %TRUE if the file exists; the check whether it's a
// directory doesn't matter since the existence test is %TRUE. With
// the current set of available tests, there's no point passing in
// more than one test at a time.
//
// Apart from %G_FILE_TEST_IS_SYMLINK all tests follow symbolic links,
// so for a symbolic link to a regular file g_file_test() will return
// %TRUE for both %G_FILE_TEST_IS_SYMLINK and %G_FILE_TEST_IS_REGULAR.
//
// Note, that for a dangling symbolic link g_file_test() will return
// %TRUE for %G_FILE_TEST_IS_SYMLINK and %FALSE for all other flags.
//
// You should never use g_file_test() to test whether it is safe
// to perform an operation, because there is always the possibility
// of the condition changing before you actually perform the operation.
// For example, you might think you could use %G_FILE_TEST_IS_SYMLINK
// to know whether it is safe to write to a file without being
// tricked into writing into a different location. It doesn't work!
// |[&lt;!-- language="C" --&gt;
//
//	// DON'T DO THIS
//	if (!g_file_test (filename, G_FILE_TEST_IS_SYMLINK))
//	  {
//	    fd = g_open (filename, O_WRONLY);
//	    // write to fd
//	  }
//
// ]|
//
// Another thing to note is that %G_FILE_TEST_EXISTS and
// %G_FILE_TEST_IS_EXECUTABLE are implemented using the access()
// system call. This usually doesn't matter, but if your program
// is setuid or setgid it means that these tests will give you
// the answer for the real user ID and group ID, rather than the
// effective user ID and group ID.
//
// On Windows, there are no symlinks, so testing for
// %G_FILE_TEST_IS_SYMLINK will always return %FALSE. Testing for
// %G_FILE_TEST_IS_EXECUTABLE will just check that the file exists and
// its name indicates that it is executable, checking for well-known
// extensions and those listed in the `PATHEXT` environment variable.
func NewFileTest(FilenameVar string, TestVar FileTest) bool {

	return xNewFileTest(FilenameVar, TestVar)

}

var xGetCurrentDir func() string

// Gets the current directory.
//
// The returned string should be freed when no longer needed.
// The encoding of the returned string is system defined.
// On Windows, it is always UTF-8.
//
// Since GLib 2.40, this function will return the value of the "PWD"
// environment variable if it is set and it happens to be the same as
// the current directory.  This can make a difference in the case that
// the current directory is the target of a symbolic link.
func GetCurrentDir() string {

	return xGetCurrentDir()

}

var xMkdirWithParents func(string, int32) int32

// Create a directory if it doesn't already exist. Create intermediate
// parent directories as needed, too.
func MkdirWithParents(PathnameVar string, ModeVar int32) int32 {

	return xMkdirWithParents(PathnameVar, ModeVar)

}

var xMkdtemp func(string) string

// Creates a temporary directory. See the mkdtemp() documentation
// on most UNIX-like systems.
//
// The parameter is a string that should follow the rules for
// mkdtemp() templates, i.e. contain the string "XXXXXX".
// g_mkdtemp() is slightly more flexible than mkdtemp() in that the
// sequence does not have to occur at the very end of the template.
// The X string will be modified to form the name of a directory that
// didn't exist.
// The string should be in the GLib file name encoding. Most importantly,
// on Windows it should be in UTF-8.
//
// If you are going to be creating a temporary directory inside the
// directory returned by g_get_tmp_dir(), you might want to use
// g_dir_make_tmp() instead.
func Mkdtemp(TmplVar string) string {

	return xMkdtemp(TmplVar)

}

var xMkdtempFull func(string, int32) string

// Creates a temporary directory. See the mkdtemp() documentation
// on most UNIX-like systems.
//
// The parameter is a string that should follow the rules for
// mkdtemp() templates, i.e. contain the string "XXXXXX".
// g_mkdtemp_full() is slightly more flexible than mkdtemp() in that the
// sequence does not have to occur at the very end of the template
// and you can pass a @mode. The X string will be modified to form
// the name of a directory that didn't exist. The string should be
// in the GLib file name encoding. Most importantly, on Windows it
// should be in UTF-8.
//
// If you are going to be creating a temporary directory inside the
// directory returned by g_get_tmp_dir(), you might want to use
// g_dir_make_tmp() instead.
func MkdtempFull(TmplVar string, ModeVar int32) string {

	return xMkdtempFull(TmplVar, ModeVar)

}

var xMkstemp func(string) int32

// Opens a temporary file. See the mkstemp() documentation
// on most UNIX-like systems.
//
// The parameter is a string that should follow the rules for
// mkstemp() templates, i.e. contain the string "XXXXXX".
// g_mkstemp() is slightly more flexible than mkstemp() in that the
// sequence does not have to occur at the very end of the template.
// The X string will be modified to form the name of a file that
// didn't exist. The string should be in the GLib file name encoding.
// Most importantly, on Windows it should be in UTF-8.
func Mkstemp(TmplVar string) int32 {

	return xMkstemp(TmplVar)

}

var xMkstempFull func(string, int32, int32) int32

// Opens a temporary file. See the mkstemp() documentation
// on most UNIX-like systems.
//
// The parameter is a string that should follow the rules for
// mkstemp() templates, i.e. contain the string "XXXXXX".
// g_mkstemp_full() is slightly more flexible than mkstemp()
// in that the sequence does not have to occur at the very end of the
// template and you can pass a @mode and additional @flags. The X
// string will be modified to form the name of a file that didn't exist.
// The string should be in the GLib file name encoding. Most importantly,
// on Windows it should be in UTF-8.
func MkstempFull(TmplVar string, FlagsVar int32, ModeVar int32) int32 {

	return xMkstempFull(TmplVar, FlagsVar, ModeVar)

}

var xPathGetBasename func(string) string

// Gets the last component of the filename.
//
// If @file_name ends with a directory separator it gets the component
// before the last slash. If @file_name consists only of directory
// separators (and on Windows, possibly a drive letter), a single
// separator is returned. If @file_name is empty, it gets ".".
func PathGetBasename(FileNameVar string) string {

	return xPathGetBasename(FileNameVar)

}

var xPathGetDirname func(string) string

// Gets the directory components of a file name. For example, the directory
// component of `/usr/bin/test` is `/usr/bin`. The directory component of `/`
// is `/`.
//
// If the file name has no directory components "." is returned.
// The returned string should be freed when no longer needed.
func PathGetDirname(FileNameVar string) string {

	return xPathGetDirname(FileNameVar)

}

var xPathIsAbsolute func(string) bool

// Returns %TRUE if the given @file_name is an absolute file name.
// Note that this is a somewhat vague concept on Windows.
//
// On POSIX systems, an absolute file name is well-defined. It always
// starts from the single root directory. For example "/usr/local".
//
// On Windows, the concepts of current drive and drive-specific
// current directory introduce vagueness. This function interprets as
// an absolute file name one that either begins with a directory
// separator such as "\Users\tml" or begins with the root on a drive,
// for example "C:\Windows". The first case also includes UNC paths
// such as "\\\\myserver\docs\foo". In all cases, either slashes or
// backslashes are accepted.
//
// Note that a file name relative to the current drive root does not
// truly specify a file uniquely over time and across processes, as
// the current drive is a per-process value and can be changed.
//
// File names relative the current directory on some specific drive,
// such as "D:foo/bar", are not interpreted as absolute by this
// function, but they obviously are not relative to the normal current
// directory as returned by getcwd() or g_get_current_dir()
// either. Such paths should be avoided, or need to be handled using
// Windows-specific code.
func PathIsAbsolute(FileNameVar string) bool {

	return xPathIsAbsolute(FileNameVar)

}

var xPathSkipRoot func(string) string

// Returns a pointer into @file_name after the root component,
// i.e. after the "/" in UNIX or "C:\" under Windows. If @file_name
// is not an absolute path it returns %NULL.
func PathSkipRoot(FileNameVar string) string {

	return xPathSkipRoot(FileNameVar)

}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GLIB"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	core.PuregoSafeRegister(&xBasename, lib, "g_basename")
	core.PuregoSafeRegister(&xBuildFilename, lib, "g_build_filename")
	core.PuregoSafeRegister(&xBuildFilenameValist, lib, "g_build_filename_valist")
	core.PuregoSafeRegister(&xBuildFilenamev, lib, "g_build_filenamev")
	core.PuregoSafeRegister(&xBuildPath, lib, "g_build_path")
	core.PuregoSafeRegister(&xBuildPathv, lib, "g_build_pathv")
	core.PuregoSafeRegister(&xCanonicalizeFilename, lib, "g_canonicalize_filename")
	core.PuregoSafeRegister(&xDirMakeTmp, lib, "g_dir_make_tmp")
	core.PuregoSafeRegister(&xFileErrorFromErrno, lib, "g_file_error_from_errno")
	core.PuregoSafeRegister(&xFileGetContents, lib, "g_file_get_contents")
	core.PuregoSafeRegister(&xFileOpenTmp, lib, "g_file_open_tmp")
	core.PuregoSafeRegister(&xFileReadLink, lib, "g_file_read_link")
	core.PuregoSafeRegister(&xFileSetContents, lib, "g_file_set_contents")
	core.PuregoSafeRegister(&xFileSetContentsFull, lib, "g_file_set_contents_full")
	core.PuregoSafeRegister(&xNewFileTest, lib, "g_file_test")
	core.PuregoSafeRegister(&xGetCurrentDir, lib, "g_get_current_dir")
	core.PuregoSafeRegister(&xMkdirWithParents, lib, "g_mkdir_with_parents")
	core.PuregoSafeRegister(&xMkdtemp, lib, "g_mkdtemp")
	core.PuregoSafeRegister(&xMkdtempFull, lib, "g_mkdtemp_full")
	core.PuregoSafeRegister(&xMkstemp, lib, "g_mkstemp")
	core.PuregoSafeRegister(&xMkstempFull, lib, "g_mkstemp_full")
	core.PuregoSafeRegister(&xPathGetBasename, lib, "g_path_get_basename")
	core.PuregoSafeRegister(&xPathGetDirname, lib, "g_path_get_dirname")
	core.PuregoSafeRegister(&xPathIsAbsolute, lib, "g_path_is_absolute")
	core.PuregoSafeRegister(&xPathSkipRoot, lib, "g_path_skip_root")

}
