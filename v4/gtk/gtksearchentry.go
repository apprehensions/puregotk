// Package gtk was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gtk

import (
	"unsafe"

	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/glib"
	"github.com/jwijenbergh/puregotk/v4/gobject"
)

// `GtkSearchEntry` is an entry widget that has been tailored for use
// as a search entry.
//
// The main API for interacting with a `GtkSearchEntry` as entry
// is the `GtkEditable` interface.
//
// ![An example GtkSearchEntry](search-entry.png)
//
// It will show an inactive symbolic “find” icon when the search
// entry is empty, and a symbolic “clear” icon when there is text.
// Clicking on the “clear” icon will empty the search entry.
//
// To make filtering appear more reactive, it is a good idea to
// not react to every change in the entry text immediately, but
// only after a short delay. To support this, `GtkSearchEntry`
// emits the [signal@Gtk.SearchEntry::search-changed] signal which
// can be used instead of the [signal@Gtk.Editable::changed] signal.
//
// The [signal@Gtk.SearchEntry::previous-match],
// [signal@Gtk.SearchEntry::next-match] and
// [signal@Gtk.SearchEntry::stop-search] signals can be used to
// implement moving between search results and ending the search.
//
// Often, `GtkSearchEntry` will be fed events by means of being
// placed inside a [class@Gtk.SearchBar]. If that is not the case,
// you can use [method@Gtk.SearchEntry.set_key_capture_widget] to
// let it capture key input from another widget.
//
// `GtkSearchEntry` provides only minimal API and should be used with
// the [iface@Gtk.Editable] API.
//
// ## CSS Nodes
//
// ```
// entry.search
// ╰── text
// ```
//
// `GtkSearchEntry` has a single CSS node with name entry that carries
// a `.search` style class, and the text node is a child of that.
//
// ## Accessibility
//
// `GtkSearchEntry` uses the %GTK_ACCESSIBLE_ROLE_SEARCH_BOX role.
type SearchEntry struct {
	Widget
}

func SearchEntryNewFromInternalPtr(ptr uintptr) *SearchEntry {
	cls := &SearchEntry{}
	cls.Ptr = ptr
	return cls
}

var xNewSearchEntry func() uintptr

// Creates a `GtkSearchEntry`.
func NewSearchEntry() *SearchEntry {
	var cls *SearchEntry

	cret := xNewSearchEntry()

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &SearchEntry{}
	cls.Ptr = cret
	return cls
}

var xSearchEntryGetKeyCaptureWidget func(uintptr) uintptr

// Gets the widget that @entry is capturing key events from.
func (x *SearchEntry) GetKeyCaptureWidget() *Widget {
	var cls *Widget

	cret := xSearchEntryGetKeyCaptureWidget(x.GoPointer())

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &Widget{}
	cls.Ptr = cret
	return cls
}

var xSearchEntryGetSearchDelay func(uintptr) uint

// Get the delay to be used between the last keypress and the
// [signal@Gtk.SearchEntry::search-changed] signal being emitted.
func (x *SearchEntry) GetSearchDelay() uint {

	cret := xSearchEntryGetSearchDelay(x.GoPointer())
	return cret
}

var xSearchEntrySetKeyCaptureWidget func(uintptr, uintptr)

// Sets @widget as the widget that @entry will capture key
// events from.
//
// Key events are consumed by the search entry to start or
// continue a search.
//
// If the entry is part of a `GtkSearchBar`, it is preferable
// to call [method@Gtk.SearchBar.set_key_capture_widget] instead,
// which will reveal the entry in addition to triggering the
// search entry.
//
// Note that despite the name of this function, the events
// are only 'captured' in the bubble phase, which means that
// editable child widgets of @widget will receive text input
// before it gets captured. If that is not desired, you can
// capture and forward the events yourself with
// [method@Gtk.EventControllerKey.forward].
func (x *SearchEntry) SetKeyCaptureWidget(WidgetVar *Widget) {

	xSearchEntrySetKeyCaptureWidget(x.GoPointer(), WidgetVar.GoPointer())

}

var xSearchEntrySetSearchDelay func(uintptr, uint)

// Set the delay to be used between the last keypress and the
// [signal@Gtk.SearchEntry::search-changed] signal being emitted.
func (x *SearchEntry) SetSearchDelay(DelayVar uint) {

	xSearchEntrySetSearchDelay(x.GoPointer(), DelayVar)

}

