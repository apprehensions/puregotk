// Package gtk was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gtk

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/gobject"
)

// `GtkInfoBar` can be show messages to the user without a dialog.
//
// ![An example GtkInfoBar](info-bar.png)
//
// It is often temporarily shown at the top or bottom of a document.
// In contrast to [class@Gtk.Dialog], which has an action area at the
// bottom, `GtkInfoBar` has an action area at the side.
//
// The API of `GtkInfoBar` is very similar to `GtkDialog`, allowing you
// to add buttons to the action area with [method@Gtk.InfoBar.add_button]
// or [ctor@Gtk.InfoBar.new_with_buttons]. The sensitivity of action widgets
// can be controlled with [method@Gtk.InfoBar.set_response_sensitive].
//
// To add widgets to the main content area of a `GtkInfoBar`, use
// [method@Gtk.InfoBar.add_child].
//
// Similar to [class@Gtk.MessageDialog], the contents of a `GtkInfoBar`
// can by classified as error message, warning, informational message, etc,
// by using [method@Gtk.InfoBar.set_message_type]. GTK may use the message
// type to determine how the message is displayed.
//
// A simple example for using a `GtkInfoBar`:
// ```c
// GtkWidget *message_label;
// GtkWidget *widget;
// GtkWidget *grid;
// GtkInfoBar *bar;
//
// // set up info bar
// widget = gtk_info_bar_new ();
// bar = GTK_INFO_BAR (widget);
// grid = gtk_grid_new ();
//
// message_label = gtk_label_new ("");
// gtk_info_bar_add_child (bar, message_label);
// gtk_info_bar_add_button (bar,
//
//	_("_OK"),
//	GTK_RESPONSE_OK);
//
// g_signal_connect (bar,
//
//	"response",
//	G_CALLBACK (gtk_widget_hide),
//	NULL);
//
// gtk_grid_attach (GTK_GRID (grid),
//
//	widget,
//	0, 2, 1, 1);
//
// // ...
//
// // show an error message
// gtk_label_set_text (GTK_LABEL (message_label), "An error occurred!");
// gtk_info_bar_set_message_type (bar, GTK_MESSAGE_ERROR);
// gtk_widget_show (bar);
// ```
//
// # GtkInfoBar as GtkBuildable
//
// `GtkInfoBar` supports a custom `&lt;action-widgets&gt;` element, which can contain
// multiple `&lt;action-widget&gt;` elements. The “response” attribute specifies a
// numeric response, and the content of the element is the id of widget
// (which should be a child of the dialogs @action_area).
//
// `GtkInfoBar` supports adding action widgets by specifying “action” as
// the “type” attribute of a `&lt;child&gt;` element. The widget will be added
// either to the action area. The response id has to be associated
// with the action widget using the `&lt;action-widgets&gt;` element.
//
// # CSS nodes
//
// `GtkInfoBar` has a single CSS node with name infobar. The node may get
// one of the style classes .info, .warning, .error or .question, depending
// on the message type.
// If the info bar shows a close button, that button will have the .close
// style class applied.
type InfoBar struct {
	Widget
}

func InfoBarNewFromInternalPtr(ptr uintptr) *InfoBar {
	cls := &InfoBar{}
	cls.Ptr = ptr
	return cls
}

var xNewInfoBar func() uintptr

// Creates a new `GtkInfoBar` object.
func NewInfoBar() *InfoBar {
	var cls *InfoBar

	cret := xNewInfoBar()

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &InfoBar{}
	cls.Ptr = cret
	return cls
}

var xNewWithButtonsInfoBar func(string, ...interface{}) uintptr

// Creates a new `GtkInfoBar` with buttons.
//
// Button text/response ID pairs should be listed, with a %NULL pointer
// ending the list. A response ID can be any positive number,
// or one of the values in the `GtkResponseType` enumeration. If the
// user clicks one of these dialog buttons, GtkInfoBar will emit
// the [signal@Gtk.InfoBar::response] signal with the corresponding
// response ID.
func NewWithButtonsInfoBar(FirstButtonTextVar string, varArgs ...interface{}) *InfoBar {
	var cls *InfoBar

	cret := xNewWithButtonsInfoBar(FirstButtonTextVar, varArgs...)

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &InfoBar{}
	cls.Ptr = cret
	return cls
}

