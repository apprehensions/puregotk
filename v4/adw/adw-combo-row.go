// Package adw was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package adw

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/gio"
	"github.com/jwijenbergh/puregotk/v4/glib"
	"github.com/jwijenbergh/puregotk/v4/gobject"
	"github.com/jwijenbergh/puregotk/v4/gtk"
)

type ComboRowClass struct {
	ParentClass uintptr

	Padding uintptr
}

// A [class@Gtk.ListBoxRow] used to choose from a list of items.
//
// &lt;picture&gt;
//
//	&lt;source srcset="combo-row-dark.png" media="(prefers-color-scheme: dark)"&gt;
//	&lt;img src="combo-row.png" alt="combo-row"&gt;
//
// &lt;/picture&gt;
//
// The `AdwComboRow` widget allows the user to choose from a list of valid
// choices. The row displays the selected choice. When activated, the row
// displays a popover which allows the user to make a new choice.
//
// Example of an `AdwComboRow` UI definition:
// ```xml
// &lt;object class="AdwComboRow"&gt;
//
//	&lt;property name="title" translatable="yes"&gt;Combo Row&lt;/property&gt;
//	&lt;property name="model"&gt;
//	  &lt;object class="GtkStringList"&gt;
//	    &lt;items&gt;
//	      &lt;item translatable="yes"&gt;Foo&lt;/item&gt;
//	      &lt;item translatable="yes"&gt;Bar&lt;/item&gt;
//	      &lt;item translatable="yes"&gt;Baz&lt;/item&gt;
//	    &lt;/items&gt;
//	  &lt;/object&gt;
//	&lt;/property&gt;
//
// &lt;/object&gt;
// ```
//
// The [property@ComboRow:selected] and [property@ComboRow:selected-item]
// properties can be used to keep track of the selected item and react to their
// changes.
//
// `AdwComboRow` mirrors [class@Gtk.DropDown], see that widget for details.
//
// `AdwComboRow` is [property@Gtk.ListBoxRow:activatable] if a model is set.
//
// ## CSS nodes
//
// `AdwComboRow` has a main CSS node with name `row` and the `.combo` style
// class.
//
// Its popover has the node named `popover` with the `.menu` style class, it
// contains a [class@Gtk.ScrolledWindow], which in turn contains a
// [class@Gtk.ListView], both are accessible via their regular nodes.
//
// ## Accessibility
//
// `AdwComboRow` uses the `GTK_ACCESSIBLE_ROLE_COMBO_BOX` role.
type ComboRow struct {
	ActionRow
}

func ComboRowNewFromInternalPtr(ptr uintptr) *ComboRow {
	cls := &ComboRow{}
	cls.Ptr = ptr
	return cls
}

var xNewComboRow func() uintptr

// Creates a new `AdwComboRow`.
func NewComboRow() *gtk.Widget {
	NewComboRowPtr := xNewComboRow()
	if NewComboRowPtr == 0 {
		return nil
	}

	gobject.IncreaseRef(NewComboRowPtr)

	NewComboRowCls := &gtk.Widget{}
	NewComboRowCls.Ptr = NewComboRowPtr
	return NewComboRowCls
}

var xComboRowGetExpression func(uintptr) uintptr

// Gets the expression used to obtain strings from items.
func (x *ComboRow) GetExpression() *gtk.Expression {

	GetExpressionPtr := xComboRowGetExpression(x.GoPointer())
	if GetExpressionPtr == 0 {
		return nil
	}

	gobject.IncreaseRef(GetExpressionPtr)

	GetExpressionCls := &gtk.Expression{}
	GetExpressionCls.Ptr = GetExpressionPtr
	return GetExpressionCls

}

var xComboRowGetFactory func(uintptr) uintptr

// Gets the factory for populating list items.
func (x *ComboRow) GetFactory() *gtk.ListItemFactory {

	GetFactoryPtr := xComboRowGetFactory(x.GoPointer())
	if GetFactoryPtr == 0 {
		return nil
	}

	gobject.IncreaseRef(GetFactoryPtr)

	GetFactoryCls := &gtk.ListItemFactory{}
	GetFactoryCls.Ptr = GetFactoryPtr
	return GetFactoryCls

}

var xComboRowGetListFactory func(uintptr) uintptr

// Gets the factory for populating list items in the popup.
func (x *ComboRow) GetListFactory() *gtk.ListItemFactory {

	GetListFactoryPtr := xComboRowGetListFactory(x.GoPointer())
	if GetListFactoryPtr == 0 {
		return nil
	}

	gobject.IncreaseRef(GetListFactoryPtr)

	GetListFactoryCls := &gtk.ListItemFactory{}
	GetListFactoryCls.Ptr = GetListFactoryPtr
	return GetListFactoryCls

}

