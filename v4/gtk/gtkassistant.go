// Package gtk was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gtk

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/gdk"
	"github.com/jwijenbergh/puregotk/v4/gio"
	"github.com/jwijenbergh/puregotk/v4/glib"
	"github.com/jwijenbergh/puregotk/v4/gobject"
	"github.com/jwijenbergh/puregotk/v4/gsk"
)

// Type of callback used to calculate the next page in a `GtkAssistant`.
//
// It’s called both for computing the next page when the user presses the
// “forward” button and for handling the behavior of the “last” button.
//
// See [method@Gtk.Assistant.set_forward_page_func].
type AssistantPageFunc func(int32, uintptr) int32

// Determines the page role inside a `GtkAssistant`.
//
// The role is used to handle buttons sensitivity and visibility.
//
// Note that an assistant needs to end its page flow with a page of type
// %GTK_ASSISTANT_PAGE_CONFIRM, %GTK_ASSISTANT_PAGE_SUMMARY or
// %GTK_ASSISTANT_PAGE_PROGRESS to be correct.
//
// The Cancel button will only be shown if the page isn’t “committed”.
// See gtk_assistant_commit() for details.
type AssistantPageType int

const (

	// The page has regular contents. Both the
	//  Back and forward buttons will be shown.
	AssistantPageContentValue AssistantPageType = 0
	// The page contains an introduction to the
	//  assistant task. Only the Forward button will be shown if there is a
	//   next page.
	AssistantPageIntroValue AssistantPageType = 1
	// The page lets the user confirm or deny the
	//  changes. The Back and Apply buttons will be shown.
	AssistantPageConfirmValue AssistantPageType = 2
	// The page informs the user of the changes
	//  done. Only the Close button will be shown.
	AssistantPageSummaryValue AssistantPageType = 3
	// Used for tasks that take a long time to
	//  complete, blocks the assistant until the page is marked as complete.
	//   Only the back button will be shown.
	AssistantPageProgressValue AssistantPageType = 4
	// Used for when other page types are not
	//  appropriate. No buttons will be shown, and the application must
	//  add its own buttons through gtk_assistant_add_action_widget().
	AssistantPageCustomValue AssistantPageType = 5
)

// `GtkAssistant` is used to represent a complex as a series of steps.
//
// ![An example GtkAssistant](assistant.png)
//
// Each step consists of one or more pages. `GtkAssistant` guides the user
// through the pages, and controls the page flow to collect the data needed
// for the operation.
//
// `GtkAssistant` handles which buttons to show and to make sensitive based
// on page sequence knowledge and the [enum@Gtk.AssistantPageType] of each
// page in addition to state information like the *completed* and *committed*
// page statuses.
//
// If you have a case that doesn’t quite fit in `GtkAssistant`s way of
// handling buttons, you can use the %GTK_ASSISTANT_PAGE_CUSTOM page
// type and handle buttons yourself.
//
// `GtkAssistant` maintains a `GtkAssistantPage` object for each added
// child, which holds additional per-child properties. You
// obtain the `GtkAssistantPage` for a child with [method@Gtk.Assistant.get_page].
//
// # GtkAssistant as GtkBuildable
//
// The `GtkAssistant` implementation of the `GtkBuildable` interface
// exposes the @action_area as internal children with the name
// “action_area”.
//
// To add pages to an assistant in `GtkBuilder`, simply add it as a
// child to the `GtkAssistant` object. If you need to set per-object
// properties, create a `GtkAssistantPage` object explicitly, and
// set the child widget as a property on it.
//
// # CSS nodes
//
// `GtkAssistant` has a single CSS node with the name window and style
// class .assistant.
type Assistant struct {
	Window
}

func AssistantNewFromInternalPtr(ptr uintptr) *Assistant {
	cls := &Assistant{}
	cls.Ptr = ptr
	return cls
}

var xNewAssistant func() uintptr

// Creates a new `GtkAssistant`.
func NewAssistant() *Widget {
	NewAssistantPtr := xNewAssistant()
	if NewAssistantPtr == 0 {
		return nil
	}

	gobject.IncreaseRef(NewAssistantPtr)

	NewAssistantCls := &Widget{}
	NewAssistantCls.Ptr = NewAssistantPtr
	return NewAssistantCls
}

