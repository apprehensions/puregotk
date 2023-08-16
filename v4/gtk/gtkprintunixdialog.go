// Package gtk was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gtk

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/gdk"
	"github.com/jwijenbergh/puregotk/v4/gobject"
	"github.com/jwijenbergh/puregotk/v4/gsk"
)

// `GtkPrintUnixDialog` implements a print dialog for platforms
// which don’t provide a native print dialog, like Unix.
//
// ![An example GtkPrintUnixDialog](printdialog.png)
//
// It can be used very much like any other GTK dialog, at the cost of
// the portability offered by the high-level printing API with
// [class@Gtk.PrintOperation].
//
// In order to print something with `GtkPrintUnixDialog`, you need to
// use [method@Gtk.PrintUnixDialog.get_selected_printer] to obtain a
// [class@Gtk.Printer] object and use it to construct a [class@Gtk.PrintJob]
// using [ctor@Gtk.PrintJob.new].
//
// `GtkPrintUnixDialog` uses the following response values:
//
// - %GTK_RESPONSE_OK: for the “Print” button
// - %GTK_RESPONSE_APPLY: for the “Preview” button
// - %GTK_RESPONSE_CANCEL: for the “Cancel” button
//
// # GtkPrintUnixDialog as GtkBuildable
//
// The `GtkPrintUnixDialog` implementation of the `GtkBuildable` interface
// exposes its @notebook internal children with the name “notebook”.
//
// An example of a `GtkPrintUnixDialog` UI definition fragment:
//
// ```xml
// &lt;object class="GtkPrintUnixDialog" id="dialog1"&gt;
//
//	&lt;child internal-child="notebook"&gt;
//	  &lt;object class="GtkNotebook" id="notebook"&gt;
//	    &lt;child&gt;
//	      &lt;object type="GtkNotebookPage"&gt;
//	        &lt;property name="tab_expand"&gt;False&lt;/property&gt;
//	        &lt;property name="tab_fill"&gt;False&lt;/property&gt;
//	        &lt;property name="tab"&gt;
//	          &lt;object class="GtkLabel" id="tablabel"&gt;
//	            &lt;property name="label"&gt;Tab label&lt;/property&gt;
//	          &lt;/object&gt;
//	        &lt;/property&gt;
//	        &lt;property name="child"&gt;
//	          &lt;object class="GtkLabel" id="tabcontent"&gt;
//	            &lt;property name="label"&gt;Content on notebook tab&lt;/property&gt;
//	          &lt;/object&gt;
//	        &lt;/property&gt;
//	      &lt;/object&gt;
//	    &lt;/child&gt;
//	  &lt;/object&gt;
//	&lt;/child&gt;
//
// &lt;/object&gt;
// ```
//
// # CSS nodes
//
// `GtkPrintUnixDialog` has a single CSS node with name window. The style classes
// dialog and print are added.
type PrintUnixDialog struct {
	Dialog
}

func PrintUnixDialogNewFromInternalPtr(ptr uintptr) *PrintUnixDialog {
	cls := &PrintUnixDialog{}
	cls.Ptr = ptr
	return cls
}

var xNewPrintUnixDialog func(string, uintptr) uintptr

// Creates a new `GtkPrintUnixDialog`.
func NewPrintUnixDialog(TitleVar string, ParentVar *Window) *Widget {
	var cls *Widget

	cret := xNewPrintUnixDialog(TitleVar, ParentVar.GoPointer())

	if cret == 0 {
		return cls
	}
	gobject.IncreaseRef(cret)
	cls = &Widget{}
	cls.Ptr = cret
	return cls
}

var xPrintUnixDialogAddCustomTab func(uintptr, uintptr, uintptr)

// Adds a custom tab to the print dialog.
func (x *PrintUnixDialog) AddCustomTab(ChildVar *Widget, TabLabelVar *Widget) {

	xPrintUnixDialogAddCustomTab(x.GoPointer(), ChildVar.GoPointer(), TabLabelVar.GoPointer())

}

var xPrintUnixDialogGetCurrentPage func(uintptr) int

// Gets the current page of the `GtkPrintUnixDialog`.
func (x *PrintUnixDialog) GetCurrentPage() int {

	cret := xPrintUnixDialogGetCurrentPage(x.GoPointer())
	return cret
}

var xPrintUnixDialogGetEmbedPageSetup func(uintptr) bool

// Gets whether to embed the page setup.
func (x *PrintUnixDialog) GetEmbedPageSetup() bool {

	cret := xPrintUnixDialogGetEmbedPageSetup(x.GoPointer())
	return cret
}

