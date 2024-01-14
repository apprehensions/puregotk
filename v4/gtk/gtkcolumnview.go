// Package gtk was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gtk

import (
	"unsafe"

	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/gio"
	"github.com/jwijenbergh/puregotk/v4/gobject"
)

type ColumnViewClass struct {
}

func (x *ColumnViewClass) GoPointer() uintptr {
	return uintptr(unsafe.Pointer(x))
}

// `GtkColumnView` presents a large dynamic list of items using multiple columns
// with headers.
//
// `GtkColumnView` uses the factories of its columns to generate a cell widget for
// each column, for each visible item and displays them together as the row for
// this item.
//
// The [property@Gtk.ColumnView:show-row-separators] and
// [property@Gtk.ColumnView:show-column-separators] properties offer a simple way
// to display separators between the rows or columns.
//
// `GtkColumnView` allows the user to select items according to the selection
// characteristics of the model. For models that allow multiple selected items,
// it is possible to turn on *rubberband selection*, using
// [property@Gtk.ColumnView:enable-rubberband].
//
// The column view supports sorting that can be customized by the user by
// clicking on column headers. To set this up, the `GtkSorter` returned by
// [method@Gtk.ColumnView.get_sorter] must be attached to a sort model for the
// data that the view is showing, and the columns must have sorters attached to
// them by calling [method@Gtk.ColumnViewColumn.set_sorter]. The initial sort
// order can be set with [method@Gtk.ColumnView.sort_by_column].
//
// The column view also supports interactive resizing and reordering of
// columns, via Drag-and-Drop of the column headers. This can be enabled or
// disabled with the [property@Gtk.ColumnView:reorderable] and
// [property@Gtk.ColumnViewColumn:resizable] properties.
//
// To learn more about the list widget framework, see the
// [overview](section-list-widget.html).
//
// # CSS nodes
//
// ```
// columnview[.column-separators][.rich-list][.navigation-sidebar][.data-table]
// ├── header
// │   ├── &lt;column header&gt;
// ┊   ┊
// │   ╰── &lt;column header&gt;
// │
// ├── listview
// │
// ┊
// ╰── [rubberband]
// ```
//
// `GtkColumnView` uses a single CSS node named columnview. It may carry the
// .column-separators style class, when [property@Gtk.ColumnView:show-column-separators]
// property is set. Header widgets appear below a node with name header.
// The rows are contained in a `GtkListView` widget, so there is a listview
// node with the same structure as for a standalone `GtkListView` widget.
// If [property@Gtk.ColumnView:show-row-separators] is set, it will be passed
// on to the list view, causing its CSS node to carry the .separators style class.
// For rubberband selection, a node with name rubberband is used.
//
// The main columnview node may also carry style classes to select
// the style of [list presentation](section-list-widget.html#list-styles):
// .rich-list, .navigation-sidebar or .data-table.
//
// # Accessibility
//
// `GtkColumnView` uses the %GTK_ACCESSIBLE_ROLE_TREE_GRID role, header title
// widgets are using the %GTK_ACCESSIBLE_ROLE_COLUMN_HEADER role. The row widgets
// are using the %GTK_ACCESSIBLE_ROLE_ROW role, and individual cells are using
// the %GTK_ACCESSIBLE_ROLE_GRID_CELL role
type ColumnView struct {
	Widget
}

func ColumnViewNewFromInternalPtr(ptr uintptr) *ColumnView {
	cls := &ColumnView{}
	cls.Ptr = ptr
	return cls
}

var xNewColumnView func(uintptr) uintptr

// Creates a new `GtkColumnView`.
//
// You most likely want to call [method@Gtk.ColumnView.append_column]
// to add columns next.
func NewColumnView(ModelVar SelectionModel) *ColumnView {
	var cls *ColumnView

	cret := xNewColumnView(ModelVar.GoPointer())

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &ColumnView{}
	cls.Ptr = cret
	return cls
}

var xColumnViewAppendColumn func(uintptr, uintptr)

// Appends the @column to the end of the columns in @self.
func (x *ColumnView) AppendColumn(ColumnVar *ColumnViewColumn) {

	xColumnViewAppendColumn(x.GoPointer(), ColumnVar.GoPointer())

}

var xColumnViewGetColumns func(uintptr) uintptr

// Gets the list of columns in this column view.
//
// This list is constant over the lifetime of @self and can be used to
// monitor changes to the columns of @self by connecting to the
// ::items-changed signal.
func (x *ColumnView) GetColumns() *gio.ListModelBase {
	var cls *gio.ListModelBase

	cret := xColumnViewGetColumns(x.GoPointer())

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &gio.ListModelBase{}
	cls.Ptr = cret
	return cls
}

var xColumnViewGetEnableRubberband func(uintptr) bool

