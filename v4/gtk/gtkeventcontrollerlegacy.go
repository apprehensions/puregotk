// Package gtk was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gtk

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/gobject"
)

type EventControllerLegacyClass struct {
}

// `GtkEventControllerLegacy` is an event controller that provides raw
// access to the event stream.
//
// It should only be used as a last resort if none of the other event
// controllers or gestures do the job.
type EventControllerLegacy struct {
	EventController
}

func EventControllerLegacyNewFromInternalPtr(ptr uintptr) *EventControllerLegacy {
	cls := &EventControllerLegacy{}
	cls.Ptr = ptr
	return cls
}

var xNewEventControllerLegacy func() uintptr

// Creates a new legacy event controller.
func NewEventControllerLegacy() *EventController {
	var cls *EventController

	cret := xNewEventControllerLegacy()

	if cret == 0 {
		return cls
	}
	cls = &EventController{}
	cls.Ptr = cret
	return cls
}

func (c *EventControllerLegacy) GoPointer() uintptr {
	return c.Ptr
}

func (c *EventControllerLegacy) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// Emitted for each GDK event delivered to @controller.
func (x *EventControllerLegacy) ConnectEvent(cb func(EventControllerLegacy, uintptr) bool) {
	fcb := func(clsPtr uintptr, EventVarp uintptr) bool {
		fa := EventControllerLegacy{}
		fa.Ptr = clsPtr

		return cb(fa, EventVarp)

	}
	gobject.ObjectConnect(x.GoPointer(), "signal::event", purego.NewCallback(fcb))
}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GTK"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xNewEventControllerLegacy, lib, "gtk_event_controller_legacy_new")

}
