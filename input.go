package xlib

/*
#cgo LDFLAGS: -lX11
#include <X11/Xlib.h>
#include <stdlib.h>
*/
import "C"
import (
	"errors"
	"unsafe"
)

/*
 * NoneCursor returns a null cursor.
 */
func NoneCursor() Cursor {
	return Cursor{ptr: C.None}
}

/*
 * GrabKey grabs keyboard input for a specific key and window.
 */
func (d *Display) GrabKey(keycode uint, modifiers uint, grabWindow Window, ownerEvents int, pointerMode int, keyboardMode int) error {
	C.XGrabKey(
		d.ptr,
		C.int(keycode),
		C.uint(modifiers),
		grabWindow.id,
		C.int(ownerEvents),
		C.int(pointerMode),
		C.int(keyboardMode))

	return d.SyncAndCheckError()
}

/*
 * UngrabKey releases a previously grabbed key.
 */
func (d *Display) UngrabKey(keycode int, modifiers uint, grabWindow Window) error {
	C.XUngrabKey(
		d.ptr,
		C.int(keycode),
		C.uint(modifiers),
		grabWindow.id)

	return d.SyncAndCheckError()
}

/*
 * GrabButton grabs mouse button input for a window.
 */
func (d *Display) GrabButton(button uint32, modifiers uint, grabWindow Window, ownerEvents int, eventMask uint32, pointerMode int, keyboardMode int, confineTo Window, cursor Cursor) error {
	C.XGrabButton(
		d.ptr,
		C.uint(button),
		C.uint(modifiers),
		grabWindow.id,
		C.int(ownerEvents),
		C.uint(eventMask),
		C.int(pointerMode),
		C.int(keyboardMode),
		confineTo.id,
		cursor.ptr)

	return d.SyncAndCheckError()
}

/*
 * UngrabButton releases a previously grabbed mouse button.
 */
func (d *Display) UngrabButton(button uint, modifiers uint, grabWindow Window) {
	C.XUngrabButton(
		d.ptr,
		C.uint(button),
		C.uint(modifiers),
		grabWindow.id)
}

/*
 * Grab the pointer
 */
func (d *Display) GrabPointer(grabWindow Window, ownerEvents int, eventMask uint32, pointerMode int, keyboardMode int, confineTo Window, cursor Cursor, time uint64) {
	C.XGrabPointer(
		d.ptr,
		grabWindow.id,
		C.int(ownerEvents),
		C.uint(eventMask),
		C.int(pointerMode),
		C.int(keyboardMode),
		confineTo.id,
		cursor.ptr,
		C.Time(time))
}

/*
 * Ungrab the pointer
 */
func (d *Display) UngrabPointer(time uint64) {
	C.XUngrabPointer(d.ptr, C.Time(time))
}

/*
 * StringToKeysym converts a string to an X11 key symbol.
 */
func StringToKeysym(str string) KeySym {
	cstr := C.CString(str)
	defer C.free(unsafe.Pointer(cstr))
	return KeySym{ptr: C.XStringToKeysym(cstr)}
}

/*
 * KeysymToKeycode converts a key symbol to a key code.
 * Returns an error if the conversion fails.
 */
func (d *Display) KeysymToKeycode(keysym KeySym) (uint, error) {
	ret := C.XKeysymToKeycode(d.ptr, keysym.ptr)
	if ret == 0 {
		return 0, errors.New("XKeysymToKeycode failed")
	}
	return uint(ret), nil
}
