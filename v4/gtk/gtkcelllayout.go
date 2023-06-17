// Package gtk was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gtk

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/glib"
	"github.com/jwijenbergh/puregotk/v4/gobject"
)

// A function which should set the value of @cell_layout’s cell renderer(s)
// as appropriate.
type CellLayoutDataFunc func(uintptr, uintptr, uintptr, *TreeIter, uintptr)

type CellLayoutIface struct {
	GIface uintptr
}

// An interface for packing cells
//
// `GtkCellLayout` is an interface to be implemented by all objects which
// want to provide a `GtkTreeViewColumn` like API for packing cells,
// setting attributes and data funcs.
//
// One of the notable features provided by implementations of
// `GtkCellLayout` are attributes. Attributes let you set the properties
// in flexible ways. They can just be set to constant values like regular
// properties. But they can also be mapped to a column of the underlying
// tree model with gtk_cell_layout_set_attributes(), which means that the value
// of the attribute can change from cell to cell as they are rendered by
// the cell renderer. Finally, it is possible to specify a function with
// gtk_cell_layout_set_cell_data_func() that is called to determine the
// value of the attribute for each cell that is rendered.
//
// # GtkCellLayouts as GtkBuildable
//
// Implementations of GtkCellLayout which also implement the GtkBuildable
// interface (`GtkCellView`, `GtkIconView`, `GtkComboBox`,
// `GtkEntryCompletion`, `GtkTreeViewColumn`) accept `GtkCellRenderer` objects
// as `&lt;child&gt;` elements in UI definitions. They support a custom `&lt;attributes&gt;`
// element for their children, which can contain multiple `&lt;attribute&gt;`
// elements. Each `&lt;attribute&gt;` element has a name attribute which specifies
// a property of the cell renderer; the content of the element is the
// attribute value.
//
// This is an example of a UI definition fragment specifying attributes:
//
// ```xml
// &lt;object class="GtkCellView"&gt;
//
//	&lt;child&gt;
//	  &lt;object class="GtkCellRendererText"/&gt;
//	  &lt;attributes&gt;
//	    &lt;attribute name="text"&gt;0&lt;/attribute&gt;
//	  &lt;/attributes&gt;
//	&lt;/child&gt;
//
// &lt;/object&gt;
// ```
//
// Furthermore for implementations of `GtkCellLayout` that use a `GtkCellArea`
// to lay out cells (all `GtkCellLayout`s in GTK use a `GtkCellArea`)
// [cell properties](class.CellArea.html#cell-properties) can also be defined
// in the format by specifying the custom `&lt;cell-packing&gt;` attribute which can
// contain multiple `&lt;property&gt;` elements.
//
// Here is a UI definition fragment specifying cell properties:
//
// ```xml
// &lt;object class="GtkTreeViewColumn"&gt;
//
//	&lt;child&gt;
//	  &lt;object class="GtkCellRendererText"/&gt;
//	  &lt;cell-packing&gt;
//	    &lt;property name="align"&gt;True&lt;/property&gt;
//	    &lt;property name="expand"&gt;False&lt;/property&gt;
//	  &lt;/cell-packing&gt;
//	&lt;/child&gt;
//
// &lt;/object&gt;
// ```
//
// # Subclassing GtkCellLayout implementations
//
// When subclassing a widget that implements `GtkCellLayout` like
// `GtkIconView` or `GtkComboBox`, there are some considerations related
// to the fact that these widgets internally use a `GtkCellArea`.
// The cell area is exposed as a construct-only property by these
// widgets. This means that it is possible to e.g. do
//
// ```c
// GtkWIdget *combo =
//
//	g_object_new (GTK_TYPE_COMBO_BOX, "cell-area", my_cell_area, NULL);
//
// ```
//
// to use a custom cell area with a combo box. But construct properties
// are only initialized after instance `init()`
// functions have run, which means that using functions which rely on
// the existence of the cell area in your subclass `init()` function will
// cause the default cell area to be instantiated. In this case, a provided
// construct property value will be ignored (with a warning, to alert
// you to the problem).
//
// ```c
// static void
// my_combo_box_init (MyComboBox *b)
//
//	{
//	  GtkCellRenderer *cell;
//
//	  cell = gtk_cell_renderer_pixbuf_new ();
//
//	  // The following call causes the default cell area for combo boxes,
//	  // a GtkCellAreaBox, to be instantiated
//	  gtk_cell_layout_pack_start (GTK_CELL_LAYOUT (b), cell, FALSE);
//	  ...
//	}
//
// GtkWidget *
// my_combo_box_new (GtkCellArea *area)
//
//	{
//	  // This call is going to cause a warning about area being ignored
//	  return g_object_new (MY_TYPE_COMBO_BOX, "cell-area", area, NULL);
//	}
//
// ```
//
// If supporting alternative cell areas with your derived widget is
// not important, then this does not have to concern you. If you want
// to support alternative cell areas, you can do so by moving the
// problematic calls out of `init()` and into a `constructor()`
// for your class.
type CellLayout interface {
	GoPointer() uintptr
	SetGoPointer(uintptr)
	AddAttribute(CellVar *CellRenderer, AttributeVar string, ColumnVar int32)
	Clear()
	ClearAttributes(CellVar *CellRenderer)
	GetArea() *CellArea
	GetCells() *glib.List
	PackEnd(CellVar *CellRenderer, ExpandVar bool)
	PackStart(CellVar *CellRenderer, ExpandVar bool)
	Reorder(CellVar *CellRenderer, PositionVar int32)
	SetAttributes(CellVar *CellRenderer, varArgs ...interface{})
	SetCellDataFunc(CellVar *CellRenderer, FuncVar CellLayoutDataFunc, FuncDataVar uintptr, DestroyVar glib.DestroyNotify)
}
type CellLayoutBase struct {
	Ptr uintptr
}

