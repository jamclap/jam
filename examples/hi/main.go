package main

import (
	"github.com/jamclap/jam/jam"
	"github.com/jamclap/jam/jam/pal"
)

func main() {
	println("Hi!")
	jam.Run(InitState)
}

type Game struct{}

func InitState(hub *jam.Hub) jam.Game {
	hub.Window.SetTitle("Hi there!")
	return &Game{}
}

func (g *Game) Update(hub *jam.Hub) {}

func (g *Game) Draw(draw *jam.Draw) {
	draw.Fill(pal.Jam.Blue1)
}
