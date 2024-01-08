// Package graphene was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package graphene

import (
	"unsafe"

	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
)

// A structure capable of holding a vector with three dimensions: x, y, and z.
//
// The contents of the #graphene_vec3_t structure are private and should
// never be accessed directly.
type Vec3 struct {
	Value uintptr
}

func (x *Vec3) GoPointer() uintptr {
	return uintptr(unsafe.Pointer(x))
}

var xAllocVec3 func() *Vec3

// Allocates a new #graphene_vec3_t structure.
//
// The contents of the returned structure are undefined.
//
// Use graphene_vec3_init() to initialize the vector.
func AllocVec3() *Vec3 {

	cret := xAllocVec3()
	return cret
}

var xVec3Add func(uintptr, *Vec3, *Vec3)

// Adds each component of the two given vectors.
func (x *Vec3) Add(BVar *Vec3, ResVar *Vec3) {

	xVec3Add(x.GoPointer(), BVar, ResVar)

}

var xVec3Cross func(uintptr, *Vec3, *Vec3)

// Computes the cross product of the two given vectors.
func (x *Vec3) Cross(BVar *Vec3, ResVar *Vec3) {

	xVec3Cross(x.GoPointer(), BVar, ResVar)

}

var xVec3Divide func(uintptr, *Vec3, *Vec3)

// Divides each component of the first operand @a by the corresponding
// component of the second operand @b, and places the results into the
// vector @res.
func (x *Vec3) Divide(BVar *Vec3, ResVar *Vec3) {

	xVec3Divide(x.GoPointer(), BVar, ResVar)

}

var xVec3Dot func(uintptr, *Vec3) float32

// Computes the dot product of the two given vectors.
func (x *Vec3) Dot(BVar *Vec3) float32 {

	cret := xVec3Dot(x.GoPointer(), BVar)
	return cret
}

var xVec3Equal func(uintptr, *Vec3) bool

// Checks whether the two given #graphene_vec3_t are equal.
func (x *Vec3) Equal(V2Var *Vec3) bool {

	cret := xVec3Equal(x.GoPointer(), V2Var)
	return cret
}

var xVec3Free func(uintptr)

// Frees the resources allocated by @v
func (x *Vec3) Free() {

	xVec3Free(x.GoPointer())

}

var xVec3GetX func(uintptr) float32

// Retrieves the first component of the given vector @v.
func (x *Vec3) GetX() float32 {

	cret := xVec3GetX(x.GoPointer())
	return cret
}

var xVec3GetXy func(uintptr, *Vec2)

// Creates a #graphene_vec2_t that contains the first and second
// components of the given #graphene_vec3_t.
func (x *Vec3) GetXy(ResVar *Vec2) {

	xVec3GetXy(x.GoPointer(), ResVar)

}

var xVec3GetXy0 func(uintptr, *Vec3)

// Creates a #graphene_vec3_t that contains the first two components of
// the given #graphene_vec3_t, and the third component set to 0.
func (x *Vec3) GetXy0(ResVar *Vec3) {

	xVec3GetXy0(x.GoPointer(), ResVar)

}

var xVec3GetXyz0 func(uintptr, *Vec4)

// Converts a #graphene_vec3_t in a #graphene_vec4_t using 0.0
// as the value for the fourth component of the resulting vector.
func (x *Vec3) GetXyz0(ResVar *Vec4) {

	xVec3GetXyz0(x.GoPointer(), ResVar)

}

var xVec3GetXyz1 func(uintptr, *Vec4)

// Converts a #graphene_vec3_t in a #graphene_vec4_t using 1.0
// as the value for the fourth component of the resulting vector.
func (x *Vec3) GetXyz1(ResVar *Vec4) {

	xVec3GetXyz1(x.GoPointer(), ResVar)

}

var xVec3GetXyzw func(uintptr, float32, *Vec4)

// Converts a #graphene_vec3_t in a #graphene_vec4_t using @w as
// the value of the fourth component of the resulting vector.
func (x *Vec3) GetXyzw(WVar float32, ResVar *Vec4) {

	xVec3GetXyzw(x.GoPointer(), WVar, ResVar)

}

var xVec3GetY func(uintptr) float32

// Retrieves the second component of the given vector @v.
func (x *Vec3) GetY() float32 {

	cret := xVec3GetY(x.GoPointer())
	return cret
}

var xVec3GetZ func(uintptr) float32

