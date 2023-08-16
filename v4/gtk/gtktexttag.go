// Package gtk was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gtk

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/gobject"
)

type TextTagClass struct {
	ParentClass uintptr

	Padding uintptr
}

type TextTagPrivate struct {
}

// A tag that can be applied to text contained in a `GtkTextBuffer`.
//
// You may wish to begin by reading the
// [text widget conceptual overview](section-text-widget.html),
// which gives an overview of all the objects and data types
// related to the text widget and how they work together.
//
// Tags should be in the [class@Gtk.TextTagTable] for a given
// `GtkTextBuffer` before using them with that buffer.
//
// [method@Gtk.TextBuffer.create_tag] is the best way to create tags.
// See “gtk4-demo” for numerous examples.
//
// For each property of `GtkTextTag`, there is a “set” property, e.g.
// “font-set” corresponds to “font”. These “set” properties reflect
// whether a property has been set or not.
//
// They are maintained by GTK and you should not set them independently.
type TextTag struct {
	gobject.Object
}

func TextTagNewFromInternalPtr(ptr uintptr) *TextTag {
	cls := &TextTag{}
	cls.Ptr = ptr
	return cls
}

var xNewTextTag func(string) uintptr

// Creates a `GtkTextTag`.
func NewTextTag(NameVar string) *TextTag {
	var cls *TextTag

	cret := xNewTextTag(NameVar)

	if cret == 0 {
		return cls
	}
	cls = &TextTag{}
	cls.Ptr = cret
	return cls
}

var xTextTagChanged func(uintptr, bool)

// Emits the [signal@Gtk.TextTagTable::tag-changed] signal on the
// `GtkTextTagTable` where the tag is included.
//
// The signal is already emitted when setting a `GtkTextTag` property.
// This function is useful for a `GtkTextTag` subclass.
func (x *TextTag) Changed(SizeChangedVar bool) {

	xTextTagChanged(x.GoPointer(), SizeChangedVar)

}

var xTextTagGetPriority func(uintptr) int

// Get the tag priority.
func (x *TextTag) GetPriority() int {

	cret := xTextTagGetPriority(x.GoPointer())
	return cret
}

var xTextTagSetPriority func(uintptr, int)

// Sets the priority of a `GtkTextTag`.
//
// Valid priorities start at 0 and go to one less than
// [method@Gtk.TextTagTable.get_size]. Each tag in a table
// has a unique priority; setting the priority of one tag shifts
// the priorities of all the other tags in the table to maintain
// a unique priority for each tag.
//
// Higher priority tags “win” if two tags both set the same text
// attribute. When adding a tag to a tag table, it will be assigned
// the highest priority in the table by default; so normally the
// precedence of a set of tags is the order in which they were added
// to the table, or created with [method@Gtk.TextBuffer.create_tag],
// which adds the tag to the buffer’s table automatically.
func (x *TextTag) SetPriority(PriorityVar int) {

	xTextTagSetPriority(x.GoPointer(), PriorityVar)

}

func (c *TextTag) GoPointer() uintptr {
	return c.Ptr
}

func (c *TextTag) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GTK"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xNewTextTag, lib, "gtk_text_tag_new")

	core.PuregoSafeRegister(&xTextTagChanged, lib, "gtk_text_tag_changed")
	core.PuregoSafeRegister(&xTextTagGetPriority, lib, "gtk_text_tag_get_priority")
	core.PuregoSafeRegister(&xTextTagSetPriority, lib, "gtk_text_tag_set_priority")

}