func (c *SearchEntry) GoPointer() uintptr {
	return c.Ptr
}

func (c *SearchEntry) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// Emitted when the entry is activated.
//
// The keybindings for this signal are all forms of the Enter key.
func (x *SearchEntry) ConnectActivate(cb *func(SearchEntry)) uint32 {
	cbPtr := uintptr(unsafe.Pointer(cb))
	if cbRefPtr, ok := glib.GetCallback(cbPtr); ok {
		return gobject.SignalConnect(x.GoPointer(), "activate", cbRefPtr)
	}

	fcb := func(clsPtr uintptr) {
		fa := SearchEntry{}
		fa.Ptr = clsPtr
		cbFn := *cb

		cbFn(fa)

	}
	cbRefPtr := purego.NewCallback(fcb)
	glib.SaveCallback(cbPtr, cbRefPtr)
	return gobject.SignalConnect(x.GoPointer(), "activate", cbRefPtr)
}

// Emitted when the user initiates a move to the next match
// for the current search string.
//
// This is a [keybinding signal](class.SignalAction.html).
//
// Applications should connect to it, to implement moving
// between matches.
//
// The default bindings for this signal is Ctrl-g.
func (x *SearchEntry) ConnectNextMatch(cb *func(SearchEntry)) uint32 {
	cbPtr := uintptr(unsafe.Pointer(cb))
	if cbRefPtr, ok := glib.GetCallback(cbPtr); ok {
		return gobject.SignalConnect(x.GoPointer(), "next-match", cbRefPtr)
	}

	fcb := func(clsPtr uintptr) {
		fa := SearchEntry{}
		fa.Ptr = clsPtr
		cbFn := *cb

		cbFn(fa)

	}
	cbRefPtr := purego.NewCallback(fcb)
	glib.SaveCallback(cbPtr, cbRefPtr)
	return gobject.SignalConnect(x.GoPointer(), "next-match", cbRefPtr)
}

// Emitted when the user initiates a move to the previous match
// for the current search string.
//
// This is a [keybinding signal](class.SignalAction.html).
//
// Applications should connect to it, to implement moving
// between matches.
//
// The default bindings for this signal is Ctrl-Shift-g.
func (x *SearchEntry) ConnectPreviousMatch(cb *func(SearchEntry)) uint32 {
	cbPtr := uintptr(unsafe.Pointer(cb))
	if cbRefPtr, ok := glib.GetCallback(cbPtr); ok {
		return gobject.SignalConnect(x.GoPointer(), "previous-match", cbRefPtr)
	}

	fcb := func(clsPtr uintptr) {
		fa := SearchEntry{}
		fa.Ptr = clsPtr
		cbFn := *cb

		cbFn(fa)

	}
	cbRefPtr := purego.NewCallback(fcb)
	glib.SaveCallback(cbPtr, cbRefPtr)
	return gobject.SignalConnect(x.GoPointer(), "previous-match", cbRefPtr)
}

// Emitted with a delay. The length of the delay can be
// changed with the [property@Gtk.SearchEntry:search-delay]
// property.
func (x *SearchEntry) ConnectSearchChanged(cb *func(SearchEntry)) uint32 {
	cbPtr := uintptr(unsafe.Pointer(cb))
	if cbRefPtr, ok := glib.GetCallback(cbPtr); ok {
		return gobject.SignalConnect(x.GoPointer(), "search-changed", cbRefPtr)
	}

	fcb := func(clsPtr uintptr) {
		fa := SearchEntry{}
		fa.Ptr = clsPtr
		cbFn := *cb

		cbFn(fa)

	}
	cbRefPtr := purego.NewCallback(fcb)
	glib.SaveCallback(cbPtr, cbRefPtr)
	return gobject.SignalConnect(x.GoPointer(), "search-changed", cbRefPtr)
}

// Emitted when the user initiated a search on the entry.
func (x *SearchEntry) ConnectSearchStarted(cb *func(SearchEntry)) uint32 {
	cbPtr := uintptr(unsafe.Pointer(cb))
	if cbRefPtr, ok := glib.GetCallback(cbPtr); ok {
		return gobject.SignalConnect(x.GoPointer(), "search-started", cbRefPtr)
	}

	fcb := func(clsPtr uintptr) {
		fa := SearchEntry{}
		fa.Ptr = clsPtr
		cbFn := *cb

		cbFn(fa)

	}
	cbRefPtr := purego.NewCallback(fcb)
	glib.SaveCallback(cbPtr, cbRefPtr)
	return gobject.SignalConnect(x.GoPointer(), "search-started", cbRefPtr)
}

