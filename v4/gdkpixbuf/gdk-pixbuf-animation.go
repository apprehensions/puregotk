// Package gdkpixbuf was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gdkpixbuf

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/gio"
	"github.com/jwijenbergh/puregotk/v4/glib"
	"github.com/jwijenbergh/puregotk/v4/gobject"
)

// Modules supporting animations must derive a type from
// #GdkPixbufAnimation, providing suitable implementations of the
// virtual functions.
type PixbufAnimationClass struct {
	ParentClass uintptr
}

// Modules supporting animations must derive a type from
// #GdkPixbufAnimationIter, providing suitable implementations of the
// virtual functions.
type PixbufAnimationIterClass struct {
	ParentClass uintptr
}

// An opaque object representing an animation.
//
// The GdkPixBuf library provides a simple mechanism to load and
// represent animations. An animation is conceptually a series of
// frames to be displayed over time.
//
// The animation may not be represented as a series of frames
// internally; for example, it may be stored as a sprite and
// instructions for moving the sprite around a background.
//
// To display an animation you don't need to understand its
// representation, however; you just ask `GdkPixbuf` what should
// be displayed at a given point in time.
type PixbufAnimation struct {
	gobject.Object
}

func PixbufAnimationNewFromInternalPtr(ptr uintptr) *PixbufAnimation {
	cls := &PixbufAnimation{}
	cls.Ptr = ptr
	return cls
}

var xNewFromFilePixbufAnimation func(string, **glib.Error) uintptr

// Creates a new animation by loading it from a file.
//
// The file format is detected automatically.
//
// If the file's format does not support multi-frame images, then an animation
// with a single frame will be created.
//
// Possible errors are in the `GDK_PIXBUF_ERROR` and `G_FILE_ERROR` domains.
func NewFromFilePixbufAnimation(FilenameVar string) (*PixbufAnimation, error) {
	var cls *PixbufAnimation
	var cerr *glib.Error

	cret := xNewFromFilePixbufAnimation(FilenameVar, &cerr)

	if cret == 0 {
		return nil, cerr
	}
	cls = &PixbufAnimation{}
	cls.Ptr = cret
	if cerr == nil {
		return cls, nil
	}
	return cls, cerr

}

var xNewFromResourcePixbufAnimation func(string, **glib.Error) uintptr

// Creates a new pixbuf animation by loading an image from an resource.
//
// The file format is detected automatically. If `NULL` is returned, then
// @error will be set.
func NewFromResourcePixbufAnimation(ResourcePathVar string) (*PixbufAnimation, error) {
	var cls *PixbufAnimation
	var cerr *glib.Error

	cret := xNewFromResourcePixbufAnimation(ResourcePathVar, &cerr)

	if cret == 0 {
		return nil, cerr
	}
	cls = &PixbufAnimation{}
	cls.Ptr = cret
	if cerr == nil {
		return cls, nil
	}
	return cls, cerr

}

var xNewFromStreamPixbufAnimation func(uintptr, uintptr, **glib.Error) uintptr

// Creates a new animation by loading it from an input stream.
//
// The file format is detected automatically.
//
// If `NULL` is returned, then @error will be set.
//
// The @cancellable can be used to abort the operation from another thread.
// If the operation was cancelled, the error `G_IO_ERROR_CANCELLED` will be
// returned. Other possible errors are in the `GDK_PIXBUF_ERROR` and
// `G_IO_ERROR` domains.
//
// The stream is not closed.
func NewFromStreamPixbufAnimation(StreamVar *gio.InputStream, CancellableVar *gio.Cancellable) (*PixbufAnimation, error) {
	var cls *PixbufAnimation
	var cerr *glib.Error

	cret := xNewFromStreamPixbufAnimation(StreamVar.GoPointer(), CancellableVar.GoPointer(), &cerr)

	if cret == 0 {
		return nil, cerr
	}
	cls = &PixbufAnimation{}
	cls.Ptr = cret
	if cerr == nil {
		return cls, nil
	}
	return cls, cerr

}

var xNewFromStreamFinishPixbufAnimation func(uintptr, **glib.Error) uintptr

// Finishes an asynchronous pixbuf animation creation operation started with
// [func@GdkPixbuf.PixbufAnimation.new_from_stream_async].
func NewFromStreamFinishPixbufAnimation(AsyncResultVar gio.AsyncResult) (*PixbufAnimation, error) {
	var cls *PixbufAnimation
	var cerr *glib.Error

	cret := xNewFromStreamFinishPixbufAnimation(AsyncResultVar.GoPointer(), &cerr)

	if cret == 0 {
		return nil, cerr
	}
	cls = &PixbufAnimation{}
	cls.Ptr = cret
	if cerr == nil {
		return cls, nil
	}
	return cls, cerr

}

var xPixbufAnimationGetHeight func(uintptr) int

// Queries the height of the bounding box of a pixbuf animation.
func (x *PixbufAnimation) GetHeight() int {

	cret := xPixbufAnimationGetHeight(x.GoPointer())
	return cret
}

