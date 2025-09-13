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

type EbitenControl struct {
	gamepad *ebiten.GamepadID
}

func (e *EbitenControl) Active(a Action) bool {
	// TODO Player 1 vs player 2 controls, etc.
	// TODO Different mapping options for keyboard?
	switch a {
	case ActionMenu:
		return ebiten.IsKeyPressed(ebiten.KeyEscape) ||
			e.gamepadPressed(ebiten.StandardGamepadButtonCenterRight)
	case ActionUp:
		return ebiten.IsKeyPressed(ebiten.KeyArrowUp) ||
			e.gamepadPressed(ebiten.StandardGamepadButtonLeftTop)
	case ActionDown:
		return ebiten.IsKeyPressed(ebiten.KeyArrowDown) ||
			e.gamepadPressed(ebiten.StandardGamepadButtonLeftBottom)
	case ActionLeft:
		return ebiten.IsKeyPressed(ebiten.KeyArrowLeft) ||
			e.gamepadPressed(ebiten.StandardGamepadButtonLeftLeft)
	case ActionRight:
		return ebiten.IsKeyPressed(ebiten.KeyArrowRight) ||
			e.gamepadPressed(ebiten.StandardGamepadButtonLeftRight)
	case ActionA:
		return ebiten.IsKeyPressed(ebiten.KeyZ) ||
			e.gamepadPressed(ebiten.StandardGamepadButtonRightBottom)
	case ActionB:
		return ebiten.IsKeyPressed(ebiten.KeyX) ||
			e.gamepadPressed(ebiten.StandardGamepadButtonRightRight)
	}
	return false
}

func (e *EbitenControl) Duration(a Action) int {
	// TODO Player 1 vs player 2 controls, etc.
	// TODO Different mapping options for keyboard?
	switch a {
	case ActionMenu:
		return max(
			inpututil.KeyPressDuration(ebiten.KeyEscape),
			e.gamepadDuration(ebiten.StandardGamepadButtonCenterRight),
		)
	case ActionUp:
		return max(
			inpututil.KeyPressDuration(ebiten.KeyArrowUp),
			e.gamepadDuration(ebiten.StandardGamepadButtonLeftTop),
		)
	case ActionDown:
		return max(
			inpututil.KeyPressDuration(ebiten.KeyArrowDown),
			e.gamepadDuration(ebiten.StandardGamepadButtonLeftBottom),
		)
	case ActionLeft:
		return max(
			inpututil.KeyPressDuration(ebiten.KeyArrowLeft),
			e.gamepadDuration(ebiten.StandardGamepadButtonLeftLeft),
		)
	case ActionRight:
		return max(
			inpututil.KeyPressDuration(ebiten.KeyArrowRight),
			e.gamepadDuration(ebiten.StandardGamepadButtonLeftRight),
		)
	case ActionA:
		return max(
			inpututil.KeyPressDuration(ebiten.KeyZ),
			e.gamepadDuration(ebiten.StandardGamepadButtonRightBottom),
		)
	case ActionB:
		return max(
			inpututil.KeyPressDuration(ebiten.KeyX),
			e.gamepadDuration(ebiten.StandardGamepadButtonRightRight),
		)
	}
	return 0
}

func JustActive(c Control, a Action) bool {
	return c.Duration(a) == 1
}

func (e *EbitenControl) gamepadDuration(b ebiten.StandardGamepadButton) int {
	if e.gamepad != nil {
		return inpututil.StandardGamepadButtonPressDuration(*e.gamepad, b)
	} else {
		return 0
	}
}

func (e *EbitenControl) gamepadPressed(b ebiten.StandardGamepadButton) bool {
	return e.gamepad != nil &&
		ebiten.IsStandardGamepadButtonPressed(*e.gamepad, b)
}
