package main

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Cocoa
#include <Cocoa/Cocoa.h>

static void makeWindowTransparent() {
	for (NSWindow *win in [NSApplication sharedApplication].windows) {
		[win setOpaque:NO];
		[win setBackgroundColor:[NSColor clearColor]];
	}
}
*/
import "C"

func makeWindowTransparent() {
	C.makeWindowTransparent()
}
