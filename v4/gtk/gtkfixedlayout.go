// Package gtk was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gtk

import (
	"unsafe"

	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/gsk"
)

type FixedLayoutChildClass struct {
	ParentClass uintptr
}

func (x *FixedLayoutChildClass) GoPointer() uintptr {
	return uintptr(unsafe.Pointer(x))
}

type FixedLayoutClass struct {
	ParentClass uintptr
}

func (x *FixedLayoutClass) GoPointer() uintptr {
	return uintptr(unsafe.Pointer(x))
}

// `GtkFixedLayout` is a layout manager which can place child widgets
// at fixed positions.
//
// Most applications should never use this layout manager; fixed positioning
// and sizing requires constant recalculations on where children need to be
// positioned and sized. Other layout managers perform this kind of work
// internally so that application developers don't need to do it. Specifically,
// widgets positioned in a fixed layout manager will need to take into account:
//
// - Themes, which may change widget sizes.
//
//   - Fonts other than the one you used to write the app will of course
//     change the size of widgets containing text; keep in mind that
//     users may use a larger font because of difficulty reading the
//     default, or they may be using a different OS that provides different
//     fonts.
//
//   - Translation of text into other languages changes its size. Also,
//     display of non-English text will use a different font in many
//     cases.
//
// In addition, `GtkFixedLayout` does not pay attention to text direction and
// thus may produce unwanted results if your app is run under right-to-left
// languages such as Hebrew or Arabic. That is: normally GTK will order
// containers appropriately depending on the text direction, e.g. to put labels
// to the right of the thing they label when using an RTL language;
// `GtkFixedLayout` won't be able to do that for you.
//
// Finally, fixed positioning makes it kind of annoying to add/remove UI
// elements, since you have to reposition all the other  elements. This is a
// long-term maintenance problem for your application.
type FixedLayout struct {
	LayoutManager
}

func FixedLayoutNewFromInternalPtr(ptr uintptr) *FixedLayout {
	cls := &FixedLayout{}
	cls.Ptr = ptr
	return cls
}

var xNewFixedLayout func() uintptr

// Creates a new `GtkFixedLayout`.
func NewFixedLayout() *FixedLayout {
	var cls *FixedLayout

	cret := xNewFixedLayout()

	if cret == 0 {
		return nil
	}
	cls = &FixedLayout{}
	cls.Ptr = cret
	return cls
}

func (c *FixedLayout) GoPointer() uintptr {
	return c.Ptr
}

func (c *FixedLayout) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// `GtkLayoutChild` subclass for children in a `GtkFixedLayout`.
type FixedLayoutChild struct {
	LayoutChild
}

func FixedLayoutChildNewFromInternalPtr(ptr uintptr) *FixedLayoutChild {
	cls := &FixedLayoutChild{}
	cls.Ptr = ptr
	return cls
}

var xFixedLayoutChildGetTransform func(uintptr) *gsk.Transform

// Retrieves the transformation of the child.
func (x *FixedLayoutChild) GetTransform() *gsk.Transform {

	cret := xFixedLayoutChildGetTransform(x.GoPointer())
	return cret
}

var xFixedLayoutChildSetTransform func(uintptr, *gsk.Transform)

// Sets the transformation of the child of a `GtkFixedLayout`.
func (x *FixedLayoutChild) SetTransform(TransformVar *gsk.Transform) {

	xFixedLayoutChildSetTransform(x.GoPointer(), TransformVar)

}

func (c *FixedLayoutChild) GoPointer() uintptr {
	return c.Ptr
}

func (c *FixedLayoutChild) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GTK"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xNewFixedLayout, lib, "gtk_fixed_layout_new")

	core.PuregoSafeRegister(&xFixedLayoutChildGetTransform, lib, "gtk_fixed_layout_child_get_transform")
	core.PuregoSafeRegister(&xFixedLayoutChildSetTransform, lib, "gtk_fixed_layout_child_set_transform")

}
