// Package adw was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package adw

import (
	"unsafe"

	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/gobject"
	"github.com/jwijenbergh/puregotk/v4/gtk"
)

type ViewSwitcherTitleClass struct {
	ParentClass uintptr
}

func (x *ViewSwitcherTitleClass) GoPointer() uintptr {
	return uintptr(unsafe.Pointer(x))
}

// A view switcher title.
//
// &lt;picture&gt;
//
//	&lt;source srcset="view-switcher-title-dark.png" media="(prefers-color-scheme: dark)"&gt;
//	&lt;img src="view-switcher-title.png" alt="view-switcher-title"&gt;
//
// &lt;/picture&gt;
//
// A widget letting you switch between multiple views contained by a
// [class@ViewStack] via an [class@ViewSwitcher].
//
// It is designed to be used as the title widget of a [class@HeaderBar], and
// will display the window's title when the window is too narrow to fit the view
// switcher e.g. on mobile phones, or if there are less than two views.
//
// In order to center the title in narrow windows, the header bar should have
// [property@HeaderBar:centering-policy] set to
// `ADW_CENTERING_POLICY_STRICT`.
//
// `AdwViewSwitcherTitle` is intended to be used together with
// [class@ViewSwitcherBar].
//
// A common use case is to bind the [property@ViewSwitcherBar:reveal] property
// to [property@ViewSwitcherTitle:title-visible] to automatically reveal the
// view switcher bar when the title label is displayed in place of the view
// switcher, as follows:
//
// ```xml
// &lt;object class="GtkWindow"&gt;
//
//	&lt;property name="titlebar"&gt;
//	  &lt;object class="AdwHeaderBar"&gt;
//	    &lt;property name="centering-policy"&gt;strict&lt;/property&gt;
//	    &lt;property name="title-widget"&gt;
//	      &lt;object class="AdwViewSwitcherTitle" id="title"&gt;
//	        &lt;property name="stack"&gt;stack&lt;/property&gt;
//	      &lt;/object&gt;
//	    &lt;/property&gt;
//	  &lt;/object&gt;
//	&lt;/property&gt;
//	&lt;property name="child"&gt;
//	  &lt;object class="GtkBox"&gt;
//	    &lt;property name="orientation"&gt;vertical&lt;/property&gt;
//	    &lt;child&gt;
//	      &lt;object class="AdwViewStack" id="stack"/&gt;
//	    &lt;/child&gt;
//	    &lt;child&gt;
//	      &lt;object class="AdwViewSwitcherBar"&gt;
//	        &lt;property name="stack"&gt;stack&lt;/property&gt;
//	        &lt;binding name="reveal"&gt;
//	          &lt;lookup name="title-visible"&gt;title&lt;/lookup&gt;
//	        &lt;/binding&gt;
//	      &lt;/object&gt;
//	    &lt;/child&gt;
//	  &lt;/object&gt;
//	&lt;/property&gt;
//
// &lt;/object&gt;
// ```
//
// ## CSS nodes
//
// `AdwViewSwitcherTitle` has a single CSS node with name `viewswitchertitle`.
type ViewSwitcherTitle struct {
	gtk.Widget
}

func ViewSwitcherTitleNewFromInternalPtr(ptr uintptr) *ViewSwitcherTitle {
	cls := &ViewSwitcherTitle{}
	cls.Ptr = ptr
	return cls
}

var xNewViewSwitcherTitle func() uintptr

// Creates a new `AdwViewSwitcherTitle`.
func NewViewSwitcherTitle() *ViewSwitcherTitle {
	var cls *ViewSwitcherTitle

	cret := xNewViewSwitcherTitle()

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &ViewSwitcherTitle{}
	cls.Ptr = cret
	return cls
}

var xViewSwitcherTitleGetStack func(uintptr) uintptr

// Gets the stack controlled by @self.
func (x *ViewSwitcherTitle) GetStack() *ViewStack {
	var cls *ViewStack

	cret := xViewSwitcherTitleGetStack(x.GoPointer())

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &ViewStack{}
	cls.Ptr = cret
	return cls
}

var xViewSwitcherTitleGetSubtitle func(uintptr) string

// Gets the subtitle of @self.
func (x *ViewSwitcherTitle) GetSubtitle() string {

	cret := xViewSwitcherTitleGetSubtitle(x.GoPointer())
	return cret
}

var xViewSwitcherTitleGetTitle func(uintptr) string

// Gets the title of @self.
func (x *ViewSwitcherTitle) GetTitle() string {

	cret := xViewSwitcherTitleGetTitle(x.GoPointer())
	return cret
}

var xViewSwitcherTitleGetTitleVisible func(uintptr) bool

// Gets whether the title of @self is currently visible.
//
// If the title is visible, it means the view switcher is hidden an it may be
// wanted to show an alternative switcher, e.g. a [class@ViewSwitcherBar].
func (x *ViewSwitcherTitle) GetTitleVisible() bool {

	cret := xViewSwitcherTitleGetTitleVisible(x.GoPointer())
	return cret
}

var xViewSwitcherTitleGetViewSwitcherEnabled func(uintptr) bool

// Gets whether @self's view switcher is enabled.
func (x *ViewSwitcherTitle) GetViewSwitcherEnabled() bool {

	cret := xViewSwitcherTitleGetViewSwitcherEnabled(x.GoPointer())
	return cret
}

var xViewSwitcherTitleSetStack func(uintptr, uintptr)

// Sets the stack controlled by @self.
func (x *ViewSwitcherTitle) SetStack(StackVar *ViewStack) {

	xViewSwitcherTitleSetStack(x.GoPointer(), StackVar.GoPointer())

}

var xViewSwitcherTitleSetSubtitle func(uintptr, string)

