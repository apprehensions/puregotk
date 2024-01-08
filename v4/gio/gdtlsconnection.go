// Package gio was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gio

import (
	"unsafe"

	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/glib"
	"github.com/jwijenbergh/puregotk/v4/gobject"
)

// Virtual method table for a #GDtlsConnection implementation.
type DtlsConnectionInterface struct {
	GIface uintptr
}

func (x *DtlsConnectionInterface) GoPointer() uintptr {
	return uintptr(unsafe.Pointer(x))
}

// #GDtlsConnection is the base DTLS connection class type, which wraps
// a #GDatagramBased and provides DTLS encryption on top of it. Its
// subclasses, #GDtlsClientConnection and #GDtlsServerConnection,
// implement client-side and server-side DTLS, respectively.
//
// For TLS support, see #GTlsConnection.
//
// As DTLS is datagram based, #GDtlsConnection implements #GDatagramBased,
// presenting a datagram-socket-like API for the encrypted connection. This
// operates over a base datagram connection, which is also a #GDatagramBased
// (#GDtlsConnection:base-socket).
//
// To close a DTLS connection, use g_dtls_connection_close().
//
// Neither #GDtlsServerConnection or #GDtlsClientConnection set the peer address
// on their base #GDatagramBased if it is a #GSocket — it is up to the caller to
// do that if they wish. If they do not, and g_socket_close() is called on the
// base socket, the #GDtlsConnection will not raise a %G_IO_ERROR_NOT_CONNECTED
// error on further I/O.
type DtlsConnection interface {
	GoPointer() uintptr
	SetGoPointer(uintptr)
	Close(CancellableVar *Cancellable) bool
	CloseAsync(IoPriorityVar int, CancellableVar *Cancellable, CallbackVar AsyncReadyCallback, UserDataVar uintptr)
	CloseFinish(ResultVar AsyncResult) bool
	EmitAcceptCertificate(PeerCertVar *TlsCertificate, ErrorsVar TlsCertificateFlags) bool
	GetCertificate() *TlsCertificate
	GetChannelBindingData(TypeVar TlsChannelBindingType, DataVar uintptr) bool
	GetCiphersuiteName() string
	GetDatabase() *TlsDatabase
	GetInteraction() *TlsInteraction
	GetNegotiatedProtocol() string
	GetPeerCertificate() *TlsCertificate
	GetPeerCertificateErrors() TlsCertificateFlags
	GetProtocolVersion() TlsProtocolVersion
	GetRehandshakeMode() TlsRehandshakeMode
	GetRequireCloseNotify() bool
	Handshake(CancellableVar *Cancellable) bool
	HandshakeAsync(IoPriorityVar int, CancellableVar *Cancellable, CallbackVar AsyncReadyCallback, UserDataVar uintptr)
	HandshakeFinish(ResultVar AsyncResult) bool
	SetAdvertisedProtocols(ProtocolsVar uintptr)
	SetCertificate(CertificateVar *TlsCertificate)
	SetDatabase(DatabaseVar *TlsDatabase)
	SetInteraction(InteractionVar *TlsInteraction)
	SetRehandshakeMode(ModeVar TlsRehandshakeMode)
	SetRequireCloseNotify(RequireCloseNotifyVar bool)
	Shutdown(ShutdownReadVar bool, ShutdownWriteVar bool, CancellableVar *Cancellable) bool
	ShutdownAsync(ShutdownReadVar bool, ShutdownWriteVar bool, IoPriorityVar int, CancellableVar *Cancellable, CallbackVar AsyncReadyCallback, UserDataVar uintptr)
	ShutdownFinish(ResultVar AsyncResult) bool
}
type DtlsConnectionBase struct {
	Ptr uintptr
}

func (x *DtlsConnectionBase) GoPointer() uintptr {
	return x.Ptr
}

func (x *DtlsConnectionBase) SetGoPointer(ptr uintptr) {
	x.Ptr = ptr
}

