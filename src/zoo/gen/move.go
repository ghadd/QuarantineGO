package gen

import (
	"fmt"
	"math/rand"
	"time"
	"zoo/animals"
	"zoo/drawers"

	"github.com/faiface/pixel"
)

var (
	dir = []pixel.Vec{
		pixel.V(-1, 0),
		pixel.V(0, 1),
		pixel.V(1, 0),
		pixel.V(0, -1)}
)

func vecToRaw(v pixel.Vec) int {
	return int(v.X*drawers.Cols) + int(v.Y)
}

func rawToVec(r int) pixel.Vec {
	return pixel.V(float64(r/drawers.Cols), float64(r%drawers.Cols))
}

func isEmpty(an *[]animals.Animal, pos pixel.Vec) bool {
	return (*an)[vecToRaw(pos)] == nil
}

func shuffle(vals1 []bool, vals2 []pixel.Vec) ([]bool, []pixel.Vec) {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	ret1 := make([]bool, len(vals1))
	ret2 := make([]pixel.Vec, len(vals2))
	n := len(vals1)
	for i := 0; i < n; i++ {
		randIndex := r.Intn(len(vals1))
		ret1[i] = vals1[randIndex]
		ret2[i] = vals2[randIndex]
		vals1 = append(vals1[:randIndex], vals1[randIndex+1:]...)
		vals2 = append(vals2[:randIndex], vals2[randIndex+1:]...)
	}

	return ret1, ret2
}

// Move moves animals
func Move(an *[]animals.Animal) {
	for i := 0; i < len(*an); i++ {
		var free []bool
		var correspondingDir []pixel.Vec

		for _, v := range dir {
			correspondingDir = append(correspondingDir, v)
		}

		current := rawToVec(i)
		for j := range dir {
			dest := current.Add(dir[j])
			if dest.X < 0 || dest.Y < 0 || dest.X >= drawers.Cols || dest.Y >= drawers.Rows {
				free = append(free, false)
				continue
			}
			free = append(free, isEmpty(an, dest))
		}

		fmt.Println(free, correspondingDir)
		free, correspondingDir = shuffle(free, correspondingDir)
		fmt.Println(free, correspondingDir)

		for j, v := range free {
			if v {
				(*an)[vecToRaw(current)], (*an)[vecToRaw(current.Add(correspondingDir[j]))] = (*an)[vecToRaw(current.Add(correspondingDir[j]))], (*an)[vecToRaw(current)]
			}
		}
	}
}
