// Package glib was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package glib

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
)

// Prototype of a #GChildWatchSource callback, called when a child
// process has exited.
//
// To interpret @wait_status, see the documentation
// for g_spawn_check_wait_status(). In particular,
// on Unix platforms, note that it is usually not equal
// to the integer passed to `exit()` or returned from `main()`.
type ChildWatchFunc func(Pid, int32, uintptr)

// Specifies the type of function passed to g_clear_handle_id().
// The implementation is expected to free the resource identified
// by @handle_id; for instance, if @handle_id is a #GSource ID,
// g_source_remove() can be used.
type ClearHandleFunc func(uint)

// Dispose function for @source. See g_source_set_dispose_function() for
// details.
type SourceDisposeFunc func(*Source)

// This is just a placeholder for #GClosureMarshal,
// which cannot be used here for dependency reasons.
type SourceDummyMarshal func()

// Specifies the type of function passed to g_timeout_add(),
// g_timeout_add_full(), g_idle_add(), and g_idle_add_full().
//
// When calling g_source_set_callback(), you may need to cast a function of a
// different type to this type. Use G_SOURCE_FUNC() to avoid warnings about
// incompatible function types.
type SourceFunc func(uintptr) bool

// The `GMainContext` struct is an opaque data
// type representing a set of sources to be handled in a main loop.
type MainContext struct {
}

// The `GMainLoop` struct is an opaque data type
// representing the main event loop of a GLib or GTK+ application.
type MainLoop struct {
}

// The `GSource` struct is an opaque data type
// representing an event source.
type Source struct {
	CallbackData uintptr

	CallbackFuncs *SourceCallbackFuncs

	SourceFuncs *SourceFuncs

	RefCount uint

	Context *MainContext

	Priority int32

	Flags uint

	SourceId uint

	PollFds *SList

	Prev *Source

	Next *Source

	Name string

	Priv *SourcePrivate
}

// The `GSourceCallbackFuncs` struct contains
// functions for managing callback objects.
type SourceCallbackFuncs struct {
}

// The `GSourceFuncs` struct contains a table of
// functions used to handle event sources in a generic manner.
//
// For idle sources, the prepare and check functions always return %TRUE
// to indicate that the source is always ready to be processed. The prepare
// function also returns a timeout value of 0 to ensure that the poll() call
// doesn't block (since that would be time wasted which could have been spent
// running the idle function).
//
// For timeout sources, the prepare and check functions both return %TRUE
// if the timeout interval has expired. The prepare function also returns
// a timeout value to ensure that the poll() call doesn't block too long
// and miss the next timeout.
//
// For file descriptor sources, the prepare function typically returns %FALSE,
// since it must wait until poll() has been called before it knows whether
// any events need to be processed. It sets the returned timeout to -1 to
// indicate that it doesn't mind how long the poll() call blocks. In the
// check function, it tests the results of the poll() call to see if the
// required condition has been met, and returns %TRUE if so.
type SourceFuncs struct {
	ClosureCallback SourceFunc

	ClosureMarshal SourceDummyMarshal
}

type SourcePrivate struct {
}

// Opaque type. See g_main_context_pusher_new() for details.
type MainContextPusher = uintptr

// Flags to pass to g_main_context_new_with_flags() which affect the behaviour
// of a #GMainContext.
type MainContextFlags int

const (

	// Default behaviour.
	GMainContextFlagsNoneValue MainContextFlags = 0
	// Assume that polling for events will
	// free the thread to process other jobs. That's useful if you're using
	// `g_main_context_{prepare,query,check,dispatch}` to integrate GMainContext in
	// other event loops.
	GMainContextFlagsOwnerlessPollingValue MainContextFlags = 1
)

var xChildWatchAdd func(Pid, uintptr, uintptr) uint

