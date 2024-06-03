package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hakuunabatata/gogo-ship/assets"
)

type Laser struct {
	image    *ebiten.Image
	position Vector
}

func NewLaser(position Vector) *Laser {
	image := assets.LaserSprite

	bounds := image.Bounds()

	halfWidth := float64(bounds.Dx())
	halfHeight := float64(bounds.Dy())

	position.X -= halfWidth
	position.Y -= halfHeight

	return &Laser{image, position}
}

func (l *Laser) Update() {
	speed := 7.0

	l.position.Y += -speed

}

func (l *Laser) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	op.GeoM.Translate(l.position.X, l.position.Y)
	screen.DrawImage(l.image, op)
}

func (l *Laser) Collider() Rect {
	bounds := l.image.Bounds()

	return NewRect(
		l.position.X,
		l.position.Y,
		float64(bounds.Dx()),
		float64(bounds.Dy()),
	)
}