// Retrieves the third component of the given vector @v.
func (x *Vec3) GetZ() float32 {

	cret := xVec3GetZ(x.GoPointer())
	return cret
}

var xVec3Init func(uintptr, float32, float32, float32) *Vec3

// Initializes a #graphene_vec3_t using the given values.
//
// This function can be called multiple times.
func (x *Vec3) Init(XVar float32, YVar float32, ZVar float32) *Vec3 {

	cret := xVec3Init(x.GoPointer(), XVar, YVar, ZVar)
	return cret
}

var xVec3InitFromFloat func(uintptr, uintptr) *Vec3

// Initializes a #graphene_vec3_t with the values from an array.
func (x *Vec3) InitFromFloat(SrcVar uintptr) *Vec3 {

	cret := xVec3InitFromFloat(x.GoPointer(), SrcVar)
	return cret
}

var xVec3InitFromVec3 func(uintptr, *Vec3) *Vec3

// Initializes a #graphene_vec3_t with the values of another
// #graphene_vec3_t.
func (x *Vec3) InitFromVec3(SrcVar *Vec3) *Vec3 {

	cret := xVec3InitFromVec3(x.GoPointer(), SrcVar)
	return cret
}

var xVec3Interpolate func(uintptr, *Vec3, float64, *Vec3)

// Linearly interpolates @v1 and @v2 using the given @factor.
func (x *Vec3) Interpolate(V2Var *Vec3, FactorVar float64, ResVar *Vec3) {

	xVec3Interpolate(x.GoPointer(), V2Var, FactorVar, ResVar)

}

var xVec3Length func(uintptr) float32

// Retrieves the length of the given vector @v.
func (x *Vec3) Length() float32 {

	cret := xVec3Length(x.GoPointer())
	return cret
}

var xVec3Max func(uintptr, *Vec3, *Vec3)

// Compares each component of the two given vectors and creates a
// vector that contains the maximum values.
func (x *Vec3) Max(BVar *Vec3, ResVar *Vec3) {

	xVec3Max(x.GoPointer(), BVar, ResVar)

}

var xVec3Min func(uintptr, *Vec3, *Vec3)

// Compares each component of the two given vectors and creates a
// vector that contains the minimum values.
func (x *Vec3) Min(BVar *Vec3, ResVar *Vec3) {

	xVec3Min(x.GoPointer(), BVar, ResVar)

}

var xVec3Multiply func(uintptr, *Vec3, *Vec3)

// Multiplies each component of the two given vectors.
func (x *Vec3) Multiply(BVar *Vec3, ResVar *Vec3) {

	xVec3Multiply(x.GoPointer(), BVar, ResVar)

}

var xVec3Near func(uintptr, *Vec3, float32) bool

// Compares the two given #graphene_vec3_t vectors and checks
// whether their values are within the given @epsilon.
func (x *Vec3) Near(V2Var *Vec3, EpsilonVar float32) bool {

	cret := xVec3Near(x.GoPointer(), V2Var, EpsilonVar)
	return cret
}

var xVec3Negate func(uintptr, *Vec3)

// Negates the given #graphene_vec3_t.
func (x *Vec3) Negate(ResVar *Vec3) {

	xVec3Negate(x.GoPointer(), ResVar)

}

var xVec3Normalize func(uintptr, *Vec3)

// Normalizes the given #graphene_vec3_t.
func (x *Vec3) Normalize(ResVar *Vec3) {

	xVec3Normalize(x.GoPointer(), ResVar)

}

var xVec3Scale func(uintptr, float32, *Vec3)

// Multiplies all components of the given vector with the given scalar @factor.
func (x *Vec3) Scale(FactorVar float32, ResVar *Vec3) {

	xVec3Scale(x.GoPointer(), FactorVar, ResVar)

}

var xVec3Subtract func(uintptr, *Vec3, *Vec3)

// Subtracts from each component of the first operand @a the
// corresponding component of the second operand @b and places
// each result into the components of @res.
func (x *Vec3) Subtract(BVar *Vec3, ResVar *Vec3) {

	xVec3Subtract(x.GoPointer(), BVar, ResVar)

}

var xVec3ToFloat func(uintptr, uintptr)

// Copies the components of a #graphene_vec3_t into the given array.
func (x *Vec3) ToFloat(DestVar uintptr) {

	xVec3ToFloat(x.GoPointer(), DestVar)

}

var xVec3One func() *Vec3

// Provides a constant pointer to a vector with three components,
// all sets to 1.
func Vec3One() *Vec3 {

	cret := xVec3One()
	return cret
}

