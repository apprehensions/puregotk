// Package adw was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package adw

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/gdk"
	"github.com/jwijenbergh/puregotk/v4/gio"
	"github.com/jwijenbergh/puregotk/v4/gobject"
	"github.com/jwijenbergh/puregotk/v4/gsk"
	"github.com/jwijenbergh/puregotk/v4/gtk"
)

type MessageDialogClass struct {
	ParentClass uintptr

	Padding uintptr
}

// Describes the possible styles of [class@MessageDialog] response buttons.
//
// See [method@MessageDialog.set_response_appearance].
type ResponseAppearance int

const (

	// the default appearance.
	ResponseDefaultValue ResponseAppearance = 0
	// used to denote important responses such as the
	//     affirmative action.
	ResponseSuggestedValue ResponseAppearance = 1
	// used to draw attention to the potentially damaging
	//     consequences of using the response. This appearance acts as a warning to
	//     the user.
	ResponseDestructiveValue ResponseAppearance = 2
)

// A dialog presenting a message or a question.
//
// &lt;picture&gt;
//
//	&lt;source srcset="message-dialog-dark.png" media="(prefers-color-scheme: dark)"&gt;
//	&lt;img src="message-dialog.png" alt="message-dialog"&gt;
//
// &lt;/picture&gt;
//
// Message dialogs have a heading, a body, an optional child widget, and one or
// multiple responses, each presented as a button.
//
// Each response has a unique string ID, and a button label. Additionally, each
// response can be enabled or disabled, and can have a suggested or destructive
// appearance.
//
// When one of the responses is activated, or the dialog is closed, the
// [signal@MessageDialog::response] signal will be emitted. This signal is
// detailed, and the detail, as well as the `response` parameter will be set to
// the ID of the activated response, or to the value of the
// [property@MessageDialog:close-response] property if the dialog had been
// closed without activating any of the responses.
//
// Response buttons can be presented horizontally or vertically depending on
// available space.
//
// When a response is activated, `AdwMessageDialog` is closed automatically.
//
// An example of using a message dialog:
//
// ```c
// GtkWidget *dialog;
//
// dialog = adw_message_dialog_new (parent, _("Replace File?"), NULL);
//
// adw_message_dialog_format_body (ADW_MESSAGE_DIALOG (dialog),
//
//	_("A file named “%s” already exists. Do you want to replace it?"),
//	filename);
//
// adw_message_dialog_add_responses (ADW_MESSAGE_DIALOG (dialog),
//
//	"cancel",  _("_Cancel"),
//	"replace", _("_Replace"),
//	NULL);
//
// adw_message_dialog_set_response_appearance (ADW_MESSAGE_DIALOG (dialog), "replace", ADW_RESPONSE_DESTRUCTIVE);
//
// adw_message_dialog_set_default_response (ADW_MESSAGE_DIALOG (dialog), "cancel");
// adw_message_dialog_set_close_response (ADW_MESSAGE_DIALOG (dialog), "cancel");
//
// g_signal_connect (dialog, "response", G_CALLBACK (response_cb), self);
//
// gtk_window_present (GTK_WINDOW (dialog));
// ```
//
// ## Async API
//
// `AdwMessageDialog` can also be used via the [method@MessageDialog.choose]
// method. This API follows the GIO async pattern, and the result can be
// obtained by calling [method@MessageDialog.choose_finish], for example:
//
// ```c
// static void
// dialog_cb (AdwMessageDialog *dialog,
//
//	GAsyncResult     *result,
//	MyWindow         *self)
//
//	{
//	  const char *response = adw_message_dialog_choose_finish (dialog, result);
//
//	  // ...
//	}
//
// static void
// show_dialog (MyWindow *self)
//
//	{
//	  GtkWidget *dialog;
//
//	  dialog = adw_message_dialog_new (GTK_WINDOW (self), _("Replace File?"), NULL);
//
//	  adw_message_dialog_format_body (ADW_MESSAGE_DIALOG (dialog),
//	                                  _("A file named “%s” already exists. Do you want to replace it?"),
//	                                  filename);
//
//	  adw_message_dialog_add_responses (ADW_MESSAGE_DIALOG (dialog),
//	                                    "cancel",  _("_Cancel"),
//	                                    "replace", _("_Replace"),
//	                                    NULL);
//
//	  adw_message_dialog_set_response_appearance (ADW_MESSAGE_DIALOG (dialog), "replace", ADW_RESPONSE_DESTRUCTIVE);
//
//	  adw_message_dialog_set_default_response (ADW_MESSAGE_DIALOG (dialog), "cancel");
//	  adw_message_dialog_set_close_response (ADW_MESSAGE_DIALOG (dialog), "cancel");
//
//	  adw_message_dialog_choose (ADW_MESSAGE_DIALOG (dialog), NULL, (GAsyncReadyCallback) dialog_cb, self);
//	}
//
// ```
//
// ## AdwMessageDialog as GtkBuildable
//
// `AdwMessageDialog` supports adding responses in UI definitions by via the
// `&lt;responses&gt;` element that may contain multiple `&lt;response&gt;` elements, each
// respresenting a response.
//
// Each of the `&lt;response&gt;` elements must have the `id` attribute specifying the
// response ID. The contents of the element are used as the response label.
//
// Response labels can be translated with the usual `translatable`, `context`
// and `comments` attributes.
//
// The `&lt;response&gt;` elements can also have `enabled` and/or `appearance`
// attributes. See [method@MessageDialog.set_response_enabled] and
// [method@MessageDialog.set_response_appearance] for details.
//
// Example of an `AdwMessageDialog` UI definition:
//
// ```xml
// &lt;object class="AdwMessageDialog" id="dialog"&gt;
//
//	&lt;property name="heading" translatable="yes"&gt;Save Changes?&lt;/property&gt;
//	&lt;property name="body" translatable="yes"&gt;Open documents contain unsaved changes. Changes which are not saved will be permanently lost.&lt;/property&gt;
//	&lt;property name="default-response"&gt;save&lt;/property&gt;
//	&lt;property name="close-response"&gt;cancel&lt;/property&gt;
//	&lt;signal name="response" handler="response_cb"/&gt;
//	&lt;responses&gt;
//	  &lt;response id="cancel" translatable="yes"&gt;_Cancel&lt;/response&gt;
//	  &lt;response id="discard" translatable="yes" appearance="destructive"&gt;_Discard&lt;/response&gt;
//	  &lt;response id="save" translatable="yes" appearance="suggested" enabled="false"&gt;_Save&lt;/response&gt;
//	&lt;/responses&gt;
//
// &lt;/object&gt;
// ```
//
// ## Accessibility
//
// `AdwMessageDialog` uses the `GTK_ACCESSIBLE_ROLE_DIALOG` role.
type MessageDialog struct {
	gtk.Window
}

