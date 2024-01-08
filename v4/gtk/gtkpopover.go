// Package gtk was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gtk

import (
	"unsafe"

	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/gdk"
	"github.com/jwijenbergh/puregotk/v4/gobject"
	"github.com/jwijenbergh/puregotk/v4/gsk"
)

type PopoverClass struct {
	ParentClass uintptr

	Reserved uintptr
}

func (x *PopoverClass) GoPointer() uintptr {
	return uintptr(unsafe.Pointer(x))
}

// `GtkPopover` is a bubble-like context popup.
//
// ![An example GtkPopover](popover.png)
//
// It is primarily meant to provide context-dependent information
// or options. Popovers are attached to a parent widget. By default,
// they point to the whole widget area, although this behavior can be
// changed with [method@Gtk.Popover.set_pointing_to].
//
// The position of a popover relative to the widget it is attached to
// can also be changed with [method@Gtk.Popover.set_position]
//
// By default, `GtkPopover` performs a grab, in order to ensure input
// events get redirected to it while it is shown, and also so the popover
// is dismissed in the expected situations (clicks outside the popover,
// or the Escape key being pressed). If no such modal behavior is desired
// on a popover, [method@Gtk.Popover.set_autohide] may be called on it to
// tweak its behavior.
//
// ## GtkPopover as menu replacement
//
// `GtkPopover` is often used to replace menus. The best was to do this
// is to use the [class@Gtk.PopoverMenu] subclass which supports being
// populated from a `GMenuModel` with [ctor@Gtk.PopoverMenu.new_from_model].
//
// ```xml
// &lt;section&gt;
//
//	&lt;attribute name="display-hint"&gt;horizontal-buttons&lt;/attribute&gt;
//	&lt;item&gt;
//	  &lt;attribute name="label"&gt;Cut&lt;/attribute&gt;
//	  &lt;attribute name="action"&gt;app.cut&lt;/attribute&gt;
//	  &lt;attribute name="verb-icon"&gt;edit-cut-symbolic&lt;/attribute&gt;
//	&lt;/item&gt;
//	&lt;item&gt;
//	  &lt;attribute name="label"&gt;Copy&lt;/attribute&gt;
//	  &lt;attribute name="action"&gt;app.copy&lt;/attribute&gt;
//	  &lt;attribute name="verb-icon"&gt;edit-copy-symbolic&lt;/attribute&gt;
//	&lt;/item&gt;
//	&lt;item&gt;
//	  &lt;attribute name="label"&gt;Paste&lt;/attribute&gt;
//	  &lt;attribute name="action"&gt;app.paste&lt;/attribute&gt;
//	  &lt;attribute name="verb-icon"&gt;edit-paste-symbolic&lt;/attribute&gt;
//	&lt;/item&gt;
//
// &lt;/section&gt;
// ```
//
// # CSS nodes
//
// ```
// popover[.menu]
// ├── arrow
// ╰── contents.background
//
//	╰── &lt;child&gt;
//
// ```
//
// The contents child node always gets the .background style class
// and the popover itself gets the .menu style class if the popover
// is menu-like (i.e. `GtkPopoverMenu`).
//
// Particular uses of `GtkPopover`, such as touch selection popups or
// magnifiers in `GtkEntry` or `GtkTextView` get style classes like
// .touch-selection or .magnifier to differentiate from plain popovers.
//
// When styling a popover directly, the popover node should usually
// not have any background. The visible part of the popover can have
// a shadow. To specify it in CSS, set the box-shadow of the contents node.
//
// Note that, in order to accomplish appropriate arrow visuals, `GtkPopover`
// uses custom drawing for the arrow node. This makes it possible for the
// arrow to change its shape dynamically, but it also limits the possibilities
// of styling it using CSS. In particular, the arrow gets drawn over the
// content node's border and shadow, so they look like one shape, which
// means that the border width of the content node and the arrow node should
// be the same. The arrow also does not support any border shape other than
// solid, no border-radius, only one border width (border-bottom-width is
// used) and no box-shadow.
type Popover struct {
	Widget
}

func PopoverNewFromInternalPtr(ptr uintptr) *Popover {
	cls := &Popover{}
	cls.Ptr = ptr
	return cls
}

var xNewPopover func() uintptr

// Creates a new `GtkPopover`.
func NewPopover() *Widget {
	var cls *Widget

	cret := xNewPopover()

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &Widget{}
	cls.Ptr = cret
	return cls
}

var xPopoverGetAutohide func(uintptr) bool