func (x *CellLayoutBase) GoPointer() uintptr {
	return x.Ptr
}

func (x *CellLayoutBase) SetGoPointer(ptr uintptr) {
	x.Ptr = ptr
}

// Adds an attribute mapping to the list in @cell_layout.
//
// The @column is the column of the model to get a value from, and the
// @attribute is the property on @cell to be set from that value. So for
// example if column 2 of the model contains strings, you could have the
// “text” attribute of a `GtkCellRendererText` get its values from column 2.
// In this context "attribute" and "property" are used interchangeably.
func (x *CellLayoutBase) AddAttribute(CellVar *CellRenderer, AttributeVar string, ColumnVar int32) {

	XGtkCellLayoutAddAttribute(x.GoPointer(), CellVar.GoPointer(), AttributeVar, ColumnVar)

}

// Unsets all the mappings on all renderers on @cell_layout and
// removes all renderers from @cell_layout.
func (x *CellLayoutBase) Clear() {

	XGtkCellLayoutClear(x.GoPointer())

}

// Clears all existing attributes previously set with
// gtk_cell_layout_set_attributes().
func (x *CellLayoutBase) ClearAttributes(CellVar *CellRenderer) {

	XGtkCellLayoutClearAttributes(x.GoPointer(), CellVar.GoPointer())

}

