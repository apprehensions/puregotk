// Package gtk was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gtk

import (
	"unsafe"

	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/gio"
	"github.com/jwijenbergh/puregotk/v4/glib"
	"github.com/jwijenbergh/puregotk/v4/gobject"
)

// Called for flow boxes that are bound to a `GListModel`.
//
// This function is called for each item that gets added to the model.
type FlowBoxCreateWidgetFunc func(uintptr, uintptr) uintptr

// A function that will be called whenever a child changes
// or is added.
//
// It lets you control if the child should be visible or not.
type FlowBoxFilterFunc func(uintptr, uintptr) bool

// A function used by gtk_flow_box_selected_foreach().
//
// It will be called on every selected child of the @box.
type FlowBoxForeachFunc func(uintptr, uintptr, uintptr)

// A function to compare two children to determine which
// should come first.
type FlowBoxSortFunc func(uintptr, uintptr, uintptr) int

type FlowBoxChildClass struct {
	ParentClass uintptr

	Padding uintptr
}

func (x *FlowBoxChildClass) GoPointer() uintptr {
	return uintptr(unsafe.Pointer(x))
}

// A `GtkFlowBox` puts child widgets in reflowing grid.
//
// For instance, with the horizontal orientation, the widgets will be
// arranged from left to right, starting a new row under the previous
// row when necessary. Reducing the width in this case will require more
// rows, so a larger height will be requested.
//
// Likewise, with the vertical orientation, the widgets will be arranged
// from top to bottom, starting a new column to the right when necessary.
// Reducing the height will require more columns, so a larger width will
// be requested.
//
// The size request of a `GtkFlowBox` alone may not be what you expect;
// if you need to be able to shrink it along both axes and dynamically
// reflow its children, you may have to wrap it in a `GtkScrolledWindow`
// to enable that.
//
// The children of a `GtkFlowBox` can be dynamically sorted and filtered.
//
// Although a `GtkFlowBox` must have only `GtkFlowBoxChild` children, you
// can add any kind of widget to it via [method@Gtk.FlowBox.insert], and a
// `GtkFlowBoxChild` widget will automatically be inserted between the box
// and the widget.
//
// Also see [class@Gtk.ListBox].
//
// # CSS nodes
//
// ```
// flowbox
// ├── flowboxchild
// │   ╰── &lt;child&gt;
// ├── flowboxchild
// │   ╰── &lt;child&gt;
// ┊
// ╰── [rubberband]
// ```
//
// `GtkFlowBox` uses a single CSS node with name flowbox. `GtkFlowBoxChild`
// uses a single CSS node with name flowboxchild. For rubberband selection,
// a subnode with name rubberband is used.
//
// # Accessibility
//
// `GtkFlowBox` uses the %GTK_ACCESSIBLE_ROLE_GRID role, and `GtkFlowBoxChild`
// uses the %GTK_ACCESSIBLE_ROLE_GRID_CELL role.
type FlowBox struct {
	Widget
}

func FlowBoxNewFromInternalPtr(ptr uintptr) *FlowBox {
	cls := &FlowBox{}
	cls.Ptr = ptr
	return cls
}

var xNewFlowBox func() uintptr

// Creates a `GtkFlowBox`.
func NewFlowBox() *FlowBox {
	var cls *FlowBox

	cret := xNewFlowBox()

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &FlowBox{}
	cls.Ptr = cret
	return cls
}

var xFlowBoxAppend func(uintptr, uintptr)

// Adds @child to the end of @self.
//
// If a sort function is set, the widget will
// actually be inserted at the calculated position.
//
// See also: [method@Gtk.FlowBox.insert].
func (x *FlowBox) Append(ChildVar *Widget) {

	xFlowBoxAppend(x.GoPointer(), ChildVar.GoPointer())

}

var xFlowBoxBindModel func(uintptr, uintptr, uintptr, uintptr, uintptr)

// Binds @model to @box.
//
// If @box was already bound to a model, that previous binding is
// destroyed.
//
// The contents of @box are cleared and then filled with widgets that
// represent items from @model. @box is updated whenever @model changes.
// If @model is %NULL, @box is left empty.
//
// It is undefined to add or remove widgets directly (for example, with
// [method@Gtk.FlowBox.insert]) while @box is bound to a model.
//
// Note that using a model is incompatible with the filtering and sorting
// functionality in `GtkFlowBox`. When using a model, filtering and sorting
// should be implemented by the model.
func (x *FlowBox) BindModel(ModelVar gio.ListModel, CreateWidgetFuncVar *FlowBoxCreateWidgetFunc, UserDataVar uintptr, UserDataFreeFuncVar *glib.DestroyNotify) {

	xFlowBoxBindModel(x.GoPointer(), ModelVar.GoPointer(), glib.NewCallback(CreateWidgetFuncVar), UserDataVar, glib.NewCallback(UserDataFreeFuncVar))

}

var xFlowBoxGetActivateOnSingleClick func(uintptr) bool

