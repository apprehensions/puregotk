// Package glib was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package glib

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
)

var xOnErrorQuery func(string)

// Prompts the user with
// `[E]xit, [H]alt, show [S]tack trace or [P]roceed`.
// This function is intended to be used for debugging use only.
// The following example shows how it can be used together with
// the g_log() functions.
//
// |[&lt;!-- language="C" --&gt;
// #include &lt;glib.h&gt;
//
// static void
// log_handler (const gchar   *log_domain,
//
//	GLogLevelFlags log_level,
//	const gchar   *message,
//	gpointer       user_data)
//
//	{
//	  g_log_default_handler (log_domain, log_level, message, user_data);
//
//	  g_on_error_query (MY_PROGRAM_NAME);
//	}
//
// int
// main (int argc, char *argv[])
//
//	{
//	  g_log_set_handler (MY_LOG_DOMAIN,
//	                     G_LOG_LEVEL_WARNING |
//	                     G_LOG_LEVEL_ERROR |
//	                     G_LOG_LEVEL_CRITICAL,
//	                     log_handler,
//	                     NULL);
//	  ...
//
// ]|
//
// If "[E]xit" is selected, the application terminates with a call
// to _exit(0).
//
// If "[S]tack" trace is selected, g_on_error_stack_trace() is called.
// This invokes gdb, which attaches to the current process and shows
// a stack trace. The prompt is then shown again.
//
// If "[P]roceed" is selected, the function returns.
//
// This function may cause different actions on non-UNIX platforms.
//
// On Windows consider using the `G_DEBUGGER` environment
// variable (see [Running GLib Applications](glib-running.html)) and
// calling g_on_error_stack_trace() instead.
func OnErrorQuery(PrgNameVar string) {

	xOnErrorQuery(PrgNameVar)

}

var xOnErrorStackTrace func(string)

// Invokes gdb, which attaches to the current process and shows a
// stack trace. Called by g_on_error_query() when the "[S]tack trace"
// option is selected. You can get the current process's program name
// with g_get_prgname(), assuming that you have called gtk_init() or
// gdk_init().
//
// This function may cause different actions on non-UNIX platforms.
//
// When running on Windows, this function is *not* called by
// g_on_error_query(). If called directly, it will raise an
// exception, which will crash the program. If the `G_DEBUGGER` environment
// variable is set, a debugger will be invoked to attach and
// handle that exception (see [Running GLib Applications](glib-running.html)).
func OnErrorStackTrace(PrgNameVar string) {

	xOnErrorStackTrace(PrgNameVar)

}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GLIB"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	core.PuregoSafeRegister(&xOnErrorQuery, lib, "g_on_error_query")
	core.PuregoSafeRegister(&xOnErrorStackTrace, lib, "g_on_error_stack_trace")

}