func MessageDialogNewFromInternalPtr(ptr uintptr) *MessageDialog {
	cls := &MessageDialog{}
	cls.Ptr = ptr
	return cls
}

var xNewMessageDialog func(uintptr, string, string) uintptr

// Creates a new `AdwMessageDialog`.
//
// @heading and @body can be set to `NULL`. This can be useful if they need to
// be formatted or use markup. In that case, set them to `NULL` and call
// [method@MessageDialog.format_body] or similar methods afterwards:
//
// ```c
// GtkWidget *dialog;
//
// dialog = adw_message_dialog_new (parent, _("Replace File?"), NULL);
// adw_message_dialog_format_body (ADW_MESSAGE_DIALOG (dialog),
//
//	_("A file named “%s” already exists.  Do you want to replace it?"),
//	filename);
//
// ```
func NewMessageDialog(ParentVar *gtk.Window, HeadingVar string, BodyVar string) *gtk.Widget {
	var cls *gtk.Widget

	cret := xNewMessageDialog(ParentVar.GoPointer(), HeadingVar, BodyVar)

	if cret == 0 {
		return cls
	}
	gobject.IncreaseRef(cret)
	cls = &gtk.Widget{}
	cls.Ptr = cret
	return cls
}

var xMessageDialogAddResponse func(uintptr, string, string)

