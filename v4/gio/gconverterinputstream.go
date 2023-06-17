// Package gio was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gio

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/glib"
	"github.com/jwijenbergh/puregotk/v4/gobject"
)

type ConverterInputStreamClass struct {
	ParentClass uintptr
}

type ConverterInputStreamPrivate struct {
}

// Converter input stream implements #GInputStream and allows
// conversion of data of various types during reading.
//
// As of GLib 2.34, #GConverterInputStream implements
// #GPollableInputStream.
type ConverterInputStream struct {
	FilterInputStream
}

func ConverterInputStreamNewFromInternalPtr(ptr uintptr) *ConverterInputStream {
	cls := &ConverterInputStream{}
	cls.Ptr = ptr
	return cls
}

var xNewConverterInputStream func(uintptr, uintptr) uintptr

// Creates a new converter input stream for the @base_stream.
func NewConverterInputStream(BaseStreamVar *InputStream, ConverterVar Converter) *InputStream {
	NewConverterInputStreamPtr := xNewConverterInputStream(BaseStreamVar.GoPointer(), ConverterVar.GoPointer())
	if NewConverterInputStreamPtr == 0 {
		return nil
	}

	NewConverterInputStreamCls := &InputStream{}
	NewConverterInputStreamCls.Ptr = NewConverterInputStreamPtr
	return NewConverterInputStreamCls
}

var xConverterInputStreamGetConverter func(uintptr) uintptr

// Gets the #GConverter that is used by @converter_stream.
func (x *ConverterInputStream) GetConverter() *ConverterBase {

	GetConverterPtr := xConverterInputStreamGetConverter(x.GoPointer())
	if GetConverterPtr == 0 {
		return nil
	}

	gobject.IncreaseRef(GetConverterPtr)

	GetConverterCls := &ConverterBase{}
	GetConverterCls.Ptr = GetConverterPtr
	return GetConverterCls

}

func (c *ConverterInputStream) GoPointer() uintptr {
	return c.Ptr
}

func (c *ConverterInputStream) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// Checks if @stream is actually pollable. Some classes may implement
// #GPollableInputStream but have only certain instances of that class
// be pollable. If this method returns %FALSE, then the behavior of
// other #GPollableInputStream methods is undefined.
//
// For any given stream, the value returned by this method is constant;
// a stream cannot switch from pollable to non-pollable or vice versa.
func (x *ConverterInputStream) CanPoll() bool {

	return XGPollableInputStreamCanPoll(x.GoPointer())

}

// Creates a #GSource that triggers when @stream can be read, or
// @cancellable is triggered or an error occurs. The callback on the
// source is of the #GPollableSourceFunc type.
//
// As with g_pollable_input_stream_is_readable(), it is possible that
// the stream may not actually be readable even after the source
// triggers, so you should use g_pollable_input_stream_read_nonblocking()
// rather than g_input_stream_read() from the callback.
func (x *ConverterInputStream) CreateSource(CancellableVar *Cancellable) *glib.Source {

	return XGPollableInputStreamCreateSource(x.GoPointer(), CancellableVar.GoPointer())

}

// Checks if @stream can be read.
//
// Note that some stream types may not be able to implement this 100%
// reliably, and it is possible that a call to g_input_stream_read()
// after this returns %TRUE would still block. To guarantee
// non-blocking behavior, you should always use
// g_pollable_input_stream_read_nonblocking(), which will return a
// %G_IO_ERROR_WOULD_BLOCK error rather than blocking.
func (x *ConverterInputStream) IsReadable() bool {

	return XGPollableInputStreamIsReadable(x.GoPointer())

}

// Attempts to read up to @count bytes from @stream into @buffer, as
// with g_input_stream_read(). If @stream is not currently readable,
// this will immediately return %G_IO_ERROR_WOULD_BLOCK, and you can
// use g_pollable_input_stream_create_source() to create a #GSource
// that will be triggered when @stream is readable.
//
// Note that since this method never blocks, you cannot actually
// use @cancellable to cancel it. However, it will return an error
// if @cancellable has already been cancelled when you call, which
// may happen if you call this method after a source triggers due
// to having been cancelled.
func (x *ConverterInputStream) ReadNonblocking(BufferVar uintptr, CountVar uint, CancellableVar *Cancellable) int {

	return XGPollableInputStreamReadNonblocking(x.GoPointer(), BufferVar, CountVar, CancellableVar.GoPointer())

}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GIO"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xNewConverterInputStream, lib, "g_converter_input_stream_new")

	core.PuregoSafeRegister(&xConverterInputStreamGetConverter, lib, "g_converter_input_stream_get_converter")

}
