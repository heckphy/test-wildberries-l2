package visitor

// интерфейс посетителя
// он будет являться принимаемым типом у методов accept элементов
type Visitor interface {
	visitForSquare(*Square)
	visitForCircle(*Circle)
	visitForRectangle(*Rectangle)
}
