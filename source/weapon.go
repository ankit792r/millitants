package source

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Weapon struct {
	active bool
	angle  float32
	rect   rl.Rectangle
}

func NewWeapon(playerPos rl.Vector2) *Weapon {
	rect := rl.Rectangle{
		X:      playerPos.X,
		Y:      playerPos.Y,
		Width:  10,
		Height: 10,
	}

	return &Weapon{
		active: true,
		rect:   rect,
	}
}

func (w *Weapon) UpdateWeapon(playerPos rl.Vector2, dt float32) {
	if !w.active {
		return
	}

	mousePos := rl.GetMousePosition()

	w.angle = float32(math.Atan2(
		float64(mousePos.Y-playerPos.Y),
		float64(mousePos.X-playerPos.X),
	))

	w.rect.X = playerPos.X
	w.rect.Y = playerPos.Y
}

func (w *Weapon) DrawWeapon(player *Player) {
	origin := rl.Vector2{
		X: 0,
		Y: player.Size.Y / 2,
	}

	rl.DrawRectanglePro(w.rect, origin, w.angle*rl.Rad2deg, rl.Blue)
}
