package main

import (
	"SdlTools/SdlButton"
	"SdlTools/SdlWindow"
	"fmt"
)

func main() {
	window := SdlWindow.New(
		800, 600, "My SdlWindow",
		"cour.ttf",
	)
	defer window.Delete()

	button := SdlButton.New(
		window.Surface(),
		window.Font(),
		func() {
			fmt.Println("Hello!")
		},
	)
	defer button.Delete()
	window.AddChild(button)

	window.Mainloop()
}
