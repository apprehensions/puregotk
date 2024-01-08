// Package gdkpixbuf was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gdkpixbuf

import (
	"unsafe"

	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/glib"
	"github.com/jwijenbergh/puregotk/v4/gmodule"
)

// Sets up the image loading state.
//
// The image loader is responsible for storing the given function pointers
// and user data, and call them when needed.
//
// The image loader should set up an internal state object, and return it
// from this function; the state object will then be updated from the
// [callback@GdkPixbuf.PixbufModuleIncrementLoadFunc] callback, and will be freed
// by [callback@GdkPixbuf.PixbufModuleStopLoadFunc] callback.
type PixbufModuleBeginLoadFunc func(uintptr, uintptr, uintptr, uintptr, **glib.Error) uintptr

// Defines the type of the function used to fill a
// #GdkPixbufFormat structure with information about a module.
type PixbufModuleFillInfoFunc func(*PixbufFormat)

// Defines the type of the function used to set the vtable of a
// #GdkPixbufModule when it is loaded.
type PixbufModuleFillVtableFunc func(*PixbufModule)

// Incrementally loads a buffer into the image data.
type PixbufModuleIncrementLoadFunc func(uintptr, uintptr, uint, **glib.Error) bool

// Loads a file from a standard C file stream into a new `GdkPixbufAnimation`.
//
// In case of error, this function should return `NULL` and set the `error` argument.
type PixbufModuleLoadAnimationFunc func(uintptr, **glib.Error) uintptr

// Loads a file from a standard C file stream into a new `GdkPixbuf`.
//
// In case of error, this function should return `NULL` and set the `error` argument.
type PixbufModuleLoadFunc func(uintptr, **glib.Error) uintptr

// Loads XPM data into a new `GdkPixbuf`.
type PixbufModuleLoadXpmDataFunc func([]string) uintptr

// Defines the type of the function that gets called once the initial
// setup of @pixbuf is done.
//
// #GdkPixbufLoader uses a function of this type to emit the
// "&lt;link linkend="GdkPixbufLoader-area-prepared"&gt;area_prepared&lt;/link&gt;"
// signal.
type PixbufModulePreparedFunc func(uintptr, uintptr, uintptr)

// Saves a `GdkPixbuf` by calling the provided function.
//
// The optional `option_keys` and `option_values` arrays contain the keys and
// values (in the same order) for attributes to be saved alongside the image
// data.
type PixbufModuleSaveCallbackFunc func(uintptr, uintptr, uintptr, uintptr, uintptr, **glib.Error) bool

// Saves a `GdkPixbuf` into a standard C file stream.
//
// The optional `param_keys` and `param_values` arrays contain the keys and
// values (in the same order) for attributes to be saved alongside the image
// data.
type PixbufModuleSaveFunc func(uintptr, uintptr, uintptr, uintptr, **glib.Error) bool

// Checks whether the given `option_key` is supported when saving.
type PixbufModuleSaveOptionSupportedFunc func(string) bool

// Defines the type of the function that gets called once the size
// of the loaded image is known.
//
// The function is expected to set @width and @height to the desired
// size to which the image should be scaled. If a module has no efficient
// way to achieve the desired scaling during the loading of the image, it may
// either ignore the size request, or only approximate it - gdk-pixbuf will
// then perform the required scaling on the completely loaded image.
//
// If the function sets @width or @height to zero, the module should interpret
// this as a hint that it will be closed soon and shouldn't allocate further
// resources. This convention is used to implement gdk_pixbuf_get_file_info()
// efficiently.
type PixbufModuleSizeFunc func(int, int, uintptr)

// Finalizes the image loading state.
//
// This function is called on success and error states.
type PixbufModuleStopLoadFunc func(uintptr, **glib.Error) bool

// Defines the type of the function that gets called every time a region
// of @pixbuf is updated.
//
// #GdkPixbufLoader uses a function of this type to emit the
// "&lt;link linkend="GdkPixbufLoader-area-updated"&gt;area_updated&lt;/link&gt;"
// signal.
type PixbufModuleUpdatedFunc func(uintptr, int, int, int, int, uintptr)

// A `GdkPixbufFormat` contains information about the image format accepted
// by a module.
//
// Only modules should access the fields directly, applications should
// use the `gdk_pixbuf_format_*` family of functions.
type PixbufFormat struct {
	Name uintptr

	Signature *PixbufModulePattern

	Domain uintptr

	Description uintptr

	MimeTypes uintptr

	Extensions uintptr

	Flags uint32

	Disabled bool

	License uintptr
}

func (x *PixbufFormat) GoPointer() uintptr {
	return uintptr(unsafe.Pointer(x))
}

var xPixbufFormatCopy func(uintptr) *PixbufFormat

// Creates a copy of `format`.
func (x *PixbufFormat) Copy() *PixbufFormat {

	cret := xPixbufFormatCopy(x.GoPointer())
	return cret
}

var xPixbufFormatFree func(uintptr)

