package xlib

/*
#cgo LDFLAGS: -lX11
#include <X11/Xlib.h>
*/
import "C"

/*
 * BlackPixel returns the black pixel value for the specified screen.
 */
func (d *Display) BlackPixel(screen int) uint64 {
	pixel := C.XBlackPixel(d.ptr, C.int(screen))
	return (uint64)(pixel)
}

/*
 * WhitePixel returns the white pixel value for the specified screen.
 */
func (d *Display) WhitePixel(screen int) uint64 {
	pixel := C.XWhitePixel(d.ptr, C.int(screen))
	return (uint64)(pixel)
}
