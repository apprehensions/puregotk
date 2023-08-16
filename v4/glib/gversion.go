// Package glib was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package glib

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
)

var xCheckVersion func(uint, uint, uint) string

// Checks that the GLib library in use is compatible with the
// given version.
//
// Generally you would pass in the constants %GLIB_MAJOR_VERSION,
// %GLIB_MINOR_VERSION, %GLIB_MICRO_VERSION as the three arguments
// to this function; that produces a check that the library in use
// is compatible with the version of GLib the application or module
// was compiled against.
//
// Compatibility is defined by two things: first the version
// of the running library is newer than the version
// `@required_major.required_minor.@required_micro`. Second
// the running library must be binary compatible with the
// version `@required_major.@required_minor.@required_micro`
// (same major version.)
func CheckVersion(RequiredMajorVar uint, RequiredMinorVar uint, RequiredMicroVar uint) string {

	cret := xCheckVersion(RequiredMajorVar, RequiredMinorVar, RequiredMicroVar)
	return cret
}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GLIB"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	core.PuregoSafeRegister(&xCheckVersion, lib, "glib_check_version")

}
