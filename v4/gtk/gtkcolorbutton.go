// Package gtk was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gtk

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/gdk"
	"github.com/jwijenbergh/puregotk/v4/gobject"
)

// The `GtkColorButton` allows to open a color chooser dialog to change
// the color.
//
// ![An example GtkColorButton](color-button.png)
//
// It is suitable widget for selecting a color in a preference dialog.
//
// # CSS nodes
//
// ```
// colorbutton
// ╰── button.color
//
//	╰── [content]
//
// ```
//
// `GtkColorButton` has a single CSS node with name colorbutton which
// contains a button node. To differentiate it from a plain `GtkButton`,
// it gets the .color style class.
type ColorButton struct {
	Widget
}

func ColorButtonNewFromInternalPtr(ptr uintptr) *ColorButton {
	cls := &ColorButton{}
	cls.Ptr = ptr
	return cls
}

var xNewColorButton func() uintptr

// Creates a new color button.
//
// This returns a widget in the form of a small button containing
// a swatch representing the current selected color. When the button
// is clicked, a color chooser dialog will open, allowing the user
// to select a color. The swatch will be updated to reflect the new
// color when the user finishes.
func NewColorButton() *Widget {
	NewColorButtonPtr := xNewColorButton()
	if NewColorButtonPtr == 0 {
		return nil
	}

	gobject.IncreaseRef(NewColorButtonPtr)

	NewColorButtonCls := &Widget{}
	NewColorButtonCls.Ptr = NewColorButtonPtr
	return NewColorButtonCls
}

var xNewWithRgbaColorButton func(*gdk.RGBA) uintptr

// Creates a new color button showing the given color.
func NewWithRgbaColorButton(RgbaVar *gdk.RGBA) *Widget {
	NewWithRgbaColorButtonPtr := xNewWithRgbaColorButton(RgbaVar)
	if NewWithRgbaColorButtonPtr == 0 {
		return nil
	}

	gobject.IncreaseRef(NewWithRgbaColorButtonPtr)

	NewWithRgbaColorButtonCls := &Widget{}
	NewWithRgbaColorButtonCls.Ptr = NewWithRgbaColorButtonPtr
	return NewWithRgbaColorButtonCls
}

var xColorButtonGetModal func(uintptr) bool

// Gets whether the dialog is modal.
func (x *ColorButton) GetModal() bool {

	return xColorButtonGetModal(x.GoPointer())

}

var xColorButtonGetTitle func(uintptr) string

// Gets the title of the color chooser dialog.
func (x *ColorButton) GetTitle() string {

	return xColorButtonGetTitle(x.GoPointer())

}

var xColorButtonSetModal func(uintptr, bool)

// Sets whether the dialog should be modal.
func (x *ColorButton) SetModal(ModalVar bool) {

	xColorButtonSetModal(x.GoPointer(), ModalVar)

}

var xColorButtonSetTitle func(uintptr, string)

// Sets the title for the color chooser dialog.
func (x *ColorButton) SetTitle(TitleVar string) {

	xColorButtonSetTitle(x.GoPointer(), TitleVar)

}

func (c *ColorButton) GoPointer() uintptr {
	return c.Ptr
}

func (c *ColorButton) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// Emitted to when the color button is activated.
//
// The `::activate` signal on `GtkMenuButton` is an action signal and
// emitting it causes the button to pop up its dialog.
func (x *ColorButton) ConnectActivate(cb func(ColorButton)) {
	fcb := func(clsPtr uintptr) {
		fa := ColorButton{}
		fa.Ptr = clsPtr

		cb(fa)

	}
	gobject.ObjectConnect(x.GoPointer(), "signal::activate", purego.NewCallback(fcb))
}

// Emitted when the user selects a color.
//
// When handling this signal, use [method@Gtk.ColorChooser.get_rgba]
// to find out which color was just selected.
//
// Note that this signal is only emitted when the user changes the color.
// If you need to react to programmatic color changes as well, use
// the notify::rgba signal.
func (x *ColorButton) ConnectColorSet(cb func(ColorButton)) {
	fcb := func(clsPtr uintptr) {
		fa := ColorButton{}
		fa.Ptr = clsPtr

		cb(fa)

	}
	gobject.ObjectConnect(x.GoPointer(), "signal::color-set", purego.NewCallback(fcb))
}

// Retrieves the `GtkAccessibleRole` for the given `GtkAccessible`.
func (x *ColorButton) GetAccessibleRole() AccessibleRole {

	return XGtkAccessibleGetAccessibleRole(x.GoPointer())

}

