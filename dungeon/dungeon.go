package dungeon

import (
	"fmt"
	"math/rand"
)

type TileType int

const (
	Floor TileType = iota
	Wall
)

type Tile struct {
	Type TileType
}

type Dungeon struct {
	Width, Height int
	Tiles         [][]Tile
}

func NewDungeon(width, height int) *Dungeon {
	dungeon := &Dungeon{
		Width:  width,
		Height: height,
		Tiles:  make([][]Tile, height),
	}

	// Initialize tiles as walls
	for y := 0; y < height; y++ {
		dungeon.Tiles[y] = make([]Tile, width)
		for x := 0; x < width; x++ {
			dungeon.Tiles[y][x] = Tile{Type: Wall}
		}
	}

	return dungeon
}

func (d *Dungeon) CarveRoom(x, y, w, h int) {
	for iy := y; iy < y+h; iy++ {
		for ix := x; ix < x+w; ix++ {
			if iy < 0 || iy >= d.Height || ix < 0 || ix >= d.Width {
				continue
			}

			d.Tiles[iy][ix] = Tile{Type: Floor}
		}
	}
}

func (d *Dungeon) GenerateRooms(maxRooms, roomMinSize, roomMaxSize int) {
	for i := 0; i < maxRooms; i++ {
		w := rand.Intn(roomMaxSize-roomMinSize) + roomMinSize
		h := rand.Intn(roomMaxSize-roomMinSize) + roomMinSize
		x := rand.Intn(d.Width - w - 1)
		y := rand.Intn(d.Height - h - 1)
		d.CarveRoom(x, y, w, h)
	}
}

func (d *Dungeon) Print() {
	for y := 0; y < d.Height; y++ {
		for x := 0; x < d.Width; x++ {
			switch d.Tiles[y][x].Type {
			case Wall:
				fmt.Print("#")
			case Floor:
				fmt.Print(".")
			}
		}
		
		fmt.Println()
	}
}
