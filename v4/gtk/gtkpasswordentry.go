// Package gtk was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gtk

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/gio"
	"github.com/jwijenbergh/puregotk/v4/gobject"
)

type PasswordEntryClass struct {
}

// `GtkPasswordEntry` is an entry that has been tailored for entering secrets.
//
// ![An example GtkPasswordEntry](password-entry.png)
//
// It does not show its contents in clear text, does not allow to copy it
// to the clipboard, and it shows a warning when Caps Lock is engaged. If
// the underlying platform allows it, `GtkPasswordEntry` will also place
// the text in a non-pageable memory area, to avoid it being written out
// to disk by the operating system.
//
// Optionally, it can offer a way to reveal the contents in clear text.
//
// `GtkPasswordEntry` provides only minimal API and should be used with
// the [iface@Gtk.Editable] API.
//
// # CSS Nodes
//
// ```
// entry.password
// ╰── text
//
//	├── image.caps-lock-indicator
//	┊
//
// ```
//
// `GtkPasswordEntry` has a single CSS node with name entry that carries
// a .passwordstyle class. The text Css node below it has a child with
// name image and style class .caps-lock-indicator for the Caps Lock
// icon, and possibly other children.
//
// # Accessibility
//
// `GtkPasswordEntry` uses the %GTK_ACCESSIBLE_ROLE_TEXT_BOX role.
type PasswordEntry struct {
	Widget
}

func PasswordEntryNewFromInternalPtr(ptr uintptr) *PasswordEntry {
	cls := &PasswordEntry{}
	cls.Ptr = ptr
	return cls
}

var xNewPasswordEntry func() uintptr

// Creates a `GtkPasswordEntry`.
func NewPasswordEntry() *Widget {
	var cls *Widget

	cret := xNewPasswordEntry()

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &Widget{}
	cls.Ptr = cret
	return cls
}

var xPasswordEntryGetExtraMenu func(uintptr) uintptr

// Gets the menu model set with gtk_password_entry_set_extra_menu().
func (x *PasswordEntry) GetExtraMenu() *gio.MenuModel {
	var cls *gio.MenuModel

	cret := xPasswordEntryGetExtraMenu(x.GoPointer())

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &gio.MenuModel{}
	cls.Ptr = cret
	return cls
}

var xPasswordEntryGetShowPeekIcon func(uintptr) bool

// Returns whether the entry is showing an icon to
// reveal the contents.
func (x *PasswordEntry) GetShowPeekIcon() bool {

	cret := xPasswordEntryGetShowPeekIcon(x.GoPointer())
	return cret
}

var xPasswordEntrySetExtraMenu func(uintptr, uintptr)

// Sets a menu model to add when constructing
// the context menu for @entry.
func (x *PasswordEntry) SetExtraMenu(ModelVar *gio.MenuModel) {

	xPasswordEntrySetExtraMenu(x.GoPointer(), ModelVar.GoPointer())

}

var xPasswordEntrySetShowPeekIcon func(uintptr, bool)

// Sets whether the entry should have a clickable icon
// to reveal the contents.
//
// Setting this to %FALSE also hides the text again.
func (x *PasswordEntry) SetShowPeekIcon(ShowPeekIconVar bool) {

	xPasswordEntrySetShowPeekIcon(x.GoPointer(), ShowPeekIconVar)

}

func (c *PasswordEntry) GoPointer() uintptr {
	return c.Ptr
}

func (c *PasswordEntry) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// Emitted when the entry is activated.
//
// The keybindings for this signal are all forms of the Enter key.
func (x *PasswordEntry) ConnectActivate(cb func(PasswordEntry)) {
	fcb := func(clsPtr uintptr) {
		fa := PasswordEntry{}
		fa.Ptr = clsPtr

		cb(fa)

	}
	gobject.ObjectConnect(x.GoPointer(), "signal::activate", purego.NewCallback(fcb))
}

// Retrieves the `GtkAccessibleRole` for the given `GtkAccessible`.
func (x *PasswordEntry) GetAccessibleRole() AccessibleRole {

	cret := XGtkAccessibleGetAccessibleRole(x.GoPointer())
	return cret
}

// Resets the accessible @property to its default value.
func (x *PasswordEntry) ResetProperty(PropertyVar AccessibleProperty) {

	XGtkAccessibleResetProperty(x.GoPointer(), PropertyVar)

}

// Resets the accessible @relation to its default value.
func (x *PasswordEntry) ResetRelation(RelationVar AccessibleRelation) {

	XGtkAccessibleResetRelation(x.GoPointer(), RelationVar)

}

