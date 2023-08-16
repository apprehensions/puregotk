// Package gtk was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gtk

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
)

type BinLayoutClass struct {
	ParentClass uintptr
}

// `GtkBinLayout` is a `GtkLayoutManager` subclass useful for create "bins" of
// widgets.
//
// `GtkBinLayout` will stack each child of a widget on top of each other,
// using the [property@Gtk.Widget:hexpand], [property@Gtk.Widget:vexpand],
// [property@Gtk.Widget:halign], and [property@Gtk.Widget:valign] properties
// of each child to determine where they should be positioned.
type BinLayout struct {
	LayoutManager
}

func BinLayoutNewFromInternalPtr(ptr uintptr) *BinLayout {
	cls := &BinLayout{}
	cls.Ptr = ptr
	return cls
}

var xNewBinLayout func() uintptr

// Creates a new `GtkBinLayout` instance.
func NewBinLayout() *LayoutManager {
	var cls *LayoutManager

	cret := xNewBinLayout()

	if cret == 0 {
		return cls
	}
	cls = &LayoutManager{}
	cls.Ptr = cret
	return cls
}

func (c *BinLayout) GoPointer() uintptr {
	return c.Ptr
}

func (c *BinLayout) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GTK"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xNewBinLayout, lib, "gtk_bin_layout_new")

}
