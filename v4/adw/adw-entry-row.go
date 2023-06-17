// Package adw was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package adw

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/glib"
	"github.com/jwijenbergh/puregotk/v4/gobject"
	"github.com/jwijenbergh/puregotk/v4/gtk"
	"github.com/jwijenbergh/puregotk/v4/pango"
)

type EntryRowClass struct {
	ParentClass uintptr
}

// A [class@Gtk.ListBoxRow] with an embedded text entry.
//
// &lt;picture&gt;
//
//	&lt;source srcset="entry-row-dark.png" media="(prefers-color-scheme: dark)"&gt;
//	&lt;img src="entry-row.png" alt="entry-row"&gt;
//
// &lt;/picture&gt;
//
// `AdwEntryRow` has a title that doubles as placeholder text. It shows an icon
// indicating that it's editable and can receive additional widgets before or
// after the editable part.
//
// If [property@EntryRow:show-apply-button] is set to `TRUE`, `AdwEntryRow` can
// show an apply button when editing its contents. This can be useful if
// changing its contents can result in an expensive operation, such as network
// activity.
//
// `AdwEntryRow` provides only minimal API and should be used with the
// [iface@Gtk.Editable] API.
//
// See also [class@PasswordEntryRow].
//
// ## AdwEntryRow as GtkBuildable
//
// The `AdwEntryRow` implementation of the [iface@Gtk.Buildable] interface
// supports adding a child at its end by specifying “suffix” or omitting the
// “type” attribute of a &lt;child&gt; element.
//
// It also supports adding a child as a prefix widget by specifying “prefix” as
// the “type” attribute of a &lt;child&gt; element.
//
// ## CSS nodes
//
// `AdwEntryRow` has a single CSS node with name `row` and the `.entry` style
// class.
type EntryRow struct {
	PreferencesRow
}

func EntryRowNewFromInternalPtr(ptr uintptr) *EntryRow {
	cls := &EntryRow{}
	cls.Ptr = ptr
	return cls
}

var xNewEntryRow func() uintptr

// Creates a new `AdwEntryRow`.
func NewEntryRow() *gtk.Widget {
	NewEntryRowPtr := xNewEntryRow()
	if NewEntryRowPtr == 0 {
		return nil
	}

	gobject.IncreaseRef(NewEntryRowPtr)

	NewEntryRowCls := &gtk.Widget{}
	NewEntryRowCls.Ptr = NewEntryRowPtr
	return NewEntryRowCls
}

var xEntryRowAddPrefix func(uintptr, uintptr)

// Adds a prefix widget to @self.
func (x *EntryRow) AddPrefix(WidgetVar *gtk.Widget) {

	xEntryRowAddPrefix(x.GoPointer(), WidgetVar.GoPointer())

}

var xEntryRowAddSuffix func(uintptr, uintptr)

// Adds a suffix widget to @self.
func (x *EntryRow) AddSuffix(WidgetVar *gtk.Widget) {

	xEntryRowAddSuffix(x.GoPointer(), WidgetVar.GoPointer())

}

var xEntryRowGetActivatesDefault func(uintptr) bool

// Gets whether activating the embedded entry can activate the default widget.
func (x *EntryRow) GetActivatesDefault() bool {

	return xEntryRowGetActivatesDefault(x.GoPointer())

}

var xEntryRowGetAttributes func(uintptr) *pango.AttrList

// Gets Pango attributes applied to the text of the embedded entry.
func (x *EntryRow) GetAttributes() *pango.AttrList {

	return xEntryRowGetAttributes(x.GoPointer())

}

var xEntryRowGetEnableEmojiCompletion func(uintptr) bool

// Gets whether to suggest emoji replacements on @self.
func (x *EntryRow) GetEnableEmojiCompletion() bool {

	return xEntryRowGetEnableEmojiCompletion(x.GoPointer())

}

var xEntryRowGetInputHints func(uintptr) gtk.InputHints

// Gets the additional input hints of @self.
func (x *EntryRow) GetInputHints() gtk.InputHints {

	return xEntryRowGetInputHints(x.GoPointer())

}

var xEntryRowGetInputPurpose func(uintptr) gtk.InputPurpose

// Gets the input purpose of @self.
func (x *EntryRow) GetInputPurpose() gtk.InputPurpose {

	return xEntryRowGetInputPurpose(x.GoPointer())

}

var xEntryRowGetShowApplyButton func(uintptr) bool

