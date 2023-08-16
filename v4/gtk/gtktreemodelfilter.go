// Package gtk was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gtk

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/gdk"
	"github.com/jwijenbergh/puregotk/v4/glib"
	"github.com/jwijenbergh/puregotk/v4/gobject"
)

// A function which calculates display values from raw values in the model.
// It must fill @value with the display value for the column @column in the
// row indicated by @iter.
//
// Since this function is called for each data access, it’s not a
// particularly efficient operation.
type TreeModelFilterModifyFunc func(uintptr, *TreeIter, *gobject.Value, int, uintptr)

// A function which decides whether the row indicated by @iter is visible.
type TreeModelFilterVisibleFunc func(uintptr, *TreeIter, uintptr) bool

type TreeModelFilterClass struct {
	ParentClass uintptr

	Padding uintptr
}

type TreeModelFilterPrivate struct {
}

// A `GtkTreeModel` which hides parts of an underlying tree model
//
// A `GtkTreeModelFilter` is a tree model which wraps another tree model,
// and can do the following things:
//
//   - Filter specific rows, based on data from a “visible column”, a column
//     storing booleans indicating whether the row should be filtered or not,
//     or based on the return value of a “visible function”, which gets a
//     model, iter and user_data and returns a boolean indicating whether the
//     row should be filtered or not.
//
//   - Modify the “appearance” of the model, using a modify function.
//     This is extremely powerful and allows for just changing some
//     values and also for creating a completely different model based
//     on the given child model.
//
//   - Set a different root node, also known as a “virtual root”. You can pass
//     in a `GtkTreePath` indicating the root node for the filter at construction
//     time.
//
// The basic API is similar to `GtkTreeModelSort`. For an example on its usage,
// see the section on `GtkTreeModelSort`.
//
// When using `GtkTreeModelFilter`, it is important to realize that
// `GtkTreeModelFilter` maintains an internal cache of all nodes which are
// visible in its clients. The cache is likely to be a subtree of the tree
// exposed by the child model. `GtkTreeModelFilter` will not cache the entire
// child model when unnecessary to not compromise the caching mechanism
// that is exposed by the reference counting scheme. If the child model
// implements reference counting, unnecessary signals may not be emitted
// because of reference counting rule 3, see the `GtkTreeModel`
// documentation. (Note that e.g. `GtkTreeStore` does not implement
// reference counting and will always emit all signals, even when
// the receiving node is not visible).
//
// Because of this, limitations for possible visible functions do apply.
// In general, visible functions should only use data or properties from
// the node for which the visibility state must be determined, its siblings
// or its parents. Usually, having a dependency on the state of any child
// node is not possible, unless references are taken on these explicitly.
// When no such reference exists, no signals may be received for these child
// nodes (see reference counting rule number 3 in the `GtkTreeModel` section).
//
// Determining the visibility state of a given node based on the state
// of its child nodes is a frequently occurring use case. Therefore,
// `GtkTreeModelFilter` explicitly supports this. For example, when a node
// does not have any children, you might not want the node to be visible.
// As soon as the first row is added to the node’s child level (or the
// last row removed), the node’s visibility should be updated.
//
// This introduces a dependency from the node on its child nodes. In order
// to accommodate this, `GtkTreeModelFilter` must make sure the necessary
// signals are received from the child model. This is achieved by building,
// for all nodes which are exposed as visible nodes to `GtkTreeModelFilter`'s
// clients, the child level (if any) and take a reference on the first node
// in this level. Furthermore, for every row-inserted, row-changed or
// row-deleted signal (also these which were not handled because the node
// was not cached), `GtkTreeModelFilter` will check if the visibility state
// of any parent node has changed.
//
// Beware, however, that this explicit support is limited to these two
// cases. For example, if you want a node to be visible only if two nodes
// in a child’s child level (2 levels deeper) are visible, you are on your
// own. In this case, either rely on `GtkTreeStore` to emit all signals
// because it does not implement reference counting, or for models that
// do implement reference counting, obtain references on these child levels
// yourself.
type TreeModelFilter struct {
	gobject.Object
}

func TreeModelFilterNewFromInternalPtr(ptr uintptr) *TreeModelFilter {
	cls := &TreeModelFilter{}
	cls.Ptr = ptr
	return cls
}

