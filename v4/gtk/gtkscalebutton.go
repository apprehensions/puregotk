// Package gtk was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gtk

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/gobject"
)

type ScaleButtonClass struct {
	ParentClass uintptr

	Padding uintptr
}

// `GtkScaleButton` provides a button which pops up a scale widget.
//
// This kind of widget is commonly used for volume controls in multimedia
// applications, and GTK provides a [class@Gtk.VolumeButton] subclass that
// is tailored for this use case.
//
// # CSS nodes
//
// `GtkScaleButton` has a single CSS node with name button. To differentiate
// it from a plain `GtkButton`, it gets the .scale style class.
type ScaleButton struct {
	Widget
}

func ScaleButtonNewFromInternalPtr(ptr uintptr) *ScaleButton {
	cls := &ScaleButton{}
	cls.Ptr = ptr
	return cls
}

var xNewScaleButton func(float64, float64, float64, uintptr) uintptr

// Creates a `GtkScaleButton`.
//
// The new scale button has a range between @min and @max,
// with a stepping of @step.
func NewScaleButton(MinVar float64, MaxVar float64, StepVar float64, IconsVar uintptr) *Widget {
	NewScaleButtonPtr := xNewScaleButton(MinVar, MaxVar, StepVar, IconsVar)
	if NewScaleButtonPtr == 0 {
		return nil
	}

	gobject.IncreaseRef(NewScaleButtonPtr)

	NewScaleButtonCls := &Widget{}
	NewScaleButtonCls.Ptr = NewScaleButtonPtr
	return NewScaleButtonCls
}

var xScaleButtonGetAdjustment func(uintptr) uintptr

// Gets the `GtkAdjustment` associated with the `GtkScaleButton`’s scale.
//
// See [method@Gtk.Range.get_adjustment] for details.
func (x *ScaleButton) GetAdjustment() *Adjustment {

	GetAdjustmentPtr := xScaleButtonGetAdjustment(x.GoPointer())
	if GetAdjustmentPtr == 0 {
		return nil
	}

	gobject.IncreaseRef(GetAdjustmentPtr)

	GetAdjustmentCls := &Adjustment{}
	GetAdjustmentCls.Ptr = GetAdjustmentPtr
	return GetAdjustmentCls

}

var xScaleButtonGetMinusButton func(uintptr) uintptr

// Retrieves the minus button of the `GtkScaleButton`.
func (x *ScaleButton) GetMinusButton() *Button {

	GetMinusButtonPtr := xScaleButtonGetMinusButton(x.GoPointer())
	if GetMinusButtonPtr == 0 {
		return nil
	}

	gobject.IncreaseRef(GetMinusButtonPtr)

	GetMinusButtonCls := &Button{}
	GetMinusButtonCls.Ptr = GetMinusButtonPtr
	return GetMinusButtonCls

}

var xScaleButtonGetPlusButton func(uintptr) uintptr

// Retrieves the plus button of the `GtkScaleButton.`
func (x *ScaleButton) GetPlusButton() *Button {

	GetPlusButtonPtr := xScaleButtonGetPlusButton(x.GoPointer())
	if GetPlusButtonPtr == 0 {
		return nil
	}

	gobject.IncreaseRef(GetPlusButtonPtr)

	GetPlusButtonCls := &Button{}
	GetPlusButtonCls.Ptr = GetPlusButtonPtr
	return GetPlusButtonCls

}

var xScaleButtonGetPopup func(uintptr) uintptr

// Retrieves the popup of the `GtkScaleButton`.
func (x *ScaleButton) GetPopup() *Widget {

	GetPopupPtr := xScaleButtonGetPopup(x.GoPointer())
	if GetPopupPtr == 0 {
		return nil
	}

	gobject.IncreaseRef(GetPopupPtr)

	GetPopupCls := &Widget{}
	GetPopupCls.Ptr = GetPopupPtr
	return GetPopupCls

}

var xScaleButtonGetValue func(uintptr) float64

// Gets the current value of the scale button.
func (x *ScaleButton) GetValue() float64 {

	return xScaleButtonGetValue(x.GoPointer())

}

var xScaleButtonSetAdjustment func(uintptr, uintptr)

// Sets the `GtkAdjustment` to be used as a model
// for the `GtkScaleButton`’s scale.
//
// See [method@Gtk.Range.set_adjustment] for details.
func (x *ScaleButton) SetAdjustment(AdjustmentVar *Adjustment) {

	xScaleButtonSetAdjustment(x.GoPointer(), AdjustmentVar.GoPointer())

}

