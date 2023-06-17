// Package gtk was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gtk

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/gdk"
	"github.com/jwijenbergh/puregotk/v4/glib"
	"github.com/jwijenbergh/puregotk/v4/gobject"
)

// A `GtkComboBoxText` is a simple variant of `GtkComboBox` for text-only
// use cases.
//
// ![An example GtkComboBoxText](combo-box-text.png)
//
// `GtkComboBoxText` hides the model-view complexity of `GtkComboBox`.
//
// To create a `GtkComboBoxText`, use [ctor@Gtk.ComboBoxText.new] or
// [ctor@Gtk.ComboBoxText.new_with_entry].
//
// You can add items to a `GtkComboBoxText` with
// [method@Gtk.ComboBoxText.append_text],
// [method@Gtk.ComboBoxText.insert_text] or
// [method@Gtk.ComboBoxText.prepend_text] and remove options with
// [method@Gtk.ComboBoxText.remove].
//
// If the `GtkComboBoxText` contains an entry (via the
// [property@Gtk.ComboBox:has-entry] property), its contents can be retrieved
// using [method@Gtk.ComboBoxText.get_active_text].
//
// You should not call [method@Gtk.ComboBox.set_model] or attempt to pack more
// cells into this combo box via its [iface@Gtk.CellLayout] interface.
//
// # GtkComboBoxText as GtkBuildable
//
// The `GtkComboBoxText` implementation of the `GtkBuildable` interface supports
// adding items directly using the &lt;items&gt; element and specifying &lt;item&gt;
// elements for each item. Each &lt;item&gt; element can specify the “id”
// corresponding to the appended text and also supports the regular
// translation attributes “translatable”, “context” and “comments”.
//
// Here is a UI definition fragment specifying `GtkComboBoxText` items:
// ```xml
// &lt;object class="GtkComboBoxText"&gt;
//
//	&lt;items&gt;
//	  &lt;item translatable="yes" id="factory"&gt;Factory&lt;/item&gt;
//	  &lt;item translatable="yes" id="home"&gt;Home&lt;/item&gt;
//	  &lt;item translatable="yes" id="subway"&gt;Subway&lt;/item&gt;
//	&lt;/items&gt;
//
// &lt;/object&gt;
// ```
//
// # CSS nodes
//
// ```
// combobox
// ╰── box.linked
//
//	├── entry.combo
//	├── button.combo
//	╰── window.popup
//
// ```
//
// `GtkComboBoxText` has a single CSS node with name combobox. It adds
// the style class .combo to the main CSS nodes of its entry and button
// children, and the .linked class to the node of its internal box.
type ComboBoxText struct {
	ComboBox
}

func ComboBoxTextNewFromInternalPtr(ptr uintptr) *ComboBoxText {
	cls := &ComboBoxText{}
	cls.Ptr = ptr
	return cls
}

var xNewComboBoxText func() uintptr

// Creates a new `GtkComboBoxText`.
func NewComboBoxText() *Widget {
	NewComboBoxTextPtr := xNewComboBoxText()
	if NewComboBoxTextPtr == 0 {
		return nil
	}

	gobject.IncreaseRef(NewComboBoxTextPtr)

	NewComboBoxTextCls := &Widget{}
	NewComboBoxTextCls.Ptr = NewComboBoxTextPtr
	return NewComboBoxTextCls
}

var xNewWithEntryComboBoxText func() uintptr

// Creates a new `GtkComboBoxText` with an entry.
func NewWithEntryComboBoxText() *Widget {
	NewWithEntryComboBoxTextPtr := xNewWithEntryComboBoxText()
	if NewWithEntryComboBoxTextPtr == 0 {
		return nil
	}

	gobject.IncreaseRef(NewWithEntryComboBoxTextPtr)

	NewWithEntryComboBoxTextCls := &Widget{}
	NewWithEntryComboBoxTextCls.Ptr = NewWithEntryComboBoxTextPtr
	return NewWithEntryComboBoxTextCls
}

var xComboBoxTextAppend func(uintptr, string, string)

// Appends @text to the list of strings stored in @combo_box.
//
// If @id is non-%NULL then it is used as the ID of the row.
//
// This is the same as calling [method@Gtk.ComboBoxText.insert]
// with a position of -1.
func (x *ComboBoxText) Append(IdVar string, TextVar string) {

	xComboBoxTextAppend(x.GoPointer(), IdVar, TextVar)

}

