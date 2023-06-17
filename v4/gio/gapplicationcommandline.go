// Package gio was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gio

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/glib"
	"github.com/jwijenbergh/puregotk/v4/gobject"
)

// The #GApplicationCommandLineClass-struct
// contains private data only.
type ApplicationCommandLineClass struct {
	ParentClass uintptr

	Padding uintptr
}

type ApplicationCommandLinePrivate struct {
}

// #GApplicationCommandLine represents a command-line invocation of
// an application.  It is created by #GApplication and emitted
// in the #GApplication::command-line signal and virtual function.
//
// The class contains the list of arguments that the program was invoked
// with.  It is also possible to query if the commandline invocation was
// local (ie: the current process is running in direct response to the
// invocation) or remote (ie: some other process forwarded the
// commandline to this process).
//
// The GApplicationCommandLine object can provide the @argc and @argv
// parameters for use with the #GOptionContext command-line parsing API,
// with the g_application_command_line_get_arguments() function. See
// [gapplication-example-cmdline3.c][gapplication-example-cmdline3]
// for an example.
//
// The exit status of the originally-invoked process may be set and
// messages can be printed to stdout or stderr of that process.  The
// lifecycle of the originally-invoked process is tied to the lifecycle
// of this object (ie: the process exits when the last reference is
// dropped).
//
// The main use for #GApplicationCommandLine (and the
// #GApplication::command-line signal) is 'Emacs server' like use cases:
// You can set the `EDITOR` environment variable to have e.g. git use
// your favourite editor to edit commit messages, and if you already
// have an instance of the editor running, the editing will happen
// in the running instance, instead of opening a new one. An important
// aspect of this use case is that the process that gets started by git
// does not return until the editing is done.
//
// Normally, the commandline is completely handled in the
// #GApplication::command-line handler. The launching instance exits
// once the signal handler in the primary instance has returned, and
// the return value of the signal handler becomes the exit status
// of the launching instance.
// |[&lt;!-- language="C" --&gt;
// static int
// command_line (GApplication            *application,
//
//	GApplicationCommandLine *cmdline)
//
//	{
//	  gchar **argv;
//	  gint argc;
//	  gint i;
//
//	  argv = g_application_command_line_get_arguments (cmdline, &amp;argc);
//
//	  g_application_command_line_print (cmdline,
//	                                    "This text is written back\n"
//	                                    "to stdout of the caller\n");
//
//	  for (i = 0; i &lt; argc; i++)
//	    g_print ("argument %d: %s\n", i, argv[i]);
//
//	  g_strfreev (argv);
//
//	  return 0;
//	}
//
// ]|
// The complete example can be found here:
// [gapplication-example-cmdline.c](https://gitlab.gnome.org/GNOME/glib/-/blob/HEAD/gio/tests/gapplication-example-cmdline.c)
//
// In more complicated cases, the handling of the commandline can be
// split between the launcher and the primary instance.
// |[&lt;!-- language="C" --&gt;
// static gboolean
//
//	test_local_cmdline (GApplication   *application,
//	                    gchar        ***arguments,
//	                    gint           *exit_status)
//
//	{
//	  gint i, j;
//	  gchar **argv;
//
//	  argv = *arguments;
//
//	  if (argv[0] == NULL)
//	    {
//	      *exit_status = 0;
//	      return FALSE;
//	    }
//
//	  i = 1;
//	  while (argv[i])
//	    {
//	      if (g_str_has_prefix (argv[i], "--local-"))
//	        {
//	          g_print ("handling argument %s locally\n", argv[i]);
//	          g_free (argv[i]);
//	          for (j = i; argv[j]; j++)
//	            argv[j] = argv[j + 1];
//	        }
//	      else
//	        {
//	          g_print ("not handling argument %s locally\n", argv[i]);
//	          i++;
//	        }
//	    }
//
//	  *exit_status = 0;
//
//	  return FALSE;
//	}
//
// static void
// test_application_class_init (TestApplicationClass *class)
//
//	{
//	  G_APPLICATION_CLASS (class)-&gt;local_command_line = test_local_cmdline;
//
//	  ...
//	}
//
// ]|
// In this example of split commandline handling, options that start
// with `--local-` are handled locally, all other options are passed
// to the #GApplication::command-line handler which runs in the primary
// instance.
//
// The complete example can be found here:
// [gapplication-example-cmdline2.c](https://gitlab.gnome.org/GNOME/glib/-/blob/HEAD/gio/tests/gapplication-example-cmdline2.c)
//
// If handling the commandline requires a lot of work, it may
// be better to defer it.
// |[&lt;!-- language="C" --&gt;
// static gboolean
// my_cmdline_handler (gpointer data)
//
//	{
//	  GApplicationCommandLine *cmdline = data;
//
//	  // do the heavy lifting in an idle
//
//	  g_application_command_line_set_exit_status (cmdline, 0);
//	  g_object_unref (cmdline); // this releases the application
//
//	  return G_SOURCE_REMOVE;
//	}
//
// static int
// command_line (GApplication            *application,
//
//	GApplicationCommandLine *cmdline)
//
//	{
//	  // keep the application running until we are done with this commandline
//	  g_application_hold (application);
//
//	  g_object_set_data_full (G_OBJECT (cmdline),
//	                          "application", application,
//	                          (GDestroyNotify)g_application_release);
//
//	  g_object_ref (cmdline);
//	  g_idle_add (my_cmdline_handler, cmdline);
//
//	  return 0;
//	}
//
// ]|
// In this example the commandline is not completely handled before
// the #GApplication::command-line handler returns. Instead, we keep
// a reference to the #GApplicationCommandLine object and handle it
// later (in this example, in an idle). Note that it is necessary to
// hold the application until you are done with the commandline.
//
// The complete example can be found here:
// [gapplication-example-cmdline3.c](https://gitlab.gnome.org/GNOME/glib/-/blob/HEAD/gio/tests/gapplication-example-cmdline3.c)
type ApplicationCommandLine struct {
	gobject.Object
}

