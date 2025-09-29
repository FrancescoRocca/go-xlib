package xlib

/*
#cgo LDFLAGS: -lX11
#include <X11/Xlib.h>
#include <X11/keysym.h>
#include <stdlib.h>
#include <unistd.h>
*/
import "C"
import (
	"errors"
	"unsafe"
)

type Display struct {
	ptr *C.Display
}

func XOpenDisplay(name string) (*Display, error) {
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
