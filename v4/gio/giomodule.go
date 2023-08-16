// Package gio was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gio

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/glib"
	"github.com/jwijenbergh/puregotk/v4/gobject"
)

type IOModuleClass struct {
}

// Represents a scope for loading IO modules. A scope can be used for blocking
// duplicate modules, or blocking a module you don't want to load.
//
// The scope can be used with g_io_modules_load_all_in_directory_with_scope()
// or g_io_modules_scan_all_in_directory_with_scope().
type IOModuleScope struct {
}

var xIoExtensionPointImplement func(string, []interface{}, string, int) *IOExtension

// Registers @type as extension for the extension point with name
// @extension_point_name.
//
// If @type has already been registered as an extension for this
// extension point, the existing #GIOExtension object is returned.
func IoExtensionPointImplement(ExtensionPointNameVar string, TypeVar []interface{}, ExtensionNameVar string, PriorityVar int) *IOExtension {

	cret := xIoExtensionPointImplement(ExtensionPointNameVar, TypeVar, ExtensionNameVar, PriorityVar)
	return cret
}

var xIoExtensionPointLookup func(string) *IOExtensionPoint

// Looks up an existing extension point.
func IoExtensionPointLookup(NameVar string) *IOExtensionPoint {

	cret := xIoExtensionPointLookup(NameVar)
	return cret
}

var xIoExtensionPointRegister func(string) *IOExtensionPoint

// Registers an extension point.
func IoExtensionPointRegister(NameVar string) *IOExtensionPoint {

	cret := xIoExtensionPointRegister(NameVar)
	return cret
}

var xIoModulesLoadAllInDirectory func(string) *glib.List

// Loads all the modules in the specified directory.
//
// If don't require all modules to be initialized (and thus registering
// all gtypes) then you can use g_io_modules_scan_all_in_directory()
// which allows delayed/lazy loading of modules.
func IoModulesLoadAllInDirectory(DirnameVar string) *glib.List {

	cret := xIoModulesLoadAllInDirectory(DirnameVar)
	return cret
}

var xIoModulesLoadAllInDirectoryWithScope func(string, *IOModuleScope) *glib.List

// Loads all the modules in the specified directory.
//
// If don't require all modules to be initialized (and thus registering
// all gtypes) then you can use g_io_modules_scan_all_in_directory()
// which allows delayed/lazy loading of modules.
func IoModulesLoadAllInDirectoryWithScope(DirnameVar string, ScopeVar *IOModuleScope) *glib.List {

	cret := xIoModulesLoadAllInDirectoryWithScope(DirnameVar, ScopeVar)
	return cret
}

var xIoModulesScanAllInDirectory func(string)

// Scans all the modules in the specified directory, ensuring that
// any extension point implemented by a module is registered.
//
// This may not actually load and initialize all the types in each
// module, some modules may be lazily loaded and initialized when
// an extension point it implements is used with e.g.
// g_io_extension_point_get_extensions() or
// g_io_extension_point_get_extension_by_name().
//
// If you need to guarantee that all types are loaded in all the modules,
// use g_io_modules_load_all_in_directory().
func IoModulesScanAllInDirectory(DirnameVar string) {

	xIoModulesScanAllInDirectory(DirnameVar)

}

var xIoModulesScanAllInDirectoryWithScope func(string, *IOModuleScope)

// Scans all the modules in the specified directory, ensuring that
// any extension point implemented by a module is registered.
//
// This may not actually load and initialize all the types in each
// module, some modules may be lazily loaded and initialized when
// an extension point it implements is used with e.g.
// g_io_extension_point_get_extensions() or
// g_io_extension_point_get_extension_by_name().
//
// If you need to guarantee that all types are loaded in all the modules,
// use g_io_modules_load_all_in_directory().
func IoModulesScanAllInDirectoryWithScope(DirnameVar string, ScopeVar *IOModuleScope) {

	xIoModulesScanAllInDirectoryWithScope(DirnameVar, ScopeVar)

}

// Provides an interface and default functions for loading and unloading
// modules. This is used internally to make GIO extensible, but can also
// be used by others to implement module loading.
type IOModule struct {
	gobject.TypeModule
}

func IOModuleNewFromInternalPtr(ptr uintptr) *IOModule {
	cls := &IOModule{}
	cls.Ptr = ptr
	return cls
}

var xNewIOModule func(string) uintptr

