// Package gtk was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gtk

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/gdk"
	"github.com/jwijenbergh/puregotk/v4/gio"
	"github.com/jwijenbergh/puregotk/v4/gobject"
)

type MountOperationClass struct {
	ParentClass uintptr
}

type MountOperationPrivate struct {
}

// `GtkMountOperation` is an implementation of `GMountOperation`.
//
// The functions and objects described here make working with GTK and
// GIO more convenient.
//
// `GtkMountOperation` is needed when mounting volumes:
// It is an implementation of `GMountOperation` that can be used with
// GIO functions for mounting volumes such as
// g_file_mount_enclosing_volume(), g_file_mount_mountable(),
// g_volume_mount(), g_mount_unmount_with_operation() and others.
//
// When necessary, `GtkMountOperation` shows dialogs to let the user
// enter passwords, ask questions or show processes blocking unmount.
type MountOperation struct {
	gio.MountOperation
}

func MountOperationNewFromInternalPtr(ptr uintptr) *MountOperation {
	cls := &MountOperation{}
	cls.Ptr = ptr
	return cls
}

var xNewMountOperation func(uintptr) uintptr

// Creates a new `GtkMountOperation`.
func NewMountOperation(ParentVar *Window) *gio.MountOperation {
	NewMountOperationPtr := xNewMountOperation(ParentVar.GoPointer())
	if NewMountOperationPtr == 0 {
		return nil
	}

	NewMountOperationCls := &gio.MountOperation{}
	NewMountOperationCls.Ptr = NewMountOperationPtr
	return NewMountOperationCls
}

var xMountOperationGetDisplay func(uintptr) uintptr

// Gets the display on which windows of the `GtkMountOperation`
// will be shown.
func (x *MountOperation) GetDisplay() *gdk.Display {

	GetDisplayPtr := xMountOperationGetDisplay(x.GoPointer())
	if GetDisplayPtr == 0 {
		return nil
	}

	gobject.IncreaseRef(GetDisplayPtr)

	GetDisplayCls := &gdk.Display{}
	GetDisplayCls.Ptr = GetDisplayPtr
	return GetDisplayCls

}

var xMountOperationGetParent func(uintptr) uintptr

// Gets the transient parent used by the `GtkMountOperation`.
func (x *MountOperation) GetParent() *Window {

	GetParentPtr := xMountOperationGetParent(x.GoPointer())
	if GetParentPtr == 0 {
		return nil
	}

	gobject.IncreaseRef(GetParentPtr)

	GetParentCls := &Window{}
	GetParentCls.Ptr = GetParentPtr
	return GetParentCls

}

var xMountOperationIsShowing func(uintptr) bool

// Returns whether the `GtkMountOperation` is currently displaying
// a window.
func (x *MountOperation) IsShowing() bool {

	return xMountOperationIsShowing(x.GoPointer())

}

var xMountOperationSetDisplay func(uintptr, uintptr)

// Sets the display to show windows of the `GtkMountOperation` on.
func (x *MountOperation) SetDisplay(DisplayVar *gdk.Display) {

	xMountOperationSetDisplay(x.GoPointer(), DisplayVar.GoPointer())

}

var xMountOperationSetParent func(uintptr, uintptr)

// Sets the transient parent for windows shown by the
// `GtkMountOperation`.
func (x *MountOperation) SetParent(ParentVar *Window) {

	xMountOperationSetParent(x.GoPointer(), ParentVar.GoPointer())

}

func (c *MountOperation) GoPointer() uintptr {
	return c.Ptr
}

func (c *MountOperation) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GTK"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xNewMountOperation, lib, "gtk_mount_operation_new")

	core.PuregoSafeRegister(&xMountOperationGetDisplay, lib, "gtk_mount_operation_get_display")
	core.PuregoSafeRegister(&xMountOperationGetParent, lib, "gtk_mount_operation_get_parent")
	core.PuregoSafeRegister(&xMountOperationIsShowing, lib, "gtk_mount_operation_is_showing")
	core.PuregoSafeRegister(&xMountOperationSetDisplay, lib, "gtk_mount_operation_set_display")
	core.PuregoSafeRegister(&xMountOperationSetParent, lib, "gtk_mount_operation_set_parent")

}
