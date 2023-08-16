// Package gtk was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gtk

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/gio"
	"github.com/jwijenbergh/puregotk/v4/glib"
	"github.com/jwijenbergh/puregotk/v4/gobject"
)

type ApplicationClass struct {
	ParentClass uintptr

	Padding uintptr
}

// Types of user actions that may be blocked by `GtkApplication`.
//
// See [method@Gtk.Application.inhibit].
type ApplicationInhibitFlags int

const (

	// Inhibit ending the user session
	//   by logging out or by shutting down the computer
	ApplicationInhibitLogoutValue ApplicationInhibitFlags = 1
	// Inhibit user switching
	ApplicationInhibitSwitchValue ApplicationInhibitFlags = 2
	// Inhibit suspending the
	//   session or computer
	ApplicationInhibitSuspendValue ApplicationInhibitFlags = 4
	// Inhibit the session being
	//   marked as idle (and possibly locked)
	ApplicationInhibitIdleValue ApplicationInhibitFlags = 8
)

// `GtkApplication` is a high-level API for writing applications.
//
// It supports many aspects of writing a GTK application in a convenient
// fashion, without enforcing a one-size-fits-all model.
//
// Currently, `GtkApplication` handles GTK initialization, application
// uniqueness, session management, provides some basic scriptability and
// desktop shell integration by exporting actions and menus and manages a
// list of toplevel windows whose life-cycle is automatically tied to the
// life-cycle of your application.
//
// While `GtkApplication` works fine with plain [class@Gtk.Window]s, it is
// recommended to use it together with [class@Gtk.ApplicationWindow].
//
// ## Automatic resources
//
// `GtkApplication` will automatically load menus from the `GtkBuilder`
// resource located at "gtk/menus.ui", relative to the application's
// resource base path (see [method@Gio.Application.set_resource_base_path]).
// The menu with the ID "menubar" is taken as the application's
// menubar. Additional menus (most interesting submenus) can be named
// and accessed via [method@Gtk.Application.get_menu_by_id] which allows for
// dynamic population of a part of the menu structure.
//
// It is also possible to provide the menubar manually using
// [method@Gtk.Application.set_menubar].
//
// `GtkApplication` will also automatically setup an icon search path for
// the default icon theme by appending "icons" to the resource base
// path. This allows your application to easily store its icons as
// resources. See [method@Gtk.IconTheme.add_resource_path] for more
// information.
//
// If there is a resource located at `gtk/help-overlay.ui` which
// defines a [class@Gtk.ShortcutsWindow] with ID `help_overlay` then
// `GtkApplication` associates an instance of this shortcuts window with
// each [class@Gtk.ApplicationWindow] and sets up the keyboard accelerator
// &lt;kbd&gt;Control&lt;/kbd&gt;+&lt;kbd&gt;?&lt;/kbd&gt; to open it. To create a menu item that
// displays the shortcuts window, associate the item with the action
// `win.show-help-overlay`.
//
// ## A simple application
//
// [A simple example](https://gitlab.gnome.org/GNOME/gtk/tree/main/examples/bp/bloatpad.c)
// is available in the GTK source code repository
//
// `GtkApplication` optionally registers with a session manager of the
// users session (if you set the [property@Gtk.Application:register-session]
// property) and offers various functionality related to the session
// life-cycle.
//
// An application can block various ways to end the session with
// the [method@Gtk.Application.inhibit] function. Typical use cases for
// this kind of inhibiting are long-running, uninterruptible operations,
// such as burning a CD or performing a disk backup. The session
// manager may not honor the inhibitor, but it can be expected to
// inform the user about the negative consequences of ending the
// session while inhibitors are present.
//
// ## See Also
//
// [HowDoI: Using GtkApplication](https://wiki.gnome.org/HowDoI/GtkApplication),
// [Getting Started with GTK: Basics](getting_started.html#basics)
type Application struct {
	gio.Application
}