// Frees the resources allocated when copying a `GdkPixbufFormat`
// using gdk_pixbuf_format_copy()
func (x *PixbufFormat) Free() {

	xPixbufFormatFree(x.GoPointer())

}

var xPixbufFormatGetDescription func(uintptr) string

// Returns a description of the format.
func (x *PixbufFormat) GetDescription() string {

	cret := xPixbufFormatGetDescription(x.GoPointer())
	return cret
}

var xPixbufFormatGetExtensions func(uintptr) uintptr

// Returns the filename extensions typically used for files in the
// given format.
func (x *PixbufFormat) GetExtensions() uintptr {

	cret := xPixbufFormatGetExtensions(x.GoPointer())
	return cret
}

var xPixbufFormatGetLicense func(uintptr) string

// Returns information about the license of the image loader for the format.
//
// The returned string should be a shorthand for a well known license, e.g.
// "LGPL", "GPL", "QPL", "GPL/QPL", or "other" to indicate some other license.
func (x *PixbufFormat) GetLicense() string {

	cret := xPixbufFormatGetLicense(x.GoPointer())
	return cret
}

var xPixbufFormatGetMimeTypes func(uintptr) uintptr

// Returns the mime types supported by the format.
func (x *PixbufFormat) GetMimeTypes() uintptr {

	cret := xPixbufFormatGetMimeTypes(x.GoPointer())
	return cret
}

var xPixbufFormatGetName func(uintptr) string

// Returns the name of the format.
func (x *PixbufFormat) GetName() string {

	cret := xPixbufFormatGetName(x.GoPointer())
	return cret
}

var xPixbufFormatIsDisabled func(uintptr) bool

// Returns whether this image format is disabled.
//
// See gdk_pixbuf_format_set_disabled().
func (x *PixbufFormat) IsDisabled() bool {

	cret := xPixbufFormatIsDisabled(x.GoPointer())
	return cret
}

var xPixbufFormatIsSaveOptionSupported func(uintptr, string) bool

// Returns `TRUE` if the save option specified by @option_key is supported when
// saving a pixbuf using the module implementing @format.
//
// See gdk_pixbuf_save() for more information about option keys.
func (x *PixbufFormat) IsSaveOptionSupported(OptionKeyVar string) bool {

	cret := xPixbufFormatIsSaveOptionSupported(x.GoPointer(), OptionKeyVar)
	return cret
}

var xPixbufFormatIsScalable func(uintptr) bool

// Returns whether this image format is scalable.
//
// If a file is in a scalable format, it is preferable to load it at
// the desired size, rather than loading it at the default size and
// scaling the resulting pixbuf to the desired size.
func (x *PixbufFormat) IsScalable() bool {

	cret := xPixbufFormatIsScalable(x.GoPointer())
	return cret
}

var xPixbufFormatIsWritable func(uintptr) bool

// Returns whether pixbufs can be saved in the given format.
func (x *PixbufFormat) IsWritable() bool {

	cret := xPixbufFormatIsWritable(x.GoPointer())
	return cret
}

var xPixbufFormatSetDisabled func(uintptr, bool)

// Disables or enables an image format.
//
// If a format is disabled, GdkPixbuf won't use the image loader for
// this format to load images.
//
// Applications can use this to avoid using image loaders with an
// inappropriate license, see gdk_pixbuf_format_get_license().
func (x *PixbufFormat) SetDisabled(DisabledVar bool) {

	xPixbufFormatSetDisabled(x.GoPointer(), DisabledVar)

}

// A `GdkPixbufModule` contains the necessary functions to load and save
// images in a certain file format.
//
// If `GdkPixbuf` has been compiled with `GModule` support, it can be extended
// by modules which can load (and perhaps also save) new image and animation
// formats.
//
// ## Implementing modules
//
// The `GdkPixbuf` interfaces needed for implementing modules are contained in
// `gdk-pixbuf-io.h` (and `gdk-pixbuf-animation.h` if the module supports
// animations). They are not covered by the same stability guarantees as the
// regular GdkPixbuf API. To underline this fact, they are protected by the
// `GDK_PIXBUF_ENABLE_BACKEND` pre-processor symbol.
//
// Each loadable module must contain a `GdkPixbufModuleFillVtableFunc` function
// named `fill_vtable`, which will get called when the module
// is loaded and must set the function pointers of the `GdkPixbufModule`.
//
// In order to make format-checking work before actually loading the modules
// (which may require calling `dlopen` to load image libraries), modules export
// their signatures (and other information) via the `fill_info` function. An
// external utility, `gdk-pixbuf-query-loaders`, uses this to create a text
// file containing a list of all available loaders and  their signatures.
// This file is then read at runtime by `GdkPixbuf` to obtain the list of
// available loaders and their signatures.
//
// Modules may only implement a subset of the functionality available via
// `GdkPixbufModule`. If a particular functionality is not implemented, the
// `fill_vtable` function will simply not set the corresponding
// function pointers of the `GdkPixbufModule` structure. If a module supports
// incremental loading (i.e. provides `begin_load`, `stop_load` and
// `load_increment`), it doesn't have to implement `load`, since `GdkPixbuf`
// can supply a generic `load` implementation wrapping the incremental loading.
//
// ## Installing modules
//
// Installing a module is a two-step process:
//
//   - copy the module file(s) to the loader directory (normally
//     `$libdir/gdk-pixbuf-2.0/$version/loaders`, unless overridden by the
//     environment variable `GDK_PIXBUF_MODULEDIR`)
//   - call `gdk-pixbuf-query-loaders` to update the module file (normally
//     `$libdir/gdk-pixbuf-2.0/$version/loaders.cache`, unless overridden
//     by the environment variable `GDK_PIXBUF_MODULE_FILE`)
type PixbufModule struct {
	ModuleName uintptr

	ModulePath uintptr

	Module *gmodule.Module

	Info *PixbufFormat

	Load PixbufModuleLoadFunc

	LoadXpmData PixbufModuleLoadXpmDataFunc

	BeginLoad PixbufModuleBeginLoadFunc

	StopLoad PixbufModuleStopLoadFunc

	LoadIncrement PixbufModuleIncrementLoadFunc

	LoadAnimation PixbufModuleLoadAnimationFunc

	Save PixbufModuleSaveFunc

	SaveToCallback PixbufModuleSaveCallbackFunc

	IsSaveOptionSupported PixbufModuleSaveOptionSupportedFunc
}

