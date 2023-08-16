// Package gtk was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gtk

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/gdk"
	"github.com/jwijenbergh/puregotk/v4/gio"
	"github.com/jwijenbergh/puregotk/v4/gobject"
	"github.com/jwijenbergh/puregotk/v4/gsk"
)

// Flags that affect how popover menus are created from
// a menu model.
type PopoverMenuFlags int

const (

	// Create submenus as nested
	//    popovers. Without this flag, submenus are created as
	//    sliding pages that replace the main menu.
	PopoverMenuNestedValue PopoverMenuFlags = 1
)

// `GtkPopoverMenu` is a subclass of `GtkPopover` that implements menu
// behavior.
//
// ![An example GtkPopoverMenu](menu.png)
//
// `GtkPopoverMenu` treats its children like menus and allows switching
// between them. It can open submenus as traditional, nested submenus,
// or in a more touch-friendly sliding fashion.
//
// `GtkPopoverMenu` is meant to be used primarily with menu models,
// using [ctor@Gtk.PopoverMenu.new_from_model]. If you need to put
// other widgets such as a `GtkSpinButton` or a `GtkSwitch` into a popover,
// you can use [method@Gtk.PopoverMenu.add_child].
//
// For more dialog-like behavior, use a plain `GtkPopover`.
//
// ## Menu models
//
// The XML format understood by `GtkBuilder` for `GMenuModel` consists
// of a toplevel `&lt;menu&gt;` element, which contains one or more `&lt;item&gt;`
// elements. Each `&lt;item&gt;` element contains `&lt;attribute&gt;` and `&lt;link&gt;`
// elements with a mandatory name attribute. `&lt;link&gt;` elements have the
// same content model as `&lt;menu&gt;`. Instead of `&lt;link name="submenu"&gt;`
// or `&lt;link name="section"&gt;`, you can use `&lt;submenu&gt;` or `&lt;section&gt;`
// elements.
//
// ```xml
// &lt;menu id='app-menu'&gt;
//
//	&lt;section&gt;
//	  &lt;item&gt;
//	    &lt;attribute name='label' translatable='yes'&gt;_New Window&lt;/attribute&gt;
//	    &lt;attribute name='action'&gt;app.new&lt;/attribute&gt;
//	  &lt;/item&gt;
//	  &lt;item&gt;
//	    &lt;attribute name='label' translatable='yes'&gt;_About Sunny&lt;/attribute&gt;
//	    &lt;attribute name='action'&gt;app.about&lt;/attribute&gt;
//	  &lt;/item&gt;
//	  &lt;item&gt;
//	    &lt;attribute name='label' translatable='yes'&gt;_Quit&lt;/attribute&gt;
//	    &lt;attribute name='action'&gt;app.quit&lt;/attribute&gt;
//	  &lt;/item&gt;
//	&lt;/section&gt;
//
// &lt;/menu&gt;
// ```
//
// Attribute values can be translated using gettext, like other `GtkBuilder`
// content. `&lt;attribute&gt;` elements can be marked for translation with a
// `translatable="yes"` attribute. It is also possible to specify message
// context and translator comments, using the context and comments attributes.
// To make use of this, the `GtkBuilder` must have been given the gettext
// domain to use.
//
// The following attributes are used when constructing menu items:
//
//   - "label": a user-visible string to display
//   - "use-markup": whether the text in the menu item includes [Pango markup](https://docs.gtk.org/Pango/pango_markup.html)
//   - "action": the prefixed name of the action to trigger
//   - "target": the parameter to use when activating the action
//   - "icon" and "verb-icon": names of icons that may be displayed
//   - "submenu-action": name of an action that may be used to track
//     whether a submenu is open
//   - "hidden-when": a string used to determine when the item will be hidden.
//     Possible values include "action-disabled", "action-missing", "macos-menubar".
//     This is mainly useful for exported menus, see [method@Gtk.Application.set_menubar].
//   - "custom": a string used to match against the ID of a custom child added with
//     [method@Gtk.PopoverMenu.add_child], [method@Gtk.PopoverMenuBar.add_child],
//     or in the ui file with `&lt;child type="ID"&gt;`.
//
// The following attributes are used when constructing sections:
//
//   - "label": a user-visible string to use as section heading
//   - "display-hint": a string used to determine special formatting for the section.
//     Possible values include "horizontal-buttons", "circular-buttons" and
//     "inline-buttons". They all indicate that section should be
//     displayed as a horizontal row of buttons.
//   - "text-direction": a string used to determine the `GtkTextDirection` to use
//     when "display-hint" is set to "horizontal-buttons". Possible values
//     include "rtl", "ltr", and "none".
//
// The following attributes are used when constructing submenus:
//
// - "label": a user-visible string to display
// - "icon": icon name to display
//
// Menu items will also show accelerators, which are usually associated
// with actions via [method@Gtk.Application.set_accels_for_action],
// [id@gtk_widget_class_add_binding_action] or
// [method@Gtk.ShortcutController.add_shortcut].
//
// # CSS Nodes
//
// `GtkPopoverMenu` is just a subclass of `GtkPopover` that adds custom content
// to it, therefore it has the same CSS nodes. It is one of the cases that add
// a .menu style class to the popover's main node.
//
// # Accessibility
//
// `GtkPopoverMenu` uses the %GTK_ACCESSIBLE_ROLE_MENU role, and its
// items use the %GTK_ACCESSIBLE_ROLE_MENU_ITEM,
// %GTK_ACCESSIBLE_ROLE_MENU_ITEM_CHECKBOX or
// %GTK_ACCESSIBLE_ROLE_MENU_ITEM_RADIO roles, depending on the
// action they are connected to.
type PopoverMenu struct {
	Popover
}

