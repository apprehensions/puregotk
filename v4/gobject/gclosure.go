// Package gobject was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gobject

import (
	"unsafe"

	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/glib"
)

// The type used for callback functions in structure definitions and function
// signatures.
//
// This doesn't mean that all callback functions must take no  parameters and
// return void. The required signature of a callback function is determined by
// the context in which is used (e.g. the signal to which it is connected).
//
// Use G_CALLBACK() to cast the callback function to a #GCallback.
type Callback func()

// The type used for marshaller functions.
type ClosureMarshal func(*Closure, *Value, uint, uintptr, uintptr, uintptr)

// The type used for the various notification callbacks which can be registered
// on closures.
type ClosureNotify func(uintptr, *Closure)

// This is the signature of va_list marshaller functions, an optional
// marshaller that can be used in some situations to avoid
// marshalling the signal argument into GValues.
type VaClosureMarshal func(*Closure, *Value, *TypeInstance, []interface{}, uintptr, int, uintptr)

// A #GCClosure is a specialization of #GClosure for C function callbacks.
type CClosure struct {
	Closure uintptr

	Callback uintptr
}

func (x *CClosure) GoPointer() uintptr {
	return uintptr(unsafe.Pointer(x))
}

// A #GClosure represents a callback supplied by the programmer.
//
// It will generally comprise a function of some kind and a marshaller
// used to call it. It is the responsibility of the marshaller to
// convert the arguments for the invocation from #GValues into
// a suitable form, perform the callback on the converted arguments,
// and transform the return value back into a #GValue.
//
// In the case of C programs, a closure usually just holds a pointer
// to a function and maybe a data argument, and the marshaller
// converts between #GValue and native C types. The GObject
// library provides the #GCClosure type for this purpose. Bindings for
// other languages need marshallers which convert between #GValues
// and suitable representations in the runtime of the language in
// order to use functions written in that language as callbacks. Use
// g_closure_set_marshal() to set the marshaller on such a custom
// closure implementation.
//
// Within GObject, closures play an important role in the
// implementation of signals. When a signal is registered, the
// @c_marshaller argument to g_signal_new() specifies the default C
// marshaller for any closure which is connected to this
// signal. GObject provides a number of C marshallers for this
// purpose, see the g_cclosure_marshal_*() functions. Additional C
// marshallers can be generated with the [glib-genmarshal][glib-genmarshal]
// utility.  Closures can be explicitly connected to signals with
// g_signal_connect_closure(), but it usually more convenient to let
// GObject create a closure automatically by using one of the
// g_signal_connect_*() functions which take a callback function/user
// data pair.
//
// Using closures has a number of important advantages over a simple
// callback function/data pointer combination:
//
//   - Closures allow the callee to get the types of the callback parameters,
//     which means that language bindings don't have to write individual glue
//     for each callback type.
//
//   - The reference counting of #GClosure makes it easy to handle reentrancy
//     right; if a callback is removed while it is being invoked, the closure
//     and its parameters won't be freed until the invocation finishes.
//
//   - g_closure_invalidate() and invalidation notifiers allow callbacks to be
//     automatically removed when the objects they point to go away.
type Closure struct {
	RefCount uint

	MetaMarshalNouse uint

	NGuards uint

	NFnotifiers uint

	NInotifiers uint

	InInotify uint

	Floating uint

	DerivativeFlag uint

	InMarshal uint

	IsInvalid uint

	Data uintptr

	Notifiers *ClosureNotifyData
}

func (x *Closure) GoPointer() uintptr {
	return uintptr(unsafe.Pointer(x))
}

var xNewClosureObject func(uint, uintptr) *Closure

// A variant of g_closure_new_simple() which stores @object in the
// @data field of the closure and calls g_object_watch_closure() on
// @object and the created closure. This function is mainly useful
// when implementing new types of closures.
func NewClosureObject(SizeofClosureVar uint, ObjectVar *Object) *Closure {

	cret := xNewClosureObject(SizeofClosureVar, ObjectVar.GoPointer())
	return cret
}

var xNewClosureSimple func(uint, uintptr) *Closure

