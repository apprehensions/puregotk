// Package gtk was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gtk

import (
	"unsafe"

	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/gio"
	"github.com/jwijenbergh/puregotk/v4/glib"
	"github.com/jwijenbergh/puregotk/v4/gobject"
)

// Meta-data to be passed to gtk_recent_manager_add_full() when
// registering a recently used resource.
type RecentData struct {
	DisplayName uintptr

	Description uintptr

	MimeType uintptr

	AppName uintptr

	AppExec uintptr

	Groups []string

	IsPrivate bool
}

func (x *RecentData) GoPointer() uintptr {
	return uintptr(unsafe.Pointer(x))
}

// `GtkRecentInfo` contains the metadata associated with an item in the
// recently used files list.
type RecentInfo struct {
}

func (x *RecentInfo) GoPointer() uintptr {
	return uintptr(unsafe.Pointer(x))
}

var xRecentInfoCreateAppInfo func(uintptr, string, **glib.Error) uintptr

// Creates a `GAppInfo` for the specified `GtkRecentInfo`
//
// In case of error, @error will be set either with a
// %GTK_RECENT_MANAGER_ERROR or a %G_IO_ERROR
func (x *RecentInfo) CreateAppInfo(AppNameVar string) (*gio.AppInfoBase, error) {
	var cls *gio.AppInfoBase
	var cerr *glib.Error

	cret := xRecentInfoCreateAppInfo(x.GoPointer(), AppNameVar, &cerr)

	if cret == 0 {
		return nil, cerr
	}
	cls = &gio.AppInfoBase{}
	cls.Ptr = cret
	if cerr == nil {
		return cls, nil
	}
	return cls, cerr

}

var xRecentInfoExists func(uintptr) bool

// Checks whether the resource pointed by @info still exists.
// At the moment this check is done only on resources pointing
// to local files.
func (x *RecentInfo) Exists() bool {

	cret := xRecentInfoExists(x.GoPointer())
	return cret
}

var xRecentInfoGetAdded func(uintptr) *glib.DateTime

// Gets the time when the resource
// was added to the recently used resources list.
func (x *RecentInfo) GetAdded() *glib.DateTime {

	cret := xRecentInfoGetAdded(x.GoPointer())
	return cret
}

var xRecentInfoGetAge func(uintptr) int

// Gets the number of days elapsed since the last update
// of the resource pointed by @info.
func (x *RecentInfo) GetAge() int {

	cret := xRecentInfoGetAge(x.GoPointer())
	return cret
}

var xRecentInfoGetApplicationInfo func(uintptr, string, string, uint, **glib.DateTime) bool

// Gets the data regarding the application that has registered the resource
// pointed by @info.
//
// If the command line contains any escape characters defined inside the
// storage specification, they will be expanded.
func (x *RecentInfo) GetApplicationInfo(AppNameVar string, AppExecVar string, CountVar uint, StampVar **glib.DateTime) bool {

	cret := xRecentInfoGetApplicationInfo(x.GoPointer(), AppNameVar, AppExecVar, CountVar, StampVar)
	return cret
}

var xRecentInfoGetApplications func(uintptr, uint) []string

// Retrieves the list of applications that have registered this resource.
func (x *RecentInfo) GetApplications(LengthVar uint) []string {

	cret := xRecentInfoGetApplications(x.GoPointer(), LengthVar)
	return cret
}

var xRecentInfoGetDescription func(uintptr) string

// Gets the (short) description of the resource.
func (x *RecentInfo) GetDescription() string {

	cret := xRecentInfoGetDescription(x.GoPointer())
	return cret
}

var xRecentInfoGetDisplayName func(uintptr) string

// Gets the name of the resource.
//
// If none has been defined, the basename
// of the resource is obtained.
func (x *RecentInfo) GetDisplayName() string {

	cret := xRecentInfoGetDisplayName(x.GoPointer())
	return cret
}

var xRecentInfoGetGicon func(uintptr) uintptr

// Retrieves the icon associated to the resource MIME type.
func (x *RecentInfo) GetGicon() *gio.IconBase {
	var cls *gio.IconBase

	cret := xRecentInfoGetGicon(x.GoPointer())

	if cret == 0 {
		return nil
	}
	cls = &gio.IconBase{}
	cls.Ptr = cret
	return cls
}

var xRecentInfoGetGroups func(uintptr, uint) []string

