// Package gio was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gio

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
)

// Class structure for #GProxyAddress.
type ProxyAddressClass struct {
	ParentClass uintptr
}

type ProxyAddressPrivate struct {
}

// Support for proxied #GInetSocketAddress.
type ProxyAddress struct {
	InetSocketAddress
}

func ProxyAddressNewFromInternalPtr(ptr uintptr) *ProxyAddress {
	cls := &ProxyAddress{}
	cls.Ptr = ptr
	return cls
}

var xNewProxyAddress func(uintptr, uint16, string, string, uint16, string, string) uintptr

// Creates a new #GProxyAddress for @inetaddr with @protocol that should
// tunnel through @dest_hostname and @dest_port.
//
// (Note that this method doesn't set the #GProxyAddress:uri or
// #GProxyAddress:destination-protocol fields; use g_object_new()
// directly if you want to set those.)
func NewProxyAddress(InetaddrVar *InetAddress, PortVar uint16, ProtocolVar string, DestHostnameVar string, DestPortVar uint16, UsernameVar string, PasswordVar string) *SocketAddress {
	var cls *SocketAddress

	cret := xNewProxyAddress(InetaddrVar.GoPointer(), PortVar, ProtocolVar, DestHostnameVar, DestPortVar, UsernameVar, PasswordVar)

	if cret == 0 {
		return cls
	}
	cls = &SocketAddress{}
	cls.Ptr = cret
	return cls
}

var xProxyAddressGetDestinationHostname func(uintptr) string

// Gets @proxy's destination hostname; that is, the name of the host
// that will be connected to via the proxy, not the name of the proxy
// itself.
func (x *ProxyAddress) GetDestinationHostname() string {

	cret := xProxyAddressGetDestinationHostname(x.GoPointer())
	return cret
}

var xProxyAddressGetDestinationPort func(uintptr) uint16

// Gets @proxy's destination port; that is, the port on the
// destination host that will be connected to via the proxy, not the
// port number of the proxy itself.
func (x *ProxyAddress) GetDestinationPort() uint16 {

	cret := xProxyAddressGetDestinationPort(x.GoPointer())
	return cret
}

var xProxyAddressGetDestinationProtocol func(uintptr) string

// Gets the protocol that is being spoken to the destination
// server; eg, "http" or "ftp".
func (x *ProxyAddress) GetDestinationProtocol() string {

	cret := xProxyAddressGetDestinationProtocol(x.GoPointer())
	return cret
}

var xProxyAddressGetPassword func(uintptr) string

// Gets @proxy's password.
func (x *ProxyAddress) GetPassword() string {

	cret := xProxyAddressGetPassword(x.GoPointer())
	return cret
}

var xProxyAddressGetProtocol func(uintptr) string

// Gets @proxy's protocol. eg, "socks" or "http"
func (x *ProxyAddress) GetProtocol() string {

	cret := xProxyAddressGetProtocol(x.GoPointer())
	return cret
}

var xProxyAddressGetUri func(uintptr) string

// Gets the proxy URI that @proxy was constructed from.
func (x *ProxyAddress) GetUri() string {

	cret := xProxyAddressGetUri(x.GoPointer())
	return cret
}

var xProxyAddressGetUsername func(uintptr) string

// Gets @proxy's username.
func (x *ProxyAddress) GetUsername() string {

	cret := xProxyAddressGetUsername(x.GoPointer())
	return cret
}

func (c *ProxyAddress) GoPointer() uintptr {
	return c.Ptr
}

func (c *ProxyAddress) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// Creates a #GSocketAddressEnumerator for @connectable.
func (x *ProxyAddress) Enumerate() *SocketAddressEnumerator {
	var cls *SocketAddressEnumerator

	cret := XGSocketConnectableEnumerate(x.GoPointer())

	if cret == 0 {
		return cls
	}
	cls = &SocketAddressEnumerator{}
	cls.Ptr = cret
	return cls
}

// Creates a #GSocketAddressEnumerator for @connectable that will
// return a #GProxyAddress for each of its addresses that you must connect
// to via a proxy.
//
// If @connectable does not implement
// g_socket_connectable_proxy_enumerate(), this will fall back to
// calling g_socket_connectable_enumerate().
func (x *ProxyAddress) ProxyEnumerate() *SocketAddressEnumerator {
	var cls *SocketAddressEnumerator

	cret := XGSocketConnectableProxyEnumerate(x.GoPointer())

	if cret == 0 {
		return cls
	}
	cls = &SocketAddressEnumerator{}
	cls.Ptr = cret
	return cls
}

// Format a #GSocketConnectable as a string. This is a human-readable format for
// use in debugging output, and is not a stable serialization format. It is not
// suitable for use in user interfaces as it exposes too much information for a
// user.
//
// If the #GSocketConnectable implementation does not support string formatting,
// the implementation’s type name will be returned as a fallback.
func (x *ProxyAddress) ToString() string {

	cret := XGSocketConnectableToString(x.GoPointer())
	return cret
}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GIO"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xNewProxyAddress, lib, "g_proxy_address_new")

	core.PuregoSafeRegister(&xProxyAddressGetDestinationHostname, lib, "g_proxy_address_get_destination_hostname")
	core.PuregoSafeRegister(&xProxyAddressGetDestinationPort, lib, "g_proxy_address_get_destination_port")
	core.PuregoSafeRegister(&xProxyAddressGetDestinationProtocol, lib, "g_proxy_address_get_destination_protocol")
	core.PuregoSafeRegister(&xProxyAddressGetPassword, lib, "g_proxy_address_get_password")
	core.PuregoSafeRegister(&xProxyAddressGetProtocol, lib, "g_proxy_address_get_protocol")
	core.PuregoSafeRegister(&xProxyAddressGetUri, lib, "g_proxy_address_get_uri")
	core.PuregoSafeRegister(&xProxyAddressGetUsername, lib, "g_proxy_address_get_username")

}
