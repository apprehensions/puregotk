// Package gtk was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gtk

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/gobject"
)

type EditableInterface struct {
	BaseIface uintptr
}

// `GtkEditable` is an interface for text editing widgets.
//
// Typical examples of editable widgets are [class@Gtk.Entry] and
// [class@Gtk.SpinButton]. It contains functions for generically manipulating
// an editable widget, a large number of action signals used for key bindings,
// and several signals that an application can connect to modify the behavior
// of a widget.
//
// As an example of the latter usage, by connecting the following handler to
// [signal@Gtk.Editable::insert-text], an application can convert all entry
// into a widget into uppercase.
//
// ## Forcing entry to uppercase.
//
// ```c
// #include &lt;ctype.h&gt;
//
// void
// insert_text_handler (GtkEditable *editable,
//
//	const char  *text,
//	int          length,
//	int         *position,
//	gpointer     data)
//
//	{
//	  char *result = g_utf8_strup (text, length);
//
//	  g_signal_handlers_block_by_func (editable,
//	                               (gpointer) insert_text_handler, data);
//	  gtk_editable_insert_text (editable, result, length, position);
//	  g_signal_handlers_unblock_by_func (editable,
//	                                     (gpointer) insert_text_handler, data);
//
//	  g_signal_stop_emission_by_name (editable, "insert_text");
//
//	  g_free (result);
//	}
//
// ```
//
// ## Implementing GtkEditable
//
// The most likely scenario for implementing `GtkEditable` on your own widget
// is that you will embed a `GtkText` inside a complex widget, and want to
// delegate the editable functionality to that text widget. `GtkEditable`
// provides some utility functions to make this easy.
//
// In your class_init function, call [func@Gtk.Editable.install_properties],
// passing the first available property ID:
//
// ```c
// static void
// my_class_init (MyClass *class)
//
//	{
//	  ...
//	  g_object_class_install_properties (object_class, NUM_PROPERTIES, props);
//	  gtk_editable_install_properties (object_clas, NUM_PROPERTIES);
//	  ...
//	}
//
// ```
//
// In your interface_init function for the `GtkEditable` interface, provide
// an implementation for the get_delegate vfunc that returns your text widget:
//
// ```c
// GtkEditable *
// get_editable_delegate (GtkEditable *editable)
//
//	{
//	  return GTK_EDITABLE (MY_WIDGET (editable)-&gt;text_widget);
//	}
//
// static void
// my_editable_init (GtkEditableInterface *iface)
//
//	{
//	  iface-&gt;get_delegate = get_editable_delegate;
//	}
//
// ```
//
// You don't need to provide any other vfuncs. The default implementations
// work by forwarding to the delegate that the GtkEditableInterface.get_delegate()
// vfunc returns.
//
// In your instance_init function, create your text widget, and then call
// [method@Gtk.Editable.init_delegate]:
//
// ```c
// static void
// my_widget_init (MyWidget *self)
//
//	{
//	  ...
//	  self-&gt;text_widget = gtk_text_new ();
//	  gtk_editable_init_delegate (GTK_EDITABLE (self));
//	  ...
//	}
//
// ```
//
// In your dispose function, call [method@Gtk.Editable.finish_delegate] before
// destroying your text widget:
//
// ```c
// static void
// my_widget_dispose (GObject *object)
//
//	{
//	  ...
//	  gtk_editable_finish_delegate (GTK_EDITABLE (self));
//	  g_clear_pointer (&amp;self-&gt;text_widget, gtk_widget_unparent);
//	  ...
//	}
//
// ```
//
// Finally, use [func@Gtk.Editable.delegate_set_property] in your `set_property`
// function (and similar for `get_property`), to set the editable properties:
//
// ```c
//
//	...
//	if (gtk_editable_delegate_set_property (object, prop_id, value, pspec))
//	  return;
//
//	switch (prop_id)
//	...
//
// ```
//
// It is important to note that if you create a `GtkEditable` that uses
// a delegate, the low level [signal@Gtk.Editable::insert-text] and
// [signal@Gtk.Editable::delete-text] signals will be propagated from the
// "wrapper" editable to the delegate, but they will not be propagated from
// the delegate to the "wrapper" editable, as they would cause an infinite
// recursion. If you wish to connect to the [signal@Gtk.Editable::insert-text]
// and [signal@Gtk.Editable::delete-text] signals, you will need to connect
// to them on the delegate obtained via [method@Gtk.Editable.get_delegate].
type Editable interface {
	GoPointer() uintptr
	SetGoPointer(uintptr)
	DeleteSelection()
	DeleteText(StartPosVar int32, EndPosVar int32)
	FinishDelegate()
	GetAlignment() float32
	GetChars(StartPosVar int32, EndPosVar int32) string
	GetDelegate() *EditableBase
	GetEditable() bool
	GetEnableUndo() bool
	GetMaxWidthChars() int32
	GetPosition() int32
	GetSelectionBounds(StartPosVar int32, EndPosVar int32) bool
	GetText() string
	GetWidthChars() int32
	InitDelegate()
	InsertText(TextVar string, LengthVar int32, PositionVar int32)
	SelectRegion(StartPosVar int32, EndPosVar int32)
	SetAlignment(XalignVar float32)
	SetEditable(IsEditableVar bool)
	SetEnableUndo(EnableUndoVar bool)
	SetMaxWidthChars(NCharsVar int32)
	SetPosition(PositionVar int32)
	SetText(TextVar string)
	SetWidthChars(NCharsVar int32)
}
type EditableBase struct {
	Ptr uintptr
}

