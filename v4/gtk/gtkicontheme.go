// Package gtk was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gtk

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/gdk"
	"github.com/jwijenbergh/puregotk/v4/gio"
	"github.com/jwijenbergh/puregotk/v4/gobject"
)

// Used to specify options for gtk_icon_theme_lookup_icon().
type IconLookupFlags int

const (

	// Try to always load regular icons, even
	//   when symbolic icon names are given
	IconLookupForceRegularValue IconLookupFlags = 1
	// Try to always load symbolic icons, even
	//   when regular icon names are given
	IconLookupForceSymbolicValue IconLookupFlags = 2
	// Starts loading the texture in the background
	//   so it is ready when later needed.
	IconLookupPreloadValue IconLookupFlags = 4
)

// Error codes for `GtkIconTheme` operations.
type IconThemeError int

const (

	// The icon specified does not exist in the theme
	IconThemeNotFoundValue IconThemeError = 0
	// An unspecified error occurred.
	IconThemeFailedValue IconThemeError = 1
)

// Contains information found when looking up an icon in `GtkIconTheme`.
//
// `GtkIconPaintable` implements `GdkPaintable`.
type IconPaintable struct {
	gobject.Object
}

func IconPaintableNewFromInternalPtr(ptr uintptr) *IconPaintable {
	cls := &IconPaintable{}
	cls.Ptr = ptr
	return cls
}

var xNewForFileIconPaintable func(uintptr, int, int) uintptr

// Creates a `GtkIconPaintable` for a file with a given size and scale.
//
// The icon can then be rendered by using it as a `GdkPaintable`.
func NewForFileIconPaintable(FileVar gio.File, SizeVar int, ScaleVar int) *IconPaintable {
	var cls *IconPaintable

	cret := xNewForFileIconPaintable(FileVar.GoPointer(), SizeVar, ScaleVar)

	if cret == 0 {
		return nil
	}
	cls = &IconPaintable{}
	cls.Ptr = cret
	return cls
}

var xIconPaintableGetFile func(uintptr) uintptr

// Gets the `GFile` that was used to load the icon.
//
// Returns %NULL if the icon was not loaded from a file.
func (x *IconPaintable) GetFile() *gio.FileBase {
	var cls *gio.FileBase

	cret := xIconPaintableGetFile(x.GoPointer())

	if cret == 0 {
		return nil
	}
	cls = &gio.FileBase{}
	cls.Ptr = cret
	return cls
}

var xIconPaintableGetIconName func(uintptr) string

// Get the icon name being used for this icon.
//
// When an icon looked up in the icon theme was not available, the
// icon theme may use fallback icons - either those specified to
// gtk_icon_theme_lookup_icon() or the always-available
// "image-missing". The icon chosen is returned by this function.
//
// If the icon was created without an icon theme, this function
// returns %NULL.
func (x *IconPaintable) GetIconName() string {

	cret := xIconPaintableGetIconName(x.GoPointer())
	return cret
}

var xIconPaintableIsSymbolic func(uintptr) bool

// Checks if the icon is symbolic or not.
//
// This currently uses only the file name and not the file contents
// for determining this. This behaviour may change in the future.
//
// Note that to render a symbolic `GtkIconPaintable` properly (with
// recoloring), you have to set its icon name on a `GtkImage`.
func (x *IconPaintable) IsSymbolic() bool {

	cret := xIconPaintableIsSymbolic(x.GoPointer())
	return cret
}

func (c *IconPaintable) GoPointer() uintptr {
	return c.Ptr
}

func (c *IconPaintable) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// Compute a concrete size for the `GdkPaintable`.
//
// Applies the sizing algorithm outlined in the
// [CSS Image spec](https://drafts.csswg.org/css-images-3/#default-sizing)
// to the given @paintable. See that link for more details.
//
// It is not necessary to call this function when both @specified_width
// and @specified_height are known, but it is useful to call this
// function in GtkWidget:measure implementations to compute the
// other dimension when only one dimension is given.
func (x *IconPaintable) ComputeConcreteSize(SpecifiedWidthVar float64, SpecifiedHeightVar float64, DefaultWidthVar float64, DefaultHeightVar float64, ConcreteWidthVar float64, ConcreteHeightVar float64) {

	gdk.XGdkPaintableComputeConcreteSize(x.GoPointer(), SpecifiedWidthVar, SpecifiedHeightVar, DefaultWidthVar, DefaultHeightVar, ConcreteWidthVar, ConcreteHeightVar)

}

