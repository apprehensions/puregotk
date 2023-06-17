// Package gio was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gio

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
)

type UnixSocketAddressClass struct {
	ParentClass uintptr
}

type UnixSocketAddressPrivate struct {
}

// Support for UNIX-domain (also known as local) sockets.
//
// UNIX domain sockets are generally visible in the filesystem.
// However, some systems support abstract socket names which are not
// visible in the filesystem and not affected by the filesystem
// permissions, visibility, etc. Currently this is only supported
// under Linux. If you attempt to use abstract sockets on other
// systems, function calls may return %G_IO_ERROR_NOT_SUPPORTED
// errors. You can use g_unix_socket_address_abstract_names_supported()
// to see if abstract names are supported.
//
// Since GLib 2.72, #GUnixSocketAddress is available on all platforms. It
// requires underlying system support (such as Windows 10 with `AF_UNIX`) at
// run time.
//
// Before GLib 2.72, `&lt;gio/gunixsocketaddress.h&gt;` belonged to the UNIX-specific
// GIO interfaces, thus you had to use the `gio-unix-2.0.pc` pkg-config file
// when using it. This is no longer necessary since GLib 2.72.
type UnixSocketAddress struct {
	SocketAddress
}

func UnixSocketAddressNewFromInternalPtr(ptr uintptr) *UnixSocketAddress {
	cls := &UnixSocketAddress{}
	cls.Ptr = ptr
	return cls
}

var xNewUnixSocketAddress func(string) uintptr

// Creates a new #GUnixSocketAddress for @path.
//
// To create abstract socket addresses, on systems that support that,
// use g_unix_socket_address_new_abstract().
func NewUnixSocketAddress(PathVar string) *SocketAddress {
	NewUnixSocketAddressPtr := xNewUnixSocketAddress(PathVar)
	if NewUnixSocketAddressPtr == 0 {
		return nil
	}

	NewUnixSocketAddressCls := &SocketAddress{}
	NewUnixSocketAddressCls.Ptr = NewUnixSocketAddressPtr
	return NewUnixSocketAddressCls
}

var xNewAbstractUnixSocketAddress func(uintptr, int32) uintptr

// Creates a new %G_UNIX_SOCKET_ADDRESS_ABSTRACT_PADDED
// #GUnixSocketAddress for @path.
func NewAbstractUnixSocketAddress(PathVar uintptr, PathLenVar int32) *SocketAddress {
	NewAbstractUnixSocketAddressPtr := xNewAbstractUnixSocketAddress(PathVar, PathLenVar)
	if NewAbstractUnixSocketAddressPtr == 0 {
		return nil
	}

	NewAbstractUnixSocketAddressCls := &SocketAddress{}
	NewAbstractUnixSocketAddressCls.Ptr = NewAbstractUnixSocketAddressPtr
	return NewAbstractUnixSocketAddressCls
}

var xNewWithTypeUnixSocketAddress func(uintptr, int32, UnixSocketAddressType) uintptr

// Creates a new #GUnixSocketAddress of type @type with name @path.
//
// If @type is %G_UNIX_SOCKET_ADDRESS_PATH, this is equivalent to
// calling g_unix_socket_address_new().
//
// If @type is %G_UNIX_SOCKET_ADDRESS_ANONYMOUS, @path and @path_len will be
// ignored.
//
// If @path_type is %G_UNIX_SOCKET_ADDRESS_ABSTRACT, then @path_len
// bytes of @path will be copied to the socket's path, and only those
// bytes will be considered part of the name. (If @path_len is -1,
// then @path is assumed to be NUL-terminated.) For example, if @path
// was "test", then calling g_socket_address_get_native_size() on the
// returned socket would return 7 (2 bytes of overhead, 1 byte for the
// abstract-socket indicator byte, and 4 bytes for the name "test").
//
// If @path_type is %G_UNIX_SOCKET_ADDRESS_ABSTRACT_PADDED, then
// @path_len bytes of @path will be copied to the socket's path, the
// rest of the path will be padded with 0 bytes, and the entire
// zero-padded buffer will be considered the name. (As above, if
// @path_len is -1, then @path is assumed to be NUL-terminated.) In
// this case, g_socket_address_get_native_size() will always return
// the full size of a `struct sockaddr_un`, although
// g_unix_socket_address_get_path_len() will still return just the
// length of @path.
//
// %G_UNIX_SOCKET_ADDRESS_ABSTRACT is preferred over
// %G_UNIX_SOCKET_ADDRESS_ABSTRACT_PADDED for new programs. Of course,
// when connecting to a server created by another process, you must
// use the appropriate type corresponding to how that process created
// its listening socket.
func NewWithTypeUnixSocketAddress(PathVar uintptr, PathLenVar int32, TypeVar UnixSocketAddressType) *SocketAddress {
	NewWithTypeUnixSocketAddressPtr := xNewWithTypeUnixSocketAddress(PathVar, PathLenVar, TypeVar)
	if NewWithTypeUnixSocketAddressPtr == 0 {
		return nil
	}

	NewWithTypeUnixSocketAddressCls := &SocketAddress{}
	NewWithTypeUnixSocketAddressCls.Ptr = NewWithTypeUnixSocketAddressPtr
	return NewWithTypeUnixSocketAddressCls
}