// Close the DTLS connection. This is equivalent to calling
// g_dtls_connection_shutdown() to shut down both sides of the connection.
//
// Closing a #GDtlsConnection waits for all buffered but untransmitted data to
// be sent before it completes. It then sends a `close_notify` DTLS alert to the
// peer and may wait for a `close_notify` to be received from the peer. It does
// not close the underlying #GDtlsConnection:base-socket; that must be closed
// separately.
//
// Once @conn is closed, all other operations will return %G_IO_ERROR_CLOSED.
// Closing a #GDtlsConnection multiple times will not return an error.
//
// #GDtlsConnections will be automatically closed when the last reference is
// dropped, but you might want to call this function to make sure resources are
// released as early as possible.
//
// If @cancellable is cancelled, the #GDtlsConnection may be left
// partially-closed and any pending untransmitted data may be lost. Call
// g_dtls_connection_close() again to complete closing the #GDtlsConnection.
func (x *DtlsConnectionBase) Close(CancellableVar *Cancellable) (bool, error) {
	var cerr *glib.Error

	cret := XGDtlsConnectionClose(x.GoPointer(), CancellableVar.GoPointer(), &cerr)
	if cerr == nil {
		return cret, nil
	}
	return cret, cerr

}

// Asynchronously close the DTLS connection. See g_dtls_connection_close() for
// more information.
func (x *DtlsConnectionBase) CloseAsync(IoPriorityVar int, CancellableVar *Cancellable, CallbackVar AsyncReadyCallback, UserDataVar uintptr) {

	XGDtlsConnectionCloseAsync(x.GoPointer(), IoPriorityVar, CancellableVar.GoPointer(), purego.NewCallback(CallbackVar), UserDataVar)

}

// Finish an asynchronous TLS close operation. See g_dtls_connection_close()
// for more information.
func (x *DtlsConnectionBase) CloseFinish(ResultVar AsyncResult) (bool, error) {
	var cerr *glib.Error

	cret := XGDtlsConnectionCloseFinish(x.GoPointer(), ResultVar.GoPointer(), &cerr)
	if cerr == nil {
		return cret, nil
	}
	return cret, cerr

}

// Used by #GDtlsConnection implementations to emit the
// #GDtlsConnection::accept-certificate signal.
func (x *DtlsConnectionBase) EmitAcceptCertificate(PeerCertVar *TlsCertificate, ErrorsVar TlsCertificateFlags) bool {

	cret := XGDtlsConnectionEmitAcceptCertificate(x.GoPointer(), PeerCertVar.GoPointer(), ErrorsVar)
	return cret
}

// Gets @conn's certificate, as set by
// g_dtls_connection_set_certificate().
func (x *DtlsConnectionBase) GetCertificate() *TlsCertificate {
	var cls *TlsCertificate

	cret := XGDtlsConnectionGetCertificate(x.GoPointer())

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &TlsCertificate{}
	cls.Ptr = cret
	return cls
}

// Query the TLS backend for TLS channel binding data of @type for @conn.
//
// This call retrieves TLS channel binding data as specified in RFC
// [5056](https://tools.ietf.org/html/rfc5056), RFC
// [5929](https://tools.ietf.org/html/rfc5929), and related RFCs.  The
// binding data is returned in @data.  The @data is resized by the callee
// using #GByteArray buffer management and will be freed when the @data
// is destroyed by g_byte_array_unref(). If @data is %NULL, it will only
// check whether TLS backend is able to fetch the data (e.g. whether @type
// is supported by the TLS backend). It does not guarantee that the data
// will be available though.  That could happen if TLS connection does not
// support @type or the binding data is not available yet due to additional
// negotiation or input required.
func (x *DtlsConnectionBase) GetChannelBindingData(TypeVar TlsChannelBindingType, DataVar uintptr) (bool, error) {
	var cerr *glib.Error

	cret := XGDtlsConnectionGetChannelBindingData(x.GoPointer(), TypeVar, DataVar, &cerr)
	if cerr == nil {
		return cret, nil
	}
	return cret, cerr

}

