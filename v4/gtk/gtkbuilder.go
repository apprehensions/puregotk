// Package gtk was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gtk

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/glib"
	"github.com/jwijenbergh/puregotk/v4/gobject"
)

type BuilderClass struct {
}

// Error codes that identify various errors that can occur while using
// `GtkBuilder`.
type BuilderError int

const (

	// A type-func attribute didn’t name
	//  a function that returns a `GType`.
	BuilderErrorInvalidTypeFunctionValue BuilderError = 0
	// The input contained a tag that `GtkBuilder`
	//  can’t handle.
	BuilderErrorUnhandledTagValue BuilderError = 1
	// An attribute that is required by
	//  `GtkBuilder` was missing.
	BuilderErrorMissingAttributeValue BuilderError = 2
	// `GtkBuilder` found an attribute that
	//  it doesn’t understand.
	BuilderErrorInvalidAttributeValue BuilderError = 3
	// `GtkBuilder` found a tag that
	//  it doesn’t understand.
	BuilderErrorInvalidTagValue BuilderError = 4
	// A required property value was
	//  missing.
	BuilderErrorMissingPropertyValueValue BuilderError = 5
	// `GtkBuilder` couldn’t parse
	//  some attribute value.
	BuilderErrorInvalidValueValue BuilderError = 6
	// The input file requires a newer version
	//  of GTK.
	BuilderErrorVersionMismatchValue BuilderError = 7
	// An object id occurred twice.
	BuilderErrorDuplicateIdValue BuilderError = 8
	// A specified object type is of the same type or
	//  derived from the type of the composite class being extended with builder XML.
	BuilderErrorObjectTypeRefusedValue BuilderError = 9
	// The wrong type was specified in a composite class’s template XML
	BuilderErrorTemplateMismatchValue BuilderError = 10
	// The specified property is unknown for the object class.
	BuilderErrorInvalidPropertyValue BuilderError = 11
	// The specified signal is unknown for the object class.
	BuilderErrorInvalidSignalValue BuilderError = 12
	// An object id is unknown.
	BuilderErrorInvalidIdValue BuilderError = 13
	// A function could not be found. This often happens
	//   when symbols are set to be kept private. Compiling code with -rdynamic or using the
	//   `gmodule-export-2.0` pkgconfig module can fix this problem.
	BuilderErrorInvalidFunctionValue BuilderError = 14
)

