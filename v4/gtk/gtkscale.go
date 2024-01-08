// Package gtk was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gtk

import (
	"unsafe"

	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/glib"
	"github.com/jwijenbergh/puregotk/v4/gobject"
	"github.com/jwijenbergh/puregotk/v4/pango"
)

type ScaleFormatValueFunc func(uintptr, float64, uintptr) string

type ScaleClass struct {
	ParentClass uintptr

	Padding uintptr
}

func (x *ScaleClass) GoPointer() uintptr {
	return uintptr(unsafe.Pointer(x))
}

// A `GtkScale` is a slider control used to select a numeric value.
//
// ![An example GtkScale](scales.png)
//
// To use it, you’ll probably want to investigate the methods on its base
// class, [class@GtkRange], in addition to the methods for `GtkScale` itself.
// To set the value of a scale, you would normally use [method@Gtk.Range.set_value].
// To detect changes to the value, you would normally use the
// [signal@Gtk.Range::value-changed] signal.
//
// Note that using the same upper and lower bounds for the `GtkScale` (through
// the `GtkRange` methods) will hide the slider itself. This is useful for
// applications that want to show an undeterminate value on the scale, without
// changing the layout of the application (such as movie or music players).
//
// # GtkScale as GtkBuildable
//
// `GtkScale` supports a custom &lt;marks&gt; element, which can contain multiple
// &lt;mark\&gt; elements. The “value” and “position” attributes have the same
// meaning as [method@Gtk.Scale.add_mark] parameters of the same name. If
// the element is not empty, its content is taken as the markup to show at
// the mark. It can be translated with the usual ”translatable” and
// “context” attributes.
//
// # CSS nodes
//
// ```
// scale[.fine-tune][.marks-before][.marks-after]
// ├── [value][.top][.right][.bottom][.left]
// ├── marks.top
// │   ├── mark
// │   ┊    ├── [label]
// │   ┊    ╰── indicator
// ┊   ┊
// │   ╰── mark
// ├── marks.bottom
// │   ├── mark
// │   ┊    ├── indicator
// │   ┊    ╰── [label]
// ┊   ┊
// │   ╰── mark
// ╰── trough
//
//	├── [fill]
//	├── [highlight]
//	╰── slider
//
// ```
//
// `GtkScale` has a main CSS node with name scale and a subnode for its contents,
// with subnodes named trough and slider.
//
// The main node gets the style class .fine-tune added when the scale is in
// 'fine-tuning' mode.
//
// If the scale has an origin (see [method@Gtk.Scale.set_has_origin]), there is
// a subnode with name highlight below the trough node that is used for rendering
// the highlighted part of the trough.
//
// If the scale is showing a fill level (see [method@Gtk.Range.set_show_fill_level]),
// there is a subnode with name fill below the trough node that is used for
// rendering the filled in part of the trough.
//
// If marks are present, there is a marks subnode before or after the trough
// node, below which each mark gets a node with name mark. The marks nodes get
// either the .top or .bottom style class.
//
// The mark node has a subnode named indicator. If the mark has text, it also
// has a subnode named label. When the mark is either above or left of the
// scale, the label subnode is the first when present. Otherwise, the indicator
// subnode is the first.
//
// The main CSS node gets the 'marks-before' and/or 'marks-after' style classes
// added depending on what marks are present.
//
// If the scale is displaying the value (see [property@Gtk.Scale:draw-value]),
// there is subnode with name value. This node will get the .top or .bottom style
// classes similar to the marks node.
//
// # Accessibility
//
// `GtkScale` uses the %GTK_ACCESSIBLE_ROLE_SLIDER role.
type Scale struct {
	Range
}

func ScaleNewFromInternalPtr(ptr uintptr) *Scale {
	cls := &Scale{}
	cls.Ptr = ptr
	return cls
}

var xNewScale func(Orientation, uintptr) uintptr

// Creates a new `GtkScale`.
func NewScale(OrientationVar Orientation, AdjustmentVar *Adjustment) *Widget {
	var cls *Widget

	cret := xNewScale(OrientationVar, AdjustmentVar.GoPointer())

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &Widget{}
	cls.Ptr = cret
	return cls
}

var xNewWithRangeScale func(Orientation, float64, float64, float64) uintptr

// Creates a new scale widget with a range from @min to @max.
//
// The returns scale will have the given orientation and will let the
// user input a number between @min and @max (including @min and @max)
// with the increment @step. @step must be nonzero; it’s the distance
// the slider moves when using the arrow keys to adjust the scale
// value.
//
// Note that the way in which the precision is derived works best if
// @step is a power of ten. If the resulting precision is not suitable
// for your needs, use [method@Gtk.Scale.set_digits] to correct it.
func NewWithRangeScale(OrientationVar Orientation, MinVar float64, MaxVar float64, StepVar float64) *Widget {
	var cls *Widget

	cret := xNewWithRangeScale(OrientationVar, MinVar, MaxVar, StepVar)

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &Widget{}
	cls.Ptr = cret
	return cls
}