// Returns the underlying `GtkCellArea` which might be @cell_layout
// if called on a `GtkCellArea` or might be %NULL if no `GtkCellArea`
// is used by @cell_layout.
func (x *CellLayoutBase) GetArea() *CellArea {

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
func (x *CellLayoutBase) GetCells() *glib.List {

	return XGtkCellLayoutGetCells(x.GoPointer())

}

// Adds the @cell to the end of @cell_layout. If @expand is %FALSE, then the
// @cell is allocated no more space than it needs. Any unused space is
// divided evenly between cells for which @expand is %TRUE.
//
// Note that reusing the same cell renderer is not supported.
func (x *CellLayoutBase) PackEnd(CellVar *CellRenderer, ExpandVar bool) {

	XGtkCellLayoutPackEnd(x.GoPointer(), CellVar.GoPointer(), ExpandVar)

}

// Packs the @cell into the beginning of @cell_layout. If @expand is %FALSE,
// then the @cell is allocated no more space than it needs. Any unused space
// is divided evenly between cells for which @expand is %TRUE.
//
// Note that reusing the same cell renderer is not supported.
func (x *CellLayoutBase) PackStart(CellVar *CellRenderer, ExpandVar bool) {

	XGtkCellLayoutPackStart(x.GoPointer(), CellVar.GoPointer(), ExpandVar)

}

// Re-inserts @cell at @position.
//
// Note that @cell has already to be packed into @cell_layout
// for this to function properly.
func (x *CellLayoutBase) Reorder(CellVar *CellRenderer, PositionVar int32) {

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
func (x *CellLayoutBase) SetAttributes(CellVar *CellRenderer, varArgs ...interface{}) {

	XGtkCellLayoutSetAttributes(x.GoPointer(), CellVar.GoPointer(), varArgs...)

}

// Sets the `GtkCellLayout`DataFunc to use for @cell_layout.
//
// This function is used instead of the standard attributes mapping
// for setting the column value, and should set the value of @cell_layout’s
// cell renderer(s) as appropriate.
//
// @func may be %NULL to remove a previously set function.
func (x *CellLayoutBase) SetCellDataFunc(CellVar *CellRenderer, FuncVar CellLayoutDataFunc, FuncDataVar uintptr, DestroyVar glib.DestroyNotify) {

	XGtkCellLayoutSetCellDataFunc(x.GoPointer(), CellVar.GoPointer(), purego.NewCallback(FuncVar), FuncDataVar, purego.NewCallback(DestroyVar))

}

var XGtkCellLayoutAddAttribute func(uintptr, uintptr, string, int32)
var XGtkCellLayoutClear func(uintptr)
var XGtkCellLayoutClearAttributes func(uintptr, uintptr)
var XGtkCellLayoutGetArea func(uintptr) uintptr
var XGtkCellLayoutGetCells func(uintptr) *glib.List
var XGtkCellLayoutPackEnd func(uintptr, uintptr, bool)
var XGtkCellLayoutPackStart func(uintptr, uintptr, bool)
var XGtkCellLayoutReorder func(uintptr, uintptr, int32)
var XGtkCellLayoutSetAttributes func(uintptr, uintptr, ...interface{})
var XGtkCellLayoutSetCellDataFunc func(uintptr, uintptr, uintptr, uintptr, uintptr)

func init() {
	lib, err := purego.Dlopen(core.GetPath("GTK"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&XGtkCellLayoutAddAttribute, lib, "gtk_cell_layout_add_attribute")
	core.PuregoSafeRegister(&XGtkCellLayoutClear, lib, "gtk_cell_layout_clear")
	core.PuregoSafeRegister(&XGtkCellLayoutClearAttributes, lib, "gtk_cell_layout_clear_attributes")
	core.PuregoSafeRegister(&XGtkCellLayoutGetArea, lib, "gtk_cell_layout_get_area")
	core.PuregoSafeRegister(&XGtkCellLayoutGetCells, lib, "gtk_cell_layout_get_cells")
	core.PuregoSafeRegister(&XGtkCellLayoutPackEnd, lib, "gtk_cell_layout_pack_end")
	core.PuregoSafeRegister(&XGtkCellLayoutPackStart, lib, "gtk_cell_layout_pack_start")
	core.PuregoSafeRegister(&XGtkCellLayoutReorder, lib, "gtk_cell_layout_reorder")
	core.PuregoSafeRegister(&XGtkCellLayoutSetAttributes, lib, "gtk_cell_layout_set_attributes")
	core.PuregoSafeRegister(&XGtkCellLayoutSetCellDataFunc, lib, "gtk_cell_layout_set_cell_data_func")

}