// A `GtkBuilder` reads XML descriptions of a user interface and
// instantiates the described objects.
//
// To create a `GtkBuilder` from a user interface description, call
// [ctor@Gtk.Builder.new_from_file], [ctor@Gtk.Builder.new_from_resource]
// or [ctor@Gtk.Builder.new_from_string].
//
// In the (unusual) case that you want to add user interface
// descriptions from multiple sources to the same `GtkBuilder` you can
// call [ctor@Gtk.Builder.new] to get an empty builder and populate it by
// (multiple) calls to [method@Gtk.Builder.add_from_file],
// [method@Gtk.Builder.add_from_resource] or
// [method@Gtk.Builder.add_from_string].
//
// A `GtkBuilder` holds a reference to all objects that it has constructed
// and drops these references when it is finalized. This finalization can
// cause the destruction of non-widget objects or widgets which are not
// contained in a toplevel window. For toplevel windows constructed by a
// builder, it is the responsibility of the user to call
// [method@Gtk.Window.destroy] to get rid of them and all the widgets
// they contain.
//
// The functions [method@Gtk.Builder.get_object] and
// [method@Gtk.Builder.get_objects] can be used to access the widgets in
// the interface by the names assigned to them inside the UI description.
// Toplevel windows returned by these functions will stay around until the
// user explicitly destroys them with [method@Gtk.Window.destroy]. Other
// widgets will either be part of a larger hierarchy constructed by the
// builder (in which case you should not have to worry about their lifecycle),
// or without a parent, in which case they have to be added to some container
// to make use of them. Non-widget objects need to be reffed with
// g_object_ref() to keep them beyond the lifespan of the builder.
//
// # GtkBuilder UI Definitions
//
// `GtkBuilder` parses textual descriptions of user interfaces which are
// specified in XML format. We refer to these descriptions as “GtkBuilder
// UI definitions” or just “UI definitions” if the context is clear.
//
// The toplevel element is `&lt;interface&gt;`. It optionally takes a “domain”
// attribute, which will make the builder look for translated strings
// using `dgettext()` in the domain specified. This can also be done by
// calling [method@Gtk.Builder.set_translation_domain] on the builder.
//
// Objects are described by `&lt;object&gt;` elements, which can contain
// `&lt;property&gt;` elements to set properties, `&lt;signal&gt;` elements which
// connect signals to handlers, and `&lt;child&gt;` elements, which describe
// child objects (most often widgets inside a container, but also e.g.
// actions in an action group, or columns in a tree model). A `&lt;child&gt;`
// element contains an `&lt;object&gt;` element which describes the child object.
//
// The target toolkit version(s) are described by `&lt;requires&gt;` elements,
// the “lib” attribute specifies the widget library in question (currently
// the only supported value is “gtk”) and the “version” attribute specifies
// the target version in the form “`&lt;major&gt;`.`&lt;minor&gt;`”. `GtkBuilder` will
// error out if the version requirements are not met.
//
// Typically, the specific kind of object represented by an `&lt;object&gt;`
// element is specified by the “class” attribute. If the type has not
// been loaded yet, GTK tries to find the `get_type()` function from the
// class name by applying heuristics. This works in most cases, but if
// necessary, it is possible to specify the name of the `get_type()`
// function explicitly with the "type-func" attribute.
//
// Objects may be given a name with the “id” attribute, which allows the
// application to retrieve them from the builder with
// [method@Gtk.Builder.get_object]. An id is also necessary to use the
// object as property value in other parts of the UI definition. GTK
// reserves ids starting and ending with `___` (three consecutive
// underscores) for its own purposes.
//
// Setting properties of objects is pretty straightforward with the
// `&lt;property&gt;` element: the “name” attribute specifies the name of the
// property, and the content of the element specifies the value.
// If the “translatable” attribute is set to a true value, GTK uses
// `gettext()` (or `dgettext()` if the builder has a translation domain set)
// to find a translation for the value. This happens before the value
// is parsed, so it can be used for properties of any type, but it is
// probably most useful for string properties. It is also possible to
// specify a context to disambiguate short strings, and comments which
// may help the translators.
//
// `GtkBuilder` can parse textual representations for the most common
// property types: characters, strings, integers, floating-point numbers,
// booleans (strings like “TRUE”, “t”, “yes”, “y”, “1” are interpreted
// as %TRUE, strings like “FALSE”, “f”, “no”, “n”, “0” are interpreted
// as %FALSE), enumerations (can be specified by their name, nick or
// integer value), flags (can be specified by their name, nick, integer
// value, optionally combined with “|”, e.g.
// “GTK_INPUT_HINT_EMOJI|GTK_INPUT_HINT_LOWERCASE”)
// and colors (in a format understood by [method@Gdk.RGBA.parse]).
//
// `GVariant`s can be specified in the format understood by
// g_variant_parse(), and pixbufs can be specified as a filename of an
// image file to load.
//
// Objects can be referred to by their name and by default refer to
// objects declared in the local XML fragment and objects exposed via
// [method@Gtk.Builder.expose_object]. In general, `GtkBuilder` allows
// forward references to objects — declared in the local XML; an object
// doesn’t have to be constructed before it can be referred to. The
// exception to this rule is that an object has to be constructed before
// it can be used as the value of a construct-only property.
//
// It is also possible to bind a property value to another object's
// property value using the attributes "bind-source" to specify the
// source object of the binding, and optionally, "bind-property" and
// "bind-flags" to specify the source property and source binding flags
// respectively. Internally, `GtkBuilder` implements this using `GBinding`
// objects. For more information see g_object_bind_property().
//
// Sometimes it is necessary to refer to widgets which have implicitly
// been constructed by GTK as part of a composite widget, to set
// properties on them or to add further children (e.g. the content area
// of a `GtkDialog`). This can be achieved by setting the “internal-child”
// property of the `&lt;child&gt;` element to a true value. Note that `GtkBuilder`
// still requires an `&lt;object&gt;` element for the internal child, even if it
// has already been constructed.
//
// A number of widgets have different places where a child can be added
// (e.g. tabs vs. page content in notebooks). This can be reflected in
// a UI definition by specifying the “type” attribute on a `&lt;child&gt;`
// The possible values for the “type” attribute are described in the
// sections describing the widget-specific portions of UI definitions.
//
// # Signal handlers and function pointers
//
// Signal handlers are set up with the `&lt;signal&gt;` element. The “name”
// attribute specifies the name of the signal, and the “handler” attribute
// specifies the function to connect to the signal.
// The remaining attributes, “after”, “swapped” and “object”, have the
// same meaning as the corresponding parameters of the
// g_signal_connect_object() or g_signal_connect_data() functions. A
// “last_modification_time” attribute is also allowed, but it does not
// have a meaning to the builder.
//
// If you rely on `GModule` support to lookup callbacks in the symbol table,
// the following details should be noted:
//
// When compiling applications for Windows, you must declare signal callbacks
// with %G_MODULE_EXPORT, or they will not be put in the symbol table.
// On Linux and Unix, this is not necessary; applications should instead
// be compiled with the -Wl,--export-dynamic `CFLAGS`, and linked against
// `gmodule-export-2.0`.
//
// # A GtkBuilder UI Definition
//
// ```xml
// &lt;interface&gt;
//
//	&lt;object class="GtkDialog" id="dialog1"&gt;
//	  &lt;child internal-child="content_area"&gt;
//	    &lt;object class="GtkBox" id="vbox1"&gt;
//	      &lt;child internal-child="action_area"&gt;
//	        &lt;object class="GtkBox" id="hbuttonbox1"&gt;
//	          &lt;child&gt;
//	            &lt;object class="GtkButton" id="ok_button"&gt;
//	              &lt;property name="label" translatable="yes"&gt;_Ok&lt;/property&gt;
//	              &lt;property name="use-underline"&gt;True&lt;/property&gt;
//	              &lt;signal name="clicked" handler="ok_button_clicked"/&gt;
//	            &lt;/object&gt;
//	          &lt;/child&gt;
//	        &lt;/object&gt;
//	      &lt;/child&gt;
//	    &lt;/object&gt;
//	  &lt;/child&gt;
//	&lt;/object&gt;
//
// &lt;/interface&gt;
// ```
//
// Beyond this general structure, several object classes define their
// own XML DTD fragments for filling in the ANY placeholders in the DTD
// above. Note that a custom element in a &lt;child&gt; element gets parsed by
// the custom tag handler of the parent object, while a custom element in
// an &lt;object&gt; element gets parsed by the custom tag handler of the object.
//
// These XML fragments are explained in the documentation of the
// respective objects.
//
// A `&lt;template&gt;` tag can be used to define a widget class’s components.
// See the [GtkWidget documentation](class.Widget.html#building-composite-widgets-from-template-xml) for details.
type Builder struct {
	gobject.Object
}

