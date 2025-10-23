package jam

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Hub struct {
	Control      *EbitenControl
	TileSheets   []*Sheet
	Window       Window
	gamepadsGone []bool
	gamepads     []ebiten.GamepadID
}

func (h *Hub) update() {
	h.updateGamepads()
	h.updateControls()
}

func (h *Hub) updateControls() {
	if len(h.gamepads) > 0 && !h.gamepadsGone[0] {
		gamepad := h.gamepads[0]
		h.Control.gamepad = &gamepad
	} else {
		h.Control.gamepad = nil
	}
}

func (h *Hub) updateGamepads() {
	// Out with the old.
	for i, id := range h.gamepads {
		if inpututil.IsGamepadJustDisconnected(id) {
			h.gamepadsGone[i] = true
		}
	}
	// In with the new.
	oldLen := len(h.gamepads)
	h.gamepads = inpututil.AppendJustConnectedGamepadIDs(h.gamepads)
	if oldLen == len(h.gamepads) {
		// Nothing new here.
		return
	}
	for i := oldLen; i < len(h.gamepads); i++ {
		h.gamepadsGone = append(h.gamepadsGone, false)
	}
	// Move new ones into empty old slots.
	// Go in reverse order for convenience. Likely only one, anyway.
	// TODO If ids persist, try to put id into same old slot.
	// TODO See ebiten.GamepadSDLID? Is that instance or kind?
	iNew := len(h.gamepads) - 1
	for i := range h.gamepads[:oldLen] {
		if iNew < oldLen {
			// No news left.
			break
		}
		if h.gamepadsGone[i] {
			h.gamepads[i] = h.gamepads[iNew]
			h.gamepadsGone[i] = false
		}
	}
	if iNew < len(h.gamepads)-1 {
		// We apparently moved at least one.
		h.gamepads = h.gamepads[:iNew+1]
		h.gamepadsGone = h.gamepadsGone[:iNew+1]
	}
}
