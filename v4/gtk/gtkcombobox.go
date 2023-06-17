// Package gtk was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gtk

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/gdk"
	"github.com/jwijenbergh/puregotk/v4/glib"
	"github.com/jwijenbergh/puregotk/v4/gobject"
)

type ComboBoxClass struct {
	ParentClass uintptr

	Padding uintptr
}

// A `GtkComboBox` is a widget that allows the user to choose from a list of
// valid choices.
//
// ![An example GtkComboBox](combo-box.png)
//
// The `GtkComboBox` displays the selected choice; when activated, the
// `GtkComboBox` displays a popup which allows the user to make a new choice.
//
// The `GtkComboBox` uses the model-view pattern; the list of valid choices
// is specified in the form of a tree model, and the display of the choices
// can be adapted to the data in the model by using cell renderers, as you
// would in a tree view. This is possible since `GtkComboBox` implements the
// [iface@Gtk.CellLayout] interface. The tree model holding the valid
// choices is not restricted to a flat list, it can be a real tree, and the
// popup will reflect the tree structure.
//
// To allow the user to enter values not in the model, the
// [property@Gtk.ComboBox:has-entry] property allows the `GtkComboBox` to
// contain a [class@Gtk.Entry]. This entry can be accessed by calling
// [method@Gtk.ComboBox.get_child] on the combo box.
//
// For a simple list of textual choices, the model-view API of `GtkComboBox`
// can be a bit overwhelming. In this case, [class@Gtk.ComboBoxText] offers
// a simple alternative. Both `GtkComboBox` and `GtkComboBoxText` can contain
// an entry.
//
// ## CSS nodes
//
// ```
// combobox
// ├── box.linked
// │   ╰── button.combo
// │       ╰── box
// │           ├── cellview
// │           ╰── arrow
// ╰── window.popup
// ```
//
// A normal combobox contains a box with the .linked class, a button
// with the .combo class and inside those buttons, there are a cellview and
// an arrow.
//
// ```
// combobox
// ├── box.linked
// │   ├── entry.combo
// │   ╰── button.combo
// │       ╰── box
// │           ╰── arrow
// ╰── window.popup
// ```
//
// A `GtkComboBox` with an entry has a single CSS node with name combobox.
// It contains a box with the .linked class. That box contains an entry and
// a button, both with the .combo class added. The button also contains another
// node with name arrow.
//
// # Accessibility
//
// `GtkComboBox` uses the %GTK_ACCESSIBLE_ROLE_COMBO_BOX role.
type ComboBox struct {
	Widget
}

func ComboBoxNewFromInternalPtr(ptr uintptr) *ComboBox {
	cls := &ComboBox{}
	cls.Ptr = ptr
	return cls
}

var xNewComboBox func() uintptr

// Creates a new empty `GtkComboBox`.
func NewComboBox() *Widget {
	NewComboBoxPtr := xNewComboBox()
	if NewComboBoxPtr == 0 {
		return nil
	}

	gobject.IncreaseRef(NewComboBoxPtr)

	NewComboBoxCls := &Widget{}
	NewComboBoxCls.Ptr = NewComboBoxPtr
	return NewComboBoxCls
}

var xNewWithEntryComboBox func() uintptr

// Creates a new empty `GtkComboBox` with an entry.
//
// In order to use a combo box with entry, you need to tell it
// which column of the model contains the text for the entry
// by calling [method@Gtk.ComboBox.set_entry_text_column].
func NewWithEntryComboBox() *Widget {
	NewWithEntryComboBoxPtr := xNewWithEntryComboBox()
	if NewWithEntryComboBoxPtr == 0 {
		return nil
	}

	gobject.IncreaseRef(NewWithEntryComboBoxPtr)

	NewWithEntryComboBoxCls := &Widget{}
	NewWithEntryComboBoxCls.Ptr = NewWithEntryComboBoxPtr
	return NewWithEntryComboBoxCls
}

var xNewWithModelComboBox func(uintptr) uintptr

