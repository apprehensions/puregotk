// Package gtk was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gtk

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/gio"
	"github.com/jwijenbergh/puregotk/v4/glib"
	"github.com/jwijenbergh/puregotk/v4/gobject"
)

// User function that is called to map an @item of the original model to
// an item expected by the map model.
//
// The returned items must conform to the item type of the model they are
// used with.
type MapListModelMapFunc func(uintptr, uintptr) uintptr

type MapListModelClass struct {
	ParentClass uintptr
}

// A `GtkMapListModel` maps the items in a list model to different items.
//
// `GtkMapListModel` uses a [callback@Gtk.MapListModelMapFunc].
//
// Example: Create a list of `GtkEventControllers`
// ```c
// static gpointer
// map_to_controllers (gpointer widget,
//
//	gpointer data)
//
//	{
//	  gpointer result = gtk_widget_observe_controllers (widget);
//	  g_object_unref (widget);
//	  return result;
//	}
//
// widgets = gtk_widget_observe_children (widget);
//
// controllers = gtk_map_list_model_new (widgets,
//
//	map_to_controllers,
//	NULL, NULL);
//
// model = gtk_flatten_list_model_new (GTK_TYPE_EVENT_CONTROLLER,
//
//	controllers);
//
// ```
//
// `GtkMapListModel` will attempt to discard the mapped objects as soon as
// they are no longer needed and recreate them if necessary.
type MapListModel struct {
	gobject.Object
}

func MapListModelNewFromInternalPtr(ptr uintptr) *MapListModel {
	cls := &MapListModel{}
	cls.Ptr = ptr
	return cls
}

var xNewMapListModel func(uintptr, uintptr, uintptr, uintptr) uintptr

// Creates a new `GtkMapListModel` for the given arguments.
func NewMapListModel(ModelVar gio.ListModel, MapFuncVar MapListModelMapFunc, UserDataVar uintptr, UserDestroyVar glib.DestroyNotify) *MapListModel {
	NewMapListModelPtr := xNewMapListModel(ModelVar.GoPointer(), purego.NewCallback(MapFuncVar), UserDataVar, purego.NewCallback(UserDestroyVar))
	if NewMapListModelPtr == 0 {
		return nil
	}

	NewMapListModelCls := &MapListModel{}
	NewMapListModelCls.Ptr = NewMapListModelPtr
	return NewMapListModelCls
}

var xMapListModelGetModel func(uintptr) uintptr

// Gets the model that is currently being mapped or %NULL if none.
func (x *MapListModel) GetModel() *gio.ListModelBase {

	GetModelPtr := xMapListModelGetModel(x.GoPointer())
	if GetModelPtr == 0 {
		return nil
	}

	gobject.IncreaseRef(GetModelPtr)

	GetModelCls := &gio.ListModelBase{}
	GetModelCls.Ptr = GetModelPtr
	return GetModelCls

}

var xMapListModelHasMap func(uintptr) bool

// Checks if a map function is currently set on @self.
func (x *MapListModel) HasMap() bool {

	return xMapListModelHasMap(x.GoPointer())

}

var xMapListModelSetMapFunc func(uintptr, uintptr, uintptr, uintptr)

// Sets the function used to map items.
//
// The function will be called whenever an item needs to be mapped
// and must return the item to use for the given input item.
//
// Note that `GtkMapListModel` may call this function multiple times
// on the same item, because it may delete items it doesn't need anymore.
//
// GTK makes no effort to ensure that @map_func conforms to the item type
// of @self. It assumes that the caller knows what they are doing and the map
// function returns items of the appropriate type.
func (x *MapListModel) SetMapFunc(MapFuncVar MapListModelMapFunc, UserDataVar uintptr, UserDestroyVar glib.DestroyNotify) {

	xMapListModelSetMapFunc(x.GoPointer(), purego.NewCallback(MapFuncVar), UserDataVar, purego.NewCallback(UserDestroyVar))

}