// Adds a response with @id and @label to @self.
//
// Responses are represented as buttons in the dialog.
//
// Response ID must be unique. It will be used in
// [signal@MessageDialog::response] to tell which response had been activated,
// as well as to inspect and modify the response later.
//
// An embedded underline in @label indicates a mnemonic.
//
// [method@MessageDialog.set_response_label] can be used to change the response
// label after it had been added.
//
// [method@MessageDialog.set_response_enabled] and
// [method@MessageDialog.set_response_appearance] can be used to customize the
// responses further.
func (x *MessageDialog) AddResponse(IdVar string, LabelVar string) {

	xMessageDialogAddResponse(x.GoPointer(), IdVar, LabelVar)

}

var xMessageDialogAddResponses func(uintptr, string, ...interface{})

// Adds multiple responses to @self.
//
// This is the same as calling [method@MessageDialog.add_response] repeatedly.
// The variable argument list should be `NULL`-terminated list of response IDs
// and labels.
//
// Example:
//
// ```c
// adw_message_dialog_add_responses (dialog,
//
//	"cancel",  _("_Cancel"),
//	"discard", _("_Discard"),
//	"save",    _("_Save"),
//	NULL);
//
// ```
func (x *MessageDialog) AddResponses(FirstIdVar string, varArgs ...interface{}) {

	xMessageDialogAddResponses(x.GoPointer(), FirstIdVar, varArgs...)

}

var xMessageDialogChoose func(uintptr, uintptr, uintptr, uintptr)

// This function shows @self to the user.
//
// The @callback will be called when the alert is dismissed. It should call
// [method@MessageDialog.choose_finish] to obtain the result.
func (x *MessageDialog) Choose(CancellableVar *gio.Cancellable, CallbackVar gio.AsyncReadyCallback, UserDataVar uintptr) {

	xMessageDialogChoose(x.GoPointer(), CancellableVar.GoPointer(), purego.NewCallback(CallbackVar), UserDataVar)

}

var xMessageDialogChooseFinish func(uintptr, uintptr) string

// Finishes the [method@MessageDialog.choose] call and returns the response ID.
func (x *MessageDialog) ChooseFinish(ResultVar gio.AsyncResult) string {

	cret := xMessageDialogChooseFinish(x.GoPointer(), ResultVar.GoPointer())
	return cret
}

var xMessageDialogFormatBody func(uintptr, string, ...interface{})

// Sets the formatted body text of @self.
//
// See [property@MessageDialog:body].
func (x *MessageDialog) FormatBody(FormatVar string, varArgs ...interface{}) {

	xMessageDialogFormatBody(x.GoPointer(), FormatVar, varArgs...)

}

var xMessageDialogFormatBodyMarkup func(uintptr, string, ...interface{})

// Sets the formatted body text of @self with Pango markup.
//
// The @format is assumed to contain Pango markup.
//
// Special XML characters in the `printf()` arguments passed to this function
// will  automatically be escaped as necessary, see
// [func@GLib.markup_printf_escaped].
//
// See [property@MessageDialog:body].
func (x *MessageDialog) FormatBodyMarkup(FormatVar string, varArgs ...interface{}) {

	xMessageDialogFormatBodyMarkup(x.GoPointer(), FormatVar, varArgs...)

}

var xMessageDialogFormatHeading func(uintptr, string, ...interface{})

// Sets the formatted heading of @self.
//
// See [property@MessageDialog:heading].
func (x *MessageDialog) FormatHeading(FormatVar string, varArgs ...interface{}) {

	xMessageDialogFormatHeading(x.GoPointer(), FormatVar, varArgs...)

}

