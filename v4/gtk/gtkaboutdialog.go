// Package gtk was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gtk

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/gdk"
	"github.com/jwijenbergh/puregotk/v4/gobject"
	"github.com/jwijenbergh/puregotk/v4/gsk"
)

// The type of license for an application.
//
// This enumeration can be expanded at later date.
type License int

const (

	// No license specified
	LicenseUnknownValue License = 0
	// A license text is going to be specified by the
	//   developer
	LicenseCustomValue License = 1
	// The GNU General Public License, version 2.0 or later
	LicenseGpl20Value License = 2
	// The GNU General Public License, version 3.0 or later
	LicenseGpl30Value License = 3
	// The GNU Lesser General Public License, version 2.1 or later
	LicenseLgpl21Value License = 4
	// The GNU Lesser General Public License, version 3.0 or later
	LicenseLgpl30Value License = 5
	// The BSD standard license
	LicenseBsdValue License = 6
	// The MIT/X11 standard license
	LicenseMitX11Value License = 7
	// The Artistic License, version 2.0
	LicenseArtisticValue License = 8
	// The GNU General Public License, version 2.0 only
	LicenseGpl20OnlyValue License = 9
	// The GNU General Public License, version 3.0 only
	LicenseGpl30OnlyValue License = 10
	// The GNU Lesser General Public License, version 2.1 only
	LicenseLgpl21OnlyValue License = 11
	// The GNU Lesser General Public License, version 3.0 only
	LicenseLgpl30OnlyValue License = 12
	// The GNU Affero General Public License, version 3.0 or later
	LicenseAgpl30Value License = 13
	// The GNU Affero General Public License, version 3.0 only
	LicenseAgpl30OnlyValue License = 14
	// The 3-clause BSD licence
	LicenseBsd3Value License = 15
	// The Apache License, version 2.0
	LicenseApache20Value License = 16
	// The Mozilla Public License, version 2.0
	LicenseMpl20Value License = 17
)

var xShowAboutDialog func(uintptr, string, ...interface{})

// A convenience function for showing an application’s about dialog.
//
// The constructed dialog is associated with the parent window and
// reused for future invocations of this function.
func ShowAboutDialog(ParentVar *Window, FirstPropertyNameVar string, varArgs ...interface{}) {

	xShowAboutDialog(ParentVar.GoPointer(), FirstPropertyNameVar, varArgs...)

}

// The `GtkAboutDialog` offers a simple way to display information about
// a program.
//
// The shown information includes the programs' logo, name, copyright,
// website and license. It is also possible to give credits to the authors,
// documenters, translators and artists who have worked on the program.
//
// An about dialog is typically opened when the user selects the `About`
// option from the `Help` menu. All parts of the dialog are optional.
//
// ![An example GtkAboutDialog](aboutdialog.png)
//
// About dialogs often contain links and email addresses. `GtkAboutDialog`
// displays these as clickable links. By default, it calls [func@Gtk.show_uri]
// when a user clicks one. The behaviour can be overridden with the
// [signal@Gtk.AboutDialog::activate-link] signal.
//
// To specify a person with an email address, use a string like
// `Edgar Allan Poe &lt;edgar@poe.com&gt;`. To specify a website with a title,
// use a string like `GTK team https://www.gtk.org`.
//
// To make constructing a `GtkAboutDialog` as convenient as possible, you can
// use the function [func@Gtk.show_about_dialog] which constructs and shows
// a dialog and keeps it around so that it can be shown again.
//
// Note that GTK sets a default title of `_("About %s")` on the dialog
// window (where `%s` is replaced by the name of the application, but in
// order to ensure proper translation of the title, applications should
// set the title property explicitly when constructing a `GtkAboutDialog`,
// as shown in the following example:
//
// ```c
// GFile *logo_file = g_file_new_for_path ("./logo.png");
// GdkTexture *example_logo = gdk_texture_new_from_file (logo_file, NULL);
// g_object_unref (logo_file);
//
// gtk_show_about_dialog (NULL,
//
//	"program-name", "ExampleCode",
//	"logo", example_logo,
//	"title", _("About ExampleCode"),
//	NULL);
//
// ```
//
// ## CSS nodes
//
// `GtkAboutDialog` has a single CSS node with the name `window` and style
// class `.aboutdialog`.
type AboutDialog struct {
	Window
}