// Returns all groups registered for the recently used item @info.
//
// The array of returned group names will be %NULL terminated, so
// length might optionally be %NULL.
func (x *RecentInfo) GetGroups(LengthVar uint) []string {

	cret := xRecentInfoGetGroups(x.GoPointer(), LengthVar)
	return cret
}

var xRecentInfoGetMimeType func(uintptr) string

// Gets the MIME type of the resource.
func (x *RecentInfo) GetMimeType() string {

	cret := xRecentInfoGetMimeType(x.GoPointer())
	return cret
}

var xRecentInfoGetModified func(uintptr) *glib.DateTime

// Gets the time when the meta-data
// for the resource was last modified.
func (x *RecentInfo) GetModified() *glib.DateTime {

	cret := xRecentInfoGetModified(x.GoPointer())
	return cret
}

var xRecentInfoGetPrivateHint func(uintptr) bool

// Gets the value of the “private” flag.
//
// Resources in the recently used list that have this flag
// set to %TRUE should only be displayed by the applications
// that have registered them.
func (x *RecentInfo) GetPrivateHint() bool {

	cret := xRecentInfoGetPrivateHint(x.GoPointer())
	return cret
}

var xRecentInfoGetShortName func(uintptr) string

// Computes a valid UTF-8 string that can be used as the
// name of the item in a menu or list.
//
// For example, calling this function on an item that refers
// to “file:///foo/bar.txt” will yield “bar.txt”.
func (x *RecentInfo) GetShortName() string {

	cret := xRecentInfoGetShortName(x.GoPointer())
	return cret
}

var xRecentInfoGetUri func(uintptr) string

// Gets the URI of the resource.
func (x *RecentInfo) GetUri() string {

	cret := xRecentInfoGetUri(x.GoPointer())
	return cret
}

var xRecentInfoGetUriDisplay func(uintptr) string

// Gets a displayable version of the resource’s URI.
//
// If the resource is local, it returns a local path; if the
// resource is not local, it returns the UTF-8 encoded content
// of [method@Gtk.RecentInfo.get_uri].
func (x *RecentInfo) GetUriDisplay() string {

	cret := xRecentInfoGetUriDisplay(x.GoPointer())
	return cret
}

var xRecentInfoGetVisited func(uintptr) *glib.DateTime

// Gets the time when the meta-data
// for the resource was last visited.
func (x *RecentInfo) GetVisited() *glib.DateTime {

	cret := xRecentInfoGetVisited(x.GoPointer())
	return cret
}

var xRecentInfoHasApplication func(uintptr, string) bool

// Checks whether an application registered this resource using @app_name.
func (x *RecentInfo) HasApplication(AppNameVar string) bool {

	cret := xRecentInfoHasApplication(x.GoPointer(), AppNameVar)
	return cret
}

var xRecentInfoHasGroup func(uintptr, string) bool

// Checks whether @group_name appears inside the groups
// registered for the recently used item @info.
func (x *RecentInfo) HasGroup(GroupNameVar string) bool {

	cret := xRecentInfoHasGroup(x.GoPointer(), GroupNameVar)
	return cret
}

var xRecentInfoIsLocal func(uintptr) bool

// Checks whether the resource is local or not by looking at the
// scheme of its URI.
func (x *RecentInfo) IsLocal() bool {

	cret := xRecentInfoIsLocal(x.GoPointer())
	return cret
}

var xRecentInfoLastApplication func(uintptr) string

// Gets the name of the last application that have registered the
// recently used resource represented by @info.
func (x *RecentInfo) LastApplication() string {

	cret := xRecentInfoLastApplication(x.GoPointer())
	return cret
}

var xRecentInfoMatch func(uintptr, *RecentInfo) bool

// Checks whether two `GtkRecentInfo` point to the same resource.
func (x *RecentInfo) Match(InfoBVar *RecentInfo) bool {

	cret := xRecentInfoMatch(x.GoPointer(), InfoBVar)
	return cret
}

var xRecentInfoRef func(uintptr) *RecentInfo

// Increases the reference count of @recent_info by one.
func (x *RecentInfo) Ref() *RecentInfo {

	cret := xRecentInfoRef(x.GoPointer())
	return cret
}

var xRecentInfoUnref func(uintptr)

