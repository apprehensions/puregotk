// Package gtk was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gtk

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/cairo"
	"github.com/jwijenbergh/puregotk/v4/glib"
	"github.com/jwijenbergh/puregotk/v4/gobject"
)

// Whenever @drawing_area needs to redraw, this function will be called.
//
// This function should exclusively redraw the contents of the drawing area
// and must not call any widget functions that cause changes.
type DrawingAreaDrawFunc func(uintptr, *cairo.Context, int, int, uintptr)

type DrawingAreaClass struct {
	ParentClass uintptr

	Padding uintptr
}

// `GtkDrawingArea` is a widget that allows drawing with cairo.
//
// ![An example GtkDrawingArea](drawingarea.png)
//
// It’s essentially a blank widget; you can draw on it. After
// creating a drawing area, the application may want to connect to:
//
//   - The [signal@Gtk.Widget::realize] signal to take any necessary actions
//     when the widget is instantiated on a particular display.
//     (Create GDK resources in response to this signal.)
//
//   - The [signal@Gtk.DrawingArea::resize] signal to take any necessary
//     actions when the widget changes size.
//
//   - Call [method@Gtk.DrawingArea.set_draw_func] to handle redrawing the
//     contents of the widget.
//
// The following code portion demonstrates using a drawing
// area to display a circle in the normal widget foreground
// color.
//
// ## Simple GtkDrawingArea usage
//
// ```c
// static void
// draw_function (GtkDrawingArea *area,
//
//	cairo_t        *cr,
//	int             width,
//	int             height,
//	gpointer        data)
//
//	{
//	  GdkRGBA color;
//	  GtkStyleContext *context;
//
//	  context = gtk_widget_get_style_context (GTK_WIDGET (area));
//
//	  cairo_arc (cr,
//	             width / 2.0, height / 2.0,
//	             MIN (width, height) / 2.0,
//	             0, 2 * G_PI);
//
//	  gtk_style_context_get_color (context,
//	                               &amp;color);
//	  gdk_cairo_set_source_rgba (cr, &amp;color);
//
//	  cairo_fill (cr);
//	}
//
// int
// main (int argc, char **argv)
//
//	{
//	  gtk_init ();
//
//	  GtkWidget *area = gtk_drawing_area_new ();
//	  gtk_drawing_area_set_content_width (GTK_DRAWING_AREA (area), 100);
//	  gtk_drawing_area_set_content_height (GTK_DRAWING_AREA (area), 100);
//	  gtk_drawing_area_set_draw_func (GTK_DRAWING_AREA (area),
//	                                  draw_function,
//	                                  NULL, NULL);
//	  return 0;
//	}
//
// ```
//
// The draw function is normally called when a drawing area first comes
// onscreen, or when it’s covered by another window and then uncovered.
// You can also force a redraw by adding to the “damage region” of the
// drawing area’s window using [method@Gtk.Widget.queue_draw].
// This will cause the drawing area to call the draw function again.
//
// The available routines for drawing are documented in the
// [Cairo documentation](https://www.cairographics.org/manual/); GDK
// offers additional API to integrate with Cairo, like [func@Gdk.cairo_set_source_rgba]
// or [func@Gdk.cairo_set_source_pixbuf].
//
// To receive mouse events on a drawing area, you will need to use
// event controllers. To receive keyboard events, you will need to set
// the “can-focus” property on the drawing area, and you should probably
// draw some user-visible indication that the drawing area is focused.
//
// If you need more complex control over your widget, you should consider
// creating your own `GtkWidget` subclass.
type DrawingArea struct {
	Widget
}

func DrawingAreaNewFromInternalPtr(ptr uintptr) *DrawingArea {
	cls := &DrawingArea{}
	cls.Ptr = ptr
	return cls
}

var xNewDrawingArea func() uintptr

// Creates a new drawing area.
func NewDrawingArea() *Widget {
	var cls *Widget

	cret := xNewDrawingArea()

	if cret == 0 {
		return cls
	}
	gobject.IncreaseRef(cret)
	cls = &Widget{}
	cls.Ptr = cret
	return cls
}

var xDrawingAreaGetContentHeight func(uintptr) int

// Retrieves the content height of the `GtkDrawingArea`.
func (x *DrawingArea) GetContentHeight() int {

	cret := xDrawingAreaGetContentHeight(x.GoPointer())
	return cret
}

var xDrawingAreaGetContentWidth func(uintptr) int

// Retrieves the content width of the `GtkDrawingArea`.
func (x *DrawingArea) GetContentWidth() int {

	cret := xDrawingAreaGetContentWidth(x.GoPointer())
	return cret
}

var xDrawingAreaSetContentHeight func(uintptr, int)

// Sets the desired height of the contents of the drawing area.
//
// Note that because widgets may be allocated larger sizes than they
// requested, it is possible that the actual height passed to your draw
// function is larger than the height set here. You can use
// [method@Gtk.Widget.set_valign] to avoid that.
//
// If the height is set to 0 (the default), the drawing area may disappear.
func (x *DrawingArea) SetContentHeight(HeightVar int) {

	xDrawingAreaSetContentHeight(x.GoPointer(), HeightVar)

}

var xDrawingAreaSetContentWidth func(uintptr, int)

