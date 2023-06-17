// Package gtk was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gtk

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
)

var xTestAccessibleAssertionMessageRole func(string, string, int32, string, string, uintptr, AccessibleRole, AccessibleRole)

func TestAccessibleAssertionMessageRole(DomainVar string, FileVar string, LineVar int32, FuncVar string, ExprVar string, AccessibleVar Accessible, ExpectedRoleVar AccessibleRole, ActualRoleVar AccessibleRole) {

	xTestAccessibleAssertionMessageRole(DomainVar, FileVar, LineVar, FuncVar, ExprVar, AccessibleVar.GoPointer(), ExpectedRoleVar, ActualRoleVar)

}

var xTestAccessibleCheckProperty func(uintptr, AccessibleProperty, ...interface{}) string

// Checks whether the accessible @property of @accessible is set to
// a specific value.
func TestAccessibleCheckProperty(AccessibleVar Accessible, PropertyVar AccessibleProperty, varArgs ...interface{}) string {

	return xTestAccessibleCheckProperty(AccessibleVar.GoPointer(), PropertyVar, varArgs...)

}

var xTestAccessibleCheckRelation func(uintptr, AccessibleRelation, ...interface{}) string

// Checks whether the accessible @relation of @accessible is set to
// a specific value.
func TestAccessibleCheckRelation(AccessibleVar Accessible, RelationVar AccessibleRelation, varArgs ...interface{}) string {

	return xTestAccessibleCheckRelation(AccessibleVar.GoPointer(), RelationVar, varArgs...)

}

var xTestAccessibleCheckState func(uintptr, AccessibleState, ...interface{}) string

// Checks whether the accessible @state of @accessible is set to
// a specific value.
func TestAccessibleCheckState(AccessibleVar Accessible, StateVar AccessibleState, varArgs ...interface{}) string {

	return xTestAccessibleCheckState(AccessibleVar.GoPointer(), StateVar, varArgs...)

}

var xTestAccessibleHasProperty func(uintptr, AccessibleProperty) bool

// Checks whether the `GtkAccessible` has @property set.
func TestAccessibleHasProperty(AccessibleVar Accessible, PropertyVar AccessibleProperty) bool {

	return xTestAccessibleHasProperty(AccessibleVar.GoPointer(), PropertyVar)

}

var xTestAccessibleHasRelation func(uintptr, AccessibleRelation) bool

// Checks whether the `GtkAccessible` has @relation set.
func TestAccessibleHasRelation(AccessibleVar Accessible, RelationVar AccessibleRelation) bool {

	return xTestAccessibleHasRelation(AccessibleVar.GoPointer(), RelationVar)

}

var xTestAccessibleHasRole func(uintptr, AccessibleRole) bool

// Checks whether the `GtkAccessible:accessible-role` of the accessible
// is @role.
func TestAccessibleHasRole(AccessibleVar Accessible, RoleVar AccessibleRole) bool {

	return xTestAccessibleHasRole(AccessibleVar.GoPointer(), RoleVar)

}

var xTestAccessibleHasState func(uintptr, AccessibleState) bool

// Checks whether the `GtkAccessible` has @state set.
func TestAccessibleHasState(AccessibleVar Accessible, StateVar AccessibleState) bool {

	return xTestAccessibleHasState(AccessibleVar.GoPointer(), StateVar)

}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GTK"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	core.PuregoSafeRegister(&xTestAccessibleAssertionMessageRole, lib, "gtk_test_accessible_assertion_message_role")
	core.PuregoSafeRegister(&xTestAccessibleCheckProperty, lib, "gtk_test_accessible_check_property")
	core.PuregoSafeRegister(&xTestAccessibleCheckRelation, lib, "gtk_test_accessible_check_relation")
	core.PuregoSafeRegister(&xTestAccessibleCheckState, lib, "gtk_test_accessible_check_state")
	core.PuregoSafeRegister(&xTestAccessibleHasProperty, lib, "gtk_test_accessible_has_property")
	core.PuregoSafeRegister(&xTestAccessibleHasRelation, lib, "gtk_test_accessible_has_relation")
	core.PuregoSafeRegister(&xTestAccessibleHasRole, lib, "gtk_test_accessible_has_role")
	core.PuregoSafeRegister(&xTestAccessibleHasState, lib, "gtk_test_accessible_has_state")

}
