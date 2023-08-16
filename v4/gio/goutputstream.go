// Package gio was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gio

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/glib"
	"github.com/jwijenbergh/puregotk/v4/gobject"
)

type OutputStreamClass struct {
	ParentClass uintptr
}

type OutputStreamPrivate struct {
}

// #GOutputStream has functions to write to a stream (g_output_stream_write()),
// to close a stream (g_output_stream_close()) and to flush pending writes
// (g_output_stream_flush()).
//
// To copy the content of an input stream to an output stream without
// manually handling the reads and writes, use g_output_stream_splice().
//
// See the documentation for #GIOStream for details of thread safety of
// streaming APIs.
//
// All of these functions have async variants too.
type OutputStream struct {
	gobject.Object
}

func OutputStreamNewFromInternalPtr(ptr uintptr) *OutputStream {
	cls := &OutputStream{}
	cls.Ptr = ptr
	return cls
}

var xOutputStreamClearPending func(uintptr)

// Clears the pending flag on @stream.
func (x *OutputStream) ClearPending() {

	xOutputStreamClearPending(x.GoPointer())

}

var xOutputStreamClose func(uintptr, uintptr, **glib.Error) bool

// Closes the stream, releasing resources related to it.
//
// Once the stream is closed, all other operations will return %G_IO_ERROR_CLOSED.
// Closing a stream multiple times will not return an error.
//
// Closing a stream will automatically flush any outstanding buffers in the
// stream.
//
// Streams will be automatically closed when the last reference
// is dropped, but you might want to call this function to make sure
// resources are released as early as possible.
//
// Some streams might keep the backing store of the stream (e.g. a file descriptor)
// open after the stream is closed. See the documentation for the individual
// stream for details.
//
// On failure the first error that happened will be reported, but the close
// operation will finish as much as possible. A stream that failed to
// close will still return %G_IO_ERROR_CLOSED for all operations. Still, it
// is important to check and report the error to the user, otherwise
// there might be a loss of data as all data might not be written.
//
// If @cancellable is not %NULL, then the operation can be cancelled by
// triggering the cancellable object from another thread. If the operation
// was cancelled, the error %G_IO_ERROR_CANCELLED will be returned.
// Cancelling a close will still leave the stream closed, but there some streams
// can use a faster close that doesn't block to e.g. check errors. On
// cancellation (as with any error) there is no guarantee that all written
// data will reach the target.
func (x *OutputStream) Close(CancellableVar *Cancellable) (bool, error) {
	var cerr *glib.Error

	cret := xOutputStreamClose(x.GoPointer(), CancellableVar.GoPointer(), &cerr)
	if cerr == nil {
		return cret, nil
	}
	return cret, cerr

}

var xOutputStreamCloseAsync func(uintptr, int, uintptr, uintptr, uintptr)

// Requests an asynchronous close of the stream, releasing resources
// related to it. When the operation is finished @callback will be
// called. You can then call g_output_stream_close_finish() to get
// the result of the operation.
//
// For behaviour details see g_output_stream_close().
//
// The asynchronous methods have a default fallback that uses threads
// to implement asynchronicity, so they are optional for inheriting
// classes. However, if you override one you must override all.
func (x *OutputStream) CloseAsync(IoPriorityVar int, CancellableVar *Cancellable, CallbackVar AsyncReadyCallback, UserDataVar uintptr) {

	xOutputStreamCloseAsync(x.GoPointer(), IoPriorityVar, CancellableVar.GoPointer(), purego.NewCallback(CallbackVar), UserDataVar)

}

var xOutputStreamCloseFinish func(uintptr, uintptr, **glib.Error) bool

// Closes an output stream.
func (x *OutputStream) CloseFinish(ResultVar AsyncResult) (bool, error) {
	var cerr *glib.Error

	cret := xOutputStreamCloseFinish(x.GoPointer(), ResultVar.GoPointer(), &cerr)
	if cerr == nil {
		return cret, nil
	}
	return cret, cerr

}

