// Package gio was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gio

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
)

// The virtual function table for #GDebugController.
type DebugControllerInterface struct {
	GIface uintptr
}

// #GDebugController is an interface to expose control of debugging features and
// debug output.
//
// It is implemented on Linux using #GDebugControllerDBus, which exposes a D-Bus
// interface to allow authenticated peers to control debug features in this
// process.
//
// Whether debug output is enabled is exposed as
// #GDebugController:debug-enabled. This controls g_log_set_debug_enabled() by
// default. Application code may connect to the #GObject::notify signal for it
// to control other parts of its debug infrastructure as necessary.
//
// If your application or service is using the default GLib log writer function,
// creating one of the built-in implementations of #GDebugController should be
// all that’s needed to dynamically enable or disable debug output.
type DebugController interface {
	GoPointer() uintptr
	SetGoPointer(uintptr)
	GetDebugEnabled() bool
	SetDebugEnabled(DebugEnabledVar bool)
}
type DebugControllerBase struct {
	Ptr uintptr
}

func (x *DebugControllerBase) GoPointer() uintptr {
	return x.Ptr
}

func (x *DebugControllerBase) SetGoPointer(ptr uintptr) {
	x.Ptr = ptr
}

// Get the value of #GDebugController:debug-enabled.
func (x *DebugControllerBase) GetDebugEnabled() bool {

	cret := XGDebugControllerGetDebugEnabled(x.GoPointer())
	return cret
}

// Set the value of #GDebugController:debug-enabled.
func (x *DebugControllerBase) SetDebugEnabled(DebugEnabledVar bool) {

	XGDebugControllerSetDebugEnabled(x.GoPointer(), DebugEnabledVar)

}

var XGDebugControllerGetDebugEnabled func(uintptr) bool
var XGDebugControllerSetDebugEnabled func(uintptr, bool)

func init() {
	lib, err := purego.Dlopen(core.GetPath("GIO"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&XGDebugControllerGetDebugEnabled, lib, "g_debug_controller_get_debug_enabled")
	core.PuregoSafeRegister(&XGDebugControllerSetDebugEnabled, lib, "g_debug_controller_set_debug_enabled")

}