var xInfoBarAddActionWidget func(uintptr, uintptr, int)

// Add an activatable widget to the action area of a `GtkInfoBar`.
//
// This also connects a signal handler that will emit the
// [signal@Gtk.InfoBar::response] signal on the message area
// when the widget is activated. The widget is appended to the
// end of the message areas action area.
func (x *InfoBar) AddActionWidget(ChildVar *Widget, ResponseIdVar int) {

	xInfoBarAddActionWidget(x.GoPointer(), ChildVar.GoPointer(), ResponseIdVar)

}

var xInfoBarAddButton func(uintptr, string, int) uintptr

// Adds a button with the given text.
//
// Clicking the button will emit the [signal@Gtk.InfoBar::response]
// signal with the given response_id. The button is appended to the
// end of the info bars's action area. The button widget is returned,
// but usually you don't need it.
func (x *InfoBar) AddButton(ButtonTextVar string, ResponseIdVar int) *Button {
	var cls *Button

	cret := xInfoBarAddButton(x.GoPointer(), ButtonTextVar, ResponseIdVar)

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &Button{}
	cls.Ptr = cret
	return cls
}

var xInfoBarAddButtons func(uintptr, string, ...interface{})

// Adds multiple buttons.
//
// This is the same as calling [method@Gtk.InfoBar.add_button]
// repeatedly. The variable argument list should be %NULL-terminated
// as with [ctor@Gtk.InfoBar.new_with_buttons]. Each button must have both
// text and response ID.
func (x *InfoBar) AddButtons(FirstButtonTextVar string, varArgs ...interface{}) {

	xInfoBarAddButtons(x.GoPointer(), FirstButtonTextVar, varArgs...)

}

var xInfoBarAddChild func(uintptr, uintptr)

// Adds a widget to the content area of the info bar.
func (x *InfoBar) AddChild(WidgetVar *Widget) {

	xInfoBarAddChild(x.GoPointer(), WidgetVar.GoPointer())

}

var xInfoBarGetMessageType func(uintptr) MessageType

// Returns the message type of the message area.
func (x *InfoBar) GetMessageType() MessageType {

	cret := xInfoBarGetMessageType(x.GoPointer())
	return cret
}

var xInfoBarGetRevealed func(uintptr) bool

// Returns whether the info bar is currently revealed.
func (x *InfoBar) GetRevealed() bool {

	cret := xInfoBarGetRevealed(x.GoPointer())
	return cret
}

var xInfoBarGetShowCloseButton func(uintptr) bool

// Returns whether the widget will display a standard close button.
func (x *InfoBar) GetShowCloseButton() bool {

	cret := xInfoBarGetShowCloseButton(x.GoPointer())
	return cret
}

var xInfoBarRemoveActionWidget func(uintptr, uintptr)

// Removes a widget from the action area of @info_bar.
//
// The widget must have been put there by a call to
// [method@Gtk.InfoBar.add_action_widget] or [method@Gtk.InfoBar.add_button].
func (x *InfoBar) RemoveActionWidget(WidgetVar *Widget) {

	xInfoBarRemoveActionWidget(x.GoPointer(), WidgetVar.GoPointer())

}

var xInfoBarRemoveChild func(uintptr, uintptr)

// Removes a widget from the content area of the info bar.
func (x *InfoBar) RemoveChild(WidgetVar *Widget) {

	xInfoBarRemoveChild(x.GoPointer(), WidgetVar.GoPointer())

}

var xInfoBarResponse func(uintptr, int)

// Emits the “response” signal with the given @response_id.
func (x *InfoBar) Response(ResponseIdVar int) {

	xInfoBarResponse(x.GoPointer(), ResponseIdVar)

}

var xInfoBarSetDefaultResponse func(uintptr, int)

