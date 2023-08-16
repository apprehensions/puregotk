// Package pango was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package pango

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/gobject"
)

// Class structure for `PangoRenderer`.
//
// The following vfuncs take user space coordinates in Pango units
// and have default implementations:
// - draw_glyphs
// - draw_rectangle
// - draw_error_underline
// - draw_shape
// - draw_glyph_item
//
// The default draw_shape implementation draws nothing.
//
// The following vfuncs take device space coordinates as doubles
// and must be implemented:
// - draw_trapezoid
// - draw_glyph
type RendererClass struct {
	ParentClass uintptr
}

type RendererPrivate struct {
}

// `PangoRenderPart` defines different items to render for such
// purposes as setting colors.
type RenderPart int

const (

	// the text itself
	RenderPartForegroundValue RenderPart = 0
	// the area behind the text
	RenderPartBackgroundValue RenderPart = 1
	// underlines
	RenderPartUnderlineValue RenderPart = 2
	// strikethrough lines
	RenderPartStrikethroughValue RenderPart = 3
	// overlines
	RenderPartOverlineValue RenderPart = 4
)

// `PangoRenderer` is a base class for objects that can render text
// provided as `PangoGlyphString` or `PangoLayout`.
//
// By subclassing `PangoRenderer` and overriding operations such as
// @draw_glyphs and @draw_rectangle, renderers for particular font
// backends and destinations can be created.
type Renderer struct {
	gobject.Object
}

func RendererNewFromInternalPtr(ptr uintptr) *Renderer {
	cls := &Renderer{}
	cls.Ptr = ptr
	return cls
}

var xRendererActivate func(uintptr)

// Does initial setup before rendering operations on @renderer.
//
// [method@Pango.Renderer.deactivate] should be called when done drawing.
// Calls such as [method@Pango.Renderer.draw_layout] automatically
// activate the layout before drawing on it.
//
// Calls to [method@Pango.Renderer.activate] and
// [method@Pango.Renderer.deactivate] can be nested and the
// renderer will only be initialized and deinitialized once.
func (x *Renderer) Activate() {

	xRendererActivate(x.GoPointer())

}

var xRendererDeactivate func(uintptr)

// Cleans up after rendering operations on @renderer.
//
// See docs for [method@Pango.Renderer.activate].
func (x *Renderer) Deactivate() {

	xRendererDeactivate(x.GoPointer())

}

var xRendererDrawErrorUnderline func(uintptr, int, int, int, int)

// Draw a squiggly line that approximately covers the given rectangle
// in the style of an underline used to indicate a spelling error.
//
// The width of the underline is rounded to an integer number
// of up/down segments and the resulting rectangle is centered
// in the original rectangle.
//
// This should be called while @renderer is already active.
// Use [method@Pango.Renderer.activate] to activate a renderer.
func (x *Renderer) DrawErrorUnderline(XVar int, YVar int, WidthVar int, HeightVar int) {

	xRendererDrawErrorUnderline(x.GoPointer(), XVar, YVar, WidthVar, HeightVar)

}

var xRendererDrawGlyph func(uintptr, uintptr, Glyph, float64, float64)

// Draws a single glyph with coordinates in device space.
func (x *Renderer) DrawGlyph(FontVar *Font, GlyphVar Glyph, XVar float64, YVar float64) {

	xRendererDrawGlyph(x.GoPointer(), FontVar.GoPointer(), GlyphVar, XVar, YVar)

}

var xRendererDrawGlyphItem func(uintptr, string, *GlyphItem, int, int)

// Draws the glyphs in @glyph_item with the specified `PangoRenderer`,
// embedding the text associated with the glyphs in the output if the
// output format supports it.
//
// This is useful for rendering text in PDF.
//
// Note that this method does not handle attributes in @glyph_item.
// If you want colors, shapes and lines handled automatically according
// to those attributes, you need to use pango_renderer_draw_layout_line()
// or pango_renderer_draw_layout().
//
// Note that @text is the start of the text for layout, which is then
// indexed by `glyph_item-&gt;item-&gt;offset`.
//
// If @text is %NULL, this simply calls [method@Pango.Renderer.draw_glyphs].
//
// The default implementation of this method simply falls back to
// [method@Pango.Renderer.draw_glyphs].
func (x *Renderer) DrawGlyphItem(TextVar string, GlyphItemVar *GlyphItem, XVar int, YVar int) {

	xRendererDrawGlyphItem(x.GoPointer(), TextVar, GlyphItemVar, XVar, YVar)

}

