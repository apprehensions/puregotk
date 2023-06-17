// Package glib was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package glib

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
)

var xBookmarkFileErrorQuark func() Quark

func BookmarkFileErrorQuark() Quark {

	return xBookmarkFileErrorQuark()

}

var xConvertErrorQuark func() Quark

func ConvertErrorQuark() Quark {

	return xConvertErrorQuark()

}

var xFileErrorQuark func() Quark

func FileErrorQuark() Quark {

	return xFileErrorQuark()

}

var xIoChannelErrorQuark func() Quark

func IoChannelErrorQuark() Quark {

	return xIoChannelErrorQuark()

}

var xKeyFileErrorQuark func() Quark

func KeyFileErrorQuark() Quark {

	return xKeyFileErrorQuark()

}

var xMarkupErrorQuark func() Quark

func MarkupErrorQuark() Quark {

	return xMarkupErrorQuark()

}

var xNumberParserErrorQuark func() Quark

func NumberParserErrorQuark() Quark {

	return xNumberParserErrorQuark()

}

var xOptionErrorQuark func() Quark

func OptionErrorQuark() Quark {

	return xOptionErrorQuark()

}

var xRegexErrorQuark func() Quark

func RegexErrorQuark() Quark {

	return xRegexErrorQuark()

}

var xShellErrorQuark func() Quark

func ShellErrorQuark() Quark {

	return xShellErrorQuark()

}

var xSpawnErrorQuark func() Quark

func SpawnErrorQuark() Quark {

	return xSpawnErrorQuark()

}

var xSpawnExitErrorQuark func() Quark

func SpawnExitErrorQuark() Quark {

	return xSpawnExitErrorQuark()

}

var xThreadErrorQuark func() Quark

func ThreadErrorQuark() Quark {

	return xThreadErrorQuark()

}

var xUnixErrorQuark func() Quark

func UnixErrorQuark() Quark {

	return xUnixErrorQuark()

}

var xUriErrorQuark func() Quark

func UriErrorQuark() Quark {

	return xUriErrorQuark()

}

var xVariantParseErrorQuark func() Quark

func VariantParseErrorQuark() Quark {

	return xVariantParseErrorQuark()

}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GLIB"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	core.PuregoSafeRegister(&xBookmarkFileErrorQuark, lib, "g_bookmark_file_error_quark")
	core.PuregoSafeRegister(&xConvertErrorQuark, lib, "g_convert_error_quark")
	core.PuregoSafeRegister(&xFileErrorQuark, lib, "g_file_error_quark")
	core.PuregoSafeRegister(&xIoChannelErrorQuark, lib, "g_io_channel_error_quark")
	core.PuregoSafeRegister(&xKeyFileErrorQuark, lib, "g_key_file_error_quark")
	core.PuregoSafeRegister(&xMarkupErrorQuark, lib, "g_markup_error_quark")
	core.PuregoSafeRegister(&xNumberParserErrorQuark, lib, "g_number_parser_error_quark")
	core.PuregoSafeRegister(&xOptionErrorQuark, lib, "g_option_error_quark")
	core.PuregoSafeRegister(&xRegexErrorQuark, lib, "g_regex_error_quark")
	core.PuregoSafeRegister(&xShellErrorQuark, lib, "g_shell_error_quark")
	core.PuregoSafeRegister(&xSpawnErrorQuark, lib, "g_spawn_error_quark")
	core.PuregoSafeRegister(&xSpawnExitErrorQuark, lib, "g_spawn_exit_error_quark")
	core.PuregoSafeRegister(&xThreadErrorQuark, lib, "g_thread_error_quark")
	core.PuregoSafeRegister(&xUnixErrorQuark, lib, "g_unix_error_quark")
	core.PuregoSafeRegister(&xUriErrorQuark, lib, "g_uri_error_quark")
	core.PuregoSafeRegister(&xVariantParseErrorQuark, lib, "g_variant_parse_error_quark")

}
