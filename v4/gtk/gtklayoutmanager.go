// Package gtk was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gtk

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/gobject"
)

// The `GtkLayoutManagerClass` structure contains only private data, and
// should only be accessed through the provided API, or when subclassing
// `GtkLayoutManager`.
type LayoutManagerClass struct {
	ParentClass uintptr

	LayoutChildType []interface{}

	Padding uintptr
}

// Layout managers are delegate classes that handle the preferred size
// and the allocation of a widget.
//
// You typically subclass `GtkLayoutManager` if you want to implement a
// layout policy for the children of a widget, or if you want to determine
// the size of a widget depending on its contents.
//
// Each `GtkWidget` can only have a `GtkLayoutManager` instance associated
// to it at any given time; it is possible, though, to replace the layout
// manager instance using [method@Gtk.Widget.set_layout_manager].
//
// ## Layout properties
//
// A layout manager can expose properties for controlling the layout of
// each child, by creating an object type derived from [class@Gtk.LayoutChild]
// and installing the properties on it as normal `GObject` properties.
//
// Each `GtkLayoutChild` instance storing the layout properties for a
// specific child is created through the [method@Gtk.LayoutManager.get_layout_child]
// method; a `GtkLayoutManager` controls the creation of its `GtkLayoutChild`
// instances by overriding the GtkLayoutManagerClass.create_layout_child()
// virtual function. The typical implementation should look like:
//
// ```c
// static GtkLayoutChild *
// create_layout_child (GtkLayoutManager *manager,
//
//	GtkWidget        *container,
//	GtkWidget        *child)
//
//	{
//	  return g_object_new (your_layout_child_get_type (),
//	                       "layout-manager", manager,
//	                       "child-widget", child,
//	                       NULL);
//	}
//
// ```
//
// The [property@Gtk.LayoutChild:layout-manager] and
// [property@Gtk.LayoutChild:child-widget] properties
// on the newly created `GtkLayoutChild` instance are mandatory. The
// `GtkLayoutManager` will cache the newly created `GtkLayoutChild` instance
// until the widget is removed from its parent, or the parent removes the
// layout manager.
//
// Each `GtkLayoutManager` instance creating a `GtkLayoutChild` should use
// [method@Gtk.LayoutManager.get_layout_child] every time it needs to query
// the layout properties; each `GtkLayoutChild` instance should call
// [method@Gtk.LayoutManager.layout_changed] every time a property is
// updated, in order to queue a new size measuring and allocation.
type LayoutManager struct {
	gobject.Object
}

func LayoutManagerNewFromInternalPtr(ptr uintptr) *LayoutManager {
	cls := &LayoutManager{}
	cls.Ptr = ptr
	return cls
}

var xLayoutManagerAllocate func(uintptr, uintptr, int32, int32, int32)

// Assigns the given @width, @height, and @baseline to
// a @widget, and computes the position and sizes of the children of
// the @widget using the layout management policy of @manager.
func (x *LayoutManager) Allocate(WidgetVar *Widget, WidthVar int32, HeightVar int32, BaselineVar int32) {

	xLayoutManagerAllocate(x.GoPointer(), WidgetVar.GoPointer(), WidthVar, HeightVar, BaselineVar)

}

var xLayoutManagerGetLayoutChild func(uintptr, uintptr) uintptr

// Retrieves a `GtkLayoutChild` instance for the `GtkLayoutManager`,
// creating one if necessary.
//
// The @child widget must be a child of the widget using @manager.
//
// The `GtkLayoutChild` instance is owned by the `GtkLayoutManager`,
// and is guaranteed to exist as long as @child is a child of the
// `GtkWidget` using the given `GtkLayoutManager`.
func (x *LayoutManager) GetLayoutChild(ChildVar *Widget) *LayoutChild {

	GetLayoutChildPtr := xLayoutManagerGetLayoutChild(x.GoPointer(), ChildVar.GoPointer())
	if GetLayoutChildPtr == 0 {
		return nil
	}

	gobject.IncreaseRef(GetLayoutChildPtr)

	GetLayoutChildCls := &LayoutChild{}
	GetLayoutChildCls.Ptr = GetLayoutChildPtr
	return GetLayoutChildCls

}

var xLayoutManagerGetRequestMode func(uintptr) SizeRequestMode

// Retrieves the request mode of @manager.
func (x *LayoutManager) GetRequestMode() SizeRequestMode {

	return xLayoutManagerGetRequestMode(x.GoPointer())

}

var xLayoutManagerGetWidget func(uintptr) uintptr

// Retrieves the `GtkWidget` using the given `GtkLayoutManager`.
func (x *LayoutManager) GetWidget() *Widget {

	GetWidgetPtr := xLayoutManagerGetWidget(x.GoPointer())
	if GetWidgetPtr == 0 {
		return nil
	}

	gobject.IncreaseRef(GetWidgetPtr)

	GetWidgetCls := &Widget{}
	GetWidgetCls.Ptr = GetWidgetPtr
	return GetWidgetCls

}

var xLayoutManagerLayoutChanged func(uintptr)

// Queues a resize on the `GtkWidget` using @manager, if any.
//
// This function should be called by subclasses of `GtkLayoutManager`
// in response to changes to their layout management policies.
func (x *LayoutManager) LayoutChanged() {

	xLayoutManagerLayoutChanged(x.GoPointer())

}

var xLayoutManagerMeasure func(uintptr, uintptr, Orientation, int32, int32, int32, int32, int32)

// Measures the size of the @widget using @manager, for the
// given @orientation and size.
//
// See the [class@Gtk.Widget] documentation on layout management for
// more details.
func (x *LayoutManager) Measure(WidgetVar *Widget, OrientationVar Orientation, ForSizeVar int32, MinimumVar int32, NaturalVar int32, MinimumBaselineVar int32, NaturalBaselineVar int32) {

	xLayoutManagerMeasure(x.GoPointer(), WidgetVar.GoPointer(), OrientationVar, ForSizeVar, MinimumVar, NaturalVar, MinimumBaselineVar, NaturalBaselineVar)

}

func (c *LayoutManager) GoPointer() uintptr {
	return c.Ptr
}

func (c *LayoutManager) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GTK"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xLayoutManagerAllocate, lib, "gtk_layout_manager_allocate")
	core.PuregoSafeRegister(&xLayoutManagerGetLayoutChild, lib, "gtk_layout_manager_get_layout_child")
	core.PuregoSafeRegister(&xLayoutManagerGetRequestMode, lib, "gtk_layout_manager_get_request_mode")
	core.PuregoSafeRegister(&xLayoutManagerGetWidget, lib, "gtk_layout_manager_get_widget")
	core.PuregoSafeRegister(&xLayoutManagerLayoutChanged, lib, "gtk_layout_manager_layout_changed")
	core.PuregoSafeRegister(&xLayoutManagerMeasure, lib, "gtk_layout_manager_measure")

}