// Sets a function to be called when the child indicated by @pid
// exits, at a default priority, %G_PRIORITY_DEFAULT.
//
// If you obtain @pid from g_spawn_async() or g_spawn_async_with_pipes()
// you will need to pass %G_SPAWN_DO_NOT_REAP_CHILD as flag to
// the spawn function for the child watching to work.
//
// Note that on platforms where #GPid must be explicitly closed
// (see g_spawn_close_pid()) @pid must not be closed while the
// source is still active. Typically, you will want to call
// g_spawn_close_pid() in the callback function for the source.
//
// GLib supports only a single callback per process id.
// On POSIX platforms, the same restrictions mentioned for
// g_child_watch_source_new() apply to this function.
//
// This internally creates a main loop source using
// g_child_watch_source_new() and attaches it to the main loop context
// using g_source_attach(). You can do these steps manually if you
// need greater control.
func ChildWatchAdd(PidVar Pid, FunctionVar ChildWatchFunc, DataVar uintptr) uint {

	return xChildWatchAdd(PidVar, purego.NewCallback(FunctionVar), DataVar)

}

var xChildWatchAddFull func(int32, Pid, uintptr, uintptr, uintptr) uint

// Sets a function to be called when the child indicated by @pid
// exits, at the priority @priority.
//
// If you obtain @pid from g_spawn_async() or g_spawn_async_with_pipes()
// you will need to pass %G_SPAWN_DO_NOT_REAP_CHILD as flag to
// the spawn function for the child watching to work.
//
// In many programs, you will want to call g_spawn_check_wait_status()
// in the callback to determine whether or not the child exited
// successfully.
//
// Also, note that on platforms where #GPid must be explicitly closed
// (see g_spawn_close_pid()) @pid must not be closed while the source
// is still active.  Typically, you should invoke g_spawn_close_pid()
// in the callback function for the source.
//
// GLib supports only a single callback per process id.
// On POSIX platforms, the same restrictions mentioned for
// g_child_watch_source_new() apply to this function.
//
// This internally creates a main loop source using
// g_child_watch_source_new() and attaches it to the main loop context
// using g_source_attach(). You can do these steps manually if you
// need greater control.
func ChildWatchAddFull(PriorityVar int32, PidVar Pid, FunctionVar ChildWatchFunc, DataVar uintptr, NotifyVar DestroyNotify) uint {

	return xChildWatchAddFull(PriorityVar, PidVar, purego.NewCallback(FunctionVar), DataVar, purego.NewCallback(NotifyVar))

}

var xChildWatchSourceNew func(Pid) *Source

// Creates a new child_watch source.
//
// The source will not initially be associated with any #GMainContext
// and must be added to one with g_source_attach() before it will be
// executed.
//
// Note that child watch sources can only be used in conjunction with
// `g_spawn...` when the %G_SPAWN_DO_NOT_REAP_CHILD flag is used.
//
// Note that on platforms where #GPid must be explicitly closed
// (see g_spawn_close_pid()) @pid must not be closed while the
// source is still active. Typically, you will want to call
// g_spawn_close_pid() in the callback function for the source.
//
// On POSIX platforms, the following restrictions apply to this API
// due to limitations in POSIX process interfaces:
//
//   - @pid must be a child of this process
//   - @pid must be positive
//   - the application must not call `waitpid` with a non-positive
//     first argument, for instance in another thread
//   - the application must not wait for @pid to exit by any other
//     mechanism, including `waitpid(pid, ...)` or a second child-watch
//     source for the same @pid
//   - the application must not ignore `SIGCHLD`
//
// If any of those conditions are not met, this and related APIs will
// not work correctly. This can often be diagnosed via a GLib warning
// stating that `ECHILD` was received by `waitpid`.
//
// Calling `waitpid` for specific processes other than @pid remains a
// valid thing to do.
func ChildWatchSourceNew(PidVar Pid) *Source {

	return xChildWatchSourceNew(PidVar)

}

var xClearHandleId func(uint, uintptr)

// Clears a numeric handler, such as a #GSource ID.
//
// @tag_ptr must be a valid pointer to the variable holding the handler.
//
// If the ID is zero then this function does nothing.
// Otherwise, clear_func() is called with the ID as a parameter, and the tag is
// set to zero.
//
// A macro is also included that allows this function to be used without
// pointer casts.
func ClearHandleId(TagPtrVar uint, ClearFuncVar ClearHandleFunc) {

	xClearHandleId(TagPtrVar, purego.NewCallback(ClearFuncVar))

}

var xGetCurrentTime func(*TimeVal)