var xTreeModelFilterClearCache func(uintptr)

// This function should almost never be called. It clears the @filter
// of any cached iterators that haven’t been reffed with
// gtk_tree_model_ref_node(). This might be useful if the child model
// being filtered is static (and doesn’t change often) and there has been
// a lot of unreffed access to nodes. As a side effect of this function,
// all unreffed iters will be invalid.
func (x *TreeModelFilter) ClearCache() {

	xTreeModelFilterClearCache(x.GoPointer())

}

var xTreeModelFilterConvertChildIterToIter func(uintptr, *TreeIter, *TreeIter) bool

// Sets @filter_iter to point to the row in @filter that corresponds to the
// row pointed at by @child_iter.  If @filter_iter was not set, %FALSE is
// returned.
func (x *TreeModelFilter) ConvertChildIterToIter(FilterIterVar *TreeIter, ChildIterVar *TreeIter) bool {

	cret := xTreeModelFilterConvertChildIterToIter(x.GoPointer(), FilterIterVar, ChildIterVar)
	return cret
}

var xTreeModelFilterConvertChildPathToPath func(uintptr, *TreePath) *TreePath

// Converts @child_path to a path relative to @filter. That is, @child_path
// points to a path in the child model. The rerturned path will point to the
// same row in the filtered model. If @child_path isn’t a valid path on the
// child model or points to a row which is not visible in @filter, then %NULL
// is returned.
func (x *TreeModelFilter) ConvertChildPathToPath(ChildPathVar *TreePath) *TreePath {

	cret := xTreeModelFilterConvertChildPathToPath(x.GoPointer(), ChildPathVar)
	return cret
}

var xTreeModelFilterConvertIterToChildIter func(uintptr, *TreeIter, *TreeIter)

// Sets @child_iter to point to the row pointed to by @filter_iter.
func (x *TreeModelFilter) ConvertIterToChildIter(ChildIterVar *TreeIter, FilterIterVar *TreeIter) {

	xTreeModelFilterConvertIterToChildIter(x.GoPointer(), ChildIterVar, FilterIterVar)

}

var xTreeModelFilterConvertPathToChildPath func(uintptr, *TreePath) *TreePath

// Converts @filter_path to a path on the child model of @filter. That is,
// @filter_path points to a location in @filter. The returned path will
// point to the same location in the model not being filtered. If @filter_path
// does not point to a location in the child model, %NULL is returned.
func (x *TreeModelFilter) ConvertPathToChildPath(FilterPathVar *TreePath) *TreePath {

	cret := xTreeModelFilterConvertPathToChildPath(x.GoPointer(), FilterPathVar)
	return cret
}

var xTreeModelFilterGetModel func(uintptr) uintptr

// Returns a pointer to the child model of @filter.
func (x *TreeModelFilter) GetModel() *TreeModelBase {
	var cls *TreeModelBase

	cret := xTreeModelFilterGetModel(x.GoPointer())

	if cret == 0 {
		return cls
	}
	gobject.IncreaseRef(cret)
	cls = &TreeModelBase{}
	cls.Ptr = cret
	return cls
}

var xTreeModelFilterRefilter func(uintptr)

// Emits ::row_changed for each row in the child model, which causes
// the filter to re-evaluate whether a row is visible or not.
func (x *TreeModelFilter) Refilter() {

	xTreeModelFilterRefilter(x.GoPointer())

}

var xTreeModelFilterSetModifyFunc func(uintptr, int, uintptr, uintptr, uintptr, uintptr)

// With the @n_columns and @types parameters, you give an array of column
// types for this model (which will be exposed to the parent model/view).
// The @func, @data and @destroy parameters are for specifying the modify
// function. The modify function will get called for each
// data access, the goal of the modify function is to return the data which
// should be displayed at the location specified using the parameters of the
// modify function.
//
// Note that gtk_tree_model_filter_set_modify_func()
// can only be called once for a given filter model.
func (x *TreeModelFilter) SetModifyFunc(NColumnsVar int, TypesVar uintptr, FuncVar TreeModelFilterModifyFunc, DataVar uintptr, DestroyVar glib.DestroyNotify) {

	xTreeModelFilterSetModifyFunc(x.GoPointer(), NColumnsVar, TypesVar, purego.NewCallback(FuncVar), DataVar, purego.NewCallback(DestroyVar))

}

