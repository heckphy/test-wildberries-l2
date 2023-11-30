package state

import (
	"fmt"
	"log"
)

/*
Состояние — это поведенческий паттерн проектирования,
который позволяет объектам менять поведение в зависимости от своего состояния.
Извне создаётся впечатление, что изменился класс объекта.

Применимость:
1. Когда есть объект, поведение которого кардинально меняться в зависимости от внутреннего состояния,
причём типов состояний много, и их код часто меняется.
Паттерн предлагает выделить в собственные классы все поля и методы,
связанные с определёнными состояниями. Первоначальный объект будет постоянно
ссылаться на один из объектов-состояний, делегируя ему часть своей работы.
Для изменения состояния в контекст достаточно будет подставить другой объект-состояние.
2. Когда код класса содержит множество больших, похожих друг на друга, условных операторов,
которые выбирают поведения в зависимости от текущих значений полей класса.
Паттерн предлагает переместить каждую ветку такого условного оператора в собственный класс.
Тут же можно поселить и все поля, связанные с данным состоянием.
3. Когда вы сознательно используете табличную машину состояний, построенную на условных операторах,
но вынуждены мириться с дублированием кода для похожих состояний и переходов.
Паттерн Состояние позволяет реализовать иерархическую машину состояний, базирующуюся на наследовании.
Вы можете отнаследовать похожие состояния от одного родительского класса и вынести туда весь дублирующий код.

Плюсы:
1. Избавляет от множества больших условных операторов машины состояний.
2. Концентрирует в одном месте код, связанный с определённым состоянием.
3. Упрощает код контекста.

Минусы:
1. Может неоправданно усложнить код, если состояний мало и они редко меняются.

Пример:
Применим паттерн проектирования Состояние в контексте торговых автоматов.
Для упрощения задачи представим, что торговый автомат может выдавать только один товар.
Также представим, что автомат может пребывать только в одном из четырех состояний:

    hasItem (имеетПредмет)
    noItem (неИмеетПредмет)
    itemRequested (выдаётПредмет)
    hasMoney (получилДеньги)

Торговый автомат может иметь различные действия. Опять-таки, для простоты оставим только четыре из них:

    Выбрать предмет
    Добавить предмет
    Ввести деньги
    Выдать предмет

Паттерн Состояние нужно использовать в случаях, когда объект может иметь много различных состояний,
которые он должен менять в зависимости от конкретного поступившего запроса.
В нашем примере, автомат может быть в одном из множества состояний, которые непрерывно меняются.
Допустим, что торговый автомат находится в режиме itemRequested.
Как только произойдет действие «Ввести деньги», он сразу же перейдет в состояние hasMoney.
В зависимости от состояния торгового автомата, в котором он находится на данный момент,
он может по-разному отвечать на одни и те же запросы. Например, если пользователь хочет купить предмет,
машина выполнит действие, если она находится в режиме hasItemState, и отклонит запрос в режиме noItemState.
Программа торгового автомата не захламлена этой логикой;
весь режимозависимый код обитает в реализациях соответствующих состояний.
*/

func main() {
	vendingMachine := newVendingMachine(1, 10)

	err := vendingMachine.requestItem()
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = vendingMachine.insertMoney(10)
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = vendingMachine.dispenseItem()
	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Println()

	err = vendingMachine.addItem(2)
	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Println()

	err = vendingMachine.requestItem()
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = vendingMachine.insertMoney(10)
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = vendingMachine.dispenseItem()
	if err != nil {
		log.Fatalf(err.Error())
	}
}