var xOutputStreamFlush func(uintptr, uintptr, **glib.Error) bool

// Forces a write of all user-space buffered data for the given
// @stream. Will block during the operation. Closing the stream will
// implicitly cause a flush.
//
// This function is optional for inherited classes.
//
// If @cancellable is not %NULL, then the operation can be cancelled by
// triggering the cancellable object from another thread. If the operation
// was cancelled, the error %G_IO_ERROR_CANCELLED will be returned.
func (x *OutputStream) Flush(CancellableVar *Cancellable) (bool, error) {
	var cerr *glib.Error

	cret := xOutputStreamFlush(x.GoPointer(), CancellableVar.GoPointer(), &cerr)
	if cerr == nil {
		return cret, nil
	}
	return cret, cerr

}

var xOutputStreamFlushAsync func(uintptr, int, uintptr, uintptr, uintptr)

// Forces an asynchronous write of all user-space buffered data for
// the given @stream.
// For behaviour details see g_output_stream_flush().
//
// When the operation is finished @callback will be
// called. You can then call g_output_stream_flush_finish() to get the
// result of the operation.
func (x *OutputStream) FlushAsync(IoPriorityVar int, CancellableVar *Cancellable, CallbackVar AsyncReadyCallback, UserDataVar uintptr) {

	xOutputStreamFlushAsync(x.GoPointer(), IoPriorityVar, CancellableVar.GoPointer(), purego.NewCallback(CallbackVar), UserDataVar)

}

var xOutputStreamFlushFinish func(uintptr, uintptr, **glib.Error) bool

// Finishes flushing an output stream.
func (x *OutputStream) FlushFinish(ResultVar AsyncResult) (bool, error) {
	var cerr *glib.Error

	cret := xOutputStreamFlushFinish(x.GoPointer(), ResultVar.GoPointer(), &cerr)
	if cerr == nil {
		return cret, nil
	}
	return cret, cerr

}

var xOutputStreamHasPending func(uintptr) bool

// Checks if an output stream has pending actions.
func (x *OutputStream) HasPending() bool {

	cret := xOutputStreamHasPending(x.GoPointer())
	return cret
}

var xOutputStreamIsClosed func(uintptr) bool

// Checks if an output stream has already been closed.
func (x *OutputStream) IsClosed() bool {

	cret := xOutputStreamIsClosed(x.GoPointer())
	return cret
}

var xOutputStreamIsClosing func(uintptr) bool

// Checks if an output stream is being closed. This can be
// used inside e.g. a flush implementation to see if the
// flush (or other i/o operation) is called from within
// the closing operation.
func (x *OutputStream) IsClosing() bool {

	cret := xOutputStreamIsClosing(x.GoPointer())
	return cret
}

var xOutputStreamPrintf func(uintptr, uint, uintptr, **glib.Error, string, ...interface{}) bool

// This is a utility function around g_output_stream_write_all(). It
// uses g_strdup_vprintf() to turn @format and @... into a string that
// is then written to @stream.
//
// See the documentation of g_output_stream_write_all() about the
// behavior of the actual write operation.
//
// Note that partial writes cannot be properly checked with this
// function due to the variable length of the written string, if you
// need precise control over partial write failures, you need to
// create you own printf()-like wrapper around g_output_stream_write()
// or g_output_stream_write_all().
func (x *OutputStream) Printf(BytesWrittenVar uint, CancellableVar *Cancellable, ErrorVar **glib.Error, FormatVar string, varArgs ...interface{}) bool {

	cret := xOutputStreamPrintf(x.GoPointer(), BytesWrittenVar, CancellableVar.GoPointer(), ErrorVar, FormatVar, varArgs...)
	return cret
}

var xOutputStreamSetPending func(uintptr) bool

// Sets @stream to have actions pending. If the pending flag is
// already set or @stream is closed, it will return %FALSE and set
// @error.
func (x *OutputStream) SetPending() (bool, error) {
	var cerr *glib.Error

	cret := xOutputStreamSetPending(x.GoPointer())
	if cerr == nil {
		return cret, nil
	}
	return cret, cerr

}