func ApplicationNewFromInternalPtr(ptr uintptr) *Application {
	cls := &Application{}
	cls.Ptr = ptr
	return cls
}

var xNewApplication func(string, gio.ApplicationFlags) uintptr

// Creates a new `GtkApplication` instance.
//
// When using `GtkApplication`, it is not necessary to call [func@Gtk.init]
// manually. It is called as soon as the application gets registered as
// the primary instance.
//
// Concretely, [func@Gtk.init] is called in the default handler for the
// `GApplication::startup` signal. Therefore, `GtkApplication` subclasses should
// always chain up in their `GApplication::startup` handler before using any GTK
// API.
//
// Note that commandline arguments are not passed to [func@Gtk.init].
//
// If `application_id` is not %NULL, then it must be valid. See
// `g_application_id_is_valid()`.
//
// If no application ID is given then some features (most notably application
// uniqueness) will be disabled.
func NewApplication(ApplicationIdVar string, FlagsVar gio.ApplicationFlags) *Application {
	var cls *Application

	cret := xNewApplication(ApplicationIdVar, FlagsVar)

	if cret == 0 {
		return nil
	}
	cls = &Application{}
	cls.Ptr = cret
	return cls
}

var xApplicationAddWindow func(uintptr, uintptr)

// Adds a window to `application`.
//
// This call can only happen after the `application` has started;
// typically, you should add new application windows in response
// to the emission of the `GApplication::activate` signal.
//
// This call is equivalent to setting the [property@Gtk.Window:application]
// property of `window` to `application`.
//
// Normally, the connection between the application and the window
// will remain until the window is destroyed, but you can explicitly
// remove it with [method@Gtk.Application.remove_window].
//
// GTK will keep the `application` running as long as it has
// any windows.
func (x *Application) AddWindow(WindowVar *Window) {

	xApplicationAddWindow(x.GoPointer(), WindowVar.GoPointer())

}

var xApplicationGetAccelsForAction func(uintptr, string) []string

// Gets the accelerators that are currently associated with
// the given action.
func (x *Application) GetAccelsForAction(DetailedActionNameVar string) []string {

	cret := xApplicationGetAccelsForAction(x.GoPointer(), DetailedActionNameVar)
	return cret
}

var xApplicationGetActionsForAccel func(uintptr, string) []string

// Returns the list of actions (possibly empty) that `accel` maps to.
//
// Each item in the list is a detailed action name in the usual form.
//
// This might be useful to discover if an accel already exists in
// order to prevent installation of a conflicting accelerator (from
// an accelerator editor or a plugin system, for example). Note that
// having more than one action per accelerator may not be a bad thing
// and might make sense in cases where the actions never appear in the
// same context.
//
// In case there are no actions for a given accelerator, an empty array
// is returned. `NULL` is never returned.
//
// It is a programmer error to pass an invalid accelerator string.
//
// If you are unsure, check it with [func@Gtk.accelerator_parse] first.
func (x *Application) GetActionsForAccel(AccelVar string) []string {

	cret := xApplicationGetActionsForAccel(x.GoPointer(), AccelVar)
	return cret
}

var xApplicationGetActiveWindow func(uintptr) uintptr

// Gets the “active” window for the application.
//
// The active window is the one that was most recently focused (within
// the application).  This window may not have the focus at the moment
// if another application has it — this is just the most
// recently-focused window within this application.
func (x *Application) GetActiveWindow() *Window {
	var cls *Window

	cret := xApplicationGetActiveWindow(x.GoPointer())

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &Window{}
	cls.Ptr = cret
	return cls
}

var xApplicationGetMenuById func(uintptr, string) uintptr

// Gets a menu from automatically loaded resources.
//
// See [the section on Automatic resources](class.Application.html#automatic-resources)
// for more information.
func (x *Application) GetMenuById(IdVar string) *gio.Menu {
	var cls *gio.Menu

	cret := xApplicationGetMenuById(x.GoPointer(), IdVar)

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &gio.Menu{}
	cls.Ptr = cret
	return cls
}

