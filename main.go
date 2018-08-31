package main

import (
	"SdlTools/SdlButton"
	"SdlTools/SdlTextArea"
	"SdlTools/SdlWindow"

	"github.com/banthar/Go-SDL/sdl"
)

func main() {
	window := SdlWindow.New(
		800, 600, "My SdlWindow",
		"cour.ttf", 14,
	)
	defer window.Delete()
	window.SetOnEventFunc(func(ev sdl.Event) {
		switch e := ev.(type) {
		case *sdl.KeyboardEvent:
			if e.Type == sdl.KEYUP {
				switch e.Keysym.Sym {
				case sdl.K_ESCAPE:
					window.Quit()
				}
			}
		}
	})

	area := SdlTextArea.New(
		window.Surface(),
		window.Font(),
	)
	defer area.Delete()
	window.AddChild(area)

	button := SdlButton.New(
		window.Surface(),
		window.Font(),
		func() {
			area.AppendText("Hello, SDL!")
		},
	)
	defer button.Delete()
	window.AddChild(button)
	button.SetXY(650, 500)
	button.SetText("Say Hello")

	window.Mainloop()
}
