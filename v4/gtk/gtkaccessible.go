// Package gtk was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gtk

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/gobject"
)

type AccessibleInterface struct {
}

// `GtkAccessible` is an interface for describing UI elements for
// Assistive Technologies.
//
// Every accessible implementation has:
//
//   - a “role”, represented by a value of the [enum@Gtk.AccessibleRole] enumeration
//   - an “attribute”, represented by a set of [enum@Gtk.AccessibleState],
//     [enum@Gtk.AccessibleProperty] and [enum@Gtk.AccessibleRelation] values
//
// The role cannot be changed after instantiating a `GtkAccessible`
// implementation.
//
// The attributes are updated every time a UI element's state changes in
// a way that should be reflected by assistive technologies. For instance,
// if a `GtkWidget` visibility changes, the %GTK_ACCESSIBLE_STATE_HIDDEN
// state will also change to reflect the [property@Gtk.Widget:visible] property.
type Accessible interface {
	GoPointer() uintptr
	SetGoPointer(uintptr)
	GetAccessibleRole() AccessibleRole
	ResetProperty(PropertyVar AccessibleProperty)
	ResetRelation(RelationVar AccessibleRelation)
	ResetState(StateVar AccessibleState)
	UpdateProperty(FirstPropertyVar AccessibleProperty, varArgs ...interface{})
	UpdatePropertyValue(NPropertiesVar int32, PropertiesVar uintptr, ValuesVar uintptr)
	UpdateRelation(FirstRelationVar AccessibleRelation, varArgs ...interface{})
	UpdateRelationValue(NRelationsVar int32, RelationsVar uintptr, ValuesVar uintptr)
	UpdateState(FirstStateVar AccessibleState, varArgs ...interface{})
	UpdateStateValue(NStatesVar int32, StatesVar uintptr, ValuesVar uintptr)
}
type AccessibleBase struct {
	Ptr uintptr
}

func (x *AccessibleBase) GoPointer() uintptr {
	return x.Ptr
}

func (x *AccessibleBase) SetGoPointer(ptr uintptr) {
	x.Ptr = ptr
}

// Retrieves the `GtkAccessibleRole` for the given `GtkAccessible`.
func (x *AccessibleBase) GetAccessibleRole() AccessibleRole {

	return XGtkAccessibleGetAccessibleRole(x.GoPointer())

}

// Resets the accessible @property to its default value.
func (x *AccessibleBase) ResetProperty(PropertyVar AccessibleProperty) {

	XGtkAccessibleResetProperty(x.GoPointer(), PropertyVar)

}

// Resets the accessible @relation to its default value.
func (x *AccessibleBase) ResetRelation(RelationVar AccessibleRelation) {

	XGtkAccessibleResetRelation(x.GoPointer(), RelationVar)

}

// Resets the accessible @state to its default value.
func (x *AccessibleBase) ResetState(StateVar AccessibleState) {

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
func (x *AccessibleBase) UpdateProperty(FirstPropertyVar AccessibleProperty, varArgs ...interface{}) {

	XGtkAccessibleUpdateProperty(x.GoPointer(), FirstPropertyVar, varArgs...)

}

// Updates an array of accessible properties.
//
// This function should be called by `GtkWidget` types whenever an accessible
// property change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *AccessibleBase) UpdatePropertyValue(NPropertiesVar int32, PropertiesVar uintptr, ValuesVar uintptr) {

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
func (x *AccessibleBase) UpdateRelation(FirstRelationVar AccessibleRelation, varArgs ...interface{}) {

	XGtkAccessibleUpdateRelation(x.GoPointer(), FirstRelationVar, varArgs...)

}

// Updates an array of accessible relations.
//
// This function should be called by `GtkWidget` types whenever an accessible
// relation change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *AccessibleBase) UpdateRelationValue(NRelationsVar int32, RelationsVar uintptr, ValuesVar uintptr) {

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
func (x *AccessibleBase) UpdateState(FirstStateVar AccessibleState, varArgs ...interface{}) {

	XGtkAccessibleUpdateState(x.GoPointer(), FirstStateVar, varArgs...)

}

// Updates an array of accessible states.
//
// This function should be called by `GtkWidget` types whenever an accessible
// state change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *AccessibleBase) UpdateStateValue(NStatesVar int32, StatesVar uintptr, ValuesVar uintptr) {

	XGtkAccessibleUpdateStateValue(x.GoPointer(), NStatesVar, StatesVar, ValuesVar)

}