// Returns whether children activate on single clicks.
func (x *FlowBox) GetActivateOnSingleClick() bool {

	cret := xFlowBoxGetActivateOnSingleClick(x.GoPointer())
	return cret
}

var xFlowBoxGetChildAtIndex func(uintptr, int) uintptr

// Gets the nth child in the @box.
func (x *FlowBox) GetChildAtIndex(IdxVar int) *FlowBoxChild {
	var cls *FlowBoxChild

	cret := xFlowBoxGetChildAtIndex(x.GoPointer(), IdxVar)

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &FlowBoxChild{}
	cls.Ptr = cret
	return cls
}

var xFlowBoxGetChildAtPos func(uintptr, int, int) uintptr

// Gets the child in the (@x, @y) position.
//
// Both @x and @y are assumed to be relative to the origin of @box.
func (x *FlowBox) GetChildAtPos(XVar int, YVar int) *FlowBoxChild {
	var cls *FlowBoxChild

	cret := xFlowBoxGetChildAtPos(x.GoPointer(), XVar, YVar)

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &FlowBoxChild{}
	cls.Ptr = cret
	return cls
}

var xFlowBoxGetColumnSpacing func(uintptr) uint

// Gets the horizontal spacing.
func (x *FlowBox) GetColumnSpacing() uint {

	cret := xFlowBoxGetColumnSpacing(x.GoPointer())
	return cret
}

var xFlowBoxGetHomogeneous func(uintptr) bool

// Returns whether the box is homogeneous.
func (x *FlowBox) GetHomogeneous() bool {

	cret := xFlowBoxGetHomogeneous(x.GoPointer())
	return cret
}

var xFlowBoxGetMaxChildrenPerLine func(uintptr) uint

// Gets the maximum number of children per line.
func (x *FlowBox) GetMaxChildrenPerLine() uint {

	cret := xFlowBoxGetMaxChildrenPerLine(x.GoPointer())
	return cret
}

var xFlowBoxGetMinChildrenPerLine func(uintptr) uint

// Gets the minimum number of children per line.
func (x *FlowBox) GetMinChildrenPerLine() uint {

	cret := xFlowBoxGetMinChildrenPerLine(x.GoPointer())
	return cret
}

var xFlowBoxGetRowSpacing func(uintptr) uint

// Gets the vertical spacing.
func (x *FlowBox) GetRowSpacing() uint {

	cret := xFlowBoxGetRowSpacing(x.GoPointer())
	return cret
}

var xFlowBoxGetSelectedChildren func(uintptr) *glib.List

// Creates a list of all selected children.
func (x *FlowBox) GetSelectedChildren() *glib.List {

	cret := xFlowBoxGetSelectedChildren(x.GoPointer())
	return cret
}

var xFlowBoxGetSelectionMode func(uintptr) SelectionMode

// Gets the selection mode of @box.
func (x *FlowBox) GetSelectionMode() SelectionMode {

	cret := xFlowBoxGetSelectionMode(x.GoPointer())
	return cret
}

var xFlowBoxInsert func(uintptr, uintptr, int)

// Inserts the @widget into @box at @position.
//
// If a sort function is set, the widget will actually be inserted
// at the calculated position.
//
// If @position is -1, or larger than the total number of children
// in the @box, then the @widget will be appended to the end.
func (x *FlowBox) Insert(WidgetVar *Widget, PositionVar int) {

	xFlowBoxInsert(x.GoPointer(), WidgetVar.GoPointer(), PositionVar)

}

var xFlowBoxInvalidateFilter func(uintptr)

// Updates the filtering for all children.
//
// Call this function when the result of the filter
// function on the @box is changed due ot an external
// factor. For instance, this would be used if the
// filter function just looked for a specific search
// term, and the entry with the string has changed.
func (x *FlowBox) InvalidateFilter() {

	xFlowBoxInvalidateFilter(x.GoPointer())

}

var xFlowBoxInvalidateSort func(uintptr)

// Updates the sorting for all children.
//
// Call this when the result of the sort function on
// @box is changed due to an external factor.
func (x *FlowBox) InvalidateSort() {

	xFlowBoxInvalidateSort(x.GoPointer())

}

var xFlowBoxPrepend func(uintptr, uintptr)

// Adds @child to the start of @self.
//
// If a sort function is set, the widget will
// actually be inserted at the calculated position.
//
// See also: [method@Gtk.FlowBox.insert].
func (x *FlowBox) Prepend(ChildVar *Widget) {

	xFlowBoxPrepend(x.GoPointer(), ChildVar.GoPointer())

}

var xFlowBoxRemove func(uintptr, uintptr)

// Removes a child from @box.
func (x *FlowBox) Remove(WidgetVar *Widget) {

	xFlowBoxRemove(x.GoPointer(), WidgetVar.GoPointer())

}

var xFlowBoxSelectAll func(uintptr)

// Select all children of @box, if the selection
// mode allows it.
func (x *FlowBox) SelectAll() {

	xFlowBoxSelectAll(x.GoPointer())

}