var xComboBoxTextAppendText func(uintptr, string)

// Appends @text to the list of strings stored in @combo_box.
//
// This is the same as calling [method@Gtk.ComboBoxText.insert_text]
// with a position of -1.
func (x *ComboBoxText) AppendText(TextVar string) {

	xComboBoxTextAppendText(x.GoPointer(), TextVar)

}

var xComboBoxTextGetActiveText func(uintptr) string

// Returns the currently active string in @combo_box.
//
// If no row is currently selected, %NULL is returned.
// If @combo_box contains an entry, this function will
// return its contents (which will not necessarily
// be an item from the list).
func (x *ComboBoxText) GetActiveText() string {

	return xComboBoxTextGetActiveText(x.GoPointer())

}

var xComboBoxTextInsert func(uintptr, int32, string, string)

// Inserts @text at @position in the list of strings stored in @combo_box.
//
// If @id is non-%NULL then it is used as the ID of the row.
// See [property@Gtk.ComboBox:id-column].
//
// If @position is negative then @text is appended.
func (x *ComboBoxText) Insert(PositionVar int32, IdVar string, TextVar string) {

	xComboBoxTextInsert(x.GoPointer(), PositionVar, IdVar, TextVar)

}

var xComboBoxTextInsertText func(uintptr, int32, string)

// Inserts @text at @position in the list of strings stored in @combo_box.
//
// If @position is negative then @text is appended.
//
// This is the same as calling [method@Gtk.ComboBoxText.insert]
// with a %NULL ID string.
func (x *ComboBoxText) InsertText(PositionVar int32, TextVar string) {

	xComboBoxTextInsertText(x.GoPointer(), PositionVar, TextVar)

}

var xComboBoxTextPrepend func(uintptr, string, string)

// Prepends @text to the list of strings stored in @combo_box.
//
// If @id is non-%NULL then it is used as the ID of the row.
//
// This is the same as calling [method@Gtk.ComboBoxText.insert]
// with a position of 0.
func (x *ComboBoxText) Prepend(IdVar string, TextVar string) {

	xComboBoxTextPrepend(x.GoPointer(), IdVar, TextVar)

}

var xComboBoxTextPrependText func(uintptr, string)

// Prepends @text to the list of strings stored in @combo_box.
//
// This is the same as calling [method@Gtk.ComboBoxText.insert_text]
// with a position of 0.
func (x *ComboBoxText) PrependText(TextVar string) {

	xComboBoxTextPrependText(x.GoPointer(), TextVar)

}

var xComboBoxTextRemove func(uintptr, int32)

// Removes the string at @position from @combo_box.
func (x *ComboBoxText) Remove(PositionVar int32) {

	xComboBoxTextRemove(x.GoPointer(), PositionVar)

}

var xComboBoxTextRemoveAll func(uintptr)

// Removes all the text entries from the combo box.
func (x *ComboBoxText) RemoveAll() {

	xComboBoxTextRemoveAll(x.GoPointer())

}

func (c *ComboBoxText) GoPointer() uintptr {
	return c.Ptr
}

func (c *ComboBoxText) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// Retrieves the `GtkAccessibleRole` for the given `GtkAccessible`.
func (x *ComboBoxText) GetAccessibleRole() AccessibleRole {

	return XGtkAccessibleGetAccessibleRole(x.GoPointer())

}

// Resets the accessible @property to its default value.
func (x *ComboBoxText) ResetProperty(PropertyVar AccessibleProperty) {

	XGtkAccessibleResetProperty(x.GoPointer(), PropertyVar)

}

// Resets the accessible @relation to its default value.
func (x *ComboBoxText) ResetRelation(RelationVar AccessibleRelation) {

	XGtkAccessibleResetRelation(x.GoPointer(), RelationVar)

}

