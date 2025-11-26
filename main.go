package main

import (
	"fmt"
	"os"

	"github.com/francescorocca/go-xlib/xlib"
)

func main() {
	display, err := xlib.OpenDisplay("")
	if display == nil {
		fmt.Fprintln(os.Stderr, "XOpenDisplay:", err)
		os.Exit(1)
	}

	/* Initialize the error handler */
	display.InitErrorHandler()

	fmt.Println("Display connected!")

	/* Make a simple window */
	screen_num := display.DefaultScreen()
	root := display.RootWindow(screen_num)
	black := display.BlackPixel(screen_num)
	white := display.WhitePixel(screen_num)

	win := display.CreateSimpleWindow(root, 100, 100, 300, 300, 0, white, black)

	/* Map window */
	display.SelectInput(win, xlib.ButtonPressMask)
	display.MapWindow(win)
	display.Flush()

	/* Event loop */
	for {
		ev := display.NextEvent()
		fmt.Println("Event type: ", ev.Type())
		if ev.Type() == xlib.ButtonPress {
			break
		}
	}

	display.Close()
}
