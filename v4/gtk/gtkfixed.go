// Package gtk was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gtk

import (
	"unsafe"

	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/gobject"
	"github.com/jwijenbergh/puregotk/v4/gsk"
)

type FixedClass struct {
	ParentClass uintptr

	Padding uintptr
}

func (x *FixedClass) GoPointer() uintptr {
	return uintptr(unsafe.Pointer(x))
}

// `GtkFixed` places its child widgets at fixed positions and with fixed sizes.
//
// `GtkFixed` performs no automatic layout management.
//
// For most applications, you should not use this container! It keeps
// you from having to learn about the other GTK containers, but it
// results in broken applications.  With `GtkFixed`, the following
// things will result in truncated text, overlapping widgets, and
// other display bugs:
//
// - Themes, which may change widget sizes.
//
//   - Fonts other than the one you used to write the app will of course
//     change the size of widgets containing text; keep in mind that
//     users may use a larger font because of difficulty reading the
//     default, or they may be using a different OS that provides different fonts.
//
//   - Translation of text into other languages changes its size. Also,
//     display of non-English text will use a different font in many
//     cases.
//
// In addition, `GtkFixed` does not pay attention to text direction and
// thus may produce unwanted results if your app is run under right-to-left
// languages such as Hebrew or Arabic. That is: normally GTK will order
// containers appropriately for the text direction, e.g. to put labels to
// the right of the thing they label when using an RTL language, but it can’t
// do that with `GtkFixed`. So if you need to reorder widgets depending on
// the text direction, you would need to manually detect it and adjust child
// positions accordingly.
//
// Finally, fixed positioning makes it kind of annoying to add/remove
// UI elements, since you have to reposition all the other elements. This
// is a long-term maintenance problem for your application.
//
// If you know none of these things are an issue for your application,
// and prefer the simplicity of `GtkFixed`, by all means use the
// widget. But you should be aware of the tradeoffs.
type Fixed struct {
	Widget
}

func FixedNewFromInternalPtr(ptr uintptr) *Fixed {
	cls := &Fixed{}
	cls.Ptr = ptr
	return cls
}

var xNewFixed func() uintptr

// Creates a new `GtkFixed`.
func NewFixed() *Fixed {
	var cls *Fixed

	cret := xNewFixed()

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &Fixed{}
	cls.Ptr = cret
	return cls
}

var xFixedGetChildPosition func(uintptr, uintptr, float64, float64)

// Retrieves the translation transformation of the
// given child `GtkWidget` in the `GtkFixed`.
//
// See also: [method@Gtk.Fixed.get_child_transform].
func (x *Fixed) GetChildPosition(WidgetVar *Widget, XVar float64, YVar float64) {

	xFixedGetChildPosition(x.GoPointer(), WidgetVar.GoPointer(), XVar, YVar)

}

var xFixedGetChildTransform func(uintptr, uintptr) *gsk.Transform

// Retrieves the transformation for @widget set using
// gtk_fixed_set_child_transform().
func (x *Fixed) GetChildTransform(WidgetVar *Widget) *gsk.Transform {

	cret := xFixedGetChildTransform(x.GoPointer(), WidgetVar.GoPointer())
	return cret
}

var xFixedMove func(uintptr, uintptr, float64, float64)

// Sets a translation transformation to the given @x and @y
// coordinates to the child @widget of the `GtkFixed`.
func (x *Fixed) Move(WidgetVar *Widget, XVar float64, YVar float64) {

	xFixedMove(x.GoPointer(), WidgetVar.GoPointer(), XVar, YVar)

}

var xFixedPut func(uintptr, uintptr, float64, float64)

// Adds a widget to a `GtkFixed` at the given position.
func (x *Fixed) Put(WidgetVar *Widget, XVar float64, YVar float64) {

	xFixedPut(x.GoPointer(), WidgetVar.GoPointer(), XVar, YVar)

}

var xFixedRemove func(uintptr, uintptr)

// Removes a child from @fixed.
func (x *Fixed) Remove(WidgetVar *Widget) {

	xFixedRemove(x.GoPointer(), WidgetVar.GoPointer())

}

