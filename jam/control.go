package jam

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Action int

const (
	ActionMenu Action = iota
	ActionUp
	ActionDown
	ActionLeft
	ActionRight
	ActionA
	ActionB
)

type Control interface {
	Active(a Action) bool
	Duration(a Action) int
}

// func (c *Control) Active(action Action) bool {
// 	return false
// }

type EbitenControl struct{}

func (e *EbitenControl) Active(a Action) bool {
	// TODO Player 1 vs player 2 controls, etc.
	// TODO Different mapping options for keyboard?
	switch a {
	case ActionMenu:
		return ebiten.IsKeyPressed(ebiten.KeyEscape)
	case ActionUp:
		return ebiten.IsKeyPressed(ebiten.KeyArrowUp)
	case ActionDown:
		return ebiten.IsKeyPressed(ebiten.KeyArrowUp)
	case ActionLeft:
		return ebiten.IsKeyPressed(ebiten.KeyArrowLeft)
	case ActionRight:
		return ebiten.IsKeyPressed(ebiten.KeyArrowRight)
	case ActionA:
		return ebiten.IsKeyPressed(ebiten.KeyZ)
	case ActionB:
		return ebiten.IsKeyPressed(ebiten.KeyX)
	}
	return false
}

func (e *EbitenControl) Duration(a Action) int {
	// TODO Player 1 vs player 2 controls, etc.
	// TODO Different mapping options for keyboard?
	switch a {
	case ActionMenu:
		return inpututil.KeyPressDuration(ebiten.KeyEscape)
	case ActionUp:
		return inpututil.KeyPressDuration(ebiten.KeyArrowUp)
	case ActionDown:
		return inpututil.KeyPressDuration(ebiten.KeyArrowUp)
	case ActionLeft:
		return inpututil.KeyPressDuration(ebiten.KeyArrowLeft)
	case ActionRight:
		return inpututil.KeyPressDuration(ebiten.KeyArrowRight)
	case ActionA:
		return inpututil.KeyPressDuration(ebiten.KeyZ)
	case ActionB:
		return inpututil.KeyPressDuration(ebiten.KeyX)
	}
	return 0
}

func JustActive(c Control, a Action) bool {
	return c.Duration(a) == 1
}
