// Package gtk was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gtk

import (
	"unsafe"

	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/glib"
	"github.com/jwijenbergh/puregotk/v4/gobject"
)

type EntryBufferClass struct {
	ParentClass uintptr
}

func (x *EntryBufferClass) GoPointer() uintptr {
	return uintptr(unsafe.Pointer(x))
}

// A `GtkEntryBuffer` hold the text displayed in a `GtkText` widget.
//
// A single `GtkEntryBuffer` object can be shared by multiple widgets
// which will then share the same text content, but not the cursor
// position, visibility attributes, icon etc.
//
// `GtkEntryBuffer` may be derived from. Such a derived class might allow
// text to be stored in an alternate location, such as non-pageable memory,
// useful in the case of important passwords. Or a derived class could
// integrate with an application’s concept of undo/redo.
type EntryBuffer struct {
	gobject.Object
}

func EntryBufferNewFromInternalPtr(ptr uintptr) *EntryBuffer {
	cls := &EntryBuffer{}
	cls.Ptr = ptr
	return cls
}

var xNewEntryBuffer func(string, int) uintptr

// Create a new `GtkEntryBuffer` object.
//
// Optionally, specify initial text to set in the buffer.
func NewEntryBuffer(InitialCharsVar string, NInitialCharsVar int) *EntryBuffer {
	var cls *EntryBuffer

	cret := xNewEntryBuffer(InitialCharsVar, NInitialCharsVar)

	if cret == 0 {
		return nil
	}
	cls = &EntryBuffer{}
	cls.Ptr = cret
	return cls
}

var xEntryBufferDeleteText func(uintptr, uint, int) uint

// Deletes a sequence of characters from the buffer.
//
// @n_chars characters are deleted starting at @position.
// If @n_chars is negative, then all characters until the
// end of the text are deleted.
//
// If @position or @n_chars are out of bounds, then they
// are coerced to sane values.
//
// Note that the positions are specified in characters,
// not bytes.
func (x *EntryBuffer) DeleteText(PositionVar uint, NCharsVar int) uint {

	cret := xEntryBufferDeleteText(x.GoPointer(), PositionVar, NCharsVar)
	return cret
}

var xEntryBufferEmitDeletedText func(uintptr, uint, uint)

// Used when subclassing `GtkEntryBuffer`.
func (x *EntryBuffer) EmitDeletedText(PositionVar uint, NCharsVar uint) {

	xEntryBufferEmitDeletedText(x.GoPointer(), PositionVar, NCharsVar)

}

var xEntryBufferEmitInsertedText func(uintptr, uint, string, uint)

// Used when subclassing `GtkEntryBuffer`.
func (x *EntryBuffer) EmitInsertedText(PositionVar uint, CharsVar string, NCharsVar uint) {

	xEntryBufferEmitInsertedText(x.GoPointer(), PositionVar, CharsVar, NCharsVar)

}

var xEntryBufferGetBytes func(uintptr) uint

// Retrieves the length in bytes of the buffer.
//
// See [method@Gtk.EntryBuffer.get_length].
func (x *EntryBuffer) GetBytes() uint {

	cret := xEntryBufferGetBytes(x.GoPointer())
	return cret
}

var xEntryBufferGetLength func(uintptr) uint

// Retrieves the length in characters of the buffer.
func (x *EntryBuffer) GetLength() uint {

	cret := xEntryBufferGetLength(x.GoPointer())
	return cret
}

var xEntryBufferGetMaxLength func(uintptr) int

// Retrieves the maximum allowed length of the text in @buffer.
func (x *EntryBuffer) GetMaxLength() int {

	cret := xEntryBufferGetMaxLength(x.GoPointer())
	return cret
}

var xEntryBufferGetText func(uintptr) string

// Retrieves the contents of the buffer.
//
// The memory pointer returned by this call will not change
// unless this object emits a signal, or is finalized.
func (x *EntryBuffer) GetText() string {

	cret := xEntryBufferGetText(x.GoPointer())
	return cret
}

var xEntryBufferInsertText func(uintptr, uint, string, int) uint

// Inserts @n_chars characters of @chars into the contents of the
// buffer, at position @position.
//
// If @n_chars is negative, then characters from chars will be inserted
// until a null-terminator is found. If @position or @n_chars are out of
// bounds, or the maximum buffer text length is exceeded, then they are
// coerced to sane values.
//
// Note that the position and length are in characters, not in bytes.
func (x *EntryBuffer) InsertText(PositionVar uint, CharsVar string, NCharsVar int) uint {

	cret := xEntryBufferInsertText(x.GoPointer(), PositionVar, CharsVar, NCharsVar)
	return cret
}