// Creates a new `GtkComboBox` with a model.
func NewWithModelComboBox(ModelVar TreeModel) *Widget {
	NewWithModelComboBoxPtr := xNewWithModelComboBox(ModelVar.GoPointer())
	if NewWithModelComboBoxPtr == 0 {
		return nil
	}

	gobject.IncreaseRef(NewWithModelComboBoxPtr)

	NewWithModelComboBoxCls := &Widget{}
	NewWithModelComboBoxCls.Ptr = NewWithModelComboBoxPtr
	return NewWithModelComboBoxCls
}

var xNewWithModelAndEntryComboBox func(uintptr) uintptr

// Creates a new empty `GtkComboBox` with an entry and a model.
//
// See also [ctor@Gtk.ComboBox.new_with_entry].
func NewWithModelAndEntryComboBox(ModelVar TreeModel) *Widget {
	NewWithModelAndEntryComboBoxPtr := xNewWithModelAndEntryComboBox(ModelVar.GoPointer())
	if NewWithModelAndEntryComboBoxPtr == 0 {
		return nil
	}

	gobject.IncreaseRef(NewWithModelAndEntryComboBoxPtr)

	NewWithModelAndEntryComboBoxCls := &Widget{}
	NewWithModelAndEntryComboBoxCls.Ptr = NewWithModelAndEntryComboBoxPtr
	return NewWithModelAndEntryComboBoxCls
}

var xComboBoxGetActive func(uintptr) int32

// Returns the index of the currently active item.
//
// If the model is a non-flat treemodel, and the active item is not
// an immediate child of the root of the tree, this function returns
// `gtk_tree_path_get_indices (path)[0]`, where `path` is the
// [struct@Gtk.TreePath] of the active item.
func (x *ComboBox) GetActive() int32 {

	return xComboBoxGetActive(x.GoPointer())

}

var xComboBoxGetActiveId func(uintptr) string

// Returns the ID of the active row of @combo_box.
//
// This value is taken from the active row and the column specified
// by the [property@Gtk.ComboBox:id-column] property of @combo_box
// (see [method@Gtk.ComboBox.set_id_column]).
//
// The returned value is an interned string which means that you can
// compare the pointer by value to other interned strings and that you
// must not free it.
//
// If the [property@Gtk.ComboBox:id-column] property of @combo_box is
// not set, or if no row is active, or if the active row has a %NULL
// ID value, then %NULL is returned.
func (x *ComboBox) GetActiveId() string {

	return xComboBoxGetActiveId(x.GoPointer())

}

var xComboBoxGetActiveIter func(uintptr, *TreeIter) bool

// Sets @iter to point to the currently active item.
//
// If no item is active, @iter is left unchanged.
func (x *ComboBox) GetActiveIter(IterVar *TreeIter) bool {

	return xComboBoxGetActiveIter(x.GoPointer(), IterVar)

}

var xComboBoxGetButtonSensitivity func(uintptr) SensitivityType

// Returns whether the combo box sets the dropdown button
// sensitive or not when there are no items in the model.
func (x *ComboBox) GetButtonSensitivity() SensitivityType {

	return xComboBoxGetButtonSensitivity(x.GoPointer())

}

var xComboBoxGetChild func(uintptr) uintptr

// Gets the child widget of @combo_box.
func (x *ComboBox) GetChild() *Widget {

	GetChildPtr := xComboBoxGetChild(x.GoPointer())
	if GetChildPtr == 0 {
		return nil
	}

	gobject.IncreaseRef(GetChildPtr)

	GetChildCls := &Widget{}
	GetChildCls.Ptr = GetChildPtr
	return GetChildCls

}

var xComboBoxGetEntryTextColumn func(uintptr) int32

// Returns the column which @combo_box is using to get the strings
// from to display in the internal entry.
func (x *ComboBox) GetEntryTextColumn() int32 {

	return xComboBoxGetEntryTextColumn(x.GoPointer())

}

var xComboBoxGetHasEntry func(uintptr) bool

// Returns whether the combo box has an entry.
func (x *ComboBox) GetHasEntry() bool {

	return xComboBoxGetHasEntry(x.GoPointer())

}

var xComboBoxGetIdColumn func(uintptr) int32

// Returns the column which @combo_box is using to get string IDs
// for values from.
func (x *ComboBox) GetIdColumn() int32 {

	return xComboBoxGetIdColumn(x.GoPointer())

}

