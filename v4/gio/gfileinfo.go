// Package gio was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gio

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/glib"
	"github.com/jwijenbergh/puregotk/v4/gobject"
)

type FileInfoClass struct {
}

// Functionality for manipulating basic metadata for files. #GFileInfo
// implements methods for getting information that all files should
// contain, and allows for manipulation of extended attributes.
//
// See [GFileAttribute][gio-GFileAttribute] for more information on how
// GIO handles file attributes.
//
// To obtain a #GFileInfo for a #GFile, use g_file_query_info() (or its
// async variant). To obtain a #GFileInfo for a file input or output
// stream, use g_file_input_stream_query_info() or
// g_file_output_stream_query_info() (or their async variants).
//
// To change the actual attributes of a file, you should then set the
// attribute in the #GFileInfo and call g_file_set_attributes_from_info()
// or g_file_set_attributes_async() on a GFile.
//
// However, not all attributes can be changed in the file. For instance,
// the actual size of a file cannot be changed via g_file_info_set_size().
// You may call g_file_query_settable_attributes() and
// g_file_query_writable_namespaces() to discover the settable attributes
// of a particular file at runtime.
//
// The direct accessors, such as g_file_info_get_name(), are slightly more
// optimized than the generic attribute accessors, such as
// g_file_info_get_attribute_byte_string().This optimization will matter
// only if calling the API in a tight loop.
//
// #GFileAttributeMatcher allows for searching through a #GFileInfo for
// attributes.
type FileInfo struct {
	gobject.Object
}

func FileInfoNewFromInternalPtr(ptr uintptr) *FileInfo {
	cls := &FileInfo{}
	cls.Ptr = ptr
	return cls
}

var xNewFileInfo func() uintptr

// Creates a new file info structure.
func NewFileInfo() *FileInfo {
	NewFileInfoPtr := xNewFileInfo()
	if NewFileInfoPtr == 0 {
		return nil
	}

	NewFileInfoCls := &FileInfo{}
	NewFileInfoCls.Ptr = NewFileInfoPtr
	return NewFileInfoCls
}

var xFileInfoClearStatus func(uintptr)

// Clears the status information from @info.
func (x *FileInfo) ClearStatus() {

	xFileInfoClearStatus(x.GoPointer())

}

var xFileInfoCopyInto func(uintptr, uintptr)

// First clears all of the [GFileAttribute][gio-GFileAttribute] of @dest_info,
// and then copies all of the file attributes from @src_info to @dest_info.
func (x *FileInfo) CopyInto(DestInfoVar *FileInfo) {

	xFileInfoCopyInto(x.GoPointer(), DestInfoVar.GoPointer())

}

var xFileInfoDup func(uintptr) uintptr

// Duplicates a file info structure.
func (x *FileInfo) Dup() *FileInfo {

	DupPtr := xFileInfoDup(x.GoPointer())
	if DupPtr == 0 {
		return nil
	}

	DupCls := &FileInfo{}
	DupCls.Ptr = DupPtr
	return DupCls

}

var xFileInfoGetAccessDateTime func(uintptr) *glib.DateTime

// Gets the access time of the current @info and returns it as a
// #GDateTime.
//
// This requires the %G_FILE_ATTRIBUTE_TIME_ACCESS attribute. If
// %G_FILE_ATTRIBUTE_TIME_ACCESS_USEC is provided, the resulting #GDateTime
// will have microsecond precision.
func (x *FileInfo) GetAccessDateTime() *glib.DateTime {

	return xFileInfoGetAccessDateTime(x.GoPointer())

}

var xFileInfoGetAttributeAsString func(uintptr, string) string

// Gets the value of a attribute, formatted as a string.
// This escapes things as needed to make the string valid
// UTF-8.
func (x *FileInfo) GetAttributeAsString(AttributeVar string) string {

	return xFileInfoGetAttributeAsString(x.GoPointer(), AttributeVar)

}

var xFileInfoGetAttributeBoolean func(uintptr, string) bool

// Gets the value of a boolean attribute. If the attribute does not
// contain a boolean value, %FALSE will be returned.
func (x *FileInfo) GetAttributeBoolean(AttributeVar string) bool {

	return xFileInfoGetAttributeBoolean(x.GoPointer(), AttributeVar)

}

var xFileInfoGetAttributeByteString func(uintptr, string) string

