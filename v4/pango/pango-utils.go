// Package pango was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package pango

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/glib"
)

var xFindParagraphBoundary func(string, int, int, int)

// Locates a paragraph boundary in @text.
//
// A boundary is caused by delimiter characters, such as
// a newline, carriage return, carriage return-newline pair,
// or Unicode paragraph separator character.
//
// The index of the run of delimiters is returned in
// @paragraph_delimiter_index. The index of the start
// of the paragrap (index after all delimiters) is stored
// in @next_paragraph_start.
//
// If no delimiters are found, both @paragraph_delimiter_index
// and @next_paragraph_start are filled with the length of @text
// (an index one off the end).
func FindParagraphBoundary(TextVar string, LengthVar int, ParagraphDelimiterIndexVar int, NextParagraphStartVar int) {

	xFindParagraphBoundary(TextVar, LengthVar, ParagraphDelimiterIndexVar, NextParagraphStartVar)

}

var xIsZeroWidth func(uint32) bool

// Checks if a character that should not be normally rendered.
//
// This includes all Unicode characters with "ZERO WIDTH" in their name,
// as well as *bidi* formatting characters, and a few other ones.
//
// This is totally different from [func@GLib.unichar_iszerowidth] and is at best misnamed.
func IsZeroWidth(ChVar uint32) bool {

	cret := xIsZeroWidth(ChVar)
	return cret
}

var xLog2visGetEmbeddingLevels func(string, int, *Direction) byte

// Return the bidirectional embedding levels of the input paragraph.
//
// The bidirectional embedding levels are defined by the [Unicode Bidirectional
// Algorithm](http://www.unicode.org/reports/tr9/).
//
// If the input base direction is a weak direction, the direction of the
// characters in the text will determine the final resolved direction.
func Log2visGetEmbeddingLevels(TextVar string, LengthVar int, PbaseDirVar *Direction) byte {

	cret := xLog2visGetEmbeddingLevels(TextVar, LengthVar, PbaseDirVar)
	return cret
}

var xParseEnum func([]interface{}, string, int, bool, string) bool

// Parses an enum type and stores the result in @value.
//
// If @str does not match the nick name of any of the possible values
// for the enum and is not an integer, %FALSE is returned, a warning
// is issued if @warn is %TRUE, and a string representing the list of
// possible values is stored in @possible_values. The list is
// slash-separated, eg. "none/start/middle/end".
//
// If failed and @possible_values is not %NULL, returned string should
// be freed using g_free().
func ParseEnum(TypeVar []interface{}, StrVar string, ValueVar int, WarnVar bool, PossibleValuesVar string) bool {

	cret := xParseEnum(TypeVar, StrVar, ValueVar, WarnVar, PossibleValuesVar)
	return cret
}

var xParseStretch func(string, *Stretch, bool) bool

// Parses a font stretch.
//
// The allowed values are
// "ultra_condensed", "extra_condensed", "condensed",
// "semi_condensed", "normal", "semi_expanded", "expanded",
// "extra_expanded" and "ultra_expanded". Case variations are
// ignored and the '_' characters may be omitted.
func ParseStretch(StrVar string, StretchVar *Stretch, WarnVar bool) bool {

	cret := xParseStretch(StrVar, StretchVar, WarnVar)
	return cret
}

var xParseStyle func(string, *Style, bool) bool

// Parses a font style.
//
// The allowed values are "normal", "italic" and "oblique", case
// variations being
// ignored.
func ParseStyle(StrVar string, StyleVar *Style, WarnVar bool) bool {

	cret := xParseStyle(StrVar, StyleVar, WarnVar)
	return cret
}

var xParseVariant func(string, *Variant, bool) bool

// Parses a font variant.
//
// The allowed values are "normal", "small-caps", "all-small-caps",
// "petite-caps", "all-petite-caps", "unicase" and "title-caps",
// case variations being ignored.
func ParseVariant(StrVar string, VariantVar *Variant, WarnVar bool) bool {

	cret := xParseVariant(StrVar, VariantVar, WarnVar)
	return cret
}

var xParseWeight func(string, *Weight, bool) bool

// Parses a font weight.
//
// The allowed values are "heavy",
// "ultrabold", "bold", "normal", "light", "ultraleight"
// and integers. Case variations are ignored.
func ParseWeight(StrVar string, WeightVar *Weight, WarnVar bool) bool {

	cret := xParseWeight(StrVar, WeightVar, WarnVar)
	return cret
}

var xQuantizeLineGeometry func(int, int)

// Quantizes the thickness and position of a line to whole device pixels.
//
// This is typically used for underline or strikethrough. The purpose of
// this function is to avoid such lines looking blurry.
//
// Care is taken to make sure @thickness is at least one pixel when this
// function returns, but returned @position may become zero as a result
// of rounding.
func QuantizeLineGeometry(ThicknessVar int, PositionVar int) {

	xQuantizeLineGeometry(ThicknessVar, PositionVar)

}

var xReadLine func(uintptr, *glib.String) int

