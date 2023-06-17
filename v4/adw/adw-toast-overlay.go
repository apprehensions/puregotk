// Package adw was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package adw

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/gobject"
	"github.com/jwijenbergh/puregotk/v4/gtk"
)

type ToastOverlayClass struct {
	ParentClass uintptr
}

// A widget showing toasts above its content.
//
// &lt;picture&gt;
//
//	&lt;source srcset="toast-overlay-dark.png" media="(prefers-color-scheme: dark)"&gt;
//	&lt;img src="toast-overlay.png" alt="toast-overlay"&gt;
//
// &lt;/picture&gt;
//
// Much like [class@Gtk.Overlay], `AdwToastOverlay` is a container with a single
// main child, on top of which it can display a [class@Toast], overlaid.
// Toasts can be shown with [method@ToastOverlay.add_toast].
//
// See [class@Toast] for details.
//
// ## CSS nodes
//
// ```
// toastoverlay
// ├── [child]
// ├── toast
// ┊   ├── widget
// ┊   │   ├── [label.heading]
//
//	│   ╰── [custom title]
//	├── [button]
//	╰── button.circular.flat
//
// ```
//
// `AdwToastOverlay`'s CSS node is called `toastoverlay`. It contains the child,
// as well as zero or more `toast` subnodes.
//
// Each of the `toast` nodes contains a `widget` subnode, optionally a `button`
// subnode, and another `button` subnode with `.circular` and `.flat` style
// classes.
//
// The `widget` subnode contains a `label` subnode with the `.heading` style
// class, or a custom widget provided by the application.
//
// ## Accessibility
//
// `AdwToastOverlay` uses the `GTK_ACCESSIBLE_ROLE_TAB_GROUP` role.
type ToastOverlay struct {
	gtk.Widget
}

func ToastOverlayNewFromInternalPtr(ptr uintptr) *ToastOverlay {
	cls := &ToastOverlay{}
	cls.Ptr = ptr
	return cls
}

var xNewToastOverlay func() uintptr

// Creates a new `AdwToastOverlay`.
func NewToastOverlay() *gtk.Widget {
	NewToastOverlayPtr := xNewToastOverlay()
	if NewToastOverlayPtr == 0 {
		return nil
	}

	gobject.IncreaseRef(NewToastOverlayPtr)

	NewToastOverlayCls := &gtk.Widget{}
	NewToastOverlayCls.Ptr = NewToastOverlayPtr
	return NewToastOverlayCls
}

var xToastOverlayAddToast func(uintptr, uintptr)

// Displays @toast.
//
// Only one toast can be shown at a time; if a toast is already being displayed,
// either @toast or the original toast will be placed in a queue, depending on
// the priority of @toast. See [property@Toast:priority].
//
// If called on a toast that's already displayed, its timeout will be reset.
//
// If called on a toast currently in the queue, the toast will be bumped
// forward to be shown as soon as possible.
func (x *ToastOverlay) AddToast(ToastVar *Toast) {

	xToastOverlayAddToast(x.GoPointer(), ToastVar.GoPointer())

}

var xToastOverlayGetChild func(uintptr) uintptr

// Gets the child widget of @self.
func (x *ToastOverlay) GetChild() *gtk.Widget {

	GetChildPtr := xToastOverlayGetChild(x.GoPointer())
	if GetChildPtr == 0 {
		return nil
	}

	gobject.IncreaseRef(GetChildPtr)

	GetChildCls := &gtk.Widget{}
	GetChildCls.Ptr = GetChildPtr
	return GetChildCls

}

var xToastOverlaySetChild func(uintptr, uintptr)

// Sets the child widget of @self.
func (x *ToastOverlay) SetChild(ChildVar *gtk.Widget) {

	xToastOverlaySetChild(x.GoPointer(), ChildVar.GoPointer())

}

func (c *ToastOverlay) GoPointer() uintptr {
	return c.Ptr
}

func (c *ToastOverlay) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// Retrieves the `GtkAccessibleRole` for the given `GtkAccessible`.
func (x *ToastOverlay) GetAccessibleRole() gtk.AccessibleRole {

	return gtk.XGtkAccessibleGetAccessibleRole(x.GoPointer())

}

// Resets the accessible @property to its default value.
func (x *ToastOverlay) ResetProperty(PropertyVar gtk.AccessibleProperty) {

	gtk.XGtkAccessibleResetProperty(x.GoPointer(), PropertyVar)

}

// Resets the accessible @relation to its default value.
func (x *ToastOverlay) ResetRelation(RelationVar gtk.AccessibleRelation) {

	gtk.XGtkAccessibleResetRelation(x.GoPointer(), RelationVar)

}

// Resets the accessible @state to its default value.
func (x *ToastOverlay) ResetState(StateVar gtk.AccessibleState) {

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
func (x *ToastOverlay) UpdateProperty(FirstPropertyVar gtk.AccessibleProperty, varArgs ...interface{}) {

	gtk.XGtkAccessibleUpdateProperty(x.GoPointer(), FirstPropertyVar, varArgs...)

}

// Updates an array of accessible properties.
//
// This function should be called by `GtkWidget` types whenever an accessible
// property change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *ToastOverlay) UpdatePropertyValue(NPropertiesVar int32, PropertiesVar uintptr, ValuesVar uintptr) {

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
func (x *ToastOverlay) UpdateRelation(FirstRelationVar gtk.AccessibleRelation, varArgs ...interface{}) {

	gtk.XGtkAccessibleUpdateRelation(x.GoPointer(), FirstRelationVar, varArgs...)

}

// Updates an array of accessible relations.
//
// This function should be called by `GtkWidget` types whenever an accessible
// relation change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *ToastOverlay) UpdateRelationValue(NRelationsVar int32, RelationsVar uintptr, ValuesVar uintptr) {

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
func (x *ToastOverlay) UpdateState(FirstStateVar gtk.AccessibleState, varArgs ...interface{}) {

	gtk.XGtkAccessibleUpdateState(x.GoPointer(), FirstStateVar, varArgs...)

}

// Updates an array of accessible states.
//
// This function should be called by `GtkWidget` types whenever an accessible
// state change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *ToastOverlay) UpdateStateValue(NStatesVar int32, StatesVar uintptr, ValuesVar uintptr) {

	gtk.XGtkAccessibleUpdateStateValue(x.GoPointer(), NStatesVar, StatesVar, ValuesVar)

}

// Gets the ID of the @buildable object.
//
// `GtkBuilder` sets the name based on the ID attribute
// of the &lt;object&gt; tag used to construct the @buildable.
func (x *ToastOverlay) GetBuildableId() string {

	return gtk.XGtkBuildableGetBuildableId(x.GoPointer())

}

func init() {
	lib, err := purego.Dlopen(core.GetPath("ADW"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xNewToastOverlay, lib, "adw_toast_overlay_new")

	core.PuregoSafeRegister(&xToastOverlayAddToast, lib, "adw_toast_overlay_add_toast")
	core.PuregoSafeRegister(&xToastOverlayGetChild, lib, "adw_toast_overlay_get_child")
	core.PuregoSafeRegister(&xToastOverlaySetChild, lib, "adw_toast_overlay_set_child")

}
