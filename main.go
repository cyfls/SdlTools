package main

import (
	"SdlTools/SdlButton"
	"SdlTools/SdlWindow"
)

func main() {
	window := SdlWindow.New(
		800, 600, "My SdlWindow",
		"cour.ttf", 14,
	)
	defer window.Delete()

	button := SdlButton.New(
		window.Surface(),
		window.Font(),
		func() {
		},
	)
	defer button.Delete()
	window.AddChild(button)
	button.SetXY(300, 300)
	button.SetText("Say Hello from the other side")

	window.Mainloop()
}
