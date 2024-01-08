// Package adw was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package adw

import (
	"unsafe"

	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/gdk"
	"github.com/jwijenbergh/puregotk/v4/gobject"
	"github.com/jwijenbergh/puregotk/v4/gtk"
)

type CarouselClass struct {
	ParentClass uintptr
}

func (x *CarouselClass) GoPointer() uintptr {
	return uintptr(unsafe.Pointer(x))
}

// A paginated scrolling widget.
//
// &lt;picture&gt;
//
//	&lt;source srcset="carousel-dark.png" media="(prefers-color-scheme: dark)"&gt;
//	&lt;img src="carousel.png" alt="carousel"&gt;
//
// &lt;/picture&gt;
//
// The `AdwCarousel` widget can be used to display a set of pages with
// swipe-based navigation between them.
//
// [class@CarouselIndicatorDots] and [class@CarouselIndicatorLines] can be used
// to provide page indicators for `AdwCarousel`.
//
// ## CSS nodes
//
// `AdwCarousel` has a single CSS node with name `carousel`.
type Carousel struct {
	gtk.Widget
}

func CarouselNewFromInternalPtr(ptr uintptr) *Carousel {
	cls := &Carousel{}
	cls.Ptr = ptr
	return cls
}

var xNewCarousel func() uintptr

// Creates a new `AdwCarousel`.
func NewCarousel() *gtk.Widget {
	var cls *gtk.Widget

	cret := xNewCarousel()

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &gtk.Widget{}
	cls.Ptr = cret
	return cls
}

var xCarouselAppend func(uintptr, uintptr)

// Appends @child to @self.
func (x *Carousel) Append(ChildVar *gtk.Widget) {

	xCarouselAppend(x.GoPointer(), ChildVar.GoPointer())

}

var xCarouselGetAllowLongSwipes func(uintptr) bool

// Gets whether to allow swiping for more than one page at a time.
func (x *Carousel) GetAllowLongSwipes() bool {

	cret := xCarouselGetAllowLongSwipes(x.GoPointer())
	return cret
}

var xCarouselGetAllowMouseDrag func(uintptr) bool

// Sets whether @self can be dragged with mouse pointer.
func (x *Carousel) GetAllowMouseDrag() bool {

	cret := xCarouselGetAllowMouseDrag(x.GoPointer())
	return cret
}

var xCarouselGetAllowScrollWheel func(uintptr) bool

// Gets whether @self will respond to scroll wheel events.
func (x *Carousel) GetAllowScrollWheel() bool {

	cret := xCarouselGetAllowScrollWheel(x.GoPointer())
	return cret
}

var xCarouselGetInteractive func(uintptr) bool

// Gets whether @self can be navigated.
func (x *Carousel) GetInteractive() bool {

	cret := xCarouselGetInteractive(x.GoPointer())
	return cret
}

var xCarouselGetNPages func(uintptr) uint

// Gets the number of pages in @self.
func (x *Carousel) GetNPages() uint {

	cret := xCarouselGetNPages(x.GoPointer())
	return cret
}

var xCarouselGetNthPage func(uintptr, uint) uintptr

// Gets the page at position @n.
func (x *Carousel) GetNthPage(NVar uint) *gtk.Widget {
	var cls *gtk.Widget

	cret := xCarouselGetNthPage(x.GoPointer(), NVar)

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &gtk.Widget{}
	cls.Ptr = cret
	return cls
}

var xCarouselGetPosition func(uintptr) float64

// Gets current scroll position in @self, unitless.
//
// 1 matches 1 page. Use [method@Carousel.scroll_to] for changing it.
func (x *Carousel) GetPosition() float64 {

	cret := xCarouselGetPosition(x.GoPointer())
	return cret
}

var xCarouselGetRevealDuration func(uintptr) uint

// Gets the page reveal duration, in milliseconds.
func (x *Carousel) GetRevealDuration() uint {

	cret := xCarouselGetRevealDuration(x.GoPointer())
	return cret
}

var xCarouselGetScrollParams func(uintptr) *SpringParams

// Gets the scroll animation spring parameters for @self.
func (x *Carousel) GetScrollParams() *SpringParams {

	cret := xCarouselGetScrollParams(x.GoPointer())
	return cret
}

var xCarouselGetSpacing func(uintptr) uint

