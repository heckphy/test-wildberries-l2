package visitor

import "fmt"

// посетитель для расчета центральных координат фигур
// каждый метод работает с определенным объект из иерархии

type MiddleCoordinates struct {
	x int
	y int
}

func (a *MiddleCoordinates) visitForSquare(s *Square) {
	// высчитываем центр квадрата,
	// а затем присваиваем это в поля x и y инстанса квадрата
	fmt.Println("Calculating middle point coordinates for square")
}

func (a *MiddleCoordinates) visitForCircle(c *Circle) {
	fmt.Println("Calculating middle point coordinates for circle")
}
func (a *MiddleCoordinates) visitForRectangle(t *Rectangle) {
	fmt.Println("Calculating middle point coordinates for rectangle")
}
