// Package gio was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gio

import (
	"unsafe"

	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/glib"
	"github.com/jwijenbergh/puregotk/v4/gobject"
)

// Defines a Unix mount entry (e.g. &lt;filename&gt;/media/cdrom&lt;/filename&gt;).
// This corresponds roughly to a mtab entry.
type UnixMountEntry struct {
}

func (x *UnixMountEntry) GoPointer() uintptr {
	return uintptr(unsafe.Pointer(x))
}

type UnixMountMonitorClass struct {
}

func (x *UnixMountMonitorClass) GoPointer() uintptr {
	return uintptr(unsafe.Pointer(x))
}

// Defines a Unix mount point (e.g. &lt;filename&gt;/dev&lt;/filename&gt;).
// This corresponds roughly to a fstab entry.
type UnixMountPoint struct {
}

func (x *UnixMountPoint) GoPointer() uintptr {
	return uintptr(unsafe.Pointer(x))
}

var xUnixMountPointCompare func(uintptr, *UnixMountPoint) int

// Compares two unix mount points.
func (x *UnixMountPoint) Compare(Mount2Var *UnixMountPoint) int {

	cret := xUnixMountPointCompare(x.GoPointer(), Mount2Var)
	return cret
}

var xUnixMountPointCopy func(uintptr) *UnixMountPoint

// Makes a copy of @mount_point.
func (x *UnixMountPoint) Copy() *UnixMountPoint {

	cret := xUnixMountPointCopy(x.GoPointer())
	return cret
}

var xUnixMountPointFree func(uintptr)

// Frees a unix mount point.
func (x *UnixMountPoint) Free() {

	xUnixMountPointFree(x.GoPointer())

}

var xUnixMountPointGetDevicePath func(uintptr) string

// Gets the device path for a unix mount point.
func (x *UnixMountPoint) GetDevicePath() string {

	cret := xUnixMountPointGetDevicePath(x.GoPointer())
	return cret
}

var xUnixMountPointGetFsType func(uintptr) string

// Gets the file system type for the mount point.
func (x *UnixMountPoint) GetFsType() string {

	cret := xUnixMountPointGetFsType(x.GoPointer())
	return cret
}

var xUnixMountPointGetMountPath func(uintptr) string

// Gets the mount path for a unix mount point.
func (x *UnixMountPoint) GetMountPath() string {

	cret := xUnixMountPointGetMountPath(x.GoPointer())
	return cret
}

var xUnixMountPointGetOptions func(uintptr) string

// Gets the options for the mount point.
func (x *UnixMountPoint) GetOptions() string {

	cret := xUnixMountPointGetOptions(x.GoPointer())
	return cret
}

var xUnixMountPointGuessCanEject func(uintptr) bool

// Guesses whether a Unix mount point can be ejected.
func (x *UnixMountPoint) GuessCanEject() bool {

	cret := xUnixMountPointGuessCanEject(x.GoPointer())
	return cret
}

var xUnixMountPointGuessIcon func(uintptr) uintptr

// Guesses the icon of a Unix mount point.
func (x *UnixMountPoint) GuessIcon() *IconBase {
	var cls *IconBase

	cret := xUnixMountPointGuessIcon(x.GoPointer())

	if cret == 0 {
		return nil
	}
	cls = &IconBase{}
	cls.Ptr = cret
	return cls
}

var xUnixMountPointGuessName func(uintptr) string

// Guesses the name of a Unix mount point.
// The result is a translated string.
func (x *UnixMountPoint) GuessName() string {

	cret := xUnixMountPointGuessName(x.GoPointer())
	return cret
}

var xUnixMountPointGuessSymbolicIcon func(uintptr) uintptr

// Guesses the symbolic icon of a Unix mount point.
func (x *UnixMountPoint) GuessSymbolicIcon() *IconBase {
	var cls *IconBase

	cret := xUnixMountPointGuessSymbolicIcon(x.GoPointer())

	if cret == 0 {
		return nil
	}
	cls = &IconBase{}
	cls.Ptr = cret
	return cls
}

var xUnixMountPointIsLoopback func(uintptr) bool

// Checks if a unix mount point is a loopback device.
func (x *UnixMountPoint) IsLoopback() bool {

	cret := xUnixMountPointIsLoopback(x.GoPointer())
	return cret
}

var xUnixMountPointIsReadonly func(uintptr) bool

// Checks if a unix mount point is read only.
func (x *UnixMountPoint) IsReadonly() bool {

	cret := xUnixMountPointIsReadonly(x.GoPointer())
	return cret
}