// Equivalent to the UNIX gettimeofday() function, but portable.
//
// You may find g_get_real_time() to be more convenient.
func GetCurrentTime(ResultVar *TimeVal) {

	xGetCurrentTime(ResultVar)

}

var xGetMonotonicTime func() int64

// Queries the system monotonic time.
//
// The monotonic clock will always increase and doesn't suffer
// discontinuities when the user (or NTP) changes the system time.  It
// may or may not continue to tick during times where the machine is
// suspended.
//
// We try to use the clock that corresponds as closely as possible to
// the passage of time as measured by system calls such as poll() but it
// may not always be possible to do this.
func GetMonotonicTime() int64 {

	return xGetMonotonicTime()

}

var xGetRealTime func() int64

// Queries the system wall-clock time.
//
// This call is functionally equivalent to g_get_current_time() except
// that the return value is often more convenient than dealing with a
// #GTimeVal.
//
// You should only use this call if you are actually interested in the real
// wall-clock time.  g_get_monotonic_time() is probably more useful for
// measuring intervals.
func GetRealTime() int64 {

	return xGetRealTime()

}

var xIdleAdd func(uintptr, uintptr) uint

// Adds a function to be called whenever there are no higher priority
// events pending to the default main loop. The function is given the
// default idle priority, %G_PRIORITY_DEFAULT_IDLE.  If the function
// returns %FALSE it is automatically removed from the list of event
// sources and will not be called again.
//
// See [memory management of sources][mainloop-memory-management] for details
// on how to handle the return value and memory management of @data.
//
// This internally creates a main loop source using g_idle_source_new()
// and attaches it to the global #GMainContext using g_source_attach(), so
// the callback will be invoked in whichever thread is running that main
// context. You can do these steps manually if you need greater control or to
// use a custom main context.
func IdleAdd(FunctionVar SourceFunc, DataVar uintptr) uint {

	return xIdleAdd(purego.NewCallback(FunctionVar), DataVar)

}

var xIdleAddFull func(int32, uintptr, uintptr, uintptr) uint

// Adds a function to be called whenever there are no higher priority
// events pending.
//
// If the function returns %G_SOURCE_REMOVE or %FALSE it is automatically
// removed from the list of event sources and will not be called again.
//
// See [memory management of sources][mainloop-memory-management] for details
// on how to handle the return value and memory management of @data.
//
// This internally creates a main loop source using g_idle_source_new()
// and attaches it to the global #GMainContext using g_source_attach(), so
// the callback will be invoked in whichever thread is running that main
// context. You can do these steps manually if you need greater control or to
// use a custom main context.
func IdleAddFull(PriorityVar int32, FunctionVar SourceFunc, DataVar uintptr, NotifyVar DestroyNotify) uint {

	return xIdleAddFull(PriorityVar, purego.NewCallback(FunctionVar), DataVar, purego.NewCallback(NotifyVar))

}

var xIdleRemoveByData func(uintptr) bool

// Removes the idle function with the given data.
func IdleRemoveByData(DataVar uintptr) bool {

	return xIdleRemoveByData(DataVar)

}

var xIdleSourceNew func() *Source

// Creates a new idle source.
//
// The source will not initially be associated with any #GMainContext
// and must be added to one with g_source_attach() before it will be
// executed. Note that the default priority for idle sources is
// %G_PRIORITY_DEFAULT_IDLE, as compared to other sources which
// have a default priority of %G_PRIORITY_DEFAULT.
func IdleSourceNew() *Source {

	return xIdleSourceNew()

}

var xMainContextDefault func() *MainContext

// Returns the global default main context. This is the main context
// used for main loop functions when a main loop is not explicitly
// specified, and corresponds to the "main" main loop. See also
// g_main_context_get_thread_default().
func MainContextDefault() *MainContext {

	return xMainContextDefault()

}

var xMainContextGetThreadDefault func() *MainContext

// Gets the thread-default #GMainContext for this thread. Asynchronous
// operations that want to be able to be run in contexts other than
// the default one should call this method or
// g_main_context_ref_thread_default() to get a #GMainContext to add
// their #GSources to. (Note that even in single-threaded
// programs applications may sometimes want to temporarily push a
// non-default context, so it is not safe to assume that this will
// always return %NULL if you are running in the default thread.)
//
// If you need to hold a reference on the context, use
// g_main_context_ref_thread_default() instead.
func MainContextGetThreadDefault() *MainContext {

	return xMainContextGetThreadDefault()

}

