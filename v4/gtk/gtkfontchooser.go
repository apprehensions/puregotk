// Package gtk was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gtk

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/glib"
	"github.com/jwijenbergh/puregotk/v4/gobject"
	"github.com/jwijenbergh/puregotk/v4/pango"
)

// The type of function that is used for deciding what fonts get
// shown in a `GtkFontChooser`.
//
// See [method@Gtk.FontChooser.set_filter_func].
type FontFilterFunc func(uintptr, uintptr, uintptr) bool

type FontChooserIface struct {
	BaseIface uintptr

	Padding uintptr
}

// `GtkFontChooser` is an interface that can be implemented by widgets
// for choosing fonts.
//
// In GTK, the main objects that implement this interface are
// [class@Gtk.FontChooserWidget], [class@Gtk.FontChooserDialog] and
// [class@Gtk.FontButton].
type FontChooser interface {
	GoPointer() uintptr
	SetGoPointer(uintptr)
	GetFont() string
	GetFontDesc() *pango.FontDescription
	GetFontFace() *pango.FontFace
	GetFontFamily() *pango.FontFamily
	GetFontFeatures() string
	GetFontMap() *pango.FontMap
	GetFontSize() int
	GetLanguage() string
	GetLevel() FontChooserLevel
	GetPreviewText() string
	GetShowPreviewEntry() bool
	SetFilterFunc(FilterVar FontFilterFunc, UserDataVar uintptr, DestroyVar glib.DestroyNotify)
	SetFont(FontnameVar string)
	SetFontDesc(FontDescVar *pango.FontDescription)
	SetFontMap(FontmapVar *pango.FontMap)
	SetLanguage(LanguageVar string)
	SetLevel(LevelVar FontChooserLevel)
	SetPreviewText(TextVar string)
	SetShowPreviewEntry(ShowPreviewEntryVar bool)
}
type FontChooserBase struct {
	Ptr uintptr
}

func (x *FontChooserBase) GoPointer() uintptr {
	return x.Ptr
}

func (x *FontChooserBase) SetGoPointer(ptr uintptr) {
	x.Ptr = ptr
}

// Gets the currently-selected font name.
//
// Note that this can be a different string than what you set with
// [method@Gtk.FontChooser.set_font], as the font chooser widget may
// normalize font names and thus return a string with a different
// structure. For example, “Helvetica Italic Bold 12” could be
// normalized to “Helvetica Bold Italic 12”.
//
// Use [method@Pango.FontDescription.equal] if you want to compare two
// font descriptions.
func (x *FontChooserBase) GetFont() string {

	cret := XGtkFontChooserGetFont(x.GoPointer())
	return cret
}

// Gets the currently-selected font.
//
// Note that this can be a different string than what you set with
// [method@Gtk.FontChooser.set_font], as the font chooser widget may
// normalize font names and thus return a string with a different
// structure. For example, “Helvetica Italic Bold 12” could be
// normalized to “Helvetica Bold Italic 12”.
//
// Use [method@Pango.FontDescription.equal] if you want to compare two
// font descriptions.
func (x *FontChooserBase) GetFontDesc() *pango.FontDescription {

	cret := XGtkFontChooserGetFontDesc(x.GoPointer())
	return cret
}

// Gets the `PangoFontFace` representing the selected font group
// details (i.e. family, slant, weight, width, etc).
//
// If the selected font is not installed, returns %NULL.
func (x *FontChooserBase) GetFontFace() *pango.FontFace {
	var cls *pango.FontFace

	cret := XGtkFontChooserGetFontFace(x.GoPointer())

	if cret == 0 {
		return cls
	}
	gobject.IncreaseRef(cret)
	cls = &pango.FontFace{}
	cls.Ptr = cret
	return cls
}

// Gets the `PangoFontFamily` representing the selected font family.
//
// Font families are a collection of font faces.
//
// If the selected font is not installed, returns %NULL.
func (x *FontChooserBase) GetFontFamily() *pango.FontFamily {
	var cls *pango.FontFamily

	cret := XGtkFontChooserGetFontFamily(x.GoPointer())

	if cret == 0 {
		return cls
	}
	gobject.IncreaseRef(cret)
	cls = &pango.FontFamily{}
	cls.Ptr = cret
	return cls
}

// Gets the currently-selected font features.
//
// The format of the returned string is compatible with the
// [CSS font-feature-settings property](https://www.w3.org/TR/css-fonts-4/#font-rend-desc).
// It can be passed to [func@Pango.AttrFontFeatures.new].
func (x *FontChooserBase) GetFontFeatures() string {

	cret := XGtkFontChooserGetFontFeatures(x.GoPointer())
	return cret
}

