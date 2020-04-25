package main

import (
	"log"
	"math"
	"strconv"

	"github.com/fogleman/gg"
	"gonum.org/v1/gonum/stat/distuv"
)

const (
	minSpeed     = 10.
	maxSpeed     = 100.
	nGradient    = 10
	nPoints      = 100
	offset       = 10
	windowWidth  = 1000
	windowHeight = 1000
)

var (
	colors  = make([]string, nGradient)
	data    = make([]float64, nPoints)
	context = gg.NewContext(windowWidth, windowHeight)
)

func sum(data []float64) (s float64) {
	for _, val := range data {
		s += float64(val)
	}
	return
}

func avg(data []float64) float64 {
	return sum(data) / float64(len(data))
}

func peakPoint(data []float64, offset int) (peakIndex int) {
	peakValue := 0.0
	for i := 0; i < len(data)-offset; i++ {
		if currAvg := avg(data[i : i+offset]); currAvg > peakValue {
			peakValue = currAvg
			peakIndex = i
		}
	}
	return
}

func byteToHexString(val byte) string {
	if val/10 == 0 {
		return "0" + strconv.FormatInt(int64(val), 0x10)
	}
	return strconv.FormatInt(int64(val), 0x10)
}

func initColors() {
	var r, g, b byte
	r, g, b = 0, 0xFF, 0
	step := byte(2 * 0xFF / nGradient)

	for i := 0; i < nGradient; i++ {
		colors[i] = byteToHexString(r) +
			byteToHexString(g) +
			byteToHexString(b)
		if r+step < 0xFF {
			r += step
		} else {
			g -= step
		}
	}
}

func getColor(val float64) string {
	partitions := (maxSpeed - minSpeed) / nGradient
	index := int(math.Floor(val-minSpeed) / partitions)
	return colors[index]
}

func fillData(deviation float64) {
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

func drawComputations(deviation, windowOffset float64) {
	peak := float64(peakPoint(data, offset))
	context.DrawRectangle(peak*windowOffset, 0, 100, 1000)

	context.SetHexColor("000000")
	context.SetLineWidth(10)
	context.Stroke()

	fontFace, err := gg.LoadFontFace("paceTracker/assets/consola.ttf", 25)
	if err != nil {
		log.Fatal(err)
	}
	context.SetFontFace(fontFace)
	context.DrawString("deviation: "+strconv.FormatFloat(deviation, 'f', 0, 64), 450., 25.)
}

func main() {
	initColors()

	for deviation := 1.; deviation <= (maxSpeed-minSpeed)/2; deviation++ {
		fillData(deviation)

		windowOffset := float64(windowWidth / nPoints)

		for i := range data {
			x, y := float64(i)*windowOffset, 0.
			context.DrawRectangle(x, y, windowOffset, windowHeight)
			context.SetHexColor(getColor(data[i]))
			context.Fill()
		}

		drawComputations(deviation, windowOffset)

		err := context.SavePNG("paceTracker/variations/" + strconv.FormatFloat(deviation, 'f', 0, 64) + ".png")
		if err != nil {
			log.Fatal(err)
		}
	}
}