// Returns whether rows can be selected by dragging with the mouse.
func (x *ColumnView) GetEnableRubberband() bool {

	cret := xColumnViewGetEnableRubberband(x.GoPointer())
	return cret
}

var xColumnViewGetModel func(uintptr) uintptr

// Gets the model that's currently used to read the items displayed.
func (x *ColumnView) GetModel() *SelectionModelBase {
	var cls *SelectionModelBase

	cret := xColumnViewGetModel(x.GoPointer())

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &SelectionModelBase{}
	cls.Ptr = cret
	return cls
}

var xColumnViewGetReorderable func(uintptr) bool

// Returns whether columns are reorderable.
func (x *ColumnView) GetReorderable() bool {

	cret := xColumnViewGetReorderable(x.GoPointer())
	return cret
}

var xColumnViewGetShowColumnSeparators func(uintptr) bool

// Returns whether the list should show separators
// between columns.
func (x *ColumnView) GetShowColumnSeparators() bool {

	cret := xColumnViewGetShowColumnSeparators(x.GoPointer())
	return cret
}

var xColumnViewGetShowRowSeparators func(uintptr) bool

// Returns whether the list should show separators
// between rows.
func (x *ColumnView) GetShowRowSeparators() bool {

	cret := xColumnViewGetShowRowSeparators(x.GoPointer())
	return cret
}

var xColumnViewGetSingleClickActivate func(uintptr) bool

// Returns whether rows will be activated on single click and
// selected on hover.
func (x *ColumnView) GetSingleClickActivate() bool {

	cret := xColumnViewGetSingleClickActivate(x.GoPointer())
	return cret
}

var xColumnViewGetSorter func(uintptr) uintptr

// Returns a special sorter that reflects the users sorting
// choices in the column view.
//
// To allow users to customizable sorting by clicking on column
// headers, this sorter needs to be set on the sort model underneath
// the model that is displayed by the view.
//
// See [method@Gtk.ColumnViewColumn.set_sorter] for setting up
// per-column sorting.
//
// Here is an example:
// ```c
// gtk_column_view_column_set_sorter (column, sorter);
// gtk_column_view_append_column (view, column);
// sorter = g_object_ref (gtk_column_view_get_sorter (view)));
// model = gtk_sort_list_model_new (store, sorter);
// selection = gtk_no_selection_new (model);
// gtk_column_view_set_model (view, selection);
// ```
func (x *ColumnView) GetSorter() *Sorter {
	var cls *Sorter

	cret := xColumnViewGetSorter(x.GoPointer())

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &Sorter{}
	cls.Ptr = cret
	return cls
}

var xColumnViewInsertColumn func(uintptr, uint, uintptr)

// Inserts a column at the given position in the columns of @self.
//
// If @column is already a column of @self, it will be repositioned.
func (x *ColumnView) InsertColumn(PositionVar uint, ColumnVar *ColumnViewColumn) {

	xColumnViewInsertColumn(x.GoPointer(), PositionVar, ColumnVar.GoPointer())

}

var xColumnViewRemoveColumn func(uintptr, uintptr)

// Removes the @column from the list of columns of @self.
func (x *ColumnView) RemoveColumn(ColumnVar *ColumnViewColumn) {

	xColumnViewRemoveColumn(x.GoPointer(), ColumnVar.GoPointer())

}

var xColumnViewSetEnableRubberband func(uintptr, bool)

// Sets whether selections can be changed by dragging with the mouse.
func (x *ColumnView) SetEnableRubberband(EnableRubberbandVar bool) {

	xColumnViewSetEnableRubberband(x.GoPointer(), EnableRubberbandVar)

}

var xColumnViewSetModel func(uintptr, uintptr)

// Sets the model to use.
//
// This must be a [iface@Gtk.SelectionModel].
func (x *ColumnView) SetModel(ModelVar SelectionModel) {

	xColumnViewSetModel(x.GoPointer(), ModelVar.GoPointer())

}

var xColumnViewSetReorderable func(uintptr, bool)

// Sets whether columns should be reorderable by dragging.
func (x *ColumnView) SetReorderable(ReorderableVar bool) {

	xColumnViewSetReorderable(x.GoPointer(), ReorderableVar)

}

var xColumnViewSetShowColumnSeparators func(uintptr, bool)

// Sets whether the list should show separators
// between columns.
func (x *ColumnView) SetShowColumnSeparators(ShowColumnSeparatorsVar bool) {

	xColumnViewSetShowColumnSeparators(x.GoPointer(), ShowColumnSeparatorsVar)

}

var xColumnViewSetShowRowSeparators func(uintptr, bool)

// Sets whether the list should show separators
// between rows.
func (x *ColumnView) SetShowRowSeparators(ShowRowSeparatorsVar bool) {

	xColumnViewSetShowRowSeparators(x.GoPointer(), ShowRowSeparatorsVar)

}

