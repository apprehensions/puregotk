// Package gio was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gio

import (
	"unsafe"

	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/glib"
	"github.com/jwijenbergh/puregotk/v4/gobject"
)

type MenuAttributeIterClass struct {
	ParentClass uintptr
}

func (x *MenuAttributeIterClass) GoPointer() uintptr {
	return uintptr(unsafe.Pointer(x))
}

type MenuAttributeIterPrivate struct {
}

func (x *MenuAttributeIterPrivate) GoPointer() uintptr {
	return uintptr(unsafe.Pointer(x))
}

type MenuLinkIterClass struct {
	ParentClass uintptr
}

func (x *MenuLinkIterClass) GoPointer() uintptr {
	return uintptr(unsafe.Pointer(x))
}

type MenuLinkIterPrivate struct {
}

func (x *MenuLinkIterPrivate) GoPointer() uintptr {
	return uintptr(unsafe.Pointer(x))
}

type MenuModelClass struct {
	ParentClass uintptr
}

func (x *MenuModelClass) GoPointer() uintptr {
	return uintptr(unsafe.Pointer(x))
}

type MenuModelPrivate struct {
}

func (x *MenuModelPrivate) GoPointer() uintptr {
	return uintptr(unsafe.Pointer(x))
}

const (
	// The menu item attribute which holds the action name of the item.  Action
	// names are namespaced with an identifier for the action group in which the
	// action resides. For example, "win." for window-specific actions and "app."
	// for application-wide actions.
	//
	// See also g_menu_model_get_item_attribute() and g_menu_item_set_attribute().
	MENU_ATTRIBUTE_ACTION string = "action"
	// The menu item attribute that holds the namespace for all action names in
	// menus that are linked from this item.
	MENU_ATTRIBUTE_ACTION_NAMESPACE string = "action-namespace"
	// The menu item attribute which holds the icon of the item.
	//
	// The icon is stored in the format returned by g_icon_serialize().
	//
	// This attribute is intended only to represent 'noun' icons such as
	// favicons for a webpage, or application icons.  It should not be used
	// for 'verbs' (ie: stock icons).
	MENU_ATTRIBUTE_ICON string = "icon"
	// The menu item attribute which holds the label of the item.
	MENU_ATTRIBUTE_LABEL string = "label"
	// The menu item attribute which holds the target with which the item's action
	// will be activated.
	//
	// See also g_menu_item_set_action_and_target()
	MENU_ATTRIBUTE_TARGET string = "target"
	// The name of the link that associates a menu item with a section.  The linked
	// menu will usually be shown in place of the menu item, using the item's label
	// as a header.
	//
	// See also g_menu_item_set_link().
	MENU_LINK_SECTION string = "section"
	// The name of the link that associates a menu item with a submenu.
	//
	// See also g_menu_item_set_link().
	MENU_LINK_SUBMENU string = "submenu"
)

// #GMenuAttributeIter is an opaque structure type.  You must access it
// using the functions below.
type MenuAttributeIter struct {
	gobject.Object
}

func MenuAttributeIterNewFromInternalPtr(ptr uintptr) *MenuAttributeIter {
	cls := &MenuAttributeIter{}
	cls.Ptr = ptr
	return cls
}

var xMenuAttributeIterGetName func(uintptr) string

// Gets the name of the attribute at the current iterator position, as
// a string.
//
// The iterator is not advanced.
func (x *MenuAttributeIter) GetName() string {

	cret := xMenuAttributeIterGetName(x.GoPointer())
	return cret
}

var xMenuAttributeIterGetNext func(uintptr, string, **glib.Variant) bool

// This function combines g_menu_attribute_iter_next() with
// g_menu_attribute_iter_get_name() and g_menu_attribute_iter_get_value().
//
// First the iterator is advanced to the next (possibly first) attribute.
// If that fails, then %FALSE is returned and there are no other
// effects.
//
// If successful, @name and @value are set to the name and value of the
// attribute that has just been advanced to.  At this point,
// g_menu_attribute_iter_get_name() and g_menu_attribute_iter_get_value() will
// return the same values again.
//
// The value returned in @name remains valid for as long as the iterator
// remains at the current position.  The value returned in @value must
// be unreffed using g_variant_unref() when it is no longer in use.
func (x *MenuAttributeIter) GetNext(OutNameVar string, ValueVar **glib.Variant) bool {

	cret := xMenuAttributeIterGetNext(x.GoPointer(), OutNameVar, ValueVar)
	return cret
}

