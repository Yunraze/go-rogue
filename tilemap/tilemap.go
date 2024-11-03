package tilemap

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/veandco/go-sdl2/sdl"
)

type PixelPosition struct {
	X, Y int
}

type TileMapPosition struct {
	Row    int `json:"row"`
	Column int `json:"column"`
}

type TileMapSprite struct {
	Name     string          `json:"name"`
	Position TileMapPosition `json:"position"`
}

type TileMap struct {
	FileName string `json:"filename"`
	Width    int    `json:"width"`
	Height   int    `json:"height"`
	TileSize struct {
		Width  int `json:"width"`
		Height int `json:"height"`
	} `json:"tile_size"`
	GapSize      int             `json:"gap_size"`
	Sprites      []TileMapSprite `json:"sprites"`
	SpriteLookup map[string]PixelPosition
}

func LoadTileMap(filename string) (*TileMap, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("Could not open file: %w", err)
	}
	defer file.Close()

	var tileMap TileMap

	if err := json.NewDecoder(file).Decode(&tileMap); err != nil {
		return nil, fmt.Errorf("Could not decode JSON: %w", err)
	}

	return &tileMap, nil
}

func (tileMap *TileMap) InitializeSpriteLookup(spriteWidth, spriteHeight int) {
	tileMap.SpriteLookup = make(map[string]PixelPosition)

	for _, sprite := range tileMap.Sprites {
		x := sprite.Position.Column*(spriteWidth+tileMap.GapSize) + tileMap.GapSize
		y := sprite.Position.Row*(spriteHeight+tileMap.GapSize) + tileMap.GapSize
		tileMap.SpriteLookup[sprite.Name] = PixelPosition{X: x, Y: y}
	}
}

func (tileMap *TileMap) DrawSpriteByName(
	spriteName string,
	spriteWidth, spriteHeight, x, y int32,
	windowSurface, spriteSurface *sdl.Surface,
) error {
	position, exists := tileMap.SpriteLookup[spriteName]
	if !exists {
		return fmt.Errorf("Sprite %s was not found in the tileMap", spriteName)
	}

	srcRect := &sdl.Rect{
		X: int32(position.X),
		Y: int32(position.Y),
		W: spriteWidth,
		H: spriteHeight,
	}

	destRect := &sdl.Rect{
		X: x,
		Y: y,
		W: spriteWidth,
		H: spriteHeight,
	}

	err := spriteSurface.Blit(srcRect, windowSurface, destRect)
	if err != nil {
		return err
	}

	return nil
}