func ApplicationCommandLineNewFromInternalPtr(ptr uintptr) *ApplicationCommandLine {
	cls := &ApplicationCommandLine{}
	cls.Ptr = ptr
	return cls
}

var xApplicationCommandLineCreateFileForArg func(uintptr, string) uintptr

// Creates a #GFile corresponding to a filename that was given as part
// of the invocation of @cmdline.
//
// This differs from g_file_new_for_commandline_arg() in that it
// resolves relative pathnames using the current working directory of
// the invoking process rather than the local process.
func (x *ApplicationCommandLine) CreateFileForArg(ArgVar string) *FileBase {

	CreateFileForArgPtr := xApplicationCommandLineCreateFileForArg(x.GoPointer(), ArgVar)
	if CreateFileForArgPtr == 0 {
		return nil
	}

	CreateFileForArgCls := &FileBase{}
	CreateFileForArgCls.Ptr = CreateFileForArgPtr
	return CreateFileForArgCls

}

var xApplicationCommandLineGetArguments func(uintptr, int32) uintptr

// Gets the list of arguments that was passed on the command line.
//
// The strings in the array may contain non-UTF-8 data on UNIX (such as
// filenames or arguments given in the system locale) but are always in
// UTF-8 on Windows.
//
// If you wish to use the return value with #GOptionContext, you must
// use g_option_context_parse_strv().
//
// The return value is %NULL-terminated and should be freed using
// g_strfreev().
func (x *ApplicationCommandLine) GetArguments(ArgcVar int32) uintptr {

	return xApplicationCommandLineGetArguments(x.GoPointer(), ArgcVar)

}

var xApplicationCommandLineGetCwd func(uintptr) string

// Gets the working directory of the command line invocation.
// The string may contain non-utf8 data.
//
// It is possible that the remote application did not send a working
// directory, so this may be %NULL.
//
// The return value should not be modified or freed and is valid for as
// long as @cmdline exists.
func (x *ApplicationCommandLine) GetCwd() string {

	return xApplicationCommandLineGetCwd(x.GoPointer())

}

