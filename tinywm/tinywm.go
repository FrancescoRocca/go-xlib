package main

import (
	"fmt"
	"go-xlib/xlib"
	"os"
)

func main() {
	dpy, err := xlib.OpenDisplay("")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	root := dpy.DefaultRootWindow()
	dpy.SelectInput(root, xlib.SubstructureNotifyMask|xlib.SubstructureRedirectMask)

	keysym, err := dpy.KeysymToKeycode(xlib.StringToKeysym("F1"))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		dpy.Close()
		return
	}

	dpy.GrabKey(keysym, xlib.Mod1Mask, root, xlib.True, xlib.GrabModeAsync, xlib.GrabModeAsync)
	dpy.GrabButton(1, xlib.Mod1Mask, root, xlib.True, xlib.ButtonPressMask|xlib.ButtonReleaseMask|xlib.PointerMotionMask, xlib.GrabModeAsync, xlib.GrabModeAsync, xlib.NoneWindow(), xlib.NoneCursor())
	dpy.GrabButton(3, xlib.Mod1Mask, root, xlib.True, xlib.ButtonPressMask|xlib.ButtonReleaseMask|xlib.PointerMotionMask, xlib.GrabModeAsync, xlib.GrabModeAsync, xlib.NoneWindow(), xlib.NoneCursor())

	var start xlib.ButtonEvent
	var attr xlib.WindowAttributes

	start.Subwindow = xlib.NoneWindow()

	for {
		ev := dpy.NextEvent()

		fmt.Println("Event Type:", ev.Type())

		if ev.Type() == xlib.KeyPress && ev.AsKeyEvent().Subwindow != xlib.NoneWindow() {
			dpy.RaiseWindow(ev.AsKeyEvent().Subwindow)
		} else if ev.Type() == xlib.ButtonPress && ev.AsButtonEvent().Subwindow != xlib.NoneWindow() {
			attr, err = dpy.GetWindowAttributes(ev.AsButtonEvent().Subwindow)
			if err != nil {
				fmt.Println(err)
			}
			start = *ev.AsButtonEvent()
		} else if ev.Type() == xlib.MotionNotify && start.Subwindow != xlib.NoneWindow() {
			xdiff := ev.AsButtonEvent().XRoot - start.XRoot
			ydiff := ev.AsButtonEvent().YRoot - start.YRoot
			var x int
			var y int
			if start.Button == 1 {
				x = attr.X + xdiff
				y = attr.Y + ydiff
			} else {
				x = attr.X
				y = attr.Y
			}

			dpy.MoveResizeWindow(start.Subwindow, x, y, 0, 0)
		} else if ev.Type() == xlib.ButtonRelease {
			start.Subwindow = xlib.NoneWindow()
		}
	}
}
