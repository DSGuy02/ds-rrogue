// Code written by Oladeji Sanyaolu (gamedata.go) 08/11/2023

package main

// GameData holds the value for the size of elements within the game
type GameData struct {
	ScreenWidth  int
	ScreenHeight int
	TileWidth    int
	TileHeight   int
	UIHeight     int
}

// NewGameData creates a fully populated GameData Struct.
func NewGameData() GameData {
	g := GameData{
		ScreenWidth:  80,
		ScreenHeight: 50,
		TileWidth:    16,
		TileHeight:   16,
		UIHeight:     10,
	}

	return g
}
