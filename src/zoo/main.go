package main

import (
	_ "image/png"
	"zoo/animals"
	"zoo/drawers"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

var (
	populationSize = 100
	population     = make([]animals.Animal, populationSize)
)

// animals enum
const (
	eParrot  = iota
	eEagle   = iota
	eWhale   = iota
	eShark   = iota
	eGiraffe = iota
	eLion    = iota

	// goes for total number of species available
	species = iota
)

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Pixel Rocks!",
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

	for !win.Closed() {
		win.Clear(colornames.Skyblue)
		imd.Draw(win)

		i := 0.
		for _, v := range sprites {
			pos := pixel.IM.Moved(pixel.V(i*60, i*60)).Scaled(pixel.V(i*60, i*60), 0.01)
			//fmt.Println(pos)
			v.Draw(win, pos)
			i++
		}
		win.Update()
	}
}

func main() {
	pixelgl.Run(run)

	// for i := 0; i < populationSize; i++ {
	// 	sp := rand.Intn(species)
	// 	switch sp {
	// 	case eParrot:
	// 		population[i] = new(bhb.Parrot)
	// 	case eEagle:
	// 		population[i] = new(bpd.Eagle)
	// 	case eWhale:
	// 		population[i] = new(fhb.Whale)
	// 	case eShark:
	// 		population[i] = new(fpd.Shark)
	// 	case eGiraffe:
	// 		population[i] = new(mhb.Giraffe)
	// 	case eLion:
	// 		population[i] = new(mpd.Lion)
	// 	}
	// }

	// for _, v := range population {
	// 	v.Sound()
	// }
}