var xFixedSetChildTransform func(uintptr, uintptr, *gsk.Transform)

// Sets the transformation for @widget.
//
// This is a convenience function that retrieves the
// [class@Gtk.FixedLayoutChild] instance associated to
// @widget and calls [method@Gtk.FixedLayoutChild.set_transform].
func (x *Fixed) SetChildTransform(WidgetVar *Widget, TransformVar *gsk.Transform) {

	xFixedSetChildTransform(x.GoPointer(), WidgetVar.GoPointer(), TransformVar)

}

func (c *Fixed) GoPointer() uintptr {
	return c.Ptr
}

func (c *Fixed) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// Retrieves the `GtkAccessibleRole` for the given `GtkAccessible`.
func (x *Fixed) GetAccessibleRole() AccessibleRole {

	cret := XGtkAccessibleGetAccessibleRole(x.GoPointer())
	return cret
}

// Resets the accessible @property to its default value.
func (x *Fixed) ResetProperty(PropertyVar AccessibleProperty) {

	XGtkAccessibleResetProperty(x.GoPointer(), PropertyVar)

}

// Resets the accessible @relation to its default value.
func (x *Fixed) ResetRelation(RelationVar AccessibleRelation) {

	XGtkAccessibleResetRelation(x.GoPointer(), RelationVar)

}

// Resets the accessible @state to its default value.
func (x *Fixed) ResetState(StateVar AccessibleState) {

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
func (x *Fixed) UpdateProperty(FirstPropertyVar AccessibleProperty, varArgs ...interface{}) {

	XGtkAccessibleUpdateProperty(x.GoPointer(), FirstPropertyVar, varArgs...)

}

// Updates an array of accessible properties.
//
// This function should be called by `GtkWidget` types whenever an accessible
// property change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *Fixed) UpdatePropertyValue(NPropertiesVar int, PropertiesVar uintptr, ValuesVar uintptr) {

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
func (x *Fixed) UpdateRelation(FirstRelationVar AccessibleRelation, varArgs ...interface{}) {

	XGtkAccessibleUpdateRelation(x.GoPointer(), FirstRelationVar, varArgs...)

}

// Updates an array of accessible relations.
//
// This function should be called by `GtkWidget` types whenever an accessible
// relation change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *Fixed) UpdateRelationValue(NRelationsVar int, RelationsVar uintptr, ValuesVar uintptr) {

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
func (x *Fixed) UpdateState(FirstStateVar AccessibleState, varArgs ...interface{}) {

	XGtkAccessibleUpdateState(x.GoPointer(), FirstStateVar, varArgs...)

}

// Updates an array of accessible states.
//
// This function should be called by `GtkWidget` types whenever an accessible
// state change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *Fixed) UpdateStateValue(NStatesVar int, StatesVar uintptr, ValuesVar uintptr) {

	XGtkAccessibleUpdateStateValue(x.GoPointer(), NStatesVar, StatesVar, ValuesVar)

}

// Gets the ID of the @buildable object.
//
// `GtkBuilder` sets the name based on the ID attribute
// of the &lt;object&gt; tag used to construct the @buildable.
func (x *Fixed) GetBuildableId() string {

	cret := XGtkBuildableGetBuildableId(x.GoPointer())
	return cret
}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GTK"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xNewFixed, lib, "gtk_fixed_new")

	core.PuregoSafeRegister(&xFixedGetChildPosition, lib, "gtk_fixed_get_child_position")
	core.PuregoSafeRegister(&xFixedGetChildTransform, lib, "gtk_fixed_get_child_transform")
	core.PuregoSafeRegister(&xFixedMove, lib, "gtk_fixed_move")
	core.PuregoSafeRegister(&xFixedPut, lib, "gtk_fixed_put")
	core.PuregoSafeRegister(&xFixedRemove, lib, "gtk_fixed_remove")
	core.PuregoSafeRegister(&xFixedSetChildTransform, lib, "gtk_fixed_set_child_transform")

}
