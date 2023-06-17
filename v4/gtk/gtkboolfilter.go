// Package gtk was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gtk

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/gobject"
)

type BoolFilterClass struct {
	ParentClass uintptr
}

// `GtkBoolFilter` evaluates a boolean `GtkExpression`
// to determine whether to include items.
type BoolFilter struct {
	Filter
}

func BoolFilterNewFromInternalPtr(ptr uintptr) *BoolFilter {
	cls := &BoolFilter{}
	cls.Ptr = ptr
	return cls
}

var xNewBoolFilter func(uintptr) uintptr

// Creates a new bool filter.
func NewBoolFilter(ExpressionVar *Expression) *BoolFilter {
	NewBoolFilterPtr := xNewBoolFilter(ExpressionVar.GoPointer())
	if NewBoolFilterPtr == 0 {
		return nil
	}

	NewBoolFilterCls := &BoolFilter{}
	NewBoolFilterCls.Ptr = NewBoolFilterPtr
	return NewBoolFilterCls
}

var xBoolFilterGetExpression func(uintptr) uintptr

// Gets the expression that the filter uses to evaluate if
// an item should be filtered.
func (x *BoolFilter) GetExpression() *Expression {

	GetExpressionPtr := xBoolFilterGetExpression(x.GoPointer())
	if GetExpressionPtr == 0 {
		return nil
	}

	gobject.IncreaseRef(GetExpressionPtr)

	GetExpressionCls := &Expression{}
	GetExpressionCls.Ptr = GetExpressionPtr
	return GetExpressionCls

}

var xBoolFilterGetInvert func(uintptr) bool

// Returns whether the filter inverts the expression.
func (x *BoolFilter) GetInvert() bool {

	return xBoolFilterGetInvert(x.GoPointer())

}

var xBoolFilterSetExpression func(uintptr, uintptr)

// Sets the expression that the filter uses to check if items
// should be filtered.
//
// The expression must have a value type of %G_TYPE_BOOLEAN.
func (x *BoolFilter) SetExpression(ExpressionVar *Expression) {

	xBoolFilterSetExpression(x.GoPointer(), ExpressionVar.GoPointer())

}

var xBoolFilterSetInvert func(uintptr, bool)

// Sets whether the filter should invert the expression.
func (x *BoolFilter) SetInvert(InvertVar bool) {

	xBoolFilterSetInvert(x.GoPointer(), InvertVar)

}

func (c *BoolFilter) GoPointer() uintptr {
	return c.Ptr
}

func (c *BoolFilter) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GTK"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xNewBoolFilter, lib, "gtk_bool_filter_new")

	core.PuregoSafeRegister(&xBoolFilterGetExpression, lib, "gtk_bool_filter_get_expression")
	core.PuregoSafeRegister(&xBoolFilterGetInvert, lib, "gtk_bool_filter_get_invert")
	core.PuregoSafeRegister(&xBoolFilterSetExpression, lib, "gtk_bool_filter_set_expression")
	core.PuregoSafeRegister(&xBoolFilterSetInvert, lib, "gtk_bool_filter_set_invert")

}