// Allocates a struct of the given size and initializes the initial
// part as a #GClosure.
//
// This function is mainly useful when implementing new types of closures:
//
// |[&lt;!-- language="C" --&gt;
// typedef struct _MyClosure MyClosure;
// struct _MyClosure
//
//	{
//	  GClosure closure;
//	  // extra data goes here
//	};
//
// static void
// my_closure_finalize (gpointer  notify_data,
//
//	GClosure *closure)
//
//	{
//	  MyClosure *my_closure = (MyClosure *)closure;
//
//	  // free extra data here
//	}
//
// MyClosure *my_closure_new (gpointer data)
//
//	{
//	  GClosure *closure;
//	  MyClosure *my_closure;
//
//	  closure = g_closure_new_simple (sizeof (MyClosure), data);
//	  my_closure = (MyClosure *) closure;
//
//	  // initialize extra data here
//
//	  g_closure_add_finalize_notifier (closure, notify_data,
//	                                   my_closure_finalize);
//	  return my_closure;
//	}
//
// ]|
func NewClosureSimple(SizeofClosureVar uint, DataVar uintptr) *Closure {

	cret := xNewClosureSimple(SizeofClosureVar, DataVar)
	return cret
}

var xClosureAddFinalizeNotifier func(uintptr, uintptr, uintptr)

// Registers a finalization notifier which will be called when the
// reference count of @closure goes down to 0.
//
// Multiple finalization notifiers on a single closure are invoked in
// unspecified order. If a single call to g_closure_unref() results in
// the closure being both invalidated and finalized, then the invalidate
// notifiers will be run before the finalize notifiers.
func (x *Closure) AddFinalizeNotifier(NotifyDataVar uintptr, NotifyFuncVar *ClosureNotify) {

	xClosureAddFinalizeNotifier(x.GoPointer(), NotifyDataVar, glib.NewCallback(NotifyFuncVar))

}

var xClosureAddInvalidateNotifier func(uintptr, uintptr, uintptr)

// Registers an invalidation notifier which will be called when the
// @closure is invalidated with g_closure_invalidate().
//
// Invalidation notifiers are invoked before finalization notifiers,
// in an unspecified order.
func (x *Closure) AddInvalidateNotifier(NotifyDataVar uintptr, NotifyFuncVar *ClosureNotify) {

	xClosureAddInvalidateNotifier(x.GoPointer(), NotifyDataVar, glib.NewCallback(NotifyFuncVar))

}

var xClosureAddMarshalGuards func(uintptr, uintptr, uintptr, uintptr, uintptr)

// Adds a pair of notifiers which get invoked before and after the
// closure callback, respectively.
//
// This is typically used to protect the extra arguments for the
// duration of the callback. See g_object_watch_closure() for an
// example of marshal guards.
func (x *Closure) AddMarshalGuards(PreMarshalDataVar uintptr, PreMarshalNotifyVar *ClosureNotify, PostMarshalDataVar uintptr, PostMarshalNotifyVar *ClosureNotify) {

	xClosureAddMarshalGuards(x.GoPointer(), PreMarshalDataVar, glib.NewCallback(PreMarshalNotifyVar), PostMarshalDataVar, glib.NewCallback(PostMarshalNotifyVar))

}

var xClosureInvalidate func(uintptr)

// Sets a flag on the closure to indicate that its calling
// environment has become invalid, and thus causes any future
// invocations of g_closure_invoke() on this @closure to be
// ignored.
//
// Also, invalidation notifiers installed on the closure will
// be called at this point. Note that unless you are holding a
// reference to the closure yourself, the invalidation notifiers may
// unref the closure and cause it to be destroyed, so if you need to
// access the closure after calling g_closure_invalidate(), make sure
// that you've previously called g_closure_ref().
//
// Note that g_closure_invalidate() will also be called when the
// reference count of a closure drops to zero (unless it has already
// been invalidated before).
func (x *Closure) Invalidate() {

	xClosureInvalidate(x.GoPointer())

}

var xClosureInvoke func(uintptr, *Value, uint, uintptr, uintptr)

// Invokes the closure, i.e. executes the callback represented by the @closure.
func (x *Closure) Invoke(ReturnValueVar *Value, NParamValuesVar uint, ParamValuesVar uintptr, InvocationHintVar uintptr) {

	xClosureInvoke(x.GoPointer(), ReturnValueVar, NParamValuesVar, ParamValuesVar, InvocationHintVar)

}

var xClosureRef func(uintptr) *Closure

// Increments the reference count on a closure to force it staying
// alive while the caller holds a pointer to it.
func (x *Closure) Ref() *Closure {

	cret := xClosureRef(x.GoPointer())
	return cret
}

var xClosureRemoveFinalizeNotifier func(uintptr, uintptr, uintptr)