var xApplicationGetMenubar func(uintptr) uintptr

// Returns the menu model that has been set with
// [method@Gtk.Application.set_menubar].
func (x *Application) GetMenubar() *gio.MenuModel {
	var cls *gio.MenuModel

	cret := xApplicationGetMenubar(x.GoPointer())

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &gio.MenuModel{}
	cls.Ptr = cret
	return cls
}

var xApplicationGetWindowById func(uintptr, uint) uintptr

// Returns the [class@Gtk.ApplicationWindow] with the given ID.
//
// The ID of a `GtkApplicationWindow` can be retrieved with
// [method@Gtk.ApplicationWindow.get_id].
func (x *Application) GetWindowById(IdVar uint) *Window {
	var cls *Window

	cret := xApplicationGetWindowById(x.GoPointer(), IdVar)

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &Window{}
	cls.Ptr = cret
	return cls
}

var xApplicationGetWindows func(uintptr) *glib.List

// Gets a list of the [class@Gtk.Window] instances associated with `application`.
//
// The list is sorted by most recently focused window, such that the first
// element is the currently focused window. (Useful for choosing a parent
// for a transient window.)
//
// The list that is returned should not be modified in any way. It will
// only remain valid until the next focus change or window creation or
// deletion.
func (x *Application) GetWindows() *glib.List {

	cret := xApplicationGetWindows(x.GoPointer())
	return cret
}

var xApplicationInhibit func(uintptr, uintptr, ApplicationInhibitFlags, string) uint

// Inform the session manager that certain types of actions should be
// inhibited.
//
// This is not guaranteed to work on all platforms and for all types of
// actions.
//
// Applications should invoke this method when they begin an operation
// that should not be interrupted, such as creating a CD or DVD. The
// types of actions that may be blocked are specified by the `flags`
// parameter. When the application completes the operation it should
// call [method@Gtk.Application.uninhibit] to remove the inhibitor. Note
// that an application can have multiple inhibitors, and all of them must
// be individually removed. Inhibitors are also cleared when the
// application exits.
//
// Applications should not expect that they will always be able to block
// the action. In most cases, users will be given the option to force
// the action to take place.
//
// The `reason` message should be short and to the point.
//
// If `window` is given, the session manager may point the user to
// this window to find out more about why the action is inhibited.
func (x *Application) Inhibit(WindowVar *Window, FlagsVar ApplicationInhibitFlags, ReasonVar string) uint {

	cret := xApplicationInhibit(x.GoPointer(), WindowVar.GoPointer(), FlagsVar, ReasonVar)
	return cret
}

var xApplicationListActionDescriptions func(uintptr) []string

// Lists the detailed action names which have associated accelerators.
//
// See [method@Gtk.Application.set_accels_for_action].
func (x *Application) ListActionDescriptions() []string {

	cret := xApplicationListActionDescriptions(x.GoPointer())
	return cret
}

var xApplicationRemoveWindow func(uintptr, uintptr)

// Remove a window from `application`.
//
// If `window` belongs to `application` then this call is equivalent to
// setting the [property@Gtk.Window:application] property of `window` to
// `NULL`.
//
// The application may stop running as a result of a call to this
// function, if `window` was the last window of the `application`.
func (x *Application) RemoveWindow(WindowVar *Window) {

	xApplicationRemoveWindow(x.GoPointer(), WindowVar.GoPointer())

}

var xApplicationSetAccelsForAction func(uintptr, string, uintptr)

// Sets zero or more keyboard accelerators that will trigger the
// given action.
//
// The first item in `accels` will be the primary accelerator, which may be
// displayed in the UI.
//
// To remove all accelerators for an action, use an empty, zero-terminated
// array for `accels`.
//
// For the `detailed_action_name`, see `g_action_parse_detailed_name()` and
// `g_action_print_detailed_name()`.
func (x *Application) SetAccelsForAction(DetailedActionNameVar string, AccelsVar uintptr) {

	xApplicationSetAccelsForAction(x.GoPointer(), DetailedActionNameVar, AccelsVar)

}