var xApplicationCommandLineGetEnviron func(uintptr) uintptr

// Gets the contents of the 'environ' variable of the command line
// invocation, as would be returned by g_get_environ(), ie as a
// %NULL-terminated list of strings in the form 'NAME=VALUE'.
// The strings may contain non-utf8 data.
//
// The remote application usually does not send an environment.  Use
// %G_APPLICATION_SEND_ENVIRONMENT to affect that.  Even with this flag
// set it is possible that the environment is still not available (due
// to invocation messages from other applications).
//
// The return value should not be modified or freed and is valid for as
// long as @cmdline exists.
//
// See g_application_command_line_getenv() if you are only interested
// in the value of a single environment variable.
func (x *ApplicationCommandLine) GetEnviron() uintptr {

	return xApplicationCommandLineGetEnviron(x.GoPointer())

}

var xApplicationCommandLineGetExitStatus func(uintptr) int32

// Gets the exit status of @cmdline.  See
// g_application_command_line_set_exit_status() for more information.
func (x *ApplicationCommandLine) GetExitStatus() int32 {

	return xApplicationCommandLineGetExitStatus(x.GoPointer())

}

var xApplicationCommandLineGetIsRemote func(uintptr) bool

// Determines if @cmdline represents a remote invocation.
func (x *ApplicationCommandLine) GetIsRemote() bool {

	return xApplicationCommandLineGetIsRemote(x.GoPointer())

}

var xApplicationCommandLineGetOptionsDict func(uintptr) *glib.VariantDict

// Gets the options there were passed to g_application_command_line().
//
// If you did not override local_command_line() then these are the same
// options that were parsed according to the #GOptionEntrys added to the
// application with g_application_add_main_option_entries() and possibly
// modified from your GApplication::handle-local-options handler.
//
// If no options were sent then an empty dictionary is returned so that
// you don't need to check for %NULL.
func (x *ApplicationCommandLine) GetOptionsDict() *glib.VariantDict {

	return xApplicationCommandLineGetOptionsDict(x.GoPointer())

}

var xApplicationCommandLineGetPlatformData func(uintptr) *glib.Variant

// Gets the platform data associated with the invocation of @cmdline.
//
// This is a #GVariant dictionary containing information about the
// context in which the invocation occurred.  It typically contains
// information like the current working directory and the startup
// notification ID.
//
// For local invocation, it will be %NULL.
func (x *ApplicationCommandLine) GetPlatformData() *glib.Variant {

	return xApplicationCommandLineGetPlatformData(x.GoPointer())

}

var xApplicationCommandLineGetStdin func(uintptr) uintptr

// Gets the stdin of the invoking process.
//
// The #GInputStream can be used to read data passed to the standard
// input of the invoking process.
// This doesn't work on all platforms.  Presently, it is only available
// on UNIX when using a D-Bus daemon capable of passing file descriptors.
// If stdin is not available then %NULL will be returned.  In the
// future, support may be expanded to other platforms.
//
// You must only call this function once per commandline invocation.
func (x *ApplicationCommandLine) GetStdin() *InputStream {

	GetStdinPtr := xApplicationCommandLineGetStdin(x.GoPointer())
	if GetStdinPtr == 0 {
		return nil
	}

	GetStdinCls := &InputStream{}
	GetStdinCls.Ptr = GetStdinPtr
	return GetStdinCls

}

var xApplicationCommandLineGetenv func(uintptr, string) string

// Gets the value of a particular environment variable of the command
// line invocation, as would be returned by g_getenv().  The strings may
// contain non-utf8 data.
//
// The remote application usually does not send an environment.  Use
// %G_APPLICATION_SEND_ENVIRONMENT to affect that.  Even with this flag
// set it is possible that the environment is still not available (due
// to invocation messages from other applications).
//
// The return value should not be modified or freed and is valid for as
// long as @cmdline exists.
func (x *ApplicationCommandLine) Getenv(NameVar string) string {

	return xApplicationCommandLineGetenv(x.GoPointer(), NameVar)

}