var xFlowBoxSelectChild func(uintptr, uintptr)

// Selects a single child of @box, if the selection
// mode allows it.
func (x *FlowBox) SelectChild(ChildVar *FlowBoxChild) {

	xFlowBoxSelectChild(x.GoPointer(), ChildVar.GoPointer())

}

var xFlowBoxSelectedForeach func(uintptr, uintptr, uintptr)

// Calls a function for each selected child.
//
// Note that the selection cannot be modified from within
// this function.
func (x *FlowBox) SelectedForeach(FuncVar *FlowBoxForeachFunc, DataVar uintptr) {

	xFlowBoxSelectedForeach(x.GoPointer(), glib.NewCallback(FuncVar), DataVar)

}

var xFlowBoxSetActivateOnSingleClick func(uintptr, bool)

// If @single is %TRUE, children will be activated when you click
// on them, otherwise you need to double-click.
func (x *FlowBox) SetActivateOnSingleClick(SingleVar bool) {

	xFlowBoxSetActivateOnSingleClick(x.GoPointer(), SingleVar)

}

var xFlowBoxSetColumnSpacing func(uintptr, uint)

// Sets the horizontal space to add between children.
func (x *FlowBox) SetColumnSpacing(SpacingVar uint) {

	xFlowBoxSetColumnSpacing(x.GoPointer(), SpacingVar)

}

var xFlowBoxSetFilterFunc func(uintptr, uintptr, uintptr, uintptr)

// By setting a filter function on the @box one can decide dynamically
// which of the children to show.
//
// For instance, to implement a search function that only shows the
// children matching the search terms.
//
// The @filter_func will be called for each child after the call, and
// it will continue to be called each time a child changes (via
// [method@Gtk.FlowBoxChild.changed]) or when
// [method@Gtk.FlowBox.invalidate_filter] is called.
//
// Note that using a filter function is incompatible with using a model
// (see [method@Gtk.FlowBox.bind_model]).
func (x *FlowBox) SetFilterFunc(FilterFuncVar *FlowBoxFilterFunc, UserDataVar uintptr, DestroyVar *glib.DestroyNotify) {

	xFlowBoxSetFilterFunc(x.GoPointer(), glib.NewCallback(FilterFuncVar), UserDataVar, glib.NewCallback(DestroyVar))

}

var xFlowBoxSetHadjustment func(uintptr, uintptr)

// Hooks up an adjustment to focus handling in @box.
//
// The adjustment is also used for autoscrolling during
// rubberband selection. See [method@Gtk.ScrolledWindow.get_hadjustment]
// for a typical way of obtaining the adjustment, and
// [method@Gtk.FlowBox.set_vadjustment] for setting the vertical
// adjustment.
//
// The adjustments have to be in pixel units and in the same
// coordinate system as the allocation for immediate children
// of the box.
func (x *FlowBox) SetHadjustment(AdjustmentVar *Adjustment) {

	xFlowBoxSetHadjustment(x.GoPointer(), AdjustmentVar.GoPointer())

}

var xFlowBoxSetHomogeneous func(uintptr, bool)

// Sets whether or not all children of @box are given
// equal space in the box.
func (x *FlowBox) SetHomogeneous(HomogeneousVar bool) {

	xFlowBoxSetHomogeneous(x.GoPointer(), HomogeneousVar)

}

var xFlowBoxSetMaxChildrenPerLine func(uintptr, uint)

// Sets the maximum number of children to request and
// allocate space for in @box’s orientation.
//
// Setting the maximum number of children per line
// limits the overall natural size request to be no more
// than @n_children children long in the given orientation.
func (x *FlowBox) SetMaxChildrenPerLine(NChildrenVar uint) {

	xFlowBoxSetMaxChildrenPerLine(x.GoPointer(), NChildrenVar)

}

var xFlowBoxSetMinChildrenPerLine func(uintptr, uint)

// Sets the minimum number of children to line up
// in @box’s orientation before flowing.
func (x *FlowBox) SetMinChildrenPerLine(NChildrenVar uint) {

	xFlowBoxSetMinChildrenPerLine(x.GoPointer(), NChildrenVar)

}

var xFlowBoxSetRowSpacing func(uintptr, uint)

// Sets the vertical space to add between children.
func (x *FlowBox) SetRowSpacing(SpacingVar uint) {

	xFlowBoxSetRowSpacing(x.GoPointer(), SpacingVar)

}

var xFlowBoxSetSelectionMode func(uintptr, SelectionMode)

// Sets how selection works in @box.
func (x *FlowBox) SetSelectionMode(ModeVar SelectionMode) {

	xFlowBoxSetSelectionMode(x.GoPointer(), ModeVar)

}

var xFlowBoxSetSortFunc func(uintptr, uintptr, uintptr, uintptr)