var xRendererDrawGlyphs func(uintptr, uintptr, *GlyphString, int, int)

// Draws the glyphs in @glyphs with the specified `PangoRenderer`.
func (x *Renderer) DrawGlyphs(FontVar *Font, GlyphsVar *GlyphString, XVar int, YVar int) {

	xRendererDrawGlyphs(x.GoPointer(), FontVar.GoPointer(), GlyphsVar, XVar, YVar)

}

var xRendererDrawLayout func(uintptr, uintptr, int, int)

// Draws @layout with the specified `PangoRenderer`.
//
// This is equivalent to drawing the lines of the layout, at their
// respective positions relative to @x, @y.
func (x *Renderer) DrawLayout(LayoutVar *Layout, XVar int, YVar int) {

	xRendererDrawLayout(x.GoPointer(), LayoutVar.GoPointer(), XVar, YVar)

}

var xRendererDrawLayoutLine func(uintptr, *LayoutLine, int, int)

// Draws @line with the specified `PangoRenderer`.
//
// This draws the glyph items that make up the line, as well as
// shapes, backgrounds and lines that are specified by the attributes
// of those items.
func (x *Renderer) DrawLayoutLine(LineVar *LayoutLine, XVar int, YVar int) {

	xRendererDrawLayoutLine(x.GoPointer(), LineVar, XVar, YVar)

}

var xRendererDrawRectangle func(uintptr, RenderPart, int, int, int, int)

// Draws an axis-aligned rectangle in user space coordinates with the
// specified `PangoRenderer`.
//
// This should be called while @renderer is already active.
// Use [method@Pango.Renderer.activate] to activate a renderer.
func (x *Renderer) DrawRectangle(PartVar RenderPart, XVar int, YVar int, WidthVar int, HeightVar int) {

	xRendererDrawRectangle(x.GoPointer(), PartVar, XVar, YVar, WidthVar, HeightVar)

}

var xRendererDrawTrapezoid func(uintptr, RenderPart, float64, float64, float64, float64, float64, float64)

// Draws a trapezoid with the parallel sides aligned with the X axis
// using the given `PangoRenderer`; coordinates are in device space.
func (x *Renderer) DrawTrapezoid(PartVar RenderPart, Y1Var float64, X11Var float64, X21Var float64, Y2Var float64, X12Var float64, X22Var float64) {

	xRendererDrawTrapezoid(x.GoPointer(), PartVar, Y1Var, X11Var, X21Var, Y2Var, X12Var, X22Var)

}

var xRendererGetAlpha func(uintptr, RenderPart) uint16

// Gets the current alpha for the specified part.
func (x *Renderer) GetAlpha(PartVar RenderPart) uint16 {

	cret := xRendererGetAlpha(x.GoPointer(), PartVar)
	return cret
}

var xRendererGetColor func(uintptr, RenderPart) *Color

// Gets the current rendering color for the specified part.
func (x *Renderer) GetColor(PartVar RenderPart) *Color {

	cret := xRendererGetColor(x.GoPointer(), PartVar)
	return cret
}

var xRendererGetLayout func(uintptr) uintptr

// Gets the layout currently being rendered using @renderer.
//
// Calling this function only makes sense from inside a subclass's
// methods, like in its draw_shape vfunc, for example.
//
// The returned layout should not be modified while still being
// rendered.
func (x *Renderer) GetLayout() *Layout {
	var cls *Layout

	cret := xRendererGetLayout(x.GoPointer())

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &Layout{}
	cls.Ptr = cret
	return cls
}

var xRendererGetLayoutLine func(uintptr) *LayoutLine

// Gets the layout line currently being rendered using @renderer.
//
// Calling this function only makes sense from inside a subclass's
// methods, like in its draw_shape vfunc, for example.
//
// The returned layout line should not be modified while still being
// rendered.
func (x *Renderer) GetLayoutLine() *LayoutLine {

	cret := xRendererGetLayoutLine(x.GoPointer())
	return cret
}

var xRendererGetMatrix func(uintptr) *Matrix