// Decreases the reference count of @info by one.
//
// If the reference count reaches zero, @info is
// deallocated, and the memory freed.
func (x *RecentInfo) Unref() {

	xRecentInfoUnref(x.GoPointer())

}

// `GtkRecentManagerClass` contains only private data.
type RecentManagerClass struct {
	ParentClass uintptr
}

func (x *RecentManagerClass) GoPointer() uintptr {
	return uintptr(unsafe.Pointer(x))
}

type RecentManagerPrivate struct {
}

func (x *RecentManagerPrivate) GoPointer() uintptr {
	return uintptr(unsafe.Pointer(x))
}

// Error codes for `GtkRecentManager` operations
type RecentManagerError int

const (

	// the URI specified does not exists in
	//   the recently used resources list.
	RecentManagerErrorNotFoundValue RecentManagerError = 0
	// the URI specified is not valid.
	RecentManagerErrorInvalidUriValue RecentManagerError = 1
	// the supplied string is not
	//   UTF-8 encoded.
	RecentManagerErrorInvalidEncodingValue RecentManagerError = 2
	// no application has registered
	//   the specified item.
	RecentManagerErrorNotRegisteredValue RecentManagerError = 3
	// failure while reading the recently used
	//   resources file.
	RecentManagerErrorReadValue RecentManagerError = 4
	// failure while writing the recently used
	//   resources file.
	RecentManagerErrorWriteValue RecentManagerError = 5
	// unspecified error.
	RecentManagerErrorUnknownValue RecentManagerError = 6
)

// `GtkRecentManager` manages and looks up recently used files.
//
// Each recently used file is identified by its URI, and has meta-data
// associated to it, like the names and command lines of the applications
// that have registered it, the number of time each application has
// registered the same file, the mime type of the file and whether
// the file should be displayed only by the applications that have
// registered it.
//
// The recently used files list is per user.
//
// `GtkRecentManager` acts like a database of all the recently
// used files. You can create new `GtkRecentManager` objects, but
// it is more efficient to use the default manager created by GTK.
//
// Adding a new recently used file is as simple as:
//
// ```c
// GtkRecentManager *manager;
//
// manager = gtk_recent_manager_get_default ();
// gtk_recent_manager_add_item (manager, file_uri);
// ```
//
// The `GtkRecentManager` will try to gather all the needed information
// from the file itself through GIO.
//
// Looking up the meta-data associated with a recently used file
// given its URI requires calling [method@Gtk.RecentManager.lookup_item]:
//
// ```c
// GtkRecentManager *manager;
// GtkRecentInfo *info;
// GError *error = NULL;
//
// manager = gtk_recent_manager_get_default ();
// info = gtk_recent_manager_lookup_item (manager, file_uri, &amp;error);
// if (error)
//
//	{
//	  g_warning ("Could not find the file: %s", error-&gt;message);
//	  g_error_free (error);
//	}
//
// else
//
//	{
//	  // Use the info object
//	  gtk_recent_info_unref (info);
//	}
//
// ```
//
// In order to retrieve the list of recently used files, you can use
// [method@Gtk.RecentManager.get_items], which returns a list of
// [struct@Gtk.RecentInfo].
//
// Note that the maximum age of the recently used files list is
// controllable through the [property@Gtk.Settings:gtk-recent-files-max-age]
// property.
type RecentManager struct {
	gobject.Object
}

func RecentManagerNewFromInternalPtr(ptr uintptr) *RecentManager {
	cls := &RecentManager{}
	cls.Ptr = ptr
	return cls
}

var xNewRecentManager func() uintptr

// Creates a new recent manager object.
//
// Recent manager objects are used to handle the list of recently used
// resources. A `GtkRecentManager` object monitors the recently used
// resources list, and emits the [signal@Gtk.RecentManager::changed]
// signal each time something inside the list changes.
//
// `GtkRecentManager` objects are expensive: be sure to create them
// only when needed. You should use [func@Gtk.RecentManager.get_default]
// instead.
func NewRecentManager() *RecentManager {
	var cls *RecentManager

	cret := xNewRecentManager()

	if cret == 0 {
		return nil
	}
	cls = &RecentManager{}
	cls.Ptr = cret
	return cls
}

var xRecentManagerAddFull func(uintptr, string, *RecentData) bool

