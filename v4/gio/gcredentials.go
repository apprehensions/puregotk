// Package gio was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gio

import (
	"unsafe"

	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/glib"
	"github.com/jwijenbergh/puregotk/v4/gobject"
)

// Class structure for #GCredentials.
type CredentialsClass struct {
}

func (x *CredentialsClass) GoPointer() uintptr {
	return uintptr(unsafe.Pointer(x))
}

// The #GCredentials type is a reference-counted wrapper for native
// credentials. This information is typically used for identifying,
// authenticating and authorizing other processes.
//
// Some operating systems supports looking up the credentials of the
// remote peer of a communication endpoint - see e.g.
// g_socket_get_credentials().
//
// Some operating systems supports securely sending and receiving
// credentials over a Unix Domain Socket, see
// #GUnixCredentialsMessage, g_unix_connection_send_credentials() and
// g_unix_connection_receive_credentials() for details.
//
// On Linux, the native credential type is a `struct ucred` - see the
// unix(7) man page for details. This corresponds to
// %G_CREDENTIALS_TYPE_LINUX_UCRED.
//
// On Apple operating systems (including iOS, tvOS, and macOS),
// the native credential type is a `struct xucred`.
// This corresponds to %G_CREDENTIALS_TYPE_APPLE_XUCRED.
//
// On FreeBSD, Debian GNU/kFreeBSD, and GNU/Hurd, the native
// credential type is a `struct cmsgcred`. This corresponds
// to %G_CREDENTIALS_TYPE_FREEBSD_CMSGCRED.
//
// On NetBSD, the native credential type is a `struct unpcbid`.
// This corresponds to %G_CREDENTIALS_TYPE_NETBSD_UNPCBID.
//
// On OpenBSD, the native credential type is a `struct sockpeercred`.
// This corresponds to %G_CREDENTIALS_TYPE_OPENBSD_SOCKPEERCRED.
//
// On Solaris (including OpenSolaris and its derivatives), the native
// credential type is a `ucred_t`. This corresponds to
// %G_CREDENTIALS_TYPE_SOLARIS_UCRED.
//
// Since GLib 2.72, on Windows, the native credentials may contain the PID of a
// process. This corresponds to %G_CREDENTIALS_TYPE_WIN32_PID.
type Credentials struct {
	gobject.Object
}

func CredentialsNewFromInternalPtr(ptr uintptr) *Credentials {
	cls := &Credentials{}
	cls.Ptr = ptr
	return cls
}

var xNewCredentials func() uintptr

// Creates a new #GCredentials object with credentials matching the
// the current process.
func NewCredentials() *Credentials {
	var cls *Credentials

	cret := xNewCredentials()

	if cret == 0 {
		return nil
	}
	cls = &Credentials{}
	cls.Ptr = cret
	return cls
}

var xCredentialsGetNative func(uintptr, CredentialsType) uintptr

// Gets a pointer to native credentials of type @native_type from
// @credentials.
//
// It is a programming error (which will cause a warning to be
// logged) to use this method if there is no #GCredentials support for
// the OS or if @native_type isn't supported by the OS.
func (x *Credentials) GetNative(NativeTypeVar CredentialsType) uintptr {

	cret := xCredentialsGetNative(x.GoPointer(), NativeTypeVar)
	return cret
}

var xCredentialsGetUnixPid func(uintptr) int

// Tries to get the UNIX process identifier from @credentials. This
// method is only available on UNIX platforms.
//
// This operation can fail if #GCredentials is not supported on the
// OS or if the native credentials type does not contain information
// about the UNIX process ID.
func (x *Credentials) GetUnixPid() (int, error) {
	var cerr *glib.Error

	cret := xCredentialsGetUnixPid(x.GoPointer())
	if cerr == nil {
		return cret, nil
	}
	return cret, cerr

}

var xCredentialsGetUnixUser func(uintptr) uint

// Tries to get the UNIX user identifier from @credentials. This
// method is only available on UNIX platforms.
//
// This operation can fail if #GCredentials is not supported on the
// OS or if the native credentials type does not contain information
// about the UNIX user.
func (x *Credentials) GetUnixUser() (uint, error) {
	var cerr *glib.Error

	cret := xCredentialsGetUnixUser(x.GoPointer())
	if cerr == nil {
		return cret, nil
	}
	return cret, cerr

}

var xCredentialsIsSameUser func(uintptr, uintptr, **glib.Error) bool

// Checks if @credentials and @other_credentials is the same user.
//
// This operation can fail if #GCredentials is not supported on the
// the OS.
func (x *Credentials) IsSameUser(OtherCredentialsVar *Credentials) (bool, error) {
	var cerr *glib.Error

	cret := xCredentialsIsSameUser(x.GoPointer(), OtherCredentialsVar.GoPointer(), &cerr)
	if cerr == nil {
		return cret, nil
	}
	return cret, cerr

}

var xCredentialsSetNative func(uintptr, CredentialsType, uintptr)

// Copies the native credentials of type @native_type from @native
// into @credentials.
//
// It is a programming error (which will cause a warning to be
// logged) to use this method if there is no #GCredentials support for
// the OS or if @native_type isn't supported by the OS.
func (x *Credentials) SetNative(NativeTypeVar CredentialsType, NativeVar uintptr) {

	xCredentialsSetNative(x.GoPointer(), NativeTypeVar, NativeVar)

}

var xCredentialsSetUnixUser func(uintptr, uint, **glib.Error) bool

// Tries to set the UNIX user identifier on @credentials. This method
// is only available on UNIX platforms.
//
// This operation can fail if #GCredentials is not supported on the
// OS or if the native credentials type does not contain information
// about the UNIX user. It can also fail if the OS does not allow the
// use of "spoofed" credentials.
func (x *Credentials) SetUnixUser(UidVar uint) (bool, error) {
	var cerr *glib.Error

	cret := xCredentialsSetUnixUser(x.GoPointer(), UidVar, &cerr)
	if cerr == nil {
		return cret, nil
	}
	return cret, cerr

}

var xCredentialsToString func(uintptr) string

// Creates a human-readable textual representation of @credentials
// that can be used in logging and debug messages. The format of the
// returned string may change in future GLib release.
func (x *Credentials) ToString() string {

	cret := xCredentialsToString(x.GoPointer())
	return cret
}

func (c *Credentials) GoPointer() uintptr {
	return c.Ptr
}

func (c *Credentials) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GIO"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xNewCredentials, lib, "g_credentials_new")

	core.PuregoSafeRegister(&xCredentialsGetNative, lib, "g_credentials_get_native")
	core.PuregoSafeRegister(&xCredentialsGetUnixPid, lib, "g_credentials_get_unix_pid")
	core.PuregoSafeRegister(&xCredentialsGetUnixUser, lib, "g_credentials_get_unix_user")
	core.PuregoSafeRegister(&xCredentialsIsSameUser, lib, "g_credentials_is_same_user")
	core.PuregoSafeRegister(&xCredentialsSetNative, lib, "g_credentials_set_native")
	core.PuregoSafeRegister(&xCredentialsSetUnixUser, lib, "g_credentials_set_unix_user")
	core.PuregoSafeRegister(&xCredentialsToString, lib, "g_credentials_to_string")

}
