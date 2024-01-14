// Package adw was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package adw

import (
	"unsafe"

	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/gobject"
	"github.com/jwijenbergh/puregotk/v4/gtk"
)

type SpringAnimationClass struct {
}

func (x *SpringAnimationClass) GoPointer() uintptr {
	return uintptr(unsafe.Pointer(x))
}

// A spring-based [class@Animation].
//
// `AdwSpringAnimation` implements an animation driven by a physical model of a
// spring described by [struct@SpringParams], with a resting position in
// [property@SpringAnimation:value-to], stretched to
// [property@SpringAnimation:value-from].
//
// Since the animation is physically simulated, spring animations don't have a
// fixed duration. The animation will stop when the simulated spring comes to a
// rest - when the amplitude of the oscillations becomes smaller than
// [property@SpringAnimation:epsilon], or immediately when it reaches
// [property@SpringAnimation:value-to] if
// [property@SpringAnimation:clamp] is set to `TRUE`. The estimated duration can
// be obtained with [property@SpringAnimation:estimated-duration].
//
// Due to the nature of spring-driven motion the animation can overshoot
// [property@SpringAnimation:value-to] before coming to a rest. Whether the
// animation will overshoot or not depends on the damping ratio of the spring.
// See [struct@SpringParams] for more information about specific damping ratio
// values.
//
// If [property@SpringAnimation:clamp] is `TRUE`, the animation will abruptly
// end as soon as it reaches the final value, preventing overshooting.
//
// Animations can have an initial velocity value, set via
// [property@SpringAnimation:initial-velocity], which adjusts the curve without
// changing the duration. This makes spring animations useful for deceleration
// at the end of gestures.
//
// If the initial and final values are equal, and the initial velocity is not 0,
// the animation value will bounce and return to its resting position.
type SpringAnimation struct {
	Animation
}

func SpringAnimationNewFromInternalPtr(ptr uintptr) *SpringAnimation {
	cls := &SpringAnimation{}
	cls.Ptr = ptr
	return cls
}

var xNewSpringAnimation func(uintptr, float64, float64, *SpringParams, uintptr) uintptr

// Creates a new `AdwSpringAnimation` on @widget.
//
// The animation will animate @target from @from to @to with the dynamics of a
// spring described by @spring_params.
func NewSpringAnimation(WidgetVar *gtk.Widget, FromVar float64, ToVar float64, SpringParamsVar *SpringParams, TargetVar *AnimationTarget) *SpringAnimation {
	var cls *SpringAnimation

	cret := xNewSpringAnimation(WidgetVar.GoPointer(), FromVar, ToVar, SpringParamsVar, TargetVar.GoPointer())

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &SpringAnimation{}
	cls.Ptr = cret
	return cls
}

var xSpringAnimationCalculateValue func(uintptr, uint) float64

// Calculates the value @self will have at @time.
//
// The time starts at 0 and ends at
// [property@SpringAnimation:estimated_duration].
//
// See also [method@SpringAnimation.calculate_velocity].
func (x *SpringAnimation) CalculateValue(TimeVar uint) float64 {

	cret := xSpringAnimationCalculateValue(x.GoPointer(), TimeVar)
	return cret
}

var xSpringAnimationCalculateVelocity func(uintptr, uint) float64

// Calculates the velocity @self will have at @time.
//
// The time starts at 0 and ends at
// [property@SpringAnimation:estimated_duration].
//
// See also [method@SpringAnimation.calculate_value].
func (x *SpringAnimation) CalculateVelocity(TimeVar uint) float64 {

	cret := xSpringAnimationCalculateVelocity(x.GoPointer(), TimeVar)
	return cret
}

var xSpringAnimationGetClamp func(uintptr) bool

// Gets whether @self should be clamped.
func (x *SpringAnimation) GetClamp() bool {

	cret := xSpringAnimationGetClamp(x.GoPointer())
	return cret
}

var xSpringAnimationGetEpsilon func(uintptr) float64

// Gets the precision of the spring.
func (x *SpringAnimation) GetEpsilon() float64 {

	cret := xSpringAnimationGetEpsilon(x.GoPointer())
	return cret
}

var xSpringAnimationGetEstimatedDuration func(uintptr) uint

// Gets the estimated duration of @self, in milliseconds.
//
// Can be [const@DURATION_INFINITE] if the spring damping is set to 0.
func (x *SpringAnimation) GetEstimatedDuration() uint {

	cret := xSpringAnimationGetEstimatedDuration(x.GoPointer())
	return cret
}

var xSpringAnimationGetInitialVelocity func(uintptr) float64

// Gets the initial velocity of @self.
func (x *SpringAnimation) GetInitialVelocity() float64 {

	cret := xSpringAnimationGetInitialVelocity(x.GoPointer())
	return cret
}

var xSpringAnimationGetSpringParams func(uintptr) *SpringParams

// Gets the physical parameters of the spring of @self.
func (x *SpringAnimation) GetSpringParams() *SpringParams {

	cret := xSpringAnimationGetSpringParams(x.GoPointer())
	return cret
}

var xSpringAnimationGetValueFrom func(uintptr) float64

// Gets the value @self will animate from.
func (x *SpringAnimation) GetValueFrom() float64 {

	cret := xSpringAnimationGetValueFrom(x.GoPointer())
	return cret
}

