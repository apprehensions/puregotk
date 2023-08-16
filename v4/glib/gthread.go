// Package glib was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package glib

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
)

// Specifies the type of the @func functions passed to g_thread_new()
// or g_thread_try_new().
type ThreadFunc func(uintptr) uintptr

// The #GCond struct is an opaque data structure that represents a
// condition. Threads can block on a #GCond if they find a certain
// condition to be false. If other threads change the state of this
// condition they signal the #GCond, and that causes the waiting
// threads to be woken up.
//
// Consider the following example of a shared variable.  One or more
// threads can wait for data to be published to the variable and when
// another thread publishes the data, it can signal one of the waiting
// threads to wake up to collect the data.
//
// Here is an example for using GCond to block a thread until a condition
// is satisfied:
// |[&lt;!-- language="C" --&gt;
//
//	gpointer current_data = NULL;
//	GMutex data_mutex;
//	GCond data_cond;
//
//	void
//	push_data (gpointer data)
//	{
//	  g_mutex_lock (&amp;data_mutex);
//	  current_data = data;
//	  g_cond_signal (&amp;data_cond);
//	  g_mutex_unlock (&amp;data_mutex);
//	}
//
//	gpointer
//	pop_data (void)
//	{
//	  gpointer data;
//
//	  g_mutex_lock (&amp;data_mutex);
//	  while (!current_data)
//	    g_cond_wait (&amp;data_cond, &amp;data_mutex);
//	  data = current_data;
//	  current_data = NULL;
//	  g_mutex_unlock (&amp;data_mutex);
//
//	  return data;
//	}
//
// ]|
// Whenever a thread calls pop_data() now, it will wait until
// current_data is non-%NULL, i.e. until some other thread
// has called push_data().
//
// The example shows that use of a condition variable must always be
// paired with a mutex.  Without the use of a mutex, there would be a
// race between the check of @current_data by the while loop in
// pop_data() and waiting. Specifically, another thread could set
// @current_data after the check, and signal the cond (with nobody
// waiting on it) before the first thread goes to sleep. #GCond is
// specifically useful for its ability to release the mutex and go
// to sleep atomically.
//
// It is also important to use the g_cond_wait() and g_cond_wait_until()
// functions only inside a loop which checks for the condition to be
// true.  See g_cond_wait() for an explanation of why the condition may
// not be true even after it returns.
//
// If a #GCond is allocated in static storage then it can be used
// without initialisation.  Otherwise, you should call g_cond_init()
// on it and g_cond_clear() when done.
//
// A #GCond should only be accessed via the g_cond_ functions.
type Cond struct {
	P uintptr

	I uintptr
}

// A #GOnce struct controls a one-time initialization function. Any
// one-time initialization function must have its own unique #GOnce
// struct.
type Once struct {
	Status OnceStatus

	Retval uintptr
}

// The #GPrivate struct is an opaque data structure to represent a
// thread-local data key. It is approximately equivalent to the
// pthread_setspecific()/pthread_getspecific() APIs on POSIX and to
// TlsSetValue()/TlsGetValue() on Windows.
//
// If you don't already know why you might want this functionality,
// then you probably don't need it.
//
// #GPrivate is a very limited resource (as far as 128 per program,
// shared between all libraries). It is also not possible to destroy a
// #GPrivate after it has been used. As such, it is only ever acceptable
// to use #GPrivate in static scope, and even then sparingly so.
//
// See G_PRIVATE_INIT() for a couple of examples.
//
// The #GPrivate structure should be considered opaque.  It should only
// be accessed via the g_private_ functions.
type Private struct {
	P uintptr

	Notify DestroyNotify

	Future uintptr
}