var xColumnViewSetSingleClickActivate func(uintptr, bool)

// Sets whether rows should be activated on single click and
// selected on hover.
func (x *ColumnView) SetSingleClickActivate(SingleClickActivateVar bool) {

	xColumnViewSetSingleClickActivate(x.GoPointer(), SingleClickActivateVar)

}

var xColumnViewSortByColumn func(uintptr, uintptr, SortType)

// Sets the sorting of the view.
//
// This function should be used to set up the initial sorting.
// At runtime, users can change the sorting of a column view
// by clicking on the list headers.
//
// This call only has an effect if the sorter returned by
// [method@Gtk.ColumnView.get_sorter] is set on a sort model,
// and [method@Gtk.ColumnViewColumn.set_sorter] has been called
// on @column to associate a sorter with the column.
//
// If @column is %NULL, the view will be unsorted.
func (x *ColumnView) SortByColumn(ColumnVar *ColumnViewColumn, DirectionVar SortType) {

	xColumnViewSortByColumn(x.GoPointer(), ColumnVar.GoPointer(), DirectionVar)

}

func (c *ColumnView) GoPointer() uintptr {
	return c.Ptr
}

func (c *ColumnView) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// Emitted when a row has been activated by the user, usually via activating
// the GtkListBase|list.activate-item action.
//
// This allows for a convenient way to handle activation in a columnview.
// See [method@Gtk.ListItem.set_activatable] for details on how to use this
// signal.
func (x *ColumnView) ConnectActivate(cb func(ColumnView, uint)) uint32 {
	fcb := func(clsPtr uintptr, PositionVarp uint) {
		fa := ColumnView{}
		fa.Ptr = clsPtr

		cb(fa, PositionVarp)

	}
	return gobject.SignalConnect(x.GoPointer(), "activate", purego.NewCallback(fcb))
}

// Retrieves the `GtkAccessibleRole` for the given `GtkAccessible`.
func (x *ColumnView) GetAccessibleRole() AccessibleRole {

	cret := XGtkAccessibleGetAccessibleRole(x.GoPointer())
	return cret
}

// Resets the accessible @property to its default value.
func (x *ColumnView) ResetProperty(PropertyVar AccessibleProperty) {

	XGtkAccessibleResetProperty(x.GoPointer(), PropertyVar)

}

// Resets the accessible @relation to its default value.
func (x *ColumnView) ResetRelation(RelationVar AccessibleRelation) {

	XGtkAccessibleResetRelation(x.GoPointer(), RelationVar)

}

// Resets the accessible @state to its default value.
func (x *ColumnView) ResetState(StateVar AccessibleState) {

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
func (x *ColumnView) UpdateProperty(FirstPropertyVar AccessibleProperty, varArgs ...interface{}) {

	XGtkAccessibleUpdateProperty(x.GoPointer(), FirstPropertyVar, varArgs...)

}

// Updates an array of accessible properties.
//
// This function should be called by `GtkWidget` types whenever an accessible
// property change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *ColumnView) UpdatePropertyValue(NPropertiesVar int, PropertiesVar uintptr, ValuesVar uintptr) {

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
func (x *ColumnView) UpdateRelation(FirstRelationVar AccessibleRelation, varArgs ...interface{}) {

	XGtkAccessibleUpdateRelation(x.GoPointer(), FirstRelationVar, varArgs...)

}

// Updates an array of accessible relations.
//
// This function should be called by `GtkWidget` types whenever an accessible
// relation change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *ColumnView) UpdateRelationValue(NRelationsVar int, RelationsVar uintptr, ValuesVar uintptr) {

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
func (x *ColumnView) UpdateState(FirstStateVar AccessibleState, varArgs ...interface{}) {

	XGtkAccessibleUpdateState(x.GoPointer(), FirstStateVar, varArgs...)

}

// Updates an array of accessible states.
//
// This function should be called by `GtkWidget` types whenever an accessible
// state change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *ColumnView) UpdateStateValue(NStatesVar int, StatesVar uintptr, ValuesVar uintptr) {

	XGtkAccessibleUpdateStateValue(x.GoPointer(), NStatesVar, StatesVar, ValuesVar)

}

// Gets the ID of the @buildable object.
//
// `GtkBuilder` sets the name based on the ID attribute
// of the &lt;object&gt; tag used to construct the @buildable.
func (x *ColumnView) GetBuildableId() string {

	cret := XGtkBuildableGetBuildableId(x.GoPointer())
	return cret
}

// Returns the size of a non-scrolling border around the
// outside of the scrollable.
//
// An example for this would be treeview headers. GTK can use
// this information to display overlaid graphics, like the
// overshoot indication, at the right position.
func (x *ColumnView) GetBorder(BorderVar *Border) bool {

	cret := XGtkScrollableGetBorder(x.GoPointer(), BorderVar)
	return cret
}

