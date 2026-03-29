//go:build darwin

package main

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Foundation -framework AppKit

#import <Foundation/Foundation.h>
#import <AppKit/AppKit.h>

void setMovableByBackground(uintptr_t window) {
    NSWindow *nswindow = (NSWindow*)window;
    [nswindow setMovableByWindowBackground:YES];
}
*/
import "C"

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver"
)

// makeWindowDraggable enables dragging the borderless mascot window by clicking
// anywhere on it, using the macOS native isMovableByWindowBackground property.
func makeWindowDraggable(w fyne.Window) {
	if nw, ok := w.(driver.NativeWindow); ok {
		nw.RunNative(func(ctx any) {
			if mac, ok := ctx.(driver.MacWindowContext); ok && mac.NSWindow != 0 {
				C.setMovableByBackground(C.uintptr_t(mac.NSWindow))
			}
		})
	}
}
