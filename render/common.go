package render

import (
	"github.com/veandco/go-sdl2/sdl"
)

const (
	fontWidth   = 10
	fontHeight  = 10
	nbCodeLines = 20
)

type KEYPressed struct {
	KeyCode uint
	Mode    uint
}

var (
	setFPS      uint64 = 60
	throttleFPS uint64 = 1000 / (setFPS + 5)
	fps         uint64
	frameCount  uint64
	lastFrame   uint64
	lastTime    uint64
	timerFPS    uint64
	Xadjust     int
	Yadjust     int
)

func getGlyph(char rune) *sdl.Rect {
	pos := int32(char - 32)
	// posy := int32(pos / 18)
	// posx := pos - int32(pos / 18)*18
	// fmt.Printf("r: %c ASCII: %d - abs: %d - x: %d - y: %d\n", char, char, pos, posx, posy)
	return &sdl.Rect{pos*7 - int32(pos/18)*126, int32(pos/18) * 9, 7, 9}
}