var xMenuAttributeIterGetValue func(uintptr) *glib.Variant

// Gets the value of the attribute at the current iterator position.
//
// The iterator is not advanced.
func (x *MenuAttributeIter) GetValue() *glib.Variant {

	cret := xMenuAttributeIterGetValue(x.GoPointer())
	return cret
}

var xMenuAttributeIterNext func(uintptr) bool

// Attempts to advance the iterator to the next (possibly first)
// attribute.
//
// %TRUE is returned on success, or %FALSE if there are no more
// attributes.
//
// You must call this function when you first acquire the iterator
// to advance it to the first attribute (and determine if the first
// attribute exists at all).
func (x *MenuAttributeIter) Next() bool {

	cret := xMenuAttributeIterNext(x.GoPointer())
	return cret
}

func (c *MenuAttributeIter) GoPointer() uintptr {
	return c.Ptr
}

func (c *MenuAttributeIter) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// #GMenuLinkIter is an opaque structure type.  You must access it using
// the functions below.
type MenuLinkIter struct {
	gobject.Object
}

func MenuLinkIterNewFromInternalPtr(ptr uintptr) *MenuLinkIter {
	cls := &MenuLinkIter{}
	cls.Ptr = ptr
	return cls
}

var xMenuLinkIterGetName func(uintptr) string

// Gets the name of the link at the current iterator position.
//
// The iterator is not advanced.
func (x *MenuLinkIter) GetName() string {

	cret := xMenuLinkIterGetName(x.GoPointer())
	return cret
}

var xMenuLinkIterGetNext func(uintptr, string, *uintptr) bool

// This function combines g_menu_link_iter_next() with
// g_menu_link_iter_get_name() and g_menu_link_iter_get_value().
//
// First the iterator is advanced to the next (possibly first) link.
// If that fails, then %FALSE is returned and there are no other effects.
//
// If successful, @out_link and @value are set to the name and #GMenuModel
// of the link that has just been advanced to.  At this point,
// g_menu_link_iter_get_name() and g_menu_link_iter_get_value() will return the
// same values again.
//
// The value returned in @out_link remains valid for as long as the iterator
// remains at the current position.  The value returned in @value must
// be unreffed using g_object_unref() when it is no longer in use.
func (x *MenuLinkIter) GetNext(OutLinkVar string, ValueVar **MenuModel) bool {

	cret := xMenuLinkIterGetNext(x.GoPointer(), OutLinkVar, gobject.ConvertPtr(ValueVar))
	return cret
}

var xMenuLinkIterGetValue func(uintptr) uintptr

// Gets the linked #GMenuModel at the current iterator position.
//
// The iterator is not advanced.
func (x *MenuLinkIter) GetValue() *MenuModel {
	var cls *MenuModel

	cret := xMenuLinkIterGetValue(x.GoPointer())

	if cret == 0 {
		return nil
	}
	cls = &MenuModel{}
	cls.Ptr = cret
	return cls
}

var xMenuLinkIterNext func(uintptr) bool

// Attempts to advance the iterator to the next (possibly first)
// link.
//
// %TRUE is returned on success, or %FALSE if there are no more links.
//
// You must call this function when you first acquire the iterator to
// advance it to the first link (and determine if the first link exists
// at all).
func (x *MenuLinkIter) Next() bool {

	cret := xMenuLinkIterNext(x.GoPointer())
	return cret
}

func (c *MenuLinkIter) GoPointer() uintptr {
	return c.Ptr
}

