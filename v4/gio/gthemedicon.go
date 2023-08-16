// Package gio was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gio

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/glib"
	"github.com/jwijenbergh/puregotk/v4/gobject"
)

type ThemedIconClass struct {
}

// #GThemedIcon is an implementation of #GIcon that supports icon themes.
// #GThemedIcon contains a list of all of the icons present in an icon
// theme, so that icons can be looked up quickly. #GThemedIcon does
// not provide actual pixmaps for icons, just the icon names.
// Ideally something like gtk_icon_theme_choose_icon() should be used to
// resolve the list of names so that fallback icons work nicely with
// themes that inherit other themes.
type ThemedIcon struct {
	gobject.Object
}

func ThemedIconNewFromInternalPtr(ptr uintptr) *ThemedIcon {
	cls := &ThemedIcon{}
	cls.Ptr = ptr
	return cls
}

var xNewThemedIcon func(string) uintptr

// Creates a new themed icon for @iconname.
func NewThemedIcon(IconnameVar string) *ThemedIcon {
	var cls *ThemedIcon

	cret := xNewThemedIcon(IconnameVar)

	if cret == 0 {
		return cls
	}
	cls = &ThemedIcon{}
	cls.Ptr = cret
	return cls
}

var xNewFromNamesThemedIcon func([]string, int) uintptr

// Creates a new themed icon for @iconnames.
func NewFromNamesThemedIcon(IconnamesVar []string, LenVar int) *ThemedIcon {
	var cls *ThemedIcon

	cret := xNewFromNamesThemedIcon(IconnamesVar, LenVar)

	if cret == 0 {
		return cls
	}
	cls = &ThemedIcon{}
	cls.Ptr = cret
	return cls
}

var xNewWithDefaultFallbacksThemedIcon func(string) uintptr

// Creates a new themed icon for @iconname, and all the names
// that can be created by shortening @iconname at '-' characters.
//
// In the following example, @icon1 and @icon2 are equivalent:
// |[&lt;!-- language="C" --&gt;
//
//	const char *names[] = {
//	  "gnome-dev-cdrom-audio",
//	  "gnome-dev-cdrom",
//	  "gnome-dev",
//	  "gnome"
//	};
//
// icon1 = g_themed_icon_new_from_names (names, 4);
// icon2 = g_themed_icon_new_with_default_fallbacks ("gnome-dev-cdrom-audio");
// ]|
func NewWithDefaultFallbacksThemedIcon(IconnameVar string) *ThemedIcon {
	var cls *ThemedIcon

	cret := xNewWithDefaultFallbacksThemedIcon(IconnameVar)

	if cret == 0 {
		return cls
	}
	cls = &ThemedIcon{}
	cls.Ptr = cret
	return cls
}

var xThemedIconAppendName func(uintptr, string)

// Append a name to the list of icons from within @icon.
//
// Note that doing so invalidates the hash computed by prior calls
// to g_icon_hash().
func (x *ThemedIcon) AppendName(IconnameVar string) {

	xThemedIconAppendName(x.GoPointer(), IconnameVar)

}

var xThemedIconGetNames func(uintptr) uintptr

// Gets the names of icons from within @icon.
func (x *ThemedIcon) GetNames() uintptr {

	cret := xThemedIconGetNames(x.GoPointer())
	return cret
}

var xThemedIconPrependName func(uintptr, string)

// Prepend a name to the list of icons from within @icon.
//
// Note that doing so invalidates the hash computed by prior calls
// to g_icon_hash().
func (x *ThemedIcon) PrependName(IconnameVar string) {

	xThemedIconPrependName(x.GoPointer(), IconnameVar)

}

func (c *ThemedIcon) GoPointer() uintptr {
	return c.Ptr
}

func (c *ThemedIcon) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// Checks if two icons are equal.
func (x *ThemedIcon) Equal(Icon2Var Icon) bool {

	cret := XGIconEqual(x.GoPointer(), Icon2Var.GoPointer())
	return cret
}

// Serializes a #GIcon into a #GVariant. An equivalent #GIcon can be retrieved
// back by calling g_icon_deserialize() on the returned value.
// As serialization will avoid using raw icon data when possible, it only
// makes sense to transfer the #GVariant between processes on the same machine,
// (as opposed to over the network), and within the same file system namespace.
func (x *ThemedIcon) Serialize() *glib.Variant {

	cret := XGIconSerialize(x.GoPointer())
	return cret
}

// Generates a textual representation of @icon that can be used for
// serialization such as when passing @icon to a different process or
// saving it to persistent storage. Use g_icon_new_for_string() to
// get @icon back from the returned string.
//
// The encoding of the returned string is proprietary to #GIcon except
// in the following two cases
//
//   - If @icon is a #GFileIcon, the returned string is a native path
//     (such as `/path/to/my icon.png`) without escaping
//     if the #GFile for @icon is a native file.  If the file is not
//     native, the returned string is the result of g_file_get_uri()
//     (such as `sftp://path/to/my%20icon.png`).
//
//   - If @icon is a #GThemedIcon with exactly one name and no fallbacks,
//     the encoding is simply the name (such as `network-server`).
func (x *ThemedIcon) ToString() string {

	cret := XGIconToString(x.GoPointer())
	return cret
}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GIO"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xNewThemedIcon, lib, "g_themed_icon_new")
	core.PuregoSafeRegister(&xNewFromNamesThemedIcon, lib, "g_themed_icon_new_from_names")
	core.PuregoSafeRegister(&xNewWithDefaultFallbacksThemedIcon, lib, "g_themed_icon_new_with_default_fallbacks")

	core.PuregoSafeRegister(&xThemedIconAppendName, lib, "g_themed_icon_append_name")
	core.PuregoSafeRegister(&xThemedIconGetNames, lib, "g_themed_icon_get_names")
	core.PuregoSafeRegister(&xThemedIconPrependName, lib, "g_themed_icon_prepend_name")

}
