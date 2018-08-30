package SdlButton

import (
	"github.com/banthar/Go-SDL/sdl"
	"github.com/banthar/Go-SDL/ttf"
)

type SdlButton struct {
	parent       *sdl.Surface
	font         *ttf.Font
	onPress      func()
	rect         *sdl.Rect
	frontColor   *sdl.Color
	backColors   map[string]uint32
	text         string
	textSurf     *sdl.Surface
	pressed      bool
	currentColor uint32
}

func New(parent *sdl.Surface, font *ttf.Font, onPress func()) *SdlButton {
	this := &SdlButton{}
	this.parent, this.font = parent, font
	this.onPress = onPress
	this.rect = &sdl.Rect{
		X: 0,
		Y: 0,
		W: 15,
		H: 15,
	}
	this.frontColor = &sdl.Color{}
	this.frontColor.R = 0
	this.frontColor.G = 0
	this.frontColor.B = 0
	this.backColors = map[string]uint32{}
	this.backColors["normal"] = sdl.MapRGB(
		this.parent.Format,
		255, 255, 255,
	)
	this.backColors["focused"] = sdl.MapRGB(
		this.parent.Format,
		255, 127, 39,
	)
	this.backColors["pressed"] = sdl.MapRGB(
		this.parent.Format,
		237, 28, 36,
	)
	this.text = "Button"
	this.textSurf = ttf.RenderText_Solid(
		this.font,
		this.text,
		*this.frontColor,
	)
	textX := this.rect.X + 10
	textY := this.rect.Y + 10
	if int32(textX)+this.textSurf.W > int32(this.rect.W-15) {
		this.rect.W = uint16(this.textSurf.W + 20)
	}
	if int32(textY)+this.textSurf.H > int32(this.rect.H-15) {
		this.rect.H = uint16(this.textSurf.H + 20)
	}
	this.pressed = false
	this.currentColor = this.backColors["normal"]
	return this
}

func (this *SdlButton) Font() *ttf.Font {
	return this.font
}

func (this *SdlButton) OnPressFunc() func() {
	return this.onPress
}

func (this *SdlButton) SetOnPressFunc(onPress func()) {
	this.onPress = onPress
}

func (this *SdlButton) XY() (int, int) {
	return int(this.rect.X), int(this.rect.Y)
}

func (this *SdlButton) SetXY(x, y int) {
	this.rect.X, this.rect.Y = int16(x), int16(y)
}

func (this *SdlButton) Size() (int, int) {
	return int(this.rect.W), int(this.rect.H)
}

func (this *SdlButton) SetSize(w, h int) {
	this.rect.W = uint16(w)
	this.rect.H = uint16(h)
}

func (this *SdlButton) InButtonArea(x, y int) bool {
	if x > int(this.rect.X) && x < int(this.rect.X+int16(this.rect.W)) {
		if y > int(this.rect.Y) && y < int(this.rect.Y+int16(this.rect.H)) {
			return true
		}
	}
	return false
}

func (this *SdlButton) BackColor(str string) uint32 {
	return this.backColors[str]
}

func (this *SdlButton) SetBackColor(str string, color uint32) {
	this.backColors[str] = color
}

func (this *SdlButton) Text() string {
	return this.text
}

func (this *SdlButton) SetText(text string) {
	this.text = text
	this.Delete()

	this.textSurf = ttf.RenderText_Solid(
		this.font,
		this.text,
		*this.frontColor,
	)

	textX := this.rect.X + 10
	textY := this.rect.Y + 10
	if int32(textX)+this.textSurf.W > int32(this.rect.W-15) {
		this.rect.W = uint16(this.textSurf.W + 20)
	}

	if int32(textY)+this.textSurf.H > int32(this.rect.H-15) {
		this.rect.H = uint16(this.textSurf.H + 20)
	}
}

func (this *SdlButton) Activate(ev sdl.Event) {
	switch e := ev.(type) {
	case *sdl.MouseButtonEvent:
		if e.Type == sdl.MOUSEBUTTONDOWN {
			this.pressed = true
			if this.InButtonArea(int(e.X), int(e.Y)) {
				this.currentColor = this.backColors["pressed"]
			} else {
				this.currentColor = this.backColors["normal"]
			}
		} else if e.Type == sdl.MOUSEBUTTONUP {
			this.pressed = false
			if this.InButtonArea(int(e.X), int(e.Y)) {
				this.currentColor = this.backColors["focused"]
				this.onPress()
			} else {
				this.currentColor = this.backColors["normal"]
			}
		}
	case *sdl.MouseMotionEvent:
		if this.InButtonArea(int(e.X), int(e.Y)) {
			if this.pressed {
				this.currentColor = this.backColors["pressed"]
			} else {
				this.currentColor = this.backColors["focused"]
			}
		} else {
			this.currentColor = this.backColors["normal"]
		}
	}
}

func (this *SdlButton) Show() {
	blackColor := sdl.MapRGB(
		this.parent.Format,
		0, 0, 0,
	)
	this.parent.FillRect(this.rect, blackColor)

	smallrect := &sdl.Rect{
		X: this.rect.X + 5,
		Y: this.rect.Y + 5,
		W: this.rect.W - 10,
		H: this.rect.H - 10,
	}
	this.parent.FillRect(smallrect, this.currentColor)

	txtrect := &sdl.Rect{
		X: smallrect.X + 5,
		Y: smallrect.X + 5,
		W: uint16(this.textSurf.W),
		H: uint16(this.textSurf.H),
	}
	this.parent.Blit(txtrect, this.textSurf, nil)
}

func (this *SdlButton) Delete() {
	this.textSurf.Free()
}