func (c *MenuLinkIter) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// #GMenuModel represents the contents of a menu -- an ordered list of
// menu items. The items are associated with actions, which can be
// activated through them. Items can be grouped in sections, and may
// have submenus associated with them. Both items and sections usually
// have some representation data, such as labels or icons. The type of
// the associated action (ie whether it is stateful, and what kind of
// state it has) can influence the representation of the item.
//
// The conceptual model of menus in #GMenuModel is hierarchical:
// sections and submenus are again represented by #GMenuModels.
// Menus themselves do not define their own roles. Rather, the role
// of a particular #GMenuModel is defined by the item that references
// it (or, in the case of the 'root' menu, is defined by the context
// in which it is used).
//
// As an example, consider the visible portions of this menu:
//
// ## An example menu # {#menu-example}
//
// ![](menu-example.png)
//
// There are 8 "menus" visible in the screenshot: one menubar, two
// submenus and 5 sections:
//
// - the toplevel menubar (containing 4 items)
// - the View submenu (containing 3 sections)
// - the first section of the View submenu (containing 2 items)
// - the second section of the View submenu (containing 1 item)
// - the final section of the View submenu (containing 1 item)
// - the Highlight Mode submenu (containing 2 sections)
// - the Sources section (containing 2 items)
// - the Markup section (containing 2 items)
//
// The [example][menu-model] illustrates the conceptual connection between
// these 8 menus. Each large block in the figure represents a menu and the
// smaller blocks within the large block represent items in that menu. Some
// items contain references to other menus.
//
// ## A menu example # {#menu-model}
//
// ![](menu-model.png)
//
// Notice that the separators visible in the [example][menu-example]
// appear nowhere in the [menu model][menu-model]. This is because
// separators are not explicitly represented in the menu model. Instead,
// a separator is inserted between any two non-empty sections of a menu.
// Section items can have labels just like any other item. In that case,
// a display system may show a section header instead of a separator.
//
// The motivation for this abstract model of application controls is
// that modern user interfaces tend to make these controls available
// outside the application. Examples include global menus, jumplists,
// dash boards, etc. To support such uses, it is necessary to 'export'
// information about actions and their representation in menus, which
// is exactly what the [GActionGroup exporter][gio-GActionGroup-exporter]
// and the [GMenuModel exporter][gio-GMenuModel-exporter] do for
// #GActionGroup and #GMenuModel. The client-side counterparts to
// make use of the exported information are #GDBusActionGroup and
// #GDBusMenuModel.
//
// The API of #GMenuModel is very generic, with iterators for the
// attributes and links of an item, see g_menu_model_iterate_item_attributes()
// and g_menu_model_iterate_item_links(). The 'standard' attributes and
// link types have predefined names: %G_MENU_ATTRIBUTE_LABEL,
// %G_MENU_ATTRIBUTE_ACTION, %G_MENU_ATTRIBUTE_TARGET, %G_MENU_LINK_SECTION
// and %G_MENU_LINK_SUBMENU.
//
// Items in a #GMenuModel represent active controls if they refer to
// an action that can get activated when the user interacts with the
// menu item. The reference to the action is encoded by the string id
// in the %G_MENU_ATTRIBUTE_ACTION attribute. An action id uniquely
// identifies an action in an action group. Which action group(s) provide
// actions depends on the context in which the menu model is used.
// E.g. when the model is exported as the application menu of a
// #GtkApplication, actions can be application-wide or window-specific
// (and thus come from two different action groups). By convention, the
// application-wide actions have names that start with "app.", while the
// names of window-specific actions start with "win.".
//
// While a wide variety of stateful actions is possible, the following
// is the minimum that is expected to be supported by all users of exported
// menu information:
// - an action with no parameter type and no state
// - an action with no parameter type and boolean state
// - an action with string parameter type and string state
//
// ## Stateless
//
// A stateless action typically corresponds to an ordinary menu item.
//
// Selecting such a menu item will activate the action (with no parameter).
//
// ## Boolean State
//
// An action with a boolean state will most typically be used with a "toggle"
// or "switch" menu item. The state can be set directly, but activating the
// action (with no parameter) results in the state being toggled.
//
// Selecting a toggle menu item will activate the action. The menu item should
// be rendered as "checked" when the state is true.
//
// ## String Parameter and State
//
// Actions with string parameters and state will most typically be used to
// represent an enumerated choice over the items available for a group of
// radio menu items. Activating the action with a string parameter is
// equivalent to setting that parameter as the state.
//
// Radio menu items, in addition to being associated with the action, will
// have a target value. Selecting that menu item will result in activation
// of the action with the target value as the parameter. The menu item should
// be rendered as "selected" when the state of the action is equal to the
// target value of the menu item.
type MenuModel struct {
	gobject.Object
}

func MenuModelNewFromInternalPtr(ptr uintptr) *MenuModel {
	cls := &MenuModel{}
	cls.Ptr = ptr
	return cls
}

