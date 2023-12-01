package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/sinisterstuf/project-scale/camera"
)

const WaterSpeed = 0.35

type Water struct {
	Level      float64
	StartLevel float64
	Image      *ebiten.Image
	Paused     bool
}

func NewWater(startLevel float64) *Water {
	return &Water{
		Level:      startLevel,
		StartLevel: startLevel,
		Image:      loadImage("assets/backdrop/Project-scale-parallax-backdrop_0000_Water-1.png"),
	}
}

func (w *Water) Update(increaseWaterLevel bool) {
	if inpututil.IsKeyJustPressed(ebiten.KeyM) {
		w.Paused = !w.Paused
	}

	if !CheatsAllowed || !w.Paused {
		increase := 1.0
		if !increaseWaterLevel {
			increase = -8.0
		}
		w.Level -= increase * WaterSpeed

		if w.Level > w.StartLevel {
			w.Level = w.StartLevel
		}
	}
}

func (w *Water) Draw(cam *camera.Camera) {
	backdropPos := cam.GetTranslation(
		&ebiten.DrawImageOptions{},
		-float64(w.Image.Bounds().Dx())/2,
		w.Level,
	)
	cam.Surface.DrawImage(w.Image, backdropPos)
}
