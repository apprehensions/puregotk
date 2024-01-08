// Package gdk was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gdk

import (
	"unsafe"

	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/glib"
)

// A `GdkContentFormatsBuilder` is an auxiliary struct used to create
// new `GdkContentFormats`, and should not be kept around.
type ContentFormatsBuilder struct {
}

func (x *ContentFormatsBuilder) GoPointer() uintptr {
	return uintptr(unsafe.Pointer(x))
}

var xNewContentFormatsBuilder func() *ContentFormatsBuilder

// Create a new `GdkContentFormatsBuilder` object.
//
// The resulting builder would create an empty `GdkContentFormats`.
// Use addition functions to add types to it.
func NewContentFormatsBuilder() *ContentFormatsBuilder {

	cret := xNewContentFormatsBuilder()
	return cret
}

var xContentFormatsBuilderAddFormats func(uintptr, *ContentFormats)

// Appends all formats from @formats to @builder, skipping those that
// already exist.
func (x *ContentFormatsBuilder) AddFormats(FormatsVar *ContentFormats) {

	xContentFormatsBuilderAddFormats(x.GoPointer(), FormatsVar)

}

var xContentFormatsBuilderAddGtype func(uintptr, []interface{})

// Appends @type to @builder if it has not already been added.
func (x *ContentFormatsBuilder) AddGtype(TypeVar []interface{}) {

	xContentFormatsBuilderAddGtype(x.GoPointer(), TypeVar)

}

var xContentFormatsBuilderAddMimeType func(uintptr, string)

// Appends @mime_type to @builder if it has not already been added.
func (x *ContentFormatsBuilder) AddMimeType(MimeTypeVar string) {

	xContentFormatsBuilderAddMimeType(x.GoPointer(), MimeTypeVar)

}

var xContentFormatsBuilderFreeToFormats func(uintptr) *ContentFormats

// Creates a new `GdkContentFormats` from the current state of the
// given @builder, and frees the @builder instance.
func (x *ContentFormatsBuilder) FreeToFormats() *ContentFormats {

	cret := xContentFormatsBuilderFreeToFormats(x.GoPointer())
	return cret
}

var xContentFormatsBuilderRef func(uintptr) *ContentFormatsBuilder

// Acquires a reference on the given @builder.
//
// This function is intended primarily for bindings.
// `GdkContentFormatsBuilder` objects should not be kept around.
func (x *ContentFormatsBuilder) Ref() *ContentFormatsBuilder {

	cret := xContentFormatsBuilderRef(x.GoPointer())
	return cret
}

var xContentFormatsBuilderToFormats func(uintptr) *ContentFormats

// Creates a new `GdkContentFormats` from the given @builder.
//
// The given `GdkContentFormatsBuilder` is reset once this function returns;
// you cannot call this function multiple times on the same @builder instance.
//
// This function is intended primarily for bindings. C code should use
// [method@Gdk.ContentFormatsBuilder.free_to_formats].
func (x *ContentFormatsBuilder) ToFormats() *ContentFormats {

	cret := xContentFormatsBuilderToFormats(x.GoPointer())
	return cret
}

var xContentFormatsBuilderUnref func(uintptr)

// Releases a reference on the given @builder.
func (x *ContentFormatsBuilder) Unref() {

	xContentFormatsBuilderUnref(x.GoPointer())

}

// An opaque type representing a list of files.
type FileList struct {
}

func (x *FileList) GoPointer() uintptr {
	return uintptr(unsafe.Pointer(x))
}

var xNewFromArrayFileList func(uintptr, uint) *FileList

// Creates a new `GdkFileList` for the given array of files.
//
// This function is meant to be used by language bindings.
func NewFromArrayFileList(FilesVar uintptr, NFilesVar uint) *FileList {

	cret := xNewFromArrayFileList(FilesVar, NFilesVar)
	return cret
}

var xNewFromListFileList func(*glib.SList) *FileList

// Creates a new files list container from a singly linked list of
// `GFile` instances.
//
// This function is meant to be used by language bindings
func NewFromListFileList(FilesVar *glib.SList) *FileList {

	cret := xNewFromListFileList(FilesVar)
	return cret
}

var xFileListGetFiles func(uintptr) *glib.SList

// Retrieves the list of files inside a `GdkFileList`.
//
// This function is meant for language bindings.
func (x *FileList) GetFiles() *glib.SList {

	cret := xFileListGetFiles(x.GoPointer())
	return cret
}

var xContentFormatsParse func(string) *ContentFormats

// Parses the given @string into `GdkContentFormats` and
// returns the formats.
//
// Strings printed via [method@Gdk.ContentFormats.to_string]
// can be read in again successfully using this function.
//
// If @string does not describe valid content formats, %NULL
// is returned.
func ContentFormatsParse(StringVar string) *ContentFormats {

	cret := xContentFormatsParse(StringVar)
	return cret
}

var xInternMimeType func(string) string

// Canonicalizes the given mime type and interns the result.
//
// If @string is not a valid mime type, %NULL is returned instead.
// See RFC 2048 for the syntax if mime types.
func InternMimeType(StringVar string) string {

	cret := xInternMimeType(StringVar)
	return cret
}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GDK"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	core.PuregoSafeRegister(&xContentFormatsParse, lib, "gdk_content_formats_parse")
	core.PuregoSafeRegister(&xInternMimeType, lib, "gdk_intern_mime_type")

	core.PuregoSafeRegister(&xNewContentFormatsBuilder, lib, "gdk_content_formats_builder_new")

	core.PuregoSafeRegister(&xContentFormatsBuilderAddFormats, lib, "gdk_content_formats_builder_add_formats")
	core.PuregoSafeRegister(&xContentFormatsBuilderAddGtype, lib, "gdk_content_formats_builder_add_gtype")
	core.PuregoSafeRegister(&xContentFormatsBuilderAddMimeType, lib, "gdk_content_formats_builder_add_mime_type")
	core.PuregoSafeRegister(&xContentFormatsBuilderFreeToFormats, lib, "gdk_content_formats_builder_free_to_formats")
	core.PuregoSafeRegister(&xContentFormatsBuilderRef, lib, "gdk_content_formats_builder_ref")
	core.PuregoSafeRegister(&xContentFormatsBuilderToFormats, lib, "gdk_content_formats_builder_to_formats")
	core.PuregoSafeRegister(&xContentFormatsBuilderUnref, lib, "gdk_content_formats_builder_unref")

	core.PuregoSafeRegister(&xNewFromArrayFileList, lib, "gdk_file_list_new_from_array")
	core.PuregoSafeRegister(&xNewFromListFileList, lib, "gdk_file_list_new_from_list")

	core.PuregoSafeRegister(&xFileListGetFiles, lib, "gdk_file_list_get_files")

}
