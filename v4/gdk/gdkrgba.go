// Package gdk was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gdk

// A `GdkRGBA` is used to represent a color, in a way that is compatible
// with cairo’s notion of color.
//
// `GdkRGBA` is a convenient way to pass colors around. It’s based on
// cairo’s way to deal with colors and mirrors its behavior. All values
// are in the range from 0.0 to 1.0 inclusive. So the color
// (0.0, 0.0, 0.0, 0.0) represents transparent black and
// (1.0, 1.0, 1.0, 1.0) is opaque white. Other values will
// be clamped to this range when drawing.
type RGBA struct {
	Red float32

	Green float32

	Blue float32

	Alpha float32
}
