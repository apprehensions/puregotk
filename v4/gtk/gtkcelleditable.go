// Package gtk was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gtk

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/gdk"
)

type CellEditableIface struct {
	GIface uintptr
}

// Interface for widgets that can be used for editing cells
//
// The `GtkCellEditable` interface must be implemented for widgets to be usable
// to edit the contents of a `GtkTreeView` cell. It provides a way to specify how
// temporary widgets should be configured for editing, get the new value, etc.
type CellEditable interface {
	GoPointer() uintptr
	SetGoPointer(uintptr)
	EditingDone()
	RemoveWidget()
	StartEditing(EventVar *gdk.Event)
}
type CellEditableBase struct {
	Ptr uintptr
}

func (x *CellEditableBase) GoPointer() uintptr {
	return x.Ptr
}

func (x *CellEditableBase) SetGoPointer(ptr uintptr) {
	x.Ptr = ptr
}

// Emits the `GtkCellEditable::editing-done` signal.
func (x *CellEditableBase) EditingDone() {

	XGtkCellEditableEditingDone(x.GoPointer())

}

// Emits the `GtkCellEditable::remove-widget` signal.
func (x *CellEditableBase) RemoveWidget() {

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
func (x *CellEditableBase) StartEditing(EventVar *gdk.Event) {

	XGtkCellEditableStartEditing(x.GoPointer(), EventVar.GoPointer())

}

var XGtkCellEditableEditingDone func(uintptr)
var XGtkCellEditableRemoveWidget func(uintptr)
var XGtkCellEditableStartEditing func(uintptr, uintptr)

func init() {
	lib, err := purego.Dlopen(core.GetPath("GTK"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&XGtkCellEditableEditingDone, lib, "gtk_cell_editable_editing_done")
	core.PuregoSafeRegister(&XGtkCellEditableRemoveWidget, lib, "gtk_cell_editable_remove_widget")
	core.PuregoSafeRegister(&XGtkCellEditableStartEditing, lib, "gtk_cell_editable_start_editing")

}