// Gets the value of a byte string attribute. If the attribute does
// not contain a byte string, %NULL will be returned.
func (x *FileInfo) GetAttributeByteString(AttributeVar string) string {

	return xFileInfoGetAttributeByteString(x.GoPointer(), AttributeVar)

}

var xFileInfoGetAttributeData func(uintptr, string, *FileAttributeType, uintptr, *FileAttributeStatus) bool

// Gets the attribute type, value and status for an attribute key.
func (x *FileInfo) GetAttributeData(AttributeVar string, TypeVar *FileAttributeType, ValuePpVar uintptr, StatusVar *FileAttributeStatus) bool {

	return xFileInfoGetAttributeData(x.GoPointer(), AttributeVar, TypeVar, ValuePpVar, StatusVar)

}

var xFileInfoGetAttributeInt32 func(uintptr, string) int32

// Gets a signed 32-bit integer contained within the attribute. If the
// attribute does not contain a signed 32-bit integer, or is invalid,
// 0 will be returned.
func (x *FileInfo) GetAttributeInt32(AttributeVar string) int32 {

	return xFileInfoGetAttributeInt32(x.GoPointer(), AttributeVar)

}

var xFileInfoGetAttributeInt64 func(uintptr, string) int64

// Gets a signed 64-bit integer contained within the attribute. If the
// attribute does not contain a signed 64-bit integer, or is invalid,
// 0 will be returned.
func (x *FileInfo) GetAttributeInt64(AttributeVar string) int64 {

	return xFileInfoGetAttributeInt64(x.GoPointer(), AttributeVar)

}

var xFileInfoGetAttributeObject func(uintptr, string) uintptr

// Gets the value of a #GObject attribute. If the attribute does
// not contain a #GObject, %NULL will be returned.
func (x *FileInfo) GetAttributeObject(AttributeVar string) *gobject.Object {

	GetAttributeObjectPtr := xFileInfoGetAttributeObject(x.GoPointer(), AttributeVar)
	if GetAttributeObjectPtr == 0 {
		return nil
	}

	gobject.IncreaseRef(GetAttributeObjectPtr)

	GetAttributeObjectCls := &gobject.Object{}
	GetAttributeObjectCls.Ptr = GetAttributeObjectPtr
	return GetAttributeObjectCls

}

var xFileInfoGetAttributeStatus func(uintptr, string) FileAttributeStatus

// Gets the attribute status for an attribute key.
func (x *FileInfo) GetAttributeStatus(AttributeVar string) FileAttributeStatus {

	return xFileInfoGetAttributeStatus(x.GoPointer(), AttributeVar)

}

var xFileInfoGetAttributeString func(uintptr, string) string

// Gets the value of a string attribute. If the attribute does
// not contain a string, %NULL will be returned.
func (x *FileInfo) GetAttributeString(AttributeVar string) string {

	return xFileInfoGetAttributeString(x.GoPointer(), AttributeVar)

}

var xFileInfoGetAttributeStringv func(uintptr, string) []string

// Gets the value of a stringv attribute. If the attribute does
// not contain a stringv, %NULL will be returned.
func (x *FileInfo) GetAttributeStringv(AttributeVar string) []string {

	return xFileInfoGetAttributeStringv(x.GoPointer(), AttributeVar)

}

var xFileInfoGetAttributeType func(uintptr, string) FileAttributeType

// Gets the attribute type for an attribute key.
func (x *FileInfo) GetAttributeType(AttributeVar string) FileAttributeType {

	return xFileInfoGetAttributeType(x.GoPointer(), AttributeVar)

}

var xFileInfoGetAttributeUint32 func(uintptr, string) uint32

// Gets an unsigned 32-bit integer contained within the attribute. If the
// attribute does not contain an unsigned 32-bit integer, or is invalid,
// 0 will be returned.
func (x *FileInfo) GetAttributeUint32(AttributeVar string) uint32 {

	return xFileInfoGetAttributeUint32(x.GoPointer(), AttributeVar)

}

var xFileInfoGetAttributeUint64 func(uintptr, string) uint64

// Gets a unsigned 64-bit integer contained within the attribute. If the
// attribute does not contain an unsigned 64-bit integer, or is invalid,
// 0 will be returned.
func (x *FileInfo) GetAttributeUint64(AttributeVar string) uint64 {

	return xFileInfoGetAttributeUint64(x.GoPointer(), AttributeVar)

}

var xFileInfoGetContentType func(uintptr) string

