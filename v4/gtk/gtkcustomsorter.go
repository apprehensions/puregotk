// Package gtk was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gtk

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/glib"
)

type CustomSorterClass struct {
	ParentClass uintptr
}

// `GtkCustomSorter` is a `GtkSorter` implementation that sorts via a callback
// function.
type CustomSorter struct {
	Sorter
}

func CustomSorterNewFromInternalPtr(ptr uintptr) *CustomSorter {
	cls := &CustomSorter{}
	cls.Ptr = ptr
	return cls
}

var xNewCustomSorter func(uintptr, uintptr, uintptr) uintptr

// Creates a new `GtkSorter` that works by calling
// @sort_func to compare items.
//
// If @sort_func is %NULL, all items are considered equal.
func NewCustomSorter(SortFuncVar glib.CompareDataFunc, UserDataVar uintptr, UserDestroyVar glib.DestroyNotify) *CustomSorter {
	var cls *CustomSorter

	cret := xNewCustomSorter(purego.NewCallback(SortFuncVar), UserDataVar, purego.NewCallback(UserDestroyVar))

	if cret == 0 {
		return cls
	}
	cls = &CustomSorter{}
	cls.Ptr = cret
	return cls
}

var xCustomSorterSetSortFunc func(uintptr, uintptr, uintptr, uintptr)

// Sets (or unsets) the function used for sorting items.
//
// If @sort_func is %NULL, all items are considered equal.
//
// If the sort func changes its sorting behavior,
// gtk_sorter_changed() needs to be called.
//
// If a previous function was set, its @user_destroy will be
// called now.
func (x *CustomSorter) SetSortFunc(SortFuncVar glib.CompareDataFunc, UserDataVar uintptr, UserDestroyVar glib.DestroyNotify) {

	xCustomSorterSetSortFunc(x.GoPointer(), purego.NewCallback(SortFuncVar), UserDataVar, purego.NewCallback(UserDestroyVar))

}

func (c *CustomSorter) GoPointer() uintptr {
	return c.Ptr
}

func (c *CustomSorter) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GTK"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xNewCustomSorter, lib, "gtk_custom_sorter_new")

	core.PuregoSafeRegister(&xCustomSorterSetSortFunc, lib, "gtk_custom_sorter_set_sort_func")

}