// Retrieves the `GtkAdjustment` used for horizontal scrolling.
func (x *ColumnView) GetHadjustment() *Adjustment {
	var cls *Adjustment

	cret := XGtkScrollableGetHadjustment(x.GoPointer())

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &Adjustment{}
	cls.Ptr = cret
	return cls
}

// Gets the horizontal `GtkScrollablePolicy`.
func (x *ColumnView) GetHscrollPolicy() ScrollablePolicy {

	cret := XGtkScrollableGetHscrollPolicy(x.GoPointer())
	return cret
}

// Retrieves the `GtkAdjustment` used for vertical scrolling.
func (x *ColumnView) GetVadjustment() *Adjustment {
	var cls *Adjustment

	cret := XGtkScrollableGetVadjustment(x.GoPointer())

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &Adjustment{}
	cls.Ptr = cret
	return cls
}

// Gets the vertical `GtkScrollablePolicy`.
func (x *ColumnView) GetVscrollPolicy() ScrollablePolicy {

	cret := XGtkScrollableGetVscrollPolicy(x.GoPointer())
	return cret
}

// Sets the horizontal adjustment of the `GtkScrollable`.
func (x *ColumnView) SetHadjustment(HadjustmentVar *Adjustment) {

	XGtkScrollableSetHadjustment(x.GoPointer(), HadjustmentVar.GoPointer())

}

// Sets the `GtkScrollablePolicy`.
//
// The policy determines whether horizontal scrolling should start
// below the minimum width or below the natural width.
func (x *ColumnView) SetHscrollPolicy(PolicyVar ScrollablePolicy) {

	XGtkScrollableSetHscrollPolicy(x.GoPointer(), PolicyVar)

}

// Sets the vertical adjustment of the `GtkScrollable`.
func (x *ColumnView) SetVadjustment(VadjustmentVar *Adjustment) {

	XGtkScrollableSetVadjustment(x.GoPointer(), VadjustmentVar.GoPointer())

}

// Sets the `GtkScrollablePolicy`.
//
// The policy determines whether vertical scrolling should start
// below the minimum height or below the natural height.
func (x *ColumnView) SetVscrollPolicy(PolicyVar ScrollablePolicy) {

	XGtkScrollableSetVscrollPolicy(x.GoPointer(), PolicyVar)

}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GTK"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xNewColumnView, lib, "gtk_column_view_new")

	core.PuregoSafeRegister(&xColumnViewAppendColumn, lib, "gtk_column_view_append_column")
	core.PuregoSafeRegister(&xColumnViewGetColumns, lib, "gtk_column_view_get_columns")
	core.PuregoSafeRegister(&xColumnViewGetEnableRubberband, lib, "gtk_column_view_get_enable_rubberband")
	core.PuregoSafeRegister(&xColumnViewGetModel, lib, "gtk_column_view_get_model")
	core.PuregoSafeRegister(&xColumnViewGetReorderable, lib, "gtk_column_view_get_reorderable")
	core.PuregoSafeRegister(&xColumnViewGetShowColumnSeparators, lib, "gtk_column_view_get_show_column_separators")
	core.PuregoSafeRegister(&xColumnViewGetShowRowSeparators, lib, "gtk_column_view_get_show_row_separators")
	core.PuregoSafeRegister(&xColumnViewGetSingleClickActivate, lib, "gtk_column_view_get_single_click_activate")
	core.PuregoSafeRegister(&xColumnViewGetSorter, lib, "gtk_column_view_get_sorter")
	core.PuregoSafeRegister(&xColumnViewInsertColumn, lib, "gtk_column_view_insert_column")
	core.PuregoSafeRegister(&xColumnViewRemoveColumn, lib, "gtk_column_view_remove_column")
	core.PuregoSafeRegister(&xColumnViewSetEnableRubberband, lib, "gtk_column_view_set_enable_rubberband")
	core.PuregoSafeRegister(&xColumnViewSetModel, lib, "gtk_column_view_set_model")
	core.PuregoSafeRegister(&xColumnViewSetReorderable, lib, "gtk_column_view_set_reorderable")
	core.PuregoSafeRegister(&xColumnViewSetShowColumnSeparators, lib, "gtk_column_view_set_show_column_separators")
	core.PuregoSafeRegister(&xColumnViewSetShowRowSeparators, lib, "gtk_column_view_set_show_row_separators")
	core.PuregoSafeRegister(&xColumnViewSetSingleClickActivate, lib, "gtk_column_view_set_single_click_activate")
	core.PuregoSafeRegister(&xColumnViewSortByColumn, lib, "gtk_column_view_sort_by_column")

}
