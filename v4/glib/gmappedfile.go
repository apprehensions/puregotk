// Package glib was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package glib

import (
	"unsafe"

	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
)

// The #GMappedFile represents a file mapping created with
// g_mapped_file_new(). It has only private members and should
// not be accessed directly.
type MappedFile struct {
}

func (x *MappedFile) GoPointer() uintptr {
	return uintptr(unsafe.Pointer(x))
}

var xNewMappedFile func(string, bool, **Error) *MappedFile

// Maps a file into memory. On UNIX, this is using the mmap() function.
//
// If @writable is %TRUE, the mapped buffer may be modified, otherwise
// it is an error to modify the mapped buffer. Modifications to the buffer
// are not visible to other processes mapping the same file, and are not
// written back to the file.
//
// Note that modifications of the underlying file might affect the contents
// of the #GMappedFile. Therefore, mapping should only be used if the file
// will not be modified, or if all modifications of the file are done
// atomically (e.g. using g_file_set_contents()).
//
// If @filename is the name of an empty, regular file, the function
// will successfully return an empty #GMappedFile. In other cases of
// size 0 (e.g. device files such as /dev/null), @error will be set
// to the #GFileError value %G_FILE_ERROR_INVAL.
func NewMappedFile(FilenameVar string, WritableVar bool) (*MappedFile, error) {
	var cerr *Error

	cret := xNewMappedFile(FilenameVar, WritableVar, &cerr)
	if cerr == nil {
		return cret, nil
	}
	return cret, cerr

}

var xNewFromFdMappedFile func(int, bool, **Error) *MappedFile

// Maps a file into memory. On UNIX, this is using the mmap() function.
//
// If @writable is %TRUE, the mapped buffer may be modified, otherwise
// it is an error to modify the mapped buffer. Modifications to the buffer
// are not visible to other processes mapping the same file, and are not
// written back to the file.
//
// Note that modifications of the underlying file might affect the contents
// of the #GMappedFile. Therefore, mapping should only be used if the file
// will not be modified, or if all modifications of the file are done
// atomically (e.g. using g_file_set_contents()).
func NewFromFdMappedFile(FdVar int, WritableVar bool) (*MappedFile, error) {
	var cerr *Error

	cret := xNewFromFdMappedFile(FdVar, WritableVar, &cerr)
	if cerr == nil {
		return cret, nil
	}
	return cret, cerr

}

var xMappedFileFree func(uintptr)

// This call existed before #GMappedFile had refcounting and is currently
// exactly the same as g_mapped_file_unref().
func (x *MappedFile) Free() {

	xMappedFileFree(x.GoPointer())

}

var xMappedFileGetBytes func(uintptr) *Bytes

// Creates a new #GBytes which references the data mapped from @file.
// The mapped contents of the file must not be modified after creating this
// bytes object, because a #GBytes should be immutable.
func (x *MappedFile) GetBytes() *Bytes {

	cret := xMappedFileGetBytes(x.GoPointer())
	return cret
}

var xMappedFileGetContents func(uintptr) string

// Returns the contents of a #GMappedFile.
//
// Note that the contents may not be zero-terminated,
// even if the #GMappedFile is backed by a text file.
//
// If the file is empty then %NULL is returned.
func (x *MappedFile) GetContents() string {

	cret := xMappedFileGetContents(x.GoPointer())
	return cret
}

var xMappedFileGetLength func(uintptr) uint

// Returns the length of the contents of a #GMappedFile.
func (x *MappedFile) GetLength() uint {

	cret := xMappedFileGetLength(x.GoPointer())
	return cret
}

var xMappedFileRef func(uintptr) *MappedFile

// Increments the reference count of @file by one.  It is safe to call
// this function from any thread.
func (x *MappedFile) Ref() *MappedFile {

	cret := xMappedFileRef(x.GoPointer())
	return cret
}

var xMappedFileUnref func(uintptr)

// Decrements the reference count of @file by one.  If the reference count
// drops to 0, unmaps the buffer of @file and frees it.
//
// It is safe to call this function from any thread.
//
// Since 2.22
func (x *MappedFile) Unref() {

	xMappedFileUnref(x.GoPointer())

}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GLIB"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xNewMappedFile, lib, "g_mapped_file_new")
	core.PuregoSafeRegister(&xNewFromFdMappedFile, lib, "g_mapped_file_new_from_fd")

	core.PuregoSafeRegister(&xMappedFileFree, lib, "g_mapped_file_free")
	core.PuregoSafeRegister(&xMappedFileGetBytes, lib, "g_mapped_file_get_bytes")
	core.PuregoSafeRegister(&xMappedFileGetContents, lib, "g_mapped_file_get_contents")
	core.PuregoSafeRegister(&xMappedFileGetLength, lib, "g_mapped_file_get_length")
	core.PuregoSafeRegister(&xMappedFileRef, lib, "g_mapped_file_ref")
	core.PuregoSafeRegister(&xMappedFileUnref, lib, "g_mapped_file_unref")

}