// Returns the name of the current DTLS ciphersuite, or %NULL if the
// connection has not handshaked or has been closed. Beware that the TLS
// backend may use any of multiple different naming conventions, because
// OpenSSL and GnuTLS have their own ciphersuite naming conventions that
// are different from each other and different from the standard, IANA-
// registered ciphersuite names. The ciphersuite name is intended to be
// displayed to the user for informative purposes only, and parsing it
// is not recommended.
func (x *DtlsConnectionBase) GetCiphersuiteName() string {

	cret := XGDtlsConnectionGetCiphersuiteName(x.GoPointer())
	return cret
}

// Gets the certificate database that @conn uses to verify
// peer certificates. See g_dtls_connection_set_database().
func (x *DtlsConnectionBase) GetDatabase() *TlsDatabase {
	var cls *TlsDatabase

	cret := XGDtlsConnectionGetDatabase(x.GoPointer())

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &TlsDatabase{}
	cls.Ptr = cret
	return cls
}

// Get the object that will be used to interact with the user. It will be used
// for things like prompting the user for passwords. If %NULL is returned, then
// no user interaction will occur for this connection.
func (x *DtlsConnectionBase) GetInteraction() *TlsInteraction {
	var cls *TlsInteraction

	cret := XGDtlsConnectionGetInteraction(x.GoPointer())

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &TlsInteraction{}
	cls.Ptr = cret
	return cls
}

// Gets the name of the application-layer protocol negotiated during
// the handshake.
//
// If the peer did not use the ALPN extension, or did not advertise a
// protocol that matched one of @conn's protocols, or the TLS backend
// does not support ALPN, then this will be %NULL. See
// g_dtls_connection_set_advertised_protocols().
func (x *DtlsConnectionBase) GetNegotiatedProtocol() string {

	cret := XGDtlsConnectionGetNegotiatedProtocol(x.GoPointer())
	return cret
}

// Gets @conn's peer's certificate after the handshake has completed
// or failed. (It is not set during the emission of
// #GDtlsConnection::accept-certificate.)
func (x *DtlsConnectionBase) GetPeerCertificate() *TlsCertificate {
	var cls *TlsCertificate

	cret := XGDtlsConnectionGetPeerCertificate(x.GoPointer())

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &TlsCertificate{}
	cls.Ptr = cret
	return cls
}

// Gets the errors associated with validating @conn's peer's
// certificate, after the handshake has completed or failed. (It is
// not set during the emission of #GDtlsConnection::accept-certificate.)
func (x *DtlsConnectionBase) GetPeerCertificateErrors() TlsCertificateFlags {

	cret := XGDtlsConnectionGetPeerCertificateErrors(x.GoPointer())
	return cret
}

// Returns the current DTLS protocol version, which may be
// %G_TLS_PROTOCOL_VERSION_UNKNOWN if the connection has not handshaked, or
// has been closed, or if the TLS backend has implemented a protocol version
// that is not a recognized #GTlsProtocolVersion.
func (x *DtlsConnectionBase) GetProtocolVersion() TlsProtocolVersion {

	cret := XGDtlsConnectionGetProtocolVersion(x.GoPointer())
	return cret
}

// Gets @conn rehandshaking mode. See
// g_dtls_connection_set_rehandshake_mode() for details.
func (x *DtlsConnectionBase) GetRehandshakeMode() TlsRehandshakeMode {

	cret := XGDtlsConnectionGetRehandshakeMode(x.GoPointer())
	return cret
}

// Tests whether or not @conn expects a proper TLS close notification
// when the connection is closed. See
// g_dtls_connection_set_require_close_notify() for details.
func (x *DtlsConnectionBase) GetRequireCloseNotify() bool {

	cret := XGDtlsConnectionGetRequireCloseNotify(x.GoPointer())
	return cret
}