var xApplicationSetMenubar func(uintptr, uintptr)

// Sets or unsets the menubar for windows of `application`.
//
// This is a menubar in the traditional sense.
//
// This can only be done in the primary instance of the application,
// after it has been registered. `GApplication::startup` is a good place
// to call this.
//
// Depending on the desktop environment, this may appear at the top of
// each window, or at the top of the screen.  In some environments, if
// both the application menu and the menubar are set, the application
// menu will be presented as if it were the first item of the menubar.
// Other environments treat the two as completely separate — for example,
// the application menu may be rendered by the desktop shell while the
// menubar (if set) remains in each individual window.
//
// Use the base `GActionMap` interface to add actions, to respond to the
// user selecting these menu items.
func (x *Application) SetMenubar(MenubarVar *gio.MenuModel) {

	xApplicationSetMenubar(x.GoPointer(), MenubarVar.GoPointer())

}

var xApplicationUninhibit func(uintptr, uint)

// Removes an inhibitor that has been previously established.
//
// See [method@Gtk.Application.inhibit].
//
// Inhibitors are also cleared when the application exits.
func (x *Application) Uninhibit(CookieVar uint) {

	xApplicationUninhibit(x.GoPointer(), CookieVar)

}

func (c *Application) GoPointer() uintptr {
	return c.Ptr
}

func (c *Application) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// Emitted when the session manager is about to end the session.
//
// This signal is only emitted if [property@Gtk.Application:register-session]
// is `TRUE`. Applications can connect to this signal and call
// [method@Gtk.Application.inhibit] with `GTK_APPLICATION_INHIBIT_LOGOUT`
// to delay the end of the session until state has been saved.
func (x *Application) ConnectQueryEnd(cb func(Application)) {
	fcb := func(clsPtr uintptr) {
		fa := Application{}
		fa.Ptr = clsPtr

		cb(fa)

	}
	gobject.ObjectConnect(x.GoPointer(), "signal::query-end", purego.NewCallback(fcb))
}

// Emitted when a [class@Gtk.Window] is added to `application` through
// [method@Gtk.Application.add_window].
func (x *Application) ConnectWindowAdded(cb func(Application, uintptr)) {
	fcb := func(clsPtr uintptr, WindowVarp uintptr) {
		fa := Application{}
		fa.Ptr = clsPtr

		cb(fa, WindowVarp)

	}
	gobject.ObjectConnect(x.GoPointer(), "signal::window-added", purego.NewCallback(fcb))
}

// Emitted when a [class@Gtk.Window] is removed from `application`.
//
// This can happen as a side-effect of the window being destroyed
// or explicitly through [method@Gtk.Application.remove_window].
func (x *Application) ConnectWindowRemoved(cb func(Application, uintptr)) {
	fcb := func(clsPtr uintptr, WindowVarp uintptr) {
		fa := Application{}
		fa.Ptr = clsPtr

		cb(fa, WindowVarp)

	}
	gobject.ObjectConnect(x.GoPointer(), "signal::window-removed", purego.NewCallback(fcb))
}

// Emits the #GActionGroup::action-added signal on @action_group.
//
// This function should only be called by #GActionGroup implementations.
func (x *Application) ActionAdded(ActionNameVar string) {

	gio.XGActionGroupActionAdded(x.GoPointer(), ActionNameVar)

}

// Emits the #GActionGroup::action-enabled-changed signal on @action_group.
//
// This function should only be called by #GActionGroup implementations.
func (x *Application) ActionEnabledChanged(ActionNameVar string, EnabledVar bool) {

	gio.XGActionGroupActionEnabledChanged(x.GoPointer(), ActionNameVar, EnabledVar)

}