var xVec3XAxis func() *Vec3

// Provides a constant pointer to a vector with three components
// with values set to (1, 0, 0).
func Vec3XAxis() *Vec3 {

	cret := xVec3XAxis()
	return cret
}

var xVec3YAxis func() *Vec3

// Provides a constant pointer to a vector with three components
// with values set to (0, 1, 0).
func Vec3YAxis() *Vec3 {

	cret := xVec3YAxis()
	return cret
}

var xVec3ZAxis func() *Vec3

// Provides a constant pointer to a vector with three components
// with values set to (0, 0, 1).
func Vec3ZAxis() *Vec3 {

	cret := xVec3ZAxis()
	return cret
}

var xVec3Zero func() *Vec3

// Provides a constant pointer to a vector with three components,
// all sets to 0.
func Vec3Zero() *Vec3 {

	cret := xVec3Zero()
	return cret
}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GRAPHENE"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	core.PuregoSafeRegister(&xVec3One, lib, "graphene_vec3_one")
	core.PuregoSafeRegister(&xVec3XAxis, lib, "graphene_vec3_x_axis")
	core.PuregoSafeRegister(&xVec3YAxis, lib, "graphene_vec3_y_axis")
	core.PuregoSafeRegister(&xVec3ZAxis, lib, "graphene_vec3_z_axis")
	core.PuregoSafeRegister(&xVec3Zero, lib, "graphene_vec3_zero")

	core.PuregoSafeRegister(&xAllocVec3, lib, "graphene_vec3_alloc")

	core.PuregoSafeRegister(&xVec3Add, lib, "graphene_vec3_add")
	core.PuregoSafeRegister(&xVec3Cross, lib, "graphene_vec3_cross")
	core.PuregoSafeRegister(&xVec3Divide, lib, "graphene_vec3_divide")
	core.PuregoSafeRegister(&xVec3Dot, lib, "graphene_vec3_dot")
	core.PuregoSafeRegister(&xVec3Equal, lib, "graphene_vec3_equal")
	core.PuregoSafeRegister(&xVec3Free, lib, "graphene_vec3_free")
	core.PuregoSafeRegister(&xVec3GetX, lib, "graphene_vec3_get_x")
	core.PuregoSafeRegister(&xVec3GetXy, lib, "graphene_vec3_get_xy")
	core.PuregoSafeRegister(&xVec3GetXy0, lib, "graphene_vec3_get_xy0")
	core.PuregoSafeRegister(&xVec3GetXyz0, lib, "graphene_vec3_get_xyz0")
	core.PuregoSafeRegister(&xVec3GetXyz1, lib, "graphene_vec3_get_xyz1")
	core.PuregoSafeRegister(&xVec3GetXyzw, lib, "graphene_vec3_get_xyzw")
	core.PuregoSafeRegister(&xVec3GetY, lib, "graphene_vec3_get_y")
	core.PuregoSafeRegister(&xVec3GetZ, lib, "graphene_vec3_get_z")
	core.PuregoSafeRegister(&xVec3Init, lib, "graphene_vec3_init")
	core.PuregoSafeRegister(&xVec3InitFromFloat, lib, "graphene_vec3_init_from_float")
	core.PuregoSafeRegister(&xVec3InitFromVec3, lib, "graphene_vec3_init_from_vec3")
	core.PuregoSafeRegister(&xVec3Interpolate, lib, "graphene_vec3_interpolate")
	core.PuregoSafeRegister(&xVec3Length, lib, "graphene_vec3_length")
	core.PuregoSafeRegister(&xVec3Max, lib, "graphene_vec3_max")
	core.PuregoSafeRegister(&xVec3Min, lib, "graphene_vec3_min")
	core.PuregoSafeRegister(&xVec3Multiply, lib, "graphene_vec3_multiply")
	core.PuregoSafeRegister(&xVec3Near, lib, "graphene_vec3_near")
	core.PuregoSafeRegister(&xVec3Negate, lib, "graphene_vec3_negate")
	core.PuregoSafeRegister(&xVec3Normalize, lib, "graphene_vec3_normalize")
	core.PuregoSafeRegister(&xVec3Scale, lib, "graphene_vec3_scale")
	core.PuregoSafeRegister(&xVec3Subtract, lib, "graphene_vec3_subtract")
	core.PuregoSafeRegister(&xVec3ToFloat, lib, "graphene_vec3_to_float")

}
