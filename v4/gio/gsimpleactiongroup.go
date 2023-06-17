// Package gio was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gio

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/glib"
	"github.com/jwijenbergh/puregotk/v4/gobject"
)

type SimpleActionGroupClass struct {
	ParentClass uintptr

	Padding uintptr
}

type SimpleActionGroupPrivate struct {
}

// #GSimpleActionGroup is a hash table filled with #GAction objects,
// implementing the #GActionGroup and #GActionMap interfaces.
type SimpleActionGroup struct {
	gobject.Object
}

func SimpleActionGroupNewFromInternalPtr(ptr uintptr) *SimpleActionGroup {
	cls := &SimpleActionGroup{}
	cls.Ptr = ptr
	return cls
}

var xNewSimpleActionGroup func() uintptr

// Creates a new, empty, #GSimpleActionGroup.
func NewSimpleActionGroup() *SimpleActionGroup {
	NewSimpleActionGroupPtr := xNewSimpleActionGroup()
	if NewSimpleActionGroupPtr == 0 {
		return nil
	}

	NewSimpleActionGroupCls := &SimpleActionGroup{}
	NewSimpleActionGroupCls.Ptr = NewSimpleActionGroupPtr
	return NewSimpleActionGroupCls
}

var xSimpleActionGroupAddEntries func(uintptr, uintptr, int32, uintptr)

// A convenience function for creating multiple #GSimpleAction instances
// and adding them to the action group.
func (x *SimpleActionGroup) AddEntries(EntriesVar uintptr, NEntriesVar int32, UserDataVar uintptr) {

	xSimpleActionGroupAddEntries(x.GoPointer(), EntriesVar, NEntriesVar, UserDataVar)

}

var xSimpleActionGroupInsert func(uintptr, uintptr)

// Adds an action to the action group.
//
// If the action group already contains an action with the same name as
// @action then the old action is dropped from the group.
//
// The action group takes its own reference on @action.
func (x *SimpleActionGroup) Insert(ActionVar Action) {

	xSimpleActionGroupInsert(x.GoPointer(), ActionVar.GoPointer())

}

var xSimpleActionGroupLookup func(uintptr, string) uintptr

// Looks up the action with the name @action_name in the group.
//
// If no such action exists, returns %NULL.
func (x *SimpleActionGroup) Lookup(ActionNameVar string) *ActionBase {

	LookupPtr := xSimpleActionGroupLookup(x.GoPointer(), ActionNameVar)
	if LookupPtr == 0 {
		return nil
	}

	gobject.IncreaseRef(LookupPtr)

	LookupCls := &ActionBase{}
	LookupCls.Ptr = LookupPtr
	return LookupCls

}

var xSimpleActionGroupRemove func(uintptr, string)

// Removes the named action from the action group.
//
// If no action of this name is in the group then nothing happens.
func (x *SimpleActionGroup) Remove(ActionNameVar string) {

	xSimpleActionGroupRemove(x.GoPointer(), ActionNameVar)

}

func (c *SimpleActionGroup) GoPointer() uintptr {
	return c.Ptr
}

func (c *SimpleActionGroup) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// Emits the #GActionGroup::action-added signal on @action_group.
//
// This function should only be called by #GActionGroup implementations.
func (x *SimpleActionGroup) ActionAdded(ActionNameVar string) {

	XGActionGroupActionAdded(x.GoPointer(), ActionNameVar)

}

// Emits the #GActionGroup::action-enabled-changed signal on @action_group.
//
// This function should only be called by #GActionGroup implementations.
func (x *SimpleActionGroup) ActionEnabledChanged(ActionNameVar string, EnabledVar bool) {

	XGActionGroupActionEnabledChanged(x.GoPointer(), ActionNameVar, EnabledVar)

}

// Emits the #GActionGroup::action-removed signal on @action_group.
//
// This function should only be called by #GActionGroup implementations.
func (x *SimpleActionGroup) ActionRemoved(ActionNameVar string) {

	XGActionGroupActionRemoved(x.GoPointer(), ActionNameVar)

}

