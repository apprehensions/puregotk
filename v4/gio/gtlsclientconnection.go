// Package gio was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gio

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/glib"
	"github.com/jwijenbergh/puregotk/v4/gobject"
)

// vtable for a #GTlsClientConnection implementation.
type TlsClientConnectionInterface struct {
	GIface uintptr
}

// #GTlsClientConnection is the client-side subclass of
// #GTlsConnection, representing a client-side TLS connection.
type TlsClientConnection interface {
	GoPointer() uintptr
	SetGoPointer(uintptr)
	CopySessionState(SourceVar TlsClientConnection)
	GetAcceptedCas() *glib.List
	GetServerIdentity() *SocketConnectableBase
	GetUseSsl3() bool
	GetValidationFlags() TlsCertificateFlags
	SetServerIdentity(IdentityVar SocketConnectable)
	SetUseSsl3(UseSsl3Var bool)
	SetValidationFlags(FlagsVar TlsCertificateFlags)
}
type TlsClientConnectionBase struct {
	Ptr uintptr
}

func (x *TlsClientConnectionBase) GoPointer() uintptr {
	return x.Ptr
}

func (x *TlsClientConnectionBase) SetGoPointer(ptr uintptr) {
	x.Ptr = ptr
}

// Possibly copies session state from one connection to another, for use
// in TLS session resumption. This is not normally needed, but may be
// used when the same session needs to be used between different
// endpoints, as is required by some protocols, such as FTP over TLS.
// @source should have already completed a handshake and, since TLS 1.3,
// it should have been used to read data at least once. @conn should not
// have completed a handshake.
//
// It is not possible to know whether a call to this function will
// actually do anything. Because session resumption is normally used
// only for performance benefit, the TLS backend might not implement
// this function. Even if implemented, it may not actually succeed in
// allowing @conn to resume @source's TLS session, because the server
// may not have sent a session resumption token to @source, or it may
// refuse to accept the token from @conn. There is no way to know
// whether a call to this function is actually successful.
//
// Using this function is not required to benefit from session
// resumption. If the TLS backend supports session resumption, the
// session will be resumed automatically if it is possible to do so
// without weakening the privacy guarantees normally provided by TLS,
// without need to call this function. For example, with TLS 1.3,
// a session ticket will be automatically copied from any
// #GTlsClientConnection that has previously received session tickets
// from the server, provided a ticket is available that has not
// previously been used for session resumption, since session ticket
// reuse would be a privacy weakness. Using this function causes the
// ticket to be copied without regard for privacy considerations.
func (x *TlsClientConnectionBase) CopySessionState(SourceVar TlsClientConnection) {

	XGTlsClientConnectionCopySessionState(x.GoPointer(), SourceVar.GoPointer())

}

// Gets the list of distinguished names of the Certificate Authorities
// that the server will accept certificates from. This will be set
// during the TLS handshake if the server requests a certificate.
// Otherwise, it will be %NULL.
//
// Each item in the list is a #GByteArray which contains the complete
// subject DN of the certificate authority.
func (x *TlsClientConnectionBase) GetAcceptedCas() *glib.List {

	cret := XGTlsClientConnectionGetAcceptedCas(x.GoPointer())
	return cret
}

// Gets @conn's expected server identity
func (x *TlsClientConnectionBase) GetServerIdentity() *SocketConnectableBase {
	var cls *SocketConnectableBase

	cret := XGTlsClientConnectionGetServerIdentity(x.GoPointer())

	if cret == 0 {
		return cls
	}
	gobject.IncreaseRef(cret)
	cls = &SocketConnectableBase{}
	cls.Ptr = cret
	return cls
}

// SSL 3.0 is no longer supported. See
// g_tls_client_connection_set_use_ssl3() for details.
func (x *TlsClientConnectionBase) GetUseSsl3() bool {

	cret := XGTlsClientConnectionGetUseSsl3(x.GoPointer())
	return cret
}

// Gets @conn's validation flags
//
// This function does not work as originally designed and is impossible
// to use correctly. See #GTlsClientConnection:validation-flags for more
// information.
func (x *TlsClientConnectionBase) GetValidationFlags() TlsCertificateFlags {

	cret := XGTlsClientConnectionGetValidationFlags(x.GoPointer())
	return cret
}