func (x *EditableBase) GoPointer() uintptr {
	return x.Ptr
}

func (x *EditableBase) SetGoPointer(ptr uintptr) {
	x.Ptr = ptr
}

// Deletes the currently selected text of the editable.
//
// This call doesn’t do anything if there is no selected text.
func (x *EditableBase) DeleteSelection() {

	XGtkEditableDeleteSelection(x.GoPointer())

}

// Deletes a sequence of characters.
//
// The characters that are deleted are those characters at positions
// from @start_pos up to, but not including @end_pos. If @end_pos is
// negative, then the characters deleted are those from @start_pos to
// the end of the text.
//
// Note that the positions are specified in characters, not bytes.
func (x *EditableBase) DeleteText(StartPosVar int32, EndPosVar int32) {

	XGtkEditableDeleteText(x.GoPointer(), StartPosVar, EndPosVar)

}

// Undoes the setup done by [method@Gtk.Editable.init_delegate].
//
// This is a helper function that should be called from dispose,
// before removing the delegate object.
func (x *EditableBase) FinishDelegate() {

	XGtkEditableFinishDelegate(x.GoPointer())

}

// Gets the alignment of the editable.
func (x *EditableBase) GetAlignment() float32 {

	return XGtkEditableGetAlignment(x.GoPointer())

}

// Retrieves a sequence of characters.
//
// The characters that are retrieved are those characters at positions
// from @start_pos up to, but not including @end_pos. If @end_pos is negative,
// then the characters retrieved are those characters from @start_pos to
// the end of the text.
//
// Note that positions are specified in characters, not bytes.
func (x *EditableBase) GetChars(StartPosVar int32, EndPosVar int32) string {

	return XGtkEditableGetChars(x.GoPointer(), StartPosVar, EndPosVar)

}

// Gets the `GtkEditable` that @editable is delegating its
// implementation to.
//
// Typically, the delegate is a [class@Gtk.Text] widget.
func (x *EditableBase) GetDelegate() *EditableBase {

	GetDelegatePtr := XGtkEditableGetDelegate(x.GoPointer())
	if GetDelegatePtr == 0 {
		return nil
	}

	gobject.IncreaseRef(GetDelegatePtr)

	GetDelegateCls := &EditableBase{}
	GetDelegateCls.Ptr = GetDelegatePtr
	return GetDelegateCls

}

// Retrieves whether @editable is editable.
func (x *EditableBase) GetEditable() bool {

	return XGtkEditableGetEditable(x.GoPointer())

}

// Gets if undo/redo actions are enabled for @editable
func (x *EditableBase) GetEnableUndo() bool {

	return XGtkEditableGetEnableUndo(x.GoPointer())

}

