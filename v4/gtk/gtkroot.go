// Package gtk was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gtk

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/gdk"
	"github.com/jwijenbergh/puregotk/v4/gobject"
)

type RootInterface struct {
}

// `GtkRoot` is the interface implemented by all widgets that can act as a toplevel
// widget.
//
// The root widget takes care of providing the connection to the windowing system
// and manages layout, drawing and event delivery for its widget hierarchy.
//
// The obvious example of a `GtkRoot` is `GtkWindow`.
//
// To get the display to which a `GtkRoot` belongs, use
// [method@Gtk.Root.get_display].
//
// `GtkRoot` also maintains the location of keyboard focus inside its widget
// hierarchy, with [method@Gtk.Root.set_focus] and [method@Gtk.Root.get_focus].
type Root interface {
	GoPointer() uintptr
	SetGoPointer(uintptr)
	GetDisplay() *gdk.Display
	GetFocus() *Widget
	SetFocus(FocusVar *Widget)
}
type RootBase struct {
	Ptr uintptr
}

func (x *RootBase) GoPointer() uintptr {
	return x.Ptr
}

func (x *RootBase) SetGoPointer(ptr uintptr) {
	x.Ptr = ptr
}

// Returns the display that this `GtkRoot` is on.
func (x *RootBase) GetDisplay() *gdk.Display {
	var cls *gdk.Display

	cret := XGtkRootGetDisplay(x.GoPointer())

	if cret == 0 {
		return cls
	}
	gobject.IncreaseRef(cret)
	cls = &gdk.Display{}
	cls.Ptr = cret
	return cls
}

// Retrieves the current focused widget within the root.
//
// Note that this is the widget that would have the focus
// if the root is active; if the root is not focused then
// `gtk_widget_has_focus (widget)` will be %FALSE for the
// widget.
func (x *RootBase) GetFocus() *Widget {
	var cls *Widget

	cret := XGtkRootGetFocus(x.GoPointer())

	if cret == 0 {
		return cls
	}
	gobject.IncreaseRef(cret)
	cls = &Widget{}
	cls.Ptr = cret
	return cls
}

// If @focus is not the current focus widget, and is focusable, sets
// it as the focus widget for the root.
//
// If @focus is %NULL, unsets the focus widget for the root.
//
// To set the focus to a particular widget in the root, it is usually
// more convenient to use [method@Gtk.Widget.grab_focus] instead of
// this function.
func (x *RootBase) SetFocus(FocusVar *Widget) {

	XGtkRootSetFocus(x.GoPointer(), FocusVar.GoPointer())

}

var XGtkRootGetDisplay func(uintptr) uintptr
var XGtkRootGetFocus func(uintptr) uintptr
var XGtkRootSetFocus func(uintptr, uintptr)

func init() {
	lib, err := purego.Dlopen(core.GetPath("GTK"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&XGtkRootGetDisplay, lib, "gtk_root_get_display")
	core.PuregoSafeRegister(&XGtkRootGetFocus, lib, "gtk_root_get_focus")
	core.PuregoSafeRegister(&XGtkRootSetFocus, lib, "gtk_root_set_focus")

}