func BuilderNewFromInternalPtr(ptr uintptr) *Builder {
	cls := &Builder{}
	cls.Ptr = ptr
	return cls
}

var xNewBuilder func() uintptr

// Creates a new empty builder object.
//
// This function is only useful if you intend to make multiple calls
// to [method@Gtk.Builder.add_from_file], [method@Gtk.Builder.add_from_resource]
// or [method@Gtk.Builder.add_from_string] in order to merge multiple UI
// descriptions into a single builder.
func NewBuilder() *Builder {
	var cls *Builder

	cret := xNewBuilder()

	if cret == 0 {
		return cls
	}
	cls = &Builder{}
	cls.Ptr = cret
	return cls
}

var xNewFromFileBuilder func(string) uintptr

// Parses the UI definition in the file @filename.
//
// If there is an error opening the file or parsing the description then
// the program will be aborted. You should only ever attempt to parse
// user interface descriptions that are shipped as part of your program.
func NewFromFileBuilder(FilenameVar string) *Builder {
	var cls *Builder

	cret := xNewFromFileBuilder(FilenameVar)

	if cret == 0 {
		return cls
	}
	cls = &Builder{}
	cls.Ptr = cret
	return cls
}

var xNewFromResourceBuilder func(string) uintptr

// Parses the UI definition at @resource_path.
//
// If there is an error locating the resource or parsing the
// description, then the program will be aborted.
func NewFromResourceBuilder(ResourcePathVar string) *Builder {
	var cls *Builder

	cret := xNewFromResourceBuilder(ResourcePathVar)

	if cret == 0 {
		return cls
	}
	cls = &Builder{}
	cls.Ptr = cret
	return cls
}

