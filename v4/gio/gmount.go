// Package gio was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gio

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/glib"
)

// Interface for implementing operations for mounts.
type MountIface struct {
	GIface uintptr
}

// The #GMount interface represents user-visible mounts. Note, when
// porting from GnomeVFS, #GMount is the moral equivalent of #GnomeVFSVolume.
//
// #GMount is a "mounted" filesystem that you can access. Mounted is in
// quotes because it's not the same as a unix mount, it might be a gvfs
// mount, but you can still access the files on it if you use GIO. Might or
// might not be related to a volume object.
//
// Unmounting a #GMount instance is an asynchronous operation. For
// more information about asynchronous operations, see #GAsyncResult
// and #GTask. To unmount a #GMount instance, first call
// g_mount_unmount_with_operation() with (at least) the #GMount instance and a
// #GAsyncReadyCallback.  The callback will be fired when the
// operation has resolved (either with success or failure), and a
// #GAsyncResult structure will be passed to the callback.  That
// callback should then call g_mount_unmount_with_operation_finish() with the #GMount
// and the #GAsyncResult data to see if the operation was completed
// successfully.  If an @error is present when g_mount_unmount_with_operation_finish()
// is called, then it will be filled with any error information.
type Mount interface {
	GoPointer() uintptr
	SetGoPointer(uintptr)
	CanEject() bool
	CanUnmount() bool
	Eject(FlagsVar MountUnmountFlags, CancellableVar *Cancellable, CallbackVar AsyncReadyCallback, UserDataVar uintptr)
	EjectFinish(ResultVar AsyncResult) bool
	EjectWithOperation(FlagsVar MountUnmountFlags, MountOperationVar *MountOperation, CancellableVar *Cancellable, CallbackVar AsyncReadyCallback, UserDataVar uintptr)
	EjectWithOperationFinish(ResultVar AsyncResult) bool
	GetDefaultLocation() *FileBase
	GetDrive() *DriveBase
	GetIcon() *IconBase
	GetName() string
	GetRoot() *FileBase
	GetSortKey() string
	GetSymbolicIcon() *IconBase
	GetUuid() string
	GetVolume() *VolumeBase
	GuessContentType(ForceRescanVar bool, CancellableVar *Cancellable, CallbackVar AsyncReadyCallback, UserDataVar uintptr)
	GuessContentTypeFinish(ResultVar AsyncResult) uintptr
	GuessContentTypeSync(ForceRescanVar bool, CancellableVar *Cancellable) uintptr
	IsShadowed() bool
	Remount(FlagsVar MountMountFlags, MountOperationVar *MountOperation, CancellableVar *Cancellable, CallbackVar AsyncReadyCallback, UserDataVar uintptr)
	RemountFinish(ResultVar AsyncResult) bool
	Shadow()
	Unmount(FlagsVar MountUnmountFlags, CancellableVar *Cancellable, CallbackVar AsyncReadyCallback, UserDataVar uintptr)
	UnmountFinish(ResultVar AsyncResult) bool
	UnmountWithOperation(FlagsVar MountUnmountFlags, MountOperationVar *MountOperation, CancellableVar *Cancellable, CallbackVar AsyncReadyCallback, UserDataVar uintptr)
	UnmountWithOperationFinish(ResultVar AsyncResult) bool
	Unshadow()
}
type MountBase struct {
	Ptr uintptr
}

func (x *MountBase) GoPointer() uintptr {
	return x.Ptr
}

func (x *MountBase) SetGoPointer(ptr uintptr) {
	x.Ptr = ptr
}

// Checks if @mount can be ejected.
func (x *MountBase) CanEject() bool {

	cret := XGMountCanEject(x.GoPointer())
	return cret
}

// Checks if @mount can be unmounted.
func (x *MountBase) CanUnmount() bool {

	cret := XGMountCanUnmount(x.GoPointer())
	return cret
}

// Ejects a mount. This is an asynchronous operation, and is
// finished by calling g_mount_eject_finish() with the @mount
// and #GAsyncResult data returned in the @callback.
func (x *MountBase) Eject(FlagsVar MountUnmountFlags, CancellableVar *Cancellable, CallbackVar AsyncReadyCallback, UserDataVar uintptr) {

	XGMountEject(x.GoPointer(), FlagsVar, CancellableVar.GoPointer(), purego.NewCallback(CallbackVar), UserDataVar)

}

