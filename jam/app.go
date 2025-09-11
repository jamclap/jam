package jam

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type App interface {
	Update(hub *Hub)
	Draw(draw *Draw)
}

type game struct {
	app  App
	draw Draw
	hub  *Hub
}

func Run(init func(hub *Hub) App) {
	ebiten.SetScreenClearedEveryFrame(false)
	ebiten.SetWindowSize(960, 540)
	hub := &Hub{}
	app := init(hub)
	app.Update(hub)
	g := &game{app: app, hub: hub}
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}

func (g *game) Update() error {
	checkFullscreen()
	g.app.Update(g.hub)
	return nil
}

func (g *game) Draw(image *ebiten.Image) {
	g.draw.Target = image
	g.app.Draw(&g.draw)
}

func (g *game) Layout(
	outsideWidth, outsideHeight int,
) (screenWidth, screenHeight int) {
	screenWidth = 240
	screenHeight = 135
	return
}

func checkFullscreen() {
	if inpututil.IsKeyJustPressed(ebiten.KeyF11) ||
		(inpututil.IsKeyJustPressed(ebiten.KeyEnter) && ebiten.IsKeyPressed(ebiten.KeyAlt)) {
		ebiten.SetFullscreen(!ebiten.IsFullscreen())
	}
}
