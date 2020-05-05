package drawers

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
)

func DrawGrid(imd *imdraw.IMDraw) {
	for i := 0.; i <= Rows; i++ {
		imd.Push(pixel.V(0., i*OffsetH))
		imd.Push(pixel.V(WindowWidth, i*OffsetH))
		imd.Line(3)
	}

	for i := 0.; i <= Cols; i++ {
		imd.Push(pixel.V(i*OffsetW, 0))
		imd.Push(pixel.V(i*OffsetH, WindowHeight))
		imd.Line(3)
	}
}