// Finishes ejecting a mount. If any errors occurred during the operation,
// @error will be set to contain the errors and %FALSE will be returned.
func (x *MountBase) EjectFinish(ResultVar AsyncResult) (bool, error) {
	var cerr *glib.Error

	cret := XGMountEjectFinish(x.GoPointer(), ResultVar.GoPointer(), &cerr)
	if cerr == nil {
		return cret, nil
	}
	return cret, cerr

}

// Ejects a mount. This is an asynchronous operation, and is
// finished by calling g_mount_eject_with_operation_finish() with the @mount
// and #GAsyncResult data returned in the @callback.
func (x *MountBase) EjectWithOperation(FlagsVar MountUnmountFlags, MountOperationVar *MountOperation, CancellableVar *Cancellable, CallbackVar AsyncReadyCallback, UserDataVar uintptr) {

	XGMountEjectWithOperation(x.GoPointer(), FlagsVar, MountOperationVar.GoPointer(), CancellableVar.GoPointer(), purego.NewCallback(CallbackVar), UserDataVar)

}

// Finishes ejecting a mount. If any errors occurred during the operation,
// @error will be set to contain the errors and %FALSE will be returned.
func (x *MountBase) EjectWithOperationFinish(ResultVar AsyncResult) (bool, error) {
	var cerr *glib.Error

	cret := XGMountEjectWithOperationFinish(x.GoPointer(), ResultVar.GoPointer(), &cerr)
	if cerr == nil {
		return cret, nil
	}
	return cret, cerr

}

// Gets the default location of @mount. The default location of the given
// @mount is a path that reflects the main entry point for the user (e.g.
// the home directory, or the root of the volume).
func (x *MountBase) GetDefaultLocation() *FileBase {
	var cls *FileBase

	cret := XGMountGetDefaultLocation(x.GoPointer())

	if cret == 0 {
		return cls
	}
	cls = &FileBase{}
	cls.Ptr = cret
	return cls
}

// Gets the drive for the @mount.
//
// This is a convenience method for getting the #GVolume and then
// using that object to get the #GDrive.
func (x *MountBase) GetDrive() *DriveBase {
	var cls *DriveBase

	cret := XGMountGetDrive(x.GoPointer())

	if cret == 0 {
		return cls
	}
	cls = &DriveBase{}
	cls.Ptr = cret
	return cls
}

// Gets the icon for @mount.
func (x *MountBase) GetIcon() *IconBase {
	var cls *IconBase

	cret := XGMountGetIcon(x.GoPointer())

	if cret == 0 {
		return cls
	}
	cls = &IconBase{}
	cls.Ptr = cret
	return cls
}

// Gets the name of @mount.
func (x *MountBase) GetName() string {

	cret := XGMountGetName(x.GoPointer())
	return cret
}

// Gets the root directory on @mount.
func (x *MountBase) GetRoot() *FileBase {
	var cls *FileBase

	cret := XGMountGetRoot(x.GoPointer())

	if cret == 0 {
		return cls
	}
	cls = &FileBase{}
	cls.Ptr = cret
	return cls
}

// Gets the sort key for @mount, if any.
func (x *MountBase) GetSortKey() string {

	cret := XGMountGetSortKey(x.GoPointer())
	return cret
}

// Gets the symbolic icon for @mount.
func (x *MountBase) GetSymbolicIcon() *IconBase {
	var cls *IconBase

	cret := XGMountGetSymbolicIcon(x.GoPointer())

	if cret == 0 {
		return cls
	}
	cls = &IconBase{}
	cls.Ptr = cret
	return cls
}

// Gets the UUID for the @mount. The reference is typically based on
// the file system UUID for the mount in question and should be
// considered an opaque string. Returns %NULL if there is no UUID
// available.
func (x *MountBase) GetUuid() string {

	cret := XGMountGetUuid(x.GoPointer())
	return cret
}

// Gets the volume for the @mount.
func (x *MountBase) GetVolume() *VolumeBase {
	var cls *VolumeBase

	cret := XGMountGetVolume(x.GoPointer())

	if cret == 0 {
		return cls
	}
	cls = &VolumeBase{}
	cls.Ptr = cret
	return cls
}