var xAssistantAddActionWidget func(uintptr, uintptr)

// Adds a widget to the action area of a `GtkAssistant`.
func (x *Assistant) AddActionWidget(ChildVar *Widget) {

	xAssistantAddActionWidget(x.GoPointer(), ChildVar.GoPointer())

}

var xAssistantAppendPage func(uintptr, uintptr) int32

// Appends a page to the @assistant.
func (x *Assistant) AppendPage(PageVar *Widget) int32 {

	return xAssistantAppendPage(x.GoPointer(), PageVar.GoPointer())

}

var xAssistantCommit func(uintptr)

// Erases the visited page history.
//
// GTK will then hide the back button on the current page,
// and removes the cancel button from subsequent pages.
//
// Use this when the information provided up to the current
// page is hereafter deemed permanent and cannot be modified
// or undone. For example, showing a progress page to track
// a long-running, unreversible operation after the user has
// clicked apply on a confirmation page.
func (x *Assistant) Commit() {

	xAssistantCommit(x.GoPointer())

}

var xAssistantGetCurrentPage func(uintptr) int32

// Returns the page number of the current page.
func (x *Assistant) GetCurrentPage() int32 {

	return xAssistantGetCurrentPage(x.GoPointer())

}

var xAssistantGetNPages func(uintptr) int32

// Returns the number of pages in the @assistant
func (x *Assistant) GetNPages() int32 {

	return xAssistantGetNPages(x.GoPointer())

}

var xAssistantGetNthPage func(uintptr, int32) uintptr

// Returns the child widget contained in page number @page_num.
func (x *Assistant) GetNthPage(PageNumVar int32) *Widget {

	GetNthPagePtr := xAssistantGetNthPage(x.GoPointer(), PageNumVar)
	if GetNthPagePtr == 0 {
		return nil
	}

	gobject.IncreaseRef(GetNthPagePtr)

	GetNthPageCls := &Widget{}
	GetNthPageCls.Ptr = GetNthPagePtr
	return GetNthPageCls

}

var xAssistantGetPage func(uintptr, uintptr) uintptr

// Returns the `GtkAssistantPage` object for @child.
func (x *Assistant) GetPage(ChildVar *Widget) *AssistantPage {

	GetPagePtr := xAssistantGetPage(x.GoPointer(), ChildVar.GoPointer())
	if GetPagePtr == 0 {
		return nil
	}

	gobject.IncreaseRef(GetPagePtr)

	GetPageCls := &AssistantPage{}
	GetPageCls.Ptr = GetPagePtr
	return GetPageCls

}

var xAssistantGetPageComplete func(uintptr, uintptr) bool

// Gets whether @page is complete.
func (x *Assistant) GetPageComplete(PageVar *Widget) bool {

	return xAssistantGetPageComplete(x.GoPointer(), PageVar.GoPointer())

}

var xAssistantGetPageTitle func(uintptr, uintptr) string

// Gets the title for @page.
func (x *Assistant) GetPageTitle(PageVar *Widget) string {

	return xAssistantGetPageTitle(x.GoPointer(), PageVar.GoPointer())

}

var xAssistantGetPageType func(uintptr, uintptr) AssistantPageType

// Gets the page type of @page.
func (x *Assistant) GetPageType(PageVar *Widget) AssistantPageType {

	return xAssistantGetPageType(x.GoPointer(), PageVar.GoPointer())

}

var xAssistantGetPages func(uintptr) uintptr

// Gets a list model of the assistant pages.
func (x *Assistant) GetPages() *gio.ListModelBase {

	GetPagesPtr := xAssistantGetPages(x.GoPointer())
	if GetPagesPtr == 0 {
		return nil
	}

	GetPagesCls := &gio.ListModelBase{}
	GetPagesCls.Ptr = GetPagesPtr
	return GetPagesCls

}

var xAssistantInsertPage func(uintptr, uintptr, int32) int32