// Creates a new GIOModule that will load the specific
// shared library when in use.
func NewIOModule(FilenameVar string) *IOModule {
	var cls *IOModule

	cret := xNewIOModule(FilenameVar)

	if cret == 0 {
		return nil
	}
	cls = &IOModule{}
	cls.Ptr = cret
	return cls
}

var xIOModuleLoad func(uintptr)

// Required API for GIO modules to implement.
//
// This function is run after the module has been loaded into GIO,
// to initialize the module. Typically, this function will call
// g_io_extension_point_implement().
//
// Since 2.56, this function should be named `g_io_&lt;modulename&gt;_load`, where
// `modulename` is the plugin’s filename with the `lib` or `libgio` prefix and
// everything after the first dot removed, and with `-` replaced with `_`
// throughout. For example, `libgiognutls-helper.so` becomes `gnutls_helper`.
// Using the new symbol names avoids name clashes when building modules
// statically. The old symbol names continue to be supported, but cannot be used
// for static builds.
func (x *IOModule) Load() {

	xIOModuleLoad(x.GoPointer())

}

var xIOModuleUnload func(uintptr)

// Required API for GIO modules to implement.
//
// This function is run when the module is being unloaded from GIO,
// to finalize the module.
//
// Since 2.56, this function should be named `g_io_&lt;modulename&gt;_unload`, where
// `modulename` is the plugin’s filename with the `lib` or `libgio` prefix and
// everything after the first dot removed, and with `-` replaced with `_`
// throughout. For example, `libgiognutls-helper.so` becomes `gnutls_helper`.
// Using the new symbol names avoids name clashes when building modules
// statically. The old symbol names continue to be supported, but cannot be used
// for static builds.
func (x *IOModule) Unload() {

	xIOModuleUnload(x.GoPointer())

}

func (c *IOModule) GoPointer() uintptr {
	return c.Ptr
}

func (c *IOModule) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// Calls the @complete_interface_info function from the
// #GTypePluginClass of @plugin. There should be no need to use this
// function outside of the GObject type system itself.
func (x *IOModule) CompleteInterfaceInfo(InstanceTypeVar []interface{}, InterfaceTypeVar []interface{}, InfoVar *gobject.InterfaceInfo) {

	gobject.XGTypePluginCompleteInterfaceInfo(x.GoPointer(), InstanceTypeVar, InterfaceTypeVar, InfoVar)

}

// Calls the @complete_type_info function from the #GTypePluginClass of @plugin.
// There should be no need to use this function outside of the GObject
// type system itself.
func (x *IOModule) CompleteTypeInfo(GTypeVar []interface{}, InfoVar *gobject.TypeInfo, ValueTableVar *gobject.TypeValueTable) {

	gobject.XGTypePluginCompleteTypeInfo(x.GoPointer(), GTypeVar, InfoVar, ValueTableVar)

}

// Calls the @unuse_plugin function from the #GTypePluginClass of
// @plugin.  There should be no need to use this function outside of
// the GObject type system itself.
func (x *IOModule) Unuse() {

	gobject.XGTypePluginUnuse(x.GoPointer())

}

// Calls the @use_plugin function from the #GTypePluginClass of
// @plugin.  There should be no need to use this function outside of
// the GObject type system itself.
func (x *IOModule) Use() {

	gobject.XGTypePluginUse(x.GoPointer())

}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GIO"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	core.PuregoSafeRegister(&xIoExtensionPointImplement, lib, "g_io_extension_point_implement")
	core.PuregoSafeRegister(&xIoExtensionPointLookup, lib, "g_io_extension_point_lookup")
	core.PuregoSafeRegister(&xIoExtensionPointRegister, lib, "g_io_extension_point_register")
	core.PuregoSafeRegister(&xIoModulesLoadAllInDirectory, lib, "g_io_modules_load_all_in_directory")
	core.PuregoSafeRegister(&xIoModulesLoadAllInDirectoryWithScope, lib, "g_io_modules_load_all_in_directory_with_scope")
	core.PuregoSafeRegister(&xIoModulesScanAllInDirectory, lib, "g_io_modules_scan_all_in_directory")
	core.PuregoSafeRegister(&xIoModulesScanAllInDirectoryWithScope, lib, "g_io_modules_scan_all_in_directory_with_scope")

	core.PuregoSafeRegister(&xNewIOModule, lib, "g_io_module_new")

	core.PuregoSafeRegister(&xIOModuleLoad, lib, "g_io_module_load")
	core.PuregoSafeRegister(&xIOModuleUnload, lib, "g_io_module_unload")

}
