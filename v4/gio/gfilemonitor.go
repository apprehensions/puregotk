// Package gio was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gio

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/gobject"
)

type FileMonitorClass struct {
	ParentClass uintptr
}

type FileMonitorPrivate struct {
}

// Monitors a file or directory for changes.
//
// To obtain a #GFileMonitor for a file or directory, use
// g_file_monitor(), g_file_monitor_file(), or
// g_file_monitor_directory().
//
// To get informed about changes to the file or directory you are
// monitoring, connect to the #GFileMonitor::changed signal. The
// signal will be emitted in the
// [thread-default main context][g-main-context-push-thread-default]
// of the thread that the monitor was created in
// (though if the global default main context is blocked, this may
// cause notifications to be blocked even if the thread-default
// context is still running).
type FileMonitor struct {
	gobject.Object
}

func FileMonitorNewFromInternalPtr(ptr uintptr) *FileMonitor {
	cls := &FileMonitor{}
	cls.Ptr = ptr
	return cls
}

var xFileMonitorCancel func(uintptr) bool

// Cancels a file monitor.
func (x *FileMonitor) Cancel() bool {

	return xFileMonitorCancel(x.GoPointer())

}

var xFileMonitorEmitEvent func(uintptr, uintptr, uintptr, FileMonitorEvent)

// Emits the #GFileMonitor::changed signal if a change
// has taken place. Should be called from file monitor
// implementations only.
//
// Implementations are responsible to call this method from the
// [thread-default main context][g-main-context-push-thread-default] of the
// thread that the monitor was created in.
func (x *FileMonitor) EmitEvent(ChildVar File, OtherFileVar File, EventTypeVar FileMonitorEvent) {

	xFileMonitorEmitEvent(x.GoPointer(), ChildVar.GoPointer(), OtherFileVar.GoPointer(), EventTypeVar)

}

var xFileMonitorIsCancelled func(uintptr) bool

// Returns whether the monitor is canceled.
func (x *FileMonitor) IsCancelled() bool {

	return xFileMonitorIsCancelled(x.GoPointer())

}

var xFileMonitorSetRateLimit func(uintptr, int32)

// Sets the rate limit to which the @monitor will report
// consecutive change events to the same file.
func (x *FileMonitor) SetRateLimit(LimitMsecsVar int32) {

	xFileMonitorSetRateLimit(x.GoPointer(), LimitMsecsVar)

}

func (c *FileMonitor) GoPointer() uintptr {
	return c.Ptr
}

func (c *FileMonitor) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// Emitted when @file has been changed.
//
// If using %G_FILE_MONITOR_WATCH_MOVES on a directory monitor, and
// the information is available (and if supported by the backend),
// @event_type may be %G_FILE_MONITOR_EVENT_RENAMED,
// %G_FILE_MONITOR_EVENT_MOVED_IN or %G_FILE_MONITOR_EVENT_MOVED_OUT.
//
// In all cases @file will be a child of the monitored directory.  For
// renames, @file will be the old name and @other_file is the new
// name.  For "moved in" events, @file is the name of the file that
// appeared and @other_file is the old name that it was moved from (in
// another directory).  For "moved out" events, @file is the name of
// the file that used to be in this directory and @other_file is the
// name of the file at its new location.
//
// It makes sense to treat %G_FILE_MONITOR_EVENT_MOVED_IN as
// equivalent to %G_FILE_MONITOR_EVENT_CREATED and
// %G_FILE_MONITOR_EVENT_MOVED_OUT as equivalent to
// %G_FILE_MONITOR_EVENT_DELETED, with extra information.
// %G_FILE_MONITOR_EVENT_RENAMED is equivalent to a delete/create
// pair.  This is exactly how the events will be reported in the case
// that the %G_FILE_MONITOR_WATCH_MOVES flag is not in use.
//
// If using the deprecated flag %G_FILE_MONITOR_SEND_MOVED flag and @event_type is
// %G_FILE_MONITOR_EVENT_MOVED, @file will be set to a #GFile containing the
// old path, and @other_file will be set to a #GFile containing the new path.
//
// In all the other cases, @other_file will be set to #NULL.
func (x *FileMonitor) ConnectChanged(cb func(FileMonitor, uintptr, uintptr, FileMonitorEvent)) {
	fcb := func(clsPtr uintptr, FileVarp uintptr, OtherFileVarp uintptr, EventTypeVarp FileMonitorEvent) {
		fa := FileMonitor{}
		fa.Ptr = clsPtr

		cb(fa, FileVarp, OtherFileVarp, EventTypeVarp)

	}
	gobject.ObjectConnect(x.GoPointer(), "signal::changed", purego.NewCallback(fcb))
}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GIO"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xFileMonitorCancel, lib, "g_file_monitor_cancel")
	core.PuregoSafeRegister(&xFileMonitorEmitEvent, lib, "g_file_monitor_emit_event")
	core.PuregoSafeRegister(&xFileMonitorIsCancelled, lib, "g_file_monitor_is_cancelled")
	core.PuregoSafeRegister(&xFileMonitorSetRateLimit, lib, "g_file_monitor_set_rate_limit")

}
