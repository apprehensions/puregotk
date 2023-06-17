// Package gtk was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gtk

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/glib"
)

// A GtkTreeIterCompareFunc should return a negative integer, zero, or a positive
// integer if @a sorts before @b, @a sorts with @b, or @a sorts after @b
// respectively.
//
// If two iters compare as equal, their order in the sorted model
// is undefined. In order to ensure that the `GtkTreeSortable` behaves as
// expected, the GtkTreeIterCompareFunc must define a partial order on
// the model, i.e. it must be reflexive, antisymmetric and transitive.
//
// For example, if @model is a product catalogue, then a compare function
// for the “price” column could be one which returns
// `price_of(@a) - price_of(@b)`.
type TreeIterCompareFunc func(uintptr, *TreeIter, *TreeIter, uintptr) int32

type TreeSortableIface struct {
	GIface uintptr
}

// The interface for sortable models used by GtkTreeView
//
// `GtkTreeSortable` is an interface to be implemented by tree models which
// support sorting. The `GtkTreeView` uses the methods provided by this interface
// to sort the model.
type TreeSortable interface {
	GoPointer() uintptr
	SetGoPointer(uintptr)
	GetSortColumnId(SortColumnIdVar int32, OrderVar *SortType) bool
	HasDefaultSortFunc() bool
	SetDefaultSortFunc(SortFuncVar TreeIterCompareFunc, UserDataVar uintptr, DestroyVar glib.DestroyNotify)
	SetSortColumnId(SortColumnIdVar int32, OrderVar SortType)
	SetSortFunc(SortColumnIdVar int32, SortFuncVar TreeIterCompareFunc, UserDataVar uintptr, DestroyVar glib.DestroyNotify)
	SortColumnChanged()
}
type TreeSortableBase struct {
	Ptr uintptr
}

func (x *TreeSortableBase) GoPointer() uintptr {
	return x.Ptr
}

func (x *TreeSortableBase) SetGoPointer(ptr uintptr) {
	x.Ptr = ptr
}

// Fills in @sort_column_id and @order with the current sort column and the
// order. It returns %TRUE unless the @sort_column_id is
// %GTK_TREE_SORTABLE_DEFAULT_SORT_COLUMN_ID or
// %GTK_TREE_SORTABLE_UNSORTED_SORT_COLUMN_ID.
func (x *TreeSortableBase) GetSortColumnId(SortColumnIdVar int32, OrderVar *SortType) bool {

	return XGtkTreeSortableGetSortColumnId(x.GoPointer(), SortColumnIdVar, OrderVar)

}

// Returns %TRUE if the model has a default sort function. This is used
// primarily by GtkTreeViewColumns in order to determine if a model can
// go back to the default state, or not.
func (x *TreeSortableBase) HasDefaultSortFunc() bool {

	return XGtkTreeSortableHasDefaultSortFunc(x.GoPointer())

}

// Sets the default comparison function used when sorting to be @sort_func.
// If the current sort column id of @sortable is
// %GTK_TREE_SORTABLE_DEFAULT_SORT_COLUMN_ID, then the model will sort using
// this function.
//
// If @sort_func is %NULL, then there will be no default comparison function.
// This means that once the model  has been sorted, it can’t go back to the
// default state. In this case, when the current sort column id of @sortable
// is %GTK_TREE_SORTABLE_DEFAULT_SORT_COLUMN_ID, the model will be unsorted.
func (x *TreeSortableBase) SetDefaultSortFunc(SortFuncVar TreeIterCompareFunc, UserDataVar uintptr, DestroyVar glib.DestroyNotify) {

	XGtkTreeSortableSetDefaultSortFunc(x.GoPointer(), purego.NewCallback(SortFuncVar), UserDataVar, purego.NewCallback(DestroyVar))

}

// Sets the current sort column to be @sort_column_id. The @sortable will
// resort itself to reflect this change, after emitting a
// `GtkTreeSortable::sort-column-changed` signal. @sort_column_id may either be
// a regular column id, or one of the following special values:
//
//   - %GTK_TREE_SORTABLE_DEFAULT_SORT_COLUMN_ID: the default sort function
//     will be used, if it is set
//
// - %GTK_TREE_SORTABLE_UNSORTED_SORT_COLUMN_ID: no sorting will occur
func (x *TreeSortableBase) SetSortColumnId(SortColumnIdVar int32, OrderVar SortType) {

	XGtkTreeSortableSetSortColumnId(x.GoPointer(), SortColumnIdVar, OrderVar)

}

// Sets the comparison function used when sorting to be @sort_func. If the
// current sort column id of @sortable is the same as @sort_column_id, then
// the model will sort using this function.
func (x *TreeSortableBase) SetSortFunc(SortColumnIdVar int32, SortFuncVar TreeIterCompareFunc, UserDataVar uintptr, DestroyVar glib.DestroyNotify) {

	XGtkTreeSortableSetSortFunc(x.GoPointer(), SortColumnIdVar, purego.NewCallback(SortFuncVar), UserDataVar, purego.NewCallback(DestroyVar))

}

// Emits a `GtkTreeSortable::sort-column-changed` signal on @sortable.
func (x *TreeSortableBase) SortColumnChanged() {

	XGtkTreeSortableSortColumnChanged(x.GoPointer())

}

var XGtkTreeSortableGetSortColumnId func(uintptr, int32, *SortType) bool
var XGtkTreeSortableHasDefaultSortFunc func(uintptr) bool
var XGtkTreeSortableSetDefaultSortFunc func(uintptr, uintptr, uintptr, uintptr)
var XGtkTreeSortableSetSortColumnId func(uintptr, int32, SortType)
var XGtkTreeSortableSetSortFunc func(uintptr, int32, uintptr, uintptr, uintptr)
var XGtkTreeSortableSortColumnChanged func(uintptr)

func init() {
	lib, err := purego.Dlopen(core.GetPath("GTK"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&XGtkTreeSortableGetSortColumnId, lib, "gtk_tree_sortable_get_sort_column_id")
	core.PuregoSafeRegister(&XGtkTreeSortableHasDefaultSortFunc, lib, "gtk_tree_sortable_has_default_sort_func")
	core.PuregoSafeRegister(&XGtkTreeSortableSetDefaultSortFunc, lib, "gtk_tree_sortable_set_default_sort_func")
	core.PuregoSafeRegister(&XGtkTreeSortableSetSortColumnId, lib, "gtk_tree_sortable_set_sort_column_id")
	core.PuregoSafeRegister(&XGtkTreeSortableSetSortFunc, lib, "gtk_tree_sortable_set_sort_func")
	core.PuregoSafeRegister(&XGtkTreeSortableSortColumnChanged, lib, "gtk_tree_sortable_sort_column_changed")

}