// Sets the subtitle of @self.
//
// The subtitle should give the user additional details.
func (x *ViewSwitcherTitle) SetSubtitle(SubtitleVar string) {

	xViewSwitcherTitleSetSubtitle(x.GoPointer(), SubtitleVar)

}

var xViewSwitcherTitleSetTitle func(uintptr, string)

// Sets the title of @self.
//
// The title typically identifies the current view or content item, and
// generally does not use the application name.
func (x *ViewSwitcherTitle) SetTitle(TitleVar string) {

	xViewSwitcherTitleSetTitle(x.GoPointer(), TitleVar)

}

var xViewSwitcherTitleSetViewSwitcherEnabled func(uintptr, bool)

// Sets whether @self's view switcher is enabled.
//
// If it is disabled, the title will be displayed instead. This allows to
// programmatically hide the view switcher even if it fits in the available
// space.
//
// This can be used e.g. to ensure the view switcher is hidden below a certain
// window width, or any other constraint you find suitable.
func (x *ViewSwitcherTitle) SetViewSwitcherEnabled(EnabledVar bool) {

	xViewSwitcherTitleSetViewSwitcherEnabled(x.GoPointer(), EnabledVar)

}

func (c *ViewSwitcherTitle) GoPointer() uintptr {
	return c.Ptr
}

func (c *ViewSwitcherTitle) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// Retrieves the `GtkAccessibleRole` for the given `GtkAccessible`.
func (x *ViewSwitcherTitle) GetAccessibleRole() gtk.AccessibleRole {

	cret := gtk.XGtkAccessibleGetAccessibleRole(x.GoPointer())
	return cret
}

// Resets the accessible @property to its default value.
func (x *ViewSwitcherTitle) ResetProperty(PropertyVar gtk.AccessibleProperty) {

	gtk.XGtkAccessibleResetProperty(x.GoPointer(), PropertyVar)

}

// Resets the accessible @relation to its default value.
func (x *ViewSwitcherTitle) ResetRelation(RelationVar gtk.AccessibleRelation) {

	gtk.XGtkAccessibleResetRelation(x.GoPointer(), RelationVar)

}

// Resets the accessible @state to its default value.
func (x *ViewSwitcherTitle) ResetState(StateVar gtk.AccessibleState) {

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
func (x *ViewSwitcherTitle) UpdateProperty(FirstPropertyVar gtk.AccessibleProperty, varArgs ...interface{}) {

	gtk.XGtkAccessibleUpdateProperty(x.GoPointer(), FirstPropertyVar, varArgs...)

}

// Updates an array of accessible properties.
//
// This function should be called by `GtkWidget` types whenever an accessible
// property change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *ViewSwitcherTitle) UpdatePropertyValue(NPropertiesVar int, PropertiesVar uintptr, ValuesVar uintptr) {

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
func (x *ViewSwitcherTitle) UpdateRelation(FirstRelationVar gtk.AccessibleRelation, varArgs ...interface{}) {

	gtk.XGtkAccessibleUpdateRelation(x.GoPointer(), FirstRelationVar, varArgs...)

}

// Updates an array of accessible relations.
//
// This function should be called by `GtkWidget` types whenever an accessible
// relation change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *ViewSwitcherTitle) UpdateRelationValue(NRelationsVar int, RelationsVar uintptr, ValuesVar uintptr) {

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
func (x *ViewSwitcherTitle) UpdateState(FirstStateVar gtk.AccessibleState, varArgs ...interface{}) {

	gtk.XGtkAccessibleUpdateState(x.GoPointer(), FirstStateVar, varArgs...)

}

// Updates an array of accessible states.
//
// This function should be called by `GtkWidget` types whenever an accessible
// state change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *ViewSwitcherTitle) UpdateStateValue(NStatesVar int, StatesVar uintptr, ValuesVar uintptr) {

	gtk.XGtkAccessibleUpdateStateValue(x.GoPointer(), NStatesVar, StatesVar, ValuesVar)

}

// Gets the ID of the @buildable object.
//
// `GtkBuilder` sets the name based on the ID attribute
// of the &lt;object&gt; tag used to construct the @buildable.
func (x *ViewSwitcherTitle) GetBuildableId() string {

	cret := gtk.XGtkBuildableGetBuildableId(x.GoPointer())
	return cret
}

func init() {
	lib, err := purego.Dlopen(core.GetPath("ADW"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xNewViewSwitcherTitle, lib, "adw_view_switcher_title_new")

	core.PuregoSafeRegister(&xViewSwitcherTitleGetStack, lib, "adw_view_switcher_title_get_stack")
	core.PuregoSafeRegister(&xViewSwitcherTitleGetSubtitle, lib, "adw_view_switcher_title_get_subtitle")
	core.PuregoSafeRegister(&xViewSwitcherTitleGetTitle, lib, "adw_view_switcher_title_get_title")
	core.PuregoSafeRegister(&xViewSwitcherTitleGetTitleVisible, lib, "adw_view_switcher_title_get_title_visible")
	core.PuregoSafeRegister(&xViewSwitcherTitleGetViewSwitcherEnabled, lib, "adw_view_switcher_title_get_view_switcher_enabled")
	core.PuregoSafeRegister(&xViewSwitcherTitleSetStack, lib, "adw_view_switcher_title_set_stack")
	core.PuregoSafeRegister(&xViewSwitcherTitleSetSubtitle, lib, "adw_view_switcher_title_set_subtitle")
	core.PuregoSafeRegister(&xViewSwitcherTitleSetTitle, lib, "adw_view_switcher_title_set_title")
	core.PuregoSafeRegister(&xViewSwitcherTitleSetViewSwitcherEnabled, lib, "adw_view_switcher_title_set_view_switcher_enabled")

}
