package composition

import (
	"errors"
	"fmt"
	"github.com/BitterLox/flux"
	"github.com/murlokswarm/app"
)

// Implemtation.
type HelloStore struct {
	flux.Store
}

func (s *HelloStore) OnDispatch(a flux.Action) error {
	if a.Name != "greet" {
		return errors.New("I only greet")
	}

	app.Log(a.Payload.(string))

	s.Emit(flux.Event{
		Name:    "greeted",
		Payload: fmt.Sprintf("Hello, %v", a.Payload),
	})
	return nil
}
