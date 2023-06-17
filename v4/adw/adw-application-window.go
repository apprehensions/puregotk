// Package adw was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package adw

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/gdk"
	"github.com/jwijenbergh/puregotk/v4/gio"
	"github.com/jwijenbergh/puregotk/v4/glib"
	"github.com/jwijenbergh/puregotk/v4/gobject"
	"github.com/jwijenbergh/puregotk/v4/gsk"
	"github.com/jwijenbergh/puregotk/v4/gtk"
)

type ApplicationWindowClass struct {
	ParentClass uintptr

	Padding uintptr
}

// A freeform application window.
//
// &lt;picture&gt;
//
//	&lt;source srcset="application-window-dark.png" media="(prefers-color-scheme: dark)"&gt;
//	&lt;img src="application-window.png" alt="application-window"&gt;
//
// &lt;/picture&gt;
//
// `AdwApplicationWindow` is a [class@Gtk.ApplicationWindow] subclass providing
// the same features as [class@Window].
//
// See [class@Window] for details.
//
// Using [property@Gtk.Application:menubar] is not supported and may result in
// visual glitches.
type ApplicationWindow struct {
	gtk.ApplicationWindow
}

func ApplicationWindowNewFromInternalPtr(ptr uintptr) *ApplicationWindow {
	cls := &ApplicationWindow{}
	cls.Ptr = ptr
	return cls
}

var xNewApplicationWindow func(uintptr) uintptr

// Creates a new `AdwApplicationWindow` for @app.
func NewApplicationWindow(AppVar *gtk.Application) *gtk.Widget {
	NewApplicationWindowPtr := xNewApplicationWindow(AppVar.GoPointer())
	if NewApplicationWindowPtr == 0 {
		return nil
	}

	gobject.IncreaseRef(NewApplicationWindowPtr)

	NewApplicationWindowCls := &gtk.Widget{}
	NewApplicationWindowCls.Ptr = NewApplicationWindowPtr
	return NewApplicationWindowCls
}

var xApplicationWindowGetContent func(uintptr) uintptr

// Gets the content widget of @self.
//
// This method should always be used instead of [method@Gtk.Window.get_child].
func (x *ApplicationWindow) GetContent() *gtk.Widget {

	GetContentPtr := xApplicationWindowGetContent(x.GoPointer())
	if GetContentPtr == 0 {
		return nil
	}

	gobject.IncreaseRef(GetContentPtr)

	GetContentCls := &gtk.Widget{}
	GetContentCls.Ptr = GetContentPtr
	return GetContentCls

}

var xApplicationWindowSetContent func(uintptr, uintptr)

// Sets the content widget of @self.
//
// This method should always be used instead of [method@Gtk.Window.set_child].
func (x *ApplicationWindow) SetContent(ContentVar *gtk.Widget) {

	xApplicationWindowSetContent(x.GoPointer(), ContentVar.GoPointer())

}

func (c *ApplicationWindow) GoPointer() uintptr {
	return c.Ptr
}

func (c *ApplicationWindow) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// Emits the #GActionGroup::action-added signal on @action_group.
//
// This function should only be called by #GActionGroup implementations.
func (x *ApplicationWindow) ActionAdded(ActionNameVar string) {

	gio.XGActionGroupActionAdded(x.GoPointer(), ActionNameVar)

}

// Emits the #GActionGroup::action-enabled-changed signal on @action_group.
//
// This function should only be called by #GActionGroup implementations.
func (x *ApplicationWindow) ActionEnabledChanged(ActionNameVar string, EnabledVar bool) {

	gio.XGActionGroupActionEnabledChanged(x.GoPointer(), ActionNameVar, EnabledVar)

}

// Emits the #GActionGroup::action-removed signal on @action_group.
//
// This function should only be called by #GActionGroup implementations.
func (x *ApplicationWindow) ActionRemoved(ActionNameVar string) {

	gio.XGActionGroupActionRemoved(x.GoPointer(), ActionNameVar)

}