var xScaleAddMark func(uintptr, float64, PositionType, string)

// Adds a mark at @value.
//
// A mark is indicated visually by drawing a tick mark next to the scale,
// and GTK makes it easy for the user to position the scale exactly at the
// marks value.
//
// If @markup is not %NULL, text is shown next to the tick mark.
//
// To remove marks from a scale, use [method@Gtk.Scale.clear_marks].
func (x *Scale) AddMark(ValueVar float64, PositionVar PositionType, MarkupVar string) {

	xScaleAddMark(x.GoPointer(), ValueVar, PositionVar, MarkupVar)

}

var xScaleClearMarks func(uintptr)

// Removes any marks that have been added.
func (x *Scale) ClearMarks() {

	xScaleClearMarks(x.GoPointer())

}

var xScaleGetDigits func(uintptr) int

// Gets the number of decimal places that are displayed in the value.
func (x *Scale) GetDigits() int {

	cret := xScaleGetDigits(x.GoPointer())
	return cret
}

var xScaleGetDrawValue func(uintptr) bool

// Returns whether the current value is displayed as a string
// next to the slider.
func (x *Scale) GetDrawValue() bool {

	cret := xScaleGetDrawValue(x.GoPointer())
	return cret
}

var xScaleGetHasOrigin func(uintptr) bool

// Returns whether the scale has an origin.
func (x *Scale) GetHasOrigin() bool {

	cret := xScaleGetHasOrigin(x.GoPointer())
	return cret
}

var xScaleGetLayout func(uintptr) uintptr

// Gets the `PangoLayout` used to display the scale.
//
// The returned object is owned by the scale so does not need
// to be freed by the caller.
func (x *Scale) GetLayout() *pango.Layout {
	var cls *pango.Layout

	cret := xScaleGetLayout(x.GoPointer())

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &pango.Layout{}
	cls.Ptr = cret
	return cls
}

var xScaleGetLayoutOffsets func(uintptr, int, int)

// Obtains the coordinates where the scale will draw the
// `PangoLayout` representing the text in the scale.
//
// Remember when using the `PangoLayout` function you need to
// convert to and from pixels using `PANGO_PIXELS()` or `PANGO_SCALE`.
//
// If the [property@GtkScale:draw-value] property is %FALSE, the return
// values are undefined.
func (x *Scale) GetLayoutOffsets(XVar int, YVar int) {

	xScaleGetLayoutOffsets(x.GoPointer(), XVar, YVar)

}

var xScaleGetValuePos func(uintptr) PositionType

// Gets the position in which the current value is displayed.
func (x *Scale) GetValuePos() PositionType {

	cret := xScaleGetValuePos(x.GoPointer())
	return cret
}

var xScaleSetDigits func(uintptr, int)

// Sets the number of decimal places that are displayed in the value.
//
// Also causes the value of the adjustment to be rounded to this number
// of digits, so the retrieved value matches the displayed one, if
// [property@GtkScale:draw-value] is %TRUE when the value changes. If
// you want to enforce rounding the value when [property@GtkScale:draw-value]
// is %FALSE, you can set [property@GtkRange:round-digits] instead.
//
// Note that rounding to a small number of digits can interfere with
// the smooth autoscrolling that is built into `GtkScale`. As an alternative,
// you can use [method@Gtk.Scale.set_format_value_func] to format the displayed
// value yourself.
func (x *Scale) SetDigits(DigitsVar int) {

	xScaleSetDigits(x.GoPointer(), DigitsVar)

}

var xScaleSetDrawValue func(uintptr, bool)

// Specifies whether the current value is displayed as a string next
// to the slider.
func (x *Scale) SetDrawValue(DrawValueVar bool) {

	xScaleSetDrawValue(x.GoPointer(), DrawValueVar)

}

var xScaleSetFormatValueFunc func(uintptr, uintptr, uintptr, uintptr)

// @func allows you to change how the scale value is displayed.
//
// The given function will return an allocated string representing
// @value. That string will then be used to display the scale's value.
//
// If #NULL is passed as @func, the value will be displayed on
// its own, rounded according to the value of the
// [property@GtkScale:digits] property.
func (x *Scale) SetFormatValueFunc(FuncVar ScaleFormatValueFunc, UserDataVar uintptr, DestroyNotifyVar glib.DestroyNotify) {

	xScaleSetFormatValueFunc(x.GoPointer(), purego.NewCallback(FuncVar), UserDataVar, purego.NewCallback(DestroyNotifyVar))

}

var xScaleSetHasOrigin func(uintptr, bool)

// Sets whether the scale has an origin.
//
// If [property@GtkScale:has-origin] is set to %TRUE (the default),
// the scale will highlight the part of the trough between the origin
// (bottom or left side) and the current value.
func (x *Scale) SetHasOrigin(HasOriginVar bool) {

	xScaleSetHasOrigin(x.GoPointer(), HasOriginVar)

}