// Gets the file's content type.
func (x *FileInfo) GetContentType() string {

	return xFileInfoGetContentType(x.GoPointer())

}

var xFileInfoGetCreationDateTime func(uintptr) *glib.DateTime

// Gets the creation time of the current @info and returns it as a
// #GDateTime.
//
// This requires the %G_FILE_ATTRIBUTE_TIME_CREATED attribute. If
// %G_FILE_ATTRIBUTE_TIME_CREATED_USEC is provided, the resulting #GDateTime
// will have microsecond precision.
func (x *FileInfo) GetCreationDateTime() *glib.DateTime {

	return xFileInfoGetCreationDateTime(x.GoPointer())

}

var xFileInfoGetDeletionDate func(uintptr) *glib.DateTime

// Returns the #GDateTime representing the deletion date of the file, as
// available in G_FILE_ATTRIBUTE_TRASH_DELETION_DATE. If the
// G_FILE_ATTRIBUTE_TRASH_DELETION_DATE attribute is unset, %NULL is returned.
func (x *FileInfo) GetDeletionDate() *glib.DateTime {

	return xFileInfoGetDeletionDate(x.GoPointer())

}

var xFileInfoGetDisplayName func(uintptr) string

// Gets a display name for a file. This is guaranteed to always be set.
func (x *FileInfo) GetDisplayName() string {

	return xFileInfoGetDisplayName(x.GoPointer())

}

var xFileInfoGetEditName func(uintptr) string

// Gets the edit name for a file.
func (x *FileInfo) GetEditName() string {

	return xFileInfoGetEditName(x.GoPointer())

}

var xFileInfoGetEtag func(uintptr) string

// Gets the [entity tag][gfile-etag] for a given
// #GFileInfo. See %G_FILE_ATTRIBUTE_ETAG_VALUE.
func (x *FileInfo) GetEtag() string {

	return xFileInfoGetEtag(x.GoPointer())

}

var xFileInfoGetFileType func(uintptr) FileType

// Gets a file's type (whether it is a regular file, symlink, etc).
// This is different from the file's content type, see g_file_info_get_content_type().
func (x *FileInfo) GetFileType() FileType {

	return xFileInfoGetFileType(x.GoPointer())

}

var xFileInfoGetIcon func(uintptr) uintptr

// Gets the icon for a file.
func (x *FileInfo) GetIcon() *IconBase {

	GetIconPtr := xFileInfoGetIcon(x.GoPointer())
	if GetIconPtr == 0 {
		return nil
	}

	gobject.IncreaseRef(GetIconPtr)

	GetIconCls := &IconBase{}
	GetIconCls.Ptr = GetIconPtr
	return GetIconCls

}

var xFileInfoGetIsBackup func(uintptr) bool

// Checks if a file is a backup file.
func (x *FileInfo) GetIsBackup() bool {

	return xFileInfoGetIsBackup(x.GoPointer())

}

var xFileInfoGetIsHidden func(uintptr) bool

// Checks if a file is hidden.
func (x *FileInfo) GetIsHidden() bool {

	return xFileInfoGetIsHidden(x.GoPointer())

}

var xFileInfoGetIsSymlink func(uintptr) bool

// Checks if a file is a symlink.
func (x *FileInfo) GetIsSymlink() bool {

	return xFileInfoGetIsSymlink(x.GoPointer())

}

var xFileInfoGetModificationDateTime func(uintptr) *glib.DateTime

// Gets the modification time of the current @info and returns it as a
// #GDateTime.
//
// This requires the %G_FILE_ATTRIBUTE_TIME_MODIFIED attribute. If
// %G_FILE_ATTRIBUTE_TIME_MODIFIED_USEC is provided, the resulting #GDateTime
// will have microsecond precision.
func (x *FileInfo) GetModificationDateTime() *glib.DateTime {

	return xFileInfoGetModificationDateTime(x.GoPointer())

}

var xFileInfoGetModificationTime func(uintptr, *glib.TimeVal)

// Gets the modification time of the current @info and sets it
// in @result.
func (x *FileInfo) GetModificationTime(ResultVar *glib.TimeVal) {

	xFileInfoGetModificationTime(x.GoPointer(), ResultVar)

}

var xFileInfoGetName func(uintptr) string

// Gets the name for a file. This is guaranteed to always be set.
func (x *FileInfo) GetName() string {

	return xFileInfoGetName(x.GoPointer())

}

var xFileInfoGetSize func(uintptr) int64