var xMainContextRefThreadDefault func() *MainContext

// Gets the thread-default #GMainContext for this thread, as with
// g_main_context_get_thread_default(), but also adds a reference to
// it with g_main_context_ref(). In addition, unlike
// g_main_context_get_thread_default(), if the thread-default context
// is the global default context, this will return that #GMainContext
// (with a ref added to it) rather than returning %NULL.
func MainContextRefThreadDefault() *MainContext {

	return xMainContextRefThreadDefault()

}

var xMainCurrentSource func() *Source

// Returns the currently firing source for this thread.
func MainCurrentSource() *Source {

	return xMainCurrentSource()

}

var xMainDepth func() int32

// Returns the depth of the stack of calls to
// g_main_context_dispatch() on any #GMainContext in the current thread.
//
//	That is, when called from the toplevel, it gives 0. When
//
// called from within a callback from g_main_context_iteration()
// (or g_main_loop_run(), etc.) it returns 1. When called from within
// a callback to a recursive call to g_main_context_iteration(),
// it returns 2. And so forth.
//
// This function is useful in a situation like the following:
// Imagine an extremely simple "garbage collected" system.
//
// |[&lt;!-- language="C" --&gt;
// static GList *free_list;
//
// gpointer
// allocate_memory (gsize size)
//
//	{
//	  gpointer result = g_malloc (size);
//	  free_list = g_list_prepend (free_list, result);
//	  return result;
//	}
//
// void
// free_allocated_memory (void)
//
//	{
//	  GList *l;
//	  for (l = free_list; l; l = l-&gt;next);
//	    g_free (l-&gt;data);
//	  g_list_free (free_list);
//	  free_list = NULL;
//	 }
//
// [...]
//
// while (TRUE);
//
//	{
//	  g_main_context_iteration (NULL, TRUE);
//	  free_allocated_memory();
//	 }
//
// ]|
//
// This works from an application, however, if you want to do the same
// thing from a library, it gets more difficult, since you no longer
// control the main loop. You might think you can simply use an idle
// function to make the call to free_allocated_memory(), but that
// doesn't work, since the idle function could be called from a
// recursive callback. This can be fixed by using g_main_depth()
//
// |[&lt;!-- language="C" --&gt;
// gpointer
// allocate_memory (gsize size)
//
//	{
//	  FreeListBlock *block = g_new (FreeListBlock, 1);
//	  block-&gt;mem = g_malloc (size);
//	  block-&gt;depth = g_main_depth ();
//	  free_list = g_list_prepend (free_list, block);
//	  return block-&gt;mem;
//	}
//
// void
// free_allocated_memory (void)
//
//	{
//	  GList *l;
//
//	  int depth = g_main_depth ();
//	  for (l = free_list; l; );
//	    {
//	      GList *next = l-&gt;next;
//	      FreeListBlock *block = l-&gt;data;
//	      if (block-&gt;depth &gt; depth)
//	        {
//	          g_free (block-&gt;mem);
//	          g_free (block);
//	          free_list = g_list_delete_link (free_list, l);
//	        }
//
//	      l = next;
//	    }
//	  }
//
// ]|
//
// There is a temptation to use g_main_depth() to solve
// problems with reentrancy. For instance, while waiting for data
// to be received from the network in response to a menu item,
// the menu item might be selected again. It might seem that
// one could make the menu item's callback return immediately
// and do nothing if g_main_depth() returns a value greater than 1.
// However, this should be avoided since the user then sees selecting
// the menu item do nothing. Furthermore, you'll find yourself adding
// these checks all over your code, since there are doubtless many,
// many things that the user could do. Instead, you can use the
// following techniques:
//
//  1. Use gtk_widget_set_sensitive() or modal dialogs to prevent
//     the user from interacting with elements while the main
//     loop is recursing.
//
//  2. Avoid main loop recursion in situations where you can't handle
//     arbitrary  callbacks. Instead, structure your code so that you
//     simply return to the main loop and then get called again when
//     there is more work to do.
func MainDepth() int32 {

	return xMainDepth()

}