// Attempts a TLS handshake on @conn.
//
// On the client side, it is never necessary to call this method;
// although the connection needs to perform a handshake after
// connecting, #GDtlsConnection will handle this for you automatically
// when you try to send or receive data on the connection. You can call
// g_dtls_connection_handshake() manually if you want to know whether
// the initial handshake succeeded or failed (as opposed to just
// immediately trying to use @conn to read or write, in which case,
// if it fails, it may not be possible to tell if it failed before
// or after completing the handshake), but beware that servers may reject
// client authentication after the handshake has completed, so a
// successful handshake does not indicate the connection will be usable.
//
// Likewise, on the server side, although a handshake is necessary at
// the beginning of the communication, you do not need to call this
// function explicitly unless you want clearer error reporting.
//
// Previously, calling g_dtls_connection_handshake() after the initial
// handshake would trigger a rehandshake; however, this usage was
// deprecated in GLib 2.60 because rehandshaking was removed from the
// TLS protocol in TLS 1.3. Since GLib 2.64, calling this function after
// the initial handshake will no longer do anything.
//
// #GDtlsConnection::accept_certificate may be emitted during the
// handshake.
func (x *DtlsConnectionBase) Handshake(CancellableVar *Cancellable) (bool, error) {
	var cerr *glib.Error

	cret := XGDtlsConnectionHandshake(x.GoPointer(), CancellableVar.GoPointer(), &cerr)
	if cerr == nil {
		return cret, nil
	}
	return cret, cerr

}

// Asynchronously performs a TLS handshake on @conn. See
// g_dtls_connection_handshake() for more information.
func (x *DtlsConnectionBase) HandshakeAsync(IoPriorityVar int, CancellableVar *Cancellable, CallbackVar AsyncReadyCallback, UserDataVar uintptr) {

	XGDtlsConnectionHandshakeAsync(x.GoPointer(), IoPriorityVar, CancellableVar.GoPointer(), purego.NewCallback(CallbackVar), UserDataVar)

}

// Finish an asynchronous TLS handshake operation. See
// g_dtls_connection_handshake() for more information.
func (x *DtlsConnectionBase) HandshakeFinish(ResultVar AsyncResult) (bool, error) {
	var cerr *glib.Error

	cret := XGDtlsConnectionHandshakeFinish(x.GoPointer(), ResultVar.GoPointer(), &cerr)
	if cerr == nil {
		return cret, nil
	}
	return cret, cerr

}

// Sets the list of application-layer protocols to advertise that the
// caller is willing to speak on this connection. The
// Application-Layer Protocol Negotiation (ALPN) extension will be
// used to negotiate a compatible protocol with the peer; use
// g_dtls_connection_get_negotiated_protocol() to find the negotiated
// protocol after the handshake.  Specifying %NULL for the the value
// of @protocols will disable ALPN negotiation.
//
// See [IANA TLS ALPN Protocol IDs](https://www.iana.org/assignments/tls-extensiontype-values/tls-extensiontype-values.xhtml#alpn-protocol-ids)
// for a list of registered protocol IDs.
func (x *DtlsConnectionBase) SetAdvertisedProtocols(ProtocolsVar uintptr) {

	XGDtlsConnectionSetAdvertisedProtocols(x.GoPointer(), ProtocolsVar)

}

// This sets the certificate that @conn will present to its peer
// during the TLS handshake. For a #GDtlsServerConnection, it is
// mandatory to set this, and that will normally be done at construct
// time.
//
// For a #GDtlsClientConnection, this is optional. If a handshake fails
// with %G_TLS_ERROR_CERTIFICATE_REQUIRED, that means that the server
// requires a certificate, and if you try connecting again, you should
// call this method first. You can call
// g_dtls_client_connection_get_accepted_cas() on the failed connection
// to get a list of Certificate Authorities that the server will
// accept certificates from.
//
// (It is also possible that a server will allow the connection with
// or without a certificate; in that case, if you don't provide a
// certificate, you can tell that the server requested one by the fact
// that g_dtls_client_connection_get_accepted_cas() will return
// non-%NULL.)
func (x *DtlsConnectionBase) SetCertificate(CertificateVar *TlsCertificate) {

	XGDtlsConnectionSetCertificate(x.GoPointer(), CertificateVar.GoPointer())

}