// Emitted when the user stops a search via keyboard input.
//
// This is a [keybinding signal](class.SignalAction.html).
//
// Applications should connect to it, to implement hiding
// the search entry in this case.
//
// The default bindings for this signal is Escape.
func (x *SearchEntry) ConnectStopSearch(cb *func(SearchEntry)) uint32 {
	cbPtr := uintptr(unsafe.Pointer(cb))
	if cbRefPtr, ok := glib.GetCallback(cbPtr); ok {
		return gobject.SignalConnect(x.GoPointer(), "stop-search", cbRefPtr)
	}

	fcb := func(clsPtr uintptr) {
		fa := SearchEntry{}
		fa.Ptr = clsPtr
		cbFn := *cb

		cbFn(fa)

	}
	cbRefPtr := purego.NewCallback(fcb)
	glib.SaveCallback(cbPtr, cbRefPtr)
	return gobject.SignalConnect(x.GoPointer(), "stop-search", cbRefPtr)
}

// Retrieves the `GtkAccessibleRole` for the given `GtkAccessible`.
func (x *SearchEntry) GetAccessibleRole() AccessibleRole {

	cret := XGtkAccessibleGetAccessibleRole(x.GoPointer())
	return cret
}

// Resets the accessible @property to its default value.
func (x *SearchEntry) ResetProperty(PropertyVar AccessibleProperty) {

	XGtkAccessibleResetProperty(x.GoPointer(), PropertyVar)

}

// Resets the accessible @relation to its default value.
func (x *SearchEntry) ResetRelation(RelationVar AccessibleRelation) {

	XGtkAccessibleResetRelation(x.GoPointer(), RelationVar)

}

// Resets the accessible @state to its default value.
func (x *SearchEntry) ResetState(StateVar AccessibleState) {

	XGtkAccessibleResetState(x.GoPointer(), StateVar)

}

// Updates a list of accessible properties.
//
// See the [enum@Gtk.AccessibleProperty] documentation for the
// value types of accessible properties.
//
// This function should be called by `GtkWidget` types whenever
// an accessible property change must be communicated to assistive
// technologies.
//
// Example:
// ```c
// value = gtk_adjustment_get_value (adjustment);
// gtk_accessible_update_property (GTK_ACCESSIBLE (spin_button),
//
//	GTK_ACCESSIBLE_PROPERTY_VALUE_NOW, value,
//	-1);
//
// ```
func (x *SearchEntry) UpdateProperty(FirstPropertyVar AccessibleProperty, varArgs ...interface{}) {

	XGtkAccessibleUpdateProperty(x.GoPointer(), FirstPropertyVar, varArgs...)

}

// Updates an array of accessible properties.
//
// This function should be called by `GtkWidget` types whenever an accessible
// property change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *SearchEntry) UpdatePropertyValue(NPropertiesVar int, PropertiesVar uintptr, ValuesVar uintptr) {

	XGtkAccessibleUpdatePropertyValue(x.GoPointer(), NPropertiesVar, PropertiesVar, ValuesVar)

}

// Updates a list of accessible relations.
//
// This function should be called by `GtkWidget` types whenever an accessible
// relation change must be communicated to assistive technologies.
//
// If the [enum@Gtk.AccessibleRelation] requires a list of references,
// you should pass each reference individually, followed by %NULL, e.g.
//
// ```c
// gtk_accessible_update_relation (accessible,
//
//	GTK_ACCESSIBLE_RELATION_CONTROLS,
//	  ref1, NULL,
//	GTK_ACCESSIBLE_RELATION_LABELLED_BY,
//	  ref1, ref2, ref3, NULL,
//	-1);
//
// ```
func (x *SearchEntry) UpdateRelation(FirstRelationVar AccessibleRelation, varArgs ...interface{}) {

	XGtkAccessibleUpdateRelation(x.GoPointer(), FirstRelationVar, varArgs...)

}

