// Package gio was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gio

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
)

// An interface for file descriptor based io objects.
type FileDescriptorBasedIface struct {
	GIface uintptr
}

// #GFileDescriptorBased is implemented by streams (implementations of
// #GInputStream or #GOutputStream) that are based on file descriptors.
//
// Note that `&lt;gio/gfiledescriptorbased.h&gt;` belongs to the UNIX-specific
// GIO interfaces, thus you have to use the `gio-unix-2.0.pc` pkg-config
// file when using it.
type FileDescriptorBased interface {
	GoPointer() uintptr
	SetGoPointer(uintptr)
	GetFd() int32
}
type FileDescriptorBasedBase struct {
	Ptr uintptr
}

func (x *FileDescriptorBasedBase) GoPointer() uintptr {
	return x.Ptr
}

func (x *FileDescriptorBasedBase) SetGoPointer(ptr uintptr) {
	x.Ptr = ptr
}

// Gets the underlying file descriptor.
func (x *FileDescriptorBasedBase) GetFd() int32 {

	return XGFileDescriptorBasedGetFd(x.GoPointer())

}

var XGFileDescriptorBasedGetFd func(uintptr) int32

func init() {
	lib, err := purego.Dlopen(core.GetPath("GIO"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&XGFileDescriptorBasedGetFd, lib, "g_file_descriptor_based_get_fd")

}