// Retrieves the desired maximum width of @editable, in characters.
func (x *EditableBase) GetMaxWidthChars() int32 {

	return XGtkEditableGetMaxWidthChars(x.GoPointer())

}

// Retrieves the current position of the cursor relative
// to the start of the content of the editable.
//
// Note that this position is in characters, not in bytes.
func (x *EditableBase) GetPosition() int32 {

	return XGtkEditableGetPosition(x.GoPointer())

}

// Retrieves the selection bound of the editable.
//
// @start_pos will be filled with the start of the selection and
// @end_pos with end. If no text was selected both will be identical
// and %FALSE will be returned.
//
// Note that positions are specified in characters, not bytes.
func (x *EditableBase) GetSelectionBounds(StartPosVar int32, EndPosVar int32) bool {

	return XGtkEditableGetSelectionBounds(x.GoPointer(), StartPosVar, EndPosVar)

}

// Retrieves the contents of @editable.
//
// The returned string is owned by GTK and must not be modified or freed.
func (x *EditableBase) GetText() string {

	return XGtkEditableGetText(x.GoPointer())

}

// Gets the number of characters of space reserved
// for the contents of the editable.
func (x *EditableBase) GetWidthChars() int32 {

	return XGtkEditableGetWidthChars(x.GoPointer())

}

// Sets up a delegate for `GtkEditable`.
//
// This is assuming that the get_delegate vfunc in the `GtkEditable`
// interface has been set up for the @editable's type.
//
// This is a helper function that should be called in instance init,
// after creating the delegate object.
func (x *EditableBase) InitDelegate() {

	XGtkEditableInitDelegate(x.GoPointer())

}

// Inserts @length bytes of @text into the contents of the
// widget, at position @position.
//
// Note that the position is in characters, not in bytes.
// The function updates @position to point after the newly
// inserted text.
func (x *EditableBase) InsertText(TextVar string, LengthVar int32, PositionVar int32) {

	XGtkEditableInsertText(x.GoPointer(), TextVar, LengthVar, PositionVar)

}

// Selects a region of text.
//
// The characters that are selected are those characters at positions
// from @start_pos up to, but not including @end_pos. If @end_pos is
// negative, then the characters selected are those characters from
// @start_pos to  the end of the text.
//
// Note that positions are specified in characters, not bytes.
func (x *EditableBase) SelectRegion(StartPosVar int32, EndPosVar int32) {

	XGtkEditableSelectRegion(x.GoPointer(), StartPosVar, EndPosVar)

}

// Sets the alignment for the contents of the editable.
//
// This controls the horizontal positioning of the contents when
// the displayed text is shorter than the width of the editable.
func (x *EditableBase) SetAlignment(XalignVar float32) {

	XGtkEditableSetAlignment(x.GoPointer(), XalignVar)

}

// Determines if the user can edit the text in the editable widget.
func (x *EditableBase) SetEditable(IsEditableVar bool) {

	XGtkEditableSetEditable(x.GoPointer(), IsEditableVar)

}

// If enabled, changes to @editable will be saved for undo/redo
// actions.
//
// This results in an additional copy of text changes and are not
// stored in secure memory. As such, undo is forcefully disabled
// when [property@Gtk.Text:visibility] is set to %FALSE.
func (x *EditableBase) SetEnableUndo(EnableUndoVar bool) {

	XGtkEditableSetEnableUndo(x.GoPointer(), EnableUndoVar)

}

// Sets the desired maximum width in characters of @editable.
func (x *EditableBase) SetMaxWidthChars(NCharsVar int32) {

	XGtkEditableSetMaxWidthChars(x.GoPointer(), NCharsVar)

}

// Sets the cursor position in the editable to the given value.
//
// The cursor is displayed before the character with the given (base 0)
// index in the contents of the editable. The value must be less than
// or equal to the number of characters in the editable. A value of -1
// indicates that the position should be set after the last character
// of the editable. Note that @position is in characters, not in bytes.
func (x *EditableBase) SetPosition(PositionVar int32) {

	XGtkEditableSetPosition(x.GoPointer(), PositionVar)

}

// Sets the text in the editable to the given value.
//
// This is replacing the current contents.
func (x *EditableBase) SetText(TextVar string) {

	XGtkEditableSetText(x.GoPointer(), TextVar)

}

