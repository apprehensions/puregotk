// Package gio was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gio

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/gobject"
)

type NetworkServiceClass struct {
	ParentClass uintptr
}

type NetworkServicePrivate struct {
}

// Like #GNetworkAddress does with hostnames, #GNetworkService
// provides an easy way to resolve a SRV record, and then attempt to
// connect to one of the hosts that implements that service, handling
// service priority/weighting, multiple IP addresses, and multiple
// address families.
//
// See #GSrvTarget for more information about SRV records, and see
// #GSocketConnectable for an example of using the connectable
// interface.
type NetworkService struct {
	gobject.Object
}

func NetworkServiceNewFromInternalPtr(ptr uintptr) *NetworkService {
	cls := &NetworkService{}
	cls.Ptr = ptr
	return cls
}

var xNewNetworkService func(string, string, string) uintptr

// Creates a new #GNetworkService representing the given @service,
// @protocol, and @domain. This will initially be unresolved; use the
// #GSocketConnectable interface to resolve it.
func NewNetworkService(ServiceVar string, ProtocolVar string, DomainVar string) *NetworkService {
	NewNetworkServicePtr := xNewNetworkService(ServiceVar, ProtocolVar, DomainVar)
	if NewNetworkServicePtr == 0 {
		return nil
	}

	NewNetworkServiceCls := &NetworkService{}
	NewNetworkServiceCls.Ptr = NewNetworkServicePtr
	return NewNetworkServiceCls
}

var xNetworkServiceGetDomain func(uintptr) string

// Gets the domain that @srv serves. This might be either UTF-8 or
// ASCII-encoded, depending on what @srv was created with.
func (x *NetworkService) GetDomain() string {

	return xNetworkServiceGetDomain(x.GoPointer())

}

var xNetworkServiceGetProtocol func(uintptr) string

// Gets @srv's protocol name (eg, "tcp").
func (x *NetworkService) GetProtocol() string {

	return xNetworkServiceGetProtocol(x.GoPointer())

}

var xNetworkServiceGetScheme func(uintptr) string

// Gets the URI scheme used to resolve proxies. By default, the service name
// is used as scheme.
func (x *NetworkService) GetScheme() string {

	return xNetworkServiceGetScheme(x.GoPointer())

}

var xNetworkServiceGetService func(uintptr) string

// Gets @srv's service name (eg, "ldap").
func (x *NetworkService) GetService() string {

	return xNetworkServiceGetService(x.GoPointer())

}

var xNetworkServiceSetScheme func(uintptr, string)

// Set's the URI scheme used to resolve proxies. By default, the service name
// is used as scheme.
func (x *NetworkService) SetScheme(SchemeVar string) {

	xNetworkServiceSetScheme(x.GoPointer(), SchemeVar)

}

func (c *NetworkService) GoPointer() uintptr {
	return c.Ptr
}

func (c *NetworkService) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// Creates a #GSocketAddressEnumerator for @connectable.
func (x *NetworkService) Enumerate() *SocketAddressEnumerator {

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
func (x *NetworkService) ProxyEnumerate() *SocketAddressEnumerator {

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
func (x *NetworkService) ToString() string {

	return XGSocketConnectableToString(x.GoPointer())

}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GIO"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xNewNetworkService, lib, "g_network_service_new")

	core.PuregoSafeRegister(&xNetworkServiceGetDomain, lib, "g_network_service_get_domain")
	core.PuregoSafeRegister(&xNetworkServiceGetProtocol, lib, "g_network_service_get_protocol")
	core.PuregoSafeRegister(&xNetworkServiceGetScheme, lib, "g_network_service_get_scheme")
	core.PuregoSafeRegister(&xNetworkServiceGetService, lib, "g_network_service_get_service")
	core.PuregoSafeRegister(&xNetworkServiceSetScheme, lib, "g_network_service_set_scheme")

}
