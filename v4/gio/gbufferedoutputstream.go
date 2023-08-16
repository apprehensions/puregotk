// Package gio was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gio

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/glib"
)

type BufferedOutputStreamClass struct {
	ParentClass uintptr
}

type BufferedOutputStreamPrivate struct {
}

// Buffered output stream implements #GFilterOutputStream and provides
// for buffered writes.
//
// By default, #GBufferedOutputStream's buffer size is set at 4 kilobytes.
//
// To create a buffered output stream, use g_buffered_output_stream_new(),
// or g_buffered_output_stream_new_sized() to specify the buffer's size
// at construction.
//
// To get the size of a buffer within a buffered input stream, use
// g_buffered_output_stream_get_buffer_size(). To change the size of a
// buffered output stream's buffer, use
// g_buffered_output_stream_set_buffer_size(). Note that the buffer's
// size cannot be reduced below the size of the data within the buffer.
type BufferedOutputStream struct {
	FilterOutputStream
}

func BufferedOutputStreamNewFromInternalPtr(ptr uintptr) *BufferedOutputStream {
	cls := &BufferedOutputStream{}
	cls.Ptr = ptr
	return cls
}

var xNewBufferedOutputStream func(uintptr) uintptr

// Creates a new buffered output stream for a base stream.
func NewBufferedOutputStream(BaseStreamVar *OutputStream) *OutputStream {
	var cls *OutputStream

	cret := xNewBufferedOutputStream(BaseStreamVar.GoPointer())

	if cret == 0 {
		return nil
	}
	cls = &OutputStream{}
	cls.Ptr = cret
	return cls
}

var xNewSizedBufferedOutputStream func(uintptr, uint) uintptr

// Creates a new buffered output stream with a given buffer size.
func NewSizedBufferedOutputStream(BaseStreamVar *OutputStream, SizeVar uint) *OutputStream {
	var cls *OutputStream

	cret := xNewSizedBufferedOutputStream(BaseStreamVar.GoPointer(), SizeVar)

	if cret == 0 {
		return nil
	}
	cls = &OutputStream{}
	cls.Ptr = cret
	return cls
}

var xBufferedOutputStreamGetAutoGrow func(uintptr) bool

// Checks if the buffer automatically grows as data is added.
func (x *BufferedOutputStream) GetAutoGrow() bool {

	cret := xBufferedOutputStreamGetAutoGrow(x.GoPointer())
	return cret
}

var xBufferedOutputStreamGetBufferSize func(uintptr) uint

// Gets the size of the buffer in the @stream.
func (x *BufferedOutputStream) GetBufferSize() uint {

	cret := xBufferedOutputStreamGetBufferSize(x.GoPointer())
	return cret
}

var xBufferedOutputStreamSetAutoGrow func(uintptr, bool)

// Sets whether or not the @stream's buffer should automatically grow.
// If @auto_grow is true, then each write will just make the buffer
// larger, and you must manually flush the buffer to actually write out
// the data to the underlying stream.
func (x *BufferedOutputStream) SetAutoGrow(AutoGrowVar bool) {

	xBufferedOutputStreamSetAutoGrow(x.GoPointer(), AutoGrowVar)

}

var xBufferedOutputStreamSetBufferSize func(uintptr, uint)

// Sets the size of the internal buffer to @size.
func (x *BufferedOutputStream) SetBufferSize(SizeVar uint) {

	xBufferedOutputStreamSetBufferSize(x.GoPointer(), SizeVar)

}

func (c *BufferedOutputStream) GoPointer() uintptr {
	return c.Ptr
}

func (c *BufferedOutputStream) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// Tests if the stream supports the #GSeekableIface.
func (x *BufferedOutputStream) CanSeek() bool {

	cret := XGSeekableCanSeek(x.GoPointer())
	return cret
}

// Tests if the length of the stream can be adjusted with
// g_seekable_truncate().
func (x *BufferedOutputStream) CanTruncate() bool {

	cret := XGSeekableCanTruncate(x.GoPointer())
	return cret
}

// Seeks in the stream by the given @offset, modified by @type.
//
// Attempting to seek past the end of the stream will have different
// results depending on if the stream is fixed-sized or resizable.  If
// the stream is resizable then seeking past the end and then writing
// will result in zeros filling the empty space.  Seeking past the end
// of a resizable stream and reading will result in EOF.  Seeking past
// the end of a fixed-sized stream will fail.
//
// Any operation that would result in a negative offset will fail.
//
// If @cancellable is not %NULL, then the operation can be cancelled by
// triggering the cancellable object from another thread. If the operation
// was cancelled, the error %G_IO_ERROR_CANCELLED will be returned.
func (x *BufferedOutputStream) Seek(OffsetVar int64, TypeVar glib.SeekType, CancellableVar *Cancellable) (bool, error) {
	var cerr *glib.Error

	cret := XGSeekableSeek(x.GoPointer(), OffsetVar, TypeVar, CancellableVar.GoPointer(), &cerr)
	if cerr == nil {
		return cret, nil
	}
	return cret, cerr

}

// Tells the current position within the stream.
func (x *BufferedOutputStream) Tell() int64 {

	cret := XGSeekableTell(x.GoPointer())
	return cret
}

// Sets the length of the stream to @offset. If the stream was previously
// larger than @offset, the extra data is discarded. If the stream was
// previously shorter than @offset, it is extended with NUL ('\0') bytes.
//
// If @cancellable is not %NULL, then the operation can be cancelled by
// triggering the cancellable object from another thread. If the operation
// was cancelled, the error %G_IO_ERROR_CANCELLED will be returned. If an
// operation was partially finished when the operation was cancelled the
// partial result will be returned, without an error.
func (x *BufferedOutputStream) Truncate(OffsetVar int64, CancellableVar *Cancellable) (bool, error) {
	var cerr *glib.Error

	cret := XGSeekableTruncate(x.GoPointer(), OffsetVar, CancellableVar.GoPointer(), &cerr)
	if cerr == nil {
		return cret, nil
	}
	return cret, cerr

}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GIO"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xNewBufferedOutputStream, lib, "g_buffered_output_stream_new")
	core.PuregoSafeRegister(&xNewSizedBufferedOutputStream, lib, "g_buffered_output_stream_new_sized")

	core.PuregoSafeRegister(&xBufferedOutputStreamGetAutoGrow, lib, "g_buffered_output_stream_get_auto_grow")
	core.PuregoSafeRegister(&xBufferedOutputStreamGetBufferSize, lib, "g_buffered_output_stream_get_buffer_size")
	core.PuregoSafeRegister(&xBufferedOutputStreamSetAutoGrow, lib, "g_buffered_output_stream_set_auto_grow")
	core.PuregoSafeRegister(&xBufferedOutputStreamSetBufferSize, lib, "g_buffered_output_stream_set_buffer_size")

}
