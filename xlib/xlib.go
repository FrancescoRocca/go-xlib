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
 * Event type constants for X11 events.
 */
const (
	None          int = C.None
	True          int = C.True
	False         int = C.False
	GrabModeSync  int = C.GrabModeSync
	GrabModeAsync int = C.GrabModeAsync
)

/*
 * X11 event types for window management and input.
 */
const (
	MapRequest       int = int(C.MapRequest)
	ConfigureRequest int = int(C.ConfigureRequest)
	MotionNotify     int = int(C.MotionNotify)
)

/*
 * X11 event types for button and keyboard input.
 */
const (
	ButtonPress   int = int(C.ButtonPress)
	ButtonRelease int = int(C.ButtonRelease)
	KeyPress      int = int(C.KeyPress)
	KeyRelease    int = int(C.KeyRelease)
)

/*
 * X11 modifier masks and event masks for input handling.
 */
const (
	Mod1Mask                 uint   = uint(C.Mod1Mask)
	Mod2Mask                 uint   = uint(C.Mod2Mask)
	Mod4Mask                 uint   = uint(C.Mod4Mask)
	LockMask                 uint   = uint(C.LockMask)
	KeyPressMask             uint32 = uint32(C.KeyPressMask)
	KeyReleaseMask           uint32 = uint32(C.KeyReleaseMask)
	ButtonPressMask          uint32 = uint32(C.ButtonPressMask)
	ButtonReleaseMask        uint32 = uint32(C.ButtonReleaseMask)
	PointerMotionMask        uint32 = uint32(C.PointerMotionMask)
	SubstructureRedirectMask uint32 = uint32(C.SubstructureRedirectMask)
	SubstructureNotifyMask   uint32 = uint32(C.SubstructureNotifyMask)
)

/*
 * Display represents an X11 display connection.
 */
type Display struct {
	ptr *C.Display
}

/*
 * Cursor represents an X11 cursor.
 */
type Cursor struct {
	ptr C.Cursor
}

/*
 * KeySym represents an X11 key symbol.
 */
type KeySym struct {
	ptr C.KeySym
}

/*
 * OpenDisplay opens a connection to the X11 display.
 * If name is empty, the default display is used.
 */
func OpenDisplay(name string) (*Display, error) {
	var cName *C.char
	if name != "" {
		cName = C.CString(name)
		defer C.free(unsafe.Pointer(cName))
	}
	d := C.XOpenDisplay(cName)
	if d == nil {
		return nil, errors.New("cannot open display")
	}
	return &Display{ptr: d}, nil
}

/*
 * Close closes the display connection.
 */
func (d *Display) Close() {
	if d.ptr == nil {
		return
	}
	C.XCloseDisplay(d.ptr)
	d.ptr = nil
}

/*
 * Flush flushes all pending requests to the X11 server.
 */
func (d *Display) Flush() {
	C.XFlush(d.ptr)
}

/*
 * DefaultScreen returns the default screen number for the display.
 */
func (d *Display) DefaultScreen() int {
	screen := C.XDefaultScreen(d.ptr)
	return (int)(screen)
}

/*
 * NoneCursor returns a null cursor.
 */
func NoneCursor() Cursor {
	return Cursor{ptr: C.None}
}

/*
 * MapWindow maps a window to the display.
 */
func (d *Display) MapWindow(w Window) {
	C.XMapWindow(d.ptr, w.id)
}

/*
 * GrabKey grabs keyboard input for a specific key and window.
 */
func (d *Display) GrabKey(keycode uint, modifiers uint, grab_window Window, owner_events int, pointer_mode int, keyboard_mode int) {
	C.XGrabKey(
		d.ptr,
		C.int(keycode),
		C.uint(modifiers),
		grab_window.id,
		C.int(owner_events),
		C.int(pointer_mode),
		C.int(keyboard_mode))
}

/*
 * UngrabKey releases a previously grabbed key.
 */
func (d *Display) UngrabKey(keycode int, modifiers uint, grab_window Window) {
	C.XUngrabKey(
		d.ptr,
		C.int(keycode),
		C.uint(modifiers),
		grab_window.id)
}

/*
 * GrabButton grabs mouse button input for a window.
 */
func (d *Display) GrabButton(button uint32, modifiers uint, grab_window Window, owner_events int, event_mask uint32, pointer_mode int, keyboard_mode int, confine_to Window, cursor Cursor) {
	C.XGrabButton(
		d.ptr,
		C.uint(button),
		C.uint(modifiers),
		grab_window.id,
		C.int(owner_events),
		C.uint(event_mask),
		C.int(pointer_mode),
		C.int(keyboard_mode),
		confine_to.id,
		cursor.ptr)
}

/*
 * UngrabButton releases a previously grabbed mouse button.
 */
func (d *Display) UngrabButton(button uint, modifiers uint, grab_window Window) {
	C.XUngrabButton(
		d.ptr,
		C.uint(button),
		C.uint(modifiers),
		grab_window.id)
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
