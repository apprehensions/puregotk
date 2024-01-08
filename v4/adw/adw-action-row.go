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

type ActionRowClass struct {
	ParentClass uintptr

	Padding uintptr
}

func (x *ActionRowClass) GoPointer() uintptr {
	return uintptr(unsafe.Pointer(x))
}

// A [class@Gtk.ListBoxRow] used to present actions.
//
// &lt;picture&gt;
//
//	&lt;source srcset="action-row-dark.png" media="(prefers-color-scheme: dark)"&gt;
//	&lt;img src="action-row.png" alt="action-row"&gt;
//
// &lt;/picture&gt;
//
// The `AdwActionRow` widget can have a title, a subtitle and an icon. The row
// can receive additional widgets at its end, or prefix widgets at its start.
//
// It is convenient to present a preference and its related actions.
//
// `AdwActionRow` is unactivatable by default, giving it an activatable widget
// will automatically make it activatable, but unsetting it won't change the
// row's activatability.
//
// ## AdwActionRow as GtkBuildable
//
// The `AdwActionRow` implementation of the [iface@Gtk.Buildable] interface
// supports adding a child at its end by specifying “suffix” or omitting the
// “type” attribute of a &lt;child&gt; element.
//
// It also supports adding a child as a prefix widget by specifying “prefix” as
// the “type” attribute of a &lt;child&gt; element.
//
// ## CSS nodes
//
// `AdwActionRow` has a main CSS node with name `row`.
//
// It contains the subnode `box.header` for its main horizontal box, and
// `box.title` for the vertical box containing the title and subtitle labels.
//
// It contains subnodes `label.title` and `label.subtitle` representing
// respectively the title label and subtitle label.
type ActionRow struct {
	PreferencesRow
}

func ActionRowNewFromInternalPtr(ptr uintptr) *ActionRow {
	cls := &ActionRow{}
	cls.Ptr = ptr
	return cls
}

var xNewActionRow func() uintptr

// Creates a new `AdwActionRow`.
func NewActionRow() *gtk.Widget {
	var cls *gtk.Widget

	cret := xNewActionRow()

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &gtk.Widget{}
	cls.Ptr = cret
	return cls
}

var xActionRowActivate func(uintptr)

// Activates @self.
func (x *ActionRow) Activate() {

	xActionRowActivate(x.GoPointer())

}

var xActionRowAddPrefix func(uintptr, uintptr)

// Adds a prefix widget to @self.
func (x *ActionRow) AddPrefix(WidgetVar *gtk.Widget) {

	xActionRowAddPrefix(x.GoPointer(), WidgetVar.GoPointer())

}

var xActionRowAddSuffix func(uintptr, uintptr)

// Adds a suffix widget to @self.
func (x *ActionRow) AddSuffix(WidgetVar *gtk.Widget) {

	xActionRowAddSuffix(x.GoPointer(), WidgetVar.GoPointer())

}

var xActionRowGetActivatableWidget func(uintptr) uintptr

// Gets the widget activated when @self is activated.
func (x *ActionRow) GetActivatableWidget() *gtk.Widget {
	var cls *gtk.Widget

	cret := xActionRowGetActivatableWidget(x.GoPointer())

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &gtk.Widget{}
	cls.Ptr = cret
	return cls
}

var xActionRowGetIconName func(uintptr) string

// Gets the icon name for @self.
func (x *ActionRow) GetIconName() string {

	cret := xActionRowGetIconName(x.GoPointer())
	return cret
}

var xActionRowGetSubtitle func(uintptr) string

// Gets the subtitle for @self.
func (x *ActionRow) GetSubtitle() string {

	cret := xActionRowGetSubtitle(x.GoPointer())
	return cret
}

var xActionRowGetSubtitleLines func(uintptr) int

// Gets the number of lines at the end of which the subtitle label will be
// ellipsized.
func (x *ActionRow) GetSubtitleLines() int {

	cret := xActionRowGetSubtitleLines(x.GoPointer())
	return cret
}

var xActionRowGetSubtitleSelectable func(uintptr) bool

// Gets whether the user can copy the subtitle from the label
func (x *ActionRow) GetSubtitleSelectable() bool {

	cret := xActionRowGetSubtitleSelectable(x.GoPointer())
	return cret
}

var xActionRowGetTitleLines func(uintptr) int

// Gets the number of lines at the end of which the title label will be
// ellipsized.
func (x *ActionRow) GetTitleLines() int {

	cret := xActionRowGetTitleLines(x.GoPointer())
	return cret
}

