package visitor

import (
	"fmt"
)

// это посетитель
// его методы позволят вычислять площадь у различных фигур
type AreaCalculator struct {
	area int
}

// реализовываем интерфейс посетителя
// каждый метод работает с определенным объект из иерархии

func (a *AreaCalculator) visitForSquare(s *Square) {
	// считаем площадь для квадрата,
	// а затем присваиваем эту площадь инстансу квадрата (в какое-нибудь поле)
	fmt.Println("Calculating area for square")
}

func (a *AreaCalculator) visitForCircle(s *Circle) {
	fmt.Println("Calculating area for circle")
}
func (a *AreaCalculator) visitForRectangle(s *Rectangle) {
	fmt.Println("Calculating area for rectangle")
}
