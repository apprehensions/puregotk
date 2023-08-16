// Package gtk was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gtk

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/gobject"
)

// `GtkOverlay` is a container which contains a single main child, on top
// of which it can place “overlay” widgets.
//
// ![An example GtkOverlay](overlay.png)
//
// The position of each overlay widget is determined by its
// [property@Gtk.Widget:halign] and [property@Gtk.Widget:valign]
// properties. E.g. a widget with both alignments set to %GTK_ALIGN_START
// will be placed at the top left corner of the `GtkOverlay` container,
// whereas an overlay with halign set to %GTK_ALIGN_CENTER and valign set
// to %GTK_ALIGN_END will be placed a the bottom edge of the `GtkOverlay`,
// horizontally centered. The position can be adjusted by setting the margin
// properties of the child to non-zero values.
//
// More complicated placement of overlays is possible by connecting
// to the [signal@Gtk.Overlay::get-child-position] signal.
//
// An overlay’s minimum and natural sizes are those of its main child.
// The sizes of overlay children are not considered when measuring these
// preferred sizes.
//
// # GtkOverlay as GtkBuildable
//
// The `GtkOverlay` implementation of the `GtkBuildable` interface
// supports placing a child as an overlay by specifying “overlay” as
// the “type” attribute of a `&lt;child&gt;` element.
//
// # CSS nodes
//
// `GtkOverlay` has a single CSS node with the name “overlay”. Overlay children
// whose alignments cause them to be positioned at an edge get the style classes
// “.left”, “.right”, “.top”, and/or “.bottom” according to their position.
type Overlay struct {
	Widget
}

func OverlayNewFromInternalPtr(ptr uintptr) *Overlay {
	cls := &Overlay{}
	cls.Ptr = ptr
	return cls
}

var xNewOverlay func() uintptr

// Creates a new `GtkOverlay`.
func NewOverlay() *Widget {
	var cls *Widget

	cret := xNewOverlay()

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &Widget{}
	cls.Ptr = cret
	return cls
}

var xOverlayAddOverlay func(uintptr, uintptr)

// Adds @widget to @overlay.
//
// The widget will be stacked on top of the main widget
// added with [method@Gtk.Overlay.set_child].
//
// The position at which @widget is placed is determined
// from its [property@Gtk.Widget:halign] and
// [property@Gtk.Widget:valign] properties.
func (x *Overlay) AddOverlay(WidgetVar *Widget) {

	xOverlayAddOverlay(x.GoPointer(), WidgetVar.GoPointer())

}

var xOverlayGetChild func(uintptr) uintptr

// Gets the child widget of @overlay.
func (x *Overlay) GetChild() *Widget {
	var cls *Widget

	cret := xOverlayGetChild(x.GoPointer())

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &Widget{}
	cls.Ptr = cret
	return cls
}

var xOverlayGetClipOverlay func(uintptr, uintptr) bool

// Gets whether @widget should be clipped within the parent.
func (x *Overlay) GetClipOverlay(WidgetVar *Widget) bool {

	cret := xOverlayGetClipOverlay(x.GoPointer(), WidgetVar.GoPointer())
	return cret
}

var xOverlayGetMeasureOverlay func(uintptr, uintptr) bool

// Gets whether @widget's size is included in the measurement of
// @overlay.
func (x *Overlay) GetMeasureOverlay(WidgetVar *Widget) bool {

	cret := xOverlayGetMeasureOverlay(x.GoPointer(), WidgetVar.GoPointer())
	return cret
}

var xOverlayRemoveOverlay func(uintptr, uintptr)

// Removes an overlay that was added with gtk_overlay_add_overlay().
func (x *Overlay) RemoveOverlay(WidgetVar *Widget) {

	xOverlayRemoveOverlay(x.GoPointer(), WidgetVar.GoPointer())

}

var xOverlaySetChild func(uintptr, uintptr)

