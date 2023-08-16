// Package gtk was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gtk

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/glib"
	"github.com/jwijenbergh/puregotk/v4/gobject"
)

type BuilderListItemFactoryClass struct {
}

// `GtkBuilderListItemFactory` is a `GtkListItemFactory` that creates
// widgets by instantiating `GtkBuilder` UI templates.
//
// The templates must be extending `GtkListItem`, and typically use
// `GtkExpression`s to obtain data from the items in the model.
//
// Example:
// ```xml
//
//	&lt;interface&gt;
//	  &lt;template class="GtkListItem"&gt;
//	    &lt;property name="child"&gt;
//	      &lt;object class="GtkLabel"&gt;
//	        &lt;property name="xalign"&gt;0&lt;/property&gt;
//	        &lt;binding name="label"&gt;
//	          &lt;lookup name="name" type="SettingsKey"&gt;
//	            &lt;lookup name="item"&gt;GtkListItem&lt;/lookup&gt;
//	          &lt;/lookup&gt;
//	        &lt;/binding&gt;
//	      &lt;/object&gt;
//	    &lt;/property&gt;
//	  &lt;/template&gt;
//	&lt;/interface&gt;
//
// ```
type BuilderListItemFactory struct {
	ListItemFactory
}

func BuilderListItemFactoryNewFromInternalPtr(ptr uintptr) *BuilderListItemFactory {
	cls := &BuilderListItemFactory{}
	cls.Ptr = ptr
	return cls
}

var xNewFromBytesBuilderListItemFactory func(uintptr, *glib.Bytes) uintptr

// Creates a new `GtkBuilderListItemFactory` that instantiates widgets
// using @bytes as the data to pass to `GtkBuilder`.
func NewFromBytesBuilderListItemFactory(ScopeVar BuilderScope, BytesVar *glib.Bytes) *ListItemFactory {
	var cls *ListItemFactory

	cret := xNewFromBytesBuilderListItemFactory(ScopeVar.GoPointer(), BytesVar)

	if cret == 0 {
		return nil
	}
	cls = &ListItemFactory{}
	cls.Ptr = cret
	return cls
}

var xNewFromResourceBuilderListItemFactory func(uintptr, string) uintptr

// Creates a new `GtkBuilderListItemFactory` that instantiates widgets
// using data read from the given @resource_path to pass to `GtkBuilder`.
func NewFromResourceBuilderListItemFactory(ScopeVar BuilderScope, ResourcePathVar string) *ListItemFactory {
	var cls *ListItemFactory

	cret := xNewFromResourceBuilderListItemFactory(ScopeVar.GoPointer(), ResourcePathVar)

	if cret == 0 {
		return nil
	}
	cls = &ListItemFactory{}
	cls.Ptr = cret
	return cls
}

var xBuilderListItemFactoryGetBytes func(uintptr) *glib.Bytes

// Gets the data used as the `GtkBuilder` UI template for constructing
// listitems.
func (x *BuilderListItemFactory) GetBytes() *glib.Bytes {

	cret := xBuilderListItemFactoryGetBytes(x.GoPointer())
	return cret
}

var xBuilderListItemFactoryGetResource func(uintptr) string

// If the data references a resource, gets the path of that resource.
func (x *BuilderListItemFactory) GetResource() string {

	cret := xBuilderListItemFactoryGetResource(x.GoPointer())
	return cret
}

var xBuilderListItemFactoryGetScope func(uintptr) uintptr

// Gets the scope used when constructing listitems.
func (x *BuilderListItemFactory) GetScope() *BuilderScopeBase {
	var cls *BuilderScopeBase

	cret := xBuilderListItemFactoryGetScope(x.GoPointer())

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &BuilderScopeBase{}
	cls.Ptr = cret
	return cls
}

func (c *BuilderListItemFactory) GoPointer() uintptr {
	return c.Ptr
}

func (c *BuilderListItemFactory) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GTK"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xNewFromBytesBuilderListItemFactory, lib, "gtk_builder_list_item_factory_new_from_bytes")
	core.PuregoSafeRegister(&xNewFromResourceBuilderListItemFactory, lib, "gtk_builder_list_item_factory_new_from_resource")

	core.PuregoSafeRegister(&xBuilderListItemFactoryGetBytes, lib, "gtk_builder_list_item_factory_get_bytes")
	core.PuregoSafeRegister(&xBuilderListItemFactoryGetResource, lib, "gtk_builder_list_item_factory_get_resource")
	core.PuregoSafeRegister(&xBuilderListItemFactoryGetScope, lib, "gtk_builder_list_item_factory_get_scope")

}
