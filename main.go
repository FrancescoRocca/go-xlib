package main

import (
	"fmt"
	"go-xlib/xlib"
	"os"
	"time"
)

func main() {
	display, err := xlib.OpenDisplay("")
	if display == nil {
		fmt.Fprintln(os.Stderr, "XOpenDisplay:", err)
		os.Exit(1)
	}

	fmt.Println("Display connected!")

	// Make a simple window
	screen_num := display.DefaultScreen()
	root := display.RootWindow(screen_num)
	black := display.BlackPixel(screen_num)
	white := display.WhitePixel(screen_num)

	win := display.CreateSimpleWindow(root, 100, 100, 300, 300, 0, white, black)

	// Map window
	display.MapWindow(win)
	display.Flush()

	// Sleep 10 seconds
	time.Sleep(10000000000)

	display.Close()
}
