// Package gtk was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gtk

import (
	"unsafe"

	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/gdk"
	"github.com/jwijenbergh/puregotk/v4/gobject"
)

type ATContextClass struct {
}

func (x *ATContextClass) GoPointer() uintptr {
	return uintptr(unsafe.Pointer(x))
}

// `GtkATContext` is an abstract class provided by GTK to communicate to
// platform-specific assistive technologies API.
//
// Each platform supported by GTK implements a `GtkATContext` subclass, and
// is responsible for updating the accessible state in response to state
// changes in `GtkAccessible`.
type ATContext struct {
	gobject.Object
}

func ATContextNewFromInternalPtr(ptr uintptr) *ATContext {
	cls := &ATContext{}
	cls.Ptr = ptr
	return cls
}

var xCreateATContext func(AccessibleRole, uintptr, uintptr) uintptr

// Creates a new `GtkATContext` instance for the given accessible role,
// accessible instance, and display connection.
//
// The `GtkATContext` implementation being instantiated will depend on the
// platform.
func CreateATContext(AccessibleRoleVar AccessibleRole, AccessibleVar Accessible, DisplayVar *gdk.Display) *ATContext {
	var cls *ATContext

	cret := xCreateATContext(AccessibleRoleVar, AccessibleVar.GoPointer(), DisplayVar.GoPointer())

	if cret == 0 {
		return nil
	}
	cls = &ATContext{}
	cls.Ptr = cret
	return cls
}

var xATContextGetAccessible func(uintptr) uintptr

// Retrieves the `GtkAccessible` using this context.
func (x *ATContext) GetAccessible() *AccessibleBase {
	var cls *AccessibleBase

	cret := xATContextGetAccessible(x.GoPointer())

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &AccessibleBase{}
	cls.Ptr = cret
	return cls
}

var xATContextGetAccessibleRole func(uintptr) AccessibleRole

// Retrieves the accessible role of this context.
func (x *ATContext) GetAccessibleRole() AccessibleRole {

	cret := xATContextGetAccessibleRole(x.GoPointer())
	return cret
}

func (c *ATContext) GoPointer() uintptr {
	return c.Ptr
}

func (c *ATContext) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// Emitted when the attributes of the accessible for the
// `GtkATContext` instance change.
func (x *ATContext) ConnectStateChange(cb func(ATContext)) uint32 {
	fcb := func(clsPtr uintptr) {
		fa := ATContext{}
		fa.Ptr = clsPtr

		cb(fa)

	}
	return gobject.SignalConnect(x.GoPointer(), "state-change", purego.NewCallback(fcb))
}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GTK"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xCreateATContext, lib, "gtk_at_context_create")

	core.PuregoSafeRegister(&xATContextGetAccessible, lib, "gtk_at_context_get_accessible")
	core.PuregoSafeRegister(&xATContextGetAccessibleRole, lib, "gtk_at_context_get_accessible_role")

}
