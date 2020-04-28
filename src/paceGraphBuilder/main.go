package main

import (
	"image/color"
	"log"
	"strconv"

	"github.com/fogleman/gg"
	"gonum.org/v1/gonum/stat/distuv"
)

const (
	minSpeed     = 10
	maxSpeed     = 100
	dataSize     = 1000
	windowWidth  = 1000
	windowHeight = 1000
	offset       = 10
)

var (
	data    = make([]float64, dataSize)
	context = gg.NewContext(windowWidth, windowHeight)
)

func generateData(deviation float64) {
	dist := distuv.Normal{
		Mu:    (maxSpeed + minSpeed) / 2,
		Sigma: deviation,
	}

	for i := range data {
		val := dist.Rand()
		for ; val < minSpeed || val > maxSpeed; val = dist.Rand() {
		}
		data[i] = val
	}
}

func sum(data []float64) (s float64) {
	for _, val := range data {
		s += float64(val)
	}
	return
}

func avg(data []float64) float64 {
	return sum(data) / float64(len(data))
}

func mapValue(value, low1, high1, low2, high2 float64) float64 {
	return low2 + (value-low1)*(high2-low2)/(high1-low1)
}

func main() {
	for deviation := 1.; deviation <= (maxSpeed-minSpeed)/2; deviation++ {
		generateData(10)

		context.DrawRectangle(0, 0, windowWidth, windowHeight)
		context.SetColor(color.White)
		context.Fill()

		for i := 0; i < len(data)-offset-1; i++ {
			x1 := mapValue(float64(i), 0, float64(len(data)), 0, windowWidth)
			y1 := windowHeight - mapValue(avg(data[i:i+offset]), minSpeed, maxSpeed, 0, windowHeight)
			x2 := mapValue(float64(i+1), 0, float64(len(data)), 0, windowWidth)
			y2 := windowHeight - mapValue(avg(data[i+1:i+1+offset]), minSpeed, maxSpeed, 0, windowHeight)

			context.DrawLine(x1, y1, x2, y2)
			context.SetColor(color.Black)
			context.SetLineWidth(3)
			context.Stroke()
		}

		fontFace, err := gg.LoadFontFace("paceGraphBuilder/assets/consola.ttf", 25)
		if err != nil {
			log.Fatal(err)
		}
		context.SetFontFace(fontFace)
		context.DrawString("deviation: "+strconv.FormatFloat(deviation, 'f', 0, 64), 450., 25.)

		context.SavePNG("paceGraphBuilder/graphs/" + strconv.FormatFloat(deviation, 'f', 0, 64) + ".png")
	}
}