// Gets an immutable paintable for the current contents displayed by @paintable.
//
// This is useful when you want to retain the current state of an animation,
// for example to take a screenshot of a running animation.
//
// If the @paintable is already immutable, it will return itself.
func (x *IconPaintable) GetCurrentImage() *gdk.PaintableBase {
	var cls *gdk.PaintableBase

	cret := gdk.XGdkPaintableGetCurrentImage(x.GoPointer())

	if cret == 0 {
		return nil
	}
	cls = &gdk.PaintableBase{}
	cls.Ptr = cret
	return cls
}

// Get flags for the paintable.
//
// This is oftentimes useful for optimizations.
//
// See [flags@Gdk.PaintableFlags] for the flags and what they mean.
func (x *IconPaintable) GetFlags() gdk.PaintableFlags {

	cret := gdk.XGdkPaintableGetFlags(x.GoPointer())
	return cret
}

// Gets the preferred aspect ratio the @paintable would like to be displayed at.
//
// The aspect ratio is the width divided by the height, so a value of 0.5
// means that the @paintable prefers to be displayed twice as high as it
// is wide. Consumers of this interface can use this to preserve aspect
// ratio when displaying the paintable.
//
// This is a purely informational value and does not in any way limit the
// values that may be passed to [method@Gdk.Paintable.snapshot].
//
// Usually when a @paintable returns nonzero values from
// [method@Gdk.Paintable.get_intrinsic_width] and
// [method@Gdk.Paintable.get_intrinsic_height] the aspect ratio
// should conform to those values, though that is not required.
//
// If the @paintable does not have a preferred aspect ratio,
// it returns 0. Negative values are never returned.
func (x *IconPaintable) GetIntrinsicAspectRatio() float64 {

	cret := gdk.XGdkPaintableGetIntrinsicAspectRatio(x.GoPointer())
	return cret
}

// Gets the preferred height the @paintable would like to be displayed at.
//
// Consumers of this interface can use this to reserve enough space to draw
// the paintable.
//
// This is a purely informational value and does not in any way limit the
// values that may be passed to [method@Gdk.Paintable.snapshot].
//
// If the @paintable does not have a preferred height, it returns 0.
// Negative values are never returned.
func (x *IconPaintable) GetIntrinsicHeight() int {

	cret := gdk.XGdkPaintableGetIntrinsicHeight(x.GoPointer())
	return cret
}

// Gets the preferred width the @paintable would like to be displayed at.
//
// Consumers of this interface can use this to reserve enough space to draw
// the paintable.
//
// This is a purely informational value and does not in any way limit the
// values that may be passed to [method@Gdk.Paintable.snapshot].
//
// If the @paintable does not have a preferred width, it returns 0.
// Negative values are never returned.
func (x *IconPaintable) GetIntrinsicWidth() int {

	cret := gdk.XGdkPaintableGetIntrinsicWidth(x.GoPointer())
	return cret
}

// Called by implementations of `GdkPaintable` to invalidate their contents.
//
// Unless the contents are invalidated, implementations must guarantee that
// multiple calls of [method@Gdk.Paintable.snapshot] produce the same output.
//
// This function will emit the [signal@Gdk.Paintable::invalidate-contents]
// signal.
//
// If a @paintable reports the %GDK_PAINTABLE_STATIC_CONTENTS flag,
// it must not call this function.
func (x *IconPaintable) InvalidateContents() {

	gdk.XGdkPaintableInvalidateContents(x.GoPointer())

}

// Called by implementations of `GdkPaintable` to invalidate their size.
//
// As long as the size is not invalidated, @paintable must return the same
// values for its intrinsic width, height and aspect ratio.
//
// This function will emit the [signal@Gdk.Paintable::invalidate-size]
// signal.
//
// If a @paintable reports the %GDK_PAINTABLE_STATIC_SIZE flag,
// it must not call this function.
func (x *IconPaintable) InvalidateSize() {

	gdk.XGdkPaintableInvalidateSize(x.GoPointer())

}

// Snapshots the given paintable with the given @width and @height.
//
// The paintable is drawn at the current (0,0) offset of the @snapshot.
// If @width and @height are not larger than zero, this function will
// do nothing.
func (x *IconPaintable) Snapshot(SnapshotVar *gdk.Snapshot, WidthVar float64, HeightVar float64) {

	gdk.XGdkPaintableSnapshot(x.GoPointer(), SnapshotVar.GoPointer(), WidthVar, HeightVar)

}