func PopoverMenuNewFromInternalPtr(ptr uintptr) *PopoverMenu {
	cls := &PopoverMenu{}
	cls.Ptr = ptr
	return cls
}

var xNewFromModelPopoverMenu func(uintptr) uintptr

// Creates a `GtkPopoverMenu` and populates it according to @model.
//
// The created buttons are connected to actions found in the
// `GtkApplicationWindow` to which the popover belongs - typically
// by means of being attached to a widget that is contained within
// the `GtkApplicationWindow`s widget hierarchy.
//
// Actions can also be added using [method@Gtk.Widget.insert_action_group]
// on the menus attach widget or on any of its parent widgets.
//
// This function creates menus with sliding submenus.
// See [ctor@Gtk.PopoverMenu.new_from_model_full] for a way
// to control this.
func NewFromModelPopoverMenu(ModelVar *gio.MenuModel) *Widget {
	var cls *Widget

	cret := xNewFromModelPopoverMenu(ModelVar.GoPointer())

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &Widget{}
	cls.Ptr = cret
	return cls
}

var xNewFromModelFullPopoverMenu func(uintptr, PopoverMenuFlags) uintptr

// Creates a `GtkPopoverMenu` and populates it according to @model.
//
// The created buttons are connected to actions found in the
// action groups that are accessible from the parent widget.
// This includes the `GtkApplicationWindow` to which the popover
// belongs. Actions can also be added using [method@Gtk.Widget.insert_action_group]
// on the parent widget or on any of its parent widgets.
//
// The only flag that is supported currently is
// %GTK_POPOVER_MENU_NESTED, which makes GTK create traditional,
// nested submenus instead of the default sliding submenus.
func NewFromModelFullPopoverMenu(ModelVar *gio.MenuModel, FlagsVar PopoverMenuFlags) *Widget {
	var cls *Widget

	cret := xNewFromModelFullPopoverMenu(ModelVar.GoPointer(), FlagsVar)

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &Widget{}
	cls.Ptr = cret
	return cls
}

var xPopoverMenuAddChild func(uintptr, uintptr, string) bool

// Adds a custom widget to a generated menu.
//
// For this to work, the menu model of @popover must have
// an item with a `custom` attribute that matches @id.
func (x *PopoverMenu) AddChild(ChildVar *Widget, IdVar string) bool {

	cret := xPopoverMenuAddChild(x.GoPointer(), ChildVar.GoPointer(), IdVar)
	return cret
}

var xPopoverMenuGetMenuModel func(uintptr) uintptr

// Returns the menu model used to populate the popover.
func (x *PopoverMenu) GetMenuModel() *gio.MenuModel {
	var cls *gio.MenuModel

	cret := xPopoverMenuGetMenuModel(x.GoPointer())

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &gio.MenuModel{}
	cls.Ptr = cret
	return cls
}

var xPopoverMenuRemoveChild func(uintptr, uintptr) bool

// Removes a widget that has previously been added with
// gtk_popover_menu_add_child().
func (x *PopoverMenu) RemoveChild(ChildVar *Widget) bool {

	cret := xPopoverMenuRemoveChild(x.GoPointer(), ChildVar.GoPointer())
	return cret
}

var xPopoverMenuSetMenuModel func(uintptr, uintptr)

// Sets a new menu model on @popover.
//
// The existing contents of @popover are removed, and
// the @popover is populated with new contents according
// to @model.
func (x *PopoverMenu) SetMenuModel(ModelVar *gio.MenuModel) {

	xPopoverMenuSetMenuModel(x.GoPointer(), ModelVar.GoPointer())

}