var xComboRowGetModel func(uintptr) uintptr

// Gets the model that provides the displayed items.
func (x *ComboRow) GetModel() *gio.ListModelBase {

	GetModelPtr := xComboRowGetModel(x.GoPointer())
	if GetModelPtr == 0 {
		return nil
	}

	gobject.IncreaseRef(GetModelPtr)

	GetModelCls := &gio.ListModelBase{}
	GetModelCls.Ptr = GetModelPtr
	return GetModelCls

}

var xComboRowGetSelected func(uintptr) uint

// Gets the position of the selected item.
func (x *ComboRow) GetSelected() uint {

	return xComboRowGetSelected(x.GoPointer())

}

var xComboRowGetSelectedItem func(uintptr) uintptr

// Gets the selected item.
func (x *ComboRow) GetSelectedItem() *gobject.Object {

	GetSelectedItemPtr := xComboRowGetSelectedItem(x.GoPointer())
	if GetSelectedItemPtr == 0 {
		return nil
	}

	gobject.IncreaseRef(GetSelectedItemPtr)

	GetSelectedItemCls := &gobject.Object{}
	GetSelectedItemCls.Ptr = GetSelectedItemPtr
	return GetSelectedItemCls

}

var xComboRowGetUseSubtitle func(uintptr) bool

// Gets whether to use the current value as the subtitle.
func (x *ComboRow) GetUseSubtitle() bool {

	return xComboRowGetUseSubtitle(x.GoPointer())

}

var xComboRowSetExpression func(uintptr, uintptr)

// Sets the expression used to obtain strings from items.
//
// The expression must have a value type of `G_TYPE_STRING`.
//
// It's used to bind strings to labels produced by the default factory if
// [property@ComboRow:factory] is not set, or when
// [property@ComboRow:use-subtitle] is set to `TRUE`.
func (x *ComboRow) SetExpression(ExpressionVar *gtk.Expression) {

	xComboRowSetExpression(x.GoPointer(), ExpressionVar.GoPointer())

}

var xComboRowSetFactory func(uintptr, uintptr)

// Sets the factory for populating list items.
//
// This factory is always used for the item in the row. It is also used for
// items in the popup unless [property@ComboRow:list-factory] is set.
func (x *ComboRow) SetFactory(FactoryVar *gtk.ListItemFactory) {

	xComboRowSetFactory(x.GoPointer(), FactoryVar.GoPointer())

}

var xComboRowSetListFactory func(uintptr, uintptr)

// Sets the factory for populating list items in the popup.
//
// If this is not set, [property@ComboRow:factory] is used.
func (x *ComboRow) SetListFactory(FactoryVar *gtk.ListItemFactory) {

	xComboRowSetListFactory(x.GoPointer(), FactoryVar.GoPointer())

}

var xComboRowSetModel func(uintptr, uintptr)

// Sets the model that provides the displayed items.
func (x *ComboRow) SetModel(ModelVar gio.ListModel) {

	xComboRowSetModel(x.GoPointer(), ModelVar.GoPointer())

}

var xComboRowSetSelected func(uintptr, uint)

// Selects the item at the given position.
func (x *ComboRow) SetSelected(PositionVar uint) {

	xComboRowSetSelected(x.GoPointer(), PositionVar)

}

var xComboRowSetUseSubtitle func(uintptr, bool)

// Sets whether to use the current value as the subtitle.
//
// If you use a custom list item factory, you will need to give the row a
// name conversion expression with [property@ComboRow:expression].
//
// If set to `TRUE`, you should not access [property@ActionRow:subtitle].
//
// The subtitle is interpreted as Pango markup if
// [property@PreferencesRow:use-markup] is set to `TRUE`.
func (x *ComboRow) SetUseSubtitle(UseSubtitleVar bool) {

	xComboRowSetUseSubtitle(x.GoPointer(), UseSubtitleVar)

}

func (c *ComboRow) GoPointer() uintptr {
	return c.Ptr
}

func (c *ComboRow) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// Retrieves the `GtkAccessibleRole` for the given `GtkAccessible`.
func (x *ComboRow) GetAccessibleRole() gtk.AccessibleRole {

	return gtk.XGtkAccessibleGetAccessibleRole(x.GoPointer())

}

// Resets the accessible @property to its default value.
func (x *ComboRow) ResetProperty(PropertyVar gtk.AccessibleProperty) {

	gtk.XGtkAccessibleResetProperty(x.GoPointer(), PropertyVar)

}

// Resets the accessible @relation to its default value.
func (x *ComboRow) ResetRelation(RelationVar gtk.AccessibleRelation) {

	gtk.XGtkAccessibleResetRelation(x.GoPointer(), RelationVar)

}

