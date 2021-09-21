package utils

var (
	buttons = []byte{
		1,   // BUTTON_1
		2,   // BUTTON_2
		16,  // BUTTON_LEFT
		32,  // BUTTON_RIGHT
		64,  // BUTTON_UP
		128, // BUTTON_DOWN
	}
)

// Gamepad handles the changes to a single gamepad
type Gamepad struct {
	gamepad   *uint8
	state     uint8
	durations map[byte]uint
}

// JustPressed returns all buttons that were pressed since the last
// update.
func (g *Gamepad) JustPressed() byte {
	result := *g.gamepad & (*g.gamepad ^ g.state)

	return result
}

// IsJustPressed checks if a specific button was pressed since the last
// update.
func (g *Gamepad) IsJustPressed(button byte) bool {
	return g.JustPressed()&button != 0
}

// IsPressed checks if a specific button is currently pressed.
func (g *Gamepad) IsPressed(button byte) bool {
	return *g.gamepad&button != 0
}

// Duration returns the amount of frames a certain button is beeing pressed.
func (g *Gamepad) Duration(button byte) uint {
	v, ok := g.durations[button]
	if !ok {
		return 0
	}

	// Add the current frame to the duration
	if g.IsPressed(button) {
		return v + 1
	}

	return 0
}

// Update updates the internal state of the gamepad.
//
// Call this at the end of the update-function!
func (g *Gamepad) Update() {
	for _, button := range buttons {
		g.durations[button]++
		if !g.IsPressed(button) {
			g.durations[button] = 0
		}
	}

	g.state = *g.gamepad
}

// Creates a new Gamepad and stores the current.
func NewGamepad(gamepad *uint8) *Gamepad {
	return &Gamepad{
		gamepad: gamepad,
		state:   *gamepad,
	}
}