// Resets the accessible @state to its default value.
func (x *ComboBoxText) ResetState(StateVar AccessibleState) {

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
func (x *ComboBoxText) UpdateProperty(FirstPropertyVar AccessibleProperty, varArgs ...interface{}) {

	XGtkAccessibleUpdateProperty(x.GoPointer(), FirstPropertyVar, varArgs...)

}

// Updates an array of accessible properties.
//
// This function should be called by `GtkWidget` types whenever an accessible
// property change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *ComboBoxText) UpdatePropertyValue(NPropertiesVar int32, PropertiesVar uintptr, ValuesVar uintptr) {

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
func (x *ComboBoxText) UpdateRelation(FirstRelationVar AccessibleRelation, varArgs ...interface{}) {

	XGtkAccessibleUpdateRelation(x.GoPointer(), FirstRelationVar, varArgs...)

}

// Updates an array of accessible relations.
//
// This function should be called by `GtkWidget` types whenever an accessible
// relation change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *ComboBoxText) UpdateRelationValue(NRelationsVar int32, RelationsVar uintptr, ValuesVar uintptr) {

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
func (x *ComboBoxText) UpdateState(FirstStateVar AccessibleState, varArgs ...interface{}) {

	XGtkAccessibleUpdateState(x.GoPointer(), FirstStateVar, varArgs...)

}

// Updates an array of accessible states.
//
// This function should be called by `GtkWidget` types whenever an accessible
// state change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *ComboBoxText) UpdateStateValue(NStatesVar int32, StatesVar uintptr, ValuesVar uintptr) {

	XGtkAccessibleUpdateStateValue(x.GoPointer(), NStatesVar, StatesVar, ValuesVar)

}

// Gets the ID of the @buildable object.
//
// `GtkBuilder` sets the name based on the ID attribute
// of the &lt;object&gt; tag used to construct the @buildable.
func (x *ComboBoxText) GetBuildableId() string {

	return XGtkBuildableGetBuildableId(x.GoPointer())

}

// Emits the `GtkCellEditable::editing-done` signal.
func (x *ComboBoxText) EditingDone() {

	XGtkCellEditableEditingDone(x.GoPointer())

}

// Emits the `GtkCellEditable::remove-widget` signal.
func (x *ComboBoxText) RemoveWidget() {

	XGtkCellEditableRemoveWidget(x.GoPointer())

}

// Begins editing on a @cell_editable.
//
// The `GtkCellRenderer` for the cell creates and returns a `GtkCellEditable` from
// gtk_cell_renderer_start_editing(), configured for the `GtkCellRenderer` type.
//
// gtk_cell_editable_start_editing() can then set up @cell_editable suitably for
// editing a cell, e.g. making the Esc key emit `GtkCellEditable::editing-done`.
//
// Note that the @cell_editable is created on-demand for the current edit; its
// lifetime is temporary and does not persist across other edits and/or cells.
func (x *ComboBoxText) StartEditing(EventVar *gdk.Event) {

	XGtkCellEditableStartEditing(x.GoPointer(), EventVar.GoPointer())

}

// Adds an attribute mapping to the list in @cell_layout.
//
// The @column is the column of the model to get a value from, and the
// @attribute is the property on @cell to be set from that value. So for
// example if column 2 of the model contains strings, you could have the
// “text” attribute of a `GtkCellRendererText` get its values from column 2.
// In this context "attribute" and "property" are used interchangeably.
func (x *ComboBoxText) AddAttribute(CellVar *CellRenderer, AttributeVar string, ColumnVar int32) {

	XGtkCellLayoutAddAttribute(x.GoPointer(), CellVar.GoPointer(), AttributeVar, ColumnVar)

}

// Unsets all the mappings on all renderers on @cell_layout and
// removes all renderers from @cell_layout.
func (x *ComboBoxText) Clear() {

	XGtkCellLayoutClear(x.GoPointer())

}

// Clears all existing attributes previously set with
// gtk_cell_layout_set_attributes().
func (x *ComboBoxText) ClearAttributes(CellVar *CellRenderer) {

	XGtkCellLayoutClearAttributes(x.GoPointer(), CellVar.GoPointer())

}

// Returns the underlying `GtkCellArea` which might be @cell_layout
// if called on a `GtkCellArea` or might be %NULL if no `GtkCellArea`
// is used by @cell_layout.
func (x *ComboBoxText) GetArea() *CellArea {

	GetAreaPtr := XGtkCellLayoutGetArea(x.GoPointer())
	if GetAreaPtr == 0 {
		return nil
	}

	gobject.IncreaseRef(GetAreaPtr)

	GetAreaCls := &CellArea{}
	GetAreaCls.Ptr = GetAreaPtr
	return GetAreaCls

}

