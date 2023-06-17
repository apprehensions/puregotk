// Package gobject was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gobject

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
)

// This function is provided by the user and should produce a copy
// of the passed in boxed structure.
type BoxedCopyFunc func(uintptr) uintptr

// This function is provided by the user and should free the boxed
// structure passed.
type BoxedFreeFunc func(uintptr)

var xBoxedCopy func([]interface{}, uintptr) uintptr

// Provide a copy of a boxed structure @src_boxed which is of type @boxed_type.
func BoxedCopy(BoxedTypeVar []interface{}, SrcBoxedVar uintptr) uintptr {

	return xBoxedCopy(BoxedTypeVar, SrcBoxedVar)

}

var xBoxedFree func([]interface{}, uintptr)

// Free the boxed structure @boxed which is of type @boxed_type.
func BoxedFree(BoxedTypeVar []interface{}, BoxedVar uintptr) {

	xBoxedFree(BoxedTypeVar, BoxedVar)

}

var xBoxedTypeRegisterStatic func(string, uintptr, uintptr) []interface{}

// This function creates a new %G_TYPE_BOXED derived type id for a new
// boxed type with name @name.
//
// Boxed type handling functions have to be provided to copy and free
// opaque boxed structures of this type.
//
// For the general case, it is recommended to use G_DEFINE_BOXED_TYPE()
// instead of calling g_boxed_type_register_static() directly. The macro
// will create the appropriate `*_get_type()` function for the boxed type.
func BoxedTypeRegisterStatic(NameVar string, BoxedCopyVar BoxedCopyFunc, BoxedFreeVar BoxedFreeFunc) []interface{} {

	return xBoxedTypeRegisterStatic(NameVar, purego.NewCallback(BoxedCopyVar), purego.NewCallback(BoxedFreeVar))

}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GOBJECT"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	core.PuregoSafeRegister(&xBoxedCopy, lib, "g_boxed_copy")
	core.PuregoSafeRegister(&xBoxedFree, lib, "g_boxed_free")
	core.PuregoSafeRegister(&xBoxedTypeRegisterStatic, lib, "g_boxed_type_register_static")

}
