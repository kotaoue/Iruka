package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

const (
	mascotChar   = "🐓"
	windowWidth  = 100
	windowHeight = 100
	fontSize     = 48.0
)

type mascotWidget struct {
	widget.BaseWidget
	text *canvas.Text
	app  fyne.App
	win  fyne.Window
}

func newMascotWidget(a fyne.App, w fyne.Window) *mascotWidget {
	m := &mascotWidget{app: a, win: w}
	m.text = canvas.NewText(mascotChar, color.Black)
	m.text.TextSize = fontSize
	m.text.Alignment = fyne.TextAlignCenter
	m.ExtendBaseWidget(m)
	return m
}

func (m *mascotWidget) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(m.text)
}

func (m *mascotWidget) MouseDown(ev *desktop.MouseEvent) {
	if ev.Button == desktop.MouseButtonSecondary {
		menu := fyne.NewMenu("Iruka",
			fyne.NewMenuItem("Quit", m.app.Quit),
		)
		widget.ShowPopUpMenuAtPosition(menu, m.win.Canvas(), ev.AbsolutePosition)
	}
}

// MouseUp is required to satisfy the desktop.Mouseable interface.
func (m *mascotWidget) MouseUp(_ *desktop.MouseEvent) {}

func main() {
	a := app.New()

	var w fyne.Window
	if drv, ok := a.Driver().(desktop.Driver); ok {
		w = drv.CreateSplashWindow()
	} else {
		w = a.NewWindow("Iruka")
	}

	mascot := newMascotWidget(a, w)
	w.SetContent(container.NewCenter(mascot))
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

	w.Show()
	makeWindowDraggable(w)
	a.Run()
}
