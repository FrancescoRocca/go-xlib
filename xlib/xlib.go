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

const (
	None                     int    = C.None
	True                     int    = C.True
	False                    int    = C.False
	GrabModeSync             int    = C.GrabModeSync
	GrabModeAsync            int    = C.GrabModeAsync
	Mod1Mask                 uint32 = uint32(C.Mod1Mask)
	ButtonPress              int    = int(C.ButtonPress)
	ButtonRelease            int    = int(C.ButtonRelease)
	KeyPress                 int    = int(C.KeyPress)
	KeyRelease               int    = int(C.KeyRelease)
	KeyPressMask             uint32 = uint32(C.KeyPressMask)
	KeyReleaseMask           uint32 = uint32(C.KeyReleaseMask)
	ButtonPressMask          uint32 = uint32(C.ButtonPressMask)
	ButtonReleaseMask        uint32 = uint32(C.ButtonReleaseMask)
	PointerMotionMask        uint32 = uint32(C.PointerMotionMask)
	MotionNotify             int    = int(C.MotionNotify)
	SubstructureRedirectMask uint32 = uint32(C.SubstructureRedirectMask)
	SubstructureNotifyMask   uint32 = uint32(C.SubstructureNotifyMask)
)

type Display struct {
	ptr *C.Display
}

type Cursor struct {
	ptr C.Cursor
}

type KeySym struct {
	ptr C.KeySym
}

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

func (d *Display) Close() {
	if d.ptr == nil {
		return
	}

	C.XCloseDisplay(d.ptr)
	d.ptr = nil
}

func (d *Display) Flush() {
	C.XFlush(d.ptr)
}

func (d *Display) DefaultScreen() int {
	screen := C.XDefaultScreen(d.ptr)

	return (int)(screen)
}

func NoneCursor() Cursor {
	return Cursor{ptr: C.None}
}

func (d *Display) MapWindow(w Window) {
	C.XMapWindow(d.ptr, w.id)
}

func (d *Display) GrabKey(keycode int, modifiers uint32, grab_window Window, owner_events int, pointer_mode int, keyboard_mode int) {
	C.XGrabKey(
		d.ptr,
		C.int(keycode),
		C.uint(modifiers),
		grab_window.id,
		C.int(owner_events),
		C.int(pointer_mode),
		C.int(keyboard_mode))
}

func (d *Display) UngrabKey(keycode int, modifiers uint, grab_window Window) {
	C.XUngrabKey(
		d.ptr,
		C.int(keycode),
		C.uint(modifiers),
		grab_window.id)
}

func (d *Display) GrabButton(button uint32, modifiers uint32, grab_window Window, owner_events int, event_mask uint32, pointer_mode int, keyboard_mode int, confine_to Window, cursor Cursor) {
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

func (d *Display) UngrabButton(button uint, modifiers uint, grab_window Window) {
	C.XUngrabButton(
		d.ptr,
		C.uint(button),
		C.uint(modifiers),
		grab_window.id)
}

func StringToKeysym(str string) KeySym {
	cstr := C.CString(str)
	defer C.free(unsafe.Pointer(cstr))
	return KeySym{ptr: C.XStringToKeysym(cstr)}
}

func (d *Display) KeysymToKeycode(keysym KeySym) (int, error) {
	ret := C.XKeysymToKeycode(d.ptr, keysym.ptr)
	if ret == 0 {
		return 0, errors.New("XKeysymToKeycode failed")
	}
	return int(ret), nil
}