var xPixbufAnimationGetIter func(uintptr, *glib.TimeVal) uintptr

// Get an iterator for displaying an animation.
//
// The iterator provides the frames that should be displayed at a
// given time.
//
// @start_time would normally come from g_get_current_time(), and marks
// the beginning of animation playback. After creating an iterator, you
// should immediately display the pixbuf returned by
// gdk_pixbuf_animation_iter_get_pixbuf(). Then, you should install
// a timeout (with g_timeout_add()) or by some other mechanism ensure
// that you'll update the image after
// gdk_pixbuf_animation_iter_get_delay_time() milliseconds. Each time
// the image is updated, you should reinstall the timeout with the new,
// possibly-changed delay time.
//
// As a shortcut, if @start_time is `NULL`, the result of
// g_get_current_time() will be used automatically.
//
// To update the image (i.e. possibly change the result of
// gdk_pixbuf_animation_iter_get_pixbuf() to a new frame of the animation),
// call gdk_pixbuf_animation_iter_advance().
//
// If you're using #GdkPixbufLoader, in addition to updating the image
// after the delay time, you should also update it whenever you
// receive the area_updated signal and
// gdk_pixbuf_animation_iter_on_currently_loading_frame() returns
// `TRUE`. In this case, the frame currently being fed into the loader
// has received new data, so needs to be refreshed. The delay time for
// a frame may also be modified after an area_updated signal, for
// example if the delay time for a frame is encoded in the data after
// the frame itself. So your timeout should be reinstalled after any
// area_updated signal.
//
// A delay time of -1 is possible, indicating "infinite".
func (x *PixbufAnimation) GetIter(StartTimeVar *glib.TimeVal) *PixbufAnimationIter {
	var cls *PixbufAnimationIter

	cret := xPixbufAnimationGetIter(x.GoPointer(), StartTimeVar)

	if cret == 0 {
		return nil
	}
	cls = &PixbufAnimationIter{}
	cls.Ptr = cret
	return cls
}

var xPixbufAnimationGetStaticImage func(uintptr) uintptr

// Retrieves a static image for the animation.
//
// If an animation is really just a plain image (has only one frame),
// this function returns that image.
//
// If the animation is an animation, this function returns a reasonable
// image to use as a static unanimated image, which might be the first
// frame, or something more sophisticated depending on the file format.
//
// If an animation hasn't loaded any frames yet, this function will
// return `NULL`.
func (x *PixbufAnimation) GetStaticImage() *Pixbuf {
	var cls *Pixbuf

	cret := xPixbufAnimationGetStaticImage(x.GoPointer())

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &Pixbuf{}
	cls.Ptr = cret
	return cls
}

var xPixbufAnimationGetWidth func(uintptr) int

// Queries the width of the bounding box of a pixbuf animation.
func (x *PixbufAnimation) GetWidth() int {

	cret := xPixbufAnimationGetWidth(x.GoPointer())
	return cret
}

var xPixbufAnimationIsStaticImage func(uintptr) bool

// Checks whether the animation is a static image.
//
// If you load a file with gdk_pixbuf_animation_new_from_file() and it
// turns out to be a plain, unanimated image, then this function will
// return `TRUE`. Use gdk_pixbuf_animation_get_static_image() to retrieve
// the image.
func (x *PixbufAnimation) IsStaticImage() bool {

	cret := xPixbufAnimationIsStaticImage(x.GoPointer())
	return cret
}

var xPixbufAnimationRef func(uintptr) uintptr

// Adds a reference to an animation.
func (x *PixbufAnimation) Ref() *PixbufAnimation {
	var cls *PixbufAnimation

	cret := xPixbufAnimationRef(x.GoPointer())

	if cret == 0 {
		return nil
	}
	cls = &PixbufAnimation{}
	cls.Ptr = cret
	return cls
}

var xPixbufAnimationUnref func(uintptr)

// Removes a reference from an animation.
func (x *PixbufAnimation) Unref() {

	xPixbufAnimationUnref(x.GoPointer())

}

func (c *PixbufAnimation) GoPointer() uintptr {
	return c.Ptr
}

func (c *PixbufAnimation) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// An opaque object representing an iterator which points to a
// certain position in an animation.
type PixbufAnimationIter struct {
	gobject.Object
}

func PixbufAnimationIterNewFromInternalPtr(ptr uintptr) *PixbufAnimationIter {
	cls := &PixbufAnimationIter{}
	cls.Ptr = ptr
	return cls
}

var xPixbufAnimationIterAdvance func(uintptr, *glib.TimeVal) bool