// Returns whether the popover is modal.
//
// See [method@Gtk.Popover.set_autohide] for the
// implications of this.
func (x *Popover) GetAutohide() bool {

	cret := xPopoverGetAutohide(x.GoPointer())
	return cret
}

var xPopoverGetCascadePopdown func(uintptr) bool

// Returns whether the popover will close after a modal child is closed.
func (x *Popover) GetCascadePopdown() bool {

	cret := xPopoverGetCascadePopdown(x.GoPointer())
	return cret
}

var xPopoverGetChild func(uintptr) uintptr

// Gets the child widget of @popover.
func (x *Popover) GetChild() *Widget {
	var cls *Widget

	cret := xPopoverGetChild(x.GoPointer())

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &Widget{}
	cls.Ptr = cret
	return cls
}

var xPopoverGetHasArrow func(uintptr) bool

// Gets whether this popover is showing an arrow
// pointing at the widget that it is relative to.
func (x *Popover) GetHasArrow() bool {

	cret := xPopoverGetHasArrow(x.GoPointer())
	return cret
}

var xPopoverGetMnemonicsVisible func(uintptr) bool

// Gets whether mnemonics are visible.
func (x *Popover) GetMnemonicsVisible() bool {

	cret := xPopoverGetMnemonicsVisible(x.GoPointer())
	return cret
}

var xPopoverGetOffset func(uintptr, int, int)

// Gets the offset previous set with gtk_popover_set_offset().
func (x *Popover) GetOffset(XOffsetVar int, YOffsetVar int) {

	xPopoverGetOffset(x.GoPointer(), XOffsetVar, YOffsetVar)

}

var xPopoverGetPointingTo func(uintptr, *gdk.Rectangle) bool

// Gets the rectangle that the popover points to.
//
// If a rectangle to point to has been set, this function will
// return %TRUE and fill in @rect with such rectangle, otherwise
// it will return %FALSE and fill in @rect with the parent
// widget coordinates.
func (x *Popover) GetPointingTo(RectVar *gdk.Rectangle) bool {

	cret := xPopoverGetPointingTo(x.GoPointer(), RectVar)
	return cret
}

var xPopoverGetPosition func(uintptr) PositionType

// Returns the preferred position of @popover.
func (x *Popover) GetPosition() PositionType {

	cret := xPopoverGetPosition(x.GoPointer())
	return cret
}

var xPopoverPopdown func(uintptr)

// Pops @popover down.
//
// This may have the side-effect of closing a parent popover
// as well. See [property@Gtk.Popover:cascade-popdown].
func (x *Popover) Popdown() {

	xPopoverPopdown(x.GoPointer())

}

var xPopoverPopup func(uintptr)

// Pops @popover up.
func (x *Popover) Popup() {

	xPopoverPopup(x.GoPointer())

}

var xPopoverPresent func(uintptr)

// Presents the popover to the user.
func (x *Popover) Present() {

	xPopoverPresent(x.GoPointer())

}

var xPopoverSetAutohide func(uintptr, bool)

// Sets whether @popover is modal.
//
// A modal popover will grab the keyboard focus on it when being
// displayed. Focus will wrap around within the popover. Clicking
// outside the popover area or pressing Esc will dismiss the popover.
//
// Called this function on an already showing popup with a new
// autohide value different from the current one, will cause the
// popup to be hidden.
func (x *Popover) SetAutohide(AutohideVar bool) {

	xPopoverSetAutohide(x.GoPointer(), AutohideVar)

}

var xPopoverSetCascadePopdown func(uintptr, bool)

// If @cascade_popdown is %TRUE, the popover will be
// closed when a child modal popover is closed.
//
// If %FALSE, @popover will stay visible.
func (x *Popover) SetCascadePopdown(CascadePopdownVar bool) {

	xPopoverSetCascadePopdown(x.GoPointer(), CascadePopdownVar)

}

var xPopoverSetChild func(uintptr, uintptr)

// Sets the child widget of @popover.
func (x *Popover) SetChild(ChildVar *Widget) {

	xPopoverSetChild(x.GoPointer(), ChildVar.GoPointer())

}

var xPopoverSetDefaultWidget func(uintptr, uintptr)

// Sets the default widget of a `GtkPopover`.
//
// The default widget is the widget that’s activated when the user
// presses Enter in a dialog (for example). This function sets or
// unsets the default widget for a `GtkPopover`.
func (x *Popover) SetDefaultWidget(WidgetVar *Widget) {

	xPopoverSetDefaultWidget(x.GoPointer(), WidgetVar.GoPointer())

}

