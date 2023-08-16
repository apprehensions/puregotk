// Package gtk was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gtk

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/glib"
	"github.com/jwijenbergh/puregotk/v4/gobject"
)

// Callback called by gtk_expression_watch() when the
// expression value changes.
type ExpressionNotify func(uintptr)

// An opaque structure representing a watched `GtkExpression`.
//
// The contents of `GtkExpressionWatch` should only be accessed through the
// provided API.
type ExpressionWatch struct {
}

var xNewParamSpecExpression func(string, string, string, gobject.ParamFlags) uintptr

// Creates a new `GParamSpec` instance for a property holding a `GtkExpression`.
//
// See `g_param_spec_internal()` for details on the property strings.
func NewParamSpecExpression(NameVar string, NickVar string, BlurbVar string, FlagsVar gobject.ParamFlags) *gobject.ParamSpec {
	var cls *gobject.ParamSpec

	cret := xNewParamSpecExpression(NameVar, NickVar, BlurbVar, FlagsVar)

	if cret == 0 {
		return cls
	}
	cls = &gobject.ParamSpec{}
	cls.Ptr = cret
	return cls
}

var xValueDupExpression func(*gobject.Value) uintptr

// Retrieves the `GtkExpression` stored inside the given `value`, and acquires
// a reference to it.
func ValueDupExpression(ValueVar *gobject.Value) *Expression {
	var cls *Expression

	cret := xValueDupExpression(ValueVar)

	if cret == 0 {
		return cls
	}
	cls = &Expression{}
	cls.Ptr = cret
	return cls
}

var xValueGetExpression func(*gobject.Value) uintptr

// Retrieves the `GtkExpression` stored inside the given `value`.
func ValueGetExpression(ValueVar *gobject.Value) *Expression {
	var cls *Expression

	cret := xValueGetExpression(ValueVar)

	if cret == 0 {
		return cls
	}
	gobject.IncreaseRef(cret)
	cls = &Expression{}
	cls.Ptr = cret
	return cls
}

var xValueSetExpression func(*gobject.Value, uintptr)

// Stores the given `GtkExpression` inside `value`.
//
// The `GValue` will acquire a reference to the `expression`.
func ValueSetExpression(ValueVar *gobject.Value, ExpressionVar *Expression) {

	xValueSetExpression(ValueVar, ExpressionVar.GoPointer())

}

var xValueTakeExpression func(*gobject.Value, uintptr)

// Stores the given `GtkExpression` inside `value`.
//
// This function transfers the ownership of the `expression` to the `GValue`.
func ValueTakeExpression(ValueVar *gobject.Value, ExpressionVar *Expression) {

	xValueTakeExpression(ValueVar, ExpressionVar.GoPointer())

}

// A variant of `GtkClosureExpression` using a C closure.
type CClosureExpression struct {
	Expression
}

func CClosureExpressionNewFromInternalPtr(ptr uintptr) *CClosureExpression {
	cls := &CClosureExpression{}
	cls.Ptr = ptr
	return cls
}

var xNewCClosureExpression func([]interface{}, uintptr, uint, uintptr, uintptr, uintptr, uintptr) uintptr

// Creates a `GtkExpression` that calls `callback_func` when it is evaluated.
//
// This function is a variant of [ctor@Gtk.ClosureExpression.new] that
// creates a `GClosure` by calling g_cclosure_new() with the given
// `callback_func`, `user_data` and `user_destroy`.
func NewCClosureExpression(ValueTypeVar []interface{}, MarshalVar gobject.ClosureMarshal, NParamsVar uint, ParamsVar uintptr, CallbackFuncVar gobject.Callback, UserDataVar uintptr, UserDestroyVar gobject.ClosureNotify) *CClosureExpression {
	var cls *CClosureExpression

	cret := xNewCClosureExpression(ValueTypeVar, purego.NewCallback(MarshalVar), NParamsVar, ParamsVar, purego.NewCallback(CallbackFuncVar), UserDataVar, purego.NewCallback(UserDestroyVar))

	if cret == 0 {
		return cls
	}
	cls = &CClosureExpression{}
	cls.Ptr = cret
	return cls
}