var xTreeModelFilterSetVisibleColumn func(uintptr, int)

// Sets @column of the child_model to be the column where @filter should
// look for visibility information. @columns should be a column of type
// %G_TYPE_BOOLEAN, where %TRUE means that a row is visible, and %FALSE
// if not.
//
// Note that gtk_tree_model_filter_set_visible_func() or
// gtk_tree_model_filter_set_visible_column() can only be called
// once for a given filter model.
func (x *TreeModelFilter) SetVisibleColumn(ColumnVar int) {

	xTreeModelFilterSetVisibleColumn(x.GoPointer(), ColumnVar)

}

var xTreeModelFilterSetVisibleFunc func(uintptr, uintptr, uintptr, uintptr)

// Sets the visible function used when filtering the @filter to be @func.
// The function should return %TRUE if the given row should be visible and
// %FALSE otherwise.
//
// If the condition calculated by the function changes over time (e.g.
// because it depends on some global parameters), you must call
// gtk_tree_model_filter_refilter() to keep the visibility information
// of the model up-to-date.
//
// Note that @func is called whenever a row is inserted, when it may still
// be empty. The visible function should therefore take special care of empty
// rows, like in the example below.
//
// |[&lt;!-- language="C" --&gt;
// static gboolean
// visible_func (GtkTreeModel *model,
//
//	GtkTreeIter  *iter,
//	gpointer      data)
//
//	{
//	  // Visible if row is non-empty and first column is “HI”
//	  char *str;
//	  gboolean visible = FALSE;
//
//	  gtk_tree_model_get (model, iter, 0, &amp;str, -1);
//	  if (str &amp;&amp; strcmp (str, "HI") == 0)
//	    visible = TRUE;
//	  g_free (str);
//
//	  return visible;
//	}
//
// ]|
//
// Note that gtk_tree_model_filter_set_visible_func() or
// gtk_tree_model_filter_set_visible_column() can only be called
// once for a given filter model.
func (x *TreeModelFilter) SetVisibleFunc(FuncVar TreeModelFilterVisibleFunc, DataVar uintptr, DestroyVar glib.DestroyNotify) {

	xTreeModelFilterSetVisibleFunc(x.GoPointer(), purego.NewCallback(FuncVar), DataVar, purego.NewCallback(DestroyVar))

}

func (c *TreeModelFilter) GoPointer() uintptr {
	return c.Ptr
}

func (c *TreeModelFilter) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// Asks the `GtkTreeDragSource` to delete the row at @path, because
// it was moved somewhere else via drag-and-drop. Returns %FALSE
// if the deletion fails because @path no longer exists, or for
// some model-specific reason. Should robustly handle a @path no
// longer found in the model!
func (x *TreeModelFilter) DragDataDelete(PathVar *TreePath) bool {

	cret := XGtkTreeDragSourceDragDataDelete(x.GoPointer(), PathVar)
	return cret
}

// Asks the `GtkTreeDragSource` to return a `GdkContentProvider` representing
// the row at @path. Should robustly handle a @path no
// longer found in the model!
func (x *TreeModelFilter) DragDataGet(PathVar *TreePath) *gdk.ContentProvider {
	var cls *gdk.ContentProvider

	cret := XGtkTreeDragSourceDragDataGet(x.GoPointer(), PathVar)

	if cret == 0 {
		return cls
	}
	cls = &gdk.ContentProvider{}
	cls.Ptr = cret
	return cls
}

// Asks the `GtkTreeDragSource` whether a particular row can be used as
// the source of a DND operation. If the source doesn’t implement
// this interface, the row is assumed draggable.
func (x *TreeModelFilter) RowDraggable(PathVar *TreePath) bool {

	cret := XGtkTreeDragSourceRowDraggable(x.GoPointer(), PathVar)
	return cret
}

// Creates a new `GtkTreeModel`, with @child_model as the child_model
// and @root as the virtual root.
func (x *TreeModelFilter) FilterNew(RootVar *TreePath) *TreeModelBase {
	var cls *TreeModelBase

	cret := XGtkTreeModelFilterNew(x.GoPointer(), RootVar)

	if cret == 0 {
		return cls
	}
	cls = &TreeModelBase{}
	cls.Ptr = cret
	return cls
}