var xActionRowRemove func(uintptr, uintptr)

// Removes a child from @self.
func (x *ActionRow) Remove(WidgetVar *gtk.Widget) {

	xActionRowRemove(x.GoPointer(), WidgetVar.GoPointer())

}

var xActionRowSetActivatableWidget func(uintptr, uintptr)

// Sets the widget to activate when @self is activated.
//
// The row can be activated either by clicking on it, calling
// [method@ActionRow.activate], or via mnemonics in the title.
// See the [property@PreferencesRow:use-underline] property to enable mnemonics.
//
// The target widget will be activated by emitting the
// [signal@Gtk.Widget::mnemonic-activate] signal on it.
func (x *ActionRow) SetActivatableWidget(WidgetVar *gtk.Widget) {

	xActionRowSetActivatableWidget(x.GoPointer(), WidgetVar.GoPointer())

}

var xActionRowSetIconName func(uintptr, string)

// Sets the icon name for @self.
func (x *ActionRow) SetIconName(IconNameVar string) {

	xActionRowSetIconName(x.GoPointer(), IconNameVar)

}

var xActionRowSetSubtitle func(uintptr, string)

// Sets the subtitle for @self.
//
// The subtitle is interpreted as Pango markup unless
// [property@PreferencesRow:use-markup] is set to `FALSE`.
func (x *ActionRow) SetSubtitle(SubtitleVar string) {

	xActionRowSetSubtitle(x.GoPointer(), SubtitleVar)

}

var xActionRowSetSubtitleLines func(uintptr, int)

// Sets the number of lines at the end of which the subtitle label will be
// ellipsized.
//
// If the value is 0, the number of lines won't be limited.
func (x *ActionRow) SetSubtitleLines(SubtitleLinesVar int) {

	xActionRowSetSubtitleLines(x.GoPointer(), SubtitleLinesVar)

}

var xActionRowSetSubtitleSelectable func(uintptr, bool)

// Sets whether the user can copy the subtitle from the label
//
// See also [property@Gtk.Label:selectable].
func (x *ActionRow) SetSubtitleSelectable(SubtitleSelectableVar bool) {

	xActionRowSetSubtitleSelectable(x.GoPointer(), SubtitleSelectableVar)

}

var xActionRowSetTitleLines func(uintptr, int)

// Sets the number of lines at the end of which the title label will be
// ellipsized.
//
// If the value is 0, the number of lines won't be limited.
func (x *ActionRow) SetTitleLines(TitleLinesVar int) {

	xActionRowSetTitleLines(x.GoPointer(), TitleLinesVar)

}

func (c *ActionRow) GoPointer() uintptr {
	return c.Ptr
}

func (c *ActionRow) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// This signal is emitted after the row has been activated.
func (x *ActionRow) ConnectActivated(cb func(ActionRow)) uint32 {
	fcb := func(clsPtr uintptr) {
		fa := ActionRow{}
		fa.Ptr = clsPtr

		cb(fa)

	}
	return gobject.SignalConnect(x.GoPointer(), "activated", purego.NewCallback(fcb))
}

// Retrieves the `GtkAccessibleRole` for the given `GtkAccessible`.
func (x *ActionRow) GetAccessibleRole() gtk.AccessibleRole {

	cret := gtk.XGtkAccessibleGetAccessibleRole(x.GoPointer())
	return cret
}

// Resets the accessible @property to its default value.
func (x *ActionRow) ResetProperty(PropertyVar gtk.AccessibleProperty) {

	gtk.XGtkAccessibleResetProperty(x.GoPointer(), PropertyVar)

}

// Resets the accessible @relation to its default value.
func (x *ActionRow) ResetRelation(RelationVar gtk.AccessibleRelation) {

	gtk.XGtkAccessibleResetRelation(x.GoPointer(), RelationVar)

}

