// Package pango was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package pango

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
)

var xLanguageFromString func(string) *Language

// Convert a language tag to a `PangoLanguage`.
//
// The language tag must be in a RFC-3066 format. `PangoLanguage` pointers
// can be efficiently copied (copy the pointer) and compared with other
// language tags (compare the pointer.)
//
// This function first canonicalizes the string by converting it to
// lowercase, mapping '_' to '-', and stripping all characters other
// than letters and '-'.
//
// Use [func@Pango.Language.get_default] if you want to get the
// `PangoLanguage` for the current locale of the process.
func LanguageFromString(LanguageVar string) *Language {

	return xLanguageFromString(LanguageVar)

}

var xLanguageGetDefault func() *Language

// Returns the `PangoLanguage` for the current locale of the process.
//
// On Unix systems, this is the return value is derived from
// `setlocale (LC_CTYPE, NULL)`, and the user can
// affect this through the environment variables LC_ALL, LC_CTYPE or
// LANG (checked in that order). The locale string typically is in
// the form lang_COUNTRY, where lang is an ISO-639 language code, and
// COUNTRY is an ISO-3166 country code. For instance, sv_FI for
// Swedish as written in Finland or pt_BR for Portuguese as written in
// Brazil.
//
// On Windows, the C library does not use any such environment
// variables, and setting them won't affect the behavior of functions
// like ctime(). The user sets the locale through the Regional Options
// in the Control Panel. The C library (in the setlocale() function)
// does not use country and language codes, but country and language
// names spelled out in English.
// However, this function does check the above environment
// variables, and does return a Unix-style locale string based on
// either said environment variables or the thread's current locale.
//
// Your application should call `setlocale(LC_ALL, "")` for the user
// settings to take effect. GTK does this in its initialization
// functions automatically (by calling gtk_set_locale()).
// See the setlocale() manpage for more details.
//
// Note that the default language can change over the life of an application.
//
// Also note that this function will not do the right thing if you
// use per-thread locales with uselocale(). In that case, you should
// just call pango_language_from_string() yourself.
func LanguageGetDefault() *Language {

	return xLanguageGetDefault()

}

var xLanguageGetPreferred func() **Language

// Returns the list of languages that the user prefers.
//
// The list is specified by the `PANGO_LANGUAGE` or `LANGUAGE`
// environment variables, in order of preference. Note that this
// list does not necessarily include the language returned by
// [func@Pango.Language.get_default].
//
// When choosing language-specific resources, such as the sample
// text returned by [method@Pango.Language.get_sample_string],
// you should first try the default language, followed by the
// languages returned by this function.
func LanguageGetPreferred() **Language {

	return xLanguageGetPreferred()

}

func init() {
	lib, err := purego.Dlopen(core.GetPath("PANGO"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	core.PuregoSafeRegister(&xLanguageFromString, lib, "pango_language_from_string")
	core.PuregoSafeRegister(&xLanguageGetDefault, lib, "pango_language_get_default")
	core.PuregoSafeRegister(&xLanguageGetPreferred, lib, "pango_language_get_preferred")

}