// Tries to guess the type of content stored on @mount. Returns one or
// more textual identifiers of well-known content types (typically
// prefixed with "x-content/"), e.g. x-content/image-dcf for camera
// memory cards. See the
// [shared-mime-info](http://www.freedesktop.org/wiki/Specifications/shared-mime-info-spec)
// specification for more on x-content types.
//
// This is an asynchronous operation (see
// g_mount_guess_content_type_sync() for the synchronous version), and
// is finished by calling g_mount_guess_content_type_finish() with the
// @mount and #GAsyncResult data returned in the @callback.
func (x *MountBase) GuessContentType(ForceRescanVar bool, CancellableVar *Cancellable, CallbackVar AsyncReadyCallback, UserDataVar uintptr) {

	XGMountGuessContentType(x.GoPointer(), ForceRescanVar, CancellableVar.GoPointer(), purego.NewCallback(CallbackVar), UserDataVar)

}

// Finishes guessing content types of @mount. If any errors occurred
// during the operation, @error will be set to contain the errors and
// %FALSE will be returned. In particular, you may get an
// %G_IO_ERROR_NOT_SUPPORTED if the mount does not support content
// guessing.
func (x *MountBase) GuessContentTypeFinish(ResultVar AsyncResult) (uintptr, error) {
	var cerr *glib.Error

	cret := XGMountGuessContentTypeFinish(x.GoPointer(), ResultVar.GoPointer(), &cerr)
	if cerr == nil {
		return cret, nil
	}
	return cret, cerr

}

// Tries to guess the type of content stored on @mount. Returns one or
// more textual identifiers of well-known content types (typically
// prefixed with "x-content/"), e.g. x-content/image-dcf for camera
// memory cards. See the
// [shared-mime-info](http://www.freedesktop.org/wiki/Specifications/shared-mime-info-spec)
// specification for more on x-content types.
//
// This is a synchronous operation and as such may block doing IO;
// see g_mount_guess_content_type() for the asynchronous version.
func (x *MountBase) GuessContentTypeSync(ForceRescanVar bool, CancellableVar *Cancellable) (uintptr, error) {
	var cerr *glib.Error

	cret := XGMountGuessContentTypeSync(x.GoPointer(), ForceRescanVar, CancellableVar.GoPointer(), &cerr)
	if cerr == nil {
		return cret, nil
	}
	return cret, cerr

}

// Determines if @mount is shadowed. Applications or libraries should
// avoid displaying @mount in the user interface if it is shadowed.
//
// A mount is said to be shadowed if there exists one or more user
// visible objects (currently #GMount objects) with a root that is
// inside the root of @mount.
//
// One application of shadow mounts is when exposing a single file
// system that is used to address several logical volumes. In this
// situation, a #GVolumeMonitor implementation would create two
// #GVolume objects (for example, one for the camera functionality of
// the device and one for a SD card reader on the device) with
// activation URIs `gphoto2://[usb:001,002]/store1/`
// and `gphoto2://[usb:001,002]/store2/`. When the
// underlying mount (with root
// `gphoto2://[usb:001,002]/`) is mounted, said
// #GVolumeMonitor implementation would create two #GMount objects
// (each with their root matching the corresponding volume activation
// root) that would shadow the original mount.
//
// The proxy monitor in GVfs 2.26 and later, automatically creates and
// manage shadow mounts (and shadows the underlying mount) if the
// activation root on a #GVolume is set.
func (x *MountBase) IsShadowed() bool {

	cret := XGMountIsShadowed(x.GoPointer())
	return cret
}

// Remounts a mount. This is an asynchronous operation, and is
// finished by calling g_mount_remount_finish() with the @mount
// and #GAsyncResults data returned in the @callback.
//
// Remounting is useful when some setting affecting the operation
// of the volume has been changed, as these may need a remount to
// take affect. While this is semantically equivalent with unmounting
// and then remounting not all backends might need to actually be
// unmounted.
func (x *MountBase) Remount(FlagsVar MountMountFlags, MountOperationVar *MountOperation, CancellableVar *Cancellable, CallbackVar AsyncReadyCallback, UserDataVar uintptr) {

	XGMountRemount(x.GoPointer(), FlagsVar, MountOperationVar.GoPointer(), CancellableVar.GoPointer(), purego.NewCallback(CallbackVar), UserDataVar)

}