// Adds a new resource, pointed by @uri, into the recently used
// resources list, using the metadata specified inside the
// `GtkRecentData` passed in @recent_data.
//
// The passed URI will be used to identify this resource inside the
// list.
//
// In order to register the new recently used resource, metadata about
// the resource must be passed as well as the URI; the metadata is
// stored in a `GtkRecentData`, which must contain the MIME
// type of the resource pointed by the URI; the name of the application
// that is registering the item, and a command line to be used when
// launching the item.
//
// Optionally, a `GtkRecentData` might contain a UTF-8 string
// to be used when viewing the item instead of the last component of
// the URI; a short description of the item; whether the item should
// be considered private - that is, should be displayed only by the
// applications that have registered it.
func (x *RecentManager) AddFull(UriVar string, RecentDataVar *RecentData) bool {

	cret := xRecentManagerAddFull(x.GoPointer(), UriVar, RecentDataVar)
	return cret
}

var xRecentManagerAddItem func(uintptr, string) bool

// Adds a new resource, pointed by @uri, into the recently used
// resources list.
//
// This function automatically retrieves some of the needed
// metadata and setting other metadata to common default values;
// it then feeds the data to [method@Gtk.RecentManager.add_full].
//
// See [method@Gtk.RecentManager.add_full] if you want to explicitly
// define the metadata for the resource pointed by @uri.
func (x *RecentManager) AddItem(UriVar string) bool {

	cret := xRecentManagerAddItem(x.GoPointer(), UriVar)
	return cret
}

var xRecentManagerGetItems func(uintptr) *glib.List

// Gets the list of recently used resources.
func (x *RecentManager) GetItems() *glib.List {

	cret := xRecentManagerGetItems(x.GoPointer())
	return cret
}

var xRecentManagerHasItem func(uintptr, string) bool

// Checks whether there is a recently used resource registered
// with @uri inside the recent manager.
func (x *RecentManager) HasItem(UriVar string) bool {

	cret := xRecentManagerHasItem(x.GoPointer(), UriVar)
	return cret
}

var xRecentManagerLookupItem func(uintptr, string, **glib.Error) *RecentInfo

// Searches for a URI inside the recently used resources list, and
// returns a `GtkRecentInfo` containing information about the resource
// like its MIME type, or its display name.
func (x *RecentManager) LookupItem(UriVar string) (*RecentInfo, error) {
	var cerr *glib.Error

	cret := xRecentManagerLookupItem(x.GoPointer(), UriVar, &cerr)
	if cerr == nil {
		return cret, nil
	}
	return cret, cerr

}

var xRecentManagerMoveItem func(uintptr, string, string, **glib.Error) bool

// Changes the location of a recently used resource from @uri to @new_uri.
//
// Please note that this function will not affect the resource pointed
// by the URIs, but only the URI used in the recently used resources list.
func (x *RecentManager) MoveItem(UriVar string, NewUriVar string) (bool, error) {
	var cerr *glib.Error

	cret := xRecentManagerMoveItem(x.GoPointer(), UriVar, NewUriVar, &cerr)
	if cerr == nil {
		return cret, nil
	}
	return cret, cerr

}

var xRecentManagerPurgeItems func(uintptr) int

// Purges every item from the recently used resources list.
func (x *RecentManager) PurgeItems() (int, error) {
	var cerr *glib.Error

	cret := xRecentManagerPurgeItems(x.GoPointer())
	if cerr == nil {
		return cret, nil
	}
	return cret, cerr

}

var xRecentManagerRemoveItem func(uintptr, string, **glib.Error) bool

// Removes a resource pointed by @uri from the recently used resources
// list handled by a recent manager.
func (x *RecentManager) RemoveItem(UriVar string) (bool, error) {
	var cerr *glib.Error

	cret := xRecentManagerRemoveItem(x.GoPointer(), UriVar, &cerr)
	if cerr == nil {
		return cret, nil
	}
	return cret, cerr

}

func (c *RecentManager) GoPointer() uintptr {
	return c.Ptr
}

func (c *RecentManager) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// Emitted when the current recently used resources manager changes
// its contents.
//
// This can happen either by calling [method@Gtk.RecentManager.add_item]
// or by another application.
func (x *RecentManager) ConnectChanged(cb func(RecentManager)) uint32 {
	fcb := func(clsPtr uintptr) {
		fa := RecentManager{}
		fa.Ptr = clsPtr

		cb(fa)

	}
	return gobject.SignalConnect(x.GoPointer(), "changed", purego.NewCallback(fcb))
}

