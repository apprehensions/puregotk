// Package gtk was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gtk

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/gobject"
)

type ListItemClass struct {
}

// `GtkListItem` is used by list widgets to represent items in a `GListModel`.
//
// The `GtkListItem`s are managed by the list widget (with its factory)
// and cannot be created by applications, but they need to be populated
// by application code. This is done by calling [method@Gtk.ListItem.set_child].
//
// `GtkListItem`s exist in 2 stages:
//
//  1. The unbound stage where the listitem is not currently connected to
//     an item in the list. In that case, the [property@Gtk.ListItem:item]
//     property is set to %NULL.
//
//  2. The bound stage where the listitem references an item from the list.
//     The [property@Gtk.ListItem:item] property is not %NULL.
type ListItem struct {
	gobject.Object
}

func ListItemNewFromInternalPtr(ptr uintptr) *ListItem {
	cls := &ListItem{}
	cls.Ptr = ptr
	return cls
}

var xListItemGetActivatable func(uintptr) bool

// Checks if a list item has been set to be activatable via
// gtk_list_item_set_activatable().
func (x *ListItem) GetActivatable() bool {

	return xListItemGetActivatable(x.GoPointer())

}

var xListItemGetChild func(uintptr) uintptr

// Gets the child previously set via gtk_list_item_set_child() or
// %NULL if none was set.
func (x *ListItem) GetChild() *Widget {

	GetChildPtr := xListItemGetChild(x.GoPointer())
	if GetChildPtr == 0 {
		return nil
	}

	gobject.IncreaseRef(GetChildPtr)

	GetChildCls := &Widget{}
	GetChildCls.Ptr = GetChildPtr
	return GetChildCls

}

var xListItemGetItem func(uintptr) uintptr

// Gets the model item that associated with @self.
//
// If @self is unbound, this function returns %NULL.
func (x *ListItem) GetItem() *gobject.Object {

	GetItemPtr := xListItemGetItem(x.GoPointer())
	if GetItemPtr == 0 {
		return nil
	}

	gobject.IncreaseRef(GetItemPtr)

	GetItemCls := &gobject.Object{}
	GetItemCls.Ptr = GetItemPtr
	return GetItemCls

}

var xListItemGetPosition func(uintptr) uint

// Gets the position in the model that @self currently displays.
//
// If @self is unbound, %GTK_INVALID_LIST_POSITION is returned.
func (x *ListItem) GetPosition() uint {

	return xListItemGetPosition(x.GoPointer())

}

var xListItemGetSelectable func(uintptr) bool

// Checks if a list item has been set to be selectable via
// gtk_list_item_set_selectable().
//
// Do not confuse this function with [method@Gtk.ListItem.get_selected].
func (x *ListItem) GetSelectable() bool {

	return xListItemGetSelectable(x.GoPointer())

}

var xListItemGetSelected func(uintptr) bool

// Checks if the item is displayed as selected.
//
// The selected state is maintained by the liste widget and its model
// and cannot be set otherwise.
func (x *ListItem) GetSelected() bool {

	return xListItemGetSelected(x.GoPointer())

}

var xListItemSetActivatable func(uintptr, bool)

// Sets @self to be activatable.
//
// If an item is activatable, double-clicking on the item, using
// the Return key or calling gtk_widget_activate() will activate
// the item. Activating instructs the containing view to handle
// activation. `GtkListView` for example will be emitting the
// [signal@Gtk.ListView::activate] signal.
//
// By default, list items are activatable.
func (x *ListItem) SetActivatable(ActivatableVar bool) {

	xListItemSetActivatable(x.GoPointer(), ActivatableVar)

}

var xListItemSetChild func(uintptr, uintptr)

// Sets the child to be used for this listitem.
//
// This function is typically called by applications when
// setting up a listitem so that the widget can be reused when
// binding it multiple times.
func (x *ListItem) SetChild(ChildVar *Widget) {

	xListItemSetChild(x.GoPointer(), ChildVar.GoPointer())

}

var xListItemSetSelectable func(uintptr, bool)

// Sets @self to be selectable.
//
// If an item is selectable, clicking on the item or using the keyboard
// will try to select or unselect the item. If this succeeds is up to
// the model to determine, as it is managing the selected state.
//
// Note that this means that making an item non-selectable has no
// influence on the selected state at all. A non-selectable item
// may still be selected.
//
// By default, list items are selectable. When rebinding them to
// a new item, they will also be reset to be selectable by GTK.
func (x *ListItem) SetSelectable(SelectableVar bool) {

	xListItemSetSelectable(x.GoPointer(), SelectableVar)

}

func (c *ListItem) GoPointer() uintptr {
	return c.Ptr
}

func (c *ListItem) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GTK"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xListItemGetActivatable, lib, "gtk_list_item_get_activatable")
	core.PuregoSafeRegister(&xListItemGetChild, lib, "gtk_list_item_get_child")
	core.PuregoSafeRegister(&xListItemGetItem, lib, "gtk_list_item_get_item")
	core.PuregoSafeRegister(&xListItemGetPosition, lib, "gtk_list_item_get_position")
	core.PuregoSafeRegister(&xListItemGetSelectable, lib, "gtk_list_item_get_selectable")
	core.PuregoSafeRegister(&xListItemGetSelected, lib, "gtk_list_item_get_selected")
	core.PuregoSafeRegister(&xListItemSetActivatable, lib, "gtk_list_item_set_activatable")
	core.PuregoSafeRegister(&xListItemSetChild, lib, "gtk_list_item_set_child")
	core.PuregoSafeRegister(&xListItemSetSelectable, lib, "gtk_list_item_set_selectable")

}
