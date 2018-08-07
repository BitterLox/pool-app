package main

import (
	"github.com/murlokswarm/app"
	"math/rand"
	"time"

	"gonum.org/v1/plot/plotter"
)

var (
	uiloop chan struct{}
)

type Container struct {
	Data plotter.Values
	To   string
	View string
}

func (h *Container) OnMount() {

	uiloop = make(chan struct{})

	go func() {
		c := time.Tick(time.Millisecond * 750)
		for _ = range c {
			select {
			default:

				rand.Seed(time.Now().UTC().UnixNano())
				data := plotter.Values{rand.Float64(), rand.Float64(), rand.Float64(), rand.Float64(), rand.Float64()}

				app.Debug("Loop Sample: %+v", data)

				h.Data = data
				h.To = "Stats"
				h.View = "<Test to='{{.To}}'>"

				elem := app.ElemByCompo(h)

				app.Debug("ELEM: %#v %v", elem, elem.ID())

				app.Render(h)
				app.Debug("Tick")
			case <-uiloop:
				return
			}
		}
	}()

}

func (h *Container) Render() string {
	return `
<body>
	<div class="container">
		<div class="header">HEADER</div>
		<div class="menu">
			<a href="/Stats">Index</a>
			<a>Stats</a>
			<a>1</a>
			<a>Index</a>
			<a>Stats</a>
			<a>1</a>
		</div>
		<div class="content">
			{{raw .View}}
			<Chart data="{{json .Data}}">
		</div>
		<div class="footer">FOOTER</div>
	</div>
</body>

	`
}

// func (h *Container) ElemByCompo() {
//
//
// }

// <div class="content">
// 	<div class="chart">
// 		<Chart data="{{json .Data}}">
// 	</div>
// </div>

// <div class="content">
// 	<Content to="{{.To}}"></Content>
// </div>

func (h *Container) OnDismount() {
	uiloop <- struct{}{}
}
