// Package gsk was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gsk

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
)

var xTransformParse func(string, **Transform) bool

// Parses the given @string into a transform and puts it in
// @out_transform.
//
// Strings printed via [method@Gsk.Transform.to_string]
// can be read in again successfully using this function.
//
// If @string does not describe a valid transform, %FALSE is
// returned and %NULL is put in @out_transform.
func TransformParse(StringVar string, OutTransformVar **Transform) bool {

	return xTransformParse(StringVar, OutTransformVar)

}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GSK"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	core.PuregoSafeRegister(&xTransformParse, lib, "gsk_transform_parse")

}