// Resets the accessible @property to its default value.
func (x *ColorButton) ResetProperty(PropertyVar AccessibleProperty) {

	XGtkAccessibleResetProperty(x.GoPointer(), PropertyVar)

}

// Resets the accessible @relation to its default value.
func (x *ColorButton) ResetRelation(RelationVar AccessibleRelation) {

	XGtkAccessibleResetRelation(x.GoPointer(), RelationVar)

}

// Resets the accessible @state to its default value.
func (x *ColorButton) ResetState(StateVar AccessibleState) {

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
func (x *ColorButton) UpdateProperty(FirstPropertyVar AccessibleProperty, varArgs ...interface{}) {

	XGtkAccessibleUpdateProperty(x.GoPointer(), FirstPropertyVar, varArgs...)

}

// Updates an array of accessible properties.
//
// This function should be called by `GtkWidget` types whenever an accessible
// property change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *ColorButton) UpdatePropertyValue(NPropertiesVar int32, PropertiesVar uintptr, ValuesVar uintptr) {

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
func (x *ColorButton) UpdateRelation(FirstRelationVar AccessibleRelation, varArgs ...interface{}) {

	XGtkAccessibleUpdateRelation(x.GoPointer(), FirstRelationVar, varArgs...)

}

// Updates an array of accessible relations.
//
// This function should be called by `GtkWidget` types whenever an accessible
// relation change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *ColorButton) UpdateRelationValue(NRelationsVar int32, RelationsVar uintptr, ValuesVar uintptr) {

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
func (x *ColorButton) UpdateState(FirstStateVar AccessibleState, varArgs ...interface{}) {

	XGtkAccessibleUpdateState(x.GoPointer(), FirstStateVar, varArgs...)

}

// Updates an array of accessible states.
//
// This function should be called by `GtkWidget` types whenever an accessible
// state change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *ColorButton) UpdateStateValue(NStatesVar int32, StatesVar uintptr, ValuesVar uintptr) {

	XGtkAccessibleUpdateStateValue(x.GoPointer(), NStatesVar, StatesVar, ValuesVar)

}

// Gets the ID of the @buildable object.
//
// `GtkBuilder` sets the name based on the ID attribute
// of the &lt;object&gt; tag used to construct the @buildable.
func (x *ColorButton) GetBuildableId() string {

	return XGtkBuildableGetBuildableId(x.GoPointer())

}

// Adds a palette to the color chooser.
//
// If @orientation is horizontal, the colors are grouped in rows,
// with @colors_per_line colors in each row. If @horizontal is %FALSE,
// the colors are grouped in columns instead.
//
// The default color palette of [class@Gtk.ColorChooserWidget] has
// 45 colors, organized in columns of 5 colors (this includes some
// grays).
//
// The layout of the color chooser widget works best when the
// palettes have 9-10 columns.
//
// Calling this function for the first time has the side effect
// of removing the default color palette from the color chooser.
//
// If @colors is %NULL, removes all previously added palettes.
func (x *ColorButton) AddPalette(OrientationVar Orientation, ColorsPerLineVar int32, NColorsVar int32, ColorsVar uintptr) {

	XGtkColorChooserAddPalette(x.GoPointer(), OrientationVar, ColorsPerLineVar, NColorsVar, ColorsVar)

}

// Gets the currently-selected color.
func (x *ColorButton) GetRgba(ColorVar *gdk.RGBA) {

	XGtkColorChooserGetRgba(x.GoPointer(), ColorVar)

}

// Returns whether the color chooser shows the alpha channel.
func (x *ColorButton) GetUseAlpha() bool {

	return XGtkColorChooserGetUseAlpha(x.GoPointer())

}

// Sets the color.
func (x *ColorButton) SetRgba(ColorVar *gdk.RGBA) {

	XGtkColorChooserSetRgba(x.GoPointer(), ColorVar)

}

// Sets whether or not the color chooser should use the alpha channel.
func (x *ColorButton) SetUseAlpha(UseAlphaVar bool) {

	XGtkColorChooserSetUseAlpha(x.GoPointer(), UseAlphaVar)

}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GTK"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xNewColorButton, lib, "gtk_color_button_new")
	core.PuregoSafeRegister(&xNewWithRgbaColorButton, lib, "gtk_color_button_new_with_rgba")

	core.PuregoSafeRegister(&xColorButtonGetModal, lib, "gtk_color_button_get_modal")
	core.PuregoSafeRegister(&xColorButtonGetTitle, lib, "gtk_color_button_get_title")
	core.PuregoSafeRegister(&xColorButtonSetModal, lib, "gtk_color_button_set_modal")
	core.PuregoSafeRegister(&xColorButtonSetTitle, lib, "gtk_color_button_set_title")

}
