// Package gio was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gio

import (
	"unsafe"

	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/glib"
	"github.com/jwijenbergh/puregotk/v4/gobject"
)

// This function type is used by g_vfs_register_uri_scheme() to make it
// possible for a client to associate an URI scheme to a different #GFile
// implementation.
//
// The client should return a reference to the new file that has been
// created for @uri, or %NULL to continue with the default implementation.
type VfsFileLookupFunc func(uintptr, string, uintptr) uintptr

type VfsClass struct {
	ParentClass uintptr
}

func (x *VfsClass) GoPointer() uintptr {
	return uintptr(unsafe.Pointer(x))
}

const (
	// Extension point for #GVfs functionality.
	// See [Extending GIO][extending-gio].
	VFS_EXTENSION_POINT_NAME string = "gio-vfs"
)

// Entry point for using GIO functionality.
type Vfs struct {
	gobject.Object
}

func VfsNewFromInternalPtr(ptr uintptr) *Vfs {
	cls := &Vfs{}
	cls.Ptr = ptr
	return cls
}

var xVfsGetFileForPath func(uintptr, string) uintptr

// Gets a #GFile for @path.
func (x *Vfs) GetFileForPath(PathVar string) *FileBase {
	var cls *FileBase

	cret := xVfsGetFileForPath(x.GoPointer(), PathVar)

	if cret == 0 {
		return nil
	}
	cls = &FileBase{}
	cls.Ptr = cret
	return cls
}

var xVfsGetFileForUri func(uintptr, string) uintptr

// Gets a #GFile for @uri.
//
// This operation never fails, but the returned object
// might not support any I/O operation if the URI
// is malformed or if the URI scheme is not supported.
func (x *Vfs) GetFileForUri(UriVar string) *FileBase {
	var cls *FileBase

	cret := xVfsGetFileForUri(x.GoPointer(), UriVar)

	if cret == 0 {
		return nil
	}
	cls = &FileBase{}
	cls.Ptr = cret
	return cls
}

var xVfsGetSupportedUriSchemes func(uintptr) uintptr

// Gets a list of URI schemes supported by @vfs.
func (x *Vfs) GetSupportedUriSchemes() uintptr {

	cret := xVfsGetSupportedUriSchemes(x.GoPointer())
	return cret
}

var xVfsIsActive func(uintptr) bool

// Checks if the VFS is active.
func (x *Vfs) IsActive() bool {

	cret := xVfsIsActive(x.GoPointer())
	return cret
}

var xVfsParseName func(uintptr, string) uintptr

// This operation never fails, but the returned object might
// not support any I/O operations if the @parse_name cannot
// be parsed by the #GVfs module.
func (x *Vfs) ParseName(ParseNameVar string) *FileBase {
	var cls *FileBase

	cret := xVfsParseName(x.GoPointer(), ParseNameVar)

	if cret == 0 {
		return nil
	}
	cls = &FileBase{}
	cls.Ptr = cret
	return cls
}

var xVfsRegisterUriScheme func(uintptr, string, uintptr, uintptr, uintptr, uintptr, uintptr, uintptr) bool

// Registers @uri_func and @parse_name_func as the #GFile URI and parse name
// lookup functions for URIs with a scheme matching @scheme.
// Note that @scheme is registered only within the running application, as
// opposed to desktop-wide as it happens with GVfs backends.
//
// When a #GFile is requested with an URI containing @scheme (e.g. through
// g_file_new_for_uri()), @uri_func will be called to allow a custom
// constructor. The implementation of @uri_func should not be blocking, and
// must not call g_vfs_register_uri_scheme() or g_vfs_unregister_uri_scheme().
//
// When g_file_parse_name() is called with a parse name obtained from such file,
// @parse_name_func will be called to allow the #GFile to be created again. In
// that case, it's responsibility of @parse_name_func to make sure the parse
// name matches what the custom #GFile implementation returned when
// g_file_get_parse_name() was previously called. The implementation of
// @parse_name_func should not be blocking, and must not call
// g_vfs_register_uri_scheme() or g_vfs_unregister_uri_scheme().
//
// It's an error to call this function twice with the same scheme. To unregister
// a custom URI scheme, use g_vfs_unregister_uri_scheme().
func (x *Vfs) RegisterUriScheme(SchemeVar string, UriFuncVar VfsFileLookupFunc, UriDataVar uintptr, UriDestroyVar glib.DestroyNotify, ParseNameFuncVar VfsFileLookupFunc, ParseNameDataVar uintptr, ParseNameDestroyVar glib.DestroyNotify) bool {

	cret := xVfsRegisterUriScheme(x.GoPointer(), SchemeVar, purego.NewCallback(UriFuncVar), UriDataVar, purego.NewCallback(UriDestroyVar), purego.NewCallback(ParseNameFuncVar), ParseNameDataVar, purego.NewCallback(ParseNameDestroyVar))
	return cret
}

var xVfsUnregisterUriScheme func(uintptr, string) bool

// Unregisters the URI handler for @scheme previously registered with
// g_vfs_register_uri_scheme().
func (x *Vfs) UnregisterUriScheme(SchemeVar string) bool {

	cret := xVfsUnregisterUriScheme(x.GoPointer(), SchemeVar)
	return cret
}

func (c *Vfs) GoPointer() uintptr {
	return c.Ptr
}

func (c *Vfs) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

var xVfsGetDefault func() uintptr

// Gets the default #GVfs for the system.
func VfsGetDefault() *Vfs {
	var cls *Vfs

	cret := xVfsGetDefault()

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &Vfs{}
	cls.Ptr = cret
	return cls
}

var xVfsGetLocal func() uintptr

// Gets the local #GVfs for the system.
func VfsGetLocal() *Vfs {
	var cls *Vfs

	cret := xVfsGetLocal()

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &Vfs{}
	cls.Ptr = cret
	return cls
}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GIO"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xVfsGetFileForPath, lib, "g_vfs_get_file_for_path")
	core.PuregoSafeRegister(&xVfsGetFileForUri, lib, "g_vfs_get_file_for_uri")
	core.PuregoSafeRegister(&xVfsGetSupportedUriSchemes, lib, "g_vfs_get_supported_uri_schemes")
	core.PuregoSafeRegister(&xVfsIsActive, lib, "g_vfs_is_active")
	core.PuregoSafeRegister(&xVfsParseName, lib, "g_vfs_parse_name")
	core.PuregoSafeRegister(&xVfsRegisterUriScheme, lib, "g_vfs_register_uri_scheme")
	core.PuregoSafeRegister(&xVfsUnregisterUriScheme, lib, "g_vfs_unregister_uri_scheme")

	core.PuregoSafeRegister(&xVfsGetDefault, lib, "g_vfs_get_default")
	core.PuregoSafeRegister(&xVfsGetLocal, lib, "g_vfs_get_local")

}