// Gets the file's size (in bytes). The size is retrieved through the value of
// the %G_FILE_ATTRIBUTE_STANDARD_SIZE attribute and is converted
// from #guint64 to #goffset before returning the result.
func (x *FileInfo) GetSize() int64 {

	return xFileInfoGetSize(x.GoPointer())

}

var xFileInfoGetSortOrder func(uintptr) int32

// Gets the value of the sort_order attribute from the #GFileInfo.
// See %G_FILE_ATTRIBUTE_STANDARD_SORT_ORDER.
func (x *FileInfo) GetSortOrder() int32 {

	return xFileInfoGetSortOrder(x.GoPointer())

}

var xFileInfoGetSymbolicIcon func(uintptr) uintptr

// Gets the symbolic icon for a file.
func (x *FileInfo) GetSymbolicIcon() *IconBase {

	GetSymbolicIconPtr := xFileInfoGetSymbolicIcon(x.GoPointer())
	if GetSymbolicIconPtr == 0 {
		return nil
	}

	gobject.IncreaseRef(GetSymbolicIconPtr)

	GetSymbolicIconCls := &IconBase{}
	GetSymbolicIconCls.Ptr = GetSymbolicIconPtr
	return GetSymbolicIconCls

}

var xFileInfoGetSymlinkTarget func(uintptr) string

// Gets the symlink target for a given #GFileInfo.
func (x *FileInfo) GetSymlinkTarget() string {

	return xFileInfoGetSymlinkTarget(x.GoPointer())

}

var xFileInfoHasAttribute func(uintptr, string) bool

// Checks if a file info structure has an attribute named @attribute.
func (x *FileInfo) HasAttribute(AttributeVar string) bool {

	return xFileInfoHasAttribute(x.GoPointer(), AttributeVar)

}

var xFileInfoHasNamespace func(uintptr, string) bool

// Checks if a file info structure has an attribute in the
// specified @name_space.
func (x *FileInfo) HasNamespace(NameSpaceVar string) bool {

	return xFileInfoHasNamespace(x.GoPointer(), NameSpaceVar)

}

var xFileInfoListAttributes func(uintptr, string) []string

// Lists the file info structure's attributes.
func (x *FileInfo) ListAttributes(NameSpaceVar string) []string {

	return xFileInfoListAttributes(x.GoPointer(), NameSpaceVar)

}

var xFileInfoRemoveAttribute func(uintptr, string)

// Removes all cases of @attribute from @info if it exists.
func (x *FileInfo) RemoveAttribute(AttributeVar string) {

	xFileInfoRemoveAttribute(x.GoPointer(), AttributeVar)

}

var xFileInfoSetAccessDateTime func(uintptr, *glib.DateTime)

// Sets the %G_FILE_ATTRIBUTE_TIME_ACCESS and
// %G_FILE_ATTRIBUTE_TIME_ACCESS_USEC attributes in the file info to the
// given date/time value.
func (x *FileInfo) SetAccessDateTime(AtimeVar *glib.DateTime) {

	xFileInfoSetAccessDateTime(x.GoPointer(), AtimeVar)

}

var xFileInfoSetAttribute func(uintptr, string, FileAttributeType, uintptr)

// Sets the @attribute to contain the given value, if possible. To unset the
// attribute, use %G_FILE_ATTRIBUTE_TYPE_INVALID for @type.
func (x *FileInfo) SetAttribute(AttributeVar string, TypeVar FileAttributeType, ValuePVar uintptr) {

	xFileInfoSetAttribute(x.GoPointer(), AttributeVar, TypeVar, ValuePVar)

}

var xFileInfoSetAttributeBoolean func(uintptr, string, bool)

// Sets the @attribute to contain the given @attr_value,
// if possible.
func (x *FileInfo) SetAttributeBoolean(AttributeVar string, AttrValueVar bool) {

	xFileInfoSetAttributeBoolean(x.GoPointer(), AttributeVar, AttrValueVar)

}

var xFileInfoSetAttributeByteString func(uintptr, string, string)

// Sets the @attribute to contain the given @attr_value,
// if possible.
func (x *FileInfo) SetAttributeByteString(AttributeVar string, AttrValueVar string) {

	xFileInfoSetAttributeByteString(x.GoPointer(), AttributeVar, AttrValueVar)

}

var xFileInfoSetAttributeInt32 func(uintptr, string, int32)