var xOutputStreamSplice func(uintptr, uintptr, OutputStreamSpliceFlags, uintptr, **glib.Error) int

// Splices an input stream into an output stream.
func (x *OutputStream) Splice(SourceVar *InputStream, FlagsVar OutputStreamSpliceFlags, CancellableVar *Cancellable) (int, error) {
	var cerr *glib.Error

	cret := xOutputStreamSplice(x.GoPointer(), SourceVar.GoPointer(), FlagsVar, CancellableVar.GoPointer(), &cerr)
	if cerr == nil {
		return cret, nil
	}
	return cret, cerr

}

var xOutputStreamSpliceAsync func(uintptr, uintptr, OutputStreamSpliceFlags, int, uintptr, uintptr, uintptr)

// Splices a stream asynchronously.
// When the operation is finished @callback will be called.
// You can then call g_output_stream_splice_finish() to get the
// result of the operation.
//
// For the synchronous, blocking version of this function, see
// g_output_stream_splice().
func (x *OutputStream) SpliceAsync(SourceVar *InputStream, FlagsVar OutputStreamSpliceFlags, IoPriorityVar int, CancellableVar *Cancellable, CallbackVar AsyncReadyCallback, UserDataVar uintptr) {

	xOutputStreamSpliceAsync(x.GoPointer(), SourceVar.GoPointer(), FlagsVar, IoPriorityVar, CancellableVar.GoPointer(), purego.NewCallback(CallbackVar), UserDataVar)

}

var xOutputStreamSpliceFinish func(uintptr, uintptr, **glib.Error) int

// Finishes an asynchronous stream splice operation.
func (x *OutputStream) SpliceFinish(ResultVar AsyncResult) (int, error) {
	var cerr *glib.Error

	cret := xOutputStreamSpliceFinish(x.GoPointer(), ResultVar.GoPointer(), &cerr)
	if cerr == nil {
		return cret, nil
	}
	return cret, cerr

}

var xOutputStreamVprintf func(uintptr, uint, uintptr, **glib.Error, string, []interface{}) bool

// This is a utility function around g_output_stream_write_all(). It
// uses g_strdup_vprintf() to turn @format and @args into a string that
// is then written to @stream.
//
// See the documentation of g_output_stream_write_all() about the
// behavior of the actual write operation.
//
// Note that partial writes cannot be properly checked with this
// function due to the variable length of the written string, if you
// need precise control over partial write failures, you need to
// create you own printf()-like wrapper around g_output_stream_write()
// or g_output_stream_write_all().
func (x *OutputStream) Vprintf(BytesWrittenVar uint, CancellableVar *Cancellable, ErrorVar **glib.Error, FormatVar string, ArgsVar []interface{}) bool {

	cret := xOutputStreamVprintf(x.GoPointer(), BytesWrittenVar, CancellableVar.GoPointer(), ErrorVar, FormatVar, ArgsVar)
	return cret
}

var xOutputStreamWrite func(uintptr, uintptr, uint, uintptr, **glib.Error) int

// Tries to write @count bytes from @buffer into the stream. Will block
// during the operation.
//
// If count is 0, returns 0 and does nothing. A value of @count
// larger than %G_MAXSSIZE will cause a %G_IO_ERROR_INVALID_ARGUMENT error.
//
// On success, the number of bytes written to the stream is returned.
// It is not an error if this is not the same as the requested size, as it
// can happen e.g. on a partial I/O error, or if there is not enough
// storage in the stream. All writes block until at least one byte
// is written or an error occurs; 0 is never returned (unless
// @count is 0).
//
// If @cancellable is not %NULL, then the operation can be cancelled by
// triggering the cancellable object from another thread. If the operation
// was cancelled, the error %G_IO_ERROR_CANCELLED will be returned. If an
// operation was partially finished when the operation was cancelled the
// partial result will be returned, without an error.
//
// On error -1 is returned and @error is set accordingly.
func (x *OutputStream) Write(BufferVar uintptr, CountVar uint, CancellableVar *Cancellable) (int, error) {
	var cerr *glib.Error

	cret := xOutputStreamWrite(x.GoPointer(), BufferVar, CountVar, CancellableVar.GoPointer(), &cerr)
	if cerr == nil {
		return cret, nil
	}
	return cret, cerr

}

