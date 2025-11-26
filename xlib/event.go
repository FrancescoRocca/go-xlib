package xlib

/*
#cgo LDFLAGS: -lX11
#include <X11/Xlib.h>
#include "event.h"
*/
import "C"

/*
 * Event represents a generic X11 event.
 */
type Event struct {
	ptr C.XEvent
}

/*
 * KeyEvent represents an X11 keyboard event (KeyPress or KeyRelease).
 */
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

/*
 * ButtonEvent represents an X11 mouse button event (ButtonPress or ButtonRelease).
 */
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

/*
 * MapRequestEvent represents a MapRequest event from a window manager.
 */
type MapRequestEvent struct {
	Type      int
	Serial    uint64
	SendEvent bool
	Display   *Display
	Parent    Window
	Window    Window
}

/*
 * ConfigureRequestEvent represents a ConfigureRequest event for window geometry changes.
 */
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

/*
 * SelectInput selects which events are reported for a window.
 */
func (d *Display) SelectInput(window Window, event_mask uint32) {
	C.XSelectInput(d.ptr, window.id, C.long(event_mask))
}

/*
 * Pending returns the number of events waiting in the event queue.
 */
func (d *Display) Pending() int {
	return int(C.XPending(d.ptr))
}

/*
 * NextEvent retrieves the next event from the queue.
 */
func (d *Display) NextEvent() Event {
	var ev Event
	C.XNextEvent(d.ptr, &ev.ptr)
	return ev
}

/*
 * Type returns the type of the event.
 */
func (e *Event) Type() int {
	return int(C.GetEventType(&e.ptr))
}

/*
 * KeyEvent returns the underlying C key event structure, or nil if not a key event.
 */
func (e *Event) KeyEvent() *C.XKeyEvent {
	return C.GetKeyEvent(&e.ptr)
}

/*
 * AsKeyEvent converts the event to a Go KeyEvent struct.
 * Returns nil if the event is not a keyboard event.
 */
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

/*
 * ButtonEvent returns the underlying C button event structure, or nil if not a button event.
 */
func (e *Event) ButtonEvent() *C.XButtonEvent {
	return C.GetButtonEvent(&e.ptr)
}

/*
 * AsButtonEvent converts the event to a Go ButtonEvent struct.
 * Returns nil if the event is not a button event.
 */
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

/*
 * MapRequestEvent returns the underlying C map request event structure, or nil if not a map request event.
 */
func (e *Event) MapRequestEvent() *C.XMapRequestEvent {
	return C.GetMapRequestEvent(&e.ptr)
}

/*
 * AsMapRequestEvent converts the event to a Go MapRequestEvent struct.
 * Returns nil if the event is not a map request event.
 */
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

/*
 * ConfigureRequestEvent returns the underlying C configure request event structure, or nil if not a configure request event.
 */
func (e *Event) ConfigureRequestEvent() *C.XConfigureRequestEvent {
	return C.GetConfigureRequestEvent(&e.ptr)
}

/*
 * AsConfigureRequestEvent converts the event to a Go ConfigureRequestEvent struct.
 * Returns nil if the event is not a configure request event.
 */
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
