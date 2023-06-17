// Package gtk was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gtk

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/gio"
)

// `GtkAppChooser` is an interface for widgets which allow the user to
// choose an application.
//
// The main objects that implement this interface are
// [class@Gtk.AppChooserWidget],
// [class@Gtk.AppChooserDialog] and [class@Gtk.AppChooserButton].
//
// Applications are represented by GIO `GAppInfo` objects here.
// GIO has a concept of recommended and fallback applications for a
// given content type. Recommended applications are those that claim
// to handle the content type itself, while fallback also includes
// applications that handle a more generic content type. GIO also
// knows the default and last-used application for a given content
// type. The `GtkAppChooserWidget` provides detailed control over
// whether the shown list of applications should include default,
// recommended or fallback applications.
//
// To obtain the application that has been selected in a `GtkAppChooser`,
// use [method@Gtk.AppChooser.get_app_info].
type AppChooser interface {
	GoPointer() uintptr
	SetGoPointer(uintptr)
	GetAppInfo() *gio.AppInfoBase
	GetContentType() string
	Refresh()
}
type AppChooserBase struct {
	Ptr uintptr
}

func (x *AppChooserBase) GoPointer() uintptr {
	return x.Ptr
}

func (x *AppChooserBase) SetGoPointer(ptr uintptr) {
	x.Ptr = ptr
}

// Returns the currently selected application.
func (x *AppChooserBase) GetAppInfo() *gio.AppInfoBase {

	GetAppInfoPtr := XGtkAppChooserGetAppInfo(x.GoPointer())
	if GetAppInfoPtr == 0 {
		return nil
	}

	GetAppInfoCls := &gio.AppInfoBase{}
	GetAppInfoCls.Ptr = GetAppInfoPtr
	return GetAppInfoCls

}

// Returns the content type for which the `GtkAppChooser`
// shows applications.
func (x *AppChooserBase) GetContentType() string {

	return XGtkAppChooserGetContentType(x.GoPointer())

}

// Reloads the list of applications.
func (x *AppChooserBase) Refresh() {

	XGtkAppChooserRefresh(x.GoPointer())

}

var XGtkAppChooserGetAppInfo func(uintptr) uintptr
var XGtkAppChooserGetContentType func(uintptr) string
var XGtkAppChooserRefresh func(uintptr)

func init() {
	lib, err := purego.Dlopen(core.GetPath("GTK"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&XGtkAppChooserGetAppInfo, lib, "gtk_app_chooser_get_app_info")
	core.PuregoSafeRegister(&XGtkAppChooserGetContentType, lib, "gtk_app_chooser_get_content_type")
	core.PuregoSafeRegister(&XGtkAppChooserRefresh, lib, "gtk_app_chooser_refresh")

}