func (c *CClosureExpression) GoPointer() uintptr {
	return c.Ptr
}

func (c *CClosureExpression) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// An expression using a custom `GClosure` to compute the value from
// its parameters.
type ClosureExpression struct {
	Expression
}

func ClosureExpressionNewFromInternalPtr(ptr uintptr) *ClosureExpression {
	cls := &ClosureExpression{}
	cls.Ptr = ptr
	return cls
}

var xNewClosureExpression func([]interface{}, *gobject.Closure, uint, uintptr) uintptr

// Creates a `GtkExpression` that calls `closure` when it is evaluated.
//
// `closure` is called with the `this` object and the results of evaluating
// the `params` expressions.
func NewClosureExpression(ValueTypeVar []interface{}, ClosureVar *gobject.Closure, NParamsVar uint, ParamsVar uintptr) *ClosureExpression {
	var cls *ClosureExpression

	cret := xNewClosureExpression(ValueTypeVar, ClosureVar, NParamsVar, ParamsVar)

	if cret == 0 {
		return cls
	}
	cls = &ClosureExpression{}
	cls.Ptr = cret
	return cls
}

func (c *ClosureExpression) GoPointer() uintptr {
	return c.Ptr
}

func (c *ClosureExpression) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// A constant value in a `GtkExpression`.
type ConstantExpression struct {
	Expression
}

func ConstantExpressionNewFromInternalPtr(ptr uintptr) *ConstantExpression {
	cls := &ConstantExpression{}
	cls.Ptr = ptr
	return cls
}

var xNewConstantExpression func([]interface{}, ...interface{}) uintptr

// Creates a `GtkExpression` that evaluates to the
// object given by the arguments.
func NewConstantExpression(ValueTypeVar []interface{}, varArgs ...interface{}) *ConstantExpression {
	var cls *ConstantExpression

	cret := xNewConstantExpression(ValueTypeVar, varArgs...)

	if cret == 0 {
		return cls
	}
	cls = &ConstantExpression{}
	cls.Ptr = cret
	return cls
}

var xNewForValueConstantExpression func(*gobject.Value) uintptr

// Creates an expression that always evaluates to the given `value`.
func NewForValueConstantExpression(ValueVar *gobject.Value) *ConstantExpression {
	var cls *ConstantExpression

	cret := xNewForValueConstantExpression(ValueVar)

	if cret == 0 {
		return cls
	}
	cls = &ConstantExpression{}
	cls.Ptr = cret
	return cls
}

var xConstantExpressionGetValue func(uintptr) *gobject.Value

// Gets the value that a constant expression evaluates to.
func (x *ConstantExpression) GetValue() *gobject.Value {

	cret := xConstantExpressionGetValue(x.GoPointer())
	return cret
}

func (c *ConstantExpression) GoPointer() uintptr {
	return c.Ptr
}

