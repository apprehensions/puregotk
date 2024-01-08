// Package gtk was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gtk

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
)

const (
	// Like [func@get_binary_age], but from the headers used at
	// application compile time, rather than from the library linked
	// against at application run time.
	BINARY_AGE int = 803
	// Like [func@get_interface_age], but from the headers used at
	// application compile time, rather than from the library linked
	// against at application run time.
	INTERFACE_AGE int = 3
	// Like [func@get_major_version], but from the headers used at
	// application compile time, rather than from the library linked
	// against at application run time.
	MAJOR_VERSION int = 4
	// Like [func@get_micro_version], but from the headers used at
	// application compile time, rather than from the library linked
	// against at application run time.
	MICRO_VERSION int = 3
	// Like [func@get_minor_version], but from the headers used at
	// application compile time, rather than from the library linked
	// against at application run time.
	MINOR_VERSION int = 8
)

var xCheckVersion func(uint, uint, uint) string

// Checks that the GTK library in use is compatible with the
// given version.
//
// Generally you would pass in the constants %GTK_MAJOR_VERSION,
// %GTK_MINOR_VERSION, %GTK_MICRO_VERSION as the three arguments
// to this function; that produces a check that the library in
// use is compatible with the version of GTK the application or
// module was compiled against.
//
// Compatibility is defined by two things: first the version
// of the running library is newer than the version
// @required_major.required_minor.@required_micro. Second
// the running library must be binary compatible with the
// version @required_major.required_minor.@required_micro
// (same major version.)
//
// This function is primarily for GTK modules; the module
// can call this function to check that it wasn’t loaded
// into an incompatible version of GTK. However, such a
// check isn’t completely reliable, since the module may be
// linked against an old version of GTK and calling the
// old version of gtk_check_version(), but still get loaded
// into an application using a newer version of GTK.
func CheckVersion(RequiredMajorVar uint, RequiredMinorVar uint, RequiredMicroVar uint) string {

	cret := xCheckVersion(RequiredMajorVar, RequiredMinorVar, RequiredMicroVar)
	return cret
}

var xGetBinaryAge func() uint

// Returns the binary age as passed to `libtool`.
//
// If `libtool` means nothing to you, don't worry about it.
func GetBinaryAge() uint {

	cret := xGetBinaryAge()
	return cret
}

var xGetInterfaceAge func() uint

// Returns the interface age as passed to `libtool`.
//
// If `libtool` means nothing to you, don't worry about it.
func GetInterfaceAge() uint {

	cret := xGetInterfaceAge()
	return cret
}

var xGetMajorVersion func() uint

// Returns the major version number of the GTK library.
//
// For example, in GTK version 3.1.5 this is 3.
//
// This function is in the library, so it represents the GTK library
// your code is running against. Contrast with the %GTK_MAJOR_VERSION
// macro, which represents the major version of the GTK headers you
// have included when compiling your code.
func GetMajorVersion() uint {

	cret := xGetMajorVersion()
	return cret
}

var xGetMicroVersion func() uint

// Returns the micro version number of the GTK library.
//
// For example, in GTK version 3.1.5 this is 5.
//
// This function is in the library, so it represents the GTK library
// your code is are running against. Contrast with the
// %GTK_MICRO_VERSION macro, which represents the micro version of the
// GTK headers you have included when compiling your code.
func GetMicroVersion() uint {

	cret := xGetMicroVersion()
	return cret
}

var xGetMinorVersion func() uint

// Returns the minor version number of the GTK library.
//
// For example, in GTK version 3.1.5 this is 1.
//
// This function is in the library, so it represents the GTK library
// your code is are running against. Contrast with the
// %GTK_MINOR_VERSION macro, which represents the minor version of the
// GTK headers you have included when compiling your code.
func GetMinorVersion() uint {

	cret := xGetMinorVersion()
	return cret
}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GTK"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	core.PuregoSafeRegister(&xCheckVersion, lib, "gtk_check_version")
	core.PuregoSafeRegister(&xGetBinaryAge, lib, "gtk_get_binary_age")
	core.PuregoSafeRegister(&xGetInterfaceAge, lib, "gtk_get_interface_age")
	core.PuregoSafeRegister(&xGetMajorVersion, lib, "gtk_get_major_version")
	core.PuregoSafeRegister(&xGetMicroVersion, lib, "gtk_get_micro_version")
	core.PuregoSafeRegister(&xGetMinorVersion, lib, "gtk_get_minor_version")

}