func (x *PixbufModule) GoPointer() uintptr {
	return uintptr(unsafe.Pointer(x))
}

// The signature prefix for a module.
//
// The signature of a module is a set of prefixes. Prefixes are encoded as
// pairs of ordinary strings, where the second string, called the mask, if
// not `NULL`, must be of the same length as the first one and may contain
// ' ', '!', 'x', 'z', and 'n' to indicate bytes that must be matched,
// not matched, "don't-care"-bytes, zeros and non-zeros, respectively.
//
// Each prefix has an associated integer that describes the relevance of
// the prefix, with 0 meaning a mismatch and 100 a "perfect match".
//
// Starting with gdk-pixbuf 2.8, the first byte of the mask may be '*',
// indicating an unanchored pattern that matches not only at the beginning,
// but also in the middle. Versions prior to 2.8 will interpret the '*'
// like an 'x'.
//
// The signature of a module is stored as an array of
// `GdkPixbufModulePatterns`. The array is terminated by a pattern
// where the `prefix` is `NULL`.
//
// ```c
//
//	GdkPixbufModulePattern *signature[] = {
//	  { "abcdx", " !x z", 100 },
//	  { "bla", NULL,  90 },
//	  { NULL, NULL, 0 }
//	};
//
// ```
//
// In the example above, the signature matches e.g. "auud\0" with
// relevance 100, and "blau" with relevance 90.
type PixbufModulePattern struct {
	Prefix uintptr

	Mask uintptr

	Relevance int
}

func (x *PixbufModulePattern) GoPointer() uintptr {
	return uintptr(unsafe.Pointer(x))
}

// Flags which allow a module to specify further details about the supported
// operations.
type PixbufFormatFlags int

const (

	// the module can write out images in the format.
	GdkPixbufFormatWritableValue PixbufFormatFlags = 1
	// the image format is scalable
	GdkPixbufFormatScalableValue PixbufFormatFlags = 2
	// the module is threadsafe. gdk-pixbuf
	//     ignores modules that are not marked as threadsafe. (Since 2.28).
	GdkPixbufFormatThreadsafeValue PixbufFormatFlags = 4
)

func init() {
	lib, err := purego.Dlopen(core.GetPath("GDKPIXBUF"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xPixbufFormatCopy, lib, "gdk_pixbuf_format_copy")
	core.PuregoSafeRegister(&xPixbufFormatFree, lib, "gdk_pixbuf_format_free")
	core.PuregoSafeRegister(&xPixbufFormatGetDescription, lib, "gdk_pixbuf_format_get_description")
	core.PuregoSafeRegister(&xPixbufFormatGetExtensions, lib, "gdk_pixbuf_format_get_extensions")
	core.PuregoSafeRegister(&xPixbufFormatGetLicense, lib, "gdk_pixbuf_format_get_license")
	core.PuregoSafeRegister(&xPixbufFormatGetMimeTypes, lib, "gdk_pixbuf_format_get_mime_types")
	core.PuregoSafeRegister(&xPixbufFormatGetName, lib, "gdk_pixbuf_format_get_name")
	core.PuregoSafeRegister(&xPixbufFormatIsDisabled, lib, "gdk_pixbuf_format_is_disabled")
	core.PuregoSafeRegister(&xPixbufFormatIsSaveOptionSupported, lib, "gdk_pixbuf_format_is_save_option_supported")
	core.PuregoSafeRegister(&xPixbufFormatIsScalable, lib, "gdk_pixbuf_format_is_scalable")
	core.PuregoSafeRegister(&xPixbufFormatIsWritable, lib, "gdk_pixbuf_format_is_writable")
	core.PuregoSafeRegister(&xPixbufFormatSetDisabled, lib, "gdk_pixbuf_format_set_disabled")

}