func (c *ConstantExpression) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// `GtkExpression` provides a way to describe references to values.
//
// An important aspect of expressions is that the value can be obtained
// from a source that is several steps away. For example, an expression
// may describe ‘the value of property A of `object1`, which is itself the
// value of a property of `object2`’. And `object1` may not even exist yet
// at the time that the expression is created. This is contrast to `GObject`
// property bindings, which can only create direct connections between
// the properties of two objects that must both exist for the duration
// of the binding.
//
// An expression needs to be "evaluated" to obtain the value that it currently
// refers to. An evaluation always happens in the context of a current object
// called `this` (it mirrors the behavior of object-oriented languages),
// which may or may not influence the result of the evaluation. Use
// [method@Gtk.Expression.evaluate] for evaluating an expression.
//
// Various methods for defining expressions exist, from simple constants via
// [ctor@Gtk.ConstantExpression.new] to looking up properties in a `GObject`
// (even recursively) via [ctor@Gtk.PropertyExpression.new] or providing
// custom functions to transform and combine expressions via
// [ctor@Gtk.ClosureExpression.new].
//
// Here is an example of a complex expression:
//
// ```c
//
//	color_expr = gtk_property_expression_new (GTK_TYPE_LIST_ITEM,
//	                                          NULL, "item");
//	expression = gtk_property_expression_new (GTK_TYPE_COLOR,
//	                                          color_expr, "name");
//
// ```
//
// when evaluated with `this` being a `GtkListItem`, it will obtain the
// "item" property from the `GtkListItem`, and then obtain the "name" property
// from the resulting object (which is assumed to be of type `GTK_TYPE_COLOR`).
//
// # A more concise way to describe this would be
//
// ```
//
//	this-&gt;item-&gt;name
//
// ```
//
// The most likely place where you will encounter expressions is in the context
// of list models and list widgets using them. For example, `GtkDropDown` is
// evaluating a `GtkExpression` to obtain strings from the items in its model
// that it can then use to match against the contents of its search entry.
// `GtkStringFilter` is using a `GtkExpression` for similar reasons.
//
// By default, expressions are not paying attention to changes and evaluation is
// just a snapshot of the current state at a given time. To get informed about
// changes, an expression needs to be "watched" via a [struct@Gtk.ExpressionWatch],
// which will cause a callback to be called whenever the value of the expression may
// have changed; [method@Gtk.Expression.watch] starts watching an expression, and
// [method@Gtk.ExpressionWatch.unwatch] stops.
//
// Watches can be created for automatically updating the property of an object,
// similar to GObject's `GBinding` mechanism, by using [method@Gtk.Expression.bind].
//
// ## GtkExpression in GObject properties
//
// In order to use a `GtkExpression` as a `GObject` property, you must use the
// [id@gtk_param_spec_expression] when creating a `GParamSpec` to install in the
// `GObject` class being defined; for instance:
//
// ```c
// obj_props[PROP_EXPRESSION] =
//
//	gtk_param_spec_expression ("expression",
//	                           "Expression",
//	                           "The expression used by the widget",
//	                           G_PARAM_READWRITE |
//	                           G_PARAM_STATIC_STRINGS |
//	                           G_PARAM_EXPLICIT_NOTIFY);
//
// ```
//
// When implementing the `GObjectClass.set_property` and `GObjectClass.get_property`
// virtual functions, you must use [id@gtk_value_get_expression], to retrieve the
// stored `GtkExpression` from the `GValue` container, and [id@gtk_value_set_expression],
// to store the `GtkExpression` into the `GValue`; for instance:
//
// ```c
//
//	// in set_property()...
//	case PROP_EXPRESSION:
//	  foo_widget_set_expression (foo, gtk_value_get_expression (value));
//	  break;
//
//	// in get_property()...
//	case PROP_EXPRESSION:
//	  gtk_value_set_expression (value, foo-&gt;expression);
//	  break;
//
// ```
//
// ## GtkExpression in .ui files
//
// `GtkBuilder` has support for creating expressions. The syntax here can be used where
// a `GtkExpression` object is needed like in a `&lt;property&gt;` tag for an expression
// property, or in a `&lt;binding name="property"&gt;` tag to bind a property to an expression.
//
// To create a property expression, use the `&lt;lookup&gt;` element. It can have a `type`
// attribute to specify the object type, and a `name` attribute to specify the property
// to look up. The content of `&lt;lookup&gt;` can either be an element specfiying the expression
// to use the object, or a string that specifies the name of the object to use.
//
// Example:
//
// ```xml
//
//	&lt;lookup name='search'&gt;string_filter&lt;/lookup&gt;
//
// ```
//
// To create a constant expression, use the `&lt;constant&gt;` element. If the type attribute
// is specified, the element content is interpreted as a value of that type. Otherwise,
// it is assumed to be an object. For instance:
//
// ```xml
//
//	&lt;constant&gt;string_filter&lt;/constant&gt;
//	&lt;constant type='gchararray'&gt;Hello, world&lt;/constant&gt;
//
// ```
//
// To create a closure expression, use the `&lt;closure&gt;` element. The `type` and `function`
// attributes specify what function to use for the closure, the content of the element
// contains the expressions for the parameters. For instance:
//
// ```xml
//
//	&lt;closure type='gchararray' function='combine_args_somehow'&gt;
//	  &lt;constant type='gchararray'&gt;File size:&lt;/constant&gt;
//	  &lt;lookup type='GFile' name='size'&gt;myfile&lt;/lookup&gt;
//	&lt;/closure&gt;
//
// ```
type Expression struct {
	Ptr uintptr
}

