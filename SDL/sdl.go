package main

import (
	"image"
	"image/color"
	"log"
	"os"
	"time"
	"unsafe"

	"github.com/golang/freetype"
	"github.com/veandco/go-sdl2/sdl"
)

var Colors [16]color.Color = [16]color.Color{
	color.RGBA{R: 0, G: 0, B: 0, A: 255},       //black
	color.RGBA{R: 72, G: 58, B: 211, A: 255},   //red
	color.RGBA{R: 163, G: 30, B: 9, A: 255},    //dk blue
	color.RGBA{R: 221, G: 84, B: 213, A: 255},  //purple
	color.RGBA{R: 57, G: 133, B: 54, A: 255},   //dk green
	color.RGBA{R: 104, G: 104, B: 104, A: 255}, //gray
	color.RGBA{R: 246, G: 68, B: 51, A: 255},   //med blue
	color.RGBA{R: 249, G: 185, B: 134, A: 255}, //lt blue
	color.RGBA{R: 33, G: 106, B: 147, A: 255},  //brown
	color.RGBA{R: 49, G: 131, B: 240, A: 255},  //orange
	color.RGBA{R: 184, G: 184, B: 184, A: 255}, //grey
	color.RGBA{R: 157, G: 175, B: 244, A: 255}, //pink
	color.RGBA{R: 64, G: 219, B: 97, A: 255},   //lt green
	color.RGBA{R: 82, G: 251, B: 254, A: 255},  //yellow
	color.RGBA{R: 210, G: 247, B: 134, A: 255}, //aqua
	color.RGBA{R: 255, G: 255, B: 255, A: 255}, //white
}

var (
	winHeight int
	winWidth  int
	emuHeight int
	emuWidth  int
	window    *sdl.Window
	w_surf    *sdl.Surface
	emul      *image.RGBA
	emul_s    *sdl.Surface
	emuRect   sdl.Rect
	renderer  *sdl.Renderer
	texture   *sdl.Texture
	mode3D    bool

	font *freetype.Context

	ShowFps  bool
	ShowCode bool
	nbFrames int = 0

	setFPS                                         uint64 = 50
	throttleFPS                                    uint64 = 1000 / (setFPS + 5)
	fps, frameCount, lastFrame, lastTime, timerFPS uint64
)

func DrawPixel(x, y int, c color.Color) {
	emul.Set(x, y, c)
}

func InitSDL2(title string, mode3D bool) {
	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		panic(err)
	}
	sdl.SetHint(sdl.HINT_RENDER_SCALE_QUALITY, "0")

	// Creation de la fenêtre et de son renderer
	window, err = sdl.CreateWindow(title, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, int32(winWidth), int32(winHeight), sdl.WINDOW_SHOWN|sdl.WINDOW_RESIZABLE)
	if mode3D {
		log.Printf("SDL2 mode: 3D (texture)\n")
		renderer, err = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
		if err != nil {
			panic(err)
		}
	} else {
		log.Printf("SDL2 mode: 2D (surface)\n")
	}

	w_surf, err = window.GetSurface()
	w_surf.SetRLE(true)

	emul = image.NewRGBA(image.Rect(0, 0, emuWidth, emuHeight))
	emul_s, _ = sdl.CreateRGBSurfaceFrom(unsafe.Pointer(&emul.Pix[0]), int32(emuWidth), int32(emuHeight), 32, 4*emuWidth, 0, 0, 0, 0)
	emul_s.SetRLE(true)
	emuRect = sdl.Rect{X: 0, Y: 0, W: int32(winWidth), H: int32(winHeight)}
}

func Init(width, height int, zoomFactor int, title string, mode bool) {
	emuHeight = height
	emuWidth = width
	winHeight = emuHeight * zoomFactor
	winWidth = emuWidth * zoomFactor
	mode3D = mode

	ShowFps = false

	log.Printf("Starting renderer using SDL2\n")

	InitSDL2(title, mode)
}

func UpdateFrame() {
	timerFPS = sdl.GetTicks64() - lastFrame
	if timerFPS < throttleFPS {
		log.Printf("timer: %d lastFrame: %d", timerFPS, lastFrame)
		// sdl.Delay(uint32(throttleFPS - timerFPS))
		return
	}
	lastFrame = sdl.GetTicks64()

	if mode3D {
		// SDL2 Texture + Render
		texture, _ = renderer.CreateTextureFromSurface(emul_s)
		renderer.Copy(texture, nil, &emuRect)
		renderer.Present()
	} else {
		// SDL2 Surface
		emul_s.BlitScaled(nil, w_surf, &emuRect)
		window.UpdateSurface()
	}
	nbFrames++
}

func fillScreen() {
	for i := 0; i < 100; i++ {
		for c := 0; c < 16; c++ {
			for y := 0; y < emuHeight; y++ {
				for x := 0; x < emuWidth; x++ {
					DrawPixel(x, y, Colors[c])
				}
			}
		}
	}
	e := &sdl.QuitEvent{Type: sdl.QUIT}
	sdl.PushEvent(e)
}

func main() {
	start := time.Now()
	Init(1024, 768, 1, "Test", true)
	lastFrame = sdl.GetTicks64()
	go fillScreen()
	for {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				elapsed := time.Since(start)
				log.Printf("Rendering took %s - (%d frames)", elapsed, nbFrames)
				os.Exit(1)
			}
		}
		UpdateFrame()
	}
}
