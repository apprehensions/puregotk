// Package gio was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gio

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/glib"
	"github.com/jwijenbergh/puregotk/v4/gobject"
)

type ListStoreClass struct {
	ParentClass uintptr
}

// #GListStore is a simple implementation of #GListModel that stores all
// items in memory.
//
// It provides insertions, deletions, and lookups in logarithmic time
// with a fast path for the common case of iterating the list linearly.
type ListStore struct {
	gobject.Object
}

func ListStoreNewFromInternalPtr(ptr uintptr) *ListStore {
	cls := &ListStore{}
	cls.Ptr = ptr
	return cls
}

var xNewListStore func([]interface{}) uintptr

// Creates a new #GListStore with items of type @item_type. @item_type
// must be a subclass of #GObject.
func NewListStore(ItemTypeVar []interface{}) *ListStore {
	NewListStorePtr := xNewListStore(ItemTypeVar)
	if NewListStorePtr == 0 {
		return nil
	}

	NewListStoreCls := &ListStore{}
	NewListStoreCls.Ptr = NewListStorePtr
	return NewListStoreCls
}

var xListStoreAppend func(uintptr, uintptr)

// Appends @item to @store. @item must be of type #GListStore:item-type.
//
// This function takes a ref on @item.
//
// Use g_list_store_splice() to append multiple items at the same time
// efficiently.
func (x *ListStore) Append(ItemVar *gobject.Object) {

	xListStoreAppend(x.GoPointer(), ItemVar.GoPointer())

}

var xListStoreFind func(uintptr, uintptr, uint) bool

// Looks up the given @item in the list store by looping over the items until
// the first occurrence of @item. If @item was not found, then @position will
// not be set, and this method will return %FALSE.
//
// If you need to compare the two items with a custom comparison function, use
// g_list_store_find_with_equal_func() with a custom #GEqualFunc instead.
func (x *ListStore) Find(ItemVar *gobject.Object, PositionVar uint) bool {

	return xListStoreFind(x.GoPointer(), ItemVar.GoPointer(), PositionVar)

}

var xListStoreFindWithEqualFunc func(uintptr, uintptr, uintptr, uint) bool

// Looks up the given @item in the list store by looping over the items and
// comparing them with @compare_func until the first occurrence of @item which
// matches. If @item was not found, then @position will not be set, and this
// method will return %FALSE.
func (x *ListStore) FindWithEqualFunc(ItemVar *gobject.Object, EqualFuncVar glib.EqualFunc, PositionVar uint) bool {

	return xListStoreFindWithEqualFunc(x.GoPointer(), ItemVar.GoPointer(), purego.NewCallback(EqualFuncVar), PositionVar)

}

var xListStoreInsert func(uintptr, uint, uintptr)

// Inserts @item into @store at @position. @item must be of type
// #GListStore:item-type or derived from it. @position must be smaller
// than the length of the list, or equal to it to append.
//
// This function takes a ref on @item.
//
// Use g_list_store_splice() to insert multiple items at the same time
// efficiently.
func (x *ListStore) Insert(PositionVar uint, ItemVar *gobject.Object) {

	xListStoreInsert(x.GoPointer(), PositionVar, ItemVar.GoPointer())

}

var xListStoreInsertSorted func(uintptr, uintptr, uintptr, uintptr) uint

// Inserts @item into @store at a position to be determined by the
// @compare_func.
//
// The list must already be sorted before calling this function or the
// result is undefined.  Usually you would approach this by only ever
// inserting items by way of this function.
//
// This function takes a ref on @item.
func (x *ListStore) InsertSorted(ItemVar *gobject.Object, CompareFuncVar glib.CompareDataFunc, UserDataVar uintptr) uint {

	return xListStoreInsertSorted(x.GoPointer(), ItemVar.GoPointer(), purego.NewCallback(CompareFuncVar), UserDataVar)

}

var xListStoreRemove func(uintptr, uint)