// Inserts a page in the @assistant at a given position.
func (x *Assistant) InsertPage(PageVar *Widget, PositionVar int32) int32 {

	return xAssistantInsertPage(x.GoPointer(), PageVar.GoPointer(), PositionVar)

}

var xAssistantNextPage func(uintptr)

// Navigate to the next page.
//
// It is a programming error to call this function when
// there is no next page.
//
// This function is for use when creating pages of the
// %GTK_ASSISTANT_PAGE_CUSTOM type.
func (x *Assistant) NextPage() {

	xAssistantNextPage(x.GoPointer())

}

var xAssistantPrependPage func(uintptr, uintptr) int32

// Prepends a page to the @assistant.
func (x *Assistant) PrependPage(PageVar *Widget) int32 {

	return xAssistantPrependPage(x.GoPointer(), PageVar.GoPointer())

}

var xAssistantPreviousPage func(uintptr)

// Navigate to the previous visited page.
//
// It is a programming error to call this function when
// no previous page is available.
//
// This function is for use when creating pages of the
// %GTK_ASSISTANT_PAGE_CUSTOM type.
func (x *Assistant) PreviousPage() {

	xAssistantPreviousPage(x.GoPointer())

}

var xAssistantRemoveActionWidget func(uintptr, uintptr)

// Removes a widget from the action area of a `GtkAssistant`.
func (x *Assistant) RemoveActionWidget(ChildVar *Widget) {

	xAssistantRemoveActionWidget(x.GoPointer(), ChildVar.GoPointer())

}

var xAssistantRemovePage func(uintptr, int32)

// Removes the @page_num’s page from @assistant.
func (x *Assistant) RemovePage(PageNumVar int32) {

	xAssistantRemovePage(x.GoPointer(), PageNumVar)

}

var xAssistantSetCurrentPage func(uintptr, int32)

// Switches the page to @page_num.
//
// Note that this will only be necessary in custom buttons,
// as the @assistant flow can be set with
// gtk_assistant_set_forward_page_func().
func (x *Assistant) SetCurrentPage(PageNumVar int32) {

	xAssistantSetCurrentPage(x.GoPointer(), PageNumVar)

}

var xAssistantSetForwardPageFunc func(uintptr, uintptr, uintptr, uintptr)

// Sets the page forwarding function to be @page_func.
//
// This function will be used to determine what will be
// the next page when the user presses the forward button.
// Setting @page_func to %NULL will make the assistant to
// use the default forward function, which just goes to the
// next visible page.
func (x *Assistant) SetForwardPageFunc(PageFuncVar AssistantPageFunc, DataVar uintptr, DestroyVar glib.DestroyNotify) {

	xAssistantSetForwardPageFunc(x.GoPointer(), purego.NewCallback(PageFuncVar), DataVar, purego.NewCallback(DestroyVar))

}

var xAssistantSetPageComplete func(uintptr, uintptr, bool)

// Sets whether @page contents are complete.
//
// This will make @assistant update the buttons state
// to be able to continue the task.
func (x *Assistant) SetPageComplete(PageVar *Widget, CompleteVar bool) {

	xAssistantSetPageComplete(x.GoPointer(), PageVar.GoPointer(), CompleteVar)

}

var xAssistantSetPageTitle func(uintptr, uintptr, string)

// Sets a title for @page.
//
// The title is displayed in the header area of the assistant
// when @page is the current page.
func (x *Assistant) SetPageTitle(PageVar *Widget, TitleVar string) {

	xAssistantSetPageTitle(x.GoPointer(), PageVar.GoPointer(), TitleVar)

}

var xAssistantSetPageType func(uintptr, uintptr, AssistantPageType)

// Sets the page type for @page.
//
// The page type determines the page behavior in the @assistant.
func (x *Assistant) SetPageType(PageVar *Widget, TypeVar AssistantPageType) {

	xAssistantSetPageType(x.GoPointer(), PageVar.GoPointer(), TypeVar)

}

var xAssistantUpdateButtonsState func(uintptr)