func AboutDialogNewFromInternalPtr(ptr uintptr) *AboutDialog {
	cls := &AboutDialog{}
	cls.Ptr = ptr
	return cls
}

var xNewAboutDialog func() uintptr

// Creates a new `GtkAboutDialog`.
func NewAboutDialog() *Widget {
	var cls *Widget

	cret := xNewAboutDialog()

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &Widget{}
	cls.Ptr = cret
	return cls
}

var xAboutDialogAddCreditSection func(uintptr, string, uintptr)

// Creates a new section in the "Credits" page.
func (x *AboutDialog) AddCreditSection(SectionNameVar string, PeopleVar uintptr) {

	xAboutDialogAddCreditSection(x.GoPointer(), SectionNameVar, PeopleVar)

}

var xAboutDialogGetArtists func(uintptr) uintptr

// Returns the names of the artists which are displayed
// in the credits page.
func (x *AboutDialog) GetArtists() uintptr {

	cret := xAboutDialogGetArtists(x.GoPointer())
	return cret
}

var xAboutDialogGetAuthors func(uintptr) uintptr

// Returns the names of the authors which are displayed
// in the credits page.
func (x *AboutDialog) GetAuthors() uintptr {

	cret := xAboutDialogGetAuthors(x.GoPointer())
	return cret
}

var xAboutDialogGetComments func(uintptr) string

// Returns the comments string.
func (x *AboutDialog) GetComments() string {

	cret := xAboutDialogGetComments(x.GoPointer())
	return cret
}

var xAboutDialogGetCopyright func(uintptr) string

// Returns the copyright string.
func (x *AboutDialog) GetCopyright() string {

	cret := xAboutDialogGetCopyright(x.GoPointer())
	return cret
}

var xAboutDialogGetDocumenters func(uintptr) uintptr

// Returns the name of the documenters which are displayed
// in the credits page.
func (x *AboutDialog) GetDocumenters() uintptr {

	cret := xAboutDialogGetDocumenters(x.GoPointer())
	return cret
}

var xAboutDialogGetLicense func(uintptr) string

// Returns the license information.
func (x *AboutDialog) GetLicense() string {

	cret := xAboutDialogGetLicense(x.GoPointer())
	return cret
}

var xAboutDialogGetLicenseType func(uintptr) License

// Retrieves the license type.
func (x *AboutDialog) GetLicenseType() License {

	cret := xAboutDialogGetLicenseType(x.GoPointer())
	return cret
}

var xAboutDialogGetLogo func(uintptr) uintptr

// Returns the paintable displayed as logo in the about dialog.
func (x *AboutDialog) GetLogo() *gdk.PaintableBase {
	var cls *gdk.PaintableBase

	cret := xAboutDialogGetLogo(x.GoPointer())

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &gdk.PaintableBase{}
	cls.Ptr = cret
	return cls
}

var xAboutDialogGetLogoIconName func(uintptr) string

// Returns the icon name displayed as logo in the about dialog.
func (x *AboutDialog) GetLogoIconName() string {

	cret := xAboutDialogGetLogoIconName(x.GoPointer())
	return cret
}

var xAboutDialogGetProgramName func(uintptr) string

// Returns the program name displayed in the about dialog.
func (x *AboutDialog) GetProgramName() string {

	cret := xAboutDialogGetProgramName(x.GoPointer())
	return cret
}

var xAboutDialogGetSystemInformation func(uintptr) string

