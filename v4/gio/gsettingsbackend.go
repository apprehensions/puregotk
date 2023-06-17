// Package gio was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gio

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/glib"
	"github.com/jwijenbergh/puregotk/v4/gobject"
)

// Class structure for #GSettingsBackend.
type SettingsBackendClass struct {
	ParentClass uintptr

	Padding uintptr
}

type SettingsBackendPrivate struct {
}

var xKeyfileSettingsBackendNew func(string, string, string) uintptr

// Creates a keyfile-backed #GSettingsBackend.
//
// The filename of the keyfile to use is given by @filename.
//
// All settings read to or written from the backend must fall under the
// path given in @root_path (which must start and end with a slash and
// not contain two consecutive slashes).  @root_path may be "/".
//
// If @root_group is non-%NULL then it specifies the name of the keyfile
// group used for keys that are written directly below @root_path.  For
// example, if @root_path is "/apps/example/" and @root_group is
// "toplevel", then settings the key "/apps/example/enabled" to a value
// of %TRUE will cause the following to appear in the keyfile:
//
// |[
//
//	[toplevel]
//	enabled=true
//
// ]|
//
// If @root_group is %NULL then it is not permitted to store keys
// directly below the @root_path.
//
// For keys not stored directly below @root_path (ie: in a sub-path),
// the name of the subpath (with the final slash stripped) is used as
// the name of the keyfile group.  To continue the example, if
// "/apps/example/profiles/default/font-size" were set to
// 12 then the following would appear in the keyfile:
//
// |[
//
//	[profiles/default]
//	font-size=12
//
// ]|
//
// The backend will refuse writes (and return writability as being
// %FALSE) for keys outside of @root_path and, in the event that
// @root_group is %NULL, also for keys directly under @root_path.
// Writes will also be refused if the backend detects that it has the
// inability to rewrite the keyfile (ie: the containing directory is not
// writable).
//
// There is no checking done for your key namespace clashing with the
// syntax of the key file format.  For example, if you have '[' or ']'
// characters in your path names or '=' in your key names you may be in
// trouble.
//
// The backend reads default values from a keyfile called `defaults` in
// the directory specified by the #GKeyfileSettingsBackend:defaults-dir property,
// and a list of locked keys from a text file with the name `locks` in
// the same location.
func KeyfileSettingsBackendNew(FilenameVar string, RootPathVar string, RootGroupVar string) *SettingsBackend {

	KeyfileSettingsBackendNewPtr := xKeyfileSettingsBackendNew(FilenameVar, RootPathVar, RootGroupVar)
	if KeyfileSettingsBackendNewPtr == 0 {
		return nil
	}

	KeyfileSettingsBackendNewCls := &SettingsBackend{}
	KeyfileSettingsBackendNewCls.Ptr = KeyfileSettingsBackendNewPtr
	return KeyfileSettingsBackendNewCls

}

var xMemorySettingsBackendNew func() uintptr

// Creates a memory-backed #GSettingsBackend.
//
// This backend allows changes to settings, but does not write them
// to any backing storage, so the next time you run your application,
// the memory backend will start out with the default values again.
func MemorySettingsBackendNew() *SettingsBackend {

	MemorySettingsBackendNewPtr := xMemorySettingsBackendNew()
	if MemorySettingsBackendNewPtr == 0 {
		return nil
	}

	MemorySettingsBackendNewCls := &SettingsBackend{}
	MemorySettingsBackendNewCls.Ptr = MemorySettingsBackendNewPtr
	return MemorySettingsBackendNewCls

}

var xNullSettingsBackendNew func() uintptr

// Creates a readonly #GSettingsBackend.
//
// This backend does not allow changes to settings, so all settings
// will always have their default values.
func NullSettingsBackendNew() *SettingsBackend {

	NullSettingsBackendNewPtr := xNullSettingsBackendNew()
	if NullSettingsBackendNewPtr == 0 {
		return nil
	}

	NullSettingsBackendNewCls := &SettingsBackend{}
	NullSettingsBackendNewCls.Ptr = NullSettingsBackendNewPtr
	return NullSettingsBackendNewCls

}

// The #GSettingsBackend interface defines a generic interface for
// non-strictly-typed data that is stored in a hierarchy. To implement
// an alternative storage backend for #GSettings, you need to implement
// the #GSettingsBackend interface and then make it implement the
// extension point %G_SETTINGS_BACKEND_EXTENSION_POINT_NAME.
//
// The interface defines methods for reading and writing values, a
// method for determining if writing of certain values will fail
// (lockdown) and a change notification mechanism.
//
// The semantics of the interface are very precisely defined and
// implementations must carefully adhere to the expectations of
// callers that are documented on each of the interface methods.
//
// Some of the #GSettingsBackend functions accept or return a #GTree.
// These trees always have strings as keys and #GVariant as values.
// g_settings_backend_create_tree() is a convenience function to create
// suitable trees.
//
// The #GSettingsBackend API is exported to allow third-party
// implementations, but does not carry the same stability guarantees
// as the public GIO API. For this reason, you have to define the
// C preprocessor symbol %G_SETTINGS_ENABLE_BACKEND before including
// `gio/gsettingsbackend.h`.
type SettingsBackend struct {
	gobject.Object
}

func SettingsBackendNewFromInternalPtr(ptr uintptr) *SettingsBackend {
	cls := &SettingsBackend{}
	cls.Ptr = ptr
	return cls
}

var xSettingsBackendChanged func(uintptr, string, uintptr)