// Forces @assistant to recompute the buttons state.
//
// GTK automatically takes care of this in most situations,
// e.g. when the user goes to a different page, or when the
// visibility or completeness of a page changes.
//
// One situation where it can be necessary to call this
// function is when changing a value on the current page
// affects the future page flow of the assistant.
func (x *Assistant) UpdateButtonsState() {

	xAssistantUpdateButtonsState(x.GoPointer())

}

func (c *Assistant) GoPointer() uintptr {
	return c.Ptr
}

func (c *Assistant) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// Emitted when the apply button is clicked.
//
// The default behavior of the `GtkAssistant` is to switch to the page
// after the current page, unless the current page is the last one.
//
// A handler for the ::apply signal should carry out the actions for
// which the wizard has collected data. If the action takes a long time
// to complete, you might consider putting a page of type
// %GTK_ASSISTANT_PAGE_PROGRESS after the confirmation page and handle
// this operation within the [signal@Gtk.Assistant::prepare] signal of
// the progress page.
func (x *Assistant) ConnectApply(cb func(Assistant)) {
	fcb := func(clsPtr uintptr) {
		fa := Assistant{}
		fa.Ptr = clsPtr

		cb(fa)

	}
	gobject.ObjectConnect(x.GoPointer(), "signal::apply", purego.NewCallback(fcb))
}

// Emitted when then the cancel button is clicked.
func (x *Assistant) ConnectCancel(cb func(Assistant)) {
	fcb := func(clsPtr uintptr) {
		fa := Assistant{}
		fa.Ptr = clsPtr

		cb(fa)

	}
	gobject.ObjectConnect(x.GoPointer(), "signal::cancel", purego.NewCallback(fcb))
}

// Emitted either when the close button of a summary page is clicked,
// or when the apply button in the last page in the flow (of type
// %GTK_ASSISTANT_PAGE_CONFIRM) is clicked.
func (x *Assistant) ConnectClose(cb func(Assistant)) {
	fcb := func(clsPtr uintptr) {
		fa := Assistant{}
		fa.Ptr = clsPtr

		cb(fa)

	}
	gobject.ObjectConnect(x.GoPointer(), "signal::close", purego.NewCallback(fcb))
}

// The action signal for the Escape binding.
func (x *Assistant) ConnectEscape(cb func(Assistant)) {
	fcb := func(clsPtr uintptr) {
		fa := Assistant{}
		fa.Ptr = clsPtr

		cb(fa)

	}
	gobject.ObjectConnect(x.GoPointer(), "signal::escape", purego.NewCallback(fcb))
}

// Emitted when a new page is set as the assistant's current page,
// before making the new page visible.
//
// A handler for this signal can do any preparations which are
// necessary before showing @page.
func (x *Assistant) ConnectPrepare(cb func(Assistant, uintptr)) {
	fcb := func(clsPtr uintptr, PageVarp uintptr) {
		fa := Assistant{}
		fa.Ptr = clsPtr

		cb(fa, PageVarp)

	}
	gobject.ObjectConnect(x.GoPointer(), "signal::prepare", purego.NewCallback(fcb))
}

// Retrieves the `GtkAccessibleRole` for the given `GtkAccessible`.
func (x *Assistant) GetAccessibleRole() AccessibleRole {

	return XGtkAccessibleGetAccessibleRole(x.GoPointer())

}

// Resets the accessible @property to its default value.
func (x *Assistant) ResetProperty(PropertyVar AccessibleProperty) {

	XGtkAccessibleResetProperty(x.GoPointer(), PropertyVar)

}

// Resets the accessible @relation to its default value.
func (x *Assistant) ResetRelation(RelationVar AccessibleRelation) {

	XGtkAccessibleResetRelation(x.GoPointer(), RelationVar)

}

