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

type transparentTheme struct {
	fyne.Theme
}

func (t *transparentTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	if name == theme.ColorNameBackground {
		return color.Transparent
	}
	return t.Theme.Color(name, variant)
}

func main() {
	a := app.New()
	a.Settings().SetTheme(&transparentTheme{Theme: theme.DefaultTheme()})
	a.Lifecycle().SetOnEnteredForeground(makeWindowTransparent)

	var w fyne.Window
	if drv, ok := a.Driver().(desktop.Driver); ok {
		w = drv.CreateSplashWindow()
	} else {
		w = a.NewWindow("Iruka")
	}

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
