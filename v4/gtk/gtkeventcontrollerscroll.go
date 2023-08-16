// Package gtk was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gtk

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/gdk"
	"github.com/jwijenbergh/puregotk/v4/gobject"
)

type EventControllerScrollClass struct {
}

// Describes the behavior of a `GtkEventControllerScroll`.
type EventControllerScrollFlags int

const (

	// Don't emit scroll.
	EventControllerScrollNoneValue EventControllerScrollFlags = 0
	// Emit scroll with vertical deltas.
	EventControllerScrollVerticalValue EventControllerScrollFlags = 1
	// Emit scroll with horizontal deltas.
	EventControllerScrollHorizontalValue EventControllerScrollFlags = 2
	// Only emit deltas that are multiples of 1.
	EventControllerScrollDiscreteValue EventControllerScrollFlags = 4
	// Emit ::decelerate after continuous scroll finishes.
	EventControllerScrollKineticValue EventControllerScrollFlags = 8
	// Emit scroll on both axes.
	EventControllerScrollBothAxesValue EventControllerScrollFlags = 3
)

// `GtkEventControllerScroll` is an event controller that handles scroll
// events.
//
// It is capable of handling both discrete and continuous scroll
// events from mice or touchpads, abstracting them both with the
// [signal@Gtk.EventControllerScroll::scroll] signal. Deltas in
// the discrete case are multiples of 1.
//
// In the case of continuous scroll events, `GtkEventControllerScroll`
// encloses all [signal@Gtk.EventControllerScroll::scroll] emissions
// between two [signal@Gtk.EventControllerScroll::scroll-begin] and
// [signal@Gtk.EventControllerScroll::scroll-end] signals.
//
// The behavior of the event controller can be modified by the flags
// given at creation time, or modified at a later point through
// [method@Gtk.EventControllerScroll.set_flags] (e.g. because the scrolling
// conditions of the widget changed).
//
// The controller can be set up to emit motion for either/both vertical
// and horizontal scroll events through %GTK_EVENT_CONTROLLER_SCROLL_VERTICAL,
// %GTK_EVENT_CONTROLLER_SCROLL_HORIZONTAL and %GTK_EVENT_CONTROLLER_SCROLL_BOTH_AXES.
// If any axis is disabled, the respective [signal@Gtk.EventControllerScroll::scroll]
// delta will be 0. Vertical scroll events will be translated to horizontal
// motion for the devices incapable of horizontal scrolling.
//
// The event controller can also be forced to emit discrete events on all
// devices through %GTK_EVENT_CONTROLLER_SCROLL_DISCRETE. This can be used
// to implement discrete actions triggered through scroll events (e.g.
// switching across combobox options).
//
// The %GTK_EVENT_CONTROLLER_SCROLL_KINETIC flag toggles the emission of the
// [signal@Gtk.EventControllerScroll::decelerate] signal, emitted at the end
// of scrolling with two X/Y velocity arguments that are consistent with the
// motion that was received.
type EventControllerScroll struct {
	EventController
}

func EventControllerScrollNewFromInternalPtr(ptr uintptr) *EventControllerScroll {
	cls := &EventControllerScroll{}
	cls.Ptr = ptr
	return cls
}

var xNewEventControllerScroll func(EventControllerScrollFlags) uintptr

// Creates a new event controller that will handle scroll events.
func NewEventControllerScroll(FlagsVar EventControllerScrollFlags) *EventController {
	var cls *EventController

	cret := xNewEventControllerScroll(FlagsVar)

	if cret == 0 {
		return cls
	}
	cls = &EventController{}
	cls.Ptr = cret
	return cls
}

var xEventControllerScrollGetFlags func(uintptr) EventControllerScrollFlags

// Gets the flags conditioning the scroll controller behavior.
func (x *EventControllerScroll) GetFlags() EventControllerScrollFlags {

	cret := xEventControllerScrollGetFlags(x.GoPointer())
	return cret
}

var xEventControllerScrollGetUnit func(uintptr) gdk.ScrollUnit

// Gets the scroll unit of the last
// [signal@Gtk.EventControllerScroll::scroll] signal received.
//
// Always returns %GDK_SCROLL_UNIT_WHEEL if the
// %GTK_EVENT_CONTROLLER_SCROLL_DISCRETE flag is set.
func (x *EventControllerScroll) GetUnit() gdk.ScrollUnit {

	cret := xEventControllerScrollGetUnit(x.GoPointer())
	return cret
}

var xEventControllerScrollSetFlags func(uintptr, EventControllerScrollFlags)

// Sets the flags conditioning scroll controller behavior.
func (x *EventControllerScroll) SetFlags(FlagsVar EventControllerScrollFlags) {

	xEventControllerScrollSetFlags(x.GoPointer(), FlagsVar)

}

func (c *EventControllerScroll) GoPointer() uintptr {
	return c.Ptr
}

func (c *EventControllerScroll) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// Emitted after scroll is finished if the
// %GTK_EVENT_CONTROLLER_SCROLL_KINETIC flag is set.
//
// @vel_x and @vel_y express the initial velocity that was
// imprinted by the scroll events. @vel_x and @vel_y are expressed in
// pixels/ms.
func (x *EventControllerScroll) ConnectDecelerate(cb func(EventControllerScroll, float64, float64)) {
	fcb := func(clsPtr uintptr, VelXVarp float64, VelYVarp float64) {
		fa := EventControllerScroll{}
		fa.Ptr = clsPtr

		cb(fa, VelXVarp, VelYVarp)

	}
	gobject.ObjectConnect(x.GoPointer(), "signal::decelerate", purego.NewCallback(fcb))
}

// Signals that the widget should scroll by the
// amount specified by @dx and @dy.
//
// For the representation unit of the deltas, see
// [method@Gtk.EventControllerScroll.get_unit].
func (x *EventControllerScroll) ConnectScroll(cb func(EventControllerScroll, float64, float64) bool) {
	fcb := func(clsPtr uintptr, DxVarp float64, DyVarp float64) bool {
		fa := EventControllerScroll{}
		fa.Ptr = clsPtr

		return cb(fa, DxVarp, DyVarp)

	}
	gobject.ObjectConnect(x.GoPointer(), "signal::scroll", purego.NewCallback(fcb))
}

// Signals that a new scrolling operation has begun.
//
// It will only be emitted on devices capable of it.
func (x *EventControllerScroll) ConnectScrollBegin(cb func(EventControllerScroll)) {
	fcb := func(clsPtr uintptr) {
		fa := EventControllerScroll{}
		fa.Ptr = clsPtr

		cb(fa)

	}
	gobject.ObjectConnect(x.GoPointer(), "signal::scroll-begin", purego.NewCallback(fcb))
}

// Signals that a scrolling operation has finished.
//
// It will only be emitted on devices capable of it.
func (x *EventControllerScroll) ConnectScrollEnd(cb func(EventControllerScroll)) {
	fcb := func(clsPtr uintptr) {
		fa := EventControllerScroll{}
		fa.Ptr = clsPtr

		cb(fa)

	}
	gobject.ObjectConnect(x.GoPointer(), "signal::scroll-end", purego.NewCallback(fcb))
}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GTK"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xNewEventControllerScroll, lib, "gtk_event_controller_scroll_new")

	core.PuregoSafeRegister(&xEventControllerScrollGetFlags, lib, "gtk_event_controller_scroll_get_flags")
	core.PuregoSafeRegister(&xEventControllerScrollGetUnit, lib, "gtk_event_controller_scroll_get_unit")
	core.PuregoSafeRegister(&xEventControllerScrollSetFlags, lib, "gtk_event_controller_scroll_set_flags")

}