var xPrintUnixDialogGetHasSelection func(uintptr) bool

// Gets whether there is a selection.
func (x *PrintUnixDialog) GetHasSelection() bool {

	cret := xPrintUnixDialogGetHasSelection(x.GoPointer())
	return cret
}

var xPrintUnixDialogGetManualCapabilities func(uintptr) PrintCapabilities

// Gets the capabilities that have been set on this `GtkPrintUnixDialog`.
func (x *PrintUnixDialog) GetManualCapabilities() PrintCapabilities {

	cret := xPrintUnixDialogGetManualCapabilities(x.GoPointer())
	return cret
}

var xPrintUnixDialogGetPageSetup func(uintptr) uintptr

// Gets the page setup that is used by the `GtkPrintUnixDialog`.
func (x *PrintUnixDialog) GetPageSetup() *PageSetup {
	var cls *PageSetup

	cret := xPrintUnixDialogGetPageSetup(x.GoPointer())

	if cret == 0 {
		return cls
	}
	gobject.IncreaseRef(cret)
	cls = &PageSetup{}
	cls.Ptr = cret
	return cls
}

var xPrintUnixDialogGetPageSetupSet func(uintptr) bool

// Gets whether a page setup was set by the user.
func (x *PrintUnixDialog) GetPageSetupSet() bool {

	cret := xPrintUnixDialogGetPageSetupSet(x.GoPointer())
	return cret
}

var xPrintUnixDialogGetSelectedPrinter func(uintptr) uintptr

// Gets the currently selected printer.
func (x *PrintUnixDialog) GetSelectedPrinter() *Printer {
	var cls *Printer

	cret := xPrintUnixDialogGetSelectedPrinter(x.GoPointer())

	if cret == 0 {
		return cls
	}
	gobject.IncreaseRef(cret)
	cls = &Printer{}
	cls.Ptr = cret
	return cls
}

var xPrintUnixDialogGetSettings func(uintptr) uintptr

// Gets a new `GtkPrintSettings` object that represents the
// current values in the print dialog.
//
// Note that this creates a new object, and you need to unref
// it if don’t want to keep it.
func (x *PrintUnixDialog) GetSettings() *PrintSettings {
	var cls *PrintSettings

	cret := xPrintUnixDialogGetSettings(x.GoPointer())

	if cret == 0 {
		return cls
	}
	cls = &PrintSettings{}
	cls.Ptr = cret
	return cls
}

var xPrintUnixDialogGetSupportSelection func(uintptr) bool

// Gets whether the print dialog allows user to print a selection.
func (x *PrintUnixDialog) GetSupportSelection() bool {

	cret := xPrintUnixDialogGetSupportSelection(x.GoPointer())
	return cret
}

var xPrintUnixDialogSetCurrentPage func(uintptr, int)

// Sets the current page number.
//
// If @current_page is not -1, this enables the current page choice
// for the range of pages to print.
func (x *PrintUnixDialog) SetCurrentPage(CurrentPageVar int) {

	xPrintUnixDialogSetCurrentPage(x.GoPointer(), CurrentPageVar)

}

var xPrintUnixDialogSetEmbedPageSetup func(uintptr, bool)

// Embed page size combo box and orientation combo box into page setup page.
func (x *PrintUnixDialog) SetEmbedPageSetup(EmbedVar bool) {

	xPrintUnixDialogSetEmbedPageSetup(x.GoPointer(), EmbedVar)

}

var xPrintUnixDialogSetHasSelection func(uintptr, bool)

// Sets whether a selection exists.
func (x *PrintUnixDialog) SetHasSelection(HasSelectionVar bool) {

	xPrintUnixDialogSetHasSelection(x.GoPointer(), HasSelectionVar)

}

var xPrintUnixDialogSetManualCapabilities func(uintptr, PrintCapabilities)

// This lets you specify the printing capabilities your application
// supports.
//
// For instance, if you can handle scaling the output then you pass
// %GTK_PRINT_CAPABILITY_SCALE. If you don’t pass that, then the dialog
// will only let you select the scale if the printing system automatically
// handles scaling.
func (x *PrintUnixDialog) SetManualCapabilities(CapabilitiesVar PrintCapabilities) {

	xPrintUnixDialogSetManualCapabilities(x.GoPointer(), CapabilitiesVar)

}

var xPrintUnixDialogSetPageSetup func(uintptr, uintptr)

// Sets the page setup of the `GtkPrintUnixDialog`.
func (x *PrintUnixDialog) SetPageSetup(PageSetupVar *PageSetup) {

	xPrintUnixDialogSetPageSetup(x.GoPointer(), PageSetupVar.GoPointer())

}

