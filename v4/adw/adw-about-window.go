// Package adw was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package adw

import (
	"unsafe"

	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/gdk"
	"github.com/jwijenbergh/puregotk/v4/gobject"
	"github.com/jwijenbergh/puregotk/v4/gsk"
	"github.com/jwijenbergh/puregotk/v4/gtk"
)

type AboutWindowClass struct {
	ParentClass uintptr
}

func (x *AboutWindowClass) GoPointer() uintptr {
	return uintptr(unsafe.Pointer(x))
}

var xShowAboutWindow func(uintptr, string, ...interface{})

// A convenience function for showing an application’s about window.
func ShowAboutWindow(ParentVar *gtk.Window, FirstPropertyNameVar string, varArgs ...interface{}) {

	xShowAboutWindow(ParentVar.GoPointer(), FirstPropertyNameVar, varArgs...)

}

// A window showing information about the application.
//
// &lt;picture&gt;
//
//	&lt;source srcset="about-window-dark.png" media="(prefers-color-scheme: dark)"&gt;
//	&lt;img src="about-window.png" alt="about-window"&gt;
//
// &lt;/picture&gt;
//
// An about window is typically opened when the user activates the `About …`
// item in the application's primary menu. All parts of the window are optional.
//
// ## Main page
//
// `AdwAboutWindow` prominently displays the application's icon, name, developer
// name and version. They can be set with the [property@AboutWindow:application-icon],
// [property@AboutWindow:application-name],
// [property@AboutWindow:developer-name] and [property@AboutWindow:version]
// respectively.
//
// ## What's New
//
// `AdwAboutWindow` provides a way for applications to display their release
// notes, set with the [property@AboutWindow:release-notes] property.
//
// Release notes are formatted the same way as
// [AppStream descriptions](https://freedesktop.org/software/appstream/docs/chap-Metadata.html#tag-description).
//
// The supported formatting options are:
//
// * Paragraph (`&lt;p&gt;`)
// * Ordered list (`&lt;ol&gt;`), with list items (`&lt;li&gt;`)
// * Unordered list (`&lt;ul&gt;`), with list items (`&lt;li&gt;`)
//
// Within paragraphs and list items, emphasis (`&lt;em&gt;`) and inline code
// (`&lt;code&gt;`) text styles are supported. The emphasis is rendered in italic,
// while inline code is shown in a monospaced font.
//
// Any text outside paragraphs or list items is ignored.
//
// Nested lists are not supported.
//
// Only one version can be shown at a time. By default, the displayed version
// number matches [property@AboutWindow:version]. Use
// [property@AboutWindow:release-notes-version] to override it.
//
// ## Details
//
// The Details page displays the application comments and links.
//
// The comments can be set with the [property@AboutWindow:comments] property.
// Unlike [property@Gtk.AboutDialog:comments], this string can be long and
// detailed. It can also contain links and Pango markup.
//
// To set the application website, use [property@AboutWindow:website].
// To add extra links below the website, use [method@AboutWindow.add_link].
//
// If the Details page doesn't have any other content besides website, the
// website will be displayed on the main page instead.
//
// ## Troubleshooting
//
// `AdwAboutWindow` displays the following two links on the main page:
//
// * Support Questions, set with the [property@AboutWindow:support-url] property,
// * Report an Issue, set with the [property@AboutWindow:issue-url] property.
//
// Additionally, applications can provide debugging information. It will be
// shown separately on the Troubleshooting page. Use the
// [property@AboutWindow:debug-info] property to specify it.
//
// It's intended to be attached to issue reports when reporting issues against
// the application. As such, it cannot contain markup or links.
//
// `AdwAboutWindow` provides a quick way to save debug information to a file.
// When saving, [property@AboutWindow:debug-info-filename] would be used as
// the suggested filename.
//
// ## Credits and Acknowledgements
//
// The Credits page has the following default sections:
//
// * Developers, set with the [property@AboutWindow:developers] property,
// * Designers, set with the [property@AboutWindow:designers] property,
// * Artists, set with the [property@AboutWindow:artists] property,
// * Documenters, set with the [property@AboutWindow:documenters] property,
// * Translators, set with the [property@AboutWindow:translator-credits] property.
//
// When setting translator credits, use the strings `"translator-credits"` or
// `"translator_credits"` and mark them as translatable.
//
// The default sections that don't contain any names won't be displayed.
//
// The Credits page can also contain an arbitrary number of extra sections below
// the default ones. Use [method@AboutWindow.add_credit_section] to add them.
//
// The Acknowledgements page can be used to acknowledge additional people and
// organizations for their non-development contributions. Use
// [method@AboutWindow.add_acknowledgement_section] to add sections to it. For
// example, it can be used to list backers in a crowdfunded project or to give
// special thanks.
//
// Each of the people or organizations can have an email address or a website
// specified. To add a email address, use a string like
// `Edgar Allan Poe &lt;edgar@poe.com&gt;`. To specify a website with a title, use a
// string like `The GNOME Project https://www.gnome.org`:
//
// &lt;picture&gt;
//
//	&lt;source srcset="about-window-credits-dark.png" media="(prefers-color-scheme: dark)"&gt;
//	&lt;img src="about-window-credits.png" alt="about-window-credits"&gt;
//
// &lt;/picture&gt;
//
// ## Legal
//
// The Legal page displays the copyright and licensing information for the
// application and other modules.
//
// The copyright string is set with the [property@AboutWindow:copyright]
// property and should be a short string of one or two lines, for example:
// `© 2022 Example`.
//
// Licensing information can be quickly set from a list of known licenses with
// the [property@AboutWindow:license-type] property. If the application's
// license is not in the list, [property@AboutWindow:license] can be used
// instead.
//
// To add information about other modules, such as application dependencies or
// data, use [method@AboutWindow.add_legal_section].
//
// ## Constructing
//
// To make constructing an `AdwAboutWindow` as convenient as possible, you can
// use the function [func@show_about_window] which constructs and shows a
// window.
//
// ```c
// static void
// show_about (GtkApplication *app)
//
//	{
//	  const char *developers[] = {
//	    "Angela Avery",
//	    NULL
//	  };
//
//	  const char *designers[] = {
//	    "GNOME Design Team",
//	    NULL
//	  };
//
//	  adw_show_about_window (gtk_application_get_active_window (app),
//	                         "application-name", _("Example"),
//	                         "application-icon", "org.example.App",
//	                         "version", "1.2.3",
//	                         "copyright", "© 2022 Angela Avery",
//	                         "issue-url", "https://gitlab.gnome.org/example/example/-/issues/new",
//	                         "license-type", GTK_LICENSE_GPL_3_0,
//	                         "developers", developers,
//	                         "designers", designers,
//	                         "translator-credits", _("translator-credits"),
//	                         NULL);
//	}
//
// ```
//
// ## CSS nodes
//
// `AdwAboutWindow` has a main CSS node with the name `window` and the
// style class `.about`.
type AboutWindow struct {
	Window
}

