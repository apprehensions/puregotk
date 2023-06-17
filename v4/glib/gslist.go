// Package glib was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package glib

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
)

// The #GSList struct is used for each element in the singly-linked
// list.
type SList struct {
	Data uintptr

	Next *SList
}

var xClearSlist func(**SList, uintptr)

// Clears a pointer to a #GSList, freeing it and, optionally, freeing its elements using @destroy.
//
// @slist_ptr must be a valid pointer. If @slist_ptr points to a null #GSList, this does nothing.
func ClearSlist(SlistPtrVar **SList, DestroyVar DestroyNotify) {

	xClearSlist(SlistPtrVar, purego.NewCallback(DestroyVar))

}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GLIB"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	core.PuregoSafeRegister(&xClearSlist, lib, "g_clear_slist")

}
