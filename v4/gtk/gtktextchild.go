// Package gtk was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gtk

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/gobject"
)

type TextChildAnchorClass struct {
	ParentClass uintptr
}

// A `GtkTextChildAnchor` is a spot in a `GtkTextBuffer` where child widgets can
// be “anchored”.
//
// The anchor can have multiple widgets anchored, to allow for multiple views.
type TextChildAnchor struct {
	gobject.Object
}

func TextChildAnchorNewFromInternalPtr(ptr uintptr) *TextChildAnchor {
	cls := &TextChildAnchor{}
	cls.Ptr = ptr
	return cls
}

var xNewTextChildAnchor func() uintptr

// Creates a new `GtkTextChildAnchor`.
//
// Usually you would then insert it into a `GtkTextBuffer` with
// [method@Gtk.TextBuffer.insert_child_anchor]. To perform the
// creation and insertion in one step, use the convenience
// function [method@Gtk.TextBuffer.create_child_anchor].
func NewTextChildAnchor() *TextChildAnchor {
	NewTextChildAnchorPtr := xNewTextChildAnchor()
	if NewTextChildAnchorPtr == 0 {
		return nil
	}

	NewTextChildAnchorCls := &TextChildAnchor{}
	NewTextChildAnchorCls.Ptr = NewTextChildAnchorPtr
	return NewTextChildAnchorCls
}

var xNewWithReplacementTextChildAnchor func(string) uintptr

// Creates a new `GtkTextChildAnchor` with the given replacement character.
//
// Usually you would then insert it into a `GtkTextBuffer` with
// [method@Gtk.TextBuffer.insert_child_anchor].
func NewWithReplacementTextChildAnchor(CharacterVar string) *TextChildAnchor {
	NewWithReplacementTextChildAnchorPtr := xNewWithReplacementTextChildAnchor(CharacterVar)
	if NewWithReplacementTextChildAnchorPtr == 0 {
		return nil
	}

	NewWithReplacementTextChildAnchorCls := &TextChildAnchor{}
	NewWithReplacementTextChildAnchorCls.Ptr = NewWithReplacementTextChildAnchorPtr
	return NewWithReplacementTextChildAnchorCls
}

var xTextChildAnchorGetDeleted func(uintptr) bool

// Determines whether a child anchor has been deleted from
// the buffer.
//
// Keep in mind that the child anchor will be unreferenced
// when removed from the buffer, so you need to hold your own
// reference (with g_object_ref()) if you plan to use this
// function — otherwise all deleted child anchors will also
// be finalized.
func (x *TextChildAnchor) GetDeleted() bool {

	return xTextChildAnchorGetDeleted(x.GoPointer())

}

var xTextChildAnchorGetWidgets func(uintptr, uint) uintptr

// Gets a list of all widgets anchored at this child anchor.
//
// The order in which the widgets are returned is not defined.
func (x *TextChildAnchor) GetWidgets(OutLenVar uint) uintptr {

	return xTextChildAnchorGetWidgets(x.GoPointer(), OutLenVar)

}

func (c *TextChildAnchor) GoPointer() uintptr {
	return c.Ptr
}

func (c *TextChildAnchor) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GTK"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xNewTextChildAnchor, lib, "gtk_text_child_anchor_new")
	core.PuregoSafeRegister(&xNewWithReplacementTextChildAnchor, lib, "gtk_text_child_anchor_new_with_replacement")

	core.PuregoSafeRegister(&xTextChildAnchorGetDeleted, lib, "gtk_text_child_anchor_get_deleted")
	core.PuregoSafeRegister(&xTextChildAnchorGetWidgets, lib, "gtk_text_child_anchor_get_widgets")

}
