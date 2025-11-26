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

type KeyEvent struct {
	Type         int
	Serial       uint64
	SendEvent    bool
	Display      *Display
	Window       Window
	Root         Window
	Subwindow    Window
	Time         uint64
	X, Y         int
	XRoot, YRoot int
	State        uint
	Keycode      uint
	SameScreen   bool
}

type ButtonEvent struct {
	Type         int
	Serial       uint64
	SendEvent    bool
	Display      *Display
	Window       Window
	Root         Window
	Subwindow    Window
	Time         uint64
	X, Y         int
	XRoot, YRoot int
	State        uint
	Button       uint
	SameScreen   int
}

type MapRequestEvent struct {
	Type      int
	Serial    uint64
	SendEvent bool
	Display   *Display
	Parent    Window
	Window    Window
}

type ConfigureRequestEvent struct {
	Type          int
	Serial        uint64
	SendEvent     bool
	Display       *Display
	Parent        Window
	Window        Window
	X, Y          int
	Width, Height int
	BorderWidth   int
	Above         Window
	Detail        int
	ValueMask     uint
}

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

func (e *Event) KeyEvent() *C.XKeyEvent {
	return C.GetKeyEvent(&e.ptr)
}

func (e *Event) AsKeyEvent() *KeyEvent {
	ke := e.KeyEvent()

	if ke == nil {
		return nil
	}

	return &KeyEvent{
		Type:       int(ke._type),
		Serial:     uint64(ke.serial),
		SendEvent:  ke.send_event != 0,
		Display:    &Display{ptr: ke.display},
		Window:     Window{id: ke.window},
		Root:       Window{id: ke.root},
		Subwindow:  Window{id: ke.subwindow},
		Time:       uint64(ke.time),
		X:          int(ke.x),
		Y:          int(ke.y),
		XRoot:      int(ke.x_root),
		YRoot:      int(ke.y_root),
		State:      uint(ke.state),
		Keycode:    uint(ke.keycode),
		SameScreen: ke.same_screen != 0,
	}
}

func (e *Event) ButtonEvent() *C.XButtonEvent {
	return C.GetButtonEvent(&e.ptr)
}

func (e *Event) AsButtonEvent() *ButtonEvent {
	be := e.ButtonEvent()

	if be == nil {
		return nil
	}

	return &ButtonEvent{
		Type:       int(be._type),
		Serial:     uint64(be.serial),
		SendEvent:  be.send_event != 0,
		Display:    &Display{ptr: be.display},
		Window:     Window{id: be.window},
		Root:       Window{id: be.root},
		Subwindow:  Window{id: be.subwindow},
		Time:       uint64(be.time),
		X:          int(be.x),
		Y:          int(be.y),
		XRoot:      int(be.x_root),
		YRoot:      int(be.y_root),
		State:      uint(be.state),
		Button:     uint(be.button),
		SameScreen: int(be.same_screen),
	}
}

func (e *Event) MapRequestEvent() *C.XMapRequestEvent {
	return C.GetMapRequestEvent(&e.ptr)
}

func (e *Event) AsMapRequestEvent() *MapRequestEvent {
	mre := e.MapRequestEvent()

	if mre == nil {
		return nil
	}

	return &MapRequestEvent{
		Type:      int(mre._type),
		Serial:    uint64(mre.serial),
		SendEvent: mre.send_event != 0,
		Display:   &Display{ptr: mre.display},
		Parent:    Window{id: mre.parent},
		Window:    Window{id: mre.window},
	}
}

func (e *Event) ConfigureRequestEvent() *C.XConfigureRequestEvent {
	return C.GetConfigureRequestEvent(&e.ptr)
}

func (e *Event) AsConfigureRequestEvent() *ConfigureRequestEvent {
	cre := e.ConfigureRequestEvent()

	if cre == nil {
		return nil
	}

	return &ConfigureRequestEvent{
		Type:        int(cre._type),
		Serial:      uint64(cre.serial),
		SendEvent:   cre.send_event != 0,
		Display:     &Display{ptr: cre.display},
		Parent:      Window{id: cre.parent},
		Window:      Window{id: cre.window},
		X:           int(cre.x),
		Y:           int(cre.y),
		Width:       int(cre.width),
		Height:      int(cre.height),
		BorderWidth: int(cre.border_width),
		Above:       Window{id: cre.above},
		Detail:      int(cre.detail),
		ValueMask:   uint(cre.value_mask),
	}
}