// Sets the @attribute to contain the given @attr_value,
// if possible.
func (x *FileInfo) SetAttributeInt32(AttributeVar string, AttrValueVar int32) {

	xFileInfoSetAttributeInt32(x.GoPointer(), AttributeVar, AttrValueVar)

}

var xFileInfoSetAttributeInt64 func(uintptr, string, int64)

// Sets the @attribute to contain the given @attr_value,
// if possible.
func (x *FileInfo) SetAttributeInt64(AttributeVar string, AttrValueVar int64) {

	xFileInfoSetAttributeInt64(x.GoPointer(), AttributeVar, AttrValueVar)

}

var xFileInfoSetAttributeMask func(uintptr, *FileAttributeMatcher)

// Sets @mask on @info to match specific attribute types.
func (x *FileInfo) SetAttributeMask(MaskVar *FileAttributeMatcher) {

	xFileInfoSetAttributeMask(x.GoPointer(), MaskVar)

}

var xFileInfoSetAttributeObject func(uintptr, string, uintptr)

// Sets the @attribute to contain the given @attr_value,
// if possible.
func (x *FileInfo) SetAttributeObject(AttributeVar string, AttrValueVar *gobject.Object) {

	xFileInfoSetAttributeObject(x.GoPointer(), AttributeVar, AttrValueVar.GoPointer())

}

var xFileInfoSetAttributeStatus func(uintptr, string, FileAttributeStatus) bool

// Sets the attribute status for an attribute key. This is only
// needed by external code that implement g_file_set_attributes_from_info()
// or similar functions.
//
// The attribute must exist in @info for this to work. Otherwise %FALSE
// is returned and @info is unchanged.
func (x *FileInfo) SetAttributeStatus(AttributeVar string, StatusVar FileAttributeStatus) bool {

	return xFileInfoSetAttributeStatus(x.GoPointer(), AttributeVar, StatusVar)

}

var xFileInfoSetAttributeString func(uintptr, string, string)

// Sets the @attribute to contain the given @attr_value,
// if possible.
func (x *FileInfo) SetAttributeString(AttributeVar string, AttrValueVar string) {

	xFileInfoSetAttributeString(x.GoPointer(), AttributeVar, AttrValueVar)

}

var xFileInfoSetAttributeStringv func(uintptr, string, []string)

// Sets the @attribute to contain the given @attr_value,
// if possible.
//
// Sinze: 2.22
func (x *FileInfo) SetAttributeStringv(AttributeVar string, AttrValueVar []string) {

	xFileInfoSetAttributeStringv(x.GoPointer(), AttributeVar, AttrValueVar)

}

var xFileInfoSetAttributeUint32 func(uintptr, string, uint32)

// Sets the @attribute to contain the given @attr_value,
// if possible.
func (x *FileInfo) SetAttributeUint32(AttributeVar string, AttrValueVar uint32) {

	xFileInfoSetAttributeUint32(x.GoPointer(), AttributeVar, AttrValueVar)

}

var xFileInfoSetAttributeUint64 func(uintptr, string, uint64)

// Sets the @attribute to contain the given @attr_value,
// if possible.
func (x *FileInfo) SetAttributeUint64(AttributeVar string, AttrValueVar uint64) {

	xFileInfoSetAttributeUint64(x.GoPointer(), AttributeVar, AttrValueVar)

}

var xFileInfoSetContentType func(uintptr, string)

// Sets the content type attribute for a given #GFileInfo.
// See %G_FILE_ATTRIBUTE_STANDARD_CONTENT_TYPE.
func (x *FileInfo) SetContentType(ContentTypeVar string) {

	xFileInfoSetContentType(x.GoPointer(), ContentTypeVar)

}

var xFileInfoSetCreationDateTime func(uintptr, *glib.DateTime)

// Sets the %G_FILE_ATTRIBUTE_TIME_CREATED and
// %G_FILE_ATTRIBUTE_TIME_CREATED_USEC attributes in the file info to the
// given date/time value.
func (x *FileInfo) SetCreationDateTime(CreationTimeVar *glib.DateTime) {

	xFileInfoSetCreationDateTime(x.GoPointer(), CreationTimeVar)

}

var xFileInfoSetDisplayName func(uintptr, string)

// Sets the display name for the current #GFileInfo.
// See %G_FILE_ATTRIBUTE_STANDARD_DISPLAY_NAME.
func (x *FileInfo) SetDisplayName(DisplayNameVar string) {

	xFileInfoSetDisplayName(x.GoPointer(), DisplayNameVar)

}

