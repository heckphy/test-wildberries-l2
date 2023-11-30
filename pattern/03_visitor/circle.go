package visitor

type Circle struct {
	radius int
}

func (c *Circle) accept(v Visitor) {
	// вызываем метод у посетителя, который работает с кругом
	v.visitForCircle(c)
}

// какой-то метод, относящийся к логике самих объектов
func (c *Circle) getType() string {
	return "Circle"
}
