package hw

import (
	"fmt"
	"math"
)

// По условиям задачи, координаты не могут быть меньше 0.
type Geom struct {
	X1, Y1, X2, Y2 float64
}

func (g *Geom) CalculateDistance() (distance float64, err error) {
	if g.X1 < 0 || g.X2 < 0 || g.Y1 < 0 || g.Y2 < 0 {
		err = fmt.Errorf("Координаты не могут быть меньше нуля")
		return
	}

	distance = math.Sqrt(math.Pow(g.X2-g.X1, 2) + math.Pow(g.Y2-g.Y1, 2))
	return
}