// Reads an entire line from a file into a buffer.
//
// Lines may be delimited with '\n', '\r', '\n\r', or '\r\n'. The delimiter
// is not written into the buffer. Text after a '#' character is treated as
// a comment and skipped. '\' can be used to escape a # character.
// '\' proceeding a line delimiter combines adjacent lines. A '\' proceeding
// any other character is ignored and written into the output buffer
// unmodified.
func ReadLine(StreamVar uintptr, StrVar *glib.String) int {

	cret := xReadLine(StreamVar, StrVar)
	return cret
}

var xScanInt func(string, int) bool

// Scans an integer.
//
// Leading white space is skipped.
func ScanInt(PosVar string, OutVar int) bool {

	cret := xScanInt(PosVar, OutVar)
	return cret
}

var xScanString func(string, *glib.String) bool

// Scans a string into a `GString` buffer.
//
// The string may either be a sequence of non-white-space characters,
// or a quoted string with '"'. Instead a quoted string, '\"' represents
// a literal quote. Leading white space outside of quotes is skipped.
func ScanString(PosVar string, OutVar *glib.String) bool {

	cret := xScanString(PosVar, OutVar)
	return cret
}

var xScanWord func(string, *glib.String) bool

// Scans a word into a `GString` buffer.
//
// A word consists of [A-Za-z_] followed by zero or more
// [A-Za-z_0-9]. Leading white space is skipped.
func ScanWord(PosVar string, OutVar *glib.String) bool {

	cret := xScanWord(PosVar, OutVar)
	return cret
}

var xSkipSpace func(string) bool

// Skips 0 or more characters of white space.
func SkipSpace(PosVar string) bool {

	cret := xSkipSpace(PosVar)
	return cret
}

var xSplitFileList func(string) []string

// Splits a %G_SEARCHPATH_SEPARATOR-separated list of files, stripping
// white space and substituting ~/ with $HOME/.
func SplitFileList(StrVar string) []string {

	cret := xSplitFileList(StrVar)
	return cret
}

var xTrimString func(string) string

// Trims leading and trailing whitespace from a string.
func TrimString(StrVar string) string {

	cret := xTrimString(StrVar)
	return cret
}

var xVersion func() int

// Returns the encoded version of Pango available at run-time.
//
// This is similar to the macro %PANGO_VERSION except that the macro
// returns the encoded version available at compile-time. A version
// number can be encoded into an integer using PANGO_VERSION_ENCODE().
func Version() int {

	cret := xVersion()
	return cret
}

var xVersionCheck func(int, int, int) string

// Checks that the Pango library in use is compatible with the
// given version.
//
// Generally you would pass in the constants %PANGO_VERSION_MAJOR,
// %PANGO_VERSION_MINOR, %PANGO_VERSION_MICRO as the three arguments
// to this function; that produces a check that the library in use at
// run-time is compatible with the version of Pango the application or
// module was compiled against.
//
// Compatibility is defined by two things: first the version
// of the running library is newer than the version
// @required_major.required_minor.@required_micro. Second
// the running library must be binary compatible with the
// version @required_major.required_minor.@required_micro
// (same major version.)
//
// For compile-time version checking use PANGO_VERSION_CHECK().
func VersionCheck(RequiredMajorVar int, RequiredMinorVar int, RequiredMicroVar int) string {

	cret := xVersionCheck(RequiredMajorVar, RequiredMinorVar, RequiredMicroVar)
	return cret
}

var xVersionString func() string

// Returns the version of Pango available at run-time.
//
// This is similar to the macro %PANGO_VERSION_STRING except that the
// macro returns the version available at compile-time.
func VersionString() string {

	cret := xVersionString()
	return cret
}

func init() {
	lib, err := purego.Dlopen(core.GetPath("PANGO"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	core.PuregoSafeRegister(&xFindParagraphBoundary, lib, "pango_find_paragraph_boundary")
	core.PuregoSafeRegister(&xIsZeroWidth, lib, "pango_is_zero_width")
	core.PuregoSafeRegister(&xLog2visGetEmbeddingLevels, lib, "pango_log2vis_get_embedding_levels")
	core.PuregoSafeRegister(&xParseEnum, lib, "pango_parse_enum")
	core.PuregoSafeRegister(&xParseStretch, lib, "pango_parse_stretch")
	core.PuregoSafeRegister(&xParseStyle, lib, "pango_parse_style")
	core.PuregoSafeRegister(&xParseVariant, lib, "pango_parse_variant")
	core.PuregoSafeRegister(&xParseWeight, lib, "pango_parse_weight")
	core.PuregoSafeRegister(&xQuantizeLineGeometry, lib, "pango_quantize_line_geometry")
	core.PuregoSafeRegister(&xReadLine, lib, "pango_read_line")
	core.PuregoSafeRegister(&xScanInt, lib, "pango_scan_int")
	core.PuregoSafeRegister(&xScanString, lib, "pango_scan_string")
	core.PuregoSafeRegister(&xScanWord, lib, "pango_scan_word")
	core.PuregoSafeRegister(&xSkipSpace, lib, "pango_skip_space")
	core.PuregoSafeRegister(&xSplitFileList, lib, "pango_split_file_list")
	core.PuregoSafeRegister(&xTrimString, lib, "pango_trim_string")
	core.PuregoSafeRegister(&xVersion, lib, "pango_version")
	core.PuregoSafeRegister(&xVersionCheck, lib, "pango_version_check")
	core.PuregoSafeRegister(&xVersionString, lib, "pango_version_string")

}