// Signals that a single key has possibly changed.  Backend
// implementations should call this if a key has possibly changed its
// value.
//
// @key must be a valid key (ie starting with a slash, not containing
// '//', and not ending with a slash).
//
// The implementation must call this function during any call to
// g_settings_backend_write(), before the call returns (except in the
// case that no keys are actually changed and it cares to detect this
// fact).  It may not rely on the existence of a mainloop for
// dispatching the signal later.
//
// The implementation may call this function at any other time it likes
// in response to other events (such as changes occurring outside of the
// program).  These calls may originate from a mainloop or may originate
// in response to any other action (including from calls to
// g_settings_backend_write()).
//
// In the case that this call is in response to a call to
// g_settings_backend_write() then @origin_tag must be set to the same
// value that was passed to that call.
func (x *SettingsBackend) Changed(KeyVar string, OriginTagVar uintptr) {

	xSettingsBackendChanged(x.GoPointer(), KeyVar, OriginTagVar)

}

var xSettingsBackendChangedTree func(uintptr, *glib.Tree, uintptr)

// This call is a convenience wrapper.  It gets the list of changes from
// @tree, computes the longest common prefix and calls
// g_settings_backend_changed().
func (x *SettingsBackend) ChangedTree(TreeVar *glib.Tree, OriginTagVar uintptr) {

	xSettingsBackendChangedTree(x.GoPointer(), TreeVar, OriginTagVar)

}

var xSettingsBackendKeysChanged func(uintptr, string, uintptr, uintptr)

// Signals that a list of keys have possibly changed.  Backend
// implementations should call this if keys have possibly changed their
// values.
//
// @path must be a valid path (ie starting and ending with a slash and
// not containing '//').  Each string in @items must form a valid key
// name when @path is prefixed to it (ie: each item must not start or
// end with '/' and must not contain '//').
//
// The meaning of this signal is that any of the key names resulting
// from the contatenation of @path with each item in @items may have
// changed.
//
// The same rules for when notifications must occur apply as per
// g_settings_backend_changed().  These two calls can be used
// interchangeably if exactly one item has changed (although in that
// case g_settings_backend_changed() is definitely preferred).
//
// For efficiency reasons, the implementation should strive for @path to
// be as long as possible (ie: the longest common prefix of all of the
// keys that were changed) but this is not strictly required.
func (x *SettingsBackend) KeysChanged(PathVar string, ItemsVar uintptr, OriginTagVar uintptr) {

	xSettingsBackendKeysChanged(x.GoPointer(), PathVar, ItemsVar, OriginTagVar)

}

var xSettingsBackendPathChanged func(uintptr, string, uintptr)

// Signals that all keys below a given path may have possibly changed.
// Backend implementations should call this if an entire path of keys
// have possibly changed their values.
//
// @path must be a valid path (ie starting and ending with a slash and
// not containing '//').
//
// The meaning of this signal is that any of the key which has a name
// starting with @path may have changed.
//
// The same rules for when notifications must occur apply as per
// g_settings_backend_changed().  This call might be an appropriate
// reasponse to a 'reset' call but implementations are also free to
// explicitly list the keys that were affected by that call if they can
// easily do so.
//
// For efficiency reasons, the implementation should strive for @path to
// be as long as possible (ie: the longest common prefix of all of the
// keys that were changed) but this is not strictly required.  As an
// example, if this function is called with the path of "/" then every
// single key in the application will be notified of a possible change.
func (x *SettingsBackend) PathChanged(PathVar string, OriginTagVar uintptr) {

	xSettingsBackendPathChanged(x.GoPointer(), PathVar, OriginTagVar)

}

var xSettingsBackendPathWritableChanged func(uintptr, string)

// Signals that the writability of all keys below a given path may have
// changed.
//
// Since GSettings performs no locking operations for itself, this call
// will always be made in response to external events.
func (x *SettingsBackend) PathWritableChanged(PathVar string) {

	xSettingsBackendPathWritableChanged(x.GoPointer(), PathVar)

}

var xSettingsBackendWritableChanged func(uintptr, string)

// Signals that the writability of a single key has possibly changed.
//
// Since GSettings performs no locking operations for itself, this call
// will always be made in response to external events.
func (x *SettingsBackend) WritableChanged(KeyVar string) {

	xSettingsBackendWritableChanged(x.GoPointer(), KeyVar)

}

func (c *SettingsBackend) GoPointer() uintptr {
	return c.Ptr
}

func (c *SettingsBackend) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GIO"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	core.PuregoSafeRegister(&xKeyfileSettingsBackendNew, lib, "g_keyfile_settings_backend_new")
	core.PuregoSafeRegister(&xMemorySettingsBackendNew, lib, "g_memory_settings_backend_new")
	core.PuregoSafeRegister(&xNullSettingsBackendNew, lib, "g_null_settings_backend_new")

	core.PuregoSafeRegister(&xSettingsBackendChanged, lib, "g_settings_backend_changed")
	core.PuregoSafeRegister(&xSettingsBackendChangedTree, lib, "g_settings_backend_changed_tree")
	core.PuregoSafeRegister(&xSettingsBackendKeysChanged, lib, "g_settings_backend_keys_changed")
	core.PuregoSafeRegister(&xSettingsBackendPathChanged, lib, "g_settings_backend_path_changed")
	core.PuregoSafeRegister(&xSettingsBackendPathWritableChanged, lib, "g_settings_backend_path_writable_changed")
	core.PuregoSafeRegister(&xSettingsBackendWritableChanged, lib, "g_settings_backend_writable_changed")

}