// Sets @conn's expected server identity, which is used both to tell
// servers on virtual hosts which certificate to present, and also
// to let @conn know what name to look for in the certificate when
// performing %G_TLS_CERTIFICATE_BAD_IDENTITY validation, if enabled.
func (x *TlsClientConnectionBase) SetServerIdentity(IdentityVar SocketConnectable) {

	XGTlsClientConnectionSetServerIdentity(x.GoPointer(), IdentityVar.GoPointer())

}

// Since GLib 2.42.1, SSL 3.0 is no longer supported.
//
// From GLib 2.42.1 through GLib 2.62, this function could be used to
// force use of TLS 1.0, the lowest-supported TLS protocol version at
// the time. In the past, this was needed to connect to broken TLS
// servers that exhibited protocol version intolerance. Such servers
// are no longer common, and using TLS 1.0 is no longer considered
// acceptable.
//
// Since GLib 2.64, this function does nothing.
func (x *TlsClientConnectionBase) SetUseSsl3(UseSsl3Var bool) {

	XGTlsClientConnectionSetUseSsl3(x.GoPointer(), UseSsl3Var)

}

// Sets @conn's validation flags, to override the default set of
// checks performed when validating a server certificate. By default,
// %G_TLS_CERTIFICATE_VALIDATE_ALL is used.
//
// This function does not work as originally designed and is impossible
// to use correctly. See #GTlsClientConnection:validation-flags for more
// information.
func (x *TlsClientConnectionBase) SetValidationFlags(FlagsVar TlsCertificateFlags) {

	XGTlsClientConnectionSetValidationFlags(x.GoPointer(), FlagsVar)

}

var XGTlsClientConnectionCopySessionState func(uintptr, uintptr)
var XGTlsClientConnectionGetAcceptedCas func(uintptr) *glib.List
var XGTlsClientConnectionGetServerIdentity func(uintptr) uintptr
var XGTlsClientConnectionGetUseSsl3 func(uintptr) bool
var XGTlsClientConnectionGetValidationFlags func(uintptr) TlsCertificateFlags
var XGTlsClientConnectionSetServerIdentity func(uintptr, uintptr)
var XGTlsClientConnectionSetUseSsl3 func(uintptr, bool)
var XGTlsClientConnectionSetValidationFlags func(uintptr, TlsCertificateFlags)

var xTlsClientConnectionNew func(uintptr, uintptr, **glib.Error) uintptr

// Creates a new #GTlsClientConnection wrapping @base_io_stream (which
// must have pollable input and output streams) which is assumed to
// communicate with the server identified by @server_identity.
//
// See the documentation for #GTlsConnection:base-io-stream for restrictions
// on when application code can run operations on the @base_io_stream after
// this function has returned.
func TlsClientConnectionNew(BaseIoStreamVar *IOStream, ServerIdentityVar SocketConnectable) (*TlsClientConnectionBase, error) {
	var cls *TlsClientConnectionBase
	var cerr *glib.Error

	cret := xTlsClientConnectionNew(BaseIoStreamVar.GoPointer(), ServerIdentityVar.GoPointer(), &cerr)

	if cret == 0 {
		return cls, cerr
	}
	cls = &TlsClientConnectionBase{}
	cls.Ptr = cret
	if cerr == nil {
		return cls, nil
	}
	return cls, cerr

}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GIO"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	core.PuregoSafeRegister(&xTlsClientConnectionNew, lib, "g_tls_client_connection_new")

	core.PuregoSafeRegister(&XGTlsClientConnectionCopySessionState, lib, "g_tls_client_connection_copy_session_state")
	core.PuregoSafeRegister(&XGTlsClientConnectionGetAcceptedCas, lib, "g_tls_client_connection_get_accepted_cas")
	core.PuregoSafeRegister(&XGTlsClientConnectionGetServerIdentity, lib, "g_tls_client_connection_get_server_identity")
	core.PuregoSafeRegister(&XGTlsClientConnectionGetUseSsl3, lib, "g_tls_client_connection_get_use_ssl3")
	core.PuregoSafeRegister(&XGTlsClientConnectionGetValidationFlags, lib, "g_tls_client_connection_get_validation_flags")
	core.PuregoSafeRegister(&XGTlsClientConnectionSetServerIdentity, lib, "g_tls_client_connection_set_server_identity")
	core.PuregoSafeRegister(&XGTlsClientConnectionSetUseSsl3, lib, "g_tls_client_connection_set_use_ssl3")
	core.PuregoSafeRegister(&XGTlsClientConnectionSetValidationFlags, lib, "g_tls_client_connection_set_validation_flags")

}