func AboutWindowNewFromInternalPtr(ptr uintptr) *AboutWindow {
	cls := &AboutWindow{}
	cls.Ptr = ptr
	return cls
}

var xNewAboutWindow func() uintptr

// Creates a new `AdwAboutWindow`.
func NewAboutWindow() *AboutWindow {
	var cls *AboutWindow

	cret := xNewAboutWindow()

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &AboutWindow{}
	cls.Ptr = cret
	return cls
}

var xAboutWindowAddAcknowledgementSection func(uintptr, string, []string)

// Adds a section to the Acknowledgements page.
//
// This can be used to acknowledge additional people and organizations for their
// non-development contributions - for example, backers in a crowdfunded
// project.
//
// Each name may contain email addresses and URLs, see the introduction for more
// details.
//
// See also:
//
// * [property@AboutWindow:developers]
// * [property@AboutWindow:designers]
// * [property@AboutWindow:artists]
// * [property@AboutWindow:documenters]
// * [property@AboutWindow:translator-credits]
// * [method@AboutWindow.add_credit_section]
func (x *AboutWindow) AddAcknowledgementSection(NameVar string, PeopleVar []string) {

	xAboutWindowAddAcknowledgementSection(x.GoPointer(), NameVar, PeopleVar)

}

var xAboutWindowAddCreditSection func(uintptr, string, []string)

// Adds an extra section to the Credits page.
//
// Extra sections are displayed below the standard categories.
//
// Each name may contain email addresses and URLs, see the introduction for more
// details.
//
// See also:
//
// * [property@AboutWindow:developers]
// * [property@AboutWindow:designers]
// * [property@AboutWindow:artists]
// * [property@AboutWindow:documenters]
// * [property@AboutWindow:translator-credits]
// * [method@AboutWindow.add_acknowledgement_section]
func (x *AboutWindow) AddCreditSection(NameVar string, PeopleVar []string) {

	xAboutWindowAddCreditSection(x.GoPointer(), NameVar, PeopleVar)

}

var xAboutWindowAddLegalSection func(uintptr, string, string, gtk.License, string)