// Gets the transformation matrix that will be applied when
// rendering.
//
// See [method@Pango.Renderer.set_matrix].
func (x *Renderer) GetMatrix() *Matrix {

	cret := xRendererGetMatrix(x.GoPointer())
	return cret
}

var xRendererPartChanged func(uintptr, RenderPart)

// Informs Pango that the way that the rendering is done
// for @part has changed.
//
// This should be called if the rendering changes in a way that would
// prevent multiple pieces being joined together into one drawing call.
// For instance, if a subclass of `PangoRenderer` was to add a stipple
// option for drawing underlines, it needs to call
//
// ```
// pango_renderer_part_changed (render, PANGO_RENDER_PART_UNDERLINE);
// ```
//
// When the stipple changes or underlines with different stipples
// might be joined together. Pango automatically calls this for
// changes to colors. (See [method@Pango.Renderer.set_color])
func (x *Renderer) PartChanged(PartVar RenderPart) {

	xRendererPartChanged(x.GoPointer(), PartVar)

}

var xRendererSetAlpha func(uintptr, RenderPart, uint16)

// Sets the alpha for part of the rendering.
//
// Note that the alpha may only be used if a color is
// specified for @part as well.
func (x *Renderer) SetAlpha(PartVar RenderPart, AlphaVar uint16) {

	xRendererSetAlpha(x.GoPointer(), PartVar, AlphaVar)

}

var xRendererSetColor func(uintptr, RenderPart, *Color)

// Sets the color for part of the rendering.
//
// Also see [method@Pango.Renderer.set_alpha].
func (x *Renderer) SetColor(PartVar RenderPart, ColorVar *Color) {

	xRendererSetColor(x.GoPointer(), PartVar, ColorVar)

}

var xRendererSetMatrix func(uintptr, *Matrix)

// Sets the transformation matrix that will be applied when rendering.
func (x *Renderer) SetMatrix(MatrixVar *Matrix) {

	xRendererSetMatrix(x.GoPointer(), MatrixVar)

}

func (c *Renderer) GoPointer() uintptr {
	return c.Ptr
}

func (c *Renderer) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

func init() {
	lib, err := purego.Dlopen(core.GetPath("PANGO"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xRendererActivate, lib, "pango_renderer_activate")
	core.PuregoSafeRegister(&xRendererDeactivate, lib, "pango_renderer_deactivate")
	core.PuregoSafeRegister(&xRendererDrawErrorUnderline, lib, "pango_renderer_draw_error_underline")
	core.PuregoSafeRegister(&xRendererDrawGlyph, lib, "pango_renderer_draw_glyph")
	core.PuregoSafeRegister(&xRendererDrawGlyphItem, lib, "pango_renderer_draw_glyph_item")
	core.PuregoSafeRegister(&xRendererDrawGlyphs, lib, "pango_renderer_draw_glyphs")
	core.PuregoSafeRegister(&xRendererDrawLayout, lib, "pango_renderer_draw_layout")
	core.PuregoSafeRegister(&xRendererDrawLayoutLine, lib, "pango_renderer_draw_layout_line")
	core.PuregoSafeRegister(&xRendererDrawRectangle, lib, "pango_renderer_draw_rectangle")
	core.PuregoSafeRegister(&xRendererDrawTrapezoid, lib, "pango_renderer_draw_trapezoid")
	core.PuregoSafeRegister(&xRendererGetAlpha, lib, "pango_renderer_get_alpha")
	core.PuregoSafeRegister(&xRendererGetColor, lib, "pango_renderer_get_color")
	core.PuregoSafeRegister(&xRendererGetLayout, lib, "pango_renderer_get_layout")
	core.PuregoSafeRegister(&xRendererGetLayoutLine, lib, "pango_renderer_get_layout_line")
	core.PuregoSafeRegister(&xRendererGetMatrix, lib, "pango_renderer_get_matrix")
	core.PuregoSafeRegister(&xRendererPartChanged, lib, "pango_renderer_part_changed")
	core.PuregoSafeRegister(&xRendererSetAlpha, lib, "pango_renderer_set_alpha")
	core.PuregoSafeRegister(&xRendererSetColor, lib, "pango_renderer_set_color")
	core.PuregoSafeRegister(&xRendererSetMatrix, lib, "pango_renderer_set_matrix")

}
