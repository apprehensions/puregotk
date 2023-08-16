// Package gio was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gio

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/glib"
	"github.com/jwijenbergh/puregotk/v4/gobject"
)

// Class structure for #GDBusInterfaceSkeleton.
type DBusInterfaceSkeletonClass struct {
	ParentClass uintptr

	VfuncPadding uintptr

	SignalPadding uintptr
}

type DBusInterfaceSkeletonPrivate struct {
}

// Abstract base class for D-Bus interfaces on the service side.
type DBusInterfaceSkeleton struct {
	gobject.Object
}

func DBusInterfaceSkeletonNewFromInternalPtr(ptr uintptr) *DBusInterfaceSkeleton {
	cls := &DBusInterfaceSkeleton{}
	cls.Ptr = ptr
	return cls
}

var xDBusInterfaceSkeletonExport func(uintptr, uintptr, string, **glib.Error) bool

// Exports @interface_ at @object_path on @connection.
//
// This can be called multiple times to export the same @interface_
// onto multiple connections however the @object_path provided must be
// the same for all connections.
//
// Use g_dbus_interface_skeleton_unexport() to unexport the object.
func (x *DBusInterfaceSkeleton) Export(ConnectionVar *DBusConnection, ObjectPathVar string) (bool, error) {
	var cerr *glib.Error

	cret := xDBusInterfaceSkeletonExport(x.GoPointer(), ConnectionVar.GoPointer(), ObjectPathVar, &cerr)
	if cerr == nil {
		return cret, nil
	}
	return cret, cerr

}

var xDBusInterfaceSkeletonFlush func(uintptr)

// If @interface_ has outstanding changes, request for these changes to be
// emitted immediately.
//
// For example, an exported D-Bus interface may queue up property
// changes and emit the
// `org.freedesktop.DBus.Properties.PropertiesChanged`
// signal later (e.g. in an idle handler). This technique is useful
// for collapsing multiple property changes into one.
func (x *DBusInterfaceSkeleton) Flush() {

	xDBusInterfaceSkeletonFlush(x.GoPointer())

}

var xDBusInterfaceSkeletonGetConnection func(uintptr) uintptr

// Gets the first connection that @interface_ is exported on, if any.
func (x *DBusInterfaceSkeleton) GetConnection() *DBusConnection {
	var cls *DBusConnection

	cret := xDBusInterfaceSkeletonGetConnection(x.GoPointer())

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &DBusConnection{}
	cls.Ptr = cret
	return cls
}

var xDBusInterfaceSkeletonGetConnections func(uintptr) *glib.List

// Gets a list of the connections that @interface_ is exported on.
func (x *DBusInterfaceSkeleton) GetConnections() *glib.List {

	cret := xDBusInterfaceSkeletonGetConnections(x.GoPointer())
	return cret
}

var xDBusInterfaceSkeletonGetFlags func(uintptr) DBusInterfaceSkeletonFlags

// Gets the #GDBusInterfaceSkeletonFlags that describes what the behavior
// of @interface_
func (x *DBusInterfaceSkeleton) GetFlags() DBusInterfaceSkeletonFlags {

	cret := xDBusInterfaceSkeletonGetFlags(x.GoPointer())
	return cret
}

var xDBusInterfaceSkeletonGetInfo func(uintptr) *DBusInterfaceInfo

// Gets D-Bus introspection information for the D-Bus interface
// implemented by @interface_.
func (x *DBusInterfaceSkeleton) GetInfo() *DBusInterfaceInfo {

	cret := xDBusInterfaceSkeletonGetInfo(x.GoPointer())
	return cret
}

var xDBusInterfaceSkeletonGetObjectPath func(uintptr) string

// Gets the object path that @interface_ is exported on, if any.
func (x *DBusInterfaceSkeleton) GetObjectPath() string {

	cret := xDBusInterfaceSkeletonGetObjectPath(x.GoPointer())
	return cret
}

var xDBusInterfaceSkeletonGetProperties func(uintptr) *glib.Variant

