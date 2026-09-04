package source

import rl "github.com/gen2brain/raylib-go/raylib"

func NewCamera(player *Player) *rl.Camera2D {
	return &rl.Camera2D{
		Target:   player.Position,
		Offset:   rl.Vector2{X: ScreenWidth / 2, Y: ScreenHeight / 2},
		Rotation: 0,
		Zoom:     2,
	}

}

func GetCameraTarget(player Player, camera rl.Camera2D) rl.Vector2 {
	playerCenter := rl.Vector2{
		X: player.Position.X + player.Size.X/2,
		Y: player.Position.Y + player.Size.Y/2,
	}

	mapWidth := getMapWidth()
	mapHeight := getMapHeight()

	// How much of the world is visible on screen.
	halfScreenWidth := float32(ScreenWidth) / (2 * camera.Zoom)
	halfScreenHeight := float32(ScreenHeight) / (2 * camera.Zoom)

	// Camera cannot go beyond the left/right boundaries.
	minX := halfScreenWidth
	maxX := mapWidth - halfScreenWidth

	// Camera cannot go beyond the top/bottom boundaries.
	minY := halfScreenHeight
	maxY := mapHeight - halfScreenHeight

	// Clamp camera to map.
	cameraX := rl.Clamp(playerCenter.X, minX, maxX)
	cameraY := rl.Clamp(playerCenter.Y, minY, maxY)

	return rl.Vector2{
		X: cameraX,
		Y: cameraY,
	}
}
