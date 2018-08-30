package SdlTextArea

import (
	"container/list"
	"log"

	"github.com/banthar/Go-SDL/sdl"
	"github.com/banthar/Go-SDL/ttf"
)

type SdlTextArea struct {
	parent     *sdl.Surface
	font       *ttf.Font
	rect       *sdl.Rect
	frontColor *sdl.Color
	backColor  uint32
	contents   *list.List
	msgSurfs   *list.List
}

func New(parent *sdl.Surface, font *ttf.Font) *SdlTextArea {
	area := &SdlTextArea{}
	area.parent = parent
	area.font = font
	area.rect = &sdl.Rect{
		X: 0,
		Y: 0,
		W: 0,
		H: 0,
	}
	area.frontColor = &sdl.Color{}
	area.frontColor.R = 255
	area.frontColor.G = 255
	area.frontColor.B = 255

	area.backColor = sdl.MapRGB(
		area.parent.Format,
		200, 200, 200,
	)

	area.contents = list.New()
	area.msgSurfs = list.New()

	return area
}

func (this *SdlTextArea) Font() *ttf.Font {
	return this.font
}

func (this *SdlTextArea) XY() (int, int) {
	return int(this.rect.X), int(this.rect.Y)
}

func (this *SdlTextArea) SetXY(x, y int) {
	this.rect.X = int16(x)
	this.rect.Y = int16(y)
}

func (this *SdlTextArea) FrontColor() *sdl.Color {
	return this.frontColor
}

func (this *SdlTextArea) SetFrontColorRGB(r, g, b int) {
	this.frontColor.R = uint8(r)
	this.frontColor.G = uint8(g)
	this.frontColor.B = uint8(b)
}

func (this *SdlTextArea) SetFrontColor(color *sdl.Color) {
	this.frontColor = color
}

func (this *SdlTextArea) BackColor() uint32 {
	return this.backColor
}

func (this *SdlTextArea) SetBackColorRGB(r, g, b int) {
	this.backColor = sdl.MapRGB(
		this.parent.Format,
		uint8(r), uint8(g), uint8(b),
	)
}

func (this *SdlTextArea) SetBackColor(color uint32) {
	this.backColor = color
}

func (this *SdlTextArea) AppendText(txt string) {
	this.contents.PushBack(txt)
	surf := ttf.RenderText_Solid(
		this.font,
		txt,
		*this.frontColor,
	)
	this.msgSurfs.PushBack(surf)
	if surf.W+5 > int32(this.rect.W) {
		this.rect.W = uint16(surf.W + 10)
	}
	globalHeight := 5
	for e := this.msgSurfs.Front(); e != nil; e = e.Next() {
		value, ok := e.Value.(*sdl.Surface)
		if !ok {
			log.Fatalln("Surface convertion failed.")
		}
		globalHeight += int(value.H + 5)
	}
	if globalHeight > int(this.rect.H) {
		this.rect.H = uint16(globalHeight + 5)
	}
}

func (this *SdlTextArea) Clear() {
	this.Delete()
	this.contents = list.New()
	this.msgSurfs = list.New()
	this.rect.W, this.rect.H = 0, 0
}

func (this *SdlTextArea) Show() {
	this.parent.FillRect(
		this.rect,
		this.backColor,
	)
	gapLen := 0
	offset := &sdl.Rect{}
	offset.X = this.rect.X + 5
	offset.Y = this.rect.Y + 5
	for e := this.msgSurfs.Front(); e != nil; e = e.Next() {
		offset.Y += int16(gapLen + 5)
		value, ok := e.Value.(*sdl.Surface)
		if !ok {
			log.Fatalln("Surface convertion failed.")
		}
		this.parent.Blit(offset, value, nil)
		gapLen = int(value.H)
	}
}

func (this *SdlTextArea) Delete() {
	for e := this.msgSurfs.Front(); e != nil; e = e.Next() {
		value, ok := e.Value.(*sdl.Surface)
		if !ok {
			log.Fatalln("Surface convertion failed.")
		}
		value.Free()
	}
}
