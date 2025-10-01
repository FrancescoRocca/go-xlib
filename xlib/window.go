package xlib

/*
#cgo LDFLAGS: -lX11
#include <X11/Xlib.h>
*/
import "C"

type Window struct {
	id C.Window
}

func (d *Display) RootWindow(screen int) Window {
	root := C.XRootWindow(d.ptr, C.int(screen))

	return Window{id: root}
}

func (d *Display) CreateSimpleWindow(parent Window, x int, y int, width uint, height uint, border_width uint, border uint64, background uint64) Window {
	win := C.XCreateSimpleWindow(d.ptr, parent.id, C.int(x), C.int(y), C.uint(width), C.uint(height), C.uint(border_width), C.ulong(border), C.ulong(background))

	return Window{id: win}
}