var xScaleButtonSetIcons func(uintptr, uintptr)

// Sets the icons to be used by the scale button.
func (x *ScaleButton) SetIcons(IconsVar uintptr) {

	xScaleButtonSetIcons(x.GoPointer(), IconsVar)

}

var xScaleButtonSetValue func(uintptr, float64)

// Sets the current value of the scale.
//
// If the value is outside the minimum or maximum range values,
// it will be clamped to fit inside them.
//
// The scale button emits the [signal@Gtk.ScaleButton::value-changed]
// signal if the value changes.
func (x *ScaleButton) SetValue(ValueVar float64) {

	xScaleButtonSetValue(x.GoPointer(), ValueVar)

}

func (c *ScaleButton) GoPointer() uintptr {
	return c.Ptr
}

func (c *ScaleButton) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// Emitted to dismiss the popup.
//
// This is a [keybinding signal](class.SignalAction.html).
//
// The default binding for this signal is &lt;kbd&gt;Escape&lt;/kbd&gt;.
func (x *ScaleButton) ConnectPopdown(cb func(ScaleButton)) {
	fcb := func(clsPtr uintptr) {
		fa := ScaleButton{}
		fa.Ptr = clsPtr

		cb(fa)

	}
	gobject.ObjectConnect(x.GoPointer(), "signal::popdown", purego.NewCallback(fcb))
}

// Emitted to popup the scale widget.
//
// This is a [keybinding signal](class.SignalAction.html).
//
// The default bindings for this signal are &lt;kbd&gt;Space&lt;/kbd&gt;,
// &lt;kbd&gt;Enter&lt;/kbd&gt; and &lt;kbd&gt;Return&lt;/kbd&gt;.
func (x *ScaleButton) ConnectPopup(cb func(ScaleButton)) {
	fcb := func(clsPtr uintptr) {
		fa := ScaleButton{}
		fa.Ptr = clsPtr

		cb(fa)

	}
	gobject.ObjectConnect(x.GoPointer(), "signal::popup", purego.NewCallback(fcb))
}

// Emitted when the value field has changed.
func (x *ScaleButton) ConnectValueChanged(cb func(ScaleButton, float64)) {
	fcb := func(clsPtr uintptr, ValueVarp float64) {
		fa := ScaleButton{}
		fa.Ptr = clsPtr

		cb(fa, ValueVarp)

	}
	gobject.ObjectConnect(x.GoPointer(), "signal::value-changed", purego.NewCallback(fcb))
}

// Retrieves the `GtkAccessibleRole` for the given `GtkAccessible`.
func (x *ScaleButton) GetAccessibleRole() AccessibleRole {

	return XGtkAccessibleGetAccessibleRole(x.GoPointer())

}

// Resets the accessible @property to its default value.
func (x *ScaleButton) ResetProperty(PropertyVar AccessibleProperty) {

	XGtkAccessibleResetProperty(x.GoPointer(), PropertyVar)

}

// Resets the accessible @relation to its default value.
func (x *ScaleButton) ResetRelation(RelationVar AccessibleRelation) {

	XGtkAccessibleResetRelation(x.GoPointer(), RelationVar)

}

// Resets the accessible @state to its default value.
func (x *ScaleButton) ResetState(StateVar AccessibleState) {

	XGtkAccessibleResetState(x.GoPointer(), StateVar)

}

// Updates a list of accessible properties.
//
// See the [enum@Gtk.AccessibleProperty] documentation for the
// value types of accessible properties.
//
// This function should be called by `GtkWidget` types whenever
// an accessible property change must be communicated to assistive
// technologies.
//
// Example:
// ```c
// value = gtk_adjustment_get_value (adjustment);
// gtk_accessible_update_property (GTK_ACCESSIBLE (spin_button),
//
//	GTK_ACCESSIBLE_PROPERTY_VALUE_NOW, value,
//	-1);
//
// ```
func (x *ScaleButton) UpdateProperty(FirstPropertyVar AccessibleProperty, varArgs ...interface{}) {

	XGtkAccessibleUpdateProperty(x.GoPointer(), FirstPropertyVar, varArgs...)

}