var xUnixSocketAddressGetAddressType func(uintptr) UnixSocketAddressType

// Gets @address's type.
func (x *UnixSocketAddress) GetAddressType() UnixSocketAddressType {

	return xUnixSocketAddressGetAddressType(x.GoPointer())

}

var xUnixSocketAddressGetIsAbstract func(uintptr) bool

// Tests if @address is abstract.
func (x *UnixSocketAddress) GetIsAbstract() bool {

	return xUnixSocketAddressGetIsAbstract(x.GoPointer())

}

var xUnixSocketAddressGetPath func(uintptr) string

// Gets @address's path, or for abstract sockets the "name".
//
// Guaranteed to be zero-terminated, but an abstract socket
// may contain embedded zeros, and thus you should use
// g_unix_socket_address_get_path_len() to get the true length
// of this string.
func (x *UnixSocketAddress) GetPath() string {

	return xUnixSocketAddressGetPath(x.GoPointer())

}

var xUnixSocketAddressGetPathLen func(uintptr) uint

// Gets the length of @address's path.
//
// For details, see g_unix_socket_address_get_path().
func (x *UnixSocketAddress) GetPathLen() uint {

	return xUnixSocketAddressGetPathLen(x.GoPointer())

}

func (c *UnixSocketAddress) GoPointer() uintptr {
	return c.Ptr
}

func (c *UnixSocketAddress) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// Creates a #GSocketAddressEnumerator for @connectable.
func (x *UnixSocketAddress) Enumerate() *SocketAddressEnumerator {

	EnumeratePtr := XGSocketConnectableEnumerate(x.GoPointer())
	if EnumeratePtr == 0 {
		return nil
	}

	EnumerateCls := &SocketAddressEnumerator{}
	EnumerateCls.Ptr = EnumeratePtr
	return EnumerateCls

}

// Creates a #GSocketAddressEnumerator for @connectable that will
// return a #GProxyAddress for each of its addresses that you must connect
// to via a proxy.
//
// If @connectable does not implement
// g_socket_connectable_proxy_enumerate(), this will fall back to
// calling g_socket_connectable_enumerate().
func (x *UnixSocketAddress) ProxyEnumerate() *SocketAddressEnumerator {

	ProxyEnumeratePtr := XGSocketConnectableProxyEnumerate(x.GoPointer())
	if ProxyEnumeratePtr == 0 {
		return nil
	}

	ProxyEnumerateCls := &SocketAddressEnumerator{}
	ProxyEnumerateCls.Ptr = ProxyEnumeratePtr
	return ProxyEnumerateCls

}

// Format a #GSocketConnectable as a string. This is a human-readable format for
// use in debugging output, and is not a stable serialization format. It is not
// suitable for use in user interfaces as it exposes too much information for a
// user.
//
// If the #GSocketConnectable implementation does not support string formatting,
// the implementation’s type name will be returned as a fallback.
func (x *UnixSocketAddress) ToString() string {

	return XGSocketConnectableToString(x.GoPointer())

}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GIO"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xNewUnixSocketAddress, lib, "g_unix_socket_address_new")
	core.PuregoSafeRegister(&xNewAbstractUnixSocketAddress, lib, "g_unix_socket_address_new_abstract")
	core.PuregoSafeRegister(&xNewWithTypeUnixSocketAddress, lib, "g_unix_socket_address_new_with_type")

	core.PuregoSafeRegister(&xUnixSocketAddressGetAddressType, lib, "g_unix_socket_address_get_address_type")
	core.PuregoSafeRegister(&xUnixSocketAddressGetIsAbstract, lib, "g_unix_socket_address_get_is_abstract")
	core.PuregoSafeRegister(&xUnixSocketAddressGetPath, lib, "g_unix_socket_address_get_path")
	core.PuregoSafeRegister(&xUnixSocketAddressGetPathLen, lib, "g_unix_socket_address_get_path_len")

}