var xFileInfoSetEditName func(uintptr, string)

// Sets the edit name for the current file.
// See %G_FILE_ATTRIBUTE_STANDARD_EDIT_NAME.
func (x *FileInfo) SetEditName(EditNameVar string) {

	xFileInfoSetEditName(x.GoPointer(), EditNameVar)

}

var xFileInfoSetFileType func(uintptr, FileType)

// Sets the file type in a #GFileInfo to @type.
// See %G_FILE_ATTRIBUTE_STANDARD_TYPE.
func (x *FileInfo) SetFileType(TypeVar FileType) {

	xFileInfoSetFileType(x.GoPointer(), TypeVar)

}

var xFileInfoSetIcon func(uintptr, uintptr)

// Sets the icon for a given #GFileInfo.
// See %G_FILE_ATTRIBUTE_STANDARD_ICON.
func (x *FileInfo) SetIcon(IconVar Icon) {

	xFileInfoSetIcon(x.GoPointer(), IconVar.GoPointer())

}

var xFileInfoSetIsHidden func(uintptr, bool)

// Sets the "is_hidden" attribute in a #GFileInfo according to @is_hidden.
// See %G_FILE_ATTRIBUTE_STANDARD_IS_HIDDEN.
func (x *FileInfo) SetIsHidden(IsHiddenVar bool) {

	xFileInfoSetIsHidden(x.GoPointer(), IsHiddenVar)

}

var xFileInfoSetIsSymlink func(uintptr, bool)

// Sets the "is_symlink" attribute in a #GFileInfo according to @is_symlink.
// See %G_FILE_ATTRIBUTE_STANDARD_IS_SYMLINK.
func (x *FileInfo) SetIsSymlink(IsSymlinkVar bool) {

	xFileInfoSetIsSymlink(x.GoPointer(), IsSymlinkVar)

}

var xFileInfoSetModificationDateTime func(uintptr, *glib.DateTime)

// Sets the %G_FILE_ATTRIBUTE_TIME_MODIFIED and
// %G_FILE_ATTRIBUTE_TIME_MODIFIED_USEC attributes in the file info to the
// given date/time value.
func (x *FileInfo) SetModificationDateTime(MtimeVar *glib.DateTime) {

	xFileInfoSetModificationDateTime(x.GoPointer(), MtimeVar)

}

var xFileInfoSetModificationTime func(uintptr, *glib.TimeVal)

// Sets the %G_FILE_ATTRIBUTE_TIME_MODIFIED and
// %G_FILE_ATTRIBUTE_TIME_MODIFIED_USEC attributes in the file info to the
// given time value.
func (x *FileInfo) SetModificationTime(MtimeVar *glib.TimeVal) {

	xFileInfoSetModificationTime(x.GoPointer(), MtimeVar)

}

var xFileInfoSetName func(uintptr, string)

// Sets the name attribute for the current #GFileInfo.
// See %G_FILE_ATTRIBUTE_STANDARD_NAME.
func (x *FileInfo) SetName(NameVar string) {

	xFileInfoSetName(x.GoPointer(), NameVar)

}

var xFileInfoSetSize func(uintptr, int64)

// Sets the %G_FILE_ATTRIBUTE_STANDARD_SIZE attribute in the file info
// to the given size.
func (x *FileInfo) SetSize(SizeVar int64) {

	xFileInfoSetSize(x.GoPointer(), SizeVar)

}

var xFileInfoSetSortOrder func(uintptr, int32)

// Sets the sort order attribute in the file info structure. See
// %G_FILE_ATTRIBUTE_STANDARD_SORT_ORDER.
func (x *FileInfo) SetSortOrder(SortOrderVar int32) {

	xFileInfoSetSortOrder(x.GoPointer(), SortOrderVar)

}

var xFileInfoSetSymbolicIcon func(uintptr, uintptr)

// Sets the symbolic icon for a given #GFileInfo.
// See %G_FILE_ATTRIBUTE_STANDARD_SYMBOLIC_ICON.
func (x *FileInfo) SetSymbolicIcon(IconVar Icon) {

	xFileInfoSetSymbolicIcon(x.GoPointer(), IconVar.GoPointer())

}

var xFileInfoSetSymlinkTarget func(uintptr, string)

