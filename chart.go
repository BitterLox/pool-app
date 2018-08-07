package main

import (
	"bytes"
	// "github.com/murlokswarm/app"

	"image/color"
	"strings"

	"github.com/bitterlox/plotters/piechart"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
	"gonum.org/v1/plot/vg/vgsvg"
)

type Chart struct {
	Data plotter.Values
}

func (c *Chart) Render() string {
	var data plotter.Values

	// elem := app.ElemByCompo(c)

	// app.Debug("ELEM: %#v %v", elem, elem.ID())

	if len(c.Data) > 0 {
		data = c.Data
	} else {
		data = plotter.Values{1, 2, 3, 4, 5}
	}

	buffer := bytes.NewBuffer([]byte{})

	p, err := plot.New()
	if err != nil {
		panic(err)
	}
	p.HideAxes()
	p.BackgroundColor = color.RGBA{0, 0, 0, 0}
	pie, err := piechart.NewPieChart(data)
	if err != nil {
		panic(err)
	}
	pie.Color = color.RGBA{234, 90, 0, 255}
	p.Add(pie)

	// Create a Canvas for writing SVG images.
	canvas := vgsvg.New(3*vg.Inch, 3*vg.Inch)

	// Draw to the Canvas.
	p.Draw(draw.New(canvas))

	// Write the Canvas to a io.Writer (in this case, os.Stdout).
	if _, err := canvas.WriteTo(buffer); err != nil {
		panic(err)
	}

	// app.Debug("Test Buffer Render() %v", buffer)

	return strings.Split(buffer.String(), "-->\n")[1]
}
