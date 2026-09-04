package source

import rl "github.com/gen2brain/raylib-go/raylib"

type scene int

const (
	sceneMenu scene = iota
	scenePlay
)

type Game struct {
	scene    scene
	quit     bool
	player   *Player
	camera   *rl.Camera2D
	levalMap *[]string

	playBtn button
	quitBtn button
}

func NewGame() *Game {
	player := NewPlayer()
	camera := NewCamera(player)
	levalMap := GetMap()

	g := &Game{
		player:   player,
		camera:   camera,
		levalMap: levalMap,
	}

	cx := float32(ScreenWidth) / 2
	cy := float32(ScreenHeight) / 2
	g.playBtn = newButton(cx, cy, 240, 64, "Play")
	g.quitBtn = newButton(cx, cy+90, 240, 64, "Quit")

	return g
}

func (g *Game) Run() {
	rl.InitWindow(ScreenWidth, ScreenHeight, "Millitants")
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)

	defer g.close()

	for !rl.WindowShouldClose() {
		g.update(rl.GetFrameTime())
		rl.BeginMode2D(*g.camera)
		defer rl.EndMode2D()
		g.draw()
	}
}

func (g *Game) update(dt float32) {
	switch g.scene {
	case sceneMenu:
		g.updateMenu()
	case scenePlay:
		g.updateGame(dt)
	}
}

func (g *Game) updateMenu() {
	if g.playBtn.released() {
		g.startGame()
	}
	if g.quitBtn.released() {
		g.quit = true
	}
}

func (g *Game) updateGame(dt float32) {
	if rl.IsKeyPressed(rl.KeyEscape) {
		g.scene = sceneMenu
		return
	}

	g.player.UpdatePlayer(g.camera, *g.levalMap, dt)
	target := GetCameraTarget(*g.player, *g.camera)

	g.camera.Target = rl.Vector2Lerp(
		g.camera.Target,
		target,
		8*dt,
	)
}

func (g *Game) draw() {
	rl.BeginDrawing()
	rl.ClearBackground(rl.DarkGray)
	// rl.DrawTexture(g.background, 0, 0, rl.White)

	switch g.scene {
	case sceneMenu:
		g.drawMenu()
	case scenePlay:
		g.drawGame()
	}

	rl.EndDrawing()
}

func (g *Game) drawMenu() {
	rl.DrawRectangle(0, 0, ScreenWidth, ScreenHeight, rl.Fade(rl.Black, 0.45))

	title := "Millitants"
	titleSize := int32(64)
	tw := rl.MeasureText(title, titleSize)
	rl.DrawText(title, (ScreenWidth-tw)/2+2, 222, titleSize, rl.Black)
	rl.DrawText(title, (ScreenWidth-tw)/2, 220, titleSize, rl.RayWhite)

	g.playBtn.draw()
	g.quitBtn.draw()
}

func (g *Game) startGame() {
	g.scene = scenePlay
}

func (g *Game) drawGame() {
	DrawMap()
	g.player.DrawPlayer()
}

func (g *Game) close() {
}