// Adds an extra section to the Legal page.
//
// Extra sections will be displayed below the application's own information.
//
// The parameters @copyright, @license_type and @license will be used to present
// the it the same way as [property@AboutWindow:copyright],
// [property@AboutWindow:license-type] and [property@AboutWindow:license] are
// for the application's own information.
//
// See those properties for more details.
//
// This can be useful to attribute the application dependencies or data.
//
// Examples:
//
// ```c
// adw_about_window_add_legal_section (ADW_ABOUT_WINDOW (about),
//
//	_("Copyright and a known license"),
//	"© 2022 Example",
//	GTK_LICENSE_LGPL_2_1,
//	NULL);
//
// adw_about_window_add_legal_section (ADW_ABOUT_WINDOW (about),
//
//	_("Copyright and custom license"),
//	"© 2022 Example",
//	GTK_LICENSE_CUSTOM,
//	"Custom license text");
//
// adw_about_window_add_legal_section (ADW_ABOUT_WINDOW (about),
//
//	_("Copyright only"),
//	"© 2022 Example",
//	GTK_LICENSE_UNKNOWN,
//	NULL);
//
// adw_about_window_add_legal_section (ADW_ABOUT_WINDOW (about),
//
//	_("Custom license only"),
//	NULL,
//	GTK_LICENSE_CUSTOM,
//	"Something completely custom here.");
//
// ```
func (x *AboutWindow) AddLegalSection(TitleVar string, CopyrightVar string, LicenseTypeVar gtk.License, LicenseVar string) {

	xAboutWindowAddLegalSection(x.GoPointer(), TitleVar, CopyrightVar, LicenseTypeVar, LicenseVar)

}

var xAboutWindowAddLink func(uintptr, string, string)

// Adds an extra link to the Details page.
//
// Extra links are displayed under the comment and website.
//
// Underlines in @title will be interpreted as indicating a mnemonic.
//
// See [property@AboutWindow:website].
func (x *AboutWindow) AddLink(TitleVar string, UrlVar string) {

	xAboutWindowAddLink(x.GoPointer(), TitleVar, UrlVar)

}

var xAboutWindowGetApplicationIcon func(uintptr) string

// Gets the name of the application icon for @self.
func (x *AboutWindow) GetApplicationIcon() string {

	cret := xAboutWindowGetApplicationIcon(x.GoPointer())
	return cret
}

var xAboutWindowGetApplicationName func(uintptr) string

// Gets the application name for @self.
func (x *AboutWindow) GetApplicationName() string {

	cret := xAboutWindowGetApplicationName(x.GoPointer())
	return cret
}

var xAboutWindowGetArtists func(uintptr) uintptr

// Gets the list of artists of the application.
func (x *AboutWindow) GetArtists() uintptr {

	cret := xAboutWindowGetArtists(x.GoPointer())
	return cret
}

var xAboutWindowGetComments func(uintptr) string

// Gets the comments about the application.
func (x *AboutWindow) GetComments() string {

	cret := xAboutWindowGetComments(x.GoPointer())
	return cret
}

var xAboutWindowGetCopyright func(uintptr) string

// Gets the copyright information for @self.
func (x *AboutWindow) GetCopyright() string {

	cret := xAboutWindowGetCopyright(x.GoPointer())
	return cret
}

var xAboutWindowGetDebugInfo func(uintptr) string

// Gets the debug information for @self.
func (x *AboutWindow) GetDebugInfo() string {

	cret := xAboutWindowGetDebugInfo(x.GoPointer())
	return cret
}

var xAboutWindowGetDebugInfoFilename func(uintptr) string

// Gets the debug information filename for @self.
func (x *AboutWindow) GetDebugInfoFilename() string {

	cret := xAboutWindowGetDebugInfoFilename(x.GoPointer())
	return cret
}

var xAboutWindowGetDesigners func(uintptr) uintptr

// Gets the list of designers of the application.
func (x *AboutWindow) GetDesigners() uintptr {

	cret := xAboutWindowGetDesigners(x.GoPointer())
	return cret
}

var xAboutWindowGetDeveloperName func(uintptr) string

// Gets the developer name for @self.
func (x *AboutWindow) GetDeveloperName() string {

	cret := xAboutWindowGetDeveloperName(x.GoPointer())
	return cret
}

var xAboutWindowGetDevelopers func(uintptr) uintptr

// Gets the list of developers of the application.
func (x *AboutWindow) GetDevelopers() uintptr {

	cret := xAboutWindowGetDevelopers(x.GoPointer())
	return cret
}

var xAboutWindowGetDocumenters func(uintptr) uintptr

// Gets the list of documenters of the application.
func (x *AboutWindow) GetDocumenters() uintptr {

	cret := xAboutWindowGetDocumenters(x.GoPointer())
	return cret
}

var xAboutWindowGetIssueUrl func(uintptr) string

// Gets the issue tracker URL for @self.
func (x *AboutWindow) GetIssueUrl() string {

	cret := xAboutWindowGetIssueUrl(x.GoPointer())
	return cret
}

var xAboutWindowGetLicense func(uintptr) string

// Gets the license for @self.
func (x *AboutWindow) GetLicense() string {

	cret := xAboutWindowGetLicense(x.GoPointer())
	return cret
}

var xAboutWindowGetLicenseType func(uintptr) gtk.License

// Gets the license type for @self.
func (x *AboutWindow) GetLicenseType() gtk.License {

	cret := xAboutWindowGetLicenseType(x.GoPointer())
	return cret
}

