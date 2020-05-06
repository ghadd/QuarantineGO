package main

import (
	_ "image/png"
	"time"
	"zoo/animals"
	"zoo/drawers"
	"zoo/gen"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

var (
	populationSize = 100
	population     = make([]animals.Animal, populationSize)

	anims = []string{
		"parrot",
		"eagle",
		"whale",
		"shark",
		"giraffe",
		"lion",
	}
)

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Zoo!",
		Bounds: pixel.R(0, 0, drawers.WindowWidth, drawers.WindowHeight),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	imd := imdraw.New(nil)

	drawers.DrawGrid(imd)
	sprites := drawers.GetSprites()
	gen.GenerateAnimals(&population)

	last := time.Now()
	for !win.Closed() {
		win.Clear(colornames.Skyblue)
		imd.Draw(win)

		dt := time.Since(last).Seconds()

		if dt > 1 {
			gen.Move(&population)
			last = time.Now()
		}

		for i := 0; i < drawers.Rows; i++ {
			for j := 0; j < drawers.Cols; j++ {
				if population[i*drawers.Cols+j] == nil {
					continue
				}

				mat := pixel.IM
				mat = mat.ScaledXY(pixel.ZV, pixel.V(drawers.OffsetW/100, drawers.OffsetH/100))
				mat = mat.Moved(pixel.V(drawers.OffsetW*(float64(i)+0.5), drawers.OffsetH*(float64(j)+0.5)))

				sprites[gen.GetKey(population[i*drawers.Cols+j])].Draw(win, mat)
			}
		}

		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
