// +build darwin,amd64

package main

import (
	"github.com/bitterlox/pool-app/markup"
	"github.com/murlokswarm/app"
	"github.com/murlokswarm/app/drivers/mac"
	"os"
)

func init() {
	app.Loggers = []app.Logger{
		app.NewLogger(os.Stdout, os.Stderr, true, true),
	}
	app.Import(&markup.Chart{})
	app.Import(&markup.Container{})

}

func main() {

	app.Run(&mac.Driver{
		OnRun: func() {
			newWindow()
		},

		OnReopen: func(hasVisibleWindow bool) {
			if !hasVisibleWindow {
				newWindow()
			}
		},
	}, app.Logs())

}

func newWindow() {
	app.NewWindow(app.WindowConfig{
		Title:           "App",
		TitlebarHidden:  true,
		Width:           1280,
		Height:          768,
		BackgroundColor: "#2576f0",
		DefaultURL:      "/markup.Container",
	})
}