var xAboutWindowGetReleaseNotes func(uintptr) string

// Gets the release notes for @self.
func (x *AboutWindow) GetReleaseNotes() string {

	cret := xAboutWindowGetReleaseNotes(x.GoPointer())
	return cret
}

var xAboutWindowGetReleaseNotesVersion func(uintptr) string

// Gets the version described by the application's release notes.
func (x *AboutWindow) GetReleaseNotesVersion() string {

	cret := xAboutWindowGetReleaseNotesVersion(x.GoPointer())
	return cret
}

var xAboutWindowGetSupportUrl func(uintptr) string

// Gets the URL of the support page for @self.
func (x *AboutWindow) GetSupportUrl() string {

	cret := xAboutWindowGetSupportUrl(x.GoPointer())
	return cret
}

var xAboutWindowGetTranslatorCredits func(uintptr) string

// Gets the translator credits string.
func (x *AboutWindow) GetTranslatorCredits() string {

	cret := xAboutWindowGetTranslatorCredits(x.GoPointer())
	return cret
}

var xAboutWindowGetVersion func(uintptr) string

// Gets the version for @self.
func (x *AboutWindow) GetVersion() string {

	cret := xAboutWindowGetVersion(x.GoPointer())
	return cret
}

var xAboutWindowGetWebsite func(uintptr) string

// Gets the application website URL for @self.
func (x *AboutWindow) GetWebsite() string {

	cret := xAboutWindowGetWebsite(x.GoPointer())
	return cret
}

var xAboutWindowSetApplicationIcon func(uintptr, string)

// Sets the name of the application icon for @self.
//
// The icon is displayed at the top of the main page.
func (x *AboutWindow) SetApplicationIcon(ApplicationIconVar string) {

	xAboutWindowSetApplicationIcon(x.GoPointer(), ApplicationIconVar)

}

var xAboutWindowSetApplicationName func(uintptr, string)

// Sets the application name for @self.
//
// The name is displayed at the top of the main page.
func (x *AboutWindow) SetApplicationName(ApplicationNameVar string) {

	xAboutWindowSetApplicationName(x.GoPointer(), ApplicationNameVar)

}

var xAboutWindowSetArtists func(uintptr, []string)

// Sets the list of artists of the application.
//
// It will be displayed on the Credits page.
//
// Each name may contain email addresses and URLs, see the introduction for more
// details.
//
// See also:
//
// * [property@AboutWindow:developers]
// * [property@AboutWindow:designers]
// * [property@AboutWindow:documenters]
// * [property@AboutWindow:translator-credits]
// * [method@AboutWindow.add_credit_section]
// * [method@AboutWindow.add_acknowledgement_section]
func (x *AboutWindow) SetArtists(ArtistsVar []string) {

	xAboutWindowSetArtists(x.GoPointer(), ArtistsVar)

}

var xAboutWindowSetComments func(uintptr, string)

// Sets the comments about the application.
//
// Comments will be shown on the Details page, above links.
//
// Unlike [property@Gtk.AboutDialog:comments], this string can be long and
// detailed. It can also contain links and Pango markup.
func (x *AboutWindow) SetComments(CommentsVar string) {

	xAboutWindowSetComments(x.GoPointer(), CommentsVar)

}

var xAboutWindowSetCopyright func(uintptr, string)

// Sets the copyright information for @self.
//
// This should be a short string of one or two lines, for example:
// `© 2022 Example`.
//
// The copyright information will be displayed on the Legal page, before the
// application license.
//
// [method@AboutWindow.add_legal_section] can be used to add copyright
// information for the application dependencies or other components.
func (x *AboutWindow) SetCopyright(CopyrightVar string) {

	xAboutWindowSetCopyright(x.GoPointer(), CopyrightVar)

}

var xAboutWindowSetDebugInfo func(uintptr, string)

// Sets the debug information for @self.
//
// Debug information will be shown on the Troubleshooting page. It's intended
// to be attached to issue reports when reporting issues against the
// application.
//
// `AdwAboutWindow` provides a quick way to save debug information to a file.
// When saving, [property@AboutWindow:debug-info-filename] would be used as
// the suggested filename.
//
// Debug information cannot contain markup or links.
func (x *AboutWindow) SetDebugInfo(DebugInfoVar string) {

	xAboutWindowSetDebugInfo(x.GoPointer(), DebugInfoVar)

}

var xAboutWindowSetDebugInfoFilename func(uintptr, string)

// Sets the debug information filename for @self.
//
// It will be used as the suggested filename when saving debug information to a
// file.
//
// See [property@AboutWindow:debug-info].
func (x *AboutWindow) SetDebugInfoFilename(FilenameVar string) {

	xAboutWindowSetDebugInfoFilename(x.GoPointer(), FilenameVar)

}

var xAboutWindowSetDesigners func(uintptr, []string)