func ExpressionNewFromInternalPtr(ptr uintptr) *Expression {
	cls := &Expression{}
	cls.Ptr = ptr
	return cls
}

var xExpressionBind func(uintptr, uintptr, string, uintptr) *ExpressionWatch

// Bind `target`'s property named `property` to `self`.
//
// The value that `self` evaluates to is set via `g_object_set()` on
// `target`. This is repeated whenever `self` changes to ensure that
// the object's property stays synchronized with `self`.
//
// If `self`'s evaluation fails, `target`'s `property` is not updated.
// You can ensure that this doesn't happen by using a fallback
// expression.
//
// Note that this function takes ownership of `self`. If you want
// to keep it around, you should [method@Gtk.Expression.ref] it beforehand.
func (x *Expression) Bind(TargetVar *gobject.Object, PropertyVar string, ThisVar *gobject.Object) *ExpressionWatch {

	cret := xExpressionBind(x.GoPointer(), TargetVar.GoPointer(), PropertyVar, ThisVar.GoPointer())
	return cret
}

var xExpressionEvaluate func(uintptr, uintptr, *gobject.Value) bool

// Evaluates the given expression and on success stores the result
// in @value.
//
// The `GType` of `value` will be the type given by
// [method@Gtk.Expression.get_value_type].
//
// It is possible that expressions cannot be evaluated - for example
// when the expression references objects that have been destroyed or
// set to `NULL`. In that case `value` will remain empty and `FALSE`
// will be returned.
func (x *Expression) Evaluate(ThisVar *gobject.Object, ValueVar *gobject.Value) bool {

	cret := xExpressionEvaluate(x.GoPointer(), ThisVar.GoPointer(), ValueVar)
	return cret
}

var xExpressionGetValueType func(uintptr) []interface{}

// Gets the `GType` that this expression evaluates to.
//
// This type is constant and will not change over the lifetime
// of this expression.
func (x *Expression) GetValueType() []interface{} {

	cret := xExpressionGetValueType(x.GoPointer())
	return cret
}

var xExpressionIsStatic func(uintptr) bool

// Checks if the expression is static.
//
// A static expression will never change its result when
// [method@Gtk.Expression.evaluate] is called on it with the same arguments.
//
// That means a call to [method@Gtk.Expression.watch] is not necessary because
// it will never trigger a notify.
func (x *Expression) IsStatic() bool {

	cret := xExpressionIsStatic(x.GoPointer())
	return cret
}

var xExpressionRef func(uintptr) uintptr

// Acquires a reference on the given `GtkExpression`.
func (x *Expression) Ref() *Expression {
	var cls *Expression

	cret := xExpressionRef(x.GoPointer())

	if cret == 0 {
		return cls
	}
	cls = &Expression{}
	cls.Ptr = cret
	return cls
}

var xExpressionUnref func(uintptr)

// Releases a reference on the given `GtkExpression`.
//
// If the reference was the last, the resources associated to the `self` are
// freed.
func (x *Expression) Unref() {

	xExpressionUnref(x.GoPointer())

}

var xExpressionWatch func(uintptr, uintptr, uintptr, uintptr, uintptr) *ExpressionWatch