var xApplicationCommandLinePrint func(uintptr, string, ...interface{})

// Formats a message and prints it using the stdout print handler in the
// invoking process.
//
// If @cmdline is a local invocation then this is exactly equivalent to
// g_print().  If @cmdline is remote then this is equivalent to calling
// g_print() in the invoking process.
func (x *ApplicationCommandLine) Print(FormatVar string, varArgs ...interface{}) {

	xApplicationCommandLinePrint(x.GoPointer(), FormatVar, varArgs...)

}

var xApplicationCommandLinePrinterr func(uintptr, string, ...interface{})

// Formats a message and prints it using the stderr print handler in the
// invoking process.
//
// If @cmdline is a local invocation then this is exactly equivalent to
// g_printerr().  If @cmdline is remote then this is equivalent to
// calling g_printerr() in the invoking process.
func (x *ApplicationCommandLine) Printerr(FormatVar string, varArgs ...interface{}) {

	xApplicationCommandLinePrinterr(x.GoPointer(), FormatVar, varArgs...)

}

var xApplicationCommandLineSetExitStatus func(uintptr, int32)

// Sets the exit status that will be used when the invoking process
// exits.
//
// The return value of the #GApplication::command-line signal is
// passed to this function when the handler returns.  This is the usual
// way of setting the exit status.
//
// In the event that you want the remote invocation to continue running
// and want to decide on the exit status in the future, you can use this
// call.  For the case of a remote invocation, the remote process will
// typically exit when the last reference is dropped on @cmdline.  The
// exit status of the remote process will be equal to the last value
// that was set with this function.
//
// In the case that the commandline invocation is local, the situation
// is slightly more complicated.  If the commandline invocation results
// in the mainloop running (ie: because the use-count of the application
// increased to a non-zero value) then the application is considered to
// have been 'successful' in a certain sense, and the exit status is
// always zero.  If the application use count is zero, though, the exit
// status of the local #GApplicationCommandLine is used.
func (x *ApplicationCommandLine) SetExitStatus(ExitStatusVar int32) {

	xApplicationCommandLineSetExitStatus(x.GoPointer(), ExitStatusVar)

}

func (c *ApplicationCommandLine) GoPointer() uintptr {
	return c.Ptr
}

func (c *ApplicationCommandLine) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GIO"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xApplicationCommandLineCreateFileForArg, lib, "g_application_command_line_create_file_for_arg")
	core.PuregoSafeRegister(&xApplicationCommandLineGetArguments, lib, "g_application_command_line_get_arguments")
	core.PuregoSafeRegister(&xApplicationCommandLineGetCwd, lib, "g_application_command_line_get_cwd")
	core.PuregoSafeRegister(&xApplicationCommandLineGetEnviron, lib, "g_application_command_line_get_environ")
	core.PuregoSafeRegister(&xApplicationCommandLineGetExitStatus, lib, "g_application_command_line_get_exit_status")
	core.PuregoSafeRegister(&xApplicationCommandLineGetIsRemote, lib, "g_application_command_line_get_is_remote")
	core.PuregoSafeRegister(&xApplicationCommandLineGetOptionsDict, lib, "g_application_command_line_get_options_dict")
	core.PuregoSafeRegister(&xApplicationCommandLineGetPlatformData, lib, "g_application_command_line_get_platform_data")
	core.PuregoSafeRegister(&xApplicationCommandLineGetStdin, lib, "g_application_command_line_get_stdin")
	core.PuregoSafeRegister(&xApplicationCommandLineGetenv, lib, "g_application_command_line_getenv")
	core.PuregoSafeRegister(&xApplicationCommandLinePrint, lib, "g_application_command_line_print")
	core.PuregoSafeRegister(&xApplicationCommandLinePrinterr, lib, "g_application_command_line_printerr")
	core.PuregoSafeRegister(&xApplicationCommandLineSetExitStatus, lib, "g_application_command_line_set_exit_status")

}