// Sets the list of designers of the application.
//
// It will be displayed on the Credits page.
//
// Each name may contain email addresses and URLs, see the introduction for more
// details.
//
// See also:
//
// * [property@AboutWindow:developers]
// * [property@AboutWindow:artists]
// * [property@AboutWindow:documenters]
// * [property@AboutWindow:translator-credits]
// * [method@AboutWindow.add_credit_section]
// * [method@AboutWindow.add_acknowledgement_section]
func (x *AboutWindow) SetDesigners(DesignersVar []string) {

	xAboutWindowSetDesigners(x.GoPointer(), DesignersVar)

}

var xAboutWindowSetDeveloperName func(uintptr, string)

// Sets the developer name for @self.
//
// The developer name is displayed on the main page, under the application name.
//
// If the application is developed by multiple people, the developer name can be
// set to values like "AppName team", "AppName developers" or
// "The AppName project", and the individual contributors can be listed on the
// Credits page, with [property@AboutWindow:developers] and related properties.
func (x *AboutWindow) SetDeveloperName(DeveloperNameVar string) {

	xAboutWindowSetDeveloperName(x.GoPointer(), DeveloperNameVar)

}

var xAboutWindowSetDevelopers func(uintptr, []string)

// Sets the list of developers of the application.
//
// It will be displayed on the Credits page.
//
// Each name may contain email addresses and URLs, see the introduction for more
// details.
//
// See also:
//
// * [property@AboutWindow:designers]
// * [property@AboutWindow:artists]
// * [property@AboutWindow:documenters]
// * [property@AboutWindow:translator-credits]
// * [method@AboutWindow.add_credit_section]
// * [method@AboutWindow.add_acknowledgement_section]
func (x *AboutWindow) SetDevelopers(DevelopersVar []string) {

	xAboutWindowSetDevelopers(x.GoPointer(), DevelopersVar)

}

var xAboutWindowSetDocumenters func(uintptr, []string)

// Sets the list of documenters of the application.
//
// It will be displayed on the Credits page.
//
// Each name may contain email addresses and URLs, see the introduction for more
// details.
//
// See also:
//
// * [property@AboutWindow:developers]
// * [property@AboutWindow:designers]
// * [property@AboutWindow:artists]
// * [property@AboutWindow:translator-credits]
// * [method@AboutWindow.add_credit_section]
// * [method@AboutWindow.add_acknowledgement_section]
func (x *AboutWindow) SetDocumenters(DocumentersVar []string) {

	xAboutWindowSetDocumenters(x.GoPointer(), DocumentersVar)

}

var xAboutWindowSetIssueUrl func(uintptr, string)

// Sets the issue tracker URL for @self.
//
// The issue tracker link is displayed on the main page.
func (x *AboutWindow) SetIssueUrl(IssueUrlVar string) {

	xAboutWindowSetIssueUrl(x.GoPointer(), IssueUrlVar)

}

var xAboutWindowSetLicense func(uintptr, string)

// Sets the license for @self.
//
// This can be used to set a custom text for the license if it can't be set via
// [property@AboutWindow:license-type].
//
// When set, [property@AboutWindow:license-type] will be set to
// `GTK_LICENSE_CUSTOM`.
//
// The license text will be displayed on the Legal page, below the copyright
// information.
//
// License text can contain Pango markup and links.
//
// [method@AboutWindow.add_legal_section] can be used to add license information
// for the application dependencies or other components.
func (x *AboutWindow) SetLicense(LicenseVar string) {

	xAboutWindowSetLicense(x.GoPointer(), LicenseVar)

}

var xAboutWindowSetLicenseType func(uintptr, gtk.License)

// Sets the license for @self from a list of known licenses.
//
// If the application's license is not in the list,
// [property@AboutWindow:license] can be used instead. The license type will be
// automatically set to `GTK_LICENSE_CUSTOM` in that case.
//
// If @license_type is `GTK_LICENSE_UNKNOWN`, no information will be displayed.
//
// If @license_type is different from `GTK_LICENSE_CUSTOM`.
// [property@AboutWindow:license] will be cleared out.
//
// The license description will be displayed on the Legal page, below the
// copyright information.
//
// [method@AboutWindow.add_legal_section] can be used to add license information
// for the application dependencies or other components.
func (x *AboutWindow) SetLicenseType(LicenseTypeVar gtk.License) {

	xAboutWindowSetLicenseType(x.GoPointer(), LicenseTypeVar)

}

var xAboutWindowSetReleaseNotes func(uintptr, string)