// Sets the %G_FILE_ATTRIBUTE_STANDARD_SYMLINK_TARGET attribute in the file info
// to the given symlink target.
func (x *FileInfo) SetSymlinkTarget(SymlinkTargetVar string) {

	xFileInfoSetSymlinkTarget(x.GoPointer(), SymlinkTargetVar)

}

var xFileInfoUnsetAttributeMask func(uintptr)

// Unsets a mask set by g_file_info_set_attribute_mask(), if one
// is set.
func (x *FileInfo) UnsetAttributeMask() {

	xFileInfoUnsetAttributeMask(x.GoPointer())

}

func (c *FileInfo) GoPointer() uintptr {
	return c.Ptr
}

func (c *FileInfo) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GIO"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xNewFileInfo, lib, "g_file_info_new")

	core.PuregoSafeRegister(&xFileInfoClearStatus, lib, "g_file_info_clear_status")
	core.PuregoSafeRegister(&xFileInfoCopyInto, lib, "g_file_info_copy_into")
	core.PuregoSafeRegister(&xFileInfoDup, lib, "g_file_info_dup")
	core.PuregoSafeRegister(&xFileInfoGetAccessDateTime, lib, "g_file_info_get_access_date_time")
	core.PuregoSafeRegister(&xFileInfoGetAttributeAsString, lib, "g_file_info_get_attribute_as_string")
	core.PuregoSafeRegister(&xFileInfoGetAttributeBoolean, lib, "g_file_info_get_attribute_boolean")
	core.PuregoSafeRegister(&xFileInfoGetAttributeByteString, lib, "g_file_info_get_attribute_byte_string")
	core.PuregoSafeRegister(&xFileInfoGetAttributeData, lib, "g_file_info_get_attribute_data")
	core.PuregoSafeRegister(&xFileInfoGetAttributeInt32, lib, "g_file_info_get_attribute_int32")
	core.PuregoSafeRegister(&xFileInfoGetAttributeInt64, lib, "g_file_info_get_attribute_int64")
	core.PuregoSafeRegister(&xFileInfoGetAttributeObject, lib, "g_file_info_get_attribute_object")
	core.PuregoSafeRegister(&xFileInfoGetAttributeStatus, lib, "g_file_info_get_attribute_status")
	core.PuregoSafeRegister(&xFileInfoGetAttributeString, lib, "g_file_info_get_attribute_string")
	core.PuregoSafeRegister(&xFileInfoGetAttributeStringv, lib, "g_file_info_get_attribute_stringv")
	core.PuregoSafeRegister(&xFileInfoGetAttributeType, lib, "g_file_info_get_attribute_type")
	core.PuregoSafeRegister(&xFileInfoGetAttributeUint32, lib, "g_file_info_get_attribute_uint32")
	core.PuregoSafeRegister(&xFileInfoGetAttributeUint64, lib, "g_file_info_get_attribute_uint64")
	core.PuregoSafeRegister(&xFileInfoGetContentType, lib, "g_file_info_get_content_type")
	core.PuregoSafeRegister(&xFileInfoGetCreationDateTime, lib, "g_file_info_get_creation_date_time")
	core.PuregoSafeRegister(&xFileInfoGetDeletionDate, lib, "g_file_info_get_deletion_date")
	core.PuregoSafeRegister(&xFileInfoGetDisplayName, lib, "g_file_info_get_display_name")
	core.PuregoSafeRegister(&xFileInfoGetEditName, lib, "g_file_info_get_edit_name")
	core.PuregoSafeRegister(&xFileInfoGetEtag, lib, "g_file_info_get_etag")
	core.PuregoSafeRegister(&xFileInfoGetFileType, lib, "g_file_info_get_file_type")
	core.PuregoSafeRegister(&xFileInfoGetIcon, lib, "g_file_info_get_icon")
	core.PuregoSafeRegister(&xFileInfoGetIsBackup, lib, "g_file_info_get_is_backup")
	core.PuregoSafeRegister(&xFileInfoGetIsHidden, lib, "g_file_info_get_is_hidden")
	core.PuregoSafeRegister(&xFileInfoGetIsSymlink, lib, "g_file_info_get_is_symlink")
	core.PuregoSafeRegister(&xFileInfoGetModificationDateTime, lib, "g_file_info_get_modification_date_time")
	core.PuregoSafeRegister(&xFileInfoGetModificationTime, lib, "g_file_info_get_modification_time")
	core.PuregoSafeRegister(&xFileInfoGetName, lib, "g_file_info_get_name")
	core.PuregoSafeRegister(&xFileInfoGetSize, lib, "g_file_info_get_size")
	core.PuregoSafeRegister(&xFileInfoGetSortOrder, lib, "g_file_info_get_sort_order")
	core.PuregoSafeRegister(&xFileInfoGetSymbolicIcon, lib, "g_file_info_get_symbolic_icon")
	core.PuregoSafeRegister(&xFileInfoGetSymlinkTarget, lib, "g_file_info_get_symlink_target")
	core.PuregoSafeRegister(&xFileInfoHasAttribute, lib, "g_file_info_has_attribute")
	core.PuregoSafeRegister(&xFileInfoHasNamespace, lib, "g_file_info_has_namespace")
	core.PuregoSafeRegister(&xFileInfoListAttributes, lib, "g_file_info_list_attributes")
	core.PuregoSafeRegister(&xFileInfoRemoveAttribute, lib, "g_file_info_remove_attribute")
	core.PuregoSafeRegister(&xFileInfoSetAccessDateTime, lib, "g_file_info_set_access_date_time")
	core.PuregoSafeRegister(&xFileInfoSetAttribute, lib, "g_file_info_set_attribute")
	core.PuregoSafeRegister(&xFileInfoSetAttributeBoolean, lib, "g_file_info_set_attribute_boolean")
	core.PuregoSafeRegister(&xFileInfoSetAttributeByteString, lib, "g_file_info_set_attribute_byte_string")
	core.PuregoSafeRegister(&xFileInfoSetAttributeInt32, lib, "g_file_info_set_attribute_int32")
	core.PuregoSafeRegister(&xFileInfoSetAttributeInt64, lib, "g_file_info_set_attribute_int64")
	core.PuregoSafeRegister(&xFileInfoSetAttributeMask, lib, "g_file_info_set_attribute_mask")
	core.PuregoSafeRegister(&xFileInfoSetAttributeObject, lib, "g_file_info_set_attribute_object")
	core.PuregoSafeRegister(&xFileInfoSetAttributeStatus, lib, "g_file_info_set_attribute_status")
	core.PuregoSafeRegister(&xFileInfoSetAttributeString, lib, "g_file_info_set_attribute_string")
	core.PuregoSafeRegister(&xFileInfoSetAttributeStringv, lib, "g_file_info_set_attribute_stringv")
	core.PuregoSafeRegister(&xFileInfoSetAttributeUint32, lib, "g_file_info_set_attribute_uint32")
	core.PuregoSafeRegister(&xFileInfoSetAttributeUint64, lib, "g_file_info_set_attribute_uint64")
	core.PuregoSafeRegister(&xFileInfoSetContentType, lib, "g_file_info_set_content_type")
	core.PuregoSafeRegister(&xFileInfoSetCreationDateTime, lib, "g_file_info_set_creation_date_time")
	core.PuregoSafeRegister(&xFileInfoSetDisplayName, lib, "g_file_info_set_display_name")
	core.PuregoSafeRegister(&xFileInfoSetEditName, lib, "g_file_info_set_edit_name")
	core.PuregoSafeRegister(&xFileInfoSetFileType, lib, "g_file_info_set_file_type")
	core.PuregoSafeRegister(&xFileInfoSetIcon, lib, "g_file_info_set_icon")
	core.PuregoSafeRegister(&xFileInfoSetIsHidden, lib, "g_file_info_set_is_hidden")
	core.PuregoSafeRegister(&xFileInfoSetIsSymlink, lib, "g_file_info_set_is_symlink")
	core.PuregoSafeRegister(&xFileInfoSetModificationDateTime, lib, "g_file_info_set_modification_date_time")
	core.PuregoSafeRegister(&xFileInfoSetModificationTime, lib, "g_file_info_set_modification_time")
	core.PuregoSafeRegister(&xFileInfoSetName, lib, "g_file_info_set_name")
	core.PuregoSafeRegister(&xFileInfoSetSize, lib, "g_file_info_set_size")
	core.PuregoSafeRegister(&xFileInfoSetSortOrder, lib, "g_file_info_set_sort_order")
	core.PuregoSafeRegister(&xFileInfoSetSymbolicIcon, lib, "g_file_info_set_symbolic_icon")
	core.PuregoSafeRegister(&xFileInfoSetSymlinkTarget, lib, "g_file_info_set_symlink_target")
	core.PuregoSafeRegister(&xFileInfoUnsetAttributeMask, lib, "g_file_info_unset_attribute_mask")

}