// Resets the accessible @state to its default value.
func (x *ComboRow) ResetState(StateVar gtk.AccessibleState) {

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
func (x *ComboRow) UpdateProperty(FirstPropertyVar gtk.AccessibleProperty, varArgs ...interface{}) {

	gtk.XGtkAccessibleUpdateProperty(x.GoPointer(), FirstPropertyVar, varArgs...)

}

// Updates an array of accessible properties.
//
// This function should be called by `GtkWidget` types whenever an accessible
// property change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *ComboRow) UpdatePropertyValue(NPropertiesVar int32, PropertiesVar uintptr, ValuesVar uintptr) {

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
func (x *ComboRow) UpdateRelation(FirstRelationVar gtk.AccessibleRelation, varArgs ...interface{}) {

	gtk.XGtkAccessibleUpdateRelation(x.GoPointer(), FirstRelationVar, varArgs...)

}

// Updates an array of accessible relations.
//
// This function should be called by `GtkWidget` types whenever an accessible
// relation change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *ComboRow) UpdateRelationValue(NRelationsVar int32, RelationsVar uintptr, ValuesVar uintptr) {

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
func (x *ComboRow) UpdateState(FirstStateVar gtk.AccessibleState, varArgs ...interface{}) {

	gtk.XGtkAccessibleUpdateState(x.GoPointer(), FirstStateVar, varArgs...)

}

// Updates an array of accessible states.
//
// This function should be called by `GtkWidget` types whenever an accessible
// state change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *ComboRow) UpdateStateValue(NStatesVar int32, StatesVar uintptr, ValuesVar uintptr) {

	gtk.XGtkAccessibleUpdateStateValue(x.GoPointer(), NStatesVar, StatesVar, ValuesVar)

}

// Gets the action name for @actionable.
func (x *ComboRow) GetActionName() string {

	return gtk.XGtkActionableGetActionName(x.GoPointer())

}

// Gets the current target value of @actionable.
func (x *ComboRow) GetActionTargetValue() *glib.Variant {

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
func (x *ComboRow) SetActionName(ActionNameVar string) {

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
func (x *ComboRow) SetActionTarget(FormatStringVar string, varArgs ...interface{}) {

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
func (x *ComboRow) SetActionTargetValue(TargetValueVar *glib.Variant) {

	gtk.XGtkActionableSetActionTargetValue(x.GoPointer(), TargetValueVar)

}

// Sets the action-name and associated string target value of an
// actionable widget.
//
// @detailed_action_name is a string in the format accepted by
// [func@Gio.Action.parse_detailed_name].
func (x *ComboRow) SetDetailedActionName(DetailedActionNameVar string) {

	gtk.XGtkActionableSetDetailedActionName(x.GoPointer(), DetailedActionNameVar)

}

// Gets the ID of the @buildable object.
//
// `GtkBuilder` sets the name based on the ID attribute
// of the &lt;object&gt; tag used to construct the @buildable.
func (x *ComboRow) GetBuildableId() string {

	return gtk.XGtkBuildableGetBuildableId(x.GoPointer())

}

func init() {
	lib, err := purego.Dlopen(core.GetPath("ADW"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xNewComboRow, lib, "adw_combo_row_new")

	core.PuregoSafeRegister(&xComboRowGetExpression, lib, "adw_combo_row_get_expression")
	core.PuregoSafeRegister(&xComboRowGetFactory, lib, "adw_combo_row_get_factory")
	core.PuregoSafeRegister(&xComboRowGetListFactory, lib, "adw_combo_row_get_list_factory")
	core.PuregoSafeRegister(&xComboRowGetModel, lib, "adw_combo_row_get_model")
	core.PuregoSafeRegister(&xComboRowGetSelected, lib, "adw_combo_row_get_selected")
	core.PuregoSafeRegister(&xComboRowGetSelectedItem, lib, "adw_combo_row_get_selected_item")
	core.PuregoSafeRegister(&xComboRowGetUseSubtitle, lib, "adw_combo_row_get_use_subtitle")
	core.PuregoSafeRegister(&xComboRowSetExpression, lib, "adw_combo_row_set_expression")
	core.PuregoSafeRegister(&xComboRowSetFactory, lib, "adw_combo_row_set_factory")
	core.PuregoSafeRegister(&xComboRowSetListFactory, lib, "adw_combo_row_set_list_factory")
	core.PuregoSafeRegister(&xComboRowSetModel, lib, "adw_combo_row_set_model")
	core.PuregoSafeRegister(&xComboRowSetSelected, lib, "adw_combo_row_set_selected")
	core.PuregoSafeRegister(&xComboRowSetUseSubtitle, lib, "adw_combo_row_set_use_subtitle")

}
