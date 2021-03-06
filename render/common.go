package render

import (
	"github.com/veandco/go-sdl2/sdl"
)

const (
	Yadjust     = 100
	fontWidth   = 8
	fontHeight  = 9
	nbCodeLines = 20

	setFPS      = 60
	throttleFPS = 1000 / setFPS
)

type KEYPressed struct {
	KeyCode uint
	Mode    uint
}

var fps, frameCount, lastFrame, lastTime, timerFPS uint32
var Xadjust int

func getGlyph(char rune) *sdl.Rect {
	pos := int32(char - 32)
	// posy := int32(pos / 18)
	// posx := pos - int32(pos / 18)*18
	// fmt.Printf("r: %c ASCII: %d - abs: %d - x: %d - y: %d\n", char, char, pos, posx, posy)
	return &sdl.Rect{pos*7 - int32(pos/18)*126, int32(pos/18) * 9, 7, 9}
}
