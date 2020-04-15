package main

import (
	"fmt"
	"math"
	"math/rand"

	"github.com/fatih/color"
)

const (
	height    = 25
	windspeed = 50
	winddir   = 'r'
)

func main() {
	magenta := color.New(color.FgMagenta)
	green := color.New(color.FgGreen)
	red := color.New(color.FgRed)
	blue := color.New(color.FgBlue)
	for i := 0; i < height; i++ {
		for j := 0; j < height*3; j++ {
			slope := 0
			switch winddir {
			case 'r':
				slope = int(windspeed / math.Pow(float64(i+2), float64(0.5)))
			case 'l':
				slope = -int((i + 1) * (height - i) * windspeed)
			default:
				slope = 0
			}

			switch {
			case j < height-i+slope || j > height+i+slope || (i > height-2 && (j < height-1 || j > height+1)):
				fmt.Print(" ")
			case j == height && i > height-6:
				magenta.Print("■")
			case (j == height-1 || j == height+1) && i > height-6:
				magenta.Print("|")
			default:
				topperCase := rand.Intn(100) // 10% of random topper generation
				switch {
				case topperCase < 5:
					red.Print("0")
				case topperCase < 10:
					blue.Print("0")
				default:
					green.Print("*")
				}
			}
		}
		fmt.Println()
	}
}
