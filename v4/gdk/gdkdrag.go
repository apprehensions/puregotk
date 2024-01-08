// Package gdk was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gdk

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/gobject"
)

// Used in `GdkDrag` to the reason of a cancelled DND operation.
type DragCancelReason int

const (

	// There is no suitable drop target.
	DragCancelNoTargetValue DragCancelReason = 0
	// Drag cancelled by the user
	DragCancelUserCancelledValue DragCancelReason = 1
	// Unspecified error.
	DragCancelErrorValue DragCancelReason = 2
)

var xDragActionIsUnique func(DragAction) bool

// Checks if @action represents a single action or includes
// multiple actions.
//
// When @action is 0 - ie no action was given, %TRUE
// is returned.
func DragActionIsUnique(ActionVar DragAction) bool {

	cret := xDragActionIsUnique(ActionVar)
	return cret
}

// The `GdkDrag` object represents the source of an ongoing DND operation.
//
// A `GdkDrag` is created when a drag is started, and stays alive for duration of
// the DND operation. After a drag has been started with [func@Gdk.Drag.begin],
// the caller gets informed about the status of the ongoing drag operation
// with signals on the `GdkDrag` object.
//
// GTK provides a higher level abstraction based on top of these functions,
// and so they are not normally needed in GTK applications. See the
// "Drag and Drop" section of the GTK documentation for more information.
type Drag struct {
	gobject.Object
}

func DragNewFromInternalPtr(ptr uintptr) *Drag {
	cls := &Drag{}
	cls.Ptr = ptr
	return cls
}

var xDragDropDone func(uintptr, bool)

// Informs GDK that the drop ended.
//
// Passing %FALSE for @success may trigger a drag cancellation
// animation.
//
// This function is called by the drag source, and should be the
// last call before dropping the reference to the @drag.
//
// The `GdkDrag` will only take the first [method@Gdk.Drag.drop_done]
// call as effective, if this function is called multiple times,
// all subsequent calls will be ignored.
func (x *Drag) DropDone(SuccessVar bool) {

	xDragDropDone(x.GoPointer(), SuccessVar)

}

var xDragGetActions func(uintptr) DragAction

// Determines the bitmask of possible actions proposed by the source.
func (x *Drag) GetActions() DragAction {

	cret := xDragGetActions(x.GoPointer())
	return cret
}

var xDragGetContent func(uintptr) uintptr

// Returns the `GdkContentProvider` associated to the `GdkDrag` object.
func (x *Drag) GetContent() *ContentProvider {
	var cls *ContentProvider

	cret := xDragGetContent(x.GoPointer())

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &ContentProvider{}
	cls.Ptr = cret
	return cls
}

var xDragGetDevice func(uintptr) uintptr

// Returns the `GdkDevice` associated to the `GdkDrag` object.
func (x *Drag) GetDevice() *Device {
	var cls *Device

	cret := xDragGetDevice(x.GoPointer())

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &Device{}
	cls.Ptr = cret
	return cls
}

var xDragGetDisplay func(uintptr) uintptr

// Gets the `GdkDisplay` that the drag object was created for.
func (x *Drag) GetDisplay() *Display {
	var cls *Display

	cret := xDragGetDisplay(x.GoPointer())

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &Display{}
	cls.Ptr = cret
	return cls
}

var xDragGetDragSurface func(uintptr) uintptr

// Returns the surface on which the drag icon should be rendered
// during the drag operation.
//
// Note that the surface may not be available until the drag operation
// has begun. GDK will move the surface in accordance with the ongoing
// drag operation. The surface is owned by @drag and will be destroyed
// when the drag operation is over.
func (x *Drag) GetDragSurface() *Surface {
	var cls *Surface

	cret := xDragGetDragSurface(x.GoPointer())

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &Surface{}
	cls.Ptr = cret
	return cls
}

var xDragGetFormats func(uintptr) *ContentFormats

// Retrieves the formats supported by this `GdkDrag` object.
func (x *Drag) GetFormats() *ContentFormats {

	cret := xDragGetFormats(x.GoPointer())
	return cret
}

var xDragGetSelectedAction func(uintptr) DragAction

// Determines the action chosen by the drag destination.
func (x *Drag) GetSelectedAction() DragAction {

	cret := xDragGetSelectedAction(x.GoPointer())
	return cret
}

var xDragGetSurface func(uintptr) uintptr

