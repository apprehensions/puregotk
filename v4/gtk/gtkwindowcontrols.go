// Package gtk was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gtk

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/gobject"
)

type WindowControlsClass struct {
	ParentClass uintptr
}

// `GtkWindowControls` shows window frame controls.
//
// Typical window frame controls are minimize, maximize and close buttons,
// and the window icon.
//
// ![An example GtkWindowControls](windowcontrols.png)
//
// `GtkWindowControls` only displays start or end side of the controls (see
// [property@Gtk.WindowControls:side]), so it's intended to be always used
// in pair with another `GtkWindowControls` for the opposite side, for example:
//
// ```xml
// &lt;object class="GtkBox"&gt;
//
//	&lt;child&gt;
//	  &lt;object class="GtkWindowControls"&gt;
//	    &lt;property name="side"&gt;start&lt;/property&gt;
//	  &lt;/object&gt;
//	&lt;/child&gt;
//
//	...
//
//	&lt;child&gt;
//	  &lt;object class="GtkWindowControls"&gt;
//	    &lt;property name="side"&gt;end&lt;/property&gt;
//	  &lt;/object&gt;
//	&lt;/child&gt;
//
// &lt;/object&gt;
// ```
//
// # CSS nodes
//
// ```
// windowcontrols
// ├── [image.icon]
// ├── [button.minimize]
// ├── [button.maximize]
// ╰── [button.close]
// ```
//
// A `GtkWindowControls`' CSS node is called windowcontrols. It contains
// subnodes corresponding to each title button. Which of the title buttons
// exist and where they are placed exactly depends on the desktop environment
// and [property@Gtk.WindowControls:decoration-layout] value.
//
// When [property@Gtk.WindowControls:empty] is %TRUE, it gets the .empty
// style class.
//
// # Accessibility
//
// `GtkWindowControls` uses the %GTK_ACCESSIBLE_ROLE_GROUP role.
type WindowControls struct {
	Widget
}

func WindowControlsNewFromInternalPtr(ptr uintptr) *WindowControls {
	cls := &WindowControls{}
	cls.Ptr = ptr
	return cls
}

var xNewWindowControls func(PackType) uintptr

// Creates a new `GtkWindowControls`.
func NewWindowControls(SideVar PackType) *Widget {
	var cls *Widget

	cret := xNewWindowControls(SideVar)

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &Widget{}
	cls.Ptr = cret
	return cls
}

var xWindowControlsGetDecorationLayout func(uintptr) string

// Gets the decoration layout of this `GtkWindowControls`.
func (x *WindowControls) GetDecorationLayout() string {

	cret := xWindowControlsGetDecorationLayout(x.GoPointer())
	return cret
}

var xWindowControlsGetEmpty func(uintptr) bool

// Gets whether the widget has any window buttons.
func (x *WindowControls) GetEmpty() bool {

	cret := xWindowControlsGetEmpty(x.GoPointer())
	return cret
}

var xWindowControlsGetSide func(uintptr) PackType

// Gets the side to which this `GtkWindowControls` instance belongs.
func (x *WindowControls) GetSide() PackType {

	cret := xWindowControlsGetSide(x.GoPointer())
	return cret
}

var xWindowControlsSetDecorationLayout func(uintptr, string)

// Sets the decoration layout for the title buttons.
//
// This overrides the [property@Gtk.Settings:gtk-decoration-layout]
// setting.
//
// The format of the string is button names, separated by commas.
// A colon separates the buttons that should appear on the left
// from those on the right. Recognized button names are minimize,
// maximize, close and icon (the window icon).
//
// For example, “icon:minimize,maximize,close” specifies a icon
// on the left, and minimize, maximize and close buttons on the right.
//
// If [property@Gtk.WindowControls:side] value is @GTK_PACK_START, @self
// will display the part before the colon, otherwise after that.
func (x *WindowControls) SetDecorationLayout(LayoutVar string) {

	xWindowControlsSetDecorationLayout(x.GoPointer(), LayoutVar)

}

