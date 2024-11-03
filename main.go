package main

import (
	"fmt"
	"os"
	"time"

	"github.com/Yunraze/go-rogue/tilemap"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

const winWidth, winHeight int32 = 800, 600

const (
	windowWidth      int32   = 800
	windowHeight     int32   = 600
	spritesPath      string  = "./sprites.png"
	targetFps        float64 = 60
	targetMsPerFrame         = 1000.0 / targetFps
)

func calculateGrid(windowWidth, windowHeight, spriteWidth, spriteHeight int) (int, int, int, int) {
	horizontalCount := windowWidth / spriteWidth
	verticalCount := windowHeight / spriteHeight

	horizontalPadding := (windowWidth - (horizontalCount * spriteWidth)) / 2
	verticalPadding := (windowHeight - (verticalCount * spriteHeight)) / 2

	return horizontalCount, verticalCount, horizontalPadding, verticalPadding
}

func run() (err error) {
	var window *sdl.Window
	var windowSurface *sdl.Surface
	var spriteSurface *sdl.Surface
	var windowWidth int = 1024
	var windowHeight int = 768

	if err = sdl.Init(sdl.INIT_VIDEO); err != nil {
		return
	}
	defer sdl.Quit()

	// Create a window for drawing the sprites on.
	if window, err = sdl.CreateWindow(
		"go-rogue",
		sdl.WINDOWPOS_UNDEFINED,
		sdl.WINDOWPOS_UNDEFINED,
		int32(windowWidth),
		int32(windowHeight),
		sdl.WINDOW_SHOWN,
	); err != nil {
		return
	}
	defer window.Destroy()

	if windowSurface, err = window.GetSurface(); err != nil {
		return
	}

	// Load a PNG spritesheet.
	if spriteSurface, err = img.Load(spritesPath); err != nil {
		return err
	}
	defer spriteSurface.Free()

	tileMap, err := tilemap.LoadTileMap("./tilemap.json")
	if err != nil {
		fmt.Println("Error loading tile map: ", err)
		os.Exit(1)
	}

	var spriteWidth int = tileMap.TileSize.Width
	var spriteHeight int = tileMap.TileSize.Height

	var horizontalCount, verticalCount, horizontalPadding, verticalPadding int = calculateGrid(
		windowWidth,
		windowHeight,
		spriteWidth,
		spriteHeight,
	)

	fmt.Printf("Number of sprites horizontally: %d\n", horizontalCount)
	fmt.Printf("Number of sprites vertically: %d\n", verticalCount)
	fmt.Printf("Horizontal padding on each side: %d pixels\n", horizontalPadding)
	fmt.Printf("Vertical padding on each side: %d pixels\n", verticalPadding)

	tileMap.InitializeSpriteLookup(spriteWidth, spriteHeight)

	// Set up frame timing.
	running := true
	startTime := time.Now()
	frameCount := 0

	for running {
		frameStart := time.Now()

		for row := 0; row < verticalCount; row++ {
			for column := 0; column < horizontalCount; column++ {
				err = tileMap.DrawSpriteByName(
					"player_male_miner",
					int32(spriteWidth),
					int32(spriteHeight),
					int32(horizontalPadding+(spriteWidth*column)),
					int32(verticalPadding+(spriteHeight*row)),
					windowSurface,
					spriteSurface,
				)

				if err != nil {
					fmt.Println("Error drawing sprite: ", err)
				}
			}
		}

		window.UpdateSurface()

		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				running = false
			}
		}

		frameCount++
		frameDuration := time.Since(frameStart).Milliseconds()

		delay := targetMsPerFrame - float64(frameDuration)
		if delay > 0 {
			sdl.Delay(uint32(delay))
		}

		if time.Since(startTime).Seconds() >= 1 {
			fps := float64(frameCount) / time.Since(startTime).Seconds()
			fmt.Printf("FPS: %.2f with a delay of %f\n", fps, delay)
			startTime = time.Now() // Reset time
			frameCount = 0         // Reset frame count
		}
	}

	return
}

func main() {
	if err := run(); err != nil {
		os.Exit(1)
	}
}