// Changes the size request of the editable to be about the
// right size for @n_chars characters.
//
// Note that it changes the size request, the size can still
// be affected by how you pack the widget into containers.
// If @n_chars is -1, the size reverts to the default size.
func (x *EditableBase) SetWidthChars(NCharsVar int32) {

	XGtkEditableSetWidthChars(x.GoPointer(), NCharsVar)

}

var XGtkEditableDeleteSelection func(uintptr)
var XGtkEditableDeleteText func(uintptr, int32, int32)
var XGtkEditableFinishDelegate func(uintptr)
var XGtkEditableGetAlignment func(uintptr) float32
var XGtkEditableGetChars func(uintptr, int32, int32) string
var XGtkEditableGetDelegate func(uintptr) uintptr
var XGtkEditableGetEditable func(uintptr) bool
var XGtkEditableGetEnableUndo func(uintptr) bool
var XGtkEditableGetMaxWidthChars func(uintptr) int32
var XGtkEditableGetPosition func(uintptr) int32
var XGtkEditableGetSelectionBounds func(uintptr, int32, int32) bool
var XGtkEditableGetText func(uintptr) string
var XGtkEditableGetWidthChars func(uintptr) int32
var XGtkEditableInitDelegate func(uintptr)
var XGtkEditableInsertText func(uintptr, string, int32, int32)
var XGtkEditableSelectRegion func(uintptr, int32, int32)
var XGtkEditableSetAlignment func(uintptr, float32)
var XGtkEditableSetEditable func(uintptr, bool)
var XGtkEditableSetEnableUndo func(uintptr, bool)
var XGtkEditableSetMaxWidthChars func(uintptr, int32)
var XGtkEditableSetPosition func(uintptr, int32)
var XGtkEditableSetText func(uintptr, string)
var XGtkEditableSetWidthChars func(uintptr, int32)

// The identifiers for [iface@Gtk.Editable] properties.
//
// See [func@Gtk.Editable.install_properties] for details on how to
// implement the `GtkEditable` interface.
type EditableProperties int

const (

	// the property id for [property@Gtk.Editable:text]
	EditablePropTextValue EditableProperties = 0
	// the property id for [property@Gtk.Editable:cursor-position]
	EditablePropCursorPositionValue EditableProperties = 1
	// the property id for [property@Gtk.Editable:selection-bound]
	EditablePropSelectionBoundValue EditableProperties = 2
	// the property id for [property@Gtk.Editable:editable]
	EditablePropEditableValue EditableProperties = 3
	// the property id for [property@Gtk.Editable:width-chars]
	EditablePropWidthCharsValue EditableProperties = 4
	// the property id for [property@Gtk.Editable:max-width-chars]
	EditablePropMaxWidthCharsValue EditableProperties = 5
	// the property id for [property@Gtk.Editable:xalign]
	EditablePropXalignValue EditableProperties = 6
	// the property id for [property@Gtk.Editable:enable-undo]
	EditablePropEnableUndoValue EditableProperties = 7
	// the number of properties
	EditableNumPropertiesValue EditableProperties = 8
)

var xEditableDelegateGetProperty func(uintptr, uint, *gobject.Value, uintptr) bool

// Gets a property of the `GtkEditable` delegate for @object.
//
// This is helper function that should be called in the `get_property`
// function of your `GtkEditable` implementation, before handling your
// own properties.
func EditableDelegateGetProperty(ObjectVar *gobject.Object, PropIdVar uint, ValueVar *gobject.Value, PspecVar *gobject.ParamSpec) bool {

	return xEditableDelegateGetProperty(ObjectVar.GoPointer(), PropIdVar, ValueVar, PspecVar.GoPointer())

}

var xEditableDelegateSetProperty func(uintptr, uint, *gobject.Value, uintptr) bool

// Sets a property on the `GtkEditable` delegate for @object.
//
// This is a helper function that should be called in the `set_property`
// function of your `GtkEditable` implementation, before handling your
// own properties.
func EditableDelegateSetProperty(ObjectVar *gobject.Object, PropIdVar uint, ValueVar *gobject.Value, PspecVar *gobject.ParamSpec) bool {

	return xEditableDelegateSetProperty(ObjectVar.GoPointer(), PropIdVar, ValueVar, PspecVar.GoPointer())

}