var xNewFromStringBuilder func(string, int) uintptr

// Parses the UI definition in @string.
//
// If @string is %NULL-terminated, then @length should be -1.
// If @length is not -1, then it is the length of @string.
//
// If there is an error parsing @string then the program will be
// aborted. You should not attempt to parse user interface description
// from untrusted sources.
func NewFromStringBuilder(StringVar string, LengthVar int) *Builder {
	var cls *Builder

	cret := xNewFromStringBuilder(StringVar, LengthVar)

	if cret == 0 {
		return cls
	}
	cls = &Builder{}
	cls.Ptr = cret
	return cls
}

var xBuilderAddFromFile func(uintptr, string, **glib.Error) bool

// Parses a file containing a UI definition and merges it with
// the current contents of @builder.
//
// This function is useful if you need to call
// [method@Gtk.Builder.set_current_object]) to add user data to
// callbacks before loading GtkBuilder UI. Otherwise, you probably
// want [ctor@Gtk.Builder.new_from_file] instead.
//
// If an error occurs, 0 will be returned and @error will be assigned a
// `GError` from the `GTK_BUILDER_ERROR`, `G_MARKUP_ERROR` or `G_FILE_ERROR`
// domains.
//
// It’s not really reasonable to attempt to handle failures of this
// call. You should not use this function with untrusted files (ie:
// files that are not part of your application). Broken `GtkBuilder`
// files can easily crash your program, and it’s possible that memory
// was leaked leading up to the reported failure. The only reasonable
// thing to do when an error is detected is to call `g_error()`.
func (x *Builder) AddFromFile(FilenameVar string) (bool, error) {
	var cerr *glib.Error

	cret := xBuilderAddFromFile(x.GoPointer(), FilenameVar, &cerr)
	if cerr == nil {
		return cret, nil
	}
	return cret, cerr

}

var xBuilderAddFromResource func(uintptr, string, **glib.Error) bool

// Parses a resource file containing a UI definition
// and merges it with the current contents of @builder.
//
// This function is useful if you need to call
// [method@Gtk.Builder.set_current_object] to add user data to
// callbacks before loading GtkBuilder UI. Otherwise, you probably
// want [ctor@Gtk.Builder.new_from_resource] instead.
//
// If an error occurs, 0 will be returned and @error will be assigned a
// `GError` from the %GTK_BUILDER_ERROR, %G_MARKUP_ERROR or %G_RESOURCE_ERROR
// domain.
//
// It’s not really reasonable to attempt to handle failures of this
// call.  The only reasonable thing to do when an error is detected is
// to call g_error().
func (x *Builder) AddFromResource(ResourcePathVar string) (bool, error) {
	var cerr *glib.Error

	cret := xBuilderAddFromResource(x.GoPointer(), ResourcePathVar, &cerr)
	if cerr == nil {
		return cret, nil
	}
	return cret, cerr

}

var xBuilderAddFromString func(uintptr, string, int, **glib.Error) bool

// Parses a string containing a UI definition and merges it
// with the current contents of @builder.
//
// This function is useful if you need to call
// [method@Gtk.Builder.set_current_object] to add user data to
// callbacks before loading `GtkBuilder` UI. Otherwise, you probably
// want [ctor@Gtk.Builder.new_from_string] instead.
//
// Upon errors %FALSE will be returned and @error will be assigned a
// `GError` from the %GTK_BUILDER_ERROR, %G_MARKUP_ERROR or
// %G_VARIANT_PARSE_ERROR domain.
//
// It’s not really reasonable to attempt to handle failures of this
// call.  The only reasonable thing to do when an error is detected is
// to call g_error().
func (x *Builder) AddFromString(BufferVar string, LengthVar int) (bool, error) {
	var cerr *glib.Error

	cret := xBuilderAddFromString(x.GoPointer(), BufferVar, LengthVar, &cerr)
	if cerr == nil {
		return cret, nil
	}
	return cret, cerr

}

var xBuilderAddObjectsFromFile func(uintptr, string, uintptr, **glib.Error) bool