// Returns the system information that is shown in the about dialog.
func (x *AboutDialog) GetSystemInformation() string {

	cret := xAboutDialogGetSystemInformation(x.GoPointer())
	return cret
}

var xAboutDialogGetTranslatorCredits func(uintptr) string

// Returns the translator credits string which is displayed
// in the credits page.
func (x *AboutDialog) GetTranslatorCredits() string {

	cret := xAboutDialogGetTranslatorCredits(x.GoPointer())
	return cret
}

var xAboutDialogGetVersion func(uintptr) string

// Returns the version string.
func (x *AboutDialog) GetVersion() string {

	cret := xAboutDialogGetVersion(x.GoPointer())
	return cret
}

var xAboutDialogGetWebsite func(uintptr) string

// Returns the website URL.
func (x *AboutDialog) GetWebsite() string {

	cret := xAboutDialogGetWebsite(x.GoPointer())
	return cret
}

var xAboutDialogGetWebsiteLabel func(uintptr) string

// Returns the label used for the website link.
func (x *AboutDialog) GetWebsiteLabel() string {

	cret := xAboutDialogGetWebsiteLabel(x.GoPointer())
	return cret
}

var xAboutDialogGetWrapLicense func(uintptr) bool

// Returns whether the license text in the about dialog is
// automatically wrapped.
func (x *AboutDialog) GetWrapLicense() bool {

	cret := xAboutDialogGetWrapLicense(x.GoPointer())
	return cret
}

var xAboutDialogSetArtists func(uintptr, uintptr)

// Sets the names of the artists to be displayed
// in the "Credits" page.
func (x *AboutDialog) SetArtists(ArtistsVar uintptr) {

	xAboutDialogSetArtists(x.GoPointer(), ArtistsVar)

}

var xAboutDialogSetAuthors func(uintptr, uintptr)

// Sets the names of the authors which are displayed
// in the "Credits" page of the about dialog.
func (x *AboutDialog) SetAuthors(AuthorsVar uintptr) {

	xAboutDialogSetAuthors(x.GoPointer(), AuthorsVar)

}

var xAboutDialogSetComments func(uintptr, string)

// Sets the comments string to display in the about dialog.
//
// This should be a short string of one or two lines.
func (x *AboutDialog) SetComments(CommentsVar string) {

	xAboutDialogSetComments(x.GoPointer(), CommentsVar)

}

var xAboutDialogSetCopyright func(uintptr, string)

// Sets the copyright string to display in the about dialog.
//
// This should be a short string of one or two lines.
func (x *AboutDialog) SetCopyright(CopyrightVar string) {

	xAboutDialogSetCopyright(x.GoPointer(), CopyrightVar)

}

var xAboutDialogSetDocumenters func(uintptr, uintptr)

// Sets the names of the documenters which are displayed
// in the "Credits" page.
func (x *AboutDialog) SetDocumenters(DocumentersVar uintptr) {

	xAboutDialogSetDocumenters(x.GoPointer(), DocumentersVar)

}

var xAboutDialogSetLicense func(uintptr, string)

// Sets the license information to be displayed in the
// about dialog.
//
// If `license` is `NULL`, the license page is hidden.
func (x *AboutDialog) SetLicense(LicenseVar string) {

	xAboutDialogSetLicense(x.GoPointer(), LicenseVar)

}

var xAboutDialogSetLicenseType func(uintptr, License)

// Sets the license of the application showing the about dialog from a
// list of known licenses.
//
// This function overrides the license set using
// [method@Gtk.AboutDialog.set_license].
func (x *AboutDialog) SetLicenseType(LicenseTypeVar License) {

	xAboutDialogSetLicenseType(x.GoPointer(), LicenseTypeVar)

}

var xAboutDialogSetLogo func(uintptr, uintptr)