var xComboBoxGetModel func(uintptr) uintptr

// Returns the `GtkTreeModel` of @combo_box.
func (x *ComboBox) GetModel() *TreeModelBase {

	GetModelPtr := xComboBoxGetModel(x.GoPointer())
	if GetModelPtr == 0 {
		return nil
	}

	gobject.IncreaseRef(GetModelPtr)

	GetModelCls := &TreeModelBase{}
	GetModelCls.Ptr = GetModelPtr
	return GetModelCls

}

var xComboBoxGetPopupFixedWidth func(uintptr) bool

// Gets whether the popup uses a fixed width.
func (x *ComboBox) GetPopupFixedWidth() bool {

	return xComboBoxGetPopupFixedWidth(x.GoPointer())

}

var xComboBoxGetRowSeparatorFunc func(uintptr) TreeViewRowSeparatorFunc

// Returns the current row separator function.
func (x *ComboBox) GetRowSeparatorFunc() TreeViewRowSeparatorFunc {

	return xComboBoxGetRowSeparatorFunc(x.GoPointer())

}

var xComboBoxPopdown func(uintptr)

// Hides the menu or dropdown list of @combo_box.
//
// This function is mostly intended for use by accessibility technologies;
// applications should have little use for it.
func (x *ComboBox) Popdown() {

	xComboBoxPopdown(x.GoPointer())

}

var xComboBoxPopup func(uintptr)

// Pops up the menu or dropdown list of @combo_box.
//
// This function is mostly intended for use by accessibility technologies;
// applications should have little use for it.
//
// Before calling this, @combo_box must be mapped, or nothing will happen.
func (x *ComboBox) Popup() {

	xComboBoxPopup(x.GoPointer())

}

var xComboBoxPopupForDevice func(uintptr, uintptr)

// Pops up the menu of @combo_box.
//
// Note that currently this does not do anything with the device, as it was
// previously only used for list-mode combo boxes, and those were removed
// in GTK 4. However, it is retained in case similar functionality is added
// back later.
func (x *ComboBox) PopupForDevice(DeviceVar *gdk.Device) {

	xComboBoxPopupForDevice(x.GoPointer(), DeviceVar.GoPointer())

}

var xComboBoxSetActive func(uintptr, int32)

// Sets the active item of @combo_box to be the item at @index.
func (x *ComboBox) SetActive(IndexVar int32) {

	xComboBoxSetActive(x.GoPointer(), IndexVar)

}

var xComboBoxSetActiveId func(uintptr, string) bool

// Changes the active row of @combo_box to the one that has an ID equal to
// @active_id.
//
// If @active_id is %NULL, the active row is unset. Rows having
// a %NULL ID string cannot be made active by this function.
//
// If the [property@Gtk.ComboBox:id-column] property of @combo_box is
// unset or if no row has the given ID then the function does nothing
// and returns %FALSE.
func (x *ComboBox) SetActiveId(ActiveIdVar string) bool {

	return xComboBoxSetActiveId(x.GoPointer(), ActiveIdVar)

}

var xComboBoxSetActiveIter func(uintptr, *TreeIter)

// Sets the current active item to be the one referenced by @iter.
//
// If @iter is %NULL, the active item is unset.
func (x *ComboBox) SetActiveIter(IterVar *TreeIter) {

	xComboBoxSetActiveIter(x.GoPointer(), IterVar)

}

var xComboBoxSetButtonSensitivity func(uintptr, SensitivityType)

// Sets whether the dropdown button of the combo box should update
// its sensitivity depending on the model contents.
func (x *ComboBox) SetButtonSensitivity(SensitivityVar SensitivityType) {

	xComboBoxSetButtonSensitivity(x.GoPointer(), SensitivityVar)

}

var xComboBoxSetChild func(uintptr, uintptr)

// Sets the child widget of @combo_box.
func (x *ComboBox) SetChild(ChildVar *Widget) {

	xComboBoxSetChild(x.GoPointer(), ChildVar.GoPointer())

}

var xComboBoxSetEntryTextColumn func(uintptr, int32)

