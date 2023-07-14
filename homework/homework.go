package hw

import (
	"fmt"
	"math"
)

// По условиям задачи, координаты не могут быть меньше 0.
type Geometric struct {
	X1, Y1, X2, Y2 float64
}

func (g *Geometric) CalculateDistance() (distance float64, err error) {
	if g.X1 < 0 || g.X2 < 0 || g.Y1 < 0 || g.Y2 < 0 {
		err = fmt.Errorf("Координаты не могут быть меньше нуля")
		return
	}

	distance = math.Sqrt(math.Pow(g.X2-g.X1, 2) + math.Pow(g.Y2-g.Y1, 2))
	return
}
//1)Изменен ресивер метода CalculateDistance с работы с копией до работы с указателем
//2)На строчке 24 оригинального кода return distance заменен на простой return, 
//  так как используется Named return values и указывать distance в return избыточно
//3)Произведена замена печати текста ошибки в консоли на возврат ошибки