// Gets all D-Bus properties for @interface_.
func (x *DBusInterfaceSkeleton) GetProperties() *glib.Variant {

	cret := xDBusInterfaceSkeletonGetProperties(x.GoPointer())
	return cret
}

var xDBusInterfaceSkeletonGetVtable func(uintptr) *DBusInterfaceVTable

// Gets the interface vtable for the D-Bus interface implemented by
// @interface_. The returned function pointers should expect @interface_
// itself to be passed as @user_data.
func (x *DBusInterfaceSkeleton) GetVtable() *DBusInterfaceVTable {

	cret := xDBusInterfaceSkeletonGetVtable(x.GoPointer())
	return cret
}

var xDBusInterfaceSkeletonHasConnection func(uintptr, uintptr) bool

// Checks if @interface_ is exported on @connection.
func (x *DBusInterfaceSkeleton) HasConnection(ConnectionVar *DBusConnection) bool {

	cret := xDBusInterfaceSkeletonHasConnection(x.GoPointer(), ConnectionVar.GoPointer())
	return cret
}

var xDBusInterfaceSkeletonSetFlags func(uintptr, DBusInterfaceSkeletonFlags)

// Sets flags describing what the behavior of @skeleton should be.
func (x *DBusInterfaceSkeleton) SetFlags(FlagsVar DBusInterfaceSkeletonFlags) {

	xDBusInterfaceSkeletonSetFlags(x.GoPointer(), FlagsVar)

}

var xDBusInterfaceSkeletonUnexport func(uintptr)

// Stops exporting @interface_ on all connections it is exported on.
//
// To unexport @interface_ from only a single connection, use
// g_dbus_interface_skeleton_unexport_from_connection()
func (x *DBusInterfaceSkeleton) Unexport() {

	xDBusInterfaceSkeletonUnexport(x.GoPointer())

}

var xDBusInterfaceSkeletonUnexportFromConnection func(uintptr, uintptr)

// Stops exporting @interface_ on @connection.
//
// To stop exporting on all connections the interface is exported on,
// use g_dbus_interface_skeleton_unexport().
func (x *DBusInterfaceSkeleton) UnexportFromConnection(ConnectionVar *DBusConnection) {

	xDBusInterfaceSkeletonUnexportFromConnection(x.GoPointer(), ConnectionVar.GoPointer())

}

func (c *DBusInterfaceSkeleton) GoPointer() uintptr {
	return c.Ptr
}

func (c *DBusInterfaceSkeleton) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// Emitted when a method is invoked by a remote caller and used to
// determine if the method call is authorized.
//
// Note that this signal is emitted in a thread dedicated to
// handling the method call so handlers are allowed to perform
// blocking IO. This means that it is appropriate to call e.g.
// [polkit_authority_check_authorization_sync()](http://hal.freedesktop.org/docs/polkit/PolkitAuthority.html#polkit-authority-check-authorization-sync)
// with the
// [POLKIT_CHECK_AUTHORIZATION_FLAGS_ALLOW_USER_INTERACTION](http://hal.freedesktop.org/docs/polkit/PolkitAuthority.html#POLKIT-CHECK-AUTHORIZATION-FLAGS-ALLOW-USER-INTERACTION:CAPS)
// flag set.
//
// If %FALSE is returned then no further handlers are run and the
// signal handler must take a reference to @invocation and finish
// handling the call (e.g. return an error via
// g_dbus_method_invocation_return_error()).
//
// Otherwise, if %TRUE is returned, signal emission continues. If no
// handlers return %FALSE, then the method is dispatched. If
// @interface has an enclosing #GDBusObjectSkeleton, then the
// #GDBusObjectSkeleton::authorize-method signal handlers run before
// the handlers for this signal.
//
// The default class handler just returns %TRUE.
//
// Please note that the common case is optimized: if no signals
// handlers are connected and the default class handler isn't
// overridden (for both @interface and the enclosing
// #GDBusObjectSkeleton, if any) and #GDBusInterfaceSkeleton:g-flags does
// not have the
// %G_DBUS_INTERFACE_SKELETON_FLAGS_HANDLE_METHOD_INVOCATIONS_IN_THREAD
// flags set, no dedicated thread is ever used and the call will be
// handled in the same thread as the object that @interface belongs
// to was exported in.
func (x *DBusInterfaceSkeleton) ConnectGAuthorizeMethod(cb func(DBusInterfaceSkeleton, uintptr) bool) {
	fcb := func(clsPtr uintptr, InvocationVarp uintptr) bool {
		fa := DBusInterfaceSkeleton{}
		fa.Ptr = clsPtr

		return cb(fa, InvocationVarp)

	}
	gobject.ObjectConnect(x.GoPointer(), "signal::g-authorize-method", purego.NewCallback(fcb))
}