var xOutputStreamWriteAll func(uintptr, uintptr, uint, uint, uintptr, **glib.Error) bool

// Tries to write @count bytes from @buffer into the stream. Will block
// during the operation.
//
// This function is similar to g_output_stream_write(), except it tries to
// write as many bytes as requested, only stopping on an error.
//
// On a successful write of @count bytes, %TRUE is returned, and @bytes_written
// is set to @count.
//
// If there is an error during the operation %FALSE is returned and @error
// is set to indicate the error status.
//
// As a special exception to the normal conventions for functions that
// use #GError, if this function returns %FALSE (and sets @error) then
// @bytes_written will be set to the number of bytes that were
// successfully written before the error was encountered.  This
// functionality is only available from C.  If you need it from another
// language then you must write your own loop around
// g_output_stream_write().
func (x *OutputStream) WriteAll(BufferVar uintptr, CountVar uint, BytesWrittenVar uint, CancellableVar *Cancellable) (bool, error) {
	var cerr *glib.Error

	cret := xOutputStreamWriteAll(x.GoPointer(), BufferVar, CountVar, BytesWrittenVar, CancellableVar.GoPointer(), &cerr)
	if cerr == nil {
		return cret, nil
	}
	return cret, cerr

}

var xOutputStreamWriteAllAsync func(uintptr, uintptr, uint, int, uintptr, uintptr, uintptr)

// Request an asynchronous write of @count bytes from @buffer into
// the stream. When the operation is finished @callback will be called.
// You can then call g_output_stream_write_all_finish() to get the result of the
// operation.
//
// This is the asynchronous version of g_output_stream_write_all().
//
// Call g_output_stream_write_all_finish() to collect the result.
//
// Any outstanding I/O request with higher priority (lower numerical
// value) will be executed before an outstanding request with lower
// priority. Default priority is %G_PRIORITY_DEFAULT.
//
// Note that no copy of @buffer will be made, so it must stay valid
// until @callback is called.
func (x *OutputStream) WriteAllAsync(BufferVar uintptr, CountVar uint, IoPriorityVar int, CancellableVar *Cancellable, CallbackVar AsyncReadyCallback, UserDataVar uintptr) {

	xOutputStreamWriteAllAsync(x.GoPointer(), BufferVar, CountVar, IoPriorityVar, CancellableVar.GoPointer(), purego.NewCallback(CallbackVar), UserDataVar)

}

var xOutputStreamWriteAllFinish func(uintptr, uintptr, uint, **glib.Error) bool

// Finishes an asynchronous stream write operation started with
// g_output_stream_write_all_async().
//
// As a special exception to the normal conventions for functions that
// use #GError, if this function returns %FALSE (and sets @error) then
// @bytes_written will be set to the number of bytes that were
// successfully written before the error was encountered.  This
// functionality is only available from C.  If you need it from another
// language then you must write your own loop around
// g_output_stream_write_async().
func (x *OutputStream) WriteAllFinish(ResultVar AsyncResult, BytesWrittenVar uint) (bool, error) {
	var cerr *glib.Error

	cret := xOutputStreamWriteAllFinish(x.GoPointer(), ResultVar.GoPointer(), BytesWrittenVar, &cerr)
	if cerr == nil {
		return cret, nil
	}
	return cret, cerr

}

var xOutputStreamWriteAsync func(uintptr, uintptr, uint, int, uintptr, uintptr, uintptr)

