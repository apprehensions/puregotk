// Package gdk was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gdk

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
)

var xKeyvalConvertCase func(uint, uint, uint)

// Obtains the upper- and lower-case versions of the keyval @symbol.
//
// Examples of keyvals are `GDK_KEY_a`, `GDK_KEY_Enter`, `GDK_KEY_F1`, etc.
func KeyvalConvertCase(SymbolVar uint, LowerVar uint, UpperVar uint) {

	xKeyvalConvertCase(SymbolVar, LowerVar, UpperVar)

}

var xKeyvalFromName func(string) uint

// Converts a key name to a key value.
//
// The names are the same as those in the
// `gdk/gdkkeysyms.h` header file
// but without the leading “GDK_KEY_”.
func KeyvalFromName(KeyvalNameVar string) uint {

	return xKeyvalFromName(KeyvalNameVar)

}

var xKeyvalIsLower func(uint) bool

// Returns %TRUE if the given key value is in lower case.
func KeyvalIsLower(KeyvalVar uint) bool {

	return xKeyvalIsLower(KeyvalVar)

}

var xKeyvalIsUpper func(uint) bool

// Returns %TRUE if the given key value is in upper case.
func KeyvalIsUpper(KeyvalVar uint) bool {

	return xKeyvalIsUpper(KeyvalVar)

}

var xKeyvalName func(uint) string

// Converts a key value into a symbolic name.
//
// The names are the same as those in the
// `gdk/gdkkeysyms.h` header file
// but without the leading “GDK_KEY_”.
func KeyvalName(KeyvalVar uint) string {

	return xKeyvalName(KeyvalVar)

}

var xKeyvalToLower func(uint) uint

// Converts a key value to lower case, if applicable.
func KeyvalToLower(KeyvalVar uint) uint {

	return xKeyvalToLower(KeyvalVar)

}

var xKeyvalToUnicode func(uint) uint32

// Convert from a GDK key symbol to the corresponding Unicode
// character.
//
// Note that the conversion does not take the current locale
// into consideration, which might be expected for particular
// keyvals, such as %GDK_KEY_KP_Decimal.
func KeyvalToUnicode(KeyvalVar uint) uint32 {

	return xKeyvalToUnicode(KeyvalVar)

}

var xKeyvalToUpper func(uint) uint

// Converts a key value to upper case, if applicable.
func KeyvalToUpper(KeyvalVar uint) uint {

	return xKeyvalToUpper(KeyvalVar)

}

var xUnicodeToKeyval func(uint32) uint

// Convert from a Unicode character to a key symbol.
func UnicodeToKeyval(WcVar uint32) uint {

	return xUnicodeToKeyval(WcVar)

}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GDK"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	core.PuregoSafeRegister(&xKeyvalConvertCase, lib, "gdk_keyval_convert_case")
	core.PuregoSafeRegister(&xKeyvalFromName, lib, "gdk_keyval_from_name")
	core.PuregoSafeRegister(&xKeyvalIsLower, lib, "gdk_keyval_is_lower")
	core.PuregoSafeRegister(&xKeyvalIsUpper, lib, "gdk_keyval_is_upper")
	core.PuregoSafeRegister(&xKeyvalName, lib, "gdk_keyval_name")
	core.PuregoSafeRegister(&xKeyvalToLower, lib, "gdk_keyval_to_lower")
	core.PuregoSafeRegister(&xKeyvalToUnicode, lib, "gdk_keyval_to_unicode")
	core.PuregoSafeRegister(&xKeyvalToUpper, lib, "gdk_keyval_to_upper")
	core.PuregoSafeRegister(&xUnicodeToKeyval, lib, "gdk_unicode_to_keyval")

}