// By setting a sort function on the @box, one can dynamically
// reorder the children of the box, based on the contents of
// the children.
//
// The @sort_func will be called for each child after the call,
// and will continue to be called each time a child changes (via
// [method@Gtk.FlowBoxChild.changed]) and when
// [method@Gtk.FlowBox.invalidate_sort] is called.
//
// Note that using a sort function is incompatible with using a model
// (see [method@Gtk.FlowBox.bind_model]).
func (x *FlowBox) SetSortFunc(SortFuncVar *FlowBoxSortFunc, UserDataVar uintptr, DestroyVar *glib.DestroyNotify) {

	xFlowBoxSetSortFunc(x.GoPointer(), glib.NewCallback(SortFuncVar), UserDataVar, glib.NewCallback(DestroyVar))

}

var xFlowBoxSetVadjustment func(uintptr, uintptr)

// Hooks up an adjustment to focus handling in @box.
//
// The adjustment is also used for autoscrolling during
// rubberband selection. See [method@Gtk.ScrolledWindow.get_vadjustment]
// for a typical way of obtaining the adjustment, and
// [method@Gtk.FlowBox.set_hadjustment] for setting the horizontal
// adjustment.
//
// The adjustments have to be in pixel units and in the same
// coordinate system as the allocation for immediate children
// of the box.
func (x *FlowBox) SetVadjustment(AdjustmentVar *Adjustment) {

	xFlowBoxSetVadjustment(x.GoPointer(), AdjustmentVar.GoPointer())

}

var xFlowBoxUnselectAll func(uintptr)

// Unselect all children of @box, if the selection
// mode allows it.
func (x *FlowBox) UnselectAll() {

	xFlowBoxUnselectAll(x.GoPointer())

}

var xFlowBoxUnselectChild func(uintptr, uintptr)

// Unselects a single child of @box, if the selection
// mode allows it.
func (x *FlowBox) UnselectChild(ChildVar *FlowBoxChild) {

	xFlowBoxUnselectChild(x.GoPointer(), ChildVar.GoPointer())

}

func (c *FlowBox) GoPointer() uintptr {
	return c.Ptr
}

func (c *FlowBox) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// Emitted when the user activates the @box.
//
// This is a [keybinding signal](class.SignalAction.html).
func (x *FlowBox) ConnectActivateCursorChild(cb *func(FlowBox)) uint32 {
	cbPtr := uintptr(unsafe.Pointer(cb))
	if cbRefPtr, ok := glib.GetCallback(cbPtr); ok {
		return gobject.SignalConnect(x.GoPointer(), "activate-cursor-child", cbRefPtr)
	}

	fcb := func(clsPtr uintptr) {
		fa := FlowBox{}
		fa.Ptr = clsPtr
		cbFn := *cb

		cbFn(fa)

	}
	cbRefPtr := purego.NewCallback(fcb)
	glib.SaveCallback(cbPtr, cbRefPtr)
	return gobject.SignalConnect(x.GoPointer(), "activate-cursor-child", cbRefPtr)
}

// Emitted when a child has been activated by the user.
func (x *FlowBox) ConnectChildActivated(cb *func(FlowBox, uintptr)) uint32 {
	cbPtr := uintptr(unsafe.Pointer(cb))
	if cbRefPtr, ok := glib.GetCallback(cbPtr); ok {
		return gobject.SignalConnect(x.GoPointer(), "child-activated", cbRefPtr)
	}

	fcb := func(clsPtr uintptr, ChildVarp uintptr) {
		fa := FlowBox{}
		fa.Ptr = clsPtr
		cbFn := *cb

		cbFn(fa, ChildVarp)

	}
	cbRefPtr := purego.NewCallback(fcb)
	glib.SaveCallback(cbPtr, cbRefPtr)
	return gobject.SignalConnect(x.GoPointer(), "child-activated", cbRefPtr)
}

// Emitted when the user initiates a cursor movement.
//
// This is a [keybinding signal](class.SignalAction.html).
// Applications should not connect to it, but may emit it with
// g_signal_emit_by_name() if they need to control the cursor
// programmatically.
//
// The default bindings for this signal come in two variants,
// the variant with the Shift modifier extends the selection,
// the variant without the Shift modifier does not.
// There are too many key combinations to list them all here.
//
//   - &lt;kbd&gt;←&lt;/kbd&gt;, &lt;kbd&gt;→&lt;/kbd&gt;, &lt;kbd&gt;↑&lt;/kbd&gt;, &lt;kbd&gt;↓&lt;/kbd&gt;
//     move by individual children
//   - &lt;kbd&gt;Home&lt;/kbd&gt;, &lt;kbd&gt;End&lt;/kbd&gt; move to the ends of the box
//   - &lt;kbd&gt;PgUp&lt;/kbd&gt;, &lt;kbd&gt;PgDn&lt;/kbd&gt; move vertically by pages
func (x *FlowBox) ConnectMoveCursor(cb *func(FlowBox, MovementStep, int, bool, bool) bool) uint32 {
	cbPtr := uintptr(unsafe.Pointer(cb))
	if cbRefPtr, ok := glib.GetCallback(cbPtr); ok {
		return gobject.SignalConnect(x.GoPointer(), "move-cursor", cbRefPtr)
	}

	fcb := func(clsPtr uintptr, StepVarp MovementStep, CountVarp int, ExtendVarp bool, ModifyVarp bool) bool {
		fa := FlowBox{}
		fa.Ptr = clsPtr
		cbFn := *cb

		return cbFn(fa, StepVarp, CountVarp, ExtendVarp, ModifyVarp)

	}
	cbRefPtr := purego.NewCallback(fcb)
	glib.SaveCallback(cbPtr, cbRefPtr)
	return gobject.SignalConnect(x.GoPointer(), "move-cursor", cbRefPtr)
}