// Gets the #GDBusObject that @interface_ belongs to, if any.
func (x *DBusInterfaceSkeleton) DupObject() *DBusObjectBase {
	var cls *DBusObjectBase

	cret := XGDbusInterfaceDupObject(x.GoPointer())

	if cret == 0 {
		return nil
	}
	cls = &DBusObjectBase{}
	cls.Ptr = cret
	return cls
}

// Gets the #GDBusObject that @interface_ belongs to, if any.
//
// It is not safe to use the returned object if @interface_ or
// the returned object is being used from other threads. See
// g_dbus_interface_dup_object() for a thread-safe alternative.
func (x *DBusInterfaceSkeleton) GetObject() *DBusObjectBase {
	var cls *DBusObjectBase

	cret := XGDbusInterfaceGetObject(x.GoPointer())

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &DBusObjectBase{}
	cls.Ptr = cret
	return cls
}

// Sets the #GDBusObject for @interface_ to @object.
//
// Note that @interface_ will hold a weak reference to @object.
func (x *DBusInterfaceSkeleton) SetObject(ObjectVar DBusObject) {

	XGDbusInterfaceSetObject(x.GoPointer(), ObjectVar.GoPointer())

}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GIO"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xDBusInterfaceSkeletonExport, lib, "g_dbus_interface_skeleton_export")
	core.PuregoSafeRegister(&xDBusInterfaceSkeletonFlush, lib, "g_dbus_interface_skeleton_flush")
	core.PuregoSafeRegister(&xDBusInterfaceSkeletonGetConnection, lib, "g_dbus_interface_skeleton_get_connection")
	core.PuregoSafeRegister(&xDBusInterfaceSkeletonGetConnections, lib, "g_dbus_interface_skeleton_get_connections")
	core.PuregoSafeRegister(&xDBusInterfaceSkeletonGetFlags, lib, "g_dbus_interface_skeleton_get_flags")
	core.PuregoSafeRegister(&xDBusInterfaceSkeletonGetInfo, lib, "g_dbus_interface_skeleton_get_info")
	core.PuregoSafeRegister(&xDBusInterfaceSkeletonGetObjectPath, lib, "g_dbus_interface_skeleton_get_object_path")
	core.PuregoSafeRegister(&xDBusInterfaceSkeletonGetProperties, lib, "g_dbus_interface_skeleton_get_properties")
	core.PuregoSafeRegister(&xDBusInterfaceSkeletonGetVtable, lib, "g_dbus_interface_skeleton_get_vtable")
	core.PuregoSafeRegister(&xDBusInterfaceSkeletonHasConnection, lib, "g_dbus_interface_skeleton_has_connection")
	core.PuregoSafeRegister(&xDBusInterfaceSkeletonSetFlags, lib, "g_dbus_interface_skeleton_set_flags")
	core.PuregoSafeRegister(&xDBusInterfaceSkeletonUnexport, lib, "g_dbus_interface_skeleton_unexport")
	core.PuregoSafeRegister(&xDBusInterfaceSkeletonUnexportFromConnection, lib, "g_dbus_interface_skeleton_unexport_from_connection")

}