// Parses a file containing a UI definition building only the
// requested objects and merges them with the current contents
// of @builder.
//
// Upon errors, 0 will be returned and @error will be assigned a
// `GError` from the %GTK_BUILDER_ERROR, %G_MARKUP_ERROR or %G_FILE_ERROR
// domain.
//
// If you are adding an object that depends on an object that is not
// its child (for instance a `GtkTreeView` that depends on its
// `GtkTreeModel`), you have to explicitly list all of them in @object_ids.
func (x *Builder) AddObjectsFromFile(FilenameVar string, ObjectIdsVar uintptr) (bool, error) {
	var cerr *glib.Error

	cret := xBuilderAddObjectsFromFile(x.GoPointer(), FilenameVar, ObjectIdsVar, &cerr)
	if cerr == nil {
		return cret, nil
	}
	return cret, cerr

}

var xBuilderAddObjectsFromResource func(uintptr, string, uintptr, **glib.Error) bool

// Parses a resource file containing a UI definition, building
// only the requested objects and merges them with the current
// contents of @builder.
//
// Upon errors, 0 will be returned and @error will be assigned a
// `GError` from the %GTK_BUILDER_ERROR, %G_MARKUP_ERROR or %G_RESOURCE_ERROR
// domain.
//
// If you are adding an object that depends on an object that is not
// its child (for instance a `GtkTreeView` that depends on its
// `GtkTreeModel`), you have to explicitly list all of them in @object_ids.
func (x *Builder) AddObjectsFromResource(ResourcePathVar string, ObjectIdsVar uintptr) (bool, error) {
	var cerr *glib.Error

	cret := xBuilderAddObjectsFromResource(x.GoPointer(), ResourcePathVar, ObjectIdsVar, &cerr)
	if cerr == nil {
		return cret, nil
	}
	return cret, cerr

}

var xBuilderAddObjectsFromString func(uintptr, string, int, uintptr, **glib.Error) bool

// Parses a string containing a UI definition, building only the
// requested objects and merges them with the current contents of
// @builder.
//
// Upon errors %FALSE will be returned and @error will be assigned a
// `GError` from the %GTK_BUILDER_ERROR or %G_MARKUP_ERROR domain.
//
// If you are adding an object that depends on an object that is not
// its child (for instance a `GtkTreeView` that depends on its
// `GtkTreeModel`), you have to explicitly list all of them in @object_ids.
func (x *Builder) AddObjectsFromString(BufferVar string, LengthVar int, ObjectIdsVar uintptr) (bool, error) {
	var cerr *glib.Error

	cret := xBuilderAddObjectsFromString(x.GoPointer(), BufferVar, LengthVar, ObjectIdsVar, &cerr)
	if cerr == nil {
		return cret, nil
	}
	return cret, cerr

}

var xBuilderCreateClosure func(uintptr, string, BuilderClosureFlags, uintptr, **glib.Error) *gobject.Closure

// Creates a closure to invoke the function called @function_name.
//
// This is using the create_closure() implementation of @builder's
// [iface@Gtk.BuilderScope].
//
// If no closure could be created, %NULL will be returned and @error
// will be set.
func (x *Builder) CreateClosure(FunctionNameVar string, FlagsVar BuilderClosureFlags, ObjectVar *gobject.Object) (*gobject.Closure, error) {
	var cerr *glib.Error

	cret := xBuilderCreateClosure(x.GoPointer(), FunctionNameVar, FlagsVar, ObjectVar.GoPointer(), &cerr)
	if cerr == nil {
		return cret, nil
	}
	return cret, cerr

}

var xBuilderExposeObject func(uintptr, string, uintptr)

// Add @object to the @builder object pool so it can be
// referenced just like any other object built by builder.
//
// Only a single object may be added using @name. However,
// it is not an error to expose the same object under multiple
// names. `gtk_builder_get_object()` may be used to determine
// if an object has already been added with @name.
func (x *Builder) ExposeObject(NameVar string, ObjectVar *gobject.Object) {

	xBuilderExposeObject(x.GoPointer(), NameVar, ObjectVar.GoPointer())

}

var xBuilderExtendWithTemplate func(uintptr, uintptr, []interface{}, string, int, **glib.Error) bool