var xUnixMountPointIsUserMountable func(uintptr) bool

// Checks if a unix mount point is mountable by the user.
func (x *UnixMountPoint) IsUserMountable() bool {

	cret := xUnixMountPointIsUserMountable(x.GoPointer())
	return cret
}

var xUnixIsMountPathSystemInternal func(string) bool

// Determines if @mount_path is considered an implementation of the
// OS. This is primarily used for hiding mountable and mounted volumes
// that only are used in the OS and has little to no relevance to the
// casual user.
func UnixIsMountPathSystemInternal(MountPathVar string) bool {

	cret := xUnixIsMountPathSystemInternal(MountPathVar)
	return cret
}

var xUnixIsSystemDevicePath func(string) bool

// Determines if @device_path is considered a block device path which is only
// used in implementation of the OS. This is primarily used for hiding
// mounted volumes that are intended as APIs for programs to read, and system
// administrators at a shell; rather than something that should, for example,
// appear in a GUI. For example, the Linux `/proc` filesystem.
//
// The list of device paths considered ‘system’ ones may change over time.
func UnixIsSystemDevicePath(DevicePathVar string) bool {

	cret := xUnixIsSystemDevicePath(DevicePathVar)
	return cret
}

var xUnixIsSystemFsType func(string) bool

// Determines if @fs_type is considered a type of file system which is only
// used in implementation of the OS. This is primarily used for hiding
// mounted volumes that are intended as APIs for programs to read, and system
// administrators at a shell; rather than something that should, for example,
// appear in a GUI. For example, the Linux `/proc` filesystem.
//
// The list of file system types considered ‘system’ ones may change over time.
func UnixIsSystemFsType(FsTypeVar string) bool {

	cret := xUnixIsSystemFsType(FsTypeVar)
	return cret
}

var xUnixMountAt func(string, uint64) *UnixMountEntry

// Gets a #GUnixMountEntry for a given mount path. If @time_read
// is set, it will be filled with a unix timestamp for checking
// if the mounts have changed since with g_unix_mounts_changed_since().
//
// If more mounts have the same mount path, the last matching mount
// is returned.
//
// This will return %NULL if there is no mount point at @mount_path.
func UnixMountAt(MountPathVar string, TimeReadVar uint64) *UnixMountEntry {

	cret := xUnixMountAt(MountPathVar, TimeReadVar)
	return cret
}

var xUnixMountCompare func(*UnixMountEntry, *UnixMountEntry) int

// Compares two unix mounts.
func UnixMountCompare(Mount1Var *UnixMountEntry, Mount2Var *UnixMountEntry) int {

	cret := xUnixMountCompare(Mount1Var, Mount2Var)
	return cret
}

var xUnixMountCopy func(*UnixMountEntry) *UnixMountEntry

// Makes a copy of @mount_entry.
func UnixMountCopy(MountEntryVar *UnixMountEntry) *UnixMountEntry {

	cret := xUnixMountCopy(MountEntryVar)
	return cret
}

var xUnixMountFor func(string, uint64) *UnixMountEntry

// Gets a #GUnixMountEntry for a given file path. If @time_read
// is set, it will be filled with a unix timestamp for checking
// if the mounts have changed since with g_unix_mounts_changed_since().
//
// If more mounts have the same mount path, the last matching mount
// is returned.
//
// This will return %NULL if looking up the mount entry fails, if
// @file_path doesn’t exist or there is an I/O error.
func UnixMountFor(FilePathVar string, TimeReadVar uint64) *UnixMountEntry {

	cret := xUnixMountFor(FilePathVar, TimeReadVar)
	return cret
}

var xUnixMountFree func(*UnixMountEntry)

// Frees a unix mount.
func UnixMountFree(MountEntryVar *UnixMountEntry) {

	xUnixMountFree(MountEntryVar)

}

var xUnixMountGetDevicePath func(*UnixMountEntry) string

// Gets the device path for a unix mount.
func UnixMountGetDevicePath(MountEntryVar *UnixMountEntry) string {

	cret := xUnixMountGetDevicePath(MountEntryVar)
	return cret
}

var xUnixMountGetFsType func(*UnixMountEntry) string

// Gets the filesystem type for the unix mount.
func UnixMountGetFsType(MountEntryVar *UnixMountEntry) string {

	cret := xUnixMountGetFsType(MountEntryVar)
	return cret
}

var xUnixMountGetMountPath func(*UnixMountEntry) string

// Gets the mount path for a unix mount.
func UnixMountGetMountPath(MountEntryVar *UnixMountEntry) string {

	cret := xUnixMountGetMountPath(MountEntryVar)
	return cret
}

