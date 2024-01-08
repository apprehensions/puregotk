// Package adw was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package adw

import (
	"unsafe"

	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/glib"
	"github.com/jwijenbergh/puregotk/v4/gobject"
	"github.com/jwijenbergh/puregotk/v4/gtk"
)

type TabButtonClass struct {
	ParentClass uintptr
}

func (x *TabButtonClass) GoPointer() uintptr {
	return uintptr(unsafe.Pointer(x))
}

// A button that displays the number of [class@TabView] pages.
//
// &lt;picture&gt;
//
//	&lt;source srcset="tab-button-dark.png" media="(prefers-color-scheme: dark)"&gt;
//	&lt;img src="tab-button.png" alt="tab-button"&gt;
//
// &lt;/picture&gt;
//
// `AdwTabButton` is a button that displays the number of pages in a given
// `AdwTabView`, as well as whether one of the inactive pages needs attention.
//
// It's intended to be used as a visible indicator when there's no visible tab
// bar, typically opening an [class@TabOverview] on click, e.g. via the
// `overview.open` action name:
//
// ```xml
// &lt;object class="AdwTabButton"&gt;
//
//	&lt;property name="view"&gt;view&lt;/property&gt;
//	&lt;property name="action-name"&gt;overview.open&lt;/property&gt;
//
// &lt;/object&gt;
// ```
//
// ## CSS nodes
//
// `AdwTabButton` has a main CSS node with name `tabbutton`.
//
// # Accessibility
//
// `AdwTabButton` uses the `GTK_ACCESSIBLE_ROLE_BUTTON` role.
type TabButton struct {
	gtk.Widget
}

func TabButtonNewFromInternalPtr(ptr uintptr) *TabButton {
	cls := &TabButton{}
	cls.Ptr = ptr
	return cls
}

var xNewTabButton func() uintptr

// Creates a new `AdwTabButton`.
func NewTabButton() *gtk.Widget {
	var cls *gtk.Widget

	cret := xNewTabButton()

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &gtk.Widget{}
	cls.Ptr = cret
	return cls
}

var xTabButtonGetView func(uintptr) uintptr

// Gets the tab view @self displays.
func (x *TabButton) GetView() *TabView {
	var cls *TabView

	cret := xTabButtonGetView(x.GoPointer())

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &TabView{}
	cls.Ptr = cret
	return cls
}

var xTabButtonSetView func(uintptr, uintptr)

// Sets the tab view to display.
func (x *TabButton) SetView(ViewVar *TabView) {

	xTabButtonSetView(x.GoPointer(), ViewVar.GoPointer())

}

func (c *TabButton) GoPointer() uintptr {
	return c.Ptr
}

func (c *TabButton) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// Emitted to animate press then release.
//
// This is an action signal. Applications should never connect to this signal,
// but use the [signal@TabButton::clicked] signal.
func (x *TabButton) ConnectActivate(cb func(TabButton)) uint32 {
	fcb := func(clsPtr uintptr) {
		fa := TabButton{}
		fa.Ptr = clsPtr

		cb(fa)

	}
	return gobject.SignalConnect(x.GoPointer(), "activate", purego.NewCallback(fcb))
}

// Emitted when the button has been activated (pressed and released).
func (x *TabButton) ConnectClicked(cb func(TabButton)) uint32 {
	fcb := func(clsPtr uintptr) {
		fa := TabButton{}
		fa.Ptr = clsPtr

		cb(fa)

	}
	return gobject.SignalConnect(x.GoPointer(), "clicked", purego.NewCallback(fcb))
}

// Retrieves the `GtkAccessibleRole` for the given `GtkAccessible`.
func (x *TabButton) GetAccessibleRole() gtk.AccessibleRole {

	cret := gtk.XGtkAccessibleGetAccessibleRole(x.GoPointer())
	return cret
}

// Resets the accessible @property to its default value.
func (x *TabButton) ResetProperty(PropertyVar gtk.AccessibleProperty) {

	gtk.XGtkAccessibleResetProperty(x.GoPointer(), PropertyVar)

}

// Resets the accessible @relation to its default value.
func (x *TabButton) ResetRelation(RelationVar gtk.AccessibleRelation) {

	gtk.XGtkAccessibleResetRelation(x.GoPointer(), RelationVar)

}

