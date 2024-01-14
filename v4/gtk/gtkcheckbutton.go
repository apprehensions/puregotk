// Package gtk was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gtk

import (
	"unsafe"

	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/glib"
	"github.com/jwijenbergh/puregotk/v4/gobject"
)

type CheckButtonClass struct {
	ParentClass uintptr

	Padding uintptr
}

func (x *CheckButtonClass) GoPointer() uintptr {
	return uintptr(unsafe.Pointer(x))
}

// A `GtkCheckButton` places a label next to an indicator.
//
// ![Example GtkCheckButtons](check-button.png)
//
// A `GtkCheckButton` is created by calling either [ctor@Gtk.CheckButton.new]
// or [ctor@Gtk.CheckButton.new_with_label].
//
// The state of a `GtkCheckButton` can be set specifically using
// [method@Gtk.CheckButton.set_active], and retrieved using
// [method@Gtk.CheckButton.get_active].
//
// # Inconsistent state
//
// In addition to "on" and "off", check buttons can be an
// "in between" state that is neither on nor off. This can be used
// e.g. when the user has selected a range of elements (such as some
// text or spreadsheet cells) that are affected by a check button,
// and the current values in that range are inconsistent.
//
// To set a `GtkCheckButton` to inconsistent state, use
// [method@Gtk.CheckButton.set_inconsistent].
//
// # Grouping
//
// Check buttons can be grouped together, to form mutually exclusive
// groups - only one of the buttons can be toggled at a time, and toggling
// another one will switch the currently toggled one off.
//
// Grouped check buttons use a different indicator, and are commonly referred
// to as *radio buttons*.
//
// ![Example GtkCheckButtons](radio-button.png)
//
// To add a `GtkCheckButton` to a group, use [method@Gtk.CheckButton.set_group].
//
// # CSS nodes
//
// ```
// checkbutton[.text-button]
// ├── check
// ╰── [label]
// ```
//
// A `GtkCheckButton` has a main node with name checkbutton. If the
// [property@Gtk.CheckButton:label] or [property@Gtk.CheckButton:child]
// properties are set, it contains a child widget. The indicator node
// is named check when no group is set, and radio if the checkbutton
// is grouped together with other checkbuttons.
//
// # Accessibility
//
// `GtkCheckButton` uses the %GTK_ACCESSIBLE_ROLE_CHECKBOX role.
type CheckButton struct {
	Widget
}

func CheckButtonNewFromInternalPtr(ptr uintptr) *CheckButton {
	cls := &CheckButton{}
	cls.Ptr = ptr
	return cls
}

var xNewCheckButton func() uintptr

// Creates a new `GtkCheckButton`.
func NewCheckButton() *CheckButton {
	var cls *CheckButton

	cret := xNewCheckButton()

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &CheckButton{}
	cls.Ptr = cret
	return cls
}

var xNewWithLabelCheckButton func(string) uintptr

// Creates a new `GtkCheckButton` with the given text.
func NewWithLabelCheckButton(LabelVar string) *CheckButton {
	var cls *CheckButton

	cret := xNewWithLabelCheckButton(LabelVar)

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &CheckButton{}
	cls.Ptr = cret
	return cls
}

var xNewWithMnemonicCheckButton func(string) uintptr

// Creates a new `GtkCheckButton` with the given text and a mnemonic.
func NewWithMnemonicCheckButton(LabelVar string) *CheckButton {
	var cls *CheckButton

	cret := xNewWithMnemonicCheckButton(LabelVar)

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &CheckButton{}
	cls.Ptr = cret
	return cls
}

var xCheckButtonGetActive func(uintptr) bool

// Returns whether the check button is active.
func (x *CheckButton) GetActive() bool {

	cret := xCheckButtonGetActive(x.GoPointer())
	return cret
}

var xCheckButtonGetChild func(uintptr) uintptr

// Gets the child widget of @button or `NULL` if [property@CheckButton:label] is set.
func (x *CheckButton) GetChild() *Widget {
	var cls *Widget

	cret := xCheckButtonGetChild(x.GoPointer())

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &Widget{}
	cls.Ptr = cret
	return cls
}

var xCheckButtonGetInconsistent func(uintptr) bool

// Returns whether the check button is in an inconsistent state.
func (x *CheckButton) GetInconsistent() bool {

	cret := xCheckButtonGetInconsistent(x.GoPointer())
	return cret
}

var xCheckButtonGetLabel func(uintptr) string

// Returns the label of the check button or `NULL` if [property@CheckButton:child] is set.
func (x *CheckButton) GetLabel() string {

	cret := xCheckButtonGetLabel(x.GoPointer())
	return cret
}

var xCheckButtonGetUseUnderline func(uintptr) bool

