// Package gio was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gio

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/gobject"
)

type FilterOutputStreamClass struct {
	ParentClass uintptr
}

// Base class for output stream implementations that perform some
// kind of filtering operation on a base stream. Typical examples
// of filtering operations are character set conversion, compression
// and byte order flipping.
type FilterOutputStream struct {
	OutputStream
}

func FilterOutputStreamNewFromInternalPtr(ptr uintptr) *FilterOutputStream {
	cls := &FilterOutputStream{}
	cls.Ptr = ptr
	return cls
}

var xFilterOutputStreamGetBaseStream func(uintptr) uintptr

// Gets the base stream for the filter stream.
func (x *FilterOutputStream) GetBaseStream() *OutputStream {
	var cls *OutputStream

	cret := xFilterOutputStreamGetBaseStream(x.GoPointer())

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &OutputStream{}
	cls.Ptr = cret
	return cls
}

var xFilterOutputStreamGetCloseBaseStream func(uintptr) bool

// Returns whether the base stream will be closed when @stream is
// closed.
func (x *FilterOutputStream) GetCloseBaseStream() bool {

	cret := xFilterOutputStreamGetCloseBaseStream(x.GoPointer())
	return cret
}

var xFilterOutputStreamSetCloseBaseStream func(uintptr, bool)

// Sets whether the base stream will be closed when @stream is closed.
func (x *FilterOutputStream) SetCloseBaseStream(CloseBaseVar bool) {

	xFilterOutputStreamSetCloseBaseStream(x.GoPointer(), CloseBaseVar)

}

func (c *FilterOutputStream) GoPointer() uintptr {
	return c.Ptr
}

func (c *FilterOutputStream) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GIO"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xFilterOutputStreamGetBaseStream, lib, "g_filter_output_stream_get_base_stream")
	core.PuregoSafeRegister(&xFilterOutputStreamGetCloseBaseStream, lib, "g_filter_output_stream_get_close_base_stream")
	core.PuregoSafeRegister(&xFilterOutputStreamSetCloseBaseStream, lib, "g_filter_output_stream_set_close_base_stream")

}