// Watch the given `expression` for changes.
//
// The @notify function will be called whenever the evaluation of `self`
// may have changed.
//
// GTK cannot guarantee that the evaluation did indeed change when the @notify
// gets invoked, but it guarantees the opposite: When it did in fact change,
// the @notify will be invoked.
func (x *Expression) Watch(ThisVar *gobject.Object, NotifyVar ExpressionNotify, UserDataVar uintptr, UserDestroyVar glib.DestroyNotify) *ExpressionWatch {

	cret := xExpressionWatch(x.GoPointer(), ThisVar.GoPointer(), purego.NewCallback(NotifyVar), UserDataVar, purego.NewCallback(UserDestroyVar))
	return cret
}

func (c *Expression) GoPointer() uintptr {
	return c.Ptr
}

func (c *Expression) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// A `GObject` value in a `GtkExpression`.
type ObjectExpression struct {
	Expression
}

func ObjectExpressionNewFromInternalPtr(ptr uintptr) *ObjectExpression {
	cls := &ObjectExpression{}
	cls.Ptr = ptr
	return cls
}

var xNewObjectExpression func(uintptr) uintptr

// Creates an expression evaluating to the given `object` with a weak reference.
//
// Once the `object` is disposed, it will fail to evaluate.
//
// This expression is meant to break reference cycles.
//
// If you want to keep a reference to `object`, use [ctor@Gtk.ConstantExpression.new].
func NewObjectExpression(ObjectVar *gobject.Object) *ObjectExpression {
	var cls *ObjectExpression

	cret := xNewObjectExpression(ObjectVar.GoPointer())

	if cret == 0 {
		return cls
	}
	cls = &ObjectExpression{}
	cls.Ptr = cret
	return cls
}

var xObjectExpressionGetObject func(uintptr) uintptr

// Gets the object that the expression evaluates to.
func (x *ObjectExpression) GetObject() *gobject.Object {
	var cls *gobject.Object

	cret := xObjectExpressionGetObject(x.GoPointer())

	if cret == 0 {
		return cls
	}
	gobject.IncreaseRef(cret)
	cls = &gobject.Object{}
	cls.Ptr = cret
	return cls
}

func (c *ObjectExpression) GoPointer() uintptr {
	return c.Ptr
}

func (c *ObjectExpression) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// A `GParamSpec` for properties holding a `GtkExpression`.
type ParamSpecExpression struct {
	gobject.ParamSpec
}

func ParamSpecExpressionNewFromInternalPtr(ptr uintptr) *ParamSpecExpression {
	cls := &ParamSpecExpression{}
	cls.Ptr = ptr
	return cls
}

func (c *ParamSpecExpression) GoPointer() uintptr {
	return c.Ptr
}

func (c *ParamSpecExpression) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// A `GObject` property value in a `GtkExpression`.
type PropertyExpression struct {
	Expression
}

func PropertyExpressionNewFromInternalPtr(ptr uintptr) *PropertyExpression {
	cls := &PropertyExpression{}
	cls.Ptr = ptr
	return cls
}

var xNewPropertyExpression func([]interface{}, uintptr, string) uintptr

// Creates an expression that looks up a property.
//
// The object to use is found by evaluating the `expression`,
// or using the `this` argument when `expression` is `NULL`.
//
// If the resulting object conforms to `this_type`, its property named
// `property_name` will be queried. Otherwise, this expression's
// evaluation will fail.
//
// The given `this_type` must have a property with `property_name`.
func NewPropertyExpression(ThisTypeVar []interface{}, ExpressionVar *Expression, PropertyNameVar string) *PropertyExpression {
	var cls *PropertyExpression

	cret := xNewPropertyExpression(ThisTypeVar, ExpressionVar.GoPointer(), PropertyNameVar)

	if cret == 0 {
		return cls
	}
	cls = &PropertyExpression{}
	cls.Ptr = cret
	return cls
}

var xNewForPspecPropertyExpression func(uintptr, uintptr) uintptr