// Emitted to select all children of the box,
// if the selection mode permits it.
//
// This is a [keybinding signal](class.SignalAction.html).
//
// The default bindings for this signal is &lt;kbd&gt;Ctrl&lt;/kbd&gt;-&lt;kbd&gt;a&lt;/kbd&gt;.
func (x *FlowBox) ConnectSelectAll(cb *func(FlowBox)) uint32 {
	cbPtr := uintptr(unsafe.Pointer(cb))
	if cbRefPtr, ok := glib.GetCallback(cbPtr); ok {
		return gobject.SignalConnect(x.GoPointer(), "select-all", cbRefPtr)
	}

	fcb := func(clsPtr uintptr) {
		fa := FlowBox{}
		fa.Ptr = clsPtr
		cbFn := *cb

		cbFn(fa)

	}
	cbRefPtr := purego.NewCallback(fcb)
	glib.SaveCallback(cbPtr, cbRefPtr)
	return gobject.SignalConnect(x.GoPointer(), "select-all", cbRefPtr)
}

// Emitted when the set of selected children changes.
//
// Use [method@Gtk.FlowBox.selected_foreach] or
// [method@Gtk.FlowBox.get_selected_children] to obtain the
// selected children.
func (x *FlowBox) ConnectSelectedChildrenChanged(cb *func(FlowBox)) uint32 {
	cbPtr := uintptr(unsafe.Pointer(cb))
	if cbRefPtr, ok := glib.GetCallback(cbPtr); ok {
		return gobject.SignalConnect(x.GoPointer(), "selected-children-changed", cbRefPtr)
	}

	fcb := func(clsPtr uintptr) {
		fa := FlowBox{}
		fa.Ptr = clsPtr
		cbFn := *cb

		cbFn(fa)

	}
	cbRefPtr := purego.NewCallback(fcb)
	glib.SaveCallback(cbPtr, cbRefPtr)
	return gobject.SignalConnect(x.GoPointer(), "selected-children-changed", cbRefPtr)
}

// Emitted to toggle the selection of the child that has the focus.
//
// This is a [keybinding signal](class.SignalAction.html).
//
// The default binding for this signal is &lt;kbd&gt;Ctrl&lt;/kbd&gt;-&lt;kbd&gt;Space&lt;/kbd&gt;.
func (x *FlowBox) ConnectToggleCursorChild(cb *func(FlowBox)) uint32 {
	cbPtr := uintptr(unsafe.Pointer(cb))
	if cbRefPtr, ok := glib.GetCallback(cbPtr); ok {
		return gobject.SignalConnect(x.GoPointer(), "toggle-cursor-child", cbRefPtr)
	}

	fcb := func(clsPtr uintptr) {
		fa := FlowBox{}
		fa.Ptr = clsPtr
		cbFn := *cb

		cbFn(fa)

	}
	cbRefPtr := purego.NewCallback(fcb)
	glib.SaveCallback(cbPtr, cbRefPtr)
	return gobject.SignalConnect(x.GoPointer(), "toggle-cursor-child", cbRefPtr)
}

// Emitted to unselect all children of the box,
// if the selection mode permits it.
//
// This is a [keybinding signal](class.SignalAction.html).
//
// The default bindings for this signal is &lt;kbd&gt;Ctrl&lt;/kbd&gt;-&lt;kbd&gt;Shift&lt;/kbd&gt;-&lt;kbd&gt;a&lt;/kbd&gt;.
func (x *FlowBox) ConnectUnselectAll(cb *func(FlowBox)) uint32 {
	cbPtr := uintptr(unsafe.Pointer(cb))
	if cbRefPtr, ok := glib.GetCallback(cbPtr); ok {
		return gobject.SignalConnect(x.GoPointer(), "unselect-all", cbRefPtr)
	}

	fcb := func(clsPtr uintptr) {
		fa := FlowBox{}
		fa.Ptr = clsPtr
		cbFn := *cb

		cbFn(fa)

	}
	cbRefPtr := purego.NewCallback(fcb)
	glib.SaveCallback(cbPtr, cbRefPtr)
	return gobject.SignalConnect(x.GoPointer(), "unselect-all", cbRefPtr)
}

