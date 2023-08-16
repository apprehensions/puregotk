// Package gio was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gio

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/glib"
	"github.com/jwijenbergh/puregotk/v4/gobject"
)

type SocketConnectionClass struct {
	ParentClass uintptr
}

type SocketConnectionPrivate struct {
}

// #GSocketConnection is a #GIOStream for a connected socket. They
// can be created either by #GSocketClient when connecting to a host,
// or by #GSocketListener when accepting a new client.
//
// The type of the #GSocketConnection object returned from these calls
// depends on the type of the underlying socket that is in use. For
// instance, for a TCP/IP connection it will be a #GTcpConnection.
//
// Choosing what type of object to construct is done with the socket
// connection factory, and it is possible for 3rd parties to register
// custom socket connection types for specific combination of socket
// family/type/protocol using g_socket_connection_factory_register_type().
//
// To close a #GSocketConnection, use g_io_stream_close(). Closing both
// substreams of the #GIOStream separately will not close the underlying
// #GSocket.
type SocketConnection struct {
	IOStream
}

func SocketConnectionNewFromInternalPtr(ptr uintptr) *SocketConnection {
	cls := &SocketConnection{}
	cls.Ptr = ptr
	return cls
}

var xSocketConnectionConnect func(uintptr, uintptr, uintptr, **glib.Error) bool

// Connect @connection to the specified remote address.
func (x *SocketConnection) Connect(AddressVar *SocketAddress, CancellableVar *Cancellable) (bool, error) {
	var cerr *glib.Error

	cret := xSocketConnectionConnect(x.GoPointer(), AddressVar.GoPointer(), CancellableVar.GoPointer(), &cerr)
	if cerr == nil {
		return cret, nil
	}
	return cret, cerr

}

var xSocketConnectionConnectAsync func(uintptr, uintptr, uintptr, uintptr, uintptr)

// Asynchronously connect @connection to the specified remote address.
//
// This clears the #GSocket:blocking flag on @connection's underlying
// socket if it is currently set.
//
// Use g_socket_connection_connect_finish() to retrieve the result.
func (x *SocketConnection) ConnectAsync(AddressVar *SocketAddress, CancellableVar *Cancellable, CallbackVar AsyncReadyCallback, UserDataVar uintptr) {

	xSocketConnectionConnectAsync(x.GoPointer(), AddressVar.GoPointer(), CancellableVar.GoPointer(), purego.NewCallback(CallbackVar), UserDataVar)

}

var xSocketConnectionConnectFinish func(uintptr, uintptr, **glib.Error) bool

// Gets the result of a g_socket_connection_connect_async() call.
func (x *SocketConnection) ConnectFinish(ResultVar AsyncResult) (bool, error) {
	var cerr *glib.Error

	cret := xSocketConnectionConnectFinish(x.GoPointer(), ResultVar.GoPointer(), &cerr)
	if cerr == nil {
		return cret, nil
	}
	return cret, cerr

}

var xSocketConnectionGetLocalAddress func(uintptr) uintptr

// Try to get the local address of a socket connection.
func (x *SocketConnection) GetLocalAddress() (*SocketAddress, error) {
	var cls *SocketAddress
	var cerr *glib.Error

	cret := xSocketConnectionGetLocalAddress(x.GoPointer())

	if cret == 0 {
		return nil, cerr
	}
	cls = &SocketAddress{}
	cls.Ptr = cret
	if cerr == nil {
		return cls, nil
	}
	return cls, cerr

}

var xSocketConnectionGetRemoteAddress func(uintptr) uintptr

// Try to get the remote address of a socket connection.
//
// Since GLib 2.40, when used with g_socket_client_connect() or
// g_socket_client_connect_async(), during emission of
// %G_SOCKET_CLIENT_CONNECTING, this function will return the remote
// address that will be used for the connection.  This allows
// applications to print e.g. "Connecting to example.com
// (10.42.77.3)...".
func (x *SocketConnection) GetRemoteAddress() (*SocketAddress, error) {
	var cls *SocketAddress
	var cerr *glib.Error

	cret := xSocketConnectionGetRemoteAddress(x.GoPointer())

	if cret == 0 {
		return nil, cerr
	}
	cls = &SocketAddress{}
	cls.Ptr = cret
	if cerr == nil {
		return cls, nil
	}
	return cls, cerr

}

var xSocketConnectionGetSocket func(uintptr) uintptr

// Gets the underlying #GSocket object of the connection.
// This can be useful if you want to do something unusual on it
// not supported by the #GSocketConnection APIs.
func (x *SocketConnection) GetSocket() *Socket {
	var cls *Socket

	cret := xSocketConnectionGetSocket(x.GoPointer())

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &Socket{}
	cls.Ptr = cret
	return cls
}

var xSocketConnectionIsConnected func(uintptr) bool

// Checks if @connection is connected. This is equivalent to calling
// g_socket_is_connected() on @connection's underlying #GSocket.
func (x *SocketConnection) IsConnected() bool {

	cret := xSocketConnectionIsConnected(x.GoPointer())
	return cret
}

func (c *SocketConnection) GoPointer() uintptr {
	return c.Ptr
}

func (c *SocketConnection) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GIO"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xSocketConnectionConnect, lib, "g_socket_connection_connect")
	core.PuregoSafeRegister(&xSocketConnectionConnectAsync, lib, "g_socket_connection_connect_async")
	core.PuregoSafeRegister(&xSocketConnectionConnectFinish, lib, "g_socket_connection_connect_finish")
	core.PuregoSafeRegister(&xSocketConnectionGetLocalAddress, lib, "g_socket_connection_get_local_address")
	core.PuregoSafeRegister(&xSocketConnectionGetRemoteAddress, lib, "g_socket_connection_get_remote_address")
	core.PuregoSafeRegister(&xSocketConnectionGetSocket, lib, "g_socket_connection_get_socket")
	core.PuregoSafeRegister(&xSocketConnectionIsConnected, lib, "g_socket_connection_is_connected")

}
