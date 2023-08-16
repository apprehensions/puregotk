// Package gio was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gio

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/glib"
	"github.com/jwijenbergh/puregotk/v4/gobject"
)

type VolumeMonitorClass struct {
	ParentClass uintptr
}

// #GVolumeMonitor is for listing the user interesting devices and volumes
// on the computer. In other words, what a file selector or file manager
// would show in a sidebar.
//
// #GVolumeMonitor is not
// [thread-default-context aware][g-main-context-push-thread-default],
// and so should not be used other than from the main thread, with no
// thread-default-context active.
//
// In order to receive updates about volumes and mounts monitored through GVFS,
// a main loop must be running.
type VolumeMonitor struct {
	gobject.Object
}

func VolumeMonitorNewFromInternalPtr(ptr uintptr) *VolumeMonitor {
	cls := &VolumeMonitor{}
	cls.Ptr = ptr
	return cls
}

var xVolumeMonitorGetConnectedDrives func(uintptr) *glib.List

// Gets a list of drives connected to the system.
//
// The returned list should be freed with g_list_free(), after
// its elements have been unreffed with g_object_unref().
func (x *VolumeMonitor) GetConnectedDrives() *glib.List {

	cret := xVolumeMonitorGetConnectedDrives(x.GoPointer())
	return cret
}

var xVolumeMonitorGetMountForUuid func(uintptr, string) uintptr

// Finds a #GMount object by its UUID (see g_mount_get_uuid())
func (x *VolumeMonitor) GetMountForUuid(UuidVar string) *MountBase {
	var cls *MountBase

	cret := xVolumeMonitorGetMountForUuid(x.GoPointer(), UuidVar)

	if cret == 0 {
		return nil
	}
	cls = &MountBase{}
	cls.Ptr = cret
	return cls
}

var xVolumeMonitorGetMounts func(uintptr) *glib.List

// Gets a list of the mounts on the system.
//
// The returned list should be freed with g_list_free(), after
// its elements have been unreffed with g_object_unref().
func (x *VolumeMonitor) GetMounts() *glib.List {

	cret := xVolumeMonitorGetMounts(x.GoPointer())
	return cret
}

var xVolumeMonitorGetVolumeForUuid func(uintptr, string) uintptr

// Finds a #GVolume object by its UUID (see g_volume_get_uuid())
func (x *VolumeMonitor) GetVolumeForUuid(UuidVar string) *VolumeBase {
	var cls *VolumeBase

	cret := xVolumeMonitorGetVolumeForUuid(x.GoPointer(), UuidVar)

	if cret == 0 {
		return nil
	}
	cls = &VolumeBase{}
	cls.Ptr = cret
	return cls
}

var xVolumeMonitorGetVolumes func(uintptr) *glib.List

// Gets a list of the volumes on the system.
//
// The returned list should be freed with g_list_free(), after
// its elements have been unreffed with g_object_unref().
func (x *VolumeMonitor) GetVolumes() *glib.List {

	cret := xVolumeMonitorGetVolumes(x.GoPointer())
	return cret
}

func (c *VolumeMonitor) GoPointer() uintptr {
	return c.Ptr
}

func (c *VolumeMonitor) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// Emitted when a drive changes.
func (x *VolumeMonitor) ConnectDriveChanged(cb func(VolumeMonitor, uintptr)) {
	fcb := func(clsPtr uintptr, DriveVarp uintptr) {
		fa := VolumeMonitor{}
		fa.Ptr = clsPtr

		cb(fa, DriveVarp)

	}
	gobject.ObjectConnect(x.GoPointer(), "signal::drive-changed", purego.NewCallback(fcb))
}

// Emitted when a drive is connected to the system.
func (x *VolumeMonitor) ConnectDriveConnected(cb func(VolumeMonitor, uintptr)) {
	fcb := func(clsPtr uintptr, DriveVarp uintptr) {
		fa := VolumeMonitor{}
		fa.Ptr = clsPtr

		cb(fa, DriveVarp)

	}
	gobject.ObjectConnect(x.GoPointer(), "signal::drive-connected", purego.NewCallback(fcb))
}

// Emitted when a drive is disconnected from the system.
func (x *VolumeMonitor) ConnectDriveDisconnected(cb func(VolumeMonitor, uintptr)) {
	fcb := func(clsPtr uintptr, DriveVarp uintptr) {
		fa := VolumeMonitor{}
		fa.Ptr = clsPtr

		cb(fa, DriveVarp)

	}
	gobject.ObjectConnect(x.GoPointer(), "signal::drive-disconnected", purego.NewCallback(fcb))
}

