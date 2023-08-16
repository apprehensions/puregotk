// Package pango was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package pango

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
)

// A `PangoTabArray` contains an array of tab stops.
//
// `PangoTabArray` can be used to set tab stops in a `PangoLayout`.
// Each tab stop has an alignment, a position, and optionally
// a character to use as decimal point.
type TabArray struct {
}

// `PangoTabAlign` specifies where the text appears relative to the tab stop
// position.
//
// Support for tab alignments other than %PANGO_TAB_LEFT was added
// in Pango 1.50.
type TabAlign int

const (

	// the text appears to the right of the tab stop position
	TabLeftValue TabAlign = 0
	// the text appears to the left of the tab stop position
	//   until the available space is filled
	TabRightValue TabAlign = 1
	// the text is centered at the tab stop position
	//   until the available space is filled
	TabCenterValue TabAlign = 2
	// text before the first occurrence of the decimal point
	//   character appears to the left of the tab stop position (until the available
	//   space is filled), the rest to the right
	TabDecimalValue TabAlign = 3
)

var xTabArrayFromString func(string) *TabArray

// Deserializes a `PangoTabArray` from a string.
//
// This is the counterpart to [method@Pango.TabArray.to_string].
// See that functions for details about the format.
func TabArrayFromString(TextVar string) *TabArray {

	cret := xTabArrayFromString(TextVar)
	return cret
}

func init() {
	lib, err := purego.Dlopen(core.GetPath("PANGO"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	core.PuregoSafeRegister(&xTabArrayFromString, lib, "pango_tab_array_from_string")

}