// Resets the accessible @state to its default value.
func (x *Assistant) ResetState(StateVar AccessibleState) {

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
func (x *Assistant) UpdateProperty(FirstPropertyVar AccessibleProperty, varArgs ...interface{}) {

	XGtkAccessibleUpdateProperty(x.GoPointer(), FirstPropertyVar, varArgs...)

}

// Updates an array of accessible properties.
//
// This function should be called by `GtkWidget` types whenever an accessible
// property change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *Assistant) UpdatePropertyValue(NPropertiesVar int32, PropertiesVar uintptr, ValuesVar uintptr) {

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
func (x *Assistant) UpdateRelation(FirstRelationVar AccessibleRelation, varArgs ...interface{}) {

	XGtkAccessibleUpdateRelation(x.GoPointer(), FirstRelationVar, varArgs...)

}

// Updates an array of accessible relations.
//
// This function should be called by `GtkWidget` types whenever an accessible
// relation change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *Assistant) UpdateRelationValue(NRelationsVar int32, RelationsVar uintptr, ValuesVar uintptr) {

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
func (x *Assistant) UpdateState(FirstStateVar AccessibleState, varArgs ...interface{}) {

	XGtkAccessibleUpdateState(x.GoPointer(), FirstStateVar, varArgs...)

}

// Updates an array of accessible states.
//
// This function should be called by `GtkWidget` types whenever an accessible
// state change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *Assistant) UpdateStateValue(NStatesVar int32, StatesVar uintptr, ValuesVar uintptr) {

	XGtkAccessibleUpdateStateValue(x.GoPointer(), NStatesVar, StatesVar, ValuesVar)

}

// Gets the ID of the @buildable object.
//
// `GtkBuilder` sets the name based on the ID attribute
// of the &lt;object&gt; tag used to construct the @buildable.
func (x *Assistant) GetBuildableId() string {

	return XGtkBuildableGetBuildableId(x.GoPointer())

}

// Returns the renderer that is used for this `GtkNative`.
func (x *Assistant) GetRenderer() *gsk.Renderer {

	GetRendererPtr := XGtkNativeGetRenderer(x.GoPointer())
	if GetRendererPtr == 0 {
		return nil
	}

	gobject.IncreaseRef(GetRendererPtr)

	GetRendererCls := &gsk.Renderer{}
	GetRendererCls.Ptr = GetRendererPtr
	return GetRendererCls

}

// Returns the surface of this `GtkNative`.
func (x *Assistant) GetSurface() *gdk.Surface {

	GetSurfacePtr := XGtkNativeGetSurface(x.GoPointer())
	if GetSurfacePtr == 0 {
		return nil
	}

	gobject.IncreaseRef(GetSurfacePtr)

	GetSurfaceCls := &gdk.Surface{}
	GetSurfaceCls.Ptr = GetSurfacePtr
	return GetSurfaceCls

}

// Retrieves the surface transform of @self.
//
// This is the translation from @self's surface coordinates into
// @self's widget coordinates.
func (x *Assistant) GetSurfaceTransform(XVar float64, YVar float64) {

	XGtkNativeGetSurfaceTransform(x.GoPointer(), XVar, YVar)

}

// Realizes a `GtkNative`.
//
// This should only be used by subclasses.
func (x *Assistant) Realize() {

	XGtkNativeRealize(x.GoPointer())

}

// Unrealizes a `GtkNative`.
//
// This should only be used by subclasses.
func (x *Assistant) Unrealize() {

	XGtkNativeUnrealize(x.GoPointer())

}

// Returns the display that this `GtkRoot` is on.
func (x *Assistant) GetDisplay() *gdk.Display {

	GetDisplayPtr := XGtkRootGetDisplay(x.GoPointer())
	if GetDisplayPtr == 0 {
		return nil
	}

	gobject.IncreaseRef(GetDisplayPtr)

	GetDisplayCls := &gdk.Display{}
	GetDisplayCls.Ptr = GetDisplayPtr
	return GetDisplayCls

}

// Retrieves the current focused widget within the root.
//
// Note that this is the widget that would have the focus
// if the root is active; if the root is not focused then
// `gtk_widget_has_focus (widget)` will be %FALSE for the
// widget.
func (x *Assistant) GetFocus() *Widget {

	GetFocusPtr := XGtkRootGetFocus(x.GoPointer())
	if GetFocusPtr == 0 {
		return nil
	}

	gobject.IncreaseRef(GetFocusPtr)

	GetFocusCls := &Widget{}
	GetFocusCls.Ptr = GetFocusPtr
	return GetFocusCls

}

