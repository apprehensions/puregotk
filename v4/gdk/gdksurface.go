// Package gdk was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gdk

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/cairo"
	"github.com/jwijenbergh/puregotk/v4/glib"
	"github.com/jwijenbergh/puregotk/v4/gobject"
)

type SurfaceClass struct {
}

// A `GdkSurface` is a rectangular region on the screen.
//
// It’s a low-level object, used to implement high-level objects
// such as [class@Gtk.Window] or [class@Gtk.Dialog] in GTK.
//
// The surfaces you see in practice are either [iface@Gdk.Toplevel] or
// [iface@Gdk.Popup], and those interfaces provide much of the required
// API to interact with these surfaces. Other, more specialized surface
// types exist, but you will rarely interact with them directly.
type Surface struct {
	gobject.Object
}

func SurfaceNewFromInternalPtr(ptr uintptr) *Surface {
	cls := &Surface{}
	cls.Ptr = ptr
	return cls
}

var xNewPopupSurface func(uintptr, bool) uintptr

// Create a new popup surface.
//
// The surface will be attached to @parent and can be positioned
// relative to it using [method@Gdk.Popup.present].
func NewPopupSurface(ParentVar *Surface, AutohideVar bool) *Surface {
	var cls *Surface

	cret := xNewPopupSurface(ParentVar.GoPointer(), AutohideVar)

	if cret == 0 {
		return nil
	}
	cls = &Surface{}
	cls.Ptr = cret
	return cls
}

var xNewToplevelSurface func(uintptr) uintptr

// Creates a new toplevel surface.
func NewToplevelSurface(DisplayVar *Display) *Surface {
	var cls *Surface

	cret := xNewToplevelSurface(DisplayVar.GoPointer())

	if cret == 0 {
		return nil
	}
	cls = &Surface{}
	cls.Ptr = cret
	return cls
}

var xSurfaceBeep func(uintptr)

// Emits a short beep associated to @surface.
//
// If the display of @surface does not support per-surface beeps,
// emits a short beep on the display just as [method@Gdk.Display.beep].
func (x *Surface) Beep() {

	xSurfaceBeep(x.GoPointer())

}

var xSurfaceCreateCairoContext func(uintptr) uintptr

// Creates a new `GdkCairoContext` for rendering on @surface.
func (x *Surface) CreateCairoContext() *CairoContext {
	var cls *CairoContext

	cret := xSurfaceCreateCairoContext(x.GoPointer())

	if cret == 0 {
		return nil
	}
	cls = &CairoContext{}
	cls.Ptr = cret
	return cls
}

var xSurfaceCreateGlContext func(uintptr) uintptr

// Creates a new `GdkGLContext` for the `GdkSurface`.
//
// The context is disconnected from any particular surface or surface.
// If the creation of the `GdkGLContext` failed, @error will be set.
// Before using the returned `GdkGLContext`, you will need to
// call [method@Gdk.GLContext.make_current] or [method@Gdk.GLContext.realize].
func (x *Surface) CreateGlContext() (*GLContext, error) {
	var cls *GLContext
	var cerr *glib.Error

	cret := xSurfaceCreateGlContext(x.GoPointer())

	if cret == 0 {
		return nil, cerr
	}
	cls = &GLContext{}
	cls.Ptr = cret
	if cerr == nil {
		return cls, nil
	}
	return cls, cerr

}

var xSurfaceCreateSimilarSurface func(uintptr, cairo.Content, int, int) *cairo.Surface

// Create a new Cairo surface that is as compatible as possible with the
// given @surface.
//
// For example the new surface will have the same fallback resolution
// and font options as @surface. Generally, the new surface will also
// use the same backend as @surface, unless that is not possible for
// some reason. The type of the returned surface may be examined with
// cairo_surface_get_type().
//
// Initially the surface contents are all 0 (transparent if contents
// have transparency, black otherwise.)
//
// This function always returns a valid pointer, but it will return a
// pointer to a “nil” surface if @other is already in an error state
// or any other error occurs.
func (x *Surface) CreateSimilarSurface(ContentVar cairo.Content, WidthVar int, HeightVar int) *cairo.Surface {

	cret := xSurfaceCreateSimilarSurface(x.GoPointer(), ContentVar, WidthVar, HeightVar)
	return cret
}