// Sets the logo in the about dialog.
func (x *AboutDialog) SetLogo(LogoVar gdk.Paintable) {

	xAboutDialogSetLogo(x.GoPointer(), LogoVar.GoPointer())

}

var xAboutDialogSetLogoIconName func(uintptr, string)

// Sets the icon name to be displayed as logo in the about dialog.
func (x *AboutDialog) SetLogoIconName(IconNameVar string) {

	xAboutDialogSetLogoIconName(x.GoPointer(), IconNameVar)

}

var xAboutDialogSetProgramName func(uintptr, string)

// Sets the name to display in the about dialog.
//
// If `name` is not set, the string returned
// by `g_get_application_name()` is used.
func (x *AboutDialog) SetProgramName(NameVar string) {

	xAboutDialogSetProgramName(x.GoPointer(), NameVar)

}

var xAboutDialogSetSystemInformation func(uintptr, string)

// Sets the system information to be displayed in the about
// dialog.
//
// If `system_information` is `NULL`, the system information
// page is hidden.
//
// See [property@Gtk.AboutDialog:system-information].
func (x *AboutDialog) SetSystemInformation(SystemInformationVar string) {

	xAboutDialogSetSystemInformation(x.GoPointer(), SystemInformationVar)

}

var xAboutDialogSetTranslatorCredits func(uintptr, string)

// Sets the translator credits string which is displayed in
// the credits page.
//
// The intended use for this string is to display the translator
// of the language which is currently used in the user interface.
// Using `gettext()`, a simple way to achieve that is to mark the
// string for translation:
//
// ```c
// GtkWidget *about = gtk_about_dialog_new ();
//
//	gtk_about_dialog_set_translator_credits (GTK_ABOUT_DIALOG (about),
//	                                         _("translator-credits"));
//
// ```
//
// It is a good idea to use the customary `msgid` “translator-credits”
// for this purpose, since translators will already know the purpose of
// that `msgid`, and since `GtkAboutDialog` will detect if “translator-credits”
// is untranslated and omit translator credits.
func (x *AboutDialog) SetTranslatorCredits(TranslatorCreditsVar string) {

	xAboutDialogSetTranslatorCredits(x.GoPointer(), TranslatorCreditsVar)

}

var xAboutDialogSetVersion func(uintptr, string)

// Sets the version string to display in the about dialog.
func (x *AboutDialog) SetVersion(VersionVar string) {

	xAboutDialogSetVersion(x.GoPointer(), VersionVar)

}

var xAboutDialogSetWebsite func(uintptr, string)

// Sets the URL to use for the website link.
func (x *AboutDialog) SetWebsite(WebsiteVar string) {

	xAboutDialogSetWebsite(x.GoPointer(), WebsiteVar)

}

var xAboutDialogSetWebsiteLabel func(uintptr, string)

// Sets the label to be used for the website link.
func (x *AboutDialog) SetWebsiteLabel(WebsiteLabelVar string) {

	xAboutDialogSetWebsiteLabel(x.GoPointer(), WebsiteLabelVar)

}

var xAboutDialogSetWrapLicense func(uintptr, bool)

// Sets whether the license text in the about dialog should be
// automatically wrapped.
func (x *AboutDialog) SetWrapLicense(WrapLicenseVar bool) {

	xAboutDialogSetWrapLicense(x.GoPointer(), WrapLicenseVar)

}

func (c *AboutDialog) GoPointer() uintptr {
	return c.Ptr
}

func (c *AboutDialog) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// Emitted every time a URL is activated.
//
// Applications may connect to it to override the default behaviour,
// which is to call [func@Gtk.show_uri].
func (x *AboutDialog) ConnectActivateLink(cb func(AboutDialog, string) bool) {
	fcb := func(clsPtr uintptr, UriVarp string) bool {
		fa := AboutDialog{}
		fa.Ptr = clsPtr

		return cb(fa, UriVarp)

	}
	gobject.ObjectConnect(x.GoPointer(), "signal::activate-link", purego.NewCallback(fcb))
}

