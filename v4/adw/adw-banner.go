// Package adw was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package adw

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/glib"
	"github.com/jwijenbergh/puregotk/v4/gobject"
	"github.com/jwijenbergh/puregotk/v4/gtk"
)

type BannerClass struct {
	ParentClass uintptr
}

// A bar with contextual information.
//
// &lt;picture&gt;
//
//	&lt;source srcset="banner-dark.png" media="(prefers-color-scheme: dark)"&gt;
//	&lt;img src="banner.png" alt="banner"&gt;
//
// &lt;/picture&gt;
//
// Banners are hidden by default, use [property@Banner:revealed] to show them.
//
// Banners have a title, set with [property@Banner:title]. Titles can be marked
// up with Pango markup, use [property@Banner:use-markup] to enable it.
//
// The title will be shown centered or left-aligned depending on available
// space.
//
// Banners can optionally have a button with text on it, set through
// [property@Banner:button-label]. The button can be used with a `GAction`,
// or with the [signal@Banner::button-clicked] signal.
//
// ## CSS nodes
//
// `AdwBanner` has a main CSS node with the name `banner`.
type Banner struct {
	gtk.Widget
}

func BannerNewFromInternalPtr(ptr uintptr) *Banner {
	cls := &Banner{}
	cls.Ptr = ptr
	return cls
}

var xNewBanner func(string) uintptr

// Creates a new `AdwBanner`.
func NewBanner(TitleVar string) *gtk.Widget {
	NewBannerPtr := xNewBanner(TitleVar)
	if NewBannerPtr == 0 {
		return nil
	}

	gobject.IncreaseRef(NewBannerPtr)

	NewBannerCls := &gtk.Widget{}
	NewBannerCls.Ptr = NewBannerPtr
	return NewBannerCls
}

var xBannerGetButtonLabel func(uintptr) string

// Gets the button label for @self.
func (x *Banner) GetButtonLabel() string {

	return xBannerGetButtonLabel(x.GoPointer())

}

var xBannerGetRevealed func(uintptr) bool

// Gets if a banner is revealed
func (x *Banner) GetRevealed() bool {

	return xBannerGetRevealed(x.GoPointer())

}

var xBannerGetTitle func(uintptr) string

// Gets the title for @self.
func (x *Banner) GetTitle() string {

	return xBannerGetTitle(x.GoPointer())

}

var xBannerGetUseMarkup func(uintptr) bool

// Gets whether to use Pango markup for the banner title.
func (x *Banner) GetUseMarkup() bool {

	return xBannerGetUseMarkup(x.GoPointer())

}

var xBannerSetButtonLabel func(uintptr, string)

// Sets the button label for @self.
//
// If set to `""` or `NULL`, the button won't be shown.
//
// The button can be used with a `GAction`, or with the
// [signal@Banner::button-clicked] signal.
func (x *Banner) SetButtonLabel(LabelVar string) {

	xBannerSetButtonLabel(x.GoPointer(), LabelVar)

}

var xBannerSetRevealed func(uintptr, bool)

// Sets whether a banner should be revealed
func (x *Banner) SetRevealed(RevealedVar bool) {

	xBannerSetRevealed(x.GoPointer(), RevealedVar)

}

var xBannerSetTitle func(uintptr, string)

// Sets the title for this banner.
//
// See also: [property@Banner:use-markup].
func (x *Banner) SetTitle(TitleVar string) {

	xBannerSetTitle(x.GoPointer(), TitleVar)

}

var xBannerSetUseMarkup func(uintptr, bool)

// Sets whether to use Pango markup for the banner title.
//
// See also [func@Pango.parse_markup].
func (x *Banner) SetUseMarkup(UseMarkupVar bool) {

	xBannerSetUseMarkup(x.GoPointer(), UseMarkupVar)

}

func (c *Banner) GoPointer() uintptr {
	return c.Ptr
}

func (c *Banner) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// This signal is emitted after the action button has been clicked.
//
// It can be used as an alternative to setting an action.
func (x *Banner) ConnectButtonClicked(cb func(Banner)) {
	fcb := func(clsPtr uintptr) {
		fa := Banner{}
		fa.Ptr = clsPtr

		cb(fa)

	}
	gobject.ObjectConnect(x.GoPointer(), "signal::button-clicked", purego.NewCallback(fcb))
}

// Retrieves the `GtkAccessibleRole` for the given `GtkAccessible`.
func (x *Banner) GetAccessibleRole() gtk.AccessibleRole {

	return gtk.XGtkAccessibleGetAccessibleRole(x.GoPointer())

}

// Resets the accessible @property to its default value.
func (x *Banner) ResetProperty(PropertyVar gtk.AccessibleProperty) {

	gtk.XGtkAccessibleResetProperty(x.GoPointer(), PropertyVar)

}

// Resets the accessible @relation to its default value.
func (x *Banner) ResetRelation(RelationVar gtk.AccessibleRelation) {

	gtk.XGtkAccessibleResetRelation(x.GoPointer(), RelationVar)

}

