// Package gio was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gio

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/gobject"
)

type TcpWrapperConnectionClass struct {
	ParentClass uintptr
}

type TcpWrapperConnectionPrivate struct {
}

// A #GTcpWrapperConnection can be used to wrap a #GIOStream that is
// based on a #GSocket, but which is not actually a
// #GSocketConnection. This is used by #GSocketClient so that it can
// always return a #GSocketConnection, even when the connection it has
// actually created is not directly a #GSocketConnection.
type TcpWrapperConnection struct {
	TcpConnection
}

func TcpWrapperConnectionNewFromInternalPtr(ptr uintptr) *TcpWrapperConnection {
	cls := &TcpWrapperConnection{}
	cls.Ptr = ptr
	return cls
}

var xNewTcpWrapperConnection func(uintptr, uintptr) uintptr

// Wraps @base_io_stream and @socket together as a #GSocketConnection.
func NewTcpWrapperConnection(BaseIoStreamVar *IOStream, SocketVar *Socket) *SocketConnection {
	var cls *SocketConnection

	cret := xNewTcpWrapperConnection(BaseIoStreamVar.GoPointer(), SocketVar.GoPointer())

	if cret == 0 {
		return nil
	}
	cls = &SocketConnection{}
	cls.Ptr = cret
	return cls
}

var xTcpWrapperConnectionGetBaseIoStream func(uintptr) uintptr

// Gets @conn's base #GIOStream
func (x *TcpWrapperConnection) GetBaseIoStream() *IOStream {
	var cls *IOStream

	cret := xTcpWrapperConnectionGetBaseIoStream(x.GoPointer())

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &IOStream{}
	cls.Ptr = cret
	return cls
}

func (c *TcpWrapperConnection) GoPointer() uintptr {
	return c.Ptr
}

func (c *TcpWrapperConnection) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GIO"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xNewTcpWrapperConnection, lib, "g_tcp_wrapper_connection_new")

	core.PuregoSafeRegister(&xTcpWrapperConnectionGetBaseIoStream, lib, "g_tcp_wrapper_connection_get_base_io_stream")

}