var xSourceRemove func(uint) bool

// Removes the source with the given ID from the default main context. You must
// use g_source_destroy() for sources added to a non-default main context.
//
// The ID of a #GSource is given by g_source_get_id(), or will be
// returned by the functions g_source_attach(), g_idle_add(),
// g_idle_add_full(), g_timeout_add(), g_timeout_add_full(),
// g_child_watch_add(), g_child_watch_add_full(), g_io_add_watch(), and
// g_io_add_watch_full().
//
// It is a programmer error to attempt to remove a non-existent source.
//
// More specifically: source IDs can be reissued after a source has been
// destroyed and therefore it is never valid to use this function with a
// source ID which may have already been removed.  An example is when
// scheduling an idle to run in another thread with g_idle_add(): the
// idle may already have run and been removed by the time this function
// is called on its (now invalid) source ID.  This source ID may have
// been reissued, leading to the operation being performed against the
// wrong source.
func SourceRemove(TagVar uint) bool {

	return xSourceRemove(TagVar)

}

var xSourceRemoveByFuncsUserData func(*SourceFuncs, uintptr) bool

// Removes a source from the default main loop context given the
// source functions and user data. If multiple sources exist with the
// same source functions and user data, only one will be destroyed.
func SourceRemoveByFuncsUserData(FuncsVar *SourceFuncs, UserDataVar uintptr) bool {

	return xSourceRemoveByFuncsUserData(FuncsVar, UserDataVar)

}

var xSourceRemoveByUserData func(uintptr) bool

// Removes a source from the default main loop context given the user
// data for the callback. If multiple sources exist with the same user
// data, only one will be destroyed.
func SourceRemoveByUserData(UserDataVar uintptr) bool {

	return xSourceRemoveByUserData(UserDataVar)

}

var xSourceSetNameById func(uint, string)

// Sets the name of a source using its ID.
//
// This is a convenience utility to set source names from the return
// value of g_idle_add(), g_timeout_add(), etc.
//
// It is a programmer error to attempt to set the name of a non-existent
// source.
//
// More specifically: source IDs can be reissued after a source has been
// destroyed and therefore it is never valid to use this function with a
// source ID which may have already been removed.  An example is when
// scheduling an idle to run in another thread with g_idle_add(): the
// idle may already have run and been removed by the time this function
// is called on its (now invalid) source ID.  This source ID may have
// been reissued, leading to the operation being performed against the
// wrong source.
func SourceSetNameById(TagVar uint, NameVar string) {

	xSourceSetNameById(TagVar, NameVar)

}

var xTimeoutAdd func(uint, uintptr, uintptr) uint

// Sets a function to be called at regular intervals, with the default
// priority, %G_PRIORITY_DEFAULT.
//
// The given @function is called repeatedly until it returns %G_SOURCE_REMOVE
// or %FALSE, at which point the timeout is automatically destroyed and the
// function will not be called again. The first call to the function will be
// at the end of the first @interval.
//
// Note that timeout functions may be delayed, due to the processing of other
// event sources. Thus they should not be relied on for precise timing.
// After each call to the timeout function, the time of the next
// timeout is recalculated based on the current time and the given interval
// (it does not try to 'catch up' time lost in delays).
//
// See [memory management of sources][mainloop-memory-management] for details
// on how to handle the return value and memory management of @data.
//
// If you want to have a timer in the "seconds" range and do not care
// about the exact time of the first call of the timer, use the
// g_timeout_add_seconds() function; this function allows for more
// optimizations and more efficient system power usage.
//
// This internally creates a main loop source using g_timeout_source_new()
// and attaches it to the global #GMainContext using g_source_attach(), so
// the callback will be invoked in whichever thread is running that main
// context. You can do these steps manually if you need greater control or to
// use a custom main context.
//
// It is safe to call this function from any thread.
//
// The interval given is in terms of monotonic time, not wall clock
// time.  See g_get_monotonic_time().
func TimeoutAdd(IntervalVar uint, FunctionVar SourceFunc, DataVar uintptr) uint {

	return xTimeoutAdd(IntervalVar, purego.NewCallback(FunctionVar), DataVar)

}

