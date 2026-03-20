package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/theme"
)

const (
	mascotChar   = "🐓"
	windowWidth  = 100
	windowHeight = 100
	fontSize     = 48.0
)

func main() {
	a := app.New()
	w := a.NewWindow("Iruka")

	text := canvas.NewText(mascotChar, color.Black)
	text.TextSize = fontSize
	text.Alignment = fyne.TextAlignCenter

	w.SetContent(container.NewCenter(text))
	w.Resize(fyne.NewSize(windowWidth, windowHeight))
	w.SetFixedSize(true)

	if desk, ok := a.(desktop.App); ok {
		m := fyne.NewMenu("Iruka",
			fyne.NewMenuItem("Show", func() { w.Show() }),
			fyne.NewMenuItemSeparator(),
			fyne.NewMenuItem("Quit", a.Quit),
		)
		desk.SetSystemTrayMenu(m)
		desk.SetSystemTrayIcon(theme.ComputerIcon())
		w.SetCloseIntercept(func() { w.Hide() })
	}

	w.Canvas().SetOnTypedKey(func(key *fyne.KeyEvent) {
		if key.Name == fyne.KeyEscape {
			a.Quit()
		}
	})

	w.ShowAndRun()
}