// Snapshots the paintable with the given colors.
//
// If less than 4 colors are provided, GTK will pad the array with default
// colors.
func (x *IconPaintable) SnapshotSymbolic(SnapshotVar *gdk.Snapshot, WidthVar float64, HeightVar float64, ColorsVar uintptr, NColorsVar uint) {

	XGtkSymbolicPaintableSnapshotSymbolic(x.GoPointer(), SnapshotVar.GoPointer(), WidthVar, HeightVar, ColorsVar, NColorsVar)

}

// `GtkIconTheme` provides a facility for loading themed icons.
//
// The main reason for using a name rather than simply providing a filename
// is to allow different icons to be used depending on what “icon theme” is
// selected by the user. The operation of icon themes on Linux and Unix
// follows the [Icon Theme Specification](http://www.freedesktop.org/Standards/icon-theme-spec)
// There is a fallback icon theme, named `hicolor`, where applications
// should install their icons, but additional icon themes can be installed
// as operating system vendors and users choose.
//
// In many cases, named themes are used indirectly, via [class@Gtk.Image]
// rather than directly, but looking up icons directly is also simple. The
// `GtkIconTheme` object acts as a database of all the icons in the current
// theme. You can create new `GtkIconTheme` objects, but it’s much more
// efficient to use the standard icon theme of the `GtkWidget` so that the
// icon information is shared with other people looking up icons.
//
// ```c
// GtkIconTheme *icon_theme;
// GtkIconPaintable *icon;
// GdkPaintable *paintable;
//
// icon_theme = gtk_icon_theme_get_for_display (gtk_widget_get_display (my_widget));
// icon = gtk_icon_theme_lookup_icon (icon_theme,
//
//	"my-icon-name", // icon name
//	48, // icon size
//	1,  // scale
//	0,  // flags);
//
// paintable = GDK_PAINTABLE (icon);
// // Use the paintable
// g_object_unref (icon);
// ```
type IconTheme struct {
	gobject.Object
}

func IconThemeNewFromInternalPtr(ptr uintptr) *IconTheme {
	cls := &IconTheme{}
	cls.Ptr = ptr
	return cls
}

var xNewIconTheme func() uintptr

// Creates a new icon theme object.
//
// Icon theme objects are used to lookup up an icon by name
// in a particular icon theme. Usually, you’ll want to use
// [func@Gtk.IconTheme.get_for_display] rather than creating
// a new icon theme object for scratch.
func NewIconTheme() *IconTheme {
	var cls *IconTheme

	cret := xNewIconTheme()

	if cret == 0 {
		return nil
	}
	cls = &IconTheme{}
	cls.Ptr = cret
	return cls
}

var xIconThemeAddResourcePath func(uintptr, string)

// Adds a resource path that will be looked at when looking
// for icons, similar to search paths.
//
// See [method@Gtk.IconTheme.set_resource_path].
//
// This function should be used to make application-specific icons
// available as part of the icon theme.
func (x *IconTheme) AddResourcePath(PathVar string) {

	xIconThemeAddResourcePath(x.GoPointer(), PathVar)

}

var xIconThemeAddSearchPath func(uintptr, string)

// Appends a directory to the search path.
//
// See [method@Gtk.IconTheme.set_search_path].
func (x *IconTheme) AddSearchPath(PathVar string) {

	xIconThemeAddSearchPath(x.GoPointer(), PathVar)

}

var xIconThemeGetDisplay func(uintptr) uintptr

// Returns the display that the `GtkIconTheme` object was
// created for.
func (x *IconTheme) GetDisplay() *gdk.Display {
	var cls *gdk.Display

	cret := xIconThemeGetDisplay(x.GoPointer())

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &gdk.Display{}
	cls.Ptr = cret
	return cls
}

var xIconThemeGetIconNames func(uintptr) []string

// Lists the names of icons in the current icon theme.
func (x *IconTheme) GetIconNames() []string {

	cret := xIconThemeGetIconNames(x.GoPointer())
	return cret
}

var xIconThemeGetIconSizes func(uintptr, string) uintptr

// Returns an array of integers describing the sizes at which
// the icon is available without scaling.
//
// A size of -1 means that the icon is available in a scalable
// format. The array is zero-terminated.
func (x *IconTheme) GetIconSizes(IconNameVar string) uintptr {

	cret := xIconThemeGetIconSizes(x.GoPointer(), IconNameVar)
	return cret
}

