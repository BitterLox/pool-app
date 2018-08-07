// +build darwin,amd64

package main

import (
	// "github.com/bitterlox/pool-app/components/markup"
	"github.com/murlokswarm/app"
	"github.com/murlokswarm/app/drivers/mac"
	"os"
)

func init() {
	app.Loggers = []app.Logger{
		app.NewLogger(os.Stdout, os.Stderr, true, true),
	}
	app.Import(&Chart{})
	app.Import(&Container{})
	app.Import(&Stats{})
	app.Import(&Test{})

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
		DefaultURL:      "/Container",
	})
}