// Returns whether underlines in the label indicate mnemonics.
func (x *CheckButton) GetUseUnderline() bool {

	cret := xCheckButtonGetUseUnderline(x.GoPointer())
	return cret
}

var xCheckButtonSetActive func(uintptr, bool)

// Changes the check buttons active state.
func (x *CheckButton) SetActive(SettingVar bool) {

	xCheckButtonSetActive(x.GoPointer(), SettingVar)

}

var xCheckButtonSetChild func(uintptr, uintptr)

// Sets the child widget of @button.
//
// Note that by using this API, you take full responsibility for setting
// up the proper accessibility label and description information for @button.
// Most likely, you'll either set the accessibility label or description
// for @button explicitly, or you'll set a labelled-by or described-by
// relations from @child to @button.
func (x *CheckButton) SetChild(ChildVar *Widget) {

	xCheckButtonSetChild(x.GoPointer(), ChildVar.GoPointer())

}

var xCheckButtonSetGroup func(uintptr, uintptr)

// Adds @self to the group of @group.
//
// In a group of multiple check buttons, only one button can be active
// at a time. The behavior of a checkbutton in a group is also commonly
// known as a *radio button*.
//
// Setting the group of a check button also changes the css name of the
// indicator widget's CSS node to 'radio'.
//
// Setting up groups in a cycle leads to undefined behavior.
//
// Note that the same effect can be achieved via the [iface@Gtk.Actionable]
// API, by using the same action with parameter type and state type 's'
// for all buttons in the group, and giving each button its own target
// value.
func (x *CheckButton) SetGroup(GroupVar *CheckButton) {

	xCheckButtonSetGroup(x.GoPointer(), GroupVar.GoPointer())

}

var xCheckButtonSetInconsistent func(uintptr, bool)

// Sets the `GtkCheckButton` to inconsistent state.
//
// You shoud turn off the inconsistent state again if the user checks
// the check button. This has to be done manually.
func (x *CheckButton) SetInconsistent(InconsistentVar bool) {

	xCheckButtonSetInconsistent(x.GoPointer(), InconsistentVar)

}

var xCheckButtonSetLabel func(uintptr, string)

// Sets the text of @self.
//
// If [property@Gtk.CheckButton:use-underline] is %TRUE, an underscore
// in @label is interpreted as mnemonic indicator, see
// [method@Gtk.CheckButton.set_use_underline] for details on this behavior.
func (x *CheckButton) SetLabel(LabelVar string) {

	xCheckButtonSetLabel(x.GoPointer(), LabelVar)

}

var xCheckButtonSetUseUnderline func(uintptr, bool)

// Sets whether underlines in the label indicate mnemonics.
//
// If @setting is %TRUE, an underscore character in @self's label
// indicates a mnemonic accelerator key. This behavior is similar
// to [property@Gtk.Label:use-underline].
func (x *CheckButton) SetUseUnderline(SettingVar bool) {

	xCheckButtonSetUseUnderline(x.GoPointer(), SettingVar)

}

func (c *CheckButton) GoPointer() uintptr {
	return c.Ptr
}

func (c *CheckButton) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// Emitted to when the check button is activated.
//
// The `::activate` signal on `GtkCheckButton` is an action signal and
// emitting it causes the button to animate press then release.
//
// Applications should never connect to this signal, but use the
// [signal@Gtk.CheckButton::toggled] signal.
func (x *CheckButton) ConnectActivate(cb func(CheckButton)) uint32 {
	fcb := func(clsPtr uintptr) {
		fa := CheckButton{}
		fa.Ptr = clsPtr

		cb(fa)

	}
	return gobject.SignalConnect(x.GoPointer(), "activate", purego.NewCallback(fcb))
}

// Emitted when the buttons's [property@Gtk.CheckButton:active]
// property changes.
func (x *CheckButton) ConnectToggled(cb func(CheckButton)) uint32 {
	fcb := func(clsPtr uintptr) {
		fa := CheckButton{}
		fa.Ptr = clsPtr

		cb(fa)

	}
	return gobject.SignalConnect(x.GoPointer(), "toggled", purego.NewCallback(fcb))
}

// Retrieves the `GtkAccessibleRole` for the given `GtkAccessible`.
func (x *CheckButton) GetAccessibleRole() AccessibleRole {

	cret := XGtkAccessibleGetAccessibleRole(x.GoPointer())
	return cret
}

// Resets the accessible @property to its default value.
func (x *CheckButton) ResetProperty(PropertyVar AccessibleProperty) {

	XGtkAccessibleResetProperty(x.GoPointer(), PropertyVar)

}

// Resets the accessible @relation to its default value.
func (x *CheckButton) ResetRelation(RelationVar AccessibleRelation) {

	XGtkAccessibleResetRelation(x.GoPointer(), RelationVar)

}