var xIconThemeGetResourcePath func(uintptr) []string

// Gets the current resource path.
//
// See [method@Gtk.IconTheme.set_resource_path].
func (x *IconTheme) GetResourcePath() []string {

	cret := xIconThemeGetResourcePath(x.GoPointer())
	return cret
}

var xIconThemeGetSearchPath func(uintptr) []string

// Gets the current search path.
//
// See [method@Gtk.IconTheme.set_search_path].
func (x *IconTheme) GetSearchPath() []string {

	cret := xIconThemeGetSearchPath(x.GoPointer())
	return cret
}

var xIconThemeGetThemeName func(uintptr) string

// Gets the current icon theme name.
//
// Returns (transfer full): the current icon theme name,
func (x *IconTheme) GetThemeName() string {

	cret := xIconThemeGetThemeName(x.GoPointer())
	return cret
}

var xIconThemeHasGicon func(uintptr, uintptr) bool

// Checks whether an icon theme includes an icon
// for a particular `GIcon`.
func (x *IconTheme) HasGicon(GiconVar gio.Icon) bool {

	cret := xIconThemeHasGicon(x.GoPointer(), GiconVar.GoPointer())
	return cret
}

var xIconThemeHasIcon func(uintptr, string) bool

// Checks whether an icon theme includes an icon
// for a particular name.
func (x *IconTheme) HasIcon(IconNameVar string) bool {

	cret := xIconThemeHasIcon(x.GoPointer(), IconNameVar)
	return cret
}

var xIconThemeLookupByGicon func(uintptr, uintptr, int, int, TextDirection, IconLookupFlags) uintptr

// Looks up a icon for a desired size and window scale.
//
// The icon can then be rendered by using it as a `GdkPaintable`,
// or you can get information such as the filename and size.
func (x *IconTheme) LookupByGicon(IconVar gio.Icon, SizeVar int, ScaleVar int, DirectionVar TextDirection, FlagsVar IconLookupFlags) *IconPaintable {
	var cls *IconPaintable

	cret := xIconThemeLookupByGicon(x.GoPointer(), IconVar.GoPointer(), SizeVar, ScaleVar, DirectionVar, FlagsVar)

	if cret == 0 {
		return nil
	}
	cls = &IconPaintable{}
	cls.Ptr = cret
	return cls
}

var xIconThemeLookupIcon func(uintptr, string, uintptr, int, int, TextDirection, IconLookupFlags) uintptr

// Looks up a named icon for a desired size and window scale,
// returning a `GtkIconPaintable`.
//
// The icon can then be rendered by using it as a `GdkPaintable`,
// or you can get information such as the filename and size.
//
// If the available @icon_name is not available and @fallbacks are
// provided, they will be tried in order.
//
// If no matching icon is found, then a paintable that renders the
// "missing icon" icon is returned. If you need to do something else
// for missing icons you need to use [method@Gtk.IconTheme.has_icon].
//
// Note that you probably want to listen for icon theme changes and
// update the icon. This is usually done by overriding the
// GtkWidgetClass.css-changed() function.
func (x *IconTheme) LookupIcon(IconNameVar string, FallbacksVar uintptr, SizeVar int, ScaleVar int, DirectionVar TextDirection, FlagsVar IconLookupFlags) *IconPaintable {
	var cls *IconPaintable

	cret := xIconThemeLookupIcon(x.GoPointer(), IconNameVar, FallbacksVar, SizeVar, ScaleVar, DirectionVar, FlagsVar)

	if cret == 0 {
		return nil
	}
	cls = &IconPaintable{}
	cls.Ptr = cret
	return cls
}

var xIconThemeSetResourcePath func(uintptr, uintptr)

// Sets the resource paths that will be looked at when
// looking for icons, similar to search paths.
//
// The resources are considered as part of the hicolor icon theme
// and must be located in subdirectories that are defined in the
// hicolor icon theme, such as `@path/16x16/actions/run.png`
// or `@path/scalable/actions/run.svg`.
//
// Icons that are directly placed in the resource path instead
// of a subdirectory are also considered as ultimate fallback,
// but they are treated like unthemed icons.
func (x *IconTheme) SetResourcePath(PathVar uintptr) {

	xIconThemeSetResourcePath(x.GoPointer(), PathVar)

}

