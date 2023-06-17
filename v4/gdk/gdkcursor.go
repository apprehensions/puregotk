// Package gdk was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gdk

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/gobject"
)

// `GdkCursor` is used to create and destroy cursors.
//
// Cursors are immutable objects, so once you created them, there is no way
// to modify them later. You should create a new cursor when you want to change
// something about it.
//
// Cursors by themselves are not very interesting: they must be bound to a
// window for users to see them. This is done with [method@Gdk.Surface.set_cursor]
// or [method@Gdk.Surface.set_device_cursor]. Applications will typically
// use higher-level GTK functions such as [method@Gtk.Widget.set_cursor] instead.
//
// Cursors are not bound to a given [class@Gdk.Display], so they can be shared.
// However, the appearance of cursors may vary when used on different
// platforms.
//
// ## Named and texture cursors
//
// There are multiple ways to create cursors. The platform's own cursors
// can be created with [ctor@Gdk.Cursor.new_from_name]. That function lists
// the commonly available names that are shared with the CSS specification.
// Other names may be available, depending on the platform in use. On some
// platforms, what images are used for named cursors may be influenced by
// the cursor theme.
//
// Another option to create a cursor is to use [ctor@Gdk.Cursor.new_from_texture]
// and provide an image to use for the cursor.
//
// To ease work with unsupported cursors, a fallback cursor can be provided.
// If a [class@Gdk.Surface] cannot use a cursor because of the reasons mentioned
// above, it will try the fallback cursor. Fallback cursors can themselves have
// fallback cursors again, so it is possible to provide a chain of progressively
// easier to support cursors. If none of the provided cursors can be supported,
// the default cursor will be the ultimate fallback.
type Cursor struct {
	gobject.Object
}

func CursorNewFromInternalPtr(ptr uintptr) *Cursor {
	cls := &Cursor{}
	cls.Ptr = ptr
	return cls
}

var xNewFromNameCursor func(string, uintptr) uintptr

// Creates a new cursor by looking up @name in the current cursor
// theme.
//
// A recommended set of cursor names that will work across different
// platforms can be found in the CSS specification:
//
// | | | | |
// | --- | --- | ---- | --- |
// | "none" | ![](default_cursor.png) "default" | ![](help_cursor.png) "help" | ![](pointer_cursor.png) "pointer" |
// | ![](context_menu_cursor.png) "context-menu" | ![](progress_cursor.png) "progress" | ![](wait_cursor.png) "wait" | ![](cell_cursor.png) "cell" |
// | ![](crosshair_cursor.png) "crosshair" | ![](text_cursor.png) "text" | ![](vertical_text_cursor.png) "vertical-text" | ![](alias_cursor.png) "alias" |
// | ![](copy_cursor.png) "copy" | ![](no_drop_cursor.png) "no-drop" | ![](move_cursor.png) "move" | ![](not_allowed_cursor.png) "not-allowed" |
// | ![](grab_cursor.png) "grab" | ![](grabbing_cursor.png) "grabbing" | ![](all_scroll_cursor.png) "all-scroll" | ![](col_resize_cursor.png) "col-resize" |
// | ![](row_resize_cursor.png) "row-resize" | ![](n_resize_cursor.png) "n-resize" | ![](e_resize_cursor.png) "e-resize" | ![](s_resize_cursor.png) "s-resize" |
// | ![](w_resize_cursor.png) "w-resize" | ![](ne_resize_cursor.png) "ne-resize" | ![](nw_resize_cursor.png) "nw-resize" | ![](sw_resize_cursor.png) "sw-resize" |
// | ![](se_resize_cursor.png) "se-resize" | ![](ew_resize_cursor.png) "ew-resize" | ![](ns_resize_cursor.png) "ns-resize" | ![](nesw_resize_cursor.png) "nesw-resize" |
// | ![](nwse_resize_cursor.png) "nwse-resize" | ![](zoom_in_cursor.png) "zoom-in" | ![](zoom_out_cursor.png) "zoom-out" | |
func NewFromNameCursor(NameVar string, FallbackVar *Cursor) *Cursor {
	NewFromNameCursorPtr := xNewFromNameCursor(NameVar, FallbackVar.GoPointer())
	if NewFromNameCursorPtr == 0 {
		return nil
	}

	NewFromNameCursorCls := &Cursor{}
	NewFromNameCursorCls.Ptr = NewFromNameCursorPtr
	return NewFromNameCursorCls
}