// Updates an array of accessible properties.
//
// This function should be called by `GtkWidget` types whenever an accessible
// property change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *ScaleButton) UpdatePropertyValue(NPropertiesVar int32, PropertiesVar uintptr, ValuesVar uintptr) {

	XGtkAccessibleUpdatePropertyValue(x.GoPointer(), NPropertiesVar, PropertiesVar, ValuesVar)

}

// Updates a list of accessible relations.
//
// This function should be called by `GtkWidget` types whenever an accessible
// relation change must be communicated to assistive technologies.
//
// If the [enum@Gtk.AccessibleRelation] requires a list of references,
// you should pass each reference individually, followed by %NULL, e.g.
//
// ```c
// gtk_accessible_update_relation (accessible,
//
//	GTK_ACCESSIBLE_RELATION_CONTROLS,
//	  ref1, NULL,
//	GTK_ACCESSIBLE_RELATION_LABELLED_BY,
//	  ref1, ref2, ref3, NULL,
//	-1);
//
// ```
func (x *ScaleButton) UpdateRelation(FirstRelationVar AccessibleRelation, varArgs ...interface{}) {

	XGtkAccessibleUpdateRelation(x.GoPointer(), FirstRelationVar, varArgs...)

}

// Updates an array of accessible relations.
//
// This function should be called by `GtkWidget` types whenever an accessible
// relation change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *ScaleButton) UpdateRelationValue(NRelationsVar int32, RelationsVar uintptr, ValuesVar uintptr) {

	XGtkAccessibleUpdateRelationValue(x.GoPointer(), NRelationsVar, RelationsVar, ValuesVar)

}

// Updates a list of accessible states. See the [enum@Gtk.AccessibleState]
// documentation for the value types of accessible states.
//
// This function should be called by `GtkWidget` types whenever an accessible
// state change must be communicated to assistive technologies.
//
// Example:
// ```c
// value = GTK_ACCESSIBLE_TRISTATE_MIXED;
// gtk_accessible_update_state (GTK_ACCESSIBLE (check_button),
//
//	GTK_ACCESSIBLE_STATE_CHECKED, value,
//	-1);
//
// ```
func (x *ScaleButton) UpdateState(FirstStateVar AccessibleState, varArgs ...interface{}) {

	XGtkAccessibleUpdateState(x.GoPointer(), FirstStateVar, varArgs...)

}

// Updates an array of accessible states.
//
// This function should be called by `GtkWidget` types whenever an accessible
// state change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *ScaleButton) UpdateStateValue(NStatesVar int32, StatesVar uintptr, ValuesVar uintptr) {

	XGtkAccessibleUpdateStateValue(x.GoPointer(), NStatesVar, StatesVar, ValuesVar)

}

// Gets the ID of the @buildable object.
//
// `GtkBuilder` sets the name based on the ID attribute
// of the &lt;object&gt; tag used to construct the @buildable.
func (x *ScaleButton) GetBuildableId() string {

	return XGtkBuildableGetBuildableId(x.GoPointer())

}

// Retrieves the orientation of the @orientable.
func (x *ScaleButton) GetOrientation() Orientation {

	return XGtkOrientableGetOrientation(x.GoPointer())

}

// Sets the orientation of the @orientable.
func (x *ScaleButton) SetOrientation(OrientationVar Orientation) {

	XGtkOrientableSetOrientation(x.GoPointer(), OrientationVar)

}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GTK"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xNewScaleButton, lib, "gtk_scale_button_new")

	core.PuregoSafeRegister(&xScaleButtonGetAdjustment, lib, "gtk_scale_button_get_adjustment")
	core.PuregoSafeRegister(&xScaleButtonGetMinusButton, lib, "gtk_scale_button_get_minus_button")
	core.PuregoSafeRegister(&xScaleButtonGetPlusButton, lib, "gtk_scale_button_get_plus_button")
	core.PuregoSafeRegister(&xScaleButtonGetPopup, lib, "gtk_scale_button_get_popup")
	core.PuregoSafeRegister(&xScaleButtonGetValue, lib, "gtk_scale_button_get_value")
	core.PuregoSafeRegister(&xScaleButtonSetAdjustment, lib, "gtk_scale_button_set_adjustment")
	core.PuregoSafeRegister(&xScaleButtonSetIcons, lib, "gtk_scale_button_set_icons")
	core.PuregoSafeRegister(&xScaleButtonSetValue, lib, "gtk_scale_button_set_value")

}