// The GRWLock struct is an opaque data structure to represent a
// reader-writer lock. It is similar to a #GMutex in that it allows
// multiple threads to coordinate access to a shared resource.
//
// The difference to a mutex is that a reader-writer lock discriminates
// between read-only ('reader') and full ('writer') access. While only
// one thread at a time is allowed write access (by holding the 'writer'
// lock via g_rw_lock_writer_lock()), multiple threads can gain
// simultaneous read-only access (by holding the 'reader' lock via
// g_rw_lock_reader_lock()).
//
// It is unspecified whether readers or writers have priority in acquiring the
// lock when a reader already holds the lock and a writer is queued to acquire
// it.
//
// Here is an example for an array with access functions:
// |[&lt;!-- language="C" --&gt;
//
//	 GRWLock lock;
//	 GPtrArray *array;
//
//	 gpointer
//	 my_array_get (guint index)
//	 {
//	   gpointer retval = NULL;
//
//	   if (!array)
//	     return NULL;
//
//	   g_rw_lock_reader_lock (&amp;lock);
//	   if (index &lt; array-&gt;len)
//	     retval = g_ptr_array_index (array, index);
//	   g_rw_lock_reader_unlock (&amp;lock);
//
//	   return retval;
//	 }
//
//	 void
//	 my_array_set (guint index, gpointer data)
//	 {
//	   g_rw_lock_writer_lock (&amp;lock);
//
//	   if (!array)
//	     array = g_ptr_array_new ();
//
//	   if (index &gt;= array-&gt;len)
//	     g_ptr_array_set_size (array, index+1);
//	   g_ptr_array_index (array, index) = data;
//
//	   g_rw_lock_writer_unlock (&amp;lock);
//	 }
//	]|
//
// This example shows an array which can be accessed by many readers
// (the my_array_get() function) simultaneously, whereas the writers
// (the my_array_set() function) will only be allowed one at a time
// and only if no readers currently access the array. This is because
// of the potentially dangerous resizing of the array. Using these
// functions is fully multi-thread safe now.
//
// If a #GRWLock is allocated in static storage then it can be used
// without initialisation.  Otherwise, you should call
// g_rw_lock_init() on it and g_rw_lock_clear() when done.
//
// A GRWLock should only be accessed with the g_rw_lock_ functions.
type RWLock struct {
	P uintptr

	I uintptr
}

// The GRecMutex struct is an opaque data structure to represent a
// recursive mutex. It is similar to a #GMutex with the difference
// that it is possible to lock a GRecMutex multiple times in the same
// thread without deadlock. When doing so, care has to be taken to
// unlock the recursive mutex as often as it has been locked.
//
// If a #GRecMutex is allocated in static storage then it can be used
// without initialisation.  Otherwise, you should call
// g_rec_mutex_init() on it and g_rec_mutex_clear() when done.
//
// A GRecMutex should only be accessed with the
// g_rec_mutex_ functions.
type RecMutex struct {
	P uintptr

	I uintptr
}

// The #GThread struct represents a running thread. This struct
// is returned by g_thread_new() or g_thread_try_new(). You can
// obtain the #GThread struct representing the current thread by
// calling g_thread_self().
//
// GThread is refcounted, see g_thread_ref() and g_thread_unref().
// The thread represented by it holds a reference while it is running,
// and g_thread_join() consumes the reference that it is given, so
// it is normally not necessary to manage GThread references
// explicitly.
//
// The structure is opaque -- none of its fields may be directly
// accessed.
type Thread struct {
}

// The #GMutex struct is an opaque data structure to represent a mutex
// (mutual exclusion). It can be used to protect data against shared
// access.
//
// Take for example the following function:
// |[&lt;!-- language="C" --&gt;
//
//	int
//	give_me_next_number (void)
//	{
//	  static int current_number = 0;
//
//	  // now do a very complicated calculation to calculate the new
//	  // number, this might for example be a random number generator
//	  current_number = calc_next_number (current_number);
//
//	  return current_number;
//	}
//
// ]|
// It is easy to see that this won't work in a multi-threaded
// application. There current_number must be protected against shared
// access. A #GMutex can be used as a solution to this problem:
// |[&lt;!-- language="C" --&gt;
//
//	int
//	give_me_next_number (void)
//	{
//	  static GMutex mutex;
//	  static int current_number = 0;
//	  int ret_val;
//
//	  g_mutex_lock (&amp;mutex);
//	  ret_val = current_number = calc_next_number (current_number);
//	  g_mutex_unlock (&amp;mutex);
//
//	  return ret_val;
//	}
//
// ]|
// Notice that the #GMutex is not initialised to any particular value.
// Its placement in static storage ensures that it will be initialised
// to all-zeros, which is appropriate.
//
// If a #GMutex is placed in other contexts (eg: embedded in a struct)
// then it must be explicitly initialised using g_mutex_init().
//
// A #GMutex should only be accessed via g_mutex_ functions.
type Mutex = uintptr

// Opaque type. See g_mutex_locker_new() for details.
type MutexLocker = uintptr

// Opaque type. See g_rw_lock_reader_locker_new() for details.
type RWLockReaderLocker = uintptr

// Opaque type. See g_rw_lock_writer_locker_new() for details.
type RWLockWriterLocker = uintptr