// Sets the child widget of @overlay.
func (x *Overlay) SetChild(ChildVar *Widget) {

	xOverlaySetChild(x.GoPointer(), ChildVar.GoPointer())

}

var xOverlaySetClipOverlay func(uintptr, uintptr, bool)

// Sets whether @widget should be clipped within the parent.
func (x *Overlay) SetClipOverlay(WidgetVar *Widget, ClipOverlayVar bool) {

	xOverlaySetClipOverlay(x.GoPointer(), WidgetVar.GoPointer(), ClipOverlayVar)

}

var xOverlaySetMeasureOverlay func(uintptr, uintptr, bool)

// Sets whether @widget is included in the measured size of @overlay.
//
// The overlay will request the size of the largest child that has
// this property set to %TRUE. Children who are not included may
// be drawn outside of @overlay's allocation if they are too large.
func (x *Overlay) SetMeasureOverlay(WidgetVar *Widget, MeasureVar bool) {

	xOverlaySetMeasureOverlay(x.GoPointer(), WidgetVar.GoPointer(), MeasureVar)

}

func (c *Overlay) GoPointer() uintptr {
	return c.Ptr
}

func (c *Overlay) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// Emitted to determine the position and size of any overlay
// child widgets.
//
// A handler for this signal should fill @allocation with
// the desired position and size for @widget, relative to
// the 'main' child of @overlay.
//
// The default handler for this signal uses the @widget's
// halign and valign properties to determine the position
// and gives the widget its natural size (except that an
// alignment of %GTK_ALIGN_FILL will cause the overlay to
// be full-width/height). If the main child is a
// `GtkScrolledWindow`, the overlays are placed relative
// to its contents.
func (x *Overlay) ConnectGetChildPosition(cb func(Overlay, uintptr, uintptr) bool) {
	fcb := func(clsPtr uintptr, WidgetVarp uintptr, AllocationVarp uintptr) bool {
		fa := Overlay{}
		fa.Ptr = clsPtr

		return cb(fa, WidgetVarp, AllocationVarp)

	}
	gobject.ObjectConnect(x.GoPointer(), "signal::get-child-position", purego.NewCallback(fcb))
}

// Retrieves the `GtkAccessibleRole` for the given `GtkAccessible`.
func (x *Overlay) GetAccessibleRole() AccessibleRole {

	cret := XGtkAccessibleGetAccessibleRole(x.GoPointer())
	return cret
}

// Resets the accessible @property to its default value.
func (x *Overlay) ResetProperty(PropertyVar AccessibleProperty) {

	XGtkAccessibleResetProperty(x.GoPointer(), PropertyVar)

}

// Resets the accessible @relation to its default value.
func (x *Overlay) ResetRelation(RelationVar AccessibleRelation) {

	XGtkAccessibleResetRelation(x.GoPointer(), RelationVar)

}

// Resets the accessible @state to its default value.
func (x *Overlay) ResetState(StateVar AccessibleState) {

	XGtkAccessibleResetState(x.GoPointer(), StateVar)

}

// Updates a list of accessible properties.
//
// See the [enum@Gtk.AccessibleProperty] documentation for the
// value types of accessible properties.
//
// This function should be called by `GtkWidget` types whenever
// an accessible property change must be communicated to assistive
// technologies.
//
// Example:
// ```c
// value = gtk_adjustment_get_value (adjustment);
// gtk_accessible_update_property (GTK_ACCESSIBLE (spin_button),
//
//	GTK_ACCESSIBLE_PROPERTY_VALUE_NOW, value,
//	-1);
//
// ```
func (x *Overlay) UpdateProperty(FirstPropertyVar AccessibleProperty, varArgs ...interface{}) {

	XGtkAccessibleUpdateProperty(x.GoPointer(), FirstPropertyVar, varArgs...)

}

