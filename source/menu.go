package source

import rl "github.com/gen2brain/raylib-go/raylib"

// Button ------------------
type Button struct {
	bounds rl.Rectangle
	label  string
}

func newButton(cx, cy, w, h float32, label string) Button {
	return Button{
		bounds: rl.Rectangle{X: cx - w/2, Y: cy - h/2, Width: w, Height: h},
		label:  label,
	}
}

func (b *Button) hovered() bool {
	return rl.CheckCollisionPointRec(rl.GetMousePosition(), b.bounds)
}

func (b *Button) released() bool {
	return b.hovered() && rl.IsMouseButtonReleased(rl.MouseButtonLeft)
}

func (b *Button) draw() {
	bg := rl.DarkGreen
	fg := rl.White
	if b.hovered() {
		bg = rl.Lime
		fg = rl.DarkGreen
	}

	rl.DrawRectangleRec(b.bounds, bg)
	rl.DrawRectangleLinesEx(b.bounds, 3, rl.RayWhite)

	size := int32(32)
	tw := rl.MeasureText(b.label, size)
	x := int32(b.bounds.X + (b.bounds.Width-float32(tw))/2)
	y := int32(b.bounds.Y + (b.bounds.Height-float32(size))/2)
	rl.DrawText(b.label, x, y, size, fg)
}

// Menu --------------
type Menu struct {
	playBtn Button
	quitBtn Button
}

func NewMenu() *Menu {
	cx := float32(ScreenWidth) / 2
	cy := float32(ScreenHeight) / 2

	return &Menu{
		playBtn: newButton(cx, cy, 240, 64, "Play"),
		quitBtn: newButton(cx, cy+90, 240, 64, "Quit"),
	}
}

func (m *Menu) UpdateMenu(game *Game) {

	if m.playBtn.released() {
		game.scene = ScenePlay
	}

	if m.quitBtn.released() {
		rl.CloseWindow()
	}
}

func (m *Menu) DrawMenu() {
	title := "Millitants"
	titleSize := int32(64)

	tw := rl.MeasureText(title, titleSize)
	x := (ScreenWidth - tw) / 2

	rl.DrawText(title, x+2, 222, titleSize, rl.Black)
	rl.DrawText(title, x, 220, titleSize, rl.RayWhite)

	m.playBtn.draw()
	m.quitBtn.draw()
}