// Retrieves the `GtkAccessibleRole` for the given `GtkAccessible`.
func (x *AboutDialog) GetAccessibleRole() AccessibleRole {

	cret := XGtkAccessibleGetAccessibleRole(x.GoPointer())
	return cret
}

// Resets the accessible @property to its default value.
func (x *AboutDialog) ResetProperty(PropertyVar AccessibleProperty) {

	XGtkAccessibleResetProperty(x.GoPointer(), PropertyVar)

}

// Resets the accessible @relation to its default value.
func (x *AboutDialog) ResetRelation(RelationVar AccessibleRelation) {

	XGtkAccessibleResetRelation(x.GoPointer(), RelationVar)

}

// Resets the accessible @state to its default value.
func (x *AboutDialog) ResetState(StateVar AccessibleState) {

	XGtkAccessibleResetState(x.GoPointer(), StateVar)

}

// Updates a list of accessible properties.
//
// See the [enum@Gtk.AccessibleProperty] documentation for the
// value types of accessible properties.
//
// This function should be called by `GtkWidget` types whenever
// an accessible property change must be communicated to assistive
// technologies.
//
// Example:
// ```c
// value = gtk_adjustment_get_value (adjustment);
// gtk_accessible_update_property (GTK_ACCESSIBLE (spin_button),
//
//	GTK_ACCESSIBLE_PROPERTY_VALUE_NOW, value,
//	-1);
//
// ```
func (x *AboutDialog) UpdateProperty(FirstPropertyVar AccessibleProperty, varArgs ...interface{}) {

	XGtkAccessibleUpdateProperty(x.GoPointer(), FirstPropertyVar, varArgs...)

}

// Updates an array of accessible properties.
//
// This function should be called by `GtkWidget` types whenever an accessible
// property change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *AboutDialog) UpdatePropertyValue(NPropertiesVar int, PropertiesVar uintptr, ValuesVar uintptr) {

	XGtkAccessibleUpdatePropertyValue(x.GoPointer(), NPropertiesVar, PropertiesVar, ValuesVar)

}

// Updates a list of accessible relations.
//
// This function should be called by `GtkWidget` types whenever an accessible
// relation change must be communicated to assistive technologies.
//
// If the [enum@Gtk.AccessibleRelation] requires a list of references,
// you should pass each reference individually, followed by %NULL, e.g.
//
// ```c
// gtk_accessible_update_relation (accessible,
//
//	GTK_ACCESSIBLE_RELATION_CONTROLS,
//	  ref1, NULL,
//	GTK_ACCESSIBLE_RELATION_LABELLED_BY,
//	  ref1, ref2, ref3, NULL,
//	-1);
//
// ```
func (x *AboutDialog) UpdateRelation(FirstRelationVar AccessibleRelation, varArgs ...interface{}) {

	XGtkAccessibleUpdateRelation(x.GoPointer(), FirstRelationVar, varArgs...)

}

// Updates an array of accessible relations.
//
// This function should be called by `GtkWidget` types whenever an accessible
// relation change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *AboutDialog) UpdateRelationValue(NRelationsVar int, RelationsVar uintptr, ValuesVar uintptr) {

	XGtkAccessibleUpdateRelationValue(x.GoPointer(), NRelationsVar, RelationsVar, ValuesVar)

}

// Updates a list of accessible states. See the [enum@Gtk.AccessibleState]
// documentation for the value types of accessible states.
//
// This function should be called by `GtkWidget` types whenever an accessible
// state change must be communicated to assistive technologies.
//
// Example:
// ```c
// value = GTK_ACCESSIBLE_TRISTATE_MIXED;
// gtk_accessible_update_state (GTK_ACCESSIBLE (check_button),
//
//	GTK_ACCESSIBLE_STATE_CHECKED, value,
//	-1);
//
// ```
func (x *AboutDialog) UpdateState(FirstStateVar AccessibleState, varArgs ...interface{}) {

	XGtkAccessibleUpdateState(x.GoPointer(), FirstStateVar, varArgs...)

}

