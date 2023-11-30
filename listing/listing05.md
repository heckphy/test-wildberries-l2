Что выведет программа? Объяснить вывод программы.

```go
package main

type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}

func test() *customError {
	{
		// do something
	}
	return nil
}

func main() {
	var err error
	err = test()
	if err != nil {
		println("error")
		return
	}
	println("ok")
}
```

Ответ:
```
Выведется "error"

test вернет указатель на customError со значением nil
Далее идет сравнение с интерфейсом.
Как было сказано в предыдущей задаче, при сравнении интерфейса с чем-то
должен учитываться динамический тип (underlying) тип интерфейса и его значение
err это переменная интерфейсного типа, динамический тип которой это *customError, а значение nil
И err сравнивается с переменной типа nil со значением nil.
Они не равны, а значит, выведется "error".
Другими словами (*customError, nil) != (nil, nil)
```
