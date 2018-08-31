package SdlWindow

import (
	"container/list"
	"log"

	"SdlTools/SdlButton"
	"SdlTools/SdlUtil"

	"github.com/banthar/Go-SDL/sdl"
	"github.com/banthar/Go-SDL/ttf"
)

type SdlWindow struct {
	surface *sdl.Surface
	font    *ttf.Font
	running bool
	draw    func(*sdl.Surface)
	onEvent func(sdl.Event)
	childs  *list.List
	waitMs  int
}

func New(width, height int, title, font string, fontSize int) *SdlWindow {
	if sdl.Init(sdl.INIT_EVERYTHING) < 0 {
		log.Fatalln(sdl.GetError())
	}
	sdl.WM_SetCaption(title, "")
	this := &SdlWindow{}
	this.surface = sdl.SetVideoMode(
		width,
		height,
		32,
		sdl.SWSURFACE,
	)
	if ttf.Init() < 0 {
		log.Fatalln("Failed to init SDL_ttf.")
	}
	this.font = ttf.OpenFont(font, fontSize)
	if this.font == nil {
		log.Fatalln("Failed to load font.")
	}
	this.running = true
	this.draw = func(surf *sdl.Surface) {
		windRect := &sdl.Rect{
			X: 0,
			Y: 0,
			W: uint16(this.surface.W),
			H: uint16(this.surface.H),
		}
		whiteColor := sdl.MapRGB(
			this.surface.Format,
			255, 255, 255,
		)
		this.surface.FillRect(windRect, whiteColor)
	}
	this.onEvent = func(e sdl.Event) {}
	this.childs = list.New()
	this.waitMs = 50
	return this
}

func (this *SdlWindow) Surface() *sdl.Surface {
	return this.surface
}

func (this *SdlWindow) Font() *ttf.Font {
	return this.font
}

func (this *SdlWindow) SetFont(font *ttf.Font) {
	this.font.Close()
	this.font = font
}

func (this *SdlWindow) Quit() {
	this.running = false
}

func (this *SdlWindow) DrawFunc() func(*sdl.Surface) {
	return this.draw
}

func (this *SdlWindow) SetDrawFunc(draw func(*sdl.Surface)) {
	this.draw = draw
}

func (this *SdlWindow) OnEventFunc() func(sdl.Event) {
	return this.onEvent
}

func (this *SdlWindow) SetOnEventFunc(onEvent func(sdl.Event)) {
	this.onEvent = onEvent
}

func (this *SdlWindow) AddChild(child interface{}) {
	this.childs.PushBack(child)
}

func (this *SdlWindow) RemoveChild(child interface{}) {
	for e := this.childs.Front(); e != nil; e = e.Next() {
		if e.Value == child {
			this.childs.Remove(e)
		}
	}
}

func (this *SdlWindow) RefreshPeriod() int {
	return this.waitMs
}

func (this *SdlWindow) SetRefreshPeriod(waitMs int) {
	this.waitMs = waitMs
}

func (this *SdlWindow) Mainloop() {
	for this.running {
		for ev := sdl.PollEvent(); ev != nil; ev = sdl.PollEvent() {
			for e := this.childs.Front(); e != nil; e = e.Next() {
				switch v := e.Value.(type) {
				case *SdlButton.SdlButton:
					v.Activate(ev)
				}
			}
			this.onEvent(ev)
			switch ev.(type) {
			case *sdl.QuitEvent:
				this.running = false
			}
		}
		sdl.Delay(uint32(this.waitMs))
		this.draw(this.surface)
		for e := this.childs.Front(); e != nil; e = e.Next() {
			switch v := e.Value.(type) {
			case SdlUtil.Showable:
				v.Show()
			}
		}
		if this.surface.Flip() < 0 {
			log.Fatalln(sdl.GetError())
		}
	}
}

func (this *SdlWindow) Delete() {
	this.font.Close()
	ttf.Quit()
	sdl.Quit()
}
