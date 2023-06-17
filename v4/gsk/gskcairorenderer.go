// Package gsk was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gsk

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
)

type CairoRendererClass struct {
}

// A GSK renderer that is using cairo.
//
// Since it is using cairo, this renderer cannot support
// 3D transformations.
type CairoRenderer struct {
	Renderer
}

func CairoRendererNewFromInternalPtr(ptr uintptr) *CairoRenderer {
	cls := &CairoRenderer{}
	cls.Ptr = ptr
	return cls
}

var xNewCairoRenderer func() uintptr

// Creates a new Cairo renderer.
//
// The Cairo renderer is the fallback renderer drawing in ways similar
// to how GTK 3 drew its content. Its primary use is as comparison tool.
//
// The Cairo renderer is incomplete. It cannot render 3D transformed
// content and will instead render an error marker. Its usage should be
// avoided.
func NewCairoRenderer() *Renderer {
	NewCairoRendererPtr := xNewCairoRenderer()
	if NewCairoRendererPtr == 0 {
		return nil
	}

	NewCairoRendererCls := &Renderer{}
	NewCairoRendererCls.Ptr = NewCairoRendererPtr
	return NewCairoRendererCls
}

func (c *CairoRenderer) GoPointer() uintptr {
	return c.Ptr
}

func (c *CairoRenderer) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GSK"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xNewCairoRenderer, lib, "gsk_cairo_renderer_new")

}