// Resets the accessible @state to its default value.
func (x *PasswordEntry) ResetState(StateVar AccessibleState) {

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
func (x *PasswordEntry) UpdateProperty(FirstPropertyVar AccessibleProperty, varArgs ...interface{}) {

	XGtkAccessibleUpdateProperty(x.GoPointer(), FirstPropertyVar, varArgs...)

}

// Updates an array of accessible properties.
//
// This function should be called by `GtkWidget` types whenever an accessible
// property change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *PasswordEntry) UpdatePropertyValue(NPropertiesVar int, PropertiesVar uintptr, ValuesVar uintptr) {

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
func (x *PasswordEntry) UpdateRelation(FirstRelationVar AccessibleRelation, varArgs ...interface{}) {

	XGtkAccessibleUpdateRelation(x.GoPointer(), FirstRelationVar, varArgs...)

}

// Updates an array of accessible relations.
//
// This function should be called by `GtkWidget` types whenever an accessible
// relation change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *PasswordEntry) UpdateRelationValue(NRelationsVar int, RelationsVar uintptr, ValuesVar uintptr) {

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
func (x *PasswordEntry) UpdateState(FirstStateVar AccessibleState, varArgs ...interface{}) {

	XGtkAccessibleUpdateState(x.GoPointer(), FirstStateVar, varArgs...)

}

// Updates an array of accessible states.
//
// This function should be called by `GtkWidget` types whenever an accessible
// state change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *PasswordEntry) UpdateStateValue(NStatesVar int, StatesVar uintptr, ValuesVar uintptr) {

	XGtkAccessibleUpdateStateValue(x.GoPointer(), NStatesVar, StatesVar, ValuesVar)

}

// Gets the ID of the @buildable object.
//
// `GtkBuilder` sets the name based on the ID attribute
// of the &lt;object&gt; tag used to construct the @buildable.
func (x *PasswordEntry) GetBuildableId() string {

	cret := XGtkBuildableGetBuildableId(x.GoPointer())
	return cret
}

// Deletes the currently selected text of the editable.
//
// This call doesn’t do anything if there is no selected text.
func (x *PasswordEntry) DeleteSelection() {

	XGtkEditableDeleteSelection(x.GoPointer())

}

// Deletes a sequence of characters.
//
// The characters that are deleted are those characters at positions
// from @start_pos up to, but not including @end_pos. If @end_pos is
// negative, then the characters deleted are those from @start_pos to
// the end of the text.
//
// Note that the positions are specified in characters, not bytes.
func (x *PasswordEntry) DeleteText(StartPosVar int, EndPosVar int) {

	XGtkEditableDeleteText(x.GoPointer(), StartPosVar, EndPosVar)

}

// Undoes the setup done by [method@Gtk.Editable.init_delegate].
//
// This is a helper function that should be called from dispose,
// before removing the delegate object.
func (x *PasswordEntry) FinishDelegate() {

	XGtkEditableFinishDelegate(x.GoPointer())

}

// Gets the alignment of the editable.
func (x *PasswordEntry) GetAlignment() float32 {

	cret := XGtkEditableGetAlignment(x.GoPointer())
	return cret
}

// Retrieves a sequence of characters.
//
// The characters that are retrieved are those characters at positions
// from @start_pos up to, but not including @end_pos. If @end_pos is negative,
// then the characters retrieved are those characters from @start_pos to
// the end of the text.
//
// Note that positions are specified in characters, not bytes.
func (x *PasswordEntry) GetChars(StartPosVar int, EndPosVar int) string {

	cret := XGtkEditableGetChars(x.GoPointer(), StartPosVar, EndPosVar)
	return cret
}

// Gets the `GtkEditable` that @editable is delegating its
// implementation to.
//
// Typically, the delegate is a [class@Gtk.Text] widget.
func (x *PasswordEntry) GetDelegate() *EditableBase {
	var cls *EditableBase

	cret := XGtkEditableGetDelegate(x.GoPointer())

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &EditableBase{}
	cls.Ptr = cret
	return cls
}

// Retrieves whether @editable is editable.
func (x *PasswordEntry) GetEditable() bool {

	cret := XGtkEditableGetEditable(x.GoPointer())
	return cret
}

// Gets if undo/redo actions are enabled for @editable
func (x *PasswordEntry) GetEnableUndo() bool {

	cret := XGtkEditableGetEnableUndo(x.GoPointer())
	return cret
}

// Retrieves the desired maximum width of @editable, in characters.
func (x *PasswordEntry) GetMaxWidthChars() int {

	cret := XGtkEditableGetMaxWidthChars(x.GoPointer())
	return cret
}

// Retrieves the current position of the cursor relative
// to the start of the content of the editable.
//
// Note that this position is in characters, not in bytes.
func (x *PasswordEntry) GetPosition() int {

	cret := XGtkEditableGetPosition(x.GoPointer())
	return cret
}