// Main private entry point for building composite components
// from template XML.
//
// Most likely you do not need to call this function in applications as
// templates are handled by `GtkWidget`.
func (x *Builder) ExtendWithTemplate(ObjectVar *gobject.Object, TemplateTypeVar []interface{}, BufferVar string, LengthVar int) (bool, error) {
	var cerr *glib.Error

	cret := xBuilderExtendWithTemplate(x.GoPointer(), ObjectVar.GoPointer(), TemplateTypeVar, BufferVar, LengthVar, &cerr)
	if cerr == nil {
		return cret, nil
	}
	return cret, cerr

}

var xBuilderGetCurrentObject func(uintptr) uintptr

// Gets the current object set via gtk_builder_set_current_object().
func (x *Builder) GetCurrentObject() *gobject.Object {
	var cls *gobject.Object

	cret := xBuilderGetCurrentObject(x.GoPointer())

	if cret == 0 {
		return cls
	}
	gobject.IncreaseRef(cret)
	cls = &gobject.Object{}
	cls.Ptr = cret
	return cls
}

var xBuilderGetObject func(uintptr, string) uintptr

// Gets the object named @name.
//
// Note that this function does not increment the reference count
// of the returned object.
func (x *Builder) GetObject(NameVar string) *gobject.Object {
	var cls *gobject.Object

	cret := xBuilderGetObject(x.GoPointer(), NameVar)

	if cret == 0 {
		return cls
	}
	gobject.IncreaseRef(cret)
	cls = &gobject.Object{}
	cls.Ptr = cret
	return cls
}

var xBuilderGetObjects func(uintptr) *glib.SList

// Gets all objects that have been constructed by @builder.
//
// Note that this function does not increment the reference
// counts of the returned objects.
func (x *Builder) GetObjects() *glib.SList {

	cret := xBuilderGetObjects(x.GoPointer())
	return cret
}

var xBuilderGetScope func(uintptr) uintptr

// Gets the scope in use that was set via gtk_builder_set_scope().
func (x *Builder) GetScope() *BuilderScopeBase {
	var cls *BuilderScopeBase

	cret := xBuilderGetScope(x.GoPointer())

	if cret == 0 {
		return cls
	}
	gobject.IncreaseRef(cret)
	cls = &BuilderScopeBase{}
	cls.Ptr = cret
	return cls
}

var xBuilderGetTranslationDomain func(uintptr) string

// Gets the translation domain of @builder.
func (x *Builder) GetTranslationDomain() string {

	cret := xBuilderGetTranslationDomain(x.GoPointer())
	return cret
}

var xBuilderGetTypeFromName func(uintptr, string) []interface{}

// Looks up a type by name.
//
// This is using the virtual function that `GtkBuilder` has
// for that purpose. This is mainly used when implementing
// the `GtkBuildable` interface on a type.
func (x *Builder) GetTypeFromName(TypeNameVar string) []interface{} {

	cret := xBuilderGetTypeFromName(x.GoPointer(), TypeNameVar)
	return cret
}

var xBuilderSetCurrentObject func(uintptr, uintptr)

// Sets the current object for the @builder.
//
// The current object can be thought of as the `this` object that the
// builder is working for and will often be used as the default object
// when an object is optional.
//
// [method@Gtk.Widget.init_template] for example will set the current
// object to the widget the template is inited for. For functions like
// [ctor@Gtk.Builder.new_from_resource], the current object will be %NULL.
func (x *Builder) SetCurrentObject(CurrentObjectVar *gobject.Object) {

	xBuilderSetCurrentObject(x.GoPointer(), CurrentObjectVar.GoPointer())

}

var xBuilderSetScope func(uintptr, uintptr)

// Sets the scope the builder should operate in.
//
// If @scope is %NULL, a new [class@Gtk.BuilderCScope] will be created.
func (x *Builder) SetScope(ScopeVar BuilderScope) {

	xBuilderSetScope(x.GoPointer(), ScopeVar.GoPointer())

}

var xBuilderSetTranslationDomain func(uintptr, string)

// Sets the translation domain of @builder.
func (x *Builder) SetTranslationDomain(DomainVar string) {

	xBuilderSetTranslationDomain(x.GoPointer(), DomainVar)

}

var xBuilderValueFromString func(uintptr, uintptr, string, *gobject.Value, **glib.Error) bool

