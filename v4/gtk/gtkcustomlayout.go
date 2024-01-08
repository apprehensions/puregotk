// Package gtk was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gtk

import (
	"unsafe"

	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
)

// A function to be used by `GtkCustomLayout` to allocate a widget.
type CustomAllocateFunc func(uintptr, int, int, int)

// A function to be used by `GtkCustomLayout` to measure a widget.
type CustomMeasureFunc func(uintptr, Orientation, int, int, int, int, int)

// Queries a widget for its preferred size request mode.
type CustomRequestModeFunc func(uintptr) SizeRequestMode

type CustomLayoutClass struct {
	ParentClass uintptr
}

func (x *CustomLayoutClass) GoPointer() uintptr {
	return uintptr(unsafe.Pointer(x))
}

// `GtkCustomLayout` uses closures for size negotiation.
//
// A `GtkCustomLayout `uses closures matching to the old `GtkWidget`
// virtual functions for size negotiation, as a convenience API to
// ease the porting towards the corresponding `GtkLayoutManager
// virtual functions.
type CustomLayout struct {
	LayoutManager
}

func CustomLayoutNewFromInternalPtr(ptr uintptr) *CustomLayout {
	cls := &CustomLayout{}
	cls.Ptr = ptr
	return cls
}

var xNewCustomLayout func(uintptr, uintptr, uintptr) uintptr

// Creates a new legacy layout manager.
//
// Legacy layout managers map to the old `GtkWidget` size negotiation
// virtual functions, and are meant to be used during the transition
// from layout containers to layout manager delegates.
func NewCustomLayout(RequestModeVar CustomRequestModeFunc, MeasureVar CustomMeasureFunc, AllocateVar CustomAllocateFunc) *LayoutManager {
	var cls *LayoutManager

	cret := xNewCustomLayout(purego.NewCallback(RequestModeVar), purego.NewCallback(MeasureVar), purego.NewCallback(AllocateVar))

	if cret == 0 {
		return nil
	}
	cls = &LayoutManager{}
	cls.Ptr = cret
	return cls
}

func (c *CustomLayout) GoPointer() uintptr {
	return c.Ptr
}

func (c *CustomLayout) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GTK"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xNewCustomLayout, lib, "gtk_custom_layout_new")

}
