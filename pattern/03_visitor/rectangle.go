package visitor

type Rectangle struct {
	l int
	b int
}

func (t *Rectangle) accept(v Visitor) {
	// вызываем метод у посетителя, который работает с прямоугольником
	v.visitForRectangle(t)
}

// какой-то метод, относящийся к логике самих объектов
func (t *Rectangle) getType() string {
	return "rectangle"
}
