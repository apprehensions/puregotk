// Package gtk was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gtk

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/gobject"
)

type NumericSorterClass struct {
	ParentClass uintptr
}

// `GtkNumericSorter` is a `GtkSorter` that compares numbers.
//
// To obtain the numbers to compare, this sorter evaluates a
// [class@Gtk.Expression].
type NumericSorter struct {
	Sorter
}

func NumericSorterNewFromInternalPtr(ptr uintptr) *NumericSorter {
	cls := &NumericSorter{}
	cls.Ptr = ptr
	return cls
}

var xNewNumericSorter func(uintptr) uintptr

// Creates a new numeric sorter using the given @expression.
//
// Smaller numbers will be sorted first. You can call
// [method@Gtk.NumericSorter.set_sort_order] to change this.
func NewNumericSorter(ExpressionVar *Expression) *NumericSorter {
	NewNumericSorterPtr := xNewNumericSorter(ExpressionVar.GoPointer())
	if NewNumericSorterPtr == 0 {
		return nil
	}

	NewNumericSorterCls := &NumericSorter{}
	NewNumericSorterCls.Ptr = NewNumericSorterPtr
	return NewNumericSorterCls
}

var xNumericSorterGetExpression func(uintptr) uintptr

// Gets the expression that is evaluated to obtain numbers from items.
func (x *NumericSorter) GetExpression() *Expression {

	GetExpressionPtr := xNumericSorterGetExpression(x.GoPointer())
	if GetExpressionPtr == 0 {
		return nil
	}

	gobject.IncreaseRef(GetExpressionPtr)

	GetExpressionCls := &Expression{}
	GetExpressionCls.Ptr = GetExpressionPtr
	return GetExpressionCls

}

var xNumericSorterGetSortOrder func(uintptr) SortType

// Gets whether this sorter will sort smaller numbers first.
func (x *NumericSorter) GetSortOrder() SortType {

	return xNumericSorterGetSortOrder(x.GoPointer())

}

var xNumericSorterSetExpression func(uintptr, uintptr)

// Sets the expression that is evaluated to obtain numbers from items.
//
// Unless an expression is set on @self, the sorter will always
// compare items as invalid.
//
// The expression must have a return type that can be compared
// numerically, such as %G_TYPE_INT or %G_TYPE_DOUBLE.
func (x *NumericSorter) SetExpression(ExpressionVar *Expression) {

	xNumericSorterSetExpression(x.GoPointer(), ExpressionVar.GoPointer())

}

var xNumericSorterSetSortOrder func(uintptr, SortType)

// Sets whether to sort smaller numbers before larger ones.
func (x *NumericSorter) SetSortOrder(SortOrderVar SortType) {

	xNumericSorterSetSortOrder(x.GoPointer(), SortOrderVar)

}

func (c *NumericSorter) GoPointer() uintptr {
	return c.Ptr
}

func (c *NumericSorter) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GTK"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xNewNumericSorter, lib, "gtk_numeric_sorter_new")

	core.PuregoSafeRegister(&xNumericSorterGetExpression, lib, "gtk_numeric_sorter_get_expression")
	core.PuregoSafeRegister(&xNumericSorterGetSortOrder, lib, "gtk_numeric_sorter_get_sort_order")
	core.PuregoSafeRegister(&xNumericSorterSetExpression, lib, "gtk_numeric_sorter_set_expression")
	core.PuregoSafeRegister(&xNumericSorterSetSortOrder, lib, "gtk_numeric_sorter_set_sort_order")

}
