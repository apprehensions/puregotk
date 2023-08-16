// Package gtk was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gtk

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/gobject"
)

type GestureDragClass struct {
}

// `GtkGestureDrag` is a `GtkGesture` implementation for drags.
//
// The drag operation itself can be tracked throughout the
// [signal@Gtk.GestureDrag::drag-begin],
// [signal@Gtk.GestureDrag::drag-update] and
// [signal@Gtk.GestureDrag::drag-end] signals, and the relevant
// coordinates can be extracted through
// [method@Gtk.GestureDrag.get_offset] and
// [method@Gtk.GestureDrag.get_start_point].
type GestureDrag struct {
	GestureSingle
}

func GestureDragNewFromInternalPtr(ptr uintptr) *GestureDrag {
	cls := &GestureDrag{}
	cls.Ptr = ptr
	return cls
}

var xNewGestureDrag func() uintptr

// Returns a newly created `GtkGesture` that recognizes drags.
func NewGestureDrag() *Gesture {
	var cls *Gesture

	cret := xNewGestureDrag()

	if cret == 0 {
		return nil
	}
	cls = &Gesture{}
	cls.Ptr = cret
	return cls
}

var xGestureDragGetOffset func(uintptr, float64, float64) bool

// Gets the offset from the start point.
//
// If the @gesture is active, this function returns %TRUE and
// fills in @x and @y with the coordinates of the current point,
// as an offset to the starting drag point.
func (x *GestureDrag) GetOffset(XVar float64, YVar float64) bool {

	cret := xGestureDragGetOffset(x.GoPointer(), XVar, YVar)
	return cret
}

var xGestureDragGetStartPoint func(uintptr, float64, float64) bool

// Gets the point where the drag started.
//
// If the @gesture is active, this function returns %TRUE
// and fills in @x and @y with the drag start coordinates,
// in widget-relative coordinates.
func (x *GestureDrag) GetStartPoint(XVar float64, YVar float64) bool {

	cret := xGestureDragGetStartPoint(x.GoPointer(), XVar, YVar)
	return cret
}

func (c *GestureDrag) GoPointer() uintptr {
	return c.Ptr
}

func (c *GestureDrag) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// Emitted whenever dragging starts.
func (x *GestureDrag) ConnectDragBegin(cb func(GestureDrag, float64, float64)) {
	fcb := func(clsPtr uintptr, StartXVarp float64, StartYVarp float64) {
		fa := GestureDrag{}
		fa.Ptr = clsPtr

		cb(fa, StartXVarp, StartYVarp)

	}
	gobject.ObjectConnect(x.GoPointer(), "signal::drag-begin", purego.NewCallback(fcb))
}

// Emitted whenever the dragging is finished.
func (x *GestureDrag) ConnectDragEnd(cb func(GestureDrag, float64, float64)) {
	fcb := func(clsPtr uintptr, OffsetXVarp float64, OffsetYVarp float64) {
		fa := GestureDrag{}
		fa.Ptr = clsPtr

		cb(fa, OffsetXVarp, OffsetYVarp)

	}
	gobject.ObjectConnect(x.GoPointer(), "signal::drag-end", purego.NewCallback(fcb))
}

// Emitted whenever the dragging point moves.
func (x *GestureDrag) ConnectDragUpdate(cb func(GestureDrag, float64, float64)) {
	fcb := func(clsPtr uintptr, OffsetXVarp float64, OffsetYVarp float64) {
		fa := GestureDrag{}
		fa.Ptr = clsPtr

		cb(fa, OffsetXVarp, OffsetYVarp)

	}
	gobject.ObjectConnect(x.GoPointer(), "signal::drag-update", purego.NewCallback(fcb))
}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GTK"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xNewGestureDrag, lib, "gtk_gesture_drag_new")

	core.PuregoSafeRegister(&xGestureDragGetOffset, lib, "gtk_gesture_drag_get_offset")
	core.PuregoSafeRegister(&xGestureDragGetStartPoint, lib, "gtk_gesture_drag_get_start_point")

}