var xNewFromTextureCursor func(uintptr, int32, int32, uintptr) uintptr

// Creates a new cursor from a `GdkTexture`.
func NewFromTextureCursor(TextureVar *Texture, HotspotXVar int32, HotspotYVar int32, FallbackVar *Cursor) *Cursor {
	NewFromTextureCursorPtr := xNewFromTextureCursor(TextureVar.GoPointer(), HotspotXVar, HotspotYVar, FallbackVar.GoPointer())
	if NewFromTextureCursorPtr == 0 {
		return nil
	}

	NewFromTextureCursorCls := &Cursor{}
	NewFromTextureCursorCls.Ptr = NewFromTextureCursorPtr
	return NewFromTextureCursorCls
}

var xCursorGetFallback func(uintptr) uintptr

// Returns the fallback for this @cursor.
//
// The fallback will be used if this cursor is not available on a given
// `GdkDisplay`. For named cursors, this can happen when using nonstandard
// names or when using an incomplete cursor theme. For textured cursors,
// this can happen when the texture is too large or when the `GdkDisplay`
// it is used on does not support textured cursors.
func (x *Cursor) GetFallback() *Cursor {

	GetFallbackPtr := xCursorGetFallback(x.GoPointer())
	if GetFallbackPtr == 0 {
		return nil
	}

	gobject.IncreaseRef(GetFallbackPtr)

	GetFallbackCls := &Cursor{}
	GetFallbackCls.Ptr = GetFallbackPtr
	return GetFallbackCls

}

var xCursorGetHotspotX func(uintptr) int32

// Returns the horizontal offset of the hotspot.
//
// The hotspot indicates the pixel that will be directly above the cursor.
//
// Note that named cursors may have a nonzero hotspot, but this function
// will only return the hotspot position for cursors created with
// [ctor@Gdk.Cursor.new_from_texture].
func (x *Cursor) GetHotspotX() int32 {

	return xCursorGetHotspotX(x.GoPointer())

}

var xCursorGetHotspotY func(uintptr) int32

// Returns the vertical offset of the hotspot.
//
// The hotspot indicates the pixel that will be directly above the cursor.
//
// Note that named cursors may have a nonzero hotspot, but this function
// will only return the hotspot position for cursors created with
// [ctor@Gdk.Cursor.new_from_texture].
func (x *Cursor) GetHotspotY() int32 {

	return xCursorGetHotspotY(x.GoPointer())

}

var xCursorGetName func(uintptr) string

// Returns the name of the cursor.
//
// If the cursor is not a named cursor, %NULL will be returned.
func (x *Cursor) GetName() string {

	return xCursorGetName(x.GoPointer())

}

var xCursorGetTexture func(uintptr) uintptr

// Returns the texture for the cursor.
//
// If the cursor is a named cursor, %NULL will be returned.
func (x *Cursor) GetTexture() *Texture {

	GetTexturePtr := xCursorGetTexture(x.GoPointer())
	if GetTexturePtr == 0 {
		return nil
	}

	gobject.IncreaseRef(GetTexturePtr)

	GetTextureCls := &Texture{}
	GetTextureCls.Ptr = GetTexturePtr
	return GetTextureCls

}

func (c *Cursor) GoPointer() uintptr {
	return c.Ptr
}

func (c *Cursor) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GDK"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xNewFromNameCursor, lib, "gdk_cursor_new_from_name")
	core.PuregoSafeRegister(&xNewFromTextureCursor, lib, "gdk_cursor_new_from_texture")

	core.PuregoSafeRegister(&xCursorGetFallback, lib, "gdk_cursor_get_fallback")
	core.PuregoSafeRegister(&xCursorGetHotspotX, lib, "gdk_cursor_get_hotspot_x")
	core.PuregoSafeRegister(&xCursorGetHotspotY, lib, "gdk_cursor_get_hotspot_y")
	core.PuregoSafeRegister(&xCursorGetName, lib, "gdk_cursor_get_name")
	core.PuregoSafeRegister(&xCursorGetTexture, lib, "gdk_cursor_get_texture")

}
