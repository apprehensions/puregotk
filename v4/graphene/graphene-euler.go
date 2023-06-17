// Package graphene was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package graphene

// Describe a rotation using Euler angles.
//
// The contents of the #graphene_euler_t structure are private
// and should never be accessed directly.
type Euler struct {
	Angles uintptr

	Order EulerOrder
}

// Specify the order of the rotations on each axis.
//
// The %GRAPHENE_EULER_ORDER_DEFAULT value is special, and is used
// as an alias for one of the other orders.
type EulerOrder int

const (

	// Rotate in the default order; the
	//   default order is one of the following enumeration values
	EulerOrderDefaultValue EulerOrder = -1
	// Rotate in the X, Y, and Z order. Deprecated in
	//   Graphene 1.10, it's an alias for %GRAPHENE_EULER_ORDER_SXYZ
	EulerOrderXyzValue EulerOrder = 0
	// Rotate in the Y, Z, and X order. Deprecated in
	//   Graphene 1.10, it's an alias for %GRAPHENE_EULER_ORDER_SYZX
	EulerOrderYzxValue EulerOrder = 1
	// Rotate in the Z, X, and Y order. Deprecated in
	//   Graphene 1.10, it's an alias for %GRAPHENE_EULER_ORDER_SZXY
	EulerOrderZxyValue EulerOrder = 2
	// Rotate in the X, Z, and Y order. Deprecated in
	//   Graphene 1.10, it's an alias for %GRAPHENE_EULER_ORDER_SXZY
	EulerOrderXzyValue EulerOrder = 3
	// Rotate in the Y, X, and Z order. Deprecated in
	//   Graphene 1.10, it's an alias for %GRAPHENE_EULER_ORDER_SYXZ
	EulerOrderYxzValue EulerOrder = 4
	// Rotate in the Z, Y, and X order. Deprecated in
	//   Graphene 1.10, it's an alias for %GRAPHENE_EULER_ORDER_SZYX
	EulerOrderZyxValue EulerOrder = 5
	// Defines a static rotation along the X, Y, and Z axes (Since: 1.10)
	EulerOrderSxyzValue EulerOrder = 6
	// Defines a static rotation along the X, Y, and X axes (Since: 1.10)
	EulerOrderSxyxValue EulerOrder = 7
	// Defines a static rotation along the X, Z, and Y axes (Since: 1.10)
	EulerOrderSxzyValue EulerOrder = 8
	// Defines a static rotation along the X, Z, and X axes (Since: 1.10)
	EulerOrderSxzxValue EulerOrder = 9
	// Defines a static rotation along the Y, Z, and X axes (Since: 1.10)
	EulerOrderSyzxValue EulerOrder = 10
	// Defines a static rotation along the Y, Z, and Y axes (Since: 1.10)
	EulerOrderSyzyValue EulerOrder = 11
	// Defines a static rotation along the Y, X, and Z axes (Since: 1.10)
	EulerOrderSyxzValue EulerOrder = 12
	// Defines a static rotation along the Y, X, and Y axes (Since: 1.10)
	EulerOrderSyxyValue EulerOrder = 13
	// Defines a static rotation along the Z, X, and Y axes (Since: 1.10)
	EulerOrderSzxyValue EulerOrder = 14
	// Defines a static rotation along the Z, X, and Z axes (Since: 1.10)
	EulerOrderSzxzValue EulerOrder = 15
	// Defines a static rotation along the Z, Y, and X axes (Since: 1.10)
	EulerOrderSzyxValue EulerOrder = 16
	// Defines a static rotation along the Z, Y, and Z axes (Since: 1.10)
	EulerOrderSzyzValue EulerOrder = 17
	// Defines a relative rotation along the Z, Y, and X axes (Since: 1.10)
	EulerOrderRzyxValue EulerOrder = 18
	// Defines a relative rotation along the X, Y, and X axes (Since: 1.10)
	EulerOrderRxyxValue EulerOrder = 19
	// Defines a relative rotation along the Y, Z, and X axes (Since: 1.10)
	EulerOrderRyzxValue EulerOrder = 20
	// Defines a relative rotation along the X, Z, and X axes (Since: 1.10)
	EulerOrderRxzxValue EulerOrder = 21
	// Defines a relative rotation along the X, Z, and Y axes (Since: 1.10)
	EulerOrderRxzyValue EulerOrder = 22
	// Defines a relative rotation along the Y, Z, and Y axes (Since: 1.10)
	EulerOrderRyzyValue EulerOrder = 23
	// Defines a relative rotation along the Z, X, and Y axes (Since: 1.10)
	EulerOrderRzxyValue EulerOrder = 24
	// Defines a relative rotation along the Y, X, and Y axes (Since: 1.10)
	EulerOrderRyxyValue EulerOrder = 25
	// Defines a relative rotation along the Y, X, and Z axes (Since: 1.10)
	EulerOrderRyxzValue EulerOrder = 26
	// Defines a relative rotation along the Z, X, and Z axes (Since: 1.10)
	EulerOrderRzxzValue EulerOrder = 27
	// Defines a relative rotation along the X, Y, and Z axes (Since: 1.10)
	EulerOrderRxyzValue EulerOrder = 28
	// Defines a relative rotation along the Z, Y, and Z axes (Since: 1.10)
	EulerOrderRzyzValue EulerOrder = 29
)
