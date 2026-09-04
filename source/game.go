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

	playBtn button
	quitBtn button
}

func NewGame() *Game {
	player := NewPlayer()
	camera := NewCamera(player)

	cx := float32(ScreenWidth) / 2
	cy := float32(ScreenHeight) / 2

	return &Game{
		scene:  SceneMenu,
		player: player,
		camera: camera,
		level:  GetMap(),

		playBtn: newButton(cx, cy, 240, 64, "Play"),
		quitBtn: newButton(cx, cy+90, 240, 64, "Quit"),
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
		g.updateMenu()

	case ScenePlay:
		g.updateGame(dt)
	}
}

func (g *Game) updateMenu() {
	if g.playBtn.released() {
		g.scene = ScenePlay
	}

	if g.quitBtn.released() {
		rl.CloseWindow()
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
		g.drawMenu()

	case ScenePlay:
		g.drawGame()
	}
}

func (g *Game) drawMenu() {
	title := "Millitants"
	titleSize := int32(64)

	tw := rl.MeasureText(title, titleSize)
	x := (ScreenWidth - tw) / 2

	rl.DrawText(title, x+2, 222, titleSize, rl.Black)
	rl.DrawText(title, x, 220, titleSize, rl.RayWhite)

	g.playBtn.draw()
	g.quitBtn.draw()
}

func (g *Game) drawGame() {
	rl.BeginMode2D(*g.camera)

	DrawMap()
	g.player.DrawPlayer()

	rl.EndMode2D()
}
