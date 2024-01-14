// Package adw was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package adw

import (
	"unsafe"

	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/gobject"
	"github.com/jwijenbergh/puregotk/v4/gtk"
)

type ViewSwitcherBarClass struct {
	ParentClass uintptr
}

func (x *ViewSwitcherBarClass) GoPointer() uintptr {
	return uintptr(unsafe.Pointer(x))
}

// A view switcher action bar.
//
// &lt;picture&gt;
//
//	&lt;source srcset="view-switcher-bar-dark.png" media="(prefers-color-scheme: dark)"&gt;
//	&lt;img src="view-switcher-bar.png" alt="view-switcher-bar"&gt;
//
// &lt;/picture&gt;
//
// An action bar letting you switch between multiple views contained in a
// [class@ViewStack], via an [class@ViewSwitcher]. It is designed to be put at
// the bottom of a window and to be revealed only on really narrow windows, e.g.
// on mobile phones. It can't be revealed if there are less than two pages.
//
// `AdwViewSwitcherBar` is intended to be used together with
// [class@ViewSwitcherTitle].
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
// `AdwViewSwitcherBar` has a single CSS node with name` viewswitcherbar`.
type ViewSwitcherBar struct {
	gtk.Widget
}

func ViewSwitcherBarNewFromInternalPtr(ptr uintptr) *ViewSwitcherBar {
	cls := &ViewSwitcherBar{}
	cls.Ptr = ptr
	return cls
}

var xNewViewSwitcherBar func() uintptr

// Creates a new `AdwViewSwitcherBar`.
func NewViewSwitcherBar() *ViewSwitcherBar {
	var cls *ViewSwitcherBar

	cret := xNewViewSwitcherBar()

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &ViewSwitcherBar{}
	cls.Ptr = cret
	return cls
}

var xViewSwitcherBarGetReveal func(uintptr) bool

// Gets whether @self should be revealed or hidden.
func (x *ViewSwitcherBar) GetReveal() bool {

	cret := xViewSwitcherBarGetReveal(x.GoPointer())
	return cret
}

var xViewSwitcherBarGetStack func(uintptr) uintptr

// Gets the stack controlled by @self.
func (x *ViewSwitcherBar) GetStack() *ViewStack {
	var cls *ViewStack

	cret := xViewSwitcherBarGetStack(x.GoPointer())

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &ViewStack{}
	cls.Ptr = cret
	return cls
}

var xViewSwitcherBarSetReveal func(uintptr, bool)

// Sets whether @self should be revealed or hidden.
func (x *ViewSwitcherBar) SetReveal(RevealVar bool) {

	xViewSwitcherBarSetReveal(x.GoPointer(), RevealVar)

}

var xViewSwitcherBarSetStack func(uintptr, uintptr)

// Sets the stack controlled by @self.
func (x *ViewSwitcherBar) SetStack(StackVar *ViewStack) {

	xViewSwitcherBarSetStack(x.GoPointer(), StackVar.GoPointer())

}

func (c *ViewSwitcherBar) GoPointer() uintptr {
	return c.Ptr
}

func (c *ViewSwitcherBar) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// Retrieves the `GtkAccessibleRole` for the given `GtkAccessible`.
func (x *ViewSwitcherBar) GetAccessibleRole() gtk.AccessibleRole {

	cret := gtk.XGtkAccessibleGetAccessibleRole(x.GoPointer())
	return cret
}

// Resets the accessible @property to its default value.
func (x *ViewSwitcherBar) ResetProperty(PropertyVar gtk.AccessibleProperty) {

	gtk.XGtkAccessibleResetProperty(x.GoPointer(), PropertyVar)

}

// Resets the accessible @relation to its default value.
func (x *ViewSwitcherBar) ResetRelation(RelationVar gtk.AccessibleRelation) {

	gtk.XGtkAccessibleResetRelation(x.GoPointer(), RelationVar)

}

// Resets the accessible @state to its default value.
func (x *ViewSwitcherBar) ResetState(StateVar gtk.AccessibleState) {

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
func (x *ViewSwitcherBar) UpdateProperty(FirstPropertyVar gtk.AccessibleProperty, varArgs ...interface{}) {

	gtk.XGtkAccessibleUpdateProperty(x.GoPointer(), FirstPropertyVar, varArgs...)

}

// Updates an array of accessible properties.
//
// This function should be called by `GtkWidget` types whenever an accessible
// property change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *ViewSwitcherBar) UpdatePropertyValue(NPropertiesVar int, PropertiesVar uintptr, ValuesVar uintptr) {

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
func (x *ViewSwitcherBar) UpdateRelation(FirstRelationVar gtk.AccessibleRelation, varArgs ...interface{}) {

	gtk.XGtkAccessibleUpdateRelation(x.GoPointer(), FirstRelationVar, varArgs...)

}

// Updates an array of accessible relations.
//
// This function should be called by `GtkWidget` types whenever an accessible
// relation change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *ViewSwitcherBar) UpdateRelationValue(NRelationsVar int, RelationsVar uintptr, ValuesVar uintptr) {

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
func (x *ViewSwitcherBar) UpdateState(FirstStateVar gtk.AccessibleState, varArgs ...interface{}) {

	gtk.XGtkAccessibleUpdateState(x.GoPointer(), FirstStateVar, varArgs...)

}

// Updates an array of accessible states.
//
// This function should be called by `GtkWidget` types whenever an accessible
// state change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *ViewSwitcherBar) UpdateStateValue(NStatesVar int, StatesVar uintptr, ValuesVar uintptr) {

	gtk.XGtkAccessibleUpdateStateValue(x.GoPointer(), NStatesVar, StatesVar, ValuesVar)

}

// Gets the ID of the @buildable object.
//
// `GtkBuilder` sets the name based on the ID attribute
// of the &lt;object&gt; tag used to construct the @buildable.
func (x *ViewSwitcherBar) GetBuildableId() string {

	cret := gtk.XGtkBuildableGetBuildableId(x.GoPointer())
	return cret
}

func init() {
	lib, err := purego.Dlopen(core.GetPath("ADW"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xNewViewSwitcherBar, lib, "adw_view_switcher_bar_new")

	core.PuregoSafeRegister(&xViewSwitcherBarGetReveal, lib, "adw_view_switcher_bar_get_reveal")
	core.PuregoSafeRegister(&xViewSwitcherBarGetStack, lib, "adw_view_switcher_bar_get_stack")
	core.PuregoSafeRegister(&xViewSwitcherBarSetReveal, lib, "adw_view_switcher_bar_set_reveal")
	core.PuregoSafeRegister(&xViewSwitcherBarSetStack, lib, "adw_view_switcher_bar_set_stack")

}