var xUnixMountGetOptions func(*UnixMountEntry) string

// Gets a comma-separated list of mount options for the unix mount. For example,
// `rw,relatime,seclabel,data=ordered`.
//
// This is similar to g_unix_mount_point_get_options(), but it takes
// a #GUnixMountEntry as an argument.
func UnixMountGetOptions(MountEntryVar *UnixMountEntry) string {

	cret := xUnixMountGetOptions(MountEntryVar)
	return cret
}

var xUnixMountGetRootPath func(*UnixMountEntry) string

// Gets the root of the mount within the filesystem. This is useful e.g. for
// mounts created by bind operation, or btrfs subvolumes.
//
// For example, the root path is equal to "/" for mount created by
// "mount /dev/sda1 /mnt/foo" and "/bar" for
// "mount --bind /mnt/foo/bar /mnt/bar".
func UnixMountGetRootPath(MountEntryVar *UnixMountEntry) string {

	cret := xUnixMountGetRootPath(MountEntryVar)
	return cret
}

var xUnixMountGuessCanEject func(*UnixMountEntry) bool

// Guesses whether a Unix mount can be ejected.
func UnixMountGuessCanEject(MountEntryVar *UnixMountEntry) bool {

	cret := xUnixMountGuessCanEject(MountEntryVar)
	return cret
}

var xUnixMountGuessIcon func(*UnixMountEntry) uintptr

// Guesses the icon of a Unix mount.
func UnixMountGuessIcon(MountEntryVar *UnixMountEntry) *IconBase {
	var cls *IconBase

	cret := xUnixMountGuessIcon(MountEntryVar)

	if cret == 0 {
		return nil
	}
	cls = &IconBase{}
	cls.Ptr = cret
	return cls
}

var xUnixMountGuessName func(*UnixMountEntry) string

// Guesses the name of a Unix mount.
// The result is a translated string.
func UnixMountGuessName(MountEntryVar *UnixMountEntry) string {

	cret := xUnixMountGuessName(MountEntryVar)
	return cret
}

var xUnixMountGuessShouldDisplay func(*UnixMountEntry) bool

// Guesses whether a Unix mount should be displayed in the UI.
func UnixMountGuessShouldDisplay(MountEntryVar *UnixMountEntry) bool {

	cret := xUnixMountGuessShouldDisplay(MountEntryVar)
	return cret
}

var xUnixMountGuessSymbolicIcon func(*UnixMountEntry) uintptr

// Guesses the symbolic icon of a Unix mount.
func UnixMountGuessSymbolicIcon(MountEntryVar *UnixMountEntry) *IconBase {
	var cls *IconBase

	cret := xUnixMountGuessSymbolicIcon(MountEntryVar)

	if cret == 0 {
		return nil
	}
	cls = &IconBase{}
	cls.Ptr = cret
	return cls
}

var xUnixMountIsReadonly func(*UnixMountEntry) bool

// Checks if a unix mount is mounted read only.
func UnixMountIsReadonly(MountEntryVar *UnixMountEntry) bool {

	cret := xUnixMountIsReadonly(MountEntryVar)
	return cret
}

var xUnixMountIsSystemInternal func(*UnixMountEntry) bool

// Checks if a Unix mount is a system mount. This is the Boolean OR of
// g_unix_is_system_fs_type(), g_unix_is_system_device_path() and
// g_unix_is_mount_path_system_internal() on @mount_entry’s properties.
//
// The definition of what a ‘system’ mount entry is may change over time as new
// file system types and device paths are ignored.
func UnixMountIsSystemInternal(MountEntryVar *UnixMountEntry) bool {

	cret := xUnixMountIsSystemInternal(MountEntryVar)
	return cret
}

var xUnixMountPointAt func(string, uint64) *UnixMountPoint

// Gets a #GUnixMountPoint for a given mount path. If @time_read is set, it
// will be filled with a unix timestamp for checking if the mount points have
// changed since with g_unix_mount_points_changed_since().
//
// If more mount points have the same mount path, the last matching mount point
// is returned.
func UnixMountPointAt(MountPathVar string, TimeReadVar uint64) *UnixMountPoint {

	cret := xUnixMountPointAt(MountPathVar, TimeReadVar)
	return cret
}

var xUnixMountPointsChangedSince func(uint64) bool

// Checks if the unix mount points have changed since a given unix time.
func UnixMountPointsChangedSince(TimeVar uint64) bool {

	cret := xUnixMountPointsChangedSince(TimeVar)
	return cret
}