// Emits the #GActionGroup::action-state-changed signal on @action_group.
//
// This function should only be called by #GActionGroup implementations.
func (x *ApplicationWindow) ActionStateChanged(ActionNameVar string, StateVar *glib.Variant) {

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
func (x *ApplicationWindow) ActivateAction(ActionNameVar string, ParameterVar *glib.Variant) {

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
func (x *ApplicationWindow) ChangeActionState(ActionNameVar string, ValueVar *glib.Variant) {

	gio.XGActionGroupChangeActionState(x.GoPointer(), ActionNameVar, ValueVar)

}

// Checks if the named action within @action_group is currently enabled.
//
// An action must be enabled in order to be activated or in order to
// have its state changed from outside callers.
func (x *ApplicationWindow) GetActionEnabled(ActionNameVar string) bool {

	return gio.XGActionGroupGetActionEnabled(x.GoPointer(), ActionNameVar)

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
func (x *ApplicationWindow) GetActionParameterType(ActionNameVar string) *glib.VariantType {

	return gio.XGActionGroupGetActionParameterType(x.GoPointer(), ActionNameVar)

}

// Queries the current state of the named action within @action_group.
//
// If the action is not stateful then %NULL will be returned.  If the
// action is stateful then the type of the return value is the type
// given by g_action_group_get_action_state_type().
//
// The return value (if non-%NULL) should be freed with
// g_variant_unref() when it is no longer required.
func (x *ApplicationWindow) GetActionState(ActionNameVar string) *glib.Variant {

	return gio.XGActionGroupGetActionState(x.GoPointer(), ActionNameVar)

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
func (x *ApplicationWindow) GetActionStateHint(ActionNameVar string) *glib.Variant {

	return gio.XGActionGroupGetActionStateHint(x.GoPointer(), ActionNameVar)

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
func (x *ApplicationWindow) GetActionStateType(ActionNameVar string) *glib.VariantType {

	return gio.XGActionGroupGetActionStateType(x.GoPointer(), ActionNameVar)

}

// Checks if the named action exists within @action_group.
func (x *ApplicationWindow) HasAction(ActionNameVar string) bool {

	return gio.XGActionGroupHasAction(x.GoPointer(), ActionNameVar)

}

// Lists the actions contained within @action_group.
//
// The caller is responsible for freeing the list with g_strfreev() when
// it is no longer required.
func (x *ApplicationWindow) ListActions() uintptr {

	return gio.XGActionGroupListActions(x.GoPointer())

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
func (x *ApplicationWindow) QueryAction(ActionNameVar string, EnabledVar bool, ParameterTypeVar **glib.VariantType, StateTypeVar **glib.VariantType, StateHintVar **glib.Variant, StateVar **glib.Variant) bool {

	return gio.XGActionGroupQueryAction(x.GoPointer(), ActionNameVar, EnabledVar, ParameterTypeVar, StateTypeVar, StateHintVar, StateVar)

}

// Adds an action to the @action_map.
//
// If the action map already contains an action with the same name
// as @action then the old action is dropped from the action map.
//
// The action map takes its own reference on @action.
func (x *ApplicationWindow) AddAction(ActionVar gio.Action) {

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
func (x *ApplicationWindow) AddActionEntries(EntriesVar uintptr, NEntriesVar int32, UserDataVar uintptr) {

	gio.XGActionMapAddActionEntries(x.GoPointer(), EntriesVar, NEntriesVar, UserDataVar)

}

// Looks up the action with the name @action_name in @action_map.
//
// If no such action exists, returns %NULL.
func (x *ApplicationWindow) LookupAction(ActionNameVar string) *gio.ActionBase {

	LookupActionPtr := gio.XGActionMapLookupAction(x.GoPointer(), ActionNameVar)
	if LookupActionPtr == 0 {
		return nil
	}

	gobject.IncreaseRef(LookupActionPtr)

	LookupActionCls := &gio.ActionBase{}
	LookupActionCls.Ptr = LookupActionPtr
	return LookupActionCls

}

// Removes the named action from the action map.
//
// If no action of this name is in the map then nothing happens.
func (x *ApplicationWindow) RemoveAction(ActionNameVar string) {

	gio.XGActionMapRemoveAction(x.GoPointer(), ActionNameVar)

}

// Retrieves the `GtkAccessibleRole` for the given `GtkAccessible`.
func (x *ApplicationWindow) GetAccessibleRole() gtk.AccessibleRole {

	return gtk.XGtkAccessibleGetAccessibleRole(x.GoPointer())

}

// Resets the accessible @property to its default value.
func (x *ApplicationWindow) ResetProperty(PropertyVar gtk.AccessibleProperty) {

	gtk.XGtkAccessibleResetProperty(x.GoPointer(), PropertyVar)

}

// Resets the accessible @relation to its default value.
func (x *ApplicationWindow) ResetRelation(RelationVar gtk.AccessibleRelation) {

	gtk.XGtkAccessibleResetRelation(x.GoPointer(), RelationVar)

}

// Resets the accessible @state to its default value.
func (x *ApplicationWindow) ResetState(StateVar gtk.AccessibleState) {

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
func (x *ApplicationWindow) UpdateProperty(FirstPropertyVar gtk.AccessibleProperty, varArgs ...interface{}) {

	gtk.XGtkAccessibleUpdateProperty(x.GoPointer(), FirstPropertyVar, varArgs...)

}

// Updates an array of accessible properties.
//
// This function should be called by `GtkWidget` types whenever an accessible
// property change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *ApplicationWindow) UpdatePropertyValue(NPropertiesVar int32, PropertiesVar uintptr, ValuesVar uintptr) {

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
func (x *ApplicationWindow) UpdateRelation(FirstRelationVar gtk.AccessibleRelation, varArgs ...interface{}) {

	gtk.XGtkAccessibleUpdateRelation(x.GoPointer(), FirstRelationVar, varArgs...)

}

// Updates an array of accessible relations.
//
// This function should be called by `GtkWidget` types whenever an accessible
// relation change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *ApplicationWindow) UpdateRelationValue(NRelationsVar int32, RelationsVar uintptr, ValuesVar uintptr) {

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
func (x *ApplicationWindow) UpdateState(FirstStateVar gtk.AccessibleState, varArgs ...interface{}) {

	gtk.XGtkAccessibleUpdateState(x.GoPointer(), FirstStateVar, varArgs...)

}

// Updates an array of accessible states.
//
// This function should be called by `GtkWidget` types whenever an accessible
// state change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *ApplicationWindow) UpdateStateValue(NStatesVar int32, StatesVar uintptr, ValuesVar uintptr) {

	gtk.XGtkAccessibleUpdateStateValue(x.GoPointer(), NStatesVar, StatesVar, ValuesVar)

}

// Gets the ID of the @buildable object.
//
// `GtkBuilder` sets the name based on the ID attribute
// of the &lt;object&gt; tag used to construct the @buildable.
func (x *ApplicationWindow) GetBuildableId() string {

	return gtk.XGtkBuildableGetBuildableId(x.GoPointer())

}

// Returns the renderer that is used for this `GtkNative`.
func (x *ApplicationWindow) GetRenderer() *gsk.Renderer {

	GetRendererPtr := gtk.XGtkNativeGetRenderer(x.GoPointer())
	if GetRendererPtr == 0 {
		return nil
	}

	gobject.IncreaseRef(GetRendererPtr)

	GetRendererCls := &gsk.Renderer{}
	GetRendererCls.Ptr = GetRendererPtr
	return GetRendererCls

}

// Returns the surface of this `GtkNative`.
func (x *ApplicationWindow) GetSurface() *gdk.Surface {

	GetSurfacePtr := gtk.XGtkNativeGetSurface(x.GoPointer())
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
func (x *ApplicationWindow) GetSurfaceTransform(XVar float64, YVar float64) {

	gtk.XGtkNativeGetSurfaceTransform(x.GoPointer(), XVar, YVar)

}

// Realizes a `GtkNative`.
//
// This should only be used by subclasses.
func (x *ApplicationWindow) Realize() {

	gtk.XGtkNativeRealize(x.GoPointer())

}

// Unrealizes a `GtkNative`.
//
// This should only be used by subclasses.
func (x *ApplicationWindow) Unrealize() {

	gtk.XGtkNativeUnrealize(x.GoPointer())

}

// Returns the display that this `GtkRoot` is on.
func (x *ApplicationWindow) GetDisplay() *gdk.Display {

	GetDisplayPtr := gtk.XGtkRootGetDisplay(x.GoPointer())
	if GetDisplayPtr == 0 {
		return nil
	}

	gobject.IncreaseRef(GetDisplayPtr)

	GetDisplayCls := &gdk.Display{}
	GetDisplayCls.Ptr = GetDisplayPtr
	return GetDisplayCls

}

// Retrieves the current focused widget within the root.
//
// Note that this is the widget that would have the focus
// if the root is active; if the root is not focused then
// `gtk_widget_has_focus (widget)` will be %FALSE for the
// widget.
func (x *ApplicationWindow) GetFocus() *gtk.Widget {

	GetFocusPtr := gtk.XGtkRootGetFocus(x.GoPointer())
	if GetFocusPtr == 0 {
		return nil
	}

	gobject.IncreaseRef(GetFocusPtr)

	GetFocusCls := &gtk.Widget{}
	GetFocusCls.Ptr = GetFocusPtr
	return GetFocusCls

}

// If @focus is not the current focus widget, and is focusable, sets
// it as the focus widget for the root.
//
// If @focus is %NULL, unsets the focus widget for the root.
//
// To set the focus to a particular widget in the root, it is usually
// more convenient to use [method@Gtk.Widget.grab_focus] instead of
// this function.
func (x *ApplicationWindow) SetFocus(FocusVar *gtk.Widget) {

	gtk.XGtkRootSetFocus(x.GoPointer(), FocusVar.GoPointer())

}

func init() {
	lib, err := purego.Dlopen(core.GetPath("ADW"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xNewApplicationWindow, lib, "adw_application_window_new")

	core.PuregoSafeRegister(&xApplicationWindowGetContent, lib, "adw_application_window_get_content")
	core.PuregoSafeRegister(&xApplicationWindowSetContent, lib, "adw_application_window_set_content")

}
