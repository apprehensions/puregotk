// Package gtk was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gtk

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/gdk"
)

type GestureSingleClass struct {
}

// `GtkGestureSingle` is a `GtkGestures` subclass optimized for singe-touch
// and mouse gestures.
//
// Under interaction, these gestures stick to the first interacting sequence,
// which is accessible through [method@Gtk.GestureSingle.get_current_sequence]
// while the gesture is being interacted with.
//
// By default gestures react to both %GDK_BUTTON_PRIMARY and touch events.
// [method@Gtk.GestureSingle.set_touch_only] can be used to change the
// touch behavior. Callers may also specify a different mouse button number
// to interact with through [method@Gtk.GestureSingle.set_button], or react
// to any mouse button by setting it to 0. While the gesture is active, the
// button being currently pressed can be known through
// [method@Gtk.GestureSingle.get_current_button].
type GestureSingle struct {
	Gesture
}

func GestureSingleNewFromInternalPtr(ptr uintptr) *GestureSingle {
	cls := &GestureSingle{}
	cls.Ptr = ptr
	return cls
}

var xGestureSingleGetButton func(uintptr) uint

// Returns the button number @gesture listens for.
//
// If this is 0, the gesture reacts to any button press.
func (x *GestureSingle) GetButton() uint {

	cret := xGestureSingleGetButton(x.GoPointer())
	return cret
}

var xGestureSingleGetCurrentButton func(uintptr) uint

// Returns the button number currently interacting
// with @gesture, or 0 if there is none.
func (x *GestureSingle) GetCurrentButton() uint {

	cret := xGestureSingleGetCurrentButton(x.GoPointer())
	return cret
}

var xGestureSingleGetCurrentSequence func(uintptr) *gdk.EventSequence

// Returns the event sequence currently interacting with @gesture.
//
// This is only meaningful if [method@Gtk.Gesture.is_active]
// returns %TRUE.
func (x *GestureSingle) GetCurrentSequence() *gdk.EventSequence {

	cret := xGestureSingleGetCurrentSequence(x.GoPointer())
	return cret
}

var xGestureSingleGetExclusive func(uintptr) bool

// Gets whether a gesture is exclusive.
//
// For more information, see [method@Gtk.GestureSingle.set_exclusive].
func (x *GestureSingle) GetExclusive() bool {

	cret := xGestureSingleGetExclusive(x.GoPointer())
	return cret
}

var xGestureSingleGetTouchOnly func(uintptr) bool

// Returns %TRUE if the gesture is only triggered by touch events.
func (x *GestureSingle) GetTouchOnly() bool {

	cret := xGestureSingleGetTouchOnly(x.GoPointer())
	return cret
}

var xGestureSingleSetButton func(uintptr, uint)

// Sets the button number @gesture listens to.
//
// If non-0, every button press from a different button
// number will be ignored. Touch events implicitly match
// with button 1.
func (x *GestureSingle) SetButton(ButtonVar uint) {

	xGestureSingleSetButton(x.GoPointer(), ButtonVar)

}

var xGestureSingleSetExclusive func(uintptr, bool)

// Sets whether @gesture is exclusive.
//
// An exclusive gesture will only handle pointer and "pointer emulated"
// touch events, so at any given time, there is only one sequence able
// to interact with those.
func (x *GestureSingle) SetExclusive(ExclusiveVar bool) {

	xGestureSingleSetExclusive(x.GoPointer(), ExclusiveVar)

}

var xGestureSingleSetTouchOnly func(uintptr, bool)

// Sets whether to handle only touch events.
//
// If @touch_only is %TRUE, @gesture will only handle events of type
// %GDK_TOUCH_BEGIN, %GDK_TOUCH_UPDATE or %GDK_TOUCH_END. If %FALSE,
// mouse events will be handled too.
func (x *GestureSingle) SetTouchOnly(TouchOnlyVar bool) {

	xGestureSingleSetTouchOnly(x.GoPointer(), TouchOnlyVar)

}

func (c *GestureSingle) GoPointer() uintptr {
	return c.Ptr
}

func (c *GestureSingle) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GTK"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xGestureSingleGetButton, lib, "gtk_gesture_single_get_button")
	core.PuregoSafeRegister(&xGestureSingleGetCurrentButton, lib, "gtk_gesture_single_get_current_button")
	core.PuregoSafeRegister(&xGestureSingleGetCurrentSequence, lib, "gtk_gesture_single_get_current_sequence")
	core.PuregoSafeRegister(&xGestureSingleGetExclusive, lib, "gtk_gesture_single_get_exclusive")
	core.PuregoSafeRegister(&xGestureSingleGetTouchOnly, lib, "gtk_gesture_single_get_touch_only")
	core.PuregoSafeRegister(&xGestureSingleSetButton, lib, "gtk_gesture_single_set_button")
	core.PuregoSafeRegister(&xGestureSingleSetExclusive, lib, "gtk_gesture_single_set_exclusive")
	core.PuregoSafeRegister(&xGestureSingleSetTouchOnly, lib, "gtk_gesture_single_set_touch_only")

}