var xScaleSetValuePos func(uintptr, PositionType)

// Sets the position in which the current value is displayed.
func (x *Scale) SetValuePos(PosVar PositionType) {

	xScaleSetValuePos(x.GoPointer(), PosVar)

}

func (c *Scale) GoPointer() uintptr {
	return c.Ptr
}

func (c *Scale) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// Retrieves the `GtkAccessibleRole` for the given `GtkAccessible`.
func (x *Scale) GetAccessibleRole() AccessibleRole {

	cret := XGtkAccessibleGetAccessibleRole(x.GoPointer())
	return cret
}

// Resets the accessible @property to its default value.
func (x *Scale) ResetProperty(PropertyVar AccessibleProperty) {

	XGtkAccessibleResetProperty(x.GoPointer(), PropertyVar)

}

// Resets the accessible @relation to its default value.
func (x *Scale) ResetRelation(RelationVar AccessibleRelation) {

	XGtkAccessibleResetRelation(x.GoPointer(), RelationVar)

}

// Resets the accessible @state to its default value.
func (x *Scale) ResetState(StateVar AccessibleState) {

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
func (x *Scale) UpdateProperty(FirstPropertyVar AccessibleProperty, varArgs ...interface{}) {

	XGtkAccessibleUpdateProperty(x.GoPointer(), FirstPropertyVar, varArgs...)

}

// Updates an array of accessible properties.
//
// This function should be called by `GtkWidget` types whenever an accessible
// property change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *Scale) UpdatePropertyValue(NPropertiesVar int, PropertiesVar uintptr, ValuesVar uintptr) {

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
func (x *Scale) UpdateRelation(FirstRelationVar AccessibleRelation, varArgs ...interface{}) {

	XGtkAccessibleUpdateRelation(x.GoPointer(), FirstRelationVar, varArgs...)

}

// Updates an array of accessible relations.
//
// This function should be called by `GtkWidget` types whenever an accessible
// relation change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *Scale) UpdateRelationValue(NRelationsVar int, RelationsVar uintptr, ValuesVar uintptr) {

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
func (x *Scale) UpdateState(FirstStateVar AccessibleState, varArgs ...interface{}) {

	XGtkAccessibleUpdateState(x.GoPointer(), FirstStateVar, varArgs...)

}

// Updates an array of accessible states.
//
// This function should be called by `GtkWidget` types whenever an accessible
// state change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *Scale) UpdateStateValue(NStatesVar int, StatesVar uintptr, ValuesVar uintptr) {

	XGtkAccessibleUpdateStateValue(x.GoPointer(), NStatesVar, StatesVar, ValuesVar)

}

// Gets the ID of the @buildable object.
//
// `GtkBuilder` sets the name based on the ID attribute
// of the &lt;object&gt; tag used to construct the @buildable.
func (x *Scale) GetBuildableId() string {

	cret := XGtkBuildableGetBuildableId(x.GoPointer())
	return cret
}

// Retrieves the orientation of the @orientable.
func (x *Scale) GetOrientation() Orientation {

	cret := XGtkOrientableGetOrientation(x.GoPointer())
	return cret
}

// Sets the orientation of the @orientable.
func (x *Scale) SetOrientation(OrientationVar Orientation) {

	XGtkOrientableSetOrientation(x.GoPointer(), OrientationVar)

}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GTK"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xNewScale, lib, "gtk_scale_new")
	core.PuregoSafeRegister(&xNewWithRangeScale, lib, "gtk_scale_new_with_range")

	core.PuregoSafeRegister(&xScaleAddMark, lib, "gtk_scale_add_mark")
	core.PuregoSafeRegister(&xScaleClearMarks, lib, "gtk_scale_clear_marks")
	core.PuregoSafeRegister(&xScaleGetDigits, lib, "gtk_scale_get_digits")
	core.PuregoSafeRegister(&xScaleGetDrawValue, lib, "gtk_scale_get_draw_value")
	core.PuregoSafeRegister(&xScaleGetHasOrigin, lib, "gtk_scale_get_has_origin")
	core.PuregoSafeRegister(&xScaleGetLayout, lib, "gtk_scale_get_layout")
	core.PuregoSafeRegister(&xScaleGetLayoutOffsets, lib, "gtk_scale_get_layout_offsets")
	core.PuregoSafeRegister(&xScaleGetValuePos, lib, "gtk_scale_get_value_pos")
	core.PuregoSafeRegister(&xScaleSetDigits, lib, "gtk_scale_set_digits")
	core.PuregoSafeRegister(&xScaleSetDrawValue, lib, "gtk_scale_set_draw_value")
	core.PuregoSafeRegister(&xScaleSetFormatValueFunc, lib, "gtk_scale_set_format_value_func")
	core.PuregoSafeRegister(&xScaleSetHasOrigin, lib, "gtk_scale_set_has_origin")
	core.PuregoSafeRegister(&xScaleSetValuePos, lib, "gtk_scale_set_value_pos")

}