// Sets the last widget in the info bar’s action area with
// the given response_id as the default widget for the dialog.
//
// Pressing “Enter” normally activates the default widget.
//
// Note that this function currently requires @info_bar to
// be added to a widget hierarchy.
func (x *InfoBar) SetDefaultResponse(ResponseIdVar int) {

	xInfoBarSetDefaultResponse(x.GoPointer(), ResponseIdVar)

}

var xInfoBarSetMessageType func(uintptr, MessageType)

// Sets the message type of the message area.
//
// GTK uses this type to determine how the message is displayed.
func (x *InfoBar) SetMessageType(MessageTypeVar MessageType) {

	xInfoBarSetMessageType(x.GoPointer(), MessageTypeVar)

}

var xInfoBarSetResponseSensitive func(uintptr, int, bool)

// Sets the sensitivity of action widgets for @response_id.
//
// Calls `gtk_widget_set_sensitive (widget, setting)` for each
// widget in the info bars’s action area with the given @response_id.
// A convenient way to sensitize/desensitize buttons.
func (x *InfoBar) SetResponseSensitive(ResponseIdVar int, SettingVar bool) {

	xInfoBarSetResponseSensitive(x.GoPointer(), ResponseIdVar, SettingVar)

}

var xInfoBarSetRevealed func(uintptr, bool)

// Sets whether the `GtkInfoBar` is revealed.
//
// Changing this will make @info_bar reveal or conceal
// itself via a sliding transition.
//
// Note: this does not show or hide @info_bar in the
// [property@Gtk.Widget:visible] sense, so revealing has no effect
// if [property@Gtk.Widget:visible] is %FALSE.
func (x *InfoBar) SetRevealed(RevealedVar bool) {

	xInfoBarSetRevealed(x.GoPointer(), RevealedVar)

}

var xInfoBarSetShowCloseButton func(uintptr, bool)

// If true, a standard close button is shown.
//
// When clicked it emits the response %GTK_RESPONSE_CLOSE.
func (x *InfoBar) SetShowCloseButton(SettingVar bool) {

	xInfoBarSetShowCloseButton(x.GoPointer(), SettingVar)

}

func (c *InfoBar) GoPointer() uintptr {
	return c.Ptr
}

func (c *InfoBar) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// Gets emitted when the user uses a keybinding to dismiss the info bar.
//
// The ::close signal is a [keybinding signal](class.SignalAction.html).
//
// The default binding for this signal is the Escape key.
func (x *InfoBar) ConnectClose(cb func(InfoBar)) uint32 {
	fcb := func(clsPtr uintptr) {
		fa := InfoBar{}
		fa.Ptr = clsPtr

		cb(fa)

	}
	return gobject.SignalConnect(x.GoPointer(), "close", purego.NewCallback(fcb))
}

// Emitted when an action widget is clicked.
//
// The signal is also emitted when the application programmer
// calls [method@Gtk.InfoBar.response]. The @response_id depends
// on which action widget was clicked.
func (x *InfoBar) ConnectResponse(cb func(InfoBar, int)) uint32 {
	fcb := func(clsPtr uintptr, ResponseIdVarp int) {
		fa := InfoBar{}
		fa.Ptr = clsPtr

		cb(fa, ResponseIdVarp)

	}
	return gobject.SignalConnect(x.GoPointer(), "response", purego.NewCallback(fcb))
}

// Retrieves the `GtkAccessibleRole` for the given `GtkAccessible`.
func (x *InfoBar) GetAccessibleRole() AccessibleRole {

	cret := XGtkAccessibleGetAccessibleRole(x.GoPointer())
	return cret
}

// Resets the accessible @property to its default value.
func (x *InfoBar) ResetProperty(PropertyVar AccessibleProperty) {

	XGtkAccessibleResetProperty(x.GoPointer(), PropertyVar)

}

// Resets the accessible @relation to its default value.
func (x *InfoBar) ResetRelation(RelationVar AccessibleRelation) {

	XGtkAccessibleResetRelation(x.GoPointer(), RelationVar)

}

