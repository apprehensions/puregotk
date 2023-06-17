// Package adw was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package adw

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/glib"
	"github.com/jwijenbergh/puregotk/v4/gobject"
	"github.com/jwijenbergh/puregotk/v4/gtk"
)

type ToastClass struct {
	ParentClass uintptr
}

// [class@Toast] behavior when another toast is already displayed.
type ToastPriority int

const (

	// the toast will be queued if another toast is
	//   already displayed.
	ToastPriorityNormalValue ToastPriority = 0
	// the toast will be displayed immediately, pushing
	//   the previous toast into the queue instead.
	ToastPriorityHighValue ToastPriority = 1
)

// A helper object for [class@ToastOverlay].
//
// Toasts are meant to be passed into [method@ToastOverlay.add_toast] as
// follows:
//
// ```c
// adw_toast_overlay_add_toast (overlay, adw_toast_new (_("Simple Toast")));
// ```
//
// &lt;picture&gt;
//
//	&lt;source srcset="toast-simple-dark.png" media="(prefers-color-scheme: dark)"&gt;
//	&lt;img src="toast-simple.png" alt="toast-simple"&gt;
//
// &lt;/picture&gt;
//
// Toasts always have a close button. They emit the [signal@Toast::dismissed]
// signal when disappearing.
//
// [property@Toast:timeout] determines how long the toast stays on screen, while
// [property@Toast:priority] determines how it behaves if another toast is
// already being displayed.
//
// [property@Toast:custom-title] can be used to replace the title label with a
// custom widget.
//
// ## Actions
//
// Toasts can have one button on them, with a label and an attached
// [iface@Gio.Action].
//
// ```c
// AdwToast *toast = adw_toast_new (_("Toast with Action"));
//
// adw_toast_set_button_label (toast, _("_Example"));
// adw_toast_set_action_name (toast, "win.example");
//
// adw_toast_overlay_add_toast (overlay, toast);
// ```
//
// &lt;picture&gt;
//
//	&lt;source srcset="toast-action-dark.png" media="(prefers-color-scheme: dark)"&gt;
//	&lt;img src="toast-action.png" alt="toast-action"&gt;
//
// &lt;/picture&gt;
//
// ## Modifying toasts
//
// Toasts can be modified after they have been shown. For this, an `AdwToast`
// reference must be kept around while the toast is visible.
//
// A common use case for this is using toasts as undo prompts that stack with
// each other, allowing to batch undo the last deleted items:
//
// ```c
//
// static void
// toast_undo_cb (GtkWidget  *sender,
//
//	const char *action,
//	GVariant   *param)
//
//	{
//	  // Undo the deletion
//	}
//
// static void
// dismissed_cb (MyWindow *self)
//
//	{
//	  self-&gt;undo_toast = NULL;
//
//	  // Permanently delete the items
//	}
//
// static void
// delete_item (MyWindow *self,
//
//	MyItem   *item)
//
//	{
//	  g_autofree char *title = NULL;
//	  int n_items;
//
//	  // Mark the item as waiting for deletion
//	  n_items = ... // The number of waiting items
//
//	  if (!self-&gt;undo_toast) {
//	    self-&gt;undo_toast = adw_toast_new_format (_("‘%s’ deleted"), ...);
//
//	    adw_toast_set_priority (self-&gt;undo_toast, ADW_TOAST_PRIORITY_HIGH);
//	    adw_toast_set_button_label (self-&gt;undo_toast, _("_Undo"));
//	    adw_toast_set_action_name (self-&gt;undo_toast, "toast.undo");
//
//	    g_signal_connect_swapped (self-&gt;undo_toast, "dismissed",
//	                              G_CALLBACK (dismissed_cb), self);
//
//	    adw_toast_overlay_add_toast (self-&gt;toast_overlay, self-&gt;undo_toast);
//
//	    return;
//	  }
//
//	  title =
//	    g_strdup_printf (ngettext ("&lt;span font_features='tnum=1'&gt;%d&lt;/span&gt; item deleted",
//	                               "&lt;span font_features='tnum=1'&gt;%d&lt;/span&gt; items deleted",
//	                               n_items), n_items);
//
//	  adw_toast_set_title (self-&gt;undo_toast, title);
//
//	  // Bump the toast timeout
//	  adw_toast_overlay_add_toast (self-&gt;toast_overlay, g_object_ref (self-&gt;undo_toast));
//	}
//
// static void
// my_window_class_init (MyWindowClass *klass)
//
//	{
//	  GtkWidgetClass *widget_class = GTK_WIDGET_CLASS (klass);
//
//	  gtk_widget_class_install_action (widget_class, "toast.undo", NULL, toast_undo_cb);
//	}
//
// ```
//
// &lt;picture&gt;
//
//	&lt;source srcset="toast-undo-dark.png" media="(prefers-color-scheme: dark)"&gt;
//	&lt;img src="toast-undo.png" alt="toast-undo"&gt;
//
// &lt;/picture&gt;
type Toast struct {
	gobject.Object
}