var xPrintUnixDialogSetSettings func(uintptr, uintptr)

// Sets the `GtkPrintSettings` for the `GtkPrintUnixDialog`.
//
// Typically, this is used to restore saved print settings
// from a previous print operation before the print dialog
// is shown.
func (x *PrintUnixDialog) SetSettings(SettingsVar *PrintSettings) {

	xPrintUnixDialogSetSettings(x.GoPointer(), SettingsVar.GoPointer())

}

var xPrintUnixDialogSetSupportSelection func(uintptr, bool)

// Sets whether the print dialog allows user to print a selection.
func (x *PrintUnixDialog) SetSupportSelection(SupportSelectionVar bool) {

	xPrintUnixDialogSetSupportSelection(x.GoPointer(), SupportSelectionVar)

}

func (c *PrintUnixDialog) GoPointer() uintptr {
	return c.Ptr
}

func (c *PrintUnixDialog) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// Retrieves the `GtkAccessibleRole` for the given `GtkAccessible`.
func (x *PrintUnixDialog) GetAccessibleRole() AccessibleRole {

	cret := XGtkAccessibleGetAccessibleRole(x.GoPointer())
	return cret
}

// Resets the accessible @property to its default value.
func (x *PrintUnixDialog) ResetProperty(PropertyVar AccessibleProperty) {

	XGtkAccessibleResetProperty(x.GoPointer(), PropertyVar)

}

// Resets the accessible @relation to its default value.
func (x *PrintUnixDialog) ResetRelation(RelationVar AccessibleRelation) {

	XGtkAccessibleResetRelation(x.GoPointer(), RelationVar)

}

// Resets the accessible @state to its default value.
func (x *PrintUnixDialog) ResetState(StateVar AccessibleState) {

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
func (x *PrintUnixDialog) UpdateProperty(FirstPropertyVar AccessibleProperty, varArgs ...interface{}) {

	XGtkAccessibleUpdateProperty(x.GoPointer(), FirstPropertyVar, varArgs...)

}

// Updates an array of accessible properties.
//
// This function should be called by `GtkWidget` types whenever an accessible
// property change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *PrintUnixDialog) UpdatePropertyValue(NPropertiesVar int, PropertiesVar uintptr, ValuesVar uintptr) {

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
func (x *PrintUnixDialog) UpdateRelation(FirstRelationVar AccessibleRelation, varArgs ...interface{}) {

	XGtkAccessibleUpdateRelation(x.GoPointer(), FirstRelationVar, varArgs...)

}

// Updates an array of accessible relations.
//
// This function should be called by `GtkWidget` types whenever an accessible
// relation change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *PrintUnixDialog) UpdateRelationValue(NRelationsVar int, RelationsVar uintptr, ValuesVar uintptr) {

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
func (x *PrintUnixDialog) UpdateState(FirstStateVar AccessibleState, varArgs ...interface{}) {

	XGtkAccessibleUpdateState(x.GoPointer(), FirstStateVar, varArgs...)

}

// Updates an array of accessible states.
//
// This function should be called by `GtkWidget` types whenever an accessible
// state change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *PrintUnixDialog) UpdateStateValue(NStatesVar int, StatesVar uintptr, ValuesVar uintptr) {

	XGtkAccessibleUpdateStateValue(x.GoPointer(), NStatesVar, StatesVar, ValuesVar)

}

// Gets the ID of the @buildable object.
//
// `GtkBuilder` sets the name based on the ID attribute
// of the &lt;object&gt; tag used to construct the @buildable.
func (x *PrintUnixDialog) GetBuildableId() string {

	cret := XGtkBuildableGetBuildableId(x.GoPointer())
	return cret
}

// Returns the renderer that is used for this `GtkNative`.
func (x *PrintUnixDialog) GetRenderer() *gsk.Renderer {
	var cls *gsk.Renderer

	cret := XGtkNativeGetRenderer(x.GoPointer())

	if cret == 0 {
		return cls
	}
	gobject.IncreaseRef(cret)
	cls = &gsk.Renderer{}
	cls.Ptr = cret
	return cls
}

// Returns the surface of this `GtkNative`.
func (x *PrintUnixDialog) GetSurface() *gdk.Surface {
	var cls *gdk.Surface

	cret := XGtkNativeGetSurface(x.GoPointer())

	if cret == 0 {
		return cls
	}
	gobject.IncreaseRef(cret)
	cls = &gdk.Surface{}
	cls.Ptr = cret
	return cls
}

