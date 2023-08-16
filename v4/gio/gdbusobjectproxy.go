// Package gio was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gio

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/glib"
	"github.com/jwijenbergh/puregotk/v4/gobject"
)

// Class structure for #GDBusObjectProxy.
type DBusObjectProxyClass struct {
	ParentClass uintptr

	Padding uintptr
}

type DBusObjectProxyPrivate struct {
}

// A #GDBusObjectProxy is an object used to represent a remote object
// with one or more D-Bus interfaces. Normally, you don't instantiate
// a #GDBusObjectProxy yourself - typically #GDBusObjectManagerClient
// is used to obtain it.
type DBusObjectProxy struct {
	gobject.Object
}

func DBusObjectProxyNewFromInternalPtr(ptr uintptr) *DBusObjectProxy {
	cls := &DBusObjectProxy{}
	cls.Ptr = ptr
	return cls
}

var xNewDBusObjectProxy func(uintptr, string) uintptr

// Creates a new #GDBusObjectProxy for the given connection and
// object path.
func NewDBusObjectProxy(ConnectionVar *DBusConnection, ObjectPathVar string) *DBusObjectProxy {
	var cls *DBusObjectProxy

	cret := xNewDBusObjectProxy(ConnectionVar.GoPointer(), ObjectPathVar)

	if cret == 0 {
		return nil
	}
	cls = &DBusObjectProxy{}
	cls.Ptr = cret
	return cls
}

var xDBusObjectProxyGetConnection func(uintptr) uintptr

// Gets the connection that @proxy is for.
func (x *DBusObjectProxy) GetConnection() *DBusConnection {
	var cls *DBusConnection

	cret := xDBusObjectProxyGetConnection(x.GoPointer())

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &DBusConnection{}
	cls.Ptr = cret
	return cls
}

func (c *DBusObjectProxy) GoPointer() uintptr {
	return c.Ptr
}

func (c *DBusObjectProxy) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// Gets the D-Bus interface with name @interface_name associated with
// @object, if any.
func (x *DBusObjectProxy) GetInterface(InterfaceNameVar string) *DBusInterfaceBase {
	var cls *DBusInterfaceBase

	cret := XGDbusObjectGetInterface(x.GoPointer(), InterfaceNameVar)

	if cret == 0 {
		return nil
	}
	cls = &DBusInterfaceBase{}
	cls.Ptr = cret
	return cls
}

// Gets the D-Bus interfaces associated with @object.
func (x *DBusObjectProxy) GetInterfaces() *glib.List {

	cret := XGDbusObjectGetInterfaces(x.GoPointer())
	return cret
}

// Gets the object path for @object.
func (x *DBusObjectProxy) GetObjectPath() string {

	cret := XGDbusObjectGetObjectPath(x.GoPointer())
	return cret
}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GIO"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xNewDBusObjectProxy, lib, "g_dbus_object_proxy_new")

	core.PuregoSafeRegister(&xDBusObjectProxyGetConnection, lib, "g_dbus_object_proxy_get_connection")

}
