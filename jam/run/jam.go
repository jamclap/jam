package run

import (
	"reflect"

	"github.com/jamclap/jam/jam"
	"github.com/jamclap/jam/jam/pal"
	"github.com/traefik/yaegi/interp"
)

var jamExports = interp.Exports{
	"github.com/jamclap/jam/jam/jam": {
		"Draw":  reflect.ValueOf((*jam.Draw)(nil)),
		"Game":  reflect.ValueOf((*jam.Game)(nil)),
		"Hub":   reflect.ValueOf((*jam.Hub)(nil)),
		"Run":   reflect.ValueOf(jam.Run),
		"_Game": reflect.ValueOf((*_jam_Game)(nil)),
	},
	"github.com/jamclap/jam/jam/pal/pal": {
		"Jam":      reflect.ValueOf(pal.Jam),
		"JamBlue1": reflect.ValueOf(pal.JamBlue1),
	},
}

type _jam_Game struct {
	IValue  interface{}
	WUpdate func(hub *jam.Hub)
	WDraw   func(draw *jam.Draw)
}

func (w _jam_Game) Update(h *jam.Hub) { w.WUpdate(h) }
func (w _jam_Game) Draw(d *jam.Draw)  { w.WDraw(d) }