// Resets the accessible @state to its default value.
func (x *Banner) ResetState(StateVar gtk.AccessibleState) {

	gtk.XGtkAccessibleResetState(x.GoPointer(), StateVar)

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
func (x *Banner) UpdateProperty(FirstPropertyVar gtk.AccessibleProperty, varArgs ...interface{}) {

	gtk.XGtkAccessibleUpdateProperty(x.GoPointer(), FirstPropertyVar, varArgs...)

}

// Updates an array of accessible properties.
//
// This function should be called by `GtkWidget` types whenever an accessible
// property change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *Banner) UpdatePropertyValue(NPropertiesVar int32, PropertiesVar uintptr, ValuesVar uintptr) {

	gtk.XGtkAccessibleUpdatePropertyValue(x.GoPointer(), NPropertiesVar, PropertiesVar, ValuesVar)

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
func (x *Banner) UpdateRelation(FirstRelationVar gtk.AccessibleRelation, varArgs ...interface{}) {

	gtk.XGtkAccessibleUpdateRelation(x.GoPointer(), FirstRelationVar, varArgs...)

}

// Updates an array of accessible relations.
//
// This function should be called by `GtkWidget` types whenever an accessible
// relation change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *Banner) UpdateRelationValue(NRelationsVar int32, RelationsVar uintptr, ValuesVar uintptr) {

	gtk.XGtkAccessibleUpdateRelationValue(x.GoPointer(), NRelationsVar, RelationsVar, ValuesVar)

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
func (x *Banner) UpdateState(FirstStateVar gtk.AccessibleState, varArgs ...interface{}) {

	gtk.XGtkAccessibleUpdateState(x.GoPointer(), FirstStateVar, varArgs...)

}

// Updates an array of accessible states.
//
// This function should be called by `GtkWidget` types whenever an accessible
// state change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *Banner) UpdateStateValue(NStatesVar int32, StatesVar uintptr, ValuesVar uintptr) {

	gtk.XGtkAccessibleUpdateStateValue(x.GoPointer(), NStatesVar, StatesVar, ValuesVar)

}

// Gets the action name for @actionable.
func (x *Banner) GetActionName() string {

	return gtk.XGtkActionableGetActionName(x.GoPointer())

}

// Gets the current target value of @actionable.
func (x *Banner) GetActionTargetValue() *glib.Variant {

	return gtk.XGtkActionableGetActionTargetValue(x.GoPointer())

}

// Specifies the name of the action with which this widget should be
// associated.
//
// If @action_name is %NULL then the widget will be unassociated from
// any previous action.
//
// Usually this function is used when the widget is located (or will be
// located) within the hierarchy of a `GtkApplicationWindow`.
//
// Names are of the form “win.save” or “app.quit” for actions on the
// containing [class@ApplicationWindow] or its associated [class@Application],
// respectively. This is the same form used for actions in the [class@Gio.Menu]
// associated with the window.
func (x *Banner) SetActionName(ActionNameVar string) {

	gtk.XGtkActionableSetActionName(x.GoPointer(), ActionNameVar)

}

// Sets the target of an actionable widget.
//
// This is a convenience function that calls [ctor@GLib.Variant.new] for
// @format_string and uses the result to call
// [method@Gtk.Actionable.set_action_target_value].
//
// If you are setting a string-valued target and want to set
// the action name at the same time, you can use
// [method@Gtk.Actionable.set_detailed_action_name].
func (x *Banner) SetActionTarget(FormatStringVar string, varArgs ...interface{}) {

	gtk.XGtkActionableSetActionTarget(x.GoPointer(), FormatStringVar, varArgs...)

}

// Sets the target value of an actionable widget.
//
// If @target_value is %NULL then the target value is unset.
//
// The target value has two purposes. First, it is used as the parameter
// to activation of the action associated with the `GtkActionable` widget.
// Second, it is used to determine if the widget should be rendered as
// “active” — the widget is active if the state is equal to the given target.
//
// Consider the example of associating a set of buttons with a [iface@Gio.Action]
// with string state in a typical “radio button” situation. Each button
// will be associated with the same action, but with a different target
// value for that action. Clicking on a particular button will activate
// the action with the target of that button, which will typically cause
// the action’s state to change to that value. Since the action’s state
// is now equal to the target value of the button, the button will now
// be rendered as active (and the other buttons, with different targets,
// rendered inactive).
func (x *Banner) SetActionTargetValue(TargetValueVar *glib.Variant) {

	gtk.XGtkActionableSetActionTargetValue(x.GoPointer(), TargetValueVar)

}

// Sets the action-name and associated string target value of an
// actionable widget.
//
// @detailed_action_name is a string in the format accepted by
// [func@Gio.Action.parse_detailed_name].
func (x *Banner) SetDetailedActionName(DetailedActionNameVar string) {

	gtk.XGtkActionableSetDetailedActionName(x.GoPointer(), DetailedActionNameVar)

}

// Gets the ID of the @buildable object.
//
// `GtkBuilder` sets the name based on the ID attribute
// of the &lt;object&gt; tag used to construct the @buildable.
func (x *Banner) GetBuildableId() string {

	return gtk.XGtkBuildableGetBuildableId(x.GoPointer())

}

func init() {
	lib, err := purego.Dlopen(core.GetPath("ADW"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xNewBanner, lib, "adw_banner_new")

	core.PuregoSafeRegister(&xBannerGetButtonLabel, lib, "adw_banner_get_button_label")
	core.PuregoSafeRegister(&xBannerGetRevealed, lib, "adw_banner_get_revealed")
	core.PuregoSafeRegister(&xBannerGetTitle, lib, "adw_banner_get_title")
	core.PuregoSafeRegister(&xBannerGetUseMarkup, lib, "adw_banner_get_use_markup")
	core.PuregoSafeRegister(&xBannerSetButtonLabel, lib, "adw_banner_set_button_label")
	core.PuregoSafeRegister(&xBannerSetRevealed, lib, "adw_banner_set_revealed")
	core.PuregoSafeRegister(&xBannerSetTitle, lib, "adw_banner_set_title")
	core.PuregoSafeRegister(&xBannerSetUseMarkup, lib, "adw_banner_set_use_markup")

}