var xSurfaceCreateVulkanContext func(uintptr) uintptr

// Creates a new `GdkVulkanContext` for rendering on @surface.
//
// If the creation of the `GdkVulkanContext` failed, @error will be set.
func (x *Surface) CreateVulkanContext() (*VulkanContext, error) {
	var cls *VulkanContext
	var cerr *glib.Error

	cret := xSurfaceCreateVulkanContext(x.GoPointer())

	if cret == 0 {
		return nil, cerr
	}
	cls = &VulkanContext{}
	cls.Ptr = cret
	if cerr == nil {
		return cls, nil
	}
	return cls, cerr

}

var xSurfaceDestroy func(uintptr)

// Destroys the window system resources associated with @surface and
// decrements @surface's reference count.
//
// The window system resources for all children of @surface are also
// destroyed, but the children’s reference counts are not decremented.
//
// Note that a surface will not be destroyed automatically when its
// reference count reaches zero. You must call this function yourself
// before that happens.
func (x *Surface) Destroy() {

	xSurfaceDestroy(x.GoPointer())

}

var xSurfaceGetCursor func(uintptr) uintptr

// Retrieves a `GdkCursor` pointer for the cursor currently set on the
// `GdkSurface`.
//
// If the return value is %NULL then there is no custom cursor set on
// the surface, and it is using the cursor for its parent surface.
//
// Use [method@Gdk.Surface.set_cursor] to unset the cursor of the surface.
func (x *Surface) GetCursor() *Cursor {
	var cls *Cursor

	cret := xSurfaceGetCursor(x.GoPointer())

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &Cursor{}
	cls.Ptr = cret
	return cls
}

var xSurfaceGetDeviceCursor func(uintptr, uintptr) uintptr

// Retrieves a `GdkCursor` pointer for the @device currently set on the
// specified `GdkSurface`.
//
// If the return value is %NULL then there is no custom cursor set on the
// specified surface, and it is using the cursor for its parent surface.
//
// Use [method@Gdk.Surface.set_cursor] to unset the cursor of the surface.
func (x *Surface) GetDeviceCursor(DeviceVar *Device) *Cursor {
	var cls *Cursor

	cret := xSurfaceGetDeviceCursor(x.GoPointer(), DeviceVar.GoPointer())

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &Cursor{}
	cls.Ptr = cret
	return cls
}

var xSurfaceGetDevicePosition func(uintptr, uintptr, float64, float64, *ModifierType) bool

// Obtains the current device position and modifier state.
//
// The position is given in coordinates relative to the upper
// left corner of @surface.
func (x *Surface) GetDevicePosition(DeviceVar *Device, XVar float64, YVar float64, MaskVar *ModifierType) bool {

	cret := xSurfaceGetDevicePosition(x.GoPointer(), DeviceVar.GoPointer(), XVar, YVar, MaskVar)
	return cret
}

var xSurfaceGetDisplay func(uintptr) uintptr

// Gets the `GdkDisplay` associated with a `GdkSurface`.
func (x *Surface) GetDisplay() *Display {
	var cls *Display

	cret := xSurfaceGetDisplay(x.GoPointer())

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &Display{}
	cls.Ptr = cret
	return cls
}

var xSurfaceGetFrameClock func(uintptr) uintptr

// Gets the frame clock for the surface.
//
// The frame clock for a surface never changes unless the surface is
// reparented to a new toplevel surface.
func (x *Surface) GetFrameClock() *FrameClock {
	var cls *FrameClock

	cret := xSurfaceGetFrameClock(x.GoPointer())

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &FrameClock{}
	cls.Ptr = cret
	return cls
}