var xMenuModelGetItemAttribute func(uintptr, int, string, string, ...interface{}) bool

// Queries item at position @item_index in @model for the attribute
// specified by @attribute.
//
// If the attribute exists and matches the #GVariantType corresponding
// to @format_string then @format_string is used to deconstruct the
// value into the positional parameters and %TRUE is returned.
//
// If the attribute does not exist, or it does exist but has the wrong
// type, then the positional parameters are ignored and %FALSE is
// returned.
//
// This function is a mix of g_menu_model_get_item_attribute_value() and
// g_variant_get(), followed by a g_variant_unref().  As such,
// @format_string must make a complete copy of the data (since the
// #GVariant may go away after the call to g_variant_unref()).  In
// particular, no '&amp;' characters are allowed in @format_string.
func (x *MenuModel) GetItemAttribute(ItemIndexVar int, AttributeVar string, FormatStringVar string, varArgs ...interface{}) bool {

	cret := xMenuModelGetItemAttribute(x.GoPointer(), ItemIndexVar, AttributeVar, FormatStringVar, varArgs...)
	return cret
}

var xMenuModelGetItemAttributeValue func(uintptr, int, string, *glib.VariantType) *glib.Variant

// Queries the item at position @item_index in @model for the attribute
// specified by @attribute.
//
// If @expected_type is non-%NULL then it specifies the expected type of
// the attribute.  If it is %NULL then any type will be accepted.
//
// If the attribute exists and matches @expected_type (or if the
// expected type is unspecified) then the value is returned.
//
// If the attribute does not exist, or does not match the expected type
// then %NULL is returned.
func (x *MenuModel) GetItemAttributeValue(ItemIndexVar int, AttributeVar string, ExpectedTypeVar *glib.VariantType) *glib.Variant {

	cret := xMenuModelGetItemAttributeValue(x.GoPointer(), ItemIndexVar, AttributeVar, ExpectedTypeVar)
	return cret
}

var xMenuModelGetItemLink func(uintptr, int, string) uintptr

// Queries the item at position @item_index in @model for the link
// specified by @link.
//
// If the link exists, the linked #GMenuModel is returned.  If the link
// does not exist, %NULL is returned.
func (x *MenuModel) GetItemLink(ItemIndexVar int, LinkVar string) *MenuModel {
	var cls *MenuModel

	cret := xMenuModelGetItemLink(x.GoPointer(), ItemIndexVar, LinkVar)

	if cret == 0 {
		return nil
	}
	cls = &MenuModel{}
	cls.Ptr = cret
	return cls
}

var xMenuModelGetNItems func(uintptr) int

// Query the number of items in @model.
func (x *MenuModel) GetNItems() int {

	cret := xMenuModelGetNItems(x.GoPointer())
	return cret
}

var xMenuModelIsMutable func(uintptr) bool

// Queries if @model is mutable.
//
// An immutable #GMenuModel will never emit the #GMenuModel::items-changed
// signal. Consumers of the model may make optimisations accordingly.
func (x *MenuModel) IsMutable() bool {

	cret := xMenuModelIsMutable(x.GoPointer())
	return cret
}

var xMenuModelItemsChanged func(uintptr, int, int, int)

// Requests emission of the #GMenuModel::items-changed signal on @model.
//
// This function should never be called except by #GMenuModel
// subclasses.  Any other calls to this function will very likely lead
// to a violation of the interface of the model.
//
// The implementation should update its internal representation of the
// menu before emitting the signal.  The implementation should further
// expect to receive queries about the new state of the menu (and
// particularly added menu items) while signal handlers are running.
//
// The implementation must dispatch this call directly from a mainloop
// entry and not in response to calls -- particularly those from the
// #GMenuModel API.  Said another way: the menu must not change while
// user code is running without returning to the mainloop.
func (x *MenuModel) ItemsChanged(PositionVar int, RemovedVar int, AddedVar int) {

	xMenuModelItemsChanged(x.GoPointer(), PositionVar, RemovedVar, AddedVar)

}

var xMenuModelIterateItemAttributes func(uintptr, int) uintptr