var xTimeoutAddFull func(int32, uint, uintptr, uintptr, uintptr) uint

// Sets a function to be called at regular intervals, with the given
// priority.  The function is called repeatedly until it returns
// %FALSE, at which point the timeout is automatically destroyed and
// the function will not be called again.  The @notify function is
// called when the timeout is destroyed.  The first call to the
// function will be at the end of the first @interval.
//
// Note that timeout functions may be delayed, due to the processing of other
// event sources. Thus they should not be relied on for precise timing.
// After each call to the timeout function, the time of the next
// timeout is recalculated based on the current time and the given interval
// (it does not try to 'catch up' time lost in delays).
//
// See [memory management of sources][mainloop-memory-management] for details
// on how to handle the return value and memory management of @data.
//
// This internally creates a main loop source using g_timeout_source_new()
// and attaches it to the global #GMainContext using g_source_attach(), so
// the callback will be invoked in whichever thread is running that main
// context. You can do these steps manually if you need greater control or to
// use a custom main context.
//
// The interval given is in terms of monotonic time, not wall clock time.
// See g_get_monotonic_time().
func TimeoutAddFull(PriorityVar int32, IntervalVar uint, FunctionVar SourceFunc, DataVar uintptr, NotifyVar DestroyNotify) uint {

	return xTimeoutAddFull(PriorityVar, IntervalVar, purego.NewCallback(FunctionVar), DataVar, purego.NewCallback(NotifyVar))

}

var xTimeoutAddSeconds func(uint, uintptr, uintptr) uint

// Sets a function to be called at regular intervals with the default
// priority, %G_PRIORITY_DEFAULT.
//
// The function is called repeatedly until it returns %G_SOURCE_REMOVE
// or %FALSE, at which point the timeout is automatically destroyed
// and the function will not be called again.
//
// This internally creates a main loop source using
// g_timeout_source_new_seconds() and attaches it to the main loop context
// using g_source_attach(). You can do these steps manually if you need
// greater control. Also see g_timeout_add_seconds_full().
//
// It is safe to call this function from any thread.
//
// Note that the first call of the timer may not be precise for timeouts
// of one second. If you need finer precision and have such a timeout,
// you may want to use g_timeout_add() instead.
//
// See [memory management of sources][mainloop-memory-management] for details
// on how to handle the return value and memory management of @data.
//
// The interval given is in terms of monotonic time, not wall clock
// time.  See g_get_monotonic_time().
func TimeoutAddSeconds(IntervalVar uint, FunctionVar SourceFunc, DataVar uintptr) uint {

	return xTimeoutAddSeconds(IntervalVar, purego.NewCallback(FunctionVar), DataVar)

}

var xTimeoutAddSecondsFull func(int32, uint, uintptr, uintptr, uintptr) uint

// Sets a function to be called at regular intervals, with @priority.
//
// The function is called repeatedly until it returns %G_SOURCE_REMOVE
// or %FALSE, at which point the timeout is automatically destroyed and
// the function will not be called again.
//
// Unlike g_timeout_add(), this function operates at whole second granularity.
// The initial starting point of the timer is determined by the implementation
// and the implementation is expected to group multiple timers together so that
// they fire all at the same time. To allow this grouping, the @interval to the
// first timer is rounded and can deviate up to one second from the specified
// interval. Subsequent timer iterations will generally run at the specified
// interval.
//
// Note that timeout functions may be delayed, due to the processing of other
// event sources. Thus they should not be relied on for precise timing.
// After each call to the timeout function, the time of the next
// timeout is recalculated based on the current time and the given @interval
//
// See [memory management of sources][mainloop-memory-management] for details
// on how to handle the return value and memory management of @data.
//
// If you want timing more precise than whole seconds, use g_timeout_add()
// instead.
//
// The grouping of timers to fire at the same time results in a more power
// and CPU efficient behavior so if your timer is in multiples of seconds
// and you don't require the first timer exactly one second from now, the
// use of g_timeout_add_seconds() is preferred over g_timeout_add().
//
// This internally creates a main loop source using
// g_timeout_source_new_seconds() and attaches it to the main loop context
// using g_source_attach(). You can do these steps manually if you need
// greater control.
//
// It is safe to call this function from any thread.
//
// The interval given is in terms of monotonic time, not wall clock
// time.  See g_get_monotonic_time().
func TimeoutAddSecondsFull(PriorityVar int32, IntervalVar uint, FunctionVar SourceFunc, DataVar uintptr, NotifyVar DestroyNotify) uint {

	return xTimeoutAddSecondsFull(PriorityVar, IntervalVar, purego.NewCallback(FunctionVar), DataVar, purego.NewCallback(NotifyVar))

}