// Emitted when the eject button is pressed on @drive.
func (x *VolumeMonitor) ConnectDriveEjectButton(cb func(VolumeMonitor, uintptr)) {
	fcb := func(clsPtr uintptr, DriveVarp uintptr) {
		fa := VolumeMonitor{}
		fa.Ptr = clsPtr

		cb(fa, DriveVarp)

	}
	gobject.ObjectConnect(x.GoPointer(), "signal::drive-eject-button", purego.NewCallback(fcb))
}

// Emitted when the stop button is pressed on @drive.
func (x *VolumeMonitor) ConnectDriveStopButton(cb func(VolumeMonitor, uintptr)) {
	fcb := func(clsPtr uintptr, DriveVarp uintptr) {
		fa := VolumeMonitor{}
		fa.Ptr = clsPtr

		cb(fa, DriveVarp)

	}
	gobject.ObjectConnect(x.GoPointer(), "signal::drive-stop-button", purego.NewCallback(fcb))
}

// Emitted when a mount is added.
func (x *VolumeMonitor) ConnectMountAdded(cb func(VolumeMonitor, uintptr)) {
	fcb := func(clsPtr uintptr, MountVarp uintptr) {
		fa := VolumeMonitor{}
		fa.Ptr = clsPtr

		cb(fa, MountVarp)

	}
	gobject.ObjectConnect(x.GoPointer(), "signal::mount-added", purego.NewCallback(fcb))
}

// Emitted when a mount changes.
func (x *VolumeMonitor) ConnectMountChanged(cb func(VolumeMonitor, uintptr)) {
	fcb := func(clsPtr uintptr, MountVarp uintptr) {
		fa := VolumeMonitor{}
		fa.Ptr = clsPtr

		cb(fa, MountVarp)

	}
	gobject.ObjectConnect(x.GoPointer(), "signal::mount-changed", purego.NewCallback(fcb))
}

// May be emitted when a mount is about to be removed.
//
// This signal depends on the backend and is only emitted if
// GIO was used to unmount.
func (x *VolumeMonitor) ConnectMountPreUnmount(cb func(VolumeMonitor, uintptr)) {
	fcb := func(clsPtr uintptr, MountVarp uintptr) {
		fa := VolumeMonitor{}
		fa.Ptr = clsPtr

		cb(fa, MountVarp)

	}
	gobject.ObjectConnect(x.GoPointer(), "signal::mount-pre-unmount", purego.NewCallback(fcb))
}

// Emitted when a mount is removed.
func (x *VolumeMonitor) ConnectMountRemoved(cb func(VolumeMonitor, uintptr)) {
	fcb := func(clsPtr uintptr, MountVarp uintptr) {
		fa := VolumeMonitor{}
		fa.Ptr = clsPtr

		cb(fa, MountVarp)

	}
	gobject.ObjectConnect(x.GoPointer(), "signal::mount-removed", purego.NewCallback(fcb))
}

// Emitted when a mountable volume is added to the system.
func (x *VolumeMonitor) ConnectVolumeAdded(cb func(VolumeMonitor, uintptr)) {
	fcb := func(clsPtr uintptr, VolumeVarp uintptr) {
		fa := VolumeMonitor{}
		fa.Ptr = clsPtr

		cb(fa, VolumeVarp)

	}
	gobject.ObjectConnect(x.GoPointer(), "signal::volume-added", purego.NewCallback(fcb))
}

// Emitted when mountable volume is changed.
func (x *VolumeMonitor) ConnectVolumeChanged(cb func(VolumeMonitor, uintptr)) {
	fcb := func(clsPtr uintptr, VolumeVarp uintptr) {
		fa := VolumeMonitor{}
		fa.Ptr = clsPtr

		cb(fa, VolumeVarp)

	}
	gobject.ObjectConnect(x.GoPointer(), "signal::volume-changed", purego.NewCallback(fcb))
}

// Emitted when a mountable volume is removed from the system.
func (x *VolumeMonitor) ConnectVolumeRemoved(cb func(VolumeMonitor, uintptr)) {
	fcb := func(clsPtr uintptr, VolumeVarp uintptr) {
		fa := VolumeMonitor{}
		fa.Ptr = clsPtr

		cb(fa, VolumeVarp)

	}
	gobject.ObjectConnect(x.GoPointer(), "signal::volume-removed", purego.NewCallback(fcb))
}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GIO"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xVolumeMonitorGetConnectedDrives, lib, "g_volume_monitor_get_connected_drives")
	core.PuregoSafeRegister(&xVolumeMonitorGetMountForUuid, lib, "g_volume_monitor_get_mount_for_uuid")
	core.PuregoSafeRegister(&xVolumeMonitorGetMounts, lib, "g_volume_monitor_get_mounts")
	core.PuregoSafeRegister(&xVolumeMonitorGetVolumeForUuid, lib, "g_volume_monitor_get_volume_for_uuid")
	core.PuregoSafeRegister(&xVolumeMonitorGetVolumes, lib, "g_volume_monitor_get_volumes")

}