// Possibly advances an animation to a new frame.
//
// Chooses the frame based on the start time passed to
// gdk_pixbuf_animation_get_iter().
//
// @current_time would normally come from g_get_current_time(), and
// must be greater than or equal to the time passed to
// gdk_pixbuf_animation_get_iter(), and must increase or remain
// unchanged each time gdk_pixbuf_animation_iter_get_pixbuf() is
// called. That is, you can't go backward in time; animations only
// play forward.
//
// As a shortcut, pass `NULL` for the current time and g_get_current_time()
// will be invoked on your behalf. So you only need to explicitly pass
// @current_time if you're doing something odd like playing the animation
// at double speed.
//
// If this function returns `FALSE`, there's no need to update the animation
// display, assuming the display had been rendered prior to advancing;
// if `TRUE`, you need to call gdk_pixbuf_animation_iter_get_pixbuf()
// and update the display with the new pixbuf.
func (x *PixbufAnimationIter) Advance(CurrentTimeVar *glib.TimeVal) bool {

	cret := xPixbufAnimationIterAdvance(x.GoPointer(), CurrentTimeVar)
	return cret
}

var xPixbufAnimationIterGetDelayTime func(uintptr) int

// Gets the number of milliseconds the current pixbuf should be displayed,
// or -1 if the current pixbuf should be displayed forever.
//
// The `g_timeout_add()` function conveniently takes a timeout in milliseconds,
// so you can use a timeout to schedule the next update.
//
// Note that some formats, like GIF, might clamp the timeout values in the
// image file to avoid updates that are just too quick. The minimum timeout
// for GIF images is currently 20 milliseconds.
func (x *PixbufAnimationIter) GetDelayTime() int {

	cret := xPixbufAnimationIterGetDelayTime(x.GoPointer())
	return cret
}

var xPixbufAnimationIterGetPixbuf func(uintptr) uintptr

// Gets the current pixbuf which should be displayed.
//
// The pixbuf might not be the same size as the animation itself
// (gdk_pixbuf_animation_get_width(), gdk_pixbuf_animation_get_height()).
//
// This pixbuf should be displayed for gdk_pixbuf_animation_iter_get_delay_time()
// milliseconds.
//
// The caller of this function does not own a reference to the returned
// pixbuf; the returned pixbuf will become invalid when the iterator
// advances to the next frame, which may happen anytime you call
// gdk_pixbuf_animation_iter_advance().
//
// Copy the pixbuf to keep it (don't just add a reference), as it may get
// recycled as you advance the iterator.
func (x *PixbufAnimationIter) GetPixbuf() *Pixbuf {
	var cls *Pixbuf

	cret := xPixbufAnimationIterGetPixbuf(x.GoPointer())

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &Pixbuf{}
	cls.Ptr = cret
	return cls
}

var xPixbufAnimationIterOnCurrentlyLoadingFrame func(uintptr) bool

// Used to determine how to respond to the area_updated signal on
// #GdkPixbufLoader when loading an animation.
//
// The `::area_updated` signal is emitted for an area of the frame currently
// streaming in to the loader. So if you're on the currently loading frame,
// you will need to redraw the screen for the updated area.
func (x *PixbufAnimationIter) OnCurrentlyLoadingFrame() bool {

	cret := xPixbufAnimationIterOnCurrentlyLoadingFrame(x.GoPointer())
	return cret
}

func (c *PixbufAnimationIter) GoPointer() uintptr {
	return c.Ptr
}

func (c *PixbufAnimationIter) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GDKPIXBUF"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xNewFromFilePixbufAnimation, lib, "gdk_pixbuf_animation_new_from_file")
	core.PuregoSafeRegister(&xNewFromResourcePixbufAnimation, lib, "gdk_pixbuf_animation_new_from_resource")
	core.PuregoSafeRegister(&xNewFromStreamPixbufAnimation, lib, "gdk_pixbuf_animation_new_from_stream")
	core.PuregoSafeRegister(&xNewFromStreamFinishPixbufAnimation, lib, "gdk_pixbuf_animation_new_from_stream_finish")

	core.PuregoSafeRegister(&xPixbufAnimationGetHeight, lib, "gdk_pixbuf_animation_get_height")
	core.PuregoSafeRegister(&xPixbufAnimationGetIter, lib, "gdk_pixbuf_animation_get_iter")
	core.PuregoSafeRegister(&xPixbufAnimationGetStaticImage, lib, "gdk_pixbuf_animation_get_static_image")
	core.PuregoSafeRegister(&xPixbufAnimationGetWidth, lib, "gdk_pixbuf_animation_get_width")
	core.PuregoSafeRegister(&xPixbufAnimationIsStaticImage, lib, "gdk_pixbuf_animation_is_static_image")
	core.PuregoSafeRegister(&xPixbufAnimationRef, lib, "gdk_pixbuf_animation_ref")
	core.PuregoSafeRegister(&xPixbufAnimationUnref, lib, "gdk_pixbuf_animation_unref")

	core.PuregoSafeRegister(&xPixbufAnimationIterAdvance, lib, "gdk_pixbuf_animation_iter_advance")
	core.PuregoSafeRegister(&xPixbufAnimationIterGetDelayTime, lib, "gdk_pixbuf_animation_iter_get_delay_time")
	core.PuregoSafeRegister(&xPixbufAnimationIterGetPixbuf, lib, "gdk_pixbuf_animation_iter_get_pixbuf")
	core.PuregoSafeRegister(&xPixbufAnimationIterOnCurrentlyLoadingFrame, lib, "gdk_pixbuf_animation_iter_on_currently_loading_frame")

}