var xMessageDialogFormatHeadingMarkup func(uintptr, string, ...interface{})

// Sets the formatted heading of @self with Pango markup.
//
// The @format is assumed to contain Pango markup.
//
// Special XML characters in the `printf()` arguments passed to this function
// will automatically be escaped as necessary, see
// [func@GLib.markup_printf_escaped].
//
// See [property@MessageDialog:heading].
func (x *MessageDialog) FormatHeadingMarkup(FormatVar string, varArgs ...interface{}) {

	xMessageDialogFormatHeadingMarkup(x.GoPointer(), FormatVar, varArgs...)

}

var xMessageDialogGetBody func(uintptr) string

// Gets the body text of @self.
func (x *MessageDialog) GetBody() string {

	cret := xMessageDialogGetBody(x.GoPointer())
	return cret
}

var xMessageDialogGetBodyUseMarkup func(uintptr) bool

// Gets whether the body text of @self includes Pango markup.
func (x *MessageDialog) GetBodyUseMarkup() bool {

	cret := xMessageDialogGetBodyUseMarkup(x.GoPointer())
	return cret
}

var xMessageDialogGetCloseResponse func(uintptr) string

// Gets the ID of the close response of @self.
func (x *MessageDialog) GetCloseResponse() string {

	cret := xMessageDialogGetCloseResponse(x.GoPointer())
	return cret
}

var xMessageDialogGetDefaultResponse func(uintptr) string

// Gets the ID of the default response of @self.
func (x *MessageDialog) GetDefaultResponse() string {

	cret := xMessageDialogGetDefaultResponse(x.GoPointer())
	return cret
}

var xMessageDialogGetExtraChild func(uintptr) uintptr

// Gets the child widget of @self.
func (x *MessageDialog) GetExtraChild() *gtk.Widget {
	var cls *gtk.Widget

	cret := xMessageDialogGetExtraChild(x.GoPointer())

	if cret == 0 {
		return cls
	}
	gobject.IncreaseRef(cret)
	cls = &gtk.Widget{}
	cls.Ptr = cret
	return cls
}

var xMessageDialogGetHeading func(uintptr) string

// Gets the heading of @self.
func (x *MessageDialog) GetHeading() string {

	cret := xMessageDialogGetHeading(x.GoPointer())
	return cret
}

var xMessageDialogGetHeadingUseMarkup func(uintptr) bool

// Gets whether the heading of @self includes Pango markup.
func (x *MessageDialog) GetHeadingUseMarkup() bool {

	cret := xMessageDialogGetHeadingUseMarkup(x.GoPointer())
	return cret
}

var xMessageDialogGetResponseAppearance func(uintptr, string) ResponseAppearance

// Gets the appearance of @response.
//
// See [method@MessageDialog.set_response_appearance].
func (x *MessageDialog) GetResponseAppearance(ResponseVar string) ResponseAppearance {

	cret := xMessageDialogGetResponseAppearance(x.GoPointer(), ResponseVar)
	return cret
}

var xMessageDialogGetResponseEnabled func(uintptr, string) bool

// Gets whether @response is enabled.
//
// See [method@MessageDialog.set_response_enabled].
func (x *MessageDialog) GetResponseEnabled(ResponseVar string) bool {

	cret := xMessageDialogGetResponseEnabled(x.GoPointer(), ResponseVar)
	return cret
}

var xMessageDialogGetResponseLabel func(uintptr, string) string

// Gets the label of @response.
//
// See [method@MessageDialog.set_response_label].
func (x *MessageDialog) GetResponseLabel(ResponseVar string) string {

	cret := xMessageDialogGetResponseLabel(x.GoPointer(), ResponseVar)
	return cret
}

var xMessageDialogHasResponse func(uintptr, string) bool

