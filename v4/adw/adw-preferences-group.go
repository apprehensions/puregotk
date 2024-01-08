// Package adw was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package adw

import (
	"unsafe"

	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/gobject"
	"github.com/jwijenbergh/puregotk/v4/gtk"
)

type PreferencesGroupClass struct {
	ParentClass uintptr

	Padding uintptr
}

func (x *PreferencesGroupClass) GoPointer() uintptr {
	return uintptr(unsafe.Pointer(x))
}

// A group of preference rows.
//
// &lt;picture&gt;
//
//	&lt;source srcset="preferences-group-dark.png" media="(prefers-color-scheme: dark)"&gt;
//	&lt;img src="preferences-group.png" alt="preferences-group"&gt;
//
// &lt;/picture&gt;
//
// An `AdwPreferencesGroup` represents a group or tightly related preferences,
// which in turn are represented by [class@PreferencesRow].
//
// To summarize the role of the preferences it gathers, a group can have both a
// title and a description. The title will be used by [class@PreferencesWindow]
// to let the user look for a preference.
//
// ## AdwPreferencesGroup as GtkBuildable
//
// The `AdwPreferencesGroup` implementation of the [iface@Gtk.Buildable] interface
// supports adding [class@PreferencesRow]s to the list by omitting "type". If "type"
// is omitted and the widget isn't a [class@PreferencesRow] the child is added to
// a box below the list.
//
// When the "type" attribute of a child is `header-suffix`, the child
// is set as the suffix on the end of the title and description.
//
// ## CSS nodes
//
// `AdwPreferencesGroup` has a single CSS node with name `preferencesgroup`.
//
// ## Accessibility
//
// `AdwPreferencesGroup` uses the `GTK_ACCESSIBLE_ROLE_GROUP` role.
type PreferencesGroup struct {
	gtk.Widget
}

func PreferencesGroupNewFromInternalPtr(ptr uintptr) *PreferencesGroup {
	cls := &PreferencesGroup{}
	cls.Ptr = ptr
	return cls
}

var xNewPreferencesGroup func() uintptr

// Creates a new `AdwPreferencesGroup`.
func NewPreferencesGroup() *gtk.Widget {
	var cls *gtk.Widget

	cret := xNewPreferencesGroup()

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &gtk.Widget{}
	cls.Ptr = cret
	return cls
}

var xPreferencesGroupAdd func(uintptr, uintptr)

// Adds a child to @self.
func (x *PreferencesGroup) Add(ChildVar *gtk.Widget) {

	xPreferencesGroupAdd(x.GoPointer(), ChildVar.GoPointer())

}

var xPreferencesGroupGetDescription func(uintptr) string

// Gets the description of @self.
func (x *PreferencesGroup) GetDescription() string {

	cret := xPreferencesGroupGetDescription(x.GoPointer())
	return cret
}

var xPreferencesGroupGetHeaderSuffix func(uintptr) uintptr

// Gets the suffix for @self's header.
func (x *PreferencesGroup) GetHeaderSuffix() *gtk.Widget {
	var cls *gtk.Widget

	cret := xPreferencesGroupGetHeaderSuffix(x.GoPointer())

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &gtk.Widget{}
	cls.Ptr = cret
	return cls
}

var xPreferencesGroupGetTitle func(uintptr) string

// Gets the title of @self.
func (x *PreferencesGroup) GetTitle() string {

	cret := xPreferencesGroupGetTitle(x.GoPointer())
	return cret
}

var xPreferencesGroupRemove func(uintptr, uintptr)

// Removes a child from @self.
func (x *PreferencesGroup) Remove(ChildVar *gtk.Widget) {

	xPreferencesGroupRemove(x.GoPointer(), ChildVar.GoPointer())

}

var xPreferencesGroupSetDescription func(uintptr, string)

// Sets the description for @self.
func (x *PreferencesGroup) SetDescription(DescriptionVar string) {

	xPreferencesGroupSetDescription(x.GoPointer(), DescriptionVar)

}

var xPreferencesGroupSetHeaderSuffix func(uintptr, uintptr)

// Sets the suffix for @self's header.
//
// Displayed above the list, next to the title and description.
//
// Suffixes are commonly used to show a button or a spinner for the whole group.
func (x *PreferencesGroup) SetHeaderSuffix(SuffixVar *gtk.Widget) {

	xPreferencesGroupSetHeaderSuffix(x.GoPointer(), SuffixVar.GoPointer())

}

var xPreferencesGroupSetTitle func(uintptr, string)

// Sets the title for @self.
func (x *PreferencesGroup) SetTitle(TitleVar string) {

	xPreferencesGroupSetTitle(x.GoPointer(), TitleVar)

}