// Request an asynchronous write of @count bytes from @buffer into
// the stream. When the operation is finished @callback will be called.
// You can then call g_output_stream_write_finish() to get the result of the
// operation.
//
// During an async request no other sync and async calls are allowed,
// and will result in %G_IO_ERROR_PENDING errors.
//
// A value of @count larger than %G_MAXSSIZE will cause a
// %G_IO_ERROR_INVALID_ARGUMENT error.
//
// On success, the number of bytes written will be passed to the
// @callback. It is not an error if this is not the same as the
// requested size, as it can happen e.g. on a partial I/O error,
// but generally we try to write as many bytes as requested.
//
// You are guaranteed that this method will never fail with
// %G_IO_ERROR_WOULD_BLOCK - if @stream can't accept more data, the
// method will just wait until this changes.
//
// Any outstanding I/O request with higher priority (lower numerical
// value) will be executed before an outstanding request with lower
// priority. Default priority is %G_PRIORITY_DEFAULT.
//
// The asynchronous methods have a default fallback that uses threads
// to implement asynchronicity, so they are optional for inheriting
// classes. However, if you override one you must override all.
//
// For the synchronous, blocking version of this function, see
// g_output_stream_write().
//
// Note that no copy of @buffer will be made, so it must stay valid
// until @callback is called. See g_output_stream_write_bytes_async()
// for a #GBytes version that will automatically hold a reference to
// the contents (without copying) for the duration of the call.
func (x *OutputStream) WriteAsync(BufferVar uintptr, CountVar uint, IoPriorityVar int, CancellableVar *Cancellable, CallbackVar AsyncReadyCallback, UserDataVar uintptr) {

	xOutputStreamWriteAsync(x.GoPointer(), BufferVar, CountVar, IoPriorityVar, CancellableVar.GoPointer(), purego.NewCallback(CallbackVar), UserDataVar)

}

var xOutputStreamWriteBytes func(uintptr, *glib.Bytes, uintptr, **glib.Error) int

// A wrapper function for g_output_stream_write() which takes a
// #GBytes as input.  This can be more convenient for use by language
// bindings or in other cases where the refcounted nature of #GBytes
// is helpful over a bare pointer interface.
//
// However, note that this function may still perform partial writes,
// just like g_output_stream_write().  If that occurs, to continue
// writing, you will need to create a new #GBytes containing just the
// remaining bytes, using g_bytes_new_from_bytes(). Passing the same
// #GBytes instance multiple times potentially can result in duplicated
// data in the output stream.
func (x *OutputStream) WriteBytes(BytesVar *glib.Bytes, CancellableVar *Cancellable) (int, error) {
	var cerr *glib.Error

	cret := xOutputStreamWriteBytes(x.GoPointer(), BytesVar, CancellableVar.GoPointer(), &cerr)
	if cerr == nil {
		return cret, nil
	}
	return cret, cerr

}

var xOutputStreamWriteBytesAsync func(uintptr, *glib.Bytes, int, uintptr, uintptr, uintptr)

// This function is similar to g_output_stream_write_async(), but
// takes a #GBytes as input.  Due to the refcounted nature of #GBytes,
// this allows the stream to avoid taking a copy of the data.
//
// However, note that this function may still perform partial writes,
// just like g_output_stream_write_async(). If that occurs, to continue
// writing, you will need to create a new #GBytes containing just the
// remaining bytes, using g_bytes_new_from_bytes(). Passing the same
// #GBytes instance multiple times potentially can result in duplicated
// data in the output stream.
//
// For the synchronous, blocking version of this function, see
// g_output_stream_write_bytes().
func (x *OutputStream) WriteBytesAsync(BytesVar *glib.Bytes, IoPriorityVar int, CancellableVar *Cancellable, CallbackVar AsyncReadyCallback, UserDataVar uintptr) {

	xOutputStreamWriteBytesAsync(x.GoPointer(), BytesVar, IoPriorityVar, CancellableVar.GoPointer(), purego.NewCallback(CallbackVar), UserDataVar)

}

var xOutputStreamWriteBytesFinish func(uintptr, uintptr, **glib.Error) int

