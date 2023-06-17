// Package gtk was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gtk

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/gobject"
)

type GestureLongPressClass struct {
}

// `GtkGestureLongPress` is a `GtkGesture` for long presses.
//
// This gesture is also known as “Press and Hold”.
//
// When the timeout is exceeded, the gesture is triggering the
// [signal@Gtk.GestureLongPress::pressed] signal.
//
// If the touchpoint is lifted before the timeout passes, or if
// it drifts too far of the initial press point, the
// [signal@Gtk.GestureLongPress::cancelled] signal will be emitted.
//
// How long the timeout is before the ::pressed signal gets emitted is
// determined by the [property@Gtk.Settings:gtk-long-press-time] setting.
// It can be modified by the [property@Gtk.GestureLongPress:delay-factor]
// property.
type GestureLongPress struct {
	GestureSingle
}

func GestureLongPressNewFromInternalPtr(ptr uintptr) *GestureLongPress {
	cls := &GestureLongPress{}
	cls.Ptr = ptr
	return cls
}

var xNewGestureLongPress func() uintptr

// Returns a newly created `GtkGesture` that recognizes long presses.
func NewGestureLongPress() *Gesture {
	NewGestureLongPressPtr := xNewGestureLongPress()
	if NewGestureLongPressPtr == 0 {
		return nil
	}

	NewGestureLongPressCls := &Gesture{}
	NewGestureLongPressCls.Ptr = NewGestureLongPressPtr
	return NewGestureLongPressCls
}

var xGestureLongPressGetDelayFactor func(uintptr) float64

// Returns the delay factor.
func (x *GestureLongPress) GetDelayFactor() float64 {

	return xGestureLongPressGetDelayFactor(x.GoPointer())

}

var xGestureLongPressSetDelayFactor func(uintptr, float64)

// Applies the given delay factor.
//
// The default long press time will be multiplied by this value.
// Valid values are in the range [0.5..2.0].
func (x *GestureLongPress) SetDelayFactor(DelayFactorVar float64) {

	xGestureLongPressSetDelayFactor(x.GoPointer(), DelayFactorVar)

}

func (c *GestureLongPress) GoPointer() uintptr {
	return c.Ptr
}

func (c *GestureLongPress) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// Emitted whenever a press moved too far, or was released
// before [signal@Gtk.GestureLongPress::pressed] happened.
func (x *GestureLongPress) ConnectCancelled(cb func(GestureLongPress)) {
	fcb := func(clsPtr uintptr) {
		fa := GestureLongPress{}
		fa.Ptr = clsPtr

		cb(fa)

	}
	gobject.ObjectConnect(x.GoPointer(), "signal::cancelled", purego.NewCallback(fcb))
}

// Emitted whenever a press goes unmoved/unreleased longer than
// what the GTK defaults tell.
func (x *GestureLongPress) ConnectPressed(cb func(GestureLongPress, float64, float64)) {
	fcb := func(clsPtr uintptr, XVarp float64, YVarp float64) {
		fa := GestureLongPress{}
		fa.Ptr = clsPtr

		cb(fa, XVarp, YVarp)

	}
	gobject.ObjectConnect(x.GoPointer(), "signal::pressed", purego.NewCallback(fcb))
}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GTK"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xNewGestureLongPress, lib, "gtk_gesture_long_press_new")

	core.PuregoSafeRegister(&xGestureLongPressGetDelayFactor, lib, "gtk_gesture_long_press_get_delay_factor")
	core.PuregoSafeRegister(&xGestureLongPressSetDelayFactor, lib, "gtk_gesture_long_press_set_delay_factor")

}