// Gets whether @self can show the apply button.
func (x *EntryRow) GetShowApplyButton() bool {

	return xEntryRowGetShowApplyButton(x.GoPointer())

}

var xEntryRowGrabFocusWithoutSelecting func(uintptr) bool

// Causes @self to have keyboard focus without selecting the text.
//
// See [method@Gtk.Text.grab_focus_without_selecting] for more information.
func (x *EntryRow) GrabFocusWithoutSelecting() bool {

	return xEntryRowGrabFocusWithoutSelecting(x.GoPointer())

}

var xEntryRowRemove func(uintptr, uintptr)

// Removes a child from @self.
func (x *EntryRow) Remove(WidgetVar *gtk.Widget) {

	xEntryRowRemove(x.GoPointer(), WidgetVar.GoPointer())

}

var xEntryRowSetActivatesDefault func(uintptr, bool)

// Sets whether activating the embedded entry can activate the default widget.
func (x *EntryRow) SetActivatesDefault(ActivatesVar bool) {

	xEntryRowSetActivatesDefault(x.GoPointer(), ActivatesVar)

}

var xEntryRowSetAttributes func(uintptr, *pango.AttrList)

// Sets Pango attributes to apply to the text of the embedded entry.
//
// The [struct@Pango.Attribute]'s `start_index` and `end_index` must refer to
// the [class@Gtk.EntryBuffer] text, i.e. without the preedit string.
func (x *EntryRow) SetAttributes(AttributesVar *pango.AttrList) {

	xEntryRowSetAttributes(x.GoPointer(), AttributesVar)

}

var xEntryRowSetEnableEmojiCompletion func(uintptr, bool)

// Sets whether to suggest emoji replacements on @self.
//
// Emoji replacement is done with :-delimited names, like `:heart:`.
func (x *EntryRow) SetEnableEmojiCompletion(EnableEmojiCompletionVar bool) {

	xEntryRowSetEnableEmojiCompletion(x.GoPointer(), EnableEmojiCompletionVar)

}

var xEntryRowSetInputHints func(uintptr, gtk.InputHints)

// Set additional input hints for @self.
//
// Input hints allow input methods to fine-tune their behavior.
//
// See also: [property@AdwEntryRow:input-purpose]
func (x *EntryRow) SetInputHints(HintsVar gtk.InputHints) {

	xEntryRowSetInputHints(x.GoPointer(), HintsVar)

}

var xEntryRowSetInputPurpose func(uintptr, gtk.InputPurpose)

// Sets the input purpose of @self.
//
// The input purpose can be used by input methods to adjust their behavior.
func (x *EntryRow) SetInputPurpose(PurposeVar gtk.InputPurpose) {

	xEntryRowSetInputPurpose(x.GoPointer(), PurposeVar)

}

var xEntryRowSetShowApplyButton func(uintptr, bool)

// Sets whether @self can show the apply button.
//
// When set to `TRUE`, typing text in the entry will reveal an apply button.
// Clicking it or pressing the &lt;kbd&gt;Enter&lt;/kbd&gt; key will hide the button and
// emit the [signal@EntryRow::apply] signal.
//
// This is useful if changing the entry contents can trigger an expensive
// operation, e.g. network activity, to avoid triggering it after typing every
// character.
func (x *EntryRow) SetShowApplyButton(ShowApplyButtonVar bool) {

	xEntryRowSetShowApplyButton(x.GoPointer(), ShowApplyButtonVar)

}

func (c *EntryRow) GoPointer() uintptr {
	return c.Ptr
}

func (c *EntryRow) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// Emitted when the apply button is pressed.
//
// See [property@EntryRow:show-apply-button].
func (x *EntryRow) ConnectApply(cb func(EntryRow)) {
	fcb := func(clsPtr uintptr) {
		fa := EntryRow{}
		fa.Ptr = clsPtr

		cb(fa)

	}
	gobject.ObjectConnect(x.GoPointer(), "signal::apply", purego.NewCallback(fcb))
}

// Emitted when the embedded entry is activated.
func (x *EntryRow) ConnectEntryActivated(cb func(EntryRow)) {
	fcb := func(clsPtr uintptr) {
		fa := EntryRow{}
		fa.Ptr = clsPtr

		cb(fa)

	}
	gobject.ObjectConnect(x.GoPointer(), "signal::entry-activated", purego.NewCallback(fcb))
}