// Retrieves the `GtkAccessibleRole` for the given `GtkAccessible`.
func (x *FlowBox) GetAccessibleRole() AccessibleRole {

	cret := XGtkAccessibleGetAccessibleRole(x.GoPointer())
	return cret
}

// Resets the accessible @property to its default value.
func (x *FlowBox) ResetProperty(PropertyVar AccessibleProperty) {

	XGtkAccessibleResetProperty(x.GoPointer(), PropertyVar)

}

// Resets the accessible @relation to its default value.
func (x *FlowBox) ResetRelation(RelationVar AccessibleRelation) {

	XGtkAccessibleResetRelation(x.GoPointer(), RelationVar)

}

// Resets the accessible @state to its default value.
func (x *FlowBox) ResetState(StateVar AccessibleState) {

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
func (x *FlowBox) UpdateProperty(FirstPropertyVar AccessibleProperty, varArgs ...interface{}) {

	XGtkAccessibleUpdateProperty(x.GoPointer(), FirstPropertyVar, varArgs...)

}

// Updates an array of accessible properties.
//
// This function should be called by `GtkWidget` types whenever an accessible
// property change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *FlowBox) UpdatePropertyValue(NPropertiesVar int, PropertiesVar uintptr, ValuesVar uintptr) {

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
func (x *FlowBox) UpdateRelation(FirstRelationVar AccessibleRelation, varArgs ...interface{}) {

	XGtkAccessibleUpdateRelation(x.GoPointer(), FirstRelationVar, varArgs...)

}

// Updates an array of accessible relations.
//
// This function should be called by `GtkWidget` types whenever an accessible
// relation change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *FlowBox) UpdateRelationValue(NRelationsVar int, RelationsVar uintptr, ValuesVar uintptr) {

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
func (x *FlowBox) UpdateState(FirstStateVar AccessibleState, varArgs ...interface{}) {

	XGtkAccessibleUpdateState(x.GoPointer(), FirstStateVar, varArgs...)

}

// Updates an array of accessible states.
//
// This function should be called by `GtkWidget` types whenever an accessible
// state change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *FlowBox) UpdateStateValue(NStatesVar int, StatesVar uintptr, ValuesVar uintptr) {

	XGtkAccessibleUpdateStateValue(x.GoPointer(), NStatesVar, StatesVar, ValuesVar)

}

// Gets the ID of the @buildable object.
//
// `GtkBuilder` sets the name based on the ID attribute
// of the &lt;object&gt; tag used to construct the @buildable.
func (x *FlowBox) GetBuildableId() string {

	cret := XGtkBuildableGetBuildableId(x.GoPointer())
	return cret
}

// Retrieves the orientation of the @orientable.
func (x *FlowBox) GetOrientation() Orientation {

	cret := XGtkOrientableGetOrientation(x.GoPointer())
	return cret
}

// Sets the orientation of the @orientable.
func (x *FlowBox) SetOrientation(OrientationVar Orientation) {

	XGtkOrientableSetOrientation(x.GoPointer(), OrientationVar)

}

// `GtkFlowBoxChild` is the kind of widget that can be added to a `GtkFlowBox`.
type FlowBoxChild struct {
	Widget
}

func FlowBoxChildNewFromInternalPtr(ptr uintptr) *FlowBoxChild {
	cls := &FlowBoxChild{}
	cls.Ptr = ptr
	return cls
}

var xNewFlowBoxChild func() uintptr

// Creates a new `GtkFlowBoxChild`.
//
// This should only be used as a child of a `GtkFlowBox`.
func NewFlowBoxChild() *FlowBoxChild {
	var cls *FlowBoxChild

	cret := xNewFlowBoxChild()

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &FlowBoxChild{}
	cls.Ptr = cret
	return cls
}

var xFlowBoxChildChanged func(uintptr)

// Marks @child as changed, causing any state that depends on this
// to be updated.
//
// This affects sorting and filtering.
//
// Note that calls to this method must be in sync with the data
// used for the sorting and filtering functions. For instance, if
// the list is mirroring some external data set, and *two* children
// changed in the external data set when you call
// gtk_flow_box_child_changed() on the first child, the sort function
// must only read the new data for the first of the two changed
// children, otherwise the resorting of the children will be wrong.
//
// This generally means that if you don’t fully control the data
// model, you have to duplicate the data that affects the sorting
// and filtering functions into the widgets themselves.
//
// Another alternative is to call [method@Gtk.FlowBox.invalidate_sort]
// on any model change, but that is more expensive.
func (x *FlowBoxChild) Changed() {

	xFlowBoxChildChanged(x.GoPointer())

}

var xFlowBoxChildGetChild func(uintptr) uintptr

// Gets the child widget of @self.
func (x *FlowBoxChild) GetChild() *Widget {
	var cls *Widget

	cret := xFlowBoxChildGetChild(x.GoPointer())

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &Widget{}
	cls.Ptr = cret
	return cls
}

var xFlowBoxChildGetIndex func(uintptr) int