var xRecentManagerGetDefault func() uintptr

// Gets a unique instance of `GtkRecentManager` that you can share
// in your application without caring about memory management.
func RecentManagerGetDefault() *RecentManager {
	var cls *RecentManager

	cret := xRecentManagerGetDefault()

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &RecentManager{}
	cls.Ptr = cret
	return cls
}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GTK"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xRecentInfoCreateAppInfo, lib, "gtk_recent_info_create_app_info")
	core.PuregoSafeRegister(&xRecentInfoExists, lib, "gtk_recent_info_exists")
	core.PuregoSafeRegister(&xRecentInfoGetAdded, lib, "gtk_recent_info_get_added")
	core.PuregoSafeRegister(&xRecentInfoGetAge, lib, "gtk_recent_info_get_age")
	core.PuregoSafeRegister(&xRecentInfoGetApplicationInfo, lib, "gtk_recent_info_get_application_info")
	core.PuregoSafeRegister(&xRecentInfoGetApplications, lib, "gtk_recent_info_get_applications")
	core.PuregoSafeRegister(&xRecentInfoGetDescription, lib, "gtk_recent_info_get_description")
	core.PuregoSafeRegister(&xRecentInfoGetDisplayName, lib, "gtk_recent_info_get_display_name")
	core.PuregoSafeRegister(&xRecentInfoGetGicon, lib, "gtk_recent_info_get_gicon")
	core.PuregoSafeRegister(&xRecentInfoGetGroups, lib, "gtk_recent_info_get_groups")
	core.PuregoSafeRegister(&xRecentInfoGetMimeType, lib, "gtk_recent_info_get_mime_type")
	core.PuregoSafeRegister(&xRecentInfoGetModified, lib, "gtk_recent_info_get_modified")
	core.PuregoSafeRegister(&xRecentInfoGetPrivateHint, lib, "gtk_recent_info_get_private_hint")
	core.PuregoSafeRegister(&xRecentInfoGetShortName, lib, "gtk_recent_info_get_short_name")
	core.PuregoSafeRegister(&xRecentInfoGetUri, lib, "gtk_recent_info_get_uri")
	core.PuregoSafeRegister(&xRecentInfoGetUriDisplay, lib, "gtk_recent_info_get_uri_display")
	core.PuregoSafeRegister(&xRecentInfoGetVisited, lib, "gtk_recent_info_get_visited")
	core.PuregoSafeRegister(&xRecentInfoHasApplication, lib, "gtk_recent_info_has_application")
	core.PuregoSafeRegister(&xRecentInfoHasGroup, lib, "gtk_recent_info_has_group")
	core.PuregoSafeRegister(&xRecentInfoIsLocal, lib, "gtk_recent_info_is_local")
	core.PuregoSafeRegister(&xRecentInfoLastApplication, lib, "gtk_recent_info_last_application")
	core.PuregoSafeRegister(&xRecentInfoMatch, lib, "gtk_recent_info_match")
	core.PuregoSafeRegister(&xRecentInfoRef, lib, "gtk_recent_info_ref")
	core.PuregoSafeRegister(&xRecentInfoUnref, lib, "gtk_recent_info_unref")

	core.PuregoSafeRegister(&xNewRecentManager, lib, "gtk_recent_manager_new")

	core.PuregoSafeRegister(&xRecentManagerAddFull, lib, "gtk_recent_manager_add_full")
	core.PuregoSafeRegister(&xRecentManagerAddItem, lib, "gtk_recent_manager_add_item")
	core.PuregoSafeRegister(&xRecentManagerGetItems, lib, "gtk_recent_manager_get_items")
	core.PuregoSafeRegister(&xRecentManagerHasItem, lib, "gtk_recent_manager_has_item")
	core.PuregoSafeRegister(&xRecentManagerLookupItem, lib, "gtk_recent_manager_lookup_item")
	core.PuregoSafeRegister(&xRecentManagerMoveItem, lib, "gtk_recent_manager_move_item")
	core.PuregoSafeRegister(&xRecentManagerPurgeItems, lib, "gtk_recent_manager_purge_items")
	core.PuregoSafeRegister(&xRecentManagerRemoveItem, lib, "gtk_recent_manager_remove_item")

	core.PuregoSafeRegister(&xRecentManagerGetDefault, lib, "gtk_recent_manager_get_default")

}
