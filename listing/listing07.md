Что выведет программа? Объяснить вывод программы.

```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func asChan(vs ...int) <-chan int {
	c := make(chan int)

	go func() {
		for _, v := range vs {
			c <- v
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}

		close(c)
	}()
	return c
}

func merge(a, b <-chan int) <-chan int {
	c := make(chan int)
	go func() {
		for {
			select {
			case v := <-a:
				c <- v
			case v := <-b:
				c <- v
			}
		}
	}()
	return c
}

func main() {

	a := asChan(1, 3, 5, 7)
	b := asChan(2, 4 ,6, 8)
	c := merge(a, b )
	for v := range c {
		fmt.Println(v)
	}
}
```

Ответ:
```
Сначала будут выводиться числа из каналов a и b в неопределенном порядке
Затем один из каналов закроется
С этого момента выводиться будут только нули

Когда один из каналов закрывается, бесконечный цикл с select в функции merge продолжает свою работу
И в одном из кейсов (с закрытым каналом) всегда будет происходить чтение данных (нулей),
т.к. чтение из закрытого канала всегда возвращает нулевое значение.
```
