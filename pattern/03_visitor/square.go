package visitor

type Square struct {
	side int
}

func (s *Square) accept(v Visitor) {
	// вызываем метод у посетителя, который работает с квадратом
	v.visitForSquare(s)
}

// какой-то метод, относящийся к логике самих объектов
func (s *Square) getType() string {
	return "Square"
}
