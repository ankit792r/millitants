package source

import rl "github.com/gen2brain/raylib-go/raylib"

type Player struct {
	Position rl.Vector2
	Size     rl.Vector2
	Speed    float32
}

func NewPlayer() *Player {
	return &Player{
		Position: rl.Vector2{
			X: 20 * TileSize,
			Y: 20 * TileSize,
		},
		Size: rl.Vector2{
			X: PlayerWidth,
			Y: PlayerHeight,
		},
		Speed: PlayerSpeed,
	}
}

func (p *Player) UpdatePlayer(camera *rl.Camera2D, levalMap *[]string, dt float32) {
	var movement rl.Vector2

	if rl.IsKeyDown(rl.KeyW) || rl.IsKeyDown(rl.KeyUp) {
		movement.Y -= 1
	}

	if rl.IsKeyDown(rl.KeyS) || rl.IsKeyDown(rl.KeyDown) {
		movement.Y += 1
	}

	if rl.IsKeyDown(rl.KeyA) || rl.IsKeyDown(rl.KeyLeft) {
		movement.X -= 1
	}

	if rl.IsKeyDown(rl.KeyD) || rl.IsKeyDown(rl.KeyRight) {
		movement.X += 1
	}

	if rl.IsKeyDown(rl.KeyH) {
		camera.Rotation += 2
	}

	if rl.IsKeyDown(rl.KeyL) {
		camera.Rotation -= 2
	}

	// Normalize diagonal movement.
	if movement.X != 0 || movement.Y != 0 {
		movement = rl.Vector2Normalize(movement)
	}

	movement.X *= p.Speed * dt
	movement.Y *= p.Speed * dt

	// Move X and check collision.
	p.Position.X += movement.X

	if checkPlayerCollision(p, levalMap) || checkWorldBoundary(*p) {
		p.Position.X -= movement.X
	}

	// Move Y and check collision.
	p.Position.Y += movement.Y

	if checkPlayerCollision(p, levalMap) || checkWorldBoundary(*p) {
		p.Position.Y -= movement.Y
	}
}

func checkPlayerCollision(player *Player, levelMap *[]string) bool {
	playerRect := rl.Rectangle{
		X:      player.Position.X,
		Y:      player.Position.Y,
		Width:  player.Size.X,
		Height: player.Size.Y,
	}

	for row, line := range *levelMap {
		for col, tile := range line {
			if tile != '#' {
				continue
			}

			wallRect := rl.Rectangle{
				X:      float32(col * TileSize),
				Y:      float32(row * TileSize),
				Width:  TileSize,
				Height: TileSize,
			}

			if rl.CheckCollisionRecs(playerRect, wallRect) {
				return true
			}
		}
	}

	return false
}

func checkWorldBoundary(player Player) bool {
	mapWidth := getMapWidth()
	mapHeight := getMapHeight()

	// Left
	if player.Position.X < 0 {
		return true
	}

	// Top
	if player.Position.Y < 0 {
		return true
	}

	// Right
	if player.Position.X+player.Size.X > mapWidth {
		return true
	}

	// Bottom
	if player.Position.Y+player.Size.Y > mapHeight {
		return true
	}

	return false
}

func (p *Player) DrawPlayer() {
	// Player body
	rl.DrawRectangle(
		int32(p.Position.X),
		int32(p.Position.Y),
		int32(p.Size.X),
		int32(p.Size.Y),
		rl.Red,
	)

	// Player outline
	rl.DrawRectangleLines(
		int32(p.Position.X),
		int32(p.Position.Y),
		int32(p.Size.X),
		int32(p.Size.Y),
		rl.Maroon,
	)
}