// Sets the certificate database that is used to verify peer certificates.
// This is set to the default database by default. See
// g_tls_backend_get_default_database(). If set to %NULL, then
// peer certificate validation will always set the
// %G_TLS_CERTIFICATE_UNKNOWN_CA error (meaning
// #GDtlsConnection::accept-certificate will always be emitted on
// client-side connections, unless that bit is not set in
// #GDtlsClientConnection:validation-flags).
//
// There are nonintuitive security implications when using a non-default
// database. See #GDtlsConnection:database for details.
func (x *DtlsConnectionBase) SetDatabase(DatabaseVar *TlsDatabase) {

	XGDtlsConnectionSetDatabase(x.GoPointer(), DatabaseVar.GoPointer())

}

// Set the object that will be used to interact with the user. It will be used
// for things like prompting the user for passwords.
//
// The @interaction argument will normally be a derived subclass of
// #GTlsInteraction. %NULL can also be provided if no user interaction
// should occur for this connection.
func (x *DtlsConnectionBase) SetInteraction(InteractionVar *TlsInteraction) {

	XGDtlsConnectionSetInteraction(x.GoPointer(), InteractionVar.GoPointer())

}

// Since GLib 2.64, changing the rehandshake mode is no longer supported
// and will have no effect. With TLS 1.3, rehandshaking has been removed from
// the TLS protocol, replaced by separate post-handshake authentication and
// rekey operations.
func (x *DtlsConnectionBase) SetRehandshakeMode(ModeVar TlsRehandshakeMode) {

	XGDtlsConnectionSetRehandshakeMode(x.GoPointer(), ModeVar)

}

// Sets whether or not @conn expects a proper TLS close notification
// before the connection is closed. If this is %TRUE (the default),
// then @conn will expect to receive a TLS close notification from its
// peer before the connection is closed, and will return a
// %G_TLS_ERROR_EOF error if the connection is closed without proper
// notification (since this may indicate a network error, or
// man-in-the-middle attack).
//
// In some protocols, the application will know whether or not the
// connection was closed cleanly based on application-level data
// (because the application-level data includes a length field, or is
// somehow self-delimiting); in this case, the close notify is
// redundant and may be omitted. You
// can use g_dtls_connection_set_require_close_notify() to tell @conn
// to allow an "unannounced" connection close, in which case the close
// will show up as a 0-length read, as in a non-TLS
// #GDatagramBased, and it is up to the application to check that
// the data has been fully received.
//
// Note that this only affects the behavior when the peer closes the
// connection; when the application calls g_dtls_connection_close_async() on
// @conn itself, this will send a close notification regardless of the
// setting of this property. If you explicitly want to do an unclean
// close, you can close @conn's #GDtlsConnection:base-socket rather
// than closing @conn itself.
func (x *DtlsConnectionBase) SetRequireCloseNotify(RequireCloseNotifyVar bool) {

	XGDtlsConnectionSetRequireCloseNotify(x.GoPointer(), RequireCloseNotifyVar)

}

// Shut down part or all of a DTLS connection.
//
// If @shutdown_read is %TRUE then the receiving side of the connection is shut
// down, and further reading is disallowed. Subsequent calls to
// g_datagram_based_receive_messages() will return %G_IO_ERROR_CLOSED.
//
// If @shutdown_write is %TRUE then the sending side of the connection is shut
// down, and further writing is disallowed. Subsequent calls to
// g_datagram_based_send_messages() will return %G_IO_ERROR_CLOSED.
//
// It is allowed for both @shutdown_read and @shutdown_write to be TRUE — this
// is equivalent to calling g_dtls_connection_close().
//
// If @cancellable is cancelled, the #GDtlsConnection may be left
// partially-closed and any pending untransmitted data may be lost. Call
// g_dtls_connection_shutdown() again to complete closing the #GDtlsConnection.
func (x *DtlsConnectionBase) Shutdown(ShutdownReadVar bool, ShutdownWriteVar bool, CancellableVar *Cancellable) (bool, error) {
	var cerr *glib.Error

	cret := XGDtlsConnectionShutdown(x.GoPointer(), ShutdownReadVar, ShutdownWriteVar, CancellableVar.GoPointer(), &cerr)
	if cerr == nil {
		return cret, nil
	}
	return cret, cerr

}