// Sets the desired width of the contents of the drawing area.
//
// Note that because widgets may be allocated larger sizes than they
// requested, it is possible that the actual width passed to your draw
// function is larger than the width set here. You can use
// [method@Gtk.Widget.set_halign] to avoid that.
//
// If the width is set to 0 (the default), the drawing area may disappear.
func (x *DrawingArea) SetContentWidth(WidthVar int) {

	xDrawingAreaSetContentWidth(x.GoPointer(), WidthVar)

}

var xDrawingAreaSetDrawFunc func(uintptr, uintptr, uintptr, uintptr)

// Setting a draw function is the main thing you want to do when using
// a drawing area.
//
// The draw function is called whenever GTK needs to draw the contents
// of the drawing area to the screen.
//
// The draw function will be called during the drawing stage of GTK.
// In the drawing stage it is not allowed to change properties of any
// GTK widgets or call any functions that would cause any properties
// to be changed. You should restrict yourself exclusively to drawing
// your contents in the draw function.
//
// If what you are drawing does change, call [method@Gtk.Widget.queue_draw]
// on the drawing area. This will cause a redraw and will call @draw_func again.
func (x *DrawingArea) SetDrawFunc(DrawFuncVar DrawingAreaDrawFunc, UserDataVar uintptr, DestroyVar glib.DestroyNotify) {

	xDrawingAreaSetDrawFunc(x.GoPointer(), purego.NewCallback(DrawFuncVar), UserDataVar, purego.NewCallback(DestroyVar))

}

func (c *DrawingArea) GoPointer() uintptr {
	return c.Ptr
}

func (c *DrawingArea) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// Emitted once when the widget is realized, and then each time the widget
// is changed while realized.
//
// This is useful in order to keep state up to date with the widget size,
// like for instance a backing surface.
func (x *DrawingArea) ConnectResize(cb func(DrawingArea, int, int)) {
	fcb := func(clsPtr uintptr, WidthVarp int, HeightVarp int) {
		fa := DrawingArea{}
		fa.Ptr = clsPtr

		cb(fa, WidthVarp, HeightVarp)

	}
	gobject.ObjectConnect(x.GoPointer(), "signal::resize", purego.NewCallback(fcb))
}

// Retrieves the `GtkAccessibleRole` for the given `GtkAccessible`.
func (x *DrawingArea) GetAccessibleRole() AccessibleRole {

	cret := XGtkAccessibleGetAccessibleRole(x.GoPointer())
	return cret
}

// Resets the accessible @property to its default value.
func (x *DrawingArea) ResetProperty(PropertyVar AccessibleProperty) {

	XGtkAccessibleResetProperty(x.GoPointer(), PropertyVar)

}

// Resets the accessible @relation to its default value.
func (x *DrawingArea) ResetRelation(RelationVar AccessibleRelation) {

	XGtkAccessibleResetRelation(x.GoPointer(), RelationVar)

}

// Resets the accessible @state to its default value.
func (x *DrawingArea) ResetState(StateVar AccessibleState) {

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
func (x *DrawingArea) UpdateProperty(FirstPropertyVar AccessibleProperty, varArgs ...interface{}) {

	XGtkAccessibleUpdateProperty(x.GoPointer(), FirstPropertyVar, varArgs...)

}

// Updates an array of accessible properties.
//
// This function should be called by `GtkWidget` types whenever an accessible
// property change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *DrawingArea) UpdatePropertyValue(NPropertiesVar int, PropertiesVar uintptr, ValuesVar uintptr) {

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
func (x *DrawingArea) UpdateRelation(FirstRelationVar AccessibleRelation, varArgs ...interface{}) {

	XGtkAccessibleUpdateRelation(x.GoPointer(), FirstRelationVar, varArgs...)

}

// Updates an array of accessible relations.
//
// This function should be called by `GtkWidget` types whenever an accessible
// relation change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *DrawingArea) UpdateRelationValue(NRelationsVar int, RelationsVar uintptr, ValuesVar uintptr) {

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
func (x *DrawingArea) UpdateState(FirstStateVar AccessibleState, varArgs ...interface{}) {

	XGtkAccessibleUpdateState(x.GoPointer(), FirstStateVar, varArgs...)

}

// Updates an array of accessible states.
//
// This function should be called by `GtkWidget` types whenever an accessible
// state change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *DrawingArea) UpdateStateValue(NStatesVar int, StatesVar uintptr, ValuesVar uintptr) {

	XGtkAccessibleUpdateStateValue(x.GoPointer(), NStatesVar, StatesVar, ValuesVar)

}

// Gets the ID of the @buildable object.
//
// `GtkBuilder` sets the name based on the ID attribute
// of the &lt;object&gt; tag used to construct the @buildable.
func (x *DrawingArea) GetBuildableId() string {

	cret := XGtkBuildableGetBuildableId(x.GoPointer())
	return cret
}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GTK"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xNewDrawingArea, lib, "gtk_drawing_area_new")

	core.PuregoSafeRegister(&xDrawingAreaGetContentHeight, lib, "gtk_drawing_area_get_content_height")
	core.PuregoSafeRegister(&xDrawingAreaGetContentWidth, lib, "gtk_drawing_area_get_content_width")
	core.PuregoSafeRegister(&xDrawingAreaSetContentHeight, lib, "gtk_drawing_area_set_content_height")
	core.PuregoSafeRegister(&xDrawingAreaSetContentWidth, lib, "gtk_drawing_area_set_content_width")
	core.PuregoSafeRegister(&xDrawingAreaSetDrawFunc, lib, "gtk_drawing_area_set_draw_func")

}