// Sets the model column which @combo_box should use to get strings
// from to be @text_column.
//
// For this column no separate
// [class@Gtk.CellRenderer] is needed.
//
// The column @text_column in the model of @combo_box must be of
// type %G_TYPE_STRING.
//
// This is only relevant if @combo_box has been created with
// [property@Gtk.ComboBox:has-entry] as %TRUE.
func (x *ComboBox) SetEntryTextColumn(TextColumnVar int32) {

	xComboBoxSetEntryTextColumn(x.GoPointer(), TextColumnVar)

}

var xComboBoxSetIdColumn func(uintptr, int32)

// Sets the model column which @combo_box should use to get string IDs
// for values from.
//
// The column @id_column in the model of @combo_box must be of type
// %G_TYPE_STRING.
func (x *ComboBox) SetIdColumn(IdColumnVar int32) {

	xComboBoxSetIdColumn(x.GoPointer(), IdColumnVar)

}

var xComboBoxSetModel func(uintptr, uintptr)

// Sets the model used by @combo_box to be @model.
//
// Will unset a previously set model (if applicable). If model is %NULL,
// then it will unset the model.
//
// Note that this function does not clear the cell renderers, you have to
// call [method@Gtk.CellLayout.clear] yourself if you need to set up different
// cell renderers for the new model.
func (x *ComboBox) SetModel(ModelVar TreeModel) {

	xComboBoxSetModel(x.GoPointer(), ModelVar.GoPointer())

}

var xComboBoxSetPopupFixedWidth func(uintptr, bool)

// Specifies whether the popup’s width should be a fixed width.
//
// If @fixed is %TRUE, the popup's width is set to match the
// allocated width of the combo box.
func (x *ComboBox) SetPopupFixedWidth(FixedVar bool) {

	xComboBoxSetPopupFixedWidth(x.GoPointer(), FixedVar)

}

var xComboBoxSetRowSeparatorFunc func(uintptr, uintptr, uintptr, uintptr)

// Sets the row separator function, which is used to determine
// whether a row should be drawn as a separator.
//
// If the row separator function is %NULL, no separators are drawn.
// This is the default value.
func (x *ComboBox) SetRowSeparatorFunc(FuncVar TreeViewRowSeparatorFunc, DataVar uintptr, DestroyVar glib.DestroyNotify) {

	xComboBoxSetRowSeparatorFunc(x.GoPointer(), purego.NewCallback(FuncVar), DataVar, purego.NewCallback(DestroyVar))

}

func (c *ComboBox) GoPointer() uintptr {
	return c.Ptr
}

func (c *ComboBox) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// Emitted to when the combo box is activated.
//
// The `::activate` signal on `GtkComboBox` is an action signal and
// emitting it causes the combo box to pop up its dropdown.
func (x *ComboBox) ConnectActivate(cb func(ComboBox)) {
	fcb := func(clsPtr uintptr) {
		fa := ComboBox{}
		fa.Ptr = clsPtr

		cb(fa)

	}
	gobject.ObjectConnect(x.GoPointer(), "signal::activate", purego.NewCallback(fcb))
}

// Emitted when the active item is changed.
//
// The can be due to the user selecting a different item from the list,
// or due to a call to [method@Gtk.ComboBox.set_active_iter]. It will
// also be emitted while typing into the entry of a combo box with an entry.
func (x *ComboBox) ConnectChanged(cb func(ComboBox)) {
	fcb := func(clsPtr uintptr) {
		fa := ComboBox{}
		fa.Ptr = clsPtr

		cb(fa)

	}
	gobject.ObjectConnect(x.GoPointer(), "signal::changed", purego.NewCallback(fcb))
}