var xWindowControlsSetSide func(uintptr, PackType)

// Determines which part of decoration layout the `GtkWindowControls` uses.
//
// See [property@Gtk.WindowControls:decoration-layout].
func (x *WindowControls) SetSide(SideVar PackType) {

	xWindowControlsSetSide(x.GoPointer(), SideVar)

}

func (c *WindowControls) GoPointer() uintptr {
	return c.Ptr
}

func (c *WindowControls) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// Retrieves the `GtkAccessibleRole` for the given `GtkAccessible`.
func (x *WindowControls) GetAccessibleRole() AccessibleRole {

	cret := XGtkAccessibleGetAccessibleRole(x.GoPointer())
	return cret
}

// Resets the accessible @property to its default value.
func (x *WindowControls) ResetProperty(PropertyVar AccessibleProperty) {

	XGtkAccessibleResetProperty(x.GoPointer(), PropertyVar)

}

// Resets the accessible @relation to its default value.
func (x *WindowControls) ResetRelation(RelationVar AccessibleRelation) {

	XGtkAccessibleResetRelation(x.GoPointer(), RelationVar)

}

// Resets the accessible @state to its default value.
func (x *WindowControls) ResetState(StateVar AccessibleState) {

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
func (x *WindowControls) UpdateProperty(FirstPropertyVar AccessibleProperty, varArgs ...interface{}) {

	XGtkAccessibleUpdateProperty(x.GoPointer(), FirstPropertyVar, varArgs...)

}

// Updates an array of accessible properties.
//
// This function should be called by `GtkWidget` types whenever an accessible
// property change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *WindowControls) UpdatePropertyValue(NPropertiesVar int, PropertiesVar uintptr, ValuesVar uintptr) {

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
func (x *WindowControls) UpdateRelation(FirstRelationVar AccessibleRelation, varArgs ...interface{}) {

	XGtkAccessibleUpdateRelation(x.GoPointer(), FirstRelationVar, varArgs...)

}

// Updates an array of accessible relations.
//
// This function should be called by `GtkWidget` types whenever an accessible
// relation change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *WindowControls) UpdateRelationValue(NRelationsVar int, RelationsVar uintptr, ValuesVar uintptr) {

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
func (x *WindowControls) UpdateState(FirstStateVar AccessibleState, varArgs ...interface{}) {

	XGtkAccessibleUpdateState(x.GoPointer(), FirstStateVar, varArgs...)

}

// Updates an array of accessible states.
//
// This function should be called by `GtkWidget` types whenever an accessible
// state change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *WindowControls) UpdateStateValue(NStatesVar int, StatesVar uintptr, ValuesVar uintptr) {

	XGtkAccessibleUpdateStateValue(x.GoPointer(), NStatesVar, StatesVar, ValuesVar)

}

// Gets the ID of the @buildable object.
//
// `GtkBuilder` sets the name based on the ID attribute
// of the &lt;object&gt; tag used to construct the @buildable.
func (x *WindowControls) GetBuildableId() string {

	cret := XGtkBuildableGetBuildableId(x.GoPointer())
	return cret
}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GTK"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xNewWindowControls, lib, "gtk_window_controls_new")

	core.PuregoSafeRegister(&xWindowControlsGetDecorationLayout, lib, "gtk_window_controls_get_decoration_layout")
	core.PuregoSafeRegister(&xWindowControlsGetEmpty, lib, "gtk_window_controls_get_empty")
	core.PuregoSafeRegister(&xWindowControlsGetSide, lib, "gtk_window_controls_get_side")
	core.PuregoSafeRegister(&xWindowControlsSetDecorationLayout, lib, "gtk_window_controls_set_decoration_layout")
	core.PuregoSafeRegister(&xWindowControlsSetSide, lib, "gtk_window_controls_set_side")

}
