package xlib

/*
#cgo LDFLAGS: -lX11
#include <X11/Xlib.h>
*/
import "C"
import "errors"

/*
 * Window configuration value mask constants.
 */
const (
	CWX           uint = uint(C.CWX)
	CWY           uint = uint(C.CWY)
	CWWidth       uint = uint(C.CWWidth)
	CWHeight      uint = uint(C.CWHeight)
	CWBorderWidth uint = uint(C.CWBorderWidth)
	CWSibling     uint = uint(C.CWSibling)
	CWStackMode   uint = uint(C.CWStackMode)
)

/*
 * Window represents an X11 window.
 */
type Window struct {
	id C.Window
}

/*
 * WindowAttributes contains attributes of an X11 window.
 */
type WindowAttributes struct {
	ptr           C.XWindowAttributes
	X, Y          int
	Width, Height int
	BorderWidth   int
	Depth         int
	Root          Window
}

/*
 * WindowChanges specifies changes to apply to a window's configuration.
 */
type WindowChanges struct {
	X, Y          int
	Width, Height int
	BorderWidth   int
	Sibling       Window
	StackMode     int
}

/*
 * NoneWindow returns a null window.
 */
func NoneWindow() Window {
	return Window{id: C.None}
}

/*
 * RootWindow returns the root window of the specified screen.
 */
func (d *Display) RootWindow(screen int) Window {
	root := C.XRootWindow(d.ptr, C.int(screen))
	return Window{id: root}
}

/*
 * DefaultRootWindow returns the default root window for the display.
 */
func (d *Display) DefaultRootWindow() Window {
	win := C.XDefaultRootWindow(d.ptr)
	return Window{id: win}
}

/*
 * ConfigureWindow changes the configuration of a window.
 * The value_mask parameter determines which fields in changes are used.
 */
func (d *Display) ConfigureWindow(window Window, value_mask uint, changes WindowChanges) error {
	cChanges := C.XWindowChanges{
		x:            C.int(changes.X),
		y:            C.int(changes.Y),
		width:        C.int(changes.Width),
		height:       C.int(changes.Height),
		border_width: C.int(changes.BorderWidth),
		sibling:      C.Window(changes.Sibling.id),
		stack_mode:   C.int(changes.StackMode),
	}
	C.XConfigureWindow(d.ptr, C.Window(window.id), C.uint(value_mask), &cChanges)

	return d.SyncAndCheckError()
}

/*
 * CreateSimpleWindow creates a simple window with the specified parameters.
 */
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

/*
 * GetWindowAttributes retrieves the attributes of a window.
 * Returns an error if the operation fails.
 */
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

/*
 * MoveResizeWindow moves and resizes a window.
 */
func (d *Display) MoveResizeWindow(window Window, x int, y int, width uint, height uint) {
	C.XMoveResizeWindow(
		d.ptr,
		window.id,
		C.int(x),
		C.int(y),
		C.uint(width),
		C.uint(height))
}

/*
 * RaiseWindow raises a window to the top of the stacking order.
 * Returns an error if the operation fails.
 */
func (d *Display) RaiseWindow(window Window) error {
	C.XRaiseWindow(d.ptr, window.id)

	return d.SyncAndCheckError()
}
