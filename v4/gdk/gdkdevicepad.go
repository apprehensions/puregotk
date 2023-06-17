// Package gdk was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gdk

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
)

type DevicePadInterface struct {
}

// `GdkDevicePad` is an interface implemented by devices of type
// %GDK_SOURCE_TABLET_PAD
//
// It allows querying the features provided by the pad device.
//
// Tablet pads may contain one or more groups, each containing a subset
// of the buttons/rings/strips available. [method@Gdk.DevicePad.get_n_groups]
// can be used to obtain the number of groups, [method@Gdk.DevicePad.get_n_features]
// and [method@Gdk.DevicePad.get_feature_group] can be combined to find out
// the number of buttons/rings/strips the device has, and how are they grouped.
//
// Each of those groups have different modes, which may be used to map each
// individual pad feature to multiple actions. Only one mode is effective
// (current) for each given group, different groups may have different
// current modes. The number of available modes in a group can be found
// out through [method@Gdk.DevicePad.get_group_n_modes], and the current mode
// for a given group will be notified through events of type `GDK_PAD_GROUP_MODE`.
type DevicePad interface {
	GoPointer() uintptr
	SetGoPointer(uintptr)
	GetFeatureGroup(FeatureVar DevicePadFeature, FeatureIdxVar int32) int32
	GetGroupNModes(GroupIdxVar int32) int32
	GetNFeatures(FeatureVar DevicePadFeature) int32
	GetNGroups() int32
}
type DevicePadBase struct {
	Ptr uintptr
}

func (x *DevicePadBase) GoPointer() uintptr {
	return x.Ptr
}

func (x *DevicePadBase) SetGoPointer(ptr uintptr) {
	x.Ptr = ptr
}

// Returns the group the given @feature and @idx belong to.
//
// f the feature or index do not exist in @pad, -1 is returned.
func (x *DevicePadBase) GetFeatureGroup(FeatureVar DevicePadFeature, FeatureIdxVar int32) int32 {

	return XGdkDevicePadGetFeatureGroup(x.GoPointer(), FeatureVar, FeatureIdxVar)

}

// Returns the number of modes that @group may have.
func (x *DevicePadBase) GetGroupNModes(GroupIdxVar int32) int32 {

	return XGdkDevicePadGetGroupNModes(x.GoPointer(), GroupIdxVar)

}

// Returns the number of features a tablet pad has.
func (x *DevicePadBase) GetNFeatures(FeatureVar DevicePadFeature) int32 {

	return XGdkDevicePadGetNFeatures(x.GoPointer(), FeatureVar)

}

// Returns the number of groups this pad device has.
//
// Pads have at least one group. A pad group is a subcollection of
// buttons/strip/rings that is affected collectively by a same
// current mode.
func (x *DevicePadBase) GetNGroups() int32 {

	return XGdkDevicePadGetNGroups(x.GoPointer())

}

var XGdkDevicePadGetFeatureGroup func(uintptr, DevicePadFeature, int32) int32
var XGdkDevicePadGetGroupNModes func(uintptr, int32) int32
var XGdkDevicePadGetNFeatures func(uintptr, DevicePadFeature) int32
var XGdkDevicePadGetNGroups func(uintptr) int32

// A pad feature.
type DevicePadFeature int

const (

	// a button
	DevicePadFeatureButtonValue DevicePadFeature = 0
	// a ring-shaped interactive area
	DevicePadFeatureRingValue DevicePadFeature = 1
	// a straight interactive area
	DevicePadFeatureStripValue DevicePadFeature = 2
)

func init() {
	lib, err := purego.Dlopen(core.GetPath("GDK"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&XGdkDevicePadGetFeatureGroup, lib, "gdk_device_pad_get_feature_group")
	core.PuregoSafeRegister(&XGdkDevicePadGetGroupNModes, lib, "gdk_device_pad_get_group_n_modes")
	core.PuregoSafeRegister(&XGdkDevicePadGetNFeatures, lib, "gdk_device_pad_get_n_features")
	core.PuregoSafeRegister(&XGdkDevicePadGetNGroups, lib, "gdk_device_pad_get_n_groups")

}
