// Package gtk was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gtk

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/glib"
	"github.com/jwijenbergh/puregotk/v4/gobject"
)

type WindowGroupClass struct {
	ParentClass uintptr
}

// `GtkWindowGroup` makes group of windows behave like separate applications.
//
// It achieves this by limiting the effect of GTK grabs and modality
// to windows in the same group.
//
// A window can be a member in at most one window group at a time.
// Windows that have not been explicitly assigned to a group are
// implicitly treated like windows of the default window group.
//
// `GtkWindowGroup` objects are referenced by each window in the group,
// so once you have added all windows to a `GtkWindowGroup`, you can drop
// the initial reference to the window group with g_object_unref(). If the
// windows in the window group are subsequently destroyed, then they will
// be removed from the window group and drop their references on the window
// group; when all window have been removed, the window group will be
// freed.
type WindowGroup struct {
	gobject.Object
}

func WindowGroupNewFromInternalPtr(ptr uintptr) *WindowGroup {
	cls := &WindowGroup{}
	cls.Ptr = ptr
	return cls
}

var xNewWindowGroup func() uintptr

// Creates a new `GtkWindowGroup` object.
//
// Modality of windows only affects windows
// within the same `GtkWindowGroup`.
func NewWindowGroup() *WindowGroup {
	var cls *WindowGroup

	cret := xNewWindowGroup()

	if cret == 0 {
		return nil
	}
	cls = &WindowGroup{}
	cls.Ptr = cret
	return cls
}

var xWindowGroupAddWindow func(uintptr, uintptr)

// Adds a window to a `GtkWindowGroup`.
func (x *WindowGroup) AddWindow(WindowVar *Window) {

	xWindowGroupAddWindow(x.GoPointer(), WindowVar.GoPointer())

}

var xWindowGroupListWindows func(uintptr) *glib.List

// Returns a list of the `GtkWindows` that belong to @window_group.
func (x *WindowGroup) ListWindows() *glib.List {

	cret := xWindowGroupListWindows(x.GoPointer())
	return cret
}

var xWindowGroupRemoveWindow func(uintptr, uintptr)

// Removes a window from a `GtkWindowGroup`.
func (x *WindowGroup) RemoveWindow(WindowVar *Window) {

	xWindowGroupRemoveWindow(x.GoPointer(), WindowVar.GoPointer())

}

func (c *WindowGroup) GoPointer() uintptr {
	return c.Ptr
}

func (c *WindowGroup) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GTK"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xNewWindowGroup, lib, "gtk_window_group_new")

	core.PuregoSafeRegister(&xWindowGroupAddWindow, lib, "gtk_window_group_add_window")
	core.PuregoSafeRegister(&xWindowGroupListWindows, lib, "gtk_window_group_list_windows")
	core.PuregoSafeRegister(&xWindowGroupRemoveWindow, lib, "gtk_window_group_remove_window")

}