var xSurfaceGetHeight func(uintptr) int

// Returns the height of the given @surface.
//
// Surface size is reported in ”application pixels”, not
// ”device pixels” (see [method@Gdk.Surface.get_scale_factor]).
func (x *Surface) GetHeight() int {

	cret := xSurfaceGetHeight(x.GoPointer())
	return cret
}

var xSurfaceGetMapped func(uintptr) bool

// Checks whether the surface has been mapped.
//
// A surface is mapped with [method@Gdk.Toplevel.present]
// or [method@Gdk.Popup.present].
func (x *Surface) GetMapped() bool {

	cret := xSurfaceGetMapped(x.GoPointer())
	return cret
}

var xSurfaceGetScaleFactor func(uintptr) int

// Returns the internal scale factor that maps from surface coordinates
// to the actual device pixels.
//
// On traditional systems this is 1, but on very high density outputs
// this can be a higher value (often 2). A higher value means that drawing
// is automatically scaled up to a higher resolution, so any code doing
// drawing will automatically look nicer. However, if you are supplying
// pixel-based data the scale value can be used to determine whether to
// use a pixel resource with higher resolution data.
//
// The scale of a surface may change during runtime.
func (x *Surface) GetScaleFactor() int {

	cret := xSurfaceGetScaleFactor(x.GoPointer())
	return cret
}

var xSurfaceGetWidth func(uintptr) int

// Returns the width of the given @surface.
//
// Surface size is reported in ”application pixels”, not
// ”device pixels” (see [method@Gdk.Surface.get_scale_factor]).
func (x *Surface) GetWidth() int {

	cret := xSurfaceGetWidth(x.GoPointer())
	return cret
}

var xSurfaceHide func(uintptr)

// Hide the surface.
//
// For toplevel surfaces, withdraws them, so they will no longer be
// known to the window manager; for all surfaces, unmaps them, so
// they won’t be displayed. Normally done automatically as
// part of [method@Gtk.Widget.hide].
func (x *Surface) Hide() {

	xSurfaceHide(x.GoPointer())

}

var xSurfaceIsDestroyed func(uintptr) bool

// Check to see if a surface is destroyed.
func (x *Surface) IsDestroyed() bool {

	cret := xSurfaceIsDestroyed(x.GoPointer())
	return cret
}

var xSurfaceQueueRender func(uintptr)

// Forces a [signal@Gdk.Surface::render] signal emission for @surface
// to be scheduled.
//
// This function is useful for implementations that track invalid
// regions on their own.
func (x *Surface) QueueRender() {

	xSurfaceQueueRender(x.GoPointer())

}

var xSurfaceRequestLayout func(uintptr)

// Request a layout phase from the surface's frame clock.
//
// See [method@Gdk.FrameClock.request_phase].
func (x *Surface) RequestLayout() {

	xSurfaceRequestLayout(x.GoPointer())

}

var xSurfaceSetCursor func(uintptr, uintptr)

// Sets the default mouse pointer for a `GdkSurface`.
//
// Passing %NULL for the @cursor argument means that @surface will use
// the cursor of its parent surface. Most surfaces should use this default.
// Note that @cursor must be for the same display as @surface.
//
// Use [ctor@Gdk.Cursor.new_from_name] or [ctor@Gdk.Cursor.new_from_texture]
// to create the cursor. To make the cursor invisible, use %GDK_BLANK_CURSOR.
func (x *Surface) SetCursor(CursorVar *Cursor) {

	xSurfaceSetCursor(x.GoPointer(), CursorVar.GoPointer())

}

var xSurfaceSetDeviceCursor func(uintptr, uintptr, uintptr)

// Sets a specific `GdkCursor` for a given device when it gets inside @surface.
//
// Passing %NULL for the @cursor argument means that @surface will use the
// cursor of its parent surface. Most surfaces should use this default.
//
// Use [ctor@Gdk.Cursor.new_from_name] or [ctor@Gdk.Cursor.new_from_texture]
// to create the cursor. To make the cursor invisible, use %GDK_BLANK_CURSOR.
func (x *Surface) SetDeviceCursor(DeviceVar *Device, CursorVar *Cursor) {

	xSurfaceSetDeviceCursor(x.GoPointer(), DeviceVar.GoPointer(), CursorVar.GoPointer())

}

