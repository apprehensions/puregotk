// Package gdk was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gdk

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/cairo"
	"github.com/jwijenbergh/puregotk/v4/gdkpixbuf"
)

var xPixbufGetFromSurface func(*cairo.Surface, int, int, int, int) uintptr

// Transfers image data from a `cairo_surface_t` and converts it
// to a `GdkPixbuf`.
//
// This allows you to efficiently read individual pixels from cairo surfaces.
//
// This function will create an RGB pixbuf with 8 bits per channel.
// The pixbuf will contain an alpha channel if the @surface contains one.
func PixbufGetFromSurface(SurfaceVar *cairo.Surface, SrcXVar int, SrcYVar int, WidthVar int, HeightVar int) *gdkpixbuf.Pixbuf {
	var cls *gdkpixbuf.Pixbuf

	cret := xPixbufGetFromSurface(SurfaceVar, SrcXVar, SrcYVar, WidthVar, HeightVar)

	if cret == 0 {
		return cls
	}
	cls = &gdkpixbuf.Pixbuf{}
	cls.Ptr = cret
	return cls
}

var xPixbufGetFromTexture func(uintptr) uintptr

// Creates a new pixbuf from @texture.
//
// This should generally not be used in newly written code as later
// stages will almost certainly convert the pixbuf back into a texture
// to draw it on screen.
func PixbufGetFromTexture(TextureVar *Texture) *gdkpixbuf.Pixbuf {
	var cls *gdkpixbuf.Pixbuf

	cret := xPixbufGetFromTexture(TextureVar.GoPointer())

	if cret == 0 {
		return cls
	}
	cls = &gdkpixbuf.Pixbuf{}
	cls.Ptr = cret
	return cls
}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GDK"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	core.PuregoSafeRegister(&xPixbufGetFromSurface, lib, "gdk_pixbuf_get_from_surface")
	core.PuregoSafeRegister(&xPixbufGetFromTexture, lib, "gdk_pixbuf_get_from_texture")

}