// Retrieves the `GtkAccessibleRole` for the given `GtkAccessible`.
func (x *EntryRow) GetAccessibleRole() gtk.AccessibleRole {

	return gtk.XGtkAccessibleGetAccessibleRole(x.GoPointer())

}

// Resets the accessible @property to its default value.
func (x *EntryRow) ResetProperty(PropertyVar gtk.AccessibleProperty) {

	gtk.XGtkAccessibleResetProperty(x.GoPointer(), PropertyVar)

}

// Resets the accessible @relation to its default value.
func (x *EntryRow) ResetRelation(RelationVar gtk.AccessibleRelation) {

	gtk.XGtkAccessibleResetRelation(x.GoPointer(), RelationVar)

}

// Resets the accessible @state to its default value.
func (x *EntryRow) ResetState(StateVar gtk.AccessibleState) {

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
func (x *EntryRow) UpdateProperty(FirstPropertyVar gtk.AccessibleProperty, varArgs ...interface{}) {

	gtk.XGtkAccessibleUpdateProperty(x.GoPointer(), FirstPropertyVar, varArgs...)

}

// Updates an array of accessible properties.
//
// This function should be called by `GtkWidget` types whenever an accessible
// property change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *EntryRow) UpdatePropertyValue(NPropertiesVar int32, PropertiesVar uintptr, ValuesVar uintptr) {

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
func (x *EntryRow) UpdateRelation(FirstRelationVar gtk.AccessibleRelation, varArgs ...interface{}) {

	gtk.XGtkAccessibleUpdateRelation(x.GoPointer(), FirstRelationVar, varArgs...)

}

// Updates an array of accessible relations.
//
// This function should be called by `GtkWidget` types whenever an accessible
// relation change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *EntryRow) UpdateRelationValue(NRelationsVar int32, RelationsVar uintptr, ValuesVar uintptr) {

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
func (x *EntryRow) UpdateState(FirstStateVar gtk.AccessibleState, varArgs ...interface{}) {

	gtk.XGtkAccessibleUpdateState(x.GoPointer(), FirstStateVar, varArgs...)

}

// Updates an array of accessible states.
//
// This function should be called by `GtkWidget` types whenever an accessible
// state change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *EntryRow) UpdateStateValue(NStatesVar int32, StatesVar uintptr, ValuesVar uintptr) {

	gtk.XGtkAccessibleUpdateStateValue(x.GoPointer(), NStatesVar, StatesVar, ValuesVar)

}

// Gets the action name for @actionable.
func (x *EntryRow) GetActionName() string {

	return gtk.XGtkActionableGetActionName(x.GoPointer())

}

