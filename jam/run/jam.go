package run

import (
	"reflect"

	"github.com/jamclap/jam/jam"
	"github.com/jamclap/jam/jam/pal"
	"github.com/traefik/yaegi/interp"
)

var jamExports = interp.Exports{
	"embed/embed": {},
	"github.com/jamclap/jam/jam/jam": {
		"Draw":    reflect.ValueOf((*jam.Draw)(nil)),
		"Game":    reflect.ValueOf((*jam.Game)(nil)),
		"Hub":     reflect.ValueOf((*jam.Hub)(nil)),
		"LoadMap": reflect.ValueOf(jam.LoadMap),
		"Run":     reflect.ValueOf(jam.Run),
		"XY":      reflect.ValueOf(jam.XY[int]), // TODO Generic???
		"_Game":   reflect.ValueOf((*_jam_Game)(nil)),
	},
	"github.com/jamclap/jam/jam/pal/pal": {
		"Jam": reflect.ValueOf(pal.Jam),
	},
}

type _jam_Game struct {
	IValue  interface{}
	WUpdate func(hub *jam.Hub)
	WDraw   func(draw *jam.Draw)
}

func (w _jam_Game) Update(h *jam.Hub) { w.WUpdate(h) }
func (w _jam_Game) Draw(d *jam.Draw)  { w.WDraw(d) }
