// Package gio was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gio

import (
	"unsafe"

	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
)

// Provides an interface for returning a #GSocketAddressEnumerator
// and #GProxyAddressEnumerator
type SocketConnectableIface struct {
	GIface uintptr
}

func (x *SocketConnectableIface) GoPointer() uintptr {
	return uintptr(unsafe.Pointer(x))
}

// Objects that describe one or more potential socket endpoints
// implement #GSocketConnectable. Callers can then use
// g_socket_connectable_enumerate() to get a #GSocketAddressEnumerator
// to try out each socket address in turn until one succeeds, as shown
// in the sample code below.
//
// |[&lt;!-- language="C" --&gt;
// MyConnectionType *
// connect_to_host (const char    *hostname,
//
//	guint16        port,
//	GCancellable  *cancellable,
//	GError       **error)
//
//	{
//	  MyConnection *conn = NULL;
//	  GSocketConnectable *addr;
//	  GSocketAddressEnumerator *enumerator;
//	  GSocketAddress *sockaddr;
//	  GError *conn_error = NULL;
//
//	  addr = g_network_address_new (hostname, port);
//	  enumerator = g_socket_connectable_enumerate (addr);
//	  g_object_unref (addr);
//
//	  // Try each sockaddr until we succeed. Record the first connection error,
//	  // but not any further ones (since they'll probably be basically the same
//	  // as the first).
//	  while (!conn &amp;&amp; (sockaddr = g_socket_address_enumerator_next (enumerator, cancellable, error))
//	    {
//	      conn = connect_to_sockaddr (sockaddr, conn_error ? NULL : &amp;conn_error);
//	      g_object_unref (sockaddr);
//	    }
//	  g_object_unref (enumerator);
//
//	  if (conn)
//	    {
//	      if (conn_error)
//	        {
//	          // We couldn't connect to the first address, but we succeeded
//	          // in connecting to a later address.
//	          g_error_free (conn_error);
//	        }
//	      return conn;
//	    }
//	  else if (error)
//	    {
//	      /// Either initial lookup failed, or else the caller cancelled us.
//	      if (conn_error)
//	        g_error_free (conn_error);
//	      return NULL;
//	    }
//	  else
//	    {
//	      g_error_propagate (error, conn_error);
//	      return NULL;
//	    }
//	}
//
// ]|
type SocketConnectable interface {
	GoPointer() uintptr
	SetGoPointer(uintptr)
	Enumerate() *SocketAddressEnumerator
	ProxyEnumerate() *SocketAddressEnumerator
	ToString() string
}
type SocketConnectableBase struct {
	Ptr uintptr
}

func (x *SocketConnectableBase) GoPointer() uintptr {
	return x.Ptr
}

func (x *SocketConnectableBase) SetGoPointer(ptr uintptr) {
	x.Ptr = ptr
}

// Creates a #GSocketAddressEnumerator for @connectable.
func (x *SocketConnectableBase) Enumerate() *SocketAddressEnumerator {
	var cls *SocketAddressEnumerator

	cret := XGSocketConnectableEnumerate(x.GoPointer())

	if cret == 0 {
		return nil
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
func (x *SocketConnectableBase) ProxyEnumerate() *SocketAddressEnumerator {
	var cls *SocketAddressEnumerator

	cret := XGSocketConnectableProxyEnumerate(x.GoPointer())

	if cret == 0 {
		return nil
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
func (x *SocketConnectableBase) ToString() string {

	cret := XGSocketConnectableToString(x.GoPointer())
	return cret
}

var XGSocketConnectableEnumerate func(uintptr) uintptr
var XGSocketConnectableProxyEnumerate func(uintptr) uintptr
var XGSocketConnectableToString func(uintptr) string

func init() {
	lib, err := purego.Dlopen(core.GetPath("GIO"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&XGSocketConnectableEnumerate, lib, "g_socket_connectable_enumerate")
	core.PuregoSafeRegister(&XGSocketConnectableProxyEnumerate, lib, "g_socket_connectable_proxy_enumerate")
	core.PuregoSafeRegister(&XGSocketConnectableToString, lib, "g_socket_connectable_to_string")

}