// Opaque type. See g_rec_mutex_locker_new() for details.
type RecMutexLocker = uintptr

// The possible statuses of a one-time initialization function
// controlled by a #GOnce struct.
type OnceStatus int

const (

	// the function has not been called yet.
	GOnceStatusNotcalledValue OnceStatus = 0
	// the function call is currently in progress.
	GOnceStatusProgressValue OnceStatus = 1
	// the function has been called.
	GOnceStatusReadyValue OnceStatus = 2
)

// Possible errors of thread related functions.
type ThreadError int

const (

	// a thread couldn't be created due to resource
	//                        shortage. Try again later.
	GThreadErrorAgainValue ThreadError = 0
)

var xGetNumProcessors func() uint

// Determine the approximate number of threads that the system will
// schedule simultaneously for this process.  This is intended to be
// used as a parameter to g_thread_pool_new() for CPU bound tasks and
// similar cases.
func GetNumProcessors() uint {

	cret := xGetNumProcessors()
	return cret
}

var xOnceInitEnter func(uintptr) bool

// Function to be called when starting a critical initialization
// section. The argument @location must point to a static
// 0-initialized variable that will be set to a value other than 0 at
// the end of the initialization section. In combination with
// g_once_init_leave() and the unique address @value_location, it can
// be ensured that an initialization section will be executed only once
// during a program's life time, and that concurrent threads are
// blocked until initialization completed. To be used in constructs
// like this:
//
// |[&lt;!-- language="C" --&gt;
//
//	static gsize initialization_value = 0;
//
//	if (g_once_init_enter (&amp;initialization_value))
//	  {
//	    gsize setup_value = 42; // initialization code here
//
//	    g_once_init_leave (&amp;initialization_value, setup_value);
//	  }
//
//	// use initialization_value here
//
// ]|
//
// While @location has a `volatile` qualifier, this is a historical artifact and
// the pointer passed to it should not be `volatile`.
func OnceInitEnter(LocationVar uintptr) bool {

	cret := xOnceInitEnter(LocationVar)
	return cret
}

var xOnceInitLeave func(uintptr, uint)

// Counterpart to g_once_init_enter(). Expects a location of a static
// 0-initialized initialization variable, and an initialization value
// other than 0. Sets the variable to the initialization value, and
// releases concurrent threads blocking in g_once_init_enter() on this
// initialization variable.
//
// While @location has a `volatile` qualifier, this is a historical artifact and
// the pointer passed to it should not be `volatile`.
func OnceInitLeave(LocationVar uintptr, ResultVar uint) {

	xOnceInitLeave(LocationVar, ResultVar)

}

var xThreadExit func(uintptr)

// Terminates the current thread.
//
// If another thread is waiting for us using g_thread_join() then the
// waiting thread will be woken up and get @retval as the return value
// of g_thread_join().
//
// Calling g_thread_exit() with a parameter @retval is equivalent to
// returning @retval from the function @func, as given to g_thread_new().
//
// You must only call g_thread_exit() from a thread that you created
// yourself with g_thread_new() or related APIs. You must not call
// this function from a thread created with another threading library
// or or from within a #GThreadPool.
func ThreadExit(RetvalVar uintptr) {

	xThreadExit(RetvalVar)

}

var xThreadSelf func() *Thread

// This function returns the #GThread corresponding to the
// current thread. Note that this function does not increase
// the reference count of the returned struct.
//
// This function will return a #GThread even for threads that
// were not created by GLib (i.e. those created by other threading
// APIs). This may be useful for thread identification purposes
// (i.e. comparisons) but you must not use GLib functions (such
// as g_thread_join()) on these threads.
func ThreadSelf() *Thread {

	cret := xThreadSelf()
	return cret
}

var xThreadYield func()

// Causes the calling thread to voluntarily relinquish the CPU, so
// that other threads can run.
//
// This function is often used as a method to make busy wait less evil.
func ThreadYield() {

	xThreadYield()

}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GLIB"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	core.PuregoSafeRegister(&xGetNumProcessors, lib, "g_get_num_processors")
	core.PuregoSafeRegister(&xOnceInitEnter, lib, "g_once_init_enter")
	core.PuregoSafeRegister(&xOnceInitLeave, lib, "g_once_init_leave")
	core.PuregoSafeRegister(&xThreadExit, lib, "g_thread_exit")
	core.PuregoSafeRegister(&xThreadSelf, lib, "g_thread_self")
	core.PuregoSafeRegister(&xThreadYield, lib, "g_thread_yield")

}