func (c *PreferencesGroup) GoPointer() uintptr {
	return c.Ptr
}

func (c *PreferencesGroup) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// Retrieves the `GtkAccessibleRole` for the given `GtkAccessible`.
func (x *PreferencesGroup) GetAccessibleRole() gtk.AccessibleRole {

	cret := gtk.XGtkAccessibleGetAccessibleRole(x.GoPointer())
	return cret
}

// Resets the accessible @property to its default value.
func (x *PreferencesGroup) ResetProperty(PropertyVar gtk.AccessibleProperty) {

	gtk.XGtkAccessibleResetProperty(x.GoPointer(), PropertyVar)

}

// Resets the accessible @relation to its default value.
func (x *PreferencesGroup) ResetRelation(RelationVar gtk.AccessibleRelation) {

	gtk.XGtkAccessibleResetRelation(x.GoPointer(), RelationVar)

}

// Resets the accessible @state to its default value.
func (x *PreferencesGroup) ResetState(StateVar gtk.AccessibleState) {

	gtk.XGtkAccessibleResetState(x.GoPointer(), StateVar)

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
func (x *PreferencesGroup) UpdateProperty(FirstPropertyVar gtk.AccessibleProperty, varArgs ...interface{}) {

	gtk.XGtkAccessibleUpdateProperty(x.GoPointer(), FirstPropertyVar, varArgs...)

}

// Updates an array of accessible properties.
//
// This function should be called by `GtkWidget` types whenever an accessible
// property change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *PreferencesGroup) UpdatePropertyValue(NPropertiesVar int, PropertiesVar uintptr, ValuesVar uintptr) {

	gtk.XGtkAccessibleUpdatePropertyValue(x.GoPointer(), NPropertiesVar, PropertiesVar, ValuesVar)

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
func (x *PreferencesGroup) UpdateRelation(FirstRelationVar gtk.AccessibleRelation, varArgs ...interface{}) {

	gtk.XGtkAccessibleUpdateRelation(x.GoPointer(), FirstRelationVar, varArgs...)

}

// Updates an array of accessible relations.
//
// This function should be called by `GtkWidget` types whenever an accessible
// relation change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *PreferencesGroup) UpdateRelationValue(NRelationsVar int, RelationsVar uintptr, ValuesVar uintptr) {

	gtk.XGtkAccessibleUpdateRelationValue(x.GoPointer(), NRelationsVar, RelationsVar, ValuesVar)

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
func (x *PreferencesGroup) UpdateState(FirstStateVar gtk.AccessibleState, varArgs ...interface{}) {

	gtk.XGtkAccessibleUpdateState(x.GoPointer(), FirstStateVar, varArgs...)

}

// Updates an array of accessible states.
//
// This function should be called by `GtkWidget` types whenever an accessible
// state change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *PreferencesGroup) UpdateStateValue(NStatesVar int, StatesVar uintptr, ValuesVar uintptr) {

	gtk.XGtkAccessibleUpdateStateValue(x.GoPointer(), NStatesVar, StatesVar, ValuesVar)

}

// Gets the ID of the @buildable object.
//
// `GtkBuilder` sets the name based on the ID attribute
// of the &lt;object&gt; tag used to construct the @buildable.
func (x *PreferencesGroup) GetBuildableId() string {

	cret := gtk.XGtkBuildableGetBuildableId(x.GoPointer())
	return cret
}

func init() {
	lib, err := purego.Dlopen(core.GetPath("ADW"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xNewPreferencesGroup, lib, "adw_preferences_group_new")

	core.PuregoSafeRegister(&xPreferencesGroupAdd, lib, "adw_preferences_group_add")
	core.PuregoSafeRegister(&xPreferencesGroupGetDescription, lib, "adw_preferences_group_get_description")
	core.PuregoSafeRegister(&xPreferencesGroupGetHeaderSuffix, lib, "adw_preferences_group_get_header_suffix")
	core.PuregoSafeRegister(&xPreferencesGroupGetTitle, lib, "adw_preferences_group_get_title")
	core.PuregoSafeRegister(&xPreferencesGroupRemove, lib, "adw_preferences_group_remove")
	core.PuregoSafeRegister(&xPreferencesGroupSetDescription, lib, "adw_preferences_group_set_description")
	core.PuregoSafeRegister(&xPreferencesGroupSetHeaderSuffix, lib, "adw_preferences_group_set_header_suffix")
	core.PuregoSafeRegister(&xPreferencesGroupSetTitle, lib, "adw_preferences_group_set_title")

}