// Removes a finalization notifier.
//
// Notice that notifiers are automatically removed after they are run.
func (x *Closure) RemoveFinalizeNotifier(NotifyDataVar uintptr, NotifyFuncVar *ClosureNotify) {

	xClosureRemoveFinalizeNotifier(x.GoPointer(), NotifyDataVar, glib.NewCallback(NotifyFuncVar))

}

var xClosureRemoveInvalidateNotifier func(uintptr, uintptr, uintptr)

// Removes an invalidation notifier.
//
// Notice that notifiers are automatically removed after they are run.
func (x *Closure) RemoveInvalidateNotifier(NotifyDataVar uintptr, NotifyFuncVar *ClosureNotify) {

	xClosureRemoveInvalidateNotifier(x.GoPointer(), NotifyDataVar, glib.NewCallback(NotifyFuncVar))

}

var xClosureSetMarshal func(uintptr, uintptr)

// Sets the marshaller of @closure.
//
// The `marshal_data` of @marshal provides a way for a meta marshaller to
// provide additional information to the marshaller.
//
// For GObject's C predefined marshallers (the `g_cclosure_marshal_*()`
// functions), what it provides is a callback function to use instead of
// @closure-&gt;callback.
//
// See also: g_closure_set_meta_marshal()
func (x *Closure) SetMarshal(MarshalVar *ClosureMarshal) {

	xClosureSetMarshal(x.GoPointer(), glib.NewCallback(MarshalVar))

}

var xClosureSetMetaMarshal func(uintptr, uintptr, uintptr)

// Sets the meta marshaller of @closure.
//
// A meta marshaller wraps the @closure's marshal and modifies the way
// it is called in some fashion. The most common use of this facility
// is for C callbacks.
//
// The same marshallers (generated by [glib-genmarshal][glib-genmarshal]),
// are used everywhere, but the way that we get the callback function
// differs. In most cases we want to use the @closure's callback, but in
// other cases we want to use some different technique to retrieve the
// callback function.
//
// For example, class closures for signals (see
// g_signal_type_cclosure_new()) retrieve the callback function from a
// fixed offset in the class structure.  The meta marshaller retrieves
// the right callback and passes it to the marshaller as the
// @marshal_data argument.
func (x *Closure) SetMetaMarshal(MarshalDataVar uintptr, MetaMarshalVar *ClosureMarshal) {

	xClosureSetMetaMarshal(x.GoPointer(), MarshalDataVar, glib.NewCallback(MetaMarshalVar))

}

var xClosureSink func(uintptr)

// Takes over the initial ownership of a closure.
//
// Each closure is initially created in a "floating" state, which means
// that the initial reference count is not owned by any caller.
//
// This function checks to see if the object is still floating, and if so,
// unsets the floating state and decreases the reference count. If the
// closure is not floating, g_closure_sink() does nothing.
//
// The reason for the existence of the floating state is to prevent
// cumbersome code sequences like:
//
// |[&lt;!-- language="C" --&gt;
// closure = g_cclosure_new (cb_func, cb_data);
// g_source_set_closure (source, closure);
// g_closure_unref (closure); // GObject doesn't really need this
// ]|
//
// Because g_source_set_closure() (and similar functions) take ownership of the
// initial reference count, if it is unowned, we instead can write:
//
// |[&lt;!-- language="C" --&gt;
// g_source_set_closure (source, g_cclosure_new (cb_func, cb_data));
// ]|
//
// Generally, this function is used together with g_closure_ref(). An example
// of storing a closure for later notification looks like:
//
// |[&lt;!-- language="C" --&gt;
// static GClosure *notify_closure = NULL;
// void
// foo_notify_set_closure (GClosure *closure)
//
//	{
//	  if (notify_closure)
//	    g_closure_unref (notify_closure);
//	  notify_closure = closure;
//	  if (notify_closure)
//	    {
//	      g_closure_ref (notify_closure);
//	      g_closure_sink (notify_closure);
//	    }
//	}
//
// ]|
//
// Because g_closure_sink() may decrement the reference count of a closure
// (if it hasn't been called on @closure yet) just like g_closure_unref(),
// g_closure_ref() should be called prior to this function.
func (x *Closure) Sink() {

	xClosureSink(x.GoPointer())

}

var xClosureUnref func(uintptr)

// Decrements the reference count of a closure after it was previously
// incremented by the same caller.
//
// If no other callers are using the closure, then the closure will be
// destroyed and freed.
func (x *Closure) Unref() {

	xClosureUnref(x.GoPointer())

}

