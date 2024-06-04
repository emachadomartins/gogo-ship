package game

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hakuunabatata/gogo-ship/assets"
)

type Game struct {
	player           *Player
	lasers           []*Laser
	meteors          []*Meteor
	meteorSpawnTimer *Timer
	score            int
	lost             bool
	record           int
}

func NewGame() *Game {
	g := &Game{
		meteorSpawnTimer: NewTimer(24),
	}
	player := NewPlayer(g)
	g.player = player

	return g
}

func (g *Game) Update() error {
	if !g.lost {
		g.player.Update()

		for _, l := range g.lasers {
			l.Update()
		}

		g.meteorSpawnTimer.Update()
		if g.meteorSpawnTimer.IsReady() {
			g.meteorSpawnTimer.Reset()

			m := NewMeteor()
			g.meteors = append(g.meteors, m)
		}

		for _, m := range g.meteors {
			m.Update()
		}

		for _, m := range g.meteors {
			if m.Collider().Intersects(g.player.Collider()) {
				g.lost = true
			}
		}

		for i, m := range g.meteors {
			for j, l := range g.lasers {
				if m.Collider().Intersects(l.Collider()) {
					g.meteors = append(g.meteors[:i], g.meteors[i+1:]...)
					g.lasers = append(g.lasers[:j], g.lasers[j+1:]...)
					g.score += 1

					if g.score > g.record {
						g.record = g.score
					}
				}
			}
		}
	} else {
		if ebiten.IsKeyPressed(ebiten.KeySpace) {
			g.Reset()
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if !g.lost {

		g.player.Draw(screen)
		text.Draw(screen, fmt.Sprintf("Record: %d", g.record), assets.FontUi, 20, 50, color.White)
		text.Draw(screen, fmt.Sprintf("Score: %d", g.score), assets.FontUi, 20, 100, color.White)
		for _, l := range g.lasers {
			l.Draw(screen)
		}

		for _, m := range g.meteors {
			m.Draw(screen)
		}

	} else {
		text.Draw(screen, "You Lost!\nPress 'space'\nto try again", assets.FontUi, screenWidth/2-150, 200, color.White)
	}

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func (g *Game) AddLasers(laser *Laser) {
	g.lasers = append(g.lasers, laser)
}

func (g *Game) Reset() {
	g.player = NewPlayer(g)
	g.meteors = nil
	g.lasers = nil
	g.meteorSpawnTimer.Reset()
	g.score = 0
	g.lost = false
}