// Gets whether @self has a response with the ID @response.
func (x *MessageDialog) HasResponse(ResponseVar string) bool {

	cret := xMessageDialogHasResponse(x.GoPointer(), ResponseVar)
	return cret
}

var xMessageDialogResponse func(uintptr, string)

// Emits the [signal@MessageDialog::response] signal with the given response ID.
//
// Used to indicate that the user has responded to the dialog in some way.
func (x *MessageDialog) Response(ResponseVar string) {

	xMessageDialogResponse(x.GoPointer(), ResponseVar)

}

var xMessageDialogSetBody func(uintptr, string)

// Sets the body text of @self.
func (x *MessageDialog) SetBody(BodyVar string) {

	xMessageDialogSetBody(x.GoPointer(), BodyVar)

}

var xMessageDialogSetBodyUseMarkup func(uintptr, bool)

// Sets whether the body text of @self includes Pango markup.
//
// See [func@Pango.parse_markup].
func (x *MessageDialog) SetBodyUseMarkup(UseMarkupVar bool) {

	xMessageDialogSetBodyUseMarkup(x.GoPointer(), UseMarkupVar)

}

var xMessageDialogSetCloseResponse func(uintptr, string)

// Sets the ID of the close response of @self.
//
// It will be passed to [signal@MessageDialog::response] if the window is
// closed by pressing &lt;kbd&gt;Escape&lt;/kbd&gt; or with a system action.
//
// It doesn't have to correspond to any of the responses in the dialog.
//
// The default close response is `close`.
func (x *MessageDialog) SetCloseResponse(ResponseVar string) {

	xMessageDialogSetCloseResponse(x.GoPointer(), ResponseVar)

}

var xMessageDialogSetDefaultResponse func(uintptr, string)

// Sets the ID of the default response of @self.
//
// If set, pressing &lt;kbd&gt;Enter&lt;/kbd&gt; will activate the corresponding button.
//
// If set to `NULL` or to a non-existent response ID, pressing &lt;kbd&gt;Enter&lt;/kbd&gt;
// will do nothing.
func (x *MessageDialog) SetDefaultResponse(ResponseVar string) {

	xMessageDialogSetDefaultResponse(x.GoPointer(), ResponseVar)

}

var xMessageDialogSetExtraChild func(uintptr, uintptr)

// Sets the child widget of @self.
//
// The child widget is displayed below the heading and body.
func (x *MessageDialog) SetExtraChild(ChildVar *gtk.Widget) {

	xMessageDialogSetExtraChild(x.GoPointer(), ChildVar.GoPointer())

}

var xMessageDialogSetHeading func(uintptr, string)

// Sets the heading of @self.
func (x *MessageDialog) SetHeading(HeadingVar string) {

	xMessageDialogSetHeading(x.GoPointer(), HeadingVar)

}

var xMessageDialogSetHeadingUseMarkup func(uintptr, bool)

// Sets whether the heading of @self includes Pango markup.
//
// See [func@Pango.parse_markup].
func (x *MessageDialog) SetHeadingUseMarkup(UseMarkupVar bool) {

	xMessageDialogSetHeadingUseMarkup(x.GoPointer(), UseMarkupVar)

}

var xMessageDialogSetResponseAppearance func(uintptr, string, ResponseAppearance)

// Sets the appearance for @response.
//
// &lt;picture&gt;
//
//	&lt;source srcset="message-dialog-appearance-dark.png" media="(prefers-color-scheme: dark)"&gt;
//	&lt;img src="message-dialog-appearance.png" alt="message-dialog-appearance"&gt;
//
// &lt;/picture&gt;
//
// Use `ADW_RESPONSE_SUGGESTED` to mark important responses such as the
// affirmative action, like the Save button in the example.
//
// Use `ADW_RESPONSE_DESTRUCTIVE` to draw attention to the potentially damaging
// consequences of using @response. This appearance acts as a warning to the
// user. The Discard button in the example is using this appearance.
//
// The default appearance is `ADW_RESPONSE_DEFAULT`.
//
// Negative responses like Cancel or Close should use the default appearance.
func (x *MessageDialog) SetResponseAppearance(ResponseVar string, AppearanceVar ResponseAppearance) {

	xMessageDialogSetResponseAppearance(x.GoPointer(), ResponseVar, AppearanceVar)

}