// Returns the cell renderers which have been added to @cell_layout.
func (x *ComboBoxText) GetCells() *glib.List {

	return XGtkCellLayoutGetCells(x.GoPointer())

}

// Adds the @cell to the end of @cell_layout. If @expand is %FALSE, then the
// @cell is allocated no more space than it needs. Any unused space is
// divided evenly between cells for which @expand is %TRUE.
//
// Note that reusing the same cell renderer is not supported.
func (x *ComboBoxText) PackEnd(CellVar *CellRenderer, ExpandVar bool) {

	XGtkCellLayoutPackEnd(x.GoPointer(), CellVar.GoPointer(), ExpandVar)

}

// Packs the @cell into the beginning of @cell_layout. If @expand is %FALSE,
// then the @cell is allocated no more space than it needs. Any unused space
// is divided evenly between cells for which @expand is %TRUE.
//
// Note that reusing the same cell renderer is not supported.
func (x *ComboBoxText) PackStart(CellVar *CellRenderer, ExpandVar bool) {

	XGtkCellLayoutPackStart(x.GoPointer(), CellVar.GoPointer(), ExpandVar)

}

// Re-inserts @cell at @position.
//
// Note that @cell has already to be packed into @cell_layout
// for this to function properly.
func (x *ComboBoxText) Reorder(CellVar *CellRenderer, PositionVar int32) {

	XGtkCellLayoutReorder(x.GoPointer(), CellVar.GoPointer(), PositionVar)

}

// Sets the attributes in the parameter list as the attributes
// of @cell_layout.
//
// See [method@Gtk.CellLayout.add_attribute] for more details.
//
// The attributes should be in attribute/column order, as in
// gtk_cell_layout_add_attribute(). All existing attributes are
// removed, and replaced with the new attributes.
func (x *ComboBoxText) SetAttributes(CellVar *CellRenderer, varArgs ...interface{}) {

	XGtkCellLayoutSetAttributes(x.GoPointer(), CellVar.GoPointer(), varArgs...)

}

// Sets the `GtkCellLayout`DataFunc to use for @cell_layout.
//
// This function is used instead of the standard attributes mapping
// for setting the column value, and should set the value of @cell_layout’s
// cell renderer(s) as appropriate.
//
// @func may be %NULL to remove a previously set function.
func (x *ComboBoxText) SetCellDataFunc(CellVar *CellRenderer, FuncVar CellLayoutDataFunc, FuncDataVar uintptr, DestroyVar glib.DestroyNotify) {

	XGtkCellLayoutSetCellDataFunc(x.GoPointer(), CellVar.GoPointer(), purego.NewCallback(FuncVar), FuncDataVar, purego.NewCallback(DestroyVar))

}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GTK"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xNewComboBoxText, lib, "gtk_combo_box_text_new")
	core.PuregoSafeRegister(&xNewWithEntryComboBoxText, lib, "gtk_combo_box_text_new_with_entry")

	core.PuregoSafeRegister(&xComboBoxTextAppend, lib, "gtk_combo_box_text_append")
	core.PuregoSafeRegister(&xComboBoxTextAppendText, lib, "gtk_combo_box_text_append_text")
	core.PuregoSafeRegister(&xComboBoxTextGetActiveText, lib, "gtk_combo_box_text_get_active_text")
	core.PuregoSafeRegister(&xComboBoxTextInsert, lib, "gtk_combo_box_text_insert")
	core.PuregoSafeRegister(&xComboBoxTextInsertText, lib, "gtk_combo_box_text_insert_text")
	core.PuregoSafeRegister(&xComboBoxTextPrepend, lib, "gtk_combo_box_text_prepend")
	core.PuregoSafeRegister(&xComboBoxTextPrependText, lib, "gtk_combo_box_text_prepend_text")
	core.PuregoSafeRegister(&xComboBoxTextRemove, lib, "gtk_combo_box_text_remove")
	core.PuregoSafeRegister(&xComboBoxTextRemoveAll, lib, "gtk_combo_box_text_remove_all")

}
