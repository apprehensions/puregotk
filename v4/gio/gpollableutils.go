// Package gio was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gio

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/glib"
	"github.com/jwijenbergh/puregotk/v4/gobject"
)

var xPollableSourceNew func(uintptr) *glib.Source

// Utility method for #GPollableInputStream and #GPollableOutputStream
// implementations. Creates a new #GSource that expects a callback of
// type #GPollableSourceFunc. The new source does not actually do
// anything on its own; use g_source_add_child_source() to add other
// sources to it to cause it to trigger.
func PollableSourceNew(PollableStreamVar *gobject.Object) *glib.Source {

	return xPollableSourceNew(PollableStreamVar.GoPointer())

}

var xPollableSourceNewFull func(uintptr, *glib.Source, uintptr) *glib.Source

// Utility method for #GPollableInputStream and #GPollableOutputStream
// implementations. Creates a new #GSource, as with
// g_pollable_source_new(), but also attaching @child_source (with a
// dummy callback), and @cancellable, if they are non-%NULL.
func PollableSourceNewFull(PollableStreamVar *gobject.Object, ChildSourceVar *glib.Source, CancellableVar *Cancellable) *glib.Source {

	return xPollableSourceNewFull(PollableStreamVar.GoPointer(), ChildSourceVar, CancellableVar.GoPointer())

}

var xPollableStreamRead func(uintptr, uintptr, uint, bool, uintptr) int

// Tries to read from @stream, as with g_input_stream_read() (if
// @blocking is %TRUE) or g_pollable_input_stream_read_nonblocking()
// (if @blocking is %FALSE). This can be used to more easily share
// code between blocking and non-blocking implementations of a method.
//
// If @blocking is %FALSE, then @stream must be a
// #GPollableInputStream for which g_pollable_input_stream_can_poll()
// returns %TRUE, or else the behavior is undefined. If @blocking is
// %TRUE, then @stream does not need to be a #GPollableInputStream.
func PollableStreamRead(StreamVar *InputStream, BufferVar uintptr, CountVar uint, BlockingVar bool, CancellableVar *Cancellable) int {

	return xPollableStreamRead(StreamVar.GoPointer(), BufferVar, CountVar, BlockingVar, CancellableVar.GoPointer())

}

var xPollableStreamWrite func(uintptr, uintptr, uint, bool, uintptr) int

// Tries to write to @stream, as with g_output_stream_write() (if
// @blocking is %TRUE) or g_pollable_output_stream_write_nonblocking()
// (if @blocking is %FALSE). This can be used to more easily share
// code between blocking and non-blocking implementations of a method.
//
// If @blocking is %FALSE, then @stream must be a
// #GPollableOutputStream for which
// g_pollable_output_stream_can_poll() returns %TRUE or else the
// behavior is undefined. If @blocking is %TRUE, then @stream does not
// need to be a #GPollableOutputStream.
func PollableStreamWrite(StreamVar *OutputStream, BufferVar uintptr, CountVar uint, BlockingVar bool, CancellableVar *Cancellable) int {

	return xPollableStreamWrite(StreamVar.GoPointer(), BufferVar, CountVar, BlockingVar, CancellableVar.GoPointer())

}

var xPollableStreamWriteAll func(uintptr, uintptr, uint, bool, uint, uintptr) bool

// Tries to write @count bytes to @stream, as with
// g_output_stream_write_all(), but using g_pollable_stream_write()
// rather than g_output_stream_write().
//
// On a successful write of @count bytes, %TRUE is returned, and
// @bytes_written is set to @count.
//
// If there is an error during the operation (including
// %G_IO_ERROR_WOULD_BLOCK in the non-blocking case), %FALSE is
// returned and @error is set to indicate the error status,
// @bytes_written is updated to contain the number of bytes written
// into the stream before the error occurred.
//
// As with g_pollable_stream_write(), if @blocking is %FALSE, then
// @stream must be a #GPollableOutputStream for which
// g_pollable_output_stream_can_poll() returns %TRUE or else the
// behavior is undefined. If @blocking is %TRUE, then @stream does not
// need to be a #GPollableOutputStream.
func PollableStreamWriteAll(StreamVar *OutputStream, BufferVar uintptr, CountVar uint, BlockingVar bool, BytesWrittenVar uint, CancellableVar *Cancellable) bool {

	return xPollableStreamWriteAll(StreamVar.GoPointer(), BufferVar, CountVar, BlockingVar, BytesWrittenVar, CancellableVar.GoPointer())

}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GIO"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	core.PuregoSafeRegister(&xPollableSourceNew, lib, "g_pollable_source_new")
	core.PuregoSafeRegister(&xPollableSourceNewFull, lib, "g_pollable_source_new_full")
	core.PuregoSafeRegister(&xPollableStreamRead, lib, "g_pollable_stream_read")
	core.PuregoSafeRegister(&xPollableStreamWrite, lib, "g_pollable_stream_write")
	core.PuregoSafeRegister(&xPollableStreamWriteAll, lib, "g_pollable_stream_write_all")

}