// Emitted to allow changing how the text in a combo box's entry is displayed.
//
// See [property@Gtk.ComboBox:has-entry].
//
// Connect a signal handler which returns an allocated string representing
// @path. That string will then be used to set the text in the combo box's
// entry. The default signal handler uses the text from the
// [property@Gtk.ComboBox:entry-text-column] model column.
//
// Here's an example signal handler which fetches data from the model and
// displays it in the entry.
// ```c
// static char *
// format_entry_text_callback (GtkComboBox *combo,
//
//	const char *path,
//	gpointer     user_data)
//
//	{
//	  GtkTreeIter iter;
//	  GtkTreeModel model;
//	  double       value;
//
//	  model = gtk_combo_box_get_model (combo);
//
//	  gtk_tree_model_get_iter_from_string (model, &amp;iter, path);
//	  gtk_tree_model_get (model, &amp;iter,
//	                      THE_DOUBLE_VALUE_COLUMN, &amp;value,
//	                      -1);
//
//	  return g_strdup_printf ("%g", value);
//	}
//
// ```
func (x *ComboBox) ConnectFormatEntryText(cb func(ComboBox, string) string) {
	fcb := func(clsPtr uintptr, PathVarp string) string {
		fa := ComboBox{}
		fa.Ptr = clsPtr

		return cb(fa, PathVarp)

	}
	gobject.ObjectConnect(x.GoPointer(), "signal::format-entry-text", purego.NewCallback(fcb))
}

// Emitted to move the active selection.
//
// This is an [keybinding signal](class.SignalAction.html).
func (x *ComboBox) ConnectMoveActive(cb func(ComboBox, ScrollType)) {
	fcb := func(clsPtr uintptr, ScrollTypeVarp ScrollType) {
		fa := ComboBox{}
		fa.Ptr = clsPtr

		cb(fa, ScrollTypeVarp)

	}
	gobject.ObjectConnect(x.GoPointer(), "signal::move-active", purego.NewCallback(fcb))
}

// Emitted to popdown the combo box list.
//
// This is an [keybinding signal](class.SignalAction.html).
//
// The default bindings for this signal are Alt+Up and Escape.
func (x *ComboBox) ConnectPopdown(cb func(ComboBox) bool) {
	fcb := func(clsPtr uintptr) bool {
		fa := ComboBox{}
		fa.Ptr = clsPtr

		return cb(fa)

	}
	gobject.ObjectConnect(x.GoPointer(), "signal::popdown", purego.NewCallback(fcb))
}

// Emitted to popup the combo box list.
//
// This is an [keybinding signal](class.SignalAction.html).
//
// The default binding for this signal is Alt+Down.
func (x *ComboBox) ConnectPopup(cb func(ComboBox)) {
	fcb := func(clsPtr uintptr) {
		fa := ComboBox{}
		fa.Ptr = clsPtr

		cb(fa)

	}
	gobject.ObjectConnect(x.GoPointer(), "signal::popup", purego.NewCallback(fcb))
}

// Retrieves the `GtkAccessibleRole` for the given `GtkAccessible`.
func (x *ComboBox) GetAccessibleRole() AccessibleRole {

	return XGtkAccessibleGetAccessibleRole(x.GoPointer())

}

// Resets the accessible @property to its default value.
func (x *ComboBox) ResetProperty(PropertyVar AccessibleProperty) {

	XGtkAccessibleResetProperty(x.GoPointer(), PropertyVar)

}

// Resets the accessible @relation to its default value.
func (x *ComboBox) ResetRelation(RelationVar AccessibleRelation) {

	XGtkAccessibleResetRelation(x.GoPointer(), RelationVar)

}

// Resets the accessible @state to its default value.
func (x *ComboBox) ResetState(StateVar AccessibleState) {

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
func (x *ComboBox) UpdateProperty(FirstPropertyVar AccessibleProperty, varArgs ...interface{}) {

	XGtkAccessibleUpdateProperty(x.GoPointer(), FirstPropertyVar, varArgs...)

}

// Updates an array of accessible properties.
//
// This function should be called by `GtkWidget` types whenever an accessible
// property change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *ComboBox) UpdatePropertyValue(NPropertiesVar int32, PropertiesVar uintptr, ValuesVar uintptr) {

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
func (x *ComboBox) UpdateRelation(FirstRelationVar AccessibleRelation, varArgs ...interface{}) {

	XGtkAccessibleUpdateRelation(x.GoPointer(), FirstRelationVar, varArgs...)

}

// Updates an array of accessible relations.
//
// This function should be called by `GtkWidget` types whenever an accessible
// relation change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *ComboBox) UpdateRelationValue(NRelationsVar int32, RelationsVar uintptr, ValuesVar uintptr) {

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
func (x *ComboBox) UpdateState(FirstStateVar AccessibleState, varArgs ...interface{}) {

	XGtkAccessibleUpdateState(x.GoPointer(), FirstStateVar, varArgs...)

}

// Updates an array of accessible states.
//
// This function should be called by `GtkWidget` types whenever an accessible
// state change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *ComboBox) UpdateStateValue(NStatesVar int32, StatesVar uintptr, ValuesVar uintptr) {

	XGtkAccessibleUpdateStateValue(x.GoPointer(), NStatesVar, StatesVar, ValuesVar)

}