var xPopoverSetHasArrow func(uintptr, bool)

// Sets whether this popover should draw an arrow
// pointing at the widget it is relative to.
func (x *Popover) SetHasArrow(HasArrowVar bool) {

	xPopoverSetHasArrow(x.GoPointer(), HasArrowVar)

}

var xPopoverSetMnemonicsVisible func(uintptr, bool)

// Sets whether mnemonics should be visible.
func (x *Popover) SetMnemonicsVisible(MnemonicsVisibleVar bool) {

	xPopoverSetMnemonicsVisible(x.GoPointer(), MnemonicsVisibleVar)

}

var xPopoverSetOffset func(uintptr, int, int)

// Sets the offset to use when calculating the position
// of the popover.
//
// These values are used when preparing the [struct@Gdk.PopupLayout]
// for positioning the popover.
func (x *Popover) SetOffset(XOffsetVar int, YOffsetVar int) {

	xPopoverSetOffset(x.GoPointer(), XOffsetVar, YOffsetVar)

}

var xPopoverSetPointingTo func(uintptr, *gdk.Rectangle)

// Sets the rectangle that @popover points to.
//
// This is in the coordinate space of the @popover parent.
func (x *Popover) SetPointingTo(RectVar *gdk.Rectangle) {

	xPopoverSetPointingTo(x.GoPointer(), RectVar)

}

var xPopoverSetPosition func(uintptr, PositionType)

// Sets the preferred position for @popover to appear.
//
// If the @popover is currently visible, it will be immediately
// updated.
//
// This preference will be respected where possible, although
// on lack of space (eg. if close to the window edges), the
// `GtkPopover` may choose to appear on the opposite side.
func (x *Popover) SetPosition(PositionVar PositionType) {

	xPopoverSetPosition(x.GoPointer(), PositionVar)

}

func (c *Popover) GoPointer() uintptr {
	return c.Ptr
}

func (c *Popover) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// Emitted whend the user activates the default widget.
//
// This is a [keybinding signal](class.SignalAction.html).
func (x *Popover) ConnectActivateDefault(cb func(Popover)) uint32 {
	fcb := func(clsPtr uintptr) {
		fa := Popover{}
		fa.Ptr = clsPtr

		cb(fa)

	}
	return gobject.SignalConnect(x.GoPointer(), "activate-default", purego.NewCallback(fcb))
}

// Emitted when the popover is closed.
func (x *Popover) ConnectClosed(cb func(Popover)) uint32 {
	fcb := func(clsPtr uintptr) {
		fa := Popover{}
		fa.Ptr = clsPtr

		cb(fa)

	}
	return gobject.SignalConnect(x.GoPointer(), "closed", purego.NewCallback(fcb))
}

// Retrieves the `GtkAccessibleRole` for the given `GtkAccessible`.
func (x *Popover) GetAccessibleRole() AccessibleRole {

	cret := XGtkAccessibleGetAccessibleRole(x.GoPointer())
	return cret
}

// Resets the accessible @property to its default value.
func (x *Popover) ResetProperty(PropertyVar AccessibleProperty) {

	XGtkAccessibleResetProperty(x.GoPointer(), PropertyVar)

}

// Resets the accessible @relation to its default value.
func (x *Popover) ResetRelation(RelationVar AccessibleRelation) {

	XGtkAccessibleResetRelation(x.GoPointer(), RelationVar)

}

// Resets the accessible @state to its default value.
func (x *Popover) ResetState(StateVar AccessibleState) {

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
func (x *Popover) UpdateProperty(FirstPropertyVar AccessibleProperty, varArgs ...interface{}) {

	XGtkAccessibleUpdateProperty(x.GoPointer(), FirstPropertyVar, varArgs...)

}

// Updates an array of accessible properties.
//
// This function should be called by `GtkWidget` types whenever an accessible
// property change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *Popover) UpdatePropertyValue(NPropertiesVar int, PropertiesVar uintptr, ValuesVar uintptr) {

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
func (x *Popover) UpdateRelation(FirstRelationVar AccessibleRelation, varArgs ...interface{}) {

	XGtkAccessibleUpdateRelation(x.GoPointer(), FirstRelationVar, varArgs...)

}

// Updates an array of accessible relations.
//
// This function should be called by `GtkWidget` types whenever an accessible
// relation change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *Popover) UpdateRelationValue(NRelationsVar int, RelationsVar uintptr, ValuesVar uintptr) {

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
func (x *Popover) UpdateState(FirstStateVar AccessibleState, varArgs ...interface{}) {

	XGtkAccessibleUpdateState(x.GoPointer(), FirstStateVar, varArgs...)

}

// Updates an array of accessible states.
//
// This function should be called by `GtkWidget` types whenever an accessible
// state change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *Popover) UpdateStateValue(NStatesVar int, StatesVar uintptr, ValuesVar uintptr) {

	XGtkAccessibleUpdateStateValue(x.GoPointer(), NStatesVar, StatesVar, ValuesVar)

}

