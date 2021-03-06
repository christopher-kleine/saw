# SAW - **S**ome **A**dditions (to) **W**ASM-4

A collection of some functions that make the use of the awesome [WASM-4](https://wasm4.org) a bit easier.

## Features:

- [ ] Simple Gamepad API
	- [x] JustPressed
	- [x] IsJustPressed
	- [x] IsPressed
	- [x] Duration
	- [ ] JustReleased
	- [ ] IsJustReleased
- [ ] Simple PRNG
	- [x] Int
	- [x] Intn
	- [ ] Uint
	- [ ] Uintn
- [ ] Tileset-Functions
	- [ ] DrawTile
- [ ] Animations
- [ ] Particle Generators

Hello-World in WASM-4 without utils:
```go
package main

import "cart/w4"

var smiley = [8]byte{
	0b11000011,
	0b10000001,
	0b00100100,
	0b00100100,
	0b00000000,
	0b00100100,
	0b10011001,
	0b11000011,
}

//go:export update
func update() {
	*w4.DRAW_COLORS = 2
	w4.Text("Hello from Go!", 10, 10)

	var gamepad = *w4.GAMEPAD1
	if gamepad&w4.BUTTON_1 != 0 {
		*w4.DRAW_COLORS = 4
	}

	w4.Blit(&smiley[0], 76, 76, 8, 8, w4.BLIT_1BPP)
	w4.Text("Press X to blink", 16, 90)
}
```

The same using utils:
```go
package main

import "cart/w4"
import "christopher-kleine/saw"

var smiley = [8]byte{
	0b11000011,
	0b10000001,
	0b00100100,
	0b00100100,
	0b00000000,
	0b00100100,
	0b10011001,
	0b11000011,
}
var gamepad = saw.NewGamepad(w4.GAMEPAD1)

//go:export update
func update() {
	*w4.DRAW_COLORS = 2
	w4.Text("Hello from Go!", 10, 10)

	if gamepad.IsPressed(w4.BUTTON_1) {
		*w4.DRAW_COLORS = 4
	}

	w4.Blit(&smiley[0], 76, 76, 8, 8, w4.BLIT_1BPP)
	w4.Text("Press X to blink", 16, 90)
}
```