// Asynchronously shut down part or all of the DTLS connection. See
// g_dtls_connection_shutdown() for more information.
func (x *DtlsConnectionBase) ShutdownAsync(ShutdownReadVar bool, ShutdownWriteVar bool, IoPriorityVar int, CancellableVar *Cancellable, CallbackVar AsyncReadyCallback, UserDataVar uintptr) {

	XGDtlsConnectionShutdownAsync(x.GoPointer(), ShutdownReadVar, ShutdownWriteVar, IoPriorityVar, CancellableVar.GoPointer(), purego.NewCallback(CallbackVar), UserDataVar)

}

// Finish an asynchronous TLS shutdown operation. See
// g_dtls_connection_shutdown() for more information.
func (x *DtlsConnectionBase) ShutdownFinish(ResultVar AsyncResult) (bool, error) {
	var cerr *glib.Error

	cret := XGDtlsConnectionShutdownFinish(x.GoPointer(), ResultVar.GoPointer(), &cerr)
	if cerr == nil {
		return cret, nil
	}
	return cret, cerr

}

var XGDtlsConnectionClose func(uintptr, uintptr, **glib.Error) bool
var XGDtlsConnectionCloseAsync func(uintptr, int, uintptr, uintptr, uintptr)
var XGDtlsConnectionCloseFinish func(uintptr, uintptr, **glib.Error) bool
var XGDtlsConnectionEmitAcceptCertificate func(uintptr, uintptr, TlsCertificateFlags) bool
var XGDtlsConnectionGetCertificate func(uintptr) uintptr
var XGDtlsConnectionGetChannelBindingData func(uintptr, TlsChannelBindingType, uintptr, **glib.Error) bool
var XGDtlsConnectionGetCiphersuiteName func(uintptr) string
var XGDtlsConnectionGetDatabase func(uintptr) uintptr
var XGDtlsConnectionGetInteraction func(uintptr) uintptr
var XGDtlsConnectionGetNegotiatedProtocol func(uintptr) string
var XGDtlsConnectionGetPeerCertificate func(uintptr) uintptr
var XGDtlsConnectionGetPeerCertificateErrors func(uintptr) TlsCertificateFlags
var XGDtlsConnectionGetProtocolVersion func(uintptr) TlsProtocolVersion
var XGDtlsConnectionGetRehandshakeMode func(uintptr) TlsRehandshakeMode
var XGDtlsConnectionGetRequireCloseNotify func(uintptr) bool
var XGDtlsConnectionHandshake func(uintptr, uintptr, **glib.Error) bool
var XGDtlsConnectionHandshakeAsync func(uintptr, int, uintptr, uintptr, uintptr)
var XGDtlsConnectionHandshakeFinish func(uintptr, uintptr, **glib.Error) bool
var XGDtlsConnectionSetAdvertisedProtocols func(uintptr, uintptr)
var XGDtlsConnectionSetCertificate func(uintptr, uintptr)
var XGDtlsConnectionSetDatabase func(uintptr, uintptr)
var XGDtlsConnectionSetInteraction func(uintptr, uintptr)
var XGDtlsConnectionSetRehandshakeMode func(uintptr, TlsRehandshakeMode)
var XGDtlsConnectionSetRequireCloseNotify func(uintptr, bool)
var XGDtlsConnectionShutdown func(uintptr, bool, bool, uintptr, **glib.Error) bool
var XGDtlsConnectionShutdownAsync func(uintptr, bool, bool, int, uintptr, uintptr, uintptr)
var XGDtlsConnectionShutdownFinish func(uintptr, uintptr, **glib.Error) bool