// Demarshals a value from a string.
//
// This function calls g_value_init() on the @value argument,
// so it need not be initialised beforehand.
//
// Can handle char, uchar, boolean, int, uint, long,
// ulong, enum, flags, float, double, string, `GdkRGBA` and
// `GtkAdjustment` type values.
//
// Upon errors %FALSE will be returned and @error will be
// assigned a `GError` from the %GTK_BUILDER_ERROR domain.
func (x *Builder) ValueFromString(PspecVar *gobject.ParamSpec, StringVar string, ValueVar *gobject.Value) (bool, error) {
	var cerr *glib.Error

	cret := xBuilderValueFromString(x.GoPointer(), PspecVar.GoPointer(), StringVar, ValueVar, &cerr)
	if cerr == nil {
		return cret, nil
	}
	return cret, cerr

}

var xBuilderValueFromStringType func(uintptr, []interface{}, string, *gobject.Value, **glib.Error) bool

// Demarshals a value from a string.
//
// Unlike [method@Gtk.Builder.value_from_string], this function
// takes a `GType` instead of `GParamSpec`.
//
// Calls g_value_init() on the @value argument, so it
// need not be initialised beforehand.
//
// Upon errors %FALSE will be returned and @error will be
// assigned a `GError` from the %GTK_BUILDER_ERROR domain.
func (x *Builder) ValueFromStringType(TypeVar []interface{}, StringVar string, ValueVar *gobject.Value) (bool, error) {
	var cerr *glib.Error

	cret := xBuilderValueFromStringType(x.GoPointer(), TypeVar, StringVar, ValueVar, &cerr)
	if cerr == nil {
		return cret, nil
	}
	return cret, cerr

}

func (c *Builder) GoPointer() uintptr {
	return c.Ptr
}

func (c *Builder) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GTK"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xNewBuilder, lib, "gtk_builder_new")
	core.PuregoSafeRegister(&xNewFromFileBuilder, lib, "gtk_builder_new_from_file")
	core.PuregoSafeRegister(&xNewFromResourceBuilder, lib, "gtk_builder_new_from_resource")
	core.PuregoSafeRegister(&xNewFromStringBuilder, lib, "gtk_builder_new_from_string")

	core.PuregoSafeRegister(&xBuilderAddFromFile, lib, "gtk_builder_add_from_file")
	core.PuregoSafeRegister(&xBuilderAddFromResource, lib, "gtk_builder_add_from_resource")
	core.PuregoSafeRegister(&xBuilderAddFromString, lib, "gtk_builder_add_from_string")
	core.PuregoSafeRegister(&xBuilderAddObjectsFromFile, lib, "gtk_builder_add_objects_from_file")
	core.PuregoSafeRegister(&xBuilderAddObjectsFromResource, lib, "gtk_builder_add_objects_from_resource")
	core.PuregoSafeRegister(&xBuilderAddObjectsFromString, lib, "gtk_builder_add_objects_from_string")
	core.PuregoSafeRegister(&xBuilderCreateClosure, lib, "gtk_builder_create_closure")
	core.PuregoSafeRegister(&xBuilderExposeObject, lib, "gtk_builder_expose_object")
	core.PuregoSafeRegister(&xBuilderExtendWithTemplate, lib, "gtk_builder_extend_with_template")
	core.PuregoSafeRegister(&xBuilderGetCurrentObject, lib, "gtk_builder_get_current_object")
	core.PuregoSafeRegister(&xBuilderGetObject, lib, "gtk_builder_get_object")
	core.PuregoSafeRegister(&xBuilderGetObjects, lib, "gtk_builder_get_objects")
	core.PuregoSafeRegister(&xBuilderGetScope, lib, "gtk_builder_get_scope")
	core.PuregoSafeRegister(&xBuilderGetTranslationDomain, lib, "gtk_builder_get_translation_domain")
	core.PuregoSafeRegister(&xBuilderGetTypeFromName, lib, "gtk_builder_get_type_from_name")
	core.PuregoSafeRegister(&xBuilderSetCurrentObject, lib, "gtk_builder_set_current_object")
	core.PuregoSafeRegister(&xBuilderSetScope, lib, "gtk_builder_set_scope")
	core.PuregoSafeRegister(&xBuilderSetTranslationDomain, lib, "gtk_builder_set_translation_domain")
	core.PuregoSafeRegister(&xBuilderValueFromString, lib, "gtk_builder_value_from_string")
	core.PuregoSafeRegister(&xBuilderValueFromStringType, lib, "gtk_builder_value_from_string_type")

}
