// Package gtk was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gtk

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/glib"
	"github.com/jwijenbergh/puregotk/v4/gobject"
)

// A cell area that renders GtkCellRenderers into a row or a column
//
// The `GtkCellAreaBox` renders cell renderers into a row or a column
// depending on its `GtkOrientation`.
//
// GtkCellAreaBox uses a notion of packing. Packing
// refers to adding cell renderers with reference to a particular position
// in a `GtkCellAreaBox`. There are two reference positions: the
// start and the end of the box.
// When the `GtkCellAreaBox` is oriented in the %GTK_ORIENTATION_VERTICAL
// orientation, the start is defined as the top of the box and the end is
// defined as the bottom. In the %GTK_ORIENTATION_HORIZONTAL orientation
// start is defined as the left side and the end is defined as the right
// side.
//
// Alignments of `GtkCellRenderer`s rendered in adjacent rows can be
// configured by configuring the `GtkCellAreaBox` align child cell property
// with gtk_cell_area_cell_set_property() or by specifying the "align"
// argument to gtk_cell_area_box_pack_start() and gtk_cell_area_box_pack_end().
type CellAreaBox struct {
	CellArea
}

func CellAreaBoxNewFromInternalPtr(ptr uintptr) *CellAreaBox {
	cls := &CellAreaBox{}
	cls.Ptr = ptr
	return cls
}

var xNewCellAreaBox func() uintptr

// Creates a new `GtkCellAreaBox`.
func NewCellAreaBox() *CellArea {
	var cls *CellArea

	cret := xNewCellAreaBox()

	if cret == 0 {
		return cls
	}
	gobject.IncreaseRef(cret)
	cls = &CellArea{}
	cls.Ptr = cret
	return cls
}

var xCellAreaBoxGetSpacing func(uintptr) int

// Gets the spacing added between cell renderers.
func (x *CellAreaBox) GetSpacing() int {

	cret := xCellAreaBoxGetSpacing(x.GoPointer())
	return cret
}

var xCellAreaBoxPackEnd func(uintptr, uintptr, bool, bool, bool)

// Adds @renderer to @box, packed with reference to the end of @box.
//
// The @renderer is packed after (away from end of) any other
// `GtkCellRenderer` packed with reference to the end of @box.
func (x *CellAreaBox) PackEnd(RendererVar *CellRenderer, ExpandVar bool, AlignVar bool, FixedVar bool) {

	xCellAreaBoxPackEnd(x.GoPointer(), RendererVar.GoPointer(), ExpandVar, AlignVar, FixedVar)

}

var xCellAreaBoxPackStart func(uintptr, uintptr, bool, bool, bool)

// Adds @renderer to @box, packed with reference to the start of @box.
//
// The @renderer is packed after any other `GtkCellRenderer` packed
// with reference to the start of @box.
func (x *CellAreaBox) PackStart(RendererVar *CellRenderer, ExpandVar bool, AlignVar bool, FixedVar bool) {

	xCellAreaBoxPackStart(x.GoPointer(), RendererVar.GoPointer(), ExpandVar, AlignVar, FixedVar)

}

var xCellAreaBoxSetSpacing func(uintptr, int)

// Sets the spacing to add between cell renderers in @box.
func (x *CellAreaBox) SetSpacing(SpacingVar int) {

	xCellAreaBoxSetSpacing(x.GoPointer(), SpacingVar)

}

func (c *CellAreaBox) GoPointer() uintptr {
	return c.Ptr
}

func (c *CellAreaBox) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// Gets the ID of the @buildable object.
//
// `GtkBuilder` sets the name based on the ID attribute
// of the &lt;object&gt; tag used to construct the @buildable.
func (x *CellAreaBox) GetBuildableId() string {

	cret := XGtkBuildableGetBuildableId(x.GoPointer())
	return cret
}