// Creates an expression that looks up a property.
//
// The object to use is found by evaluating the `expression`,
// or using the `this` argument when `expression` is `NULL`.
//
// If the resulting object conforms to `this_type`, its
// property specified by `pspec` will be queried.
// Otherwise, this expression's evaluation will fail.
func NewForPspecPropertyExpression(ExpressionVar *Expression, PspecVar *gobject.ParamSpec) *PropertyExpression {
	var cls *PropertyExpression

	cret := xNewForPspecPropertyExpression(ExpressionVar.GoPointer(), PspecVar.GoPointer())

	if cret == 0 {
		return cls
	}
	cls = &PropertyExpression{}
	cls.Ptr = cret
	return cls
}

var xPropertyExpressionGetExpression func(uintptr) uintptr

// Gets the expression specifying the object of
// a property expression.
func (x *PropertyExpression) GetExpression() *Expression {
	var cls *Expression

	cret := xPropertyExpressionGetExpression(x.GoPointer())

	if cret == 0 {
		return cls
	}
	gobject.IncreaseRef(cret)
	cls = &Expression{}
	cls.Ptr = cret
	return cls
}

var xPropertyExpressionGetPspec func(uintptr) uintptr

// Gets the `GParamSpec` specifying the property of
// a property expression.
func (x *PropertyExpression) GetPspec() *gobject.ParamSpec {
	var cls *gobject.ParamSpec

	cret := xPropertyExpressionGetPspec(x.GoPointer())

	if cret == 0 {
		return cls
	}
	gobject.IncreaseRef(cret)
	cls = &gobject.ParamSpec{}
	cls.Ptr = cret
	return cls
}

func (c *PropertyExpression) GoPointer() uintptr {
	return c.Ptr
}

func (c *PropertyExpression) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GTK"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	core.PuregoSafeRegister(&xNewParamSpecExpression, lib, "gtk_param_spec_expression")
	core.PuregoSafeRegister(&xValueDupExpression, lib, "gtk_value_dup_expression")
	core.PuregoSafeRegister(&xValueGetExpression, lib, "gtk_value_get_expression")
	core.PuregoSafeRegister(&xValueSetExpression, lib, "gtk_value_set_expression")
	core.PuregoSafeRegister(&xValueTakeExpression, lib, "gtk_value_take_expression")

	core.PuregoSafeRegister(&xNewCClosureExpression, lib, "gtk_cclosure_expression_new")

	core.PuregoSafeRegister(&xNewClosureExpression, lib, "gtk_closure_expression_new")

	core.PuregoSafeRegister(&xNewConstantExpression, lib, "gtk_constant_expression_new")
	core.PuregoSafeRegister(&xNewForValueConstantExpression, lib, "gtk_constant_expression_new_for_value")

	core.PuregoSafeRegister(&xConstantExpressionGetValue, lib, "gtk_constant_expression_get_value")

	core.PuregoSafeRegister(&xExpressionBind, lib, "gtk_expression_bind")
	core.PuregoSafeRegister(&xExpressionEvaluate, lib, "gtk_expression_evaluate")
	core.PuregoSafeRegister(&xExpressionGetValueType, lib, "gtk_expression_get_value_type")
	core.PuregoSafeRegister(&xExpressionIsStatic, lib, "gtk_expression_is_static")
	core.PuregoSafeRegister(&xExpressionRef, lib, "gtk_expression_ref")
	core.PuregoSafeRegister(&xExpressionUnref, lib, "gtk_expression_unref")
	core.PuregoSafeRegister(&xExpressionWatch, lib, "gtk_expression_watch")

	core.PuregoSafeRegister(&xNewObjectExpression, lib, "gtk_object_expression_new")

	core.PuregoSafeRegister(&xObjectExpressionGetObject, lib, "gtk_object_expression_get_object")

	core.PuregoSafeRegister(&xNewPropertyExpression, lib, "gtk_property_expression_new")
	core.PuregoSafeRegister(&xNewForPspecPropertyExpression, lib, "gtk_property_expression_new_for_pspec")

	core.PuregoSafeRegister(&xPropertyExpressionGetExpression, lib, "gtk_property_expression_get_expression")
	core.PuregoSafeRegister(&xPropertyExpressionGetPspec, lib, "gtk_property_expression_get_pspec")

}
