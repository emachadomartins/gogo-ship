package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hakuunabatata/gogo-ship/assets"
)

type Player struct {
	image *ebiten.Image
}

func NewPlayer() *Player {
	image := assets.PlayerSprite
	return &Player{
		image: image,
	}
}

func (p *Player) Update() {

}

func (p *Player) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	op.GeoM.Translate(100, 100)
	screen.DrawImage(p.image, op)
}
