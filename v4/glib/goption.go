// Package glib was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package glib

// The type of function to be passed as callback for %G_OPTION_ARG_CALLBACK
// options.
type OptionArgFunc func(string, string, uintptr) bool

// The type of function to be used as callback when a parse error occurs.
type OptionErrorFunc func(*OptionContext, *OptionGroup, uintptr)

// The type of function that can be called before and after parsing.
type OptionParseFunc func(*OptionContext, *OptionGroup, uintptr) bool

// A `GOptionContext` struct defines which options
// are accepted by the commandline option parser. The struct has only private
// fields and should not be directly accessed.
type OptionContext struct {
}

// A GOptionEntry struct defines a single option. To have an effect, they
// must be added to a #GOptionGroup with g_option_context_add_main_entries()
// or g_option_group_add_entries().
type OptionEntry struct {
	LongName string

	ShortName byte

	Flags int32

	Arg OptionArg

	ArgData uintptr

	Description string

	ArgDescription string
}

// A `GOptionGroup` struct defines the options in a single
// group. The struct has only private fields and should not be directly accessed.
//
// All options in a group share the same translation function. Libraries which
// need to parse commandline options are expected to provide a function for
// getting a `GOptionGroup` holding their options, which
// the application can then add to its #GOptionContext.
type OptionGroup struct {
}

// Flags which modify individual options.
type OptionFlags int

const (

	// No flags. Since: 2.42.
	GOptionFlagNoneValue OptionFlags = 0
	// The option doesn't appear in `--help` output.
	GOptionFlagHiddenValue OptionFlags = 1
	// The option appears in the main section of the
	//     `--help` output, even if it is defined in a group.
	GOptionFlagInMainValue OptionFlags = 2
	// For options of the %G_OPTION_ARG_NONE kind, this
	//     flag indicates that the sense of the option is reversed.
	GOptionFlagReverseValue OptionFlags = 4
	// For options of the %G_OPTION_ARG_CALLBACK kind,
	//     this flag indicates that the callback does not take any argument
	//     (like a %G_OPTION_ARG_NONE option). Since 2.8
	GOptionFlagNoArgValue OptionFlags = 8
	// For options of the %G_OPTION_ARG_CALLBACK
	//     kind, this flag indicates that the argument should be passed to the
	//     callback in the GLib filename encoding rather than UTF-8. Since 2.8
	GOptionFlagFilenameValue OptionFlags = 16
	// For options of the %G_OPTION_ARG_CALLBACK
	//     kind, this flag indicates that the argument supply is optional.
	//     If no argument is given then data of %GOptionParseFunc will be
	//     set to NULL. Since 2.8
	GOptionFlagOptionalArgValue OptionFlags = 32
	// This flag turns off the automatic conflict
	//     resolution which prefixes long option names with `groupname-` if
	//     there is a conflict. This option should only be used in situations
	//     where aliasing is necessary to model some legacy commandline interface.
	//     It is not safe to use this option, unless all option groups are under
	//     your direct control. Since 2.8.
	GOptionFlagNoaliasValue OptionFlags = 64
)

// The #GOptionArg enum values determine which type of extra argument the
// options expect to find. If an option expects an extra argument, it can
// be specified in several ways; with a short option: `-x arg`, with a long
// option: `--name arg` or combined in a single argument: `--name=arg`.
type OptionArg int

const (

	// No extra argument. This is useful for simple flags.
	GOptionArgNoneValue OptionArg = 0
	// The option takes a UTF-8 string argument.
	GOptionArgStringValue OptionArg = 1
	// The option takes an integer argument.
	GOptionArgIntValue OptionArg = 2
	// The option provides a callback (of type
	//     #GOptionArgFunc) to parse the extra argument.
	GOptionArgCallbackValue OptionArg = 3
	// The option takes a filename as argument, which will
	//        be in the GLib filename encoding rather than UTF-8.
	GOptionArgFilenameValue OptionArg = 4
	// The option takes a string argument, multiple
	//     uses of the option are collected into an array of strings.
	GOptionArgStringArrayValue OptionArg = 5
	// The option takes a filename as argument,
	//     multiple uses of the option are collected into an array of strings.
	GOptionArgFilenameArrayValue OptionArg = 6
	// The option takes a double argument. The argument
	//     can be formatted either for the user's locale or for the "C" locale.
	//     Since 2.12
	GOptionArgDoubleValue OptionArg = 7
	// The option takes a 64-bit integer. Like
	//     %G_OPTION_ARG_INT but for larger numbers. The number can be in
	//     decimal base, or in hexadecimal (when prefixed with `0x`, for
	//     example, `0xffffffff`). Since 2.12
	GOptionArgInt64Value OptionArg = 8
)

// Error codes returned by option parsing.
type OptionError int

const (

	// An option was not known to the parser.
	//  This error will only be reported, if the parser hasn't been instructed
	//  to ignore unknown options, see g_option_context_set_ignore_unknown_options().
	GOptionErrorUnknownOptionValue OptionError = 0
	// A value couldn't be parsed.
	GOptionErrorBadValueValue OptionError = 1
	// A #GOptionArgFunc callback failed.
	GOptionErrorFailedValue OptionError = 2
)
