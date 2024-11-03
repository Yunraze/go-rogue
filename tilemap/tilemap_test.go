package tilemap

import (
	"fmt"
	"testing"
)

func TestLoadingTilemap(t *testing.T) {
	tileMap, err := LoadTileMap("../tilemap.json")
	if err != nil {
		fmt.Println("Error loading tile map:", err)
		return
	}

	// Print loaded tile map data for demonstration
	fmt.Printf("Tile Map: %dx%d\n", tileMap.Width, tileMap.Height)
	fmt.Printf("Tile Size: %dx%d\n", tileMap.TileSize.Width, tileMap.TileSize.Height)
	fmt.Printf("Gap Size: %d\n", tileMap.GapSize)
	fmt.Println("Sprites:")

	for _, sprite := range tileMap.Sprites {
		fmt.Printf(" - %s: (Row: %d, Column: %d)\n", sprite.Name, sprite.Position.Row, sprite.Position.Column)
	}
}