// Gets the current target value of @actionable.
func (x *EntryRow) GetActionTargetValue() *glib.Variant {

	return gtk.XGtkActionableGetActionTargetValue(x.GoPointer())

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
func (x *EntryRow) SetActionName(ActionNameVar string) {

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
func (x *EntryRow) SetActionTarget(FormatStringVar string, varArgs ...interface{}) {

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
func (x *EntryRow) SetActionTargetValue(TargetValueVar *glib.Variant) {

	gtk.XGtkActionableSetActionTargetValue(x.GoPointer(), TargetValueVar)

}

// Sets the action-name and associated string target value of an
// actionable widget.
//
// @detailed_action_name is a string in the format accepted by
// [func@Gio.Action.parse_detailed_name].
func (x *EntryRow) SetDetailedActionName(DetailedActionNameVar string) {

	gtk.XGtkActionableSetDetailedActionName(x.GoPointer(), DetailedActionNameVar)

}

// Gets the ID of the @buildable object.
//
// `GtkBuilder` sets the name based on the ID attribute
// of the &lt;object&gt; tag used to construct the @buildable.
func (x *EntryRow) GetBuildableId() string {

	return gtk.XGtkBuildableGetBuildableId(x.GoPointer())

}

// Deletes the currently selected text of the editable.
//
// This call doesn’t do anything if there is no selected text.
func (x *EntryRow) DeleteSelection() {

	gtk.XGtkEditableDeleteSelection(x.GoPointer())

}

// Deletes a sequence of characters.
//
// The characters that are deleted are those characters at positions
// from @start_pos up to, but not including @end_pos. If @end_pos is
// negative, then the characters deleted are those from @start_pos to
// the end of the text.
//
// Note that the positions are specified in characters, not bytes.
func (x *EntryRow) DeleteText(StartPosVar int32, EndPosVar int32) {

	gtk.XGtkEditableDeleteText(x.GoPointer(), StartPosVar, EndPosVar)

}

// Undoes the setup done by [method@Gtk.Editable.init_delegate].
//
// This is a helper function that should be called from dispose,
// before removing the delegate object.
func (x *EntryRow) FinishDelegate() {

	gtk.XGtkEditableFinishDelegate(x.GoPointer())

}

// Gets the alignment of the editable.
func (x *EntryRow) GetAlignment() float32 {

	return gtk.XGtkEditableGetAlignment(x.GoPointer())

}

// Retrieves a sequence of characters.
//
// The characters that are retrieved are those characters at positions
// from @start_pos up to, but not including @end_pos. If @end_pos is negative,
// then the characters retrieved are those characters from @start_pos to
// the end of the text.
//
// Note that positions are specified in characters, not bytes.
func (x *EntryRow) GetChars(StartPosVar int32, EndPosVar int32) string {

	return gtk.XGtkEditableGetChars(x.GoPointer(), StartPosVar, EndPosVar)

}

// Gets the `GtkEditable` that @editable is delegating its
// implementation to.
//
// Typically, the delegate is a [class@Gtk.Text] widget.
func (x *EntryRow) GetDelegate() *gtk.EditableBase {

	GetDelegatePtr := gtk.XGtkEditableGetDelegate(x.GoPointer())
	if GetDelegatePtr == 0 {
		return nil
	}

	gobject.IncreaseRef(GetDelegatePtr)

	GetDelegateCls := &gtk.EditableBase{}
	GetDelegateCls.Ptr = GetDelegatePtr
	return GetDelegateCls

}

// Retrieves whether @editable is editable.
func (x *EntryRow) GetEditable() bool {

	return gtk.XGtkEditableGetEditable(x.GoPointer())

}

// Gets if undo/redo actions are enabled for @editable
func (x *EntryRow) GetEnableUndo() bool {

	return gtk.XGtkEditableGetEnableUndo(x.GoPointer())

}

// Retrieves the desired maximum width of @editable, in characters.
func (x *EntryRow) GetMaxWidthChars() int32 {

	return gtk.XGtkEditableGetMaxWidthChars(x.GoPointer())

}

// Retrieves the current position of the cursor relative
// to the start of the content of the editable.
//
// Note that this position is in characters, not in bytes.
func (x *EntryRow) GetPosition() int32 {

	return gtk.XGtkEditableGetPosition(x.GoPointer())

}

// Retrieves the selection bound of the editable.
//
// @start_pos will be filled with the start of the selection and
// @end_pos with end. If no text was selected both will be identical
// and %FALSE will be returned.
//
// Note that positions are specified in characters, not bytes.
func (x *EntryRow) GetSelectionBounds(StartPosVar int32, EndPosVar int32) bool {

	return gtk.XGtkEditableGetSelectionBounds(x.GoPointer(), StartPosVar, EndPosVar)

}

// Retrieves the contents of @editable.
//
// The returned string is owned by GTK and must not be modified or freed.
func (x *EntryRow) GetText() string {

	return gtk.XGtkEditableGetText(x.GoPointer())

}

// Gets the number of characters of space reserved
// for the contents of the editable.
func (x *EntryRow) GetWidthChars() int32 {

	return gtk.XGtkEditableGetWidthChars(x.GoPointer())

}

// Sets up a delegate for `GtkEditable`.
//
// This is assuming that the get_delegate vfunc in the `GtkEditable`
// interface has been set up for the @editable's type.
//
// This is a helper function that should be called in instance init,
// after creating the delegate object.
func (x *EntryRow) InitDelegate() {

	gtk.XGtkEditableInitDelegate(x.GoPointer())

}

// Inserts @length bytes of @text into the contents of the
// widget, at position @position.
//
// Note that the position is in characters, not in bytes.
// The function updates @position to point after the newly
// inserted text.
func (x *EntryRow) InsertText(TextVar string, LengthVar int32, PositionVar int32) {

	gtk.XGtkEditableInsertText(x.GoPointer(), TextVar, LengthVar, PositionVar)

}

// Selects a region of text.
//
// The characters that are selected are those characters at positions
// from @start_pos up to, but not including @end_pos. If @end_pos is
// negative, then the characters selected are those characters from
// @start_pos to  the end of the text.
//
// Note that positions are specified in characters, not bytes.
func (x *EntryRow) SelectRegion(StartPosVar int32, EndPosVar int32) {

	gtk.XGtkEditableSelectRegion(x.GoPointer(), StartPosVar, EndPosVar)

}

// Sets the alignment for the contents of the editable.
//
// This controls the horizontal positioning of the contents when
// the displayed text is shorter than the width of the editable.
func (x *EntryRow) SetAlignment(XalignVar float32) {

	gtk.XGtkEditableSetAlignment(x.GoPointer(), XalignVar)

}

// Determines if the user can edit the text in the editable widget.
func (x *EntryRow) SetEditable(IsEditableVar bool) {

	gtk.XGtkEditableSetEditable(x.GoPointer(), IsEditableVar)

}

// If enabled, changes to @editable will be saved for undo/redo
// actions.
//
// This results in an additional copy of text changes and are not
// stored in secure memory. As such, undo is forcefully disabled
// when [property@Gtk.Text:visibility] is set to %FALSE.
func (x *EntryRow) SetEnableUndo(EnableUndoVar bool) {

	gtk.XGtkEditableSetEnableUndo(x.GoPointer(), EnableUndoVar)

}

// Sets the desired maximum width in characters of @editable.
func (x *EntryRow) SetMaxWidthChars(NCharsVar int32) {

	gtk.XGtkEditableSetMaxWidthChars(x.GoPointer(), NCharsVar)

}

// Sets the cursor position in the editable to the given value.
//
// The cursor is displayed before the character with the given (base 0)
// index in the contents of the editable. The value must be less than
// or equal to the number of characters in the editable. A value of -1
// indicates that the position should be set after the last character
// of the editable. Note that @position is in characters, not in bytes.
func (x *EntryRow) SetPosition(PositionVar int32) {

	gtk.XGtkEditableSetPosition(x.GoPointer(), PositionVar)

}

// Sets the text in the editable to the given value.
//
// This is replacing the current contents.
func (x *EntryRow) SetText(TextVar string) {

	gtk.XGtkEditableSetText(x.GoPointer(), TextVar)

}

// Changes the size request of the editable to be about the
// right size for @n_chars characters.
//
// Note that it changes the size request, the size can still
// be affected by how you pack the widget into containers.
// If @n_chars is -1, the size reverts to the default size.
func (x *EntryRow) SetWidthChars(NCharsVar int32) {

	gtk.XGtkEditableSetWidthChars(x.GoPointer(), NCharsVar)

}

func init() {
	lib, err := purego.Dlopen(core.GetPath("ADW"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xNewEntryRow, lib, "adw_entry_row_new")

	core.PuregoSafeRegister(&xEntryRowAddPrefix, lib, "adw_entry_row_add_prefix")
	core.PuregoSafeRegister(&xEntryRowAddSuffix, lib, "adw_entry_row_add_suffix")
	core.PuregoSafeRegister(&xEntryRowGetActivatesDefault, lib, "adw_entry_row_get_activates_default")
	core.PuregoSafeRegister(&xEntryRowGetAttributes, lib, "adw_entry_row_get_attributes")
	core.PuregoSafeRegister(&xEntryRowGetEnableEmojiCompletion, lib, "adw_entry_row_get_enable_emoji_completion")
	core.PuregoSafeRegister(&xEntryRowGetInputHints, lib, "adw_entry_row_get_input_hints")
	core.PuregoSafeRegister(&xEntryRowGetInputPurpose, lib, "adw_entry_row_get_input_purpose")
	core.PuregoSafeRegister(&xEntryRowGetShowApplyButton, lib, "adw_entry_row_get_show_apply_button")
	core.PuregoSafeRegister(&xEntryRowGrabFocusWithoutSelecting, lib, "adw_entry_row_grab_focus_without_selecting")
	core.PuregoSafeRegister(&xEntryRowRemove, lib, "adw_entry_row_remove")
	core.PuregoSafeRegister(&xEntryRowSetActivatesDefault, lib, "adw_entry_row_set_activates_default")
	core.PuregoSafeRegister(&xEntryRowSetAttributes, lib, "adw_entry_row_set_attributes")
	core.PuregoSafeRegister(&xEntryRowSetEnableEmojiCompletion, lib, "adw_entry_row_set_enable_emoji_completion")
	core.PuregoSafeRegister(&xEntryRowSetInputHints, lib, "adw_entry_row_set_input_hints")
	core.PuregoSafeRegister(&xEntryRowSetInputPurpose, lib, "adw_entry_row_set_input_purpose")
	core.PuregoSafeRegister(&xEntryRowSetShowApplyButton, lib, "adw_entry_row_set_show_apply_button")

}
