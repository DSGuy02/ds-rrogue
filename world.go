// Code written by Oladeji Sanyaolu (world.go) 10/11/2023

package main

import (
	"log"

	"github.com/bytearena/ecs"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var position *ecs.Component
var renderable *ecs.Component

func InitializeWorld() (*ecs.Manager, map[string]ecs.Tag) {
	tags := make(map[string]ecs.Tag)
	manager := ecs.NewManager()

	playerImg, _, err := ebitenutil.NewImageFromFile("assets/player.png")
	if err != nil {
		log.Fatal(err)
	}

	player := manager.NewComponent()
	position = manager.NewComponent()
	renderable = manager.NewComponent()
	movable := manager.NewComponent()

	manager.NewEntity().
		AddComponent(player, Player{}).
		AddComponent(renderable, Renderable{}).
		AddComponent(movable, Movable{}).
		AddComponent(position, &Position{
			X: 40,
			Y: 25,
		}).
		AddComponent(renderable, &Renderable{
			Image: playerImg,
		})

	players := ecs.BuildTag(player, position)
	tags["players"] = players

	renderables := ecs.BuildTag(renderable, position)
	tags["renderables"] = renderables

	return manager, tags
}