// Updates an array of accessible states.
//
// This function should be called by `GtkWidget` types whenever an accessible
// state change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *AboutDialog) UpdateStateValue(NStatesVar int, StatesVar uintptr, ValuesVar uintptr) {

	XGtkAccessibleUpdateStateValue(x.GoPointer(), NStatesVar, StatesVar, ValuesVar)

}

// Gets the ID of the @buildable object.
//
// `GtkBuilder` sets the name based on the ID attribute
// of the &lt;object&gt; tag used to construct the @buildable.
func (x *AboutDialog) GetBuildableId() string {

	cret := XGtkBuildableGetBuildableId(x.GoPointer())
	return cret
}

// Returns the renderer that is used for this `GtkNative`.
func (x *AboutDialog) GetRenderer() *gsk.Renderer {
	var cls *gsk.Renderer

	cret := XGtkNativeGetRenderer(x.GoPointer())

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &gsk.Renderer{}
	cls.Ptr = cret
	return cls
}

// Returns the surface of this `GtkNative`.
func (x *AboutDialog) GetSurface() *gdk.Surface {
	var cls *gdk.Surface

	cret := XGtkNativeGetSurface(x.GoPointer())

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &gdk.Surface{}
	cls.Ptr = cret
	return cls
}

// Retrieves the surface transform of @self.
//
// This is the translation from @self's surface coordinates into
// @self's widget coordinates.
func (x *AboutDialog) GetSurfaceTransform(XVar float64, YVar float64) {

	XGtkNativeGetSurfaceTransform(x.GoPointer(), XVar, YVar)

}

// Realizes a `GtkNative`.
//
// This should only be used by subclasses.
func (x *AboutDialog) Realize() {

	XGtkNativeRealize(x.GoPointer())

}

// Unrealizes a `GtkNative`.
//
// This should only be used by subclasses.
func (x *AboutDialog) Unrealize() {

	XGtkNativeUnrealize(x.GoPointer())

}

// Returns the display that this `GtkRoot` is on.
func (x *AboutDialog) GetDisplay() *gdk.Display {
	var cls *gdk.Display

	cret := XGtkRootGetDisplay(x.GoPointer())

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &gdk.Display{}
	cls.Ptr = cret
	return cls
}

// Retrieves the current focused widget within the root.
//
// Note that this is the widget that would have the focus
// if the root is active; if the root is not focused then
// `gtk_widget_has_focus (widget)` will be %FALSE for the
// widget.
func (x *AboutDialog) GetFocus() *Widget {
	var cls *Widget

	cret := XGtkRootGetFocus(x.GoPointer())

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &Widget{}
	cls.Ptr = cret
	return cls
}

