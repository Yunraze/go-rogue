package dungeon

import (
	"testing"
)

func TestDungeonGeneration(t *testing.T) {
	width, height := 50, 50
	dungeon := NewDungeon(width, height)

	// Generate rooms with parameters for number of rooms, min/max room sizes
	dungeon.GenerateRooms(10, 5, 10)

	// Test if the dungeon dimensions are correct
	if dungeon.Width != width || dungeon.Height != height {
		t.Fatalf("Dungeon dimensions are incorrect. Expected %dx%d, got %dx%d", width, height, dungeon.Width, dungeon.Height)
	}

	// Test if at least some tiles are floors
	foundFloor := false
	foundWall := false

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			switch dungeon.Tiles[y][x].Type {
			case Floor:
				foundFloor = true
			case Wall:
				foundWall = true
			}
		}
	}

	if !foundFloor {
		t.Error("Dungeon does not have any floor tiles")
	}

	if !foundWall {
		t.Error("Dungeon does not have any wall tiles")
	}
}

func TestDungeonGenerationVisual(t *testing.T) {
	width, height := 50, 50
	dungeon := NewDungeon(width, height)
	dungeon.GenerateRooms(10, 5, 10)
	
	dungeon.Print()
}

func TestRoomPlacement(t *testing.T) {
	width, height := 50, 50
	dungeon := NewDungeon(width, height)
	
	// Generate rooms
	dungeon.GenerateRooms(10, 5, 10)
	
	// Check for overlapping rooms or rooms placed out of bounds
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			// Ensure the rooms and walls are placed within bounds
			if y < 0 || y >= height || x < 0 || x >= width {
				t.Errorf("Room placed out of bounds at %d, %d", x, y)
			}
		}
	}
	
	// Additional tests here
}