// Finishes remounting a mount. If any errors occurred during the operation,
// @error will be set to contain the errors and %FALSE will be returned.
func (x *MountBase) RemountFinish(ResultVar AsyncResult) (bool, error) {
	var cerr *glib.Error

	cret := XGMountRemountFinish(x.GoPointer(), ResultVar.GoPointer(), &cerr)
	if cerr == nil {
		return cret, nil
	}
	return cret, cerr

}

// Increments the shadow count on @mount. Usually used by
// #GVolumeMonitor implementations when creating a shadow mount for
// @mount, see g_mount_is_shadowed() for more information. The caller
// will need to emit the #GMount::changed signal on @mount manually.
func (x *MountBase) Shadow() {

	XGMountShadow(x.GoPointer())

}

// Unmounts a mount. This is an asynchronous operation, and is
// finished by calling g_mount_unmount_finish() with the @mount
// and #GAsyncResult data returned in the @callback.
func (x *MountBase) Unmount(FlagsVar MountUnmountFlags, CancellableVar *Cancellable, CallbackVar AsyncReadyCallback, UserDataVar uintptr) {

	XGMountUnmount(x.GoPointer(), FlagsVar, CancellableVar.GoPointer(), purego.NewCallback(CallbackVar), UserDataVar)

}

// Finishes unmounting a mount. If any errors occurred during the operation,
// @error will be set to contain the errors and %FALSE will be returned.
func (x *MountBase) UnmountFinish(ResultVar AsyncResult) (bool, error) {
	var cerr *glib.Error

	cret := XGMountUnmountFinish(x.GoPointer(), ResultVar.GoPointer(), &cerr)
	if cerr == nil {
		return cret, nil
	}
	return cret, cerr

}

// Unmounts a mount. This is an asynchronous operation, and is
// finished by calling g_mount_unmount_with_operation_finish() with the @mount
// and #GAsyncResult data returned in the @callback.
func (x *MountBase) UnmountWithOperation(FlagsVar MountUnmountFlags, MountOperationVar *MountOperation, CancellableVar *Cancellable, CallbackVar AsyncReadyCallback, UserDataVar uintptr) {

	XGMountUnmountWithOperation(x.GoPointer(), FlagsVar, MountOperationVar.GoPointer(), CancellableVar.GoPointer(), purego.NewCallback(CallbackVar), UserDataVar)

}

// Finishes unmounting a mount. If any errors occurred during the operation,
// @error will be set to contain the errors and %FALSE will be returned.
func (x *MountBase) UnmountWithOperationFinish(ResultVar AsyncResult) (bool, error) {
	var cerr *glib.Error

	cret := XGMountUnmountWithOperationFinish(x.GoPointer(), ResultVar.GoPointer(), &cerr)
	if cerr == nil {
		return cret, nil
	}
	return cret, cerr

}

// Decrements the shadow count on @mount. Usually used by
// #GVolumeMonitor implementations when destroying a shadow mount for
// @mount, see g_mount_is_shadowed() for more information. The caller
// will need to emit the #GMount::changed signal on @mount manually.
func (x *MountBase) Unshadow() {

	XGMountUnshadow(x.GoPointer())

}