// Updates an array of accessible properties.
//
// This function should be called by `GtkWidget` types whenever an accessible
// property change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *Overlay) UpdatePropertyValue(NPropertiesVar int, PropertiesVar uintptr, ValuesVar uintptr) {

	XGtkAccessibleUpdatePropertyValue(x.GoPointer(), NPropertiesVar, PropertiesVar, ValuesVar)

}

// Updates a list of accessible relations.
//
// This function should be called by `GtkWidget` types whenever an accessible
// relation change must be communicated to assistive technologies.
//
// If the [enum@Gtk.AccessibleRelation] requires a list of references,
// you should pass each reference individually, followed by %NULL, e.g.
//
// ```c
// gtk_accessible_update_relation (accessible,
//
//	GTK_ACCESSIBLE_RELATION_CONTROLS,
//	  ref1, NULL,
//	GTK_ACCESSIBLE_RELATION_LABELLED_BY,
//	  ref1, ref2, ref3, NULL,
//	-1);
//
// ```
func (x *Overlay) UpdateRelation(FirstRelationVar AccessibleRelation, varArgs ...interface{}) {

	XGtkAccessibleUpdateRelation(x.GoPointer(), FirstRelationVar, varArgs...)

}

// Updates an array of accessible relations.
//
// This function should be called by `GtkWidget` types whenever an accessible
// relation change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *Overlay) UpdateRelationValue(NRelationsVar int, RelationsVar uintptr, ValuesVar uintptr) {

	XGtkAccessibleUpdateRelationValue(x.GoPointer(), NRelationsVar, RelationsVar, ValuesVar)

}

// Updates a list of accessible states. See the [enum@Gtk.AccessibleState]
// documentation for the value types of accessible states.
//
// This function should be called by `GtkWidget` types whenever an accessible
// state change must be communicated to assistive technologies.
//
// Example:
// ```c
// value = GTK_ACCESSIBLE_TRISTATE_MIXED;
// gtk_accessible_update_state (GTK_ACCESSIBLE (check_button),
//
//	GTK_ACCESSIBLE_STATE_CHECKED, value,
//	-1);
//
// ```
func (x *Overlay) UpdateState(FirstStateVar AccessibleState, varArgs ...interface{}) {

	XGtkAccessibleUpdateState(x.GoPointer(), FirstStateVar, varArgs...)

}

// Updates an array of accessible states.
//
// This function should be called by `GtkWidget` types whenever an accessible
// state change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *Overlay) UpdateStateValue(NStatesVar int, StatesVar uintptr, ValuesVar uintptr) {

	XGtkAccessibleUpdateStateValue(x.GoPointer(), NStatesVar, StatesVar, ValuesVar)

}

// Gets the ID of the @buildable object.
//
// `GtkBuilder` sets the name based on the ID attribute
// of the &lt;object&gt; tag used to construct the @buildable.
func (x *Overlay) GetBuildableId() string {

	cret := XGtkBuildableGetBuildableId(x.GoPointer())
	return cret
}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GTK"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xNewOverlay, lib, "gtk_overlay_new")

	core.PuregoSafeRegister(&xOverlayAddOverlay, lib, "gtk_overlay_add_overlay")
	core.PuregoSafeRegister(&xOverlayGetChild, lib, "gtk_overlay_get_child")
	core.PuregoSafeRegister(&xOverlayGetClipOverlay, lib, "gtk_overlay_get_clip_overlay")
	core.PuregoSafeRegister(&xOverlayGetMeasureOverlay, lib, "gtk_overlay_get_measure_overlay")
	core.PuregoSafeRegister(&xOverlayRemoveOverlay, lib, "gtk_overlay_remove_overlay")
	core.PuregoSafeRegister(&xOverlaySetChild, lib, "gtk_overlay_set_child")
	core.PuregoSafeRegister(&xOverlaySetClipOverlay, lib, "gtk_overlay_set_clip_overlay")
	core.PuregoSafeRegister(&xOverlaySetMeasureOverlay, lib, "gtk_overlay_set_measure_overlay")

}
