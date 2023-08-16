// Package gtk was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gtk

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/gobject"
)

type CenterLayoutClass struct {
	ParentClass uintptr
}

// `GtkCenterLayout` is a layout manager that manages up to three children.
//
// The start widget is allocated at the start of the layout (left in
// left-to-right locales and right in right-to-left ones), and the end
// widget at the end.
//
// The center widget is centered regarding the full width of the layout's.
type CenterLayout struct {
	LayoutManager
}

func CenterLayoutNewFromInternalPtr(ptr uintptr) *CenterLayout {
	cls := &CenterLayout{}
	cls.Ptr = ptr
	return cls
}

var xNewCenterLayout func() uintptr

// Creates a new `GtkCenterLayout`.
func NewCenterLayout() *LayoutManager {
	var cls *LayoutManager

	cret := xNewCenterLayout()

	if cret == 0 {
		return cls
	}
	cls = &LayoutManager{}
	cls.Ptr = cret
	return cls
}

var xCenterLayoutGetBaselinePosition func(uintptr) BaselinePosition

// Returns the baseline position of the layout.
func (x *CenterLayout) GetBaselinePosition() BaselinePosition {

	cret := xCenterLayoutGetBaselinePosition(x.GoPointer())
	return cret
}

var xCenterLayoutGetCenterWidget func(uintptr) uintptr

// Returns the center widget of the layout.
func (x *CenterLayout) GetCenterWidget() *Widget {
	var cls *Widget

	cret := xCenterLayoutGetCenterWidget(x.GoPointer())

	if cret == 0 {
		return cls
	}
	gobject.IncreaseRef(cret)
	cls = &Widget{}
	cls.Ptr = cret
	return cls
}

var xCenterLayoutGetEndWidget func(uintptr) uintptr

// Returns the end widget of the layout.
func (x *CenterLayout) GetEndWidget() *Widget {
	var cls *Widget

	cret := xCenterLayoutGetEndWidget(x.GoPointer())

	if cret == 0 {
		return cls
	}
	gobject.IncreaseRef(cret)
	cls = &Widget{}
	cls.Ptr = cret
	return cls
}

var xCenterLayoutGetOrientation func(uintptr) Orientation

// Gets the current orienration of the layout manager.
func (x *CenterLayout) GetOrientation() Orientation {

	cret := xCenterLayoutGetOrientation(x.GoPointer())
	return cret
}

var xCenterLayoutGetStartWidget func(uintptr) uintptr

// Returns the start widget fo the layout.
func (x *CenterLayout) GetStartWidget() *Widget {
	var cls *Widget

	cret := xCenterLayoutGetStartWidget(x.GoPointer())

	if cret == 0 {
		return cls
	}
	gobject.IncreaseRef(cret)
	cls = &Widget{}
	cls.Ptr = cret
	return cls
}

var xCenterLayoutSetBaselinePosition func(uintptr, BaselinePosition)

// Sets the new baseline position of @self
func (x *CenterLayout) SetBaselinePosition(BaselinePositionVar BaselinePosition) {

	xCenterLayoutSetBaselinePosition(x.GoPointer(), BaselinePositionVar)

}

var xCenterLayoutSetCenterWidget func(uintptr, uintptr)

// Sets the new center widget of @self.
//
// To remove the existing center widget, pass %NULL.
func (x *CenterLayout) SetCenterWidget(WidgetVar *Widget) {

	xCenterLayoutSetCenterWidget(x.GoPointer(), WidgetVar.GoPointer())

}

var xCenterLayoutSetEndWidget func(uintptr, uintptr)

// Sets the new end widget of @self.
//
// To remove the existing center widget, pass %NULL.
func (x *CenterLayout) SetEndWidget(WidgetVar *Widget) {

	xCenterLayoutSetEndWidget(x.GoPointer(), WidgetVar.GoPointer())

}

var xCenterLayoutSetOrientation func(uintptr, Orientation)

// Sets the orientation of @self.
func (x *CenterLayout) SetOrientation(OrientationVar Orientation) {

	xCenterLayoutSetOrientation(x.GoPointer(), OrientationVar)

}

var xCenterLayoutSetStartWidget func(uintptr, uintptr)

// Sets the new start widget of @self.
//
// To remove the existing start widget, pass %NULL.
func (x *CenterLayout) SetStartWidget(WidgetVar *Widget) {

	xCenterLayoutSetStartWidget(x.GoPointer(), WidgetVar.GoPointer())

}

func (c *CenterLayout) GoPointer() uintptr {
	return c.Ptr
}

func (c *CenterLayout) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GTK"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xNewCenterLayout, lib, "gtk_center_layout_new")

	core.PuregoSafeRegister(&xCenterLayoutGetBaselinePosition, lib, "gtk_center_layout_get_baseline_position")
	core.PuregoSafeRegister(&xCenterLayoutGetCenterWidget, lib, "gtk_center_layout_get_center_widget")
	core.PuregoSafeRegister(&xCenterLayoutGetEndWidget, lib, "gtk_center_layout_get_end_widget")
	core.PuregoSafeRegister(&xCenterLayoutGetOrientation, lib, "gtk_center_layout_get_orientation")
	core.PuregoSafeRegister(&xCenterLayoutGetStartWidget, lib, "gtk_center_layout_get_start_widget")
	core.PuregoSafeRegister(&xCenterLayoutSetBaselinePosition, lib, "gtk_center_layout_set_baseline_position")
	core.PuregoSafeRegister(&xCenterLayoutSetCenterWidget, lib, "gtk_center_layout_set_center_widget")
	core.PuregoSafeRegister(&xCenterLayoutSetEndWidget, lib, "gtk_center_layout_set_end_widget")
	core.PuregoSafeRegister(&xCenterLayoutSetOrientation, lib, "gtk_center_layout_set_orientation")
	core.PuregoSafeRegister(&xCenterLayoutSetStartWidget, lib, "gtk_center_layout_set_start_widget")

}
