// Package glib was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package glib

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
)

// A GQuark is a non-zero integer which uniquely identifies a
// particular string. A GQuark value of zero is associated to %NULL.
type Quark = uint32

var xInternStaticString func(string) string

// Returns a canonical representation for @string. Interned strings
// can be compared for equality by comparing the pointers, instead of
// using strcmp(). g_intern_static_string() does not copy the string,
// therefore @string must not be freed or modified.
//
// This function must not be used before library constructors have finished
// running. In particular, this means it cannot be used to initialize global
// variables in C++.
func InternStaticString(StringVar string) string {

	cret := xInternStaticString(StringVar)
	return cret
}

var xInternString func(string) string

// Returns a canonical representation for @string. Interned strings
// can be compared for equality by comparing the pointers, instead of
// using strcmp().
//
// This function must not be used before library constructors have finished
// running. In particular, this means it cannot be used to initialize global
// variables in C++.
func InternString(StringVar string) string {

	cret := xInternString(StringVar)
	return cret
}

var xQuarkFromStaticString func(string) Quark

// Gets the #GQuark identifying the given (static) string. If the
// string does not currently have an associated #GQuark, a new #GQuark
// is created, linked to the given string.
//
// Note that this function is identical to g_quark_from_string() except
// that if a new #GQuark is created the string itself is used rather
// than a copy. This saves memory, but can only be used if the string
// will continue to exist until the program terminates. It can be used
// with statically allocated strings in the main program, but not with
// statically allocated memory in dynamically loaded modules, if you
// expect to ever unload the module again (e.g. do not use this
// function in GTK+ theme engines).
//
// This function must not be used before library constructors have finished
// running. In particular, this means it cannot be used to initialize global
// variables in C++.
func QuarkFromStaticString(StringVar string) Quark {

	cret := xQuarkFromStaticString(StringVar)
	return cret
}

var xQuarkFromString func(string) Quark

// Gets the #GQuark identifying the given string. If the string does
// not currently have an associated #GQuark, a new #GQuark is created,
// using a copy of the string.
//
// This function must not be used before library constructors have finished
// running. In particular, this means it cannot be used to initialize global
// variables in C++.
func QuarkFromString(StringVar string) Quark {

	cret := xQuarkFromString(StringVar)
	return cret
}

var xQuarkToString func(Quark) string

// Gets the string associated with the given #GQuark.
func QuarkToString(QuarkVar Quark) string {

	cret := xQuarkToString(QuarkVar)
	return cret
}

var xQuarkTryString func(string) Quark

// Gets the #GQuark associated with the given string, or 0 if string is
// %NULL or it has no associated #GQuark.
//
// If you want the GQuark to be created if it doesn't already exist,
// use g_quark_from_string() or g_quark_from_static_string().
//
// This function must not be used before library constructors have finished
// running.
func QuarkTryString(StringVar string) Quark {

	cret := xQuarkTryString(StringVar)
	return cret
}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GLIB"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	core.PuregoSafeRegister(&xInternStaticString, lib, "g_intern_static_string")
	core.PuregoSafeRegister(&xInternString, lib, "g_intern_string")
	core.PuregoSafeRegister(&xQuarkFromStaticString, lib, "g_quark_from_static_string")
	core.PuregoSafeRegister(&xQuarkFromString, lib, "g_quark_from_string")
	core.PuregoSafeRegister(&xQuarkToString, lib, "g_quark_to_string")
	core.PuregoSafeRegister(&xQuarkTryString, lib, "g_quark_try_string")

}