// Gets the current index of the @child in its `GtkFlowBox` container.
func (x *FlowBoxChild) GetIndex() int {

	cret := xFlowBoxChildGetIndex(x.GoPointer())
	return cret
}

var xFlowBoxChildIsSelected func(uintptr) bool

// Returns whether the @child is currently selected in its
// `GtkFlowBox` container.
func (x *FlowBoxChild) IsSelected() bool {

	cret := xFlowBoxChildIsSelected(x.GoPointer())
	return cret
}

var xFlowBoxChildSetChild func(uintptr, uintptr)

// Sets the child widget of @self.
func (x *FlowBoxChild) SetChild(ChildVar *Widget) {

	xFlowBoxChildSetChild(x.GoPointer(), ChildVar.GoPointer())

}

func (c *FlowBoxChild) GoPointer() uintptr {
	return c.Ptr
}

func (c *FlowBoxChild) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// Emitted when the user activates a child widget in a `GtkFlowBox`.
//
// This can be happen either by clicking or double-clicking,
// or via a keybinding.
//
// This is a [keybinding signal](class.SignalAction.html),
// but it can be used by applications for their own purposes.
//
// The default bindings are &lt;kbd&gt;Space&lt;/kbd&gt; and &lt;kbd&gt;Enter&lt;/kbd&gt;.
func (x *FlowBoxChild) ConnectActivate(cb *func(FlowBoxChild)) uint32 {
	cbPtr := uintptr(unsafe.Pointer(cb))
	if cbRefPtr, ok := glib.GetCallback(cbPtr); ok {
		return gobject.SignalConnect(x.GoPointer(), "activate", cbRefPtr)
	}

	fcb := func(clsPtr uintptr) {
		fa := FlowBoxChild{}
		fa.Ptr = clsPtr
		cbFn := *cb

		cbFn(fa)

	}
	cbRefPtr := purego.NewCallback(fcb)
	glib.SaveCallback(cbPtr, cbRefPtr)
	return gobject.SignalConnect(x.GoPointer(), "activate", cbRefPtr)
}

// Retrieves the `GtkAccessibleRole` for the given `GtkAccessible`.
func (x *FlowBoxChild) GetAccessibleRole() AccessibleRole {

	cret := XGtkAccessibleGetAccessibleRole(x.GoPointer())
	return cret
}

// Resets the accessible @property to its default value.
func (x *FlowBoxChild) ResetProperty(PropertyVar AccessibleProperty) {

	XGtkAccessibleResetProperty(x.GoPointer(), PropertyVar)

}

// Resets the accessible @relation to its default value.
func (x *FlowBoxChild) ResetRelation(RelationVar AccessibleRelation) {

	XGtkAccessibleResetRelation(x.GoPointer(), RelationVar)

}

// Resets the accessible @state to its default value.
func (x *FlowBoxChild) ResetState(StateVar AccessibleState) {

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
func (x *FlowBoxChild) UpdateProperty(FirstPropertyVar AccessibleProperty, varArgs ...interface{}) {

	XGtkAccessibleUpdateProperty(x.GoPointer(), FirstPropertyVar, varArgs...)

}

// Updates an array of accessible properties.
//
// This function should be called by `GtkWidget` types whenever an accessible
// property change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *FlowBoxChild) UpdatePropertyValue(NPropertiesVar int, PropertiesVar uintptr, ValuesVar uintptr) {

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
func (x *FlowBoxChild) UpdateRelation(FirstRelationVar AccessibleRelation, varArgs ...interface{}) {

	XGtkAccessibleUpdateRelation(x.GoPointer(), FirstRelationVar, varArgs...)

}

// Updates an array of accessible relations.
//
// This function should be called by `GtkWidget` types whenever an accessible
// relation change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *FlowBoxChild) UpdateRelationValue(NRelationsVar int, RelationsVar uintptr, ValuesVar uintptr) {

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
func (x *FlowBoxChild) UpdateState(FirstStateVar AccessibleState, varArgs ...interface{}) {

	XGtkAccessibleUpdateState(x.GoPointer(), FirstStateVar, varArgs...)

}

// Updates an array of accessible states.
//
// This function should be called by `GtkWidget` types whenever an accessible
// state change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *FlowBoxChild) UpdateStateValue(NStatesVar int, StatesVar uintptr, ValuesVar uintptr) {

	XGtkAccessibleUpdateStateValue(x.GoPointer(), NStatesVar, StatesVar, ValuesVar)

}