// Resets the accessible @state to its default value.
func (x *CheckButton) ResetState(StateVar AccessibleState) {

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
func (x *CheckButton) UpdateProperty(FirstPropertyVar AccessibleProperty, varArgs ...interface{}) {

	XGtkAccessibleUpdateProperty(x.GoPointer(), FirstPropertyVar, varArgs...)

}

// Updates an array of accessible properties.
//
// This function should be called by `GtkWidget` types whenever an accessible
// property change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *CheckButton) UpdatePropertyValue(NPropertiesVar int, PropertiesVar uintptr, ValuesVar uintptr) {

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
func (x *CheckButton) UpdateRelation(FirstRelationVar AccessibleRelation, varArgs ...interface{}) {

	XGtkAccessibleUpdateRelation(x.GoPointer(), FirstRelationVar, varArgs...)

}

// Updates an array of accessible relations.
//
// This function should be called by `GtkWidget` types whenever an accessible
// relation change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *CheckButton) UpdateRelationValue(NRelationsVar int, RelationsVar uintptr, ValuesVar uintptr) {

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
func (x *CheckButton) UpdateState(FirstStateVar AccessibleState, varArgs ...interface{}) {

	XGtkAccessibleUpdateState(x.GoPointer(), FirstStateVar, varArgs...)

}

// Updates an array of accessible states.
//
// This function should be called by `GtkWidget` types whenever an accessible
// state change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *CheckButton) UpdateStateValue(NStatesVar int, StatesVar uintptr, ValuesVar uintptr) {

	XGtkAccessibleUpdateStateValue(x.GoPointer(), NStatesVar, StatesVar, ValuesVar)

}

// Gets the action name for @actionable.
func (x *CheckButton) GetActionName() string {

	cret := XGtkActionableGetActionName(x.GoPointer())
	return cret
}

// Gets the current target value of @actionable.
func (x *CheckButton) GetActionTargetValue() *glib.Variant {

	cret := XGtkActionableGetActionTargetValue(x.GoPointer())
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
func (x *CheckButton) SetActionName(ActionNameVar string) {

	XGtkActionableSetActionName(x.GoPointer(), ActionNameVar)

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
func (x *CheckButton) SetActionTarget(FormatStringVar string, varArgs ...interface{}) {

	XGtkActionableSetActionTarget(x.GoPointer(), FormatStringVar, varArgs...)

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
func (x *CheckButton) SetActionTargetValue(TargetValueVar *glib.Variant) {

	XGtkActionableSetActionTargetValue(x.GoPointer(), TargetValueVar)

}

// Sets the action-name and associated string target value of an
// actionable widget.
//
// @detailed_action_name is a string in the format accepted by
// [func@Gio.Action.parse_detailed_name].
func (x *CheckButton) SetDetailedActionName(DetailedActionNameVar string) {

	XGtkActionableSetDetailedActionName(x.GoPointer(), DetailedActionNameVar)

}

// Gets the ID of the @buildable object.
//
// `GtkBuilder` sets the name based on the ID attribute
// of the &lt;object&gt; tag used to construct the @buildable.
func (x *CheckButton) GetBuildableId() string {

	cret := XGtkBuildableGetBuildableId(x.GoPointer())
	return cret
}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GTK"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xNewCheckButton, lib, "gtk_check_button_new")
	core.PuregoSafeRegister(&xNewWithLabelCheckButton, lib, "gtk_check_button_new_with_label")
	core.PuregoSafeRegister(&xNewWithMnemonicCheckButton, lib, "gtk_check_button_new_with_mnemonic")

	core.PuregoSafeRegister(&xCheckButtonGetActive, lib, "gtk_check_button_get_active")
	core.PuregoSafeRegister(&xCheckButtonGetChild, lib, "gtk_check_button_get_child")
	core.PuregoSafeRegister(&xCheckButtonGetInconsistent, lib, "gtk_check_button_get_inconsistent")
	core.PuregoSafeRegister(&xCheckButtonGetLabel, lib, "gtk_check_button_get_label")
	core.PuregoSafeRegister(&xCheckButtonGetUseUnderline, lib, "gtk_check_button_get_use_underline")
	core.PuregoSafeRegister(&xCheckButtonSetActive, lib, "gtk_check_button_set_active")
	core.PuregoSafeRegister(&xCheckButtonSetChild, lib, "gtk_check_button_set_child")
	core.PuregoSafeRegister(&xCheckButtonSetGroup, lib, "gtk_check_button_set_group")
	core.PuregoSafeRegister(&xCheckButtonSetInconsistent, lib, "gtk_check_button_set_inconsistent")
	core.PuregoSafeRegister(&xCheckButtonSetLabel, lib, "gtk_check_button_set_label")
	core.PuregoSafeRegister(&xCheckButtonSetUseUnderline, lib, "gtk_check_button_set_use_underline")

}