func ToastNewFromInternalPtr(ptr uintptr) *Toast {
	cls := &Toast{}
	cls.Ptr = ptr
	return cls
}

var xNewToast func(string) uintptr

// Creates a new `AdwToast`.
//
// The toast will use @title as its title.
//
// @title can be marked up with the Pango text markup language.
func NewToast(TitleVar string) *Toast {
	NewToastPtr := xNewToast(TitleVar)
	if NewToastPtr == 0 {
		return nil
	}

	NewToastCls := &Toast{}
	NewToastCls.Ptr = NewToastPtr
	return NewToastCls
}

var xNewFormatToast func(string, ...interface{}) uintptr

// Creates a new `AdwToast`.
//
// The toast will use the format string as its title.
//
// See also: [ctor@Toast.new]
func NewFormatToast(FormatVar string, varArgs ...interface{}) *Toast {
	NewFormatToastPtr := xNewFormatToast(FormatVar, varArgs...)
	if NewFormatToastPtr == 0 {
		return nil
	}

	NewFormatToastCls := &Toast{}
	NewFormatToastCls.Ptr = NewFormatToastPtr
	return NewFormatToastCls
}

var xToastDismiss func(uintptr)

// Dismisses @self.
//
// Does nothing if @self has already been dismissed, or hasn't been added to an
// [class@ToastOverlay].
func (x *Toast) Dismiss() {

	xToastDismiss(x.GoPointer())

}

var xToastGetActionName func(uintptr) string

// Gets the name of the associated action.
func (x *Toast) GetActionName() string {

	return xToastGetActionName(x.GoPointer())

}

var xToastGetActionTargetValue func(uintptr) *glib.Variant

// Gets the parameter for action invocations.
func (x *Toast) GetActionTargetValue() *glib.Variant {

	return xToastGetActionTargetValue(x.GoPointer())

}

var xToastGetButtonLabel func(uintptr) string

// Gets the label to show on the button.
func (x *Toast) GetButtonLabel() string {

	return xToastGetButtonLabel(x.GoPointer())

}

var xToastGetCustomTitle func(uintptr) uintptr

// Gets the custom title widget of @self.
func (x *Toast) GetCustomTitle() *gtk.Widget {

	GetCustomTitlePtr := xToastGetCustomTitle(x.GoPointer())
	if GetCustomTitlePtr == 0 {
		return nil
	}

	gobject.IncreaseRef(GetCustomTitlePtr)

	GetCustomTitleCls := &gtk.Widget{}
	GetCustomTitleCls.Ptr = GetCustomTitlePtr
	return GetCustomTitleCls

}

var xToastGetPriority func(uintptr) ToastPriority

// Gets priority for @self.
func (x *Toast) GetPriority() ToastPriority {

	return xToastGetPriority(x.GoPointer())

}

var xToastGetTimeout func(uintptr) uint

// Gets timeout for @self.
func (x *Toast) GetTimeout() uint {

	return xToastGetTimeout(x.GoPointer())

}

var xToastGetTitle func(uintptr) string

// Gets the title that will be displayed on the toast.
//
// If a custom title has been set with [method@Adw.Toast.set_custom_title]
// the return value will be %NULL.
func (x *Toast) GetTitle() string {

	return xToastGetTitle(x.GoPointer())

}

var xToastSetActionName func(uintptr, string)

// Sets the name of the associated action.
//
// It will be activated when clicking the button.
//
// See [property@Toast:action-target].
func (x *Toast) SetActionName(ActionNameVar string) {

	xToastSetActionName(x.GoPointer(), ActionNameVar)

}

var xToastSetActionTarget func(uintptr, string, ...interface{})

// Sets the parameter for action invocations.
//
// This is a convenience function that calls [ctor@GLib.Variant.new] for
// @format_string and uses the result to call
// [method@Toast.set_action_target_value].
//
// If you are setting a string-valued target and want to set
// the action name at the same time, you can use
// [method@Toast.set_detailed_action_name].
func (x *Toast) SetActionTarget(FormatStringVar string, varArgs ...interface{}) {

	xToastSetActionTarget(x.GoPointer(), FormatStringVar, varArgs...)

}

var xToastSetActionTargetValue func(uintptr, *glib.Variant)

// Sets the parameter for action invocations.
//
// If the @action_target variant has a floating reference this function
// will sink it.
func (x *Toast) SetActionTargetValue(ActionTargetVar *glib.Variant) {

	xToastSetActionTargetValue(x.GoPointer(), ActionTargetVar)

}

var xToastSetButtonLabel func(uintptr, string)

// Sets the label to show on the button.
//
// Underlines in the button text can be used to indicate a mnemonic.
//
// If set to `NULL`, the button won't be shown.
//
// See [property@Toast:action-name].
func (x *Toast) SetButtonLabel(ButtonLabelVar string) {

	xToastSetButtonLabel(x.GoPointer(), ButtonLabelVar)

}

var xToastSetCustomTitle func(uintptr, uintptr)