// Gets the custom font map of this font chooser widget,
// or %NULL if it does not have one.
func (x *FontChooserBase) GetFontMap() *pango.FontMap {
	var cls *pango.FontMap

	cret := XGtkFontChooserGetFontMap(x.GoPointer())

	if cret == 0 {
		return cls
	}
	cls = &pango.FontMap{}
	cls.Ptr = cret
	return cls
}

// The selected font size.
func (x *FontChooserBase) GetFontSize() int {

	cret := XGtkFontChooserGetFontSize(x.GoPointer())
	return cret
}

// Gets the language that is used for font features.
func (x *FontChooserBase) GetLanguage() string {

	cret := XGtkFontChooserGetLanguage(x.GoPointer())
	return cret
}

// Returns the current level of granularity for selecting fonts.
func (x *FontChooserBase) GetLevel() FontChooserLevel {

	cret := XGtkFontChooserGetLevel(x.GoPointer())
	return cret
}

// Gets the text displayed in the preview area.
func (x *FontChooserBase) GetPreviewText() string {

	cret := XGtkFontChooserGetPreviewText(x.GoPointer())
	return cret
}

// Returns whether the preview entry is shown or not.
func (x *FontChooserBase) GetShowPreviewEntry() bool {

	cret := XGtkFontChooserGetShowPreviewEntry(x.GoPointer())
	return cret
}

// Adds a filter function that decides which fonts to display
// in the font chooser.
func (x *FontChooserBase) SetFilterFunc(FilterVar FontFilterFunc, UserDataVar uintptr, DestroyVar glib.DestroyNotify) {

	XGtkFontChooserSetFilterFunc(x.GoPointer(), purego.NewCallback(FilterVar), UserDataVar, purego.NewCallback(DestroyVar))

}

// Sets the currently-selected font.
func (x *FontChooserBase) SetFont(FontnameVar string) {

	XGtkFontChooserSetFont(x.GoPointer(), FontnameVar)

}

// Sets the currently-selected font from @font_desc.
func (x *FontChooserBase) SetFontDesc(FontDescVar *pango.FontDescription) {

	XGtkFontChooserSetFontDesc(x.GoPointer(), FontDescVar)

}

// Sets a custom font map to use for this font chooser widget.
//
// A custom font map can be used to present application-specific
// fonts instead of or in addition to the normal system fonts.
//
// ```c
// FcConfig *config;
// PangoFontMap *fontmap;
//
// config = FcInitLoadConfigAndFonts ();
// FcConfigAppFontAddFile (config, my_app_font_file);
//
// fontmap = pango_cairo_font_map_new_for_font_type (CAIRO_FONT_TYPE_FT);
// pango_fc_font_map_set_config (PANGO_FC_FONT_MAP (fontmap), config);
//
// gtk_font_chooser_set_font_map (font_chooser, fontmap);
// ```
//
// Note that other GTK widgets will only be able to use the
// application-specific font if it is present in the font map they use:
//
// ```c
// context = gtk_widget_get_pango_context (label);
// pango_context_set_font_map (context, fontmap);
// ```
func (x *FontChooserBase) SetFontMap(FontmapVar *pango.FontMap) {

	XGtkFontChooserSetFontMap(x.GoPointer(), FontmapVar.GoPointer())

}

// Sets the language to use for font features.
func (x *FontChooserBase) SetLanguage(LanguageVar string) {

	XGtkFontChooserSetLanguage(x.GoPointer(), LanguageVar)

}

// Sets the desired level of granularity for selecting fonts.
func (x *FontChooserBase) SetLevel(LevelVar FontChooserLevel) {

	XGtkFontChooserSetLevel(x.GoPointer(), LevelVar)

}

// Sets the text displayed in the preview area.
//
// The @text is used to show how the selected font looks.
func (x *FontChooserBase) SetPreviewText(TextVar string) {

	XGtkFontChooserSetPreviewText(x.GoPointer(), TextVar)

}

// Shows or hides the editable preview entry.
func (x *FontChooserBase) SetShowPreviewEntry(ShowPreviewEntryVar bool) {

	XGtkFontChooserSetShowPreviewEntry(x.GoPointer(), ShowPreviewEntryVar)

}

