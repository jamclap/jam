package jam

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game interface {
	Update(hub *Hub)
	Draw(draw *Draw)
}

type ebitenGame struct {
	game Game
	draw Draw
	hub  *Hub
}

func Run(init func(hub *Hub) Game) {
	ebiten.SetScreenClearedEveryFrame(false)
	ebiten.SetWindowSize(960, 540)
	hub := &Hub{
		Control: &EbitenControl{},
		Window:  &EbitenWindow{},
	}
	app := init(hub)
	app.Update(hub)
	e := &ebitenGame{game: app, hub: hub}
	if err := ebiten.RunGame(e); err != nil {
		log.Fatal(err)
	}
}

func (e *ebitenGame) Update() error {
	checkFullscreen()
	e.hub.update()
	e.game.Update(e.hub)
	return nil
}

func (e *ebitenGame) Draw(image *ebiten.Image) {
	e.draw.Target = image
	e.game.Draw(&e.draw)
}

func (e *ebitenGame) Layout(
	outsideWidth, outsideHeight int,
) (screenWidth, screenHeight int) {
	screenWidth = 160
	screenHeight = 90
	return
}

func checkFullscreen() {
	if inpututil.IsKeyJustPressed(ebiten.KeyF11) ||
		(inpututil.IsKeyJustPressed(ebiten.KeyEnter) && ebiten.IsKeyPressed(ebiten.KeyAlt)) {
		ebiten.SetFullscreen(!ebiten.IsFullscreen())
	}
}
