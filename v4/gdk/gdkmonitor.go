// Package gdk was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gdk

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/gobject"
)

type MonitorClass struct {
}

// This enumeration describes how the red, green and blue components
// of physical pixels on an output device are laid out.
type SubpixelLayout int

const (

	// The layout is not known
	SubpixelLayoutUnknownValue SubpixelLayout = 0
	// Not organized in this way
	SubpixelLayoutNoneValue SubpixelLayout = 1
	// The layout is horizontal, the order is RGB
	SubpixelLayoutHorizontalRgbValue SubpixelLayout = 2
	// The layout is horizontal, the order is BGR
	SubpixelLayoutHorizontalBgrValue SubpixelLayout = 3
	// The layout is vertical, the order is RGB
	SubpixelLayoutVerticalRgbValue SubpixelLayout = 4
	// The layout is vertical, the order is BGR
	SubpixelLayoutVerticalBgrValue SubpixelLayout = 5
)

// `GdkMonitor` objects represent the individual outputs that are
// associated with a `GdkDisplay`.
//
// `GdkDisplay` keeps a `GListModel` to enumerate and monitor
// monitors with [method@Gdk.Display.get_monitors]. You can use
// [method@Gdk.Display.get_monitor_at_surface] to find a particular
// monitor.
type Monitor struct {
	gobject.Object
}

func MonitorNewFromInternalPtr(ptr uintptr) *Monitor {
	cls := &Monitor{}
	cls.Ptr = ptr
	return cls
}

var xMonitorGetConnector func(uintptr) string

// Gets the name of the monitor's connector, if available.
func (x *Monitor) GetConnector() string {

	cret := xMonitorGetConnector(x.GoPointer())
	return cret
}

var xMonitorGetDisplay func(uintptr) uintptr

// Gets the display that this monitor belongs to.
func (x *Monitor) GetDisplay() *Display {
	var cls *Display

	cret := xMonitorGetDisplay(x.GoPointer())

	if cret == 0 {
		return cls
	}
	gobject.IncreaseRef(cret)
	cls = &Display{}
	cls.Ptr = cret
	return cls
}

var xMonitorGetGeometry func(uintptr, *Rectangle)

// Retrieves the size and position of the monitor within the
// display coordinate space.
//
// The returned geometry is in  ”application pixels”, not in
// ”device pixels” (see [method@Gdk.Monitor.get_scale_factor]).
func (x *Monitor) GetGeometry(GeometryVar *Rectangle) {

	xMonitorGetGeometry(x.GoPointer(), GeometryVar)

}

var xMonitorGetHeightMm func(uintptr) int

// Gets the height in millimeters of the monitor.
func (x *Monitor) GetHeightMm() int {

	cret := xMonitorGetHeightMm(x.GoPointer())
	return cret
}

var xMonitorGetManufacturer func(uintptr) string

// Gets the name or PNP ID of the monitor's manufacturer.
//
// Note that this value might also vary depending on actual
// display backend.
//
// The PNP ID registry is located at
// [https://uefi.org/pnp_id_list](https://uefi.org/pnp_id_list).
func (x *Monitor) GetManufacturer() string {

	cret := xMonitorGetManufacturer(x.GoPointer())
	return cret
}

var xMonitorGetModel func(uintptr) string

// Gets the string identifying the monitor model, if available.
func (x *Monitor) GetModel() string {

	cret := xMonitorGetModel(x.GoPointer())
	return cret
}

var xMonitorGetRefreshRate func(uintptr) int

// Gets the refresh rate of the monitor, if available.
//
// The value is in milli-Hertz, so a refresh rate of 60Hz
// is returned as 60000.
func (x *Monitor) GetRefreshRate() int {

	cret := xMonitorGetRefreshRate(x.GoPointer())
	return cret
}

var xMonitorGetScaleFactor func(uintptr) int

// Gets the internal scale factor that maps from monitor coordinates
// to device pixels.
//
// On traditional systems this is 1, but on very high density outputs
// it can be a higher value (often 2).
//
// This can be used if you want to create pixel based data for a
// particular monitor, but most of the time you’re drawing to a surface
// where it is better to use [method@Gdk.Surface.get_scale_factor] instead.
func (x *Monitor) GetScaleFactor() int {

	cret := xMonitorGetScaleFactor(x.GoPointer())
	return cret
}

var xMonitorGetSubpixelLayout func(uintptr) SubpixelLayout

// Gets information about the layout of red, green and blue
// primaries for pixels.
func (x *Monitor) GetSubpixelLayout() SubpixelLayout {

	cret := xMonitorGetSubpixelLayout(x.GoPointer())
	return cret
}

var xMonitorGetWidthMm func(uintptr) int

// Gets the width in millimeters of the monitor.
func (x *Monitor) GetWidthMm() int {

	cret := xMonitorGetWidthMm(x.GoPointer())
	return cret
}

var xMonitorIsValid func(uintptr) bool

// Returns %TRUE if the @monitor object corresponds to a
// physical monitor.
//
// The @monitor becomes invalid when the physical monitor
// is unplugged or removed.
func (x *Monitor) IsValid() bool {

	cret := xMonitorIsValid(x.GoPointer())
	return cret
}

func (c *Monitor) GoPointer() uintptr {
	return c.Ptr
}

func (c *Monitor) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// Emitted when the output represented by @monitor gets disconnected.
func (x *Monitor) ConnectInvalidate(cb func(Monitor)) {
	fcb := func(clsPtr uintptr) {
		fa := Monitor{}
		fa.Ptr = clsPtr

		cb(fa)

	}
	gobject.ObjectConnect(x.GoPointer(), "signal::invalidate", purego.NewCallback(fcb))
}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GDK"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xMonitorGetConnector, lib, "gdk_monitor_get_connector")
	core.PuregoSafeRegister(&xMonitorGetDisplay, lib, "gdk_monitor_get_display")
	core.PuregoSafeRegister(&xMonitorGetGeometry, lib, "gdk_monitor_get_geometry")
	core.PuregoSafeRegister(&xMonitorGetHeightMm, lib, "gdk_monitor_get_height_mm")
	core.PuregoSafeRegister(&xMonitorGetManufacturer, lib, "gdk_monitor_get_manufacturer")
	core.PuregoSafeRegister(&xMonitorGetModel, lib, "gdk_monitor_get_model")
	core.PuregoSafeRegister(&xMonitorGetRefreshRate, lib, "gdk_monitor_get_refresh_rate")
	core.PuregoSafeRegister(&xMonitorGetScaleFactor, lib, "gdk_monitor_get_scale_factor")
	core.PuregoSafeRegister(&xMonitorGetSubpixelLayout, lib, "gdk_monitor_get_subpixel_layout")
	core.PuregoSafeRegister(&xMonitorGetWidthMm, lib, "gdk_monitor_get_width_mm")
	core.PuregoSafeRegister(&xMonitorIsValid, lib, "gdk_monitor_is_valid")

}