var xMapListModelSetModel func(uintptr, uintptr)

// Sets the model to be mapped.
//
// GTK makes no effort to ensure that @model conforms to the item type
// expected by the map function. It assumes that the caller knows what
// they are doing and have set up an appropriate map function.
func (x *MapListModel) SetModel(ModelVar gio.ListModel) {

	xMapListModelSetModel(x.GoPointer(), ModelVar.GoPointer())

}

func (c *MapListModel) GoPointer() uintptr {
	return c.Ptr
}

func (c *MapListModel) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// Get the item at @position.
//
// If @position is greater than the number of items in @list, %NULL is
// returned.
//
// %NULL is never returned for an index that is smaller than the length
// of the list.
//
// See also: g_list_model_get_n_items()
func (x *MapListModel) GetItem(PositionVar uint) uintptr {

	return gio.XGListModelGetItem(x.GoPointer(), PositionVar)

}

// Gets the type of the items in @list.
//
// All items returned from g_list_model_get_item() are of the type
// returned by this function, or a subtype, or if the type is an
// interface, they are an implementation of that interface.
//
// The item type of a #GListModel can not change during the life of the
// model.
func (x *MapListModel) GetItemType() []interface{} {

	return gio.XGListModelGetItemType(x.GoPointer())

}

// Gets the number of items in @list.
//
// Depending on the model implementation, calling this function may be
// less efficient than iterating the list with increasing values for
// @position until g_list_model_get_item() returns %NULL.
func (x *MapListModel) GetNItems() uint {

	return gio.XGListModelGetNItems(x.GoPointer())

}

// Get the item at @position.
//
// If @position is greater than the number of items in @list, %NULL is
// returned.
//
// %NULL is never returned for an index that is smaller than the length
// of the list.
//
// This function is meant to be used by language bindings in place
// of g_list_model_get_item().
//
// See also: g_list_model_get_n_items()
func (x *MapListModel) GetObject(PositionVar uint) *gobject.Object {

	GetObjectPtr := gio.XGListModelGetObject(x.GoPointer(), PositionVar)
	if GetObjectPtr == 0 {
		return nil
	}

	GetObjectCls := &gobject.Object{}
	GetObjectCls.Ptr = GetObjectPtr
	return GetObjectCls

}

// Emits the #GListModel::items-changed signal on @list.
//
// This function should only be called by classes implementing
// #GListModel. It has to be called after the internal representation
// of @list has been updated, because handlers connected to this signal
// might query the new state of the list.
//
// Implementations must only make changes to the model (as visible to
// its consumer) in places that will not cause problems for that
// consumer.  For models that are driven directly by a write API (such
// as #GListStore), changes can be reported in response to uses of that
// API.  For models that represent remote data, changes should only be
// made from a fresh mainloop dispatch.  It is particularly not
// permitted to make changes in response to a call to the #GListModel
// consumer API.
//
// Stated another way: in general, it is assumed that code making a
// series of accesses to the model via the API, without returning to the
// mainloop, and without calling other code, will continue to view the
// same contents of the model.
func (x *MapListModel) ItemsChanged(PositionVar uint, RemovedVar uint, AddedVar uint) {

	gio.XGListModelItemsChanged(x.GoPointer(), PositionVar, RemovedVar, AddedVar)

}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GTK"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xNewMapListModel, lib, "gtk_map_list_model_new")

	core.PuregoSafeRegister(&xMapListModelGetModel, lib, "gtk_map_list_model_get_model")
	core.PuregoSafeRegister(&xMapListModelHasMap, lib, "gtk_map_list_model_has_map")
	core.PuregoSafeRegister(&xMapListModelSetMapFunc, lib, "gtk_map_list_model_set_map_func")
	core.PuregoSafeRegister(&xMapListModelSetModel, lib, "gtk_map_list_model_set_model")

}
