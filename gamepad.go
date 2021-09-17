package utils

// Gamepad handles the changes to a single gamepad
type Gamepad struct {
	gamepad *uint8
	state   uint8
}

// JustPressedButtons returns all buttons that were pressed since the last
// update.
func (g *Gamepad) JustPressedButtons() byte {
	result := *g.gamepad & (*g.gamepad ^ g.state)

	return result
}

// IsJustPressed checks if a specific button was pressed since the last
// update.
func (g *Gamepad) IsJustPressed(button byte) bool {
	return g.JustPressedButtons()&button != 0
}

// IsPressed checks if a specific button is currently pressed.
func (g *Gamepad) IsButtonJustPressed(button byte) bool {
	return *g.gamepad&button != 0
}

// Update updates the internal state of the gamepad.
//
// Call this at the end of the update-function!
func (g *Gamepad) Update() {
	g.state = *g.gamepad
}

// Creates a new Gamepad and stores the current.
func NewGamepad(gamepad *uint8) *Gamepad {
	return &Gamepad{
		gamepad: gamepad,
		state:   *gamepad,
	}
}
