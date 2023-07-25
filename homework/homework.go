package hw

import (
	"math"
)

type Geometric struct {
	X1, Y1, X2, Y2 float64
}

func (g *Geometric) CalculateDistance() float64 {
	return math.Sqrt(math.Pow(g.X2-g.X1, 2) + math.Pow(g.Y2-g.Y1, 2))
}

//1)Изменен ресивер метода CalculateDistance с работы с копией до работы с указателем
//2)Удалены лишние переменные код сокращен до одной строчки
