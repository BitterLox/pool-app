package composition

import (
	// "bytes"
	// "encoding/json"
	"github.com/BitterLox/flux"
	"github.com/murlokswarm/app"
	"github.com/wcharczuk/go-chart"
	"time"
)

type Test struct {
	A int
	B int
	C int
}

type Value struct {
	Label string
	Value float64
}

type Values []Value

func (v Values) convertData() []chart.Value {
	result := make([]chart.Value, len(v))
	for i := 0; i < len(v); i++ {
		result = append(result, chart.Value{Value: v[i].Value, Label: v[i].Label})
	}
	return result
}

type Hello struct {
	Name string
}

func (h *Hello) OnMount() {
	app.Log("%v", h)

	// h.Name = "daaaam"

	// h.Name = &Values{
	// 	{Value: 5, Label: "one"},
	// 	{Value: 5, Label: "two"},
	// 	{Value: 4, Label: "three"},
	// 	{Value: 4, Label: "four"},
	// 	{Value: 3, Label: "five"},
	// 	{Value: 3, Label: "six"},
	// 	{Value: 1, Label: "seven"},
	// }

	// test := chartData.convertData()
	//
	// app.Debug("convertData test: %#v %T", test, test)

	// data, err := json.Marshal(chartData)
	//
	// if err != nil {
	// 	app.Debug("Error marshaling data: %v", err)
	// }

	// h.Name = &data

	app.Debug("OnMount() Hello")
	app.Debug("OnMount() Hello: %v", h.Name)

	uiloop := make(chan struct{})

	go func() {
		c := time.Tick(500 * time.Millisecond)
		for _ = range c {
			select {
			default:
				app.Debug("Tick")

				flux.Dispatch(flux.Action{
					Name:    "greet",
					Payload: "Emanuele",
				})

				// app.Render(h)
			case <-uiloop:
				return
			}
		}
	}()
}

func (h *Hello) Render() string {
	return `
<div>
    <h1>
				{{json .Name}}
				<composition.World name="">
    </h1>
</div>
    `
}

type World struct {
	Name string
}

func (w *World) OnMount() {
	app.Debug("OnMount() World")
	app.Debug("OnMount() World: %v", w.Name)
	w.Name = "ding"
	app.Render(w)
}

func (w *World) OnStoreEvent(e flux.Event) {
	if e.Name != "greeted" {
		return
	}

	// Handling events from helloStore.
	if greet, ok := e.Payload.(string); ok {
		app.Log("Store event: %v", greet)
		// w.Name = greet
		app.Log("World: %#v", w)
		app.Render(w)
	}
}

func (w *World) Render() string {
	return `
<span>
    {{.Name}}
</span>
    `
}
