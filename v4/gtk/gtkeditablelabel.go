// Package gtk was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gtk

import (
	"unsafe"

	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/gobject"
)

type EditableLabelClass struct {
	ParentClass uintptr
}

func (x *EditableLabelClass) GoPointer() uintptr {
	return uintptr(unsafe.Pointer(x))
}

// A `GtkEditableLabel` is a label that allows users to
// edit the text by switching to an “edit mode”.
//
// ![An example GtkEditableLabel](editable-label.png)
//
// `GtkEditableLabel` does not have API of its own, but it
// implements the [iface@Gtk.Editable] interface.
//
// The default bindings for activating the edit mode is
// to click or press the Enter key. The default bindings
// for leaving the edit mode are the Enter key (to save
// the results) or the Escape key (to cancel the editing).
//
// # CSS nodes
//
// ```
// editablelabel[.editing]
// ╰── stack
//
//	├── label
//	╰── text
//
// ```
//
// `GtkEditableLabel` has a main node with the name editablelabel.
// When the entry is in editing mode, it gets the .editing style
// class.
//
// For all the subnodes added to the text node in various situations,
// see [class@Gtk.Text].
type EditableLabel struct {
	Widget
}

func EditableLabelNewFromInternalPtr(ptr uintptr) *EditableLabel {
	cls := &EditableLabel{}
	cls.Ptr = ptr
	return cls
}

var xNewEditableLabel func(string) uintptr

// Creates a new `GtkEditableLabel` widget.
func NewEditableLabel(StrVar string) *Widget {
	var cls *Widget

	cret := xNewEditableLabel(StrVar)

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &Widget{}
	cls.Ptr = cret
	return cls
}

var xEditableLabelGetEditing func(uintptr) bool

// Returns whether the label is currently in “editing mode”.
func (x *EditableLabel) GetEditing() bool {

	cret := xEditableLabelGetEditing(x.GoPointer())
	return cret
}

var xEditableLabelStartEditing func(uintptr)

// Switches the label into “editing mode”.
func (x *EditableLabel) StartEditing() {

	xEditableLabelStartEditing(x.GoPointer())

}

var xEditableLabelStopEditing func(uintptr, bool)

// Switches the label out of “editing mode”.
//
// If @commit is %TRUE, the resulting text is kept as the
// [property@Gtk.Editable:text] property value, otherwise the
// resulting text is discarded and the label will keep its
// previous [property@Gtk.Editable:text] property value.
func (x *EditableLabel) StopEditing(CommitVar bool) {

	xEditableLabelStopEditing(x.GoPointer(), CommitVar)

}

func (c *EditableLabel) GoPointer() uintptr {
	return c.Ptr
}

func (c *EditableLabel) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// Retrieves the `GtkAccessibleRole` for the given `GtkAccessible`.
func (x *EditableLabel) GetAccessibleRole() AccessibleRole {

	cret := XGtkAccessibleGetAccessibleRole(x.GoPointer())
	return cret
}

// Resets the accessible @property to its default value.
func (x *EditableLabel) ResetProperty(PropertyVar AccessibleProperty) {

	XGtkAccessibleResetProperty(x.GoPointer(), PropertyVar)

}

// Resets the accessible @relation to its default value.
func (x *EditableLabel) ResetRelation(RelationVar AccessibleRelation) {

	XGtkAccessibleResetRelation(x.GoPointer(), RelationVar)

}

// Resets the accessible @state to its default value.
func (x *EditableLabel) ResetState(StateVar AccessibleState) {

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
func (x *EditableLabel) UpdateProperty(FirstPropertyVar AccessibleProperty, varArgs ...interface{}) {

	XGtkAccessibleUpdateProperty(x.GoPointer(), FirstPropertyVar, varArgs...)

}

// Updates an array of accessible properties.
//
// This function should be called by `GtkWidget` types whenever an accessible
// property change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *EditableLabel) UpdatePropertyValue(NPropertiesVar int, PropertiesVar uintptr, ValuesVar uintptr) {

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
func (x *EditableLabel) UpdateRelation(FirstRelationVar AccessibleRelation, varArgs ...interface{}) {

	XGtkAccessibleUpdateRelation(x.GoPointer(), FirstRelationVar, varArgs...)

}

// Updates an array of accessible relations.
//
// This function should be called by `GtkWidget` types whenever an accessible
// relation change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *EditableLabel) UpdateRelationValue(NRelationsVar int, RelationsVar uintptr, ValuesVar uintptr) {

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
func (x *EditableLabel) UpdateState(FirstStateVar AccessibleState, varArgs ...interface{}) {

	XGtkAccessibleUpdateState(x.GoPointer(), FirstStateVar, varArgs...)

}

// Updates an array of accessible states.
//
// This function should be called by `GtkWidget` types whenever an accessible
// state change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *EditableLabel) UpdateStateValue(NStatesVar int, StatesVar uintptr, ValuesVar uintptr) {

	XGtkAccessibleUpdateStateValue(x.GoPointer(), NStatesVar, StatesVar, ValuesVar)

}

// Gets the ID of the @buildable object.
//
// `GtkBuilder` sets the name based on the ID attribute
// of the &lt;object&gt; tag used to construct the @buildable.
func (x *EditableLabel) GetBuildableId() string {

	cret := XGtkBuildableGetBuildableId(x.GoPointer())
	return cret
}