// Finishes a stream write-from-#GBytes operation.
func (x *OutputStream) WriteBytesFinish(ResultVar AsyncResult) (int, error) {
	var cerr *glib.Error

	cret := xOutputStreamWriteBytesFinish(x.GoPointer(), ResultVar.GoPointer(), &cerr)
	if cerr == nil {
		return cret, nil
	}
	return cret, cerr

}

var xOutputStreamWriteFinish func(uintptr, uintptr, **glib.Error) int

// Finishes a stream write operation.
func (x *OutputStream) WriteFinish(ResultVar AsyncResult) (int, error) {
	var cerr *glib.Error

	cret := xOutputStreamWriteFinish(x.GoPointer(), ResultVar.GoPointer(), &cerr)
	if cerr == nil {
		return cret, nil
	}
	return cret, cerr

}

var xOutputStreamWritev func(uintptr, uintptr, uint, uint, uintptr, **glib.Error) bool

// Tries to write the bytes contained in the @n_vectors @vectors into the
// stream. Will block during the operation.
//
// If @n_vectors is 0 or the sum of all bytes in @vectors is 0, returns 0 and
// does nothing.
//
// On success, the number of bytes written to the stream is returned.
// It is not an error if this is not the same as the requested size, as it
// can happen e.g. on a partial I/O error, or if there is not enough
// storage in the stream. All writes block until at least one byte
// is written or an error occurs; 0 is never returned (unless
// @n_vectors is 0 or the sum of all bytes in @vectors is 0).
//
// If @cancellable is not %NULL, then the operation can be cancelled by
// triggering the cancellable object from another thread. If the operation
// was cancelled, the error %G_IO_ERROR_CANCELLED will be returned. If an
// operation was partially finished when the operation was cancelled the
// partial result will be returned, without an error.
//
// Some implementations of g_output_stream_writev() may have limitations on the
// aggregate buffer size, and will return %G_IO_ERROR_INVALID_ARGUMENT if these
// are exceeded. For example, when writing to a local file on UNIX platforms,
// the aggregate buffer size must not exceed %G_MAXSSIZE bytes.
func (x *OutputStream) Writev(VectorsVar uintptr, NVectorsVar uint, BytesWrittenVar uint, CancellableVar *Cancellable) (bool, error) {
	var cerr *glib.Error

	cret := xOutputStreamWritev(x.GoPointer(), VectorsVar, NVectorsVar, BytesWrittenVar, CancellableVar.GoPointer(), &cerr)
	if cerr == nil {
		return cret, nil
	}
	return cret, cerr

}

var xOutputStreamWritevAll func(uintptr, uintptr, uint, uint, uintptr, **glib.Error) bool

// Tries to write the bytes contained in the @n_vectors @vectors into the
// stream. Will block during the operation.
//
// This function is similar to g_output_stream_writev(), except it tries to
// write as many bytes as requested, only stopping on an error.
//
// On a successful write of all @n_vectors vectors, %TRUE is returned, and
// @bytes_written is set to the sum of all the sizes of @vectors.
//
// If there is an error during the operation %FALSE is returned and @error
// is set to indicate the error status.
//
// As a special exception to the normal conventions for functions that
// use #GError, if this function returns %FALSE (and sets @error) then
// @bytes_written will be set to the number of bytes that were
// successfully written before the error was encountered.  This
// functionality is only available from C. If you need it from another
// language then you must write your own loop around
// g_output_stream_write().
//
// The content of the individual elements of @vectors might be changed by this
// function.
func (x *OutputStream) WritevAll(VectorsVar uintptr, NVectorsVar uint, BytesWrittenVar uint, CancellableVar *Cancellable) (bool, error) {
	var cerr *glib.Error

	cret := xOutputStreamWritevAll(x.GoPointer(), VectorsVar, NVectorsVar, BytesWrittenVar, CancellableVar.GoPointer(), &cerr)
	if cerr == nil {
		return cret, nil
	}
	return cret, cerr

}

var xOutputStreamWritevAllAsync func(uintptr, uintptr, uint, int, uintptr, uintptr, uintptr)