var xIconThemeSetSearchPath func(uintptr, uintptr)

// Sets the search path for the icon theme object.
//
// When looking for an icon theme, GTK will search for a subdirectory
// of one or more of the directories in @path with the same name
// as the icon theme containing an index.theme file. (Themes from
// multiple of the path elements are combined to allow themes to be
// extended by adding icons in the user’s home directory.)
//
// In addition if an icon found isn’t found either in the current
// icon theme or the default icon theme, and an image file with
// the right name is found directly in one of the elements of
// @path, then that image will be used for the icon name.
// (This is legacy feature, and new icons should be put
// into the fallback icon theme, which is called hicolor,
// rather than directly on the icon path.)
func (x *IconTheme) SetSearchPath(PathVar uintptr) {

	xIconThemeSetSearchPath(x.GoPointer(), PathVar)

}

var xIconThemeSetThemeName func(uintptr, string)

// Sets the name of the icon theme that the `GtkIconTheme` object uses
// overriding system configuration.
//
// This function cannot be called on the icon theme objects returned
// from [func@Gtk.IconTheme.get_for_display].
func (x *IconTheme) SetThemeName(ThemeNameVar string) {

	xIconThemeSetThemeName(x.GoPointer(), ThemeNameVar)

}

func (c *IconTheme) GoPointer() uintptr {
	return c.Ptr
}

func (c *IconTheme) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// Emitted when the icon theme changes.
//
// This can happen becuase current icon theme is switched or
// because GTK detects that a change has occurred in the
// contents of the current icon theme.
func (x *IconTheme) ConnectChanged(cb func(IconTheme)) {
	fcb := func(clsPtr uintptr) {
		fa := IconTheme{}
		fa.Ptr = clsPtr

		cb(fa)

	}
	gobject.ObjectConnect(x.GoPointer(), "signal::changed", purego.NewCallback(fcb))
}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GTK"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xNewForFileIconPaintable, lib, "gtk_icon_paintable_new_for_file")

	core.PuregoSafeRegister(&xIconPaintableGetFile, lib, "gtk_icon_paintable_get_file")
	core.PuregoSafeRegister(&xIconPaintableGetIconName, lib, "gtk_icon_paintable_get_icon_name")
	core.PuregoSafeRegister(&xIconPaintableIsSymbolic, lib, "gtk_icon_paintable_is_symbolic")

	core.PuregoSafeRegister(&xNewIconTheme, lib, "gtk_icon_theme_new")

	core.PuregoSafeRegister(&xIconThemeAddResourcePath, lib, "gtk_icon_theme_add_resource_path")
	core.PuregoSafeRegister(&xIconThemeAddSearchPath, lib, "gtk_icon_theme_add_search_path")
	core.PuregoSafeRegister(&xIconThemeGetDisplay, lib, "gtk_icon_theme_get_display")
	core.PuregoSafeRegister(&xIconThemeGetIconNames, lib, "gtk_icon_theme_get_icon_names")
	core.PuregoSafeRegister(&xIconThemeGetIconSizes, lib, "gtk_icon_theme_get_icon_sizes")
	core.PuregoSafeRegister(&xIconThemeGetResourcePath, lib, "gtk_icon_theme_get_resource_path")
	core.PuregoSafeRegister(&xIconThemeGetSearchPath, lib, "gtk_icon_theme_get_search_path")
	core.PuregoSafeRegister(&xIconThemeGetThemeName, lib, "gtk_icon_theme_get_theme_name")
	core.PuregoSafeRegister(&xIconThemeHasGicon, lib, "gtk_icon_theme_has_gicon")
	core.PuregoSafeRegister(&xIconThemeHasIcon, lib, "gtk_icon_theme_has_icon")
	core.PuregoSafeRegister(&xIconThemeLookupByGicon, lib, "gtk_icon_theme_lookup_by_gicon")
	core.PuregoSafeRegister(&xIconThemeLookupIcon, lib, "gtk_icon_theme_lookup_icon")
	core.PuregoSafeRegister(&xIconThemeSetResourcePath, lib, "gtk_icon_theme_set_resource_path")
	core.PuregoSafeRegister(&xIconThemeSetSearchPath, lib, "gtk_icon_theme_set_search_path")
	core.PuregoSafeRegister(&xIconThemeSetThemeName, lib, "gtk_icon_theme_set_theme_name")

}