// Gets spacing between pages in pixels.
func (x *Carousel) GetSpacing() uint {

	cret := xCarouselGetSpacing(x.GoPointer())
	return cret
}

var xCarouselInsert func(uintptr, uintptr, int)

// Inserts @child into @self at position @position.
//
// If position is -1, or larger than the number of pages,
// @child will be appended to the end.
func (x *Carousel) Insert(ChildVar *gtk.Widget, PositionVar int) {

	xCarouselInsert(x.GoPointer(), ChildVar.GoPointer(), PositionVar)

}

var xCarouselPrepend func(uintptr, uintptr)

// Prepends @child to @self.
func (x *Carousel) Prepend(ChildVar *gtk.Widget) {

	xCarouselPrepend(x.GoPointer(), ChildVar.GoPointer())

}

var xCarouselRemove func(uintptr, uintptr)

// Removes @child from @self.
func (x *Carousel) Remove(ChildVar *gtk.Widget) {

	xCarouselRemove(x.GoPointer(), ChildVar.GoPointer())

}

var xCarouselReorder func(uintptr, uintptr, int)

// Moves @child into position @position.
//
// If position is -1, or larger than the number of pages, @child will be moved
// at the end.
func (x *Carousel) Reorder(ChildVar *gtk.Widget, PositionVar int) {

	xCarouselReorder(x.GoPointer(), ChildVar.GoPointer(), PositionVar)

}

var xCarouselScrollTo func(uintptr, uintptr, bool)

// Scrolls to @widget.
//
// If @animate is `TRUE`, the transition will be animated.
func (x *Carousel) ScrollTo(WidgetVar *gtk.Widget, AnimateVar bool) {

	xCarouselScrollTo(x.GoPointer(), WidgetVar.GoPointer(), AnimateVar)

}

var xCarouselSetAllowLongSwipes func(uintptr, bool)

// Sets whether to allow swiping for more than one page at a time.
//
// If @allow_long_swipes is `FALSE`, each swipe can only move to the adjacent
// pages.
func (x *Carousel) SetAllowLongSwipes(AllowLongSwipesVar bool) {

	xCarouselSetAllowLongSwipes(x.GoPointer(), AllowLongSwipesVar)

}

var xCarouselSetAllowMouseDrag func(uintptr, bool)

// Sets whether @self can be dragged with mouse pointer.
//
// If @allow_mouse_drag is `FALSE`, dragging is only available on touch.
func (x *Carousel) SetAllowMouseDrag(AllowMouseDragVar bool) {

	xCarouselSetAllowMouseDrag(x.GoPointer(), AllowMouseDragVar)

}

var xCarouselSetAllowScrollWheel func(uintptr, bool)

// Sets whether @self will respond to scroll wheel events.
//
// If @allow_scroll_wheel is `FALSE`, wheel events will be ignored.
func (x *Carousel) SetAllowScrollWheel(AllowScrollWheelVar bool) {

	xCarouselSetAllowScrollWheel(x.GoPointer(), AllowScrollWheelVar)

}

var xCarouselSetInteractive func(uintptr, bool)

// Sets whether @self can be navigated.
//
// This can be used to temporarily disable the carousel to only allow navigating
// it in a certain state.
func (x *Carousel) SetInteractive(InteractiveVar bool) {

	xCarouselSetInteractive(x.GoPointer(), InteractiveVar)

}

var xCarouselSetRevealDuration func(uintptr, uint)

// Sets the page reveal duration, in milliseconds.
//
// Reveal duration is used when animating adding or removing pages.
func (x *Carousel) SetRevealDuration(RevealDurationVar uint) {

	xCarouselSetRevealDuration(x.GoPointer(), RevealDurationVar)

}

var xCarouselSetScrollParams func(uintptr, *SpringParams)

// Sets the scroll animation spring parameters for @self.
//
// The default value is equivalent to:
//
// ```c
// adw_spring_params_new (1, 0.5, 500)
// ```
func (x *Carousel) SetScrollParams(ParamsVar *SpringParams) {

	xCarouselSetScrollParams(x.GoPointer(), ParamsVar)

}

var xCarouselSetSpacing func(uintptr, uint)

// Sets spacing between pages in pixels.
func (x *Carousel) SetSpacing(SpacingVar uint) {

	xCarouselSetSpacing(x.GoPointer(), SpacingVar)

}