// Gets the ID of the @buildable object.
//
// `GtkBuilder` sets the name based on the ID attribute
// of the &lt;object&gt; tag used to construct the @buildable.
func (x *Popover) GetBuildableId() string {

	cret := XGtkBuildableGetBuildableId(x.GoPointer())
	return cret
}

// Returns the renderer that is used for this `GtkNative`.
func (x *Popover) GetRenderer() *gsk.Renderer {
	var cls *gsk.Renderer

	cret := XGtkNativeGetRenderer(x.GoPointer())

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &gsk.Renderer{}
	cls.Ptr = cret
	return cls
}

// Returns the surface of this `GtkNative`.
func (x *Popover) GetSurface() *gdk.Surface {
	var cls *gdk.Surface

	cret := XGtkNativeGetSurface(x.GoPointer())

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &gdk.Surface{}
	cls.Ptr = cret
	return cls
}

// Retrieves the surface transform of @self.
//
// This is the translation from @self's surface coordinates into
// @self's widget coordinates.
func (x *Popover) GetSurfaceTransform(XVar float64, YVar float64) {

	XGtkNativeGetSurfaceTransform(x.GoPointer(), XVar, YVar)

}

// Realizes a `GtkNative`.
//
// This should only be used by subclasses.
func (x *Popover) Realize() {

	XGtkNativeRealize(x.GoPointer())

}

// Unrealizes a `GtkNative`.
//
// This should only be used by subclasses.
func (x *Popover) Unrealize() {

	XGtkNativeUnrealize(x.GoPointer())

}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GTK"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xNewPopover, lib, "gtk_popover_new")

	core.PuregoSafeRegister(&xPopoverGetAutohide, lib, "gtk_popover_get_autohide")
	core.PuregoSafeRegister(&xPopoverGetCascadePopdown, lib, "gtk_popover_get_cascade_popdown")
	core.PuregoSafeRegister(&xPopoverGetChild, lib, "gtk_popover_get_child")
	core.PuregoSafeRegister(&xPopoverGetHasArrow, lib, "gtk_popover_get_has_arrow")
	core.PuregoSafeRegister(&xPopoverGetMnemonicsVisible, lib, "gtk_popover_get_mnemonics_visible")
	core.PuregoSafeRegister(&xPopoverGetOffset, lib, "gtk_popover_get_offset")
	core.PuregoSafeRegister(&xPopoverGetPointingTo, lib, "gtk_popover_get_pointing_to")
	core.PuregoSafeRegister(&xPopoverGetPosition, lib, "gtk_popover_get_position")
	core.PuregoSafeRegister(&xPopoverPopdown, lib, "gtk_popover_popdown")
	core.PuregoSafeRegister(&xPopoverPopup, lib, "gtk_popover_popup")
	core.PuregoSafeRegister(&xPopoverPresent, lib, "gtk_popover_present")
	core.PuregoSafeRegister(&xPopoverSetAutohide, lib, "gtk_popover_set_autohide")
	core.PuregoSafeRegister(&xPopoverSetCascadePopdown, lib, "gtk_popover_set_cascade_popdown")
	core.PuregoSafeRegister(&xPopoverSetChild, lib, "gtk_popover_set_child")
	core.PuregoSafeRegister(&xPopoverSetDefaultWidget, lib, "gtk_popover_set_default_widget")
	core.PuregoSafeRegister(&xPopoverSetHasArrow, lib, "gtk_popover_set_has_arrow")
	core.PuregoSafeRegister(&xPopoverSetMnemonicsVisible, lib, "gtk_popover_set_mnemonics_visible")
	core.PuregoSafeRegister(&xPopoverSetOffset, lib, "gtk_popover_set_offset")
	core.PuregoSafeRegister(&xPopoverSetPointingTo, lib, "gtk_popover_set_pointing_to")
	core.PuregoSafeRegister(&xPopoverSetPosition, lib, "gtk_popover_set_position")

}
