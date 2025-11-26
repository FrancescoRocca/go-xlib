package xlib

/*
#cgo LDFLAGS: -lX11
#include <X11/Xlib.h>
*/
import "C"
import "errors"

type Window struct {
	id C.Window
}

type WindowAttributes struct {
	ptr           C.XWindowAttributes
	X, Y          int
	Width, Height int
	BorderWidth   int
	Depth         int
	Root          Window
}

func NoneWindow() Window {
	return Window{id: C.None}
}

func (d *Display) RootWindow(screen int) Window {
	root := C.XRootWindow(d.ptr, C.int(screen))

	return Window{id: root}
}

func (d *Display) DefaultRootWindow() Window {
	win := C.XDefaultRootWindow(d.ptr)

	return Window{id: win}
}

func (d *Display) CreateSimpleWindow(parent Window, x int, y int, width uint, height uint, border_width uint, border uint64, background uint64) Window {
	win := C.XCreateSimpleWindow(
		d.ptr,
		parent.id,
		C.int(x),
		C.int(y),
		C.uint(width),
		C.uint(height),
		C.uint(border_width),
		C.ulong(border),
		C.ulong(background))

	return Window{id: win}
}

func (d *Display) GetWindowAttributes(window Window) (WindowAttributes, error) {
	var wa C.XWindowAttributes

	status := C.XGetWindowAttributes(d.ptr, window.id, &wa)
	if status == 0 {
		return WindowAttributes{}, errors.New("XGetWindowAttributes failed")
	}

	return WindowAttributes{
		ptr:         wa,
		X:           int(wa.x),
		Y:           int(wa.y),
		Width:       int(wa.width),
		Height:      int(wa.height),
		BorderWidth: int(wa.border_width),
		Depth:       int(wa.depth),
		Root:        Window{id: wa.root},
	}, nil
}

func (d *Display) MoveResizeWindow(window Window, x int, y int, width uint, height uint) {
	C.XMoveResizeWindow(
		d.ptr,
		window.id,
		C.int(x),
		C.int(y),
		C.uint(width),
		C.uint(height))
}

func (d *Display) RaiseWindow(window Window) error {
	ret := C.XRaiseWindow(d.ptr, window.id)
	if ret != 0 {
		return errors.New("XRaiseWindow failed")
	}

	return nil
}
