// Package gtk was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gtk

import (
	"unsafe"

	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/gobject"
)

type BoxClass struct {
	ParentClass uintptr

	Padding uintptr
}

func (x *BoxClass) GoPointer() uintptr {
	return uintptr(unsafe.Pointer(x))
}

// The `GtkBox` widget arranges child widgets into a single row or column.
//
// ![An example GtkBox](box.png)
//
// Whether it is a row or column depends on the value of its
// [property@Gtk.Orientable:orientation] property. Within the other
// dimension, all children are allocated the same size. Of course, the
// [property@Gtk.Widget:halign] and [property@Gtk.Widget:valign] properties
// can be used on the children to influence their allocation.
//
// Use repeated calls to [method@Gtk.Box.append] to pack widgets into a
// `GtkBox` from start to end. Use [method@Gtk.Box.remove] to remove widgets
// from the `GtkBox`. [method@Gtk.Box.insert_child_after] can be used to add
// a child at a particular position.
//
// Use [method@Gtk.Box.set_homogeneous] to specify whether or not all children
// of the `GtkBox` are forced to get the same amount of space.
//
// Use [method@Gtk.Box.set_spacing] to determine how much space will be minimally
// placed between all children in the `GtkBox`. Note that spacing is added
// *between* the children.
//
// Use [method@Gtk.Box.reorder_child_after] to move a child to a different
// place in the box.
//
// # CSS nodes
//
// `GtkBox` uses a single CSS node with name box.
//
// # Accessibility
//
// `GtkBox` uses the %GTK_ACCESSIBLE_ROLE_GROUP role.
type Box struct {
	Widget
}

func BoxNewFromInternalPtr(ptr uintptr) *Box {
	cls := &Box{}
	cls.Ptr = ptr
	return cls
}

var xNewBox func(Orientation, int) uintptr

// Creates a new `GtkBox`.
func NewBox(OrientationVar Orientation, SpacingVar int) *Box {
	var cls *Box

	cret := xNewBox(OrientationVar, SpacingVar)

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &Box{}
	cls.Ptr = cret
	return cls
}

var xBoxAppend func(uintptr, uintptr)

// Adds @child as the last child to @box.
func (x *Box) Append(ChildVar *Widget) {

	xBoxAppend(x.GoPointer(), ChildVar.GoPointer())

}

var xBoxGetBaselinePosition func(uintptr) BaselinePosition

// Gets the value set by gtk_box_set_baseline_position().
func (x *Box) GetBaselinePosition() BaselinePosition {

	cret := xBoxGetBaselinePosition(x.GoPointer())
	return cret
}

var xBoxGetHomogeneous func(uintptr) bool

// Returns whether the box is homogeneous (all children are the
// same size).
func (x *Box) GetHomogeneous() bool {

	cret := xBoxGetHomogeneous(x.GoPointer())
	return cret
}

var xBoxGetSpacing func(uintptr) int

// Gets the value set by gtk_box_set_spacing().
func (x *Box) GetSpacing() int {

	cret := xBoxGetSpacing(x.GoPointer())
	return cret
}

var xBoxInsertChildAfter func(uintptr, uintptr, uintptr)

// Inserts @child in the position after @sibling in the list
// of @box children.
//
// If @sibling is %NULL, insert @child at the first position.
func (x *Box) InsertChildAfter(ChildVar *Widget, SiblingVar *Widget) {

	xBoxInsertChildAfter(x.GoPointer(), ChildVar.GoPointer(), SiblingVar.GoPointer())

}

var xBoxPrepend func(uintptr, uintptr)

// Adds @child as the first child to @box.
func (x *Box) Prepend(ChildVar *Widget) {

	xBoxPrepend(x.GoPointer(), ChildVar.GoPointer())

}

var xBoxRemove func(uintptr, uintptr)

// Removes a child widget from @box.
//
// The child must have been added before with
// [method@Gtk.Box.append], [method@Gtk.Box.prepend], or
// [method@Gtk.Box.insert_child_after].
func (x *Box) Remove(ChildVar *Widget) {

	xBoxRemove(x.GoPointer(), ChildVar.GoPointer())

}

var xBoxReorderChildAfter func(uintptr, uintptr, uintptr)

// Moves @child to the position after @sibling in the list
// of @box children.
//
// If @sibling is %NULL, move @child to the first position.
func (x *Box) ReorderChildAfter(ChildVar *Widget, SiblingVar *Widget) {

	xBoxReorderChildAfter(x.GoPointer(), ChildVar.GoPointer(), SiblingVar.GoPointer())

}

var xBoxSetBaselinePosition func(uintptr, BaselinePosition)

// Sets the baseline position of a box.
//
// This affects only horizontal boxes with at least one baseline
// aligned child. If there is more vertical space available than
// requested, and the baseline is not allocated by the parent then
// @position is used to allocate the baseline with respect to the
// extra space available.
func (x *Box) SetBaselinePosition(PositionVar BaselinePosition) {

	xBoxSetBaselinePosition(x.GoPointer(), PositionVar)

}

