// Package gtk was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gtk

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/gobject"
)

type AdjustmentClass struct {
	ParentClass uintptr
}

// `GtkAdjustment` is a model for a numeric value.
//
// The `GtkAdjustment has an associated lower and upper bound.
// It also contains step and page increments, and a page size.
//
// Adjustments are used within several GTK widgets, including
// [class@Gtk.SpinButton], [class@Gtk.Viewport], [class@Gtk.Scrollbar]
// and [class@Gtk.Scale].
//
// The `GtkAdjustment` object does not update the value itself. Instead
// it is left up to the owner of the `GtkAdjustment` to control the value.
type Adjustment struct {
	gobject.InitiallyUnowned
}

func AdjustmentNewFromInternalPtr(ptr uintptr) *Adjustment {
	cls := &Adjustment{}
	cls.Ptr = ptr
	return cls
}

var xNewAdjustment func(float64, float64, float64, float64, float64, float64) uintptr

// Creates a new `GtkAdjustment`.
func NewAdjustment(ValueVar float64, LowerVar float64, UpperVar float64, StepIncrementVar float64, PageIncrementVar float64, PageSizeVar float64) *Adjustment {
	NewAdjustmentPtr := xNewAdjustment(ValueVar, LowerVar, UpperVar, StepIncrementVar, PageIncrementVar, PageSizeVar)
	if NewAdjustmentPtr == 0 {
		return nil
	}

	gobject.IncreaseRef(NewAdjustmentPtr)

	NewAdjustmentCls := &Adjustment{}
	NewAdjustmentCls.Ptr = NewAdjustmentPtr
	return NewAdjustmentCls
}

var xAdjustmentClampPage func(uintptr, float64, float64)

// Updates the value property to ensure that the range
// between @lower and @upper is in the current page.
//
// The current page goes from `value` to `value` + `page-size`.
// If the range is larger than the page size, then only the
// start of it will be in the current page.
//
// A [signal@Gtk.Adjustment::value-changed] signal will be emitted
// if the value is changed.
func (x *Adjustment) ClampPage(LowerVar float64, UpperVar float64) {

	xAdjustmentClampPage(x.GoPointer(), LowerVar, UpperVar)

}

var xAdjustmentConfigure func(uintptr, float64, float64, float64, float64, float64, float64)

// Sets all properties of the adjustment at once.
//
// Use this function to avoid multiple emissions of the
// [signal@Gtk.Adjustment::changed] signal. See
// [method@Gtk.Adjustment.set_lower] for an alternative
// way of compressing multiple emissions of
// [signal@Gtk.Adjustment::changed] into one.
func (x *Adjustment) Configure(ValueVar float64, LowerVar float64, UpperVar float64, StepIncrementVar float64, PageIncrementVar float64, PageSizeVar float64) {

	xAdjustmentConfigure(x.GoPointer(), ValueVar, LowerVar, UpperVar, StepIncrementVar, PageIncrementVar, PageSizeVar)

}

var xAdjustmentGetLower func(uintptr) float64

// Retrieves the minimum value of the adjustment.
func (x *Adjustment) GetLower() float64 {

	return xAdjustmentGetLower(x.GoPointer())

}

var xAdjustmentGetMinimumIncrement func(uintptr) float64

// Gets the smaller of step increment and page increment.
func (x *Adjustment) GetMinimumIncrement() float64 {

	return xAdjustmentGetMinimumIncrement(x.GoPointer())

}

var xAdjustmentGetPageIncrement func(uintptr) float64

// Retrieves the page increment of the adjustment.
func (x *Adjustment) GetPageIncrement() float64 {

	return xAdjustmentGetPageIncrement(x.GoPointer())

}

var xAdjustmentGetPageSize func(uintptr) float64

// Retrieves the page size of the adjustment.
func (x *Adjustment) GetPageSize() float64 {

	return xAdjustmentGetPageSize(x.GoPointer())

}

var xAdjustmentGetStepIncrement func(uintptr) float64

// Retrieves the step increment of the adjustment.
func (x *Adjustment) GetStepIncrement() float64 {

	return xAdjustmentGetStepIncrement(x.GoPointer())

}

var xAdjustmentGetUpper func(uintptr) float64

// Retrieves the maximum value of the adjustment.
func (x *Adjustment) GetUpper() float64 {

	return xAdjustmentGetUpper(x.GoPointer())

}

var xAdjustmentGetValue func(uintptr) float64

// Gets the current value of the adjustment.
func (x *Adjustment) GetValue() float64 {

	return xAdjustmentGetValue(x.GoPointer())

}

var xAdjustmentSetLower func(uintptr, float64)

// Sets the minimum value of the adjustment.
//
// When setting multiple adjustment properties via their individual
// setters, multiple [signal@Gtk.Adjustment::changed] signals will
// be emitted. However, since the emission of the
// [signal@Gtk.Adjustment::changed] signal is tied to the emission
// of the ::notify signals of the changed properties, it’s possible
// to compress the [signal@Gtk.Adjustment::changed] signals into one
// by calling g_object_freeze_notify() and g_object_thaw_notify()
// around the calls to the individual setters.
//
// Alternatively, using a single g_object_set() for all the properties
// to change, or using [method@Gtk.Adjustment.configure] has the same effect.
func (x *Adjustment) SetLower(LowerVar float64) {

	xAdjustmentSetLower(x.GoPointer(), LowerVar)

}

var xAdjustmentSetPageIncrement func(uintptr, float64)

