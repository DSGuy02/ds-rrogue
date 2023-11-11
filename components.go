// Code written by Oladeji Sanyaolu (components.go) 09/11/2023

package main

import "github.com/hajimehoshi/ebiten/v2"

type Player struct{}

type Monster struct{}

type Position struct {
	X int
	Y int
}

type Renderable struct {
	Image *ebiten.Image
}

type Movable struct{}