// Calls @func on each node in model in a depth-first fashion.
//
// If @func returns %TRUE, then the tree ceases to be walked,
// and gtk_tree_model_foreach() returns.
func (x *TreeModelFilter) Foreach(FuncVar TreeModelForeachFunc, UserDataVar uintptr) {

	XGtkTreeModelForeach(x.GoPointer(), purego.NewCallback(FuncVar), UserDataVar)

}

// Gets the value of one or more cells in the row referenced by @iter.
//
// The variable argument list should contain integer column numbers,
// each column number followed by a place to store the value being
// retrieved.  The list is terminated by a -1. For example, to get a
// value from column 0 with type %G_TYPE_STRING, you would
// write: `gtk_tree_model_get (model, iter, 0, &amp;place_string_here, -1)`,
// where `place_string_here` is a #gchararray
// to be filled with the string.
//
// Returned values with type %G_TYPE_OBJECT have to be unreferenced,
// values with type %G_TYPE_STRING or %G_TYPE_BOXED have to be freed.
// Other values are passed by value.
func (x *TreeModelFilter) Get(IterVar *TreeIter, varArgs ...interface{}) {

	XGtkTreeModelGet(x.GoPointer(), IterVar, varArgs...)

}

// Returns the type of the column.
func (x *TreeModelFilter) GetColumnType(IndexVar int) []interface{} {

	cret := XGtkTreeModelGetColumnType(x.GoPointer(), IndexVar)
	return cret
}

// Returns a set of flags supported by this interface.
//
// The flags are a bitwise combination of `GtkTreeModel`Flags.
// The flags supported should not change during the lifetime
// of the @tree_model.
func (x *TreeModelFilter) GetFlags() TreeModelFlags {

	cret := XGtkTreeModelGetFlags(x.GoPointer())
	return cret
}

// Sets @iter to a valid iterator pointing to @path.
//
// If @path does not exist, @iter is set to an invalid
// iterator and %FALSE is returned.
func (x *TreeModelFilter) GetIter(IterVar *TreeIter, PathVar *TreePath) bool {

	cret := XGtkTreeModelGetIter(x.GoPointer(), IterVar, PathVar)
	return cret
}

// Initializes @iter with the first iterator in the tree
// (the one at the path "0").
//
// Returns %FALSE if the tree is empty, %TRUE otherwise.
func (x *TreeModelFilter) GetIterFirst(IterVar *TreeIter) bool {

	cret := XGtkTreeModelGetIterFirst(x.GoPointer(), IterVar)
	return cret
}

// Sets @iter to a valid iterator pointing to @path_string, if it
// exists.
//
// Otherwise, @iter is left invalid and %FALSE is returned.
func (x *TreeModelFilter) GetIterFromString(IterVar *TreeIter, PathStringVar string) bool {

	cret := XGtkTreeModelGetIterFromString(x.GoPointer(), IterVar, PathStringVar)
	return cret
}

// Returns the number of columns supported by @tree_model.
func (x *TreeModelFilter) GetNColumns() int {

	cret := XGtkTreeModelGetNColumns(x.GoPointer())
	return cret
}

// Returns a newly-created `GtkTreePath` referenced by @iter.
//
// This path should be freed with gtk_tree_path_free().
func (x *TreeModelFilter) GetPath(IterVar *TreeIter) *TreePath {

	cret := XGtkTreeModelGetPath(x.GoPointer(), IterVar)
	return cret
}

// Generates a string representation of the iter.
//
// This string is a “:” separated list of numbers.
// For example, “4:10:0:3” would be an acceptable
// return value for this string.
func (x *TreeModelFilter) GetStringFromIter(IterVar *TreeIter) string {

	cret := XGtkTreeModelGetStringFromIter(x.GoPointer(), IterVar)
	return cret
}

// Gets the value of one or more cells in the row referenced by @iter.
//
// See [method@Gtk.TreeModel.get], this version takes a va_list
// for language bindings to use.
func (x *TreeModelFilter) GetValist(IterVar *TreeIter, VarArgsVar []interface{}) {

	XGtkTreeModelGetValist(x.GoPointer(), IterVar, VarArgsVar)

}

