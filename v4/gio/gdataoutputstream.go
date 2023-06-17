// Package gio was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gio

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/glib"
)

type DataOutputStreamClass struct {
	ParentClass uintptr
}

type DataOutputStreamPrivate struct {
}

// Data output stream implements #GOutputStream and includes functions for
// writing data directly to an output stream.
type DataOutputStream struct {
	FilterOutputStream
}

func DataOutputStreamNewFromInternalPtr(ptr uintptr) *DataOutputStream {
	cls := &DataOutputStream{}
	cls.Ptr = ptr
	return cls
}

var xNewDataOutputStream func(uintptr) uintptr

// Creates a new data output stream for @base_stream.
func NewDataOutputStream(BaseStreamVar *OutputStream) *DataOutputStream {
	NewDataOutputStreamPtr := xNewDataOutputStream(BaseStreamVar.GoPointer())
	if NewDataOutputStreamPtr == 0 {
		return nil
	}

	NewDataOutputStreamCls := &DataOutputStream{}
	NewDataOutputStreamCls.Ptr = NewDataOutputStreamPtr
	return NewDataOutputStreamCls
}

var xDataOutputStreamGetByteOrder func(uintptr) DataStreamByteOrder

// Gets the byte order for the stream.
func (x *DataOutputStream) GetByteOrder() DataStreamByteOrder {

	return xDataOutputStreamGetByteOrder(x.GoPointer())

}

var xDataOutputStreamPutByte func(uintptr, byte, uintptr) bool

// Puts a byte into the output stream.
func (x *DataOutputStream) PutByte(DataVar byte, CancellableVar *Cancellable) bool {

	return xDataOutputStreamPutByte(x.GoPointer(), DataVar, CancellableVar.GoPointer())

}

var xDataOutputStreamPutInt16 func(uintptr, int16, uintptr) bool

// Puts a signed 16-bit integer into the output stream.
func (x *DataOutputStream) PutInt16(DataVar int16, CancellableVar *Cancellable) bool {

	return xDataOutputStreamPutInt16(x.GoPointer(), DataVar, CancellableVar.GoPointer())

}

var xDataOutputStreamPutInt32 func(uintptr, int32, uintptr) bool

// Puts a signed 32-bit integer into the output stream.
func (x *DataOutputStream) PutInt32(DataVar int32, CancellableVar *Cancellable) bool {

	return xDataOutputStreamPutInt32(x.GoPointer(), DataVar, CancellableVar.GoPointer())

}

var xDataOutputStreamPutInt64 func(uintptr, int64, uintptr) bool

// Puts a signed 64-bit integer into the stream.
func (x *DataOutputStream) PutInt64(DataVar int64, CancellableVar *Cancellable) bool {

	return xDataOutputStreamPutInt64(x.GoPointer(), DataVar, CancellableVar.GoPointer())

}

var xDataOutputStreamPutString func(uintptr, string, uintptr) bool

// Puts a string into the output stream.
func (x *DataOutputStream) PutString(StrVar string, CancellableVar *Cancellable) bool {

	return xDataOutputStreamPutString(x.GoPointer(), StrVar, CancellableVar.GoPointer())

}

var xDataOutputStreamPutUint16 func(uintptr, uint16, uintptr) bool

// Puts an unsigned 16-bit integer into the output stream.
func (x *DataOutputStream) PutUint16(DataVar uint16, CancellableVar *Cancellable) bool {

	return xDataOutputStreamPutUint16(x.GoPointer(), DataVar, CancellableVar.GoPointer())

}

var xDataOutputStreamPutUint32 func(uintptr, uint32, uintptr) bool

// Puts an unsigned 32-bit integer into the stream.
func (x *DataOutputStream) PutUint32(DataVar uint32, CancellableVar *Cancellable) bool {

	return xDataOutputStreamPutUint32(x.GoPointer(), DataVar, CancellableVar.GoPointer())

}

var xDataOutputStreamPutUint64 func(uintptr, uint64, uintptr) bool

// Puts an unsigned 64-bit integer into the stream.
func (x *DataOutputStream) PutUint64(DataVar uint64, CancellableVar *Cancellable) bool {

	return xDataOutputStreamPutUint64(x.GoPointer(), DataVar, CancellableVar.GoPointer())

}

var xDataOutputStreamSetByteOrder func(uintptr, DataStreamByteOrder)

// Sets the byte order of the data output stream to @order.
func (x *DataOutputStream) SetByteOrder(OrderVar DataStreamByteOrder) {

	xDataOutputStreamSetByteOrder(x.GoPointer(), OrderVar)

}

func (c *DataOutputStream) GoPointer() uintptr {
	return c.Ptr
}

func (c *DataOutputStream) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// Tests if the stream supports the #GSeekableIface.
func (x *DataOutputStream) CanSeek() bool {

	return XGSeekableCanSeek(x.GoPointer())

}

// Tests if the length of the stream can be adjusted with
// g_seekable_truncate().
func (x *DataOutputStream) CanTruncate() bool {

	return XGSeekableCanTruncate(x.GoPointer())

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
func (x *DataOutputStream) Seek(OffsetVar int64, TypeVar glib.SeekType, CancellableVar *Cancellable) bool {

	return XGSeekableSeek(x.GoPointer(), OffsetVar, TypeVar, CancellableVar.GoPointer())

}

// Tells the current position within the stream.
func (x *DataOutputStream) Tell() int64 {

	return XGSeekableTell(x.GoPointer())

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
func (x *DataOutputStream) Truncate(OffsetVar int64, CancellableVar *Cancellable) bool {

	return XGSeekableTruncate(x.GoPointer(), OffsetVar, CancellableVar.GoPointer())

}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GIO"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xNewDataOutputStream, lib, "g_data_output_stream_new")

	core.PuregoSafeRegister(&xDataOutputStreamGetByteOrder, lib, "g_data_output_stream_get_byte_order")
	core.PuregoSafeRegister(&xDataOutputStreamPutByte, lib, "g_data_output_stream_put_byte")
	core.PuregoSafeRegister(&xDataOutputStreamPutInt16, lib, "g_data_output_stream_put_int16")
	core.PuregoSafeRegister(&xDataOutputStreamPutInt32, lib, "g_data_output_stream_put_int32")
	core.PuregoSafeRegister(&xDataOutputStreamPutInt64, lib, "g_data_output_stream_put_int64")
	core.PuregoSafeRegister(&xDataOutputStreamPutString, lib, "g_data_output_stream_put_string")
	core.PuregoSafeRegister(&xDataOutputStreamPutUint16, lib, "g_data_output_stream_put_uint16")
	core.PuregoSafeRegister(&xDataOutputStreamPutUint32, lib, "g_data_output_stream_put_uint32")
	core.PuregoSafeRegister(&xDataOutputStreamPutUint64, lib, "g_data_output_stream_put_uint64")
	core.PuregoSafeRegister(&xDataOutputStreamSetByteOrder, lib, "g_data_output_stream_set_byte_order")

}