var xUnixMountPointsGet func(uint64) *glib.List

// Gets a #GList of #GUnixMountPoint containing the unix mount points.
// If @time_read is set, it will be filled with the mount timestamp,
// allowing for checking if the mounts have changed with
// g_unix_mount_points_changed_since().
func UnixMountPointsGet(TimeReadVar uint64) *glib.List {

	cret := xUnixMountPointsGet(TimeReadVar)
	return cret
}

var xUnixMountsChangedSince func(uint64) bool

// Checks if the unix mounts have changed since a given unix time.
func UnixMountsChangedSince(TimeVar uint64) bool {

	cret := xUnixMountsChangedSince(TimeVar)
	return cret
}

var xUnixMountsGet func(uint64) *glib.List

// Gets a #GList of #GUnixMountEntry containing the unix mounts.
// If @time_read is set, it will be filled with the mount
// timestamp, allowing for checking if the mounts have changed
// with g_unix_mounts_changed_since().
func UnixMountsGet(TimeReadVar uint64) *glib.List {

	cret := xUnixMountsGet(TimeReadVar)
	return cret
}

// Watches #GUnixMounts for changes.
type UnixMountMonitor struct {
	gobject.Object
}

func UnixMountMonitorNewFromInternalPtr(ptr uintptr) *UnixMountMonitor {
	cls := &UnixMountMonitor{}
	cls.Ptr = ptr
	return cls
}

var xNewUnixMountMonitor func() uintptr

// Deprecated alias for g_unix_mount_monitor_get().
//
// This function was never a true constructor, which is why it was
// renamed.
func NewUnixMountMonitor() *UnixMountMonitor {
	var cls *UnixMountMonitor

	cret := xNewUnixMountMonitor()

	if cret == 0 {
		return nil
	}
	cls = &UnixMountMonitor{}
	cls.Ptr = cret
	return cls
}

var xUnixMountMonitorSetRateLimit func(uintptr, int)

// This function does nothing.
//
// Before 2.44, this was a partially-effective way of controlling the
// rate at which events would be reported under some uncommon
// circumstances.  Since @mount_monitor is a singleton, it also meant
// that calling this function would have side effects for other users of
// the monitor.
func (x *UnixMountMonitor) SetRateLimit(LimitMsecVar int) {

	xUnixMountMonitorSetRateLimit(x.GoPointer(), LimitMsecVar)

}

func (c *UnixMountMonitor) GoPointer() uintptr {
	return c.Ptr
}

func (c *UnixMountMonitor) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// Emitted when the unix mount points have changed.
func (x *UnixMountMonitor) ConnectMountpointsChanged(cb func(UnixMountMonitor)) uint32 {
	fcb := func(clsPtr uintptr) {
		fa := UnixMountMonitor{}
		fa.Ptr = clsPtr

		cb(fa)

	}
	return gobject.SignalConnect(x.GoPointer(), "mountpoints-changed", purego.NewCallback(fcb))
}

// Emitted when the unix mounts have changed.
func (x *UnixMountMonitor) ConnectMountsChanged(cb func(UnixMountMonitor)) uint32 {
	fcb := func(clsPtr uintptr) {
		fa := UnixMountMonitor{}
		fa.Ptr = clsPtr

		cb(fa)

	}
	return gobject.SignalConnect(x.GoPointer(), "mounts-changed", purego.NewCallback(fcb))
}

var xUnixMountMonitorGet func() uintptr