// Adds an attribute mapping to the list in @cell_layout.
//
// The @column is the column of the model to get a value from, and the
// @attribute is the property on @cell to be set from that value. So for
// example if column 2 of the model contains strings, you could have the
// “text” attribute of a `GtkCellRendererText` get its values from column 2.
// In this context "attribute" and "property" are used interchangeably.
func (x *CellAreaBox) AddAttribute(CellVar *CellRenderer, AttributeVar string, ColumnVar int) {

	XGtkCellLayoutAddAttribute(x.GoPointer(), CellVar.GoPointer(), AttributeVar, ColumnVar)

}

// Unsets all the mappings on all renderers on @cell_layout and
// removes all renderers from @cell_layout.
func (x *CellAreaBox) Clear() {

	XGtkCellLayoutClear(x.GoPointer())

}

// Clears all existing attributes previously set with
// gtk_cell_layout_set_attributes().
func (x *CellAreaBox) ClearAttributes(CellVar *CellRenderer) {

	XGtkCellLayoutClearAttributes(x.GoPointer(), CellVar.GoPointer())

}

// Returns the underlying `GtkCellArea` which might be @cell_layout
// if called on a `GtkCellArea` or might be %NULL if no `GtkCellArea`
// is used by @cell_layout.
func (x *CellAreaBox) GetArea() *CellArea {
	var cls *CellArea

	cret := XGtkCellLayoutGetArea(x.GoPointer())

	if cret == 0 {
		return cls
	}
	gobject.IncreaseRef(cret)
	cls = &CellArea{}
	cls.Ptr = cret
	return cls
}

// Returns the cell renderers which have been added to @cell_layout.
func (x *CellAreaBox) GetCells() *glib.List {

	cret := XGtkCellLayoutGetCells(x.GoPointer())
	return cret
}

// Re-inserts @cell at @position.
//
// Note that @cell has already to be packed into @cell_layout
// for this to function properly.
func (x *CellAreaBox) Reorder(CellVar *CellRenderer, PositionVar int) {

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
func (x *CellAreaBox) SetAttributes(CellVar *CellRenderer, varArgs ...interface{}) {

	XGtkCellLayoutSetAttributes(x.GoPointer(), CellVar.GoPointer(), varArgs...)

}

// Sets the `GtkCellLayout`DataFunc to use for @cell_layout.
//
// This function is used instead of the standard attributes mapping
// for setting the column value, and should set the value of @cell_layout’s
// cell renderer(s) as appropriate.
//
// @func may be %NULL to remove a previously set function.
func (x *CellAreaBox) SetCellDataFunc(CellVar *CellRenderer, FuncVar CellLayoutDataFunc, FuncDataVar uintptr, DestroyVar glib.DestroyNotify) {

	XGtkCellLayoutSetCellDataFunc(x.GoPointer(), CellVar.GoPointer(), purego.NewCallback(FuncVar), FuncDataVar, purego.NewCallback(DestroyVar))

}

// Retrieves the orientation of the @orientable.
func (x *CellAreaBox) GetOrientation() Orientation {

	cret := XGtkOrientableGetOrientation(x.GoPointer())
	return cret
}

// Sets the orientation of the @orientable.
func (x *CellAreaBox) SetOrientation(OrientationVar Orientation) {

	XGtkOrientableSetOrientation(x.GoPointer(), OrientationVar)

}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GTK"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xNewCellAreaBox, lib, "gtk_cell_area_box_new")

	core.PuregoSafeRegister(&xCellAreaBoxGetSpacing, lib, "gtk_cell_area_box_get_spacing")
	core.PuregoSafeRegister(&xCellAreaBoxPackEnd, lib, "gtk_cell_area_box_pack_end")
	core.PuregoSafeRegister(&xCellAreaBoxPackStart, lib, "gtk_cell_area_box_pack_start")
	core.PuregoSafeRegister(&xCellAreaBoxSetSpacing, lib, "gtk_cell_area_box_set_spacing")

}