// Returns the `GdkSurface` where the drag originates.
func (x *Drag) GetSurface() *Surface {
	var cls *Surface

	cret := xDragGetSurface(x.GoPointer())

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &Surface{}
	cls.Ptr = cret
	return cls
}

var xDragSetHotspot func(uintptr, int, int)

// Sets the position of the drag surface that will be kept
// under the cursor hotspot.
//
// Initially, the hotspot is at the top left corner of the drag surface.
func (x *Drag) SetHotspot(HotXVar int, HotYVar int) {

	xDragSetHotspot(x.GoPointer(), HotXVar, HotYVar)

}

func (c *Drag) GoPointer() uintptr {
	return c.Ptr
}

func (c *Drag) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// Emitted when the drag operation is cancelled.
func (x *Drag) ConnectCancel(cb func(Drag, DragCancelReason)) uint32 {
	fcb := func(clsPtr uintptr, ReasonVarp DragCancelReason) {
		fa := Drag{}
		fa.Ptr = clsPtr

		cb(fa, ReasonVarp)

	}
	return gobject.SignalConnect(x.GoPointer(), "cancel", purego.NewCallback(fcb))
}

// Emitted when the destination side has finished reading all data.
//
// The drag object can now free all miscellaneous data.
func (x *Drag) ConnectDndFinished(cb func(Drag)) uint32 {
	fcb := func(clsPtr uintptr) {
		fa := Drag{}
		fa.Ptr = clsPtr

		cb(fa)

	}
	return gobject.SignalConnect(x.GoPointer(), "dnd-finished", purego.NewCallback(fcb))
}

// Emitted when the drop operation is performed on an accepting client.
func (x *Drag) ConnectDropPerformed(cb func(Drag)) uint32 {
	fcb := func(clsPtr uintptr) {
		fa := Drag{}
		fa.Ptr = clsPtr

		cb(fa)

	}
	return gobject.SignalConnect(x.GoPointer(), "drop-performed", purego.NewCallback(fcb))
}

var xDragBegin func(uintptr, uintptr, uintptr, DragAction, float64, float64) uintptr

// Starts a drag and creates a new drag context for it.
//
// This function is called by the drag source. After this call, you
// probably want to set up the drag icon using the surface returned
// by [method@Gdk.Drag.get_drag_surface].
//
// This function returns a reference to the [class@Gdk.Drag] object,
// but GTK keeps its own reference as well, as long as the DND operation
// is going on.
//
// Note: if @actions include %GDK_ACTION_MOVE, you need to listen for
// the [signal@Gdk.Drag::dnd-finished] signal and delete the data at
// the source if [method@Gdk.Drag.get_selected_action] returns
// %GDK_ACTION_MOVE.
func DragBegin(SurfaceVar *Surface, DeviceVar *Device, ContentVar *ContentProvider, ActionsVar DragAction, DxVar float64, DyVar float64) *Drag {
	var cls *Drag

	cret := xDragBegin(SurfaceVar.GoPointer(), DeviceVar.GoPointer(), ContentVar.GoPointer(), ActionsVar, DxVar, DyVar)

	if cret == 0 {
		return nil
	}
	cls = &Drag{}
	cls.Ptr = cret
	return cls
}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GDK"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	core.PuregoSafeRegister(&xDragActionIsUnique, lib, "gdk_drag_action_is_unique")

	core.PuregoSafeRegister(&xDragDropDone, lib, "gdk_drag_drop_done")
	core.PuregoSafeRegister(&xDragGetActions, lib, "gdk_drag_get_actions")
	core.PuregoSafeRegister(&xDragGetContent, lib, "gdk_drag_get_content")
	core.PuregoSafeRegister(&xDragGetDevice, lib, "gdk_drag_get_device")
	core.PuregoSafeRegister(&xDragGetDisplay, lib, "gdk_drag_get_display")
	core.PuregoSafeRegister(&xDragGetDragSurface, lib, "gdk_drag_get_drag_surface")
	core.PuregoSafeRegister(&xDragGetFormats, lib, "gdk_drag_get_formats")
	core.PuregoSafeRegister(&xDragGetSelectedAction, lib, "gdk_drag_get_selected_action")
	core.PuregoSafeRegister(&xDragGetSurface, lib, "gdk_drag_get_surface")
	core.PuregoSafeRegister(&xDragSetHotspot, lib, "gdk_drag_set_hotspot")

	core.PuregoSafeRegister(&xDragBegin, lib, "gdk_drag_begin")

}