var xTimeoutSourceNew func(uint) *Source

// Creates a new timeout source.
//
// The source will not initially be associated with any #GMainContext
// and must be added to one with g_source_attach() before it will be
// executed.
//
// The interval given is in terms of monotonic time, not wall clock
// time.  See g_get_monotonic_time().
func TimeoutSourceNew(IntervalVar uint) *Source {

	return xTimeoutSourceNew(IntervalVar)

}

var xTimeoutSourceNewSeconds func(uint) *Source

// Creates a new timeout source.
//
// The source will not initially be associated with any #GMainContext
// and must be added to one with g_source_attach() before it will be
// executed.
//
// The scheduling granularity/accuracy of this timeout source will be
// in seconds.
//
// The interval given is in terms of monotonic time, not wall clock time.
// See g_get_monotonic_time().
func TimeoutSourceNewSeconds(IntervalVar uint) *Source {

	return xTimeoutSourceNewSeconds(IntervalVar)

}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GLIB"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	core.PuregoSafeRegister(&xChildWatchAdd, lib, "g_child_watch_add")
	core.PuregoSafeRegister(&xChildWatchAddFull, lib, "g_child_watch_add_full")
	core.PuregoSafeRegister(&xChildWatchSourceNew, lib, "g_child_watch_source_new")
	core.PuregoSafeRegister(&xClearHandleId, lib, "g_clear_handle_id")
	core.PuregoSafeRegister(&xGetCurrentTime, lib, "g_get_current_time")
	core.PuregoSafeRegister(&xGetMonotonicTime, lib, "g_get_monotonic_time")
	core.PuregoSafeRegister(&xGetRealTime, lib, "g_get_real_time")
	core.PuregoSafeRegister(&xIdleAdd, lib, "g_idle_add")
	core.PuregoSafeRegister(&xIdleAddFull, lib, "g_idle_add_full")
	core.PuregoSafeRegister(&xIdleRemoveByData, lib, "g_idle_remove_by_data")
	core.PuregoSafeRegister(&xIdleSourceNew, lib, "g_idle_source_new")
	core.PuregoSafeRegister(&xMainContextDefault, lib, "g_main_context_default")
	core.PuregoSafeRegister(&xMainContextGetThreadDefault, lib, "g_main_context_get_thread_default")
	core.PuregoSafeRegister(&xMainContextRefThreadDefault, lib, "g_main_context_ref_thread_default")
	core.PuregoSafeRegister(&xMainCurrentSource, lib, "g_main_current_source")
	core.PuregoSafeRegister(&xMainDepth, lib, "g_main_depth")
	core.PuregoSafeRegister(&xSourceRemove, lib, "g_source_remove")
	core.PuregoSafeRegister(&xSourceRemoveByFuncsUserData, lib, "g_source_remove_by_funcs_user_data")
	core.PuregoSafeRegister(&xSourceRemoveByUserData, lib, "g_source_remove_by_user_data")
	core.PuregoSafeRegister(&xSourceSetNameById, lib, "g_source_set_name_by_id")
	core.PuregoSafeRegister(&xTimeoutAdd, lib, "g_timeout_add")
	core.PuregoSafeRegister(&xTimeoutAddFull, lib, "g_timeout_add_full")
	core.PuregoSafeRegister(&xTimeoutAddSeconds, lib, "g_timeout_add_seconds")
	core.PuregoSafeRegister(&xTimeoutAddSecondsFull, lib, "g_timeout_add_seconds_full")
	core.PuregoSafeRegister(&xTimeoutSourceNew, lib, "g_timeout_source_new")
	core.PuregoSafeRegister(&xTimeoutSourceNewSeconds, lib, "g_timeout_source_new_seconds")

}