var xSpringAnimationGetValueTo func(uintptr) float64

// Gets the value @self will animate to.
func (x *SpringAnimation) GetValueTo() float64 {

	cret := xSpringAnimationGetValueTo(x.GoPointer())
	return cret
}

var xSpringAnimationGetVelocity func(uintptr) float64

// Gets the current velocity of @self.
func (x *SpringAnimation) GetVelocity() float64 {

	cret := xSpringAnimationGetVelocity(x.GoPointer())
	return cret
}

var xSpringAnimationSetClamp func(uintptr, bool)

// Sets whether @self should be clamped.
//
// If set to `TRUE`, the animation will abruptly end as soon as it reaches the
// final value, preventing overshooting.
//
// It won't prevent overshooting [property@SpringAnimation:value-from] if a
// relative negative [property@SpringAnimation:initial-velocity] is set.
func (x *SpringAnimation) SetClamp(ClampVar bool) {

	xSpringAnimationSetClamp(x.GoPointer(), ClampVar)

}

var xSpringAnimationSetEpsilon func(uintptr, float64)

// Sets the precision of the spring.
//
// The level of precision used to determine when the animation has come to a
// rest, that is, when the amplitude of the oscillations becomes smaller than
// this value.
//
// If the epsilon value is too small, the animation will take a long time to
// stop after the animated value has stopped visibly changing.
//
// If the epsilon value is too large, the animation will end prematurely.
//
// The default value is 0.001.
func (x *SpringAnimation) SetEpsilon(EpsilonVar float64) {

	xSpringAnimationSetEpsilon(x.GoPointer(), EpsilonVar)

}

var xSpringAnimationSetInitialVelocity func(uintptr, float64)

// Sets the initial velocity of @self.
//
// Initial velocity affects only the animation curve, but not its duration.
func (x *SpringAnimation) SetInitialVelocity(VelocityVar float64) {

	xSpringAnimationSetInitialVelocity(x.GoPointer(), VelocityVar)

}

var xSpringAnimationSetSpringParams func(uintptr, *SpringParams)

// Sets the physical parameters of the spring of @self.
func (x *SpringAnimation) SetSpringParams(SpringParamsVar *SpringParams) {

	xSpringAnimationSetSpringParams(x.GoPointer(), SpringParamsVar)

}

var xSpringAnimationSetValueFrom func(uintptr, float64)

// Sets the value @self will animate from.
//
// The animation will start at this value and end at
// [property@SpringAnimation:value-to].
func (x *SpringAnimation) SetValueFrom(ValueVar float64) {

	xSpringAnimationSetValueFrom(x.GoPointer(), ValueVar)

}

var xSpringAnimationSetValueTo func(uintptr, float64)

// Sets the value @self will animate to.
//
// The animation will start at [property@SpringAnimation:value-from] and end at
// this value.
func (x *SpringAnimation) SetValueTo(ValueVar float64) {

	xSpringAnimationSetValueTo(x.GoPointer(), ValueVar)

}

func (c *SpringAnimation) GoPointer() uintptr {
	return c.Ptr
}

func (c *SpringAnimation) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

func init() {
	lib, err := purego.Dlopen(core.GetPath("ADW"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xNewSpringAnimation, lib, "adw_spring_animation_new")

	core.PuregoSafeRegister(&xSpringAnimationCalculateValue, lib, "adw_spring_animation_calculate_value")
	core.PuregoSafeRegister(&xSpringAnimationCalculateVelocity, lib, "adw_spring_animation_calculate_velocity")
	core.PuregoSafeRegister(&xSpringAnimationGetClamp, lib, "adw_spring_animation_get_clamp")
	core.PuregoSafeRegister(&xSpringAnimationGetEpsilon, lib, "adw_spring_animation_get_epsilon")
	core.PuregoSafeRegister(&xSpringAnimationGetEstimatedDuration, lib, "adw_spring_animation_get_estimated_duration")
	core.PuregoSafeRegister(&xSpringAnimationGetInitialVelocity, lib, "adw_spring_animation_get_initial_velocity")
	core.PuregoSafeRegister(&xSpringAnimationGetSpringParams, lib, "adw_spring_animation_get_spring_params")
	core.PuregoSafeRegister(&xSpringAnimationGetValueFrom, lib, "adw_spring_animation_get_value_from")
	core.PuregoSafeRegister(&xSpringAnimationGetValueTo, lib, "adw_spring_animation_get_value_to")
	core.PuregoSafeRegister(&xSpringAnimationGetVelocity, lib, "adw_spring_animation_get_velocity")
	core.PuregoSafeRegister(&xSpringAnimationSetClamp, lib, "adw_spring_animation_set_clamp")
	core.PuregoSafeRegister(&xSpringAnimationSetEpsilon, lib, "adw_spring_animation_set_epsilon")
	core.PuregoSafeRegister(&xSpringAnimationSetInitialVelocity, lib, "adw_spring_animation_set_initial_velocity")
	core.PuregoSafeRegister(&xSpringAnimationSetSpringParams, lib, "adw_spring_animation_set_spring_params")
	core.PuregoSafeRegister(&xSpringAnimationSetValueFrom, lib, "adw_spring_animation_set_value_from")
	core.PuregoSafeRegister(&xSpringAnimationSetValueTo, lib, "adw_spring_animation_set_value_to")

}