var xMessageDialogSetResponseEnabled func(uintptr, string, bool)

// Sets whether @response is enabled.
//
// If @response is not enabled, the corresponding button will have
// [property@Gtk.Widget:sensitive] set to `FALSE` and it can't be activated as
// a default response.
//
// @response can still be used as [property@MessageDialog:close-response] while
// it's not enabled.
//
// Responses are enabled by default.
func (x *MessageDialog) SetResponseEnabled(ResponseVar string, EnabledVar bool) {

	xMessageDialogSetResponseEnabled(x.GoPointer(), ResponseVar, EnabledVar)

}

var xMessageDialogSetResponseLabel func(uintptr, string, string)

// Sets the label of @response to @label.
//
// Labels are displayed on the dialog buttons. An embedded underline in @label
// indicates a mnemonic.
func (x *MessageDialog) SetResponseLabel(ResponseVar string, LabelVar string) {

	xMessageDialogSetResponseLabel(x.GoPointer(), ResponseVar, LabelVar)

}

func (c *MessageDialog) GoPointer() uintptr {
	return c.Ptr
}

func (c *MessageDialog) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// This signal is emitted when the dialog is closed.
//
// @response will be set to the response ID of the button that had been
// activated.
//
// if the dialog was closed by pressing &lt;kbd&gt;Escape&lt;/kbd&gt; or with a system
// action, @response will be set to the value of
// [property@MessageDialog:close-response].
func (x *MessageDialog) ConnectResponse(cb func(MessageDialog, string)) {
	fcb := func(clsPtr uintptr, ResponseVarp string) {
		fa := MessageDialog{}
		fa.Ptr = clsPtr

		cb(fa, ResponseVarp)

	}
	gobject.ObjectConnect(x.GoPointer(), "signal::response", purego.NewCallback(fcb))
}

// Retrieves the `GtkAccessibleRole` for the given `GtkAccessible`.
func (x *MessageDialog) GetAccessibleRole() gtk.AccessibleRole {

	cret := gtk.XGtkAccessibleGetAccessibleRole(x.GoPointer())
	return cret
}

// Resets the accessible @property to its default value.
func (x *MessageDialog) ResetProperty(PropertyVar gtk.AccessibleProperty) {

	gtk.XGtkAccessibleResetProperty(x.GoPointer(), PropertyVar)

}

// Resets the accessible @relation to its default value.
func (x *MessageDialog) ResetRelation(RelationVar gtk.AccessibleRelation) {

	gtk.XGtkAccessibleResetRelation(x.GoPointer(), RelationVar)

}

// Resets the accessible @state to its default value.
func (x *MessageDialog) ResetState(StateVar gtk.AccessibleState) {

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
func (x *MessageDialog) UpdateProperty(FirstPropertyVar gtk.AccessibleProperty, varArgs ...interface{}) {

	gtk.XGtkAccessibleUpdateProperty(x.GoPointer(), FirstPropertyVar, varArgs...)

}

// Updates an array of accessible properties.
//
// This function should be called by `GtkWidget` types whenever an accessible
// property change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *MessageDialog) UpdatePropertyValue(NPropertiesVar int, PropertiesVar uintptr, ValuesVar uintptr) {

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
func (x *MessageDialog) UpdateRelation(FirstRelationVar gtk.AccessibleRelation, varArgs ...interface{}) {

	gtk.XGtkAccessibleUpdateRelation(x.GoPointer(), FirstRelationVar, varArgs...)

}

// Updates an array of accessible relations.
//
// This function should be called by `GtkWidget` types whenever an accessible
// relation change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *MessageDialog) UpdateRelationValue(NRelationsVar int, RelationsVar uintptr, ValuesVar uintptr) {

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
func (x *MessageDialog) UpdateState(FirstStateVar gtk.AccessibleState, varArgs ...interface{}) {

	gtk.XGtkAccessibleUpdateState(x.GoPointer(), FirstStateVar, varArgs...)

}

// Updates an array of accessible states.
//
// This function should be called by `GtkWidget` types whenever an accessible
// state change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *MessageDialog) UpdateStateValue(NStatesVar int, StatesVar uintptr, ValuesVar uintptr) {

	gtk.XGtkAccessibleUpdateStateValue(x.GoPointer(), NStatesVar, StatesVar, ValuesVar)

}