var xEditableInstallProperties func(*gobject.ObjectClass, uint) uint

// Overrides the `GtkEditable` properties for @class.
//
// This is a helper function that should be called in class_init,
// after installing your own properties.
//
// Note that your class must have "text", "cursor-position",
// "selection-bound", "editable", "width-chars", "max-width-chars",
// "xalign" and "enable-undo" properties for this function to work.
//
// To handle the properties in your set_property and get_property
// functions, you can either use [func@Gtk.Editable.delegate_set_property]
// and [func@Gtk.Editable.delegate_get_property] (if you are using
// a delegate), or remember the @first_prop offset and add it to the
// values in the [enum@Gtk.EditableProperties] enumeration to get the
// property IDs for these properties.
func EditableInstallProperties(ObjectClassVar *gobject.ObjectClass, FirstPropVar uint) uint {

	return xEditableInstallProperties(ObjectClassVar, FirstPropVar)

}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GTK"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	core.PuregoSafeRegister(&xEditableDelegateGetProperty, lib, "gtk_editable_delegate_get_property")
	core.PuregoSafeRegister(&xEditableDelegateSetProperty, lib, "gtk_editable_delegate_set_property")
	core.PuregoSafeRegister(&xEditableInstallProperties, lib, "gtk_editable_install_properties")

	core.PuregoSafeRegister(&XGtkEditableDeleteSelection, lib, "gtk_editable_delete_selection")
	core.PuregoSafeRegister(&XGtkEditableDeleteText, lib, "gtk_editable_delete_text")
	core.PuregoSafeRegister(&XGtkEditableFinishDelegate, lib, "gtk_editable_finish_delegate")
	core.PuregoSafeRegister(&XGtkEditableGetAlignment, lib, "gtk_editable_get_alignment")
	core.PuregoSafeRegister(&XGtkEditableGetChars, lib, "gtk_editable_get_chars")
	core.PuregoSafeRegister(&XGtkEditableGetDelegate, lib, "gtk_editable_get_delegate")
	core.PuregoSafeRegister(&XGtkEditableGetEditable, lib, "gtk_editable_get_editable")
	core.PuregoSafeRegister(&XGtkEditableGetEnableUndo, lib, "gtk_editable_get_enable_undo")
	core.PuregoSafeRegister(&XGtkEditableGetMaxWidthChars, lib, "gtk_editable_get_max_width_chars")
	core.PuregoSafeRegister(&XGtkEditableGetPosition, lib, "gtk_editable_get_position")
	core.PuregoSafeRegister(&XGtkEditableGetSelectionBounds, lib, "gtk_editable_get_selection_bounds")
	core.PuregoSafeRegister(&XGtkEditableGetText, lib, "gtk_editable_get_text")
	core.PuregoSafeRegister(&XGtkEditableGetWidthChars, lib, "gtk_editable_get_width_chars")
	core.PuregoSafeRegister(&XGtkEditableInitDelegate, lib, "gtk_editable_init_delegate")
	core.PuregoSafeRegister(&XGtkEditableInsertText, lib, "gtk_editable_insert_text")
	core.PuregoSafeRegister(&XGtkEditableSelectRegion, lib, "gtk_editable_select_region")
	core.PuregoSafeRegister(&XGtkEditableSetAlignment, lib, "gtk_editable_set_alignment")
	core.PuregoSafeRegister(&XGtkEditableSetEditable, lib, "gtk_editable_set_editable")
	core.PuregoSafeRegister(&XGtkEditableSetEnableUndo, lib, "gtk_editable_set_enable_undo")
	core.PuregoSafeRegister(&XGtkEditableSetMaxWidthChars, lib, "gtk_editable_set_max_width_chars")
	core.PuregoSafeRegister(&XGtkEditableSetPosition, lib, "gtk_editable_set_position")
	core.PuregoSafeRegister(&XGtkEditableSetText, lib, "gtk_editable_set_text")
	core.PuregoSafeRegister(&XGtkEditableSetWidthChars, lib, "gtk_editable_set_width_chars")

}
