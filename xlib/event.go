package xlib

/*
#cgo LDFLAGS: -lX11
#include <X11/Xlib.h>
#include "event.h"
*/
import "C"

type Event struct {
	ptr C.XEvent
}

const (
	ButtonPress = int(C.ButtonPress)
	KeyPress    = int(C.KeyPress)
)

const (
	ButtonPressMask = uint32(C.ButtonPressMask)
)

func (d *Display) SelectInput(window Window, event_mask uint32) {
	C.XSelectInput(d.ptr, window.id, C.long(event_mask))
}

func (d *Display) Pending() int {
	return int(C.XPending(d.ptr))
}

func (d *Display) NextEvent() Event {
	var ev Event
	C.XNextEvent(d.ptr, &ev.ptr)

	return ev
}

func (e *Event) Type() int {
	return int(C.GetEventType(&e.ptr))
}
