// Package gtk was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gtk

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/gdk"
	"github.com/jwijenbergh/puregotk/v4/gobject"
	"github.com/jwijenbergh/puregotk/v4/gsk"
)

type EmojiChooserClass struct {
}

// The `GtkEmojiChooser` is used by text widgets such as `GtkEntry` or
// `GtkTextView` to let users insert Emoji characters.
//
// ![An example GtkEmojiChooser](emojichooser.png)
//
// `GtkEmojiChooser` emits the [signal@Gtk.EmojiChooser::emoji-picked]
// signal when an Emoji is selected.
//
// # CSS nodes
//
// ```
// popover
// ├── box.emoji-searchbar
// │   ╰── entry.search
// ╰── box.emoji-toolbar
//
//	├── button.image-button.emoji-section
//	├── ...
//	╰── button.image-button.emoji-section
//
// ```
//
// Every `GtkEmojiChooser` consists of a main node called popover.
// The contents of the popover are largely implementation defined
// and supposed to inherit general styles.
// The top searchbar used to search emoji and gets the .emoji-searchbar
// style class itself.
// The bottom toolbar used to switch between different emoji categories
// consists of buttons with the .emoji-section style class and gets the
// .emoji-toolbar style class itself.
type EmojiChooser struct {
	Popover
}

func EmojiChooserNewFromInternalPtr(ptr uintptr) *EmojiChooser {
	cls := &EmojiChooser{}
	cls.Ptr = ptr
	return cls
}

var xNewEmojiChooser func() uintptr

// Creates a new `GtkEmojiChooser`.
func NewEmojiChooser() *Widget {
	NewEmojiChooserPtr := xNewEmojiChooser()
	if NewEmojiChooserPtr == 0 {
		return nil
	}

	gobject.IncreaseRef(NewEmojiChooserPtr)

	NewEmojiChooserCls := &Widget{}
	NewEmojiChooserCls.Ptr = NewEmojiChooserPtr
	return NewEmojiChooserCls
}

func (c *EmojiChooser) GoPointer() uintptr {
	return c.Ptr
}

func (c *EmojiChooser) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// Emitted when the user selects an Emoji.
func (x *EmojiChooser) ConnectEmojiPicked(cb func(EmojiChooser, string)) {
	fcb := func(clsPtr uintptr, TextVarp string) {
		fa := EmojiChooser{}
		fa.Ptr = clsPtr

		cb(fa, TextVarp)

	}
	gobject.ObjectConnect(x.GoPointer(), "signal::emoji-picked", purego.NewCallback(fcb))
}

// Retrieves the `GtkAccessibleRole` for the given `GtkAccessible`.
func (x *EmojiChooser) GetAccessibleRole() AccessibleRole {

	return XGtkAccessibleGetAccessibleRole(x.GoPointer())

}

// Resets the accessible @property to its default value.
func (x *EmojiChooser) ResetProperty(PropertyVar AccessibleProperty) {

	XGtkAccessibleResetProperty(x.GoPointer(), PropertyVar)

}

// Resets the accessible @relation to its default value.
func (x *EmojiChooser) ResetRelation(RelationVar AccessibleRelation) {

	XGtkAccessibleResetRelation(x.GoPointer(), RelationVar)

}

// Resets the accessible @state to its default value.
func (x *EmojiChooser) ResetState(StateVar AccessibleState) {

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
func (x *EmojiChooser) UpdateProperty(FirstPropertyVar AccessibleProperty, varArgs ...interface{}) {

	XGtkAccessibleUpdateProperty(x.GoPointer(), FirstPropertyVar, varArgs...)

}

// Updates an array of accessible properties.
//
// This function should be called by `GtkWidget` types whenever an accessible
// property change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *EmojiChooser) UpdatePropertyValue(NPropertiesVar int32, PropertiesVar uintptr, ValuesVar uintptr) {

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
func (x *EmojiChooser) UpdateRelation(FirstRelationVar AccessibleRelation, varArgs ...interface{}) {

	XGtkAccessibleUpdateRelation(x.GoPointer(), FirstRelationVar, varArgs...)

}

// Updates an array of accessible relations.
//
// This function should be called by `GtkWidget` types whenever an accessible
// relation change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *EmojiChooser) UpdateRelationValue(NRelationsVar int32, RelationsVar uintptr, ValuesVar uintptr) {

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
func (x *EmojiChooser) UpdateState(FirstStateVar AccessibleState, varArgs ...interface{}) {

	XGtkAccessibleUpdateState(x.GoPointer(), FirstStateVar, varArgs...)

}

// Updates an array of accessible states.
//
// This function should be called by `GtkWidget` types whenever an accessible
// state change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *EmojiChooser) UpdateStateValue(NStatesVar int32, StatesVar uintptr, ValuesVar uintptr) {

	XGtkAccessibleUpdateStateValue(x.GoPointer(), NStatesVar, StatesVar, ValuesVar)

}

// Gets the ID of the @buildable object.
//
// `GtkBuilder` sets the name based on the ID attribute
// of the &lt;object&gt; tag used to construct the @buildable.
func (x *EmojiChooser) GetBuildableId() string {

	return XGtkBuildableGetBuildableId(x.GoPointer())

}

// Returns the renderer that is used for this `GtkNative`.
func (x *EmojiChooser) GetRenderer() *gsk.Renderer {

	GetRendererPtr := XGtkNativeGetRenderer(x.GoPointer())
	if GetRendererPtr == 0 {
		return nil
	}

	gobject.IncreaseRef(GetRendererPtr)

	GetRendererCls := &gsk.Renderer{}
	GetRendererCls.Ptr = GetRendererPtr
	return GetRendererCls

}

// Returns the surface of this `GtkNative`.
func (x *EmojiChooser) GetSurface() *gdk.Surface {

	GetSurfacePtr := XGtkNativeGetSurface(x.GoPointer())
	if GetSurfacePtr == 0 {
		return nil
	}

	gobject.IncreaseRef(GetSurfacePtr)

	GetSurfaceCls := &gdk.Surface{}
	GetSurfaceCls.Ptr = GetSurfacePtr
	return GetSurfaceCls

}

// Retrieves the surface transform of @self.
//
// This is the translation from @self's surface coordinates into
// @self's widget coordinates.
func (x *EmojiChooser) GetSurfaceTransform(XVar float64, YVar float64) {

	XGtkNativeGetSurfaceTransform(x.GoPointer(), XVar, YVar)

}

// Realizes a `GtkNative`.
//
// This should only be used by subclasses.
func (x *EmojiChooser) Realize() {

	XGtkNativeRealize(x.GoPointer())

}

// Unrealizes a `GtkNative`.
//
// This should only be used by subclasses.
func (x *EmojiChooser) Unrealize() {

	XGtkNativeUnrealize(x.GoPointer())

}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GTK"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xNewEmojiChooser, lib, "gtk_emoji_chooser_new")

}
