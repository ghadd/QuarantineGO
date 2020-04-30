package main

import (
	"image"
	"image/color"
	"log"
	"math"
	"os"

	"github.com/fogleman/gg"
)

const (
	windowWidth  = 1280
	windowHeight = 720

	topBorder    = 90
	bottomBorder = 630
	leftBorder   = 80
	rightBorder  = 1200

	ballRadius = 20
)

var (
	context = gg.NewContext(windowWidth, windowHeight)

	board image.Image
	ball  image.Image

	phase1 = color.RGBA{0xff, 0, 0, 0xff}
	phase2 = color.RGBA{0xff, 0xff, 0, 0xff}
	phase3 = color.RGBA{0, 0xff, 0, 0xff}
	phase4 = color.RGBA{0, 0, 0xff, 0xff}

	pools = make([]*point, 6)
)

func load(filename string, dest *image.Image) {
	reader, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()

	*dest, _, err = image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}
}

func loadAssets() {
	load("billiard/assets/board.png", &board)
	load("billiard/assets/ball.png", &ball)

	if board == nil {
		log.Fatal("nilll")
	}
}

func init() {
	pools[0] = newPoint(topBorder, leftBorder)
	pools[1] = newPoint(topBorder, windowWidth/2)
	pools[2] = newPoint(topBorder, rightBorder)
	pools[3] = newPoint(bottomBorder, leftBorder)
	pools[4] = newPoint(bottomBorder, windowWidth/2)
	pools[5] = newPoint(bottomBorder, rightBorder)
}

type point struct {
	x, y float64
}

func newPoint(x, y float64) (p *point) {
	p = new(point)
	p.x = x
	p.y = y
	return
}

func radians(degrees float64) float64 {
	return math.Pi * degrees / 180
}

func mapValue(value, min1, max1, min2, max2 float64) float64 {
	return min2 + (max2-min2)*((value-min1)/(max1-min1))
}

func getColor(value, max float64, transparency uint8) (col color.RGBA) {
	if v := value / max; v < 1/4. {
		col = color.RGBA{phase1.R, phase1.G, phase1.B, transparency}
		col.G += uint8(mapValue(value, 0, max/4, 0, 0xff))
	} else if v < 1/2. {
		col = color.RGBA{phase2.R, phase2.G, phase2.B, transparency}
		col.R -= uint8(mapValue(value, max/4, max/2, 0, 0xff))
	} else if v < 3/4. {
		col = color.RGBA{phase3.R, phase3.G, phase3.B, transparency}
		col.B += uint8(mapValue(value, max/2, 3*max/4, 0, 0xff))
	} else {
		col = color.RGBA{phase4.R, phase4.G, phase4.B, transparency}
		col.G -= uint8(mapValue(value, 3*max/4, max, 0, 0xff))
	}

	return col
}

func dist(a, b point) float64 {
	return math.Sqrt(math.Pow(a.x-b.x, 2) + math.Pow(a.y-b.y, 2))
}

func inPool(position *point) bool {
	offset := 30.
	for i := 0; i < len(pools); i++ {
		if dist(*position, *pools[i]) < offset {
			return true
		}
	}
	return false
}

func draw(position *point, velocity, acceleration, angle float64) {
	var gone float64 = 0
	var distance float64 = math.Pow(velocity, 2) / (2 * acceleration)
	step := 0

	for ; velocity > 0; step++ {
		var bounce bool = position.x > rightBorder-ballRadius ||
			position.x < leftBorder+ballRadius ||
			position.y > bottomBorder-ballRadius ||
			position.y < topBorder+ballRadius

		if bounce {
			if inPool(position) {
				break
			}
			context.DrawImage(ball, int(position.x-20), int(position.y-20))
			if (position.y < topBorder+ballRadius && angle > 0 && angle < radians(180)) ||
				(position.y > bottomBorder-ballRadius && angle > radians(180) && angle < radians(360)) {
				angle = radians(360) - angle
			} else if (position.x > rightBorder-ballRadius && angle >= 0 && angle <= radians(90)) ||
				(position.x < leftBorder+ballRadius && angle > radians(90) && angle < radians(180)) {
				angle = radians(180) - angle
			} else if (position.x > rightBorder-ballRadius && angle >= radians(270) && angle <= radians(360)) ||
				(position.x < leftBorder+ballRadius && angle >= radians(180) && angle <= radians(270)) {
				angle = radians(540) - angle
			}
		}
		temp := newPoint(position.x, position.y)

		position.x += math.Cos(angle) * velocity
		position.y -= math.Sin(angle) * velocity

		context.SetLineWidth(5)
		context.DrawLine(temp.x, temp.y, position.x, position.y)

		context.SetColor(getColor(gone, distance, 0xff))
		context.Stroke()

		if step%50 == 0 {
			context.DrawCircle(position.x, position.y, 20)
			context.Fill()
		}
		context.DrawImage(ball, windowWidth/3-20, windowHeight/2-20)

		gone += velocity
		velocity -= acceleration
	}
}

func main() {
	loadAssets()

	////

	context.DrawImage(board, 0, 0)

	pos := newPoint(windowWidth/3, windowHeight/2)
	velocity := 10.
	acceleration := 0.005
	angle := radians(10)

	draw(pos, velocity, acceleration, angle)

	err := context.SavePNG("billiard/data/board.png")
	if err != nil {
		log.Fatal(err)
	}

	////
}