var XGMountCanEject func(uintptr) bool
var XGMountCanUnmount func(uintptr) bool
var XGMountEject func(uintptr, MountUnmountFlags, uintptr, uintptr, uintptr)
var XGMountEjectFinish func(uintptr, uintptr, **glib.Error) bool
var XGMountEjectWithOperation func(uintptr, MountUnmountFlags, uintptr, uintptr, uintptr, uintptr)
var XGMountEjectWithOperationFinish func(uintptr, uintptr, **glib.Error) bool
var XGMountGetDefaultLocation func(uintptr) uintptr
var XGMountGetDrive func(uintptr) uintptr
var XGMountGetIcon func(uintptr) uintptr
var XGMountGetName func(uintptr) string
var XGMountGetRoot func(uintptr) uintptr
var XGMountGetSortKey func(uintptr) string
var XGMountGetSymbolicIcon func(uintptr) uintptr
var XGMountGetUuid func(uintptr) string
var XGMountGetVolume func(uintptr) uintptr
var XGMountGuessContentType func(uintptr, bool, uintptr, uintptr, uintptr)
var XGMountGuessContentTypeFinish func(uintptr, uintptr, **glib.Error) uintptr
var XGMountGuessContentTypeSync func(uintptr, bool, uintptr, **glib.Error) uintptr
var XGMountIsShadowed func(uintptr) bool
var XGMountRemount func(uintptr, MountMountFlags, uintptr, uintptr, uintptr, uintptr)
var XGMountRemountFinish func(uintptr, uintptr, **glib.Error) bool
var XGMountShadow func(uintptr)
var XGMountUnmount func(uintptr, MountUnmountFlags, uintptr, uintptr, uintptr)
var XGMountUnmountFinish func(uintptr, uintptr, **glib.Error) bool
var XGMountUnmountWithOperation func(uintptr, MountUnmountFlags, uintptr, uintptr, uintptr, uintptr)
var XGMountUnmountWithOperationFinish func(uintptr, uintptr, **glib.Error) bool
var XGMountUnshadow func(uintptr)

func init() {
	lib, err := purego.Dlopen(core.GetPath("GIO"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&XGMountCanEject, lib, "g_mount_can_eject")
	core.PuregoSafeRegister(&XGMountCanUnmount, lib, "g_mount_can_unmount")
	core.PuregoSafeRegister(&XGMountEject, lib, "g_mount_eject")
	core.PuregoSafeRegister(&XGMountEjectFinish, lib, "g_mount_eject_finish")
	core.PuregoSafeRegister(&XGMountEjectWithOperation, lib, "g_mount_eject_with_operation")
	core.PuregoSafeRegister(&XGMountEjectWithOperationFinish, lib, "g_mount_eject_with_operation_finish")
	core.PuregoSafeRegister(&XGMountGetDefaultLocation, lib, "g_mount_get_default_location")
	core.PuregoSafeRegister(&XGMountGetDrive, lib, "g_mount_get_drive")
	core.PuregoSafeRegister(&XGMountGetIcon, lib, "g_mount_get_icon")
	core.PuregoSafeRegister(&XGMountGetName, lib, "g_mount_get_name")
	core.PuregoSafeRegister(&XGMountGetRoot, lib, "g_mount_get_root")
	core.PuregoSafeRegister(&XGMountGetSortKey, lib, "g_mount_get_sort_key")
	core.PuregoSafeRegister(&XGMountGetSymbolicIcon, lib, "g_mount_get_symbolic_icon")
	core.PuregoSafeRegister(&XGMountGetUuid, lib, "g_mount_get_uuid")
	core.PuregoSafeRegister(&XGMountGetVolume, lib, "g_mount_get_volume")
	core.PuregoSafeRegister(&XGMountGuessContentType, lib, "g_mount_guess_content_type")
	core.PuregoSafeRegister(&XGMountGuessContentTypeFinish, lib, "g_mount_guess_content_type_finish")
	core.PuregoSafeRegister(&XGMountGuessContentTypeSync, lib, "g_mount_guess_content_type_sync")
	core.PuregoSafeRegister(&XGMountIsShadowed, lib, "g_mount_is_shadowed")
	core.PuregoSafeRegister(&XGMountRemount, lib, "g_mount_remount")
	core.PuregoSafeRegister(&XGMountRemountFinish, lib, "g_mount_remount_finish")
	core.PuregoSafeRegister(&XGMountShadow, lib, "g_mount_shadow")
	core.PuregoSafeRegister(&XGMountUnmount, lib, "g_mount_unmount")
	core.PuregoSafeRegister(&XGMountUnmountFinish, lib, "g_mount_unmount_finish")
	core.PuregoSafeRegister(&XGMountUnmountWithOperation, lib, "g_mount_unmount_with_operation")
	core.PuregoSafeRegister(&XGMountUnmountWithOperationFinish, lib, "g_mount_unmount_with_operation_finish")
	core.PuregoSafeRegister(&XGMountUnshadow, lib, "g_mount_unshadow")

}
