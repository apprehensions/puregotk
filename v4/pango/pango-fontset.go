// Package pango was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package pango

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/gobject"
)

// Callback used by pango_fontset_foreach() when enumerating
// fonts in a fontset.
type FontsetForeachFunc func(uintptr, uintptr, uintptr) bool

// The `PangoFontsetClass` structure holds the virtual functions for
// a particular `PangoFontset` implementation.
type FontsetClass struct {
	ParentClass uintptr
}

type FontsetSimpleClass struct {
}

// A `PangoFontset` represents a set of `PangoFont` to use when rendering text.
//
// A `PAngoFontset` is the result of resolving a `PangoFontDescription`
// against a particular `PangoContext`. It has operations for finding the
// component font for a particular Unicode character, and for finding a
// composite set of metrics for the entire fontset.
type Fontset struct {
	gobject.Object
}

func FontsetNewFromInternalPtr(ptr uintptr) *Fontset {
	cls := &Fontset{}
	cls.Ptr = ptr
	return cls
}

var xFontsetForeach func(uintptr, uintptr, uintptr)

// Iterates through all the fonts in a fontset, calling @func for
// each one.
//
// If @func returns %TRUE, that stops the iteration.
func (x *Fontset) Foreach(FuncVar FontsetForeachFunc, DataVar uintptr) {

	xFontsetForeach(x.GoPointer(), purego.NewCallback(FuncVar), DataVar)

}

var xFontsetGetFont func(uintptr, uint) uintptr

// Returns the font in the fontset that contains the best glyph for a
// Unicode character.
func (x *Fontset) GetFont(WcVar uint) *Font {
	var cls *Font

	cret := xFontsetGetFont(x.GoPointer(), WcVar)

	if cret == 0 {
		return cls
	}
	cls = &Font{}
	cls.Ptr = cret
	return cls
}

var xFontsetGetMetrics func(uintptr) *FontMetrics

// Get overall metric information for the fonts in the fontset.
func (x *Fontset) GetMetrics() *FontMetrics {

	cret := xFontsetGetMetrics(x.GoPointer())
	return cret
}

func (c *Fontset) GoPointer() uintptr {
	return c.Ptr
}

func (c *Fontset) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// `PangoFontsetSimple` is a implementation of the abstract
// `PangoFontset` base class as an array of fonts.
//
// When creating a `PangoFontsetSimple`, you have to provide
// the array of fonts that make up the fontset.
type FontsetSimple struct {
	Fontset
}

func FontsetSimpleNewFromInternalPtr(ptr uintptr) *FontsetSimple {
	cls := &FontsetSimple{}
	cls.Ptr = ptr
	return cls
}

var xNewFontsetSimple func(*Language) uintptr

// Creates a new `PangoFontsetSimple` for the given language.
func NewFontsetSimple(LanguageVar *Language) *FontsetSimple {
	var cls *FontsetSimple

	cret := xNewFontsetSimple(LanguageVar)

	if cret == 0 {
		return cls
	}
	cls = &FontsetSimple{}
	cls.Ptr = cret
	return cls
}

var xFontsetSimpleAppend func(uintptr, uintptr)

// Adds a font to the fontset.
//
// The fontset takes ownership of @font.
func (x *FontsetSimple) Append(FontVar *Font) {

	xFontsetSimpleAppend(x.GoPointer(), FontVar.GoPointer())

}

var xFontsetSimpleSize func(uintptr) int

// Returns the number of fonts in the fontset.
func (x *FontsetSimple) Size() int {

	cret := xFontsetSimpleSize(x.GoPointer())
	return cret
}

func (c *FontsetSimple) GoPointer() uintptr {
	return c.Ptr
}

func (c *FontsetSimple) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

func init() {
	lib, err := purego.Dlopen(core.GetPath("PANGO"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xFontsetForeach, lib, "pango_fontset_foreach")
	core.PuregoSafeRegister(&xFontsetGetFont, lib, "pango_fontset_get_font")
	core.PuregoSafeRegister(&xFontsetGetMetrics, lib, "pango_fontset_get_metrics")

	core.PuregoSafeRegister(&xNewFontsetSimple, lib, "pango_fontset_simple_new")

	core.PuregoSafeRegister(&xFontsetSimpleAppend, lib, "pango_fontset_simple_append")
	core.PuregoSafeRegister(&xFontsetSimpleSize, lib, "pango_fontset_simple_size")

}