// Emits the #GActionGroup::action-removed signal on @action_group.
//
// This function should only be called by #GActionGroup implementations.
func (x *Application) ActionRemoved(ActionNameVar string) {

	gio.XGActionGroupActionRemoved(x.GoPointer(), ActionNameVar)

}

// Emits the #GActionGroup::action-state-changed signal on @action_group.
//
// This function should only be called by #GActionGroup implementations.
func (x *Application) ActionStateChanged(ActionNameVar string, StateVar *glib.Variant) {

	gio.XGActionGroupActionStateChanged(x.GoPointer(), ActionNameVar, StateVar)

}

// Activate the named action within @action_group.
//
// If the action is expecting a parameter, then the correct type of
// parameter must be given as @parameter.  If the action is expecting no
// parameters then @parameter must be %NULL.  See
// g_action_group_get_action_parameter_type().
//
// If the #GActionGroup implementation supports asynchronous remote
// activation over D-Bus, this call may return before the relevant
// D-Bus traffic has been sent, or any replies have been received. In
// order to block on such asynchronous activation calls,
// g_dbus_connection_flush() should be called prior to the code, which
// depends on the result of the action activation. Without flushing
// the D-Bus connection, there is no guarantee that the action would
// have been activated.
//
// The following code which runs in a remote app instance, shows an
// example of a "quit" action being activated on the primary app
// instance over D-Bus. Here g_dbus_connection_flush() is called
// before `exit()`. Without g_dbus_connection_flush(), the "quit" action
// may fail to be activated on the primary instance.
//
// |[&lt;!-- language="C" --&gt;
// // call "quit" action on primary instance
// g_action_group_activate_action (G_ACTION_GROUP (app), "quit", NULL);
//
// // make sure the action is activated now
// g_dbus_connection_flush (...);
//
// g_debug ("application has been terminated. exiting.");
//
// exit (0);
// ]|
func (x *Application) ActivateAction(ActionNameVar string, ParameterVar *glib.Variant) {

	gio.XGActionGroupActivateAction(x.GoPointer(), ActionNameVar, ParameterVar)

}

// Request for the state of the named action within @action_group to be
// changed to @value.
//
// The action must be stateful and @value must be of the correct type.
// See g_action_group_get_action_state_type().
//
// This call merely requests a change.  The action may refuse to change
// its state or may change its state to something other than @value.
// See g_action_group_get_action_state_hint().
//
// If the @value GVariant is floating, it is consumed.
func (x *Application) ChangeActionState(ActionNameVar string, ValueVar *glib.Variant) {

	gio.XGActionGroupChangeActionState(x.GoPointer(), ActionNameVar, ValueVar)

}

// Checks if the named action within @action_group is currently enabled.
//
// An action must be enabled in order to be activated or in order to
// have its state changed from outside callers.
func (x *Application) GetActionEnabled(ActionNameVar string) bool {

	cret := gio.XGActionGroupGetActionEnabled(x.GoPointer(), ActionNameVar)
	return cret
}

// Queries the type of the parameter that must be given when activating
// the named action within @action_group.
//
// When activating the action using g_action_group_activate_action(),
// the #GVariant given to that function must be of the type returned
// by this function.
//
// In the case that this function returns %NULL, you must not give any
// #GVariant, but %NULL instead.
//
// The parameter type of a particular action will never change but it is
// possible for an action to be removed and for a new action to be added
// with the same name but a different parameter type.
func (x *Application) GetActionParameterType(ActionNameVar string) *glib.VariantType {

	cret := gio.XGActionGroupGetActionParameterType(x.GoPointer(), ActionNameVar)
	return cret
}

// Queries the current state of the named action within @action_group.
//
// If the action is not stateful then %NULL will be returned.  If the
// action is stateful then the type of the return value is the type
// given by g_action_group_get_action_state_type().
//
// The return value (if non-%NULL) should be freed with
// g_variant_unref() when it is no longer required.
func (x *Application) GetActionState(ActionNameVar string) *glib.Variant {

	cret := gio.XGActionGroupGetActionState(x.GoPointer(), ActionNameVar)
	return cret
}

