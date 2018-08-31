package main

import (
	"SdlTools/SdlButton"
	"SdlTools/SdlWindow"
	"container/ring"
	"log"
)

func main() {
	// It's easy for SdlTools to create a window and a button on it,
	// Just like this:
	window := SdlWindow.New(
		800, 600, "My SdlWindow",
		"cour.ttf", 14,
	)
	defer window.Delete()

	button := SdlButton.New(
		window.Surface(),
		window.Font(),
	)
	defer button.Delete()
	window.AddChild(button)

	// The code below is just for fun
	// If you only need to create a window and a button on it,
	// you just have to input the code up the comment.
	lyrics := ring.New(7)
	lystr := []string{
		"Hello from the other side",
		"I must've called a thousand times",
		"To tell you I'm sorry",
		"For everything that I've done",
		"But when I call you",
		"Never",
		"Seem to be home",
	}
	for i := 0; i < 7; i++ {
		lyrics.Value = lystr[i]
		lyrics = lyrics.Next()
	}
	button.SetXY(300, 300)
	button.SetText("Click me to play Adele's <<Hello>> ~")
	button.SetOnPressFunc(
		func() {
			value, ok := lyrics.Value.(string)
			if !ok {
				log.Fatalln("String convertion failed")
			}
			button.SetText(value)
			lyrics = lyrics.Next()
		},
	)

	// Oh and this is important.
	// If you forget to input this line of code,
	// the window will close meanwhile it shows.
	window.Mainloop()
}
