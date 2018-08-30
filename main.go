package main

import (
	"fmt"
	"log"

	"SdlTools/SdlButton"

	"github.com/banthar/Go-SDL/sdl"
	"github.com/banthar/Go-SDL/ttf"
)

const (
	WINDOW_WIDTH  = 800
	WINDOW_HEIGHT = 600
	WINDOW_TITLE  = "My SDL Button"
)

var (
	/*
	 *
	 * SDL Variables
	 */
	window *sdl.Surface

	/*
	 *
	 * SDL_ttf Variables
	 */
	font *ttf.Font

	/*
	 *
	 * Mainloop Variables
	 */
	running bool

	/*
	 *
	 * Other Variables
	 */
	buttons []*SdlButton.SdlButton
)

func draw() {
	windRect := &sdl.Rect{
		X: 0,
		Y: 0,
		W: WINDOW_WIDTH,
		H: WINDOW_HEIGHT,
	}
	whiteColor := sdl.MapRGB(
		window.Format,
		255, 255, 255,
	)
	window.FillRect(windRect, whiteColor)

	for _, v := range buttons {
		v.Show()
	}

	if window.Flip() < 0 {
		log.Fatalln(sdl.GetError())
	}
}

func mainloop() {
	running = true
	for running {
		for ev := sdl.PollEvent(); ev != nil; ev = sdl.PollEvent() {
			for _, v := range buttons {
				v.Activate(ev)
			}
			switch e := ev.(type) {
			case *sdl.QuitEvent:
				running = false
			case *sdl.KeyboardEvent:
				if e.Type == sdl.KEYUP {
					if e.Keysym.Sym == sdl.K_ESCAPE {
						running = false
					}
				}
			}
		}
		sdl.Delay(50)
		draw()
	}
}

func main() {
	/*
	 *
	 * SDL Initialization
	 */
	if sdl.Init(sdl.INIT_EVERYTHING) < 0 {
		log.Fatalln(sdl.GetError())
	}
	defer sdl.Quit()
	sdl.WM_SetCaption(WINDOW_TITLE, "")
	window = sdl.SetVideoMode(
		WINDOW_WIDTH,
		WINDOW_HEIGHT,
		32,
		sdl.SWSURFACE,
	)
	if window == nil {
		log.Fatalln(sdl.GetError())
	}

	/*
	 *
	 * SDL_ttf Initialization
	 */
	if ttf.Init() < 0 {
		log.Fatalln("Failed to init SDL_ttf.")
	}
	defer ttf.Quit()
	font = ttf.OpenFont("cour.ttf", 12)
	if font == nil {
		log.Fatalln("Failed to load font.")
	}
	defer font.Close()

	/*
	 *
	 * Other Initialization
	 */
	buttons = []*SdlButton.SdlButton{}

	button := SdlButton.New(window, font, func() {
		fmt.Println("Hello From Button")
	})
	defer button.Delete()
	buttons = append(buttons, button)
	button.SetText("Hello from the other side~~~")

	/*
	 *
	 * Start Mainloop
	 */
	mainloop()
}
