Что выведет программа? Объяснить вывод программы. Объяснить внутреннее устройство интерфейсов и их отличие от пустых интерфейсов.

```go
package main

import (
	"fmt"
	"os"
)

func Foo() error {
	var err *os.PathError = nil
	return err
}

func main() {
	err := Foo()
	fmt.Println(err)
	fmt.Println(err == nil)
}
```

Ответ:
```
Выведется "nil false"

Функция Foo возвращает nil, который кастится к интерфейсу error, т.о.
функция возвращает интерфейс error. А интерфейс это структура с 2 полями:
указатель на метаданные типа и на значение
В данном случае типом будет *os.PathError, а значение nil
Далее в функции main мы сравниваем этот интерфейс с обычным nil'ом
А они не равны. Это как бы 2 разных nil'а
Интерфейс будет равен nil'у только если И его динамический тип И значени равны nil
Вообще при сравнении с интерфейсом нужно всегда учитывать underlying тип, а не только значение
Вместо *os.PathError мог быть любой тип, который может принимать значение nil,
например, непроинициализированный map
Разобраться в ситуации поможет следующий код:
fmt.Printf("%T %v", err, err)
fmt.Printf("%T %v", nil, nil)

Вывод будет таким:
*fs.PathError nil
nil nil

Мы вывели тип и хранимое значение и видим, что они разные

Интерфейс в исходниках представлен структурой iface из 2-х полей:
tab *itab и data unsafe.Pointer
структура itab хранит метаданные об интерфейсе и о типе (в поле *_type)
data это указатель на значение

Пустой интерфейс как особый случай представлен упрощенной структурой eface из 2-х полей:
_type *_type и data unsafe.Pointer
Теперь вместо *itab идет сразу информация о типе (дескриптор типа)
Пустой интерфейс это просто обертка, говорящая "по такому адресу хранится значение такого типа"
```