// Package adw was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package adw

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/gtk"
)

type ClampLayoutClass struct {
	ParentClass uintptr
}

// A layout manager constraining its children to a given size.
//
// &lt;picture&gt;
//
//	&lt;source srcset="clamp-wide-dark.png" media="(prefers-color-scheme: dark)"&gt;
//	&lt;img src="clamp-wide.png" alt="clamp-wide"&gt;
//
// &lt;/picture&gt;
// &lt;picture&gt;
//
//	&lt;source srcset="clamp-narrow-dark.png" media="(prefers-color-scheme: dark)"&gt;
//	&lt;img src="clamp-narrow.png" alt="clamp-narrow"&gt;
//
// &lt;/picture&gt;
//
// `AdwClampLayout` constraints the size of the widgets it contains to a given
// maximum size. It will constrain the width if it is horizontal, or the height
// if it is vertical. The expansion of the children from their minimum to their
// maximum size is eased out for a smooth transition.
//
// If a child requires more than the requested maximum size, it will be
// allocated the minimum size it can fit in instead.
//
// Each child will get the style  classes .large when it reached its maximum
// size, .small when it's allocated the full size, .medium in-between, or none
// if it hasn't been allocated yet.
type ClampLayout struct {
	gtk.LayoutManager
}

func ClampLayoutNewFromInternalPtr(ptr uintptr) *ClampLayout {
	cls := &ClampLayout{}
	cls.Ptr = ptr
	return cls
}

var xNewClampLayout func() uintptr

// Creates a new `AdwClampLayout`.
func NewClampLayout() *gtk.LayoutManager {
	NewClampLayoutPtr := xNewClampLayout()
	if NewClampLayoutPtr == 0 {
		return nil
	}

	NewClampLayoutCls := &gtk.LayoutManager{}
	NewClampLayoutCls.Ptr = NewClampLayoutPtr
	return NewClampLayoutCls
}

var xClampLayoutGetMaximumSize func(uintptr) int32

// Gets the maximum size allocated to the children.
func (x *ClampLayout) GetMaximumSize() int32 {

	return xClampLayoutGetMaximumSize(x.GoPointer())

}

var xClampLayoutGetTighteningThreshold func(uintptr) int32

// Gets the size above which the children are clamped.
func (x *ClampLayout) GetTighteningThreshold() int32 {

	return xClampLayoutGetTighteningThreshold(x.GoPointer())

}

var xClampLayoutSetMaximumSize func(uintptr, int32)

// Sets the maximum size allocated to the children.
//
// It is the width if the layout is horizontal, or the height if it is vertical.
func (x *ClampLayout) SetMaximumSize(MaximumSizeVar int32) {

	xClampLayoutSetMaximumSize(x.GoPointer(), MaximumSizeVar)

}

var xClampLayoutSetTighteningThreshold func(uintptr, int32)

// Sets the size above which the children are clamped.
//
// Starting from this size, the layout will tighten its grip on the children,
// slowly allocating less and less of the available size up to the maximum
// allocated size. Below that threshold and below the maximum size, the children
// will be allocated all the available size.
//
// If the threshold is greater than the maximum size to allocate to the
// children, they will be allocated the whole size up to the maximum. If the
// threshold is lower than the minimum size to allocate to the children, that
// size will be used as the tightening threshold.
//
// Effectively, tightening the grip on a child before it reaches its maximum
// size makes transitions to and from the maximum size smoother when resizing.
func (x *ClampLayout) SetTighteningThreshold(TighteningThresholdVar int32) {

	xClampLayoutSetTighteningThreshold(x.GoPointer(), TighteningThresholdVar)

}

func (c *ClampLayout) GoPointer() uintptr {
	return c.Ptr
}

func (c *ClampLayout) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// Retrieves the orientation of the @orientable.
func (x *ClampLayout) GetOrientation() gtk.Orientation {

	return gtk.XGtkOrientableGetOrientation(x.GoPointer())

}

// Sets the orientation of the @orientable.
func (x *ClampLayout) SetOrientation(OrientationVar gtk.Orientation) {

	gtk.XGtkOrientableSetOrientation(x.GoPointer(), OrientationVar)

}

func init() {
	lib, err := purego.Dlopen(core.GetPath("ADW"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xNewClampLayout, lib, "adw_clamp_layout_new")

	core.PuregoSafeRegister(&xClampLayoutGetMaximumSize, lib, "adw_clamp_layout_get_maximum_size")
	core.PuregoSafeRegister(&xClampLayoutGetTighteningThreshold, lib, "adw_clamp_layout_get_tightening_threshold")
	core.PuregoSafeRegister(&xClampLayoutSetMaximumSize, lib, "adw_clamp_layout_set_maximum_size")
	core.PuregoSafeRegister(&xClampLayoutSetTighteningThreshold, lib, "adw_clamp_layout_set_tightening_threshold")

}