// Request an asynchronous write of the bytes contained in the @n_vectors @vectors into
// the stream. When the operation is finished @callback will be called.
// You can then call g_output_stream_writev_all_finish() to get the result of the
// operation.
//
// This is the asynchronous version of g_output_stream_writev_all().
//
// Call g_output_stream_writev_all_finish() to collect the result.
//
// Any outstanding I/O request with higher priority (lower numerical
// value) will be executed before an outstanding request with lower
// priority. Default priority is %G_PRIORITY_DEFAULT.
//
// Note that no copy of @vectors will be made, so it must stay valid
// until @callback is called. The content of the individual elements
// of @vectors might be changed by this function.
func (x *OutputStream) WritevAllAsync(VectorsVar uintptr, NVectorsVar uint, IoPriorityVar int, CancellableVar *Cancellable, CallbackVar AsyncReadyCallback, UserDataVar uintptr) {

	xOutputStreamWritevAllAsync(x.GoPointer(), VectorsVar, NVectorsVar, IoPriorityVar, CancellableVar.GoPointer(), purego.NewCallback(CallbackVar), UserDataVar)

}

var xOutputStreamWritevAllFinish func(uintptr, uintptr, uint, **glib.Error) bool

// Finishes an asynchronous stream write operation started with
// g_output_stream_writev_all_async().
//
// As a special exception to the normal conventions for functions that
// use #GError, if this function returns %FALSE (and sets @error) then
// @bytes_written will be set to the number of bytes that were
// successfully written before the error was encountered.  This
// functionality is only available from C.  If you need it from another
// language then you must write your own loop around
// g_output_stream_writev_async().
func (x *OutputStream) WritevAllFinish(ResultVar AsyncResult, BytesWrittenVar uint) (bool, error) {
	var cerr *glib.Error

	cret := xOutputStreamWritevAllFinish(x.GoPointer(), ResultVar.GoPointer(), BytesWrittenVar, &cerr)
	if cerr == nil {
		return cret, nil
	}
	return cret, cerr

}

var xOutputStreamWritevAsync func(uintptr, uintptr, uint, int, uintptr, uintptr, uintptr)

// Request an asynchronous write of the bytes contained in @n_vectors @vectors into
// the stream. When the operation is finished @callback will be called.
// You can then call g_output_stream_writev_finish() to get the result of the
// operation.
//
// During an async request no other sync and async calls are allowed,
// and will result in %G_IO_ERROR_PENDING errors.
//
// On success, the number of bytes written will be passed to the
// @callback. It is not an error if this is not the same as the
// requested size, as it can happen e.g. on a partial I/O error,
// but generally we try to write as many bytes as requested.
//
// You are guaranteed that this method will never fail with
// %G_IO_ERROR_WOULD_BLOCK — if @stream can't accept more data, the
// method will just wait until this changes.
//
// Any outstanding I/O request with higher priority (lower numerical
// value) will be executed before an outstanding request with lower
// priority. Default priority is %G_PRIORITY_DEFAULT.
//
// The asynchronous methods have a default fallback that uses threads
// to implement asynchronicity, so they are optional for inheriting
// classes. However, if you override one you must override all.
//
// For the synchronous, blocking version of this function, see
// g_output_stream_writev().
//
// Note that no copy of @vectors will be made, so it must stay valid
// until @callback is called.
func (x *OutputStream) WritevAsync(VectorsVar uintptr, NVectorsVar uint, IoPriorityVar int, CancellableVar *Cancellable, CallbackVar AsyncReadyCallback, UserDataVar uintptr) {

	xOutputStreamWritevAsync(x.GoPointer(), VectorsVar, NVectorsVar, IoPriorityVar, CancellableVar.GoPointer(), purego.NewCallback(CallbackVar), UserDataVar)

}

var xOutputStreamWritevFinish func(uintptr, uintptr, uint, **glib.Error) bool