// If @focus is not the current focus widget, and is focusable, sets
// it as the focus widget for the root.
//
// If @focus is %NULL, unsets the focus widget for the root.
//
// To set the focus to a particular widget in the root, it is usually
// more convenient to use [method@Gtk.Widget.grab_focus] instead of
// this function.
func (x *Assistant) SetFocus(FocusVar *Widget) {

	XGtkRootSetFocus(x.GoPointer(), FocusVar.GoPointer())

}

// `GtkAssistantPage` is an auxiliary object used by `GtkAssistant.
type AssistantPage struct {
	gobject.Object
}

func AssistantPageNewFromInternalPtr(ptr uintptr) *AssistantPage {
	cls := &AssistantPage{}
	cls.Ptr = ptr
	return cls
}

var xAssistantPageGetChild func(uintptr) uintptr

// Returns the child to which @page belongs.
func (x *AssistantPage) GetChild() *Widget {

	GetChildPtr := xAssistantPageGetChild(x.GoPointer())
	if GetChildPtr == 0 {
		return nil
	}

	gobject.IncreaseRef(GetChildPtr)

	GetChildCls := &Widget{}
	GetChildCls.Ptr = GetChildPtr
	return GetChildCls

}

func (c *AssistantPage) GoPointer() uintptr {
	return c.Ptr
}

func (c *AssistantPage) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GTK"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xNewAssistant, lib, "gtk_assistant_new")

	core.PuregoSafeRegister(&xAssistantAddActionWidget, lib, "gtk_assistant_add_action_widget")
	core.PuregoSafeRegister(&xAssistantAppendPage, lib, "gtk_assistant_append_page")
	core.PuregoSafeRegister(&xAssistantCommit, lib, "gtk_assistant_commit")
	core.PuregoSafeRegister(&xAssistantGetCurrentPage, lib, "gtk_assistant_get_current_page")
	core.PuregoSafeRegister(&xAssistantGetNPages, lib, "gtk_assistant_get_n_pages")
	core.PuregoSafeRegister(&xAssistantGetNthPage, lib, "gtk_assistant_get_nth_page")
	core.PuregoSafeRegister(&xAssistantGetPage, lib, "gtk_assistant_get_page")
	core.PuregoSafeRegister(&xAssistantGetPageComplete, lib, "gtk_assistant_get_page_complete")
	core.PuregoSafeRegister(&xAssistantGetPageTitle, lib, "gtk_assistant_get_page_title")
	core.PuregoSafeRegister(&xAssistantGetPageType, lib, "gtk_assistant_get_page_type")
	core.PuregoSafeRegister(&xAssistantGetPages, lib, "gtk_assistant_get_pages")
	core.PuregoSafeRegister(&xAssistantInsertPage, lib, "gtk_assistant_insert_page")
	core.PuregoSafeRegister(&xAssistantNextPage, lib, "gtk_assistant_next_page")
	core.PuregoSafeRegister(&xAssistantPrependPage, lib, "gtk_assistant_prepend_page")
	core.PuregoSafeRegister(&xAssistantPreviousPage, lib, "gtk_assistant_previous_page")
	core.PuregoSafeRegister(&xAssistantRemoveActionWidget, lib, "gtk_assistant_remove_action_widget")
	core.PuregoSafeRegister(&xAssistantRemovePage, lib, "gtk_assistant_remove_page")
	core.PuregoSafeRegister(&xAssistantSetCurrentPage, lib, "gtk_assistant_set_current_page")
	core.PuregoSafeRegister(&xAssistantSetForwardPageFunc, lib, "gtk_assistant_set_forward_page_func")
	core.PuregoSafeRegister(&xAssistantSetPageComplete, lib, "gtk_assistant_set_page_complete")
	core.PuregoSafeRegister(&xAssistantSetPageTitle, lib, "gtk_assistant_set_page_title")
	core.PuregoSafeRegister(&xAssistantSetPageType, lib, "gtk_assistant_set_page_type")
	core.PuregoSafeRegister(&xAssistantUpdateButtonsState, lib, "gtk_assistant_update_buttons_state")

	core.PuregoSafeRegister(&xAssistantPageGetChild, lib, "gtk_assistant_page_get_child")

}