// Resets the accessible @state to its default value.
func (x *TabButton) ResetState(StateVar gtk.AccessibleState) {

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
func (x *TabButton) UpdateProperty(FirstPropertyVar gtk.AccessibleProperty, varArgs ...interface{}) {

	gtk.XGtkAccessibleUpdateProperty(x.GoPointer(), FirstPropertyVar, varArgs...)

}

// Updates an array of accessible properties.
//
// This function should be called by `GtkWidget` types whenever an accessible
// property change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *TabButton) UpdatePropertyValue(NPropertiesVar int, PropertiesVar uintptr, ValuesVar uintptr) {

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
func (x *TabButton) UpdateRelation(FirstRelationVar gtk.AccessibleRelation, varArgs ...interface{}) {

	gtk.XGtkAccessibleUpdateRelation(x.GoPointer(), FirstRelationVar, varArgs...)

}

// Updates an array of accessible relations.
//
// This function should be called by `GtkWidget` types whenever an accessible
// relation change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *TabButton) UpdateRelationValue(NRelationsVar int, RelationsVar uintptr, ValuesVar uintptr) {

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
func (x *TabButton) UpdateState(FirstStateVar gtk.AccessibleState, varArgs ...interface{}) {

	gtk.XGtkAccessibleUpdateState(x.GoPointer(), FirstStateVar, varArgs...)

}

// Updates an array of accessible states.
//
// This function should be called by `GtkWidget` types whenever an accessible
// state change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *TabButton) UpdateStateValue(NStatesVar int, StatesVar uintptr, ValuesVar uintptr) {

	gtk.XGtkAccessibleUpdateStateValue(x.GoPointer(), NStatesVar, StatesVar, ValuesVar)

}

// Gets the action name for @actionable.
func (x *TabButton) GetActionName() string {

	cret := gtk.XGtkActionableGetActionName(x.GoPointer())
	return cret
}

// Gets the current target value of @actionable.
func (x *TabButton) GetActionTargetValue() *glib.Variant {

	cret := gtk.XGtkActionableGetActionTargetValue(x.GoPointer())
	return cret
}

// Specifies the name of the action with which this widget should be
// associated.
//
// If @action_name is %NULL then the widget will be unassociated from
// any previous action.
//
// Usually this function is used when the widget is located (or will be
// located) within the hierarchy of a `GtkApplicationWindow`.
//
// Names are of the form “win.save” or “app.quit” for actions on the
// containing [class@ApplicationWindow] or its associated [class@Application],
// respectively. This is the same form used for actions in the [class@Gio.Menu]
// associated with the window.
func (x *TabButton) SetActionName(ActionNameVar string) {

	gtk.XGtkActionableSetActionName(x.GoPointer(), ActionNameVar)

}

// Sets the target of an actionable widget.
//
// This is a convenience function that calls [ctor@GLib.Variant.new] for
// @format_string and uses the result to call
// [method@Gtk.Actionable.set_action_target_value].
//
// If you are setting a string-valued target and want to set
// the action name at the same time, you can use
// [method@Gtk.Actionable.set_detailed_action_name].
func (x *TabButton) SetActionTarget(FormatStringVar string, varArgs ...interface{}) {

	gtk.XGtkActionableSetActionTarget(x.GoPointer(), FormatStringVar, varArgs...)

}

// Sets the target value of an actionable widget.
//
// If @target_value is %NULL then the target value is unset.
//
// The target value has two purposes. First, it is used as the parameter
// to activation of the action associated with the `GtkActionable` widget.
// Second, it is used to determine if the widget should be rendered as
// “active” — the widget is active if the state is equal to the given target.
//
// Consider the example of associating a set of buttons with a [iface@Gio.Action]
// with string state in a typical “radio button” situation. Each button
// will be associated with the same action, but with a different target
// value for that action. Clicking on a particular button will activate
// the action with the target of that button, which will typically cause
// the action’s state to change to that value. Since the action’s state
// is now equal to the target value of the button, the button will now
// be rendered as active (and the other buttons, with different targets,
// rendered inactive).
func (x *TabButton) SetActionTargetValue(TargetValueVar *glib.Variant) {

	gtk.XGtkActionableSetActionTargetValue(x.GoPointer(), TargetValueVar)

}

// Sets the action-name and associated string target value of an
// actionable widget.
//
// @detailed_action_name is a string in the format accepted by
// [func@Gio.Action.parse_detailed_name].
func (x *TabButton) SetDetailedActionName(DetailedActionNameVar string) {

	gtk.XGtkActionableSetDetailedActionName(x.GoPointer(), DetailedActionNameVar)

}

// Gets the ID of the @buildable object.
//
// `GtkBuilder` sets the name based on the ID attribute
// of the &lt;object&gt; tag used to construct the @buildable.
func (x *TabButton) GetBuildableId() string {

	cret := gtk.XGtkBuildableGetBuildableId(x.GoPointer())
	return cret
}

func init() {
	lib, err := purego.Dlopen(core.GetPath("ADW"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xNewTabButton, lib, "adw_tab_button_new")

	core.PuregoSafeRegister(&xTabButtonGetView, lib, "adw_tab_button_get_view")
	core.PuregoSafeRegister(&xTabButtonSetView, lib, "adw_tab_button_set_view")

}