// Resets the accessible @state to its default value.
func (x *ActionRow) ResetState(StateVar gtk.AccessibleState) {

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
func (x *ActionRow) UpdateProperty(FirstPropertyVar gtk.AccessibleProperty, varArgs ...interface{}) {

	gtk.XGtkAccessibleUpdateProperty(x.GoPointer(), FirstPropertyVar, varArgs...)

}

// Updates an array of accessible properties.
//
// This function should be called by `GtkWidget` types whenever an accessible
// property change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *ActionRow) UpdatePropertyValue(NPropertiesVar int, PropertiesVar uintptr, ValuesVar uintptr) {

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
func (x *ActionRow) UpdateRelation(FirstRelationVar gtk.AccessibleRelation, varArgs ...interface{}) {

	gtk.XGtkAccessibleUpdateRelation(x.GoPointer(), FirstRelationVar, varArgs...)

}

// Updates an array of accessible relations.
//
// This function should be called by `GtkWidget` types whenever an accessible
// relation change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *ActionRow) UpdateRelationValue(NRelationsVar int, RelationsVar uintptr, ValuesVar uintptr) {

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
func (x *ActionRow) UpdateState(FirstStateVar gtk.AccessibleState, varArgs ...interface{}) {

	gtk.XGtkAccessibleUpdateState(x.GoPointer(), FirstStateVar, varArgs...)

}

// Updates an array of accessible states.
//
// This function should be called by `GtkWidget` types whenever an accessible
// state change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *ActionRow) UpdateStateValue(NStatesVar int, StatesVar uintptr, ValuesVar uintptr) {

	gtk.XGtkAccessibleUpdateStateValue(x.GoPointer(), NStatesVar, StatesVar, ValuesVar)

}

// Gets the action name for @actionable.
func (x *ActionRow) GetActionName() string {

	cret := gtk.XGtkActionableGetActionName(x.GoPointer())
	return cret
}

// Gets the current target value of @actionable.
func (x *ActionRow) GetActionTargetValue() *glib.Variant {

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
func (x *ActionRow) SetActionName(ActionNameVar string) {

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
func (x *ActionRow) SetActionTarget(FormatStringVar string, varArgs ...interface{}) {

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
func (x *ActionRow) SetActionTargetValue(TargetValueVar *glib.Variant) {

	gtk.XGtkActionableSetActionTargetValue(x.GoPointer(), TargetValueVar)

}

// Sets the action-name and associated string target value of an
// actionable widget.
//
// @detailed_action_name is a string in the format accepted by
// [func@Gio.Action.parse_detailed_name].
func (x *ActionRow) SetDetailedActionName(DetailedActionNameVar string) {

	gtk.XGtkActionableSetDetailedActionName(x.GoPointer(), DetailedActionNameVar)

}

// Gets the ID of the @buildable object.
//
// `GtkBuilder` sets the name based on the ID attribute
// of the &lt;object&gt; tag used to construct the @buildable.
func (x *ActionRow) GetBuildableId() string {

	cret := gtk.XGtkBuildableGetBuildableId(x.GoPointer())
	return cret
}

func init() {
	lib, err := purego.Dlopen(core.GetPath("ADW"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xNewActionRow, lib, "adw_action_row_new")

	core.PuregoSafeRegister(&xActionRowActivate, lib, "adw_action_row_activate")
	core.PuregoSafeRegister(&xActionRowAddPrefix, lib, "adw_action_row_add_prefix")
	core.PuregoSafeRegister(&xActionRowAddSuffix, lib, "adw_action_row_add_suffix")
	core.PuregoSafeRegister(&xActionRowGetActivatableWidget, lib, "adw_action_row_get_activatable_widget")
	core.PuregoSafeRegister(&xActionRowGetIconName, lib, "adw_action_row_get_icon_name")
	core.PuregoSafeRegister(&xActionRowGetSubtitle, lib, "adw_action_row_get_subtitle")
	core.PuregoSafeRegister(&xActionRowGetSubtitleLines, lib, "adw_action_row_get_subtitle_lines")
	core.PuregoSafeRegister(&xActionRowGetSubtitleSelectable, lib, "adw_action_row_get_subtitle_selectable")
	core.PuregoSafeRegister(&xActionRowGetTitleLines, lib, "adw_action_row_get_title_lines")
	core.PuregoSafeRegister(&xActionRowRemove, lib, "adw_action_row_remove")
	core.PuregoSafeRegister(&xActionRowSetActivatableWidget, lib, "adw_action_row_set_activatable_widget")
	core.PuregoSafeRegister(&xActionRowSetIconName, lib, "adw_action_row_set_icon_name")
	core.PuregoSafeRegister(&xActionRowSetSubtitle, lib, "adw_action_row_set_subtitle")
	core.PuregoSafeRegister(&xActionRowSetSubtitleLines, lib, "adw_action_row_set_subtitle_lines")
	core.PuregoSafeRegister(&xActionRowSetSubtitleSelectable, lib, "adw_action_row_set_subtitle_selectable")
	core.PuregoSafeRegister(&xActionRowSetTitleLines, lib, "adw_action_row_set_title_lines")

}