func (c *PopoverMenu) GoPointer() uintptr {
	return c.Ptr
}

func (c *PopoverMenu) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// Retrieves the `GtkAccessibleRole` for the given `GtkAccessible`.
func (x *PopoverMenu) GetAccessibleRole() AccessibleRole {

	cret := XGtkAccessibleGetAccessibleRole(x.GoPointer())
	return cret
}

// Resets the accessible @property to its default value.
func (x *PopoverMenu) ResetProperty(PropertyVar AccessibleProperty) {

	XGtkAccessibleResetProperty(x.GoPointer(), PropertyVar)

}

// Resets the accessible @relation to its default value.
func (x *PopoverMenu) ResetRelation(RelationVar AccessibleRelation) {

	XGtkAccessibleResetRelation(x.GoPointer(), RelationVar)

}

// Resets the accessible @state to its default value.
func (x *PopoverMenu) ResetState(StateVar AccessibleState) {

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
func (x *PopoverMenu) UpdateProperty(FirstPropertyVar AccessibleProperty, varArgs ...interface{}) {

	XGtkAccessibleUpdateProperty(x.GoPointer(), FirstPropertyVar, varArgs...)

}

// Updates an array of accessible properties.
//
// This function should be called by `GtkWidget` types whenever an accessible
// property change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *PopoverMenu) UpdatePropertyValue(NPropertiesVar int, PropertiesVar uintptr, ValuesVar uintptr) {

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
func (x *PopoverMenu) UpdateRelation(FirstRelationVar AccessibleRelation, varArgs ...interface{}) {

	XGtkAccessibleUpdateRelation(x.GoPointer(), FirstRelationVar, varArgs...)

}

// Updates an array of accessible relations.
//
// This function should be called by `GtkWidget` types whenever an accessible
// relation change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *PopoverMenu) UpdateRelationValue(NRelationsVar int, RelationsVar uintptr, ValuesVar uintptr) {

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
func (x *PopoverMenu) UpdateState(FirstStateVar AccessibleState, varArgs ...interface{}) {

	XGtkAccessibleUpdateState(x.GoPointer(), FirstStateVar, varArgs...)

}

// Updates an array of accessible states.
//
// This function should be called by `GtkWidget` types whenever an accessible
// state change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *PopoverMenu) UpdateStateValue(NStatesVar int, StatesVar uintptr, ValuesVar uintptr) {

	XGtkAccessibleUpdateStateValue(x.GoPointer(), NStatesVar, StatesVar, ValuesVar)

}

// Gets the ID of the @buildable object.
//
// `GtkBuilder` sets the name based on the ID attribute
// of the &lt;object&gt; tag used to construct the @buildable.
func (x *PopoverMenu) GetBuildableId() string {

	cret := XGtkBuildableGetBuildableId(x.GoPointer())
	return cret
}

// Returns the renderer that is used for this `GtkNative`.
func (x *PopoverMenu) GetRenderer() *gsk.Renderer {
	var cls *gsk.Renderer

	cret := XGtkNativeGetRenderer(x.GoPointer())

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &gsk.Renderer{}
	cls.Ptr = cret
	return cls
}

// Returns the surface of this `GtkNative`.
func (x *PopoverMenu) GetSurface() *gdk.Surface {
	var cls *gdk.Surface

	cret := XGtkNativeGetSurface(x.GoPointer())

	if cret == 0 {
		return nil
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
func (x *PopoverMenu) GetSurfaceTransform(XVar float64, YVar float64) {

	XGtkNativeGetSurfaceTransform(x.GoPointer(), XVar, YVar)

}

// Realizes a `GtkNative`.
//
// This should only be used by subclasses.
func (x *PopoverMenu) Realize() {

	XGtkNativeRealize(x.GoPointer())

}

// Unrealizes a `GtkNative`.
//
// This should only be used by subclasses.
func (x *PopoverMenu) Unrealize() {

	XGtkNativeUnrealize(x.GoPointer())

}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GTK"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xNewFromModelPopoverMenu, lib, "gtk_popover_menu_new_from_model")
	core.PuregoSafeRegister(&xNewFromModelFullPopoverMenu, lib, "gtk_popover_menu_new_from_model_full")

	core.PuregoSafeRegister(&xPopoverMenuAddChild, lib, "gtk_popover_menu_add_child")
	core.PuregoSafeRegister(&xPopoverMenuGetMenuModel, lib, "gtk_popover_menu_get_menu_model")
	core.PuregoSafeRegister(&xPopoverMenuRemoveChild, lib, "gtk_popover_menu_remove_child")
	core.PuregoSafeRegister(&xPopoverMenuSetMenuModel, lib, "gtk_popover_menu_set_menu_model")

}
