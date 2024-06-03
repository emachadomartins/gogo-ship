package game

import "github.com/hajimehoshi/ebiten/v2"

type Game struct {
	player *Player
	lasers []*Laser
}

func NewGame() *Game {
	g := &Game{}
	player := NewPlayer(g)
	g.player = player

	return g
}

func (g *Game) Update() error {
	g.player.Update()

	for _, l := range g.lasers {
		l.Update()
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.player.Draw(screen)

	for _, l := range g.lasers {
		l.Draw(screen)
	}

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func (g *Game) AddLasers(laser *Laser) {
	g.lasers = append(g.lasers, laser)
}