// Updates an array of accessible relations.
//
// This function should be called by `GtkWidget` types whenever an accessible
// relation change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *SearchEntry) UpdateRelationValue(NRelationsVar int, RelationsVar uintptr, ValuesVar uintptr) {

	XGtkAccessibleUpdateRelationValue(x.GoPointer(), NRelationsVar, RelationsVar, ValuesVar)

}

// Updates a list of accessible states. See the [enum@Gtk.AccessibleState]
// documentation for the value types of accessible states.
//
// This function should be called by `GtkWidget` types whenever an accessible
// state change must be communicated to assistive technologies.
//
// Example:
// ```c
// value = GTK_ACCESSIBLE_TRISTATE_MIXED;
// gtk_accessible_update_state (GTK_ACCESSIBLE (check_button),
//
//	GTK_ACCESSIBLE_STATE_CHECKED, value,
//	-1);
//
// ```
func (x *SearchEntry) UpdateState(FirstStateVar AccessibleState, varArgs ...interface{}) {

	XGtkAccessibleUpdateState(x.GoPointer(), FirstStateVar, varArgs...)

}

// Updates an array of accessible states.
//
// This function should be called by `GtkWidget` types whenever an accessible
// state change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *SearchEntry) UpdateStateValue(NStatesVar int, StatesVar uintptr, ValuesVar uintptr) {

	XGtkAccessibleUpdateStateValue(x.GoPointer(), NStatesVar, StatesVar, ValuesVar)

}

// Gets the ID of the @buildable object.
//
// `GtkBuilder` sets the name based on the ID attribute
// of the &lt;object&gt; tag used to construct the @buildable.
func (x *SearchEntry) GetBuildableId() string {

	cret := XGtkBuildableGetBuildableId(x.GoPointer())
	return cret
}