// Finishes a stream writev operation.
func (x *OutputStream) WritevFinish(ResultVar AsyncResult, BytesWrittenVar uint) (bool, error) {
	var cerr *glib.Error

	cret := xOutputStreamWritevFinish(x.GoPointer(), ResultVar.GoPointer(), BytesWrittenVar, &cerr)
	if cerr == nil {
		return cret, nil
	}
	return cret, cerr

}

func (c *OutputStream) GoPointer() uintptr {
	return c.Ptr
}

func (c *OutputStream) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GIO"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xOutputStreamClearPending, lib, "g_output_stream_clear_pending")
	core.PuregoSafeRegister(&xOutputStreamClose, lib, "g_output_stream_close")
	core.PuregoSafeRegister(&xOutputStreamCloseAsync, lib, "g_output_stream_close_async")
	core.PuregoSafeRegister(&xOutputStreamCloseFinish, lib, "g_output_stream_close_finish")
	core.PuregoSafeRegister(&xOutputStreamFlush, lib, "g_output_stream_flush")
	core.PuregoSafeRegister(&xOutputStreamFlushAsync, lib, "g_output_stream_flush_async")
	core.PuregoSafeRegister(&xOutputStreamFlushFinish, lib, "g_output_stream_flush_finish")
	core.PuregoSafeRegister(&xOutputStreamHasPending, lib, "g_output_stream_has_pending")
	core.PuregoSafeRegister(&xOutputStreamIsClosed, lib, "g_output_stream_is_closed")
	core.PuregoSafeRegister(&xOutputStreamIsClosing, lib, "g_output_stream_is_closing")
	core.PuregoSafeRegister(&xOutputStreamPrintf, lib, "g_output_stream_printf")
	core.PuregoSafeRegister(&xOutputStreamSetPending, lib, "g_output_stream_set_pending")
	core.PuregoSafeRegister(&xOutputStreamSplice, lib, "g_output_stream_splice")
	core.PuregoSafeRegister(&xOutputStreamSpliceAsync, lib, "g_output_stream_splice_async")
	core.PuregoSafeRegister(&xOutputStreamSpliceFinish, lib, "g_output_stream_splice_finish")
	core.PuregoSafeRegister(&xOutputStreamVprintf, lib, "g_output_stream_vprintf")
	core.PuregoSafeRegister(&xOutputStreamWrite, lib, "g_output_stream_write")
	core.PuregoSafeRegister(&xOutputStreamWriteAll, lib, "g_output_stream_write_all")
	core.PuregoSafeRegister(&xOutputStreamWriteAllAsync, lib, "g_output_stream_write_all_async")
	core.PuregoSafeRegister(&xOutputStreamWriteAllFinish, lib, "g_output_stream_write_all_finish")
	core.PuregoSafeRegister(&xOutputStreamWriteAsync, lib, "g_output_stream_write_async")
	core.PuregoSafeRegister(&xOutputStreamWriteBytes, lib, "g_output_stream_write_bytes")
	core.PuregoSafeRegister(&xOutputStreamWriteBytesAsync, lib, "g_output_stream_write_bytes_async")
	core.PuregoSafeRegister(&xOutputStreamWriteBytesFinish, lib, "g_output_stream_write_bytes_finish")
	core.PuregoSafeRegister(&xOutputStreamWriteFinish, lib, "g_output_stream_write_finish")
	core.PuregoSafeRegister(&xOutputStreamWritev, lib, "g_output_stream_writev")
	core.PuregoSafeRegister(&xOutputStreamWritevAll, lib, "g_output_stream_writev_all")
	core.PuregoSafeRegister(&xOutputStreamWritevAllAsync, lib, "g_output_stream_writev_all_async")
	core.PuregoSafeRegister(&xOutputStreamWritevAllFinish, lib, "g_output_stream_writev_all_finish")
	core.PuregoSafeRegister(&xOutputStreamWritevAsync, lib, "g_output_stream_writev_async")
	core.PuregoSafeRegister(&xOutputStreamWritevFinish, lib, "g_output_stream_writev_finish")

}
