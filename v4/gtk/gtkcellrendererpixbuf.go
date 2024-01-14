// Package gtk was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gtk

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/gobject"
)

// Renders a pixbuf in a cell
//
// A `GtkCellRendererPixbuf` can be used to render an image in a cell. It allows
// to render either a given `GdkPixbuf` (set via the
// `GtkCellRendererPixbuf:pixbuf` property) or a named icon (set via the
// `GtkCellRendererPixbuf:icon-name` property).
//
// To support the tree view, `GtkCellRendererPixbuf` also supports rendering two
// alternative pixbufs, when the `GtkCellRenderer:is-expander` property is %TRUE.
// If the `GtkCellRenderer:is-expanded property` is %TRUE and the
// `GtkCellRendererPixbuf:pixbuf-expander-open` property is set to a pixbuf, it
// renders that pixbuf, if the `GtkCellRenderer:is-expanded` property is %FALSE
// and the `GtkCellRendererPixbuf:pixbuf-expander-closed` property is set to a
// pixbuf, it renders that one.
type CellRendererPixbuf struct {
	CellRenderer
}

func CellRendererPixbufNewFromInternalPtr(ptr uintptr) *CellRendererPixbuf {
	cls := &CellRendererPixbuf{}
	cls.Ptr = ptr
	return cls
}

var xNewCellRendererPixbuf func() uintptr

// Creates a new `GtkCellRendererPixbuf`. Adjust rendering
// parameters using object properties. Object properties can be set
// globally (with g_object_set()). Also, with `GtkTreeViewColumn`, you
// can bind a property to a value in a `GtkTreeModel`. For example, you
// can bind the “pixbuf” property on the cell renderer to a pixbuf value
// in the model, thus rendering a different image in each row of the
// `GtkTreeView`.
func NewCellRendererPixbuf() *CellRendererPixbuf {
	var cls *CellRendererPixbuf

	cret := xNewCellRendererPixbuf()

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &CellRendererPixbuf{}
	cls.Ptr = cret
	return cls
}

func (c *CellRendererPixbuf) GoPointer() uintptr {
	return c.Ptr
}

func (c *CellRendererPixbuf) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GTK"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xNewCellRendererPixbuf, lib, "gtk_cell_renderer_pixbuf_new")

}