var xBoxSetHomogeneous func(uintptr, bool)

// Sets whether or not all children of @box are given equal space
// in the box.
func (x *Box) SetHomogeneous(HomogeneousVar bool) {

	xBoxSetHomogeneous(x.GoPointer(), HomogeneousVar)

}

var xBoxSetSpacing func(uintptr, int)

// Sets the number of pixels to place between children of @box.
func (x *Box) SetSpacing(SpacingVar int) {

	xBoxSetSpacing(x.GoPointer(), SpacingVar)

}

func (c *Box) GoPointer() uintptr {
	return c.Ptr
}

func (c *Box) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// Retrieves the `GtkAccessibleRole` for the given `GtkAccessible`.
func (x *Box) GetAccessibleRole() AccessibleRole {

	cret := XGtkAccessibleGetAccessibleRole(x.GoPointer())
	return cret
}

// Resets the accessible @property to its default value.
func (x *Box) ResetProperty(PropertyVar AccessibleProperty) {

	XGtkAccessibleResetProperty(x.GoPointer(), PropertyVar)

}

// Resets the accessible @relation to its default value.
func (x *Box) ResetRelation(RelationVar AccessibleRelation) {

	XGtkAccessibleResetRelation(x.GoPointer(), RelationVar)

}

// Resets the accessible @state to its default value.
func (x *Box) ResetState(StateVar AccessibleState) {

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
func (x *Box) UpdateProperty(FirstPropertyVar AccessibleProperty, varArgs ...interface{}) {

	XGtkAccessibleUpdateProperty(x.GoPointer(), FirstPropertyVar, varArgs...)

}

// Updates an array of accessible properties.
//
// This function should be called by `GtkWidget` types whenever an accessible
// property change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *Box) UpdatePropertyValue(NPropertiesVar int, PropertiesVar uintptr, ValuesVar uintptr) {

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
func (x *Box) UpdateRelation(FirstRelationVar AccessibleRelation, varArgs ...interface{}) {

	XGtkAccessibleUpdateRelation(x.GoPointer(), FirstRelationVar, varArgs...)

}

// Updates an array of accessible relations.
//
// This function should be called by `GtkWidget` types whenever an accessible
// relation change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *Box) UpdateRelationValue(NRelationsVar int, RelationsVar uintptr, ValuesVar uintptr) {

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
func (x *Box) UpdateState(FirstStateVar AccessibleState, varArgs ...interface{}) {

	XGtkAccessibleUpdateState(x.GoPointer(), FirstStateVar, varArgs...)

}

// Updates an array of accessible states.
//
// This function should be called by `GtkWidget` types whenever an accessible
// state change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *Box) UpdateStateValue(NStatesVar int, StatesVar uintptr, ValuesVar uintptr) {

	XGtkAccessibleUpdateStateValue(x.GoPointer(), NStatesVar, StatesVar, ValuesVar)

}

// Gets the ID of the @buildable object.
//
// `GtkBuilder` sets the name based on the ID attribute
// of the &lt;object&gt; tag used to construct the @buildable.
func (x *Box) GetBuildableId() string {

	cret := XGtkBuildableGetBuildableId(x.GoPointer())
	return cret
}

// Retrieves the orientation of the @orientable.
func (x *Box) GetOrientation() Orientation {

	cret := XGtkOrientableGetOrientation(x.GoPointer())
	return cret
}

// Sets the orientation of the @orientable.
func (x *Box) SetOrientation(OrientationVar Orientation) {

	XGtkOrientableSetOrientation(x.GoPointer(), OrientationVar)

}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GTK"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xNewBox, lib, "gtk_box_new")

	core.PuregoSafeRegister(&xBoxAppend, lib, "gtk_box_append")
	core.PuregoSafeRegister(&xBoxGetBaselinePosition, lib, "gtk_box_get_baseline_position")
	core.PuregoSafeRegister(&xBoxGetHomogeneous, lib, "gtk_box_get_homogeneous")
	core.PuregoSafeRegister(&xBoxGetSpacing, lib, "gtk_box_get_spacing")
	core.PuregoSafeRegister(&xBoxInsertChildAfter, lib, "gtk_box_insert_child_after")
	core.PuregoSafeRegister(&xBoxPrepend, lib, "gtk_box_prepend")
	core.PuregoSafeRegister(&xBoxRemove, lib, "gtk_box_remove")
	core.PuregoSafeRegister(&xBoxReorderChildAfter, lib, "gtk_box_reorder_child_after")
	core.PuregoSafeRegister(&xBoxSetBaselinePosition, lib, "gtk_box_set_baseline_position")
	core.PuregoSafeRegister(&xBoxSetHomogeneous, lib, "gtk_box_set_homogeneous")
	core.PuregoSafeRegister(&xBoxSetSpacing, lib, "gtk_box_set_spacing")

}