// Deletes the currently selected text of the editable.
//
// This call doesn’t do anything if there is no selected text.
func (x *SearchEntry) DeleteSelection() {

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
func (x *SearchEntry) DeleteText(StartPosVar int, EndPosVar int) {

	XGtkEditableDeleteText(x.GoPointer(), StartPosVar, EndPosVar)

}

// Undoes the setup done by [method@Gtk.Editable.init_delegate].
//
// This is a helper function that should be called from dispose,
// before removing the delegate object.
func (x *SearchEntry) FinishDelegate() {

	XGtkEditableFinishDelegate(x.GoPointer())

}

// Gets the alignment of the editable.
func (x *SearchEntry) GetAlignment() float32 {

	cret := XGtkEditableGetAlignment(x.GoPointer())
	return cret
}

// Retrieves a sequence of characters.
//
// The characters that are retrieved are those characters at positions
// from @start_pos up to, but not including @end_pos. If @end_pos is negative,
// then the characters retrieved are those characters from @start_pos to
// the end of the text.
//
// Note that positions are specified in characters, not bytes.
func (x *SearchEntry) GetChars(StartPosVar int, EndPosVar int) string {

	cret := XGtkEditableGetChars(x.GoPointer(), StartPosVar, EndPosVar)
	return cret
}

// Gets the `GtkEditable` that @editable is delegating its
// implementation to.
//
// Typically, the delegate is a [class@Gtk.Text] widget.
func (x *SearchEntry) GetDelegate() *EditableBase {
	var cls *EditableBase

	cret := XGtkEditableGetDelegate(x.GoPointer())

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &EditableBase{}
	cls.Ptr = cret
	return cls
}

// Retrieves whether @editable is editable.
func (x *SearchEntry) GetEditable() bool {

	cret := XGtkEditableGetEditable(x.GoPointer())
	return cret
}

// Gets if undo/redo actions are enabled for @editable
func (x *SearchEntry) GetEnableUndo() bool {

	cret := XGtkEditableGetEnableUndo(x.GoPointer())
	return cret
}

// Retrieves the desired maximum width of @editable, in characters.
func (x *SearchEntry) GetMaxWidthChars() int {

	cret := XGtkEditableGetMaxWidthChars(x.GoPointer())
	return cret
}

// Retrieves the current position of the cursor relative
// to the start of the content of the editable.
//
// Note that this position is in characters, not in bytes.
func (x *SearchEntry) GetPosition() int {

	cret := XGtkEditableGetPosition(x.GoPointer())
	return cret
}

// Retrieves the selection bound of the editable.
//
// @start_pos will be filled with the start of the selection and
// @end_pos with end. If no text was selected both will be identical
// and %FALSE will be returned.
//
// Note that positions are specified in characters, not bytes.
func (x *SearchEntry) GetSelectionBounds(StartPosVar int, EndPosVar int) bool {

	cret := XGtkEditableGetSelectionBounds(x.GoPointer(), StartPosVar, EndPosVar)
	return cret
}

// Retrieves the contents of @editable.
//
// The returned string is owned by GTK and must not be modified or freed.
func (x *SearchEntry) GetText() string {

	cret := XGtkEditableGetText(x.GoPointer())
	return cret
}

// Gets the number of characters of space reserved
// for the contents of the editable.
func (x *SearchEntry) GetWidthChars() int {

	cret := XGtkEditableGetWidthChars(x.GoPointer())
	return cret
}

// Sets up a delegate for `GtkEditable`.
//
// This is assuming that the get_delegate vfunc in the `GtkEditable`
// interface has been set up for the @editable's type.
//
// This is a helper function that should be called in instance init,
// after creating the delegate object.
func (x *SearchEntry) InitDelegate() {

	XGtkEditableInitDelegate(x.GoPointer())

}

// Inserts @length bytes of @text into the contents of the
// widget, at position @position.
//
// Note that the position is in characters, not in bytes.
// The function updates @position to point after the newly
// inserted text.
func (x *SearchEntry) InsertText(TextVar string, LengthVar int, PositionVar int) {

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
func (x *SearchEntry) SelectRegion(StartPosVar int, EndPosVar int) {

	XGtkEditableSelectRegion(x.GoPointer(), StartPosVar, EndPosVar)

}

// Sets the alignment for the contents of the editable.
//
// This controls the horizontal positioning of the contents when
// the displayed text is shorter than the width of the editable.
func (x *SearchEntry) SetAlignment(XalignVar float32) {

	XGtkEditableSetAlignment(x.GoPointer(), XalignVar)

}

// Determines if the user can edit the text in the editable widget.
func (x *SearchEntry) SetEditable(IsEditableVar bool) {

	XGtkEditableSetEditable(x.GoPointer(), IsEditableVar)

}

// If enabled, changes to @editable will be saved for undo/redo
// actions.
//
// This results in an additional copy of text changes and are not
// stored in secure memory. As such, undo is forcefully disabled
// when [property@Gtk.Text:visibility] is set to %FALSE.
func (x *SearchEntry) SetEnableUndo(EnableUndoVar bool) {

	XGtkEditableSetEnableUndo(x.GoPointer(), EnableUndoVar)

}

// Sets the desired maximum width in characters of @editable.
func (x *SearchEntry) SetMaxWidthChars(NCharsVar int) {

	XGtkEditableSetMaxWidthChars(x.GoPointer(), NCharsVar)

}

// Sets the cursor position in the editable to the given value.
//
// The cursor is displayed before the character with the given (base 0)
// index in the contents of the editable. The value must be less than
// or equal to the number of characters in the editable. A value of -1
// indicates that the position should be set after the last character
// of the editable. Note that @position is in characters, not in bytes.
func (x *SearchEntry) SetPosition(PositionVar int) {

	XGtkEditableSetPosition(x.GoPointer(), PositionVar)

}

// Sets the text in the editable to the given value.
//
// This is replacing the current contents.
func (x *SearchEntry) SetText(TextVar string) {

	XGtkEditableSetText(x.GoPointer(), TextVar)

}

// Changes the size request of the editable to be about the
// right size for @n_chars characters.
//
// Note that it changes the size request, the size can still
// be affected by how you pack the widget into containers.
// If @n_chars is -1, the size reverts to the default size.
func (x *SearchEntry) SetWidthChars(NCharsVar int) {

	XGtkEditableSetWidthChars(x.GoPointer(), NCharsVar)

}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GTK"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xNewSearchEntry, lib, "gtk_search_entry_new")

	core.PuregoSafeRegister(&xSearchEntryGetKeyCaptureWidget, lib, "gtk_search_entry_get_key_capture_widget")
	core.PuregoSafeRegister(&xSearchEntryGetSearchDelay, lib, "gtk_search_entry_get_search_delay")
	core.PuregoSafeRegister(&xSearchEntrySetKeyCaptureWidget, lib, "gtk_search_entry_set_key_capture_widget")
	core.PuregoSafeRegister(&xSearchEntrySetSearchDelay, lib, "gtk_search_entry_set_search_delay")

}