var xSurfaceSetInputRegion func(uintptr, *cairo.Region)

// Apply the region to the surface for the purpose of event
// handling.
//
// Mouse events which happen while the pointer position corresponds
// to an unset bit in the mask will be passed on the surface below
// @surface.
//
// An input region is typically used with RGBA surfaces. The alpha
// channel of the surface defines which pixels are invisible and
// allows for nicely antialiased borders, and the input region
// controls where the surface is “clickable”.
//
// Use [method@Gdk.Display.supports_input_shapes] to find out if
// a particular backend supports input regions.
func (x *Surface) SetInputRegion(RegionVar *cairo.Region) {

	xSurfaceSetInputRegion(x.GoPointer(), RegionVar)

}

var xSurfaceSetOpaqueRegion func(uintptr, *cairo.Region)

// Marks a region of the `GdkSurface` as opaque.
//
// For optimisation purposes, compositing window managers may
// like to not draw obscured regions of surfaces, or turn off blending
// during for these regions. With RGB windows with no transparency,
// this is just the shape of the window, but with ARGB32 windows, the
// compositor does not know what regions of the window are transparent
// or not.
//
// This function only works for toplevel surfaces.
//
// GTK will update this property automatically if the @surface background
// is opaque, as we know where the opaque regions are. If your surface
// background is not opaque, please update this property in your
// [vfunc@Gtk.Widget.css_changed] handler.
func (x *Surface) SetOpaqueRegion(RegionVar *cairo.Region) {

	xSurfaceSetOpaqueRegion(x.GoPointer(), RegionVar)

}

var xSurfaceTranslateCoordinates func(uintptr, uintptr, float64, float64) bool

// Translates coordinates between two surfaces.
//
// Note that this only works if @to and @from are popups or
// transient-for to the same toplevel (directly or indirectly).
func (x *Surface) TranslateCoordinates(ToVar *Surface, XVar float64, YVar float64) bool {

	cret := xSurfaceTranslateCoordinates(x.GoPointer(), ToVar.GoPointer(), XVar, YVar)
	return cret
}

func (c *Surface) GoPointer() uintptr {
	return c.Ptr
}

func (c *Surface) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// Emitted when @surface starts being present on the monitor.
func (x *Surface) ConnectEnterMonitor(cb func(Surface, uintptr)) {
	fcb := func(clsPtr uintptr, MonitorVarp uintptr) {
		fa := Surface{}
		fa.Ptr = clsPtr

		cb(fa, MonitorVarp)

	}
	gobject.ObjectConnect(x.GoPointer(), "signal::enter-monitor", purego.NewCallback(fcb))
}

// Emitted when GDK receives an input event for @surface.
func (x *Surface) ConnectEvent(cb func(Surface, *Event) bool) {
	fcb := func(clsPtr uintptr, EventVarp uintptr) bool {
		fa := Surface{}
		fa.Ptr = clsPtr

		return cb(fa, EventNewFromInternalPtr(EventVarp))

	}
	gobject.ObjectConnect(x.GoPointer(), "signal::event", purego.NewCallback(fcb))
}

// Emitted when the size of @surface is changed, or when relayout should
// be performed.
//
// Surface size is reported in ”application pixels”, not
// ”device pixels” (see gdk_surface_get_scale_factor()).
func (x *Surface) ConnectLayout(cb func(Surface, int, int)) {
	fcb := func(clsPtr uintptr, WidthVarp int, HeightVarp int) {
		fa := Surface{}
		fa.Ptr = clsPtr

		cb(fa, WidthVarp, HeightVarp)

	}
	gobject.ObjectConnect(x.GoPointer(), "signal::layout", purego.NewCallback(fcb))
}