// Resets the accessible @state to its default value.
func (x *InfoBar) ResetState(StateVar AccessibleState) {

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
func (x *InfoBar) UpdateProperty(FirstPropertyVar AccessibleProperty, varArgs ...interface{}) {

	XGtkAccessibleUpdateProperty(x.GoPointer(), FirstPropertyVar, varArgs...)

}

// Updates an array of accessible properties.
//
// This function should be called by `GtkWidget` types whenever an accessible
// property change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *InfoBar) UpdatePropertyValue(NPropertiesVar int, PropertiesVar uintptr, ValuesVar uintptr) {

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
func (x *InfoBar) UpdateRelation(FirstRelationVar AccessibleRelation, varArgs ...interface{}) {

	XGtkAccessibleUpdateRelation(x.GoPointer(), FirstRelationVar, varArgs...)

}

// Updates an array of accessible relations.
//
// This function should be called by `GtkWidget` types whenever an accessible
// relation change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *InfoBar) UpdateRelationValue(NRelationsVar int, RelationsVar uintptr, ValuesVar uintptr) {

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
func (x *InfoBar) UpdateState(FirstStateVar AccessibleState, varArgs ...interface{}) {

	XGtkAccessibleUpdateState(x.GoPointer(), FirstStateVar, varArgs...)

}

// Updates an array of accessible states.
//
// This function should be called by `GtkWidget` types whenever an accessible
// state change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *InfoBar) UpdateStateValue(NStatesVar int, StatesVar uintptr, ValuesVar uintptr) {

	XGtkAccessibleUpdateStateValue(x.GoPointer(), NStatesVar, StatesVar, ValuesVar)

}

// Gets the ID of the @buildable object.
//
// `GtkBuilder` sets the name based on the ID attribute
// of the &lt;object&gt; tag used to construct the @buildable.
func (x *InfoBar) GetBuildableId() string {

	cret := XGtkBuildableGetBuildableId(x.GoPointer())
	return cret
}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GTK"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xNewInfoBar, lib, "gtk_info_bar_new")
	core.PuregoSafeRegister(&xNewWithButtonsInfoBar, lib, "gtk_info_bar_new_with_buttons")

	core.PuregoSafeRegister(&xInfoBarAddActionWidget, lib, "gtk_info_bar_add_action_widget")
	core.PuregoSafeRegister(&xInfoBarAddButton, lib, "gtk_info_bar_add_button")
	core.PuregoSafeRegister(&xInfoBarAddButtons, lib, "gtk_info_bar_add_buttons")
	core.PuregoSafeRegister(&xInfoBarAddChild, lib, "gtk_info_bar_add_child")
	core.PuregoSafeRegister(&xInfoBarGetMessageType, lib, "gtk_info_bar_get_message_type")
	core.PuregoSafeRegister(&xInfoBarGetRevealed, lib, "gtk_info_bar_get_revealed")
	core.PuregoSafeRegister(&xInfoBarGetShowCloseButton, lib, "gtk_info_bar_get_show_close_button")
	core.PuregoSafeRegister(&xInfoBarRemoveActionWidget, lib, "gtk_info_bar_remove_action_widget")
	core.PuregoSafeRegister(&xInfoBarRemoveChild, lib, "gtk_info_bar_remove_child")
	core.PuregoSafeRegister(&xInfoBarResponse, lib, "gtk_info_bar_response")
	core.PuregoSafeRegister(&xInfoBarSetDefaultResponse, lib, "gtk_info_bar_set_default_response")
	core.PuregoSafeRegister(&xInfoBarSetMessageType, lib, "gtk_info_bar_set_message_type")
	core.PuregoSafeRegister(&xInfoBarSetResponseSensitive, lib, "gtk_info_bar_set_response_sensitive")
	core.PuregoSafeRegister(&xInfoBarSetRevealed, lib, "gtk_info_bar_set_revealed")
	core.PuregoSafeRegister(&xInfoBarSetShowCloseButton, lib, "gtk_info_bar_set_show_close_button")

}
