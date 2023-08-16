// Package gtk was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gtk

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/gobject"
)

type GestureZoomClass struct {
}

// `GtkGestureZoom` is a `GtkGesture` for 2-finger pinch/zoom gestures.
//
// Whenever the distance between both tracked sequences changes, the
// [signal@Gtk.GestureZoom::scale-changed] signal is emitted to report
// the scale factor.
type GestureZoom struct {
	Gesture
}

func GestureZoomNewFromInternalPtr(ptr uintptr) *GestureZoom {
	cls := &GestureZoom{}
	cls.Ptr = ptr
	return cls
}

var xNewGestureZoom func() uintptr

// Returns a newly created `GtkGesture` that recognizes
// pinch/zoom gestures.
func NewGestureZoom() *Gesture {
	var cls *Gesture

	cret := xNewGestureZoom()

	if cret == 0 {
		return nil
	}
	cls = &Gesture{}
	cls.Ptr = cret
	return cls
}

var xGestureZoomGetScaleDelta func(uintptr) float64

// Gets the scale delta.
//
// If @gesture is active, this function returns the zooming
// difference since the gesture was recognized (hence the
// starting point is considered 1:1). If @gesture is not
// active, 1 is returned.
func (x *GestureZoom) GetScaleDelta() float64 {

	cret := xGestureZoomGetScaleDelta(x.GoPointer())
	return cret
}

func (c *GestureZoom) GoPointer() uintptr {
	return c.Ptr
}

func (c *GestureZoom) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// Emitted whenever the distance between both tracked sequences changes.
func (x *GestureZoom) ConnectScaleChanged(cb func(GestureZoom, float64)) {
	fcb := func(clsPtr uintptr, ScaleVarp float64) {
		fa := GestureZoom{}
		fa.Ptr = clsPtr

		cb(fa, ScaleVarp)

	}
	gobject.ObjectConnect(x.GoPointer(), "signal::scale-changed", purego.NewCallback(fcb))
}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GTK"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xNewGestureZoom, lib, "gtk_gesture_zoom_new")

	core.PuregoSafeRegister(&xGestureZoomGetScaleDelta, lib, "gtk_gesture_zoom_get_scale_delta")

}
