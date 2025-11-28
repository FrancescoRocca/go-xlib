package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/francescorocca/go-xlib"
)

func spawnLauncher() error {
	if err := exec.Command("dmenu_run").Start(); err == nil {
		return nil
	} else {
		fmt.Println(err)
	}

	if err := exec.Command("rofi", "-show", "run").Start(); err == nil {
		return nil
	} else {
		fmt.Println(err)
	}

	if err := exec.Command("xterm").Start(); err == nil {
		return nil
	} else {
		fmt.Println(err)
		return err
	}
}

func grabKeyWithMasks(dpy *xlib.Display, keycode uint, mods uint) {
	masks := []uint{0, xlib.Mod2Mask, xlib.LockMask, xlib.Mod2Mask | xlib.LockMask}
	for _, m := range masks {
		dpy.GrabKey(keycode, mods|m, dpy.DefaultRootWindow(), xlib.True, xlib.GrabModeAsync, xlib.GrabModeAsync)
	}
}

func main() {
	dpy, err := xlib.OpenDisplay("")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	defer dpy.Close()

	dpy.InitErrorHandler()

	root := dpy.DefaultRootWindow()
	dpy.SelectInput(root, xlib.SubstructureNotifyMask|xlib.SubstructureRedirectMask)

	ret, err := dpy.KeysymToKeycode(xlib.StringToKeysym("Return"))
	if err != nil {
		fmt.Fprintln(os.Stderr, "KeysymToKeycode(Return):", err)
	} else {
		grabKeyWithMasks(dpy, ret, xlib.Mod4Mask)
	}

	dpy.GrabButton(
		1,
		xlib.Mod1Mask,
		root,
		xlib.True,
		xlib.ButtonPressMask|xlib.ButtonReleaseMask|xlib.PointerMotionMask,
		xlib.GrabModeAsync,
		xlib.GrabModeAsync,
		xlib.NoneWindow(),
		xlib.NoneCursor(),
	)

	dpy.GrabButton(
		3,
		xlib.Mod1Mask,
		root,
		xlib.True,
		xlib.ButtonPressMask|xlib.ButtonReleaseMask|xlib.PointerMotionMask,
		xlib.GrabModeAsync,
		xlib.GrabModeAsync,
		xlib.NoneWindow(),
		xlib.NoneCursor(),
	)

	var start xlib.ButtonEvent
	var attr xlib.WindowAttributes
	start.Subwindow = xlib.NoneWindow()

	for {
		ev := dpy.NextEvent()

		switch ev.Type() {
		case xlib.KeyPress:
			ke := ev.AsKeyEvent()
			fmt.Printf("KeyPress keycode=%d state=0x%x\n", ke.Keycode, ke.State)
			fmt.Printf("\t(ret=%d, Mod4Mask=0x%x)\n", ret, xlib.Mod4Mask)

			if ke.Keycode == ret && (ke.State&xlib.Mod4Mask) != 0 {
				fmt.Println("WIN+Enter detected -> launcher")
				if err := spawnLauncher(); err != nil {
					fmt.Fprintln(os.Stderr, "Spawn error:", err)
				}
			}

			if ke.Subwindow != xlib.NoneWindow() {
				dpy.RaiseWindow(ke.Subwindow)
			}

		case xlib.ButtonPress:
			be := ev.AsButtonEvent()
			if be.Subwindow != xlib.NoneWindow() {
				attr, err = dpy.GetWindowAttributes(be.Subwindow)
				if err != nil {
					fmt.Println(err)
				}
				start = *be
			}

		case xlib.MotionNotify:
			if start.Subwindow != xlib.NoneWindow() {
				me := ev.AsButtonEvent()
				xdiff := me.XRoot - start.XRoot
				ydiff := me.YRoot - start.YRoot
				x := attr.X
				y := attr.Y
				if start.Button == 1 {
					x += xdiff
					y += ydiff
				}
				dpy.MoveResizeWindow(start.Subwindow, x, y, 0, 0)
			}

		case xlib.ButtonRelease:
			start.Subwindow = xlib.NoneWindow()

		case xlib.MapRequest:
			mr := ev.AsMapRequestEvent()
			dpy.MapWindow(mr.Window)

		case xlib.ConfigureRequest:
			cr := ev.AsConfigureRequestEvent()

			var mask uint
			var changes xlib.WindowChanges

			if (cr.ValueMask & xlib.CWX) != 0 {
				mask |= xlib.CWX
				changes.X = cr.X
			}
			if (cr.ValueMask & xlib.CWY) != 0 {
				mask |= xlib.CWY
				changes.Y = cr.Y
			}
			if (cr.ValueMask & xlib.CWWidth) != 0 {
				mask |= xlib.CWWidth
				changes.Width = cr.Width
			}
			if (cr.ValueMask & xlib.CWHeight) != 0 {
				mask |= xlib.CWHeight
				changes.Height = cr.Height
			}
			if (cr.ValueMask & xlib.CWBorderWidth) != 0 {
				mask |= xlib.CWBorderWidth
				changes.BorderWidth = cr.BorderWidth
			}
			if (cr.ValueMask & xlib.CWSibling) != 0 {
				mask |= xlib.CWSibling
				changes.Sibling = cr.Above
			}
			if (cr.ValueMask & xlib.CWStackMode) != 0 {
				mask |= xlib.CWStackMode
				changes.StackMode = cr.Detail
			}

			dpy.ConfigureWindow(cr.Window, mask, changes)
		}
	}
}
