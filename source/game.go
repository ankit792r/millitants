package source

import rl "github.com/gen2brain/raylib-go/raylib"

type Scene int

const (
	SceneMenu Scene = iota
	ScenePlay
)

type Game struct {
	scene Scene

	player *Player
	camera *rl.Camera2D
	level  *[]string

	menu *Menu
}

func NewGame() *Game {
	player := NewPlayer()
	camera := NewCamera(player)
	menu := NewMenu()

	return &Game{
		scene:  SceneMenu,
		player: player,
		camera: camera,
		level:  GetMap(),
		menu:   menu,
	}
}

func (g *Game) Run() {
	rl.InitWindow(ScreenWidth, ScreenHeight, "Millitants")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		dt := rl.GetFrameTime()

		g.update(dt)
		g.draw()
	}
}

func (g *Game) update(dt float32) {
	switch g.scene {
	case SceneMenu:
		g.menu.UpdateMenu(g)

	case ScenePlay:
		g.updateGame(dt)
	}
}

func (g *Game) updateGame(dt float32) {
	if rl.IsKeyPressed(rl.KeyEscape) {
		g.scene = SceneMenu
		return
	}

	g.player.UpdatePlayer(g.camera, g.level, dt)

	target := GetCameraTarget(*g.player, *g.camera)

	g.camera.Target = rl.Vector2Lerp(
		g.camera.Target,
		target,
		8*dt,
	)
}

func (g *Game) draw() {
	rl.BeginDrawing()
	defer rl.EndDrawing()

	rl.ClearBackground(rl.DarkGray)

	switch g.scene {
	case SceneMenu:
		g.menu.DrawMenu()

	case ScenePlay:
		g.drawGame()
	}
}

func (g *Game) drawGame() {
	rl.BeginMode2D(*g.camera)

	DrawMap()
	g.player.DrawPlayer()

	rl.EndMode2D()
}