// Sets the release notes for @self.
//
// Release notes are displayed on the the What's New page.
//
// Release notes are formatted the same way as
// [AppStream descriptions](https://freedesktop.org/software/appstream/docs/chap-Metadata.html#tag-description).
//
// The supported formatting options are:
//
// * Paragraph (`&lt;p&gt;`)
// * Ordered list (`&lt;ol&gt;`), with list items (`&lt;li&gt;`)
// * Unordered list (`&lt;ul&gt;`), with list items (`&lt;li&gt;`)
//
// Within paragraphs and list items, emphasis (`&lt;em&gt;`) and inline code
// (`&lt;code&gt;`) text styles are supported. The emphasis is rendered in italic,
// while inline code is shown in a monospaced font.
//
// Any text outside paragraphs or list items is ignored.
//
// Nested lists are not supported.
//
// `AdwAboutWindow` displays the version above the release notes. If set, the
// [property@AboutWindow:release-notes-version] of the property will be used
// as the version; otherwise, [property@AboutWindow:version] is used.
func (x *AboutWindow) SetReleaseNotes(ReleaseNotesVar string) {

	xAboutWindowSetReleaseNotes(x.GoPointer(), ReleaseNotesVar)

}

var xAboutWindowSetReleaseNotesVersion func(uintptr, string)

// Sets the version described by the application's release notes.
//
// The release notes version is displayed on the What's New page, above the
// release notes.
//
// If not set, [property@AboutWindow:version] will be used instead.
//
// For example, an application with the current version 2.0.2 might want to
// keep the release notes from 2.0.0, and set the release notes version
// accordingly.
//
// See [property@AboutWindow:release-notes].
func (x *AboutWindow) SetReleaseNotesVersion(VersionVar string) {

	xAboutWindowSetReleaseNotesVersion(x.GoPointer(), VersionVar)

}

var xAboutWindowSetSupportUrl func(uintptr, string)

// Sets the URL of the support page for @self.
//
// The support page link is displayed on the main page.
func (x *AboutWindow) SetSupportUrl(SupportUrlVar string) {

	xAboutWindowSetSupportUrl(x.GoPointer(), SupportUrlVar)

}

var xAboutWindowSetTranslatorCredits func(uintptr, string)

// Sets the translator credits string.
//
// It will be displayed on the Credits page.
//
// This string should be `"translator-credits"` or `"translator_credits"` and
// should be marked as translatable.
//
// The string may contain email addresses and URLs, see the introduction for
// more details.
//
// See also:
//
// * [property@AboutWindow:developers]
// * [property@AboutWindow:designers]
// * [property@AboutWindow:artists]
// * [property@AboutWindow:documenters]
// * [method@AboutWindow.add_credit_section]
// * [method@AboutWindow.add_acknowledgement_section]
func (x *AboutWindow) SetTranslatorCredits(TranslatorCreditsVar string) {

	xAboutWindowSetTranslatorCredits(x.GoPointer(), TranslatorCreditsVar)

}

var xAboutWindowSetVersion func(uintptr, string)

// Sets the version for @self.
//
// The version is displayed on the main page.
//
// If [property@AboutWindow:release-notes-version] is not set, the version will
// also be displayed above the release notes on the What's New page.
func (x *AboutWindow) SetVersion(VersionVar string) {

	xAboutWindowSetVersion(x.GoPointer(), VersionVar)

}

var xAboutWindowSetWebsite func(uintptr, string)

// Sets the application website URL for @self.
//
// Website is displayed on the Details page, below comments, or on the main page
// if the Details page doesn't have any other content.
//
// Applications can add other links below, see [method@AboutWindow.add_link].
func (x *AboutWindow) SetWebsite(WebsiteVar string) {

	xAboutWindowSetWebsite(x.GoPointer(), WebsiteVar)

}

func (c *AboutWindow) GoPointer() uintptr {
	return c.Ptr
}

func (c *AboutWindow) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// Emitted when a URL is activated.
//
// Applications may connect to it to override the default behavior, which is
// to call [func@Gtk.show_uri].
func (x *AboutWindow) ConnectActivateLink(cb func(AboutWindow, string) bool) uint32 {
	fcb := func(clsPtr uintptr, UriVarp string) bool {
		fa := AboutWindow{}
		fa.Ptr = clsPtr

		return cb(fa, UriVarp)

	}
	return gobject.SignalConnect(x.GoPointer(), "activate-link", purego.NewCallback(fcb))
}

// Retrieves the `GtkAccessibleRole` for the given `GtkAccessible`.
func (x *AboutWindow) GetAccessibleRole() gtk.AccessibleRole {

	cret := gtk.XGtkAccessibleGetAccessibleRole(x.GoPointer())
	return cret
}

// Resets the accessible @property to its default value.
func (x *AboutWindow) ResetProperty(PropertyVar gtk.AccessibleProperty) {

	gtk.XGtkAccessibleResetProperty(x.GoPointer(), PropertyVar)

}

// Resets the accessible @relation to its default value.
func (x *AboutWindow) ResetRelation(RelationVar gtk.AccessibleRelation) {

	gtk.XGtkAccessibleResetRelation(x.GoPointer(), RelationVar)

}