// If @focus is not the current focus widget, and is focusable, sets
// it as the focus widget for the root.
//
// If @focus is %NULL, unsets the focus widget for the root.
//
// To set the focus to a particular widget in the root, it is usually
// more convenient to use [method@Gtk.Widget.grab_focus] instead of
// this function.
func (x *AboutDialog) SetFocus(FocusVar *Widget) {

	XGtkRootSetFocus(x.GoPointer(), FocusVar.GoPointer())

}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GTK"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	core.PuregoSafeRegister(&xShowAboutDialog, lib, "gtk_show_about_dialog")

	core.PuregoSafeRegister(&xNewAboutDialog, lib, "gtk_about_dialog_new")

	core.PuregoSafeRegister(&xAboutDialogAddCreditSection, lib, "gtk_about_dialog_add_credit_section")
	core.PuregoSafeRegister(&xAboutDialogGetArtists, lib, "gtk_about_dialog_get_artists")
	core.PuregoSafeRegister(&xAboutDialogGetAuthors, lib, "gtk_about_dialog_get_authors")
	core.PuregoSafeRegister(&xAboutDialogGetComments, lib, "gtk_about_dialog_get_comments")
	core.PuregoSafeRegister(&xAboutDialogGetCopyright, lib, "gtk_about_dialog_get_copyright")
	core.PuregoSafeRegister(&xAboutDialogGetDocumenters, lib, "gtk_about_dialog_get_documenters")
	core.PuregoSafeRegister(&xAboutDialogGetLicense, lib, "gtk_about_dialog_get_license")
	core.PuregoSafeRegister(&xAboutDialogGetLicenseType, lib, "gtk_about_dialog_get_license_type")
	core.PuregoSafeRegister(&xAboutDialogGetLogo, lib, "gtk_about_dialog_get_logo")
	core.PuregoSafeRegister(&xAboutDialogGetLogoIconName, lib, "gtk_about_dialog_get_logo_icon_name")
	core.PuregoSafeRegister(&xAboutDialogGetProgramName, lib, "gtk_about_dialog_get_program_name")
	core.PuregoSafeRegister(&xAboutDialogGetSystemInformation, lib, "gtk_about_dialog_get_system_information")
	core.PuregoSafeRegister(&xAboutDialogGetTranslatorCredits, lib, "gtk_about_dialog_get_translator_credits")
	core.PuregoSafeRegister(&xAboutDialogGetVersion, lib, "gtk_about_dialog_get_version")
	core.PuregoSafeRegister(&xAboutDialogGetWebsite, lib, "gtk_about_dialog_get_website")
	core.PuregoSafeRegister(&xAboutDialogGetWebsiteLabel, lib, "gtk_about_dialog_get_website_label")
	core.PuregoSafeRegister(&xAboutDialogGetWrapLicense, lib, "gtk_about_dialog_get_wrap_license")
	core.PuregoSafeRegister(&xAboutDialogSetArtists, lib, "gtk_about_dialog_set_artists")
	core.PuregoSafeRegister(&xAboutDialogSetAuthors, lib, "gtk_about_dialog_set_authors")
	core.PuregoSafeRegister(&xAboutDialogSetComments, lib, "gtk_about_dialog_set_comments")
	core.PuregoSafeRegister(&xAboutDialogSetCopyright, lib, "gtk_about_dialog_set_copyright")
	core.PuregoSafeRegister(&xAboutDialogSetDocumenters, lib, "gtk_about_dialog_set_documenters")
	core.PuregoSafeRegister(&xAboutDialogSetLicense, lib, "gtk_about_dialog_set_license")
	core.PuregoSafeRegister(&xAboutDialogSetLicenseType, lib, "gtk_about_dialog_set_license_type")
	core.PuregoSafeRegister(&xAboutDialogSetLogo, lib, "gtk_about_dialog_set_logo")
	core.PuregoSafeRegister(&xAboutDialogSetLogoIconName, lib, "gtk_about_dialog_set_logo_icon_name")
	core.PuregoSafeRegister(&xAboutDialogSetProgramName, lib, "gtk_about_dialog_set_program_name")
	core.PuregoSafeRegister(&xAboutDialogSetSystemInformation, lib, "gtk_about_dialog_set_system_information")
	core.PuregoSafeRegister(&xAboutDialogSetTranslatorCredits, lib, "gtk_about_dialog_set_translator_credits")
	core.PuregoSafeRegister(&xAboutDialogSetVersion, lib, "gtk_about_dialog_set_version")
	core.PuregoSafeRegister(&xAboutDialogSetWebsite, lib, "gtk_about_dialog_set_website")
	core.PuregoSafeRegister(&xAboutDialogSetWebsiteLabel, lib, "gtk_about_dialog_set_website_label")
	core.PuregoSafeRegister(&xAboutDialogSetWrapLicense, lib, "gtk_about_dialog_set_wrap_license")

}
