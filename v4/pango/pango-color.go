// Package pango was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package pango

import (
	"unsafe"

	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
)

// The `PangoColor` structure is used to
// represent a color in an uncalibrated RGB color-space.
type Color struct {
	Red uint16

	Green uint16

	Blue uint16
}

func (x *Color) GoPointer() uintptr {
	return uintptr(unsafe.Pointer(x))
}

var xColorCopy func(uintptr) *Color

// Creates a copy of @src.
//
// The copy should be freed with [method@Pango.Color.free].
// Primarily used by language bindings, not that useful
// otherwise (since colors can just be copied by assignment
// in C).
func (x *Color) Copy() *Color {

	cret := xColorCopy(x.GoPointer())
	return cret
}

var xColorFree func(uintptr)

// Frees a color allocated by [method@Pango.Color.copy].
func (x *Color) Free() {

	xColorFree(x.GoPointer())

}

var xColorParse func(uintptr, string) bool

// Fill in the fields of a color from a string specification.
//
// The string can either one of a large set of standard names.
// (Taken from the CSS Color [specification](https://www.w3.org/TR/css-color-4/#named-colors),
// or it can be a value in the form `#rgb`, `#rrggbb`,
// `#rrrgggbbb` or `#rrrrggggbbbb`, where `r`, `g` and `b`
// are hex digits of the red, green, and blue components
// of the color, respectively. (White in the four forms is
// `#fff`, `#ffffff`, `#fffffffff` and `#ffffffffffff`.)
func (x *Color) Parse(SpecVar string) bool {

	cret := xColorParse(x.GoPointer(), SpecVar)
	return cret
}

var xColorParseWithAlpha func(uintptr, uint16, string) bool

// Fill in the fields of a color from a string specification.
//
// The string can either one of a large set of standard names.
// (Taken from the CSS Color [specification](https://www.w3.org/TR/css-color-4/#named-colors),
// or it can be a hexadecimal value in the form `#rgb`,
// `#rrggbb`, `#rrrgggbbb` or `#rrrrggggbbbb` where `r`, `g`
// and `b` are hex digits of the red, green, and blue components
// of the color, respectively. (White in the four forms is
// `#fff`, `#ffffff`, `#fffffffff` and `#ffffffffffff`.)
//
// Additionally, parse strings of the form `#rgba`, `#rrggbbaa`,
// `#rrrrggggbbbbaaaa`, if @alpha is not %NULL, and set @alpha
// to the value specified by the hex digits for `a`. If no alpha
// component is found in @spec, @alpha is set to 0xffff (for a
// solid color).
func (x *Color) ParseWithAlpha(AlphaVar uint16, SpecVar string) bool {

	cret := xColorParseWithAlpha(x.GoPointer(), AlphaVar, SpecVar)
	return cret
}

var xColorToString func(uintptr) string

// Returns a textual specification of @color.
//
// The string is in the hexadecimal form `#rrrrggggbbbb`,
// where `r`, `g` and `b` are hex digits representing the
// red, green, and blue components respectively.
func (x *Color) ToString() string {

	cret := xColorToString(x.GoPointer())
	return cret
}

func init() {
	lib, err := purego.Dlopen(core.GetPath("PANGO"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xColorCopy, lib, "pango_color_copy")
	core.PuregoSafeRegister(&xColorFree, lib, "pango_color_free")
	core.PuregoSafeRegister(&xColorParse, lib, "pango_color_parse")
	core.PuregoSafeRegister(&xColorParseWithAlpha, lib, "pango_color_parse_with_alpha")
	core.PuregoSafeRegister(&xColorToString, lib, "pango_color_to_string")

}