var xEntryBufferSetMaxLength func(uintptr, int)

// Sets the maximum allowed length of the contents of the buffer.
//
// If the current contents are longer than the given length, then
// they will be truncated to fit.
func (x *EntryBuffer) SetMaxLength(MaxLengthVar int) {

	xEntryBufferSetMaxLength(x.GoPointer(), MaxLengthVar)

}

var xEntryBufferSetText func(uintptr, string, int)

// Sets the text in the buffer.
//
// This is roughly equivalent to calling
// [method@Gtk.EntryBuffer.delete_text] and
// [method@Gtk.EntryBuffer.insert_text].
//
// Note that @n_chars is in characters, not in bytes.
func (x *EntryBuffer) SetText(CharsVar string, NCharsVar int) {

	xEntryBufferSetText(x.GoPointer(), CharsVar, NCharsVar)

}

func (c *EntryBuffer) GoPointer() uintptr {
	return c.Ptr
}

func (c *EntryBuffer) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// The text is altered in the default handler for this signal.
//
// If you want access to the text after the text has been modified,
// use %G_CONNECT_AFTER.
func (x *EntryBuffer) ConnectDeletedText(cb *func(EntryBuffer, uint, uint)) uint32 {
	cbPtr := uintptr(unsafe.Pointer(cb))
	if cbRefPtr, ok := glib.GetCallback(cbPtr); ok {
		return gobject.SignalConnect(x.GoPointer(), "deleted-text", cbRefPtr)
	}

	fcb := func(clsPtr uintptr, PositionVarp uint, NCharsVarp uint) {
		fa := EntryBuffer{}
		fa.Ptr = clsPtr
		cbFn := *cb

		cbFn(fa, PositionVarp, NCharsVarp)

	}
	cbRefPtr := purego.NewCallback(fcb)
	glib.SaveCallback(cbPtr, cbRefPtr)
	return gobject.SignalConnect(x.GoPointer(), "deleted-text", cbRefPtr)
}

// This signal is emitted after text is inserted into the buffer.
func (x *EntryBuffer) ConnectInsertedText(cb *func(EntryBuffer, uint, string, uint)) uint32 {
	cbPtr := uintptr(unsafe.Pointer(cb))
	if cbRefPtr, ok := glib.GetCallback(cbPtr); ok {
		return gobject.SignalConnect(x.GoPointer(), "inserted-text", cbRefPtr)
	}

	fcb := func(clsPtr uintptr, PositionVarp uint, CharsVarp string, NCharsVarp uint) {
		fa := EntryBuffer{}
		fa.Ptr = clsPtr
		cbFn := *cb

		cbFn(fa, PositionVarp, CharsVarp, NCharsVarp)

	}
	cbRefPtr := purego.NewCallback(fcb)
	glib.SaveCallback(cbPtr, cbRefPtr)
	return gobject.SignalConnect(x.GoPointer(), "inserted-text", cbRefPtr)
}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GTK"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xNewEntryBuffer, lib, "gtk_entry_buffer_new")

	core.PuregoSafeRegister(&xEntryBufferDeleteText, lib, "gtk_entry_buffer_delete_text")
	core.PuregoSafeRegister(&xEntryBufferEmitDeletedText, lib, "gtk_entry_buffer_emit_deleted_text")
	core.PuregoSafeRegister(&xEntryBufferEmitInsertedText, lib, "gtk_entry_buffer_emit_inserted_text")
	core.PuregoSafeRegister(&xEntryBufferGetBytes, lib, "gtk_entry_buffer_get_bytes")
	core.PuregoSafeRegister(&xEntryBufferGetLength, lib, "gtk_entry_buffer_get_length")
	core.PuregoSafeRegister(&xEntryBufferGetMaxLength, lib, "gtk_entry_buffer_get_max_length")
	core.PuregoSafeRegister(&xEntryBufferGetText, lib, "gtk_entry_buffer_get_text")
	core.PuregoSafeRegister(&xEntryBufferInsertText, lib, "gtk_entry_buffer_insert_text")
	core.PuregoSafeRegister(&xEntryBufferSetMaxLength, lib, "gtk_entry_buffer_set_max_length")
	core.PuregoSafeRegister(&xEntryBufferSetText, lib, "gtk_entry_buffer_set_text")

}
