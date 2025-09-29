package main

import (
	"fmt"
	"go-xlib/xlib"
	"os"
)

func main() {
	display, err := xlib.XOpenDisplay("")
	if display == nil {
		fmt.Fprintln(os.Stderr, "XOpenDisplay:", err)
		os.Exit(1)
	}

	fmt.Println("Display connected!")

	display.Close()
}
