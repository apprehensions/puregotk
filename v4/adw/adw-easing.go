// Package adw was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package adw

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
)

// Describes the available easing functions for use with
// [class@TimedAnimation].
//
// New values may be added to this enumeration over time.
type Easing int

const (

	// Linear tweening.
	LinearValue Easing = 0
	// Quadratic tweening.
	EaseInQuadValue Easing = 1
	// Quadratic tweening, inverse of `ADW_EASE_IN_QUAD`.
	EaseOutQuadValue Easing = 2
	// Quadratic tweening, combining `ADW_EASE_IN_QUAD` and
	//   `ADW_EASE_OUT_QUAD`.
	EaseInOutQuadValue Easing = 3
	// Cubic tweening.
	EaseInCubicValue Easing = 4
	// Cubic tweening, inverse of `ADW_EASE_IN_CUBIC`.
	EaseOutCubicValue Easing = 5
	// Cubic tweening, combining `ADW_EASE_IN_CUBIC` and
	//   `ADW_EASE_OUT_CUBIC`.
	EaseInOutCubicValue Easing = 6
	// Quartic tweening.
	EaseInQuartValue Easing = 7
	// Quartic tweening, inverse of `ADW_EASE_IN_QUART`.
	EaseOutQuartValue Easing = 8
	// Quartic tweening, combining `ADW_EASE_IN_QUART` and
	//   `ADW_EASE_OUT_QUART`.
	EaseInOutQuartValue Easing = 9
	// Quintic tweening.
	EaseInQuintValue Easing = 10
	// Quintic tweening, inverse of `ADW_EASE_IN_QUINT`.
	EaseOutQuintValue Easing = 11
	// Quintic tweening, combining `ADW_EASE_IN_QUINT` and
	//   `ADW_EASE_OUT_QUINT`.
	EaseInOutQuintValue Easing = 12
	// Sine wave tweening.
	EaseInSineValue Easing = 13
	// Sine wave tweening, inverse of `ADW_EASE_IN_SINE`.
	EaseOutSineValue Easing = 14
	// Sine wave tweening, combining `ADW_EASE_IN_SINE` and
	//   `ADW_EASE_OUT_SINE`.
	EaseInOutSineValue Easing = 15
	// Exponential tweening.
	EaseInExpoValue Easing = 16
	// Exponential tweening, inverse of `ADW_EASE_IN_EXPO`.
	EaseOutExpoValue Easing = 17
	// Exponential tweening, combining `ADW_EASE_IN_EXPO` and
	//   `ADW_EASE_OUT_EXPO`.
	EaseInOutExpoValue Easing = 18
	// Circular tweening.
	EaseInCircValue Easing = 19
	// Circular tweening, inverse of `ADW_EASE_IN_CIRC`.
	EaseOutCircValue Easing = 20
	// Circular tweening, combining `ADW_EASE_IN_CIRC` and
	//   `ADW_EASE_OUT_CIRC`.
	EaseInOutCircValue Easing = 21
	// Elastic tweening, with offshoot on start.
	EaseInElasticValue Easing = 22
	// Elastic tweening, with offshoot on end, inverse of
	//   `ADW_EASE_IN_ELASTIC`.
	EaseOutElasticValue Easing = 23
	// Elastic tweening, with offshoot on both ends,
	//   combining `ADW_EASE_IN_ELASTIC` and `ADW_EASE_OUT_ELASTIC`.
	EaseInOutElasticValue Easing = 24
	// Overshooting cubic tweening, with backtracking on start.
	EaseInBackValue Easing = 25
	// Overshooting cubic tweening, with backtracking on end,
	//   inverse of `ADW_EASE_IN_BACK`.
	EaseOutBackValue Easing = 26
	// Overshooting cubic tweening, with backtracking on both
	//   ends, combining `ADW_EASE_IN_BACK` and `ADW_EASE_OUT_BACK`.
	EaseInOutBackValue Easing = 27
	// Exponentially decaying parabolic (bounce) tweening,
	//   on start.
	EaseInBounceValue Easing = 28
	// Exponentially decaying parabolic (bounce) tweening,
	//   with bounce on end, inverse of `ADW_EASE_IN_BOUNCE`.
	EaseOutBounceValue Easing = 29
	// Exponentially decaying parabolic (bounce) tweening,
	//   with bounce on both ends, combining `ADW_EASE_IN_BOUNCE` and
	//   `ADW_EASE_OUT_BOUNCE`.
	EaseInOutBounceValue Easing = 30
)

var xEasingEase func(Easing, float64) float64

// Computes easing with @easing for @value.
//
// @value should generally be in the [0, 1] range.
func EasingEase(SelfVar Easing, ValueVar float64) float64 {

	return xEasingEase(SelfVar, ValueVar)

}

func init() {
	lib, err := purego.Dlopen(core.GetPath("ADW"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	core.PuregoSafeRegister(&xEasingEase, lib, "adw_easing_ease")

}