func (c *Carousel) GoPointer() uintptr {
	return c.Ptr
}

func (c *Carousel) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// This signal is emitted after a page has been changed.
//
// It can be used to implement "infinite scrolling" by amending the pages
// after every scroll. Note that an empty carousel is indicated by
// `(int)index == -1`.
func (x *Carousel) ConnectPageChanged(cb func(Carousel, uint)) uint32 {
	fcb := func(clsPtr uintptr, IndexVarp uint) {
		fa := Carousel{}
		fa.Ptr = clsPtr

		cb(fa, IndexVarp)

	}
	return gobject.SignalConnect(x.GoPointer(), "page-changed", purego.NewCallback(fcb))
}

// Gets the progress @self will snap back to after the gesture is canceled.
func (x *Carousel) GetCancelProgress() float64 {

	cret := XAdwSwipeableGetCancelProgress(x.GoPointer())
	return cret
}

// Gets the swipe distance of @self.
//
// This corresponds to how many pixels 1 unit represents.
func (x *Carousel) GetDistance() float64 {

	cret := XAdwSwipeableGetDistance(x.GoPointer())
	return cret
}

// Gets the current progress of @self.
func (x *Carousel) GetProgress() float64 {

	cret := XAdwSwipeableGetProgress(x.GoPointer())
	return cret
}

// Gets the snap points of @self.
//
// Each snap point represents a progress value that is considered acceptable to
// end the swipe on.
func (x *Carousel) GetSnapPoints(NSnapPointsVar int) uintptr {

	cret := XAdwSwipeableGetSnapPoints(x.GoPointer(), NSnapPointsVar)
	return cret
}

// Gets the area @self can start a swipe from for the given direction and
// gesture type.
//
// This can be used to restrict swipes to only be possible from a certain area,
// for example, to only allow edge swipes, or to have a draggable element and
// ignore swipes elsewhere.
//
// If not implemented, the default implementation returns the allocation of
// @self, allowing swipes from anywhere.
func (x *Carousel) GetSwipeArea(NavigationDirectionVar NavigationDirection, IsDragVar bool, RectVar *gdk.Rectangle) {

	XAdwSwipeableGetSwipeArea(x.GoPointer(), NavigationDirectionVar, IsDragVar, RectVar)

}

// Retrieves the `GtkAccessibleRole` for the given `GtkAccessible`.
func (x *Carousel) GetAccessibleRole() gtk.AccessibleRole {

	cret := gtk.XGtkAccessibleGetAccessibleRole(x.GoPointer())
	return cret
}

// Resets the accessible @property to its default value.
func (x *Carousel) ResetProperty(PropertyVar gtk.AccessibleProperty) {

	gtk.XGtkAccessibleResetProperty(x.GoPointer(), PropertyVar)

}

// Resets the accessible @relation to its default value.
func (x *Carousel) ResetRelation(RelationVar gtk.AccessibleRelation) {

	gtk.XGtkAccessibleResetRelation(x.GoPointer(), RelationVar)

}

// Resets the accessible @state to its default value.
func (x *Carousel) ResetState(StateVar gtk.AccessibleState) {

	gtk.XGtkAccessibleResetState(x.GoPointer(), StateVar)

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
func (x *Carousel) UpdateProperty(FirstPropertyVar gtk.AccessibleProperty, varArgs ...interface{}) {

	gtk.XGtkAccessibleUpdateProperty(x.GoPointer(), FirstPropertyVar, varArgs...)

}

// Updates an array of accessible properties.
//
// This function should be called by `GtkWidget` types whenever an accessible
// property change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *Carousel) UpdatePropertyValue(NPropertiesVar int, PropertiesVar uintptr, ValuesVar uintptr) {

	gtk.XGtkAccessibleUpdatePropertyValue(x.GoPointer(), NPropertiesVar, PropertiesVar, ValuesVar)

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
func (x *Carousel) UpdateRelation(FirstRelationVar gtk.AccessibleRelation, varArgs ...interface{}) {

	gtk.XGtkAccessibleUpdateRelation(x.GoPointer(), FirstRelationVar, varArgs...)

}

// Updates an array of accessible relations.
//
// This function should be called by `GtkWidget` types whenever an accessible
// relation change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *Carousel) UpdateRelationValue(NRelationsVar int, RelationsVar uintptr, ValuesVar uintptr) {

	gtk.XGtkAccessibleUpdateRelationValue(x.GoPointer(), NRelationsVar, RelationsVar, ValuesVar)

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
func (x *Carousel) UpdateState(FirstStateVar gtk.AccessibleState, varArgs ...interface{}) {

	gtk.XGtkAccessibleUpdateState(x.GoPointer(), FirstStateVar, varArgs...)

}

// Updates an array of accessible states.
//
// This function should be called by `GtkWidget` types whenever an accessible
// state change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *Carousel) UpdateStateValue(NStatesVar int, StatesVar uintptr, ValuesVar uintptr) {

	gtk.XGtkAccessibleUpdateStateValue(x.GoPointer(), NStatesVar, StatesVar, ValuesVar)

}