// Retrieves the surface transform of @self.
//
// This is the translation from @self's surface coordinates into
// @self's widget coordinates.
func (x *PrintUnixDialog) GetSurfaceTransform(XVar float64, YVar float64) {

	XGtkNativeGetSurfaceTransform(x.GoPointer(), XVar, YVar)

}

// Realizes a `GtkNative`.
//
// This should only be used by subclasses.
func (x *PrintUnixDialog) Realize() {

	XGtkNativeRealize(x.GoPointer())

}

// Unrealizes a `GtkNative`.
//
// This should only be used by subclasses.
func (x *PrintUnixDialog) Unrealize() {

	XGtkNativeUnrealize(x.GoPointer())

}

// Returns the display that this `GtkRoot` is on.
func (x *PrintUnixDialog) GetDisplay() *gdk.Display {
	var cls *gdk.Display

	cret := XGtkRootGetDisplay(x.GoPointer())

	if cret == 0 {
		return cls
	}
	gobject.IncreaseRef(cret)
	cls = &gdk.Display{}
	cls.Ptr = cret
	return cls
}

// Retrieves the current focused widget within the root.
//
// Note that this is the widget that would have the focus
// if the root is active; if the root is not focused then
// `gtk_widget_has_focus (widget)` will be %FALSE for the
// widget.
func (x *PrintUnixDialog) GetFocus() *Widget {
	var cls *Widget

	cret := XGtkRootGetFocus(x.GoPointer())

	if cret == 0 {
		return cls
	}
	gobject.IncreaseRef(cret)
	cls = &Widget{}
	cls.Ptr = cret
	return cls
}

// If @focus is not the current focus widget, and is focusable, sets
// it as the focus widget for the root.
//
// If @focus is %NULL, unsets the focus widget for the root.
//
// To set the focus to a particular widget in the root, it is usually
// more convenient to use [method@Gtk.Widget.grab_focus] instead of
// this function.
func (x *PrintUnixDialog) SetFocus(FocusVar *Widget) {

	XGtkRootSetFocus(x.GoPointer(), FocusVar.GoPointer())

}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GTK"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xNewPrintUnixDialog, lib, "gtk_print_unix_dialog_new")

	core.PuregoSafeRegister(&xPrintUnixDialogAddCustomTab, lib, "gtk_print_unix_dialog_add_custom_tab")
	core.PuregoSafeRegister(&xPrintUnixDialogGetCurrentPage, lib, "gtk_print_unix_dialog_get_current_page")
	core.PuregoSafeRegister(&xPrintUnixDialogGetEmbedPageSetup, lib, "gtk_print_unix_dialog_get_embed_page_setup")
	core.PuregoSafeRegister(&xPrintUnixDialogGetHasSelection, lib, "gtk_print_unix_dialog_get_has_selection")
	core.PuregoSafeRegister(&xPrintUnixDialogGetManualCapabilities, lib, "gtk_print_unix_dialog_get_manual_capabilities")
	core.PuregoSafeRegister(&xPrintUnixDialogGetPageSetup, lib, "gtk_print_unix_dialog_get_page_setup")
	core.PuregoSafeRegister(&xPrintUnixDialogGetPageSetupSet, lib, "gtk_print_unix_dialog_get_page_setup_set")
	core.PuregoSafeRegister(&xPrintUnixDialogGetSelectedPrinter, lib, "gtk_print_unix_dialog_get_selected_printer")
	core.PuregoSafeRegister(&xPrintUnixDialogGetSettings, lib, "gtk_print_unix_dialog_get_settings")
	core.PuregoSafeRegister(&xPrintUnixDialogGetSupportSelection, lib, "gtk_print_unix_dialog_get_support_selection")
	core.PuregoSafeRegister(&xPrintUnixDialogSetCurrentPage, lib, "gtk_print_unix_dialog_set_current_page")
	core.PuregoSafeRegister(&xPrintUnixDialogSetEmbedPageSetup, lib, "gtk_print_unix_dialog_set_embed_page_setup")
	core.PuregoSafeRegister(&xPrintUnixDialogSetHasSelection, lib, "gtk_print_unix_dialog_set_has_selection")
	core.PuregoSafeRegister(&xPrintUnixDialogSetManualCapabilities, lib, "gtk_print_unix_dialog_set_manual_capabilities")
	core.PuregoSafeRegister(&xPrintUnixDialogSetPageSetup, lib, "gtk_print_unix_dialog_set_page_setup")
	core.PuregoSafeRegister(&xPrintUnixDialogSetSettings, lib, "gtk_print_unix_dialog_set_settings")
	core.PuregoSafeRegister(&xPrintUnixDialogSetSupportSelection, lib, "gtk_print_unix_dialog_set_support_selection")

}
