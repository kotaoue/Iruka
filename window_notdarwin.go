//go:build !darwin

package main

import "fyne.io/fyne/v2"

// makeWindowDraggable is a no-op on non-Darwin platforms.
func makeWindowDraggable(_ fyne.Window) {}
