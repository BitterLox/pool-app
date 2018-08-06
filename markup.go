package main

import (
	"bytes"
	"encoding/json"
	"github.com/murlokswarm/app"
	"github.com/wcharczuk/go-chart"
	"time"
)

var (
	uiloop chan struct{}
)

type Container struct {
	Data []byte
}

func (h *Container) OnMount() {

	app.Log("%s", "OnMount() main")
	app.Debug("%v", h.Data)

	sample := []chart.Value{
		{Value: 5, Label: "one"},
		{Value: 5, Label: "two"},
		{Value: 4, Label: "three"},
		{Value: 4, Label: "four"},
		{Value: 3, Label: "five"},
		{Value: 3, Label: "six"},
		{Value: 1, Label: "seven"},
	}
	app.Debug("OnMount Sample: %+v", sample)

	data, err := json.Marshal(sample)

	if err != nil {
		app.Debug("Error marshaling: %v", err)
	}

	h.Data = data

	app.Log("OnMount() main 2")
	app.Debug("%v 2", h.Data)

	uiloop = make(chan struct{})

	go func() {
		c := time.Tick(time.Second)
		for _ = range c {
			select {
			default:
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
		<div class="menu"></div>
		<div class="content">
		<div class="chart">
			<Chart data="{{json .Data}}">
		</div>
		</div>
		<div class="footer">FOOTER</div>
	</div>
</body>

	`
}

func (h *Container) OnDismount() {
	uiloop <- struct{}{}
}

/////////////////////
// Chart component //
/////////////////////

type Chart struct {
	Data []byte
}

func (c *Chart) OnMount() {
	app.Log("OnMount() pie")
	app.Debug("%v", c.Data)
}

func (c *Chart) Render() string {
	app.Log("Render() pie")
	app.Debug("ChartData: %v", c.Data)

	test := []chart.Value{}

	if err := json.Unmarshal(c.Data, &test); err != nil {
		app.Debug("Error unmarshling: %v", err)
	}

	app.Debug("Chart Test %v", test)

	buffer := bytes.NewBuffer([]byte{})

	pie := chart.PieChart{
		Width:  300,
		Height: 300,
		Background: chart.Style{
			FillColor:   chart.ColorTransparent,
			StrokeColor: chart.ColorTransparent,
		},
		Canvas: chart.Style{
			FillColor:   chart.ColorTransparent,
			StrokeColor: chart.ColorTransparent,
		},
		SliceStyle: chart.Style{
			StrokeWidth: 1.0,
			StrokeColor: chart.ColorTransparent,
		},
		Values: chart.Values{
			{Value: 5, Label: "Blue"},
			{Value: 5, Label: "Green"},
			{Value: 4, Label: "Gray"},
			{Value: 4, Label: "Orange"},
			{Value: 3, Label: "Deep Blue"},
			{Value: 3, Label: "??"},
			{Value: 1, Label: "!!"},
		},
	}
	app.Debug("%+v", buffer)
	err := pie.Render(chart.SVG, buffer)
	if err != nil {
		app.Debug("Error rendering pie chart: %v\n", err)
		return buffer.String()
	}
	return buffer.String()
}