// Creates a #GMenuAttributeIter to iterate over the attributes of
// the item at position @item_index in @model.
//
// You must free the iterator with g_object_unref() when you are done.
func (x *MenuModel) IterateItemAttributes(ItemIndexVar int) *MenuAttributeIter {
	var cls *MenuAttributeIter

	cret := xMenuModelIterateItemAttributes(x.GoPointer(), ItemIndexVar)

	if cret == 0 {
		return nil
	}
	cls = &MenuAttributeIter{}
	cls.Ptr = cret
	return cls
}

var xMenuModelIterateItemLinks func(uintptr, int) uintptr

// Creates a #GMenuLinkIter to iterate over the links of the item at
// position @item_index in @model.
//
// You must free the iterator with g_object_unref() when you are done.
func (x *MenuModel) IterateItemLinks(ItemIndexVar int) *MenuLinkIter {
	var cls *MenuLinkIter

	cret := xMenuModelIterateItemLinks(x.GoPointer(), ItemIndexVar)

	if cret == 0 {
		return nil
	}
	cls = &MenuLinkIter{}
	cls.Ptr = cret
	return cls
}

func (c *MenuModel) GoPointer() uintptr {
	return c.Ptr
}

func (c *MenuModel) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// Emitted when a change has occurred to the menu.
//
// The only changes that can occur to a menu is that items are removed
// or added.  Items may not change (except by being removed and added
// back in the same location).  This signal is capable of describing
// both of those changes (at the same time).
//
// The signal means that starting at the index @position, @removed
// items were removed and @added items were added in their place.  If
// @removed is zero then only items were added.  If @added is zero
// then only items were removed.
//
// As an example, if the menu contains items a, b, c, d (in that
// order) and the signal (2, 1, 3) occurs then the new composition of
// the menu will be a, b, _, _, _, d (with each _ representing some
// new item).
//
// Signal handlers may query the model (particularly the added items)
// and expect to see the results of the modification that is being
// reported.  The signal is emitted after the modification.
func (x *MenuModel) ConnectItemsChanged(cb func(MenuModel, int, int, int)) uint32 {
	fcb := func(clsPtr uintptr, PositionVarp int, RemovedVarp int, AddedVarp int) {
		fa := MenuModel{}
		fa.Ptr = clsPtr

		cb(fa, PositionVarp, RemovedVarp, AddedVarp)

	}
	return gobject.SignalConnect(x.GoPointer(), "items-changed", purego.NewCallback(fcb))
}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GIO"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xMenuAttributeIterGetName, lib, "g_menu_attribute_iter_get_name")
	core.PuregoSafeRegister(&xMenuAttributeIterGetNext, lib, "g_menu_attribute_iter_get_next")
	core.PuregoSafeRegister(&xMenuAttributeIterGetValue, lib, "g_menu_attribute_iter_get_value")
	core.PuregoSafeRegister(&xMenuAttributeIterNext, lib, "g_menu_attribute_iter_next")

	core.PuregoSafeRegister(&xMenuLinkIterGetName, lib, "g_menu_link_iter_get_name")
	core.PuregoSafeRegister(&xMenuLinkIterGetNext, lib, "g_menu_link_iter_get_next")
	core.PuregoSafeRegister(&xMenuLinkIterGetValue, lib, "g_menu_link_iter_get_value")
	core.PuregoSafeRegister(&xMenuLinkIterNext, lib, "g_menu_link_iter_next")

	core.PuregoSafeRegister(&xMenuModelGetItemAttribute, lib, "g_menu_model_get_item_attribute")
	core.PuregoSafeRegister(&xMenuModelGetItemAttributeValue, lib, "g_menu_model_get_item_attribute_value")
	core.PuregoSafeRegister(&xMenuModelGetItemLink, lib, "g_menu_model_get_item_link")
	core.PuregoSafeRegister(&xMenuModelGetNItems, lib, "g_menu_model_get_n_items")
	core.PuregoSafeRegister(&xMenuModelIsMutable, lib, "g_menu_model_is_mutable")
	core.PuregoSafeRegister(&xMenuModelItemsChanged, lib, "g_menu_model_items_changed")
	core.PuregoSafeRegister(&xMenuModelIterateItemAttributes, lib, "g_menu_model_iterate_item_attributes")
	core.PuregoSafeRegister(&xMenuModelIterateItemLinks, lib, "g_menu_model_iterate_item_links")

}
