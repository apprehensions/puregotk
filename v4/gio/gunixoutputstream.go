// Package gio was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gio

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/glib"
)

type UnixOutputStreamClass struct {
	ParentClass uintptr
}

type UnixOutputStreamPrivate struct {
}

// #GUnixOutputStream implements #GOutputStream for writing to a UNIX
// file descriptor, including asynchronous operations. (If the file
// descriptor refers to a socket or pipe, this will use poll() to do
// asynchronous I/O. If it refers to a regular file, it will fall back
// to doing asynchronous I/O in another thread.)
//
// Note that `&lt;gio/gunixoutputstream.h&gt;` belongs to the UNIX-specific GIO
// interfaces, thus you have to use the `gio-unix-2.0.pc` pkg-config file
// when using it.
type UnixOutputStream struct {
	OutputStream
}

func UnixOutputStreamNewFromInternalPtr(ptr uintptr) *UnixOutputStream {
	cls := &UnixOutputStream{}
	cls.Ptr = ptr
	return cls
}

var xNewUnixOutputStream func(int, bool) uintptr

// Creates a new #GUnixOutputStream for the given @fd.
//
// If @close_fd, is %TRUE, the file descriptor will be closed when
// the output stream is destroyed.
func NewUnixOutputStream(FdVar int, CloseFdVar bool) *OutputStream {
	var cls *OutputStream

	cret := xNewUnixOutputStream(FdVar, CloseFdVar)

	if cret == 0 {
		return cls
	}
	cls = &OutputStream{}
	cls.Ptr = cret
	return cls
}

var xUnixOutputStreamGetCloseFd func(uintptr) bool

// Returns whether the file descriptor of @stream will be
// closed when the stream is closed.
func (x *UnixOutputStream) GetCloseFd() bool {

	cret := xUnixOutputStreamGetCloseFd(x.GoPointer())
	return cret
}

var xUnixOutputStreamGetFd func(uintptr) int

// Return the UNIX file descriptor that the stream writes to.
func (x *UnixOutputStream) GetFd() int {

	cret := xUnixOutputStreamGetFd(x.GoPointer())
	return cret
}

var xUnixOutputStreamSetCloseFd func(uintptr, bool)

// Sets whether the file descriptor of @stream shall be closed
// when the stream is closed.
func (x *UnixOutputStream) SetCloseFd(CloseFdVar bool) {

	xUnixOutputStreamSetCloseFd(x.GoPointer(), CloseFdVar)

}

func (c *UnixOutputStream) GoPointer() uintptr {
	return c.Ptr
}

func (c *UnixOutputStream) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// Checks if @stream is actually pollable. Some classes may implement
// #GPollableOutputStream but have only certain instances of that
// class be pollable. If this method returns %FALSE, then the behavior
// of other #GPollableOutputStream methods is undefined.
//
// For any given stream, the value returned by this method is constant;
// a stream cannot switch from pollable to non-pollable or vice versa.
func (x *UnixOutputStream) CanPoll() bool {

	cret := XGPollableOutputStreamCanPoll(x.GoPointer())
	return cret
}

// Creates a #GSource that triggers when @stream can be written, or
// @cancellable is triggered or an error occurs. The callback on the
// source is of the #GPollableSourceFunc type.
//
// As with g_pollable_output_stream_is_writable(), it is possible that
// the stream may not actually be writable even after the source
// triggers, so you should use g_pollable_output_stream_write_nonblocking()
// rather than g_output_stream_write() from the callback.
func (x *UnixOutputStream) CreateSource(CancellableVar *Cancellable) *glib.Source {

	cret := XGPollableOutputStreamCreateSource(x.GoPointer(), CancellableVar.GoPointer())
	return cret
}

// Checks if @stream can be written.
//
// Note that some stream types may not be able to implement this 100%
// reliably, and it is possible that a call to g_output_stream_write()
// after this returns %TRUE would still block. To guarantee
// non-blocking behavior, you should always use
// g_pollable_output_stream_write_nonblocking(), which will return a
// %G_IO_ERROR_WOULD_BLOCK error rather than blocking.
func (x *UnixOutputStream) IsWritable() bool {

	cret := XGPollableOutputStreamIsWritable(x.GoPointer())
	return cret
}

// Attempts to write up to @count bytes from @buffer to @stream, as
// with g_output_stream_write(). If @stream is not currently writable,
// this will immediately return %G_IO_ERROR_WOULD_BLOCK, and you can
// use g_pollable_output_stream_create_source() to create a #GSource
// that will be triggered when @stream is writable.
//
// Note that since this method never blocks, you cannot actually
// use @cancellable to cancel it. However, it will return an error
// if @cancellable has already been cancelled when you call, which
// may happen if you call this method after a source triggers due
// to having been cancelled.
//
// Also note that if %G_IO_ERROR_WOULD_BLOCK is returned some underlying
// transports like D/TLS require that you re-send the same @buffer and
// @count in the next write call.
func (x *UnixOutputStream) WriteNonblocking(BufferVar uintptr, CountVar uint, CancellableVar *Cancellable) (int, error) {
	var cerr *glib.Error

	cret := XGPollableOutputStreamWriteNonblocking(x.GoPointer(), BufferVar, CountVar, CancellableVar.GoPointer(), &cerr)
	if cerr == nil {
		return cret, nil
	}
	return cret, cerr

}

// Attempts to write the bytes contained in the @n_vectors @vectors to @stream,
// as with g_output_stream_writev(). If @stream is not currently writable,
// this will immediately return %@G_POLLABLE_RETURN_WOULD_BLOCK, and you can
// use g_pollable_output_stream_create_source() to create a #GSource
// that will be triggered when @stream is writable. @error will *not* be
// set in that case.
//
// Note that since this method never blocks, you cannot actually
// use @cancellable to cancel it. However, it will return an error
// if @cancellable has already been cancelled when you call, which
// may happen if you call this method after a source triggers due
// to having been cancelled.
//
// Also note that if %G_POLLABLE_RETURN_WOULD_BLOCK is returned some underlying
// transports like D/TLS require that you re-send the same @vectors and
// @n_vectors in the next write call.
func (x *UnixOutputStream) WritevNonblocking(VectorsVar uintptr, NVectorsVar uint, BytesWrittenVar uint, CancellableVar *Cancellable) (PollableReturn, error) {
	var cerr *glib.Error

	cret := XGPollableOutputStreamWritevNonblocking(x.GoPointer(), VectorsVar, NVectorsVar, BytesWrittenVar, CancellableVar.GoPointer(), &cerr)
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

	core.PuregoSafeRegister(&xNewUnixOutputStream, lib, "g_unix_output_stream_new")

	core.PuregoSafeRegister(&xUnixOutputStreamGetCloseFd, lib, "g_unix_output_stream_get_close_fd")
	core.PuregoSafeRegister(&xUnixOutputStreamGetFd, lib, "g_unix_output_stream_get_fd")
	core.PuregoSafeRegister(&xUnixOutputStreamSetCloseFd, lib, "g_unix_output_stream_set_close_fd")

}
