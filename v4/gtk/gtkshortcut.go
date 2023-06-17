// Package gtk was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gtk

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/glib"
	"github.com/jwijenbergh/puregotk/v4/gobject"
)

type ShortcutClass struct {
	ParentClass uintptr
}

// A `GtkShortcut` describes a keyboard shortcut.
//
// It contains a description of how to trigger the shortcut via a
// [class@Gtk.ShortcutTrigger] and a way to activate the shortcut
// on a widget via a [class@Gtk.ShortcutAction].
//
// The actual work is usually done via [class@Gtk.ShortcutController],
// which decides if and when to activate a shortcut. Using that controller
// directly however is rarely necessary as various higher level
// convenience APIs exist on `GtkWidget`s that make it easier to use
// shortcuts in GTK.
//
// `GtkShortcut` does provide functionality to make it easy for users
// to work with shortcuts, either by providing informational strings
// for display purposes or by allowing shortcuts to be configured.
type Shortcut struct {
	gobject.Object
}

func ShortcutNewFromInternalPtr(ptr uintptr) *Shortcut {
	cls := &Shortcut{}
	cls.Ptr = ptr
	return cls
}

var xNewShortcut func(uintptr, uintptr) uintptr

// Creates a new `GtkShortcut` that is triggered by
// @trigger and then activates @action.
func NewShortcut(TriggerVar *ShortcutTrigger, ActionVar *ShortcutAction) *Shortcut {
	NewShortcutPtr := xNewShortcut(TriggerVar.GoPointer(), ActionVar.GoPointer())
	if NewShortcutPtr == 0 {
		return nil
	}

	NewShortcutCls := &Shortcut{}
	NewShortcutCls.Ptr = NewShortcutPtr
	return NewShortcutCls
}

var xNewWithArgumentsShortcut func(uintptr, uintptr, string, ...interface{}) uintptr

// Creates a new `GtkShortcut` that is triggered by @trigger and then activates
// @action with arguments given by @format_string.
func NewWithArgumentsShortcut(TriggerVar *ShortcutTrigger, ActionVar *ShortcutAction, FormatStringVar string, varArgs ...interface{}) *Shortcut {
	NewWithArgumentsShortcutPtr := xNewWithArgumentsShortcut(TriggerVar.GoPointer(), ActionVar.GoPointer(), FormatStringVar, varArgs...)
	if NewWithArgumentsShortcutPtr == 0 {
		return nil
	}

	NewWithArgumentsShortcutCls := &Shortcut{}
	NewWithArgumentsShortcutCls.Ptr = NewWithArgumentsShortcutPtr
	return NewWithArgumentsShortcutCls
}

var xShortcutGetAction func(uintptr) uintptr

// Gets the action that is activated by this shortcut.
func (x *Shortcut) GetAction() *ShortcutAction {

	GetActionPtr := xShortcutGetAction(x.GoPointer())
	if GetActionPtr == 0 {
		return nil
	}

	gobject.IncreaseRef(GetActionPtr)

	GetActionCls := &ShortcutAction{}
	GetActionCls.Ptr = GetActionPtr
	return GetActionCls

}

var xShortcutGetArguments func(uintptr) *glib.Variant

// Gets the arguments that are passed when activating the shortcut.
func (x *Shortcut) GetArguments() *glib.Variant {

	return xShortcutGetArguments(x.GoPointer())

}

var xShortcutGetTrigger func(uintptr) uintptr

// Gets the trigger used to trigger @self.
func (x *Shortcut) GetTrigger() *ShortcutTrigger {

	GetTriggerPtr := xShortcutGetTrigger(x.GoPointer())
	if GetTriggerPtr == 0 {
		return nil
	}

	gobject.IncreaseRef(GetTriggerPtr)

	GetTriggerCls := &ShortcutTrigger{}
	GetTriggerCls.Ptr = GetTriggerPtr
	return GetTriggerCls

}

var xShortcutSetAction func(uintptr, uintptr)

// Sets the new action for @self to be @action.
func (x *Shortcut) SetAction(ActionVar *ShortcutAction) {

	xShortcutSetAction(x.GoPointer(), ActionVar.GoPointer())

}

var xShortcutSetArguments func(uintptr, *glib.Variant)

// Sets the arguments to pass when activating the shortcut.
func (x *Shortcut) SetArguments(ArgsVar *glib.Variant) {

	xShortcutSetArguments(x.GoPointer(), ArgsVar)

}

var xShortcutSetTrigger func(uintptr, uintptr)

// Sets the new trigger for @self to be @trigger.
func (x *Shortcut) SetTrigger(TriggerVar *ShortcutTrigger) {

	xShortcutSetTrigger(x.GoPointer(), TriggerVar.GoPointer())

}

func (c *Shortcut) GoPointer() uintptr {
	return c.Ptr
}

func (c *Shortcut) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GTK"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xNewShortcut, lib, "gtk_shortcut_new")
	core.PuregoSafeRegister(&xNewWithArgumentsShortcut, lib, "gtk_shortcut_new_with_arguments")

	core.PuregoSafeRegister(&xShortcutGetAction, lib, "gtk_shortcut_get_action")
	core.PuregoSafeRegister(&xShortcutGetArguments, lib, "gtk_shortcut_get_arguments")
	core.PuregoSafeRegister(&xShortcutGetTrigger, lib, "gtk_shortcut_get_trigger")
	core.PuregoSafeRegister(&xShortcutSetAction, lib, "gtk_shortcut_set_action")
	core.PuregoSafeRegister(&xShortcutSetArguments, lib, "gtk_shortcut_set_arguments")
	core.PuregoSafeRegister(&xShortcutSetTrigger, lib, "gtk_shortcut_set_trigger")

}