var XGtkAccessibleGetAccessibleRole func(uintptr) AccessibleRole
var XGtkAccessibleResetProperty func(uintptr, AccessibleProperty)
var XGtkAccessibleResetRelation func(uintptr, AccessibleRelation)
var XGtkAccessibleResetState func(uintptr, AccessibleState)
var XGtkAccessibleUpdateProperty func(uintptr, AccessibleProperty, ...interface{})
var XGtkAccessibleUpdatePropertyValue func(uintptr, int32, uintptr, uintptr)
var XGtkAccessibleUpdateRelation func(uintptr, AccessibleRelation, ...interface{})
var XGtkAccessibleUpdateRelationValue func(uintptr, int32, uintptr, uintptr)
var XGtkAccessibleUpdateState func(uintptr, AccessibleState, ...interface{})
var XGtkAccessibleUpdateStateValue func(uintptr, int32, uintptr, uintptr)

var xAccessiblePropertyInitValue func(AccessibleProperty, *gobject.Value)

func AccessiblePropertyInitValue(PropertyVar AccessibleProperty, ValueVar *gobject.Value) {

	xAccessiblePropertyInitValue(PropertyVar, ValueVar)

}

var xAccessibleRelationInitValue func(AccessibleRelation, *gobject.Value)

func AccessibleRelationInitValue(RelationVar AccessibleRelation, ValueVar *gobject.Value) {

	xAccessibleRelationInitValue(RelationVar, ValueVar)

}

var xAccessibleStateInitValue func(AccessibleState, *gobject.Value)

func AccessibleStateInitValue(StateVar AccessibleState, ValueVar *gobject.Value) {

	xAccessibleStateInitValue(StateVar, ValueVar)

}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GTK"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	core.PuregoSafeRegister(&xAccessiblePropertyInitValue, lib, "gtk_accessible_property_init_value")
	core.PuregoSafeRegister(&xAccessibleRelationInitValue, lib, "gtk_accessible_relation_init_value")
	core.PuregoSafeRegister(&xAccessibleStateInitValue, lib, "gtk_accessible_state_init_value")

	core.PuregoSafeRegister(&XGtkAccessibleGetAccessibleRole, lib, "gtk_accessible_get_accessible_role")
	core.PuregoSafeRegister(&XGtkAccessibleResetProperty, lib, "gtk_accessible_reset_property")
	core.PuregoSafeRegister(&XGtkAccessibleResetRelation, lib, "gtk_accessible_reset_relation")
	core.PuregoSafeRegister(&XGtkAccessibleResetState, lib, "gtk_accessible_reset_state")
	core.PuregoSafeRegister(&XGtkAccessibleUpdateProperty, lib, "gtk_accessible_update_property")
	core.PuregoSafeRegister(&XGtkAccessibleUpdatePropertyValue, lib, "gtk_accessible_update_property_value")
	core.PuregoSafeRegister(&XGtkAccessibleUpdateRelation, lib, "gtk_accessible_update_relation")
	core.PuregoSafeRegister(&XGtkAccessibleUpdateRelationValue, lib, "gtk_accessible_update_relation_value")
	core.PuregoSafeRegister(&XGtkAccessibleUpdateState, lib, "gtk_accessible_update_state")
	core.PuregoSafeRegister(&XGtkAccessibleUpdateStateValue, lib, "gtk_accessible_update_state_value")

}