// Removes the item from @store that is at @position. @position must be
// smaller than the current length of the list.
//
// Use g_list_store_splice() to remove multiple items at the same time
// efficiently.
func (x *ListStore) Remove(PositionVar uint) {

	xListStoreRemove(x.GoPointer(), PositionVar)

}

var xListStoreRemoveAll func(uintptr)

// Removes all items from @store.
func (x *ListStore) RemoveAll() {

	xListStoreRemoveAll(x.GoPointer())

}

var xListStoreSort func(uintptr, uintptr, uintptr)

// Sort the items in @store according to @compare_func.
func (x *ListStore) Sort(CompareFuncVar glib.CompareDataFunc, UserDataVar uintptr) {

	xListStoreSort(x.GoPointer(), purego.NewCallback(CompareFuncVar), UserDataVar)

}

var xListStoreSplice func(uintptr, uint, uint, uintptr, uint)

// Changes @store by removing @n_removals items and adding @n_additions
// items to it. @additions must contain @n_additions items of type
// #GListStore:item-type.  %NULL is not permitted.
//
// This function is more efficient than g_list_store_insert() and
// g_list_store_remove(), because it only emits
// #GListModel::items-changed once for the change.
//
// This function takes a ref on each item in @additions.
//
// The parameters @position and @n_removals must be correct (ie:
// @position + @n_removals must be less than or equal to the length of
// the list at the time this function is called).
func (x *ListStore) Splice(PositionVar uint, NRemovalsVar uint, AdditionsVar uintptr, NAdditionsVar uint) {

	xListStoreSplice(x.GoPointer(), PositionVar, NRemovalsVar, AdditionsVar, NAdditionsVar)

}

func (c *ListStore) GoPointer() uintptr {
	return c.Ptr
}

func (c *ListStore) SetGoPointer(ptr uintptr) {
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
func (x *ListStore) GetItem(PositionVar uint) uintptr {

	return XGListModelGetItem(x.GoPointer(), PositionVar)

}

// Gets the type of the items in @list.
//
// All items returned from g_list_model_get_item() are of the type
// returned by this function, or a subtype, or if the type is an
// interface, they are an implementation of that interface.
//
// The item type of a #GListModel can not change during the life of the
// model.
func (x *ListStore) GetItemType() []interface{} {

	return XGListModelGetItemType(x.GoPointer())

}

// Gets the number of items in @list.
//
// Depending on the model implementation, calling this function may be
// less efficient than iterating the list with increasing values for
// @position until g_list_model_get_item() returns %NULL.
func (x *ListStore) GetNItems() uint {

	return XGListModelGetNItems(x.GoPointer())

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
func (x *ListStore) GetObject(PositionVar uint) *gobject.Object {

	GetObjectPtr := XGListModelGetObject(x.GoPointer(), PositionVar)
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
func (x *ListStore) ItemsChanged(PositionVar uint, RemovedVar uint, AddedVar uint) {

	XGListModelItemsChanged(x.GoPointer(), PositionVar, RemovedVar, AddedVar)

}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GIO"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xNewListStore, lib, "g_list_store_new")

	core.PuregoSafeRegister(&xListStoreAppend, lib, "g_list_store_append")
	core.PuregoSafeRegister(&xListStoreFind, lib, "g_list_store_find")
	core.PuregoSafeRegister(&xListStoreFindWithEqualFunc, lib, "g_list_store_find_with_equal_func")
	core.PuregoSafeRegister(&xListStoreInsert, lib, "g_list_store_insert")
	core.PuregoSafeRegister(&xListStoreInsertSorted, lib, "g_list_store_insert_sorted")
	core.PuregoSafeRegister(&xListStoreRemove, lib, "g_list_store_remove")
	core.PuregoSafeRegister(&xListStoreRemoveAll, lib, "g_list_store_remove_all")
	core.PuregoSafeRegister(&xListStoreSort, lib, "g_list_store_sort")
	core.PuregoSafeRegister(&xListStoreSplice, lib, "g_list_store_splice")

}