type ClosureNotifyData struct {
	Data uintptr

	Notify ClosureNotify
}

func (x *ClosureNotifyData) GoPointer() uintptr {
	return uintptr(unsafe.Pointer(x))
}

var xCclosureMarshalGeneric func(*Closure, *Value, uint, *Value, uintptr, uintptr)

// A generic marshaller function implemented via
// [libffi](http://sourceware.org/libffi/).
//
// Normally this function is not passed explicitly to g_signal_new(),
// but used automatically by GLib when specifying a %NULL marshaller.
func CclosureMarshalGeneric(ClosureVar *Closure, ReturnGvalueVar *Value, NParamValuesVar uint, ParamValuesVar *Value, InvocationHintVar uintptr, MarshalDataVar uintptr) {

	xCclosureMarshalGeneric(ClosureVar, ReturnGvalueVar, NParamValuesVar, ParamValuesVar, InvocationHintVar, MarshalDataVar)

}

var xCclosureNew func(uintptr, uintptr, uintptr) *Closure

// Creates a new closure which invokes @callback_func with @user_data as
// the last parameter.
//
// @destroy_data will be called as a finalize notifier on the #GClosure.
func CclosureNew(CallbackFuncVar *Callback, UserDataVar uintptr, DestroyDataVar *ClosureNotify) *Closure {

	cret := xCclosureNew(glib.NewCallback(CallbackFuncVar), UserDataVar, glib.NewCallback(DestroyDataVar))
	return cret
}

var xCclosureNewSwap func(uintptr, uintptr, uintptr) *Closure

// Creates a new closure which invokes @callback_func with @user_data as
// the first parameter.
//
// @destroy_data will be called as a finalize notifier on the #GClosure.
func CclosureNewSwap(CallbackFuncVar *Callback, UserDataVar uintptr, DestroyDataVar *ClosureNotify) *Closure {

	cret := xCclosureNewSwap(glib.NewCallback(CallbackFuncVar), UserDataVar, glib.NewCallback(DestroyDataVar))
	return cret
}

var xSignalTypeCclosureNew func([]interface{}, uint) *Closure

// Creates a new closure which invokes the function found at the offset
// @struct_offset in the class structure of the interface or classed type
// identified by @itype.
func SignalTypeCclosureNew(ItypeVar []interface{}, StructOffsetVar uint) *Closure {

	cret := xSignalTypeCclosureNew(ItypeVar, StructOffsetVar)
	return cret
}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GOBJECT"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	core.PuregoSafeRegister(&xCclosureMarshalGeneric, lib, "g_cclosure_marshal_generic")
	core.PuregoSafeRegister(&xCclosureNew, lib, "g_cclosure_new")
	core.PuregoSafeRegister(&xCclosureNewSwap, lib, "g_cclosure_new_swap")
	core.PuregoSafeRegister(&xSignalTypeCclosureNew, lib, "g_signal_type_cclosure_new")

	core.PuregoSafeRegister(&xNewClosureObject, lib, "g_closure_new_object")
	core.PuregoSafeRegister(&xNewClosureSimple, lib, "g_closure_new_simple")

	core.PuregoSafeRegister(&xClosureAddFinalizeNotifier, lib, "g_closure_add_finalize_notifier")
	core.PuregoSafeRegister(&xClosureAddInvalidateNotifier, lib, "g_closure_add_invalidate_notifier")
	core.PuregoSafeRegister(&xClosureAddMarshalGuards, lib, "g_closure_add_marshal_guards")
	core.PuregoSafeRegister(&xClosureInvalidate, lib, "g_closure_invalidate")
	core.PuregoSafeRegister(&xClosureInvoke, lib, "g_closure_invoke")
	core.PuregoSafeRegister(&xClosureRef, lib, "g_closure_ref")
	core.PuregoSafeRegister(&xClosureRemoveFinalizeNotifier, lib, "g_closure_remove_finalize_notifier")
	core.PuregoSafeRegister(&xClosureRemoveInvalidateNotifier, lib, "g_closure_remove_invalidate_notifier")
	core.PuregoSafeRegister(&xClosureSetMarshal, lib, "g_closure_set_marshal")
	core.PuregoSafeRegister(&xClosureSetMetaMarshal, lib, "g_closure_set_meta_marshal")
	core.PuregoSafeRegister(&xClosureSink, lib, "g_closure_sink")
	core.PuregoSafeRegister(&xClosureUnref, lib, "g_closure_unref")

}
