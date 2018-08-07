package main

// import (
// 	"github.com/murlokswarm/app"
// )

type Test struct {
	To string
}

func (t *Test) Render() string {
	return `
    <div>
      <a href="{{.To}}">Index</a>
    </div>
  `
}
