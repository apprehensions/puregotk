// Package gtk was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gtk

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/gobject"
)

type TextMarkClass struct {
	ParentClass uintptr

	Padding uintptr
}

// A `GtkTextMark` is a position in a `GtkTextbuffer` that is preserved
// across modifications.
//
// You may wish to begin by reading the
// [text widget conceptual overview](section-text-widget.html),
// which gives an overview of all the objects and data types
// related to the text widget and how they work together.
//
// A `GtkTextMark` is like a bookmark in a text buffer; it preserves
// a position in the text. You can convert the mark to an iterator using
// [method@Gtk.TextBuffer.get_iter_at_mark]. Unlike iterators, marks remain
// valid across buffer mutations, because their behavior is defined when
// text is inserted or deleted. When text containing a mark is deleted,
// the mark remains in the position originally occupied by the deleted
// text. When text is inserted at a mark, a mark with “left gravity” will
// be moved to the beginning of the newly-inserted text, and a mark with
// “right gravity” will be moved to the end.
//
// Note that “left” and “right” here refer to logical direction (left
// is the toward the start of the buffer); in some languages such as
// Hebrew the logically-leftmost text is not actually on the left when
// displayed.
//
// Marks are reference counted, but the reference count only controls
// the validity of the memory; marks can be deleted from the buffer at
// any time with [method@Gtk.TextBuffer.delete_mark]. Once deleted from
// the buffer, a mark is essentially useless.
//
// Marks optionally have names; these can be convenient to avoid passing
// the `GtkTextMark` object around.
//
// Marks are typically created using the [method@Gtk.TextBuffer.create_mark]
// function.
type TextMark struct {
	gobject.Object
}

func TextMarkNewFromInternalPtr(ptr uintptr) *TextMark {
	cls := &TextMark{}
	cls.Ptr = ptr
	return cls
}

var xNewTextMark func(string, bool) uintptr

// Creates a text mark.
//
// Add it to a buffer using [method@Gtk.TextBuffer.add_mark].
// If @name is %NULL, the mark is anonymous; otherwise, the mark can be
// retrieved by name using [method@Gtk.TextBuffer.get_mark]. If a mark
// has left gravity, and text is inserted at the mark’s current location,
// the mark will be moved to the left of the newly-inserted text. If the
// mark has right gravity (@left_gravity = %FALSE), the mark will end up
// on the right of newly-inserted text. The standard left-to-right cursor
// is a mark with right gravity (when you type, the cursor stays on the
// right side of the text you’re typing).
func NewTextMark(NameVar string, LeftGravityVar bool) *TextMark {
	NewTextMarkPtr := xNewTextMark(NameVar, LeftGravityVar)
	if NewTextMarkPtr == 0 {
		return nil
	}

	NewTextMarkCls := &TextMark{}
	NewTextMarkCls.Ptr = NewTextMarkPtr
	return NewTextMarkCls
}

var xTextMarkGetBuffer func(uintptr) uintptr

// Gets the buffer this mark is located inside.
//
// Returns %NULL if the mark is deleted.
func (x *TextMark) GetBuffer() *TextBuffer {

	GetBufferPtr := xTextMarkGetBuffer(x.GoPointer())
	if GetBufferPtr == 0 {
		return nil
	}

	gobject.IncreaseRef(GetBufferPtr)

	GetBufferCls := &TextBuffer{}
	GetBufferCls.Ptr = GetBufferPtr
	return GetBufferCls

}

var xTextMarkGetDeleted func(uintptr) bool

// Returns %TRUE if the mark has been removed from its buffer.
//
// See [method@Gtk.TextBuffer.add_mark] for a way to add it
// to a buffer again.
func (x *TextMark) GetDeleted() bool {

	return xTextMarkGetDeleted(x.GoPointer())

}

var xTextMarkGetLeftGravity func(uintptr) bool

// Determines whether the mark has left gravity.
func (x *TextMark) GetLeftGravity() bool {

	return xTextMarkGetLeftGravity(x.GoPointer())

}

var xTextMarkGetName func(uintptr) string

// Returns the mark name.
//
// Returns %NULL for anonymous marks.
func (x *TextMark) GetName() string {

	return xTextMarkGetName(x.GoPointer())

}

var xTextMarkGetVisible func(uintptr) bool

// Returns %TRUE if the mark is visible.
//
// A cursor is displayed for visible marks.
func (x *TextMark) GetVisible() bool {

	return xTextMarkGetVisible(x.GoPointer())

}

var xTextMarkSetVisible func(uintptr, bool)

func (x *TextMark) SetVisible(SettingVar bool) {

	xTextMarkSetVisible(x.GoPointer(), SettingVar)

}

func (c *TextMark) GoPointer() uintptr {
	return c.Ptr
}

func (c *TextMark) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GTK"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xNewTextMark, lib, "gtk_text_mark_new")

	core.PuregoSafeRegister(&xTextMarkGetBuffer, lib, "gtk_text_mark_get_buffer")
	core.PuregoSafeRegister(&xTextMarkGetDeleted, lib, "gtk_text_mark_get_deleted")
	core.PuregoSafeRegister(&xTextMarkGetLeftGravity, lib, "gtk_text_mark_get_left_gravity")
	core.PuregoSafeRegister(&xTextMarkGetName, lib, "gtk_text_mark_get_name")
	core.PuregoSafeRegister(&xTextMarkGetVisible, lib, "gtk_text_mark_get_visible")
	core.PuregoSafeRegister(&xTextMarkSetVisible, lib, "gtk_text_mark_set_visible")

}