// Sets the page increment of the adjustment.
//
// See [method@Gtk.Adjustment.set_lower] about how to compress
// multiple emissions of the [signal@Gtk.Adjustment::changed]
// signal when setting multiple adjustment properties.
func (x *Adjustment) SetPageIncrement(PageIncrementVar float64) {

	xAdjustmentSetPageIncrement(x.GoPointer(), PageIncrementVar)

}

var xAdjustmentSetPageSize func(uintptr, float64)

// Sets the page size of the adjustment.
//
// See [method@Gtk.Adjustment.set_lower] about how to compress
// multiple emissions of the [signal@Gtk.Adjustment::changed]
// signal when setting multiple adjustment properties.
func (x *Adjustment) SetPageSize(PageSizeVar float64) {

	xAdjustmentSetPageSize(x.GoPointer(), PageSizeVar)

}

var xAdjustmentSetStepIncrement func(uintptr, float64)

// Sets the step increment of the adjustment.
//
// See [method@Gtk.Adjustment.set_lower] about how to compress
// multiple emissions of the [signal@Gtk.Adjustment::changed]
// signal when setting multiple adjustment properties.
func (x *Adjustment) SetStepIncrement(StepIncrementVar float64) {

	xAdjustmentSetStepIncrement(x.GoPointer(), StepIncrementVar)

}

var xAdjustmentSetUpper func(uintptr, float64)

// Sets the maximum value of the adjustment.
//
// Note that values will be restricted by `upper - page-size`
// if the page-size property is nonzero.
//
// See [method@Gtk.Adjustment.set_lower] about how to compress
// multiple emissions of the [signal@Gtk.Adjustment::changed]
// signal when setting multiple adjustment properties.
func (x *Adjustment) SetUpper(UpperVar float64) {

	xAdjustmentSetUpper(x.GoPointer(), UpperVar)

}

var xAdjustmentSetValue func(uintptr, float64)

// Sets the `GtkAdjustment` value.
//
// The value is clamped to lie between [property@Gtk.Adjustment:lower]
// and [property@Gtk.Adjustment:upper].
//
// Note that for adjustments which are used in a `GtkScrollbar`,
// the effective range of allowed values goes from
// [property@Gtk.Adjustment:lower] to
// [property@Gtk.Adjustment:upper] - [property@Gtk.Adjustment:page-size].
func (x *Adjustment) SetValue(ValueVar float64) {

	xAdjustmentSetValue(x.GoPointer(), ValueVar)

}

func (c *Adjustment) GoPointer() uintptr {
	return c.Ptr
}

func (c *Adjustment) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// Emitted when one or more of the `GtkAdjustment` properties have been
// changed.
//
// Note that the [property@Gtk.Adjustment:value] property is
// covered by the [signal@Gtk.Adjustment::value-changed] signal.
func (x *Adjustment) ConnectChanged(cb func(Adjustment)) {
	fcb := func(clsPtr uintptr) {
		fa := Adjustment{}
		fa.Ptr = clsPtr

		cb(fa)

	}
	gobject.ObjectConnect(x.GoPointer(), "signal::changed", purego.NewCallback(fcb))
}

// Emitted when the value has been changed.
func (x *Adjustment) ConnectValueChanged(cb func(Adjustment)) {
	fcb := func(clsPtr uintptr) {
		fa := Adjustment{}
		fa.Ptr = clsPtr

		cb(fa)

	}
	gobject.ObjectConnect(x.GoPointer(), "signal::value-changed", purego.NewCallback(fcb))
}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GTK"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xNewAdjustment, lib, "gtk_adjustment_new")

	core.PuregoSafeRegister(&xAdjustmentClampPage, lib, "gtk_adjustment_clamp_page")
	core.PuregoSafeRegister(&xAdjustmentConfigure, lib, "gtk_adjustment_configure")
	core.PuregoSafeRegister(&xAdjustmentGetLower, lib, "gtk_adjustment_get_lower")
	core.PuregoSafeRegister(&xAdjustmentGetMinimumIncrement, lib, "gtk_adjustment_get_minimum_increment")
	core.PuregoSafeRegister(&xAdjustmentGetPageIncrement, lib, "gtk_adjustment_get_page_increment")
	core.PuregoSafeRegister(&xAdjustmentGetPageSize, lib, "gtk_adjustment_get_page_size")
	core.PuregoSafeRegister(&xAdjustmentGetStepIncrement, lib, "gtk_adjustment_get_step_increment")
	core.PuregoSafeRegister(&xAdjustmentGetUpper, lib, "gtk_adjustment_get_upper")
	core.PuregoSafeRegister(&xAdjustmentGetValue, lib, "gtk_adjustment_get_value")
	core.PuregoSafeRegister(&xAdjustmentSetLower, lib, "gtk_adjustment_set_lower")
	core.PuregoSafeRegister(&xAdjustmentSetPageIncrement, lib, "gtk_adjustment_set_page_increment")
	core.PuregoSafeRegister(&xAdjustmentSetPageSize, lib, "gtk_adjustment_set_page_size")
	core.PuregoSafeRegister(&xAdjustmentSetStepIncrement, lib, "gtk_adjustment_set_step_increment")
	core.PuregoSafeRegister(&xAdjustmentSetUpper, lib, "gtk_adjustment_set_upper")
	core.PuregoSafeRegister(&xAdjustmentSetValue, lib, "gtk_adjustment_set_value")

}