func init() {
	lib, err := purego.Dlopen(core.GetPath("GIO"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&XGDtlsConnectionClose, lib, "g_dtls_connection_close")
	core.PuregoSafeRegister(&XGDtlsConnectionCloseAsync, lib, "g_dtls_connection_close_async")
	core.PuregoSafeRegister(&XGDtlsConnectionCloseFinish, lib, "g_dtls_connection_close_finish")
	core.PuregoSafeRegister(&XGDtlsConnectionEmitAcceptCertificate, lib, "g_dtls_connection_emit_accept_certificate")
	core.PuregoSafeRegister(&XGDtlsConnectionGetCertificate, lib, "g_dtls_connection_get_certificate")
	core.PuregoSafeRegister(&XGDtlsConnectionGetChannelBindingData, lib, "g_dtls_connection_get_channel_binding_data")
	core.PuregoSafeRegister(&XGDtlsConnectionGetCiphersuiteName, lib, "g_dtls_connection_get_ciphersuite_name")
	core.PuregoSafeRegister(&XGDtlsConnectionGetDatabase, lib, "g_dtls_connection_get_database")
	core.PuregoSafeRegister(&XGDtlsConnectionGetInteraction, lib, "g_dtls_connection_get_interaction")
	core.PuregoSafeRegister(&XGDtlsConnectionGetNegotiatedProtocol, lib, "g_dtls_connection_get_negotiated_protocol")
	core.PuregoSafeRegister(&XGDtlsConnectionGetPeerCertificate, lib, "g_dtls_connection_get_peer_certificate")
	core.PuregoSafeRegister(&XGDtlsConnectionGetPeerCertificateErrors, lib, "g_dtls_connection_get_peer_certificate_errors")
	core.PuregoSafeRegister(&XGDtlsConnectionGetProtocolVersion, lib, "g_dtls_connection_get_protocol_version")
	core.PuregoSafeRegister(&XGDtlsConnectionGetRehandshakeMode, lib, "g_dtls_connection_get_rehandshake_mode")
	core.PuregoSafeRegister(&XGDtlsConnectionGetRequireCloseNotify, lib, "g_dtls_connection_get_require_close_notify")
	core.PuregoSafeRegister(&XGDtlsConnectionHandshake, lib, "g_dtls_connection_handshake")
	core.PuregoSafeRegister(&XGDtlsConnectionHandshakeAsync, lib, "g_dtls_connection_handshake_async")
	core.PuregoSafeRegister(&XGDtlsConnectionHandshakeFinish, lib, "g_dtls_connection_handshake_finish")
	core.PuregoSafeRegister(&XGDtlsConnectionSetAdvertisedProtocols, lib, "g_dtls_connection_set_advertised_protocols")
	core.PuregoSafeRegister(&XGDtlsConnectionSetCertificate, lib, "g_dtls_connection_set_certificate")
	core.PuregoSafeRegister(&XGDtlsConnectionSetDatabase, lib, "g_dtls_connection_set_database")
	core.PuregoSafeRegister(&XGDtlsConnectionSetInteraction, lib, "g_dtls_connection_set_interaction")
	core.PuregoSafeRegister(&XGDtlsConnectionSetRehandshakeMode, lib, "g_dtls_connection_set_rehandshake_mode")
	core.PuregoSafeRegister(&XGDtlsConnectionSetRequireCloseNotify, lib, "g_dtls_connection_set_require_close_notify")
	core.PuregoSafeRegister(&XGDtlsConnectionShutdown, lib, "g_dtls_connection_shutdown")
	core.PuregoSafeRegister(&XGDtlsConnectionShutdownAsync, lib, "g_dtls_connection_shutdown_async")
	core.PuregoSafeRegister(&XGDtlsConnectionShutdownFinish, lib, "g_dtls_connection_shutdown_finish")

}