var XGtkFontChooserGetFont func(uintptr) string
var XGtkFontChooserGetFontDesc func(uintptr) *pango.FontDescription
var XGtkFontChooserGetFontFace func(uintptr) uintptr
var XGtkFontChooserGetFontFamily func(uintptr) uintptr
var XGtkFontChooserGetFontFeatures func(uintptr) string
var XGtkFontChooserGetFontMap func(uintptr) uintptr
var XGtkFontChooserGetFontSize func(uintptr) int
var XGtkFontChooserGetLanguage func(uintptr) string
var XGtkFontChooserGetLevel func(uintptr) FontChooserLevel
var XGtkFontChooserGetPreviewText func(uintptr) string
var XGtkFontChooserGetShowPreviewEntry func(uintptr) bool
var XGtkFontChooserSetFilterFunc func(uintptr, uintptr, uintptr, uintptr)
var XGtkFontChooserSetFont func(uintptr, string)
var XGtkFontChooserSetFontDesc func(uintptr, *pango.FontDescription)
var XGtkFontChooserSetFontMap func(uintptr, uintptr)
var XGtkFontChooserSetLanguage func(uintptr, string)
var XGtkFontChooserSetLevel func(uintptr, FontChooserLevel)
var XGtkFontChooserSetPreviewText func(uintptr, string)
var XGtkFontChooserSetShowPreviewEntry func(uintptr, bool)

// Specifies the granularity of font selection
// that is desired in a `GtkFontChooser`.
//
// This enumeration may be extended in the future; applications should
// ignore unknown values.
type FontChooserLevel int

const (

	// Allow selecting a font family
	FontChooserLevelFamilyValue FontChooserLevel = 0
	// Allow selecting a specific font face
	FontChooserLevelStyleValue FontChooserLevel = 1
	// Allow selecting a specific font size
	FontChooserLevelSizeValue FontChooserLevel = 2
	// Allow changing OpenType font variation axes
	FontChooserLevelVariationsValue FontChooserLevel = 4
	// Allow selecting specific OpenType font features
	FontChooserLevelFeaturesValue FontChooserLevel = 8
)

func init() {
	lib, err := purego.Dlopen(core.GetPath("GTK"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&XGtkFontChooserGetFont, lib, "gtk_font_chooser_get_font")
	core.PuregoSafeRegister(&XGtkFontChooserGetFontDesc, lib, "gtk_font_chooser_get_font_desc")
	core.PuregoSafeRegister(&XGtkFontChooserGetFontFace, lib, "gtk_font_chooser_get_font_face")
	core.PuregoSafeRegister(&XGtkFontChooserGetFontFamily, lib, "gtk_font_chooser_get_font_family")
	core.PuregoSafeRegister(&XGtkFontChooserGetFontFeatures, lib, "gtk_font_chooser_get_font_features")
	core.PuregoSafeRegister(&XGtkFontChooserGetFontMap, lib, "gtk_font_chooser_get_font_map")
	core.PuregoSafeRegister(&XGtkFontChooserGetFontSize, lib, "gtk_font_chooser_get_font_size")
	core.PuregoSafeRegister(&XGtkFontChooserGetLanguage, lib, "gtk_font_chooser_get_language")
	core.PuregoSafeRegister(&XGtkFontChooserGetLevel, lib, "gtk_font_chooser_get_level")
	core.PuregoSafeRegister(&XGtkFontChooserGetPreviewText, lib, "gtk_font_chooser_get_preview_text")
	core.PuregoSafeRegister(&XGtkFontChooserGetShowPreviewEntry, lib, "gtk_font_chooser_get_show_preview_entry")
	core.PuregoSafeRegister(&XGtkFontChooserSetFilterFunc, lib, "gtk_font_chooser_set_filter_func")
	core.PuregoSafeRegister(&XGtkFontChooserSetFont, lib, "gtk_font_chooser_set_font")
	core.PuregoSafeRegister(&XGtkFontChooserSetFontDesc, lib, "gtk_font_chooser_set_font_desc")
	core.PuregoSafeRegister(&XGtkFontChooserSetFontMap, lib, "gtk_font_chooser_set_font_map")
	core.PuregoSafeRegister(&XGtkFontChooserSetLanguage, lib, "gtk_font_chooser_set_language")
	core.PuregoSafeRegister(&XGtkFontChooserSetLevel, lib, "gtk_font_chooser_set_level")
	core.PuregoSafeRegister(&XGtkFontChooserSetPreviewText, lib, "gtk_font_chooser_set_preview_text")
	core.PuregoSafeRegister(&XGtkFontChooserSetShowPreviewEntry, lib, "gtk_font_chooser_set_show_preview_entry")

}
