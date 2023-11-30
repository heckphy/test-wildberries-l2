package visitor

// общий интерфейс для всей иерархии объектов (ну в Go на самом деле нет иерархий)
type Shape interface {
	getType() string
	accept(Visitor)
}