// Initializes and sets @value to that at @column.
//
// When done with @value, g_value_unset() needs to be called
// to free any allocated memory.
func (x *TreeModelFilter) GetValue(IterVar *TreeIter, ColumnVar int, ValueVar *gobject.Value) {

	XGtkTreeModelGetValue(x.GoPointer(), IterVar, ColumnVar, ValueVar)

}

// Sets @iter to point to the first child of @parent.
//
// If @parent has no children, %FALSE is returned and @iter is
// set to be invalid. @parent will remain a valid node after this
// function has been called.
//
// If @parent is %NULL returns the first node, equivalent to
// `gtk_tree_model_get_iter_first (tree_model, iter);`
func (x *TreeModelFilter) IterChildren(IterVar *TreeIter, ParentVar *TreeIter) bool {

	cret := XGtkTreeModelIterChildren(x.GoPointer(), IterVar, ParentVar)
	return cret
}

// Returns %TRUE if @iter has children, %FALSE otherwise.
func (x *TreeModelFilter) IterHasChild(IterVar *TreeIter) bool {

	cret := XGtkTreeModelIterHasChild(x.GoPointer(), IterVar)
	return cret
}

// Returns the number of children that @iter has.
//
// As a special case, if @iter is %NULL, then the number
// of toplevel nodes is returned.
func (x *TreeModelFilter) IterNChildren(IterVar *TreeIter) int {

	cret := XGtkTreeModelIterNChildren(x.GoPointer(), IterVar)
	return cret
}

// Sets @iter to point to the node following it at the current level.
//
// If there is no next @iter, %FALSE is returned and @iter is set
// to be invalid.
func (x *TreeModelFilter) IterNext(IterVar *TreeIter) bool {

	cret := XGtkTreeModelIterNext(x.GoPointer(), IterVar)
	return cret
}

// Sets @iter to be the child of @parent, using the given index.
//
// The first index is 0. If @n is too big, or @parent has no children,
// @iter is set to an invalid iterator and %FALSE is returned. @parent
// will remain a valid node after this function has been called. As a
// special case, if @parent is %NULL, then the @n-th root node
// is set.
func (x *TreeModelFilter) IterNthChild(IterVar *TreeIter, ParentVar *TreeIter, NVar int) bool {

	cret := XGtkTreeModelIterNthChild(x.GoPointer(), IterVar, ParentVar, NVar)
	return cret
}

// Sets @iter to be the parent of @child.
//
// If @child is at the toplevel, and doesn’t have a parent, then
// @iter is set to an invalid iterator and %FALSE is returned.
// @child will remain a valid node after this function has been
// called.
//
// @iter will be initialized before the lookup is performed, so @child
// and @iter cannot point to the same memory location.
func (x *TreeModelFilter) IterParent(IterVar *TreeIter, ChildVar *TreeIter) bool {

	cret := XGtkTreeModelIterParent(x.GoPointer(), IterVar, ChildVar)
	return cret
}

// Sets @iter to point to the previous node at the current level.
//
// If there is no previous @iter, %FALSE is returned and @iter is
// set to be invalid.
func (x *TreeModelFilter) IterPrevious(IterVar *TreeIter) bool {

	cret := XGtkTreeModelIterPrevious(x.GoPointer(), IterVar)
	return cret
}

// Lets the tree ref the node.
//
// This is an optional method for models to implement.
// To be more specific, models may ignore this call as it exists
// primarily for performance reasons.
//
// This function is primarily meant as a way for views to let
// caching models know when nodes are being displayed (and hence,
// whether or not to cache that node). Being displayed means a node
// is in an expanded branch, regardless of whether the node is currently
// visible in the viewport. For example, a file-system based model
// would not want to keep the entire file-hierarchy in memory,
// just the sections that are currently being displayed by
// every current view.
//
// A model should be expected to be able to get an iter independent
// of its reffed state.
func (x *TreeModelFilter) RefNode(IterVar *TreeIter) {

	XGtkTreeModelRefNode(x.GoPointer(), IterVar)

}

// Emits the ::row-changed signal on @tree_model.
//
// See [signal@Gtk.TreeModel::row-changed].
func (x *TreeModelFilter) RowChanged(PathVar *TreePath, IterVar *TreeIter) {

	XGtkTreeModelRowChanged(x.GoPointer(), PathVar, IterVar)

}

