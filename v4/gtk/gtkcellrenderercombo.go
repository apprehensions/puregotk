// Package gtk was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gtk

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/gobject"
)

// Renders a combobox in a cell
//
// `GtkCellRendererCombo` renders text in a cell like `GtkCellRendererText` from
// which it is derived. But while `GtkCellRendererText` offers a simple entry to
// edit the text, `GtkCellRendererCombo` offers a `GtkComboBox`
// widget to edit the text. The values to display in the combo box are taken from
// the tree model specified in the `GtkCellRendererCombo`:model property.
//
// The combo cell renderer takes care of adding a text cell renderer to the combo
// box and sets it to display the column specified by its
// `GtkCellRendererCombo`:text-column property. Further properties of the combo box
// can be set in a handler for the `GtkCellRenderer::editing-started` signal.
type CellRendererCombo struct {
	CellRendererText
}

func CellRendererComboNewFromInternalPtr(ptr uintptr) *CellRendererCombo {
	cls := &CellRendererCombo{}
	cls.Ptr = ptr
	return cls
}

var xNewCellRendererCombo func() uintptr

// Creates a new `GtkCellRendererCombo`.
// Adjust how text is drawn using object properties.
// Object properties can be set globally (with g_object_set()).
// Also, with `GtkTreeViewColumn`, you can bind a property to a value
// in a `GtkTreeModel`. For example, you can bind the “text” property
// on the cell renderer to a string value in the model, thus rendering
// a different string in each row of the `GtkTreeView`.
func NewCellRendererCombo() *CellRendererCombo {
	var cls *CellRendererCombo

	cret := xNewCellRendererCombo()

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &CellRendererCombo{}
	cls.Ptr = cret
	return cls
}

func (c *CellRendererCombo) GoPointer() uintptr {
	return c.Ptr
}

func (c *CellRendererCombo) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// This signal is emitted each time after the user selected an item in
// the combo box, either by using the mouse or the arrow keys.  Contrary
// to GtkComboBox, GtkCellRendererCombo::changed is not emitted for
// changes made to a selected item in the entry.  The argument @new_iter
// corresponds to the newly selected item in the combo box and it is relative
// to the GtkTreeModel set via the model property on GtkCellRendererCombo.
//
// Note that as soon as you change the model displayed in the tree view,
// the tree view will immediately cease the editing operating.  This
// means that you most probably want to refrain from changing the model
// until the combo cell renderer emits the edited or editing_canceled signal.
func (x *CellRendererCombo) ConnectChanged(cb func(CellRendererCombo, string, uintptr)) uint32 {
	fcb := func(clsPtr uintptr, PathStringVarp string, NewIterVarp uintptr) {
		fa := CellRendererCombo{}
		fa.Ptr = clsPtr

		cb(fa, PathStringVarp, NewIterVarp)

	}
	return gobject.SignalConnect(x.GoPointer(), "changed", purego.NewCallback(fcb))
}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GTK"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xNewCellRendererCombo, lib, "gtk_cell_renderer_combo_new")

}