// Gets the ID of the @buildable object.
//
// `GtkBuilder` sets the name based on the ID attribute
// of the &lt;object&gt; tag used to construct the @buildable.
func (x *ComboBox) GetBuildableId() string {

	return XGtkBuildableGetBuildableId(x.GoPointer())

}

// Emits the `GtkCellEditable::editing-done` signal.
func (x *ComboBox) EditingDone() {

	XGtkCellEditableEditingDone(x.GoPointer())

}

// Emits the `GtkCellEditable::remove-widget` signal.
func (x *ComboBox) RemoveWidget() {

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
func (x *ComboBox) StartEditing(EventVar *gdk.Event) {

	XGtkCellEditableStartEditing(x.GoPointer(), EventVar.GoPointer())

}

// Adds an attribute mapping to the list in @cell_layout.
//
// The @column is the column of the model to get a value from, and the
// @attribute is the property on @cell to be set from that value. So for
// example if column 2 of the model contains strings, you could have the
// “text” attribute of a `GtkCellRendererText` get its values from column 2.
// In this context "attribute" and "property" are used interchangeably.
func (x *ComboBox) AddAttribute(CellVar *CellRenderer, AttributeVar string, ColumnVar int32) {

	XGtkCellLayoutAddAttribute(x.GoPointer(), CellVar.GoPointer(), AttributeVar, ColumnVar)

}

// Unsets all the mappings on all renderers on @cell_layout and
// removes all renderers from @cell_layout.
func (x *ComboBox) Clear() {

	XGtkCellLayoutClear(x.GoPointer())

}

// Clears all existing attributes previously set with
// gtk_cell_layout_set_attributes().
func (x *ComboBox) ClearAttributes(CellVar *CellRenderer) {

	XGtkCellLayoutClearAttributes(x.GoPointer(), CellVar.GoPointer())

}

// Returns the underlying `GtkCellArea` which might be @cell_layout
// if called on a `GtkCellArea` or might be %NULL if no `GtkCellArea`
// is used by @cell_layout.
func (x *ComboBox) GetArea() *CellArea {

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
func (x *ComboBox) GetCells() *glib.List {

	return XGtkCellLayoutGetCells(x.GoPointer())

}

// Adds the @cell to the end of @cell_layout. If @expand is %FALSE, then the
// @cell is allocated no more space than it needs. Any unused space is
// divided evenly between cells for which @expand is %TRUE.
//
// Note that reusing the same cell renderer is not supported.
func (x *ComboBox) PackEnd(CellVar *CellRenderer, ExpandVar bool) {

	XGtkCellLayoutPackEnd(x.GoPointer(), CellVar.GoPointer(), ExpandVar)

}

// Packs the @cell into the beginning of @cell_layout. If @expand is %FALSE,
// then the @cell is allocated no more space than it needs. Any unused space
// is divided evenly between cells for which @expand is %TRUE.
//
// Note that reusing the same cell renderer is not supported.
func (x *ComboBox) PackStart(CellVar *CellRenderer, ExpandVar bool) {

	XGtkCellLayoutPackStart(x.GoPointer(), CellVar.GoPointer(), ExpandVar)

}

// Re-inserts @cell at @position.
//
// Note that @cell has already to be packed into @cell_layout
// for this to function properly.
func (x *ComboBox) Reorder(CellVar *CellRenderer, PositionVar int32) {

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
func (x *ComboBox) SetAttributes(CellVar *CellRenderer, varArgs ...interface{}) {

	XGtkCellLayoutSetAttributes(x.GoPointer(), CellVar.GoPointer(), varArgs...)

}

// Sets the `GtkCellLayout`DataFunc to use for @cell_layout.
//
// This function is used instead of the standard attributes mapping
// for setting the column value, and should set the value of @cell_layout’s
// cell renderer(s) as appropriate.
//
// @func may be %NULL to remove a previously set function.
func (x *ComboBox) SetCellDataFunc(CellVar *CellRenderer, FuncVar CellLayoutDataFunc, FuncDataVar uintptr, DestroyVar glib.DestroyNotify) {

	XGtkCellLayoutSetCellDataFunc(x.GoPointer(), CellVar.GoPointer(), purego.NewCallback(FuncVar), FuncDataVar, purego.NewCallback(DestroyVar))

}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GTK"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xNewComboBox, lib, "gtk_combo_box_new")
	core.PuregoSafeRegister(&xNewWithEntryComboBox, lib, "gtk_combo_box_new_with_entry")
	core.PuregoSafeRegister(&xNewWithModelComboBox, lib, "gtk_combo_box_new_with_model")
	core.PuregoSafeRegister(&xNewWithModelAndEntryComboBox, lib, "gtk_combo_box_new_with_model_and_entry")

	core.PuregoSafeRegister(&xComboBoxGetActive, lib, "gtk_combo_box_get_active")
	core.PuregoSafeRegister(&xComboBoxGetActiveId, lib, "gtk_combo_box_get_active_id")
	core.PuregoSafeRegister(&xComboBoxGetActiveIter, lib, "gtk_combo_box_get_active_iter")
	core.PuregoSafeRegister(&xComboBoxGetButtonSensitivity, lib, "gtk_combo_box_get_button_sensitivity")
	core.PuregoSafeRegister(&xComboBoxGetChild, lib, "gtk_combo_box_get_child")
	core.PuregoSafeRegister(&xComboBoxGetEntryTextColumn, lib, "gtk_combo_box_get_entry_text_column")
	core.PuregoSafeRegister(&xComboBoxGetHasEntry, lib, "gtk_combo_box_get_has_entry")
	core.PuregoSafeRegister(&xComboBoxGetIdColumn, lib, "gtk_combo_box_get_id_column")
	core.PuregoSafeRegister(&xComboBoxGetModel, lib, "gtk_combo_box_get_model")
	core.PuregoSafeRegister(&xComboBoxGetPopupFixedWidth, lib, "gtk_combo_box_get_popup_fixed_width")
	core.PuregoSafeRegister(&xComboBoxGetRowSeparatorFunc, lib, "gtk_combo_box_get_row_separator_func")
	core.PuregoSafeRegister(&xComboBoxPopdown, lib, "gtk_combo_box_popdown")
	core.PuregoSafeRegister(&xComboBoxPopup, lib, "gtk_combo_box_popup")
	core.PuregoSafeRegister(&xComboBoxPopupForDevice, lib, "gtk_combo_box_popup_for_device")
	core.PuregoSafeRegister(&xComboBoxSetActive, lib, "gtk_combo_box_set_active")
	core.PuregoSafeRegister(&xComboBoxSetActiveId, lib, "gtk_combo_box_set_active_id")
	core.PuregoSafeRegister(&xComboBoxSetActiveIter, lib, "gtk_combo_box_set_active_iter")
	core.PuregoSafeRegister(&xComboBoxSetButtonSensitivity, lib, "gtk_combo_box_set_button_sensitivity")
	core.PuregoSafeRegister(&xComboBoxSetChild, lib, "gtk_combo_box_set_child")
	core.PuregoSafeRegister(&xComboBoxSetEntryTextColumn, lib, "gtk_combo_box_set_entry_text_column")
	core.PuregoSafeRegister(&xComboBoxSetIdColumn, lib, "gtk_combo_box_set_id_column")
	core.PuregoSafeRegister(&xComboBoxSetModel, lib, "gtk_combo_box_set_model")
	core.PuregoSafeRegister(&xComboBoxSetPopupFixedWidth, lib, "gtk_combo_box_set_popup_fixed_width")
	core.PuregoSafeRegister(&xComboBoxSetRowSeparatorFunc, lib, "gtk_combo_box_set_row_separator_func")

}