// Resets the accessible @state to its default value.
func (x *AboutWindow) ResetState(StateVar gtk.AccessibleState) {

	gtk.XGtkAccessibleResetState(x.GoPointer(), StateVar)

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
func (x *AboutWindow) UpdateProperty(FirstPropertyVar gtk.AccessibleProperty, varArgs ...interface{}) {

	gtk.XGtkAccessibleUpdateProperty(x.GoPointer(), FirstPropertyVar, varArgs...)

}

// Updates an array of accessible properties.
//
// This function should be called by `GtkWidget` types whenever an accessible
// property change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *AboutWindow) UpdatePropertyValue(NPropertiesVar int, PropertiesVar uintptr, ValuesVar uintptr) {

	gtk.XGtkAccessibleUpdatePropertyValue(x.GoPointer(), NPropertiesVar, PropertiesVar, ValuesVar)

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
func (x *AboutWindow) UpdateRelation(FirstRelationVar gtk.AccessibleRelation, varArgs ...interface{}) {

	gtk.XGtkAccessibleUpdateRelation(x.GoPointer(), FirstRelationVar, varArgs...)

}

// Updates an array of accessible relations.
//
// This function should be called by `GtkWidget` types whenever an accessible
// relation change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *AboutWindow) UpdateRelationValue(NRelationsVar int, RelationsVar uintptr, ValuesVar uintptr) {

	gtk.XGtkAccessibleUpdateRelationValue(x.GoPointer(), NRelationsVar, RelationsVar, ValuesVar)

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
func (x *AboutWindow) UpdateState(FirstStateVar gtk.AccessibleState, varArgs ...interface{}) {

	gtk.XGtkAccessibleUpdateState(x.GoPointer(), FirstStateVar, varArgs...)

}

// Updates an array of accessible states.
//
// This function should be called by `GtkWidget` types whenever an accessible
// state change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *AboutWindow) UpdateStateValue(NStatesVar int, StatesVar uintptr, ValuesVar uintptr) {

	gtk.XGtkAccessibleUpdateStateValue(x.GoPointer(), NStatesVar, StatesVar, ValuesVar)

}

// Gets the ID of the @buildable object.
//
// `GtkBuilder` sets the name based on the ID attribute
// of the &lt;object&gt; tag used to construct the @buildable.
func (x *AboutWindow) GetBuildableId() string {

	cret := gtk.XGtkBuildableGetBuildableId(x.GoPointer())
	return cret
}

// Returns the renderer that is used for this `GtkNative`.
func (x *AboutWindow) GetRenderer() *gsk.Renderer {
	var cls *gsk.Renderer

	cret := gtk.XGtkNativeGetRenderer(x.GoPointer())

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &gsk.Renderer{}
	cls.Ptr = cret
	return cls
}

