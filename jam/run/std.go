package run

import (
	"image/color"
	"reflect"

	"github.com/traefik/yaegi/interp"
)

var MiniStdExports = interp.Exports{
	"image/color/color": {
		"Color":  reflect.ValueOf((*color.Color)(nil)),
		"_Color": reflect.ValueOf((*_image_color_Color)(nil)),
	},
}

type _image_color_Color struct {
	IValue interface{}
	WRGBA  func() (r uint32, g uint32, b uint32, a uint32)
}

func (W _image_color_Color) RGBA() (r uint32, g uint32, b uint32, a uint32) { return W.WRGBA() }