// Emitted when @surface stops being present on the monitor.
func (x *Surface) ConnectLeaveMonitor(cb func(Surface, uintptr)) {
	fcb := func(clsPtr uintptr, MonitorVarp uintptr) {
		fa := Surface{}
		fa.Ptr = clsPtr

		cb(fa, MonitorVarp)

	}
	gobject.ObjectConnect(x.GoPointer(), "signal::leave-monitor", purego.NewCallback(fcb))
}

// Emitted when part of the surface needs to be redrawn.
func (x *Surface) ConnectRender(cb func(Surface, uintptr) bool) {
	fcb := func(clsPtr uintptr, RegionVarp uintptr) bool {
		fa := Surface{}
		fa.Ptr = clsPtr

		return cb(fa, RegionVarp)

	}
	gobject.ObjectConnect(x.GoPointer(), "signal::render", purego.NewCallback(fcb))
}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GDK"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xNewPopupSurface, lib, "gdk_surface_new_popup")
	core.PuregoSafeRegister(&xNewToplevelSurface, lib, "gdk_surface_new_toplevel")

	core.PuregoSafeRegister(&xSurfaceBeep, lib, "gdk_surface_beep")
	core.PuregoSafeRegister(&xSurfaceCreateCairoContext, lib, "gdk_surface_create_cairo_context")
	core.PuregoSafeRegister(&xSurfaceCreateGlContext, lib, "gdk_surface_create_gl_context")
	core.PuregoSafeRegister(&xSurfaceCreateSimilarSurface, lib, "gdk_surface_create_similar_surface")
	core.PuregoSafeRegister(&xSurfaceCreateVulkanContext, lib, "gdk_surface_create_vulkan_context")
	core.PuregoSafeRegister(&xSurfaceDestroy, lib, "gdk_surface_destroy")
	core.PuregoSafeRegister(&xSurfaceGetCursor, lib, "gdk_surface_get_cursor")
	core.PuregoSafeRegister(&xSurfaceGetDeviceCursor, lib, "gdk_surface_get_device_cursor")
	core.PuregoSafeRegister(&xSurfaceGetDevicePosition, lib, "gdk_surface_get_device_position")
	core.PuregoSafeRegister(&xSurfaceGetDisplay, lib, "gdk_surface_get_display")
	core.PuregoSafeRegister(&xSurfaceGetFrameClock, lib, "gdk_surface_get_frame_clock")
	core.PuregoSafeRegister(&xSurfaceGetHeight, lib, "gdk_surface_get_height")
	core.PuregoSafeRegister(&xSurfaceGetMapped, lib, "gdk_surface_get_mapped")
	core.PuregoSafeRegister(&xSurfaceGetScaleFactor, lib, "gdk_surface_get_scale_factor")
	core.PuregoSafeRegister(&xSurfaceGetWidth, lib, "gdk_surface_get_width")
	core.PuregoSafeRegister(&xSurfaceHide, lib, "gdk_surface_hide")
	core.PuregoSafeRegister(&xSurfaceIsDestroyed, lib, "gdk_surface_is_destroyed")
	core.PuregoSafeRegister(&xSurfaceQueueRender, lib, "gdk_surface_queue_render")
	core.PuregoSafeRegister(&xSurfaceRequestLayout, lib, "gdk_surface_request_layout")
	core.PuregoSafeRegister(&xSurfaceSetCursor, lib, "gdk_surface_set_cursor")
	core.PuregoSafeRegister(&xSurfaceSetDeviceCursor, lib, "gdk_surface_set_device_cursor")
	core.PuregoSafeRegister(&xSurfaceSetInputRegion, lib, "gdk_surface_set_input_region")
	core.PuregoSafeRegister(&xSurfaceSetOpaqueRegion, lib, "gdk_surface_set_opaque_region")
	core.PuregoSafeRegister(&xSurfaceTranslateCoordinates, lib, "gdk_surface_translate_coordinates")

}