// Gets the ID of the @buildable object.
//
// `GtkBuilder` sets the name based on the ID attribute
// of the &lt;object&gt; tag used to construct the @buildable.
func (x *FlowBoxChild) GetBuildableId() string {

	cret := XGtkBuildableGetBuildableId(x.GoPointer())
	return cret
}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GTK"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xNewFlowBox, lib, "gtk_flow_box_new")

	core.PuregoSafeRegister(&xFlowBoxAppend, lib, "gtk_flow_box_append")
	core.PuregoSafeRegister(&xFlowBoxBindModel, lib, "gtk_flow_box_bind_model")
	core.PuregoSafeRegister(&xFlowBoxGetActivateOnSingleClick, lib, "gtk_flow_box_get_activate_on_single_click")
	core.PuregoSafeRegister(&xFlowBoxGetChildAtIndex, lib, "gtk_flow_box_get_child_at_index")
	core.PuregoSafeRegister(&xFlowBoxGetChildAtPos, lib, "gtk_flow_box_get_child_at_pos")
	core.PuregoSafeRegister(&xFlowBoxGetColumnSpacing, lib, "gtk_flow_box_get_column_spacing")
	core.PuregoSafeRegister(&xFlowBoxGetHomogeneous, lib, "gtk_flow_box_get_homogeneous")
	core.PuregoSafeRegister(&xFlowBoxGetMaxChildrenPerLine, lib, "gtk_flow_box_get_max_children_per_line")
	core.PuregoSafeRegister(&xFlowBoxGetMinChildrenPerLine, lib, "gtk_flow_box_get_min_children_per_line")
	core.PuregoSafeRegister(&xFlowBoxGetRowSpacing, lib, "gtk_flow_box_get_row_spacing")
	core.PuregoSafeRegister(&xFlowBoxGetSelectedChildren, lib, "gtk_flow_box_get_selected_children")
	core.PuregoSafeRegister(&xFlowBoxGetSelectionMode, lib, "gtk_flow_box_get_selection_mode")
	core.PuregoSafeRegister(&xFlowBoxInsert, lib, "gtk_flow_box_insert")
	core.PuregoSafeRegister(&xFlowBoxInvalidateFilter, lib, "gtk_flow_box_invalidate_filter")
	core.PuregoSafeRegister(&xFlowBoxInvalidateSort, lib, "gtk_flow_box_invalidate_sort")
	core.PuregoSafeRegister(&xFlowBoxPrepend, lib, "gtk_flow_box_prepend")
	core.PuregoSafeRegister(&xFlowBoxRemove, lib, "gtk_flow_box_remove")
	core.PuregoSafeRegister(&xFlowBoxSelectAll, lib, "gtk_flow_box_select_all")
	core.PuregoSafeRegister(&xFlowBoxSelectChild, lib, "gtk_flow_box_select_child")
	core.PuregoSafeRegister(&xFlowBoxSelectedForeach, lib, "gtk_flow_box_selected_foreach")
	core.PuregoSafeRegister(&xFlowBoxSetActivateOnSingleClick, lib, "gtk_flow_box_set_activate_on_single_click")
	core.PuregoSafeRegister(&xFlowBoxSetColumnSpacing, lib, "gtk_flow_box_set_column_spacing")
	core.PuregoSafeRegister(&xFlowBoxSetFilterFunc, lib, "gtk_flow_box_set_filter_func")
	core.PuregoSafeRegister(&xFlowBoxSetHadjustment, lib, "gtk_flow_box_set_hadjustment")
	core.PuregoSafeRegister(&xFlowBoxSetHomogeneous, lib, "gtk_flow_box_set_homogeneous")
	core.PuregoSafeRegister(&xFlowBoxSetMaxChildrenPerLine, lib, "gtk_flow_box_set_max_children_per_line")
	core.PuregoSafeRegister(&xFlowBoxSetMinChildrenPerLine, lib, "gtk_flow_box_set_min_children_per_line")
	core.PuregoSafeRegister(&xFlowBoxSetRowSpacing, lib, "gtk_flow_box_set_row_spacing")
	core.PuregoSafeRegister(&xFlowBoxSetSelectionMode, lib, "gtk_flow_box_set_selection_mode")
	core.PuregoSafeRegister(&xFlowBoxSetSortFunc, lib, "gtk_flow_box_set_sort_func")
	core.PuregoSafeRegister(&xFlowBoxSetVadjustment, lib, "gtk_flow_box_set_vadjustment")
	core.PuregoSafeRegister(&xFlowBoxUnselectAll, lib, "gtk_flow_box_unselect_all")
	core.PuregoSafeRegister(&xFlowBoxUnselectChild, lib, "gtk_flow_box_unselect_child")

	core.PuregoSafeRegister(&xNewFlowBoxChild, lib, "gtk_flow_box_child_new")

	core.PuregoSafeRegister(&xFlowBoxChildChanged, lib, "gtk_flow_box_child_changed")
	core.PuregoSafeRegister(&xFlowBoxChildGetChild, lib, "gtk_flow_box_child_get_child")
	core.PuregoSafeRegister(&xFlowBoxChildGetIndex, lib, "gtk_flow_box_child_get_index")
	core.PuregoSafeRegister(&xFlowBoxChildIsSelected, lib, "gtk_flow_box_child_is_selected")
	core.PuregoSafeRegister(&xFlowBoxChildSetChild, lib, "gtk_flow_box_child_set_child")

}