// Emits the #GActionGroup::action-state-changed signal on @action_group.
//
// This function should only be called by #GActionGroup implementations.
func (x *SimpleActionGroup) ActionStateChanged(ActionNameVar string, StateVar *glib.Variant) {

	XGActionGroupActionStateChanged(x.GoPointer(), ActionNameVar, StateVar)

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
func (x *SimpleActionGroup) ActivateAction(ActionNameVar string, ParameterVar *glib.Variant) {

	XGActionGroupActivateAction(x.GoPointer(), ActionNameVar, ParameterVar)

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
func (x *SimpleActionGroup) ChangeActionState(ActionNameVar string, ValueVar *glib.Variant) {

	XGActionGroupChangeActionState(x.GoPointer(), ActionNameVar, ValueVar)

}

// Checks if the named action within @action_group is currently enabled.
//
// An action must be enabled in order to be activated or in order to
// have its state changed from outside callers.
func (x *SimpleActionGroup) GetActionEnabled(ActionNameVar string) bool {

	return XGActionGroupGetActionEnabled(x.GoPointer(), ActionNameVar)

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
func (x *SimpleActionGroup) GetActionParameterType(ActionNameVar string) *glib.VariantType {

	return XGActionGroupGetActionParameterType(x.GoPointer(), ActionNameVar)

}

// Queries the current state of the named action within @action_group.
//
// If the action is not stateful then %NULL will be returned.  If the
// action is stateful then the type of the return value is the type
// given by g_action_group_get_action_state_type().
//
// The return value (if non-%NULL) should be freed with
// g_variant_unref() when it is no longer required.
func (x *SimpleActionGroup) GetActionState(ActionNameVar string) *glib.Variant {

	return XGActionGroupGetActionState(x.GoPointer(), ActionNameVar)

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
func (x *SimpleActionGroup) GetActionStateHint(ActionNameVar string) *glib.Variant {

	return XGActionGroupGetActionStateHint(x.GoPointer(), ActionNameVar)

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
func (x *SimpleActionGroup) GetActionStateType(ActionNameVar string) *glib.VariantType {

	return XGActionGroupGetActionStateType(x.GoPointer(), ActionNameVar)

}

// Checks if the named action exists within @action_group.
func (x *SimpleActionGroup) HasAction(ActionNameVar string) bool {

	return XGActionGroupHasAction(x.GoPointer(), ActionNameVar)

}

// Lists the actions contained within @action_group.
//
// The caller is responsible for freeing the list with g_strfreev() when
// it is no longer required.
func (x *SimpleActionGroup) ListActions() uintptr {

	return XGActionGroupListActions(x.GoPointer())

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
func (x *SimpleActionGroup) QueryAction(ActionNameVar string, EnabledVar bool, ParameterTypeVar **glib.VariantType, StateTypeVar **glib.VariantType, StateHintVar **glib.Variant, StateVar **glib.Variant) bool {

	return XGActionGroupQueryAction(x.GoPointer(), ActionNameVar, EnabledVar, ParameterTypeVar, StateTypeVar, StateHintVar, StateVar)

}

// Adds an action to the @action_map.
//
// If the action map already contains an action with the same name
// as @action then the old action is dropped from the action map.
//
// The action map takes its own reference on @action.
func (x *SimpleActionGroup) AddAction(ActionVar Action) {

	XGActionMapAddAction(x.GoPointer(), ActionVar.GoPointer())

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
func (x *SimpleActionGroup) AddActionEntries(EntriesVar uintptr, NEntriesVar int32, UserDataVar uintptr) {

	XGActionMapAddActionEntries(x.GoPointer(), EntriesVar, NEntriesVar, UserDataVar)

}

// Looks up the action with the name @action_name in @action_map.
//
// If no such action exists, returns %NULL.
func (x *SimpleActionGroup) LookupAction(ActionNameVar string) *ActionBase {

	LookupActionPtr := XGActionMapLookupAction(x.GoPointer(), ActionNameVar)
	if LookupActionPtr == 0 {
		return nil
	}

	gobject.IncreaseRef(LookupActionPtr)

	LookupActionCls := &ActionBase{}
	LookupActionCls.Ptr = LookupActionPtr
	return LookupActionCls

}

// Removes the named action from the action map.
//
// If no action of this name is in the map then nothing happens.
func (x *SimpleActionGroup) RemoveAction(ActionNameVar string) {

	XGActionMapRemoveAction(x.GoPointer(), ActionNameVar)

}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GIO"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xNewSimpleActionGroup, lib, "g_simple_action_group_new")

	core.PuregoSafeRegister(&xSimpleActionGroupAddEntries, lib, "g_simple_action_group_add_entries")
	core.PuregoSafeRegister(&xSimpleActionGroupInsert, lib, "g_simple_action_group_insert")
	core.PuregoSafeRegister(&xSimpleActionGroupLookup, lib, "g_simple_action_group_lookup")
	core.PuregoSafeRegister(&xSimpleActionGroupRemove, lib, "g_simple_action_group_remove")

}