// Emits the ::row-deleted signal on @tree_model.
//
// See [signal@Gtk.TreeModel::row-deleted].
//
// This should be called by models after a row has been removed.
// The location pointed to by @path should be the location that
// the row previously was at. It may not be a valid location anymore.
//
// Nodes that are deleted are not unreffed, this means that any
// outstanding references on the deleted node should not be released.
func (x *TreeModelFilter) RowDeleted(PathVar *TreePath) {

	XGtkTreeModelRowDeleted(x.GoPointer(), PathVar)

}

// Emits the ::row-has-child-toggled signal on @tree_model.
//
// See [signal@Gtk.TreeModel::row-has-child-toggled].
//
// This should be called by models after the child
// state of a node changes.
func (x *TreeModelFilter) RowHasChildToggled(PathVar *TreePath, IterVar *TreeIter) {

	XGtkTreeModelRowHasChildToggled(x.GoPointer(), PathVar, IterVar)

}

// Emits the ::row-inserted signal on @tree_model.
//
// See [signal@Gtk.TreeModel::row-inserted].
func (x *TreeModelFilter) RowInserted(PathVar *TreePath, IterVar *TreeIter) {

	XGtkTreeModelRowInserted(x.GoPointer(), PathVar, IterVar)

}

// Emits the ::rows-reordered signal on @tree_model.
//
// See [signal@Gtk.TreeModel::rows-reordered].
//
// This should be called by models when their rows have been
// reordered.
func (x *TreeModelFilter) RowsReordered(PathVar *TreePath, IterVar *TreeIter, NewOrderVar int) {

	XGtkTreeModelRowsReordered(x.GoPointer(), PathVar, IterVar, NewOrderVar)

}

// Emits the ::rows-reordered signal on @tree_model.
//
// See [signal@Gtk.TreeModel::rows-reordered].
//
// This should be called by models when their rows have been
// reordered.
func (x *TreeModelFilter) RowsReorderedWithLength(PathVar *TreePath, IterVar *TreeIter, NewOrderVar uintptr, LengthVar int) {

	XGtkTreeModelRowsReorderedWithLength(x.GoPointer(), PathVar, IterVar, NewOrderVar, LengthVar)

}

// Lets the tree unref the node.
//
// This is an optional method for models to implement.
// To be more specific, models may ignore this call as it exists
// primarily for performance reasons. For more information on what
// this means, see gtk_tree_model_ref_node().
//
// Please note that nodes that are deleted are not unreffed.
func (x *TreeModelFilter) UnrefNode(IterVar *TreeIter) {

	XGtkTreeModelUnrefNode(x.GoPointer(), IterVar)

}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GTK"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xTreeModelFilterClearCache, lib, "gtk_tree_model_filter_clear_cache")
	core.PuregoSafeRegister(&xTreeModelFilterConvertChildIterToIter, lib, "gtk_tree_model_filter_convert_child_iter_to_iter")
	core.PuregoSafeRegister(&xTreeModelFilterConvertChildPathToPath, lib, "gtk_tree_model_filter_convert_child_path_to_path")
	core.PuregoSafeRegister(&xTreeModelFilterConvertIterToChildIter, lib, "gtk_tree_model_filter_convert_iter_to_child_iter")
	core.PuregoSafeRegister(&xTreeModelFilterConvertPathToChildPath, lib, "gtk_tree_model_filter_convert_path_to_child_path")
	core.PuregoSafeRegister(&xTreeModelFilterGetModel, lib, "gtk_tree_model_filter_get_model")
	core.PuregoSafeRegister(&xTreeModelFilterRefilter, lib, "gtk_tree_model_filter_refilter")
	core.PuregoSafeRegister(&xTreeModelFilterSetModifyFunc, lib, "gtk_tree_model_filter_set_modify_func")
	core.PuregoSafeRegister(&xTreeModelFilterSetVisibleColumn, lib, "gtk_tree_model_filter_set_visible_column")
	core.PuregoSafeRegister(&xTreeModelFilterSetVisibleFunc, lib, "gtk_tree_model_filter_set_visible_func")

}