// Requests a hint about the valid range of values for the state of the
// named action within @action_group.
//
// If %NULL is returned it either means that the action is not stateful
// or that there is no hint about the valid range of values for the
// state of the action.
//
// If a #GVariant array is returned then each item in the array is a
// possible value for the state.  If a #GVariant pair (ie: two-tuple) is
// returned then the tuple specifies the inclusive lower and upper bound
// of valid values for the state.
//
// In any case, the information is merely a hint.  It may be possible to
// have a state value outside of the hinted range and setting a value
// within the range may fail.
//
// The return value (if non-%NULL) should be freed with
// g_variant_unref() when it is no longer required.
func (x *Application) GetActionStateHint(ActionNameVar string) *glib.Variant {

	cret := gio.XGActionGroupGetActionStateHint(x.GoPointer(), ActionNameVar)
	return cret
}

// Queries the type of the state of the named action within
// @action_group.
//
// If the action is stateful then this function returns the
// #GVariantType of the state.  All calls to
// g_action_group_change_action_state() must give a #GVariant of this
// type and g_action_group_get_action_state() will return a #GVariant
// of the same type.
//
// If the action is not stateful then this function will return %NULL.
// In that case, g_action_group_get_action_state() will return %NULL
// and you must not call g_action_group_change_action_state().
//
// The state type of a particular action will never change but it is
// possible for an action to be removed and for a new action to be added
// with the same name but a different state type.
func (x *Application) GetActionStateType(ActionNameVar string) *glib.VariantType {

	cret := gio.XGActionGroupGetActionStateType(x.GoPointer(), ActionNameVar)
	return cret
}

// Checks if the named action exists within @action_group.
func (x *Application) HasAction(ActionNameVar string) bool {

	cret := gio.XGActionGroupHasAction(x.GoPointer(), ActionNameVar)
	return cret
}

// Lists the actions contained within @action_group.
//
// The caller is responsible for freeing the list with g_strfreev() when
// it is no longer required.
func (x *Application) ListActions() uintptr {

	cret := gio.XGActionGroupListActions(x.GoPointer())
	return cret
}

// Queries all aspects of the named action within an @action_group.
//
// This function acquires the information available from
// g_action_group_has_action(), g_action_group_get_action_enabled(),
// g_action_group_get_action_parameter_type(),
// g_action_group_get_action_state_type(),
// g_action_group_get_action_state_hint() and
// g_action_group_get_action_state() with a single function call.
//
// This provides two main benefits.
//
// The first is the improvement in efficiency that comes with not having
// to perform repeated lookups of the action in order to discover
// different things about it.  The second is that implementing
// #GActionGroup can now be done by only overriding this one virtual
// function.
//
// The interface provides a default implementation of this function that
// calls the individual functions, as required, to fetch the
// information.  The interface also provides default implementations of
// those functions that call this function.  All implementations,
// therefore, must override either this function or all of the others.
//
// If the action exists, %TRUE is returned and any of the requested
// fields (as indicated by having a non-%NULL reference passed in) are
// filled.  If the action doesn't exist, %FALSE is returned and the
// fields may or may not have been modified.
func (x *Application) QueryAction(ActionNameVar string, EnabledVar bool, ParameterTypeVar **glib.VariantType, StateTypeVar **glib.VariantType, StateHintVar **glib.Variant, StateVar **glib.Variant) bool {

	cret := gio.XGActionGroupQueryAction(x.GoPointer(), ActionNameVar, EnabledVar, ParameterTypeVar, StateTypeVar, StateHintVar, StateVar)
	return cret
}

