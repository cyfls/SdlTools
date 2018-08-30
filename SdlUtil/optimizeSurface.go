package SdlUtil

import (
	"log"

	"github.com/banthar/Go-SDL/sdl"
)

func LoadImage(file string) *sdl.Surface {
	tempImage := sdl.Load(file)
	if tempImage == nil {
		log.Fatalln(sdl.GetError())
	}
	defer tempImage.Free()

	ret := sdl.DisplayFormat(tempImage)
	return ret
}