// Gets the #GUnixMountMonitor for the current thread-default main
// context.
//
// The mount monitor can be used to monitor for changes to the list of
// mounted filesystems as well as the list of mount points (ie: fstab
// entries).
//
// You must only call g_object_unref() on the return value from under
// the same main context as you called this function.
func UnixMountMonitorGet() *UnixMountMonitor {
	var cls *UnixMountMonitor

	cret := xUnixMountMonitorGet()

	if cret == 0 {
		return nil
	}
	cls = &UnixMountMonitor{}
	cls.Ptr = cret
	return cls
}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GIO"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	core.PuregoSafeRegister(&xUnixIsMountPathSystemInternal, lib, "g_unix_is_mount_path_system_internal")
	core.PuregoSafeRegister(&xUnixIsSystemDevicePath, lib, "g_unix_is_system_device_path")
	core.PuregoSafeRegister(&xUnixIsSystemFsType, lib, "g_unix_is_system_fs_type")
	core.PuregoSafeRegister(&xUnixMountAt, lib, "g_unix_mount_at")
	core.PuregoSafeRegister(&xUnixMountCompare, lib, "g_unix_mount_compare")
	core.PuregoSafeRegister(&xUnixMountCopy, lib, "g_unix_mount_copy")
	core.PuregoSafeRegister(&xUnixMountFor, lib, "g_unix_mount_for")
	core.PuregoSafeRegister(&xUnixMountFree, lib, "g_unix_mount_free")
	core.PuregoSafeRegister(&xUnixMountGetDevicePath, lib, "g_unix_mount_get_device_path")
	core.PuregoSafeRegister(&xUnixMountGetFsType, lib, "g_unix_mount_get_fs_type")
	core.PuregoSafeRegister(&xUnixMountGetMountPath, lib, "g_unix_mount_get_mount_path")
	core.PuregoSafeRegister(&xUnixMountGetOptions, lib, "g_unix_mount_get_options")
	core.PuregoSafeRegister(&xUnixMountGetRootPath, lib, "g_unix_mount_get_root_path")
	core.PuregoSafeRegister(&xUnixMountGuessCanEject, lib, "g_unix_mount_guess_can_eject")
	core.PuregoSafeRegister(&xUnixMountGuessIcon, lib, "g_unix_mount_guess_icon")
	core.PuregoSafeRegister(&xUnixMountGuessName, lib, "g_unix_mount_guess_name")
	core.PuregoSafeRegister(&xUnixMountGuessShouldDisplay, lib, "g_unix_mount_guess_should_display")
	core.PuregoSafeRegister(&xUnixMountGuessSymbolicIcon, lib, "g_unix_mount_guess_symbolic_icon")
	core.PuregoSafeRegister(&xUnixMountIsReadonly, lib, "g_unix_mount_is_readonly")
	core.PuregoSafeRegister(&xUnixMountIsSystemInternal, lib, "g_unix_mount_is_system_internal")
	core.PuregoSafeRegister(&xUnixMountPointAt, lib, "g_unix_mount_point_at")
	core.PuregoSafeRegister(&xUnixMountPointsChangedSince, lib, "g_unix_mount_points_changed_since")
	core.PuregoSafeRegister(&xUnixMountPointsGet, lib, "g_unix_mount_points_get")
	core.PuregoSafeRegister(&xUnixMountsChangedSince, lib, "g_unix_mounts_changed_since")
	core.PuregoSafeRegister(&xUnixMountsGet, lib, "g_unix_mounts_get")

	core.PuregoSafeRegister(&xUnixMountPointCompare, lib, "g_unix_mount_point_compare")
	core.PuregoSafeRegister(&xUnixMountPointCopy, lib, "g_unix_mount_point_copy")
	core.PuregoSafeRegister(&xUnixMountPointFree, lib, "g_unix_mount_point_free")
	core.PuregoSafeRegister(&xUnixMountPointGetDevicePath, lib, "g_unix_mount_point_get_device_path")
	core.PuregoSafeRegister(&xUnixMountPointGetFsType, lib, "g_unix_mount_point_get_fs_type")
	core.PuregoSafeRegister(&xUnixMountPointGetMountPath, lib, "g_unix_mount_point_get_mount_path")
	core.PuregoSafeRegister(&xUnixMountPointGetOptions, lib, "g_unix_mount_point_get_options")
	core.PuregoSafeRegister(&xUnixMountPointGuessCanEject, lib, "g_unix_mount_point_guess_can_eject")
	core.PuregoSafeRegister(&xUnixMountPointGuessIcon, lib, "g_unix_mount_point_guess_icon")
	core.PuregoSafeRegister(&xUnixMountPointGuessName, lib, "g_unix_mount_point_guess_name")
	core.PuregoSafeRegister(&xUnixMountPointGuessSymbolicIcon, lib, "g_unix_mount_point_guess_symbolic_icon")
	core.PuregoSafeRegister(&xUnixMountPointIsLoopback, lib, "g_unix_mount_point_is_loopback")
	core.PuregoSafeRegister(&xUnixMountPointIsReadonly, lib, "g_unix_mount_point_is_readonly")
	core.PuregoSafeRegister(&xUnixMountPointIsUserMountable, lib, "g_unix_mount_point_is_user_mountable")

	core.PuregoSafeRegister(&xNewUnixMountMonitor, lib, "g_unix_mount_monitor_new")

	core.PuregoSafeRegister(&xUnixMountMonitorSetRateLimit, lib, "g_unix_mount_monitor_set_rate_limit")

	core.PuregoSafeRegister(&xUnixMountMonitorGet, lib, "g_unix_mount_monitor_get")

}
