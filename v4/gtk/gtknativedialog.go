// Package gtk was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gtk

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/gobject"
)

// Class structure for `GtkNativeDialog`.
type NativeDialogClass struct {
	ParentClass uintptr
}

// Native dialogs are platform dialogs that don't use `GtkDialog`.
//
// They are used in order to integrate better with a platform, by
// looking the same as other native applications and supporting
// platform specific features.
//
// The [class@Gtk.Dialog] functions cannot be used on such objects,
// but we need a similar API in order to drive them. The `GtkNativeDialog`
// object is an API that allows you to do this. It allows you to set
// various common properties on the dialog, as well as show and hide
// it and get a [signal@Gtk.NativeDialog::response] signal when the user
// finished with the dialog.
//
// Note that unlike `GtkDialog`, `GtkNativeDialog` objects are not
// toplevel widgets, and GTK does not keep them alive. It is your
// responsibility to keep a reference until you are done with the
// object.
type NativeDialog struct {
	gobject.Object
}

func NativeDialogNewFromInternalPtr(ptr uintptr) *NativeDialog {
	cls := &NativeDialog{}
	cls.Ptr = ptr
	return cls
}

var xNativeDialogDestroy func(uintptr)

// Destroys a dialog.
//
// When a dialog is destroyed, it will break any references it holds
// to other objects.
//
// If it is visible it will be hidden and any underlying window system
// resources will be destroyed.
//
// Note that this does not release any reference to the object (as opposed
// to destroying a `GtkWindow`) because there is no reference from the
// windowing system to the `GtkNativeDialog`.
func (x *NativeDialog) Destroy() {

	xNativeDialogDestroy(x.GoPointer())

}

var xNativeDialogGetModal func(uintptr) bool

// Returns whether the dialog is modal.
func (x *NativeDialog) GetModal() bool {

	return xNativeDialogGetModal(x.GoPointer())

}

var xNativeDialogGetTitle func(uintptr) string

// Gets the title of the `GtkNativeDialog`.
func (x *NativeDialog) GetTitle() string {

	return xNativeDialogGetTitle(x.GoPointer())

}

var xNativeDialogGetTransientFor func(uintptr) uintptr

// Fetches the transient parent for this window.
func (x *NativeDialog) GetTransientFor() *Window {

	GetTransientForPtr := xNativeDialogGetTransientFor(x.GoPointer())
	if GetTransientForPtr == 0 {
		return nil
	}

	gobject.IncreaseRef(GetTransientForPtr)

	GetTransientForCls := &Window{}
	GetTransientForCls.Ptr = GetTransientForPtr
	return GetTransientForCls

}

var xNativeDialogGetVisible func(uintptr) bool

// Determines whether the dialog is visible.
func (x *NativeDialog) GetVisible() bool {

	return xNativeDialogGetVisible(x.GoPointer())

}

var xNativeDialogHide func(uintptr)

// Hides the dialog if it is visible, aborting any interaction.
//
// Once this is called the [signal@Gtk.NativeDialog::response] signal
// will *not* be emitted until after the next call to
// [method@Gtk.NativeDialog.show].
//
// If the dialog is not visible this does nothing.
func (x *NativeDialog) Hide() {

	xNativeDialogHide(x.GoPointer())

}

var xNativeDialogSetModal func(uintptr, bool)

// Sets a dialog modal or non-modal.
//
// Modal dialogs prevent interaction with other windows in the same
// application. To keep modal dialogs on top of main application
// windows, use [method@Gtk.NativeDialog.set_transient_for] to make
// the dialog transient for the parent; most window managers will
// then disallow lowering the dialog below the parent.
func (x *NativeDialog) SetModal(ModalVar bool) {

	xNativeDialogSetModal(x.GoPointer(), ModalVar)

}

var xNativeDialogSetTitle func(uintptr, string)

// Sets the title of the `GtkNativeDialog.`
func (x *NativeDialog) SetTitle(TitleVar string) {

	xNativeDialogSetTitle(x.GoPointer(), TitleVar)

}

var xNativeDialogSetTransientFor func(uintptr, uintptr)

// Dialog windows should be set transient for the main application
// window they were spawned from.
//
// This allows window managers to e.g. keep the dialog on top of the
// main window, or center the dialog over the main window.
//
// Passing %NULL for @parent unsets the current transient window.
func (x *NativeDialog) SetTransientFor(ParentVar *Window) {

	xNativeDialogSetTransientFor(x.GoPointer(), ParentVar.GoPointer())

}

var xNativeDialogShow func(uintptr)

// Shows the dialog on the display.
//
// When the user accepts the state of the dialog the dialog will
// be automatically hidden and the [signal@Gtk.NativeDialog::response]
// signal will be emitted.
//
// Multiple calls while the dialog is visible will be ignored.
func (x *NativeDialog) Show() {

	xNativeDialogShow(x.GoPointer())

}

func (c *NativeDialog) GoPointer() uintptr {
	return c.Ptr
}

func (c *NativeDialog) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// Emitted when the user responds to the dialog.
//
// When this is called the dialog has been hidden.
//
// If you call [method@Gtk.NativeDialog.hide] before the user
// responds to the dialog this signal will not be emitted.
func (x *NativeDialog) ConnectResponse(cb func(NativeDialog, int32)) {
	fcb := func(clsPtr uintptr, ResponseIdVarp int32) {
		fa := NativeDialog{}
		fa.Ptr = clsPtr

		cb(fa, ResponseIdVarp)

	}
	gobject.ObjectConnect(x.GoPointer(), "signal::response", purego.NewCallback(fcb))
}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GTK"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xNativeDialogDestroy, lib, "gtk_native_dialog_destroy")
	core.PuregoSafeRegister(&xNativeDialogGetModal, lib, "gtk_native_dialog_get_modal")
	core.PuregoSafeRegister(&xNativeDialogGetTitle, lib, "gtk_native_dialog_get_title")
	core.PuregoSafeRegister(&xNativeDialogGetTransientFor, lib, "gtk_native_dialog_get_transient_for")
	core.PuregoSafeRegister(&xNativeDialogGetVisible, lib, "gtk_native_dialog_get_visible")
	core.PuregoSafeRegister(&xNativeDialogHide, lib, "gtk_native_dialog_hide")
	core.PuregoSafeRegister(&xNativeDialogSetModal, lib, "gtk_native_dialog_set_modal")
	core.PuregoSafeRegister(&xNativeDialogSetTitle, lib, "gtk_native_dialog_set_title")
	core.PuregoSafeRegister(&xNativeDialogSetTransientFor, lib, "gtk_native_dialog_set_transient_for")
	core.PuregoSafeRegister(&xNativeDialogShow, lib, "gtk_native_dialog_show")

}