// Deletes the currently selected text of the editable.
//
// This call doesn’t do anything if there is no selected text.
func (x *EditableLabel) DeleteSelection() {

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
func (x *EditableLabel) DeleteText(StartPosVar int, EndPosVar int) {

	XGtkEditableDeleteText(x.GoPointer(), StartPosVar, EndPosVar)

}

// Undoes the setup done by [method@Gtk.Editable.init_delegate].
//
// This is a helper function that should be called from dispose,
// before removing the delegate object.
func (x *EditableLabel) FinishDelegate() {

	XGtkEditableFinishDelegate(x.GoPointer())

}

// Gets the alignment of the editable.
func (x *EditableLabel) GetAlignment() float32 {

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
func (x *EditableLabel) GetChars(StartPosVar int, EndPosVar int) string {

	cret := XGtkEditableGetChars(x.GoPointer(), StartPosVar, EndPosVar)
	return cret
}

// Gets the `GtkEditable` that @editable is delegating its
// implementation to.
//
// Typically, the delegate is a [class@Gtk.Text] widget.
func (x *EditableLabel) GetDelegate() *EditableBase {
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
func (x *EditableLabel) GetEditable() bool {

	cret := XGtkEditableGetEditable(x.GoPointer())
	return cret
}

// Gets if undo/redo actions are enabled for @editable
func (x *EditableLabel) GetEnableUndo() bool {

	cret := XGtkEditableGetEnableUndo(x.GoPointer())
	return cret
}

// Retrieves the desired maximum width of @editable, in characters.
func (x *EditableLabel) GetMaxWidthChars() int {

	cret := XGtkEditableGetMaxWidthChars(x.GoPointer())
	return cret
}

// Retrieves the current position of the cursor relative
// to the start of the content of the editable.
//
// Note that this position is in characters, not in bytes.
func (x *EditableLabel) GetPosition() int {

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
func (x *EditableLabel) GetSelectionBounds(StartPosVar int, EndPosVar int) bool {

	cret := XGtkEditableGetSelectionBounds(x.GoPointer(), StartPosVar, EndPosVar)
	return cret
}

// Retrieves the contents of @editable.
//
// The returned string is owned by GTK and must not be modified or freed.
func (x *EditableLabel) GetText() string {

	cret := XGtkEditableGetText(x.GoPointer())
	return cret
}

// Gets the number of characters of space reserved
// for the contents of the editable.
func (x *EditableLabel) GetWidthChars() int {

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
func (x *EditableLabel) InitDelegate() {

	XGtkEditableInitDelegate(x.GoPointer())

}

// Inserts @length bytes of @text into the contents of the
// widget, at position @position.
//
// Note that the position is in characters, not in bytes.
// The function updates @position to point after the newly
// inserted text.
func (x *EditableLabel) InsertText(TextVar string, LengthVar int, PositionVar int) {

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
func (x *EditableLabel) SelectRegion(StartPosVar int, EndPosVar int) {

	XGtkEditableSelectRegion(x.GoPointer(), StartPosVar, EndPosVar)

}

// Sets the alignment for the contents of the editable.
//
// This controls the horizontal positioning of the contents when
// the displayed text is shorter than the width of the editable.
func (x *EditableLabel) SetAlignment(XalignVar float32) {

	XGtkEditableSetAlignment(x.GoPointer(), XalignVar)

}

// Determines if the user can edit the text in the editable widget.
func (x *EditableLabel) SetEditable(IsEditableVar bool) {

	XGtkEditableSetEditable(x.GoPointer(), IsEditableVar)

}

// If enabled, changes to @editable will be saved for undo/redo
// actions.
//
// This results in an additional copy of text changes and are not
// stored in secure memory. As such, undo is forcefully disabled
// when [property@Gtk.Text:visibility] is set to %FALSE.
func (x *EditableLabel) SetEnableUndo(EnableUndoVar bool) {

	XGtkEditableSetEnableUndo(x.GoPointer(), EnableUndoVar)

}

// Sets the desired maximum width in characters of @editable.
func (x *EditableLabel) SetMaxWidthChars(NCharsVar int) {

	XGtkEditableSetMaxWidthChars(x.GoPointer(), NCharsVar)

}

// Sets the cursor position in the editable to the given value.
//
// The cursor is displayed before the character with the given (base 0)
// index in the contents of the editable. The value must be less than
// or equal to the number of characters in the editable. A value of -1
// indicates that the position should be set after the last character
// of the editable. Note that @position is in characters, not in bytes.
func (x *EditableLabel) SetPosition(PositionVar int) {

	XGtkEditableSetPosition(x.GoPointer(), PositionVar)

}

// Sets the text in the editable to the given value.
//
// This is replacing the current contents.
func (x *EditableLabel) SetText(TextVar string) {

	XGtkEditableSetText(x.GoPointer(), TextVar)

}

// Changes the size request of the editable to be about the
// right size for @n_chars characters.
//
// Note that it changes the size request, the size can still
// be affected by how you pack the widget into containers.
// If @n_chars is -1, the size reverts to the default size.
func (x *EditableLabel) SetWidthChars(NCharsVar int) {

	XGtkEditableSetWidthChars(x.GoPointer(), NCharsVar)

}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GTK"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xNewEditableLabel, lib, "gtk_editable_label_new")

	core.PuregoSafeRegister(&xEditableLabelGetEditing, lib, "gtk_editable_label_get_editing")
	core.PuregoSafeRegister(&xEditableLabelStartEditing, lib, "gtk_editable_label_start_editing")
	core.PuregoSafeRegister(&xEditableLabelStopEditing, lib, "gtk_editable_label_stop_editing")

}