// Adds an action to the @action_map.
//
// If the action map already contains an action with the same name
// as @action then the old action is dropped from the action map.
//
// The action map takes its own reference on @action.
func (x *Application) AddAction(ActionVar gio.Action) {

	gio.XGActionMapAddAction(x.GoPointer(), ActionVar.GoPointer())

}

// A convenience function for creating multiple #GSimpleAction instances
// and adding them to a #GActionMap.
//
// Each action is constructed as per one #GActionEntry.
//
// |[&lt;!-- language="C" --&gt;
// static void
// activate_quit (GSimpleAction *simple,
//
//	GVariant      *parameter,
//	gpointer       user_data)
//
//	{
//	  exit (0);
//	}
//
// static void
// activate_print_string (GSimpleAction *simple,
//
//	GVariant      *parameter,
//	gpointer       user_data)
//
//	{
//	  g_print ("%s\n", g_variant_get_string (parameter, NULL));
//	}
//
// static GActionGroup *
// create_action_group (void)
//
//	{
//	  const GActionEntry entries[] = {
//	    { "quit",         activate_quit              },
//	    { "print-string", activate_print_string, "s" }
//	  };
//	  GSimpleActionGroup *group;
//
//	  group = g_simple_action_group_new ();
//	  g_action_map_add_action_entries (G_ACTION_MAP (group), entries, G_N_ELEMENTS (entries), NULL);
//
//	  return G_ACTION_GROUP (group);
//	}
//
// ]|
func (x *Application) AddActionEntries(EntriesVar uintptr, NEntriesVar int, UserDataVar uintptr) {

	gio.XGActionMapAddActionEntries(x.GoPointer(), EntriesVar, NEntriesVar, UserDataVar)

}

// Looks up the action with the name @action_name in @action_map.
//
// If no such action exists, returns %NULL.
func (x *Application) LookupAction(ActionNameVar string) *gio.ActionBase {
	var cls *gio.ActionBase

	cret := gio.XGActionMapLookupAction(x.GoPointer(), ActionNameVar)

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &gio.ActionBase{}
	cls.Ptr = cret
	return cls
}

// Removes the named action from the action map.
//
// If no action of this name is in the map then nothing happens.
func (x *Application) RemoveAction(ActionNameVar string) {

	gio.XGActionMapRemoveAction(x.GoPointer(), ActionNameVar)

}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GTK"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xNewApplication, lib, "gtk_application_new")

	core.PuregoSafeRegister(&xApplicationAddWindow, lib, "gtk_application_add_window")
	core.PuregoSafeRegister(&xApplicationGetAccelsForAction, lib, "gtk_application_get_accels_for_action")
	core.PuregoSafeRegister(&xApplicationGetActionsForAccel, lib, "gtk_application_get_actions_for_accel")
	core.PuregoSafeRegister(&xApplicationGetActiveWindow, lib, "gtk_application_get_active_window")
	core.PuregoSafeRegister(&xApplicationGetMenuById, lib, "gtk_application_get_menu_by_id")
	core.PuregoSafeRegister(&xApplicationGetMenubar, lib, "gtk_application_get_menubar")
	core.PuregoSafeRegister(&xApplicationGetWindowById, lib, "gtk_application_get_window_by_id")
	core.PuregoSafeRegister(&xApplicationGetWindows, lib, "gtk_application_get_windows")
	core.PuregoSafeRegister(&xApplicationInhibit, lib, "gtk_application_inhibit")
	core.PuregoSafeRegister(&xApplicationListActionDescriptions, lib, "gtk_application_list_action_descriptions")
	core.PuregoSafeRegister(&xApplicationRemoveWindow, lib, "gtk_application_remove_window")
	core.PuregoSafeRegister(&xApplicationSetAccelsForAction, lib, "gtk_application_set_accels_for_action")
	core.PuregoSafeRegister(&xApplicationSetMenubar, lib, "gtk_application_set_menubar")
	core.PuregoSafeRegister(&xApplicationUninhibit, lib, "gtk_application_uninhibit")

}