// Gets the ID of the @buildable object.
//
// `GtkBuilder` sets the name based on the ID attribute
// of the &lt;object&gt; tag used to construct the @buildable.
func (x *MessageDialog) GetBuildableId() string {

	cret := gtk.XGtkBuildableGetBuildableId(x.GoPointer())
	return cret
}

// Returns the renderer that is used for this `GtkNative`.
func (x *MessageDialog) GetRenderer() *gsk.Renderer {
	var cls *gsk.Renderer

	cret := gtk.XGtkNativeGetRenderer(x.GoPointer())

	if cret == 0 {
		return cls
	}
	gobject.IncreaseRef(cret)
	cls = &gsk.Renderer{}
	cls.Ptr = cret
	return cls
}

// Returns the surface of this `GtkNative`.
func (x *MessageDialog) GetSurface() *gdk.Surface {
	var cls *gdk.Surface

	cret := gtk.XGtkNativeGetSurface(x.GoPointer())

	if cret == 0 {
		return cls
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
func (x *MessageDialog) GetSurfaceTransform(XVar float64, YVar float64) {

	gtk.XGtkNativeGetSurfaceTransform(x.GoPointer(), XVar, YVar)

}

// Realizes a `GtkNative`.
//
// This should only be used by subclasses.
func (x *MessageDialog) Realize() {

	gtk.XGtkNativeRealize(x.GoPointer())

}

// Unrealizes a `GtkNative`.
//
// This should only be used by subclasses.
func (x *MessageDialog) Unrealize() {

	gtk.XGtkNativeUnrealize(x.GoPointer())

}

// Returns the display that this `GtkRoot` is on.
func (x *MessageDialog) GetDisplay() *gdk.Display {
	var cls *gdk.Display

	cret := gtk.XGtkRootGetDisplay(x.GoPointer())

	if cret == 0 {
		return cls
	}
	gobject.IncreaseRef(cret)
	cls = &gdk.Display{}
	cls.Ptr = cret
	return cls
}

// Retrieves the current focused widget within the root.
//
// Note that this is the widget that would have the focus
// if the root is active; if the root is not focused then
// `gtk_widget_has_focus (widget)` will be %FALSE for the
// widget.
func (x *MessageDialog) GetFocus() *gtk.Widget {
	var cls *gtk.Widget

	cret := gtk.XGtkRootGetFocus(x.GoPointer())

	if cret == 0 {
		return cls
	}
	gobject.IncreaseRef(cret)
	cls = &gtk.Widget{}
	cls.Ptr = cret
	return cls
}

// If @focus is not the current focus widget, and is focusable, sets
// it as the focus widget for the root.
//
// If @focus is %NULL, unsets the focus widget for the root.
//
// To set the focus to a particular widget in the root, it is usually
// more convenient to use [method@Gtk.Widget.grab_focus] instead of
// this function.
func (x *MessageDialog) SetFocus(FocusVar *gtk.Widget) {

	gtk.XGtkRootSetFocus(x.GoPointer(), FocusVar.GoPointer())

}

func init() {
	lib, err := purego.Dlopen(core.GetPath("ADW"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xNewMessageDialog, lib, "adw_message_dialog_new")

	core.PuregoSafeRegister(&xMessageDialogAddResponse, lib, "adw_message_dialog_add_response")
	core.PuregoSafeRegister(&xMessageDialogAddResponses, lib, "adw_message_dialog_add_responses")
	core.PuregoSafeRegister(&xMessageDialogChoose, lib, "adw_message_dialog_choose")
	core.PuregoSafeRegister(&xMessageDialogChooseFinish, lib, "adw_message_dialog_choose_finish")
	core.PuregoSafeRegister(&xMessageDialogFormatBody, lib, "adw_message_dialog_format_body")
	core.PuregoSafeRegister(&xMessageDialogFormatBodyMarkup, lib, "adw_message_dialog_format_body_markup")
	core.PuregoSafeRegister(&xMessageDialogFormatHeading, lib, "adw_message_dialog_format_heading")
	core.PuregoSafeRegister(&xMessageDialogFormatHeadingMarkup, lib, "adw_message_dialog_format_heading_markup")
	core.PuregoSafeRegister(&xMessageDialogGetBody, lib, "adw_message_dialog_get_body")
	core.PuregoSafeRegister(&xMessageDialogGetBodyUseMarkup, lib, "adw_message_dialog_get_body_use_markup")
	core.PuregoSafeRegister(&xMessageDialogGetCloseResponse, lib, "adw_message_dialog_get_close_response")
	core.PuregoSafeRegister(&xMessageDialogGetDefaultResponse, lib, "adw_message_dialog_get_default_response")
	core.PuregoSafeRegister(&xMessageDialogGetExtraChild, lib, "adw_message_dialog_get_extra_child")
	core.PuregoSafeRegister(&xMessageDialogGetHeading, lib, "adw_message_dialog_get_heading")
	core.PuregoSafeRegister(&xMessageDialogGetHeadingUseMarkup, lib, "adw_message_dialog_get_heading_use_markup")
	core.PuregoSafeRegister(&xMessageDialogGetResponseAppearance, lib, "adw_message_dialog_get_response_appearance")
	core.PuregoSafeRegister(&xMessageDialogGetResponseEnabled, lib, "adw_message_dialog_get_response_enabled")
	core.PuregoSafeRegister(&xMessageDialogGetResponseLabel, lib, "adw_message_dialog_get_response_label")
	core.PuregoSafeRegister(&xMessageDialogHasResponse, lib, "adw_message_dialog_has_response")
	core.PuregoSafeRegister(&xMessageDialogResponse, lib, "adw_message_dialog_response")
	core.PuregoSafeRegister(&xMessageDialogSetBody, lib, "adw_message_dialog_set_body")
	core.PuregoSafeRegister(&xMessageDialogSetBodyUseMarkup, lib, "adw_message_dialog_set_body_use_markup")
	core.PuregoSafeRegister(&xMessageDialogSetCloseResponse, lib, "adw_message_dialog_set_close_response")
	core.PuregoSafeRegister(&xMessageDialogSetDefaultResponse, lib, "adw_message_dialog_set_default_response")
	core.PuregoSafeRegister(&xMessageDialogSetExtraChild, lib, "adw_message_dialog_set_extra_child")
	core.PuregoSafeRegister(&xMessageDialogSetHeading, lib, "adw_message_dialog_set_heading")
	core.PuregoSafeRegister(&xMessageDialogSetHeadingUseMarkup, lib, "adw_message_dialog_set_heading_use_markup")
	core.PuregoSafeRegister(&xMessageDialogSetResponseAppearance, lib, "adw_message_dialog_set_response_appearance")
	core.PuregoSafeRegister(&xMessageDialogSetResponseEnabled, lib, "adw_message_dialog_set_response_enabled")
	core.PuregoSafeRegister(&xMessageDialogSetResponseLabel, lib, "adw_message_dialog_set_response_label")

}