// Gets the ID of the @buildable object.
//
// `GtkBuilder` sets the name based on the ID attribute
// of the &lt;object&gt; tag used to construct the @buildable.
func (x *Carousel) GetBuildableId() string {

	cret := gtk.XGtkBuildableGetBuildableId(x.GoPointer())
	return cret
}

// Retrieves the orientation of the @orientable.
func (x *Carousel) GetOrientation() gtk.Orientation {

	cret := gtk.XGtkOrientableGetOrientation(x.GoPointer())
	return cret
}

// Sets the orientation of the @orientable.
func (x *Carousel) SetOrientation(OrientationVar gtk.Orientation) {

	gtk.XGtkOrientableSetOrientation(x.GoPointer(), OrientationVar)

}

func init() {
	lib, err := purego.Dlopen(core.GetPath("ADW"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xNewCarousel, lib, "adw_carousel_new")

	core.PuregoSafeRegister(&xCarouselAppend, lib, "adw_carousel_append")
	core.PuregoSafeRegister(&xCarouselGetAllowLongSwipes, lib, "adw_carousel_get_allow_long_swipes")
	core.PuregoSafeRegister(&xCarouselGetAllowMouseDrag, lib, "adw_carousel_get_allow_mouse_drag")
	core.PuregoSafeRegister(&xCarouselGetAllowScrollWheel, lib, "adw_carousel_get_allow_scroll_wheel")
	core.PuregoSafeRegister(&xCarouselGetInteractive, lib, "adw_carousel_get_interactive")
	core.PuregoSafeRegister(&xCarouselGetNPages, lib, "adw_carousel_get_n_pages")
	core.PuregoSafeRegister(&xCarouselGetNthPage, lib, "adw_carousel_get_nth_page")
	core.PuregoSafeRegister(&xCarouselGetPosition, lib, "adw_carousel_get_position")
	core.PuregoSafeRegister(&xCarouselGetRevealDuration, lib, "adw_carousel_get_reveal_duration")
	core.PuregoSafeRegister(&xCarouselGetScrollParams, lib, "adw_carousel_get_scroll_params")
	core.PuregoSafeRegister(&xCarouselGetSpacing, lib, "adw_carousel_get_spacing")
	core.PuregoSafeRegister(&xCarouselInsert, lib, "adw_carousel_insert")
	core.PuregoSafeRegister(&xCarouselPrepend, lib, "adw_carousel_prepend")
	core.PuregoSafeRegister(&xCarouselRemove, lib, "adw_carousel_remove")
	core.PuregoSafeRegister(&xCarouselReorder, lib, "adw_carousel_reorder")
	core.PuregoSafeRegister(&xCarouselScrollTo, lib, "adw_carousel_scroll_to")
	core.PuregoSafeRegister(&xCarouselSetAllowLongSwipes, lib, "adw_carousel_set_allow_long_swipes")
	core.PuregoSafeRegister(&xCarouselSetAllowMouseDrag, lib, "adw_carousel_set_allow_mouse_drag")
	core.PuregoSafeRegister(&xCarouselSetAllowScrollWheel, lib, "adw_carousel_set_allow_scroll_wheel")
	core.PuregoSafeRegister(&xCarouselSetInteractive, lib, "adw_carousel_set_interactive")
	core.PuregoSafeRegister(&xCarouselSetRevealDuration, lib, "adw_carousel_set_reveal_duration")
	core.PuregoSafeRegister(&xCarouselSetScrollParams, lib, "adw_carousel_set_scroll_params")
	core.PuregoSafeRegister(&xCarouselSetSpacing, lib, "adw_carousel_set_spacing")

}