// Sets the custom title widget of @self.
//
// It will be displayed instead of the title if set. In this case,
// [property@Toast:title] is ignored.
//
// Setting a custom title will unset [property@Toast:title].
func (x *Toast) SetCustomTitle(WidgetVar *gtk.Widget) {

	xToastSetCustomTitle(x.GoPointer(), WidgetVar.GoPointer())

}

var xToastSetDetailedActionName func(uintptr, string)

// Sets the action name and its parameter.
//
// @detailed_action_name is a string in the format accepted by
// [func@Gio.Action.parse_detailed_name].
func (x *Toast) SetDetailedActionName(DetailedActionNameVar string) {

	xToastSetDetailedActionName(x.GoPointer(), DetailedActionNameVar)

}

var xToastSetPriority func(uintptr, ToastPriority)

// Sets priority for @self.
//
// Priority controls how the toast behaves when another toast is already
// being displayed.
//
// If @priority is `ADW_TOAST_PRIORITY_NORMAL`, the toast will be queued.
//
// If @priority is `ADW_TOAST_PRIORITY_HIGH`, the toast will be displayed
// immediately, pushing the previous toast into the queue instead.
func (x *Toast) SetPriority(PriorityVar ToastPriority) {

	xToastSetPriority(x.GoPointer(), PriorityVar)

}

var xToastSetTimeout func(uintptr, uint)

// Sets timeout for @self.
//
// If @timeout is 0, the toast is displayed indefinitely until manually
// dismissed.
//
// Toasts cannot disappear while being hovered, pressed (on touchscreen), or
// have keyboard focus inside them.
func (x *Toast) SetTimeout(TimeoutVar uint) {

	xToastSetTimeout(x.GoPointer(), TimeoutVar)

}

var xToastSetTitle func(uintptr, string)

// Sets the title that will be displayed on the toast.
//
// The title can be marked up with the Pango text markup language.
//
// Setting a title will unset [property@Toast:custom-title].
//
// If [property@Toast:custom-title] is set, it will be used instead.
func (x *Toast) SetTitle(TitleVar string) {

	xToastSetTitle(x.GoPointer(), TitleVar)

}

func (c *Toast) GoPointer() uintptr {
	return c.Ptr
}

func (c *Toast) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// Emitted after the button has been clicked.
//
// It can be used as an alternative to setting an action.
func (x *Toast) ConnectButtonClicked(cb func(Toast)) {
	fcb := func(clsPtr uintptr) {
		fa := Toast{}
		fa.Ptr = clsPtr

		cb(fa)

	}
	gobject.ObjectConnect(x.GoPointer(), "signal::button-clicked", purego.NewCallback(fcb))
}

// Emitted when the toast has been dismissed.
func (x *Toast) ConnectDismissed(cb func(Toast)) {
	fcb := func(clsPtr uintptr) {
		fa := Toast{}
		fa.Ptr = clsPtr

		cb(fa)

	}
	gobject.ObjectConnect(x.GoPointer(), "signal::dismissed", purego.NewCallback(fcb))
}

func init() {
	lib, err := purego.Dlopen(core.GetPath("ADW"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xNewToast, lib, "adw_toast_new")
	core.PuregoSafeRegister(&xNewFormatToast, lib, "adw_toast_new_format")

	core.PuregoSafeRegister(&xToastDismiss, lib, "adw_toast_dismiss")
	core.PuregoSafeRegister(&xToastGetActionName, lib, "adw_toast_get_action_name")
	core.PuregoSafeRegister(&xToastGetActionTargetValue, lib, "adw_toast_get_action_target_value")
	core.PuregoSafeRegister(&xToastGetButtonLabel, lib, "adw_toast_get_button_label")
	core.PuregoSafeRegister(&xToastGetCustomTitle, lib, "adw_toast_get_custom_title")
	core.PuregoSafeRegister(&xToastGetPriority, lib, "adw_toast_get_priority")
	core.PuregoSafeRegister(&xToastGetTimeout, lib, "adw_toast_get_timeout")
	core.PuregoSafeRegister(&xToastGetTitle, lib, "adw_toast_get_title")
	core.PuregoSafeRegister(&xToastSetActionName, lib, "adw_toast_set_action_name")
	core.PuregoSafeRegister(&xToastSetActionTarget, lib, "adw_toast_set_action_target")
	core.PuregoSafeRegister(&xToastSetActionTargetValue, lib, "adw_toast_set_action_target_value")
	core.PuregoSafeRegister(&xToastSetButtonLabel, lib, "adw_toast_set_button_label")
	core.PuregoSafeRegister(&xToastSetCustomTitle, lib, "adw_toast_set_custom_title")
	core.PuregoSafeRegister(&xToastSetDetailedActionName, lib, "adw_toast_set_detailed_action_name")
	core.PuregoSafeRegister(&xToastSetPriority, lib, "adw_toast_set_priority")
	core.PuregoSafeRegister(&xToastSetTimeout, lib, "adw_toast_set_timeout")
	core.PuregoSafeRegister(&xToastSetTitle, lib, "adw_toast_set_title")

}