// Retrieves the selection bound of the editable.
//
// @start_pos will be filled with the start of the selection and
// @end_pos with end. If no text was selected both will be identical
// and %FALSE will be returned.
//
// Note that positions are specified in characters, not bytes.
func (x *PasswordEntry) GetSelectionBounds(StartPosVar int, EndPosVar int) bool {

	cret := XGtkEditableGetSelectionBounds(x.GoPointer(), StartPosVar, EndPosVar)
	return cret
}

// Retrieves the contents of @editable.
//
// The returned string is owned by GTK and must not be modified or freed.
func (x *PasswordEntry) GetText() string {

	cret := XGtkEditableGetText(x.GoPointer())
	return cret
}

// Gets the number of characters of space reserved
// for the contents of the editable.
func (x *PasswordEntry) GetWidthChars() int {

	cret := XGtkEditableGetWidthChars(x.GoPointer())
	return cret
}

// Sets up a delegate for `GtkEditable`.
//
// This is assuming that the get_delegate vfunc in the `GtkEditable`
// interface has been set up for the @editable's type.
//
// This is a helper function that should be called in instance init,
// after creating the delegate object.
func (x *PasswordEntry) InitDelegate() {

	XGtkEditableInitDelegate(x.GoPointer())

}

// Inserts @length bytes of @text into the contents of the
// widget, at position @position.
//
// Note that the position is in characters, not in bytes.
// The function updates @position to point after the newly
// inserted text.
func (x *PasswordEntry) InsertText(TextVar string, LengthVar int, PositionVar int) {

	XGtkEditableInsertText(x.GoPointer(), TextVar, LengthVar, PositionVar)

}

// Selects a region of text.
//
// The characters that are selected are those characters at positions
// from @start_pos up to, but not including @end_pos. If @end_pos is
// negative, then the characters selected are those characters from
// @start_pos to  the end of the text.
//
// Note that positions are specified in characters, not bytes.
func (x *PasswordEntry) SelectRegion(StartPosVar int, EndPosVar int) {

	XGtkEditableSelectRegion(x.GoPointer(), StartPosVar, EndPosVar)

}

// Sets the alignment for the contents of the editable.
//
// This controls the horizontal positioning of the contents when
// the displayed text is shorter than the width of the editable.
func (x *PasswordEntry) SetAlignment(XalignVar float32) {

	XGtkEditableSetAlignment(x.GoPointer(), XalignVar)

}

// Determines if the user can edit the text in the editable widget.
func (x *PasswordEntry) SetEditable(IsEditableVar bool) {

	XGtkEditableSetEditable(x.GoPointer(), IsEditableVar)

}

// If enabled, changes to @editable will be saved for undo/redo
// actions.
//
// This results in an additional copy of text changes and are not
// stored in secure memory. As such, undo is forcefully disabled
// when [property@Gtk.Text:visibility] is set to %FALSE.
func (x *PasswordEntry) SetEnableUndo(EnableUndoVar bool) {

	XGtkEditableSetEnableUndo(x.GoPointer(), EnableUndoVar)

}

// Sets the desired maximum width in characters of @editable.
func (x *PasswordEntry) SetMaxWidthChars(NCharsVar int) {

	XGtkEditableSetMaxWidthChars(x.GoPointer(), NCharsVar)

}

// Sets the cursor position in the editable to the given value.
//
// The cursor is displayed before the character with the given (base 0)
// index in the contents of the editable. The value must be less than
// or equal to the number of characters in the editable. A value of -1
// indicates that the position should be set after the last character
// of the editable. Note that @position is in characters, not in bytes.
func (x *PasswordEntry) SetPosition(PositionVar int) {

	XGtkEditableSetPosition(x.GoPointer(), PositionVar)

}

// Sets the text in the editable to the given value.
//
// This is replacing the current contents.
func (x *PasswordEntry) SetText(TextVar string) {

	XGtkEditableSetText(x.GoPointer(), TextVar)

}

// Changes the size request of the editable to be about the
// right size for @n_chars characters.
//
// Note that it changes the size request, the size can still
// be affected by how you pack the widget into containers.
// If @n_chars is -1, the size reverts to the default size.
func (x *PasswordEntry) SetWidthChars(NCharsVar int) {

	XGtkEditableSetWidthChars(x.GoPointer(), NCharsVar)

}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GTK"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xNewPasswordEntry, lib, "gtk_password_entry_new")

	core.PuregoSafeRegister(&xPasswordEntryGetExtraMenu, lib, "gtk_password_entry_get_extra_menu")
	core.PuregoSafeRegister(&xPasswordEntryGetShowPeekIcon, lib, "gtk_password_entry_get_show_peek_icon")
	core.PuregoSafeRegister(&xPasswordEntrySetExtraMenu, lib, "gtk_password_entry_set_extra_menu")
	core.PuregoSafeRegister(&xPasswordEntrySetShowPeekIcon, lib, "gtk_password_entry_set_show_peek_icon")

}
