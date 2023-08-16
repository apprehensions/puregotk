// Package gtk was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gtk

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/gobject"
)

// The `GtkScrollbar` widget is a horizontal or vertical scrollbar.
//
// ![An example GtkScrollbar](scrollbar.png)
//
// Its position and movement are controlled by the adjustment that is passed to
// or created by [ctor@Gtk.Scrollbar.new]. See [class@Gtk.Adjustment] for more
// details. The [property@Gtk.Adjustment:value] field sets the position of the
// thumb and must be between [property@Gtk.Adjustment:lower] and
// [property@Gtk.Adjustment:upper] - [property@Gtk.Adjustment:page-size].
// The [property@Gtk.Adjustment:page-size] represents the size of the visible
// scrollable area.
//
// The fields [property@Gtk.Adjustment:step-increment] and
// [property@Gtk.Adjustment:page-increment] fields are added to or subtracted
// from the [property@Gtk.Adjustment:value] when the user asks to move by a step
// (using e.g. the cursor arrow keys) or by a page (using e.g. the Page Down/Up
// keys).
//
// # CSS nodes
//
// ```
// scrollbar
// ╰── range[.fine-tune]
//
//	╰── trough
//	    ╰── slider
//
// ```
//
// `GtkScrollbar` has a main CSS node with name scrollbar and a subnode for its
// contents. The main node gets the .horizontal or .vertical style classes applied,
// depending on the scrollbar's orientation.
//
// The range node gets the style class .fine-tune added when the scrollbar is
// in 'fine-tuning' mode.
//
// Other style classes that may be added to scrollbars inside
// [class@Gtk.ScrolledWindow] include the positional classes (.left, .right,
// .top, .bottom) and style classes related to overlay scrolling (.overlay-indicator,
// .dragging, .hovering).
//
// # Accessibility
//
// `GtkScrollbar` uses the %GTK_ACCESSIBLE_ROLE_SCROLLBAR role.
type Scrollbar struct {
	Widget
}

func ScrollbarNewFromInternalPtr(ptr uintptr) *Scrollbar {
	cls := &Scrollbar{}
	cls.Ptr = ptr
	return cls
}

var xNewScrollbar func(Orientation, uintptr) uintptr

// Creates a new scrollbar with the given orientation.
func NewScrollbar(OrientationVar Orientation, AdjustmentVar *Adjustment) *Widget {
	var cls *Widget

	cret := xNewScrollbar(OrientationVar, AdjustmentVar.GoPointer())

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &Widget{}
	cls.Ptr = cret
	return cls
}

var xScrollbarGetAdjustment func(uintptr) uintptr

// Returns the scrollbar's adjustment.
func (x *Scrollbar) GetAdjustment() *Adjustment {
	var cls *Adjustment

	cret := xScrollbarGetAdjustment(x.GoPointer())

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &Adjustment{}
	cls.Ptr = cret
	return cls
}

var xScrollbarSetAdjustment func(uintptr, uintptr)

// Makes the scrollbar use the given adjustment.
func (x *Scrollbar) SetAdjustment(AdjustmentVar *Adjustment) {

	xScrollbarSetAdjustment(x.GoPointer(), AdjustmentVar.GoPointer())

}

func (c *Scrollbar) GoPointer() uintptr {
	return c.Ptr
}

func (c *Scrollbar) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// Retrieves the `GtkAccessibleRole` for the given `GtkAccessible`.
func (x *Scrollbar) GetAccessibleRole() AccessibleRole {

	cret := XGtkAccessibleGetAccessibleRole(x.GoPointer())
	return cret
}

// Resets the accessible @property to its default value.
func (x *Scrollbar) ResetProperty(PropertyVar AccessibleProperty) {

	XGtkAccessibleResetProperty(x.GoPointer(), PropertyVar)

}

// Resets the accessible @relation to its default value.
func (x *Scrollbar) ResetRelation(RelationVar AccessibleRelation) {

	XGtkAccessibleResetRelation(x.GoPointer(), RelationVar)

}

// Resets the accessible @state to its default value.
func (x *Scrollbar) ResetState(StateVar AccessibleState) {

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
func (x *Scrollbar) UpdateProperty(FirstPropertyVar AccessibleProperty, varArgs ...interface{}) {

	XGtkAccessibleUpdateProperty(x.GoPointer(), FirstPropertyVar, varArgs...)

}

// Updates an array of accessible properties.
//
// This function should be called by `GtkWidget` types whenever an accessible
// property change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *Scrollbar) UpdatePropertyValue(NPropertiesVar int, PropertiesVar uintptr, ValuesVar uintptr) {

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
func (x *Scrollbar) UpdateRelation(FirstRelationVar AccessibleRelation, varArgs ...interface{}) {

	XGtkAccessibleUpdateRelation(x.GoPointer(), FirstRelationVar, varArgs...)

}

// Updates an array of accessible relations.
//
// This function should be called by `GtkWidget` types whenever an accessible
// relation change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *Scrollbar) UpdateRelationValue(NRelationsVar int, RelationsVar uintptr, ValuesVar uintptr) {

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
func (x *Scrollbar) UpdateState(FirstStateVar AccessibleState, varArgs ...interface{}) {

	XGtkAccessibleUpdateState(x.GoPointer(), FirstStateVar, varArgs...)

}

// Updates an array of accessible states.
//
// This function should be called by `GtkWidget` types whenever an accessible
// state change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *Scrollbar) UpdateStateValue(NStatesVar int, StatesVar uintptr, ValuesVar uintptr) {

	XGtkAccessibleUpdateStateValue(x.GoPointer(), NStatesVar, StatesVar, ValuesVar)

}

// Gets the ID of the @buildable object.
//
// `GtkBuilder` sets the name based on the ID attribute
// of the &lt;object&gt; tag used to construct the @buildable.
func (x *Scrollbar) GetBuildableId() string {

	cret := XGtkBuildableGetBuildableId(x.GoPointer())
	return cret
}

// Retrieves the orientation of the @orientable.
func (x *Scrollbar) GetOrientation() Orientation {

	cret := XGtkOrientableGetOrientation(x.GoPointer())
	return cret
}

// Sets the orientation of the @orientable.
func (x *Scrollbar) SetOrientation(OrientationVar Orientation) {

	XGtkOrientableSetOrientation(x.GoPointer(), OrientationVar)

}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GTK"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xNewScrollbar, lib, "gtk_scrollbar_new")

	core.PuregoSafeRegister(&xScrollbarGetAdjustment, lib, "gtk_scrollbar_get_adjustment")
	core.PuregoSafeRegister(&xScrollbarSetAdjustment, lib, "gtk_scrollbar_set_adjustment")

}