// Returns the surface of this `GtkNative`.
func (x *AboutWindow) GetSurface() *gdk.Surface {
	var cls *gdk.Surface

	cret := gtk.XGtkNativeGetSurface(x.GoPointer())

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
func (x *AboutWindow) GetSurfaceTransform(XVar float64, YVar float64) {

	gtk.XGtkNativeGetSurfaceTransform(x.GoPointer(), XVar, YVar)

}

// Realizes a `GtkNative`.
//
// This should only be used by subclasses.
func (x *AboutWindow) Realize() {

	gtk.XGtkNativeRealize(x.GoPointer())

}

// Unrealizes a `GtkNative`.
//
// This should only be used by subclasses.
func (x *AboutWindow) Unrealize() {

	gtk.XGtkNativeUnrealize(x.GoPointer())

}

// Returns the display that this `GtkRoot` is on.
func (x *AboutWindow) GetDisplay() *gdk.Display {
	var cls *gdk.Display

	cret := gtk.XGtkRootGetDisplay(x.GoPointer())

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
func (x *AboutWindow) GetFocus() *gtk.Widget {
	var cls *gtk.Widget

	cret := gtk.XGtkRootGetFocus(x.GoPointer())

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &gtk.Widget{}
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
func (x *AboutWindow) SetFocus(FocusVar *gtk.Widget) {

	gtk.XGtkRootSetFocus(x.GoPointer(), FocusVar.GoPointer())

}

func init() {
	lib, err := purego.Dlopen(core.GetPath("ADW"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	core.PuregoSafeRegister(&xShowAboutWindow, lib, "adw_show_about_window")

	core.PuregoSafeRegister(&xNewAboutWindow, lib, "adw_about_window_new")

	core.PuregoSafeRegister(&xAboutWindowAddAcknowledgementSection, lib, "adw_about_window_add_acknowledgement_section")
	core.PuregoSafeRegister(&xAboutWindowAddCreditSection, lib, "adw_about_window_add_credit_section")
	core.PuregoSafeRegister(&xAboutWindowAddLegalSection, lib, "adw_about_window_add_legal_section")
	core.PuregoSafeRegister(&xAboutWindowAddLink, lib, "adw_about_window_add_link")
	core.PuregoSafeRegister(&xAboutWindowGetApplicationIcon, lib, "adw_about_window_get_application_icon")
	core.PuregoSafeRegister(&xAboutWindowGetApplicationName, lib, "adw_about_window_get_application_name")
	core.PuregoSafeRegister(&xAboutWindowGetArtists, lib, "adw_about_window_get_artists")
	core.PuregoSafeRegister(&xAboutWindowGetComments, lib, "adw_about_window_get_comments")
	core.PuregoSafeRegister(&xAboutWindowGetCopyright, lib, "adw_about_window_get_copyright")
	core.PuregoSafeRegister(&xAboutWindowGetDebugInfo, lib, "adw_about_window_get_debug_info")
	core.PuregoSafeRegister(&xAboutWindowGetDebugInfoFilename, lib, "adw_about_window_get_debug_info_filename")
	core.PuregoSafeRegister(&xAboutWindowGetDesigners, lib, "adw_about_window_get_designers")
	core.PuregoSafeRegister(&xAboutWindowGetDeveloperName, lib, "adw_about_window_get_developer_name")
	core.PuregoSafeRegister(&xAboutWindowGetDevelopers, lib, "adw_about_window_get_developers")
	core.PuregoSafeRegister(&xAboutWindowGetDocumenters, lib, "adw_about_window_get_documenters")
	core.PuregoSafeRegister(&xAboutWindowGetIssueUrl, lib, "adw_about_window_get_issue_url")
	core.PuregoSafeRegister(&xAboutWindowGetLicense, lib, "adw_about_window_get_license")
	core.PuregoSafeRegister(&xAboutWindowGetLicenseType, lib, "adw_about_window_get_license_type")
	core.PuregoSafeRegister(&xAboutWindowGetReleaseNotes, lib, "adw_about_window_get_release_notes")
	core.PuregoSafeRegister(&xAboutWindowGetReleaseNotesVersion, lib, "adw_about_window_get_release_notes_version")
	core.PuregoSafeRegister(&xAboutWindowGetSupportUrl, lib, "adw_about_window_get_support_url")
	core.PuregoSafeRegister(&xAboutWindowGetTranslatorCredits, lib, "adw_about_window_get_translator_credits")
	core.PuregoSafeRegister(&xAboutWindowGetVersion, lib, "adw_about_window_get_version")
	core.PuregoSafeRegister(&xAboutWindowGetWebsite, lib, "adw_about_window_get_website")
	core.PuregoSafeRegister(&xAboutWindowSetApplicationIcon, lib, "adw_about_window_set_application_icon")
	core.PuregoSafeRegister(&xAboutWindowSetApplicationName, lib, "adw_about_window_set_application_name")
	core.PuregoSafeRegister(&xAboutWindowSetArtists, lib, "adw_about_window_set_artists")
	core.PuregoSafeRegister(&xAboutWindowSetComments, lib, "adw_about_window_set_comments")
	core.PuregoSafeRegister(&xAboutWindowSetCopyright, lib, "adw_about_window_set_copyright")
	core.PuregoSafeRegister(&xAboutWindowSetDebugInfo, lib, "adw_about_window_set_debug_info")
	core.PuregoSafeRegister(&xAboutWindowSetDebugInfoFilename, lib, "adw_about_window_set_debug_info_filename")
	core.PuregoSafeRegister(&xAboutWindowSetDesigners, lib, "adw_about_window_set_designers")
	core.PuregoSafeRegister(&xAboutWindowSetDeveloperName, lib, "adw_about_window_set_developer_name")
	core.PuregoSafeRegister(&xAboutWindowSetDevelopers, lib, "adw_about_window_set_developers")
	core.PuregoSafeRegister(&xAboutWindowSetDocumenters, lib, "adw_about_window_set_documenters")
	core.PuregoSafeRegister(&xAboutWindowSetIssueUrl, lib, "adw_about_window_set_issue_url")
	core.PuregoSafeRegister(&xAboutWindowSetLicense, lib, "adw_about_window_set_license")
	core.PuregoSafeRegister(&xAboutWindowSetLicenseType, lib, "adw_about_window_set_license_type")
	core.PuregoSafeRegister(&xAboutWindowSetReleaseNotes, lib, "adw_about_window_set_release_notes")
	core.PuregoSafeRegister(&xAboutWindowSetReleaseNotesVersion, lib, "adw_about_window_set_release_notes_version")
	core.PuregoSafeRegister(&xAboutWindowSetSupportUrl, lib, "adw_about_window_set_support_url")
	core.PuregoSafeRegister(&xAboutWindowSetTranslatorCredits, lib, "adw_about_window_set_translator_credits")
	core.PuregoSafeRegister(&xAboutWindowSetVersion, lib, "adw_about_window_set_version")
	core.PuregoSafeRegister(&xAboutWindowSetWebsite, lib, "adw_about_window_set_website")

}
