// Package gio was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gio

import (
	"unsafe"

	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/gobject"
)

type ThreadedSocketServiceClass struct {
	ParentClass uintptr
}

func (x *ThreadedSocketServiceClass) GoPointer() uintptr {
	return uintptr(unsafe.Pointer(x))
}

type ThreadedSocketServicePrivate struct {
}

func (x *ThreadedSocketServicePrivate) GoPointer() uintptr {
	return uintptr(unsafe.Pointer(x))
}

// A #GThreadedSocketService is a simple subclass of #GSocketService
// that handles incoming connections by creating a worker thread and
// dispatching the connection to it by emitting the
// #GThreadedSocketService::run signal in the new thread.
//
// The signal handler may perform blocking IO and need not return
// until the connection is closed.
//
// The service is implemented using a thread pool, so there is a
// limited amount of threads available to serve incoming requests.
// The service automatically stops the #GSocketService from accepting
// new connections when all threads are busy.
//
// As with #GSocketService, you may connect to #GThreadedSocketService::run,
// or subclass and override the default handler.
type ThreadedSocketService struct {
	SocketService
}

func ThreadedSocketServiceNewFromInternalPtr(ptr uintptr) *ThreadedSocketService {
	cls := &ThreadedSocketService{}
	cls.Ptr = ptr
	return cls
}

var xNewThreadedSocketService func(int) uintptr

// Creates a new #GThreadedSocketService with no listeners. Listeners
// must be added with one of the #GSocketListener "add" methods.
func NewThreadedSocketService(MaxThreadsVar int) *SocketService {
	var cls *SocketService

	cret := xNewThreadedSocketService(MaxThreadsVar)

	if cret == 0 {
		return nil
	}
	cls = &SocketService{}
	cls.Ptr = cret
	return cls
}

func (c *ThreadedSocketService) GoPointer() uintptr {
	return c.Ptr
}

func (c *ThreadedSocketService) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// The ::run signal is emitted in a worker thread in response to an
// incoming connection. This thread is dedicated to handling
// @connection and may perform blocking IO. The signal handler need
// not return until the connection is closed.
func (x *ThreadedSocketService) ConnectRun(cb func(ThreadedSocketService, uintptr, uintptr) bool) uint32 {
	fcb := func(clsPtr uintptr, ConnectionVarp uintptr, SourceObjectVarp uintptr) bool {
		fa := ThreadedSocketService{}
		fa.Ptr = clsPtr

		return cb(fa, ConnectionVarp, SourceObjectVarp)

	}
	return gobject.SignalConnect(x.GoPointer(), "run", purego.NewCallback(fcb))
}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GIO"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xNewThreadedSocketService, lib, "g_threaded_socket_service_new")

}
